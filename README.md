# gotem


## migration
migrate -source "file://migration/scripts" -database "mysql://root:root@tcp(localhost:3305)/gotem_db" up 1
migrate -source "file://migration/scripts" -database "mysql://root:root@tcp(localhost:3305)/gotem_db" down 1
migrate -source "file://migration/scripts" -database "mysql://root:root@tcp(localhost:3305)/gotem_db" version