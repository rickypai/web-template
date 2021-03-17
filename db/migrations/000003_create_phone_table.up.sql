BEGIN;

CREATE TABLE IF NOT EXISTS phone(
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(300) UNIQUE NOT NULL,
  manufacturer_id INTEGER REFERENCES manufacturer(id) NOT NULL,
  os_id INTEGER REFERENCES os(id) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  modified_at TIMESTAMP WITH TIME ZONE NOT NULL
);

COMMIT;
