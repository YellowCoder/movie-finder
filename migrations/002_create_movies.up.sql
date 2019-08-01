CREATE TABLE movies (
  id INTEGER PRIMARY KEY,
  category_id INTEGER REFERENCES categories(id),
  name CHARACTER VARYING NOT NULL,
  slug CHARACTER VARYING NOT NULL,
  url CHARACTER VARYING NOT NULL,
  rate FLOAT NOT NULL
);

CREATE INDEX index_category_id_on_movies ON movies(category_id);
