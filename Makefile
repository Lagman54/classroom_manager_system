migration_up:
	migrate -path internal/classroom-app/migration/ -database "postgresql://postgres:s123@localhost:5432/classroom_app?sslmode=disable" -verbose up

migration_down:
	migrate -path internal/classroom-app/migration/ -database "postgresql://postgres:s123@localhost:5432/classroom_app?sslmode=disable" -verbose down

migration_fix:
	migrate -path internal/classroom-app/migration/ -database "postgresql://postgres:s123@localhost:5432/classroom_app?sslmode=disable" force 1