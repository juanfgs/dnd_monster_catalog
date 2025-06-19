package proficiency 

import(
	"context"
	"database/sql"
)

type Repository interface {
	Create(ctx context.Context, s *Proficiency) error
	Associate(ctx context.Context, proficiencyID string, monsterID string, value int64) 
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, p *Proficiency) error {
	q := `INSERT INTO proficiencies (name ) VALUES ($1) ON CONFLICT (name) DO NOTHING RETURNING id`  
	err := r.db.QueryRowContext(ctx, q, p.Name).Scan(&p.ID)
	return err
}

func (r *repository) Associate(ctx context.Context, proficiencyID string, monsterID string, value int64) {
	q := `INSERT INTO monster_proficiency (proficiency_id,monster_id, value  ) VALUES ($1,$2,$3) RETURNING id`  
	r.db.QueryContext(ctx, q, proficiencyID,monsterID, value)
}
