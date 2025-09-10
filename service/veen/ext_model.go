package veen

// 官方没有的  写在这里
// 批量创建弹性公网IP===================================
type BatchCreateEIPInstancesReq struct {
	// 弹性公网 IP 所属的项目。如果不指定该参数或参数值为空字符串，采用默认值 default。
	Project string `json:"project" query:"project"`
	// 弹性公网 IP 的需求列表
	EipRequirementList []EipRequirementItem `json:"eip_requirement_list" query:"eip_requirement_list" validate:"required"`
	// 弹性公网 IP 的名称。命名规则如下：
	// 允许 5~20 个字符。
	// 支持汉字、大写字母、小写字母、数字和特殊字符 ()`~!@#$%^&*-+=_|{}[]:;'<>,.?/。
	// 不能包含双引号（"）、反斜线（ \）和 空格，且不能以正斜线（/）开头。
	// 系统会自动在您指定的名称后面添加节点信息，最终生成的名称的结构为：<指定的名称>-<节点名称>。当您创建多个弹性公网IP时，系统还会为名称添加数字后缀、以区分不同的弹性公网IP。第一个创建出来的实例的名称不会被添加数字后缀。例如，如果您在北京电信节点下创建了3个弹性公网IP并指定名称为“dev_a”，最终系统生成的名称将分别是dev_a-**bjct、dev_a-**bjct-1和dev_a-**bjct-2。
	EipName string `json:"eip_name,omitempty" query:"eip_name"`
	// 弹性公网 IP 的类型。取值范围：IPv4：IPv4 类型。 IPv6：IPv6 类型。  IPv6 类型的弹性公网 IP 只能绑定到支持 IPv4/IPv6 双栈的负载均衡实例或边缘实例。
	EipType string `json:"eip_type" query:"eip_type" validate:"required"`
	// 弹性公网 IP 的网络类型。取值范围：public：公网类型
	NetworkType string `json:"network_type" query:"network_type" validate:"required"`
	// 带宽峰值。取值范围：20 ~ 5000。单位：Mbps。 取值只能是 5 的倍数。
	BandwidthPeak int `json:"bandwidth_peak" query:"bandwidth_peak" validate:"required"`
	// 弹性公网 IP 的描述。最多可输入 80 个字符。
	Desc string `json:"desc" query:"desc"`
}

type EipRequirementItem struct {
	// 弹性公网 IP 所在节点的名称。您可以调用 ListAvailableResourceInfo 接口查询支持的节点的信息。
	ClusterName string `json:"cluster_name" query:"cluster_name" validate:"required"`
	// 线路类型。取值范围：CMCC：中国移动。CUCC：中国联通。CTCC：中国电信
	// 当您在多线节点创建弹性公网 IP 时，您可以设置该参数来指定线路类型。例如，如果您将参数值设置为 CMCC，那么系统仅会创建一个中国移动线路的弹性公网 IP。
	// 如果您指定了 ip_pool_id 参数，isp 参数无需指定。
	ISP string `json:"isp" query:"isp"`
	// 弹性公网 IP 的数量。 eip_num 和 ip_addrs 参数不能同时指定。
	EipNum int `json:"eip_num" query:"eip_num"`
	// IP 地址池 ID。您可以调用 ListIpPools 接口获取 IP 地址池的 ID。 ip_pool_id 和 ip_addrs 参数必须搭配使用。
	IpPoolID string `json:"ip_pool_id" query:"ip_pool_id"`
	// 需要从 IP 地址池中分配的 IP 地址的列表
	// ip_addrs 和 ip_pool_id 参数必须搭配使用。
	// eip_num 和 ip_addrs 参数不能同时指定
	IpAddrs []string `json:"ip_addrs" query:"ip_addrs"`
}

type BatchCreateEIPInstancesResp struct {
	ResponseMetadata VolcResponseMetadata          `json:"ResponseMetadata"`
	Result           BatchCreateEIPInstancesResult `json:"Result"`
}

