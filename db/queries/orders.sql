-- name: GetOrderByID :one
SELECT
    o.order_id,
    o.customer_id,
    o.locale,
    o.internal_signature,
    o.delivery_service,
    o.track_number,
    o.entry,
    o.shard_key,
    o.off_shard,
    o.sm_id,
    o.created_at,

    p.transaction,
    p.request_id,
    p.currency,
    p.provider,
    p.amount,
    p.payment_dt,
    p.bank,
    p.delivery_cost,
    p.goods_total,
    p.custom_fee,

    d.name,
    d.phone,
    d.zip,
    d.city,
    d.address,
    d.region,
    d.email,

    (
        SELECT json_agg(json_build_object(
            'item_id', oi.item_id,
            'chrt_id', oi.chrt_id,
            'track_number', oi.track_number,
            'price', oi.price,
            'rid', oi.rid,
            'name', oi.name,
            'sale', oi.sale,
            'size', oi.size,
            'total_price', oi.total_price,
            'nm_id', oi.nm_id,
            'brand', oi.brand,
            'status', oi.status
        ))
        FROM order_items oi
        WHERE oi.order_id = o.order_id
    ) AS items
FROM orders AS o
LEFT JOIN payments AS p ON p.order_id = o.order_id
LEFT JOIN deliveries AS d ON d.order_id = o.order_id
WHERE o.order_id = $1;

-- name: GetOrders :many
SELECT
    o.order_id,
    o.customer_id,
    o.locale,
    o.internal_signature,
    o.delivery_service,
    o.track_number,
    o.entry,
    o.shard_key,
    o.off_shard,
    o.sm_id,
    o.created_at,

    p.transaction,
    p.request_id,
    p.currency,
    p.provider,
    p.amount,
    p.payment_dt,
    p.bank,
    p.delivery_cost,
    p.goods_total,
    p.custom_fee,

    d.name,
    d.phone,
    d.zip,
    d.city,
    d.address,
    d.region,
    d.email,

    (
        SELECT json_agg(json_build_object(
            'item_id', oi.item_id,
            'chrt_id', oi.chrt_id,
            'track_number', oi.track_number,
            'price', oi.price,
            'rid', oi.rid,
            'name', oi.name,
            'sale', oi.sale,
            'size', oi.size,
            'total_price', oi.total_price,
            'nm_id', oi.nm_id,
            'brand', oi.brand,
            'status', oi.status
        ))
        FROM order_items oi
        WHERE oi.order_id = o.order_id
    ) AS items
FROM orders AS o
LEFT JOIN payments AS p ON p.order_id = o.order_id
LEFT JOIN deliveries AS d ON d.order_id = o.order_id;

-- name: CreateOrder :one
INSERT INTO orders (customer_id, locale, internal_signature, delivery_service, track_number, entry, shard_key, off_shard, sm_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING order_id;

-- -- name: CreateOrder :one
-- WITH
-- new_order AS (
--   INSERT INTO orders (customer_id, locale, internal_signature, delivery_service, track_number, entry, shard_key, off_shard, sm_id)
--   VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
--   RETURNING order_id
-- ),
-- new_payment AS (
--   INSERT INTO payments (order_id, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
--   SELECT order_id, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19
--   FROM new_order
-- ),
-- new_delivery AS (
--   INSERT INTO deliveries (order_id, name, phone, zip, city, address, region, email)
--   SELECT order_id, $20, $21, $22, $23, $24, $25, $26
--   FROM new_order
-- ),
-- new_order_item AS (
--   INSERT INTO order_items (order_id, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)
--   SELECT order_id, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37
--   FROM new_order
-- )
-- SELECT order_id FROM new_order;