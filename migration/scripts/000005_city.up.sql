CREATE SEQUENCE m_city_seq;

CREATE TABLE m_city (
  id int check (id > 0) NOT NULL DEFAULT NEXTVAL ('m_city_seq'),
  created_at timestamp(0) NULL DEFAULT NULL,
  updated_at timestamp(0) NULL DEFAULT NULL,
  deleted_at timestamp(0) NULL DEFAULT NULL,
  city_name varchar(255) NOT NULL,
  city_code varchar(6) NOT NULL,
  province_id int check (province_id > 0) NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_m_city_province_id FOREIGN KEY (province_id) REFERENCES m_province (id) ON DELETE RESTRICT ON UPDATE CASCADE,
  UNIQUE(city_code, province_id)
);

CREATE INDEX m_city_deleted_at ON m_city (deleted_at);
CREATE INDEX m_city_index_province_id ON m_city (province_id ASC);
