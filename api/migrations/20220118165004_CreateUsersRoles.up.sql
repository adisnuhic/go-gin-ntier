CREATE TABLE users_roles (
  role_id INT unsigned NOT NULL,
  user_id  INT unsigned NOT NULL,
  FOREIGN KEY(role_id) REFERENCES roles(id) ON UPDATE CASCADE,
  FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
  PRIMARY KEY(role_id,user_id)
);