type BatchCreateEIPInstancesResult struct {
	Eips []BatchCreateEIPInstancesItem `json:"eips" query:"eips"` // 弹性公网 IP 的列表
}

type BatchCreateEIPInstancesItem struct {
	// 账号 ID。
	AccountIdentity int `json:"account_identity"`
	// 子用户 ID
	UserIdentity int `json:"user_identity"`
	// 弹性公网 IP 的 ID
	EipIdentity string `json:"eip_identity"`
	// 弹性公网 IP 的名称
	EipName string `json:"eip_name"`
	// 弹性公网 IP 所在的节点
	Cluster Cluster `json:"cluster"`
	// 弹性公网 IP 的类型：IPv4   IPv6
	EipType string `json:"eip_type"`
	// 弹性公网 IP 的地址
	EipAddr string `json:"eip_addr"`
	// 弹性公网 IP 的带宽峰值。
	BandwidthPeak int `json:"bandwidth_peak"`
	// 弹性公网 IP 的线路类型。
	Isp string `json:"isp"`
	// 弹性公网 IP 的网络类型  public：公网
	NetworkType string `json:"network_type"`
	// 弹性公网 IP 绑定的资源。
	BinderResource BinderResource `json:"binder_resource"`
	// 弹性公网 IP 的状态：
	//creating：创建中。
	//unbound：未绑定。
	//binding：绑定中。
	//bound：已绑定。
	//unbinding：解绑中。
	//deleting：删除中。
	Status string `json:"status"`
	// 弹性公网 IP 的描述
	Desc string `json:"desc"`
	// 弹性公网 IP 实例计费配置
	BillingConfig BillingConfig `json:"billing_config"`
	// 1694505600
	CreateTime int `json:"create_time"`
	UpdateTime int `json:"update_time"`
}

// 弹性公网 IP 所在的节点。
type Cluster struct {
	// 弹性公网 IP 所属节点的名称。
	ClusterName string `json:"cluster_name"`
	// 弹性公网 IP 所属节点的国家。
	Country string `json:"country"`
	// 弹性公网 IP 所属节点的区域。
	Region string `json:"region"`
	// 弹性公网 IP 所属节点的省份 zj
	Province string `json:"province"`
	// 弹性公网 IP 所属节点的城市
	City string `json:"city"`
	// 弹性公网 IP 所属节点的线路类型 CTCC
	Isp string `json:"isp"`
	// 弹性公网 IP 所属节点的计费类别：
	// 1：北上广
	// 2：地区中心
	// 3：一般城市
	Level string `json:"level"`
	// 弹性公网 IP 所属节点的别名  浙江宁波电信01
	Alias string `json:"alias"`
}

type BinderResource struct {
	// 弹性公网 IP 绑定的资源的类型：
	// lb：四层负载均衡实例。
	// lb7：七层负载均衡实例。
	// veen：边缘实例。
	ResourceType string `json:"resource_type"`
	// 弹性公网 IP 绑定的资源的 ID
	ResourceIdentity string `json:"resource_identity"`
	// 弹性公网 IP 绑定的资源的名称
	ResourceName string `json:"resource_name"`
	// 弹性公网 IP 绑定的资源所属私有网络的 ID
	ResourceVpcIdentity string `json:"resource_vpc_identity"`
	// 弹性公网 IP 绑定的资源所属私有网络的名称  宁波电信01-默认VPC
	ResourceVpcName string `json:"resource_vpc_name"`
	// 弹性公网 IP 绑定的资源所属子网的 CIDR 地址段。 ["172.**.**.192/26"]
	ResourceVpcSubnetCidrs []string `json:"resource_vpc_subnet_cidrs"`
}

