-- Create Monsters table

CREATE TABLE IF NOT EXISTS monsters(
       index  VARCHAR (255) UNIQUE NOT NULL,
       name  VARCHAR (255) UNIQUE NOT NULL,
       size  VARCHAR (255) UNIQUE NOT NULL,
       alignment VARCHAR (255) UNIQUE NOT NULL,
       hit_points INTEGER NOT NULL,
       hit_dice VARCHAR(255) NOT NULL,
       hit_points_roll VARCHAR(255) NOT NULL,
       languages VARCHAR(255) NOT NULL,
       proficiency_bonus INTEGER NOT NULL,
       xp INTEGER NOT NULL
);
