-- name: CreateOrder :one
INSERT INTO "order" ("totalPrice", "userId", "status")
VALUES ($1, $2, $3)
RETURNING *;

-- name: CreateOrderItem :exec
INSERT INTO "order_item" ("orderId", "menuItemId", "qty", "unitPrice")
VALUES ($1, $2, $3, $4);

-- name: GetAllOrders :many
SELECT * FROM "order" ORDER BY "createdAt" DESC;


-- name: GetOrderById :one
SELECT * FROM "order" WHERE "id"=$1;

-- name: GetOrderItems :many
SELECT * FROM "order_item" WHERE "orderId" = $1;

-- name: VerfiyOrder :exec
UPDATE "order" SET "status"='complete' WHERE "id"=$1;

-- name: GetOrderDetails :many
SELECT
    o."id" as order_id,
    o."totalPrice" as order_total,
    o."status",
    o."userId",
    o."createdAt",
    oi."id" as order_item_id,
    oi."menuItemId",
    oi."qty",
    oi."unitPrice",
    oi."totalPrice"
FROM "order" o
JOIN "order_item" oi ON o."id" = oi."orderId"
WHERE o."id" = $1;

-- name: GetOrdersByUserId :many
SELECT * FROM "order"
WHERE "userId" = $1
ORDER BY "createdAt" DESC;

-- name: GetUserOrdersWithItems :many
SELECT
    o."id" as order_id,
    o."status",
    o."totalPrice" as order_total,
    o."createdAt",
    oi."menuItemId",
    oi."qty",
    oi."unitPrice",
    oi."totalPrice"
FROM "order" o
JOIN "order_item" oi ON o."id" = oi."orderId"
WHERE o."userId" = $1
ORDER BY o."createdAt" DESC;

