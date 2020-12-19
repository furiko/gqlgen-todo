
-- +migrate Up
CREATE TABLE todos(
    id varchar(15),
    text varchar(200),
    done boolean,
    user_id varchar(15),
    PRIMARY KEY (id)
);

INSERT INTO todos (id, text, done, user_id) VALUES
    ('1', 'test 1', false, '1234567'),
    ('2', 'test 2', false, '1234568'),
    ('3', 'test 3', false, '1234569'),
    ('4', 'test 4', false, '1234560'),
    ('5', 'test 5', false, '1234561');

-- +migrate Down
DROP TABLE IF EXISTS todos;
