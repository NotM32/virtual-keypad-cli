package query

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type GraphQLConfig struct {
	Name       string   `json:"name"`
	SchemaPath string   `json:"schemaPath"`
	Includes   []string `json:"includes"`
	Extensions struct {
		Endpoints struct {
			Default struct {
				Url     string            `json:"url"`
				Headers map[string]string `json:"headers"`
			} `json:"default"`
		} `json:"endpoints"`
	} `json:"extensions"`
}

func LoadGraphQLConfig() (GraphQLConfig, error) {
	var gqlConfig GraphQLConfig

	gqlConfigFile, err := os.Open(".graphqlconfig")
	if err != nil {
		fmt.Println(".graphqlconfig couldn't be opened: ", err)
		return  gqlConfig, err
	}
	defer gqlConfigFile.Close()

	gqlConfigContent, err := io.ReadAll(gqlConfigFile)
	if err != nil {
		fmt.Println(".graphqlconfig couldn't be read: ", err)
		return  gqlConfig, err
	}

	err = json.Unmarshal(gqlConfigContent, &gqlConfig)
	if err != nil {
		fmt.Println("Failed to unmarshal .graphqlconfig: ", err)
		return  gqlConfig, err
	}

	return gqlConfig, nil
}
