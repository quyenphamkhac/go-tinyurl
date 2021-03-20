CREATE KEYSPACE tinyurl
WITH REPLICATION = { 'class' : 'SimpleStrategy','replication_factor' : 1 };

USE tinyurl;

CREATE TABLE users (
  id uuid,
  username text,
  name text,
  hashed_password text,
  email text,
  creation_date timestamp,
  last_login timestamp,
  PRIMARY KEY (id)
);

CREATE TABLE urls (
  hash text,
  original_url text,
  creation_date timestamp,
  expiration_date timestamp,
  user_id text,
  PRIMARY KEY (user_id, hash)
);