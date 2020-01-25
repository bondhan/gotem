CREATE SEQUENCE m_variants_seq;
CREATE TYPE size AS ENUM('S', 'M', 'L', 'XL', 'XXL', 'LL', 'LLL');

CREATE TABLE m_variants (
  id int check (id > 0) NOT NULL DEFAULT NEXTVAL ('m_variants_seq'),
  created_at timestamp(0) NULL DEFAULT NULL,
  updated_at timestamp(0) NULL DEFAULT NULL,
  deleted_at timestamp(0) NULL DEFAULT NULL,
  size size default NULL,
  PRIMARY KEY (id)
);

CREATE INDEX m_variants_deleted_at ON m_variants(deleted_at);