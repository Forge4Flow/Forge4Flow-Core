# Running Warrant with SQLite

This guide covers how to set up SQLite as a datastore/eventstore for Warrant.

Note: Please first refer to the [development guide](/development.md) to ensure that your Go environment is set up and you have checked out the Warrant source or [downloaded a binary](https://github.com/forge4flow/forge4flow-core/releases).

## Install SQLite

Many operating systems (like MacOS) come with SQLite pre-installed. If you already have SQLite installed, you can skip to the next step. If you don't already have SQLite installed, [install it](https://www.tutorialspoint.com/sqlite/sqlite_installation.htm). Once installed, you should be able to run the following command to print the currently installed version of SQLite:

```bash
sqlite3 --version
```

## Warrant configuration

The Warrant server requires certain configuration, defined either within a `forge4flow.yaml` file (located within the same directory as the binary) or via environment variables. This configuration includes some common variables and some SQLite specific variables. Here's a sample config:

### Sample `forge4flow.yaml` config

```yaml
port: 8000
logLevel: 1
enableAccessLog: true
autoMigrate: true
authentication:
  apiKey: replace_with_api_key
datastore:
  sqlite:
    database: warrant
    inMemory: true
eventstore:
  synchronizeEvents: false
  sqlite:
    database: warrantEvents
    inMemory: true
```

Note: By default, SQLite will create a database file for both the database and eventstore. The filenames are configurable using the `database` property under `datastore` and `eventstore`. Specifying the `inMemory` option under `datastore` or `eventstore` will bypass creation of a database file and run the SQLite database completely in memory. When running Warrant with the `inMemory` configuration, **any data in Warrant will be lost once the Warrant process is shutdown/killed**.

The `synchronizeEvents` attribute in the eventstore section is false by default. Setting it to true means that all events will be tracked in order within the same transaction (helpful for testing locally).

Unlike `mysql` and `postgresql`, `sqlite` currently does not support manually running db migrations on the command line via golang-migrate. Therefore, you should keep `autoMigrate` set to true in your Warrant config so that the server runs migrations as part of startup.
