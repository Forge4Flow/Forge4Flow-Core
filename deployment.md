# Deployment examples

Sample deployment configurations for running Forge4Flow-Core.

## Docker Compose

### Using MySQL as the datastore & eventstore

This guide will cover how to self-host Forge4Flow-Core with MySQL as the datastore and eventstore. Note that Forge4Flow-Core only supports versions of MySQL >= 8.0.32.

The following [Docker Compose](https://docs.docker.com/compose/) manifest will create a MySQL database, setup the database schema required by Forge4Flow-Core, and start Forge4Flow-Core. You can also accomplish this by running Forge4Flow-Core with [Kubernetes](https://kubernetes.io/):

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
    image: Forge4Flow-Core/Forge4Flow-Core
    ports:
      - 8000:8000
    depends_on:
      database:
        condition: service_healthy
    environment:
      Forge4Flow-Core_PORT: 8000
      Forge4Flow-Core_LOGLEVEL: 1
      Forge4Flow-Core_ENABLEACCESSLOG: true
      Forge4Flow-Core_AUTOMIGRATE: true
      Forge4Flow-Core_AUTHENTICATION_APIKEY: replace_with_api_key
      Forge4Flow-Core_DATASTORE_MYSQL_USERNAME: replace_with_username
      Forge4Flow-Core_DATASTORE_MYSQL_PASSWORD: replace_with_password
      Forge4Flow-Core_DATASTORE_MYSQL_HOSTNAME: database
      Forge4Flow-Core_DATASTORE_MYSQL_DATABASE: Forge4Flow-Core
      Forge4Flow-Core_EVENTSTORE_SYNCHRONIZEEVENTS: false
      Forge4Flow-Core_EVENTSTORE_MYSQL_USERNAME: replace_with_username
      Forge4Flow-Core_EVENTSTORE_MYSQL_PASSWORD: replace_with_password
      Forge4Flow-Core_EVENTSTORE_MYSQL_HOSTNAME: database
      Forge4Flow-Core_EVENTSTORE_MYSQL_DATABASE: Forge4Flow-CoreEvents
```
