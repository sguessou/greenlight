run:
	go run ./cmd/api

psql:
	psql ${GREENLIGHT_DB_DSN}

migration:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

up:
	@echo 'running up migrations...'
	migrate -path ./migrations -database ${GREENLIGHT_DB_DSN} up
