GO_VERSION_SHORT:=$(shell echo `go version` | sed -E 's/.* go(.*) .*/\1/g')
ifneq ("1.19","$(shell printf "$(GO_VERSION_SHORT)\n1.19" | sort -V | head -1)")
$(error NEED GO VERSION >= 1.19. Found: $(GO_VERSION_SHORT))
endif

export GO111MODULE=on

SERVICE_NAME=transactions
SERVICE_PATH=.
ENV = .env

.PHONY: run
run:
	docker compose -f ./docker-compose.yaml --env-file $(ENV) up

.PHONY: run-svc
run-svc: #  make run-svc svc=redis
	docker compose -f ./docker-compose.yaml --env-file $(ENV) up $(svc)

.PHONY: stop
stop:
	docker compose -f ./docker-compose.yaml --env-file $(ENV) down

.PHONY: lint
lint:
	go vet ./...
	golangci-lint run

.PHONY: lint-fix
lint-fix:
	golangci-lint run --fix

.PHONY: test
test:
	go test -v -race -timeout 30s -coverprofile cover.out ./...
	@go tool cover -func cover.out | grep total | awk '{print $3}'

.PHONY: proto
proto:
	@buf lint
	@buf format -w .
	@buf build && buf generate proto


.PHONY: deps
deps: .deps

.deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.0
	# go install github.com/bufbuild/buf/cmd/buf@1.8.0
	go install google.golang.org/protobuf/cmd/protoc-gen-go@1.5.2
	go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@1.1.0
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
	go install github.com/envoyproxy/protoc-gen-validate@v0.9.0
	go install github.com/bufbuild/buf/cmd/protoc-gen-buf-lint@v0.56.0
	go mod download

.PHONY: build
.build:
		go mod download && CGO_ENABLED=0  go build \
			-o ./bin/balance ./cmd/balance/main.go
		go mod download && CGO_ENABLED=0  go build \
        			-o ./bin/user ./cmd/user/main.go

.PHONY: upgrade
upgrade:
	go get -u -t ./... && go mod tidy -v


.PHONY: mocks
mocks:
	cd tx-user/internal/handler; go generate;
	cd tx-balance/internal/handler; go generate;