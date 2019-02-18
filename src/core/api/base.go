// Copyright 2018 Project Harbor Authors
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

package api

import (
	"net/http"

	"github.com/ghodss/yaml"
	"github.com/goharbor/harbor/src/common/api"
	"github.com/goharbor/harbor/src/common/security"
	"github.com/goharbor/harbor/src/common/utils/log"
	"github.com/goharbor/harbor/src/core/config"
	"github.com/goharbor/harbor/src/core/filter"
	"github.com/goharbor/harbor/src/core/promgr"
	"github.com/goharbor/harbor/src/core/utils"
	commonhttp "github.com/goharbor/harbor/src/common/http"
)

const (
	yamlFileContentType = "application/x-yaml"
)

// BaseController ...
type BaseController struct {
	api.BaseAPI
	// SecurityCtx is the security context used to authN &authZ
	SecurityCtx security.Context
	// ProjectMgr is the project manager which abstracts the operations
	// related to projects
	ProjectMgr promgr.ProjectManager
}

const (
	// ReplicationJobType ...
	ReplicationJobType = "replication"
	// ScanJobType ...
	ScanJobType = "scan"
)

// Prepare inits security context and project manager from request
// context
func (b *BaseController) Prepare() {
	ctx, err := filter.GetSecurityContext(b.Ctx.Request)
	if err != nil {
		log.Errorf("failed to get security context: %v", err)
		b.CustomAbort(http.StatusInternalServerError, "")
	}
	b.SecurityCtx = ctx

	pm, err := filter.GetProjectManager(b.Ctx.Request)
	if err != nil {
		log.Errorf("failed to get project manager: %v", err)
		b.CustomAbort(http.StatusInternalServerError, "")
	}
	b.ProjectMgr = pm
}

// RenderFormatedError renders errors with well formted style `{"error": "This is an error"}`
func (b *BaseController) RenderFormatedError(statusCode, errorCode int, err error) {
	formatedErr := utils.WrapError(errorCode, err)
	log.Errorf("%s %s failed with error: %s", b.Ctx.Request.Method, b.Ctx.Request.URL.String(), formatedErr.Error())
	b.RenderError(statusCode, formatedErr.Error())
}

func (b *BaseController) RenderError(code int, text string) {
	http.Error(b.Ctx.ResponseWriter, text, code)
}

// ParseAndHandleError : if the err is an instance of utils/error.Error,
// return the status code and the detail message contained in err, otherwise
// return 500
func (b *BaseController) ParseAndHandleError(text string, err error) {
	if err == nil {
		return
	}
	log.Errorf("%s: %v", text, err)
	if e, ok := err.(*commonhttp.Error); ok {
		b.RenderError(e.Code, e.Message)
		return
	}
	b.RenderError(http.StatusInternalServerError, "")
}

// SendUnAuthorizedError sends unauthorized error to the client.
func (b *BaseController) SendUnAuthorizedError(err error) {
	b.RenderFormatedError(http.StatusUnauthorized, http.StatusUnauthorized, err)
}

// SendConflictError sends conflict error to the client.
func (b *BaseController) SendConflictError(err error) {
	b.RenderFormatedError(http.StatusConflict, http.StatusConflict, err)
}

// SendNotFoundError sends not found error to the client.
func (b *BaseController) SendNotFoundError(err error) {
	b.RenderFormatedError(http.StatusNotFound, http.StatusNotFound, err)
}

// SendBadRequestError sends bad request error to the client.
func (b *BaseController) SendBadRequestError(err error) {
	b.RenderFormatedError(http.StatusBadRequest, http.StatusBadRequest, err)
}

// SendInternalServerError sends internal server error to the client.
func (b *BaseController) SendInternalServerError(err error) {
	b.RenderFormatedError(http.StatusInternalServerError, http.StatusInternalServerError, err)
}

// SendForbiddenError sends forbidden error to the client.
func (b *BaseController) SendForbiddenError(err error) {
	b.RenderFormatedError(http.StatusForbidden, http.StatusForbidden, err)
}

// SendPreconditionFailedError sends forbidden error to the client.
func (b *BaseController) SendPreconditionFailedError(err error) {
	b.RenderFormatedError(http.StatusPreconditionFailed, http.StatusPreconditionFailed, err)
}

// WriteJSONData writes the JSON data to the client.
func (b *BaseController) WriteJSONData(object interface{}) {
	b.Data["json"] = object
	b.ServeJSON()
}

// WriteYamlData writes the yaml data to the client.
func (b *BaseController) WriteYamlData(object interface{}) {
	yData, err := yaml.Marshal(object)
	if err != nil {
		b.SendInternalServerError(err)
		return
	}

	w := b.Ctx.ResponseWriter
	w.Header().Set("Content-Type", yamlFileContentType)
	w.WriteHeader(http.StatusOK)
	w.Write(yData)
}

// Init related objects/configurations for the API controllers
func Init() error {
	registerHealthCheckers()
	// If chart repository is not enabled then directly return
	if !config.WithChartMuseum() {
		return nil
	}

	chartCtl, err := initializeChartController()
	if err != nil {
		return err
	}

	chartController = chartCtl

	return nil
}
