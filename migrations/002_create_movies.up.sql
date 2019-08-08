CREATE TABLE movies (
  id SERIAL PRIMARY KEY,
  category_id INTEGER REFERENCES categories(id),
  name CHARACTER VARYING,
  slug CHARACTER VARYING,
  url CHARACTER VARYING,
  rate FLOAT
);

CREATE INDEX index_category_id_on_movies ON movies(category_id);
