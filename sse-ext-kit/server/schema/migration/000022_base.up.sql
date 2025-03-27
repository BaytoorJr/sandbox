ALTER TABLE remote_payment
    ADD COLUMN IF NOT EXISTS qr_id uuid,
    ADD COLUMN IF NOT EXISTS transaction_id bigint not null default 0;