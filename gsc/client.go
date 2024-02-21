package gsc

import (
	"github.com/valyala/fasthttp"
	"log"
	"strings"
)

// 提供与 gae 服务端通信的服务
type RpcRequest struct {
	Host     string
	Token    string
	AppId    string
	AppKey   string
	Endpoint string
	Header   map[string]string
}

func NewRpcClient(id string, key string, host string, token string) *RpcRequest {
	return &RpcRequest{
		Host:     host,
		Header:   map[string]string{},
		Endpoint: "",
		AppId:    id,
		AppKey:   key,
		Token:    token,
	}
}

func NewRpcMasterClient(key string, host string, token string) *RpcRequest {
	return NewRpcClient("0", key, host, token)
}

func reset(r *RpcRequest) *RpcRequest {
	r.Header = map[string]string{}
	return r
}

func (r *RpcRequest) SetEndpoint(serviceName string, className string) *RpcRequest {
	r.Endpoint = serviceName + "/" + className
	return r
}

func (r *RpcRequest) AddHeader(key string, value string) *RpcRequest {
	r.Header[key] = value
	return r
}

func (r *RpcRequest) Call(actionName string, args []string) *RpcResponse {
	defer reset(r)

	req := fasthttp.AcquireRequest()
	// 构建请求头
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("appID", r.AppId)
	req.Header.Set("appKey", r.AppKey)
	if len(r.Token) > 0 {
		req.Header.Set("GrapeSID", r.Token)
	}

	// 构建请求参数
	req.SetRequestURI(r.Host + "/" + r.Endpoint + "/" + actionName)
	GscParameter := buildParameterBody(args)
	req.SetBodyString(GscParameter)

	res := fasthttp.AcquireResponse()

	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()

	err := fasthttp.Do(req, res)
	if err != nil {
		println(err.Error())
		return NewFailRpcResponse("request error")
	}
	statusCode := res.StatusCode()
	if statusCode != fasthttp.StatusOK {
		log.Fatalf("Unexpected status code: %d. Expecting %d", statusCode, fasthttp.StatusOK)
	}
	body := res.Body()
	b, err := NewRpcResponseBody(string(body))
	if err != nil {
		return NewFailRpcResponse("response error")
	}
	return NewRpcResponse(b)
}

func buildParameterBody(args []string) string {
	if len(args) > 0 {
		return "gsc-post:" + strings.Join(args, ":,")
	} else {
		return ""
	}
}
