run:
	go run ./app/main.go
# test:
# 	clear
# 	go test ./... -coverprofile=cover.out -p 1
# 	go tool cover -func=cover.out
# ifeq ($(mode), html)
# 	go tool cover -html=cover.out
# endif
build:
	cd ./app && go build -o ../
	