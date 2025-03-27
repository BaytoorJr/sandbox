CREATE TABLE IF NOT EXISTS service_mapping (
    service_type VARCHAR(255) NOT NULL,
    service_id UUID NOT NULL,
    PRIMARY KEY (service_type, service_id)
);

create table if not exists remote_payment (
    id varchar(255) not null primary key,
    type varchar(255) not null default 'qr',
    service_mapping_type VARCHAR(255) NOT NULL,
    service_mapping_id UUID NOT NULL,
    company_id UUID not null,
    manager_id UUID,
    account_number varchar(255) not null,
    amount NUMERIC(18, 2) not null,
    receipt_number varchar(255),
    payer_short_name varchar(255) not null,
    payer_phone_number varchar(255) not null,
    description varchar(255),
    status varchar(255) not null default 'NEW',
    created_at timestamp without time zone,
    CONSTRAINT fk_service_mapping FOREIGN KEY (service_mapping_type, service_mapping_id)
        REFERENCES service_mapping (service_type, service_id)
);

CREATE OR REPLACE FUNCTION validate_service_mapping()
    RETURNS TRIGGER AS $$
BEGIN
    IF NEW.service_type = 'selling_point' THEN
        IF NOT EXISTS (SELECT 1 FROM selling_point WHERE id = NEW.service_id) THEN
            RAISE EXCEPTION 'Invalid service_id: % for service_type: selling_point', NEW.service_id;
        END IF;
    ELSIF NEW.service_type = 'company_service' THEN
        IF NOT EXISTS (SELECT 1 FROM company_services WHERE id = NEW.service_id) THEN
            RAISE EXCEPTION 'Invalid service_id: % for service_type: company_service', NEW.service_id;
        END IF;
    ELSE
        RAISE EXCEPTION 'Invalid service_type: %', NEW.service_type;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER validate_service_mapping_trigger
    BEFORE INSERT OR UPDATE ON service_mapping
    FOR EACH ROW EXECUTE FUNCTION validate_service_mapping();