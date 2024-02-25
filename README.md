# Classrom management sytem


```
Author: Aibar Kudiyarkhan 
ID: 22B030553

Application to create classrooms with teachers and students
```

## Classroom REST API
```
POST /class
GET /class/:id
PUT /class/:id
DELTE /class/:id
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
}

Table classroom_user {
  userId integer [ref: > users.id]
  classId integer [ref: > classroom.id]
  roleId integer
}

Table task {
  id integer [primary key]
  header string
  description string
  created_at timestamp
}

Table classroom_task {
  class_id integer [ref: > classroom.id ]
  task_id integer [ref: > task.id]
}

```