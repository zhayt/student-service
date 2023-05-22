CREATE TABLE IF NOT EXISTS student(
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL UNIQUE,
    email VARCHAR NOT NULL UNIQUE,
    password VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS gender (
    id SERIAL PRIMARY KEY,
    name varchar NOT NULL
);

CREATE TABLE IF NOT EXISTS student_personal_info (
    id SERIAL PRIMARY KEY,
    student_id INTEGER NOT NULL UNIQUE,
    gender_id INTEGER NOT NULL,
    full_name VARCHAR NOT NULL,
    about_student TEXT NOT NULL,
    country VARCHAR NOT NULL,
    region VARCHAR NOT NULL,
    city VARCHAR NOT NULL,
    birthday_date TIMESTAMP NOT NULL,
    phone_number VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS image (
    id SERIAL PRIMARY KEY,
    student_id INTEGER NOT NULL UNIQUE,
    image_url varchar NOT NULL
);
