CREATE SEQUENCE m_province_seq;

CREATE TABLE m_province (
  id int check (id > 0) NOT NULL DEFAULT NEXTVAL ('m_province_seq'),
  created_at timestamp(0) NULL DEFAULT NULL,
  updated_at timestamp(0) NULL DEFAULT NULL,
  deleted_at timestamp(0) NULL DEFAULT NULL,
  province_name varchar(255) NOT NULL,
  province_code varchar(3) NOT NULL,
  country_id int check (country_id > 0) NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_m_province FOREIGN KEY (country_id) REFERENCES m_country (id) ON DELETE RESTRICT ON UPDATE CASCADE
);

CREATE INDEX m_province_deleted_at ON m_province (deleted_at);
CREATE INDEX m_province_index_country_id ON m_province (country_id ASC);
