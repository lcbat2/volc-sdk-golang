package veen

import (
	"errors"
	"net/url"
	"strconv"
)

func (v *Veen) BatchCreateEIPInstances(req *BatchCreateEIPInstancesReq) (*BatchCreateEIPInstancesResp, error) {
	resp := &BatchCreateEIPInstancesResp{}
	if err := v.post("BatchCreateEIPInstances", req, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

func (v *Veen) DeleteEIPInstance(req *DeleteEIPInstanceReq) (*DeleteEIPInstanceResp, error) {
	resp := &DeleteEIPInstanceResp{}
	if err := v.post("DeleteEIPInstance", req, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}
func (v *Veen) SetEIPInstanceName(req *SetEIPInstanceNameReq) (*SetEIPInstanceNameResp, error) {
	resp := &SetEIPInstanceNameResp{}
	if err := v.post("SetEIPInstanceName", req, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}
func (v *Veen) SetEIPInstanceDesc(req *SetEIPInstanceDescReq) (*SetEIPInstanceDescResp, error) {
	resp := &SetEIPInstanceDescResp{}
	if err := v.post("SetEIPInstanceDesc", req, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}
func (v *Veen) SetEIPInstanceBandwidthPeak(req *SetEIPInstanceBandwidthPeakReq) (*SetEIPInstanceBandwidthPeakResp, error) {
	resp := &SetEIPInstanceBandwidthPeakResp{}
	if err := v.post("SetEIPInstanceBandwidthPeak", req, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

func (v *Veen) BindEIPToVeen(req *BindEIPToVeenReq) (*BindEIPToVeenResp, error) {
	resp := &BindEIPToVeenResp{}
	if err := v.post("BindEIPToVeen", req, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}
func (v *Veen) UnbindEIPFromVeen(req *UnbindEIPFromVeenReq) (*UnbindEIPFromVeenResp, error) {
	resp := &UnbindEIPFromVeenResp{}
	if err := v.post("UnbindEIPFromVeen", req, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}
func (v *Veen) ListEIPInstances(req *ListEIPInstancesReq) (*ListEIPInstancesResp, error) {
	resp := &ListEIPInstancesResp{}
	query := url.Values{}
	query.Set("page", strconv.Itoa(int(req.Page)))
	query.Set("limit", strconv.Itoa(int(req.Limit)))
	query.Set("order_by", strconv.Itoa(int(req.OrderBy)))
	if req.EipIdentityList != "" {
		query.Set("eip_identity_list", req.EipIdentityList)
	}
	if req.ClusterNames != "" {
		query.Set("cluster_names", req.ClusterNames)
	}
	if req.WithBinderInfo != nil && *req.WithBinderInfo {
		query.Set("with_binder_info", "true")
	} else if req.WithBinderInfo != nil && !*req.WithBinderInfo {
		query.Set("with_binder_info", "false")
	}
	if err := v.get("ListEIPInstances", query, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}
func (v *Veen) GetEIPInstance(req *GetEIPInstanceReq) (*GetEIPInstanceResp, error) {
	resp := &GetEIPInstanceResp{}
	query := url.Values{}
	if req.EipIdentity != "" {
		query.Set("eip_identity", req.EipIdentity)
	}
	if req.WithBinderInfo != nil && *req.WithBinderInfo {
		query.Set("with_binder_info", "true")
	} else if req.WithBinderInfo != nil && !*req.WithBinderInfo {
		query.Set("with_binder_info", "false")
	}
	if err := v.get("GetEIPInstance", query, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

func (v *Veen) ListVPCInstances(req *ListVPCInstancesReq) (*ListVPCInstancesResp, error) {
	resp := &ListVPCInstancesResp{}
	query := url.Values{}
	query.Set("page", strconv.Itoa(int(req.Page)))
	query.Set("limit", strconv.Itoa(int(req.Limit)))
	query.Set("order_by", strconv.Itoa(int(req.OrderBy)))
	if req.VpcIdentityList != "" {
		query.Set("vpc_identity_list", req.VpcIdentityList)
	}
	if req.ClusterNames != "" {
		query.Set("cluster_names", req.ClusterNames)
	}
	if req.IsVlanVpc {
		query.Set("is_vlan_vpc", "true")
	} else {
		query.Set("is_vlan_vpc", "false")
	}
	if req.IsDefaultVpc {
		query.Set("is_default_vpc", "true")
	} else {
		query.Set("is_default_vpc", "false")
	}
	if req.IsCustomVpc {
		query.Set("is_custom_vpc", "true")
	} else {
		query.Set("is_custom_vpc", "false")
	}
	if req.WithResourceStatistic {
		query.Set("with_resource_statistic", "true")
	} else {
		query.Set("with_resource_statistic", "false")
	}
	if err := v.get("ListVPCInstances", query, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

func (v *Veen) CreateCustomVPCInstance(req *CreateCustomVPCInstanceReq) (*CreateCustomVPCInstanceResp, error) {
	resp := &CreateCustomVPCInstanceResp{}
	if err := v.post("CreateCustomVPCInstance", req, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

func (v *Veen) SetVPCInstanceName(req *SetVPCInstanceNameReq) (*SetVPCInstanceNameResp, error) {
	resp := &SetVPCInstanceNameResp{}
	if err := v.post("SetVPCInstanceName", req, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}
func (v *Veen) SetVPCInstanceDesc(req *SetVPCInstanceDescReq) (*SetVPCInstanceDescResp, error) {
	resp := &SetVPCInstanceDescResp{}
	if err := v.post("SetVPCInstanceDesc", req, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

func (v *Veen) DeleteCustomVPCInstances(req *DeleteCustomVPCInstancesReq) (*DeleteCustomVPCInstancesResp, error) {
	resp := &DeleteCustomVPCInstancesResp{}
	if err := v.post("DeleteCustomVPCInstances", req, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

func (v *Veen) ListSubnetInstances(req *ListSubnetInstancesReq) (*ListSubnetInstancesResp, error) {
	resp := &ListSubnetInstancesResp{}
	query := url.Values{}
	query.Set("page", strconv.Itoa(int(req.Page)))
	query.Set("limit", strconv.Itoa(int(req.Limit)))
	query.Set("order_by", strconv.Itoa(int(req.OrderBy)))
	if req.SubnetIdentityList != "" {
		query.Set("subnet_identity_list", req.SubnetIdentityList)
	}
	if req.ClusterNames != "" {
		query.Set("cluster_names", req.ClusterNames)
	}
	if req.StatusList != "" {
		query.Set("status_list", req.StatusList)
	}
	if err := v.get("ListSubnetInstances", query, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

func (v *Veen) CreateSubnetsForCustomVPC(req *CreateSubnetsForCustomVPCReq) (*CreateSubnetsForCustomVPCResp, error) {
	resp := &CreateSubnetsForCustomVPCResp{}
	if err := v.post("CreateSubnetsForCustomVPC", req, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}
func (v *Veen) SetSubnetNameAndDesc(req *SetSubnetNameAndDescReq) (*SetSubnetNameAndDescResp, error) {
	resp := &SetSubnetNameAndDescResp{}
	if err := v.post("SetSubnetNameAndDesc", req, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

func (v *Veen) DeleteSubnetsForCustomVPC(req *DeleteSubnetsForCustomVPCReq) (*DeleteSubnetsForCustomVPCResp, error) {
	resp := &DeleteSubnetsForCustomVPCResp{}
	if err := v.post("DeleteSubnetsForCustomVPC", req, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

// 绑定单个弹性公网 IP 到私网 IP 地址
func (v *Veen) BindEipToInternalIP(req *BindEipToInternalIPReq) (*BindEipToInternalIPResp, error) {
	resp := &BindEipToInternalIPResp{}
	if err := v.post("BindEipToInternalIP", req, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

// 设置弹性公网 IP 的共享带宽峰值
func (v *Veen) SetBoundEipShareBandwidthPeak(req *SetBoundEipShareBandwidthPeakReq) (*SetBoundEipShareBandwidthPeakResp, error) {
	resp := &SetBoundEipShareBandwidthPeakResp{}
	if err := v.post("SetBoundEipShareBandwidthPeak", req, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

// 获取私网 IP 地址列表
func (v *Veen) ListInstanceInternalIps(req *ListInstanceInternalIpsReq) (*ListInstanceInternalIpsResp, error) {
	resp := &ListInstanceInternalIpsResp{}
	if req.InstanceIdentity == "" {
		return nil, errors.New("instance_identity is required")
	}
	query := url.Values{}
	query.Set("instance_identity", req.InstanceIdentity)
	if err := v.get("ListInstanceInternalIps", query, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}

// 批量解绑弹性公网 IP
func (v *Veen) BatchUnbindEipFromInternalIP(req *BatchUnbindEipFromInternalIPReq) (*BatchUnbindEipFromInternalIPResp, error) {
	resp := &BatchUnbindEipFromInternalIPResp{}
	if err := v.post("BatchUnbindEipFromInternalIP", req, resp); err != nil {
		return nil, err
	}
	if resp.ResponseMetadata.Error != nil {
		return nil, packErrorInfo(resp.ResponseMetadata)
	}
	return resp, nil
}
