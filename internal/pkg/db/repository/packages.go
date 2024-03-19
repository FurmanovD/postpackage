package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/FurmanovD/postpackage/internal/pkg/db/model"
)

type packageRepositoryImpl struct {
	db *sqlx.DB
}

// ================= interface instantiation =================================================

func NewPackagesRepository(db *sqlx.DB) PackagesRepository {
	return &packageRepositoryImpl{
		db: db,
	}
}

// ================= interface methods =================================================
func (r *packageRepositoryImpl) Insert(ctx context.Context, pkg model.Package) (*model.Package, error) {
	pkg.CreatedAt = time.Now() // NOTE(DF): add fnTimeNow dependency instead
	pkg.ID = 0

	res, err := r.db.NamedExecContext(ctx,
		`INSERT INTO packages (id, name, items_per_package, created_at)
		VALUES (:id, :name, :items_per_package, :created_at);`,
		pkg,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
		}

		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	resPkg, err := r.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return resPkg, nil
}

func (r *packageRepositoryImpl) Get(ctx context.Context, id int64) (*model.Package, error) {
	res := &model.Package{}

	err := r.db.GetContext(ctx, res, "SELECT * FROM packages WHERE id = ? AND deleted_at IS NULL;", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
		}

		return nil, err
	}

	return res, nil
}

func (r *packageRepositoryImpl) List(ctx context.Context) ([]model.Package, error) {
	// TODO(DF): add pagination
	res := make([]model.Package, 0)

	err := r.db.SelectContext(ctx, &res, "SELECT * FROM packages;")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
		}

		return nil, err
	}

	return res, nil
}

func (r *packageRepositoryImpl) Update(ctx context.Context, pkg model.Package) (*model.Package, error) {
	updatedAt := time.Now() // NOTE(DF): add fnTimeNow dependency instead
	pkg.UpdatedAt = &updatedAt

	_, err := r.db.NamedExecContext(ctx,
		`UPDATE packages SET name=:name, items_per_package=:items_per_package, updated_at=:updated_at WHERE id = :id;`,
		pkg,
	)
	if err != nil {
		return nil, err
	}

	resPkg, err := r.Get(ctx, pkg.ID)
	if err != nil {
		return nil, err
	}

	return resPkg, nil
}

func (r *packageRepositoryImpl) Delete(ctx context.Context, id int64) error {
	deletedAt := time.Now() // NOTE(DF): add fnTimeNow dependency instead

	args := map[string]interface{}{
		"deleted_at": deletedAt,
		"id":         id,
	}

	_, err := r.db.NamedExecContext(ctx,
		`UPDATE packages SET deleted_at=:deleted_at WHERE deleted_at IS NULL AND id = :id;`,
		args,
	)
	if err != nil {
		return nil
	}

	return nil
}
