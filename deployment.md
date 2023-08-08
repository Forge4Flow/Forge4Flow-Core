# Deployment examples

Sample deployment configurations for running Forge4Flow-Core.

## Docker Compose

### Using MySQL as the datastore & eventstore

This guide will cover how to self-host Forge4Flow-Core with MySQL as the datastore and eventstore. Note that Forge4Flow-Core only supports versions of MySQL >= 8.0.32.

The following [Docker Compose](https://docs.docker.com/compose/) manifest will create a MySQL database, set up the database schema required by Forge4Flow-Core, and start Forge4Flow-Core. You can also accomplish this by running Forge4Flow-Core with [Kubernetes](https://kubernetes.io/):

```yaml
version: "3.9"
services:
  database:
    image: mysql:8.0.32
    environment:
      MYSQL_USER: replace_with_username
      MYSQL_PASSWORD: replace_with_password
    ports:
      - 3306:3306
    healthcheck:
      test: ["CMD", "mysqladmin", "ping", "-h", "localhost"]
      timeout: 5s
      retries: 10

  web:
    image: forge4flow/forge4flow-core:latest
    ports:
      - 8000:8000
    depends_on:
      database:
        condition: service_healthy
    environment:
      FORGE4FLOW_PORT: 8000
      FORGE4FLOW_COREINSTALL: true
      FORGE4FLOW_FLOWNETWORK: emulator"
      FORGE4FLOW_ADMINACCOUNT: 0xf8d6e0586b0a20c7"
      FORGE4FLOW_LOGLEVEL: 1
      FORGE4FLOW_ENABLEACCESSLOG: true
      FORGE4FLOW_AUTOMIGRATE: true
      FORGE4FLOW_AUTHENTICATION_APIKEY: replace_with_api_key"
      FORGE4FLOW_AUTHENTICATION_AUTOREGISTER: true
      FORGE4FLOW_DATASTORE_MYSQL_USERNAME: replace_with_username"
      FORGE4FLOW_DATASTORE_MYSQL_PASSWORD: replace_with_password"
      FORGE4FLOW_DATASTORE_MYSQL_HOSTNAME: replace_with_hostname"
      FORGE4FLOW_DATASTORE_MYSQL_DATABASE: warrant"
      FORGE4FLOW_EVENTSTORE_SYNCHRONIZEEVENTS: false
      FORGE4FLOW_EVENTSTORE_MYSQL_USERNAME: replace_with_username"
      FORGE4FLOW_EVENTSTORE_MYSQL_PASSWORD: replace_with_password"
      FORGE4FLOW_EVENTSTORE_MYSQL_HOSTNAME: replace_with_hostname"
      FORGE4FLOW_EVENTSTORE_MYSQL_DATABASE: warrantEvents"
```
