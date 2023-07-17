BEGIN;

CREATE TABLE IF NOT EXISTS flowEvent (
  id int NOT NULL AUTO_INCREMENT,
  type varchar(255) NOT NULL,
  lastBlockHeight bigint NOT NULL DEFAULT 0,
  createdAt timestamp(6) NULL DEFAULT CURRENT_TIMESTAMP(6),
  updatedAt timestamp(6) NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  deletedAt timestamp(6) NULL DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY flowEvent_uk_flowEvent_type (type)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

COMMIT;