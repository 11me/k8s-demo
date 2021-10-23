DROP SCHEMA IF EXISTS demo CASCADE;
CREATE SCHEMA demo;

CREATE TABLE demo.user
(
  id BIGSERIAL NOT NULL PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

INSERT INTO demo.user VALUES
(
  1,
  'Frodo'
);

INSERT INTO demo.user VALUES
(
  2,
  'Legolas'
);

INSERT INTO demo.user VALUES
(
  3,
  'Aragorn'
);
