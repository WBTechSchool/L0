-- migrate:up

CREATE TABLE orders (
  order_id BIGSERIAL PRIMARY KEY,
  customer_id BIGINT REFERENCES customers(customer_id) ON DELETE NO ACTION NOT NULL,
  locale VARCHAR(16) NOT NULL,
  internal_signature VARCHAR(16) NOT NULL,
  delivery_service VARCHAR(16) NOT NULL,
  track_number VARCHAR(64) NOT NULL,
  entry VARCHAR(16) NOT NULL,
  shard_key VARCHAR(16) NOT NULL,
  off_shard VARCHAR(16) NOT NULL,
  sm_id BIGINT NOT NULL,
  created_at TIMESTAMPTZ DEFAULT current_timestamp
);

-- migrate:down

DROP TABLE orders;