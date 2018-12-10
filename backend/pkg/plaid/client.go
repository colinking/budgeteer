package plaid

import (
	"encoding/json"
	"errors"

	plaidgo "github.com/plaid/plaid-go/plaid"
)

type Client struct {
	*plaidgo.Client

	publicKey string
}

func NewClient(client *plaidgo.Client, publicKey string) *Client {
	return &Client{
		Client:    client,
		publicKey: publicKey,
	}
}

type Institution struct {
	Credentials []plaidgo.Credential `json:"credentials"`
	HasMFA      bool                 `json:"has_mfa"`
	ID          string               `json:"institution_id"`
	MFA         []string             `json:"mfa"`
	Name        string               `json:"name"`
	BrandName   string               `json:"brand_name"`
	Products    []string             `json:"products"`
	Colors      InstitutionColors    `json:"colors"`
	Logo        string               `json:"logo"`
	URL         string               `json:"url"`
}

type InstitutionColors struct {
	Dark    string `json:"dark"`
	Darker  string `json:"darker"`
	Light   string `json:"light"`
	Primary string `json:"primary"`
}

type getInstitutionByIDWithDisplayRequest struct {
	ID        string                                      `json:"institution_id"`
	PublicKey string                                      `json:"public_key"`
	Options   GetInstitutionByIDWithDisplayRequestOptions `json:"options"`
}

type GetInstitutionByIDWithDisplayResponse struct {
	plaidgo.APIResponse
	Institution Institution `json:"institution"`
}

type GetInstitutionByIDWithDisplayRequestOptions struct {
	IncludeDisplayData bool `json:"include_display_data"`
}

// GetInstitutionByID returns information for a single institution given an ID.
// See https://plaid.com/docs/api/#institutions-by-id.
// Forked from: https://github.com/plaid/plaid-go/blob/d56a645/plaid/institutions.go#L57-L75
func (c *Client) GetInstitutionByIDWithDisplay(id string) (resp GetInstitutionByIDWithDisplayResponse, err error) {
	if id == "" {
		return resp, errors.New("/institutions/get_by_id - institution id must be specified")
	}

	jsonBody, err := json.Marshal(getInstitutionByIDWithDisplayRequest{
		ID:        id,
		PublicKey: c.publicKey,
		Options: GetInstitutionByIDWithDisplayRequestOptions{
			IncludeDisplayData: true,
		},
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/institutions/get_by_id", jsonBody, &resp)
	return resp, err
}
