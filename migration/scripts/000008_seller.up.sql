CREATE SEQUENCE m_seller_seq;

CREATE TABLE m_seller (
  id int check (id > 0) NOT NULL DEFAULT NEXTVAL ('m_seller_seq'),
  created_at timestamp(0) NULL DEFAULT NULL,
  updated_at timestamp(0) NULL DEFAULT NULL,
  deleted_at timestamp(0) NULL DEFAULT NULL,
  user_id int check (user_id > 0) NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_m_seller_user_id FOREIGN KEY (user_id) REFERENCES m_user (id) ON DELETE RESTRICT ON UPDATE CASCADE
) ;

CREATE INDEX m_seller_deleted_at ON m_seller (deleted_at);