-- migrate:up

CREATE TABLE payments (
  payment_id BIGSERIAL PRIMARY KEY,
  order_id BIGINT UNIQUE REFERENCES orders(order_id) ON DELETE NO ACTION NOT NULL,
  transaction VARCHAR(255) NOT NULL,
  request_id VARCHAR(16) NOT NULL,
  currency VARCHAR(12) NOT NULL,
  provider VARCHAR(12) NOT NULL,
  amount BIGINT NOT NULL,
  payment_dt BIGINT NOT NULL,
  bank VARCHAR(12) NOT NULL,
  delivery_cost BIGINT NOT NULL,
  goods_total BIGINT NOT NULL,
  custom_fee BIGINT NOT NULL
)

-- migrate:down

DROP TABLE payments;