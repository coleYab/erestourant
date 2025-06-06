CREATE TABLE "order" (
    "id" UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    "totalPrice" DOUBLE PRECISION NOT NULL,
    "userId" UUID NOT NULL REFERENCES "user" ("id") ON DELETE CASCADE,
    "status" TEXT NOT NULL,
    "createdAt" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "order_item" (
    "id" UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    "orderId" UUID NOT NULL REFERENCES "order"("id") ON DELETE CASCADE,
    "menuItemId" UUID NOT NULL REFERENCES "menu_item"("id"),
    "qty" INTEGER NOT NULL CHECK ("qty" > 0),
    "unitPrice" DOUBLE PRECISION NOT NULL,
    "totalPrice" DOUBLE PRECISION GENERATED ALWAYS AS ("unitPrice" * "qty") STORED
);

