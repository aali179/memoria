CREATE TABLE "public"."page_images" (
    "page_id" bigint NOT NULL,
    "image_id" bigint NOT NULL,
    CONSTRAINT "page_images_pkey" PRIMARY KEY ("page_id", "image_id")
) WITH (oids = false);

ALTER TABLE ONLY "public"."page_images" ADD CONSTRAINT "fk_page_images_image" FOREIGN KEY (image_id) REFERENCES images(id) NOT DEFERRABLE;
ALTER TABLE ONLY "public"."page_images" ADD CONSTRAINT "fk_page_images_page" FOREIGN KEY (page_id) REFERENCES pages(id) NOT DEFERRABLE;