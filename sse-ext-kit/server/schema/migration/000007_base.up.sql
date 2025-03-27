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
    END;
$$;


DO $$
    BEGIN
        -- Check if the old constraint exists
        IF EXISTS (
            SELECT 1
            FROM pg_constraint cons
                     JOIN pg_class tbl ON cons.conrelid = tbl.oid
                     JOIN pg_namespace ns ON tbl.relnamespace = ns.oid
            WHERE cons.conname = 'unique_account_number_counterparty_id'
              AND ns.nspname = 'public'  -- Replace with your schema if not 'public'
        ) THEN
            -- Drop the old constraint
            ALTER TABLE counterparty_accounts
                DROP CONSTRAINT unique_account_number_counterparty_id;
        END IF;

        -- Check if the new constraint already exists
        IF NOT EXISTS (
            SELECT 1
            FROM pg_constraint cons
                     JOIN pg_class tbl ON cons.conrelid = tbl.oid
                     JOIN pg_namespace ns ON tbl.relnamespace = ns.oid
            WHERE cons.conname = 'unique_account_number_counterparty_id_active'
              AND ns.nspname = 'public'  -- Replace with your schema if not 'public'
        ) THEN
            -- Add the new constraint
            ALTER TABLE counterparty_accounts
                ADD CONSTRAINT unique_account_number_counterparty_id_active UNIQUE (counterparty_id, number, active);
        END IF;
    END;
$$;

-- Drop the trigger if it exists
DO $$
    BEGIN
        IF EXISTS (
            SELECT 1
            FROM pg_trigger
            WHERE tgname = 'trigger_check_unique_account_globally'
        ) THEN
            EXECUTE 'DROP TRIGGER trigger_check_unique_account_globally ON counterparty_accounts';
        END IF;
    END
$$;

-- Create or replace the function
CREATE OR REPLACE FUNCTION check_unique_account_globally()
    RETURNS TRIGGER AS $$
BEGIN
    -- Проверка уникальности счета по компании
    IF EXISTS (
        SELECT 1 FROM counterparty_accounts ca
        WHERE ca.number = NEW.number
          AND ca.bank_code = NEW.bank_code
          AND ca.currency = NEW.currency
          AND ca.active = TRUE
          AND ca.counterparty_id != NEW.counterparty_id
    ) THEN
        RAISE EXCEPTION USING ERRCODE = '23505',
            MESSAGE = 'Account number ' || NEW.number || ' already exists in the system';
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create the trigger
CREATE TRIGGER trigger_check_unique_account_globally
    BEFORE INSERT OR UPDATE ON counterparty_accounts
    FOR EACH ROW EXECUTE FUNCTION check_unique_account_globally();