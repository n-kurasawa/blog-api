.PHONY: build clean deploy undeploy

build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/query cmd/lambda.go

clean:
	rm -rf ./bin ./vendor

deploy: clean build
	sls deploy --verbose

undeploy:
	sls remove --verbose
