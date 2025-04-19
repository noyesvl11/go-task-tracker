CREATE TABLE IF NOT EXISTS courses (
                                       id SERIAL PRIMARY KEY,
                                       title VARCHAR(255) NOT NULL,
                                       description TEXT
    );
