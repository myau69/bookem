CREATE TABLE rooms (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL CHECK (length(btrim(name)) BETWEEN 1 AND 100),
    description TEXT,
    capacity INTEGER CHECK (capacity > 0),
    created_at TIMESTAMPTZ NOT NULL
);