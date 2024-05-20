CREATE EXTENSION "uuid-ossp";

CREATE TABLE mst_authors (
    id uuid PRIMARY KEY NOT NULL DEFAULT  uuid_generate_v4(),
    fullname varchar(50) NULL,
    email varchar(100) NULL UNIQUE,
    password varchar(100) NULL,
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    role varchar(100) NULL
);

CREATE TABLE trx_tasks (
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    title varchar(50) NULL,
    content text NULL ,
    author_id uuid NULL,
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP
 );

-- DML

INSERT INTO mst_authors (fullname, email, password, role) VALUES 
('Damar', 'JkSsA@example.com', 'secret', 'user'),
('Laras', 'LarasW@example.com', 'secret', 'user'),
('Andi', 'AndiS@example.com', 'secret', 'user'),
('Budi', 'BudiT@example.com', 'secret', 'user'),
('Citra', 'CitraM@example.com', 'secret', 'user'),
('Dewi', 'DewiK@example.com', 'secret', 'user'),
('Eka', 'EkaJ@example.com', 'secret', 'user'),
('Farah', 'FarahH@example.com', 'secret', 'user'),
('Gita', 'GitaF@example.com', 'secret', 'user'),
('Hadi', 'HadiR@example.com', 'secret', 'user'),
('Ika', 'IkaD@example.com', 'secret', 'user'),
('Joko', 'JokoN@example.com', 'secret', 'user'),
('Kiki', 'KikiP@example.com', 'secret', 'user'),
('Lina', 'LinaL@example.com', 'secret', 'user'),
('Maya', 'MayaG@example.com', 'secret', 'user'),
('Nina', 'NinaB@example.com', 'secret', 'user'),
('Oni', 'OniQ@example.com', 'secret', 'user'),
('Putri', 'PutriS@example.com', 'secret', 'user'),
('Rian', 'RianT@example.com', 'secret', 'user'),
('Sari', 'SariY@example.com', 'secret', 'user');