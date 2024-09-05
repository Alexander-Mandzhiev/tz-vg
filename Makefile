create migrate:
	migrate -path ./db/migrations -database 'postgres://postgres:root@localhost:5432/tasks?sslmode=disable' up
drop migrate:
	migrate -path ./db/migrations -database 'postgres://postgres:root@localhost:5432/tasks?sslmode=disable' down