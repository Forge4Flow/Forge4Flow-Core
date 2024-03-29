package service

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/forge4flow/forge4flow-core/pkg/config"
	"github.com/forge4flow/forge4flow-core/pkg/stats"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
)

type HealthCheck struct {
	Status string `json:"status"`
}

type RouteHandler[T Service] struct {
	svc     T
	handler func(svc T, w http.ResponseWriter, r *http.Request) error
}

func NewRouteHandler[T Service](svc T, handler func(svc T, w http.ResponseWriter, r *http.Request) error) RouteHandler[T] {
	return RouteHandler[T]{
		svc:     svc,
		handler: handler,
	}
}

func (rh RouteHandler[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := rh.handler(rh.svc, w, r)
	if err != nil {
		// Write err response to client
		SendErrorResponse(w, err)

		// Log and send err to Sentry
		logEvent := hlog.FromRequest(r).Error().Stack().Err(err)
		if apiError, ok := err.(Error); ok {
			// Add additional context to log if ApiError
			logEvent = logEvent.Str("apiError", apiError.GetTag()).
				Int("statusCode", apiError.GetStatus())
		}

		// Log event
		logEvent.Msg("ERROR")
	}
}

func NewRouter(config config.Config, pathPrefix string, svcs []Service, authMiddleware AuthMiddlewareFunc, routerMiddlewares []Middleware, requestMiddlewares []Middleware) (*mux.Router, error) {
	router := mux.NewRouter()

	// Setup default middleware
	logger := zerolog.New(os.Stderr).
		With().
		Timestamp().
		Logger().
		Level(zerolog.Level(config.GetLogLevel()))
	if logger.GetLevel() == zerolog.DebugLevel {
		logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	router.Use(hlog.NewHandler(logger))
	router.Use(requestStatsMiddleware)
	if config.GetEnableAccessLog() {
		router.Use(accessLogMiddleware)
	}
	router.Use(hlog.RequestIDHandler("requestId", "Warrant-Request-Id"))
	router.Use(hlog.URLHandler("uri"))
	router.Use(hlog.MethodHandler("method"))
	router.Use(hlog.ProtoHandler("protocol"))

	// Enable CORS for all origins
	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}),
	)
	router.Use(corsMiddleware)

	// Setup router middlewares, which will be run on ALL
	// requests, even if they are to non-existent endpoints.
	for _, routerMiddleware := range routerMiddlewares {
		router.Use(mux.MiddlewareFunc(routerMiddleware))
	}

	routes := make([]Route, 0)
	for _, svc := range svcs {
		svcRoutes, err := svc.Routes()
		if err != nil {
			log.Fatal().Err(err).Msg("Could not setup routes for service")
		}

		routes = append(routes, svcRoutes...)
	}

	// Setup routes
	for _, route := range routes {
		routePattern := fmt.Sprintf("%s%s", pathPrefix, route.GetPattern())
		middlewareWrappedHandler := ChainMiddleware(route.GetHandler(), requestMiddlewares...)

		var err error
		if route.GetOverrideAuthMiddlewareFunc() != nil {
			middlewareWrappedHandler, err = route.GetOverrideAuthMiddlewareFunc()(config, middlewareWrappedHandler, svcs...)
		} else {
			middlewareWrappedHandler, err = authMiddleware(config, middlewareWrappedHandler, svcs...)
		}
		if err != nil {
			return nil, err
		}

		router.Handle(routePattern, middlewareWrappedHandler).Methods(route.GetMethod())
	}

	router.Handle("/health", http.HandlerFunc(healthCheckHandler)).Methods("GET")

	// Configure catch all handler for 404s
	router.PathPrefix("/").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		SendErrorResponse(w, NewRecordNotFoundError("Endpoint", r.URL.Path))
	}))

	return router, nil
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	status := HealthCheck{Status: "Healthy"}
	SendJSONResponse(w, status)
}

// Create & inject a 'per-request' stats object into request context
func requestStatsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqStats := stats.RequestStats{
			Queries: make([]stats.QueryStat, 0),
		}
		newContext := context.WithValue(r.Context(), stats.RequestStatsKey{}, &reqStats)
		next.ServeHTTP(w, r.WithContext(newContext))
	})
}

func accessLogMiddleware(next http.Handler) http.Handler {
	return hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		logger := hlog.FromRequest(r)
		logEvent := logger.Info().
			Str("method", r.Method).
			Str("protocol", r.Proto).
			Stringer("uri", r.URL).
			Int("statusCode", status).
			Int("size", size).
			Dur("duration", duration).
			Str("clientIp", GetClientIpAddress(r))

		if referer := r.Referer(); referer != "" {
			logEvent = logEvent.Str("referer", referer)
		}

		if userAgent := r.UserAgent(); userAgent != "" {
			logEvent = logEvent.Str("userAgent", userAgent)
		}

		if duration.Milliseconds() >= 500 {
			logEvent = logEvent.Bool("slow", true)
		}

		reqStats, ok := r.Context().Value(stats.RequestStatsKey{}).(*stats.RequestStats)
		if ok {
			logEvent = logEvent.Int("numQueries", reqStats.NumQueries)
			if logger.GetLevel() <= zerolog.DebugLevel {
				logEvent.Object("requestStats", reqStats)
			}
		}

		logEvent.Msg("ACCESS")
	})(next)
}

func GetClientIpAddress(r *http.Request) string {
	clientIpAddress := r.Header.Get("X-Forwarded-For")
	if clientIpAddress == "" {
		return strings.Split(r.RemoteAddr, ":")[0]
	}

	return clientIpAddress
}
