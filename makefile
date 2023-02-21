update:
	swag init -g cmd/main.go --output docs/swagger
	curl -X POST https://converter.swagger.io/api/convert \
		-d @docs/swagger/swagger.json \
		--header 'Content-Type: application/json' > docs/swagger/openapi.json

build: update
	go mod tidy
	go build cmd/main.go

run: update
	go run cmd/main.go