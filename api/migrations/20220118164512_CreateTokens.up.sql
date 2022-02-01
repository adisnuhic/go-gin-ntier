CREATE TABLE tokens (
  id  INT NOT NULL  AUTO_INCREMENT,
  user_id INT unsigned NOT NULL,
  token VARCHAR(255) NOT NULL,
  token_type INT unsigned NOT NULL,
  meta VARCHAR(255) NULL,
  expires_at TIMESTAMP NOT NULL,
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now(),
  PRIMARY KEY(id),
  FOREIGN KEY (token_type) REFERENCES token_types(id),
  FOREIGN KEY (user_id) REFERENCES users(id)
);