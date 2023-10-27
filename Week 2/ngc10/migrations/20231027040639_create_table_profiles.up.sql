CREATE TABLE profiles (
  id SERIAL PRIMARY KEY,
  first_name VARCHAR NOT NULL,
  last_name VARCHAR NOT NULL,
  address VARCHAR NOT NULL,
  phone_number VARCHAR NOT NULL,
  user_id INT,
  FOREIGN KEY (user_id) REFERENCES users(id)
);