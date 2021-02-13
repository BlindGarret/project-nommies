install:
	go install -v

tidy:
	go mod tidy

test:
	go test ./...

gen-spec:
	swagger-cli bundle ./open-api/swagger.yml --outfile ./open-api/generated.yml --type yaml

validate: gen-spec
	swagger validate ./open-api/generated.yml

gen: validate
	swagger generate server \
		--target=./generated/api \
		--spec=./open-api/generated.yml \
		--exclude-main \
		--name=recipes

clean-gen: 
	rm -rf ./generated/api/*

infrastructure:
	docker-compose up -d

.PHONY: install tidy test validate gen clean-gen infrastructure