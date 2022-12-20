# database name
DB_NAME ?= postgres

# database type
DB_TYPE ?= postgres

# database username
DB_USER ?= postgres

# database password
DB_PWD ?= mysecretpassword

# psql URL
IP=`docker inspect --format '{{ .NetworkSettings.IPAddress }}' test-postgres`

PSQLURL ?= $(DB_TYPE)://$(DB_USER):$(DB_PWD)@$(IP):5432/$(DB_NAME)

# sqlc yaml file
SQLC_YAML ?= ./sqlc.yaml

.PHONY : postgresup postgresdown psql createdb teardown_recreate generate

postgresup:
	docker run --name test-postgres -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres

postgresdown:
	docker stop test-postgres  || true && 	docker rm test-postgres || true

psql:
	docker run -v $(PWD):/usr/share/chapter5 -it --rm jbergknoff/postgresql-client -p 5432:5432 $(PSQLURL)

# task to create database without typing it manually
createdb:
	echo $(PWD)
	docker run -v $(PWD):/usr/share/chapter5 -it --rm jbergknoff/postgresql-client -p 5432:5432 $(PSQLURL) -c "\i /usr/share/chapter5/db/schema.sql"

teardown_recreate: postgresdown postgresup
	sleep 5
	$(MAKE) createdb

generate:
	@echo "Generating Go models with sqlc "
	sqlc generate -f $(SQLC_YAML)