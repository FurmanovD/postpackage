SHELL := /bin/sh

export APP_NAME=postpackage
export TEST_CONTAINER=${APP_NAME}-tests
# constructed
export APP_VERSION=$(shell cat ./VERSION)
export IMAGE_NAME ?= ${APP_NAME}
export IMAGE_TAG = ${IMAGE_NAME}:${APP_VERSION}
export IMAGE_TAG_LATEST = ${IMAGE_NAME}:latest
# construct a build information to set it to a binary file variables at 'go build' step:
export APP_VERSION=$(shell cat ./VERSION)
export BUILD_TIME=$$(date -u "+%F_%T")
export GIT_COMMIT=$$(git log -1 --format="%H")

# builds an image and a binary file
build: clean
	@echo "building ${APP_NAME}"
	@echo "APP_VERSION=${APP_VERSION}"
	@echo "BUILD_TIME=${BUILD_TIME}"
	@echo "GIT_COMMIT=${GIT_COMMIT}"
	APP_NAME=${APP_NAME} docker build --no-cache \
		--build-arg APP_NAME=${APP_NAME} \
		--build-arg APP_VERSION=${APP_VERSION} \
		--build-arg BUILD_TIME=${BUILD_TIME} \
		--build-arg GIT_COMMIT=${GIT_COMMIT} \
		--tag ${IMAGE_TAG} \
		--tag ${IMAGE_TAG_LATEST} \
		--file Dockerfile.build .

	CID=$$(docker create ${IMAGE_TAG}) && \
	docker cp $${CID}:/app/${APP_NAME} ${APP_NAME} && \
	docker rm $${CID}

clean: stop
	@echo "cleaning"
	docker rmi ${IMAGE_TAG} || echo ""
	docker rmi ${IMAGE_TAG_LATEST} || echo ""
	rm -f ${APP_NAME} || echo ""
	@echo "clean complete"

start: stop build
	@echo "starting..."
	IMAGE_TAG=${IMAGE_TAG} docker-compose build api && \
	IMAGE_TAG=${IMAGE_TAG} docker-compose up -d

stop:
	docker-compose -f docker-compose.yaml down --remove-orphans

deps-start:
	IMAGE_TAG=${IMAGE_TAG} docker-compose -f docker-compose.yaml up -d mysqldb

test: 
	go test ./... -v -cover -failfast
	
test-docker: 
	@docker rmi ${TEST_CONTAINER} || echo ""
	TEST_CONTAINER=${TEST_CONTAINER} docker-compose -f docker-compose.yaml -f docker-compose.tests.yaml up tests \
	&& docker-compose -f docker-compose.yaml -f docker-compose.tests.yaml down

.PHONY: build clean start stop deps-start test test-docker