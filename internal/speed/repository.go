package speed 

import(
	"context"
	"database/sql"
)

type Repository interface {
	Create(ctx context.Context, s *Speed) error
	Associate(ctx context.Context, speedID string, monsterID string, value int64, unit string) error
	Fetch(ctx context.Context, monsterID string) ([]Speed, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, s *Speed) error {
	q := `WITH ins AS (
                  INSERT INTO speeds ( type )
                         VALUES ($1 ) ON CONFLICT ( type) DO NOTHING RETURNING id
                  ) SELECT id FROM ins
              UNION SELECT id FROM speeds WHERE  type = $1 LIMIT 1;`  
	err := r.db.QueryRowContext(ctx, q, s.Type).Scan(&s.ID)
	return err
}

func (r *repository) Associate(ctx context.Context, speedID string, monsterID string, value int64, unit string) error {
	q := `INSERT INTO monster_speed (speed_id,monster_id, value, unit  ) VALUES ($1,$2,$3,$4)`  
	_, err := r.db.ExecContext(ctx, q, speedID,monsterID, value, unit)
	return err
}

func (r *repository) Fetch(ctx context.Context, monsterID string) ([]Speed, error) {
	speeds := make([]Speed, 0)

	q := `SELECT s.id, s.type, ms.value,ms.unit FROM monster_speed AS ms INNER JOIN speeds AS s ON (ms.speed_id = s.id) WHERE ms.monster_id = $1`  
	rows, err := r.db.QueryContext(ctx, q, monsterID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var speed Speed 
		if err = scanColumns(rows, &speed); err != nil {
			return nil, err
		}
		speeds = append(speeds, speed)
	}
	// Check for errors from iterating over rows.
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return speeds, nil
}

func scanColumns(rows *sql.Rows, s *Speed) error{
	err := rows.Scan(
		&s.ID,
		&s.Type,
		&s.Value,
		&s.Unit,
	)
	return err
}

