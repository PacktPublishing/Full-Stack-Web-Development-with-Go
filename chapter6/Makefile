# database name
DB_NAME ?= fullstackdb

# database type
DB_TYPE ?= postgres

# database username
DB_USER ?= local

# database password
DB_PWD ?= asecurepassword

# psql URL
IP=`docker inspect --format '{{ .NetworkSettings.IPAddress }}' test-postgres`

PSQLURL ?= $(DB_TYPE)://$(DB_USER):$(DB_PWD)@$(IP):5432/$(DB_NAME)

.PHONY : postgresup postgresdown psql createdb teardown_recreate generate

postgresup:
	docker run --name test-postgres -e POSTGRES_PASSWORD=$(DB_PWD) -p 5432:5432 -d postgres

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
	go generate