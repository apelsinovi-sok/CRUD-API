## Run Migration
goose postgres "user=user password=password dbname=database sslmode=disable" up

## Reset migration
goose postgres "user=user password=password dbname=database sslmode=disable" down
