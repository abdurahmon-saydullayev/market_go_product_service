package service

import (
	"context"
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

type ProductService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*product_service.UnimplementedProductServiceServer
}

func NewProductService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ProductService {
	return &ProductService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *ProductService) Create(ctx context.Context, req *product_service.CreateProduct) (resp *product_service.Product, err error) {

	i.log.Info("---CreateProduct------>", logger.Any("req", req))

	pKey, err := i.strg.Product().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateProduct->Product->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Product().GetByID(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyProduct->Product->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ProductService) GetByID(ctx context.Context, req *product_service.ProductPK) (resp *product_service.Product, err error) {

	i.log.Info("---GetProductByID------>", logger.Any("req", req))

	resp, err = i.strg.Product().GetByID(ctx, req)
	if err != nil {
		i.log.Error("!!!GetProductByID->Product->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ProductService) GetList(ctx context.Context, req *product_service.GetListProductRequest) (resp *product_service.GetListProductResponse, err error) {

	i.log.Info("---GetProducts------>", logger.Any("req", req))

	resp, err = i.strg.Product().GetList(ctx, req)
	if err != nil {
		i.log.Error("!!!GetProducts->Product->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ProductService) Update(ctx context.Context, req *product_service.UpdateProduct) (resp *product_service.Product, err error) {

	i.log.Info("---UpdateProduct------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Product().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateProduct--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Product().GetByID(ctx, &product_service.ProductPK{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetProduct->Product->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *ProductService) UpdatePatch(ctx context.Context, req *product_service.UpdatePatchProduct) (resp *product_service.Product, err error) {

	i.log.Info("---UpdatePatchProduct------>", logger.Any("req", req))

	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.Product().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchProduct--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Product().GetByID(ctx, &product_service.ProductPK{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetProduct->Product->Get--->", logger.Error(err))

		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *ProductService) Delete(ctx context.Context, req *product_service.ProductPK) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteProduct------>", logger.Any("req", req))

	err = i.strg.Product().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteProduct->Product->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &empty.Empty{}, nil
}
