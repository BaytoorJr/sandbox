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


CREATE TABLE IF NOT EXISTS counterparty
(
    id uuid primary key not null,
    company_id uuid not null,
    uin varchar(12) not null,
    name varchar(255) not null,
    active bool default true,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,

    CONSTRAINT fk_company_id FOREIGN KEY (company_id) REFERENCES company_info (id)
);

CREATE TABLE IF NOT EXISTS counterparty_accounts
(
    id uuid primary key not null,
    counterparty_id uuid not null,
    type varchar(255) not null,
    number varchar(255) not null,
    bank_code varchar(255) not null,
    currency varchar(255) not null,
    active bool default true,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,

    CONSTRAINT  fk_counterparty_id FOREIGN KEY (counterparty_id) REFERENCES counterparty (id)
);

DO $$
    BEGIN
        IF NOT EXISTS (
            SELECT 1
            FROM pg_constraint cons
                     JOIN pg_class tbl ON cons.conrelid = tbl.oid
                     JOIN pg_namespace ns ON tbl.relnamespace = ns.oid
            WHERE cons.conname = 'unique_company_uin_name'
              AND ns.nspname = 'public'  -- Replace with your schema if not 'public'
        ) THEN
            -- Add the constraint
            ALTER TABLE counterparty
                ADD CONSTRAINT unique_company_uin_name UNIQUE (company_id, uin, name);
        END IF;
    END $$;


DO $$
    BEGIN
        IF NOT EXISTS (
            SELECT 1
            FROM pg_constraint cons
                     JOIN pg_class tbl ON cons.conrelid = tbl.oid
                     JOIN pg_namespace ns ON tbl.relnamespace = ns.oid
            WHERE cons.conname = 'unique_account_number_counterparty_id'
              AND ns.nspname = 'public'  -- Replace with your schema if not 'public'
        ) THEN
            ALTER TABLE counterparty_accounts
            ADD CONSTRAINT unique_account_number_counterparty_id UNIQUE(counterparty_id, number);
        end if;
END;$$