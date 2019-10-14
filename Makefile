.PHONY: deploy
deploy:
	go test ./...
	go build -o shoshka-go
	git add shoshka-go
	git commit -m 'Compiled binary'
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
	go mod vendor
	go test ./...

.DEFAULT_GOAL := scrutinizer
