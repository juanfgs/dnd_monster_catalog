-- Create Monsters table
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS monsters(
       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
       index  VARCHAR (255) UNIQUE NOT NULL,
       name  VARCHAR (255)  NOT NULL,
       size  VARCHAR (255)  NOT NULL,
       alignment VARCHAR (255)  NOT NULL,
       hit_points INTEGER NOT NULL,
       hit_dice VARCHAR(255) NOT NULL,
       hit_points_roll VARCHAR(255) NOT NULL,
       languages VARCHAR(255) NOT NULL,
       challenge_rating DOUBLE PRECISION,
       proficiency_bonus INTEGER NOT NULL,
       xp INTEGER NOT NULL
);
