package stats 

import(
	"context"
	"database/sql"
)

type Repository interface {
	Create(ctx context.Context, s *Stats, monsterID string) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, s *Stats, monsterID string) error {
	q := `INSERT INTO stats (strength, dexterity, constitution, intelligence, wisdom, charisma, monster_id ) VALUES ($1, $2,$3,$4,$5,$6, $7) RETURNING id`  
	err := r.db.QueryRowContext(ctx, q,
		s.Strength,
		s.Dexterity,
		s.Constitution,
		s.Intelligence,
		s.Wisdom,
		s.Charisma,
		monsterID,
	).Scan(&s.ID)
	return err
}
