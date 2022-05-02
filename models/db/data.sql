--
-- PostgreSQL database cluster dump
--

-- Started on 2022-04-30 23:30:54

SET default_transaction_read_only = off;

SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;

--
-- Roles
--

CREATE ROLE user01;
ALTER ROLE user01 WITH NOSUPERUSER INHERIT NOCREATEROLE NOCREATEDB LOGIN NOREPLICATION NOBYPASSRLS PASSWORD 'SCRAM-SHA-256$4096:6zoLhI9FR2DxhtEx0+bgEQ==$3r8Ti2CUoVYLMOD0Sx30yaefYXB3kcUVn2X1x74PBGM=:K9pW3L/PWK3I++HJsPs7yVa199mxaOlHUESnvSTC7po=';






--
-- Databases
--

--
-- Database "template1" dump
--

\connect template1

--
-- PostgreSQL database dump
--

-- Dumped from database version 14.2 (Debian 14.2-1.pgdg110+1)
-- Dumped by pg_dump version 14.1

-- Started on 2022-05-02 09:50:35

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

-- Completed on 2022-05-02 09:50:35

--
-- PostgreSQL database dump complete
--

--
-- Database "postgres" dump
--

\connect postgres

--
-- PostgreSQL database dump
--

-- Dumped from database version 14.2 (Debian 14.2-1.pgdg110+1)
-- Dumped by pg_dump version 14.1

-- Started on 2022-05-02 09:50:35

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

-- Completed on 2022-05-02 09:50:35

--
-- PostgreSQL database dump complete
--

--
-- Database "practice-sales-backend" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 14.2 (Debian 14.2-1.pgdg110+1)
-- Dumped by pg_dump version 14.1

-- Started on 2022-05-02 09:50:36

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 3381 (class 1262 OID 16385)
-- Name: practice-sales-backend; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE "practice-sales-backend" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.utf8';


ALTER DATABASE "practice-sales-backend" OWNER TO postgres;

\connect -reuse-previous=on "dbname='practice-sales-backend'"

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 832 (class 1247 OID 16387)
-- Name: promotion_type; Type: TYPE; Schema: public; Owner: user01
--

CREATE TYPE public.promotion_type AS ENUM (
    'percentage_off',
    'exchange'
);


ALTER TYPE public.promotion_type OWNER TO user01;

--
-- TOC entry 835 (class 1247 OID 16392)
-- Name: vip_type; Type: TYPE; Schema: public; Owner: user01
--

CREATE TYPE public.vip_type AS ENUM (
    'Normal',
    'VIP1',
    'VIP2',
    'VIP3'
);


ALTER TYPE public.vip_type OWNER TO user01;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 217 (class 1259 OID 16486)
-- Name: order_items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.order_items (
    order_id integer NOT NULL,
    product_no integer NOT NULL,
    quantity integer
);


ALTER TABLE public.order_items OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 16474)
-- Name: orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.orders (
    order_id integer NOT NULL,
    uid integer,
    cost_coin integer NOT NULL,
    cost_point integer NOT NULL,
    "time" timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.orders OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 16473)
-- Name: orders_order_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.orders_order_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.orders_order_id_seq OWNER TO postgres;

--
-- TOC entry 3383 (class 0 OID 0)
-- Dependencies: 215
-- Name: orders_order_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.orders_order_id_seq OWNED BY public.orders.order_id;


--
-- TOC entry 209 (class 1259 OID 16409)
-- Name: products; Type: TABLE; Schema: public; Owner: user01
--

CREATE TABLE public.products (
    product_no integer NOT NULL,
    name text,
    price numeric
);


ALTER TABLE public.products OWNER TO user01;

--
-- TOC entry 210 (class 1259 OID 16414)
-- Name: products_product_no_seq; Type: SEQUENCE; Schema: public; Owner: user01
--

CREATE SEQUENCE public.products_product_no_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.products_product_no_seq OWNER TO user01;

--
-- TOC entry 3384 (class 0 OID 0)
-- Dependencies: 210
-- Name: products_product_no_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: user01
--

ALTER SEQUENCE public.products_product_no_seq OWNED BY public.products.product_no;


--
-- TOC entry 219 (class 1259 OID 16528)
-- Name: promotion_items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.promotion_items (
    id integer NOT NULL,
    p_no integer,
    promotion_type public.promotion_type NOT NULL,
    vip_type public.vip_type NOT NULL,
    value integer DEFAULT 0 NOT NULL
);


ALTER TABLE public.promotion_items OWNER TO postgres;

--
-- TOC entry 218 (class 1259 OID 16527)
-- Name: promotion_items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.promotion_items_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.promotion_items_id_seq OWNER TO postgres;

--
-- TOC entry 3385 (class 0 OID 0)
-- Dependencies: 218
-- Name: promotion_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.promotion_items_id_seq OWNED BY public.promotion_items.id;


