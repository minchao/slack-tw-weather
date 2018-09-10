.PHONY: build clean deps deploy lint

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/slash slash/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean deps lint build
	sls deploy --verbose

deps:
	which dep || (curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh)
	which gometalinter || (go get -u -v github.com/alecthomas/gometalinter && gometalinter --install)
	dep ensure -v

lint:
	gometalinter --exclude=vendor --disable-all --enable=vet --enable=vetshadow --enable=gofmt ./...
