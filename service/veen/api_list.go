package veen

import (
	"github.com/volcengine/volc-sdk-golang/base"
	"net/http"
	"net/url"
)

var (
	// 注册api
	ApiInfoList = map[string]*base.ApiInfo{
		"CreateCloudServer": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateCloudServer"},
				"Version": []string{ServiceVersion},
			},
		},
		"ListCloudServers": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListCloudServers"},
				"Version": []string{ServiceVersion},
			},
		},
		"GetCloudServer": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetCloudServer"},
				"Version": []string{ServiceVersion},
			},
		},
		"StartCloudServer": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartCloudServer"},
				"Version": []string{ServiceVersion},
			},
		},
		"StopCloudServer": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopCloudServer"},
				"Version": []string{ServiceVersion},
			},
		},
		"RebootCloudServer": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RebootCloudServer"},
				"Version": []string{ServiceVersion},
			},
		},
		"DeleteCloudServer": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteCloudServer"},
				"Version": []string{ServiceVersion},
			},
		},
		"ListInstanceTypes": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListInstanceTypes"},
				"Version": []string{ServiceVersion},
			},
		},
		"ListAvailableResourceInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListAvailableResourceInfo"},
				"Version": []string{ServiceVersion},
			},
		},
		"CreateInstance": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateInstance"},
				"Version": []string{ServiceVersion},
			},
		},
		"ListInstances": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListInstances"},
				"Version": []string{ServiceVersion},
			},
		},
		"GetInstance": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetInstance"},
				"Version": []string{ServiceVersion},
			},
		},
		"StartInstances": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartInstances"},
				"Version": []string{ServiceVersion},
			},
		},
		"StopInstances": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopInstances"},
				"Version": []string{ServiceVersion},
			},
		},
		"RebootInstances": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RebootInstances"},
				"Version": []string{ServiceVersion},
			},
		},
		"OfflineInstances": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"OfflineInstances"},
				"Version": []string{ServiceVersion},
			},
		},
		"SetInstanceName": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SetInstanceName"},
				"Version": []string{ServiceVersion},
			},
		},
		"ResetLoginCredential": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ResetLoginCredential"},
				"Version": []string{ServiceVersion},
			},
		},
		"GetInstanceCloudDiskInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetInstanceCloudDiskInfo"},
				"Version": []string{ServiceVersion},
			},
		},
		"ScaleInstanceCloudDiskCapacity": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ScaleInstanceCloudDiskCapacity"},
				"Version": []string{ServiceVersion},
			},
		},
		"CreateEbsInstances": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateEbsInstances"},
				"Version": []string{ServiceVersion},
			},
		},
		"ListEbsInstances": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListEbsInstances"},
				"Version": []string{ServiceVersion},
			},
		},
		"GetEbsInstance": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetEbsInstance"},
				"Version": []string{ServiceVersion},
			},
		},
		"ScaleEbsInstanceCapacity": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ScaleEbsInstanceCapacity"},
				"Version": []string{ServiceVersion},
			},
		},
		"AttachEbs": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AttachEbs"},
				"Version": []string{ServiceVersion},
			},
		},
		"DetachEbs": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DetachEbs"},
				"Version": []string{ServiceVersion},
			},
		},
		"DeleteEbsInstance": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteEbsInstance"},
				"Version": []string{ServiceVersion},
			},
		},
		"BatchResetSystem": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BatchResetSystem"},
				"Version": []string{ServiceVersion},
			},
		},

		// 扩展 api 接口
		"BatchCreateEIPInstances": { // 批量创建弹性公网IP
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BatchCreateEIPInstances"},
				"Version": []string{ServiceVersion},
			},
		},
		"DeleteEIPInstance": { // 释放弹性公网IP
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteEIPInstance"},
				"Version": []string{ServiceVersion},
			},
		},
		"SetEIPInstanceName": { // 修改弹性公网IP名称
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SetEIPInstanceName"},
				"Version": []string{ServiceVersion},
			},
		},
		"SetEIPInstanceDesc": { // 修改弹性公网IP描述
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SetEIPInstanceDesc"},
				"Version": []string{ServiceVersion},
			},
		},
		"SetEIPInstanceBandwidthPeak": { // 修改弹性公网IP带宽
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SetEIPInstanceBandwidthPeak"},
				"Version": []string{ServiceVersion},
			},
		},
		"BindEIPToVeen": { // 绑定边缘实例
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BindEIPToVeen"},
				"Version": []string{ServiceVersion},
			},
		},
		"UnbindEIPFromVeen": { // 解除已绑定的边缘实例
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UnbindEIPFromVeen"},
				"Version": []string{ServiceVersion},
			},
		},
		"ListEIPInstances": { // 获取弹性公网IP列表
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListEIPInstances"},
				"Version": []string{ServiceVersion},
			},
		},
		"GetEIPInstance": { // 获取弹性公网IP详情
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetEIPInstance"},
				"Version": []string{ServiceVersion},
			},
		},
		"ListVPCInstances": { // 获取私有网络列表
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListVPCInstances"},
				"Version": []string{ServiceVersion},
			},
		},
		"SetVPCInstanceName": { // 修改私有网络名称
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SetVPCInstanceName"},
				"Version": []string{ServiceVersion},
			},
		},
		"SetVPCInstanceDesc": { // 修改私有网络描述
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SetVPCInstanceDesc"},
				"Version": []string{ServiceVersion},
			},
		},
		"CreateCustomVPCInstance": { // 创建自定义私有网络
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateCustomVPCInstance"},
				"Version": []string{ServiceVersion},
			},
		},
		"DeleteCustomVPCInstances": { // 删除自定义私有网络
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteCustomVPCInstances"},
				"Version": []string{ServiceVersion},
			},
		},
		"ListSubnetInstances": { // 获取子网列表
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListSubnetInstances"},
				"Version": []string{ServiceVersion},
			},
		},
		"CreateSubnetsForCustomVPC": { // 创建子网
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateSubnetsForCustomVPC"},
				"Version": []string{ServiceVersion},
			},
		},
		"SetSubnetNameAndDesc": { // 修改子网名称和描述
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SetSubnetNameAndDesc"},
				"Version": []string{ServiceVersion},
			},
		},
		"DeleteSubnetsForCustomVPC": { // 删除子网
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteSubnetsForCustomVPC"},
				"Version": []string{ServiceVersion},
			},
		},
		"BindEipToInternalIP": { // 绑定单个弹性公网 IP 到私网 IP 地址
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BindEipToInternalIP"},
				"Version": []string{ServiceVersion},
			},
		},
		"SetBoundEipShareBandwidthPeak": { // 设置弹性公网 IP 的共享带宽峰值
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SetBoundEipShareBandwidthPeak"},
				"Version": []string{ServiceVersion},
			},
		},
		"ListInstanceInternalIps": { // 获取私网 IP 地址列表
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListInstanceInternalIps"},
				"Version": []string{ServiceVersion},
			},
		},
		"BatchUnbindEipFromInternalIP": { // 批量解绑弹性公网 IP
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BatchUnbindEipFromInternalIP"},
				"Version": []string{ServiceVersion},
			},
		},
	}
)
