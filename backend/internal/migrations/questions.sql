create table questions (
  id serial PRIMARY KEY,
  type_user varchar(10) NOT NULL,
  question varchar(2500) NOT NULL,
  email varchar(126) NOT NULL,
  add_date timestamp without time zone NOT NULL,
  tel varchar(12) DEFAULT '',
  first_name varchar(50) NOT NULL,
  second_name varchar(50) DEFAULT '',
  third_name varchar(50) DEFAULT ''
)