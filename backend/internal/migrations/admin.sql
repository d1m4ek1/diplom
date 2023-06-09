create table admin (
  id serial PRIMARY KEY,
  login varchar(24) NOT NULL,
  password text NOT NULL,
  token varchar(256) DEFAULT ''
);

create table site_content (
  id serial PRIMARY KEY,
  add_date timestamp without time zone NOT NULL,
  header text DEFAULT '',
  content text DEFAULT ''
);

create table actual_site_content (
  id serial PRIMARY KEY,
  id_site_content int NOT NULL
);

INSERT INTO admin (login, password) VALUES ('admin', 'admin_root')