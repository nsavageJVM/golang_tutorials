package rest

import (
	"log"
	"golang.org/x/oauth2"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

)


// upper case first letter === public
func TestScope(mssg string ) {

	fmt.Println("TestScope :", mssg)
}

func GetGoogleSignInToken() {


	url := "https://accounts.google.com/o/oauth2/auth"
	fmt.Println("URL:>", url)
	resp, err := http.Get( url)
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func GetGoogleAPIToken () {

	conf := &oauth2.Config{
		ClientID:     "<your client id>",
		ClientSecret: "<your client secret>",
		RedirectURL:  "<your callback url>",
		Scopes: []string{
			"openid email profile https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: oauth2.Endpoint{
		AuthURL:  "https://accounts.google.com/o/oauth2/auth",
		TokenURL: "https://accounts.google.com/o/oauth2/token",
		},
	}

	// AuthCodeURL returns a URL to OAuth 2.0 provider's consent page
	// that asks for permissions for the required scopes explicitly.
	//
	// State is a token to protect the user from CSRF attacks. You must
	// always provide a non-zero string and validate that it matches the
	// the state query parameter on your redirect callback.
	// See http://tools.ietf.org/html/rfc6749#section-10.12 for more info.
	//
	// Opts may include AccessTypeOnline or AccessTypeOffline, as well
	// as ApprovalForce.
//	authUrl := conf.AuthCodeURL("state")
//
//	u, _ := url.Parse(authUrl)
//
//	fmt.Printf("Visit the URL for the auth dialog: %v \n", u )
//	// Handle the exchange code to initiate a transport.
	tok, err := conf.Exchange(oauth2.NoContext, "4/cBaIyVG3bsCfpramR-fid_BQLny659kUODyaPfdrh6Y")
	if err != nil {
		log.Fatal(err)
	}
	client := conf.Client(oauth2.NoContext, tok)
	client.Get("...")

	fmt.Printf("\n got a token: %v", tok)


}

func urlEncoded(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}