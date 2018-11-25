package user

import (
	"github.com/colinking/budgeteer/backend/pkg/auth"
	"github.com/colinking/budgeteer/backend/pkg/db"
	"github.com/colinking/budgeteer/backend/pkg/gen/userpb"
	"golang.org/x/net/context"
	"google.golang.org/grpc/grpclog"
)

// Service contains all User-related handlers.
type Service struct {
	db db.Database
}

func (s Service) Login(ctx context.Context, req *userpb.LoginRequest) (*userpb.LoginResponse, error) {
	// TODO: move out to auth-ed middleware
	authID, err := auth.GetAuthId(ctx)
	if err != nil {
		return nil, err
	}

	grpclog.Infof("Received request from: %s\n", authID)

	resp := s.db.UpsertUser(&db.UpsertUserInput{
		AuthID:     authID,
		FirstName:  req.User.FirstName,
		LastName:   req.User.LastName,
		PictureURL: req.User.PictureURL,
		Email:      req.User.Email,
	})

	return &userpb.LoginResponse{
		New: resp.IsNew,
	}, nil
}

// ServiceConfig specifies the configuration for a new User Service.
type ServiceConfig struct {
	Database db.Database
}

// Validate implementation of proto interface.
var _ userpb.UserServiceServer = &Service{}

// New returns a new instance of a Plaid service client.
func New(c *ServiceConfig) *Service {
	return &Service{
		db: c.Database,
	}
}
