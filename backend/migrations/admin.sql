create table admin (
  id serial PRIMARY KEY,
  name varchar(24) NOT NULL,
  password text NOT NULL,
  token varchar(256) DEFAULT ''
);

create table site_content (
  id serial PRIMARY KEY,
  add_date timestamp without time zone NOT NULL,
  content text DEFAULT ''
);

INSERT INTO admin (name, password) VALUES ('admin', 'admin_root')