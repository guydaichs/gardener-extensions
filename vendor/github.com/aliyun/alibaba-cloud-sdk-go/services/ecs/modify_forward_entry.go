package ecs

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

// ModifyForwardEntry invokes the ecs.ModifyForwardEntry API synchronously
// api document: https://help.aliyun.com/api/ecs/modifyforwardentry.html
func (client *Client) ModifyForwardEntry(request *ModifyForwardEntryRequest) (response *ModifyForwardEntryResponse, err error) {
	response = CreateModifyForwardEntryResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyForwardEntryWithChan invokes the ecs.ModifyForwardEntry API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifyforwardentry.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyForwardEntryWithChan(request *ModifyForwardEntryRequest) (<-chan *ModifyForwardEntryResponse, <-chan error) {
	responseChan := make(chan *ModifyForwardEntryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyForwardEntry(request)
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

// ModifyForwardEntryWithCallback invokes the ecs.ModifyForwardEntry API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifyforwardentry.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyForwardEntryWithCallback(request *ModifyForwardEntryRequest, callback func(response *ModifyForwardEntryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyForwardEntryResponse
		var err error
		defer close(result)
		response, err = client.ModifyForwardEntry(request)
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

// ModifyForwardEntryRequest is the request struct for api ModifyForwardEntry
type ModifyForwardEntryRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	IpProtocol           string           `position:"Query" name:"IpProtocol"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	ForwardTableId       string           `position:"Query" name:"ForwardTableId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InternalIp           string           `position:"Query" name:"InternalIp"`
	ForwardEntryId       string           `position:"Query" name:"ForwardEntryId"`
	InternalPort         string           `position:"Query" name:"InternalPort"`
	ExternalIp           string           `position:"Query" name:"ExternalIp"`
	ExternalPort         string           `position:"Query" name:"ExternalPort"`
}

// ModifyForwardEntryResponse is the response struct for api ModifyForwardEntry
type ModifyForwardEntryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyForwardEntryRequest creates a request to invoke ModifyForwardEntry API
func CreateModifyForwardEntryRequest() (request *ModifyForwardEntryRequest) {
	request = &ModifyForwardEntryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyForwardEntry", "ecs", "openAPI")
	return
}

// CreateModifyForwardEntryResponse creates a response to parse from ModifyForwardEntry response
func CreateModifyForwardEntryResponse() (response *ModifyForwardEntryResponse) {
	response = &ModifyForwardEntryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
