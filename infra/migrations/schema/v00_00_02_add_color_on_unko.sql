-- +migrate Up
ALTER TABLE
    unko
ADD
    color varchar(256) not null
;

-- +migrate Down
ALTER TABLE
    unko
DROP COLUMN
    color
;
