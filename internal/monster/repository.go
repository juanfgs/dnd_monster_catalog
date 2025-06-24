package monster

import(
	"github.com/juanfgs/dnd-monster-library/internal/stats"
	"database/sql"
	"context"
)

type Repository interface {
	Index(ctx context.Context) ([]Monster, error)
	Create(ctx context.Context, m *Monster) error
	FindByChallengeRating(ctx context.Context, minChallengeRating float64, maxChallengeRating float64, quantity int64 ) ([]Monster, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

// TODO: Implement pagination
func (r *repository) Index(ctx context.Context) ([]Monster, error) {
	q := `SELECT * FROM monsters AS m LEFT JOIN stats AS s ON(s.monster_id = m.id)`  
	return r.buildResponse(ctx, q, []any{})
}

func (r *repository) Create(ctx context.Context, m *Monster) error {
	q := `INSERT INTO monsters (index, name, size, alignment, hit_points, hit_dice, hit_points_roll, languages, challenge_rating, proficiency_bonus, xp) VALUES ($1, $2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING id`  
	err := r.db.QueryRowContext(ctx, q,
		m.Index,
		m.Name,
		m.Size,
		m.Alignment,
		m.HitPoints,
		m.HitDice,
		m.HitPointsRoll,
		m.Languages,
		m.ChallengeRating,
		m.ProficiencyBonus,
		m.XP,
	).Scan(&m.ID)
	return err
}

func (r *repository) FindByChallengeRating(ctx context.Context, minChallengeRating float64, maxChallengeRating float64, quantity int64) ([]Monster, error) {
	params := []any{minChallengeRating, maxChallengeRating, quantity}
	q := `SELECT * FROM monsters AS m LEFT JOIN stats AS s ON(s.monster_id = m.id) WHERE challenge_rating >= $1 AND challenge_rating <= $2 ORDER BY RANDOM() LIMIT $3`   
	return r.buildResponse(ctx, q, params)
}


func (r *repository) buildResponse(ctx context.Context, q string, params []any ) ([]Monster, error) {
	monsters := make([]Monster, 0)
	rows, err := r.db.QueryContext(ctx, q, params...)
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
		&m.ChallengeRating,
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

