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

type productRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *productRepo {
	return &productRepo{
		db: db,
	}
}

func (c *productRepo) Create(ctx context.Context, req *product_service.CreateProduct) (resp *product_service.ProductPK, err error) {
	id := uuid.New().String()

	query := `
		INSERT INTO "product" (
			id,
			photo,
			name,
			category_id,
			barcode,
			price,
			created_at,
			updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
	`

	_, err = c.db.Exec(
		ctx,
		query,
		id,
		req.Photo,
		req.Name,
		req.CategoryId,
		req.Barcode,
		req.Price,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &product_service.ProductPK{Id: id}, nil
}

func (c *productRepo) GetByID(ctx context.Context, req *product_service.ProductPK) (order *product_service.Product, err error) {
	query := `
		SELECT 
			id,
			photo,
			name,
			category_id,
			barcode,
			price,
			created_at,
			updated_at
		FROM "product"
		WHERE id = $1;
	`
	var (
		id          sql.NullString
		photo       sql.NullString
		name        sql.NullString
		category_id sql.NullString
		barcode     sql.NullString
		price       sql.NullFloat64
		created_at  sql.NullString
		updated_at  sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&photo,
		&name,
		&category_id,
		&barcode,
		&price,
		&created_at,
		&updated_at,
	)
	if err != nil {
		return order, err
	}

	order = &product_service.Product{
		Id:         id.String,
		Photo:      photo.String,
		Name:       name.String,
		CategoryId: category_id.String,
		Barcode:    barcode.String,
		Price:      float32(price.Float64),
		CreatedAt:  created_at.String,
		UpdatedAt:  updated_at.String,
	}

	return
}

func (c *productRepo) GetList(ctx context.Context, req *product_service.GetListProductRequest) (resp *product_service.GetListProductResponse, err error) {
	resp = &product_service.GetListProductResponse{}

	var (
		query  string
		limit  = ""
		offset = " OFFSET 0 "
		params = make(map[string]interface{})
		filter = " WHERE TRUE "
		sort   = " ORDER BY created_at DESC"
	)

	query = `
	   SELECT 
	   		COUNT(*) OVER(),
			    id,
			    photo,
			    name,
			    category_id,
			    barcode,
			    price,
			    created_at,
			    updated_at
		FROM "product"
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
			id          sql.NullString
			photo       sql.NullString
			name        sql.NullString
			category_id sql.NullString
			barcode     sql.NullString
			price       sql.NullFloat64
			created_at  sql.NullString
			updated_at  sql.NullString
		)

		err := rows.Scan(
			&id,
			&photo,
			&name,
			&category_id,
			&barcode,
			&price,
			&created_at,
			&updated_at,
		)
		if err != nil {
			return resp, err
		}

		resp.Products = append(resp.Products, &product_service.Product{
			Id:         id.String,
			Photo:      photo.String,
			Name:       name.String,
			CategoryId: category_id.String,
			Barcode:    barcode.String,
			Price:      float32(price.Float64),
			CreatedAt:  created_at.String,
			UpdatedAt:  updated_at.String,
		})
	}

	return
}

func (c *productRepo) Update(ctx context.Context, req *product_service.UpdateProduct) (resp int64, err error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			"product"
		SET
			photo = :photo,
			name= :name,
			category_id = :category_id,
			barcode = :barcode,
			price = :price,
			updated_at = now()
		WHERE id = :id
	`
	params = map[string]interface{}{
		"id":          req.GetId(),
		"photo":       req.GetPhoto(),
		"name":        req.GetName(),
		"category_id": req.GetCategoryId(),
		"barcode":     req.GetBarcode(),
		"price":       req.GetPrice(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *productRepo) UpdatePatch(ctx context.Context, req *models.UpdatePatchRequest) (resp int64, err error) {
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
			"product"
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

func (c *productRepo) Delete(ctx context.Context, req *product_service.ProductPK) error {
	query := `DELETE FROM "product" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)
	if err != nil {
		return err
	}

	return nil
}
