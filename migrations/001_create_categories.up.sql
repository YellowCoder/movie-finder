CREATE TABLE categories (
  id INTEGER PRIMARY KEY,
  name CHARACTER VARYING NOT NULL,
  slug CHARACTER VARYING NOT NULL
);

CREATE INDEX index_slug_on_categories ON categories(slug);
