# Running Warrant with PostgreSQL

This guide covers how to set up PostgreSQL as a datastore/eventstore for Warrant.

Note: Please first refer to the [development guide](/development.md) to ensure that your Go environment is set up and you have checked out the Warrant source or [downloaded a binary](https://github.com/forge4flow/forge4flow-core/releases).

## Install PostgreSQL

Install and run the [PostgreSQL Installer](https://www.postgresql.org/download/) for your OS to install and start PostgreSQL. For MacOS users, we recommend [installing PostgreSQL using homebrew](https://formulae.brew.sh/formula/postgresql@14).

## Warrant configuration

The Warrant server requires certain configuration, defined either within a `forge4flow.yaml` file (located within the same directory as the binary) or via environment variables. This configuration includes some common variables and some PostgreSQL specific variables. Here's a sample config:

### Sample `forge4flow.yaml` config

```yaml
port: 8000
logLevel: 1
enableAccessLog: true
autoMigrate: true
authentication:
  apiKey: replace_with_api_key
datastore:
  postgres:
    username: replace_with_username
    password: replace_with_password
    hostname: localhost
    database: warrant
    sslmode: disable
eventstore:
  synchronizeEvents: false
  postgres:
    username: replace_with_username
    password: replace_with_password
    hostname: localhost
    database: warrant_events
    sslmode: disable
```

Note: You must use 2 different databases for `datastore` and `eventstore`. You can create the databases via the postgres command line and configure them as the `database` attribute under datastore and eventstore.

The `synchronizeEvents` attribute in the eventstore section is false by default. Setting it to true means that all events will be tracked in order within the same transaction (helpful for testing locally).

## Running db migrations

Warrant uses [golang-migrate](https://github.com/golang-migrate/migrate) to manage sql db migrations. If the `autoMigrate` config flag is set to true, the server will automatically run migrations on start. If you prefer managing migrations and upgrades manually, please set the `autoMigrate` flag to false.

You can [install golang-migrate yourself](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) and run the PostgreSQL migrations manually:

```shell
migrate -path ./migrations/datastore/postgres/ -database postgres://username:password@hostname/warrant up
migrate -path ./migrations/eventstore/postgres/ -database postgres://username:password@hostname/warrant_events up
```
