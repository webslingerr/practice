CREATE OR REPLACE FUNCTION add_columns_orders() RETURNS TRIGGER LANGUAGE PLPGSQL
AS
$$
DECLARE 
    product_price NUMERIC;
BEGIN
    SELECT 
        price
    INTO product_price
    FROM products;

    NEW.price = product_price;
    NEW.total_price = NEW.quantity * product_price;
    RETURN NEW;
END;
$$;

CREATE TRIGGER add_columns_orders_tg
    BEFORE INSERT OR UPDATE
ON orders
FOR EACH ROW 
    EXECUTE PROCEDURE add_columns_orders();
