package sql

const schema = `
CREATE TABLE IF NOT EXISTS book (
    id uuid,
    name text,
    author text
);`