--
-- TOC entry 211 (class 1259 OID 16420)
-- Name: promotions; Type: TABLE; Schema: public; Owner: user01
--

CREATE TABLE public.promotions (
    p_no integer NOT NULL,
    event_name character varying(100) NOT NULL,
    event_content text,
    start_time timestamp without time zone,
    end_time timestamp without time zone
);


ALTER TABLE public.promotions OWNER TO user01;

--
-- TOC entry 212 (class 1259 OID 16425)
-- Name: promotions_p_no_seq; Type: SEQUENCE; Schema: public; Owner: user01
--

CREATE SEQUENCE public.promotions_p_no_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.promotions_p_no_seq OWNER TO user01;

--
-- TOC entry 3386 (class 0 OID 0)
-- Dependencies: 212
-- Name: promotions_p_no_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: user01
--

ALTER SEQUENCE public.promotions_p_no_seq OWNED BY public.promotions.p_no;


--
-- TOC entry 213 (class 1259 OID 16426)
-- Name: users; Type: TABLE; Schema: public; Owner: user01
--

CREATE TABLE public.users (
    uid integer NOT NULL,
    username character varying(30) NOT NULL,
    password character varying(64) NOT NULL,
    coin integer DEFAULT 0 NOT NULL,
    point integer DEFAULT 0 NOT NULL,
    vip_type public.vip_type DEFAULT 'Normal'::public.vip_type NOT NULL,
    accumulated_spent integer DEFAULT 0 NOT NULL,
    CONSTRAINT positive_coin CHECK ((coin >= 0)),
    CONSTRAINT positive_point CHECK ((point >= 0))
);


ALTER TABLE public.users OWNER TO user01;

--
-- TOC entry 214 (class 1259 OID 16435)
-- Name: users_uid_seq; Type: SEQUENCE; Schema: public; Owner: user01
--

CREATE SEQUENCE public.users_uid_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_uid_seq OWNER TO user01;

--
-- TOC entry 3387 (class 0 OID 0)
-- Dependencies: 214
-- Name: users_uid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: user01
--

ALTER SEQUENCE public.users_uid_seq OWNED BY public.users.uid;


--
-- TOC entry 3206 (class 2604 OID 16477)
-- Name: orders order_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders ALTER COLUMN order_id SET DEFAULT nextval('public.orders_order_id_seq'::regclass);


--
-- TOC entry 3197 (class 2604 OID 16437)
-- Name: products product_no; Type: DEFAULT; Schema: public; Owner: user01
--

ALTER TABLE ONLY public.products ALTER COLUMN product_no SET DEFAULT nextval('public.products_product_no_seq'::regclass);


--
-- TOC entry 3208 (class 2604 OID 16531)
-- Name: promotion_items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.promotion_items ALTER COLUMN id SET DEFAULT nextval('public.promotion_items_id_seq'::regclass);


--
-- TOC entry 3198 (class 2604 OID 16439)
-- Name: promotions p_no; Type: DEFAULT; Schema: public; Owner: user01
--

ALTER TABLE ONLY public.promotions ALTER COLUMN p_no SET DEFAULT nextval('public.promotions_p_no_seq'::regclass);


--
-- TOC entry 3203 (class 2604 OID 16440)
-- Name: users uid; Type: DEFAULT; Schema: public; Owner: user01
--

ALTER TABLE ONLY public.users ALTER COLUMN uid SET DEFAULT nextval('public.users_uid_seq'::regclass);


--
-- TOC entry 3373 (class 0 OID 16486)
-- Dependencies: 217
-- Data for Name: order_items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.order_items (order_id, product_no, quantity) FROM stdin;
1	2	5
1	3	1
2	3	2
\.


--
-- TOC entry 3372 (class 0 OID 16474)
-- Dependencies: 216
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.orders (order_id, uid, cost_coin, cost_point, "time") FROM stdin;
1	1	115	10	2022-04-30 16:44:38
2	2	110	38	2022-05-01 00:49:57
\.


--
-- TOC entry 3365 (class 0 OID 16409)
-- Dependencies: 209
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: user01
--

COPY public.products (product_no, name, price) FROM stdin;
1	價格1元	1
2	價格10元	10
3	價格100元	100
4	價格1000元	1000
\.


--
-- TOC entry 3375 (class 0 OID 16528)
-- Dependencies: 219
-- Data for Name: promotion_items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.promotion_items (id, p_no, promotion_type, vip_type, value) FROM stdin;
1	1	percentage_off	Normal	0
2	1	percentage_off	VIP1	5
3	1	percentage_off	VIP2	10
4	1	percentage_off	VIP3	15
5	1	exchange	Normal	100
6	1	exchange	VIP1	100
7	1	exchange	VIP2	100
8	1	exchange	VIP3	100
9	2	percentage_off	Normal	10
10	2	exchange	Normal	200
12	2	exchange	VIP1	200
11	2	percentage_off	VIP1	12
13	2	percentage_off	VIP2	15
14	2	exchange	VIP2	200
15	2	percentage_off	VIP3	18
16	2	exchange	VIP3	200
\.


