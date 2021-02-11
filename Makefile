install:
	go install -v

tidy:
	go mod tidy

test:
	go test ./...

validate:
	swagger validate ./oas/swagger.yml

gen: validate
	swagger generate server \
		--target=./generated/api \
		--spec=./oas/swagger.yml \
		--exclude-main \
		--name=recipes

clean-gen: 
	rm -rf ./generated/api/*

infrastructure:
	docker-compose up -d

.PHONY: install validate gen tidy infrastructure test