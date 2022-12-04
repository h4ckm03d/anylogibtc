.PHONY: dep gen docker migrate test 

test:
	go test -mod=readonly -v ./... -covermode=count -coverprofile=profile.out && go tool cover -func=profile.out

migrate:
	soda migrate -e development

# setup dependencies
dep:
	go mod tidy
	go install -tags sqlite github.com/gobuffalo/pop/v6/soda@latest

gen:
	go generate ./...

compose:
	docker compose up -d