--
-- PostgreSQL database cluster dump
--

SET default_transaction_read_only = off;

SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;

--
-- Roles
--

CREATE ROLE hew;
ALTER ROLE hew WITH SUPERUSER INHERIT CREATEROLE CREATEDB LOGIN REPLICATION BYPASSRLS;
--
-- Database creation
--

CREATE DATABASE hew WITH TEMPLATE = template0 OWNER = postgres;
REVOKE CONNECT,TEMPORARY ON DATABASE template1 FROM PUBLIC;
GRANT CONNECT ON DATABASE template1 TO PUBLIC;


\connect hew

SET default_transaction_read_only = off;

--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.1
-- Dumped by pg_dump version 9.6.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: activities; Type: TABLE; Schema: public; Owner: hew
--

CREATE TABLE activities (
    id integer NOT NULL,
    name character varying(255)
);


ALTER TABLE activities OWNER TO hew;

--
-- Name: activities_id_seq; Type: SEQUENCE; Schema: public; Owner: hew
--

CREATE SEQUENCE activities_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE activities_id_seq OWNER TO hew;

--
-- Name: activities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hew
--

ALTER SEQUENCE activities_id_seq OWNED BY activities.id;


--
-- Name: menus; Type: TABLE; Schema: public; Owner: hew
--

CREATE TABLE menus (
    id integer NOT NULL,
    name character varying(255),
    price numeric
);


ALTER TABLE menus OWNER TO hew;

--
-- Name: menus_id_seq; Type: SEQUENCE; Schema: public; Owner: hew
--

CREATE SEQUENCE menus_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE menus_id_seq OWNER TO hew;

--
-- Name: menus_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hew
--

ALTER SEQUENCE menus_id_seq OWNED BY menus.id;


--
-- Name: organizations; Type: TABLE; Schema: public; Owner: hew
--

CREATE TABLE organizations (
    id integer NOT NULL,
    name character varying(200)
);


ALTER TABLE organizations OWNER TO hew;

--
-- Name: organizations_id_seq; Type: SEQUENCE; Schema: public; Owner: hew
--

CREATE SEQUENCE organizations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE organizations_id_seq OWNER TO hew;

--
-- Name: organizations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hew
--

ALTER SEQUENCE organizations_id_seq OWNED BY organizations.id;


--
-- Name: restuarants; Type: TABLE; Schema: public; Owner: hew
--

CREATE TABLE restuarants (
    id integer NOT NULL,
    name character varying(500)
);


ALTER TABLE restuarants OWNER TO hew;

--
-- Name: restuarants_id_seq; Type: SEQUENCE; Schema: public; Owner: hew
--

CREATE SEQUENCE restuarants_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE restuarants_id_seq OWNER TO hew;

--
-- Name: restuarants_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hew
--

ALTER SEQUENCE restuarants_id_seq OWNED BY restuarants.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: hew
--

CREATE TABLE users (
    id integer NOT NULL,
    username character varying(200),
    password character varying(100)
);


ALTER TABLE users OWNER TO hew;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: hew
--

CREATE SEQUENCE users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE users_id_seq OWNER TO hew;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hew
--

ALTER SEQUENCE users_id_seq OWNED BY users.id;


--
-- Name: activities id; Type: DEFAULT; Schema: public; Owner: hew
--

ALTER TABLE ONLY activities ALTER COLUMN id SET DEFAULT nextval('activities_id_seq'::regclass);


--
-- Name: menus id; Type: DEFAULT; Schema: public; Owner: hew
--

ALTER TABLE ONLY menus ALTER COLUMN id SET DEFAULT nextval('menus_id_seq'::regclass);


--
-- Name: organizations id; Type: DEFAULT; Schema: public; Owner: hew
--

ALTER TABLE ONLY organizations ALTER COLUMN id SET DEFAULT nextval('organizations_id_seq'::regclass);


--
-- Name: restuarants id; Type: DEFAULT; Schema: public; Owner: hew
--

ALTER TABLE ONLY restuarants ALTER COLUMN id SET DEFAULT nextval('restuarants_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: hew
--

ALTER TABLE ONLY users ALTER COLUMN id SET DEFAULT nextval('users_id_seq'::regclass);


--
-- Data for Name: activities; Type: TABLE DATA; Schema: public; Owner: hew
--

COPY activities (id, name) FROM stdin;
\.


--
-- Name: activities_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hew
--

SELECT pg_catalog.setval('activities_id_seq', 1, false);


--
-- Data for Name: menus; Type: TABLE DATA; Schema: public; Owner: hew
--

COPY menus (id, name, price) FROM stdin;
\.


--
-- Name: menus_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hew
--

SELECT pg_catalog.setval('menus_id_seq', 1, false);


--
-- Data for Name: organizations; Type: TABLE DATA; Schema: public; Owner: hew
--

COPY organizations (id, name) FROM stdin;
\.


--
-- Name: organizations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hew
--

SELECT pg_catalog.setval('organizations_id_seq', 1, false);


--
-- Data for Name: restuarants; Type: TABLE DATA; Schema: public; Owner: hew
--

COPY restuarants (id, name) FROM stdin;
\.


--
-- Name: restuarants_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hew
--

SELECT pg_catalog.setval('restuarants_id_seq', 1, false);


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: hew
--

COPY users (id, username, password) FROM stdin;
\.


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hew
--

SELECT pg_catalog.setval('users_id_seq', 1, false);


--
-- Name: menus menus_pkey; Type: CONSTRAINT; Schema: public; Owner: hew
--

ALTER TABLE ONLY menus
    ADD CONSTRAINT menus_pkey PRIMARY KEY (id);


--
-- Name: organizations organizations_pkey; Type: CONSTRAINT; Schema: public; Owner: hew
--

ALTER TABLE ONLY organizations
    ADD CONSTRAINT organizations_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: hew
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

\connect postgres

SET default_transaction_read_only = off;

--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.1
-- Dumped by pg_dump version 9.6.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: postgres; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON DATABASE postgres IS 'default administrative connection database';


--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


--
-- PostgreSQL database dump complete
--

\connect template1

SET default_transaction_read_only = off;

--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.1
-- Dumped by pg_dump version 9.6.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: template1; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON DATABASE template1 IS 'default template for new databases';


--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


--
-- PostgreSQL database dump complete
--

--
-- PostgreSQL database cluster dump complete
--

