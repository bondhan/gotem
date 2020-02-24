CREATE SEQUENCE m_user_seq;

CREATE TABLE m_user (
  id int check (id > 0) NOT NULL DEFAULT NEXTVAL ('m_user_seq'),
  created_at timestamp(0) NULL DEFAULT NULL,
  updated_at timestamp(0) NULL DEFAULT NULL,
  deleted_at timestamp(0) NULL DEFAULT NULL,
  name varchar(255) NOT NULL,
  mobile varchar(20) NOT NULL,
  email varchar(255) DEFAULT NULL,
  address varchar(255) DEFAULT NULL,
  zipcode_id int check (zipcode_id > 0) NULL,
  password varchar(255) NOT NULL,
  salt varchar(255) NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_m_user_zipcode_id FOREIGN KEY (zipcode_id) REFERENCES m_zipcode (id) ON DELETE RESTRICT ON UPDATE CASCADE
);

CREATE INDEX m_user_deleted_at ON m_user (deleted_at);