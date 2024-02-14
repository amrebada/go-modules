dev:
	air -c .air.toml server
swagger:
	go run main.go swagger
test:
	go test ./...
migrate:
	go run main.go migrate
build:
	GOOS=linux go build -o app -a main.go
build_mac:
	GOOS=darwin go build -o app -a main.go
rename:
	go run .scripts/rename.go -name=$(name) -old=go-template