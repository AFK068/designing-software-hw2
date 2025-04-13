.PHONY: run
run:
	@go run cmd/run/main.go

.PHONY: imports
imports:
	@goimports-reviser -project-name github.com/AFK068/bot -file-path ./... -separate-named

.PHONY: generate_openapi
generate_openapi:
	@mkdir -p internal/api/openapi/hsezoo/v1
	@oapi-codegen -package v1 \
		-generate server,types \
		api/openapi/v1/hsezoo.yaml > internal/api/openapi/hsezoo/v1/hsezoo-api.gen.go