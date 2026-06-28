package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/sameer2006-s/grpc-url-shortner/internal/model"
)

type PostgresRepository struct {
	db *pgx.Conn
}

type LinkRepository interface {
    Save(link model.Link) error

    Get(code string) (model.Link,bool)

    IncrementClicks(code string,) error
}

func NewPostgresRepository(db *pgx.Conn) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Save(link model.Link) error {
	_, err := r.db.Exec(context.Background(), `
INSERT INTO links
(
	id,
	short_code,
	url,
	clicks,
	created_at
)
VALUES
(
	$1,
	$2,
	$3,
	$4,
	$5
)
`, uuid.New(), link.ShortCode, link.URL , link.Clicks, link.CreatedAt)

	return err
}

func (r *PostgresRepository) Get(code string) (model.Link, bool) {
	var link model.Link

	err := r.db.QueryRow(context.Background(), `
SELECT
short_code,
url,
clicks,
created_at
FROM links
WHERE short_code=$1
`, code).Scan(&link.ShortCode, &link.URL, &link.Clicks, &link.CreatedAt)

	if err != nil {
		return model.Link{}, false
	}

	return link, true
}

func (r *PostgresRepository) IncrementClicks(code string,) error {
    _, err :=
        r.db.Exec(
            context.Background(),
            `
UPDATE links
SET clicks = clicks + 1
WHERE short_code=$1
`,
            code,
        )

    return err
}