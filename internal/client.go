package internal

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

	"xuekewang/internal/util/random"

	"github.com/golang-module/dongle"
)

type SdkClient struct {
	appId     string
	appSecret string
	apiAddr   string

	apiClient *http.Client
}

type ApiParams map[string]interface{}

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

func (cli *SdkClient) createRequest(method, action string, params ApiParams, body interface{}) (req *http.Request, err error) {
	action = strings.TrimSpace(action)
	requestUrl := fmt.Sprintf("%s%s", cli.apiAddr, action)

	if params == nil {
		params = make(ApiParams)
	}

	if len(params) > 0 {
		requestUrl += "?" + params.SortQueryString()
	}

	fmt.Printf("requestUrl -> %+v\n", requestUrl)

	var bodyBuf []byte
	if body != nil {
		if bodyBuf, err = json.Marshal(body); err != nil {
			return nil, err
		}
		fmt.Printf("body -> %+v\n", bodyBuf)
		params["xop_body"] = string(bodyBuf)
	}

	params["Xop-App-Id"] = cli.appId
	params["Xop-Timestamp"] = fmt.Sprint(time.Now().Unix())
	// params["Xop-Timestamp"] = 1645151617
	params["Xop-Nonce"] = random.RandomString(36)
	// params["Xop-Nonce"] = "01e7bd52ee7b45328630fe39d7f295ad"
	params["xop_url"] = action

	plainStr := params.SortQueryString() + "&secret=" + cli.appSecret
	fmt.Printf("plainStr -> %+v\n", plainStr)

	base64Str := base64.StdEncoding.EncodeToString([]byte(plainStr))
	fmt.Printf("base64Str -> %+v\n", base64Str)

	sha1Str := dongle.Encrypt.FromString(base64Str).BySha1().String()
	fmt.Printf("sha1Str -> %+v\n", sha1Str)

	sign := sha1Str
	fmt.Printf("sign -> %+v\n", sign)

	if body != nil {
		if req, err = http.NewRequest(method, requestUrl, bytes.NewReader(bodyBuf)); err != nil {
			return nil, err
		}
	} else {
		if req, err = http.NewRequest(method, requestUrl, nil); err != nil {
			return nil, err
		}
	}

	req.Header.Add("Xop-App-Id", fmt.Sprint(params["Xop-App-Id"]))
	req.Header.Add("Xop-Timestamp", fmt.Sprint(params["Xop-Timestamp"]))
	req.Header.Add("Xop-Nonce", fmt.Sprint(params["Xop-Nonce"]))
	req.Header.Add("Xop-Sign", sign)
	req.Header.Add("xop_url", fmt.Sprint(params["xop-url"]))

	return req, nil
}

func (cli *SdkClient) requestJSON(method, action string, params ApiParams, body io.Reader, resultRef interface{}) error {
	req, err := cli.createRequest(method, action, params, body)
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

	fmt.Printf("method=%s action=%s result=%s\n", method, action, resultBuf)

	return json.Unmarshal(resultBuf, resultRef)
}

type ApiBaseResult struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// ======================================================

// 获取学科列表
func (cli *SdkClient) GetSubjects() (res struct {
	ApiBaseResult
	Data []interface{}
}, err error) {
	err = cli.requestJSON("GET", "/xopqbm/subjects", nil, nil, &res)
	return
}

// 获取课程列表
func (cli *SdkClient) GetCoursesAll() (res struct {
	ApiBaseResult
	Data []interface{}
}, err error) {
	err = cli.requestJSON("GET", "/xopqbm/courses/all", nil, nil, &res)
	return
}

// 获取教材列表
func (cli *SdkClient) GetTextBooks() (res struct {
	ApiBaseResult
	Data []interface{}
}, err error) {
	err = cli.requestJSON("GET", "/xopqbm/textbooks", ApiParams{
		"course_id":  27,
		"grade_id":   nil,
		"page_index": 1,
		"page_size":  10,
		"version_id": nil,
	}, nil, &res)
	return
}
