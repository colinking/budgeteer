package plaid

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"

	"github.com/colinking/budgeteer/backend/pkg/db"
	proto "github.com/colinking/budgeteer/backend/pkg/proto/plaid"
	"github.com/plaid/plaid-go/plaid"
)

// Service contains all Plaid-related handlers.
type Service struct {
	client *plaid.Client
	db     db.Database
}

// ServiceConfig specifies the configuration for a new Plaid Service.
type ServiceConfig struct {
	ClientID string
	PublicKey string
	Secret string
	Env string
	Database db.Database
}

// Validate implementation of proto interface.
var _ proto.PlaidServer = &Service{}

// New returns a new instance of a Plaid service client.
func New(c *ServiceConfig) *Service {
	client := newClient(c.ClientID, c.PublicKey, c.Secret, c.Env)

	return &Service{
		client: client,
		db:     c.Database,
	}
}

// ExchangeToken converts a public token into an access token and stores it in the DB.
func (s *Service) ExchangeToken(ctx context.Context, in *proto.ExchangeTokenRequest) (*proto.ExchangeTokenResponse, error) {
	res, err := s.client.ExchangePublicToken(in.Token)
	if err != nil {
		grpclog.Error(err)
		return nil, grpc.Errorf(codes.InvalidArgument, "token could not be converted to an access token")
	}

	grpclog.Infof("Exchanged public token (%s) for access token (%s)", in.Token, res.AccessToken)
	s.db.SaveToken("1234", res.AccessToken)

	return &proto.ExchangeTokenResponse{
		AccessToken: res.AccessToken,
		ItemId:      res.ItemID,
	}, nil
}

// GetTransactions gets the transactions for a given user.
func (s *Service) GetTransactions(ctx context.Context, in *proto.GetTransactionsRequest) (*proto.GetTransactionsResponse, error) {
	startDate, endDate := "2018-08-01", "2018-08-31"

	accessToken := s.db.GetToken("1234")

	res, err := s.client.GetTransactions(accessToken, startDate, endDate)
	if err != nil {
		grpclog.Error(err)
		return nil, grpc.Errorf(codes.NotFound, "could not fetch transactions")
	}

	transactions := []*proto.Transaction{}
	for _, tx := range res.Transactions {
		transactions = append(transactions, toTransaction(tx))
	}

	grpclog.Infof("Found %d products for access token: %s\n", res.TotalTransactions, accessToken)

	return &proto.GetTransactionsResponse{
		Transactions: transactions,
	}, nil
}
