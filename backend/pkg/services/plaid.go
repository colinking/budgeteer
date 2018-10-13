package services

import (
	"context"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"

	"github.com/colinking/budgeteer/backend/pkg/db"
	proto "github.com/colinking/budgeteer/backend/pkg/proto/plaid"
	"github.com/colinking/budgeteer/backend/pkg/services/utils"
	"github.com/plaid/plaid-go/plaid"
)

type PlaidService struct {
	client *plaid.Client
	db     db.Database
}

// Validate implementation of proto interface.
var _ proto.PlaidServer = &PlaidService{}

// New returns a new instance of a plaid service client.
func New(clientID, publicKey, secret, env string) *PlaidService {
	client, err := plaid.NewClient(plaid.ClientOptions{
		ClientID:    clientID,
		Secret:      secret,
		PublicKey:   publicKey,
		Environment: utils.ToEnvironment(env),
		HTTPClient:  &http.Client{},
	})

	if err != nil {
		grpclog.Fatalf("failed to initialize Plaid client: %v", err)
	}

	return &PlaidService{
		client: client,
		db:     db.New(),
	}
}

// ExchangeToken converts a public token into an access token and stores it in the DB.
func (s *PlaidService) ExchangeToken(ctx context.Context, in *proto.ExchangeTokenRequest) (*proto.ExchangeTokenResponse, error) {
	res, err := s.client.ExchangePublicToken(in.Token)
	if err != nil {
		grpclog.Error(err)
		return nil, grpc.Errorf(codes.InvalidArgument, "token could not be converted to an access token")
	}

	grpclog.Infof("Exchanged public token (%s) for access token (%s)", in.Token, res.AccessToken)
	s.db.SaveToken(res.AccessToken)

	return &proto.ExchangeTokenResponse{
		AccessToken: res.AccessToken,
		ItemId:      res.ItemID,
	}, nil
}

// GetTransactions gets the transactions for a given user.
func (s *PlaidService) GetTransactions(ctx context.Context, in *proto.GetTransactionsRequest) (*proto.GetTransactionsResponse, error) {
	startDate, endDate := "2018-08-01", "2018-08-31"

	accessToken := s.db.GetToken()

	res, err := s.client.GetTransactions(accessToken, startDate, endDate)
	if err != nil {
		grpclog.Error(err)
		return nil, grpc.Errorf(codes.NotFound, "could not fetch transactions")
	}

	transactions := []*proto.Transaction{}
	for _, tx := range res.Transactions {
		transactions = append(transactions, utils.ToTransaction(tx))
	}

	grpclog.Infof("Found %d products for access token: %s\n", res.TotalTransactions, accessToken)

	return &proto.GetTransactionsResponse{
		Transactions: transactions,
	}, nil
}
