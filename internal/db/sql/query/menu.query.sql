-- name: CreateMenu :one
INSERT INTO "menu_item" ("name", "description", "price", "qty")
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetMenuById :one
SELECT * FROM "menu_item" WHERE "id" = $1;

-- name: DeleteMenuById :exec
DELETE FROM "menu_item" WHERE "id" = $1;


-- name: UpdateMenuItemQty :exec
UPDATE "menu_item" SET "qty"=$2 WHERE "id"=$1;

-- name: GetAllMenu :many
SELECT * FROM "menu_item";

-- name: UpdateMenuById :one
UPDATE "menu_item"
SET "name" = $2,
    "description" = $3,
    "price" = $4,
    "qty" = $5
WHERE "id" = $1
RETURNING *;

-- name: UpdateMenuQtyById :one
UPDATE "menu_item"
SET "qty" = $2
WHERE "id" = $1
RETURNING *;

-- name: GetMenuByName :many
SELECT * FROM "menu_item"
WHERE "name" ILIKE '%' || $1 || '%';

-- name: GetAvailableMenuItems :many
SELECT * FROM "menu_item"
WHERE "qty" IS NULL OR "qty" > 0;

-- name: ReduceQtyById :exec
UPDATE "menu_item"
SET "qty" = "qty" - $2
WHERE "id" = $1 AND ("qty" IS NULL OR "qty" >= $2);

