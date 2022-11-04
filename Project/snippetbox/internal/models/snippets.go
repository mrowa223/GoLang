package models

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *pgxpool.Pool
}

var ctx = context.Background()

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	// stmt := `insert into snippets (title, content, created, expires)
	// values($1, $2, now(), (now() + interval '1 DAY' * $3)) returning id`

	var id int
	query := `INSERT INTO snippets (title, content, created, expires)
	VALUES($1, $2, NOW(), NOW() + INTERVAL '1 DAY' * $3)
	RETURNING id`
	row := m.DB.QueryRow(ctx, query, title, content, expires)

	if err := row.Scan(&id); err != nil {
		return 0, nil
	}

	return id, nil
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	stmt := `select id, title, content, created, expires from snippets 
	where expires > now() and id = $1`

	row := m.DB.QueryRow(ctx, stmt, id)
	s := &Snippet{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	stmt := `select id, title, content, created, expires from snippets 
	where expires > now() order by id desc limit 10`

	rows, err := m.DB.Query(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	snippets := []*Snippet{}

	for rows.Next() {
		s := &Snippet{}
		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return snippets, nil
}
