deploy:
	go test ./...
	go build -o shoshka-go
	git push origin master

test:
	go test ./...

dev:
	go test ./...
	go run main.go
