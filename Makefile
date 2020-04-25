GOBIN=$(shell go env GOROOT)/bin/go
TARGET = oea-go
DOCKER_TAG = $(TARGET):latest
TARGET_LOCAL_PATH = docker/$(TARGET)
#TARGET_HOST = artsaf.in
#TARGET_PATH = $(TARGET)/
VERSION = $(shell git rev-parse --short HEAD)

.PHONY: all

all: image

assets:
	test -f resources/bootstrap.min.css || curl -s https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css -o resources/bootstrap.min.css

build: assets
	$(shell go env GOPATH)/bin/go-bindata -o "common/bindata.go" -pkg "common" resources/ resources/partials/

	CGO_ENABLED=0 GOOS=linux \
	$(GOBIN) build -i -installsuffix cgo -o "$(TARGET_LOCAL_PATH)" -ldflags "-X main.AppVersion=$(VERSION)"

image: build
	docker build -t $(DOCKER_TAG) docker/

clean:
	rm -fv $(TARGET_LOCAL_PATH)

#deploy: $(TARGET)
#	ssh $(TARGET_HOST) mkdir -pv $(TARGET_PATH)
#	rsync $(TARGET) config.yaml .env $(TARGET_HOST):$(TARGET_PATH)
