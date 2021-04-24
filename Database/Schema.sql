CREATE TABLE base_table{
    id uuid PRIMARY KEY uuid_generate_v1(),
    created_at TIMESTAMP NOT NULL,
    updated at TIMESTAMP NOT NULL
}
CREATE TABLE user_account{
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL
} INHERITS (base_table)
CREATE TABLE item{
    title VARCHAR(255) NOT NULL,
    notes text,
    seller_id uuid,
    price_in_cents INTEGER,
    FOREIGN KEY(seller_id) REFERENCES user_account(id) ON DELETE CASCADE
} INHERITS (base_table)
CREATE TABLE purchase{
    buyer_id uuid,
    item_id uuid,
    seller_id uuid,
    price_in_cents INTEGER,
    title VARCHAR(255) NOT NULL,
    FOREIGN KEY(buyer_id) REFERENCES user_account(id),FOREIGN KEY(item_id) REFERENCES item(id) ON DELETE CASCADE
} INHERITS (base_table)