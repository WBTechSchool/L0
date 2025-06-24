-- migrate:up

CREATE TABLE customers (
  customer_id BIGSERIAL PRIMARY KEY,
  first_name VARCHAR(64) NOT NULL,
  created_at TIMESTAMPTZ DEFAULT current_timestamp
);

INSERT INTO customers (first_name)
VALUES ('Test User');

-- migrate:down

DROP TABLE customers;