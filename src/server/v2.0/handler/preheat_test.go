package handler

import (
	"reflect"
	"testing"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/goharbor/harbor/src/pkg/p2p/preheat/models/policy"
	"github.com/goharbor/harbor/src/pkg/p2p/preheat/provider"
	"github.com/goharbor/harbor/src/server/v2.0/models"
	"github.com/stretchr/testify/assert"
)

func Test_convertProvidersToFrontend(t *testing.T) {
	backend, _ := provider.ListProviders()
	tests := []struct {
		name         string
		backend      []*provider.Metadata
		wantFrontend []*models.Metadata
	}{
		{"",
			backend,
			[]*models.Metadata{
				{ID: "dragonfly", Icon: "https://raw.githubusercontent.com/alibaba/Dragonfly/master/docs/images/logo.png", Maintainers: []string{"Jin Zhang/taiyun.zj@alibaba-inc.com"}, Name: "Dragonfly", Source: "https://github.com/alibaba/Dragonfly", Version: "0.10.1"},
				{Icon: "https://github.com/uber/kraken/blob/master/assets/kraken-logo-color.svg", ID: "kraken", Maintainers: []string{"mmpei/peimingming@corp.netease.com"}, Name: "Kraken", Source: "https://github.com/uber/kraken", Version: "0.1.3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFrontend := convertProvidersToFrontend(tt.backend); !reflect.DeepEqual(gotFrontend, tt.wantFrontend) {
				t.Errorf("convertProvidersToFrontend() = %#v, want %#v", gotFrontend, tt.wantFrontend)
			}
		})
	}
}

func Test_convertPolicyToPayload(t *testing.T) {
	tests := []struct {
		name      string
		input     *policy.Schema
		expect    *models.PreheatPolicy
		shouldErr bool
	}{
		{
			name:      "should error",
			input:     nil,
			expect:    nil,
			shouldErr: true,
		},
		{
			name: "should success",
			input: &policy.Schema{
				ID:          0,
				Name:        "abc",
				Description: "test case",
				ProjectID:   0,
				ProviderID:  0,
				Filters:     nil,
				FiltersStr:  "",
				Trigger:     nil,
				TriggerStr:  "",
				Enabled:     false,
				CreatedAt:   time.Time{},
				UpdatedTime: time.Time{},
			},
			expect: &models.PreheatPolicy{
				CreationTime: strfmt.DateTime{},
				Description:  "test case",
				Enabled:      false,
				Filters:      "",
				ID:           0,
				Name:         "abc",
				ProjectID:    0,
				ProviderID:   0,
				Trigger:      "",
				UpdateTime:   strfmt.DateTime{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := convertPolicyToPayload(tt.input)
			if !tt.shouldErr {
				if !assert.Equal(t, tt.expect, actual) {
					t.Errorf("convertPolicyToPayload() = %#v, want %#v", actual, tt.expect)
				}
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func Test_convertParamPolicyToModelPolicy(t *testing.T) {
	tests := []struct {
		name      string
		input     *models.PreheatPolicy
		expect    *policy.Schema
		shouldErr bool
	}{
		{
			name:      "should err",
			input:     nil,
			expect:    nil,
			shouldErr: true,
		},
		{
			name: "should success",
			input: &models.PreheatPolicy{
				CreationTime: strfmt.DateTime{},
				Description:  "test case",
				Enabled:      false,
				Filters:      "",
				ID:           0,
				Name:         "abc",
				ProjectID:    0,
				ProviderID:   0,
				Trigger:      "",
				UpdateTime:   strfmt.DateTime{},
			},
			expect: &policy.Schema{
				ID:          0,
				Name:        "abc",
				Description: "test case",
				ProjectID:   0,
				ProviderID:  0,
				Filters:     nil,
				FiltersStr:  "",
				Trigger:     nil,
				TriggerStr:  "",
				Enabled:     false,
				CreatedAt:   time.Time{},
				UpdatedTime: time.Time{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := convertParamPolicyToModelPolicy(tt.input)
			if !tt.shouldErr {
				if !assert.Equal(t, tt.expect, actual) {
					t.Errorf("convertParamPolicyToModelPolicy() = %#v, want %#v", actual, tt.expect)
				}
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}
