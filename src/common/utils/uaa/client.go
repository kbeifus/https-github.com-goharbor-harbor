// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package uaa

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/goharbor/harbor/src/common/utils/log"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	AuthURLSuffix =     1
	TokenURLSuffix =    2
	UserInfoURLSuffix = 3
	UsersURLSuffix =    4
)

var uaaTransport = &http.Transport{Proxy: http.ProxyFromEnvironment}

// Client provides funcs to interact with UAA.
type Client interface {
	// PasswordAuth accepts username and password, return a token if it's valid.
	PasswordAuth(username, password string) (*oauth2.Token, error)
	// GetUserInfoByToken send the token to OIDC endpoint to get user info, currently it's also used to validate the token.
	GetUserInfo(token string) (*UserInfo, error)
	// SearchUser searches a user based on user name.
	SearchUser(name string) ([]*SearchUserEntry, error)
	// UpdateConfig updates the config of the current client
	UpdateConfig(cfg *ClientConfig) error
}

// ClientConfig values to initialize UAA Client
type ClientConfig struct {
	ClientID      string
	ClientSecret  string
	Endpoint      string
	Realm         string
	SkipTLSVerify bool
	// Absolut path for CA root used to communicate with UAA, only effective when skipTLSVerify set to false.
	CARootPath string
}

