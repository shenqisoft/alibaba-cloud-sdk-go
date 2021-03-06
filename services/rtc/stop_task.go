package rtc

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

// StopTask invokes the rtc.StopTask API synchronously
// api document: https://help.aliyun.com/api/rtc/stoptask.html
func (client *Client) StopTask(request *StopTaskRequest) (response *StopTaskResponse, err error) {
	response = CreateStopTaskResponse()
	err = client.DoAction(request, response)
	return
}

// StopTaskWithChan invokes the rtc.StopTask API asynchronously
// api document: https://help.aliyun.com/api/rtc/stoptask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StopTaskWithChan(request *StopTaskRequest) (<-chan *StopTaskResponse, <-chan error) {
	responseChan := make(chan *StopTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopTask(request)
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

// StopTaskWithCallback invokes the rtc.StopTask API asynchronously
// api document: https://help.aliyun.com/api/rtc/stoptask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StopTaskWithCallback(request *StopTaskRequest, callback func(response *StopTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopTaskResponse
		var err error
		defer close(result)
		response, err = client.StopTask(request)
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

// StopTaskRequest is the request struct for api StopTask
type StopTaskRequest struct {
	*requests.RpcRequest
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
	AppId     string           `position:"Query" name:"AppId"`
	ChannelId string           `position:"Query" name:"ChannelId"`
	TaskId    requests.Integer `position:"Query" name:"TaskId"`
}

// StopTaskResponse is the response struct for api StopTask
type StopTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStopTaskRequest creates a request to invoke StopTask API
func CreateStopTaskRequest() (request *StopTaskRequest) {
	request = &StopTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "StopTask", "rtc", "openAPI")
	return
}

// CreateStopTaskResponse creates a response to parse from StopTask response
func CreateStopTaskResponse() (response *StopTaskResponse) {
	response = &StopTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
