package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/LFreedomDev/xuekewang/util/random"

	"github.com/golang-module/dongle"
)

type SdkClient struct {
	appId     string
	appSecret string
	apiAddr   string

	apiClient *http.Client
}

type ApiParams map[string]interface{}

func NewApiParamsFromObject(value interface{}) ApiParams {
	buf, _ := json.Marshal(value)
	var res ApiParams
	json.Unmarshal(buf, &res)
	return res
}

func (sm ApiParams) SortQueryString() string {

	keys := make([]string, 0, len(sm))
	for k := range sm {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	params := []string{}
	for _, key := range keys {
		if val := sm[key]; val == nil {
			params = append(params, fmt.Sprintf("%s=", key))
		} else {
			params = append(params, fmt.Sprintf("%s=%v", key, val))
		}
	}

	return strings.Join(params, "&")
}

func NewSdkClient(appId, appSecret string, timeout time.Duration) *SdkClient {
	res := &SdkClient{
		appId:     appId,
		appSecret: appSecret,
		apiClient: &http.Client{
			Timeout: timeout,
		},
	}
	if res.apiAddr == "" {
		res.apiAddr = "https://openapi.xkw.com"
	} else {
		res.apiAddr = strings.TrimSpace(res.apiAddr)
		res.apiAddr = strings.TrimRight(res.apiAddr, "/")
	}
	return res
}

func (cli *SdkClient) createRequest(method, action string, queryParams ApiParams, reqBody interface{}) (req *http.Request, err error) {
	action = strings.TrimSpace(action)
	requestUrl := fmt.Sprintf("%s%s", cli.apiAddr, action)

	if queryParams == nil {
		queryParams = make(ApiParams)
	}

	if len(queryParams) > 0 {
		requestUrl += "?" + queryParams.SortQueryString()
	}

	// fmt.Printf("requestUrl -> %+v\n", requestUrl)

	var bodyBuf []byte
	if reqBody != nil {
		if bodyBuf, err = json.Marshal(reqBody); err != nil {
			return nil, err
		}
		// fmt.Printf("body -> %+v\n", bodyBuf)
		queryParams["xop_body"] = string(bodyBuf)
	}

	queryParams["Xop-App-Id"] = cli.appId
	queryParams["Xop-Timestamp"] = fmt.Sprint(time.Now().Unix())
	queryParams["Xop-Nonce"] = random.RandomString(36)
	queryParams["xop_url"] = action

	plainStr := queryParams.SortQueryString() + "&secret=" + cli.appSecret
	// fmt.Printf("plainStr -> %+v\n", plainStr)

	base64Str := base64.StdEncoding.EncodeToString([]byte(plainStr))
	// fmt.Printf("base64Str -> %+v\n", base64Str)

	sha1Str := dongle.Encrypt.FromString(base64Str).BySha1().String()
	// fmt.Printf("sha1Str -> %+v\n", sha1Str)

	sign := sha1Str
	// fmt.Printf("sign -> %+v\n", sign)

	if reqBody != nil {
		if req, err = http.NewRequest(method, requestUrl, bytes.NewReader(bodyBuf)); err != nil {
			return nil, err
		}
	} else {
		if req, err = http.NewRequest(method, requestUrl, nil); err != nil {
			return nil, err
		}
	}

	req.Header.Add("Xop-App-Id", fmt.Sprint(queryParams["Xop-App-Id"]))
	req.Header.Add("Xop-Timestamp", fmt.Sprint(queryParams["Xop-Timestamp"]))
	req.Header.Add("Xop-Nonce", fmt.Sprint(queryParams["Xop-Nonce"]))
	req.Header.Add("Xop-Sign", sign)
	req.Header.Add("xop_url", fmt.Sprint(queryParams["xop-url"]))

	return req, nil
}

func (cli *SdkClient) requestJSON(method, action string, queryParams ApiParams, reqBody interface{}, resultRef interface{}) error {
	req, err := cli.createRequest(method, action, queryParams, reqBody)
	if err != nil {
		return err
	}

	resp, err := cli.apiClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	resultBuf, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// fmt.Printf("method=%s action=%s result=%s\n", method, action, resultBuf)

	return json.Unmarshal(resultBuf, resultRef)
}

type ApiBaseResult struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func (res ApiBaseResult) Error() error {
	if res.Code != 2000000 {
		return fmt.Errorf("%v - %v", res.Code, res.Message)
	}
	return nil
}
