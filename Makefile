migration_up:
	migrate -path internal/migration/ -database "postgresql://postgres:s123@localhost:5432/classroom_app?sslmode=disable" -verbose up

migration_down:
	migrate -path internal/migration/ -database "postgresql://postgres:s123@localhost:5432/classroom_app?sslmode=disable" -verbose down

migration_fix:
	-path internal/migration/ -database "postgresql://postgres:s123@localhost:5432/classroom_app?sslmode=disable" force VERSION