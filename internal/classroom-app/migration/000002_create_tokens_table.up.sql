CREATE TABLE IF NOT EXISTS tokens
(
    hash    bytea PRIMARY KEY,
    user_id int                         NOT NULL references users on delete cascade ,
    expiry  timestamp(0) with time zone NOT NULL,
    scope   text                        NOT NULL
);