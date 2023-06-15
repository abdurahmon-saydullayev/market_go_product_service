package grpc

import (
	"product_service/config"
	"product_service/genproto/product_service"
	"product_service/grpc/client"
	"product_service/grpc/service"
	"product_service/pkg/logger"
	"product_service/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) (grpcServer *grpc.Server) {

	grpcServer = grpc.NewServer()

	product_service.RegisterProductServiceServer(grpcServer, service.NewProductService(cfg, log, strg, srvc))
	product_service.RegisterCategoryServiceServer(grpcServer, service.NewCategoryService(cfg, log, strg, srvc))

	reflection.Register(grpcServer)
	return
}
