CREATE TABLE token_types (
  id INT unsigned NOT NULL  AUTO_INCREMENT,
  type VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now(),
  PRIMARY KEY(id)
);