CREATE TABLE IF NOT EXISTS sources(  
    id int NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    resource VARCHAR(25),
    url VARCHAR(255),
    is_active BOOLEAN DEFAULT false
);
CREATE UNIQUE INDEX sources_url_unique ON sources (url);
INSERT INTO sources (resource, url, is_active) 
OVERRIDING SYSTEM VALUE VALUES
('reddit', 'https://www.reddit.com/r/golang/', true),
('reddit', 'https://www.reddit.com/r/StartledCats/', true),
('reddit', 'https://www.reddit.com/r/ProgrammerHumor/', true),
('reddit', 'https://www.reddit.com/r/CozyPlaces/', true);

CREATE TABLE IF NOT EXISTS channels(  
    id int NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    tg_id VARCHAR(25),
    name VARCHAR(250)
);

CREATE TABLE IF NOT EXISTS sources_channels(  
    source_id int,
    channel_id int,
    PRIMARY KEY(source_id, channel_id)
);

INSERT INTO channels (tg_id, name) VALUES
('-1002131623767','Reddit tech news'),
('-1002138236338','Prokrastinators')

INSERT INTO sources_channels (source_id, channel_id) VALUES
(1,1),
(2,2),
(3,1),
(4,2),
(4,1);



SELECT s.id, s.resource, s.url, s.is_active, c.name
FROM sources s
INNER JOIN sources_channels sc ON sc.source_id = s.id
INNER JOIN channels c ON c.id = sc.channel_id
WHERE s.is_active = true