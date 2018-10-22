.PHONY: build clean coverage coverage-view deps deploy lint test

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/weather cmd/weather/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/funcRadar cmd/funcs/radar/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

coverage:
	go test -v ./... -race -coverprofile=coverage.out -covermode=atomic

coverage-view:
	go tool cover -html=coverage.out

deploy: clean deps lint test build
	sls deploy --verbose

deps:
	which dep || (curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh)
	which gometalinter || (go get -u -v github.com/alecthomas/gometalinter && gometalinter --install)
	dep ensure -v

lint:
	gometalinter --exclude=vendor --disable-all --enable=vet --enable=vetshadow --enable=gofmt ./...

test:
	go test -v ./... -race
