--
-- PostgreSQL database dump
--

-- Dumped from database version 15.1 (Ubuntu 15.1-1.pgdg22.04+1)
-- Dumped by pg_dump version 15.1 (Ubuntu 15.1-1.pgdg22.04+1)

-- Started on 2022-12-15 22:30:14 WIB

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 215 (class 1259 OID 16632)
-- Name: source_of_funds; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.source_of_funds (
    id integer NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE public.source_of_funds OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 16631)
-- Name: source_of_funds_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.source_of_funds_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.source_of_funds_id_seq OWNER TO postgres;

--
-- TOC entry 3414 (class 0 OID 0)
-- Dependencies: 214
-- Name: source_of_funds_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.source_of_funds_id_seq OWNED BY public.source_of_funds.id;


--
-- TOC entry 222 (class 1259 OID 16741)
-- Name: transactions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transactions (
    id integer NOT NULL,
    amount bigint,
    source_wallet_id integer,
    destination_wallet_id integer,
    description character varying(35),
    transaction_type character varying(12),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone,
    source_of_fund_id integer
);


ALTER TABLE public.transactions OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 16740)
-- Name: transactions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.transactions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.transactions_id_seq OWNER TO postgres;

--
-- TOC entry 3415 (class 0 OID 0)
-- Dependencies: 221
-- Name: transactions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.transactions_id_seq OWNED BY public.transactions.id;


--
-- TOC entry 217 (class 1259 OID 16687)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    email character varying NOT NULL,
    password character varying,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 16686)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- TOC entry 3416 (class 0 OID 0)
-- Dependencies: 216
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 220 (class 1259 OID 16701)
-- Name: wallets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wallets (
    id integer NOT NULL,
    amount bigint,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone,
    user_id integer
);


ALTER TABLE public.wallets OWNER TO postgres;

--
-- TOC entry 218 (class 1259 OID 16699)
-- Name: wallet_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wallet_id_seq
    START WITH 777000
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wallet_id_seq OWNER TO postgres;

--
-- TOC entry 3417 (class 0 OID 0)
-- Dependencies: 218
-- Name: wallet_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wallet_id_seq OWNED BY public.wallets.id;


--
-- TOC entry 219 (class 1259 OID 16700)
-- Name: wallets_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wallets_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wallets_id_seq OWNER TO postgres;

--
-- TOC entry 3418 (class 0 OID 0)
-- Dependencies: 219
-- Name: wallets_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wallets_id_seq OWNED BY public.wallets.id;


--
-- TOC entry 3232 (class 2604 OID 16635)
-- Name: source_of_funds id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.source_of_funds ALTER COLUMN id SET DEFAULT nextval('public.source_of_funds_id_seq'::regclass);


--
-- TOC entry 3239 (class 2604 OID 16744)
-- Name: transactions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions ALTER COLUMN id SET DEFAULT nextval('public.transactions_id_seq'::regclass);


--
-- TOC entry 3233 (class 2604 OID 16690)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 3236 (class 2604 OID 16714)
-- Name: wallets id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wallets ALTER COLUMN id SET DEFAULT nextval('public.wallet_id_seq'::regclass);


--
-- TOC entry 3401 (class 0 OID 16632)
-- Dependencies: 215
-- Data for Name: source_of_funds; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.source_of_funds (id, name) FROM stdin;
1	Bank Transfer
2	Credit Card
3	Cash
\.


--
-- TOC entry 3408 (class 0 OID 16741)
-- Dependencies: 222
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.transactions (id, amount, source_wallet_id, destination_wallet_id, description, transaction_type, created_at, updated_at, deleted_at, source_of_fund_id) FROM stdin;
4	100000	\N	777001	Top Up from Bank Transfer	Topup	2022-12-14 23:06:00.965627	2022-12-14 23:06:00.965627	\N	1
5	200000	\N	777001	Top Up from Credit Card	Topup	2022-12-15 11:21:50.941741	2022-12-15 11:21:50.941741	\N	2
6	20000	777001	777003	hutang kemarin	Transfer	2022-12-15 12:04:45.558253	2022-12-15 12:04:45.558253	\N	\N
7	250000	\N	777002	Top Up from Bank Transfer	Topup	2022-12-15 22:04:48.540443	2022-12-15 22:04:48.540443	\N	1
8	350000	\N	777002	Top Up from Bank Transfer	Topup	2022-12-15 22:06:17.374716	2022-12-15 22:06:17.374716	\N	1
9	150000	\N	777002	Top Up from Bank Transfer	Topup	2022-12-15 22:07:38.717672	2022-12-15 22:07:38.717672	\N	1
10	150000	\N	777004	Top Up from Cash	Topup	2022-12-15 22:11:37.750605	2022-12-15 22:11:37.750605	\N	3
11	1020000	\N	777005	Top Up from Bank Transfer	Topup	2022-12-15 22:12:54.99433	2022-12-15 22:12:54.99433	\N	1
12	105000	777005	777003	food payment	Transfer	2022-12-15 22:13:50.353214	2022-12-15 22:13:50.353214	\N	\N
13	12000	777005	777005	Family Chicken	Transfer	2022-12-15 22:14:11.411244	2022-12-15 22:14:11.411244	\N	\N
14	11000	777004	777001	Bike Fuel	Transfer	2022-12-15 22:16:06.156406	2022-12-15 22:16:06.156406	\N	\N
15	215000	777002	777004	Cosmetics	Transfer	2022-12-15 22:17:08.529508	2022-12-15 22:17:08.529508	\N	\N
16	300000	\N	777005	Top Up from Cash	Topup	2022-12-15 22:19:14.373401	2022-12-15 22:19:14.373401	\N	3
17	420000	\N	777003	Top Up from Bank Transfer	Topup	2022-12-15 22:21:07.591498	2022-12-15 22:21:07.591498	\N	1
18	38000	777003	777002	Coffee	Transfer	2022-12-15 22:21:41.695825	2022-12-15 22:21:41.695825	\N	\N
19	22000	777003	777001	snack	Transfer	2022-12-15 22:22:25.993193	2022-12-15 22:22:25.993193	\N	\N
20	100000	\N	777002	Top Up from Bank Transfer	Topup	2022-12-15 22:24:42.084366	2022-12-15 22:24:42.084366	\N	1
21	180000	\N	777004	Top Up from Credit Card	Topup	2022-12-15 22:25:30.464729	2022-12-15 22:25:30.464729	\N	2
22	15000	777004	777005	ongkos	Transfer	2022-12-15 22:26:42.730081	2022-12-15 22:26:42.730081	\N	\N
23	16000	777004	777002		Transfer	2022-12-15 22:29:19.84904	2022-12-15 22:29:19.84904	\N	\N
\.


