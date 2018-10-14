package dynamodb

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/colinking/budgeteer/backend/pkg/db"
	"google.golang.org/grpc/grpclog"
)

type Config struct {
	Port int
}

type Database struct {
	db *dynamodb.DynamoDB
}

// New initializes a new DynamoDB database connection.
func New(c *Config) (db.Database, error) {
	return openConnection(c.Port)
}

func openConnection(port int) (db.Database, error) {
	endpoint := fmt.Sprintf("http://localhost:%d", port)
	s, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-west-2"),
		Endpoint: aws.String(endpoint)})
	if err != nil {
		return nil, err
	}

	grpclog.Infof("Opened DynamoDB connection: %s", endpoint)

	return &Database{
		db: dynamodb.New(s),
	}, nil
}

func (d *Database) StoreUser(u *db.User) {
	panic("implement me")
}

func (d *Database) GetUser(userID string) *db.User {
	panic("implement me")
}

func (d *Database) GetUserByEmail(email string) *db.User {
	panic("implement me")
}

func (d *Database) SaveToken(userID string, token string) {
	panic("implement me")
}

func (d *Database) GetToken(userID string) (token string) {
	panic("implement me")
}