// UserInfo represent the JSON object of a userinfo response from UAA.
// As the response varies, this struct will contain only a subset of attributes
// that may be used in Harbor
type UserInfo struct {
	UserID   string `json:"user_id"`
	Sub      string `json:"sub"`
	UserName string `json:"user_name"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

// SearchUserEmailEntry ...
type SearchUserEmailEntry struct {
	Value   string `json:"value"`
	Primary bool   `json:"primary"`
}

// SearchUserEntry is the struct of an entry of user within search result.
type SearchUserEntry struct {
	ID       string                 `json:"id"`
	ExtID    string                 `json:"externalId"`
	UserName string                 `json:"userName"`
	Emails   []SearchUserEmailEntry `json:"emails"`
	Groups   []interface{}
}

type SearchUserRes struct {
	Resources    []*SearchUserEntry `json:"resources"`
	TotalResults int                `json:"totalResults"`
	Schemas      []string           `json:"schemas"`
}

// DefaultClient leverages oauth2 package for oauth features
type defaultClient struct {
	httpClient *http.Client
	oauth2Cfg  *oauth2.Config
	twoLegCfg  *clientcredentials.Config
	endpoint   string
        Realm      string
	//TODO: add public key, etc...
}

func (dc *defaultClient) PasswordAuth(username, password string) (*oauth2.Token, error) {
	return dc.oauth2Cfg.PasswordCredentialsToken(dc.prepareCtx(), username, password)
}

func (dc *defaultClient) GetUserInfo(token string) (*UserInfo, error) {
	userInfoURL := dc.endpoint + dc.GetPrefix(UserInfoURLSuffix)
	req, err := http.NewRequest(http.MethodGet, userInfoURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "bearer "+token)
	resp, err := dc.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	info := &UserInfo{}
	if err := json.Unmarshal(data, info); err != nil {
		return nil, err
	}
	return info, nil
}

func (dc *defaultClient) SearchUser(username string) ([]*SearchUserEntry, error) {

	token, err := dc.twoLegCfg.Token(dc.prepareCtx())
	if err != nil {
	        log.Debugf("erro dc.twoLegCfg.Token %s %v", username, err)
		return nil, err
	}
	url := dc.endpoint + dc.GetPrefix(UsersURLSuffix)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	log.Debugf("erro SearchUser %s %v", username, err)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	//q.Add("filter", fmt.Sprintf("Username eq '%s'", username))
	q.Add("username", username)
	req.URL.RawQuery = q.Encode()
	token.SetAuthHeader(req)
	log.Debugf("request URL: %s", req.URL)
	resp, err := dc.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Debugf("Unexpected status code for searching user in UAA: %d, response: %s", resp.StatusCode, string(bytes))
		return nil, fmt.Errorf("Unexpected status code for searching user in UAA: %d, response: %s", resp.StatusCode, string(bytes))
	}

        if dc.Realm !="" {
            return Parse_rsp_on_realm(bytes)
        }

	res := &SearchUserRes{}
	if err := json.Unmarshal(bytes, res); err != nil {
		return nil, err
	}
	return res.Resources, nil
}

func (dc *defaultClient) prepareCtx() context.Context {
	return context.WithValue(context.Background(), oauth2.HTTPClient, dc.httpClient)
}

func (dc *defaultClient) UpdateConfig(cfg *ClientConfig) error {
	url := cfg.Endpoint
	if !strings.Contains(url, "://") {
		url = "https://" + url
	}
	url = strings.TrimSuffix(url, "/")
	dc.endpoint = url
        dc.Realm = cfg.Realm
	tc := &tls.Config{
		InsecureSkipVerify: cfg.SkipTLSVerify,
	}
	if !cfg.SkipTLSVerify && len(cfg.CARootPath) > 0 {
		if _, err := os.Stat(cfg.CARootPath); !os.IsNotExist(err) {
			content, err := ioutil.ReadFile(cfg.CARootPath)
			if err != nil {
				return err
			}
			pool := x509.NewCertPool()
			// Do not throw error if the certificate is malformed, so we can put a place holder.
			if ok := pool.AppendCertsFromPEM(content); !ok {
				log.Warningf("Failed to append certificate to cert pool, cert path: %s", cfg.CARootPath)
			} else {
				tc.RootCAs = pool
			}
		} else {
			log.Warningf("The root certificate file %s is not found, skip configuring root cert in UAA client.", cfg.CARootPath)
		}
	}
	uaaTransport.TLSClientConfig = tc
	dc.httpClient.Transport = uaaTransport
	// dc.httpClient.Transport = transport.

	oc := &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL: url + dc.GetPrefix(TokenURLSuffix),
			AuthURL:  url + dc.GetPrefix(AuthURLSuffix),
		},
	}

	cc := &clientcredentials.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		TokenURL:     url + dc.GetPrefix(TokenURLSuffix),
	}
	dc.oauth2Cfg = oc
	dc.twoLegCfg = cc
	return nil
}

// NewDefaultClient creates an instance of defaultClient.
func NewDefaultClient(cfg *ClientConfig) (Client, error) {
	hc := &http.Client{}
	c := &defaultClient{httpClient: hc}
	if err := c.UpdateConfig(cfg); err != nil {
		return nil, err
	}
	return c, nil
}

func (dc *defaultClient) GetPrefix(prefix_type int) (string) {

    if dc.Realm !="" {
        switch prefix_type {
        case AuthURLSuffix:
            return "/realms/"+dc.Realm +"/protocol/openid-connect/auth"
        case TokenURLSuffix:
            return "/realms/"+dc.Realm +"/protocol/openid-connect/token"
        case UserInfoURLSuffix:
            return "/realms/"+dc.Realm +"/protocol/openid-connect/userinfo"
        case UsersURLSuffix:
            return "/admin/realms/"+dc.Realm +"/users"
        }
   } else {
      switch prefix_type {
      case AuthURLSuffix:
            return "/oauth/authorize"
      case TokenURLSuffix:
            return "/oauth/token"
      case UserInfoURLSuffix:
            return "/userinfo"
      case UsersURLSuffix:
            return "/Users"
     }
   }
   return ""

}

func Parse_rsp_on_realm(in_buff []byte) ([]*SearchUserEntry, error) {

        var res [] SearchUserEntry
        if err := json.Unmarshal(in_buff, &res); err != nil {
                log.Debugf("SearchUser error: erro %s %v", string(in_buff), err)
                return nil, err
        }

        var ret_val []*SearchUserEntry
        for index, _ := range res {
                ret_val = append(ret_val, &res[index])
        }
        return ret_val, nil

}

