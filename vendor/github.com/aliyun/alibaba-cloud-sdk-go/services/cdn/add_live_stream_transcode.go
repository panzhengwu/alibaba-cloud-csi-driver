package cdn

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

// AddLiveStreamTranscode invokes the cdn.AddLiveStreamTranscode API synchronously
// api document: https://help.aliyun.com/api/cdn/addlivestreamtranscode.html
func (client *Client) AddLiveStreamTranscode(request *AddLiveStreamTranscodeRequest) (response *AddLiveStreamTranscodeResponse, err error) {
	response = CreateAddLiveStreamTranscodeResponse()
	err = client.DoAction(request, response)
	return
}

// AddLiveStreamTranscodeWithChan invokes the cdn.AddLiveStreamTranscode API asynchronously
// api document: https://help.aliyun.com/api/cdn/addlivestreamtranscode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddLiveStreamTranscodeWithChan(request *AddLiveStreamTranscodeRequest) (<-chan *AddLiveStreamTranscodeResponse, <-chan error) {
	responseChan := make(chan *AddLiveStreamTranscodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddLiveStreamTranscode(request)
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

// AddLiveStreamTranscodeWithCallback invokes the cdn.AddLiveStreamTranscode API asynchronously
// api document: https://help.aliyun.com/api/cdn/addlivestreamtranscode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddLiveStreamTranscodeWithCallback(request *AddLiveStreamTranscodeRequest, callback func(response *AddLiveStreamTranscodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddLiveStreamTranscodeResponse
		var err error
		defer close(result)
		response, err = client.AddLiveStreamTranscode(request)
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

// AddLiveStreamTranscodeRequest is the request struct for api AddLiveStreamTranscode
type AddLiveStreamTranscodeRequest struct {
	*requests.RpcRequest
	Template      string           `position:"Query" name:"Template"`
	App           string           `position:"Query" name:"App"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	OwnerAccount  string           `position:"Query" name:"OwnerAccount"`
	Domain        string           `position:"Query" name:"Domain"`
	Record        string           `position:"Query" name:"Record"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	Snapshot      string           `position:"Query" name:"Snapshot"`
}

// AddLiveStreamTranscodeResponse is the response struct for api AddLiveStreamTranscode
type AddLiveStreamTranscodeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddLiveStreamTranscodeRequest creates a request to invoke AddLiveStreamTranscode API
func CreateAddLiveStreamTranscodeRequest() (request *AddLiveStreamTranscodeRequest) {
	request = &AddLiveStreamTranscodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "AddLiveStreamTranscode", "", "")
	return
}

// CreateAddLiveStreamTranscodeResponse creates a response to parse from AddLiveStreamTranscode response
func CreateAddLiveStreamTranscodeResponse() (response *AddLiveStreamTranscodeResponse) {
	response = &AddLiveStreamTranscodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
