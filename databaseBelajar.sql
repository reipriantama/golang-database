--
-- PostgreSQL database dump
--

-- Dumped from database version 16.3 (Homebrew)
-- Dumped by pg_dump version 16.3 (Homebrew)

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
-- Name: product_category; Type: TYPE; Schema: public; Owner: plabs.id
--

CREATE TYPE public.product_category AS ENUM (
    'Makanan',
    'Minuman',
    'Lain-lain'
);


ALTER TYPE public.product_category OWNER TO "plabs.id";

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: products; Type: TABLE; Schema: public; Owner: plabs.id
--

CREATE TABLE public.products (
    id character varying(10) NOT NULL,
    name character varying(100) NOT NULL,
    description text,
    price integer NOT NULL,
    quantity integer DEFAULT 0 NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    category public.product_category
);


ALTER TABLE public.products OWNER TO "plabs.id";

--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: plabs.id
--

COPY public.products (id, name, description, price, quantity, created_at, category) FROM stdin;
P0001	Mie Ayam Original	\N	15000	100	2024-05-17 13:53:24.618705	Makanan
P0002	Mie Ayam Bakso Tahu	Mie Ayam Original + Bakso Tahu	20000	100	2024-05-17 13:53:24.618705	Makanan
P0003	Mie Ayam Ceker	Mie Ayam + Ceker	20000	100	2024-05-17 13:53:24.618705	Makanan
P0004	Mie Ayam Special	\N	30000	100	2024-05-17 13:53:24.618705	Makanan
P0005	Mie Ayam Yamin	\N	15000	100	2024-05-17 13:53:24.618705	Makanan
P0006	Es teh tawar	\N	10000	100	2024-05-20 11:35:36.394123	Minuman
P0007	Es Campur	\N	20000	100	2024-05-20 11:35:36.394123	Minuman
P0008	Es Jeruk	\N	10000	100	2024-05-20 11:35:36.394123	Minuman
\.


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: plabs.id
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

