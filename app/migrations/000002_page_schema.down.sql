DROP TABLE IF EXISTS "page_images";
CREATE TABLE "public"."page_images" (
    "page_id" bigint NOT NULL,
    "image_id" bigint NOT NULL,
    CONSTRAINT "page_images_pkey" PRIMARY KEY ("page_id", "image_id")
) WITH (oids = false);
