# Go Modules

**go modules** is a boilerplate for **Nest.js** like, it has modules, controllers, REST generator and much more...

## What is included:
1. `dotenv` support
1. [gorm](https://gorm.io/) database support with entities, auto migration
1. `postgresql` database
1. [gofiber](https://gofiber.io/) support for web server, middlewares ...
1.  🔎 `swagger generating` *comming soon*


## Prequisites
- nodemon `npm install -g nodemon`
- go (https://golang.org/doc/install)
- make

## Scripts
- development: `make dev`
- rename name of the project: update `Makefile` with replace <name> -> `make rename`
- build for linux: `make build`
- build for mac: `make build_mac`
- run database migration: `make migrate`
- run tests: `make test`
- generate swagger file: `make swagger` *comming soon*
