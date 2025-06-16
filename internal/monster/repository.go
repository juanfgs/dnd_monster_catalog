package monster

import(
	"context"
	"database/sql"
)

type Repository interface {
	Create(ctx context.Context, m *Monster) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, m *Monster) error {
	q := `INSERT INTO monsters (index, name, size, alignment, hit_points, hit_dice, hit_points_roll, languages, proficiency_bonus, xp) VALUES ($1, $2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id`  
	err := r.db.QueryRowContext(ctx, q,
		m.Index,
		m.Name,
		m.Size,
		m.Alignment,
		m.HitPoints,
		m.HitDice,
		m.HitPointsRoll,
		m.Languages,
		m.ProficiencyBonus,
		m.XP,
	).Scan(&m.ID)
	return err
}



