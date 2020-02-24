CREATE SEQUENCE m_zipcode_seq;

CREATE TABLE m_zipcode (
  id int check (id > 0) NOT NULL DEFAULT NEXTVAL ('m_zipcode_seq'),
  created_at timestamp(0) NULL DEFAULT NULL,
  updated_at timestamp(0) NULL DEFAULT NULL,
  deleted_at timestamp(0) NULL DEFAULT NULL,
  zipcode varchar(6) NOT NULL,
  city_id int check (city_id > 0) NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_m_zipcode_city_id FOREIGN KEY (city_id) REFERENCES m_city (id) ON DELETE RESTRICT ON UPDATE CASCADE,
  UNIQUE(zipcode,city_id)
);

CREATE INDEX m_zipcode_deleted_at ON m_zipcode (deleted_at);
