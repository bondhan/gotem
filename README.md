# gotem
a (try to be complete) go template. Includes:
- echo framework (interceptor, middleware)
- redis 
- rabbitmq
- db migration
- app server

# running

make the db persistence:
- `docker volume create --name=postgresql`
- `docker volume create --name=postgresql_data`

starting the service as service:
- `docker-compose up -d`

What it will do:
 - run postgres
 - run redis
 - run rabbitmq
 - apply database migration (up)
 - run app server

## migration

if you need manual migration:

- Migrate version up <br>
`migrate -database "postgres://root:root@localhost:54321/gotem_db?sslmode=disable" -path migration/scripts/ up`

- Migrate version down <br>
`migrate -database "postgres://root:root@localhost:54321/gotem_db?sslmode=disable" -path migration/scripts/ down`

## healthcheck

- Get the container ID from gotem app:<br>
`bondhan@syuhada:~/go/src/github.com/bondhan/gotem$ docker ps -aqf "name=gotem_gotem"`<br>
`c33e4831e74d`

- Inspect the health status:<br>
`docker inspect c33 | jq '.[].State.Health'`