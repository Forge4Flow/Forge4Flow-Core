# Running Forge4Flow with MySQL

This guide covers how to set up MySQL as a datastore/eventstore for Forge4Flow.

Note: Please first refer to the [development guide](/development.md) to ensure that your Go environment is set up and you have checked out the Forge4Flow source or [downloaded a binary](https://github.com/forge4flow/forge4flow-core/releases).

## Install MySQL

Follow the [MySQL Installation Guide](https://dev.mysql.com/doc/mysql-installation-excerpt/8.0/en/) for your OS to install and start MySQL. For MacOS users, we recommend [installing MySQL using homebrew](https://formulae.brew.sh/formula/mysql).

## Forge4Flow configuration

The Forge4Flow server requires certain configuration, defined either within a `forge4flow.yaml` file (located within the same directory as the binary) or via environment variables. This configuration includes some common variables and some MySQL specific variables. Here's a sample config:

### Sample `forge4flow.yaml` config

```yaml
port: 8000
coreInstall: true
flowNetwork: emulator
adminAccount: "0xf8d6e0586b0a20c7"
logLevel: 1
enableAccessLog: true
autoMigrate: true
authentication:
  apiKey: your_api_key
  autoRegister: true
datastore:
  mysql:
    username: forge4flow
    password: forge4flow
    hostname: 127.0.0.1
    database: forge4flow
eventstore:
  synchronizeEvents: false
  mysql:
    username: forge4flow
    password: forge4flow
    hostname: 127.0.0.1
    database: forge4flowEvents
```

Note: You must use 2 different databases for `datastore` and `eventstore`. You can create the databases via the mysql command line and configure them as the `database` attribute under datastore and eventstore.

The `synchronizeEvents` attribute in the eventstore section is false by default. Setting it to true means that all events will be tracked in order within the same transaction (helpful for testing locally).

## Running db migrations

Forge4Flow uses [golang-migrate](https://github.com/golang-migrate/migrate) to manage sql db migrations. If the `autoMigrate` config flag is set to true, the server will automatically run migrations on start. If you prefer managing migrations and upgrades manually, please set the `autoMigrate` flag to false.

You can [install golang-migrate yourself](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) and run the MySQL migrations manually:

```shell
migrate -path ./migrations/datastore/mysql/ -database mysql://username:password@hostname/forge4flow up
migrate -path ./migrations/eventstore/mysql/ -database mysql://username:password@hostname/forge4flowEvents up
```