--
-- TOC entry 3403 (class 0 OID 16687)
-- Dependencies: 217
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, email, password, created_at, updated_at, deleted_at) FROM stdin;
2	don@email.com	$2a$04$nSJlGx.mzV.vG/kqXgD/we3ZG3Ouz6yHrEn9KW4ojyupOCPUu.PP2	2022-12-14 13:31:01.258216	2022-12-14 13:31:01.258216	\N
3	adel@email.com	$2a$04$ZEsQSI7/NPnV09tBFrKXneUgiinkidljDgQwYjzfMTYRYDH9kqKHe	2022-12-14 15:47:54.288535	2022-12-14 15:47:54.288535	\N
4	ship@email.com	$2a$04$DYbwpnV3atVuw3jSvp3w6.tFWLr/z7qtsFZUvrLnWVBsn9VymUFNa	2022-12-14 15:48:22.628029	2022-12-14 15:48:22.628029	\N
5	myadmin@email.com	$2a$04$gSIMdTyd/l3lTkLApm4CruVqKF0OOBlQ7BTIGXxzA0VyUl89d3hWO	2022-12-14 15:48:40.133506	2022-12-14 15:48:40.133506	\N
6	brad@email.com	$2a$04$KkQuZ27Ko0IcGXkzrKyOT.Q7Jw8x0kl7rNim0IdBIyNJockm48G2y	2022-12-14 15:49:16.66963	2022-12-14 15:49:16.66963	\N
\.


--
-- TOC entry 3406 (class 0 OID 16701)
-- Dependencies: 220
-- Data for Name: wallets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wallets (id, amount, created_at, updated_at, deleted_at, user_id) FROM stdin;
777003	485000	2022-12-14 15:48:22.628235	2022-12-15 22:22:25.992412	\N	4
777001	313000	2022-12-14 13:31:01.258483	2022-12-15 22:22:25.992874	\N	2
777005	1230000	2022-12-14 15:49:16.669856	2022-12-15 22:26:42.729761	\N	6
777004	503000	2022-12-14 15:48:40.133718	2022-12-15 22:29:19.848075	\N	5
777002	689000	2022-12-14 15:47:54.288952	2022-12-15 22:29:19.848591	\N	3
\.


--
-- TOC entry 3419 (class 0 OID 0)
-- Dependencies: 214
-- Name: source_of_funds_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.source_of_funds_id_seq', 3, true);


--
-- TOC entry 3420 (class 0 OID 0)
-- Dependencies: 221
-- Name: transactions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.transactions_id_seq', 23, true);


--
-- TOC entry 3421 (class 0 OID 0)
-- Dependencies: 216
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 6, true);


--
-- TOC entry 3422 (class 0 OID 0)
-- Dependencies: 218
-- Name: wallet_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wallet_id_seq', 777005, true);


--
-- TOC entry 3423 (class 0 OID 0)
-- Dependencies: 219
-- Name: wallets_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wallets_id_seq', 1, false);


--
-- TOC entry 3243 (class 2606 OID 16639)
-- Name: source_of_funds source_of_funds_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.source_of_funds
    ADD CONSTRAINT source_of_funds_pkey PRIMARY KEY (id);


--
-- TOC entry 3253 (class 2606 OID 16748)
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- TOC entry 3249 (class 2606 OID 16765)
-- Name: wallets user_id_unique; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wallets
    ADD CONSTRAINT user_id_unique UNIQUE (user_id);


--
-- TOC entry 3245 (class 2606 OID 16698)
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- TOC entry 3247 (class 2606 OID 16696)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 3251 (class 2606 OID 16708)
-- Name: wallets wallets_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wallets
    ADD CONSTRAINT wallets_pkey PRIMARY KEY (id);


--
-- TOC entry 3255 (class 2606 OID 16754)
-- Name: transactions fk_destination_wallet; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT fk_destination_wallet FOREIGN KEY (destination_wallet_id) REFERENCES public.wallets(id);


--
-- TOC entry 3256 (class 2606 OID 16759)
-- Name: transactions fk_source_of_fund; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT fk_source_of_fund FOREIGN KEY (source_of_fund_id) REFERENCES public.source_of_funds(id);


--
-- TOC entry 3257 (class 2606 OID 16749)
-- Name: transactions fk_source_wallet; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT fk_source_wallet FOREIGN KEY (source_wallet_id) REFERENCES public.wallets(id);


--
-- TOC entry 3254 (class 2606 OID 16709)
-- Name: wallets fk_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wallets
    ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.users(id);


-- Completed on 2022-12-15 22:30:14 WIB

--
-- PostgreSQL database dump complete
--

