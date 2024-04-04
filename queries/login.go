package query

import (
	"context"
	"fmt"

	graphql "github.com/hasura/go-graphql-client"
)

type LoginAuthenticateMutation struct {
	Authenticate struct {
		Typename                    string `graphql:"__typename"`
		AuthenticateSuccessResponse struct {
			Users []struct {
				AuthToken string
				User      struct {
					Id string
				}
			} `graphql:"users(accessibleType: [DEALER, SUPERVISOR])"`
		} `graphql:"... on AuthenticateSuccessResponse"`
	} `graphql:"authenticate(email: $email, password: $PW, clientApplication: VIRTUAL_KEYPAD, forceSendToEmail: $forceSendToEmail)"`
}

// LoginAuthenticate function calls signin API and returns the auth token
func LoginAuthenticate(email string, password string, endpoint string) string {
	client := graphql.NewClient(endpoint, nil)

	variables := map[string]interface{}{
		"email":            email,
		"PW":               password,
		"forceSendToEmail": false,
	}

	var mutation LoginAuthenticateMutation
	err := client.Mutate(context.Background(), &mutation, variables)
	if err != nil {
		fmt.Println("Failed to authenticate: ", err)
	}

	return mutation.Authenticate.AuthenticateSuccessResponse.Users[0].AuthToken
}
