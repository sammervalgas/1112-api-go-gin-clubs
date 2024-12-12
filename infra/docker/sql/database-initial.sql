-- Criação da tabela
CREATE TABLE clubs (
    id SERIAL PRIMARY KEY,
    address VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    pid UUID NOT NULL,
    website VARCHAR(255) NOT NULL
);

-- Inserção de registros
INSERT INTO clubs (address, name, pid, website) VALUES
('7th Street', 'Hot Club', '3d01b30c-ad24-4f94-a945-26f4e2b95054', 'hotclubs.com.br'),
('8th Avenue', 'Jazz Lounge', 'd27c996b-3bd2-45e3-9016-f9b034e1a945', 'jazzlounge.com'),
('Main Street', 'Blues House', 'bc12f4a7-5289-48c1-b0e4-2d87b15cd84a', 'blueshouse.net'),
('Broadway', 'Rock Arena', 'f9d84a25-1ed5-4778-9a68-334e92c1c8c6', 'rockarena.org'),
('Central Park', 'Classical Hall', 'e5f2399d-ef2d-41a1-b682-439dff2b67a3', 'classicalhall.com'),
('Elm Street', 'Reggae Spot', '12e83bd2-4fb3-48fc-917a-e0ff3cd30bcd', 'reggaespot.io'),
('10th Boulevard', 'Dance Haven', 'a7b8f224-03a1-4925-96f5-e2a91b398d37', 'dancehaven.com'),
('Sunset Boulevard', 'Pop World', '748cd946-71a5-4c36-8b9e-6d8d2d446e8b', 'popworld.co'),
('Ocean Drive', 'Salsa Palace', 'f6c3b94b-74fc-44e9-a520-f98b9c4cb8c2', 'salsapalace.org'),
('Highland Road', 'Hip-Hop Hub', 'a45d3921-cf59-4b16-8219-321afbe5f613', 'hiphophub.com'),
('Park Avenue', 'Electronic Beats', 'b9e72d4f-1c35-4e89-a876-1d4f2c8e9b3a', 'electronicbeats.com'),
('River Street', 'Country Music Bar', 'c8f91e25-6d4a-4b37-9c52-3f8a7d9e4b5c', 'countrymusicbar.net');
