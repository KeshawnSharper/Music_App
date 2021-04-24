CREATE TABLE base_table{
    created_at TIMESTAMP NOT NULL,
    updated at TIMESTAMP NOT NULL
}
CREATE TABLE user_account{
    id uuid PRIMARY KEY uuid_generate_v1(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL
} INHERITS (base_table)
CREATE TABLE item{
    id uuid PRIMARY KEY uuid_generate_v1(),
    title VARCHAR(255) NOT NULL,
    notes text,
    seller_id uuid,
    price_in_cents INTEGER,
    FOREIGN KEY(seller_id) REFERENCES user_account(id) ON DELETE CASCADE
} INHERITS (base_table)