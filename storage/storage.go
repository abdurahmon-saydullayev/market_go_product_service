package storage

import (
	"context"
	"product_service/genproto/product_service"
	"product_service/models"
)

type StorageI interface {
	CloseDB()
	Category() CategoryRepoI
	Product() ProductRepoI
}

type ProductRepoI interface {
	Create(context.Context, *product_service.CreateProduct) (*product_service.ProductPK, error)
	GetByID(context.Context, *product_service.ProductPK) (*product_service.Product, error)
	GetList(context.Context, *product_service.GetListProductRequest) (*product_service.GetListProductResponse, error)
	Update(context.Context, *product_service.UpdateProduct) (int64, error)
	UpdatePatch(ctx context.Context, req *models.UpdatePatchRequest) (resp int64, err error)
	Delete(context.Context, *product_service.ProductPK) error
}

type CategoryRepoI interface {
	Create(context.Context, *product_service.CreateCategory) (*product_service.CategoryPK, error)
	GetByID(context.Context, *product_service.CategoryPK) (*product_service.Category, error)
	GetList(context.Context, *product_service.GetListCategoryRequest) (*product_service.GetListCategoryResponse, error)
	Update(context.Context, *product_service.UpdateCategory) (int64, error)
	UpdatePatch(ctx context.Context, req *models.UpdatePatchRequest) (resp int64, err error)
	Delete(context.Context, *product_service.CategoryPK) error
}
