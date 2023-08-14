dev:
	nodemon --exec go run main.go -env dev --signal SIGTERM
swagger:
	go run main.go -sw
test:
	go test ./...
migrate:
	go run main.go -m true
build:
	GOOS=linux go build -o app -a main.go
build_mac:
	GOOS=darwin go build -o app -a main.go
rename:
	go run .scripts/rename.go -name=$(name) -old=go-template