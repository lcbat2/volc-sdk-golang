package veen

import (
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
