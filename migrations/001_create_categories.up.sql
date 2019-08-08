CREATE TABLE categories (
  id SERIAL PRIMARY KEY,
  name CHARACTER VARYING,
  slug CHARACTER VARYING
);

CREATE INDEX index_slug_on_categories ON categories(slug);
