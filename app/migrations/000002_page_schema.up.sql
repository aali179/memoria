
CREATE SEQUENCE pages_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."pages" (
    "id" bigint DEFAULT nextval('pages_id_seq') NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "number" text,
    "heading_one" text,
    "heading_two" text,
    "heading_three" text,
    "scrapbook_id" bigint,
    CONSTRAINT "pages_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "idx_pages_deleted_at" ON "public"."pages" USING btree ("deleted_at");

ALTER TABLE ONLY "public"."pages" ADD CONSTRAINT "fk_scrapbooks_page" FOREIGN KEY (scrapbook_id) REFERENCES scrapbooks(id) NOT DEFERRABLE;
