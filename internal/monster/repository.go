package monster

import(
	"github.com/juanfgs/dnd-monster-library/internal/stats"
	"database/sql"
	"context"
)

type Repository interface {
	Index(ctx context.Context) ([]Monster, error)
	Create(ctx context.Context, m *Monster) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

// TODO: Implement pagination
func (r *repository) Index(ctx context.Context) ([]Monster, error) {
	monsters := make([]Monster, 0)
	q := `SELECT * FROM monsters AS m LEFT JOIN stats AS s ON(s.monster_id = m.id)`  
	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var monster Monster
		if err = scanColumns(rows, &monster); err != nil {
			return nil, err
		}
		monsters = append(monsters, monster)
	}
	// Check for errors from iterating over rows.
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return monsters, nil
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

func scanColumns(rows *sql.Rows, m *Monster) error{
	var s stats.Stats 
	err := rows.Scan(
		&m.ID,
		&m.Index,
		&m.Name,
		&m.Size,
		&m.Alignment,
		&m.HitPoints,
		&m.HitDice,
		&m.HitPointsRoll,
		&m.Languages,
		&m.ProficiencyBonus,
		&m.XP,
		&s.ID,
		&s.Strength,
		&s.Dexterity,
		&s.Constitution,
		&s.Intelligence,
		&s.Wisdom,
		&s.Charisma,
		&s.MonsterID,
	)
	m.Stats = &s
	return err
}

