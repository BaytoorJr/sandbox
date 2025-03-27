DO $$
    BEGIN
        IF EXISTS (
            SELECT 1
            FROM pg_constraint
            WHERE conname = 'unique_selling_point_profile'
        ) THEN
            ALTER TABLE managers
                DROP CONSTRAINT unique_selling_point_profile;
        END IF;
    END $$;
