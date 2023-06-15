package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"product_service/genproto/product_service"
	"product_service/models"
	"product_service/pkg/helper"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type categoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) *categoryRepo {
	return &categoryRepo{
		db: db,
	}
}

func (c *categoryRepo) Create(ctx context.Context, req *product_service.CreateCategory) (resp *product_service.CategoryPK, err error) {
	id := uuid.New().String()

	query := `
		INSERT INTO "category" (
			id,
			name,
			parent,
			created_at,
			updated_at
		) VALUES ($1, $2, $3, NOW(), NOW())
	`

	_, err = c.db.Exec(
		ctx,
		query,
		id,
		req.Name,
		req.Parent,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &product_service.CategoryPK{Id: id}, nil
}

func (c *categoryRepo) GetByID(ctx context.Context, req *product_service.CategoryPK) (order *product_service.Category, err error) {
	query := `
		SELECT 
		    id,
			name,
			parent,
			created_at,
			updated_at
		FROM "category"
		WHERE id = $1;
	`
	var (
		id         sql.NullString
		name       sql.NullString
		parent     sql.NullString
		created_at sql.NullString
		updated_at sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&parent,
		&created_at,
		&updated_at,
	)
	if err != nil {
		return order, err
	}

	order = &product_service.Category{
		Id:        id.String,
		Name:      name.String,
		Parent:    parent.String,
		CreatedAt: created_at.String,
		UpdatedAt: updated_at.String,
	}

	return
}

func (c *categoryRepo) GetList(ctx context.Context, req *product_service.GetListCategoryRequest) (resp *product_service.GetListCategoryResponse, err error) {
	resp = &product_service.GetListCategoryResponse{}

	var (
		query  string
		limit  = ""
		offset = " OFFSET 0 "
		params = make(map[string]interface{})
		filter = " WHERE TRUE "
		sort   = " ORDER BY created_at DESC "
	)

	query = `
	   SELECT 
	   		COUNT(*) OVER(),
			   id,
			   name,
			   parent,
			   created_at,
			   updated_at
		FROM "category"
	`
	if len(req.GetSearch()) > 0 {
		filter += " AND name ILIKE '%' || '" + req.Search + "' || '%' "
	}
	if req.GetLimit() > 0 {
		limit = " LIMIT :limit"
		params["limit"] = req.Limit
	}
	if req.GetOffset() > 0 {
		offset = " OFFSET :offset"
		params["offset"] = req.Offset
	}
	query += filter + sort + offset + limit

	query, args := helper.ReplaceQueryParams(query, params)

	rows, err := c.db.Query(ctx, query, args...)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id         sql.NullString
			name       sql.NullString
			parent     sql.NullString
			created_at sql.NullString
			updated_at sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&parent,
			&created_at,
			&updated_at,
		)
		if err != nil {
			return resp, err
		}

		resp.Categorys = append(resp.Categorys, &product_service.Category{
			Id:        id.String,
			Name:      name.String,
			Parent:    parent.String,
			CreatedAt: created_at.String,
			UpdatedAt: updated_at.String,
		})
	}

	return
}

func (c *categoryRepo) Update(ctx context.Context, req *product_service.UpdateCategory) (resp int64, err error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			"category"
		SET
			name = :name,
			parent= :parent,
			updated_at = now()
		WHERE id = :id
	`
	params = map[string]interface{}{
		"id":     req.GetId(),
		"name":   req.GetName(),
		"parent": req.GetParent(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *categoryRepo) UpdatePatch(ctx context.Context, req *models.UpdatePatchRequest) (resp int64, err error) {
	var (
		set   = " SET "
		ind   = 0
		query string
	)

	if len(req.Fields) == 0 {
		err = errors.New("no updates provided")
		return
	}

	req.Fields["id"] = req.Id

	for key := range req.Fields {
		set += fmt.Sprintf(" %s = :%s ", key, key)
		if ind != len(req.Fields)-1 {
			set += ", "
		}
		ind++
	}

	query = `
		UPDATE
			"category"
	` + set + ` , updated_at = now()
		WHERE
			id = :id
	`

	query, args := helper.ReplaceQueryParams(query, req.Fields)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), err
}

func (c *categoryRepo) Delete(ctx context.Context, req *product_service.CategoryPK) error {
	query := `DELETE FROM "category" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)
	if err != nil {
		return err
	}

	return nil
}
