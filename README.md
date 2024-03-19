# postpackage
A service to use to manage package's post options.

## Makefile targets
Build a docker image and a local executable for linux: `make build`

Start the system in a docker containers: `make start`

Stop containers: `make stop`

Run unit tests locally: `make test`

Run unit-tests in a docker container to avoid side-effects of the local configuration, etc.: `make test-docker` 

## Endpoints
All endpoints are located under "api/v1/" prefix 

To test different endpoints, import a Postman collection from `api/postman/postpackage.postman_collection.json`.
Also, you can use respective shell script from `api/curl` folder.
