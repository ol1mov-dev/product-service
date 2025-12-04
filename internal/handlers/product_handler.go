package handlers

import (
	"context"

	productV1 "github.com/ol1mov-dev/protos/pkg/product/v1"
)

type ProductServerHandler struct {
	productV1.UnimplementedProductV1ServiceServer
}

func (s *ProductServerHandler) CreateProduct(ctx context.Context, req *productV1.CreateProductRequest) (*productV1.Product, error) {
	return &productV1.Product{}, nil
}

func (s *ProductServerHandler) UpdateProduct(ctx context.Context, req *productV1.UpdateProductRequest) (*productV1.Product, error) {
	return &productV1.Product{}, nil
}

func (s *ProductServerHandler) DeleteProduct(ctx context.Context, req *productV1.DeleteProductRequest) (*productV1.Product, error) {
	return &productV1.Product{}, nil
}

func (s *ProductServerHandler) GetProduct(ctx context.Context, req *productV1.GetProductRequest) (*productV1.Product, error) {
	return &productV1.Product{}, nil
}
