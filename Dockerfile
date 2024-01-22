# Stage 1: Build the backend
FROM golang:1.20 AS backend-builder

WORKDIR /forge4flow
COPY . .
RUN go mod download

RUN GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -v -o forge4flow-core cmd/forge4flow-core/main.go

# Stage 2: Import OpenFaaS of-watchdog
FROM ghcr.io/openfaas/of-watchdog:0.9.13 as watchdog

# Stage 3: Create the final image
FROM alpine:latest

RUN addgroup -S forge4flow-core && adduser -S forge4flow-core -G forge4flow-core
USER forge4flow-core

# Install the watchdog from the base image
COPY --from=watchdog /fwatchdog /usr/bin/

# Install Forge4Flow-Core from the backend-builder
WORKDIR /forge4flow
COPY --from=backend-builder /forge4flow/forge4flow-core .

# Now set the watchdog as the start-up process
# Along with the HTTP mode, and an upstream URL to 
# where your HTTP server will be running from the original
# image.
ENV mode="http"
ENV upstream_url="http://127.0.0.1:8000"

# Set fprocess to forge4flow-core application
ENV fprocess="./forge4flow-core"

EXPOSE 8080

# Start watchdog and forge4flow-core
CMD ["fwatchdog"]