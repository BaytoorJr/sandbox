CREATE TABLE IF NOT EXISTS company_info
(
    "id"                UUID PRIMARY KEY            NOT NULL,
    "user_id"           UUID UNIQUE                 NOT NULL,
    "uin"               varchar(12) UNIQUE          NOT NULL,
    "short_name"        varchar(100)                NOT NULL,
    "full_name"         varchar(150)                NOT NULL,
    "kato"              varchar(30)                 NOT NULL,
    "region"            varchar(250)                NOT NULL,
    "city"              varchar(250)                NOT NULL,
    "street"            varchar(250)                NOT NULL,
    "address_string"    varchar(500)                NOT NULL,
    "address_type_code" varchar(70)                 NOT NULL,
    "address_type_name" varchar(200)                NOT NULL,
    "classifier_code"   varchar(70)                 NOT NULL,
    "classifier_name"   varchar(500)                NOT NULL,
    "register_date"     timestamp without time zone NOT NULL,
    "created_at"        timestamp without time zone NOT NULL,
    "updated_at"        timestamp without time zone NOT NULL
);

CREATE TABLE IF NOT EXISTS license
(
    "id"                    UUID PRIMARY KEY            NOT NULL,
    "company_id"            UUID UNIQUE                 NOT NULL,
    "nikad"                 varchar(30)                 NOT NULL,
    "nikad_request"         varchar(30)                 NOT NULL,
    "licensiar_name_ru"     varchar(255)                NOT NULL,
    "activity_type_code"    varchar(50)                 NOT NULL,
    "activity_type_name_ru" varchar(255)                NOT NULL,
    "validity_start_date"   timestamp without time zone,
    "validity_end_date"     timestamp without time zone,
    "issue_date"            timestamp without time zone NOT NULL,
    "created_at"            timestamp without time zone NOT NULL,
    "updated_at"            timestamp without time zone NOT NULL,

    CONSTRAINT fk_company_id FOREIGN KEY (company_id) REFERENCES company_info (id)
);

CREATE INDEX IF NOT EXISTS idx_license_company_id ON license (company_id);

CREATE INDEX IF NOT EXISTS idx_company_user_id ON company_info (user_id);
