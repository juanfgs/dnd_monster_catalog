-- Create Stats table
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS stats(
       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
       strength INTEGER NOT NULL,
       dexterity INTEGER NOT NULL,
       constitution INTEGER NOT NULL,
       intelligence INTEGER NOT NULL,
       wisdom INTEGER NOT NULL,
       charisma INTEGER NOT NULL,
       monster_id UUID NOT NULL REFERENCES monsters(id) ON DELETE CASCADE
);
