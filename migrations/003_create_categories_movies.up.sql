CREATE TABLE categories_movies (
  category_id INTEGER REFERENCES categories(id),
  movie_id INTEGER REFERENCES movies(id)
);
