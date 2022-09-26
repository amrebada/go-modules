dev:
	nodemon --exec go run main.go -env dev --signal SIGTERM
test:
	go test ./...
migrate:
	go run main.go -migrate true
build:
	GOOS=linux go build -o app -a main.go
build_mac:
	GOOS=darwin go build -o app -a main.go
rename:
	/bin/bash ./.scripts/rename.sh . go-template <name>