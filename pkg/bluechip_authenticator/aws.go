package bluechip_authenticator

// Copy code from https://github.com/aws/aws-sdk-go-v2/issues/1922#issuecomment-1432648495

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

const (
	k8sHeader   = "x-k8s-aws-id"
	tokenPrefix = "k8s-aws-v1."
)

type STSTokenRetriever struct {
	PresignClient *sts.PresignClient
}

func NewSTSTokenRetriver(client *sts.PresignClient) STSTokenRetriever {
	return STSTokenRetriever{PresignClient: client}
}

func (s *STSTokenRetriever) GetToken(ctx context.Context, clusterName string) (string, error) {
	out, err := s.PresignClient.PresignGetCallerIdentity(ctx, &sts.GetCallerIdentityInput{}, func(opt *sts.PresignOptions) {
		opt.Presigner = newCustomHTTPPresignerV4(opt.Presigner, map[string]string{
			k8sHeader:       clusterName,
			"X-Amz-Expires": "60",
		})
	})
	if err != nil {
		return "", err
	}
	token := fmt.Sprintf("%s%s", tokenPrefix, base64.RawURLEncoding.EncodeToString([]byte(out.URL))) //RawURLEncoding
	return token, nil
}

type customHTTPPresignerV4 struct {
	client  sts.HTTPPresignerV4
	headers map[string]string
}

func newCustomHTTPPresignerV4(client sts.HTTPPresignerV4, headers map[string]string) sts.HTTPPresignerV4 {
	return &customHTTPPresignerV4{
		client:  client,
		headers: headers,
	}
}

func (p *customHTTPPresignerV4) PresignHTTP(
	ctx context.Context, credentials aws.Credentials, r *http.Request,
	payloadHash string, service string, region string, signingTime time.Time,
	optFns ...func(*v4.SignerOptions),
) (url string, signedHeader http.Header, err error) {
	for key, val := range p.headers {
		r.Header.Add(key, val)
	}
	return p.client.PresignHTTP(ctx, credentials, r, payloadHash, service, region, signingTime, optFns...)
}
