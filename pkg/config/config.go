package config

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/spf13/viper"
)

const (
	DefaultMySQLDatastoreMigrationSource     = "github://auth4flow/auth4flow-core/migrations/datastore/mysql"
	DefaultMySQLEventstoreMigrationSource    = "github://auth4flow/auth4flow-core/migrations/eventstore/mysql"
	DefaultPostgresDatastoreMigrationSource  = "github://auth4flow/auth4flow-core/migrations/datastore/postgres"
	DefaultPostgresEventstoreMigrationSource = "github://auth4flow/auth4flow-core/migrations/eventstore/postgres"
	DefaultSQLiteDatastoreMigrationSource    = "github://auth4flow/auth4flow-core/migrations/datastore/sqlite"
	DefaultSQLiteEventstoreMigrationSource   = "github://auth4flow/auth4flow-core/migrations/eventstore/sqlite"
	DefaultAuthenticationUserIdClaim         = "sub"
	DefaultSessionTokenLength                = 32
	DefaultSessionIdleTimeout                = 15 * time.Minute
	DefaultSessionExpTimeout                 = time.Hour
	PrefixAuth4Flow                          = "auth4flow"
	ConfigFileName                           = "auth4flow.yaml"
)

type Config interface {
	GetPort() int
	GetLogLevel() int8
	GetEnableAccessLog() bool
	GetAutoMigrate() bool
	GetDatastore() *DatastoreConfig
	GetEventstore() *EventstoreConfig
}

type Auth4FlowConfig struct {
	Port            int               `mapstructure:"port"`
	LogLevel        int8              `mapstructure:"logLevel"`
	EnableAccessLog bool              `mapstructure:"enableAccessLog"`
	AutoMigrate     bool              `mapstructure:"autoMigrate"`
	Datastore       *DatastoreConfig  `mapstructure:"datastore"`
	Eventstore      *EventstoreConfig `mapstructure:"eventstore"`
	Authentication  *AuthConfig       `mapstructure:"authentication"`
}

func (warrantConfig Auth4FlowConfig) GetPort() int {
	return warrantConfig.Port
}

func (warrantConfig Auth4FlowConfig) GetLogLevel() int8 {
	return warrantConfig.LogLevel
}

func (warrantConfig Auth4FlowConfig) GetEnableAccessLog() bool {
	return warrantConfig.EnableAccessLog
}

func (warrantConfig Auth4FlowConfig) GetAutoMigrate() bool {
	return warrantConfig.AutoMigrate
}

func (warrantConfig Auth4FlowConfig) GetDatastore() *DatastoreConfig {
	return warrantConfig.Datastore
}

func (warrantConfig Auth4FlowConfig) GetEventstore() *EventstoreConfig {
	return warrantConfig.Eventstore
}

func (warrantConfig Auth4FlowConfig) GetAuthentication() *AuthConfig {
	return warrantConfig.Authentication
}

type DatastoreConfig struct {
	MySQL    *MySQLConfig    `mapstructure:"mysql"`
	Postgres *PostgresConfig `mapstructure:"postgres"`
	SQLite   *SQLiteConfig   `mapstructure:"sqlite"`
}

type MySQLConfig struct {
	Username           string `mapstructure:"username"`
	Password           string `mapstructure:"password"`
	Hostname           string `mapstructure:"hostname"`
	Database           string `mapstructure:"database"`
	MigrationSource    string `mapstructure:"migrationSource"`
	MaxIdleConnections int    `mapstructure:"maxIdleConnections"`
	MaxOpenConnections int    `mapstructure:"maxOpenConnections"`
}

type PostgresConfig struct {
	Username           string `mapstructure:"username"`
	Password           string `mapstructure:"password"`
	Hostname           string `mapstructure:"hostname"`
	Database           string `mapstructure:"database"`
	SSLMode            string `mapstructure:"sslmode"`
	MigrationSource    string `mapstructure:"migrationSource"`
	MaxIdleConnections int    `mapstructure:"maxIdleConnections"`
	MaxOpenConnections int    `mapstructure:"maxOpenConnections"`
}

type SQLiteConfig struct {
	Database           string `mapstructure:"database"`
	InMemory           bool   `mapstructure:"inMemory"`
	MigrationSource    string `mapstructure:"migrationSource"`
	MaxIdleConnections int    `mapstructure:"maxIdleConnections"`
	MaxOpenConnections int    `mapstructure:"maxOpenConnections"`
}

type EventstoreConfig struct {
	MySQL             *MySQLConfig    `mapstructure:"mysql"`
	Postgres          *PostgresConfig `mapstructure:"postgres"`
	SQLite            *SQLiteConfig   `mapstructure:"sqlite"`
	SynchronizeEvents bool            `mapstructure:"synchronizeEvents"`
}

