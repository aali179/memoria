CREATE SEQUENCE images_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."images" (
    "id" bigint DEFAULT nextval('images_id_seq') NOT NULL,
    "description" text,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "file" text,
    "deleted_at" timestamptz,
    CONSTRAINT "images_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "idx_images_deleted_at" ON "public"."images" USING btree ("deleted_at");
