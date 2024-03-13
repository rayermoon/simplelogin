/**
  This is the SQL script that will be used to initialize the database schema.
  We will evaluate you based on how well you design your database.
  1. How you design the tables.
  2. How you choose the data types and keys.
  3. How you name the fields.
  In this assignment we will use PostgreSQL as the database.
  */
CREATE TABLE sp_users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    full_name VARCHAR(100) NOT NULL,
    gender VARCHAR(255) NOT NULL,
    occupation VARCHAR(50),
    password VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20) unique NOT null,
    login_count int default 0
);

CREATE INDEX idx_email ON sp_users (email);
CREATE INDEX idx_phone_number ON sp_users (phone_number);
 

