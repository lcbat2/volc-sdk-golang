package businessSecurity

import (
	"fmt"
	"github.com/volcengine/volc-sdk-golang/base"
	"net/http"
	"net/url"
	"time"
)

type BusinessSecurity struct {
	*base.Client
	retry bool
}

var DefaultInstance = NewInstance()

func NewInstance() *BusinessSecurity {
	instance := &BusinessSecurity{
		Client: base.NewClient(ServiceInfoMap[base.RegionCnNorth1], ApiInfoList),
		retry:  true,
	}
	return instance
}

func GetServiceInfo(region string, host string, timeout time.Duration) *base.ServiceInfo {
	return &base.ServiceInfo{
		Timeout: timeout,
		Host:    host,
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
		Scheme:      "https",
		Credentials: base.Credentials{Region: region, Service: "BusinessSecurity"},
	}
}

func (p *BusinessSecurity) Retry() bool {
	return p.retry
}

func (p *BusinessSecurity) CloseRetry() {
	p.retry = false
}

func (p *BusinessSecurity) SetRegion(region string) error {
	serviceInfo, ok := ServiceInfoMap[region]
	if !ok {
		return fmt.Errorf("region does not spport or unknown region")
	}
	p.ServiceInfo = serviceInfo
	p.SetScheme("http")
	return nil
}

