APP_NAME=go-example-app

go-example-app:
	go build -o dist/${APP_NAME} .

releaser:
	/usr/local/bin/goreleaser release --skip-publish

clean:
	go clean
