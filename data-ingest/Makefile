test:
	- go test ./... -coverprofile cover.out
	- go tool cover -func cover.out | grep total

tidy:
	- go fmt ./...
