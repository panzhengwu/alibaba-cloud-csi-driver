package mts

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

// AddCensorPipeline invokes the mts.AddCensorPipeline API synchronously
// api document: https://help.aliyun.com/api/mts/addcensorpipeline.html
func (client *Client) AddCensorPipeline(request *AddCensorPipelineRequest) (response *AddCensorPipelineResponse, err error) {
	response = CreateAddCensorPipelineResponse()
	err = client.DoAction(request, response)
	return
}

// AddCensorPipelineWithChan invokes the mts.AddCensorPipeline API asynchronously
// api document: https://help.aliyun.com/api/mts/addcensorpipeline.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddCensorPipelineWithChan(request *AddCensorPipelineRequest) (<-chan *AddCensorPipelineResponse, <-chan error) {
	responseChan := make(chan *AddCensorPipelineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddCensorPipeline(request)
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

// AddCensorPipelineWithCallback invokes the mts.AddCensorPipeline API asynchronously
// api document: https://help.aliyun.com/api/mts/addcensorpipeline.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddCensorPipelineWithCallback(request *AddCensorPipelineRequest, callback func(response *AddCensorPipelineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddCensorPipelineResponse
		var err error
		defer close(result)
		response, err = client.AddCensorPipeline(request)
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

// AddCensorPipelineRequest is the request struct for api AddCensorPipeline
type AddCensorPipelineRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Name                 string           `position:"Query" name:"Name"`
	NotifyConfig         string           `position:"Query" name:"NotifyConfig"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Priority             requests.Integer `position:"Query" name:"Priority"`
}

// AddCensorPipelineResponse is the response struct for api AddCensorPipeline
type AddCensorPipelineResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Pipeline  Pipeline `json:"Pipeline" xml:"Pipeline"`
}

// CreateAddCensorPipelineRequest creates a request to invoke AddCensorPipeline API
func CreateAddCensorPipelineRequest() (request *AddCensorPipelineRequest) {
	request = &AddCensorPipelineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "AddCensorPipeline", "mts", "openAPI")
	return
}

// CreateAddCensorPipelineResponse creates a response to parse from AddCensorPipeline response
func CreateAddCensorPipelineResponse() (response *AddCensorPipelineResponse) {
	response = &AddCensorPipelineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
