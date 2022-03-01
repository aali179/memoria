CREATE SEQUENCE maps_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."maps" (
    "id" bigint DEFAULT nextval('maps_id_seq') NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "location" text,
    "image_id" bigint,
    "page_id" bigint,
    CONSTRAINT "maps_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "idx_maps_deleted_at" ON "public"."maps" USING btree ("deleted_at");

ALTER TABLE ONLY "public"."maps" ADD CONSTRAINT "fk_images_map" FOREIGN KEY (image_id) REFERENCES images(id) NOT DEFERRABLE;
ALTER TABLE ONLY "public"."maps" ADD CONSTRAINT "fk_pages_map" FOREIGN KEY (page_id) REFERENCES pages(id) NOT DEFERRABLE;