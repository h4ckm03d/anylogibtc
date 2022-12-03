.PHONY: dep gen docker migrate test 

test:
	go test -mod=readonly -v ./... -covermode=count -coverprofile=profile.out && go tool cover -func=profile.out

migrate:
	# todo

# setup dependencies
dep:
	go mod tidy 

gen:
	go generate ./...

compose:
	docker compose up -d