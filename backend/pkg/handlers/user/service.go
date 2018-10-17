package user

import (
	"github.com/colinking/budgeteer/backend/pkg/db"
	proto "github.com/colinking/budgeteer/backend/pkg/proto/user"
	"golang.org/x/net/context"
)

// Service contains all User-related handlers.
type Service struct {
	db     db.Database
}

func (s Service) UserLogin(context.Context, *proto.UserLoginRequest) (*proto.UserLoginResponse, error) {
	// Get user by auth0 user id
	

	//	TODO, actually implement
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

