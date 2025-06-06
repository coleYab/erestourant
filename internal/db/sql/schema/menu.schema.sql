CREATE TABLE "menu_item" (
    "id" UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    "name" VARCHAR(255) NOT NULL,
    "description" TEXT NOT NULL,
    "price" DOUBLE NOT NULL,
    "qty" INTEGER NOT NULL
);

