# Server configuration

Forge4Flow-Core requires certain configuration variables to be set via either a `forge4flow.yaml` config file or via environment variables. There is a set of common variables as well as datastore and eventstore-specific configuration.

## Common variables

| Variable                      | Description                                                                                                                                                    | Required?        | Default  | YAML                                             | ENV VAR                                  |
| ----------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------- | -------- | ------------------------------------------------ | ---------------------------------------- |
| `port`                        | Port where the server runs.                                                                                                                                    | no               | 8000     | `port: VALUE`                                    | `FORGE4FLOW_PORT=VALUE`                  |
| `logLevel`                    | Log level (e.g. Debug, Info etc.) for the server. Forge4Flow uses zerolog, valid log levels are defined [here](https://github.com/rs/zerolog#leveled-logging). | no               | 0        | `logLevel: VALUE`                                | `FORGE4FLOW_LOGLEVEL=VALUE`              |
| `enableAccessLog`             | Determines whether the built-in request logger is enabled or not.                                                                                              | no               | true     | `enableAccessLog: VALUE`                         | `FORGE4FLOW_ENABLEACCESSLOG=VALUE`       |
| `autoMigrate`                 | If set to `true`, the server will apply datastore and eventstore migrations before starting up.                                                                | no               | false    | `autoMigrate: VALUE`                             | `FORGE4FLOW_AUTOMIGRATE=VALUE`           |
| `coreInstall`                 | if set to `true`, the server will setup the objects required to use the `Forge4Flow-Dashboard` -- Will be deprecated in future release                         | no               | false    | `coreInstall: VALUE`                             | `FORGE4FLOW_COREINSTALL=VALUE`           |
| `flowNetwork`                 | String value of the Flow network to connect to, valid options are `emulator`, `testnet`, and `mainnet`                                                         | no               | emulator | `flowNetwork: VALUE`                             | `FORGE4FLOW_FLOWNETWORK=VALUE`           |
| `adminAccount`                | Used in connection with `coreInstall`, setups the initial admin account for the `Forge4Flow-Dashboard` -- Will be deprecated in future release                 | w/ `coreInstall` | -        | `adminAccount: VALUE`                            | `FORGE4FLOW_ADMINACCOUNT=VALUE`          |
| `authentication.autoRegister` | If set to `true`, the server with auto create users on their initial login.                                                                                    | no               | false    | `authentication:`<br>&emsp;`autoRegister: VALUE` | `FORGE4FLOW_AUTHENTICATION_AUTOREGISTER` |

## Server Side - API Key Authentication

By default, you must configure an API key that Forge4Flow will use to authenticate all requests. You should follow standard security practices for generating and storing your API key.

| Variable                | Description                                                                                              | Required? | Default | YAML                                       | ENV VAR                                  |
| ----------------------- | -------------------------------------------------------------------------------------------------------- | --------- | ------- | ------------------------------------------ | ---------------------------------------- |
| `authentication.apiKey` | The unique API key that all clients must pass to the server via the `Authorization: ApiKey VALUE` header | yes       | -       | `authentication:`<br>&emsp;`apiKey: VALUE` | `FORGE4FLOW_AUTHENTICATION_APIKEY=VALUE` |

## Client Side - Session Token Authentication

> TODO: Details about the client side auth configuration

## Set up datastore & eventstore

Forge4Flow is a stateful service that runs with an accompanying `datastore` and `eventstore` (for tracking resource & access events). Currently, only `MySQL` is supported, with `PostgreSQL` and `SQLite` (file and in-memory) support in the works. Refer to these guides to set up your desired database(s):

- [MySQL](/migrations/datastore/mysql/README.md)

Note: It's possible to use different dbs for the `datastore` and `eventstore` (e.g. mysql for datastore and sqlite for eventstore) but we recommend using the same type of db during development for simplicity.

Here is an example of a full server config using `mysql` for both the datastore and eventstore:

### Sample `forge4flow.yaml` config (place file in same dir as server binary)

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

### Sample environment variables config

```shell
export FORGE4FLOW_PORT=8000
export FORGE4FLOW_COREINSTALL=true
export FORGE4FLOW_FLOWNETWORK="emulator"
export FORGE4FLOW_ADMINACCOUNT="0xf8d6e0586b0a20c7"
export FORGE4FLOW_LOGLEVEL=1
export FORGE4FLOW_ENABLEACCESSLOG=true
export FORGE4FLOW_AUTOMIGRATE=true
export FORGE4FLOW_AUTHENTICATION_APIKEY="replace_with_api_key"
export FORGE4FLOW_AUTHENTICATION_AUTOREGISTER=true
export FORGE4FLOW_DATASTORE_MYSQL_USERNAME="replace_with_username"
export FORGE4FLOW_DATASTORE_MYSQL_PASSWORD="replace_with_password"
export FORGE4FLOW_DATASTORE_MYSQL_HOSTNAME="replace_with_hostname"
export FORGE4FLOW_DATASTORE_MYSQL_DATABASE="forge4flow"
export FORGE4FLOW_EVENTSTORE_SYNCHRONIZEEVENTS=false
export FORGE4FLOW_EVENTSTORE_MYSQL_USERNAME="replace_with_username"
export FORGE4FLOW_EVENTSTORE_MYSQL_PASSWORD="replace_with_password"
export FORGE4FLOW_EVENTSTORE_MYSQL_HOSTNAME="replace_with_hostname"
export FORGE4FLOW_EVENTSTORE_MYSQL_DATABASE="forge4flowEvents"
```
