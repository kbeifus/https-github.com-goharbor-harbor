// Copyright 2018 The Harbor Authors
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

package source

import (
	"github.com/goharbor/harbor/src/replication/models"
)

// DefaultFilterChain provides a default implement for interface FilterChain
type DefaultFilterChain struct {
	filters []Filter
}

// NewDefaultFilterChain returns an instance of DefaultFilterChain
func NewDefaultFilterChain(filters []Filter) *DefaultFilterChain {
	return &DefaultFilterChain{
		filters: filters,
	}
}

// Build nil implement now
func (d *DefaultFilterChain) Build(filters []Filter) error {
	return nil
}

// Filters returns the filter list
func (d *DefaultFilterChain) Filters() []Filter {
	return d.filters
}

// DoFilter does the filter works for filterItems
func (d *DefaultFilterChain) DoFilter(filterItems []models.FilterItem) []models.FilterItem {
	if len(filterItems) == 0 {
		return []models.FilterItem{}
	}

	for _, filter := range d.filters {
		convertor := filter.GetConvertor()
		if convertor != nil {
			filterItems = convertor.Convert(filterItems)
		}
		filterItems = filter.DoFilter(filterItems)
	}
	return filterItems
}
