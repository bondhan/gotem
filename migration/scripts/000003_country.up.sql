CREATE SEQUENCE m_country_seq;

CREATE TABLE m_country (
  id int check (id > 0) NOT NULL DEFAULT NEXTVAL ('m_country_seq'),
  created_at timestamp(0) NULL DEFAULT NULL,
  updated_at timestamp(0) NULL DEFAULT NULL,
  deleted_at timestamp(0) NULL DEFAULT NULL,
  country_name varchar(255) NOT NULL,
  country_code varchar(6) NOT NULL,
  PRIMARY KEY (id)
);

CREATE INDEX m_country_deleted_at ON m_country (deleted_at);
