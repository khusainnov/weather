m-up:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5434/postgres?sslmode=disable' up
m-down:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5434/postgres?sslmode=disable' down

d-up:
	docker run --name=edulab -e POSTGRES_PASSWORD='qwerty' -p 5434:5432 -d --rm postgres

d-exec:
	docker exec -it e39640a90ec3 /bin/bash
d-stop:
	docker stop e39640a90ec3