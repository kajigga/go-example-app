APP_NAME=todoer

todoer:
	go build -o dist/${APP_NAME} .

releaser:
	/usr/local/bin/goreleaser release --skip-publish --rm-dist

clean:
	go clean
