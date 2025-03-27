ALTER TABLE mundea_contracts
    ADD COLUMN IF NOT EXISTS external_id uuid not null default '00000000-0000-0000-0000-000000000000',
    ADD COLUMN IF NOT EXISTS type varchar(255);
