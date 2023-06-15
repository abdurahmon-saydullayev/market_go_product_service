package service

import (
	"context"
	"fmt"
	"product_service/config"
	"product_service/genproto/product_service"
	"product_service/grpc/client"
	"product_service/models"
	"product_service/pkg/logger"
	"product_service/storage"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CategoryService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*product_service.UnimplementedCategoryServiceServer
}

func NewCategoryService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *CategoryService {
	return &CategoryService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *CategoryService) Create(ctx context.Context, req *product_service.CreateCategory) (resp *product_service.Category, err error) {

	i.log.Info("---CreateCategory------>", logger.Any("req", req))

	pKey, err := i.strg.Category().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateCategory->Category->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Category().GetByID(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyCategory->Category->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *CategoryService) GetByID(ctx context.Context, req *product_service.CategoryPK) (resp *product_service.Category, err error) {

	i.log.Info("---GetCategoryByID------>", logger.Any("req", req))

	resp, err = i.strg.Category().GetByID(ctx, req)
	if err != nil {
		i.log.Error("!!!GetCategoryByID->Category->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *CategoryService) GetList(ctx context.Context, req *product_service.GetListCategoryRequest) (resp *product_service.GetListCategoryResponse, err error) {

	i.log.Info("---GetCategorys------>", logger.Any("req", req))
	fmt.Println("slkd")
	resp, err = i.strg.Category().GetList(ctx, req)
	if err != nil {
		i.log.Error("!!!GetCategorys->Category->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *CategoryService) Update(ctx context.Context, req *product_service.UpdateCategory) (resp *product_service.Category, err error) {

	i.log.Info("---UpdateCategory------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Category().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateCategory--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Category().GetByID(ctx, &product_service.CategoryPK{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetCategory->Category->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *CategoryService) UpdatePatch(ctx context.Context, req *product_service.UpdatePatchCategory) (resp *product_service.Category, err error) {

	i.log.Info("---UpdatePatchCategory------>", logger.Any("req", req))

	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.Category().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchCategory--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Category().GetByID(ctx, &product_service.CategoryPK{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetCategory->Category->Get--->", logger.Error(err))

		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *CategoryService) Delete(ctx context.Context, req *product_service.CategoryPK) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteCategory------>", logger.Any("req", req))

	err = i.strg.Category().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteCategory->Category->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &empty.Empty{}, nil
}
