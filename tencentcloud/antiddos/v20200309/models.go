// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20200309

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AnycastOutPackRelation struct {
	// Application bandwidth (in Mbps).
	// Note: This field may return null, indicating that no valid values can be obtained.
	NormalBandwidth *uint64 `json:"NormalBandwidth,omitnil" name:"NormalBandwidth"`

	// Number of forwarding rules
	// Note: This field may return null, indicating that no valid values can be obtained.
	ForwardRulesLimit *uint64 `json:"ForwardRulesLimit,omitnil" name:"ForwardRulesLimit"`

	// Auto-renewal flag
	// Note: This field may return null, indicating that no valid values can be obtained.
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// Expiration date
	// Note: This field may return null, indicating that no valid values can be obtained.
	CurDeadline *string `json:"CurDeadline,omitnil" name:"CurDeadline"`
}

// Predefined struct for user
type AssociateDDoSEipAddressRequestParams struct {
	// Anti-DDoS instance ID (only Anti-DDoS Advanced). For example, `bgpip-0000011x`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// EIP of the Anti-DDoS instance ID
	Eip *string `json:"Eip,omitnil" name:"Eip"`

	// Instance ID to bind. For example, `ins-11112222`. It can be queried in the console or obtained from `InstanceId` returned by `DescribeInstances`.
	CvmInstanceID *string `json:"CvmInstanceID,omitnil" name:"CvmInstanceID"`

	// Region of the CVM instance. For example, `ap-hongkong`.
	CvmRegion *string `json:"CvmRegion,omitnil" name:"CvmRegion"`
}

type AssociateDDoSEipAddressRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS instance ID (only Anti-DDoS Advanced). For example, `bgpip-0000011x`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// EIP of the Anti-DDoS instance ID
	Eip *string `json:"Eip,omitnil" name:"Eip"`

	// Instance ID to bind. For example, `ins-11112222`. It can be queried in the console or obtained from `InstanceId` returned by `DescribeInstances`.
	CvmInstanceID *string `json:"CvmInstanceID,omitnil" name:"CvmInstanceID"`

	// Region of the CVM instance. For example, `ap-hongkong`.
	CvmRegion *string `json:"CvmRegion,omitnil" name:"CvmRegion"`
}

func (r *AssociateDDoSEipAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateDDoSEipAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Eip")
	delete(f, "CvmInstanceID")
	delete(f, "CvmRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateDDoSEipAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateDDoSEipAddressResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AssociateDDoSEipAddressResponse struct {
	*tchttp.BaseResponse
	Response *AssociateDDoSEipAddressResponseParams `json:"Response"`
}

func (r *AssociateDDoSEipAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateDDoSEipAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateDDoSEipLoadBalancerRequestParams struct {
	// Anti-DDoS instance ID (only Anti-DDoS Advanced). For example, `bgpip-0000011x`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// EIP of the Anti-DDoS instance ID.
	Eip *string `json:"Eip,omitnil" name:"Eip"`

	// ID of the CLB to bind, such as `lb-0000002i`. It can be queried in the console or obtained from `LoadBalancerId` returned by the `DescribeLoadBalancers` API.
	LoadBalancerID *string `json:"LoadBalancerID,omitnil" name:"LoadBalancerID"`

	// Region of the CLB instance, such as `ap-hongkong`.
	LoadBalancerRegion *string `json:"LoadBalancerRegion,omitnil" name:"LoadBalancerRegion"`

	// CLB private IP
	Vip *string `json:"Vip,omitnil" name:"Vip"`
}

type AssociateDDoSEipLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS instance ID (only Anti-DDoS Advanced). For example, `bgpip-0000011x`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// EIP of the Anti-DDoS instance ID.
	Eip *string `json:"Eip,omitnil" name:"Eip"`

	// ID of the CLB to bind, such as `lb-0000002i`. It can be queried in the console or obtained from `LoadBalancerId` returned by the `DescribeLoadBalancers` API.
	LoadBalancerID *string `json:"LoadBalancerID,omitnil" name:"LoadBalancerID"`

	// Region of the CLB instance, such as `ap-hongkong`.
	LoadBalancerRegion *string `json:"LoadBalancerRegion,omitnil" name:"LoadBalancerRegion"`

	// CLB private IP
	Vip *string `json:"Vip,omitnil" name:"Vip"`
}

func (r *AssociateDDoSEipLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateDDoSEipLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Eip")
	delete(f, "LoadBalancerID")
	delete(f, "LoadBalancerRegion")
	delete(f, "Vip")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateDDoSEipLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateDDoSEipLoadBalancerResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AssociateDDoSEipLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *AssociateDDoSEipLoadBalancerResponseParams `json:"Response"`
}

func (r *AssociateDDoSEipLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateDDoSEipLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BGPIPInstance struct {
	// Anti-DDoS instance details
	InstanceDetail *InstanceRelation `json:"InstanceDetail,omitnil" name:"InstanceDetail"`

	// Anti-DDoS instance specifications
	SpecificationLimit *BGPIPInstanceSpecification `json:"SpecificationLimit,omitnil" name:"SpecificationLimit"`

	// Anti-DDoS instance usage statistics
	Usage *BGPIPInstanceUsages `json:"Usage,omitnil" name:"Usage"`

	// Region of the Anti-DDoS instance
	Region *RegionInfo `json:"Region,omitnil" name:"Region"`

	// Status of the Anti-DDoS instance. Valid values:
	// `idle`: running
	// `attacking`: under attacks
	// `blocking`: blocked
	// `creating`: creating
	// `deblocking`: unblocking
	// `isolate`: reprocessed and isolated
	Status *string `json:"Status,omitnil" name:"Status"`

	// Purchase time
	ExpiredTime *string `json:"ExpiredTime,omitnil" name:"ExpiredTime"`

	// Expired At
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// Name of the Anti-DDoS instance
	Name *string `json:"Name,omitnil" name:"Name"`

	// Package details of the Anti-DDoS instance.
	// Note: This field is `null` for an Anti-DDoS instance without using a package.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	PackInfo *PackInfo `json:"PackInfo,omitnil" name:"PackInfo"`

	// Non-BGP package details of the Anti-DDoS instance.
	// Note: This field is `null` for an Anti-DDoS instance without using a non-BGP package.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	StaticPackRelation *StaticPackRelation `json:"StaticPackRelation,omitnil" name:"StaticPackRelation"`

	// Specifies the ISP. `0`: Chinese mainland ISPs (default); `1`：Radware；`2`: Tencent; `3`: NSFOCUS. Note that `1`, `2` and `3` are used for services outside the Chinese mainland.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ZoneId *uint64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// Used to differentiate clusters
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	Tgw *uint64 `json:"Tgw,omitnil" name:"Tgw"`

	// EIP states: `CREATING`, `BINDING`, `BIND`, `UNBINDING`, `UNBIND`, `OFFLINING`, and `BIND_ENI`. The EIP must be bound to an Anti-DDoS Advanced instance.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	EipAddressStatus *string `json:"EipAddressStatus,omitnil" name:"EipAddressStatus"`

	// Whether it is an Anti-DDoS EIP instance. `1`: Yes; `0`: No.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	EipFlag *int64 `json:"EipFlag,omitnil" name:"EipFlag"`

	// EIP package details of the Anti-DDoS Advanced instance.
	// Note: This field is `null` for an Anti-DDoS Advanced instance without using an EIP package.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	EipAddressPackRelation *EipAddressPackRelation `json:"EipAddressPackRelation,omitnil" name:"EipAddressPackRelation"`

	// Details of the Anti-DDoS Advanced instance bound to the EIP.
	// Note: This field is `null` if the EIP is not bound to an Anti-DDoS Advanced instance.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	EipAddressInfo *EipAddressRelation `json:"EipAddressInfo,omitnil" name:"EipAddressInfo"`

	// Recommended domain name for clients to access.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Whether to enable Sec-MCA. Valid values: `1` (enabled) and `0` (disabled).
	DamDDoSStatus *uint64 `json:"DamDDoSStatus,omitnil" name:"DamDDoSStatus"`

	// Whether it’s an IPv6 address. `1`: Yes; `0`: No.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	V6Flag *uint64 `json:"V6Flag,omitnil" name:"V6Flag"`

	// Whether it’s an Anti-DDoS Advanced instance from Tencent Cloud channels. `1`: Yes; `0`: No.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	BGPIPChannelFlag *uint64 `json:"BGPIPChannelFlag,omitnil" name:"BGPIPChannelFlag"`

	// Tag that the Anti-DDoS Advanced instance is associated with
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	TagInfoList []*TagInfo `json:"TagInfoList,omitnil" name:"TagInfoList"`

	// All-out package details of the instance
	// When an all-out package is not used by the instance, this field is `null`.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	AnycastOutPackRelation *AnycastOutPackRelation `json:"AnycastOutPackRelation,omitnil" name:"AnycastOutPackRelation"`

	// Edition of the instance
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	InstanceVersion *uint64 `json:"InstanceVersion,omitnil" name:"InstanceVersion"`

	// Convoy instance ID
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ConvoyId *string `json:"ConvoyId,omitnil" name:"ConvoyId"`

	// Pay-as-you-go bandwidth
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ElasticBandwidth *uint64 `json:"ElasticBandwidth,omitnil" name:"ElasticBandwidth"`

	// Whether it’s the IP broadcasted by EdgeOne. Values: `1` (yes), `0` (no)
	// Note: This field may return null, indicating that no valid values can be obtained.
	EOFlag *uint64 `json:"EOFlag,omitnil" name:"EOFlag"`
}

type BGPIPInstanceSpecification struct {
	// Base protection bandwidth (in Mbps)
	ProtectBandwidth *uint64 `json:"ProtectBandwidth,omitnil" name:"ProtectBandwidth"`

	// CC protection bandwidth (in QPS)
	ProtectCCQPS *uint64 `json:"ProtectCCQPS,omitnil" name:"ProtectCCQPS"`

	// Normal application bandwidth (in Mbps)
	NormalBandwidth *uint64 `json:"NormalBandwidth,omitnil" name:"NormalBandwidth"`

	// Number of forwarding rules
	ForwardRulesLimit *uint64 `json:"ForwardRulesLimit,omitnil" name:"ForwardRulesLimit"`

	// Auto-renewal status. Valid values:
	// `0`: disabled
	// `1`: enabled
	// ]
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// Anti-DDoS Advanced line. Valid values:
	// `1`: BGP
	// `2`: China Telecom
	// `3`: China Unicom
	// `4`: China Mobile
	// `99`: third-party line
	// ]
	Line *uint64 `json:"Line,omitnil" name:"Line"`

	// Elastic protection bandwidth (in Mbps)
	ElasticBandwidth *uint64 `json:"ElasticBandwidth,omitnil" name:"ElasticBandwidth"`
}

type BGPIPInstanceUsages struct {
	// Number of used port rules
	PortRulesUsage *uint64 `json:"PortRulesUsage,omitnil" name:"PortRulesUsage"`

	// Number of used domain name rules
	DomainRulesUsage *uint64 `json:"DomainRulesUsage,omitnil" name:"DomainRulesUsage"`

	// Number of attack times in the last 7 days
	Last7DayAttackCount *uint64 `json:"Last7DayAttackCount,omitnil" name:"Last7DayAttackCount"`
}

type BGPInstance struct {
	// Details of the Anti-DDoS Pro instance
	InstanceDetail *InstanceRelation `json:"InstanceDetail,omitnil" name:"InstanceDetail"`

	// Specifications of the Anti-DDoS Pro instance
	SpecificationLimit *BGPInstanceSpecification `json:"SpecificationLimit,omitnil" name:"SpecificationLimit"`

	// Usage statistics of the Anti-DDoS Pro instance
	Usage *BGPInstanceUsages `json:"Usage,omitnil" name:"Usage"`

	// Region of the Anti-DDoS Pro instance
	Region *RegionInfo `json:"Region,omitnil" name:"Region"`

	// Status of the Anti-DDoS Pro instance. Valid values:
	// `idle`: The instance is running normally.
	// `attacking`: The instance is under attack.
	// `blocking`: The instance is blocked.
	// `creating`: The instance is being created.
	// `deblocking`: Unblocking the instance
	// `isolate`: The instance is being isolated.
	Status *string `json:"Status,omitnil" name:"Status"`

	// Purchase time
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// Expiration time
	ExpiredTime *string `json:"ExpiredTime,omitnil" name:"ExpiredTime"`

	// Name of the Anti-DDoS Pro instance
	Name *string `json:"Name,omitnil" name:"Name"`

	// Details of the package to which the Anti-DDoS Pro instance belongs.
	// When the package provided is not used by the instance, this field is `null`.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	PackInfo *PackInfo `json:"PackInfo,omitnil" name:"PackInfo"`

	// Details of the cloud product used by the EIP bound to the Anti-DDoS Pro instance
	EipProductInfos []*EipProductInfo `json:"EipProductInfos,omitnil" name:"EipProductInfos"`

	// Binding status of the Anti-DDoS Pro instance
	// `idle`: The instance is bound.
	//  `bounding`: Binding the instance.
	// `failed`: Failed to bind
	// ]
	BoundStatus *string `json:"BoundStatus,omitnil" name:"BoundStatus"`

	// Layer-4 protection level
	DDoSLevel *string `json:"DDoSLevel,omitnil" name:"DDoSLevel"`

	// Status of CC protection
	CCEnable *int64 `json:"CCEnable,omitnil" name:"CCEnable"`

	// Tags associated with the resource
	TagInfoList []*TagInfo `json:"TagInfoList,omitnil" name:"TagInfoList"`

	// New edition of Anti-DDoS Pro
	IpCountNewFlag *uint64 `json:"IpCountNewFlag,omitnil" name:"IpCountNewFlag"`

	// The version of attack defense package
	VitalityVersion *uint64 `json:"VitalityVersion,omitnil" name:"VitalityVersion"`

	// Network line
	// Note: This field may return null, indicating that no valid values can be obtained.
	Line *uint64 `json:"Line,omitnil" name:"Line"`

	// Whether to enable elastic bandwidth
	ElasticServiceBandwidth *uint64 `json:"ElasticServiceBandwidth,omitnil" name:"ElasticServiceBandwidth"`

	// Bandwidth quota given away by Tencent Cloud
	GiftServiceBandWidth *int64 `json:"GiftServiceBandWidth,omitnil" name:"GiftServiceBandWidth"`
}

type BGPInstanceSpecification struct {
	// Base protection bandwidth (in Gbps)
	ProtectBandwidth *uint64 `json:"ProtectBandwidth,omitnil" name:"ProtectBandwidth"`

	// Number of protection chances
	ProtectCountLimit *uint64 `json:"ProtectCountLimit,omitnil" name:"ProtectCountLimit"`

	// Number of protected IPs
	ProtectIPNumberLimit *uint64 `json:"ProtectIPNumberLimit,omitnil" name:"ProtectIPNumberLimit"`

	// Auto-renewal status. Values:
	// `0`: Disabled
	// `1`: Enabled
	// ]
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// Protection type of Anti-DDoS Pro. Valid values: `0` (general protection) and `1` (Lighthouse-based protection).
	// Note: This field may return null, indicating that no valid values can be obtained.
	UnionPackFlag *uint64 `json:"UnionPackFlag,omitnil" name:"UnionPackFlag"`

	// Application bandwidth
	// Note: This field may return null, indicating that no valid values can be obtained.
	ServiceBandWidth *uint64 `json:"ServiceBandWidth,omitnil" name:"ServiceBandWidth"`

	// Whether it’s an Anti-DDoS Pro Premium edition. Values: `0` (General edition); `1` (Premium edition).
	// Note: This field may return null, indicating that no valid values can be obtained.
	BattleEditionFlag *uint64 `json:"BattleEditionFlag,omitnil" name:"BattleEditionFlag"`

	// Whether it’s an Anti-DDoS Pro Standard edition. Values: `0` (General edition); `1` (Standard edition).
	// Note: This field may return null, indicating that no valid values can be obtained.
	ChannelEditionFlag *uint64 `json:"ChannelEditionFlag,omitnil" name:"ChannelEditionFlag"`

	// Whether it’s an Anti-DDoS Pro Enterprise edition. Values: `0` (General edition); `1` (Enterprise edition).
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnterpriseFlag *uint64 `json:"EnterpriseFlag,omitnil" name:"EnterpriseFlag"`

	// Elastic bandwidth threshold of the Anti-DDoS Pro Enterprise edition.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ElasticLimit *uint64 `json:"ElasticLimit,omitnil" name:"ElasticLimit"`

	// Protection bandwidth after the plan downgrade, in Gbps. 
	// Note: This field may return `null`, indicating that no valid value was found.
	DownGradeProtect *uint64 `json:"DownGradeProtect,omitnil" name:"DownGradeProtect"`
}

type BGPInstanceUsages struct {
	// Number of used protection chances
	ProtectCountUsage *uint64 `json:"ProtectCountUsage,omitnil" name:"ProtectCountUsage"`

	// Number of protected IPs
	ProtectIPNumberUsage *uint64 `json:"ProtectIPNumberUsage,omitnil" name:"ProtectIPNumberUsage"`

	// Number of attack times in the last 7 days
	Last7DayAttackCount *uint64 `json:"Last7DayAttackCount,omitnil" name:"Last7DayAttackCount"`
}

type BlackWhiteIpRelation struct {
	// IP address
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// IP type. Valid values: `black` (blocklisted IP), `white`(allowlisted IP).
	Type *string `json:"Type,omitnil" name:"Type"`

	// Anti-DDoS instance configured
	InstanceDetailList []*InstanceRelation `json:"InstanceDetailList,omitnil" name:"InstanceDetailList"`

	// IP mask. `0` indicates a 32-bit IP.
	Mask *uint64 `json:"Mask,omitnil" name:"Mask"`

	// Modification time
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`
}

type BoundIpInfo struct {
	// IP address
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Category of product that can be bound. Valid values: `public` (CVM and CLB), `bm` (BM), `eni` (ENI), `vpngw` (VPN gateway), `natgw` (NAT gateway), `waf` (WAF), `fpc` (financial products), `gaap` (GAAP), and `other` (hosted IP). This field is required when you perform binding.
	BizType *string `json:"BizType,omitnil" name:"BizType"`

	// Anti-DDoS instance ID of the IP. This field is required only when the instance is bound to an IP. For example, this field InstanceId will be `eni-*` if the instance ID is bound to an ENI IP; `none` if there is no instance to bind to a managed IP.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Sub-product category. Valid values: `cvm` (CVM), `lb` (Load balancer), `eni` (ENI), `vpngw` (VPN gateway), `natgw` (NAT gateway), `waf` (WAF), `fpc` (financial products), `gaap` (GAAP), `eip` (BM EIP) and `other` (managed IP). This field is required when you perform binding.
	DeviceType *string `json:"DeviceType,omitnil" name:"DeviceType"`

	// ISP. Valid values: `0` (China Telecom), `1` (China Unicom), `2` (China Mobile), and `5` (BGP). This field is required when you perform binding.
	IspCode *uint64 `json:"IspCode,omitnil" name:"IspCode"`
}

type CCLevelPolicy struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Ip
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Protocol
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// Domain name
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Protection level. Values: `default`, `loose` and `strict`.
	Level *string `json:"Level,omitnil" name:"Level"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Modification time
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`
}

type CCPrecisionPlyRecord struct {
	// Type of the configuration field. Only `value` is supported.
	FieldType *string `json:"FieldType,omitnil" name:"FieldType"`

	// Configuration field. Valid values: `cgi`, `ua`, `cookie`, `referer`, `accept`, and `srcip`.
	FieldName *string `json:"FieldName,omitnil" name:"FieldName"`

	// Value of the configuration field
	Value *string `json:"Value,omitnil" name:"Value"`

	// Filters values of configuration fields. `equal`: The value matches the configuration field. `not_equal`: The value does not match the configuration field. `include`: The value is included.
	ValueOperator *string `json:"ValueOperator,omitnil" name:"ValueOperator"`
}

type CCPrecisionPolicy struct {
	// Policy ID
	PolicyId *string `json:"PolicyId,omitnil" name:"PolicyId"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// IP address
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Protocol
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// Domain name
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Action of limiting request frequency. Valid values: `alg` (limit request frequency via verification codes) and `drop`(drop requests).
	PolicyAction *string `json:"PolicyAction,omitnil" name:"PolicyAction"`

	// List of policies
	PolicyList []*CCPrecisionPlyRecord `json:"PolicyList,omitnil" name:"PolicyList"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Modification time
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`
}

type CCReqLimitPolicyRecord struct {
	// Sampling interval (in seconds). Valid values: `1`, `10`, `30`, and `60`.
	Period *uint64 `json:"Period,omitnil" name:"Period"`

	// Number of requests. Value range: 1-2000.
	RequestNum *uint64 `json:"RequestNum,omitnil" name:"RequestNum"`

	// Action of limiting request frequency. Valid values: `alg` (limit request frequency via verification codes) and `drop`(drop requests).
	Action *string `json:"Action,omitnil" name:"Action"`

	// Sets an interval of the frequency limit policy. Value range: 1-86400 (in seconds).
	ExecuteDuration *uint64 `json:"ExecuteDuration,omitnil" name:"ExecuteDuration"`

	// Filters values of configuration fields. `include`: The value is included. `equal`: The value matches the configuration field.
	Mode *string `json:"Mode,omitnil" name:"Mode"`

	// URI, which cannot be used together with `User-Agent` and `Cookie`.
	Uri *string `json:"Uri,omitnil" name:"Uri"`

	// User-Agent, which cannot be used together with `Uri` and `Cookie`.
	UserAgent *string `json:"UserAgent,omitnil" name:"UserAgent"`

	// Cookie, which cannot be used together with `Uri` and `User-Agent`.
	Cookie *string `json:"Cookie,omitnil" name:"Cookie"`
}

type CCThresholdPolicy struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// IP address
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Protocol
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// Domain name
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Cleansing threshold
	Threshold *int64 `json:"Threshold,omitnil" name:"Threshold"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Modification time
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`
}

type CcBlackWhiteIpPolicy struct {
	// Policy ID
	PolicyId *string `json:"PolicyId,omitnil" name:"PolicyId"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// IP address
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Domain name
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Protocol
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// IP type. Valid values: `black` (blocklisted IP), `white`(allowlisted IP).
	Type *string `json:"Type,omitnil" name:"Type"`

	// Blocklist/Allowlist IP address
	BlackWhiteIp *string `json:"BlackWhiteIp,omitnil" name:"BlackWhiteIp"`

	// Mask
	Mask *uint64 `json:"Mask,omitnil" name:"Mask"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Modification time
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`
}

type CcGeoIPBlockConfig struct {
	// Region type. Valid values:
	// `oversea`: Outside the Chinese mainland.
	// `china`: The Chinese mainland.
	// `customized`: User-specified region.
	// ]
	RegionType *string `json:"RegionType,omitnil" name:"RegionType"`

	// Blocking action. Valid values:
	// `drop`: Block the request.
	// `alg`: Verify the request via CAPTCHA.
	// ]
	Action *string `json:"Action,omitnil" name:"Action"`

	// Configuration ID, which is generated after a configuration is added. This field is only required to modify or delete a configuration.
	Id *string `json:"Id,omitnil" name:"Id"`

	// This field is required when RegionType is `customized`; it can be left empty when RegionType is `china` or `oversea`.
	AreaList []*int64 `json:"AreaList,omitnil" name:"AreaList"`
}

type CcGeoIpPolicyNew struct {
	// Policy ID
	PolicyId *string `json:"PolicyId,omitnil" name:"PolicyId"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// IP address
	Ip *string `json:"Ip,omitnil" name:"Ip"`


	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Protocol. Valid values: `HTTP` and `HTTPS`.
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// Action. Valid values: `drop` and `alg`.
	Action *string `json:"Action,omitnil" name:"Action"`

	// Region type. Valid values: `china`, `oversea` and `customized`.
	RegionType *string `json:"RegionType,omitnil" name:"RegionType"`

	// ID list of regions to be blocked
	AreaList []*uint64 `json:"AreaList,omitnil" name:"AreaList"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Modification time
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`
}

type CertIdInsL7Rules struct {
	// List of rules configured for certificates
	L7Rules []*InsL7Rules `json:"L7Rules,omitnil" name:"L7Rules"`

	// Certificate ID
	CertId *string `json:"CertId,omitnil" name:"CertId"`
}

// Predefined struct for user
type CreateBlackWhiteIpListRequestParams struct {
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// List of IPs
	IpList []*string `json:"IpList,omitnil" name:"IpList"`

	// IP type. Valid values: `black` (blocklisted IP), `white`(allowlisted IP).
	Type *string `json:"Type,omitnil" name:"Type"`
}

type CreateBlackWhiteIpListRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// List of IPs
	IpList []*string `json:"IpList,omitnil" name:"IpList"`

	// IP type. Valid values: `black` (blocklisted IP), `white`(allowlisted IP).
	Type *string `json:"Type,omitnil" name:"Type"`
}

