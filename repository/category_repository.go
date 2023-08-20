package repository

import (
	"context"
	"database/sql"
	"golang-restful-api/model/domain"
)

type CategoryRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindByID(ctx context.Context, tx *sql.Tx, categoryID int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