// 弹性公网 IP 实例计费配置
type BillingConfig struct {
	// IP 计费方式：MonthlyPeak：按月计费。 DailyPeak：按日计费。如需使用该计费方式，请提交工单申请。
	IPBillingMethod string `json:"ip_billing_method"`
	// 带宽计费方式：
	// MonthlyP95：按月 95 峰值计费。
	// DailyPeak：按日峰值计费。如需使用该计费方式，请提交工单申请。
	// HourUsed：按流量计费。如需使用该计费方式，请提交工单申请。
	BandwidthBillingMethod string `json:"bandwidth_billing_method"`
}

// 释放弹性公网IP ===========================
type DeleteEIPInstanceReq struct {
	// 弹性公网 IP 的 ID。您可以通过 ListEIPInstances 接口查询弹性公网 IP 的 ID。
	EipIdentity string `json:"eip_identity" query:"eip_identity" validate:"required"`
}
type DeleteEIPInstanceResp struct {
	ResponseMetadata VolcResponseMetadata    `json:"ResponseMetadata"`
	Result           DeleteEIPInstanceResult `json:"Result"`
}

type DeleteEIPInstanceResult struct{}

// 修改弹性公网IP名称 ===========================
type SetEIPInstanceNameReq struct {
	// 弹性公网 IP 的 ID。您可以通过 ListEIPInstances 接口查询弹性公网 IP 的 ID。
	EipIdentity string `json:"eip_identity" query:"eip_identity" validate:"required"`
	// 弹性公网 IP 的名称。命名规则如下：
	// 允许 5~20 个字符。
	// 支持中文、大写字母、小写字母、数字。
	// 支持特殊字符 ()`~!@#$%^&*-+=_|{}[]:;'<>,.?/。 |
	// 不能包含双引号（"）、反斜线（ \）和 空格，且不能以正斜线（/）开头。
	Name string `json:"name" query:"name" validate:"required"`
}
type SetEIPInstanceNameResp struct {
	ResponseMetadata VolcResponseMetadata     `json:"ResponseMetadata"`
	Result           SetEIPInstanceNameResult `json:"Result"`
}

type SetEIPInstanceNameResult struct{}

// 修改弹性公网IP描述 ===========================
type SetEIPInstanceDescReq struct {
	// 弹性公网 IP 的 ID。您可以通过 ListEIPInstances 接口查询弹性公网 IP 的 ID。
	EipIdentity string `json:"eip_identity" query:"eip_identity" validate:"required"`
	// 弹性公网 IP 的描述。最多可输入 80 个字符。
	// 如果不指定该参数或参数值为空字符串，原来的描述将被清空
	Desc string `json:"desc" query:"desc"`
}
type SetEIPInstanceDescResp struct {
	ResponseMetadata VolcResponseMetadata     `json:"ResponseMetadata"`
	Result           SetEIPInstanceDescResult `json:"Result"`
}

type SetEIPInstanceDescResult struct{}

// 修改弹性公网IP带宽 ===========================
type SetEIPInstanceBandwidthPeakReq struct {
	// 弹性公网 IP 的 ID。您可以通过 ListEIPInstances 接口查询弹性公网 IP 的 ID。
	EipIdentity string `json:"eip_identity" query:"eip_identity" validate:"required"`
	// 带宽峰值。取值范围：20 ~ 5000。单位：Mbps。
	// 取值只能是 5 的倍数。
	BandwidthPeak int `json:"bandwidth_peak" query:"bandwidth_peak" validate:"required"`
}
type SetEIPInstanceBandwidthPeakResp struct {
	ResponseMetadata VolcResponseMetadata              `json:"ResponseMetadata"`
	Result           SetEIPInstanceBandwidthPeakResult `json:"Result"`
}

type SetEIPInstanceBandwidthPeakResult struct{}

