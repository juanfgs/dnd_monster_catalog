-- Create proficiency tables
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS armor_classes(
       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
       type VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS monster_armor_class(
       armor_class_id UUID NOT NULL,
       monster_id UUID NOT NULL,
       value INTEGER NOT NULL,
       PRIMARY KEY (monster_id, armor_class_id),
       FOREIGN KEY (monster_id) REFERENCES monsters(id) ON DELETE CASCADE,
       FOREIGN KEY (armor_class_id) REFERENCES armor_classes(id) ON DELETE CASCADE
);
