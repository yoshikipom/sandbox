\echo :target_id
UPDATE person SET deleted = true WHERE id = :target_id;
