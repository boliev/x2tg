CREATE TABLE IF NOT EXISTS sources(  
    id int NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    resource VARCHAR(25),
    url VARCHAR(255),
    is_active BOOLEAN DEFAULT false
);
 CREATE UNIQUE INDEX sources_url_unique ON sources (url);
 INSERT INTO sources (resource, url, is_active) VALUES
 ('reddit', 'https://www.reddit.com/r/golang/', true),
 ('reddit', 'https://www.reddit.com/r/StartledCats/', true),
 ('reddit', 'https://www.reddit.com/r/ProgrammerHumor/', true);