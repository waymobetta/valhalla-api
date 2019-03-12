all: deploy

.PHONY: vendor
vendor:
	@go mod vendor

.PHONY: test
test:
	@go test -v ./...

.PHONY: test/local
test/local:
	(. .env.local && go test -v controllers/* $(ARGS))

.PHONY: test/auth
test/auth:
	@go test -v auth/*.go $(ARGS)

.PHONY: test/controllers
test/controllers:
	@go test -v controllers/*.go $(ARGS)

.PHONY: start
start:
	@go run cmd/valhalla/main.go

start/local:
	@(. .env.local && go run cmd/valhalla/main.go)

start/staging:
	@(. .env.staging && go run cmd/valhalla/main.go)

start/prod:
	@(. .env.prod && go run cmd/.go)

.PHONY: start/docs
start/docs:
	@(cd web/documentation && python -m SimpleHTTPServer 8000)

deploy: build compress
	@echo "deploying..\n"
	@MAKE done

.PHONY: build
build:
	@echo "building.."
	@command go build -ldflags "-s -w" -o cmd/valhalla cmd/main.go
	$(MAKE) compress
	
.PHONY: build/docker
build/docker:
	@docker build -t valhalla/api:latest .

.PHONY: compress
compress:
	@echo "compressing.."
	@command upx cmd/valhalla

.PHONY: done
done:
	@echo "done"

.PHONY: goa
goa:
	@goagen bootstrap -d github.com/waymobetta/valhalla-api/design
	@rm main.go
	@rm helse.go
	@rm godkjent.go
	@MAKE swagger

.PHONY: swagger
swagger:
	@echo "generating swagger spec"
	@(goagen swagger -d github.com/waymobetta/valhalla-api/design && cp swagger/swagger.json web/documentation/swagger.json)
