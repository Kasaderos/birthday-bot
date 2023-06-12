CREATE TABLE IF NOT EXISTS  users (
  id SERIAL PRIMARY KEY,
  first_name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  birthday DATE NOT NULL,
  telegram_chat_id INT NOT NULL
);