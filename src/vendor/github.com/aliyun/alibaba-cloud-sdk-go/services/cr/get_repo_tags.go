package cr

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// GetRepoTags invokes the cr.GetRepoTags API synchronously
// api document: https://help.aliyun.com/api/cr/getrepotags.html
func (client *Client) GetRepoTags(request *GetRepoTagsRequest) (response *GetRepoTagsResponse, err error) {
	response = CreateGetRepoTagsResponse()
	err = client.DoAction(request, response)
	return
}

// GetRepoTagsWithChan invokes the cr.GetRepoTags API asynchronously
// api document: https://help.aliyun.com/api/cr/getrepotags.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRepoTagsWithChan(request *GetRepoTagsRequest) (<-chan *GetRepoTagsResponse, <-chan error) {
	responseChan := make(chan *GetRepoTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRepoTags(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// GetRepoTagsWithCallback invokes the cr.GetRepoTags API asynchronously
// api document: https://help.aliyun.com/api/cr/getrepotags.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRepoTagsWithCallback(request *GetRepoTagsRequest, callback func(response *GetRepoTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRepoTagsResponse
		var err error
		defer close(result)
		response, err = client.GetRepoTags(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// GetRepoTagsRequest is the request struct for api GetRepoTags
type GetRepoTagsRequest struct {
	*requests.RoaRequest
	RepoNamespace string           `position:"Path" name:"RepoNamespace"`
	RepoName      string           `position:"Path" name:"RepoName"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	Page          requests.Integer `position:"Query" name:"Page"`
}

// GetRepoTagsResponse is the response struct for api GetRepoTags
type GetRepoTagsResponse struct {
	*responses.BaseResponse
}

// CreateGetRepoTagsRequest creates a request to invoke GetRepoTags API
func CreateGetRepoTagsRequest() (request *GetRepoTagsRequest) {
	request = &GetRepoTagsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("cr", "2016-06-07", "GetRepoTags", "/repos/[RepoNamespace]/[RepoName]/tags", "cr", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetRepoTagsResponse creates a response to parse from GetRepoTags response
func CreateGetRepoTagsResponse() (response *GetRepoTagsResponse) {
	response = &GetRepoTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
