package user

import (
	"github.com/colinking/budgeteer/backend/pkg/db"
	proto "github.com/colinking/budgeteer/backend/pkg/proto/user"
	"github.com/go-chi/jwtauth"
	"golang.org/x/net/context"
)

// Service contains all User-related handlers.
type Service struct {
	db     db.Database
}

func getAuthId(ctx context.Context) (string) {
	_, claims, _ := jwtauth.FromContext(ctx)
	return claims["sub"].(string)
}

func (s Service) UserLogin(ctx context.Context, req *proto.UserLoginRequest) (*proto.UserLoginResponse, error) {
	// Get user by auth0 user id
	// TODO: actually implement by querying DB
	getAuthId(ctx)

	return &proto.UserLoginResponse{
		New: true,
	}, nil
}

// ServiceConfig specifies the configuration for a new User Service.
type ServiceConfig struct {
	Database db.Database
}

// Validate implementation of proto interface.
var _ proto.UserServiceServer = &Service{}

// New returns a new instance of a Plaid service client.
func New(c *ServiceConfig) *Service {
	return &Service{
		db:     c.Database,
	}
}

