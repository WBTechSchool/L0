-- migrate:up

CREATE TABLE order_items (
  item_id BIGSERIAL PRIMARY KEY,
  order_id BIGINT REFERENCES orders(order_id) ON DELETE NO ACTION NOT NULL,
  chrt_id BIGINT NOT NULL,
  track_number VARCHAR(64) NOT NULL,
  price BIGINT NOT NULL,
  rid VARCHAR(64) NOT NULL,
  name VARCHAR(64) NOT NULL,
  sale BIGINT NOT NULL,
  size VARCHAR(12) NOT NULL,
  total_price BIGINT NOT NULL,
  nm_id BIGINT NOT NULL,
  brand VARCHAR(64) NOT NULL,
  status INT NOT NULL
)

-- migrate:down

DROP TABLE order_items;