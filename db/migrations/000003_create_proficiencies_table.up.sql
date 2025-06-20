-- Create proficiency tables
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS proficiencies(
       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
       name VARCHAR(255) UNIQUE NOT NULL,
       type VARCHAR(255) NOT NULL,
       UNIQUE (type, name)
);

CREATE TABLE IF NOT EXISTS monster_proficiency(
       proficiency_id UUID NOT NULL,
       monster_id UUID NOT NULL,
       value INTEGER NOT NULL,
       PRIMARY KEY (monster_id, proficiency_id),
       FOREIGN KEY (monster_id) REFERENCES monsters(id) ON DELETE CASCADE,
       FOREIGN KEY (proficiency_id) REFERENCES proficiencies(id) ON DELETE CASCADE
);
