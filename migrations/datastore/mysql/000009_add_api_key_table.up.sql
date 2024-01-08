BEGIN;

CREATE TABLE IF NOT EXISTS apiKey (
  id int NOT NULL AUTO_INCREMENT,
  objectId int NOT NULL,
  displayName varchar(255) NOT NULL,
  apiKey varchar(255) NOT NULL,
  expDate timestamp(6) NULL DEFAULT NULL,
  createdAt timestamp(6) NULL DEFAULT CURRENT_TIMESTAMP(6),
  updatedAt timestamp(6) NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  deletedAt timestamp(6) NULL DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY apiKey_uk_apiKey (apikey),
  KEY objectId (objectId),
  CONSTRAINT apiKey_fk_object_id FOREIGN KEY (objectId) REFERENCES object (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

COMMIT;