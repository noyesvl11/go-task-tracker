CREATE TABLE IF NOT EXISTS tasks (
                                     id SERIAL PRIMARY KEY,
                                     title TEXT NOT NULL,
                                     description TEXT,
                                     status INTEGER NOT NULL
);

CREATE TABLE public.users (
                              id SERIAL PRIMARY KEY,
                              username VARCHAR(255) NOT NULL,
                              password VARCHAR(255) NOT NULL,
                              role VARCHAR(50) NOT NULL DEFAULT 'student'
);
