package services

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	proto "github.com/colinking/budgeteer/backend/pkg/proto/budgeteer"
)

type budgetService struct{}

func NewBudgetService() *budgetService {
	return &budgetService{}
}

var purchases = []*proto.Purchase{
	{
		Id:     "1",
		Amount: 10.0,
	},
}

// Validate implementation of interface
var _ proto.BudgetServiceServer = &budgetService{}

func (s *budgetService) GetPurchase(ctx context.Context, in *proto.GetPurchasesRequest) (*proto.Purchase, error) {
	for _, purchase := range purchases {
		if purchase.Id == in.Id {
			return purchase, nil
		}
	}

	return nil, grpc.Errorf(codes.NotFound, "purchase could not be found")
}
