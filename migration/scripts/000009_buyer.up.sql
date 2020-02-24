CREATE SEQUENCE m_buyer_seq;

CREATE TABLE m_buyer (
  id int check (id > 0) NOT NULL DEFAULT NEXTVAL ('m_buyer_seq'),
  created_at timestamp(0) NULL DEFAULT NULL,
  updated_at timestamp(0) NULL DEFAULT NULL,
  deleted_at timestamp(0) NULL DEFAULT NULL,
  user_id int check (user_id > 0) NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_m_buyer_user_id FOREIGN KEY (user_id) REFERENCES m_user (id) ON DELETE RESTRICT ON UPDATE CASCADE
) ;

CREATE INDEX m_buyer_deleted_at ON m_buyer (deleted_at);