CREATE SEQUENCE m_products_seq;
CREATE TYPE category AS ENUM('shirt', 'trousers', 'pants', 'short');

CREATE TABLE m_products (
  id int check (id > 0) NOT NULL DEFAULT NEXTVAL ('m_products_seq'),
  created_at timestamp(0) NULL DEFAULT NULL,
  updated_at timestamp(0) NULL DEFAULT NULL,
  deleted_at timestamp(0) NULL DEFAULT NULL,
  name varchar(255) NOT NULL,
  category category NOT NULL,
  variant_id int check (variant_id > 0) NOT NULL,
  price DECIMAL(14, 2) DEFAULT NULL,
  "desc" text DEFAULT NULL,
  quantity int DEFAULT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_m_products_variant_id FOREIGN KEY (variant_id) REFERENCES m_variants (id) ON DELETE RESTRICT ON UPDATE CASCADE
);

CREATE INDEX m_products_deleted_at ON m_products (deleted_at);
