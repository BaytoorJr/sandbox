ALTER TABLE mundea_contracts
    ADD COLUMN IF NOT EXISTS pay_type  int not null default 0,
    DROP COLUMN IF EXISTS description,
    DROP COLUMN IF EXISTS knp,
    DROP COLUMN IF EXISTS kbe,
    DROP COLUMN IF EXISTS oked;
