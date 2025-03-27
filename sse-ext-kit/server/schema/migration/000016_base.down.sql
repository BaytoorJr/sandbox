ALTER TABLE mundea_contracts
    ADD COLUMN IF NOT EXISTS kbe varchar(255) not null default '';