# gotem
a (try to be complete) go template. Includes:
- echo framework (interceptor, middleware)
- redis 
- rabbitmq
- db migration
- app server

# running

`docker-compose up -d`

What it will do:
 - run postgres
 - run redis
 - run rabbitmq
 - apply database migration (up)
 - run app server

## migration

if you need manual migration:

Migrate version up <br>
`migrate -database "postgres://root:root@localhost:54321/gotem_db?sslmode=disable" -path migration/scripts/ up`

Migrate version down <br>
`migrate -database "postgres://root:root@localhost:54321/gotem_db?sslmode=disable" -path migration/scripts/ down`