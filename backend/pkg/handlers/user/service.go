package user

import (
	"github.com/colinking/budgeteer/backend/pkg/auth"
	"github.com/colinking/budgeteer/backend/pkg/db"
	"github.com/colinking/budgeteer/backend/pkg/gen/userpb"
	plaidgo "github.com/plaid/plaid-go/plaid"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

// Service contains all User-related handlers.
type Service struct {
	db     db.Database
	client *plaidgo.Client
}

func (s Service) Get(ctx context.Context, req *userpb.GetRequest) (*userpb.GetResponse, error) {
	authID, err := auth.GetAuthId(ctx)
	if err != nil {
		return nil, err
	}

	out := s.db.GetUser(&db.GetUserInput{
		ID: authID,
	})

	return &userpb.GetResponse{
		User: db.FromUser(out.User),
	}, nil
}

func (s Service) AddItem(ctx context.Context, req *userpb.AddItemRequest) (*userpb.AddItemResponse, error) {
	authID, err := auth.GetAuthId(ctx)
	if err != nil {
		return nil, err
	}

	tokenRes, err := s.client.ExchangePublicToken(req.Token)
	if err != nil {
		grpclog.Error(err)
		return nil, status.Errorf(codes.InvalidArgument, "token could not be converted to an access token: %s", err)
	}

	grpclog.Infof("Exchanged public token (%s) for access token (%s)", req.Token, tokenRes.AccessToken)
	out := s.db.AddItem(&db.AddItemInput{
		AuthID:           authID,
		PlaidID:          tokenRes.ItemID,
		PlaidAccessToken: tokenRes.AccessToken,
	})

	accountsRes, err := s.client.GetAccounts(tokenRes.AccessToken)
	if err != nil {
		grpclog.Error(err)
		return nil, status.Errorf(codes.InvalidArgument, "unable to fetch Plaid accounts: %s", err)
	}

	addAccountsRes := s.db.AddAccounts(&db.AddAccountsInput{
		ItemID:   tokenRes.ItemID,
		Accounts: db.ToAccounts(accountsRes.Accounts),
	})

	return &userpb.AddItemResponse{
		New:  out.IsNew,
		User: db.FromUser(addAccountsRes.User),
	}, nil
}

func (s Service) Login(ctx context.Context, req *userpb.LoginRequest) (*userpb.LoginResponse, error) {
	authID, err := auth.GetAuthId(ctx)
	if err != nil {
		return nil, err
	}

	resp := s.db.UpsertUser(&db.UpsertUserInput{
		AuthID:     authID,
		Name:       req.Name,
		PictureURL: req.PictureURL,
		Email:      req.Email,
	})

	return &userpb.LoginResponse{
		New:  resp.IsNew,
		User: db.FromUser(resp.User),
	}, nil
}

// ServiceConfig specifies the configuration for a new User Service.
type ServiceConfig struct {
	Database db.Database
	Client   *plaidgo.Client
}

// Validate implementation of proto interface.
var _ userpb.UserServiceServer = &Service{}

// New returns a new instance of a User service client.
func New(c *ServiceConfig) *Service {
	return &Service{
		db:     c.Database,
		client: c.Client,
	}
}

//// GetTransactions gets the transactions for a given user.
//func (s *Service) GetTransactions(ctx context.Context, in *plaidpb.GetTransactionsRequest) (*plaidpb.GetTransactionsResponse, error) {
//	startDate, endDate := "2018-08-01", "2018-08-31"
//
//	accessToken := s.db.GetToken("1234")
//
//	res, err := s.client.GetTransactions(accessToken, startDate, endDate)
//	if err != nil {
//		grpclog.Error(err)
//		return nil, status.Errorf(codes.NotFound, "could not fetch transactions")
//	}
//
//	var transactions []*plaidpb.Transaction
//	for _, tx := range res.Transactions {
//		transactions = append(transactions, plaid.ToTransaction(tx))
//	}
//
//	grpclog.Infof("Found %d products for access token: %s\n", res.TotalTransactions, accessToken)
//
//	return &plaidpb.GetTransactionsResponse{
//		Transactions: transactions,
//	}, nil
//}
//
//func (s *Service) GetAccounts(context.Context, *plaidpb.GetAccountsRequest) (*plaidpb.GetAccountsResponse, error) {
//	panic("implement me")
//}
