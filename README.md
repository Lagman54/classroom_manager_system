# Classroom management system


```
Author: Aibar Kudiyarkhan 
ID: 22B030553
```

## How to run the app
```
go run ./cmd/classroom-app `
-dsn="postgres://postgres:s123@localhost:5432/classroom_app?sslmode=disable" `
-migrations=file://internal/classroom-app/migration `
-fill=true `
-env=development `
-port=8081
```

### List of flags
```
dsn — postgress connection string with username, password, address, port, database name, and SSL mode. Default: Value is not correct by security reasons.

migrations — Path to folder with migration files. If not provided, migrations do not applied.

fill — Fill database with dummy data. Default: false.

env - App running mode. Default: development

port - App port. Default: 8081
```

## Connect to server
```
https://octopus-app-a8j68.ondigitalocean.app/
```

## Classroom REST API
```
GET /classes
POST /class
GET /class/:id
PUT /class/:id
DELETE /class/:id
GET /class/:id/tasks

POST /task
GET /task/:id
PUT /task/:id
DELETE /task/:id
```

## Add write permission example
```sql
INSERT INTO users_permissions
SELECT 1, permissions.id
FROM permissions
WHERE permissions.code = 'task:write';
```

## DB Structure

```
// Use DBML to define your database structure
// Docs: https://dbml.dbdiagram.io/docs

Table users {
  id integer [primary key]
  first_name varchar
  last_name varchar
}

Table classroom {
  id integer [primary key]
  name varchar
  description varchar
  created_at timestamp
}

Table task {
  id integer [primary key]
  header string
  description string
  created_at timestamp
  updated_at timestamp
}

Table classroom_task {
  class_id integer [ref: > classroom.id ]
  task_id integer [ref: > task.id]
}

```