CREATE TYPE promotion_type AS ENUM ('percentage_off', 'exchange');

-- CREATE TABLE discounts (
--   discount_id SERIAL PRIMARY KEY,
--   discount_value integer Default 0 NOT NULL,
--   type promotion_type NOT NULL
-- );


-- CREATE TABLE vips (
--   vip_no SERIAL PRIMARY KEY,
--   name varchar(10) NOT NULL,
--   discount_id integer REFERENCES discounts ON DELETE RESTRICT
-- );

CREATE TYPE vip_type AS ENUM ('Normal', 'VIP1', 'VIP2', 'VIP3');
CREATE TABLE promotions (
  p_no SERIAL PRIMARY KEY,
  event_name varchar(100) NOT NULL,
  event_content text,
  start_time timestamp,
  end_time timestamp
);

CREATE TABLE promotion_item (
  id SERIAL PRIMARY KEY,
  p_no integer REFERENCES promotions,
  promotion_type promotion_type NOT NULL,
  vip_type vip_type NOT NULL,
  value integer DEFAULT 0 NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
  uid SERIAL PRIMARY KEY,
  username varchar(30) NOT NULL,
  password varchar(64) NOT NULL,
  salt varchar(50) NOT NULL,
  coin integer NOT NULL,
  point integer NOT NULL,
  vip_type vip_type DEFAULT 'Normal' NOT NULL,
  accumulated_spent integer DEFAULT 0 NOT NULL

  CONSTRAINT positive_coin CHECK (coin >= 0)
  CONSTRAINT positive_point CHECK (point >= 0)
);

CREATE FUNCTION buy() RETURNS integer AS $$
    SELECT 1 AS result;
$$ LANGUAGE SQL;


CREATE TABLE products (
  product_no SERIAL PRIMARY KEY,
  name text,
  price numeric
);

CREATE TABLE IF NOT EXISTS orders (
  order_id SERIAL PRIMARY KEY,
  cost_coin integer NOT NULL,
  cost_point integer NOT NULL
);


CREATE TABLE order_items (
  order_id integer REFERENCES orders ON DELETE CASCADE,
  product_no integer REFERENCES products ON DELETE RESTRICT,
  quantity integer,
  PRIMARY KEY (product_no, order_id)
);