func (r *CreateBlackWhiteIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBlackWhiteIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IpList")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBlackWhiteIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBlackWhiteIpListResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateBlackWhiteIpListResponse struct {
	*tchttp.BaseResponse
	Response *CreateBlackWhiteIpListResponseParams `json:"Response"`
}

func (r *CreateBlackWhiteIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBlackWhiteIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBoundIPRequestParams struct {
	// Anti-DDoS service type. `bgp`: Anti-DDoS Pro (Single IP); `bgp-multip`: Anti-DDoS Pro (Multi-IP)
	Business *string `json:"Business,omitnil" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// Array of IPs to bind to the Anti-DDoS instance. For Anti-DDoS Pro Single IP instance, the array contains only one IP. If there are no IPs to bind, it is empty; however, either `BoundDevList` or `UnBoundDevList` must not be empty.
	BoundDevList []*BoundIpInfo `json:"BoundDevList,omitnil" name:"BoundDevList"`

	// Array of IPs to unbind from the Anti-DDoS instance. For Anti-DDoS Pro Single IP instance, the array contains only one IP; if there are no IPs to unbind, it is empty; however, either `BoundDevList` or `UnBoundDevList` must not be empty.
	UnBoundDevList []*BoundIpInfo `json:"UnBoundDevList,omitnil" name:"UnBoundDevList"`

	// Disused
	CopyPolicy *string `json:"CopyPolicy,omitnil" name:"CopyPolicy"`
}

type CreateBoundIPRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS service type. `bgp`: Anti-DDoS Pro (Single IP); `bgp-multip`: Anti-DDoS Pro (Multi-IP)
	Business *string `json:"Business,omitnil" name:"Business"`

	// Anti-DDoS instance ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// Array of IPs to bind to the Anti-DDoS instance. For Anti-DDoS Pro Single IP instance, the array contains only one IP. If there are no IPs to bind, it is empty; however, either `BoundDevList` or `UnBoundDevList` must not be empty.
	BoundDevList []*BoundIpInfo `json:"BoundDevList,omitnil" name:"BoundDevList"`

	// Array of IPs to unbind from the Anti-DDoS instance. For Anti-DDoS Pro Single IP instance, the array contains only one IP; if there are no IPs to unbind, it is empty; however, either `BoundDevList` or `UnBoundDevList` must not be empty.
	UnBoundDevList []*BoundIpInfo `json:"UnBoundDevList,omitnil" name:"UnBoundDevList"`

	// Disused
	CopyPolicy *string `json:"CopyPolicy,omitnil" name:"CopyPolicy"`
}

func (r *CreateBoundIPRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBoundIPRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "BoundDevList")
	delete(f, "UnBoundDevList")
	delete(f, "CopyPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBoundIPRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBoundIPResponseParams struct {
	// Success code
	Success *SuccessCode `json:"Success,omitnil" name:"Success"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateBoundIPResponse struct {
	*tchttp.BaseResponse
	Response *CreateBoundIPResponseParams `json:"Response"`
}

func (r *CreateBoundIPResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBoundIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCCPrecisionPolicyRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// IP address
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Protocol. Valid values: `HTTP` and `HTTPS`.
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// Domain name
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Action of limiting request frequency. Valid values: `alg` (limit request frequency via verification codes) and `drop`(drop requests).
	PolicyAction *string `json:"PolicyAction,omitnil" name:"PolicyAction"`

	// Policy records
	PolicyList []*CCPrecisionPlyRecord `json:"PolicyList,omitnil" name:"PolicyList"`
}

type CreateCCPrecisionPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// IP address
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Protocol. Valid values: `HTTP` and `HTTPS`.
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// Domain name
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Action of limiting request frequency. Valid values: `alg` (limit request frequency via verification codes) and `drop`(drop requests).
	PolicyAction *string `json:"PolicyAction,omitnil" name:"PolicyAction"`

	// Policy records
	PolicyList []*CCPrecisionPlyRecord `json:"PolicyList,omitnil" name:"PolicyList"`
}

func (r *CreateCCPrecisionPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCCPrecisionPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Ip")
	delete(f, "Protocol")
	delete(f, "Domain")
	delete(f, "PolicyAction")
	delete(f, "PolicyList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCCPrecisionPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCCPrecisionPolicyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCCPrecisionPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateCCPrecisionPolicyResponseParams `json:"Response"`
}

func (r *CreateCCPrecisionPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCCPrecisionPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCCReqLimitPolicyRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// IP address
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Protocol. Valid values: `HTTP` and `HTTPS`.
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// Domain name
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Configuration field
	Policy *CCReqLimitPolicyRecord `json:"Policy,omitnil" name:"Policy"`

	// Whether it’s a global CC frequency limit
	IsGlobal *int64 `json:"IsGlobal,omitnil" name:"IsGlobal"`
}

type CreateCCReqLimitPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// IP address
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Protocol. Valid values: `HTTP` and `HTTPS`.
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// Domain name
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Configuration field
	Policy *CCReqLimitPolicyRecord `json:"Policy,omitnil" name:"Policy"`

	// Whether it’s a global CC frequency limit
	IsGlobal *int64 `json:"IsGlobal,omitnil" name:"IsGlobal"`
}

func (r *CreateCCReqLimitPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCCReqLimitPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Ip")
	delete(f, "Protocol")
	delete(f, "Domain")
	delete(f, "Policy")
	delete(f, "IsGlobal")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCCReqLimitPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCCReqLimitPolicyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCCReqLimitPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateCCReqLimitPolicyResponseParams `json:"Response"`
}

func (r *CreateCCReqLimitPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCCReqLimitPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCcBlackWhiteIpListRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// List of IPs
	IpList []*IpSegment `json:"IpList,omitnil" name:"IpList"`

	// IP permission. Valid values: `black` (blocked IP), `white` (allowed IP).
	Type *string `json:"Type,omitnil" name:"Type"`

	// IP address
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Domain name
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Protocol
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`
}

type CreateCcBlackWhiteIpListRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// List of IPs
	IpList []*IpSegment `json:"IpList,omitnil" name:"IpList"`

	// IP permission. Valid values: `black` (blocked IP), `white` (allowed IP).
	Type *string `json:"Type,omitnil" name:"Type"`

	// IP address
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Domain name
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Protocol
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`
}

func (r *CreateCcBlackWhiteIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCcBlackWhiteIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IpList")
	delete(f, "Type")
	delete(f, "Ip")
	delete(f, "Domain")
	delete(f, "Protocol")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCcBlackWhiteIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCcBlackWhiteIpListResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCcBlackWhiteIpListResponse struct {
	*tchttp.BaseResponse
	Response *CreateCcBlackWhiteIpListResponseParams `json:"Response"`
}

func (r *CreateCcBlackWhiteIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCcBlackWhiteIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCcGeoIPBlockConfigRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// IP address
	IP *string `json:"IP,omitnil" name:"IP"`

	// Domain name
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Protocol type
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// CC regional blocking configuration
	CcGeoIPBlockConfig *CcGeoIPBlockConfig `json:"CcGeoIPBlockConfig,omitnil" name:"CcGeoIPBlockConfig"`
}

type CreateCcGeoIPBlockConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// IP address
	IP *string `json:"IP,omitnil" name:"IP"`

	// Domain name
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Protocol type
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// CC regional blocking configuration
	CcGeoIPBlockConfig *CcGeoIPBlockConfig `json:"CcGeoIPBlockConfig,omitnil" name:"CcGeoIPBlockConfig"`
}

func (r *CreateCcGeoIPBlockConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCcGeoIPBlockConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IP")
	delete(f, "Domain")
	delete(f, "Protocol")
	delete(f, "CcGeoIPBlockConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCcGeoIPBlockConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCcGeoIPBlockConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCcGeoIPBlockConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateCcGeoIPBlockConfigResponseParams `json:"Response"`
}

func (r *CreateCcGeoIPBlockConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCcGeoIPBlockConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDDoSAIRequestParams struct {
	// List of Anti-DDoS instance IDs
	InstanceIdList []*string `json:"InstanceIdList,omitnil" name:"InstanceIdList"`

	// AI protection switch. Valid values:
	// `on`: enabled
	// `off`: disabled
	// ]
	DDoSAI *string `json:"DDoSAI,omitnil" name:"DDoSAI"`
}

type CreateDDoSAIRequest struct {
	*tchttp.BaseRequest
	
	// List of Anti-DDoS instance IDs
	InstanceIdList []*string `json:"InstanceIdList,omitnil" name:"InstanceIdList"`

	// AI protection switch. Valid values:
	// `on`: enabled
	// `off`: disabled
	// ]
	DDoSAI *string `json:"DDoSAI,omitnil" name:"DDoSAI"`
}

func (r *CreateDDoSAIRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDDoSAIRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIdList")
	delete(f, "DDoSAI")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDDoSAIRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDDoSAIResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDDoSAIResponse struct {
	*tchttp.BaseResponse
	Response *CreateDDoSAIResponseParams `json:"Response"`
}

func (r *CreateDDoSAIResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDDoSAIResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDDoSGeoIPBlockConfigRequestParams struct {
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Region blocking configuration. The configuration ID should be cleared when you set this parameter.
	DDoSGeoIPBlockConfig *DDoSGeoIPBlockConfig `json:"DDoSGeoIPBlockConfig,omitnil" name:"DDoSGeoIPBlockConfig"`
}

type CreateDDoSGeoIPBlockConfigRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Region blocking configuration. The configuration ID should be cleared when you set this parameter.
	DDoSGeoIPBlockConfig *DDoSGeoIPBlockConfig `json:"DDoSGeoIPBlockConfig,omitnil" name:"DDoSGeoIPBlockConfig"`
}

func (r *CreateDDoSGeoIPBlockConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDDoSGeoIPBlockConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DDoSGeoIPBlockConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDDoSGeoIPBlockConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDDoSGeoIPBlockConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDDoSGeoIPBlockConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateDDoSGeoIPBlockConfigResponseParams `json:"Response"`
}

func (r *CreateDDoSGeoIPBlockConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDDoSGeoIPBlockConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDDoSSpeedLimitConfigRequestParams struct {
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Access rate limit configuration. The configuration ID should be cleared when you set this parameter.
	DDoSSpeedLimitConfig *DDoSSpeedLimitConfig `json:"DDoSSpeedLimitConfig,omitnil" name:"DDoSSpeedLimitConfig"`
}

type CreateDDoSSpeedLimitConfigRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Access rate limit configuration. The configuration ID should be cleared when you set this parameter.
	DDoSSpeedLimitConfig *DDoSSpeedLimitConfig `json:"DDoSSpeedLimitConfig,omitnil" name:"DDoSSpeedLimitConfig"`
}

func (r *CreateDDoSSpeedLimitConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDDoSSpeedLimitConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DDoSSpeedLimitConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDDoSSpeedLimitConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDDoSSpeedLimitConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDDoSSpeedLimitConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateDDoSSpeedLimitConfigResponseParams `json:"Response"`
}

func (r *CreateDDoSSpeedLimitConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDDoSSpeedLimitConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDefaultAlarmThresholdRequestParams struct {
	// Default alarm threshold configuration
	DefaultAlarmConfig *DefaultAlarmThreshold `json:"DefaultAlarmConfig,omitnil" name:"DefaultAlarmConfig"`

	// Product category. Valid values:
	// `bgp`: Anti-DDoS Pro
	// `bgpip`: Anti-DDoS Advanced
	// ]
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`
}

type CreateDefaultAlarmThresholdRequest struct {
	*tchttp.BaseRequest
	
	// Default alarm threshold configuration
	DefaultAlarmConfig *DefaultAlarmThreshold `json:"DefaultAlarmConfig,omitnil" name:"DefaultAlarmConfig"`

	// Product category. Valid values:
	// `bgp`: Anti-DDoS Pro
	// `bgpip`: Anti-DDoS Advanced
	// ]
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`
}

func (r *CreateDefaultAlarmThresholdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDefaultAlarmThresholdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DefaultAlarmConfig")
	delete(f, "InstanceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDefaultAlarmThresholdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDefaultAlarmThresholdResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDefaultAlarmThresholdResponse struct {
	*tchttp.BaseResponse
	Response *CreateDefaultAlarmThresholdResponseParams `json:"Response"`
}

func (r *CreateDefaultAlarmThresholdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDefaultAlarmThresholdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIPAlarmThresholdConfigRequestParams struct {
	// List of IP alarm threshold configurations
	IpAlarmThresholdConfigList []*IPAlarmThresholdRelation `json:"IpAlarmThresholdConfigList,omitnil" name:"IpAlarmThresholdConfigList"`
}

type CreateIPAlarmThresholdConfigRequest struct {
	*tchttp.BaseRequest
	
	// List of IP alarm threshold configurations
	IpAlarmThresholdConfigList []*IPAlarmThresholdRelation `json:"IpAlarmThresholdConfigList,omitnil" name:"IpAlarmThresholdConfigList"`
}

func (r *CreateIPAlarmThresholdConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIPAlarmThresholdConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IpAlarmThresholdConfigList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIPAlarmThresholdConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIPAlarmThresholdConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateIPAlarmThresholdConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateIPAlarmThresholdConfigResponseParams `json:"Response"`
}

func (r *CreateIPAlarmThresholdConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIPAlarmThresholdConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateL7RuleCertsRequestParams struct {
	// SSL certificate ID
	CertId *string `json:"CertId,omitnil" name:"CertId"`

	// List of Layer-7 domain name forwarding rules
	L7Rules []*InsL7Rules `json:"L7Rules,omitnil" name:"L7Rules"`
}

type CreateL7RuleCertsRequest struct {
	*tchttp.BaseRequest
	
	// SSL certificate ID
	CertId *string `json:"CertId,omitnil" name:"CertId"`

	// List of Layer-7 domain name forwarding rules
	L7Rules []*InsL7Rules `json:"L7Rules,omitnil" name:"L7Rules"`
}

func (r *CreateL7RuleCertsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7RuleCertsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertId")
	delete(f, "L7Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateL7RuleCertsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateL7RuleCertsResponseParams struct {
	// Success code
	Success *SuccessCode `json:"Success,omitnil" name:"Success"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateL7RuleCertsResponse struct {
	*tchttp.BaseResponse
	Response *CreateL7RuleCertsResponseParams `json:"Response"`
}

func (r *CreateL7RuleCertsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateL7RuleCertsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNewL7RulesRequestParams struct {
	// List of rules
	Rules []*L7RuleEntry `json:"Rules,omitnil" name:"Rules"`

	// Anti-DDoS service type (`bgpip`: Anti-DDoS Advanced)
	Business *string `json:"Business,omitnil" name:"Business"`

	// List of resource IDs
	IdList []*string `json:"IdList,omitnil" name:"IdList"`

	// List of resource IPs
	VipList []*string `json:"VipList,omitnil" name:"VipList"`
}

type CreateNewL7RulesRequest struct {
	*tchttp.BaseRequest
	
	// List of rules
	Rules []*L7RuleEntry `json:"Rules,omitnil" name:"Rules"`

	// Anti-DDoS service type (`bgpip`: Anti-DDoS Advanced)
	Business *string `json:"Business,omitnil" name:"Business"`

	// List of resource IDs
	IdList []*string `json:"IdList,omitnil" name:"IdList"`

	// List of resource IPs
	VipList []*string `json:"VipList,omitnil" name:"VipList"`
}

func (r *CreateNewL7RulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNewL7RulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Rules")
	delete(f, "Business")
	delete(f, "IdList")
	delete(f, "VipList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNewL7RulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNewL7RulesResponseParams struct {
	// Success code
	Success *SuccessCode `json:"Success,omitnil" name:"Success"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateNewL7RulesResponse struct {
	*tchttp.BaseResponse
	Response *CreateNewL7RulesResponseParams `json:"Response"`
}

func (r *CreateNewL7RulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNewL7RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePacketFilterConfigRequestParams struct {
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Feature filtering rules
	PacketFilterConfig *PacketFilterConfig `json:"PacketFilterConfig,omitnil" name:"PacketFilterConfig"`
}

type CreatePacketFilterConfigRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Feature filtering rules
	PacketFilterConfig *PacketFilterConfig `json:"PacketFilterConfig,omitnil" name:"PacketFilterConfig"`
}

func (r *CreatePacketFilterConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePacketFilterConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PacketFilterConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePacketFilterConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePacketFilterConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreatePacketFilterConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreatePacketFilterConfigResponseParams `json:"Response"`
}

func (r *CreatePacketFilterConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePacketFilterConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProtocolBlockConfigRequestParams struct {
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Protocol blocking configuration
	ProtocolBlockConfig *ProtocolBlockConfig `json:"ProtocolBlockConfig,omitnil" name:"ProtocolBlockConfig"`
}

type CreateProtocolBlockConfigRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Protocol blocking configuration
	ProtocolBlockConfig *ProtocolBlockConfig `json:"ProtocolBlockConfig,omitnil" name:"ProtocolBlockConfig"`
}

func (r *CreateProtocolBlockConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProtocolBlockConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ProtocolBlockConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProtocolBlockConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProtocolBlockConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateProtocolBlockConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateProtocolBlockConfigResponseParams `json:"Response"`
}

func (r *CreateProtocolBlockConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProtocolBlockConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSchedulingDomainRequestParams struct {
	// Indicates whether a hybrid cloud product is used.
	// `hybrid`: Anti-DDoS Service Platform
	// For other products, leave this field empty.
	Product *string `json:"Product,omitnil" name:"Product"`
}

type CreateSchedulingDomainRequest struct {
	*tchttp.BaseRequest
	
	// Indicates whether a hybrid cloud product is used.
	// `hybrid`: Anti-DDoS Service Platform
	// For other products, leave this field empty.
	Product *string `json:"Product,omitnil" name:"Product"`
}

func (r *CreateSchedulingDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSchedulingDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSchedulingDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSchedulingDomainResponseParams struct {
	// Created domain name
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateSchedulingDomainResponse struct {
	*tchttp.BaseResponse
	Response *CreateSchedulingDomainResponseParams `json:"Response"`
}

func (r *CreateSchedulingDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSchedulingDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWaterPrintConfigRequestParams struct {
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Watermark configuration
	WaterPrintConfig *WaterPrintConfig `json:"WaterPrintConfig,omitnil" name:"WaterPrintConfig"`
}

type CreateWaterPrintConfigRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Watermark configuration
	WaterPrintConfig *WaterPrintConfig `json:"WaterPrintConfig,omitnil" name:"WaterPrintConfig"`
}

func (r *CreateWaterPrintConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWaterPrintConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "WaterPrintConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWaterPrintConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWaterPrintConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateWaterPrintConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateWaterPrintConfigResponseParams `json:"Response"`
}

func (r *CreateWaterPrintConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWaterPrintConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWaterPrintKeyRequestParams struct {
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type CreateWaterPrintKeyRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *CreateWaterPrintKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWaterPrintKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWaterPrintKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWaterPrintKeyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateWaterPrintKeyResponse struct {
	*tchttp.BaseResponse
	Response *CreateWaterPrintKeyResponseParams `json:"Response"`
}

func (r *CreateWaterPrintKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWaterPrintKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DDoSAIRelation struct {
	// AI protection switch. Valid values:
	// `on`: enabled
	// `off`: disabled
	// ]
	DDoSAI *string `json:"DDoSAI,omitnil" name:"DDoSAI"`

	// Anti-DDoS instance configured
	InstanceDetailList []*InstanceRelation `json:"InstanceDetailList,omitnil" name:"InstanceDetailList"`
}

type DDoSGeoIPBlockConfig struct {
	// Region type. Valid values:
	// `oversea`: Outside the Chinese mainland
	// `china`: The Chinese mainland
	// `customized`: User-specified region
	// ]
	RegionType *string `json:"RegionType,omitnil" name:"RegionType"`

	// Blocking action. Valid values:
	// `drop`: the request is blocked.
	// `trans`: the request is allowed.
	// ]
	Action *string `json:"Action,omitnil" name:"Action"`

	// Configuration ID, which is generated after a configuration is added. This field is only required to modify or delete a configuration.
	Id *string `json:"Id,omitnil" name:"Id"`

	// When `RegionType = customized`, AreaList is required and contains up to 128 areas.
	AreaList []*int64 `json:"AreaList,omitnil" name:"AreaList"`
}

type DDoSGeoIPBlockConfigRelation struct {
	// Anti-DDoS region blocking configuration
	GeoIPBlockConfig *DDoSGeoIPBlockConfig `json:"GeoIPBlockConfig,omitnil" name:"GeoIPBlockConfig"`

	// Anti-DDoS instance configured
	InstanceDetailList []*InstanceRelation `json:"InstanceDetailList,omitnil" name:"InstanceDetailList"`
}

type DDoSSpeedLimitConfig struct {
	// Rate limit mode. Valid values:
	// `1`: rate limit based on the real server IP
	// `2`: rate limit based on the destination port
	// ]
	Mode *uint64 `json:"Mode,omitnil" name:"Mode"`

	// Rate limit value. This field contains at least one valid rate limit type. Note that only up to one value of each type is supported.
	SpeedValues []*SpeedValue `json:"SpeedValues,omitnil" name:"SpeedValues"`

	// This field is replaced with a new field DstPortList.
	DstPortScopes []*PortSegment `json:"DstPortScopes,omitnil" name:"DstPortScopes"`


	Id *string `json:"Id,omitnil" name:"Id"`

	// IP protocol number. Valid values:
	// `ALL`: all protocols
	// `TCP`: TCP protocol
	// `UDP`: UDP protocol
	// `SMP`: SMP protocol
	// `1;2–100`: user-defined protocol with up to 8 ranges
	// ]
	// Note: For custom protocol ranges, only protocol number is supported. Multiple ranges are separated by ";". If the value is `ALL`, any other protocol or protocol number should be excluded.
	ProtocolList *string `json:"ProtocolList,omitnil" name:"ProtocolList"`

	// Port range list, which contains up to 8 ranges. Use ";" to separate multiple ports and "–" to indicate a range of ports, as described in the following formats: `0–65535`, `80;443;1000–2000`.
	DstPortList *string `json:"DstPortList,omitnil" name:"DstPortList"`
}

type DDoSSpeedLimitConfigRelation struct {
	// Anti-DDoS access rate limit configuration
	SpeedLimitConfig *DDoSSpeedLimitConfig `json:"SpeedLimitConfig,omitnil" name:"SpeedLimitConfig"`

	// Anti-DDoS instance configured
	InstanceDetailList []*InstanceRelation `json:"InstanceDetailList,omitnil" name:"InstanceDetailList"`
}

type DefaultAlarmThreshold struct {
	// Alarm threshold type. Valid values:
	// `1`: alarm threshold for inbound traffic
	// `2`: alarm threshold for cleansing attack traffic
	// ]
	AlarmType *uint64 `json:"AlarmType,omitnil" name:"AlarmType"`

	// Alarm threshold (Mbps). The value should be greater than or equal to 0. Note that the alarm threshold configuration will be removed if you pass the parameter for input and set it to 0.
	AlarmThreshold *uint64 `json:"AlarmThreshold,omitnil" name:"AlarmThreshold"`
}

// Predefined struct for user
type DeleteCCLevelPolicyRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Target IP of the policy
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Domain name
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Value: `http`
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`
}

type DeleteCCLevelPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Target IP of the policy
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Domain name
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Value: `http`
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`
}

func (r *DeleteCCLevelPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCCLevelPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Ip")
	delete(f, "Domain")
	delete(f, "Protocol")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCCLevelPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCCLevelPolicyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteCCLevelPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCCLevelPolicyResponseParams `json:"Response"`
}

func (r *DeleteCCLevelPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCCLevelPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCCPrecisionPolicyRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Policy ID
	PolicyId *string `json:"PolicyId,omitnil" name:"PolicyId"`
}

type DeleteCCPrecisionPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Policy ID
	PolicyId *string `json:"PolicyId,omitnil" name:"PolicyId"`
}

func (r *DeleteCCPrecisionPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCCPrecisionPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCCPrecisionPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCCPrecisionPolicyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteCCPrecisionPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCCPrecisionPolicyResponseParams `json:"Response"`
}

func (r *DeleteCCPrecisionPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCCPrecisionPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCCThresholdPolicyRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Target IP of the policy
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Domain name
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Value: `http`
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`
}

type DeleteCCThresholdPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Target IP of the policy
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Domain name
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Value: `http`
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`
}

func (r *DeleteCCThresholdPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCCThresholdPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Ip")
	delete(f, "Domain")
	delete(f, "Protocol")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCCThresholdPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCCThresholdPolicyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteCCThresholdPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCCThresholdPolicyResponseParams `json:"Response"`
}

func (r *DeleteCCThresholdPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCCThresholdPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCcBlackWhiteIpListRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Policy ID
	PolicyId *string `json:"PolicyId,omitnil" name:"PolicyId"`
}

type DeleteCcBlackWhiteIpListRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Policy ID
	PolicyId *string `json:"PolicyId,omitnil" name:"PolicyId"`
}

func (r *DeleteCcBlackWhiteIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCcBlackWhiteIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCcBlackWhiteIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCcBlackWhiteIpListResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteCcBlackWhiteIpListResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCcBlackWhiteIpListResponseParams `json:"Response"`
}

func (r *DeleteCcBlackWhiteIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCcBlackWhiteIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCcGeoIPBlockConfigRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Region blocking configuration. The configuration ID cannot be empty when you set this parameter.
	CcGeoIPBlockConfig *CcGeoIPBlockConfig `json:"CcGeoIPBlockConfig,omitnil" name:"CcGeoIPBlockConfig"`
}

type DeleteCcGeoIPBlockConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Region blocking configuration. The configuration ID cannot be empty when you set this parameter.
	CcGeoIPBlockConfig *CcGeoIPBlockConfig `json:"CcGeoIPBlockConfig,omitnil" name:"CcGeoIPBlockConfig"`
}

func (r *DeleteCcGeoIPBlockConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCcGeoIPBlockConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "CcGeoIPBlockConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCcGeoIPBlockConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCcGeoIPBlockConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteCcGeoIPBlockConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCcGeoIPBlockConfigResponseParams `json:"Response"`
}

func (r *DeleteCcGeoIPBlockConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCcGeoIPBlockConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDDoSGeoIPBlockConfigRequestParams struct {
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Region blocking configuration. The configuration ID cannot be empty when you set this parameter.
	DDoSGeoIPBlockConfig *DDoSGeoIPBlockConfig `json:"DDoSGeoIPBlockConfig,omitnil" name:"DDoSGeoIPBlockConfig"`
}

type DeleteDDoSGeoIPBlockConfigRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Region blocking configuration. The configuration ID cannot be empty when you set this parameter.
	DDoSGeoIPBlockConfig *DDoSGeoIPBlockConfig `json:"DDoSGeoIPBlockConfig,omitnil" name:"DDoSGeoIPBlockConfig"`
}

func (r *DeleteDDoSGeoIPBlockConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDDoSGeoIPBlockConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DDoSGeoIPBlockConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDDoSGeoIPBlockConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDDoSGeoIPBlockConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteDDoSGeoIPBlockConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDDoSGeoIPBlockConfigResponseParams `json:"Response"`
}

func (r *DeleteDDoSGeoIPBlockConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDDoSGeoIPBlockConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDDoSSpeedLimitConfigRequestParams struct {
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Access rate limit configuration. The configuration ID cannot be empty when you set this parameter.
	DDoSSpeedLimitConfig *DDoSSpeedLimitConfig `json:"DDoSSpeedLimitConfig,omitnil" name:"DDoSSpeedLimitConfig"`
}

type DeleteDDoSSpeedLimitConfigRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Access rate limit configuration. The configuration ID cannot be empty when you set this parameter.
	DDoSSpeedLimitConfig *DDoSSpeedLimitConfig `json:"DDoSSpeedLimitConfig,omitnil" name:"DDoSSpeedLimitConfig"`
}

func (r *DeleteDDoSSpeedLimitConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDDoSSpeedLimitConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DDoSSpeedLimitConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDDoSSpeedLimitConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDDoSSpeedLimitConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteDDoSSpeedLimitConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDDoSSpeedLimitConfigResponseParams `json:"Response"`
}

func (r *DeleteDDoSSpeedLimitConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDDoSSpeedLimitConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePacketFilterConfigRequestParams struct {
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Feature filtering configuration
	PacketFilterConfig *PacketFilterConfig `json:"PacketFilterConfig,omitnil" name:"PacketFilterConfig"`
}

type DeletePacketFilterConfigRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Feature filtering configuration
	PacketFilterConfig *PacketFilterConfig `json:"PacketFilterConfig,omitnil" name:"PacketFilterConfig"`
}

func (r *DeletePacketFilterConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePacketFilterConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PacketFilterConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePacketFilterConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePacketFilterConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeletePacketFilterConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeletePacketFilterConfigResponseParams `json:"Response"`
}

func (r *DeletePacketFilterConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePacketFilterConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWaterPrintConfigRequestParams struct {
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DeleteWaterPrintConfigRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DeleteWaterPrintConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWaterPrintConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWaterPrintConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWaterPrintConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteWaterPrintConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWaterPrintConfigResponseParams `json:"Response"`
}

func (r *DeleteWaterPrintConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWaterPrintConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWaterPrintKeyRequestParams struct {
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Watermark key ID
	KeyId *string `json:"KeyId,omitnil" name:"KeyId"`
}

type DeleteWaterPrintKeyRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Watermark key ID
	KeyId *string `json:"KeyId,omitnil" name:"KeyId"`
}

func (r *DeleteWaterPrintKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWaterPrintKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWaterPrintKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWaterPrintKeyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteWaterPrintKeyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWaterPrintKeyResponseParams `json:"Response"`
}

func (r *DeleteWaterPrintKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWaterPrintKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBasicDeviceStatusRequestParams struct {
	// List of IP resources
	IpList []*string `json:"IpList,omitnil" name:"IpList"`
}

type DescribeBasicDeviceStatusRequest struct {
	*tchttp.BaseRequest
	
	// List of IP resources
	IpList []*string `json:"IpList,omitnil" name:"IpList"`
}

func (r *DescribeBasicDeviceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBasicDeviceStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IpList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBasicDeviceStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBasicDeviceStatusResponseParams struct {
	// Status of the specified Anti-DDoS resource. Valid values:
	// `1`: The IP is blocked.
	// `2`: The P is normal.
	// `3`: The IP is being attacked.
	Data []*KeyValue `json:"Data,omitnil" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBasicDeviceStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBasicDeviceStatusResponseParams `json:"Response"`
}

func (r *DescribeBasicDeviceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBasicDeviceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBgpBizTrendRequestParams struct {
	// Anti-DDoS service code. `bgp-multip` indicates Anti-DDos Pro.
	Business *string `json:"Business,omitnil" name:"Business"`

	// Beginning of the time range for the query, such as `2020-09-22 00:00:00`.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End of the time range for the query, such as `2020-09-22 00:00:00`.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Statistical metric. Values: `intraffic`, `outtraffic`, `inpkg`, and `outpkg`.
	MetricName *string `json:"MetricName,omitnil" name:"MetricName"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// `0`: Fixed time. `1`: Custom time.
	Flag *uint64 `json:"Flag,omitnil" name:"Flag"`
}

type DescribeBgpBizTrendRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS service code. `bgp-multip` indicates Anti-DDos Pro.
	Business *string `json:"Business,omitnil" name:"Business"`

	// Beginning of the time range for the query, such as `2020-09-22 00:00:00`.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End of the time range for the query, such as `2020-09-22 00:00:00`.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Statistical metric. Values: `intraffic`, `outtraffic`, `inpkg`, and `outpkg`.
	MetricName *string `json:"MetricName,omitnil" name:"MetricName"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// `0`: Fixed time. `1`: Custom time.
	Flag *uint64 `json:"Flag,omitnil" name:"Flag"`
}

func (r *DescribeBgpBizTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBgpBizTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricName")
	delete(f, "InstanceId")
	delete(f, "Flag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBgpBizTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBgpBizTrendResponseParams struct {
	// Values of the samples
	DataList []*uint64 `json:"DataList,omitnil" name:"DataList"`

	// Number of samples
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// Statistical metric
	MetricName *string `json:"MetricName,omitnil" name:"MetricName"`

	// Maximum value of the arrays returned
	MaxData *uint64 `json:"MaxData,omitnil" name:"MaxData"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBgpBizTrendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBgpBizTrendResponseParams `json:"Response"`
}

func (r *DescribeBgpBizTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBgpBizTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBizHttpStatusRequestParams struct {
	// Statistical mode. Value: `sum`.
	Statistics *string `json:"Statistics,omitnil" name:"Statistics"`

	// Anti-DDoS service type (`bgpip`: Anti-DDoS Advanced)
	Business *string `json:"Business,omitnil" name:"Business"`

	// Statistical period in seconds. Valid values: `60`, `300`, `1800`, `3600`, `21600`, and `86400`.
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// Statistics start time, such as `2020-02-01 12:04:12`
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// Statistics end time, such as `2020-02-03 18:03:23`
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// The resource ID.
	Id *string `json:"Id,omitnil" name:"Id"`

	// Specific domain name query
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Protocol and port list, which is valid when the metric is `connum`, `new_conn` or `inactive_conn`. Valid protocols: `TCP`, `UDP`, `HTTP`, `HTTPS`
	ProtoInfo []*ProtocolPort `json:"ProtoInfo,omitnil" name:"ProtoInfo"`
}

type DescribeBizHttpStatusRequest struct {
	*tchttp.BaseRequest
	
	// Statistical mode. Value: `sum`.
	Statistics *string `json:"Statistics,omitnil" name:"Statistics"`

	// Anti-DDoS service type (`bgpip`: Anti-DDoS Advanced)
	Business *string `json:"Business,omitnil" name:"Business"`

	// Statistical period in seconds. Valid values: `60`, `300`, `1800`, `3600`, `21600`, and `86400`.
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// Statistics start time, such as `2020-02-01 12:04:12`
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// Statistics end time, such as `2020-02-03 18:03:23`
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// The resource ID.
	Id *string `json:"Id,omitnil" name:"Id"`

	// Specific domain name query
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Protocol and port list, which is valid when the metric is `connum`, `new_conn` or `inactive_conn`. Valid protocols: `TCP`, `UDP`, `HTTP`, `HTTPS`
	ProtoInfo []*ProtocolPort `json:"ProtoInfo,omitnil" name:"ProtoInfo"`
}

func (r *DescribeBizHttpStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBizHttpStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Statistics")
	delete(f, "Business")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Id")
	delete(f, "Domain")
	delete(f, "ProtoInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBizHttpStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBizHttpStatusResponseParams struct {
	// Statistics on the HTTP status codes of business traffic
	HttpStatusMap *HttpStatusMap `json:"HttpStatusMap,omitnil" name:"HttpStatusMap"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBizHttpStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBizHttpStatusResponseParams `json:"Response"`
}

func (r *DescribeBizHttpStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBizHttpStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBizTrendRequestParams struct {
	// Statistical method. Valid values: `max`, `min`, `avg`, `sum`. It can only be `max` if the statistical dimension is traffic rate or packet rate.
	Statistics *string `json:"Statistics,omitnil" name:"Statistics"`

	// Anti-DDoS service type (`bgpip`: Anti-DDoS Advanced)
	Business *string `json:"Business,omitnil" name:"Business"`

	// Sampling interval in seconds. Valid values: `60`, `300`, `1800`, `3600`, `21600`, `86400`
	Period *uint64 `json:"Period,omitnil" name:"Period"`

	// Beginning of the time range for the query, such as `2020-09-22 00:00:00`.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End of the time range for the query, such as `2020-09-22 00:00:00`.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Instance ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// Metric. Valid values: `connum`, `new_conn`, `inactive_conn`, `intraffic`, `outtraffic`, `inpkg`, `outpkg`, `qps`
	MetricName *string `json:"MetricName,omitnil" name:"MetricName"`

	// You can query data by specifying a domain name when the metric is `qps`.
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Protocol and port list, which is valid when the metric is `connum`, `new_conn` or `inactive_conn`. Valid protocols: `TCP`, `UDP`, `HTTP`, `HTTPS`
	ProtoInfo []*ProtocolPort `json:"ProtoInfo,omitnil" name:"ProtoInfo"`
}

type DescribeBizTrendRequest struct {
	*tchttp.BaseRequest
	
	// Statistical method. Valid values: `max`, `min`, `avg`, `sum`. It can only be `max` if the statistical dimension is traffic rate or packet rate.
	Statistics *string `json:"Statistics,omitnil" name:"Statistics"`

	// Anti-DDoS service type (`bgpip`: Anti-DDoS Advanced)
	Business *string `json:"Business,omitnil" name:"Business"`

	// Sampling interval in seconds. Valid values: `60`, `300`, `1800`, `3600`, `21600`, `86400`
	Period *uint64 `json:"Period,omitnil" name:"Period"`

	// Beginning of the time range for the query, such as `2020-09-22 00:00:00`.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End of the time range for the query, such as `2020-09-22 00:00:00`.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Instance ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// Metric. Valid values: `connum`, `new_conn`, `inactive_conn`, `intraffic`, `outtraffic`, `inpkg`, `outpkg`, `qps`
	MetricName *string `json:"MetricName,omitnil" name:"MetricName"`

	// You can query data by specifying a domain name when the metric is `qps`.
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Protocol and port list, which is valid when the metric is `connum`, `new_conn` or `inactive_conn`. Valid protocols: `TCP`, `UDP`, `HTTP`, `HTTPS`
	ProtoInfo []*ProtocolPort `json:"ProtoInfo,omitnil" name:"ProtoInfo"`
}

func (r *DescribeBizTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBizTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Statistics")
	delete(f, "Business")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Id")
	delete(f, "MetricName")
	delete(f, "Domain")
	delete(f, "ProtoInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBizTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBizTrendResponseParams struct {
	// Value at a time point on the curve
	DataList []*float64 `json:"DataList,omitnil" name:"DataList"`

	// Statistical dimension
	MetricName *string `json:"MetricName,omitnil" name:"MetricName"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBizTrendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBizTrendResponseParams `json:"Response"`
}

func (r *DescribeBizTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBizTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBlackWhiteIpListRequestParams struct {
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeBlackWhiteIpListRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeBlackWhiteIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlackWhiteIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBlackWhiteIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBlackWhiteIpListResponseParams struct {
	// IP blocklist
	BlackIpList []*string `json:"BlackIpList,omitnil" name:"BlackIpList"`

	// IP allowlist
	WhiteIpList []*string `json:"WhiteIpList,omitnil" name:"WhiteIpList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBlackWhiteIpListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBlackWhiteIpListResponseParams `json:"Response"`
}

func (r *DescribeBlackWhiteIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlackWhiteIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCLevelListRequestParams struct {
	// Anti-DDoS service code. `bgp-multip` indicates Anti-DDos Pro.
	Business *string `json:"Business,omitnil" name:"Business"`

	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of results returned in one page
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// ID of the specified instance
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeCCLevelListRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS service code. `bgp-multip` indicates Anti-DDos Pro.
	Business *string `json:"Business,omitnil" name:"Business"`

	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of results returned in one page
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// ID of the specified instance
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeCCLevelListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCLevelListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCCLevelListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCLevelListResponseParams struct {
	// Total number of level-defining policies
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// Total number of level-defining policies
	LevelList []*CCLevelPolicy `json:"LevelList,omitnil" name:"LevelList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCCLevelListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCCLevelListResponseParams `json:"Response"`
}

func (r *DescribeCCLevelListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCLevelListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCLevelPolicyRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// IP
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Domain name
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Protocol. Values: `HTTP`，`HTTPS`
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`
}

type DescribeCCLevelPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// IP
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Domain name
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Protocol. Values: `HTTP`，`HTTPS`
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`
}

func (r *DescribeCCLevelPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCLevelPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Ip")
	delete(f, "Domain")
	delete(f, "Protocol")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCCLevelPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCLevelPolicyResponseParams struct {
	// CC protection level. Vaules: `loose`, `strict`, `normal`, `emergency`, `sup_loose` (super loose), `default` (used when the frequency limit is not configured) and `customized`
	Level *string `json:"Level,omitnil" name:"Level"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCCLevelPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCCLevelPolicyResponseParams `json:"Response"`
}

func (r *DescribeCCLevelPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCLevelPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCPrecisionPlyListRequestParams struct {
	// Anti-DDoS service type. Valid values: `bgpip-multip` (Anti-DDoS Pro) and `bgpip` (Anti-DDoS Advanced).
	Business *string `json:"Business,omitnil" name:"Business"`

	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of results returned in one page
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// IP address, which is required when an Anti-DDoS Advanced instance is used.
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Domain name, which is required when an Anti-DDoS Advanced instance is used.
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Protocol, which is required when an Anti-DDoS Advanced instance is used.
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`
}

type DescribeCCPrecisionPlyListRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS service type. Valid values: `bgpip-multip` (Anti-DDoS Pro) and `bgpip` (Anti-DDoS Advanced).
	Business *string `json:"Business,omitnil" name:"Business"`

	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of results returned in one page
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// IP address, which is required when an Anti-DDoS Advanced instance is used.
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Domain name, which is required when an Anti-DDoS Advanced instance is used.
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Protocol, which is required when an Anti-DDoS Advanced instance is used.
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`
}

func (r *DescribeCCPrecisionPlyListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCPrecisionPlyListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstanceId")
	delete(f, "Ip")
	delete(f, "Domain")
	delete(f, "Protocol")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCCPrecisionPlyListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCPrecisionPlyListResponseParams struct {
	// Number of policy lists
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// Information of the policy list
	PrecisionPolicyList []*CCPrecisionPolicy `json:"PrecisionPolicyList,omitnil" name:"PrecisionPolicyList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCCPrecisionPlyListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCCPrecisionPlyListResponseParams `json:"Response"`
}

func (r *DescribeCCPrecisionPlyListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCPrecisionPlyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCThresholdListRequestParams struct {
	// Anti-DDoS service code. `bgp-multip` indicates Anti-DDos Pro.
	Business *string `json:"Business,omitnil" name:"Business"`

	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of results returned in one page
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// ID of the specified instance
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeCCThresholdListRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS service code. `bgp-multip` indicates Anti-DDos Pro.
	Business *string `json:"Business,omitnil" name:"Business"`

	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of results returned in one page
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// ID of the specified instance
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeCCThresholdListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCThresholdListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCCThresholdListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCThresholdListResponseParams struct {
	// Total number of cleansing threshold policies
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// Details of cleansing threshold policies
	ThresholdList []*CCThresholdPolicy `json:"ThresholdList,omitnil" name:"ThresholdList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCCThresholdListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCCThresholdListResponseParams `json:"Response"`
}

func (r *DescribeCCThresholdListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCThresholdListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCTrendRequestParams struct {
	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (Single IP); `bgp-multip`: Anti-DDoS Pro (Multi-IP); `net`: Anti-DDoS Ultimate; `basic`: Anti-DDoS Basic
	Business *string `json:"Business,omitnil" name:"Business"`

	// Instance IP
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Sampling period. Valid values: `300` (5 minutes), `3600` (one hour), `86400` (one day)
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// Beginning of the time range for the query
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End of the time range for the query
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Metric. Valid values: `inqps` (total QPS peaks), `dropqps` (attack QPS peaks), `incount` (total number of requests), and `dropcount` (number of attack requests).
	MetricName *string `json:"MetricName,omitnil" name:"MetricName"`

	// (Optional) Domain name
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Instance ID. Leave this field empty when `Business` is `basic`, as basic protection does not require an instance.
	Id *string `json:"Id,omitnil" name:"Id"`
}

type DescribeCCTrendRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (Single IP); `bgp-multip`: Anti-DDoS Pro (Multi-IP); `net`: Anti-DDoS Ultimate; `basic`: Anti-DDoS Basic
	Business *string `json:"Business,omitnil" name:"Business"`

	// Instance IP
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Sampling period. Valid values: `300` (5 minutes), `3600` (one hour), `86400` (one day)
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// Beginning of the time range for the query
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End of the time range for the query
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Metric. Valid values: `inqps` (total QPS peaks), `dropqps` (attack QPS peaks), `incount` (total number of requests), and `dropcount` (number of attack requests).
	MetricName *string `json:"MetricName,omitnil" name:"MetricName"`

	// (Optional) Domain name
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Instance ID. Leave this field empty when `Business` is `basic`, as basic protection does not require an instance.
	Id *string `json:"Id,omitnil" name:"Id"`
}

func (r *DescribeCCTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Ip")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricName")
	delete(f, "Domain")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCCTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCTrendResponseParams struct {
	// Number of values returned
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (Single IP); `bgp-multip`: Anti-DDoS Pro (Multi-IP); `net`: Anti-DDoS Ultimate; `basic`: Anti-DDoS Basic
	Business *string `json:"Business,omitnil" name:"Business"`

	// Instance IP
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Sampling period. Valid values: `300` (5 minutes), `3600` (one hour), `86400` (one day)
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// Beginning of the time range for the query
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End of the time range for the query
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Value array
	Data []*uint64 `json:"Data,omitnil" name:"Data"`

	// Instance ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Id *string `json:"Id,omitnil" name:"Id"`

	// Metric. Valid values: `inqps` (total QPS peaks), `dropqps` (attack QPS peaks), `incount` (total number of requests), and `dropcount` (number of attack requests).
	MetricName *string `json:"MetricName,omitnil" name:"MetricName"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCCTrendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCCTrendResponseParams `json:"Response"`
}

func (r *DescribeCCTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcBlackWhiteIpListRequestParams struct {
	// Anti-DDoS service type. Valid values: `bgpip-multip` (Anti-DDoS Pro) and `bgpip` (Anti-DDoS Advanced).
	Business *string `json:"Business,omitnil" name:"Business"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of results returned in one page
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// IP address, which is required when an Anti-DDoS Advanced instance is used.
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Domain name, which is required when an Anti-DDoS Advanced instance is used.
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Protocol, which is required when an Anti-DDoS Advanced instance is used.
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// Specifies a blocklist/allowlist IP.
	FilterIp *string `json:"FilterIp,omitnil" name:"FilterIp"`

	// Specifies whether is an IP blocklist or IP allowlist.
	FilterType *string `json:"FilterType,omitnil" name:"FilterType"`
}

type DescribeCcBlackWhiteIpListRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS service type. Valid values: `bgpip-multip` (Anti-DDoS Pro) and `bgpip` (Anti-DDoS Advanced).
	Business *string `json:"Business,omitnil" name:"Business"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of results returned in one page
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// IP address, which is required when an Anti-DDoS Advanced instance is used.
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Domain name, which is required when an Anti-DDoS Advanced instance is used.
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Protocol, which is required when an Anti-DDoS Advanced instance is used.
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// Specifies a blocklist/allowlist IP.
	FilterIp *string `json:"FilterIp,omitnil" name:"FilterIp"`

	// Specifies whether is an IP blocklist or IP allowlist.
	FilterType *string `json:"FilterType,omitnil" name:"FilterType"`
}

func (r *DescribeCcBlackWhiteIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcBlackWhiteIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Ip")
	delete(f, "Domain")
	delete(f, "Protocol")
	delete(f, "FilterIp")
	delete(f, "FilterType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCcBlackWhiteIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcBlackWhiteIpListResponseParams struct {
	// Number of policy lists
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// Information of the policy list
	CcBlackWhiteIpList []*CcBlackWhiteIpPolicy `json:"CcBlackWhiteIpList,omitnil" name:"CcBlackWhiteIpList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCcBlackWhiteIpListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCcBlackWhiteIpListResponseParams `json:"Response"`
}

func (r *DescribeCcBlackWhiteIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcBlackWhiteIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcGeoIPBlockConfigListRequestParams struct {
	// Anti-DDoS service type. Valid values: `bgpip-multip` (Anti-DDoS Pro) and `bgpip` (Anti-DDoS Advanced).
	Business *string `json:"Business,omitnil" name:"Business"`

	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of results returned in one page
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// IP address, which is required when an Anti-DDoS Advanced instance is used.
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Domain name, which is required when an Anti-DDoS Advanced instance is used.
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Protocol, which is required when an Anti-DDoS Advanced instance is used.
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`
}

type DescribeCcGeoIPBlockConfigListRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS service type. Valid values: `bgpip-multip` (Anti-DDoS Pro) and `bgpip` (Anti-DDoS Advanced).
	Business *string `json:"Business,omitnil" name:"Business"`

	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of results returned in one page
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// IP address, which is required when an Anti-DDoS Advanced instance is used.
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Domain name, which is required when an Anti-DDoS Advanced instance is used.
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Protocol, which is required when an Anti-DDoS Advanced instance is used.
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`
}

func (r *DescribeCcGeoIPBlockConfigListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcGeoIPBlockConfigListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstanceId")
	delete(f, "Ip")
	delete(f, "Domain")
	delete(f, "Protocol")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCcGeoIPBlockConfigListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcGeoIPBlockConfigListResponseParams struct {
	// Number of policy lists
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// Information of the policy list
	CcGeoIpPolicyList []*CcGeoIpPolicyNew `json:"CcGeoIpPolicyList,omitnil" name:"CcGeoIpPolicyList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCcGeoIPBlockConfigListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCcGeoIPBlockConfigListResponseParams `json:"Response"`
}

func (r *DescribeCcGeoIPBlockConfigListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcGeoIPBlockConfigListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSTrendRequestParams struct {
	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (Single IP); `bgp-multip`: Anti-DDoS Pro (Multi-IP); `net`: Anti-DDoS Ultimate; `basic`: Anti-DDoS Basic
	Business *string `json:"Business,omitnil" name:"Business"`

	// Instance IP
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Sampling period. Valid values: `300` (5 minutes), `3600` (one hour), `86400` (one day)
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// Beginning of the time range for the query
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End of the time range for the query
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Metric. Valid values: `bps`: attack traffic bandwidth; `pps`: attack packet rate
	MetricName *string `json:"MetricName,omitnil" name:"MetricName"`

	// Instance ID. Leave this field empty when `Business` is `basic`, as basic protection does not require an instance.
	Id *string `json:"Id,omitnil" name:"Id"`
}

type DescribeDDoSTrendRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (Single IP); `bgp-multip`: Anti-DDoS Pro (Multi-IP); `net`: Anti-DDoS Ultimate; `basic`: Anti-DDoS Basic
	Business *string `json:"Business,omitnil" name:"Business"`

	// Instance IP
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Sampling period. Valid values: `300` (5 minutes), `3600` (one hour), `86400` (one day)
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// Beginning of the time range for the query
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End of the time range for the query
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Metric. Valid values: `bps`: attack traffic bandwidth; `pps`: attack packet rate
	MetricName *string `json:"MetricName,omitnil" name:"MetricName"`

	// Instance ID. Leave this field empty when `Business` is `basic`, as basic protection does not require an instance.
	Id *string `json:"Id,omitnil" name:"Id"`
}

func (r *DescribeDDoSTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Ip")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricName")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSTrendResponseParams struct {
	// Number of values returned
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// Anti-DDoS service type. `bgpip`: Anti-DDoS Advanced; `bgp`: Anti-DDoS Pro (Single IP); `bgp-multip`: Anti-DDoS Pro (Multi-IP); `net`: Anti-DDoS Ultimate; `basic`: Anti-DDoS Basic
	Business *string `json:"Business,omitnil" name:"Business"`

	// Instance IP
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Sampling period. Valid values: `300` (5 minutes), `3600` (one hour), `86400` (one day)
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// Beginning of the time range for the query
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End of the time range for the query
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Value array. The unit for attack traffic bandwidth is Mbps, and that for the packet rate is pps.
	Data []*uint64 `json:"Data,omitnil" name:"Data"`

	// Instance ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Id *string `json:"Id,omitnil" name:"Id"`

	// Metric. Valid values: `bps`: attack traffic bandwidth; `pps`: attack packet rate
	MetricName *string `json:"MetricName,omitnil" name:"MetricName"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDDoSTrendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSTrendResponseParams `json:"Response"`
}

func (r *DescribeDDoSTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDefaultAlarmThresholdRequestParams struct {
	// Product category. Valid values:
	// `bgp`: Anti-DDoS Pro
	// `bgpip`: Anti-DDoS Advanced
	// ]
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// Alarm threshold type filter. Valid values:
	// `1`: alarm threshold for inbound traffic
	// `2`: alarm threshold for cleansing attack traffic
	// ]
	FilterAlarmType *int64 `json:"FilterAlarmType,omitnil" name:"FilterAlarmType"`
}

type DescribeDefaultAlarmThresholdRequest struct {
	*tchttp.BaseRequest
	
	// Product category. Valid values:
	// `bgp`: Anti-DDoS Pro
	// `bgpip`: Anti-DDoS Advanced
	// ]
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// Alarm threshold type filter. Valid values:
	// `1`: alarm threshold for inbound traffic
	// `2`: alarm threshold for cleansing attack traffic
	// ]
	FilterAlarmType *int64 `json:"FilterAlarmType,omitnil" name:"FilterAlarmType"`
}

func (r *DescribeDefaultAlarmThresholdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefaultAlarmThresholdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceType")
	delete(f, "FilterAlarmType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDefaultAlarmThresholdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDefaultAlarmThresholdResponseParams struct {
	// Default alarm threshold configuration
	DefaultAlarmConfigList []*DefaultAlarmThreshold `json:"DefaultAlarmConfigList,omitnil" name:"DefaultAlarmConfigList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDefaultAlarmThresholdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDefaultAlarmThresholdResponseParams `json:"Response"`
}

func (r *DescribeDefaultAlarmThresholdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefaultAlarmThresholdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL7RulesBySSLCertIdRequestParams struct {
	// Domain name status. Valid values: `bindable`, `binded`, `opened`, `closed`, `all` (all states).
	Status *string `json:"Status,omitnil" name:"Status"`

	// List of certificate IDs
	CertIds []*string `json:"CertIds,omitnil" name:"CertIds"`
}

type DescribeL7RulesBySSLCertIdRequest struct {
	*tchttp.BaseRequest
	
	// Domain name status. Valid values: `bindable`, `binded`, `opened`, `closed`, `all` (all states).
	Status *string `json:"Status,omitnil" name:"Status"`

	// List of certificate IDs
	CertIds []*string `json:"CertIds,omitnil" name:"CertIds"`
}

func (r *DescribeL7RulesBySSLCertIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL7RulesBySSLCertIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "CertIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeL7RulesBySSLCertIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeL7RulesBySSLCertIdResponseParams struct {
	// Certificate rule set
	CertSet []*CertIdInsL7Rules `json:"CertSet,omitnil" name:"CertSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeL7RulesBySSLCertIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeL7RulesBySSLCertIdResponseParams `json:"Response"`
}

func (r *DescribeL7RulesBySSLCertIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeL7RulesBySSLCertIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListBGPIPInstancesRequestParams struct {
	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of items per page. The default value is 20 when `Limit = 0`. The maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// IP filter
	FilterIp *string `json:"FilterIp,omitnil" name:"FilterIp"`

	// Anti-DDoS instance ID filter. For example, you can filter the Anti-DDoS Advanced instance ID by `bgpip-00000001`.
	FilterInstanceId *string `json:"FilterInstanceId,omitnil" name:"FilterInstanceId"`

	// Anti-DDoS Advanced line filter. Valid values:
	// `1`: BGP line
	// `2`: China Telecom
	// `3`: China Unicom
	// `4`: China Mobile
	// `99`: third-party line
	// ]
	FilterLine *uint64 `json:"FilterLine,omitnil" name:"FilterLine"`

	// Region filter. For example, `ap-guangzhou`.
	FilterRegion *string `json:"FilterRegion,omitnil" name:"FilterRegion"`

	// Name filter
	FilterName *string `json:"FilterName,omitnil" name:"FilterName"`

	// Whether to obtain only Anti-DDoS EIP instances. `1`: Yes; `0`: No.
	FilterEipType *int64 `json:"FilterEipType,omitnil" name:"FilterEipType"`

	// Anti-DDoS Advanced instance binding status filter. Valid values: `BINDING`, `BIND`, `UNBINDING`, `UNBIND`. This filter is only valid when `FilterEipType = 1`.
	FilterEipEipAddressStatus []*string `json:"FilterEipEipAddressStatus,omitnil" name:"FilterEipEipAddressStatus"`

	// Whether to obtain only Anti-DDoS instances with Sec-MCA enabled. Valid values: `1` (only obtain Anti-DDoS instances with Sec-MCA enabled) and `0` (obtain other Anti-DDoS instances).
	FilterDamDDoSStatus *int64 `json:"FilterDamDDoSStatus,omitnil" name:"FilterDamDDoSStatus"`

	// Filters by the status of bound resources. Values: `idle` (normal), `attacking` (being attacked), `blocking` (being blocked), `trial` (in trial)
	FilterStatus *string `json:"FilterStatus,omitnil" name:"FilterStatus"`

	// Filters by the instance CNAME
	FilterCname *string `json:"FilterCname,omitnil" name:"FilterCname"`

	// Filters by the instance ID
	FilterInstanceIdList []*string `json:"FilterInstanceIdList,omitnil" name:"FilterInstanceIdList"`

	// Searches by tag
	FilterTag *TagFilter `json:"FilterTag,omitnil" name:"FilterTag"`

	// Filters by package type.
	FilterPackType []*string `json:"FilterPackType,omitnil" name:"FilterPackType"`

	// Filters out Convoy instances
	FilterConvoy *uint64 `json:"FilterConvoy,omitnil" name:"FilterConvoy"`
}

type DescribeListBGPIPInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of items per page. The default value is 20 when `Limit = 0`. The maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// IP filter
	FilterIp *string `json:"FilterIp,omitnil" name:"FilterIp"`

	// Anti-DDoS instance ID filter. For example, you can filter the Anti-DDoS Advanced instance ID by `bgpip-00000001`.
	FilterInstanceId *string `json:"FilterInstanceId,omitnil" name:"FilterInstanceId"`

	// Anti-DDoS Advanced line filter. Valid values:
	// `1`: BGP line
	// `2`: China Telecom
	// `3`: China Unicom
	// `4`: China Mobile
	// `99`: third-party line
	// ]
	FilterLine *uint64 `json:"FilterLine,omitnil" name:"FilterLine"`

	// Region filter. For example, `ap-guangzhou`.
	FilterRegion *string `json:"FilterRegion,omitnil" name:"FilterRegion"`

	// Name filter
	FilterName *string `json:"FilterName,omitnil" name:"FilterName"`

	// Whether to obtain only Anti-DDoS EIP instances. `1`: Yes; `0`: No.
	FilterEipType *int64 `json:"FilterEipType,omitnil" name:"FilterEipType"`

	// Anti-DDoS Advanced instance binding status filter. Valid values: `BINDING`, `BIND`, `UNBINDING`, `UNBIND`. This filter is only valid when `FilterEipType = 1`.
	FilterEipEipAddressStatus []*string `json:"FilterEipEipAddressStatus,omitnil" name:"FilterEipEipAddressStatus"`

	// Whether to obtain only Anti-DDoS instances with Sec-MCA enabled. Valid values: `1` (only obtain Anti-DDoS instances with Sec-MCA enabled) and `0` (obtain other Anti-DDoS instances).
	FilterDamDDoSStatus *int64 `json:"FilterDamDDoSStatus,omitnil" name:"FilterDamDDoSStatus"`

	// Filters by the status of bound resources. Values: `idle` (normal), `attacking` (being attacked), `blocking` (being blocked), `trial` (in trial)
	FilterStatus *string `json:"FilterStatus,omitnil" name:"FilterStatus"`

	// Filters by the instance CNAME
	FilterCname *string `json:"FilterCname,omitnil" name:"FilterCname"`

	// Filters by the instance ID
	FilterInstanceIdList []*string `json:"FilterInstanceIdList,omitnil" name:"FilterInstanceIdList"`

	// Searches by tag
	FilterTag *TagFilter `json:"FilterTag,omitnil" name:"FilterTag"`

	// Filters by package type.
	FilterPackType []*string `json:"FilterPackType,omitnil" name:"FilterPackType"`

	// Filters out Convoy instances
	FilterConvoy *uint64 `json:"FilterConvoy,omitnil" name:"FilterConvoy"`
}

func (r *DescribeListBGPIPInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListBGPIPInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterIp")
	delete(f, "FilterInstanceId")
	delete(f, "FilterLine")
	delete(f, "FilterRegion")
	delete(f, "FilterName")
	delete(f, "FilterEipType")
	delete(f, "FilterEipEipAddressStatus")
	delete(f, "FilterDamDDoSStatus")
	delete(f, "FilterStatus")
	delete(f, "FilterCname")
	delete(f, "FilterInstanceIdList")
	delete(f, "FilterTag")
	delete(f, "FilterPackType")
	delete(f, "FilterConvoy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListBGPIPInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListBGPIPInstancesResponseParams struct {
	// Total number of lists
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// List of Anti-DDoS Advanced instances
	InstanceList []*BGPIPInstance `json:"InstanceList,omitnil" name:"InstanceList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeListBGPIPInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListBGPIPInstancesResponseParams `json:"Response"`
}

func (r *DescribeListBGPIPInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListBGPIPInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListBGPInstancesRequestParams struct {
	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of items per page. The default value is 20 when `Limit = 0`. The maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Filters by IP.
	FilterIp *string `json:"FilterIp,omitnil" name:"FilterIp"`

	// Anti-DDoS instance ID filter. For example, `bgp-00000001`.
	FilterInstanceId *string `json:"FilterInstanceId,omitnil" name:"FilterInstanceId"`

	// Filters by region. For example, `ap-guangzhou`.
	FilterRegion *string `json:"FilterRegion,omitnil" name:"FilterRegion"`

	// Filters by name.
	FilterName *string `json:"FilterName,omitnil" name:"FilterName"`

	// Line filter. Valid values: 1: BGP; 2: Non-BGP.
	FilterLine *uint64 `json:"FilterLine,omitnil" name:"FilterLine"`

	// Filters by instance status. `idle`: Running; `attacking`: Being attacked; `blocking`: Being blocked.
	FilterStatus *string `json:"FilterStatus,omitnil" name:"FilterStatus"`

	// Filters by binding status. `bounding`: The instance is bound; `failed`: The binding failed.
	FilterBoundStatus *string `json:"FilterBoundStatus,omitnil" name:"FilterBoundStatus"`

	// Array of instance IDs
	FilterInstanceIdList []*string `json:"FilterInstanceIdList,omitnil" name:"FilterInstanceIdList"`

	// Enterprise edition. Values: `1` (the Convoy package included), `2` (the Convoy package not included)
	FilterEnterpriseFlag *uint64 `json:"FilterEnterpriseFlag,omitnil" name:"FilterEnterpriseFlag"`

	// Whether it’s a Lighthouse edition
	FilterLightFlag *uint64 `json:"FilterLightFlag,omitnil" name:"FilterLightFlag"`

	// Whether it’s a Channel edition
	FilterChannelFlag *uint64 `json:"FilterChannelFlag,omitnil" name:"FilterChannelFlag"`

	// Filters by tag
	FilterTag *TagFilter `json:"FilterTag,omitnil" name:"FilterTag"`

	// Filters out trial instances. Values: `1` (emergency protection instances), `2` (PLG instances)
	FilterTrialFlag *uint64 `json:"FilterTrialFlag,omitnil" name:"FilterTrialFlag"`

	// Filters out Convoy instances
	FilterConvoy *uint64 `json:"FilterConvoy,omitnil" name:"FilterConvoy"`

	// Whether to exclude the advanced information (such as `InstanceList[0].Usage`). Values: `true` (exclude), `false` (do not exclude). The default value is `false`.
	ExcludeAdvancedInfo *bool `json:"ExcludeAdvancedInfo,omitnil" name:"ExcludeAdvancedInfo"`
}

type DescribeListBGPInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of items per page. The default value is 20 when `Limit = 0`. The maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Filters by IP.
	FilterIp *string `json:"FilterIp,omitnil" name:"FilterIp"`

	// Anti-DDoS instance ID filter. For example, `bgp-00000001`.
	FilterInstanceId *string `json:"FilterInstanceId,omitnil" name:"FilterInstanceId"`

	// Filters by region. For example, `ap-guangzhou`.
	FilterRegion *string `json:"FilterRegion,omitnil" name:"FilterRegion"`

	// Filters by name.
	FilterName *string `json:"FilterName,omitnil" name:"FilterName"`

	// Line filter. Valid values: 1: BGP; 2: Non-BGP.
	FilterLine *uint64 `json:"FilterLine,omitnil" name:"FilterLine"`

	// Filters by instance status. `idle`: Running; `attacking`: Being attacked; `blocking`: Being blocked.
	FilterStatus *string `json:"FilterStatus,omitnil" name:"FilterStatus"`

	// Filters by binding status. `bounding`: The instance is bound; `failed`: The binding failed.
	FilterBoundStatus *string `json:"FilterBoundStatus,omitnil" name:"FilterBoundStatus"`

	// Array of instance IDs
	FilterInstanceIdList []*string `json:"FilterInstanceIdList,omitnil" name:"FilterInstanceIdList"`

	// Enterprise edition. Values: `1` (the Convoy package included), `2` (the Convoy package not included)
	FilterEnterpriseFlag *uint64 `json:"FilterEnterpriseFlag,omitnil" name:"FilterEnterpriseFlag"`

	// Whether it’s a Lighthouse edition
	FilterLightFlag *uint64 `json:"FilterLightFlag,omitnil" name:"FilterLightFlag"`

	// Whether it’s a Channel edition
	FilterChannelFlag *uint64 `json:"FilterChannelFlag,omitnil" name:"FilterChannelFlag"`

	// Filters by tag
	FilterTag *TagFilter `json:"FilterTag,omitnil" name:"FilterTag"`

	// Filters out trial instances. Values: `1` (emergency protection instances), `2` (PLG instances)
	FilterTrialFlag *uint64 `json:"FilterTrialFlag,omitnil" name:"FilterTrialFlag"`

	// Filters out Convoy instances
	FilterConvoy *uint64 `json:"FilterConvoy,omitnil" name:"FilterConvoy"`

	// Whether to exclude the advanced information (such as `InstanceList[0].Usage`). Values: `true` (exclude), `false` (do not exclude). The default value is `false`.
	ExcludeAdvancedInfo *bool `json:"ExcludeAdvancedInfo,omitnil" name:"ExcludeAdvancedInfo"`
}

func (r *DescribeListBGPInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListBGPInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterIp")
	delete(f, "FilterInstanceId")
	delete(f, "FilterRegion")
	delete(f, "FilterName")
	delete(f, "FilterLine")
	delete(f, "FilterStatus")
	delete(f, "FilterBoundStatus")
	delete(f, "FilterInstanceIdList")
	delete(f, "FilterEnterpriseFlag")
	delete(f, "FilterLightFlag")
	delete(f, "FilterChannelFlag")
	delete(f, "FilterTag")
	delete(f, "FilterTrialFlag")
	delete(f, "FilterConvoy")
	delete(f, "ExcludeAdvancedInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListBGPInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListBGPInstancesResponseParams struct {
	// Total number of Anti-DDoS Pro instances
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// List of Anti-DDoS Pro instances
	InstanceList []*BGPInstance `json:"InstanceList,omitnil" name:"InstanceList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeListBGPInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListBGPInstancesResponseParams `json:"Response"`
}

func (r *DescribeListBGPInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListBGPInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListBlackWhiteIpListRequestParams struct {
	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of items per page. The default value is 20 when Limit = 0. The maximum value is 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Anti-DDoS instance ID filter. Anti-DDoS instance prefix wildcard search is supported. For example, you can filter Anti-DDoS Pro instances by `bgp-*`.
	FilterInstanceId *string `json:"FilterInstanceId,omitnil" name:"FilterInstanceId"`

	// IP filter
	FilterIp *string `json:"FilterIp,omitnil" name:"FilterIp"`
}

type DescribeListBlackWhiteIpListRequest struct {
	*tchttp.BaseRequest
	
	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of items per page. The default value is 20 when Limit = 0. The maximum value is 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Anti-DDoS instance ID filter. Anti-DDoS instance prefix wildcard search is supported. For example, you can filter Anti-DDoS Pro instances by `bgp-*`.
	FilterInstanceId *string `json:"FilterInstanceId,omitnil" name:"FilterInstanceId"`

	// IP filter
	FilterIp *string `json:"FilterIp,omitnil" name:"FilterIp"`
}

func (r *DescribeListBlackWhiteIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListBlackWhiteIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterInstanceId")
	delete(f, "FilterIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListBlackWhiteIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListBlackWhiteIpListResponseParams struct {
	// Total number of lists
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// IP blocklist/allowlist
	IpList []*BlackWhiteIpRelation `json:"IpList,omitnil" name:"IpList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeListBlackWhiteIpListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListBlackWhiteIpListResponseParams `json:"Response"`
}

func (r *DescribeListBlackWhiteIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListBlackWhiteIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListDDoSAIRequestParams struct {
	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of items per page. The default value is 20 when Limit = 0. The maximum value is 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Anti-DDoS instance ID filter. Anti-DDoS instance prefix wildcard search is supported. For example, you can filter Anti-DDoS Pro instances by `bgp-*`.
	FilterInstanceId *string `json:"FilterInstanceId,omitnil" name:"FilterInstanceId"`

	// IP filter
	FilterIp *string `json:"FilterIp,omitnil" name:"FilterIp"`
}

type DescribeListDDoSAIRequest struct {
	*tchttp.BaseRequest
	
	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of items per page. The default value is 20 when Limit = 0. The maximum value is 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Anti-DDoS instance ID filter. Anti-DDoS instance prefix wildcard search is supported. For example, you can filter Anti-DDoS Pro instances by `bgp-*`.
	FilterInstanceId *string `json:"FilterInstanceId,omitnil" name:"FilterInstanceId"`

	// IP filter
	FilterIp *string `json:"FilterIp,omitnil" name:"FilterIp"`
}

func (r *DescribeListDDoSAIRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListDDoSAIRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterInstanceId")
	delete(f, "FilterIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListDDoSAIRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListDDoSAIResponseParams struct {
	// Total number of lists
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// List of AI protection switches
	ConfigList []*DDoSAIRelation `json:"ConfigList,omitnil" name:"ConfigList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeListDDoSAIResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListDDoSAIResponseParams `json:"Response"`
}

func (r *DescribeListDDoSAIResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListDDoSAIResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListDDoSGeoIPBlockConfigRequestParams struct {
	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of items per page. The default value is 20 when Limit = 0. The maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Anti-DDoS instance ID filter. Anti-DDoS instance prefix wildcard search is supported. For example, you can filter Anti-DDoS Pro instances by `bgp-*`.
	FilterInstanceId *string `json:"FilterInstanceId,omitnil" name:"FilterInstanceId"`

	// IP filter
	FilterIp *string `json:"FilterIp,omitnil" name:"FilterIp"`
}

type DescribeListDDoSGeoIPBlockConfigRequest struct {
	*tchttp.BaseRequest
	
	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of items per page. The default value is 20 when Limit = 0. The maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Anti-DDoS instance ID filter. Anti-DDoS instance prefix wildcard search is supported. For example, you can filter Anti-DDoS Pro instances by `bgp-*`.
	FilterInstanceId *string `json:"FilterInstanceId,omitnil" name:"FilterInstanceId"`

	// IP filter
	FilterIp *string `json:"FilterIp,omitnil" name:"FilterIp"`
}

func (r *DescribeListDDoSGeoIPBlockConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListDDoSGeoIPBlockConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterInstanceId")
	delete(f, "FilterIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListDDoSGeoIPBlockConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListDDoSGeoIPBlockConfigResponseParams struct {
	// Total number of lists
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// List of Anti-DDoS region blocking configurations
	ConfigList []*DDoSGeoIPBlockConfigRelation `json:"ConfigList,omitnil" name:"ConfigList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeListDDoSGeoIPBlockConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListDDoSGeoIPBlockConfigResponseParams `json:"Response"`
}

func (r *DescribeListDDoSGeoIPBlockConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListDDoSGeoIPBlockConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListDDoSSpeedLimitConfigRequestParams struct {
	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of items per page. The default value is 20 when Limit = 0. The maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Anti-DDoS instance ID filter. Anti-DDoS instance prefix wildcard search is supported. For example, you can filter Anti-DDoS Pro instances by `bgp-*`.
	FilterInstanceId *string `json:"FilterInstanceId,omitnil" name:"FilterInstanceId"`

	// IP filter
	FilterIp *string `json:"FilterIp,omitnil" name:"FilterIp"`
}

type DescribeListDDoSSpeedLimitConfigRequest struct {
	*tchttp.BaseRequest
	
	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of items per page. The default value is 20 when Limit = 0. The maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Anti-DDoS instance ID filter. Anti-DDoS instance prefix wildcard search is supported. For example, you can filter Anti-DDoS Pro instances by `bgp-*`.
	FilterInstanceId *string `json:"FilterInstanceId,omitnil" name:"FilterInstanceId"`

	// IP filter
	FilterIp *string `json:"FilterIp,omitnil" name:"FilterIp"`
}

func (r *DescribeListDDoSSpeedLimitConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListDDoSSpeedLimitConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterInstanceId")
	delete(f, "FilterIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListDDoSSpeedLimitConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListDDoSSpeedLimitConfigResponseParams struct {
	// Total number of lists
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// List of access rate limit configurations
	ConfigList []*DDoSSpeedLimitConfigRelation `json:"ConfigList,omitnil" name:"ConfigList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeListDDoSSpeedLimitConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListDDoSSpeedLimitConfigResponseParams `json:"Response"`
}

func (r *DescribeListDDoSSpeedLimitConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListDDoSSpeedLimitConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListIPAlarmConfigRequestParams struct {
	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of items per page. The default value is 20 when Limit = 0. The maximum value is 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Anti-DDoS instance ID filter. Anti-DDoS instance prefix wildcard search is supported. For example, you can filter Anti-DDoS Pro instances by `bgp-*`.
	FilterInstanceId *string `json:"FilterInstanceId,omitnil" name:"FilterInstanceId"`

	// Alarm threshold type filter. Valid values:
	// `1`: alarm threshold for inbound traffic
	// `2`: alarm threshold for cleansing attack traffic
	// ]
	FilterAlarmType *int64 `json:"FilterAlarmType,omitnil" name:"FilterAlarmType"`

	// IP filter
	FilterIp *string `json:"FilterIp,omitnil" name:"FilterIp"`

	// CNAME of the Anti-DDoS Advanced instance
	FilterCname *string `json:"FilterCname,omitnil" name:"FilterCname"`
}

type DescribeListIPAlarmConfigRequest struct {
	*tchttp.BaseRequest
	
	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of items per page. The default value is 20 when Limit = 0. The maximum value is 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Anti-DDoS instance ID filter. Anti-DDoS instance prefix wildcard search is supported. For example, you can filter Anti-DDoS Pro instances by `bgp-*`.
	FilterInstanceId *string `json:"FilterInstanceId,omitnil" name:"FilterInstanceId"`

	// Alarm threshold type filter. Valid values:
	// `1`: alarm threshold for inbound traffic
	// `2`: alarm threshold for cleansing attack traffic
	// ]
	FilterAlarmType *int64 `json:"FilterAlarmType,omitnil" name:"FilterAlarmType"`

	// IP filter
	FilterIp *string `json:"FilterIp,omitnil" name:"FilterIp"`

	// CNAME of the Anti-DDoS Advanced instance
	FilterCname *string `json:"FilterCname,omitnil" name:"FilterCname"`
}

func (r *DescribeListIPAlarmConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListIPAlarmConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterInstanceId")
	delete(f, "FilterAlarmType")
	delete(f, "FilterIp")
	delete(f, "FilterCname")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListIPAlarmConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListIPAlarmConfigResponseParams struct {
	// Total number of lists
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// List of IP alarm threshold configurations
	ConfigList []*IPAlarmThresholdRelation `json:"ConfigList,omitnil" name:"ConfigList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeListIPAlarmConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListIPAlarmConfigResponseParams `json:"Response"`
}

func (r *DescribeListIPAlarmConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListIPAlarmConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListListenerRequestParams struct {

}

type DescribeListListenerRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeListListenerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListListenerResponseParams struct {
	// List of layer-4 forwarding listeners
	Layer4Listeners []*Layer4Rule `json:"Layer4Listeners,omitnil" name:"Layer4Listeners"`

	// List of layer-7 forwarding listeners
	Layer7Listeners []*Layer7Rule `json:"Layer7Listeners,omitnil" name:"Layer7Listeners"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeListListenerResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListListenerResponseParams `json:"Response"`
}

func (r *DescribeListListenerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListPacketFilterConfigRequestParams struct {
	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of items per page. The default value is 20 when Limit = 0. The maximum value is 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Anti-DDoS instance ID filter. Anti-DDoS instance prefix wildcard search is supported. For example, you can filter Anti-DDoS Pro instances by `bgp-*`.
	FilterInstanceId *string `json:"FilterInstanceId,omitnil" name:"FilterInstanceId"`

	// IP filter
	FilterIp *string `json:"FilterIp,omitnil" name:"FilterIp"`
}

type DescribeListPacketFilterConfigRequest struct {
	*tchttp.BaseRequest
	
	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of items per page. The default value is 20 when Limit = 0. The maximum value is 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Anti-DDoS instance ID filter. Anti-DDoS instance prefix wildcard search is supported. For example, you can filter Anti-DDoS Pro instances by `bgp-*`.
	FilterInstanceId *string `json:"FilterInstanceId,omitnil" name:"FilterInstanceId"`

	// IP filter
	FilterIp *string `json:"FilterIp,omitnil" name:"FilterIp"`
}

func (r *DescribeListPacketFilterConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListPacketFilterConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterInstanceId")
	delete(f, "FilterIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListPacketFilterConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListPacketFilterConfigResponseParams struct {
	// Total number of lists
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// Feature filtering configuration
	ConfigList []*PacketFilterRelation `json:"ConfigList,omitnil" name:"ConfigList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeListPacketFilterConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListPacketFilterConfigResponseParams `json:"Response"`
}

func (r *DescribeListPacketFilterConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListPacketFilterConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListProtectThresholdConfigRequestParams struct {
	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of items per page. The default value is 20 when `Limit = 0`. The maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Anti-DDoS instance ID filter. Anti-DDoS instance prefix wildcard search is supported. For example, you can filter Anti-DDoS Pro instances by `bgp-*`.
	FilterInstanceId *string `json:"FilterInstanceId,omitnil" name:"FilterInstanceId"`

	// IP filter
	FilterIp *string `json:"FilterIp,omitnil" name:"FilterIp"`

	// Domain name filter for querying CC protection thresholds of domain names and protocols
	FilterDomain *string `json:"FilterDomain,omitnil" name:"FilterDomain"`

	// Protocol filter for querying CC protection thresholds of domain names and protocols
	FilterProtocol *string `json:"FilterProtocol,omitnil" name:"FilterProtocol"`
}

type DescribeListProtectThresholdConfigRequest struct {
	*tchttp.BaseRequest
	
	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of items per page. The default value is 20 when `Limit = 0`. The maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Anti-DDoS instance ID filter. Anti-DDoS instance prefix wildcard search is supported. For example, you can filter Anti-DDoS Pro instances by `bgp-*`.
	FilterInstanceId *string `json:"FilterInstanceId,omitnil" name:"FilterInstanceId"`

	// IP filter
	FilterIp *string `json:"FilterIp,omitnil" name:"FilterIp"`

	// Domain name filter for querying CC protection thresholds of domain names and protocols
	FilterDomain *string `json:"FilterDomain,omitnil" name:"FilterDomain"`

	// Protocol filter for querying CC protection thresholds of domain names and protocols
	FilterProtocol *string `json:"FilterProtocol,omitnil" name:"FilterProtocol"`
}

func (r *DescribeListProtectThresholdConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListProtectThresholdConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterInstanceId")
	delete(f, "FilterIp")
	delete(f, "FilterDomain")
	delete(f, "FilterProtocol")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListProtectThresholdConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListProtectThresholdConfigResponseParams struct {
	// Total number of lists
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// List of protection threshold configurations
	ConfigList []*ProtectThresholdRelation `json:"ConfigList,omitnil" name:"ConfigList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeListProtectThresholdConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListProtectThresholdConfigResponseParams `json:"Response"`
}

func (r *DescribeListProtectThresholdConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListProtectThresholdConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListProtocolBlockConfigRequestParams struct {
	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of items per page. The default value is 20 when `Limit = 0`. The maximum value is 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Anti-DDoS instance ID filter. Anti-DDoS instance prefix wildcard search is supported. For example, you can filter Anti-DDoS Pro instances by `bgp-*`.
	FilterInstanceId *string `json:"FilterInstanceId,omitnil" name:"FilterInstanceId"`

	// IP filter
	FilterIp *string `json:"FilterIp,omitnil" name:"FilterIp"`
}

type DescribeListProtocolBlockConfigRequest struct {
	*tchttp.BaseRequest
	
	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of items per page. The default value is 20 when `Limit = 0`. The maximum value is 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Anti-DDoS instance ID filter. Anti-DDoS instance prefix wildcard search is supported. For example, you can filter Anti-DDoS Pro instances by `bgp-*`.
	FilterInstanceId *string `json:"FilterInstanceId,omitnil" name:"FilterInstanceId"`

	// IP filter
	FilterIp *string `json:"FilterIp,omitnil" name:"FilterIp"`
}

func (r *DescribeListProtocolBlockConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListProtocolBlockConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterInstanceId")
	delete(f, "FilterIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListProtocolBlockConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListProtocolBlockConfigResponseParams struct {
	// Total number of lists
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// Protocol blocking configuration
	ConfigList []*ProtocolBlockRelation `json:"ConfigList,omitnil" name:"ConfigList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeListProtocolBlockConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListProtocolBlockConfigResponseParams `json:"Response"`
}

func (r *DescribeListProtocolBlockConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListProtocolBlockConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListSchedulingDomainRequestParams struct {
	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of items per page. The default value is 20 when `Limit = 0`. The maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Scheduling domain name filter
	FilterDomain *string `json:"FilterDomain,omitnil" name:"FilterDomain"`
}

type DescribeListSchedulingDomainRequest struct {
	*tchttp.BaseRequest
	
	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of items per page. The default value is 20 when `Limit = 0`. The maximum value is 100.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Scheduling domain name filter
	FilterDomain *string `json:"FilterDomain,omitnil" name:"FilterDomain"`
}

func (r *DescribeListSchedulingDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListSchedulingDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterDomain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListSchedulingDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListSchedulingDomainResponseParams struct {
	// Total number of lists
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// List of scheduling domain names
	DomainList []*SchedulingDomainInfo `json:"DomainList,omitnil" name:"DomainList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeListSchedulingDomainResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListSchedulingDomainResponseParams `json:"Response"`
}

func (r *DescribeListSchedulingDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListSchedulingDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListWaterPrintConfigRequestParams struct {
	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of items per page. The default value is 20 when `Limit = 0`. The maximum value is 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Anti-DDoS instance ID filter. Anti-DDoS instance prefix wildcard search is supported. For example, you can filter Anti-DDoS Pro instances by `bgp-*`.
	FilterInstanceId *string `json:"FilterInstanceId,omitnil" name:"FilterInstanceId"`

	// IP filter
	FilterIp *string `json:"FilterIp,omitnil" name:"FilterIp"`
}

type DescribeListWaterPrintConfigRequest struct {
	*tchttp.BaseRequest
	
	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of items per page. The default value is 20 when `Limit = 0`. The maximum value is 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Anti-DDoS instance ID filter. Anti-DDoS instance prefix wildcard search is supported. For example, you can filter Anti-DDoS Pro instances by `bgp-*`.
	FilterInstanceId *string `json:"FilterInstanceId,omitnil" name:"FilterInstanceId"`

	// IP filter
	FilterIp *string `json:"FilterIp,omitnil" name:"FilterIp"`
}

func (r *DescribeListWaterPrintConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListWaterPrintConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterInstanceId")
	delete(f, "FilterIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListWaterPrintConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListWaterPrintConfigResponseParams struct {
	// Total number of lists
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// List of watermark configurations
	ConfigList []*WaterPrintRelation `json:"ConfigList,omitnil" name:"ConfigList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeListWaterPrintConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListWaterPrintConfigResponseParams `json:"Response"`
}

func (r *DescribeListWaterPrintConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListWaterPrintConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNewL7RulesErrHealthRequestParams struct {
	// Anti-DDoS service type (`bgpip`: Anti-DDoS Advanced)
	Business *string `json:"Business,omitnil" name:"Business"`

	// List of rule IDs
	RuleIdList []*string `json:"RuleIdList,omitnil" name:"RuleIdList"`
}

type DescribeNewL7RulesErrHealthRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS service type (`bgpip`: Anti-DDoS Advanced)
	Business *string `json:"Business,omitnil" name:"Business"`

	// List of rule IDs
	RuleIdList []*string `json:"RuleIdList,omitnil" name:"RuleIdList"`
}

func (r *DescribeNewL7RulesErrHealthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNewL7RulesErrHealthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "RuleIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNewL7RulesErrHealthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNewL7RulesErrHealthResponseParams struct {
	// List of rules with exceptions. `Key`: Rule ID, `Value`: Exception IPs and error message. 
	ErrHealths []*KeyValue `json:"ErrHealths,omitnil" name:"ErrHealths"`

	// Total number of rules with exceptions
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeNewL7RulesErrHealthResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNewL7RulesErrHealthResponseParams `json:"Response"`
}

func (r *DescribeNewL7RulesErrHealthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNewL7RulesErrHealthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNewL7RulesRequestParams struct {
	// Anti-DDoS service type (`bgpip`: Anti-DDoS Advanced)
	Business *string `json:"Business,omitnil" name:"Business"`

	// (Optional) Searches by rule status. Valid values: `0` (Successfully configured), `1` (Being configured), `2` (Configuration failed), `3` (Being deleted), `5` (Deletion failed), `6` (awaiting configuration), `7` (awaiting deletion), and `8` (awaiting certificate configuration).
	StatusList []*uint64 `json:"StatusList,omitnil" name:"StatusList"`

	// (Optional) Searches by domain name.
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// (Optional) Searches by IP.
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Number of items in a page. Returned results are not paged if you enter “0”.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// (Optional) Searches by forwarding protocol. Values: [http, https, http/https]
	ProtocolList []*string `json:"ProtocolList,omitnil" name:"ProtocolList"`

	// CNAME of the Anti-DDoS Advanced instance
	Cname *string `json:"Cname,omitnil" name:"Cname"`
}

type DescribeNewL7RulesRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS service type (`bgpip`: Anti-DDoS Advanced)
	Business *string `json:"Business,omitnil" name:"Business"`

	// (Optional) Searches by rule status. Valid values: `0` (Successfully configured), `1` (Being configured), `2` (Configuration failed), `3` (Being deleted), `5` (Deletion failed), `6` (awaiting configuration), `7` (awaiting deletion), and `8` (awaiting certificate configuration).
	StatusList []*uint64 `json:"StatusList,omitnil" name:"StatusList"`

	// (Optional) Searches by domain name.
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// (Optional) Searches by IP.
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Number of items in a page. Returned results are not paged if you enter “0”.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Starting offset of the page. Value: (number of pages – 1) * items per page.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// (Optional) Searches by forwarding protocol. Values: [http, https, http/https]
	ProtocolList []*string `json:"ProtocolList,omitnil" name:"ProtocolList"`

	// CNAME of the Anti-DDoS Advanced instance
	Cname *string `json:"Cname,omitnil" name:"Cname"`
}

func (r *DescribeNewL7RulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNewL7RulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "StatusList")
	delete(f, "Domain")
	delete(f, "Ip")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ProtocolList")
	delete(f, "Cname")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNewL7RulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNewL7RulesResponseParams struct {
	// List of forwarding rules
	Rules []*NewL7RuleEntry `json:"Rules,omitnil" name:"Rules"`

	// List of health check settings
	Healths []*L7RuleHealth `json:"Healths,omitnil" name:"Healths"`

	// Total number of rules
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeNewL7RulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNewL7RulesResponseParams `json:"Response"`
}

func (r *DescribeNewL7RulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNewL7RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOverviewDDoSEventListRequestParams struct {
	// Start time
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Filters by the attack status. `start`: The attack is ongoing; `end`: The attack ends.
	AttackStatus *string `json:"AttackStatus,omitnil" name:"AttackStatus"`

	// The offset value
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Total number of records
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeOverviewDDoSEventListRequest struct {
	*tchttp.BaseRequest
	
	// Start time
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Filters by the attack status. `start`: The attack is ongoing; `end`: The attack ends.
	AttackStatus *string `json:"AttackStatus,omitnil" name:"AttackStatus"`

	// The offset value
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Total number of records
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeOverviewDDoSEventListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOverviewDDoSEventListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "AttackStatus")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOverviewDDoSEventListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOverviewDDoSEventListResponseParams struct {
	// Total number of records
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// Event list
	EventList []*OverviewDDoSEvent `json:"EventList,omitnil" name:"EventList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeOverviewDDoSEventListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOverviewDDoSEventListResponseParams `json:"Response"`
}

func (r *DescribeOverviewDDoSEventListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOverviewDDoSEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePendingRiskInfoRequestParams struct {

}

type DescribePendingRiskInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribePendingRiskInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePendingRiskInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePendingRiskInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePendingRiskInfoResponseParams struct {
	// Whether the user is a paid user. Values: `true`, `false`.
	IsPaidUsr *bool `json:"IsPaidUsr,omitnil" name:"IsPaidUsr"`

	// Number of resources being attacked
	AttackingCount *int64 `json:"AttackingCount,omitnil" name:"AttackingCount"`

	// Number of resource blocked
	BlockingCount *int64 `json:"BlockingCount,omitnil" name:"BlockingCount"`

	// Number of expired resources
	ExpiredCount *int64 `json:"ExpiredCount,omitnil" name:"ExpiredCount"`

	// Total pending risk events
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribePendingRiskInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribePendingRiskInfoResponseParams `json:"Response"`
}

func (r *DescribePendingRiskInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePendingRiskInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateDDoSEipAddressRequestParams struct {
	// Anti-DDoS instance ID (only Anti-DDoS Advanced). For example, `bgpip-0000011x`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// EIP of the Anti-DDoS instance ID
	Eip *string `json:"Eip,omitnil" name:"Eip"`
}

type DisassociateDDoSEipAddressRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS instance ID (only Anti-DDoS Advanced). For example, `bgpip-0000011x`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// EIP of the Anti-DDoS instance ID
	Eip *string `json:"Eip,omitnil" name:"Eip"`
}

func (r *DisassociateDDoSEipAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateDDoSEipAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Eip")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateDDoSEipAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateDDoSEipAddressResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DisassociateDDoSEipAddressResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateDDoSEipAddressResponseParams `json:"Response"`
}

func (r *DisassociateDDoSEipAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateDDoSEipAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EipAddressPackRelation struct {
	// Number of package IPs
	IpCount *uint64 `json:"IpCount,omitnil" name:"IpCount"`

	// Auto-renewal flag
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// Current expiration time
	CurDeadline *string `json:"CurDeadline,omitnil" name:"CurDeadline"`
}

type EipAddressRelation struct {
	// Region of the Anti-DDoS instance bound to the EIP. For example, hk indicates Hong Kong.
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	EipAddressRegion *string `json:"EipAddressRegion,omitnil" name:"EipAddressRegion"`

	// ID of the bound resource. For example, an ID may be bound to an CVM instance.
	// Note: This is field may return `null`, indicating that no valid value can be obtained.
	EipBoundRscIns *string `json:"EipBoundRscIns,omitnil" name:"EipBoundRscIns"`

	// ID of the bound ENI
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	EipBoundRscEni *string `json:"EipBoundRscEni,omitnil" name:"EipBoundRscEni"`

	// Private IP of the bound resource
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	EipBoundRscVip *string `json:"EipBoundRscVip,omitnil" name:"EipBoundRscVip"`

	// Modification time
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`
}

type EipProductInfo struct {
	// IP address
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Cloud product category. Valid values:
	// `public`: CVM
	// `bm`: BM
	// `eni`: ENI
	// `vpngw`: VPN gateway
	//  `natgw`: NAT gateway
	// `waf`: WAF
	// `fpc`: financial products
	// `gaap`: GAAP 
	// `other`: hosted IP
	// ]
	BizType *string `json:"BizType,omitnil" name:"BizType"`

	// Cloud sub-product category. Valid values: `cvm` (CVM), `lb` (Load balancer), `eni` (ENI), `vpngw` (VPN gateway), `natgw` (NAT gateway), `waf` (WAF), `fpc` (financial products), `gaap` (GAAP), `eip` (BM EIP) and `other` (hosted IP).
	DeviceType *string `json:"DeviceType,omitnil" name:"DeviceType"`

	// Cloud instance ID of the IP. This field InstanceId will be `eni-*` if the instance ID is bound to an ENI IP; `none` if there is no instance ID to bind to a hosted IP.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type ForwardListener struct {
	// The starting port for listener forwarding. Value range: 1 to 65535.
	FrontendPort *int64 `json:"FrontendPort,omitnil" name:"FrontendPort"`

	// Forwarding protocol. Valid values:
	// `TCP`
	// `UDP`
	// ]
	ForwardProtocol *string `json:"ForwardProtocol,omitnil" name:"ForwardProtocol"`

	// The ending port for listener forwarding. Value range: 1 to 65535.
	FrontendPortEnd *int64 `json:"FrontendPortEnd,omitnil" name:"FrontendPortEnd"`
}

type HttpStatusMap struct {
	// HTTP 2xx Forwarding status code
	SourceHttp2xx []*float64 `json:"SourceHttp2xx,omitnil" name:"SourceHttp2xx"`

	// HTTP 5xx Status code
	Http5xx []*float64 `json:"Http5xx,omitnil" name:"Http5xx"`

	// HTTP 5xx Forwarding status code
	SourceHttp5xx []*float64 `json:"SourceHttp5xx,omitnil" name:"SourceHttp5xx"`

	// HTTP 404 Forwarding status code
	SourceHttp404 []*float64 `json:"SourceHttp404,omitnil" name:"SourceHttp404"`

	// HTTP 4xx Status code
	Http4xx []*float64 `json:"Http4xx,omitnil" name:"Http4xx"`

	// HTTP 4xx Forwarding status code
	SourceHttp4xx []*float64 `json:"SourceHttp4xx,omitnil" name:"SourceHttp4xx"`

	// HTTP 2xx Status code
	Http2xx []*float64 `json:"Http2xx,omitnil" name:"Http2xx"`

	// HTTP 404 Status code
	Http404 []*float64 `json:"Http404,omitnil" name:"Http404"`

	// HTTP 3xx Forwarding status code
	SourceHttp3xx []*float64 `json:"SourceHttp3xx,omitnil" name:"SourceHttp3xx"`

	// HTTP 3xx Status code
	Http3xx []*float64 `json:"Http3xx,omitnil" name:"Http3xx"`
}

type IPAlarmThresholdRelation struct {
	// Alarm threshold type. Valid values:
	// `1`: alarm threshold for inbound traffic
	// `2`: alarm threshold for cleansing attack traffic
	// ]
	AlarmType *uint64 `json:"AlarmType,omitnil" name:"AlarmType"`

	// Alarm threshold (Mbps). The value should be greater than or equal to 0. Note that the alarm threshold configuration will be removed if you pass the parameter for input and set it to 0.
	AlarmThreshold *uint64 `json:"AlarmThreshold,omitnil" name:"AlarmThreshold"`

	// Anti-DDoS instance configured
	InstanceDetailList []*InstanceRelation `json:"InstanceDetailList,omitnil" name:"InstanceDetailList"`
}

type IPLineInfo struct {
	// IP line type. Valid values:
	// `bgp`: BGP IP
	// `ctcc`: CTCC IP
	// `cucc`: CUCC IP
	// `cmcc`: CMCC IP
	// `abroad`: IP outside the Chinese mainland
	// ]
	Type *string `json:"Type,omitnil" name:"Type"`


	Eip *string `json:"Eip,omitnil" name:"Eip"`

	// CNAME of the instance
	Cname *string `json:"Cname,omitnil" name:"Cname"`

	// Flag of the instance. `0`: Anti-DDoS Pro instance; `1`: Anti-DDoS Advanced instance; `2`: Non-Anti-DDoS Advanced instance.
	ResourceFlag *int64 `json:"ResourceFlag,omitnil" name:"ResourceFlag"`
}

type InsL7Rules struct {
	// Rules can only be modified when the status is `0`, `2`, or `8`.
	// Rule status. Values: `0` (Normal), `1` (Being configured), `2` (Configuration failed), `3` (Being deleted), `5` (Failed to be deleted), `6` (Pending add), `7` (Pending delete), `8` (Pending certificate upload), `9` (Associated resource not exist), `10` (Pending modify), `11` (Being modified).
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// Domain name
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Protocol
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// Instance ID
	InsId *string `json:"InsId,omitnil" name:"InsId"`

	// User App ID
	AppId *string `json:"AppId,omitnil" name:"AppId"`

	// High-defense port
	VirtualPort *string `json:"VirtualPort,omitnil" name:"VirtualPort"`

	// Certificate ID
	SSLId *string `json:"SSLId,omitnil" name:"SSLId"`
}

type InstanceRelation struct {
	// Instance IP
	EipList []*string `json:"EipList,omitnil" name:"EipList"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type IpSegment struct {
	// IP address
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// IP mask. For a 32-bit IP address, enter `0`.
	Mask *uint64 `json:"Mask,omitnil" name:"Mask"`
}

type KeyValue struct {
	// IP
	Key *string `json:"Key,omitnil" name:"Key"`

	// Status of the IP. Values: `1` (blocked); `2` (normal); `3` (being attacked)
	Value *string `json:"Value,omitnil" name:"Value"`
}

type L4RuleSource struct {
	// IP or domain name for forwarding.
	Source *string `json:"Source,omitnil" name:"Source"`

	// Weight. Value range: [0,100].
	Weight *uint64 `json:"Weight,omitnil" name:"Weight"`

	// 8000
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Port *uint64 `json:"Port,omitnil" name:"Port"`

	// Secondary origin server. `1`: secondary origin server; `0`: general origin server.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Backup *uint64 `json:"Backup,omitnil" name:"Backup"`
}

type L7RuleEntry struct {
	// Session persistence duration, in seconds.
	KeepTime *uint64 `json:"KeepTime,omitnil" name:"KeepTime"`

	// Forwarding domain name.
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Forwarding protocol. Valid values: `http` and `https`.
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// Forwarding method. Valid values: `1` (by domain name); `2` (by IP).
	SourceType *uint64 `json:"SourceType,omitnil" name:"SourceType"`

	// Load balancing method. Valid value: `1` (weighed polling).
	LbType *uint64 `json:"LbType,omitnil" name:"LbType"`

	// List of origins
	SourceList []*L4RuleSource `json:"SourceList,omitnil" name:"SourceList"`

	// Whether session persistence is enabled. Valid values: `0` (disabled) and `1` (enabled).
	KeepEnable *uint64 `json:"KeepEnable,omitnil" name:"KeepEnable"`

	// Rule status. Valid values: `0` (the rule was successfully configured), `1` (configuring the rule), `2` (rule configuration failed), `3` (deleting the rule), `5` (failed to delete rule), `6` (rule awaiting configuration), `7` (rule awaiting deletion), and `8` (rule awaiting certificate configuration).
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// Rule ID. This field is not required for adding a rule, but is required for modifying or deleting a rule.
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// CC protection threshold based on HTTPS.
	CCThreshold *uint64 `json:"CCThreshold,omitnil" name:"CCThreshold"`

	// [Disused] When the certificate is an external certificate, the certificate key should be provided here. 
	PrivateKey *string `json:"PrivateKey,omitnil" name:"PrivateKey"`

	// CC protection status based on HTTPS. Valid values: `0` (disabled) and `1` (enabled).
	CCEnable *uint64 `json:"CCEnable,omitnil" name:"CCEnable"`

	// Whether to enable **Forward HTTPS requests via HTTP**. Valid values: `0` (disable) and `1` (enable). It defaults to `0`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	HttpsToHttpEnable *uint64 `json:"HttpsToHttpEnable,omitnil" name:"HttpsToHttpEnable"`

	// Certificate source. When the forwarding protocol is HTTPS, this field must be set to `2` (Tencent Cloud managed certificate), and for HTTP protocol, it can be set to `0`.
	CertType *uint64 `json:"CertType,omitnil" name:"CertType"`

	// [Disused] When the certificate is an external certificate, the certificate content should be provided here. 
	Cert *string `json:"Cert,omitnil" name:"Cert"`

	// CC protection level based on HTTPS.
	CCLevel *string `json:"CCLevel,omitnil" name:"CCLevel"`

	// Rule description.
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// CC protection status. Valid values: `0` (disabled) and `1` (enabled).
	CCStatus *uint64 `json:"CCStatus,omitnil" name:"CCStatus"`

	// Access port number.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VirtualPort *uint64 `json:"VirtualPort,omitnil" name:"VirtualPort"`

	// When the certificate is managed by Tencent Cloud, this field must be set to the ID of the managed certificate.
	SSLId *string `json:"SSLId,omitnil" name:"SSLId"`

	// ID of the rule
	Id *string `json:"Id,omitnil" name:"Id"`

	// Intelligent CC protection status. Valid values: `0` (disabled) and `1` (enabled).
	CCAIEnable *uint64 `json:"CCAIEnable,omitnil" name:"CCAIEnable"`
}

type L7RuleHealth struct {
	// Configuration status. Values: `0` (normal), `1` (configuration in progress) and `2` (configuration failed).
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// Switch. Values: `1`: Enable; `0`: Disable.
	Enable *uint64 `json:"Enable,omitnil" name:"Enable"`

	// ID of the rule
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// HTTP request path. The default value is /.
	Url *string `json:"Url,omitnil" name:"Url"`

	// Health check interval. Unit: second.
	Interval *uint64 `json:"Interval,omitnil" name:"Interval"`

	// Healthy threshold, which specifies the number of consecutive successful health checks.
	AliveNum *uint64 `json:"AliveNum,omitnil" name:"AliveNum"`

	// Unhealthy threshold, which specifies the number of consecutive failed health checks.
	KickNum *uint64 `json:"KickNum,omitnil" name:"KickNum"`

	// HTTP request method. Values: `HEAD` and `GET`.
	Method *string `json:"Method,omitnil" name:"Method"`

	// Status code that signifies a normal state. Values: `1` (1xx), `2` (2xx), `4` (3xx), `8` (4xx), and `16` (5xx).
	StatusCode *uint64 `json:"StatusCode,omitnil" name:"StatusCode"`

	// Whether to deploy both HTTP and HTTPS health check rules
	ProtocolFlag *uint64 `json:"ProtocolFlag,omitnil" name:"ProtocolFlag"`

	// Enables passive detection. Values: `1` (enable) and `0` (disable).
	PassiveEnable *uint64 `json:"PassiveEnable,omitnil" name:"PassiveEnable"`

	// Blocking period in the passive detection configuration
	BlockInter *uint64 `json:"BlockInter,omitnil" name:"BlockInter"`

	// Time interval between passive detections
	FailedCountInter *uint64 `json:"FailedCountInter,omitnil" name:"FailedCountInter"`

	// Unhealthy threshold in the passive detection configuration
	FailedThreshold *uint64 `json:"FailedThreshold,omitnil" name:"FailedThreshold"`

	// Status code that signals that the passive detection considers the status normal. Values: `1` (1xx), `2` (2xx), `4` (3xx), `8` (4xx), and `16` (5xx).
	PassiveStatusCode *uint64 `json:"PassiveStatusCode,omitnil" name:"PassiveStatusCode"`

	// Configuration status of the passive health check. Values: `0` (Normal), `1` (configuration in progress) and `2` (configuration failed).
	PassiveStatus *uint64 `json:"PassiveStatus,omitnil" name:"PassiveStatus"`
}

type Layer4Rule struct {
	// Real server port. Value range: 1–65535.
	BackendPort *uint64 `json:"BackendPort,omitnil" name:"BackendPort"`

	// Forwarding port. Value range: 1–65535.
	FrontendPort *uint64 `json:"FrontendPort,omitnil" name:"FrontendPort"`

	// Forwarding rule. Valid values:
	// TCP
	// UDP
	// ]
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// List of real servers
	RealServers []*SourceServer `json:"RealServers,omitnil" name:"RealServers"`

	// Information of the Anti-DDoS instance
	InstanceDetails []*InstanceRelation `json:"InstanceDetails,omitnil" name:"InstanceDetails"`

	// Information of the Anti-DDoS instance configured
	InstanceDetailRule []*RuleInstanceRelation `json:"InstanceDetailRule,omitnil" name:"InstanceDetailRule"`
}

type Layer7Rule struct {
	// Domain name
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// List of forwarding types
	ProxyTypeList []*ProxyTypeInfo `json:"ProxyTypeList,omitnil" name:"ProxyTypeList"`

	// List of real servers
	RealServers []*SourceServer `json:"RealServers,omitnil" name:"RealServers"`

	// Information of the Anti-DDoS instance
	InstanceDetails []*InstanceRelation `json:"InstanceDetails,omitnil" name:"InstanceDetails"`

	// Information of the Anti-DDoS instance configured
	InstanceDetailRule []*RuleInstanceRelation `json:"InstanceDetailRule,omitnil" name:"InstanceDetailRule"`

	// Protocol
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// Port number
	Vport *int64 `json:"Vport,omitnil" name:"Vport"`
}

type ListenerCcThreholdConfig struct {
	// Domain name
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Protocol. Value: `https`.
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// Status. Valid values: `0` (disabled), `1` (enabled).
	CCEnable *int64 `json:"CCEnable,omitnil" name:"CCEnable"`

	// CC protection threshold
	CCThreshold *int64 `json:"CCThreshold,omitnil" name:"CCThreshold"`
}

// Predefined struct for user
type ModifyCCPrecisionPolicyRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Policy ID
	PolicyId *string `json:"PolicyId,omitnil" name:"PolicyId"`

	// Specifies the action. `alg`: Verify the access request via CAPTCHA; `drop`: Drop the access request.
	PolicyAction *string `json:"PolicyAction,omitnil" name:"PolicyAction"`

	// Policy records
	PolicyList []*CCPrecisionPlyRecord `json:"PolicyList,omitnil" name:"PolicyList"`
}

type ModifyCCPrecisionPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Policy ID
	PolicyId *string `json:"PolicyId,omitnil" name:"PolicyId"`

	// Specifies the action. `alg`: Verify the access request via CAPTCHA; `drop`: Drop the access request.
	PolicyAction *string `json:"PolicyAction,omitnil" name:"PolicyAction"`

	// Policy records
	PolicyList []*CCPrecisionPlyRecord `json:"PolicyList,omitnil" name:"PolicyList"`
}

func (r *ModifyCCPrecisionPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCPrecisionPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PolicyId")
	delete(f, "PolicyAction")
	delete(f, "PolicyList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCCPrecisionPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCCPrecisionPolicyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCCPrecisionPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCCPrecisionPolicyResponseParams `json:"Response"`
}

func (r *ModifyCCPrecisionPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCCPrecisionPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCcBlackWhiteIpListRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// List of IPs
	IpList []*IpSegment `json:"IpList,omitnil" name:"IpList"`

	// IP type. Valid values: `black` (blocklisted IP), `white`(allowlisted IP).
	Type *string `json:"Type,omitnil" name:"Type"`

	// Policy ID
	PolicyId *string `json:"PolicyId,omitnil" name:"PolicyId"`
}

type ModifyCcBlackWhiteIpListRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// List of IPs
	IpList []*IpSegment `json:"IpList,omitnil" name:"IpList"`

	// IP type. Valid values: `black` (blocklisted IP), `white`(allowlisted IP).
	Type *string `json:"Type,omitnil" name:"Type"`

	// Policy ID
	PolicyId *string `json:"PolicyId,omitnil" name:"PolicyId"`
}

func (r *ModifyCcBlackWhiteIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCcBlackWhiteIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IpList")
	delete(f, "Type")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCcBlackWhiteIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCcBlackWhiteIpListResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCcBlackWhiteIpListResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCcBlackWhiteIpListResponseParams `json:"Response"`
}

func (r *ModifyCcBlackWhiteIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCcBlackWhiteIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSGeoIPBlockConfigRequestParams struct {
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Region blocking configuration. The configuration ID cannot be empty when you set this parameter.
	DDoSGeoIPBlockConfig *DDoSGeoIPBlockConfig `json:"DDoSGeoIPBlockConfig,omitnil" name:"DDoSGeoIPBlockConfig"`
}

type ModifyDDoSGeoIPBlockConfigRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Region blocking configuration. The configuration ID cannot be empty when you set this parameter.
	DDoSGeoIPBlockConfig *DDoSGeoIPBlockConfig `json:"DDoSGeoIPBlockConfig,omitnil" name:"DDoSGeoIPBlockConfig"`
}

func (r *ModifyDDoSGeoIPBlockConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSGeoIPBlockConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DDoSGeoIPBlockConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDDoSGeoIPBlockConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSGeoIPBlockConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDDoSGeoIPBlockConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDDoSGeoIPBlockConfigResponseParams `json:"Response"`
}

func (r *ModifyDDoSGeoIPBlockConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSGeoIPBlockConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSSpeedLimitConfigRequestParams struct {
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Access rate limit configuration. The configuration ID cannot be empty when you set this parameter.
	DDoSSpeedLimitConfig *DDoSSpeedLimitConfig `json:"DDoSSpeedLimitConfig,omitnil" name:"DDoSSpeedLimitConfig"`
}

type ModifyDDoSSpeedLimitConfigRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Access rate limit configuration. The configuration ID cannot be empty when you set this parameter.
	DDoSSpeedLimitConfig *DDoSSpeedLimitConfig `json:"DDoSSpeedLimitConfig,omitnil" name:"DDoSSpeedLimitConfig"`
}

func (r *ModifyDDoSSpeedLimitConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSSpeedLimitConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DDoSSpeedLimitConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDDoSSpeedLimitConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDDoSSpeedLimitConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDDoSSpeedLimitConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDDoSSpeedLimitConfigResponseParams `json:"Response"`
}

func (r *ModifyDDoSSpeedLimitConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDDoSSpeedLimitConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainUsrNameRequestParams struct {
	// User CNAME
	DomainName *string `json:"DomainName,omitnil" name:"DomainName"`

	// Domain name
	DomainUserName *string `json:"DomainUserName,omitnil" name:"DomainUserName"`
}

type ModifyDomainUsrNameRequest struct {
	*tchttp.BaseRequest
	
	// User CNAME
	DomainName *string `json:"DomainName,omitnil" name:"DomainName"`

	// Domain name
	DomainUserName *string `json:"DomainUserName,omitnil" name:"DomainUserName"`
}

func (r *ModifyDomainUsrNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainUsrNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "DomainUserName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDomainUsrNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainUsrNameResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDomainUsrNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDomainUsrNameResponseParams `json:"Response"`
}

func (r *ModifyDomainUsrNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainUsrNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNewDomainRulesRequestParams struct {
	// Anti-DDoS service type (`bgpip`: Anti-DDoS Advanced).
	Business *string `json:"Business,omitnil" name:"Business"`

	// Anti-DDoS instance ID.
	Id *string `json:"Id,omitnil" name:"Id"`

	// Domain name forwarding rule.
	Rule *NewL7RuleEntry `json:"Rule,omitnil" name:"Rule"`
}

type ModifyNewDomainRulesRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS service type (`bgpip`: Anti-DDoS Advanced).
	Business *string `json:"Business,omitnil" name:"Business"`

	// Anti-DDoS instance ID.
	Id *string `json:"Id,omitnil" name:"Id"`

	// Domain name forwarding rule.
	Rule *NewL7RuleEntry `json:"Rule,omitnil" name:"Rule"`
}

func (r *ModifyNewDomainRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNewDomainRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Business")
	delete(f, "Id")
	delete(f, "Rule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNewDomainRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNewDomainRulesResponseParams struct {
	// Success code.
	Success *SuccessCode `json:"Success,omitnil" name:"Success"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyNewDomainRulesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNewDomainRulesResponseParams `json:"Response"`
}

func (r *ModifyNewDomainRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNewDomainRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPacketFilterConfigRequestParams struct {
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Feature filtering configuration
	PacketFilterConfig *PacketFilterConfig `json:"PacketFilterConfig,omitnil" name:"PacketFilterConfig"`
}

type ModifyPacketFilterConfigRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Feature filtering configuration
	PacketFilterConfig *PacketFilterConfig `json:"PacketFilterConfig,omitnil" name:"PacketFilterConfig"`
}

func (r *ModifyPacketFilterConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPacketFilterConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PacketFilterConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPacketFilterConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPacketFilterConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyPacketFilterConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPacketFilterConfigResponseParams `json:"Response"`
}

func (r *ModifyPacketFilterConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPacketFilterConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NewL7RuleEntry struct {
	// Forwarding protocol. Valid values: `http` and `https`.
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// Forwarding domain name.
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Load balancing method. Valid value: `1` (weighed polling).
	LbType *uint64 `json:"LbType,omitnil" name:"LbType"`

	// Whether session persistence is enabled. Valid values: `0` (disabled) and `1` (enabled).
	KeepEnable *uint64 `json:"KeepEnable,omitnil" name:"KeepEnable"`

	// Session persistence duration, in seconds.
	KeepTime *uint64 `json:"KeepTime,omitnil" name:"KeepTime"`

	// Forwarding method. Valid values: `1` (by domain name); `2` (by IP).
	SourceType *uint64 `json:"SourceType,omitnil" name:"SourceType"`

	// List of origins
	SourceList []*L4RuleSource `json:"SourceList,omitnil" name:"SourceList"`

	// Region code.
	Region *uint64 `json:"Region,omitnil" name:"Region"`

	// Resource ID.
	Id *string `json:"Id,omitnil" name:"Id"`

	// Anti-DDoS instance IP address.
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Rule ID. This field is not required for adding a rule, but is required for modifying or deleting a rule.
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// Rule description.
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// Certificate source. When the forwarding protocol is HTTPS, this field must be set to `2` (Tencent Cloud managed certificate), and for HTTP protocol, it can be set to `0`.
	CertType *uint64 `json:"CertType,omitnil" name:"CertType"`

	// When the certificate is managed by Tencent Cloud, this field must be set to the ID of the managed certificate.
	SSLId *string `json:"SSLId,omitnil" name:"SSLId"`

	// [Disused] When the certificate is an external certificate, the certificate content should be provided here. 
	Cert *string `json:"Cert,omitnil" name:"Cert"`

	// [Disused] When the certificate is an external certificate, the certificate key should be provided here. 
	PrivateKey *string `json:"PrivateKey,omitnil" name:"PrivateKey"`

	// Rule status. Valid values: `0` (the rule was successfully configured), `1` (configuring the rule), `2` (rule configuration failed), `3` (deleting the rule), `5` (failed to delete rule), `6` (rule awaiting configuration), `7` (rule awaiting deletion), and `8` (rule awaiting certificate configuration).
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// CC protection status. Valid values: `0` (disabled) and `1` (enabled).
	CCStatus *uint64 `json:"CCStatus,omitnil" name:"CCStatus"`

	// CC protection status based on HTTPS. Valid values: `0` (disabled) and `1` (enabled).
	CCEnable *uint64 `json:"CCEnable,omitnil" name:"CCEnable"`

	// CC protection threshold based on HTTPS.
	CCThreshold *uint64 `json:"CCThreshold,omitnil" name:"CCThreshold"`

	// CC protection level based on HTTPS.
	CCLevel *string `json:"CCLevel,omitnil" name:"CCLevel"`

	// Modification time.
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// Whether to enable **Forward HTTPS requests via HTTP**. Valid values: `0` (disabled) and `1` (enabled). It defaults to `0`.
	HttpsToHttpEnable *uint64 `json:"HttpsToHttpEnable,omitnil" name:"HttpsToHttpEnable"`

	// Access port number.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	VirtualPort *uint64 `json:"VirtualPort,omitnil" name:"VirtualPort"`

	// Specifies whether to forcibly redirect HTTP to HTTPS. `1`: Enable. `0`: Disable.
	RewriteHttps *uint64 `json:"RewriteHttps,omitnil" name:"RewriteHttps"`

	// Returns an error code when the rule configuration fails (only valid when `Status=2`). `1001`: The certificate does not exist. `1002`: Failed to obtain the certificate. `1003`: Failed to upload the certificate. `1004`: The certificate has expired.
	ErrCode *uint64 `json:"ErrCode,omitnil" name:"ErrCode"`

	// Version
	// Note: This field may return null, indicating that no valid values can be obtained.
	Version *uint64 `json:"Version,omitnil" name:"Version"`
}

type OverviewDDoSEvent struct {
	// Event ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// IP
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// Start time
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Attack type
	AttackType *string `json:"AttackType,omitnil" name:"AttackType"`

	// Attack status. `0`: The attack is ongoing; `1`: The attack ends.
	AttackStatus *uint64 `json:"AttackStatus,omitnil" name:"AttackStatus"`

	// Attack traffic, in Mbps
	Mbps *uint64 `json:"Mbps,omitnil" name:"Mbps"`

	// Attack packets, in PPS
	Pps *uint64 `json:"Pps,omitnil" name:"Pps"`

	// Anti-DDoS service type. `bgp-multip`: Anti-DDoS Pro; `bgpip`: Anti-DDoS Advanced; `basic`: Anti-DDoS Basic.
	Business *string `json:"Business,omitnil" name:"Business"`

	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Anti-DDoS instance name
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`
}

type PackInfo struct {
	// Package type. Valid values:
	// `staticpack`: non-BGP package
	// `insurance`: guarantee package
	// ]
	PackType *string `json:"PackType,omitnil" name:"PackType"`

	// Package ID
	PackId *string `json:"PackId,omitnil" name:"PackId"`
}

type PacketFilterConfig struct {
	// Protocol. Valid values: `tcp`, `udp`, `icmp`, `all`.
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// Start source port. Value range: 0–65535.
	SportStart *int64 `json:"SportStart,omitnil" name:"SportStart"`

	// End source port. Value range: 0–65535. The value also should be greater than or equal to that of the start source port.
	SportEnd *int64 `json:"SportEnd,omitnil" name:"SportEnd"`

	// Start destination port
	DportStart *int64 `json:"DportStart,omitnil" name:"DportStart"`

	// End destination port. Value range: 1–65535. The value also should be greater than or equal to that of the start source port.
	DportEnd *int64 `json:"DportEnd,omitnil" name:"DportEnd"`

	// Minimum message length. Value range: 1–1500.
	PktlenMin *int64 `json:"PktlenMin,omitnil" name:"PktlenMin"`

	// Maximum message length. Value range: 1–1500. The value also should be greater than or equal to that of the minimum message length.
	PktlenMax *int64 `json:"PktlenMax,omitnil" name:"PktlenMax"`

	// Action. Valid values:
	// `drop`: discards the request.
	// `transmit`: allows the request.
	// `drop_black`: discards the request and adds the IP to blocklist.
	// `drop_rst`: blocks the request.
	// `drop_black_rst`: blocks the request and adds the IP to blocklist.
	// `forward`: continues protection.
	// ]
	Action *string `json:"Action,omitnil" name:"Action"`

	// Detection location:
	// `begin_l3`: IP header
	// `begin_l4`: TCP/UDP header
	// `begin_l5`: T load
	// `no_match`: no match
	// ]
	MatchBegin *string `json:"MatchBegin,omitnil" name:"MatchBegin"`

	// Detection type:
	// `sunday`: keyword
	// `pcre`: regular expression
	// ]
	MatchType *string `json:"MatchType,omitnil" name:"MatchType"`

	// Detection value. Should be in key string or regular expression. 
	// When the `MatchType` is `sunday`, enter a string or a string in hexadecimal byte code representation. For example, a string "123" corresponds to the hexadecimal byte code "\x313233".
	// When the `MatchType` is `pcre`, enter a regular expression.
	// ]
	Str *string `json:"Str,omitnil" name:"Str"`

	// Detection depth starting from the detection position. Value range: [0, 1500].
	Depth *int64 `json:"Depth,omitnil" name:"Depth"`

	// Offset starting from the detection position. Value range: [0, Depth].
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Whether the detection value is included:
	// `0`: included
	// `1`: excluded
	// ]
	IsNot *int64 `json:"IsNot,omitnil" name:"IsNot"`

	// Relationship between the first and second detection conditions:
	// `and`: under both the first and second detection conditions
	// `none`: under only the first detection condition
	// ]
	MatchLogic *string `json:"MatchLogic,omitnil" name:"MatchLogic"`

	// The second detection position:
	// `begin_l5`: load
	// `no_match`: no match
	// ]
	MatchBegin2 *string `json:"MatchBegin2,omitnil" name:"MatchBegin2"`

	// The second detection type:
	// `sunday`: keyword
	// `pcre`: regular expression
	// ]
	MatchType2 *string `json:"MatchType2,omitnil" name:"MatchType2"`

	// The second detection value. Should be in key string or regular expression.
	// When the `MatchType` is `sunday`, enter a string or a string in hexadecimal byte code representation. For example, a string "123" corresponds to the hexadecimal byte code "\x313233".
	// When the `MatchType` is `pcre`, enter a regular expression.
	// ]
	Str2 *string `json:"Str2,omitnil" name:"Str2"`

	// Detection depth starting from the second detection position. Value range: [0, 1500].
	Depth2 *int64 `json:"Depth2,omitnil" name:"Depth2"`

	// Offset starting from the second detection position. Value range: [0, Depth2].
	Offset2 *int64 `json:"Offset2,omitnil" name:"Offset2"`

	// Whether the second detection value is included:
	// `0`: included
	// `1`: excluded
	// ]
	IsNot2 *int64 `json:"IsNot2,omitnil" name:"IsNot2"`

	// A rule ID is generated after a feature filtering configuration is added successfully. Leave this field empty when adding a new feature filtering configuration.
	Id *string `json:"Id,omitnil" name:"Id"`

	// Byte threshold of the packet. Packets larger than the specified size are not returned. It must be an integer larger than 1.
	PktLenGT *int64 `json:"PktLenGT,omitnil" name:"PktLenGT"`
}

type PacketFilterRelation struct {
	// Feature filtering configuration
	PacketFilterConfig *PacketFilterConfig `json:"PacketFilterConfig,omitnil" name:"PacketFilterConfig"`

	// Anti-DDoS instance configured
	InstanceDetailList []*InstanceRelation `json:"InstanceDetailList,omitnil" name:"InstanceDetailList"`

	// Modification time
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`
}

type PortSegment struct {
	// Start port. Value range: 1–65535.
	BeginPort *uint64 `json:"BeginPort,omitnil" name:"BeginPort"`

	// End port. The value should be in the range 1–65535 and cannot be less than that of the start port.
	EndPort *uint64 `json:"EndPort,omitnil" name:"EndPort"`
}

type ProtectThresholdRelation struct {
	// DDoS protection level:
	// `low`: loose protection
	// `middle`: medium protection
	// `high`: strict protection
	// ]
	DDoSLevel *string `json:"DDoSLevel,omitnil" name:"DDoSLevel"`

	// DDoS cleansing threshold (in Mbps)
	DDoSThreshold *uint64 `json:"DDoSThreshold,omitnil" name:"DDoSThreshold"`

	// DDoS AI protection switch:
	// `on`: enabled
	// `off`: disabled
	// ]
	DDoSAI *string `json:"DDoSAI,omitnil" name:"DDoSAI"`

	// CC cleansing switch
	// `0`: disabled
	// `1`: enabled
	// ]
	CCEnable *uint64 `json:"CCEnable,omitnil" name:"CCEnable"`

	// CC cleansing threshold (in QPS)
	CCThreshold *uint64 `json:"CCThreshold,omitnil" name:"CCThreshold"`

	// Anti-DDoS instance configured
	InstanceDetailList []*InstanceRelation `json:"InstanceDetailList,omitnil" name:"InstanceDetailList"`

	// Domain name and protocol protection thresholds
	ListenerCcThresholdList []*ListenerCcThreholdConfig `json:"ListenerCcThresholdList,omitnil" name:"ListenerCcThresholdList"`

	// SYN traffic threshold
	// Note: This field may return null, indicating that no valid values can be obtained.
	SynFloodThreshold *uint64 `json:"SynFloodThreshold,omitnil" name:"SynFloodThreshold"`

	// SYN packet threshold
	// Note: This field may return null, indicating that no valid values can be obtained.
	SynFloodPktThreshold *uint64 `json:"SynFloodPktThreshold,omitnil" name:"SynFloodPktThreshold"`

	// UDP traffic threshold
	// Note: This field may return null, indicating that no valid values can be obtained.
	UdpFloodThreshold *uint64 `json:"UdpFloodThreshold,omitnil" name:"UdpFloodThreshold"`

	// UDP packet threshold
	// Note: This field may return null, indicating that no valid values can be obtained.
	UdpFloodPktThreshold *uint64 `json:"UdpFloodPktThreshold,omitnil" name:"UdpFloodPktThreshold"`

	// ACK traffic threshold
	// Note: This field may return null, indicating that no valid values can be obtained.
	AckFloodThreshold *uint64 `json:"AckFloodThreshold,omitnil" name:"AckFloodThreshold"`

	// ACK packet threshold
	// Note: This field may return null, indicating that no valid values can be obtained.
	AckFloodPktThreshold *uint64 `json:"AckFloodPktThreshold,omitnil" name:"AckFloodPktThreshold"`

	// SYNACK traffic threshold
	// Note: This field may return null, indicating that no valid values can be obtained.
	SynAckFloodThreshold *uint64 `json:"SynAckFloodThreshold,omitnil" name:"SynAckFloodThreshold"`

	// SYNACK packet threshold
	// Note: This field may return null, indicating that no valid values can be obtained.
	SynAckFloodPktThreshold *uint64 `json:"SynAckFloodPktThreshold,omitnil" name:"SynAckFloodPktThreshold"`

	// RST traffic threshold
	// Note: This field may return null, indicating that no valid values can be obtained.
	RstFloodThreshold *uint64 `json:"RstFloodThreshold,omitnil" name:"RstFloodThreshold"`

	// RST packet threshold
	// Note: This field may return null, indicating that no valid values can be obtained.
	RstFloodPktThreshold *uint64 `json:"RstFloodPktThreshold,omitnil" name:"RstFloodPktThreshold"`
}

type ProtocolBlockConfig struct {
	// TCP blocking. Valid values: `0` (disabled), `1`(enabled).
	DropTcp *int64 `json:"DropTcp,omitnil" name:"DropTcp"`

	// UDP blocking. Valid values: `0` (disabled), `1`(enabled).
	DropUdp *int64 `json:"DropUdp,omitnil" name:"DropUdp"`

	// ICMP blocking. Valid values: `0` (disabled), `1`(enabled).
	DropIcmp *int64 `json:"DropIcmp,omitnil" name:"DropIcmp"`

	// Other protocol blocking. Valid values: `0` (disabled), `1`(enabled).
	DropOther *int64 `json:"DropOther,omitnil" name:"DropOther"`

	// Null connection protection. Valid values: `0` (disabled), `1` (enabled).
	CheckExceptNullConnect *int64 `json:"CheckExceptNullConnect,omitnil" name:"CheckExceptNullConnect"`

	// PoD protection. Values: `0` (disable), `1` (enable).
	PingOfDeath *int64 `json:"PingOfDeath,omitnil" name:"PingOfDeath"`

	// Teardrop protection. Values: `0` (disable), `1` (enable).
	TearDrop *int64 `json:"TearDrop,omitnil" name:"TearDrop"`
}

type ProtocolBlockRelation struct {
	// Protocol blocking configuration
	ProtocolBlockConfig *ProtocolBlockConfig `json:"ProtocolBlockConfig,omitnil" name:"ProtocolBlockConfig"`

	// Anti-DDoS instance configured
	InstanceDetailList []*InstanceRelation `json:"InstanceDetailList,omitnil" name:"InstanceDetailList"`
}

type ProtocolPort struct {
	// Protocol. Valid values: `tcp`, `udp`
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// Port
	Port *uint64 `json:"Port,omitnil" name:"Port"`
}

type ProxyTypeInfo struct {
	// List of forwarding listening ports. Value range: 1–65535.
	ProxyPorts []*int64 `json:"ProxyPorts,omitnil" name:"ProxyPorts"`

	// Forwarding protocol:
	// `http`: HTTP protocol
	// `https`: HTTPS protocol
	// ]
	ProxyType *string `json:"ProxyType,omitnil" name:"ProxyType"`
}

type RegionInfo struct {
	// Region name, such as `ap-guangzhou`
	Region *string `json:"Region,omitnil" name:"Region"`
}

type RuleInstanceRelation struct {
	// Instance IP
	EipList []*string `json:"EipList,omitnil" name:"EipList"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance CNAME
	Cname *string `json:"Cname,omitnil" name:"Cname"`
}

type SchedulingDomainInfo struct {
	// Scheduling domain name
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// List of line IPs
	LineIPList []*IPLineInfo `json:"LineIPList,omitnil" name:"LineIPList"`

	// Scheduling mode. Valid value: `priority` (priority scheduling).
	Method *string `json:"Method,omitnil" name:"Method"`

	// TTL value recorded from the scheduling domain name resolution
	TTL *uint64 `json:"TTL,omitnil" name:"TTL"`

	// Running status. Valid values:
	// `0`: not running
	// `1`: running
	// `2`: abnormal
	// ]
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// Creation time
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// Last modification time
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// Domain name
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	UsrDomainName *string `json:"UsrDomainName,omitnil" name:"UsrDomainName"`
}

type SourceServer struct {
	// Types of the real server address, such as IP or domain name.
	RealServer *string `json:"RealServer,omitnil" name:"RealServer"`

	// Types of the real server address:
	// `1`: domain name
	// `2`: IP
	// ]
	RsType *int64 `json:"RsType,omitnil" name:"RsType"`

	// Forward weight of the real server. Value range: 1–100.
	Weight *int64 `json:"Weight,omitnil" name:"Weight"`

	// Port number. Value range: 0-65535.
	Port *int64 `json:"Port,omitnil" name:"Port"`
}

type SpeedValue struct {
	// Rate limit value types:
	// `1`: packets per second (pps)
	// `2`: bits per second (bps)
	// ]
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// Value
	Value *uint64 `json:"Value,omitnil" name:"Value"`
}

type StaticPackRelation struct {
	// Base protection bandwidth
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ProtectBandwidth *uint64 `json:"ProtectBandwidth,omitnil" name:"ProtectBandwidth"`

	// Application bandwidth
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	NormalBandwidth *uint64 `json:"NormalBandwidth,omitnil" name:"NormalBandwidth"`

	// Forwarding rules
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	ForwardRulesLimit *uint64 `json:"ForwardRulesLimit,omitnil" name:"ForwardRulesLimit"`

	// Auto-renewal flag
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// Expiration time
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CurDeadline *string `json:"CurDeadline,omitnil" name:"CurDeadline"`
}

type SuccessCode struct {
	// Description
	Message *string `json:"Message,omitnil" name:"Message"`

	// Success/Error code
	Code *string `json:"Code,omitnil" name:"Code"`
}

// Predefined struct for user
type SwitchWaterPrintConfigRequestParams struct {
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Watermark status. `1`: enabled; `0`: disabled.
	OpenStatus *int64 `json:"OpenStatus,omitnil" name:"OpenStatus"`

	// Whether to enable proxy. Values: `1` (Enable proxy and ignore IP+port verification), `0` (Do not enable proxy and IP+port verification is required)
	CloudSdkProxy *int64 `json:"CloudSdkProxy,omitnil" name:"CloudSdkProxy"`
}

type SwitchWaterPrintConfigRequest struct {
	*tchttp.BaseRequest
	
	// Anti-DDoS instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Watermark status. `1`: enabled; `0`: disabled.
	OpenStatus *int64 `json:"OpenStatus,omitnil" name:"OpenStatus"`

	// Whether to enable proxy. Values: `1` (Enable proxy and ignore IP+port verification), `0` (Do not enable proxy and IP+port verification is required)
	CloudSdkProxy *int64 `json:"CloudSdkProxy,omitnil" name:"CloudSdkProxy"`
}

func (r *SwitchWaterPrintConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchWaterPrintConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "OpenStatus")
	delete(f, "CloudSdkProxy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchWaterPrintConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchWaterPrintConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SwitchWaterPrintConfigResponse struct {
	*tchttp.BaseResponse
	Response *SwitchWaterPrintConfigResponseParams `json:"Response"`
}

func (r *SwitchWaterPrintConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchWaterPrintConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagFilter struct {
	// Tag key
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// Tag value
	TagValue []*string `json:"TagValue,omitnil" name:"TagValue"`
}

type TagInfo struct {
	// Tag key
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}

type WaterPrintConfig struct {
	// Watermark offset. Value range: [0, 100).
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Start status. Valid values:
	// `0`: manual start
	// `1`: instant start
	// ]
	OpenStatus *int64 `json:"OpenStatus,omitnil" name:"OpenStatus"`

	// List of forwarding listeners configured
	Listeners []*ForwardListener `json:"Listeners,omitnil" name:"Listeners"`

	// A list of watermark keys is generated after a watermark is added successfully. Each watermark can have up to 2 keys. When there is only one valid key, it cannot be deleted.
	Keys []*WaterPrintKey `json:"Keys,omitnil" name:"Keys"`

	// Watermark checking mode, which can be:
	// `checkall`: normal mode
	// `shortfpcheckall`: compact mode
	// ]
	Verify *string `json:"Verify,omitnil" name:"Verify"`

	// Whether to enable proxy. Values: `1` (Enable proxy and ignore IP+port verification), `0` (Do not enable proxy and IP+port verification is required)
	// Note: This field may return null, indicating that no valid values can be obtained.
	CloudSdkProxy *int64 `json:"CloudSdkProxy,omitnil" name:"CloudSdkProxy"`
}

type WaterPrintKey struct {
	// Key version
	KeyVersion *string `json:"KeyVersion,omitnil" name:"KeyVersion"`

	// Key content
	KeyContent *string `json:"KeyContent,omitnil" name:"KeyContent"`

	// Key ID
	KeyId *string `json:"KeyId,omitnil" name:"KeyId"`

	// Key status. Valid value: `1` (enabled).
	KeyOpenStatus *int64 `json:"KeyOpenStatus,omitnil" name:"KeyOpenStatus"`

	// Key creation time
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`
}

type WaterPrintRelation struct {
	// Watermark configuration
	WaterPrintConfig *WaterPrintConfig `json:"WaterPrintConfig,omitnil" name:"WaterPrintConfig"`

	// Anti-DDoS instance configured
	InstanceDetailList []*InstanceRelation `json:"InstanceDetailList,omitnil" name:"InstanceDetailList"`
}