// 绑定边缘实例 ===========================
type BindEIPToVeenReq struct {
	// 弹性公网 IP 的 ID。您可以通过 ListEIPInstances 接口查询弹性公网 IP 的 ID。
	EipIdentity string `json:"eip_identity" query:"eip_identity" validate:"required"`
	// 边缘实例的 ID。您可以通过 ListInstances 接口查询边缘实例的 ID
	VeenIdentity string `json:"veen_identity" query:"veen_identity" validate:"required"`
}
type BindEIPToVeenResp struct {
	ResponseMetadata VolcResponseMetadata `json:"ResponseMetadata"`
	Result           BindEIPToVeenResult  `json:"Result"`
}

type BindEIPToVeenResult struct{}

// 解除已绑定的边缘实例 ===========================
type UnbindEIPFromVeenReq struct {
	// 弹性公网 IP 的 ID。您可以通过 ListEIPInstances 接口查询弹性公网 IP 的 ID。
	EipIdentity string `json:"eip_identity" query:"eip_identity" validate:"required"`
	// 边缘实例的 ID。您可以通过 ListInstances 接口查询边缘实例的 ID
	VeenIdentity string `json:"veen_identity" query:"veen_identity" validate:"required"`
}
type UnbindEIPFromVeenResp struct {
	ResponseMetadata VolcResponseMetadata    `json:"ResponseMetadata"`
	Result           UnbindEIPFromVeenResult `json:"Result"`
}

type UnbindEIPFromVeenResult struct{}

// 获取弹性公网IP列表 ===========================
type ListEIPInstancesReq struct {
	// 弹性公网 IP 的 ID 列表。ID 之间用半角逗号（,）分隔
	EipIdentityList string `json:"eip_identity_list" query:"eip_identity_list"`
	// 节点名称。节点之间用半角逗号（,）分隔。
	ClusterNames string `json:"cluster_names" query:"cluster_names"`
	// 是否列出绑定的实例。
	// true：列出绑定的实例。
	// false（默认值）：不列出绑定的实例。
	WithBinderInfo *bool `json:"with_binder_info" query:"with_binder_info"`
	Pagination
}
type ListEIPInstancesResp struct {
	ResponseMetadata VolcResponseMetadata   `json:"ResponseMetadata"`
	Result           ListEIPInstancesResult `json:"Result"`
}

type ListEIPInstancesResult struct {
	EipInstances []*EIPInstance `json:"eip_instances"` // 弹性公网 IP 的列表
	TotalCount   int64          `json:"count"`         // 弹性公网 IP 的数量
}

type EIPInstance struct {
	AccountIdentity string         `json:"account_identity"`
	UserIdentity    int            `json:"user_identity"`
	EipIdentity     string         `json:"eip_identity"`
	EipName         string         `json:"eip_name"`
	Cluster         Cluster        `json:"cluster"`
	EipType         string         `json:"eip_type"`
	EipAddr         string         `json:"eip_addr"`
	BandwidthPeak   int            `json:"bandwidth_peak"`
	Isp             string         `json:"isp"`
	NetworkType     string         `json:"network_type"`
	BinderResource  BinderResource `json:"binder_resource"`
	Status          string         `json:"status"`
	Desc            string         `json:"desc"`
	BillingConfig   BillingConfig  `json:"billing_config"`
	ClusterBwpID    int            `json:"cluster_bwp_id"`
	Project         string         `json:"project"`
	CreateTime      int            `json:"create_time"`
	UpdateTime      int            `json:"update_time"`
}

// 获取弹性公网IP详情 ===========================
type GetEIPInstanceReq struct {
	//  弹性公网 IP 的 ID
	EipIdentity string `json:"eip_identity" query:"eip_identity"`
	// 是否列出绑定的实例。
	// true：列出绑定的实例。
	// false（默认值）：不列出绑定的实例。
	WithBinderInfo *bool `json:"with_binder_info" query:"with_binder_info"`
}
type GetEIPInstanceResp struct {
	ResponseMetadata VolcResponseMetadata  `json:"ResponseMetadata"`
	Result           GetEIPInstancesResult `json:"Result"`
}

type GetEIPInstancesResult struct {
	Eip EIPInstance `json:"eip"`
}
