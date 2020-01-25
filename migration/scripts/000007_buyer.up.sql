CREATE SEQUENCE m_buyer_seq;

CREATE TABLE m_buyer (
  id int check (id > 0) NOT NULL DEFAULT NEXTVAL ('m_buyer_seq'),
  created_at timestamp(0) NULL DEFAULT NULL,
  updated_at timestamp(0) NULL DEFAULT NULL,
  deleted_at timestamp(0) NULL DEFAULT NULL,
  name varchar(255) NOT NULL,
  mobile varchar(255) DEFAULT NULL,
  email varchar(255) DEFAULT NULL,
  address varchar(255) DEFAULT NULL,
  zipcode_id int check (zipcode_id > 0) NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_m_buyer_zip_code FOREIGN KEY (zipcode_id) REFERENCES m_zipcode (id) ON DELETE RESTRICT ON UPDATE CASCADE
) ;

CREATE INDEX m_buyer_deleted_at ON m_buyer (deleted_at);