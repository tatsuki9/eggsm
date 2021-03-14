package internal

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/urfave/cli"
)

// GetSecretValues ... get values from secret manager
func GetSecretValues(c *cli.Context) (map[string]interface{}, error) {
	profile := c.String("profile")
	prefix := c.String("prefix")
	env := c.String("env")
	sess := session.Must(session.NewSessionWithOptions(session.Options{Profile: profile}))
	svc := secretsmanager.New(
		sess,
		aws.NewConfig().WithRegion("ap-northeast-1"))
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(getSecretName(prefix, env)),
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

func getSecretName(prefix, env string) string {
	return fmt.Sprintf("%s.%s", prefix, env)
}
