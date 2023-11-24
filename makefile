openapi_http:
	oapi-codegen -generate types -o internal/ports/openapi_types.gen.go -package ports api/openapi/purchasing.yaml
	oapi-codegen -generate fiber -o internal/ports/openapi_api.gen.go -package ports api/openapi/purchasing.yaml

gen-sqlc:
	cd ./resources && sqlc generate