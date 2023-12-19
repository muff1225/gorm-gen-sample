# Generate Model
# https://gorm.io/gen/gen_tool.html
# generating all of associations is required multiple execute angenerate_associations/main.go  .
.PHONY: generate-models
generate-models:
	@rm -rf ./interface/db/query/*.gen.go ./interface/db/model/*.gen.go
	@go run ./hack/gormgen/generate_all/main.go
	@go run ./hack/gormgen/generate_associations/main.go

## Migrate to database from SQL files.
.PHONY: migrate-up
migrate-up:
	@migrate -path ./sql -database 'mysql://${MYSQL_ROOT_USER}:${MYSQL_ROOT_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}?charset=utf8' up