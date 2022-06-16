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

package auth

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/goharbor/harbor/src/common/http/modifier"
)

type apiKeyType = string

const (
	// APIKeyInHeader sets auth content in header
	APIKeyInHeader apiKeyType = "header"
	// APIKeyInQuery sets auth content in url query
	APIKeyInQuery apiKeyType = "query"
)

type apiKeyAuthorizer struct {
	key, value, in apiKeyType
}

// NewAPIKeyAuthorizer returns a apikey authorizer
func NewAPIKeyAuthorizer(key, value, in apiKeyType) modifier.Modifier {
	return &apiKeyAuthorizer{
		key:   key,
		value: value,
		in:    in,
	}
}

// Modify implements modifier.Modifier
func (a *apiKeyAuthorizer) Modify(r *http.Request) error {
	switch a.in {
	case APIKeyInHeader:
		r.Header.Set(a.key, a.value)
		return nil
	case APIKeyInQuery:
		values, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			return err
		}
		values.Add(a.key, a.value)
		r.URL.RawQuery = values.Encode()
		return nil
	}
	return fmt.Errorf("set api key in %s is invalid", a.in)
}
