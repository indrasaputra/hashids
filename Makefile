test:
	go test -v -race ./...

dep:
	env GO111MODULE=on go mod download

tidy:
	env GO111MODULE=on go mod tidy

vendor:
	env GO111MODULE=on go mod vendor

cover:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out