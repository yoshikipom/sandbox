CREATE TABLE person (
    id bigserial,
    name varchar(255),
    deleted bool NOT NULL DEFAULT false,
    PRIMARY KEY (id)

);

CREATE TABLE person_score (
    person_id bigint,
    item varchar(255),
    score bigint,
    FOREIGN KEY (person_id) REFERENCES person (id)
);
CREATE INDEX idx_person_score_person_id ON person_score(person_id);
