-- name: CreateDelivery :exec
INSERT INTO deliveries (order_id, name, phone, zip, city, address, region, email)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);