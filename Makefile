APP_NAME="go-example-app"

go-example-app:
	go build -o dist/${APP_NAME} .

clean:
	go clean
