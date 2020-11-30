//testing go locally
package main
 
import "fmt"
 
// func main() {
// 	fmt.Println("Hi there")
// }

import (
  //project imports
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// CREDENTIALS størs all access/consumer tokens and secret keys needed to auth against the twitter API

type Credentials struct {
	ConsumerKey			string
	ConsumerSecret		string
	AccessToken			string
	AccessTokenSecret 	string
}

// getClient is a help function that will return a twitter client which can be used to send or stream tweets
// this will ingest a pointer and it direct to the Credential struct which will contain necessary data to authenticate 
// and return a pointer to a twitter client or an error

func getClient(creds *Credentials) (*twitter.Client, error) {
	//send comsumer key and secret
	config := oauth.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	//send access token and access token secret
	token := oauth.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	//verify credentials

	verifyParams := &twitter.AccountVerifyingParams{
		SkipStatus:		twitter.Bool(true),
		InclueEmail: 	twitter.Bool(true),
	}

	//retrieve and verify if the creds used will allow login

	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}
	log.Printf(Users's ACCOUNT:\n%+v\n", user)
}