var (
	ServiceInfoMap = map[string]*base.ServiceInfo{
		base.RegionCnNorth1: {
			Timeout: 5 * time.Second,
			Host:    "riskcontrol.volcengineapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionCnNorth1, Service: "BusinessSecurity"},
		},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		"RiskDetection": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RiskDetection"},
				"Version": []string{"2021-02-02"},
			},
		},
		"AsyncRiskDetection": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AsyncRiskDetection"},
				"Version": []string{"2021-02-25"},
			},
		},
		"RiskResult": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RiskResult"},
				"Version": []string{"2021-03-10"},
			},
		},
		"DataReport": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DataReport"},
				"Version": []string{"2021-08-31"},
			},
		},
		"TextRisk": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"TextRisk"},
				"Version": []string{"2022-01-26"},
			},
		},
		"AsyncVideoRisk": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AsyncVideoRisk"},
				"Version": []string{"2021-11-29"},
			},
		},
		"VideoResult": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"VideoResult"},
				"Version": []string{"2021-11-29"},
			},
		},
		"AsyncImageRisk": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AsyncImageRisk"},
				"Version": []string{"2021-11-29"},
			},
		},
		"AsyncImageRiskV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AsyncImageRisk"},
				"Version": []string{"2022-08-26"},
			},
		},
		"ImageContentRisk": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ImageContentRisk"},
				"Version": []string{"2021-11-29"},
			},
		},
		"ImageContentRiskV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ImageContentRiskV2"},
				"Version": []string{"2021-11-29"},
			},
		},
		"GetImageResult": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageResult"},
				"Version": []string{"2021-11-29"},
			},
		},
		"GetImageResultV2": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ImageResult"},
				"Version": []string{"2022-08-26"},
			},
		},
		"AsyncAudioRisk": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AsyncAudioRisk"},
				"Version": []string{"2022-04-01"},
			},
		},
		"GetAudioResult": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetAudioResult"},
				"Version": []string{"2022-04-01"},
			},
		},
		"AudioRisk": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AudioRisk"},
				"Version": []string{"2022-04-01"},
			},
		},
		"AsyncLiveVideoRisk": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AsyncLiveVideoRisk"},
				"Version": []string{"2022-04-25"},
			},
		},
		"GetVideoLiveResult": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetVideoLiveResult"},
				"Version": []string{"2022-04-25"},
			},
		},
		"CloseVideoLiveRisk": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CloseVideoLive"},
				"Version": []string{"2022-04-25"},
			},
		},
		"AsyncLiveAudioRisk": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AsyncLiveAudioRisk"},
				"Version": []string{"2022-04-25"},
			},
		},
		"GetAudioLiveResult": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetAudioLiveResult"},
				"Version": []string{"2022-04-25"},
			},
		},
		"CloseAudioLiveRisk": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CloseAudioLive"},
				"Version": []string{"2022-04-25"},
			},
		},
		"EnableCustomContents": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"EnableCustomContents"},
				"Version": []string{"2022-04-28"},
			},
		},
		"DisableCustomContents": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DisableCustomContents"},
				"Version": []string{"2022-04-28"},
			},
		},
		"CreateCustomContents": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateCustomContents"},
				"Version": []string{"2022-01-22"},
			},
		},
		"UploadCustomContents": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UploadCustomContents"},
				"Version": []string{"2022-02-07"},
			},
		},
		"DeleteCustomContents": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteCustomContents"},
				"Version": []string{"2022-04-28"},
			},
		},
		"ElementVerify": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ElementVerify"},
				"Version": []string{"2021-11-23"},
			},
		},
		"MobileStatus": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"MobileStatus"},
				"Version": []string{"2020-12-25"},
			},
		},
		"ElementVerifyV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ElementVerify"},
				"Version": []string{"2022-04-13"},
			},
		},
		"ElementVerifyEncrypted": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ElementVerifyEncrypted"},
				"Version": []string{"2022-11-24"},
			},
		},
		"MobileStatusV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"MobileStatus"},
				"Version": []string{"2022-04-13"},
			},
		},
		"TextSliceRisk": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"TextSliceRisk"},
				"Version": []string{"2022-11-07"},
			},
		},
		"TextAsyncRisk": {
			Method: http.MethodPost,
			Path:   "/open/api/v3/async_text_risk",
			Query: url.Values{
				"Action":  []string{"TextAsyncRisk"},
				"Version": []string{"2022-11-07"},
			},
		},
		"TextResult": {
			Method: http.MethodGet,
			Path:   "/open/api/v3/text_result",
			Query: url.Values{
				"Action":  []string{"TextResult"},
				"Version": []string{"2022-11-07"},
			},
		},
		"SimpleRiskStat": {
			Method:  http.MethodGet,
			Path:    "/",
			Timeout: 10 * time.Second,
			Query: url.Values{
				"Action":  []string{"SimpleRiskStat"},
				"Version": []string{"2022-12-23"},
			},
		},
		"ContentRiskStat": {
			Method:  http.MethodGet,
			Timeout: 10 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"ContentRiskStat"},
				"Version": []string{"2022-12-23"},
			},
		},
		"DelSystemNameListItem": {
			Method:  http.MethodPost,
			Timeout: 10 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"DelSystemNameListItem"},
				"Version": []string{"2022-12-23"},
			},
		},
		"QuerySystemNameListItem": {
			Method:  http.MethodGet,
			Timeout: 10 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"QuerySystemNameListItem"},
				"Version": []string{"2022-12-23"},
			},
		},
		"CreateCustomLib": {
			Method:  http.MethodPost,
			Timeout: 10 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"CreateCustomLib"},
				"Version": []string{"2023-10-01"},
			},
		},
		"UpdateCustomLib": {
			Method:  http.MethodPost,
			Timeout: 10 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"UpdateCustomLib"},
				"Version": []string{"2023-10-01"},
			},
		},
		"ChangeCustomContentsStatus": {
			Method:  http.MethodPost,
			Timeout: 10 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"ChangeCustomContentsStatus"},
				"Version": []string{"2023-10-01"},
			},
		},
		"DeleteCustomLib": {
			Method:  http.MethodPost,
			Timeout: 10 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"DeleteCustomLib"},
				"Version": []string{"2023-10-01"},
			},
		},
		"GetCustomLib": {
			Method:  http.MethodGet,
			Timeout: 10 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"GetCustomLib"},
				"Version": []string{"2023-10-01"},
			},
		},
		"CreateAccessConfig": {
			Method:  http.MethodPost,
			Timeout: 10 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"CreateAccessConfig"},
				"Version": []string{"2023-10-01"},
			},
		},
		"UpdateAccessConfig": {
			Method:  http.MethodPost,
			Timeout: 10 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"UpdateAccessConfig"},
				"Version": []string{"2023-10-01"},
			},
		},
		"UpdateConfigStatus": {
			Method:  http.MethodPost,
			Timeout: 10 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"UpdateConfigStatus"},
				"Version": []string{"2023-10-01"},
			},
		},
		"GetAccessConfig": {
			Method:  http.MethodGet,
			Timeout: 10 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"GetAccessConfig"},
				"Version": []string{"2023-10-01"},
			},
		},
		"GetTextLibContent": {
			Method:  http.MethodGet,
			Timeout: 10 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"GetTextLibContent"},
				"Version": []string{"2023-10-01"},
			},
		},
		"DeleteTextLibContent": {
			Method:  http.MethodPost,
			Timeout: 10 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"DeleteTextLibContent"},
				"Version": []string{"2023-10-01"},
			},
		},
		"UploadTextLibContent": {
			Method:  http.MethodPost,
			Timeout: 10 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"UploadTextLibContent"},
				"Version": []string{"2023-10-01"},
			},
		},
		"GetImageLibContent": {
			Method:  http.MethodGet,
			Timeout: 10 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"GetImageLibContent"},
				"Version": []string{"2023-10-01"},
			},
		},
		"DeleteImageLibContent": {
			Method:  http.MethodPost,
			Timeout: 10 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"DeleteImageLibContent"},
				"Version": []string{"2023-10-01"},
			},
		},
		"UploadImageLibContent": {
			Method:  http.MethodPost,
			Timeout: 10 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"UploadImageLibContent"},
				"Version": []string{"2023-10-01"},
			},
		},
		"CreateApp": {
			Method:  http.MethodPost,
			Timeout: 10 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"CreateApp"},
				"Version": []string{"2022-12-23"},
			},
		},
		"ListApps": {
			Method:  http.MethodGet,
			Timeout: 10 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"ListApps"},
				"Version": []string{"2022-12-23"},
			},
		},
		"ActivateRiskSampleData": {
			Method:  http.MethodPost,
			Timeout: 10 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"ActivateRiskSampleData"},
				"Version": []string{"2023-10-01"},
			},
		},
		"ActivateRiskBasePackage": {
			Method:  http.MethodPost,
			Timeout: 10 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"ActivateRiskBasePackage"},
				"Version": []string{"2023-10-01"},
			},
		},
		"ActivateRiskResult": {
			Method:  http.MethodPost,
			Timeout: 3 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"ActivateRiskResult"},
				"Version": []string{"2023-10-01"},
			},
		},
		"CancelActivateRiskResult": {
			Method:  http.MethodPost,
			Timeout: 3 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"CancelActivateRiskResult"},
				"Version": []string{"2023-10-01"},
			},
		},
		"GetTextStatisticsOpen": {
			Method:  http.MethodPost,
			Timeout: 5 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"GetTextStatisticsOpen"},
				"Version": []string{"2022-12-23"},
			},
		},
		"GetImageStatisticsOpen": {
			Method:  http.MethodPost,
			Timeout: 5 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"GetImageStatisticsOpen"},
				"Version": []string{"2022-12-23"},
			},
		},
		"GetVideoStatisticsOpen": {
			Method:  http.MethodPost,
			Timeout: 5 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"GetVideoStatisticsOpen"},
				"Version": []string{"2022-12-23"},
			},
		},
		"GetAudioLiveStatisticsOpen": {
			Method:  http.MethodPost,
			Timeout: 5 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"GetAudioLiveStatisticsOpen"},
				"Version": []string{"2022-12-23"},
			},
		},
		"GetVideoLiveStatisticsOpen": {
			Method:  http.MethodPost,
			Timeout: 5 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"GetVideoLiveStatisticsOpen"},
				"Version": []string{"2022-12-23"},
			},
		},
		"GetAudioStatisticsOpen": {
			Method:  http.MethodPost,
			Timeout: 5 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"GetAudioStatisticsOpen"},
				"Version": []string{"2022-12-23"},
			},
		},
		"OpenProduct": {
			Method:  http.MethodPost,
			Timeout: 10 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"OpenProduct"},
				"Version": []string{"2022-12-23"},
			},
		},
		"CheckProductStatus": {
			Method:  http.MethodGet,
			Timeout: 10 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"CheckProductStatus"},
				"Version": []string{"2022-12-23"},
			},
		},
		"EnableService": {
			Method:  http.MethodPost,
			Timeout: 10 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"EnableService"},
				"Version": []string{"2022-12-23"},
			},
		},
		"CheckServiceStatus": {
			Method:  http.MethodGet,
			Timeout: 10 * time.Second,
			Path:    "/",
			Query: url.Values{
				"Action":  []string{"CheckServiceStatus"},
				"Version": []string{"2022-12-23"},
			},
		},
	}
)
