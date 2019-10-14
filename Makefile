.PHONY: deploy
deploy:
	go test ./...
	go build -o shoshka-go
	git push origin master

.PHONY: test
test:
	go test ./...

.PHONY: dev
dev:
	go test ./...
	go run main.go

.PHONY: scrutinizer
scrutinizer:
	cp .env.example .env
	cp storage/errors.example storage/errors
	cp storage/titles.example storage/titles
	go test ./...

.DEFAULT_GOAL := scrutinizer
