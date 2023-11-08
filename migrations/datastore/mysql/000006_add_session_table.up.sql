BEGIN;

CREATE TABLE IF NOT EXISTS session (
  id int NOT NULL AUTO_INCREMENT,
  sessionId varchar(64) NOT NULL,
  userId varchar(64) NOT NULL,
  lastActivity timestamp(6) NOT NULL,
  idleTimeout bigint NOT NULL,
  expTime timestamp(6) NOT NULL,
  userAgent varchar(255) NOT NULL,
  clientIp varchar(14) NOT NULL,
  createdAt timestamp(6) NULL DEFAULT CURRENT_TIMESTAMP(6),
  updatedAt timestamp(6) NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  deletedAt timestamp(6) NULL DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY session_uk_session_id (sessionId)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

COMMIT;