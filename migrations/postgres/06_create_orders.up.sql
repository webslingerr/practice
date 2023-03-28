CREATE TABLE orders (
    id VARCHAR PRIMARY KEY,
    name VARCHAR NOT NULL,
    price NUMERIC DEFAULT 0,
    quantity INT NOT NULL,
    total_price NUMERIC DEFAULT 0,
    user_id VARCHAR REFERENCES users(id),
    customer_id VARCHAR REFERENCES customers(id),
    courier_id VARCHAR REFERENCES couriers(id),
    product_id VARCHAR REFERENCES products(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);