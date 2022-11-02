run:
	go run ./app/main.go
test:
	go test ./app/... -coverprofile=cover.out -v
	go tool cover -func=cover.out
ifeq ($(mode), html)
	go tool cover -html=cover.out
endif
build:
	cd ./app && go build -o ../
	