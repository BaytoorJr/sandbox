CREATE TABLE IF NOT EXISTS selling_point
(
    "id"                UUID PRIMARY KEY            NOT NULL,
    "company_id"        UUID                        NOT NULL,
    "name"              varchar(150)                NOT NULL,
    "address"           varchar(250)                NOT NULL,
    "created_at"        timestamp without time zone NOT NULL,
    "updated_at"        timestamp without time zone NOT NULL
)