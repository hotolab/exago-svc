package job

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	. "github.com/hotolab/exago-svc/config"
)

const (
	fnPrefix = "exago-"
)

var (
	svc       *lambda.Lambda
	ErrNoData = errors.New("Empty dataset")
)

// Response contains the generic JSend response sent by Lambda functions.
type Response struct {
	Status   string                 `json:"status"`
	Data     *json.RawMessage       `json:"data"`
	Metadata map[string]interface{} `json:"_metadata"`
}

type context struct {
	Repository string `json:"repository"`
	Branch     string `json:"branch,omitempty"`
	Cleanup    bool   `json:"cleanup,omitempty"`
}

func Init() {
	creds := credentials.NewStaticCredentials(
		Config.AwsAccessKeyID,
		Config.AwsSecretAccessKey,
		"",
	)
	svc = lambda.New(
		session.New(),
		aws.NewConfig().
			WithRegion(Config.AwsRegion).
			WithCredentials(creds),
	)
}

func CallLambdaFn(fn, repo, branch string) (lrsp Response, err error) {
	payload, _ := json.Marshal(context{
		Repository: repo,
		Branch:     branch,
		Cleanup:    true,
	})
	params := &lambda.InvokeInput{
		FunctionName: aws.String(fnPrefix + fn),
		Payload:      payload,
	}

	out, err := svc.Invoke(params)
	if err != nil {
		return lrsp, err
	}

	var resp Response
	if err = json.Unmarshal(out.Payload, &resp); err != nil {
		return lrsp, err
	}

	// Data is always expected from Lambda
	if resp.Data == nil {
		return lrsp, ErrNoData
	}

	// If the Lambda request failed, return the message as an error
	if resp.Status == "fail" {
		var msg struct {
			// Message is the only expected field in Data
			Message string `json:"message"`
		}
		if err = json.Unmarshal(*resp.Data, &msg); err != nil {
			return lrsp, err
		}
		return lrsp, errors.New(msg.Message)
	}

	return resp, nil
}