package proficiency 

import(
	"context"
	"database/sql"
)

type Repository interface {
	Create(ctx context.Context, s *Proficiency) error
	Associate(ctx context.Context, proficiencyID string, monsterID string, value int64) error
	Fetch(ctx context.Context, monsterID string) ([]Proficiency, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, p *Proficiency) error {
	q := `WITH ins AS (
                  INSERT INTO proficiencies (name, type )
                         VALUES ($1, $2) ON CONFLICT (name, type) DO NOTHING RETURNING id
                  ) SELECT id FROM ins
              UNION SELECT id FROM proficiencies WHERE name = $1 AND type = $2 LIMIT 1;`  
	err := r.db.QueryRowContext(ctx, q, p.Name, p.Type).Scan(&p.ID)
	return err
}

func (r *repository) Associate(ctx context.Context, proficiencyID string, monsterID string, value int64) error {
	q := `INSERT INTO monster_proficiency (proficiency_id,monster_id, value  ) VALUES ($1,$2,$3)`  
	_, err := r.db.ExecContext(ctx, q, proficiencyID,monsterID, value)
	return err
}

func (r *repository) Fetch(ctx context.Context, monsterID string) ([]Proficiency, error) {
	proficiencies := make([]Proficiency, 0)

	q := `SELECT p.id, p.name, p.type, mp.value FROM monster_proficiency AS mp INNER JOIN proficiencies AS p ON (mp.proficiency_id = p.id) WHERE mp.monster_id = $1`  
	rows, err := r.db.QueryContext(ctx, q, monsterID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var proficiency Proficiency 
		if err = scanColumns(rows, &proficiency); err != nil {
			return nil, err
		}
		proficiencies = append(proficiencies, proficiency)
	}
	// Check for errors from iterating over rows.
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return proficiencies, nil
}

func scanColumns(rows *sql.Rows, p *Proficiency) error{
	err := rows.Scan(
		&p.ID,
		&p.Name,
		&p.Type,
		&p.Value,
	)
	return err
}

