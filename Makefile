install: 
	go install -v

tidy: 
	go mod tidy

.PHONY: install tidy