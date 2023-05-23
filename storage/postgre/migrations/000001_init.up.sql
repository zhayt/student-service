CREATE TABLE IF NOT EXISTS student
(
    id serial,
    name character varying NOT NULL,
    email character varying NOT NULL,
    password character varying NOT NULL,
    PRIMARY KEY (id)
);


CREATE TABLE IF NOT EXISTS gender (
    id SERIAL PRIMARY KEY,
    name varchar NOT NULL
);

CREATE TABLE IF NOT EXISTS image
(
    id serial,
    student_id integer NOT NULL,
    image_url character varying NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_student_id FOREIGN KEY (student_id)
        REFERENCES student (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
);


CREATE TABLE session
(
    id serial,
    student_id integer NOT NULL,
    token character varying NOT NULL,
    expiration_time timestamp without time zone NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_student_id FOREIGN KEY (student_id)
        REFERENCES student (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS student_personal_info
(
    id serial,
    student_id integer NOT NULL,
    gender_id integer NOT NULL,
    full_name character varying NOT NULL,
    about_student text NOT NULL,
    country character varying NOT NULL,
    region character varying NOT NULL,
    city character varying NOT NULL,
    birthday_date timestamp without time zone NOT NULL,
    phone_number character varying NOT NULL,
    CONSTRAINT fk_student_id FOREIGN KEY (student_id)
        REFERENCES student (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT fk_gender_id FOREIGN KEY (gender_id)
        REFERENCES gender (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
);