--
-- TOC entry 3367 (class 0 OID 16420)
-- Dependencies: 211
-- Data for Name: promotions; Type: TABLE DATA; Schema: public; Owner: user01
--

COPY public.promotions (p_no, event_name, event_content, start_time, end_time) FROM stdin;
1	預設折扣	\N	\N	\N
2	優惠測試折扣\n	新的	2022-04-29 13:00:27.409	2022-05-15 13:00:27.409
\.


--
-- TOC entry 3369 (class 0 OID 16426)
-- Dependencies: 213
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: user01
--

COPY public.users (uid, username, password, coin, point, vip_type, accumulated_spent) FROM stdin;
1	user01	$2a$10$1pZm5bEHsfz.7GhJfsCJQ.J0/1o1.QcaJYRxTL0CnZj96/4KU0vuK	885	490	Normal	115
2	user02	$2a$10$janPPGBTdMa6tz17TVG0IeTP7vPXWooHOx.H6rqu49rXm6KYMmJIq	1400	962	VIP1	100
\.


--
-- TOC entry 3388 (class 0 OID 0)
-- Dependencies: 215
-- Name: orders_order_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.orders_order_id_seq', 2, true);


--
-- TOC entry 3389 (class 0 OID 0)
-- Dependencies: 210
-- Name: products_product_no_seq; Type: SEQUENCE SET; Schema: public; Owner: user01
--

SELECT pg_catalog.setval('public.products_product_no_seq', 4, true);


--
-- TOC entry 3390 (class 0 OID 0)
-- Dependencies: 218
-- Name: promotion_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.promotion_items_id_seq', 1, false);


--
-- TOC entry 3391 (class 0 OID 0)
-- Dependencies: 212
-- Name: promotions_p_no_seq; Type: SEQUENCE SET; Schema: public; Owner: user01
--

SELECT pg_catalog.setval('public.promotions_p_no_seq', 1, true);


--
-- TOC entry 3392 (class 0 OID 0)
-- Dependencies: 214
-- Name: users_uid_seq; Type: SEQUENCE SET; Schema: public; Owner: user01
--

SELECT pg_catalog.setval('public.users_uid_seq', 2, true);


--
-- TOC entry 3219 (class 2606 OID 16490)
-- Name: order_items order_items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_items
    ADD CONSTRAINT order_items_pkey PRIMARY KEY (product_no, order_id);


--
-- TOC entry 3217 (class 2606 OID 16480)
-- Name: orders orders_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (order_id);


--
-- TOC entry 3211 (class 2606 OID 16446)
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: user01
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (product_no);


--
-- TOC entry 3221 (class 2606 OID 16534)
-- Name: promotion_items promotion_items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.promotion_items
    ADD CONSTRAINT promotion_items_pkey PRIMARY KEY (id);


--
-- TOC entry 3213 (class 2606 OID 16450)
-- Name: promotions promotions_pkey; Type: CONSTRAINT; Schema: public; Owner: user01
--

ALTER TABLE ONLY public.promotions
    ADD CONSTRAINT promotions_pkey PRIMARY KEY (p_no);


--
-- TOC entry 3215 (class 2606 OID 16452)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: user01
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (uid);


--
-- TOC entry 3223 (class 2606 OID 16491)
-- Name: order_items order_items_order_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_items
    ADD CONSTRAINT order_items_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.orders(order_id) ON DELETE CASCADE;


--
-- TOC entry 3224 (class 2606 OID 16496)
-- Name: order_items order_items_product_no_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_items
    ADD CONSTRAINT order_items_product_no_fkey FOREIGN KEY (product_no) REFERENCES public.products(product_no) ON DELETE RESTRICT;


--
-- TOC entry 3222 (class 2606 OID 16481)
-- Name: orders orders_uid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_uid_fkey FOREIGN KEY (uid) REFERENCES public.users(uid);


--
-- TOC entry 3225 (class 2606 OID 16535)
-- Name: promotion_items promotion_items_p_no_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.promotion_items
    ADD CONSTRAINT promotion_items_p_no_fkey FOREIGN KEY (p_no) REFERENCES public.promotions(p_no);


--
-- TOC entry 3382 (class 0 OID 0)
-- Dependencies: 3381
-- Name: DATABASE "practice-sales-backend"; Type: ACL; Schema: -; Owner: postgres
--

GRANT TEMPORARY ON DATABASE "practice-sales-backend" TO user01 WITH GRANT OPTION;


-- Completed on 2022-05-02 09:50:36

--
-- PostgreSQL database dump complete
--

-- Completed on 2022-05-02 09:50:36

--
-- PostgreSQL database cluster dump complete
--

