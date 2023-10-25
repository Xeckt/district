-- Servers table
CREATE TABLE IF NOT EXISTS servers (
    server_id BIGINT PRIMARY KEY,
    server_name TEXT
);

-- Users table
CREATE TABLE IF NOT EXISTS users (
    user_id BIGINT PRIMARY KEY,
    username TEXT
);

-- Metrics table for aggregated data
CREATE TABLE IF NOT EXISTS metrics (
    id SERIAL PRIMARY KEY,
    server_id BIGINT REFERENCES servers(server_id),
    user_id BIGINT REFERENCES users(user_id),
    metric_date DATE,
    message_count INT DEFAULT 0,
    user_joined_count INT DEFAULT 0,
    user_left_count INT DEFAULT 0,
    UNIQUE(server_id, user_id, metric_date)
);
