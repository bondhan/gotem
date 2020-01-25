CREATE SEQUENCE m_seller_seq;

CREATE TABLE m_seller (
  id int check (id > 0) NOT NULL DEFAULT NEXTVAL ('m_seller_seq'),
  created_at timestamp(0) NULL DEFAULT NULL,
  updated_at timestamp(0) NULL DEFAULT NULL,
  deleted_at timestamp(0) NULL DEFAULT NULL,
  name varchar(255) NOT NULL,
  address varchar(255) DEFAULT NULL,
  zipcode_id int check (zipcode_id > 0) NULL,
  PRIMARY KEY (id) ,
  CONSTRAINT fk_m_seller_zip_code FOREIGN KEY (zipcode_id) REFERENCES m_zipcode (id) ON DELETE RESTRICT ON UPDATE CASCADE
);

CREATE INDEX m_seller_deleted_at ON m_seller (deleted_at);
