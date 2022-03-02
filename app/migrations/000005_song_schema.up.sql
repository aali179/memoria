CREATE SEQUENCE songs_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."songs" (
    "id" bigint DEFAULT nextval('songs_id_seq') NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "spotify_id" text,
    "page_id" bigint,
    CONSTRAINT "songs_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "idx_songs_deleted_at" ON "public"."songs" USING btree ("deleted_at");

ALTER TABLE ONLY "public"."songs" ADD CONSTRAINT "fk_pages_song" FOREIGN KEY (page_id) REFERENCES pages(id) NOT DEFERRABLE;
