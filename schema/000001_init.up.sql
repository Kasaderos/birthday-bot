CREATE TABLE IF NOT EXISTS  users (
  id SERIAL PRIMARY KEY,
  first_name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  birthday DATE NOT NULL,
  telegram_chat_id INT NOT NULL
);

INSERT INTO users (first_name, last_name, birthday, telegram_chat_id)
VALUE ('nate', 'nate', TO_DATE('1998-08-25', 'YYYY-MM-DD'))