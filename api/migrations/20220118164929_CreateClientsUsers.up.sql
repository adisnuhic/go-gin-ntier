CREATE TABLE clients_users (
  client_id INT unsigned NOT NULL,
  user_id  INT unsigned NOT NULL,
  FOREIGN KEY(client_id) REFERENCES clients(id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
  PRIMARY KEY(client_id,user_id)
);