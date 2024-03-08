templ:
	templ generate

tailwind:
	tailwindcss -i web/src/index.css -o web/public/styles.css --minify --watch
	
postgres:
	docker run --name portfolio -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123 -d postgres:16-alpine

createdb: 
	docker exec -it portfolio createdb --username=root --owner=root myportfolio

dropdb: 
	docker exec -it portfolio dropdb myportfolio

migrateup:
	migrate -path database/schema -database "postgresql://root:123@localhost:5432/myportfolio?sslmode=disable" -verbose up

migratedown:
	migrate -path database/schema -database "postgresql://root:123@localhost:5432/myportfolio?sslmode=disable" -verbose down

sqlc: 
	sqlc generate

.PHONY: templ tailwind postgres createdb dropdb migrateup migratedown sqlc

