# Deployment examples

Sample deployment configurations for running Auth4Flow-Core.

## Docker Compose

### Using MySQL as the datastore & eventstore

This guide will cover how to self-host Auth4Flow-Core with MySQL as the datastore and eventstore. Note that Auth4Flow-Core only supports versions of MySQL >= 8.0.32.

The following [Docker Compose](https://docs.docker.com/compose/) manifest will create a MySQL database, setup the database schema required by Auth4Flow-Core, and start Auth4Flow-Core. You can also accomplish this by running Auth4Flow-Core with [Kubernetes](https://kubernetes.io/):

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
    image: Auth4Flow-Core/Auth4Flow-Core
    ports:
      - 8000:8000
    depends_on:
      database:
        condition: service_healthy
    environment:
      Auth4Flow-Core_PORT: 8000
      Auth4Flow-Core_LOGLEVEL: 1
      Auth4Flow-Core_ENABLEACCESSLOG: true
      Auth4Flow-Core_AUTOMIGRATE: true
      Auth4Flow-Core_AUTHENTICATION_APIKEY: replace_with_api_key
      Auth4Flow-Core_DATASTORE_MYSQL_USERNAME: replace_with_username
      Auth4Flow-Core_DATASTORE_MYSQL_PASSWORD: replace_with_password
      Auth4Flow-Core_DATASTORE_MYSQL_HOSTNAME: database
      Auth4Flow-Core_DATASTORE_MYSQL_DATABASE: Auth4Flow-Core
      Auth4Flow-Core_EVENTSTORE_SYNCHRONIZEEVENTS: false
      Auth4Flow-Core_EVENTSTORE_MYSQL_USERNAME: replace_with_username
      Auth4Flow-Core_EVENTSTORE_MYSQL_PASSWORD: replace_with_password
      Auth4Flow-Core_EVENTSTORE_MYSQL_HOSTNAME: database
      Auth4Flow-Core_EVENTSTORE_MYSQL_DATABASE: Auth4Flow-CoreEvents
```

### Using PostgreSQL as the datastore & eventstore

This guide will cover how to self-host Auth4Flow-Core with PostgreSQL as the datastore and eventstore. Note that Auth4Flow-Core only supports versions of PostgreSQL >= 14.7.

The following [Docker Compose](https://docs.docker.com/compose/) manifest will create a PostgreSQL database, setup the database schema required by Auth4Flow-Core, and start Auth4Flow-Core. You can also accomplish this by running Auth4Flow-Core with [Kubernetes](https://kubernetes.io/):

```yaml
version: "3.9"
services:
  database:
    image: postgres:14.7
    environment:
      POSTGRES_PASSWORD: replace_with_password
    ports:
      - 5432:5432
    healthcheck:
      test: ["CMD", "pg_isready", "-d", "Auth4Flow-Core"]
      timeout: 5s
      retries: 10

  web:
    image: Auth4Flow-Coredev/Auth4Flow-Core
    ports:
      - 8000:8000
    depends_on:
      database:
        condition: service_healthy
    environment:
      Auth4Flow-Core_PORT: 8000
      Auth4Flow-Core_LOGLEVEL: 1
      Auth4Flow-Core_ENABLEACCESSLOG: true
      Auth4Flow-Core_AUTOMIGRATE: true
      Auth4Flow-Core_AUTHENTICATION_APIKEY: replace_with_api_key
      Auth4Flow-Core_DATASTORE_POSTGRES_USERNAME: postgres
      Auth4Flow-Core_DATASTORE_POSTGRES_PASSWORD: replace_with_password
      Auth4Flow-Core_DATASTORE_POSTGRES_HOSTNAME: database
      Auth4Flow-Core_DATASTORE_POSTGRES_DATABASE: Auth4Flow-Core
      Auth4Flow-Core_DATASTORE_POSTGRES_SSLMODE: disable
      Auth4Flow-Core_EVENTSTORE_SYNCHRONIZEEVENTS: false
      Auth4Flow-Core_EVENTSTORE_POSTGRES_USERNAME: postgres
      Auth4Flow-Core_EVENTSTORE_POSTGRES_PASSWORD: replace_with_password
      Auth4Flow-Core_EVENTSTORE_POSTGRES_HOSTNAME: database
      Auth4Flow-Core_EVENTSTORE_POSTGRES_DATABASE: Auth4Flow-Core_events
      Auth4Flow-Core_EVENTSTORE_POSTGRES_SSLMODE: disable
```
