VERSION=0.1.0
BUILD_DATE=$(shell date +%Y-%m-%d.%H:%M:%S)

LDFLAGS="-X main.VERSION=${VERSION}  -X main.BUILD_TIME=${BUILD_DATE} -X 'main.GO_VERSION=`go version`' -X main.APP_ENV=prod"

LDFLAGS_DEV="-X main.VERSION=${VERSION}  -X main.BUILD_TIME=${BUILD_DATE} -X 'main.GO_VERSION=`go version`' -X main.APP_ENV=dev"

build:
	CGO_ENABLED=1 go build -mod vendor -ldflags $(LDFLAGS) -o online_file .

dev:
	CGO_ENABLED=1 go build -mod vendor -ldflags $(LDFLAGS_DEV) -o online_file .

release:
	CGO_ENABLED=1 go build -mod vendor -ldflags $(LDFLAGS) -o online_file .

clean:
	rm -f online_file