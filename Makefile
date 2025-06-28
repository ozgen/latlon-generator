build:
	@go build -o bin/latlon main.go

.PHONY: test

test:
	@for pkg in $$(go list ./...); do \
		echo "Testing $$pkg"; \
		go test -v $$pkg || exit 1; \
	done

run: build
	@./bin/app

debug:
	@dlv debug --headless --listen=:2345 --log --api-version=2 ./main.go


