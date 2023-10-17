-- DDL
CREATE DATABASE p2_ungraded_2;

CREATE TABLE IF NOT EXISTS Universe (
    ID INT AUTO_INCREMENT NOT NULL,
    Name VARCHAR(100) NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TABLE IF NOT EXISTS Heroes (
    ID INT AUTO_INCREMENT NOT NULL,
    Name VARCHAR(100) NOT NULL,
    Universe_id INT NOT NULL,
    Skill VARCHAR(255) NOT NULL,
    ImageURL VARCHAR(255) NOT NULL,
    PRIMARY KEY (ID),
    FOREIGN KEY (Universe_id) REFERENCES Universe(ID)
);

CREATE TABLE IF NOT EXISTS Villains (
    ID INT AUTO_INCREMENT NOT NULL,
    Name VARCHAR(100) NOT NULL,
    Universe_id INT NOT NULL,
    ImageURL VARCHAR(255) NOT NULL,
    PRIMARY KEY (ID),
    FOREIGN KEY (Universe_id) REFERENCES Universe(ID)
);

-- DML
INSERT INTO Universe (Name) VALUES
("Earth"), ("Galaxy"), ("Earth-616");

INSERT INTO Heroes (Name, Universe_id, Skill, ImageURL) VALUES
("Doctor Strage", 1, "Manipulate time", "https://static.wikia.nocookie.net/disney/images/d/dc/Doctor_Strange_-_Profile.png/revision/latest?cb=20220804200852"),
("Hulk", 1, "Turning into a giant", "https://www.greenscene.co.id/wp-content/uploads/2022/09/Hulk-696x497.jpg"),
("Star-Lord", 2, "Intelligence", "https://d1tgyzt3mf06m9.cloudfront.net/v3-staging/2023/05/fakta-star-lord-film-guardians-of-the-galaxy-vol-3db70a29f-1e3a-41e9-8255-915a6360f829.jpg"),
("Spider Man", 1, "Spider webs", "https://image.cnbcfm.com/api/v1/image/104375204-spideyheader.jpg?v=1529474677");

INSERT INTO Villains (Name, Universe_id, ImageURL) VALUES
("Dormammu", 2, "https://static.wikia.nocookie.net/vsbattles/images/e/ee/MCU_Dormammu-Guidebook_to_the_marvel_cinematic_universe.jpg/revision/latest?cb=20180905010811"),
("Abomination", 3, "https://static.wikia.nocookie.net/marveldatabase/images/b/b4/Emil_Blonsky_%28Earth-616%29_from_Hulk_Vol_3_2_cover.jpg/revision/latest?cb=20220905112355"),
("Ronan the Accusser", 2, "https://static1.cbrimages.com/wordpress/wp-content/uploads/2018/10/Lee-Pace-as-Ronan-e1552616857776.jpg"),
("Green Goblin", 1, "https://cdn.vox-cdn.com/thumbor/lq7SOeEhvUDEQB5sS-oR93UE02E=/0x0:4577x2412/1200x800/filters:focal(1785x473:2517x1205)/cdn.vox-cdn.com/uploads/chorus_image/image/70289439/nhw_still_120.0.jpg");