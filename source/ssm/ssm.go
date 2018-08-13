package ssm

import (
	"errors"
	"github.com/kryptn/confg/containers"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// found https://gist.github.com/miguelmota/bd8a5c7942a3587544bcc7cef6bd80de

type SsmSource struct {
	awsSession *session.Session
	ssmSvc     *ssm.SSM
}

func (ss *SsmSource) Lookup(lookup string) (interface{}, bool) {
	withDecryption := false
	param, err := ss.ssmSvc.GetParameter(&ssm.GetParameterInput{
		Name:           &lookup,
		WithDecryption: &withDecryption,
	})
	if err != nil {
		return nil, false
	}

	value := *param.Parameter.Value
	return value, true
}

func (ss *SsmSource) Gather(keys []*containers.Key) {
	for _, key := range keys {
		v, ok := ss.Lookup(key.Lookup)
		key.Inject(v, ok)
	}
}

func Get(backend *containers.Backend) (*SsmSource, error) {
	if backend.Source != "ssm" {
		return nil, errors.New("source.ssm invalid backend")
	}
	ss := SsmSource{}

	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String(backend.AwsRegion)},
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return nil, err
	}
	ss.awsSession = sess
	ss.ssmSvc = ssm.New(sess, aws.NewConfig().WithRegion(backend.AwsRegion))

	return &ss, nil

}
