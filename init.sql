--
-- CESJB — Script de inicialização do banco de dados
--

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

-- Tabela de administradores
CREATE TABLE public.admins (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    email character varying(100) NOT NULL,
    password character varying(255) NOT NULL
);

CREATE SEQUENCE public.admins_id_seq
    AS integer START WITH 1 INCREMENT BY 1
    NO MINVALUE NO MAXVALUE CACHE 1;

ALTER SEQUENCE public.admins_id_seq OWNED BY public.admins.id;
ALTER TABLE ONLY public.admins ALTER COLUMN id SET DEFAULT nextval('public.admins_id_seq'::regclass);
ALTER TABLE ONLY public.admins ADD CONSTRAINT admins_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.admins ADD CONSTRAINT admins_email_key UNIQUE (email);

-- Tabela de associados
CREATE TABLE public.associates (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    cpf character varying(14) NOT NULL,
    email character varying(100),
    tel character varying(20),
    date_of_birth timestamp without time zone,
    association_date timestamp without time zone,
    address character varying(200),
    status boolean DEFAULT true,
    "position" character varying(50)
);

CREATE SEQUENCE public.associates_id_seq
    AS integer START WITH 1 INCREMENT BY 1
    NO MINVALUE NO MAXVALUE CACHE 1;

ALTER SEQUENCE public.associates_id_seq OWNED BY public.associates.id;
ALTER TABLE ONLY public.associates ALTER COLUMN id SET DEFAULT nextval('public.associates_id_seq'::regclass);
ALTER TABLE ONLY public.associates ADD CONSTRAINT associates_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.associates ADD CONSTRAINT associates_cpf_key UNIQUE (cpf);

-- Tabela de pagamentos
CREATE TABLE public.payments (
    id integer NOT NULL,
    associate_id integer NOT NULL,
    payment_date date NOT NULL,
    value numeric(10,2) NOT NULL,
    status boolean DEFAULT false NOT NULL,
    competence date
);

CREATE SEQUENCE public.payments_id_seq
    AS integer START WITH 1 INCREMENT BY 1
    NO MINVALUE NO MAXVALUE CACHE 1;

ALTER SEQUENCE public.payments_id_seq OWNED BY public.payments.id;
ALTER TABLE ONLY public.payments ALTER COLUMN id SET DEFAULT nextval('public.payments_id_seq'::regclass);
ALTER TABLE ONLY public.payments ADD CONSTRAINT payments_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.payments ADD CONSTRAINT unique_associate_competence UNIQUE (associate_id, competence);
ALTER TABLE ONLY public.payments ADD CONSTRAINT fk_associate FOREIGN KEY (associate_id) REFERENCES public.associates(id) ON DELETE CASCADE;

-- Admin padrão
INSERT INTO public.admins (name, email, password)
VALUES ('testando', 'testando@email.com', '$2a$10$Gh0/peL365JiNGKjL1PHVun5tHU0Ed9.5v5eHDybgr3T33VTaxWqG');
