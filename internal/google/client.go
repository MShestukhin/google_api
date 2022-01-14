package google

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

type GoogleAdsClient struct {
	credentials    *oauth2.Config
	token          *oauth2.Token
	conn           *grpc.ClientConn
	developerToken string
	ctx            context.Context
}

type GoogleAdsClientParams struct {
	ClientID       string
	ClientSecret   string
	DeveloperToken string
	RefreshToken   string
}

func RefreshToken(params *GoogleAdsClientParams) (string, error) {
	credentials := NewCredentials(params.ClientID, params.ClientSecret)
	initialToken := NewPartialToken(params.RefreshToken)

	newToken, err := refreshToken(credentials, initialToken)
	if err != nil {
		return "", err
	}
	return newToken.AccessToken, nil
}

// NewClient creates a new client with specified credential params
func NewClient(token string, developerToken string) (*GoogleAdsClient, error) {
	conn, ctx, err := NewGrpcConnection(NewToken(token), developerToken)

	if err != nil {
		return nil, err
	}
	c := &GoogleAdsClient{}
	c.conn = conn
	c.ctx = ctx

	return c, nil
}

func (g *GoogleAdsClient) Conn() *grpc.ClientConn {
	return g.conn
}

func (g *GoogleAdsClient) Context() context.Context {
	return g.ctx
}

func (g *GoogleAdsClient) TokenIsValid() bool {
	return g.token.Valid()
}

const (
	address string = "googleads.googleapis.com:443"
)

func NewCredentials(clientID, clientSecret string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     google.Endpoint,
	}
}

func NewPartialToken(refreshToken string) *oauth2.Token {
	return &oauth2.Token{
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
	}
}

func NewToken(token string) *oauth2.Token {
	return &oauth2.Token{
		AccessToken: token,
	}
}

func refreshToken(creds *oauth2.Config, token *oauth2.Token, ctx ...context.Context) (*oauth2.Token, error) {
	defaultCtx := context.Background()
	if ctx != nil {
		defaultCtx = ctx[0]
	}
	tokenSource := creds.TokenSource(defaultCtx, token)
	newToken, err := tokenSource.Token()
	if err != nil {
		return nil, err
	}
	return newToken, nil
}

func NewGrpcConnection(token *oauth2.Token, developerToken string) (*grpc.ClientConn, context.Context, error) {
	headers := createHeaders(token.AccessToken, developerToken)
	ctx := metadata.NewOutgoingContext(context.Background(), headers)

	transportCreds := grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, ""))
	conn, err := grpc.Dial(address, transportCreds)
	if err != nil {
		return nil, nil, err
	}
	return conn, ctx, nil
}

func createHeaders(accessToken string, developerToken string) metadata.MD {
	return metadata.Pairs(
		"Authorization", fmt.Sprintf("Bearer %s", accessToken),
		"developer-token", developerToken,
	)
}
