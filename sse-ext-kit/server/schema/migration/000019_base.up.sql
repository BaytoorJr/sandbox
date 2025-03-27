DO $$
    BEGIN
        IF NOT EXISTS (
            SELECT 1
            FROM pg_constraint
            WHERE conname = 'unique_selling_point_profile'
        ) THEN
            ALTER TABLE managers
                ADD CONSTRAINT unique_selling_point_profile UNIQUE (selling_point_id, profile_id);
        END IF;
    END $$;
