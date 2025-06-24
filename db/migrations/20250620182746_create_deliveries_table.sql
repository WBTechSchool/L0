-- migrate:up

CREATE TABLE deliveries (
  delivery_id BIGSERIAL PRIMARY KEY,
  order_id BIGINT UNIQUE REFERENCES orders(order_id) ON DELETE NO ACTION NOT NULL,
  name VARCHAR(64) NOT NULL,
  phone VARCHAR(64) NOT NULL,
  zip VARCHAR(16) NOT NULL,
  city VARCHAR(128) NOT NULL,
  address VARCHAR(256) NOT NULL,
  region VARCHAR(256) NOT NULL,
  email VARCHAR(64) NOT NULL
);

-- migrate:down

DROP TABLE deliveries;