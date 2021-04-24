CREATE TABLE base_table{
    created_at TIMESTAMP NOT NULL,
    updated at TIMESTAMP NOT NULL
}
CREATE TABLE user_account{
    id uuid PRIMARY KEY uuid_generate_v1(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL
} INHERITS (base_table)