type AuthConfig struct {
	ApiKey             string              `mapstructure:"apiKey"`
	SessionTokenLength int64               `mapstructure:"sessionTokenLength"`
	SessionIdleTimeout int64               `mapstructure:"sessionIdleTimeout"`
	SessionExpTimeout  int64               `mapstructure:"sessionExpTimeout"`
	Provider           *AuthProviderConfig `mapstructure:"providers"`
}

type AuthProviderConfig struct {
	Name          string `mapstructure:"name"`
	PublicKey     string `mapstructure:"publicKey"`
	UserIdClaim   string `mapstructure:"userIdClaim"`
	TenantIdClaim string `mapstructure:"tenantIdClaim"`
}

func NewConfig() Auth4FlowConfig {
	viper.SetConfigFile(ConfigFileName)
	viper.SetDefault("port", 8000)
	viper.SetDefault("logLevel", zerolog.DebugLevel)
	viper.SetDefault("enableAccessLog", true)
	viper.SetDefault("autoMigrate", false)
	viper.SetDefault("datastore.mysql.migrationSource", DefaultMySQLDatastoreMigrationSource)
	viper.SetDefault("datastore.postgres.migrationSource", DefaultPostgresDatastoreMigrationSource)
	viper.SetDefault("datastore.sqlite.migrationSource", DefaultSQLiteDatastoreMigrationSource)
	viper.SetDefault("eventstore.mysql.migrationSource", DefaultMySQLEventstoreMigrationSource)
	viper.SetDefault("eventstore.postgres.migrationSource", DefaultPostgresEventstoreMigrationSource)
	viper.SetDefault("eventstore.sqlite.migrationSource", DefaultSQLiteEventstoreMigrationSource)
	viper.SetDefault("eventstore.synchronizeEvents", false)
	viper.SetDefault("authentication.provider.userIdClaim", DefaultAuthenticationUserIdClaim)
	viper.SetDefault("authentication.sessionTokenLength", DefaultSessionTokenLength)
	viper.SetDefault("authentication.sessionIdleTimeout", int64(DefaultSessionIdleTimeout.Seconds()))
	viper.SetDefault("authentication.sessionExpTimeout", int64(DefaultSessionExpTimeout.Seconds()))

	// If config file exists, use it
	_, err := os.ReadFile(ConfigFileName)
	if err == nil {
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal().Err(err).Msg("Error while reading warrant.yaml. Shutting down.")
		}
	} else {
		if os.IsNotExist(err) {
			log.Info().Msg("Could not find warrant.yaml. Attempting to use environment variables.")
		} else {
			log.Fatal().Err(err).Msg("Error while reading warrant.yaml. Shutting down.")
		}
	}

	var config Auth4FlowConfig
	// If available, use env vars for config
	for _, fieldName := range getFlattenedStructFields(reflect.TypeOf(config)) {
		envKey := strings.ToUpper(fmt.Sprintf("%s_%s", PrefixAuth4Flow, strings.ReplaceAll(fieldName, ".", "_")))
		envVar := os.Getenv(envKey)
		if envVar != "" {
			viper.Set(fieldName, envVar)
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal().Err(err).Msg("Error while creating config. Shutting down.")
	}

	// Configure logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.DurationFieldUnit = time.Millisecond
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.Level(config.GetLogLevel()))
	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	if config.GetAuthentication() == nil || config.GetAuthentication().ApiKey == "" {
		log.Fatal().Msg("Must provide an API key to authenticate incoming requests to Warrant.")
	}

	return config
}

func getFlattenedStructFields(t reflect.Type) []string {
	return getFlattenedStructFieldsHelper(t, []string{})
}

func getFlattenedStructFieldsHelper(t reflect.Type, prefixes []string) []string {
	unwrappedT := t
	if t.Kind() == reflect.Pointer {
		unwrappedT = t.Elem()
	}

	flattenedFields := make([]string, 0)
	for i := 0; i < unwrappedT.NumField(); i++ {
		field := unwrappedT.Field(i)
		fieldName := field.Tag.Get("mapstructure")
		switch field.Type.Kind() {
		case reflect.Struct, reflect.Pointer:
			flattenedFields = append(flattenedFields, getFlattenedStructFieldsHelper(field.Type, append(prefixes, fieldName))...)
		default:
			flattenedField := fieldName
			if len(prefixes) > 0 {
				flattenedField = fmt.Sprintf("%s.%s", strings.Join(prefixes, "."), fieldName)
			}
			flattenedFields = append(flattenedFields, flattenedField)
		}
	}

	return flattenedFields
}
