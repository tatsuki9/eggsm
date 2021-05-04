package internal

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/urfave/cli"
)

// GetSecretValues ... get values from secret manager
func GetSecretValues(c *cli.Context) (map[string]interface{}, error) {
	secretID := c.String("secret_id")
	sess := session.Must(session.NewSession())
	svc := secretsmanager.New(
		sess,
		aws.NewConfig().WithRegion("ap-northeast-1"))
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretID),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}
	result, err := svc.GetSecretValue(input)
	if err != nil {
		return nil, err
	}
	secretString := aws.StringValue(result.SecretString)
	res := make(map[string]interface{})
	if err := json.Unmarshal([]byte(secretString), &res); err != nil {
		return nil, err
	}
	return res, nil
}
