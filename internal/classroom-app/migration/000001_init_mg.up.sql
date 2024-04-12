CREATE TABLE IF NOT EXISTS classroom
(
    id          serial PRIMARY KEY,
    name        varchar NOT NULL,
    description varchar NOT NULL,
    created_at timestamp(0) with time zone DEFAULT now()
);

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

CREATE TABLE IF NOT EXISTS classroom_user
(
    user_id  int references users (id),
    class_id int references classroom (id),
    role_id  int NOT NULL,
    primary key (user_id, class_id)
);

CREATE TABLE IF NOT EXISTS task
(
    id          serial primary key,
    header      varchar                     NOT NULL,
    description varchar                     NOT NULL,
    created_at  timestamp(0) with time zone NOT NULL DEFAULT now(),
    updated_at  timestamp(0) with time zone NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS classroom_task
(
    class_id int references classroom (id),
    task_id  int references task (id),
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