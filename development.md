# Local development - TODO this document is a work in progress and should not be used at this time

## Set up Go

Forge4Flow-Core is written in Go. Prior to cloning the repo and making any code changes, please ensure that your local Go environment is set up. Refer to the appropriate instructions for your platform [here](https://go.dev/).

## Fork & clone repository

We follow GitHub's fork & pull request model. If you're looking to make code changes, it's easier to do so on your own fork and then contribute pull requests back to the Forge4Flow-Core repo. You can create your own fork of the repo [here](https://github.com/forge4flow/forge4flow-core/fork).

If you'd just like to check out the source to build & run, you can clone the repo directly:

```shell
git clone git@github.com:forge4flow/forge4flow-core.git
```

## Server configuration

To set up your server config file, refer to our [configuration doc](/configuration.md).

## Start development server

After the datastore, eventstore and configuration are set, build & start the server:

```shell
cd cmd/forge4flow-core
go run .
```

## Make requests

Once the server is running, you can make API requests using curl, any of the [Forge4Flow-Core SDKs](/README.md#sdks), or your favorite API client:

```shell
curl -g "http://localhost:port/v1/object-types" -H "Authorization: ApiKey YOUR_KEY"
```

# Running tests

TODO