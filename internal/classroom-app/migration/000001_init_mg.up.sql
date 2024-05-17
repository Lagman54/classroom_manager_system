CREATE TABLE IF NOT EXISTS classroom
(
    id          serial PRIMARY KEY,
    name        varchar NOT NULL,
    description varchar NOT NULL,
    created_at timestamp(0) with time zone DEFAULT now()
);

CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users
(
    id            serial PRIMARY KEY,
    created_at    timestamp(0) with time zone DEFAULT now(),
    first_name    varchar NOT NULL,
    last_name     varchar NOT NULL,
    email         citext  NOT NULL unique,
    password_hash bytea   NOT NULL,
    activated     bool    NOT NULL
);

CREATE TABLE IF NOT EXISTS task
(
    id          serial primary key,
    header      varchar                     NOT NULL,
    description varchar                     NOT NULL,
    created_at  timestamp(0) with time zone DEFAULT now(),
    updated_at  timestamp(0) with time zone DEFAULT now()
);

CREATE TABLE IF NOT EXISTS classroom_task
(
    class_id int references classroom (id) on delete CASCADE,
    task_id  int references task (id) on delete CASCADE,
    primary key (class_id, task_id)
);

CREATE OR REPLACE FUNCTION update_timestamp()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_task_timestamp
    BEFORE UPDATE
    ON task
    FOR EACH ROW
EXECUTE PROCEDURE
    update_timestamp();