package repository

import (
	"context"
	"github.com/Cloude77/mailing-system/api-service/internal/model"
)

func (p *Postgres) CreateMailing(ctx context.Context, m *model.Mailing) error {
	query := `
        INSERT INTO mailings (message, group_name, send_time, periodicity)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at
    `
	err := p.Pool.QueryRow(ctx, query,
		m.Message, m.GroupName, m.SendTime, m.Periodicity,
	).Scan(&m.ID, &m.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
