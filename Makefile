.PHONY: dep gen docker migrate test 

test:
	go test -mod=readonly -v ./... -covermode=count -coverprofile=profile.out && go tool cover -func=profile.out

integration-test:
	go test --tags=integration -mod=readonly -v ./... 

docker:
	docker compose up -d

# setup dependencies
dep:
	go mod tidy

fmt:
	gofumpt -w .

gen:
	go generate ./...

compose:
	docker compose up -d