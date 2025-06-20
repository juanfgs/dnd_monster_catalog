package armor 

import(
	"context"
	"database/sql"
)

type Repository interface {
	Create(ctx context.Context, ac *ArmorClass) error
	Associate(ctx context.Context, proficiencyID string, monsterID string, value int64) error
	Fetch(ctx context.Context, monsterID string) ([]ArmorClass, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, ac *ArmorClass) error {
	q := `WITH ins AS (
                  INSERT INTO armor_classes ( type )
                         VALUES ($1 ) ON CONFLICT ( type) DO NOTHING RETURNING id
                  ) SELECT id FROM ins
              UNION SELECT id FROM armor_classes WHERE  type = $1 LIMIT 1;`  
	err := r.db.QueryRowContext(ctx, q, ac.Type).Scan(&ac.ID)
	return err
}

func (r *repository) Associate(ctx context.Context, armorClassID string, monsterID string, value int64) error {
	q := `INSERT INTO monster_armor_class (armor_class_id,monster_id, value  ) VALUES ($1,$2,$3)`  
	_, err := r.db.ExecContext(ctx, q, armorClassID,monsterID, value)
	return err
}

func (r *repository) Fetch(ctx context.Context, monsterID string) ([]ArmorClass, error) {
	armorClasses := make([]ArmorClass, 0)

	q := `SELECT ac.id, ac.type, mac.value FROM monster_armor_class AS mac INNER JOIN armor_classes AS ac ON (mac.armor_class_id = ac.id) WHERE mac.monster_id = $1`  
	rows, err := r.db.QueryContext(ctx, q, monsterID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var armorClass ArmorClass 
		if err = scanColumns(rows, &armorClass); err != nil {
			return nil, err
		}
		armorClasses = append(armorClasses, armorClass)
	}
	// Check for errors from iterating over rows.
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return armorClasses, nil
}

func scanColumns(rows *sql.Rows, ac *ArmorClass) error{
	err := rows.Scan(
		&ac.ID,
		&ac.Type,
		&ac.Value,
	)
	return err
}

