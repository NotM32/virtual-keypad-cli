package main

import (
	"fmt"
	"os"

	query "github.com/NotM32/godmp/queries"
	"github.com/NotM32/godmp/cmd"
)

var (
	email          = os.Getenv("DMP_EMAIL")
	password       = os.Getenv("DMP_PASSWORD")
	userCode       = os.Getenv("DMP_PASSCODE")
	panel          = os.Getenv("DMP_PANEL")
	panelSecondary = os.Getenv("DMP_PANEL_SECONDARY")
)

func main() {
	gqlConfig, err := loadGraphQLConfig()
	if err != nil {
		fmt.Println("Error reading GraphQLConfig: ", err)
		panic(err)
	}

	authToken := query.LoginAuthenticate(email, password, gqlConfig.Extensions.Endpoints.Default.Url)

	cmd.Execute()

	doors := query.GetDoors(authToken, userCode, panel, panelSecondary)
	doors[2].Access()
}
