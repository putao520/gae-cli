package gsc

import "encoding/json"

type RpcResponseBody struct {
	Code    int         `json:"errorcode"`
	Message string      `json:"message"`
	Data    interface{} `json:"record"`
}

type RpcResponse struct {
	Body RpcResponseBody
}

func NewRpcResponseBody(responseString string) (*RpcResponseBody, error) {
	var response RpcResponseBody

	err := json.Unmarshal([]byte(responseString), &response)
	if err != nil {
		return nil, err
	}

	return &response, nil

}

func NewRpcResponse(response *RpcResponseBody) *RpcResponse {
	return &RpcResponse{
		Body: *response,
	}
}

func NewFailRpcResponse(errorMessage string) *RpcResponse {
	r := &RpcResponseBody{
		Code:    1,
		Message: errorMessage,
		Data:    nil,
	}
	return NewRpcResponse(r)
}

func (r *RpcResponse) GetMessage() string {
	return r.Body.Message
}

// 判断返回是否成功
func (r *RpcResponse) GetStatus() bool {
	return r.Body.Code == 0
}

// 获得返回的数据,字符串类型
func (r *RpcResponse) AsString() string {
	return r.Body.Data.(string)
}

// 获得返回的数据,整型类型
func (r *RpcResponse) AsInt() int {
	return r.Body.Data.(int)
}

// 获得返回的数据,浮点类型
func (r *RpcResponse) AsFloat() float64 {
	return r.Body.Data.(float64)
}

// 获得返回的数据,布尔类型
func (r *RpcResponse) AsBool() bool {
	return r.Body.Data.(bool)
}

// 获得返回的数据,数组类型
func (r *RpcResponse) AsArray() []interface{} {
	return r.Body.Data.([]interface{})
}

// 获得返回的数据,对象类型
func (r *RpcResponse) AsObject() map[string]interface{} {
	return r.Body.Data.(map[string]interface{})
}
