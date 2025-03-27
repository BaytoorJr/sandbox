DO $$
    BEGIN
        IF NOT EXISTS (
            SELECT 1
            FROM pg_constraint
            WHERE conname = 'unique_profile_selling_point'
              AND conrelid = 'managers'::regclass
        ) THEN
            ALTER TABLE managers
                ADD CONSTRAINT unique_profile_selling_point
                    UNIQUE (profile_id, selling_point_id);
        END IF;
    END $$;