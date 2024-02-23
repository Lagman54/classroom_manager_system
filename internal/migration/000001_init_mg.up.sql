CREATE TABLE IF NOT EXISTS classroom
(
    id serial PRIMARY KEY,
    name varchar NOT NULL
);

CREATE TABLE IF NOT EXISTS users
(
    id serial PRIMARY KEY,
    first_name varchar NOT NULL,
    last_name varchar NOT NULL
);

CREATE TABLE IF NOT EXISTS classroom_user(
    user_id int,
    class_id int,
    role_id int,
    primary key (user_id, class_id)
);

CREATE TABLE IF NOT EXISTS task(
    id serial primary key,
    header varchar,
    description varchar,
    created_at timestamp
);

CREATE TABLE IF NOT EXISTS classroom_task(
    class_id int,
    task_id int,
    primary key (class_id, task_id)
);