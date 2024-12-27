-- Drop table if it already exists
DROP TABLE IF EXISTS example_table;

-- Create a table with a variety of PostgreSQL column types
CREATE TABLE example_table (
    id SERIAL PRIMARY KEY,                -- Auto-incrementing integer
    name VARCHAR(100),                    -- Variable-length string
    description TEXT,                     -- Longer text field
    age SMALLINT,                         -- Small integer
    height NUMERIC(5, 2),                 -- Fixed precision number
    active BOOLEAN,                       -- Boolean value
    created_at TIMESTAMP DEFAULT now(),   -- Timestamp with timezone
    updated_at DATE DEFAULT CURRENT_DATE, -- Date
    preferences JSONB,                    -- JSON data
    ip_address INET,                      -- IP address
    coordinates POINT                     -- Geometric point
);

-- Insert 1000 rows of random data
INSERT INTO example_table (name, description, age, height, active, preferences, ip_address, coordinates)
SELECT
    'Name_' || i,                                      -- Random name
    'Description for entry ' || i,                    -- Random description
    (RANDOM() * 80 + 20)::SMALLINT,                   -- Random age between 20 and 100
    (RANDOM() * 50 + 150)::NUMERIC(5, 2),             -- Random height between 150.00 and 200.00
    (RANDOM() > 0.5),                                 -- Random boolean
    jsonb_build_object('likes', i % 10, 'type', 'X'), -- Random JSON data
    ('192.168.' || (RANDOM() * 100)::INT || '.' || (RANDOM() * 100)::INT)::inet, -- Random IP address, cast to inet
    POINT(RANDOM() * 100, RANDOM() * 100)             -- Random point coordinates
FROM generate_series(1, 1000) AS i;                   -- Generate 1000 rows