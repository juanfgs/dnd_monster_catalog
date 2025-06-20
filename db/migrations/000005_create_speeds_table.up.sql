-- Create proficiency tables
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS speeds(
       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
       type VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS monster_speed(
       speed_id UUID NOT NULL,
       monster_id UUID NOT NULL,
       value INTEGER NOT NULL,
       unit VARCHAR(255) NOT NULL,
       PRIMARY KEY (monster_id, speed_id),
       FOREIGN KEY (monster_id) REFERENCES monsters(id) ON DELETE CASCADE,
       FOREIGN KEY (speed_id) REFERENCES speeds(id) ON DELETE CASCADE
);
