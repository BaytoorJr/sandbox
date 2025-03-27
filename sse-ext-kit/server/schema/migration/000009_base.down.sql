ALTER TABLE license
ADD CONSTRAINT license_company_id_key UNIQUE (company_id);

ALTER TABLE license
DROP COLUMN IF EXISTS guid;