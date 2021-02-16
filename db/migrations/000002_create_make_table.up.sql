BEGIN;

CREATE TABLE IF NOT EXISTS make(
  id SERIAL PRIMARY KEY,
  name VARCHAR(300) UNIQUE NOT NULL,
  created_at TIME WITH TIME ZONE NOT NULL,
  modified_at TIME WITH TIME ZONE NOT NULL
);

COMMIT;