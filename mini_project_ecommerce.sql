--
-- PostgreSQL database dump
--

-- Dumped from database version 16.5 (Ubuntu 16.5-1.pgdg24.04+1)
-- Dumped by pg_dump version 17.1 (Ubuntu 17.1-1.pgdg24.04+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: addresses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.addresses (
    id integer NOT NULL,
    user_id integer NOT NULL,
    address text NOT NULL,
    is_main boolean DEFAULT false,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp without time zone
);


ALTER TABLE public.addresses OWNER TO postgres;

--
-- Name: addresses_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.addresses_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.addresses_id_seq OWNER TO postgres;

--
-- Name: addresses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.addresses_id_seq OWNED BY public.addresses.id;


--
-- Name: categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.categories (
    id integer NOT NULL,
    name character varying,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


ALTER TABLE public.categories OWNER TO postgres;

--
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.categories_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.categories_id_seq OWNER TO postgres;

--
-- Name: categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;


--
-- Name: checkout_items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.checkout_items (
    id integer NOT NULL,
    checkout_id integer,
    product_variant_id integer,
    qty integer,
    total integer,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


ALTER TABLE public.checkout_items OWNER TO postgres;

--
-- Name: checkout_items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.checkout_items_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.checkout_items_id_seq OWNER TO postgres;

--
-- Name: checkout_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.checkout_items_id_seq OWNED BY public.checkout_items.id;


--
-- Name: checkouts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.checkouts (
    id integer NOT NULL,
    user_id integer,
    total_amount integer,
    payment character varying,
    payment_status character varying,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


ALTER TABLE public.checkouts OWNER TO postgres;

--
-- Name: checkouts_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.checkouts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.checkouts_id_seq OWNER TO postgres;

--
-- Name: checkouts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.checkouts_id_seq OWNED BY public.checkouts.id;


--
-- Name: images; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.images (
    id integer NOT NULL,
    product_id integer,
    image_url character varying,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


ALTER TABLE public.images OWNER TO postgres;

--
-- Name: images_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.images_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.images_id_seq OWNER TO postgres;

--
-- Name: images_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.images_id_seq OWNED BY public.images.id;


--
-- Name: previews; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.previews (
    id integer NOT NULL,
    checkout_item_id integer,
    rating double precision,
    comentars text,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


ALTER TABLE public.previews OWNER TO postgres;

--
-- Name: previews_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.previews_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.previews_id_seq OWNER TO postgres;

--
-- Name: previews_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.previews_id_seq OWNED BY public.previews.id;


--
-- Name: product_varians; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.product_varians (
    id integer NOT NULL,
    product_id integer,
    size character varying,
    color character varying,
    stock integer,
    image_url character varying,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


ALTER TABLE public.product_varians OWNER TO postgres;

--
-- Name: product_varians_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.product_varians_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.product_varians_id_seq OWNER TO postgres;

--
-- Name: product_varians_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.product_varians_id_seq OWNED BY public.product_varians.id;


--
-- Name: products; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.products (
    id integer NOT NULL,
    name character varying,
    image_url character varying,
    price integer,
    discount integer,
    category_id integer,
    description text,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


ALTER TABLE public.products OWNER TO postgres;

--
-- Name: products_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.products_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.products_id_seq OWNER TO postgres;

--
-- Name: products_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.products_id_seq OWNED BY public.products.id;


--
-- Name: promotions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.promotions (
    id integer NOT NULL,
    title character varying NOT NULL,
    subtitle character varying,
    image_url character varying,
    path_url character varying,
    start_date date NOT NULL,
    end_date date NOT NULL,
    is_promo boolean DEFAULT false NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp without time zone
);


ALTER TABLE public.promotions OWNER TO postgres;

--
-- Name: promotions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.promotions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.promotions_id_seq OWNER TO postgres;

--
-- Name: promotions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.promotions_id_seq OWNED BY public.promotions.id;


--
-- Name: recomendeds; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.recomendeds (
    id integer NOT NULL,
    title character varying(255) NOT NULL,
    image_url character varying,
    subtitle character varying(255),
    product_id integer NOT NULL,
    status boolean DEFAULT false,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp without time zone
);


ALTER TABLE public.recomendeds OWNER TO postgres;

--
-- Name: recomendeds_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.recomendeds_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.recomendeds_id_seq OWNER TO postgres;

--
-- Name: recomendeds_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.recomendeds_id_seq OWNED BY public.recomendeds.id;


--
-- Name: shopping_carts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.shopping_carts (
    id integer NOT NULL,
    product_variant_id integer,
    user_id integer,
    qty integer DEFAULT 1,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp without time zone
);


ALTER TABLE public.shopping_carts OWNER TO postgres;

--
-- Name: shopping_carts_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.shopping_carts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.shopping_carts_id_seq OWNER TO postgres;

--
-- Name: shopping_carts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.shopping_carts_id_seq OWNED BY public.shopping_carts.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying,
    username character varying(50),
    phone character varying,
    email character varying,
    password character varying,
    token character varying,
    expired timestamp without time zone,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: wishlists; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wishlists (
    id integer NOT NULL,
    product_variant_id integer,
    user_id integer,
    created_at timestamp without time zone,
    deleted_at timestamp without time zone
);


ALTER TABLE public.wishlists OWNER TO postgres;

--
-- Name: wishlists_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wishlists_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.wishlists_id_seq OWNER TO postgres;

--
-- Name: wishlists_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wishlists_id_seq OWNED BY public.wishlists.id;


--
-- Name: addresses id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.addresses ALTER COLUMN id SET DEFAULT nextval('public.addresses_id_seq'::regclass);


--
-- Name: categories id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);


--
-- Name: checkout_items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checkout_items ALTER COLUMN id SET DEFAULT nextval('public.checkout_items_id_seq'::regclass);


--
-- Name: checkouts id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checkouts ALTER COLUMN id SET DEFAULT nextval('public.checkouts_id_seq'::regclass);


--
-- Name: images id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.images ALTER COLUMN id SET DEFAULT nextval('public.images_id_seq'::regclass);


--
-- Name: previews id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.previews ALTER COLUMN id SET DEFAULT nextval('public.previews_id_seq'::regclass);


--
-- Name: product_varians id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_varians ALTER COLUMN id SET DEFAULT nextval('public.product_varians_id_seq'::regclass);


--
-- Name: products id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products ALTER COLUMN id SET DEFAULT nextval('public.products_id_seq'::regclass);


--
-- Name: promotions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.promotions ALTER COLUMN id SET DEFAULT nextval('public.promotions_id_seq'::regclass);


--
-- Name: recomendeds id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.recomendeds ALTER COLUMN id SET DEFAULT nextval('public.recomendeds_id_seq'::regclass);


--
-- Name: shopping_carts id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shopping_carts ALTER COLUMN id SET DEFAULT nextval('public.shopping_carts_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: wishlists id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wishlists ALTER COLUMN id SET DEFAULT nextval('public.wishlists_id_seq'::regclass);


--
-- Data for Name: addresses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.addresses (id, user_id, address, is_main, created_at, updated_at, deleted_at) FROM stdin;
1	1	123 Main Street, City A	t	2024-11-21 22:52:13.120058	2024-11-21 22:52:13.120058	\N
2	2	456 Oak Avenue, City B	t	2024-11-21 22:52:16.329574	2024-11-21 22:52:16.329574	\N
3	2	789 Pine Road, City B	f	2024-11-21 22:52:16.329574	2024-11-21 22:52:16.329574	\N
4	3	111 Maple Lane, City C	t	2024-11-21 22:52:19.08705	2024-11-21 22:52:19.08705	\N
5	4	222 Cedar Street, City D	t	2024-11-21 22:52:22.762853	2024-11-21 22:52:22.762853	\N
6	4	333 Birch Boulevard, City D	f	2024-11-21 22:52:22.762853	2024-11-21 22:52:22.762853	\N
7	4	444 Spruce Terrace, City D	f	2024-11-21 22:52:22.762853	2024-11-21 22:52:22.762853	\N
8	5	jl. Cilamaya Wetan	f	2024-11-21 22:52:25.040112	2024-11-23 18:50:38.704332	\N
10	5	Jalan Kembang	t	2024-11-23 16:27:06.404663	2024-11-23 18:50:38.704332	\N
19	11	Jalan Kembang	f	2024-11-24 12:18:09.99145	2024-11-24 20:39:58.957552	\N
\.


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.categories (id, name, created_at, updated_at, deleted_at) FROM stdin;
1	Electronics	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
2	Clothing	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
3	Home Appliances	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
4	Books	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
5	Toys	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
\.


--
-- Data for Name: checkout_items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.checkout_items (id, checkout_id, product_variant_id, qty, total, created_at, updated_at, deleted_at) FROM stdin;
46	14	2	2	20400000	2024-11-24 17:49:36.021756	2024-11-24 17:49:36.021756	\N
47	14	3	4	570000	2024-11-24 17:49:36.021756	2024-11-24 17:49:36.021756	\N
\.


--
-- Data for Name: checkouts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.checkouts (id, user_id, total_amount, payment, payment_status, created_at, updated_at, deleted_at) FROM stdin;
1	1	10000000	Credit Card	Paid	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
2	2	150000	Cash	Pending	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
3	3	5000000	Debit	Paid	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
4	4	3000000	E-Wallet	Paid	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
5	5	80000	Bank Transfer	Pending	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
6	5	37062500	COD	Paid	2024-11-22 22:47:02.427684	2024-11-22 22:47:02.427684	\N
7	5	35400000	COD	Paid	2024-11-22 23:07:48.489066	2024-11-22 23:07:48.489066	\N
10	11	15080000	COD	Paid	2024-11-24 16:30:50.086978	2024-11-24 16:30:50.086978	\N
13	11	320000	COD	Paid	2024-11-24 17:43:39.653134	2024-11-24 17:43:39.653134	\N
14	11	20970000	COD	Paid	2024-11-24 17:49:36.021756	2024-11-24 17:49:36.021756	\N
\.


--
-- Data for Name: images; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.images (id, product_id, image_url, created_at, updated_at, deleted_at) FROM stdin;
1	1	http://example.com/smartphone-front.jpg	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
2	1	http://example.com/smartphone-back.jpg	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
3	2	http://example.com/laptop-front.jpg	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
4	3	http://example.com/tshirt-front.jpg	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
5	4	http://example.com/vacuum.jpg	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
\.


--
-- Data for Name: previews; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.previews (id, checkout_item_id, rating, comentars, created_at, updated_at, deleted_at) FROM stdin;
1	1	4.5	Amazing product!	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
2	2	4	Good quality, arrived late.	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
3	3	5	Excellent performance!	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
4	4	3.5	Works well but noisy.	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
5	5	4.8	Loved the story!	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
6	6	4.5	Great product!	2024-11-06 10:00:00	2024-11-06 10:30:00	\N
7	7	5	Excellent quality!	2024-11-06 11:00:00	2024-11-06 11:30:00	\N
8	8	3.8	Good but a bit expensive.	2024-11-07 12:00:00	2024-11-07 12:30:00	\N
9	9	4	Satisfied with the purchase.	2024-11-07 13:00:00	2024-11-07 13:30:00	\N
10	10	5	Fantastic! Highly recommend.	2024-11-08 14:00:00	2024-11-08 14:30:00	\N
11	11	4.7	Very good, fast delivery.	2024-11-08 15:00:00	2024-11-08 15:30:00	\N
12	12	3.5	Average quality.	2024-11-09 16:00:00	2024-11-09 16:30:00	\N
13	13	4.2	Pretty good, worth it.	2024-11-09 17:00:00	2024-11-09 17:30:00	\N
14	14	4.9	Exceeded expectations.	2024-11-10 18:00:00	2024-11-10 18:30:00	\N
15	15	4.8	Would buy again!	2024-11-10 19:00:00	2024-11-10 19:30:00	\N
16	6	4.5	Great product!	2024-11-06 10:00:00	2024-11-06 10:30:00	\N
17	7	5	Excellent quality!	2024-11-06 11:00:00	2024-11-06 11:30:00	\N
18	8	3.8	Good but a bit expensive.	2024-11-07 12:00:00	2024-11-07 12:30:00	\N
19	9	4	Satisfied with the purchase.	2024-11-07 13:00:00	2024-11-07 13:30:00	\N
20	10	5	Fantastic! Highly recommend.	2024-11-08 14:00:00	2024-11-08 14:30:00	\N
21	11	4.7	Very good, fast delivery.	2024-11-08 15:00:00	2024-11-08 15:30:00	\N
22	12	3.5	Average quality.	2024-11-09 16:00:00	2024-11-09 16:30:00	\N
23	13	4.2	Pretty good, worth it.	2024-11-09 17:00:00	2024-11-09 17:30:00	\N
24	14	4.9	Exceeded expectations.	2024-11-10 18:00:00	2024-11-10 18:30:00	\N
25	15	4.8	Would buy again!	2024-11-10 19:00:00	2024-11-10 19:30:00	\N
\.


--
-- Data for Name: product_varians; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.product_varians (id, product_id, size, color, stock, image_url, created_at, updated_at, deleted_at) FROM stdin;
1	1	6.5 inch	Black	50	http://example.com/smartphone-black.jpg	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
4	4	Standard	White	30	http://example.com/vacuum-white.jpg	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
6	1	L	Black	10	\N	\N	\N	\N
5	5	Hardcover	N/A	196	http://example.com/novel-hardcover.jpg	2024-11-19 15:06:28.853789	2024-11-24 17:43:39.653134	\N
2	2	15 inch	Silver	18	http://example.com/laptop-silver.jpg	2024-11-19 15:06:28.853789	2024-11-24 17:49:36.021756	\N
3	3	M	Blue	96	http://example.com/tshirt-blue.jpg	2024-11-19 15:06:28.853789	2024-11-24 17:49:36.021756	\N
\.


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.products (id, name, image_url, price, discount, category_id, description, created_at, updated_at, deleted_at) FROM stdin;
1	Smartphone	http://example.com/smartphone.jpg	5000000	10	1	A high-end smartphone.	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
2	Laptop	http://example.com/laptop.jpg	12000000	15	1	A powerful laptop.	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
4	Vacuum Cleaner	http://example.com/vacuum.jpg	3000000	20	3	High-suction vacuum cleaner.	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
5	Novel	http://example.com/novel.jpg	80000	0	4	A bestselling novel.	2024-10-19 15:06:28.853	2024-11-19 15:06:28.853789	\N
3	T-shirt	http://example.com/tshirt.jpg	150000	5	2	A comfortable cotton t-shirt.	2024-10-19 15:06:28.853	2024-11-19 15:06:28.853789	\N
\.


--
-- Data for Name: promotions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.promotions (id, title, subtitle, image_url, path_url, start_date, end_date, is_promo, created_at, updated_at, deleted_at) FROM stdin;
2	Promo 2	Gratis ongkir!	https://example.com/image2.jpg	https://example.com/promo/2	2024-11-20	2024-11-27	t	2024-11-20 07:50:45.039406	2024-11-20 07:50:45.039406	\N
3	Promo 3	Beli 1 gratis 1	https://example.com/image3.jpg	https://example.com/promo/3	2024-11-20	2024-11-27	t	2024-11-20 07:50:45.039406	2024-11-20 07:50:45.039406	\N
4	Promo 4	Hanya hari ini	https://example.com/image4.jpg	https://example.com/promo/4	2024-11-20	2024-11-27	t	2024-11-20 07:50:45.039406	2024-11-20 07:50:45.039406	\N
5	Promo 5	Diskon spesial	https://example.com/image5.jpg	https://example.com/promo/5	2024-11-20	2024-11-27	t	2024-11-20 07:50:45.039406	2024-11-20 07:50:45.039406	\N
6	Promo Lama 1	Promo habis	https://example.com/old1.jpg	https://example.com/promo/old1	2023-01-01	2023-01-07	f	2024-11-20 07:50:52.892836	2024-11-20 07:50:52.892836	\N
7	Promo Lama 2	Sudah tidak aktif	https://example.com/old2.jpg	https://example.com/promo/old2	2023-02-01	2023-02-07	f	2024-11-20 07:50:52.892836	2024-11-20 07:50:52.892836	\N
8	Promo Lama 3	Event berakhir	https://example.com/old3.jpg	https://example.com/promo/old3	2023-03-01	2023-03-07	f	2024-11-20 07:50:52.892836	2024-11-20 07:50:52.892836	\N
9	Promo Lama 4	Promo ini kadaluarsa	https://example.com/old4.jpg	https://example.com/promo/old4	2023-04-01	2023-04-07	f	2024-11-20 07:50:52.892836	2024-11-20 07:50:52.892836	\N
10	Promo Lama 5	Tidak berlaku	https://example.com/old5.jpg	https://example.com/promo/old5	2023-05-01	2023-05-07	f	2024-11-20 07:50:52.892836	2024-11-20 07:50:52.892836	\N
11	Promo 1	Diskon besar-besaran	https://example.com/image1.jpg	https://example.com/promo/1	2024-11-10	2024-11-17	t	2024-11-20 07:56:41.37815	2024-11-20 07:56:41.37815	\N
12	Promo 2	Gratis ongkir!	https://example.com/image2.jpg	https://example.com/promo/2	2024-11-12	2024-11-19	t	2024-11-20 07:56:41.37815	2024-11-20 07:56:41.37815	\N
13	Promo 3	Beli 1 gratis 1	https://example.com/image3.jpg	https://example.com/promo/3	2024-11-08	2024-11-15	t	2024-11-20 07:56:41.37815	2024-11-20 07:56:41.37815	\N
14	Promo 4	Hanya hari ini	https://example.com/image4.jpg	https://example.com/promo/4	2024-11-14	2024-11-21	t	2024-11-20 07:56:41.37815	2024-11-20 07:56:41.37815	\N
15	Promo 5	Diskon spesial	https://example.com/image5.jpg	https://example.com/promo/5	2024-11-09	2024-11-16	t	2024-11-20 07:56:41.37815	2024-11-20 07:56:41.37815	\N
20	Banner 5	Ikuti terus update kami	https://example.com/banner5.jpg	https://example.com/banner/5	2024-11-09	2024-11-16	f	2024-11-20 07:56:48.833157	2024-11-20 07:56:48.833157	\N
1	Promo 1	Diskon besar-besaran	https://example.com/image1.jpg	https://example.com/promo/1	2024-11-23	2024-11-30	t	2024-11-20 07:50:45.039406	2024-11-20 07:50:45.039406	\N
16	Banner 1	Informasi menarik	https://example.com/banner1.jpg	https://example.com/banner/1	2024-11-23	2024-11-30	f	2024-11-20 07:56:48.833157	2024-11-20 07:56:48.833157	\N
17	Banner 2	Promo akan datang	https://example.com/banner2.jpg	https://example.com/banner/2	2024-11-23	2024-11-30	f	2024-11-20 07:56:48.833157	2024-11-20 07:56:48.833157	\N
18	Banner 3	Stay tuned!	https://example.com/banner3.jpg	https://example.com/banner/3	2024-11-25	2024-12-02	f	2024-11-20 07:56:48.833157	2024-11-20 07:56:48.833157	\N
19	Banner 4	Jangan lewatkan!	https://example.com/banner4.jpg	https://example.com/banner/4	2024-11-22	2024-11-29	f	2024-11-20 07:56:48.833157	2024-11-20 07:56:48.833157	\N
21	Banner 1	Informasi menarik	https://example.com/banner1.jpg	https://example.com/banner/1	2024-11-29	2024-12-06	f	2024-11-20 10:14:15.909535	2024-11-20 10:14:15.909535	\N
22	Banner 2	Promo akan datang	https://example.com/banner2.jpg	https://example.com/banner/2	2024-11-29	2024-12-06	f	2024-11-20 10:14:15.909535	2024-11-20 10:14:15.909535	\N
\.


--
-- Data for Name: recomendeds; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.recomendeds (id, title, image_url, subtitle, product_id, status, created_at, updated_at, deleted_at) FROM stdin;
1	Top Deals	https://example.com/image1.jpg	Best offers of the week	1	f	2024-11-20 15:42:18.238484	2024-11-20 15:42:18.238484	\N
2	New Arrivals	https://example.com/image2.jpg	Check out the latest products	2	t	2024-11-20 15:42:18.238484	2024-11-20 15:42:18.238484	\N
3	Limited Edition	https://example.com/image3.jpg	Exclusive products just for you	3	t	2024-11-20 15:42:18.238484	2024-11-20 15:42:18.238484	\N
4	Seasonal Discounts	https://example.com/image4.jpg	Grab discounts this season	4	f	2024-11-20 15:42:18.238484	2024-11-20 15:42:18.238484	\N
5	Editor’s Pick	https://example.com/image5.jpg	Curated collection of top picks	5	t	2024-11-20 15:42:18.238484	2024-11-20 15:42:18.238484	\N
\.


--
-- Data for Name: shopping_carts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.shopping_carts (id, product_variant_id, user_id, qty, created_at, updated_at, deleted_at) FROM stdin;
1	1	1	2	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
2	2	2	1	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
3	3	3	3	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
4	4	4	1	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, name, username, phone, email, password, token, expired, created_at, updated_at, deleted_at) FROM stdin;
2	Jane Smith	janesmith	08123456788	jane@example.com	password456	token456	2024-11-26 15:06:28.853789	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
3	Alice Brown	aliceb	08123456787	alice@example.com	password789	token789	2024-11-26 15:06:28.853789	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
4	Bob White	bobw	08123456786	bob@example.com	password321	token321	2024-11-26 15:06:28.853789	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
5	Fikri		085791526831	fikri123@mail.com	citykadubanen	a83202b8-5710-480f-9635-951834468ea4	2024-11-23 12:34:57.978334	2024-11-19 15:06:28.853789	2024-11-23 15:46:00.903259	\N
10	fikri	fikri72	08579152683	fikri@xample.com	123kadubanen	\N	\N	\N	\N	\N
7	fikri	fikri121	+62956372891	fikri@example.com	password123	\N	\N	\N	\N	\N
1	John Doe	johndoe	08123456789	john@example.com	password123	3e3e4a79-4bbd-4ee2-800f-a9477a307f9a	2024-11-21 00:03:00.021956	2024-11-19 15:06:28.853789	2024-11-19 15:06:28.853789	\N
11	Deni	Deni351	085791526821	deni123@mail.com	kadubanen	8922f89a-56aa-42e2-b428-92e756b3b2f0	2024-11-24 14:27:41.786904	2024-11-23 20:40:51.17568	2024-11-24 18:25:12.073562	\N
\.


--
-- Data for Name: wishlists; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wishlists (id, product_variant_id, user_id, created_at, deleted_at) FROM stdin;
5	5	5	2024-11-19 15:06:28.853789	\N
24	4	11	2024-11-24 11:02:24.82844	\N
\.


--
-- Name: addresses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.addresses_id_seq', 19, true);


--
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.categories_id_seq', 5, true);


--
-- Name: checkout_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.checkout_items_id_seq', 47, true);


--
-- Name: checkouts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.checkouts_id_seq', 14, true);


--
-- Name: images_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.images_id_seq', 5, true);


--
-- Name: previews_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.previews_id_seq', 25, true);


--
-- Name: product_varians_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.product_varians_id_seq', 5, true);


--
-- Name: products_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.products_id_seq', 5, true);


--
-- Name: promotions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.promotions_id_seq', 22, true);


--
-- Name: recomendeds_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.recomendeds_id_seq', 5, true);


--
-- Name: shopping_carts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.shopping_carts_id_seq', 28, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 11, true);


--
-- Name: wishlists_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wishlists_id_seq', 24, true);


--
-- Name: addresses addresses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.addresses
    ADD CONSTRAINT addresses_pkey PRIMARY KEY (id);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- Name: checkout_items checkout_items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checkout_items
    ADD CONSTRAINT checkout_items_pkey PRIMARY KEY (id);


--
-- Name: checkouts checkouts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checkouts
    ADD CONSTRAINT checkouts_pkey PRIMARY KEY (id);


--
-- Name: images images_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.images
    ADD CONSTRAINT images_pkey PRIMARY KEY (id);


--
-- Name: previews previews_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.previews
    ADD CONSTRAINT previews_pkey PRIMARY KEY (id);


--
-- Name: product_varians product_varians_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_varians
    ADD CONSTRAINT product_varians_pkey PRIMARY KEY (id);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- Name: promotions promotions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.promotions
    ADD CONSTRAINT promotions_pkey PRIMARY KEY (id);


--
-- Name: recomendeds recomendeds_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.recomendeds
    ADD CONSTRAINT recomendeds_pkey PRIMARY KEY (id);


--
-- Name: shopping_carts shopping_carts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shopping_carts
    ADD CONSTRAINT shopping_carts_pkey PRIMARY KEY (id);


--
-- Name: shopping_carts unique_product_user; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shopping_carts
    ADD CONSTRAINT unique_product_user UNIQUE (product_variant_id, user_id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_phone_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_phone_key UNIQUE (phone);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: wishlists wishlists_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wishlists
    ADD CONSTRAINT wishlists_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

