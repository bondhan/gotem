# gotem
a go template

## migration

Migrate 1 version up <br>
`migrate -source "file://migration/scripts" -database "mysql://root:root@tcp(localhost:3305)/gotem_db" up 1`

Migrate 1 version down <br>
`migrate -source "file://migration/scripts" -database "mysql://root:root@tcp(localhost:3305)/gotem_db" down 1`

Get migration version<br>
`migrate -source "file://migration/scripts" -database "mysql://root:root@tcp(localhost:3305)/gotem_db" version`

Force version to n<br>
`migrate -source "file://migration/scripts" -database "mysql://root:root@tcp(localhost:3305)/gotem_db" force n`