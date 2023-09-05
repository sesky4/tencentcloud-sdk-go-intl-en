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

package v20170312

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type Account struct {
	// Account name
	User *string `json:"User,omitnil" name:"User"`

	// Host address
	Host *string `json:"Host,omitnil" name:"Host"`
}

// Predefined struct for user
type AssociateSecurityGroupsRequestParams struct {
	// Database engine name. Valid value: `mariadb`.
	Product *string `json:"Product,omitnil" name:"Product"`

	// ID of the security group to be associated in the format of sg-efil73jd.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// ID(s) of the instance(s) to be associated in the format of tdsql-lesecurk. You can specify multiple instances.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Database engine name. Valid value: `mariadb`.
	Product *string `json:"Product,omitnil" name:"Product"`

	// ID of the security group to be associated in the format of sg-efil73jd.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// ID(s) of the instance(s) to be associated in the format of tdsql-lesecurk. You can specify multiple instances.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

func (r *AssociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "SecurityGroupId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateSecurityGroupsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AssociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *AssociateSecurityGroupsResponseParams `json:"Response"`
}

func (r *AssociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelDcnJobRequestParams struct {
	// Disaster recovery instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type CancelDcnJobRequest struct {
	*tchttp.BaseRequest
	
	// Disaster recovery instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *CancelDcnJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelDcnJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelDcnJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelDcnJobResponseParams struct {
	// Task ID
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CancelDcnJobResponse struct {
	*tchttp.BaseResponse
	Response *CancelDcnJobResponseParams `json:"Response"`
}

func (r *CancelDcnJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelDcnJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloneAccountRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Source user account name
	SrcUser *string `json:"SrcUser,omitnil" name:"SrcUser"`

	// Source user host
	SrcHost *string `json:"SrcHost,omitnil" name:"SrcHost"`

	// Target user account name
	DstUser *string `json:"DstUser,omitnil" name:"DstUser"`

	// Target user host
	DstHost *string `json:"DstHost,omitnil" name:"DstHost"`

	// Target account description
	DstDesc *string `json:"DstDesc,omitnil" name:"DstDesc"`
}

type CloneAccountRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Source user account name
	SrcUser *string `json:"SrcUser,omitnil" name:"SrcUser"`

	// Source user host
	SrcHost *string `json:"SrcHost,omitnil" name:"SrcHost"`

	// Target user account name
	DstUser *string `json:"DstUser,omitnil" name:"DstUser"`

	// Target user host
	DstHost *string `json:"DstHost,omitnil" name:"DstHost"`

	// Target account description
	DstDesc *string `json:"DstDesc,omitnil" name:"DstDesc"`
}

func (r *CloneAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SrcUser")
	delete(f, "SrcHost")
	delete(f, "DstUser")
	delete(f, "DstHost")
	delete(f, "DstDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloneAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloneAccountResponseParams struct {
	// Async task flow ID.
	FlowId *uint64 `json:"FlowId,omitnil" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CloneAccountResponse struct {
	*tchttp.BaseResponse
	Response *CloneAccountResponseParams `json:"Response"`
}

func (r *CloneAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseDBExtranetAccessRequestParams struct {
	// ID of instance for which to disable public network access. The ID is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Whether IPv6 is used. Default value: 0
	Ipv6Flag *int64 `json:"Ipv6Flag,omitnil" name:"Ipv6Flag"`
}

type CloseDBExtranetAccessRequest struct {
	*tchttp.BaseRequest
	
	// ID of instance for which to disable public network access. The ID is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Whether IPv6 is used. Default value: 0
	Ipv6Flag *int64 `json:"Ipv6Flag,omitnil" name:"Ipv6Flag"`
}

func (r *CloseDBExtranetAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseDBExtranetAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Ipv6Flag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseDBExtranetAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseDBExtranetAccessResponseParams struct {
	// Async task ID. The task status can be queried through the `DescribeFlow` API.
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CloseDBExtranetAccessResponse struct {
	*tchttp.BaseResponse
	Response *CloseDBExtranetAccessResponseParams `json:"Response"`
}

func (r *CloseDBExtranetAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseDBExtranetAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ColumnPrivilege struct {
	// Database name
	Database *string `json:"Database,omitnil" name:"Database"`

	// Table name
	Table *string `json:"Table,omitnil" name:"Table"`

	// Column name
	Column *string `json:"Column,omitnil" name:"Column"`

	// Permission information
	Privileges []*string `json:"Privileges,omitnil" name:"Privileges"`
}

type ConstraintRange struct {
	// Minimum value when the constraint type is `section`
	Min *string `json:"Min,omitnil" name:"Min"`

	// Maximum value when the constraint type is `section`
	Max *string `json:"Max,omitnil" name:"Max"`
}

// Predefined struct for user
type CopyAccountPrivilegesRequestParams struct {
	// Instance ID, which is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Source username
	SrcUserName *string `json:"SrcUserName,omitnil" name:"SrcUserName"`

	// Access host allowed for source user
	SrcHost *string `json:"SrcHost,omitnil" name:"SrcHost"`

	// Target username
	DstUserName *string `json:"DstUserName,omitnil" name:"DstUserName"`

	// Access host allowed for target user
	DstHost *string `json:"DstHost,omitnil" name:"DstHost"`

	// `ReadOnly` attribute of source account
	SrcReadOnly *string `json:"SrcReadOnly,omitnil" name:"SrcReadOnly"`

	// `ReadOnly` attribute of target account
	DstReadOnly *string `json:"DstReadOnly,omitnil" name:"DstReadOnly"`
}

type CopyAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Source username
	SrcUserName *string `json:"SrcUserName,omitnil" name:"SrcUserName"`

	// Access host allowed for source user
	SrcHost *string `json:"SrcHost,omitnil" name:"SrcHost"`

	// Target username
	DstUserName *string `json:"DstUserName,omitnil" name:"DstUserName"`

	// Access host allowed for target user
	DstHost *string `json:"DstHost,omitnil" name:"DstHost"`

	// `ReadOnly` attribute of source account
	SrcReadOnly *string `json:"SrcReadOnly,omitnil" name:"SrcReadOnly"`

	// `ReadOnly` attribute of target account
	DstReadOnly *string `json:"DstReadOnly,omitnil" name:"DstReadOnly"`
}

func (r *CopyAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyAccountPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SrcUserName")
	delete(f, "SrcHost")
	delete(f, "DstUserName")
	delete(f, "DstHost")
	delete(f, "SrcReadOnly")
	delete(f, "DstReadOnly")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CopyAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CopyAccountPrivilegesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CopyAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *CopyAccountPrivilegesResponseParams `json:"Response"`
}

func (r *CopyAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccountRequestParams struct {
	// Instance ID, which is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Login username, which can contain 1-32 letters, digits, underscores, and hyphens.
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Host that can be logged in to, which is in the same format as the host of the MySQL account and supports wildcards, such as %, 10.%, and 10.20.%.
	Host *string `json:"Host,omitnil" name:"Host"`

	// Account password. It must contain 8-32 characters in all of the following four types: lowercase letters, uppercase letters, digits, and symbols (()~!@#$%^&*-+=_|{}[]:<>,.?/), and cannot start with a slash (/).
	Password *string `json:"Password,omitnil" name:"Password"`

	// Whether to create a read-only account. 0: no; 1: for the account's SQL requests, the secondary will be used first, and if it is unavailable, the primary will be used; 2: the secondary will be used first, and if it is unavailable, the operation will fail.
	ReadOnly *int64 `json:"ReadOnly,omitnil" name:"ReadOnly"`

	// Account remarks, which can contain 0-256 letters, digits, and common symbols.
	Description *string `json:"Description,omitnil" name:"Description"`

	// Determines whether the secondary is unavailable based on the passed-in time
	DelayThresh *int64 `json:"DelayThresh,omitnil" name:"DelayThresh"`

	// Whether to specify a replica server for read-only account. Valid values: `0` (No replica server is specified, which means that the proxy will select another available replica server to keep connection with the client if the current replica server doesn’t meet the requirement). `1` (The replica server is specified, which means that the connection will be disconnected if the specified replica server doesn’t meet the requirement.)
	SlaveConst *int64 `json:"SlaveConst,omitnil" name:"SlaveConst"`

	// Maximum number of connections. If left empty or `0` is passed in, the connections will be unlimited. This parameter configuration is not supported for kernel version 10.1.
	MaxUserConnections *uint64 `json:"MaxUserConnections,omitnil" name:"MaxUserConnections"`
}

type CreateAccountRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Login username, which can contain 1-32 letters, digits, underscores, and hyphens.
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Host that can be logged in to, which is in the same format as the host of the MySQL account and supports wildcards, such as %, 10.%, and 10.20.%.
	Host *string `json:"Host,omitnil" name:"Host"`

	// Account password. It must contain 8-32 characters in all of the following four types: lowercase letters, uppercase letters, digits, and symbols (()~!@#$%^&*-+=_|{}[]:<>,.?/), and cannot start with a slash (/).
	Password *string `json:"Password,omitnil" name:"Password"`

	// Whether to create a read-only account. 0: no; 1: for the account's SQL requests, the secondary will be used first, and if it is unavailable, the primary will be used; 2: the secondary will be used first, and if it is unavailable, the operation will fail.
	ReadOnly *int64 `json:"ReadOnly,omitnil" name:"ReadOnly"`

	// Account remarks, which can contain 0-256 letters, digits, and common symbols.
	Description *string `json:"Description,omitnil" name:"Description"`

	// Determines whether the secondary is unavailable based on the passed-in time
	DelayThresh *int64 `json:"DelayThresh,omitnil" name:"DelayThresh"`

	// Whether to specify a replica server for read-only account. Valid values: `0` (No replica server is specified, which means that the proxy will select another available replica server to keep connection with the client if the current replica server doesn’t meet the requirement). `1` (The replica server is specified, which means that the connection will be disconnected if the specified replica server doesn’t meet the requirement.)
	SlaveConst *int64 `json:"SlaveConst,omitnil" name:"SlaveConst"`

	// Maximum number of connections. If left empty or `0` is passed in, the connections will be unlimited. This parameter configuration is not supported for kernel version 10.1.
	MaxUserConnections *uint64 `json:"MaxUserConnections,omitnil" name:"MaxUserConnections"`
}

func (r *CreateAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Host")
	delete(f, "Password")
	delete(f, "ReadOnly")
	delete(f, "Description")
	delete(f, "DelayThresh")
	delete(f, "SlaveConst")
	delete(f, "MaxUserConnections")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccountResponseParams struct {
	// Instance ID, which is passed through from the input parameters.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Username, which is passed through from the input parameters.
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Host allowed for access, which is passed through from the input parameters.
	Host *string `json:"Host,omitnil" name:"Host"`

	// Passed through from the input parameters.
	ReadOnly *int64 `json:"ReadOnly,omitnil" name:"ReadOnly"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateAccountResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccountResponseParams `json:"Response"`
}

func (r *CreateAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceRequestParams struct {
	// AZs to deploy instance nodes. You can specify up to two AZs (one as primary AZ and another as replica AZ). When the shard specification is 1-primary-2-replica, the primary and one of the replicas are deployed in the primary AZ.
	Zones []*string `json:"Zones,omitnil" name:"Zones"`

	// Number of nodes, which can be obtained 
	//  by querying the instance specification through the `DescribeDBInstanceSpecs` API.
	NodeCount *int64 `json:"NodeCount,omitnil" name:"NodeCount"`

	// Memory size in GB, which can be obtained 
	//  by querying the instance specification through the `DescribeDBInstanceSpecs` API.
	Memory *int64 `json:"Memory,omitnil" name:"Memory"`

	// Storage capacity in GB. The maximum and minimum storage space can be obtained 
	//  by querying instance specification through the `DescribeDBInstanceSpecs` API.
	Storage *int64 `json:"Storage,omitnil" name:"Storage"`

	// Validity period in months
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// The number of instances to be purchased. Only one instance is queried for price by default.
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// Whether to automatically use vouchers. This option is disabled by default.
	AutoVoucher *bool `json:"AutoVoucher,omitnil" name:"AutoVoucher"`

	// Voucher ID list. Currently, you can specify only one voucher.
	VoucherIds []*string `json:"VoucherIds,omitnil" name:"VoucherIds"`

	// VPC ID. If this parameter is not passed in, the instance will be created on the classic network.
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// VPC subnet ID, which is required when `VpcId` is specified.
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// Project ID, which can be obtained through the `DescribeProjects` API. If this parameter is not passed in, the instance will be associated with the default project.
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Database engine version. Valid values: `5.7`, `8.0`, `10.0`, `10.1`.
	DbVersionId *string `json:"DbVersionId,omitnil" name:"DbVersionId"`

	// Name of the instance, which can be customized.
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// List of security group IDs
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`

	// Auto-renewal flag. Valid values: `1` (auto-renewal), `2` (no renewal upon expiration).
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// Whether IPv6 is supported. Valid values: `0` (unsupported), `1` (supported).
	Ipv6Flag *int64 `json:"Ipv6Flag,omitnil" name:"Ipv6Flag"`

	// Array of tag key-value pairs
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil" name:"ResourceTags"`

	// List of parameters. Valid values: `character_set_server` (character set; required); `lower_case_table_names` (table name case sensitivity; required; `0`: case-sensitive; `1`: case-insensitive); `innodb_page_size` (InnoDB data page size; default size: 16 KB); `sync_mode` (sync mode; `0`: async; `1`: strong sync; `2`: downgradable strong sync; default value: `2`).
	InitParams []*DBParamValue `json:"InitParams,omitnil" name:"InitParams"`

	// DCN source region
	DcnRegion *string `json:"DcnRegion,omitnil" name:"DcnRegion"`

	// DCN source instance ID
	DcnInstanceId *string `json:"DcnInstanceId,omitnil" name:"DcnInstanceId"`
}

type CreateDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// AZs to deploy instance nodes. You can specify up to two AZs (one as primary AZ and another as replica AZ). When the shard specification is 1-primary-2-replica, the primary and one of the replicas are deployed in the primary AZ.
	Zones []*string `json:"Zones,omitnil" name:"Zones"`

	// Number of nodes, which can be obtained 
	//  by querying the instance specification through the `DescribeDBInstanceSpecs` API.
	NodeCount *int64 `json:"NodeCount,omitnil" name:"NodeCount"`

	// Memory size in GB, which can be obtained 
	//  by querying the instance specification through the `DescribeDBInstanceSpecs` API.
	Memory *int64 `json:"Memory,omitnil" name:"Memory"`

	// Storage capacity in GB. The maximum and minimum storage space can be obtained 
	//  by querying instance specification through the `DescribeDBInstanceSpecs` API.
	Storage *int64 `json:"Storage,omitnil" name:"Storage"`

	// Validity period in months
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// The number of instances to be purchased. Only one instance is queried for price by default.
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// Whether to automatically use vouchers. This option is disabled by default.
	AutoVoucher *bool `json:"AutoVoucher,omitnil" name:"AutoVoucher"`

	// Voucher ID list. Currently, you can specify only one voucher.
	VoucherIds []*string `json:"VoucherIds,omitnil" name:"VoucherIds"`

	// VPC ID. If this parameter is not passed in, the instance will be created on the classic network.
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// VPC subnet ID, which is required when `VpcId` is specified.
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// Project ID, which can be obtained through the `DescribeProjects` API. If this parameter is not passed in, the instance will be associated with the default project.
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Database engine version. Valid values: `5.7`, `8.0`, `10.0`, `10.1`.
	DbVersionId *string `json:"DbVersionId,omitnil" name:"DbVersionId"`

	// Name of the instance, which can be customized.
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// List of security group IDs
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`

	// Auto-renewal flag. Valid values: `1` (auto-renewal), `2` (no renewal upon expiration).
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// Whether IPv6 is supported. Valid values: `0` (unsupported), `1` (supported).
	Ipv6Flag *int64 `json:"Ipv6Flag,omitnil" name:"Ipv6Flag"`

	// Array of tag key-value pairs
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil" name:"ResourceTags"`

	// List of parameters. Valid values: `character_set_server` (character set; required); `lower_case_table_names` (table name case sensitivity; required; `0`: case-sensitive; `1`: case-insensitive); `innodb_page_size` (InnoDB data page size; default size: 16 KB); `sync_mode` (sync mode; `0`: async; `1`: strong sync; `2`: downgradable strong sync; default value: `2`).
	InitParams []*DBParamValue `json:"InitParams,omitnil" name:"InitParams"`

	// DCN source region
	DcnRegion *string `json:"DcnRegion,omitnil" name:"DcnRegion"`

	// DCN source instance ID
	DcnInstanceId *string `json:"DcnInstanceId,omitnil" name:"DcnInstanceId"`
}

func (r *CreateDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zones")
	delete(f, "NodeCount")
	delete(f, "Memory")
	delete(f, "Storage")
	delete(f, "Period")
	delete(f, "Count")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "ProjectId")
	delete(f, "DbVersionId")
	delete(f, "InstanceName")
	delete(f, "SecurityGroupIds")
	delete(f, "AutoRenewFlag")
	delete(f, "Ipv6Flag")
	delete(f, "ResourceTags")
	delete(f, "InitParams")
	delete(f, "DcnRegion")
	delete(f, "DcnInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceResponseParams struct {
	// Order ID, which is used for calling the `DescribeOrders` API.
	//  The parameter can be used to either query order details or call the user account APIs to make another payment when this payment fails.
	DealName *string `json:"DealName,omitnil" name:"DealName"`

	// IDs of the instances you have purchased in this order. If no instance IDs are returned, you can query them with the `DescribeOrders` API. You can also use the `DescribeDBInstances` API to check whether an instance has been created successfully.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBInstanceResponseParams `json:"Response"`
}

func (r *CreateDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHourDBInstanceRequestParams struct {
	// AZs to deploy instance nodes. You can specify up to two AZs.
	Zones []*string `json:"Zones,omitnil" name:"Zones"`

	// Number of nodes.
	NodeCount *int64 `json:"NodeCount,omitnil" name:"NodeCount"`

	// Memory size in GB.
	Memory *int64 `json:"Memory,omitnil" name:"Memory"`

	// Storage size in GB.
	Storage *int64 `json:"Storage,omitnil" name:"Storage"`

	// Number of instances to purchase.
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// Project ID. If this parameter is not passed in, the default project will be used.
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Unique ID of the network. If this parameter is not passed in, the classic network will be used.
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// Unique ID of the subnet. If `VpcId` is specified, this parameter is required.
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// Database engine version. Valid values: `5.7`, `8.0`, `10.0`, `10.1`.
	DbVersionId *string `json:"DbVersionId,omitnil" name:"DbVersionId"`

	// Custom name of the instance.
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Security group ID. If this parameter is not passed in, no security groups will be associated when the instance is created.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`

	// Whether IPv6 is supported. Valid values: `0` (unsupported), `1` (supported).
	Ipv6Flag *int64 `json:"Ipv6Flag,omitnil" name:"Ipv6Flag"`

	// Array of tag key-value pairs.
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil" name:"ResourceTags"`

	// If you create a disaster recovery instance, you need to use this parameter to specify the region of the associated primary instance so that the disaster recovery instance can sync data with the primary instance over the Data Communication Network (DCN).
	DcnRegion *string `json:"DcnRegion,omitnil" name:"DcnRegion"`

	// If you create a disaster recovery instance, you need to use this parameter to specify the ID of the associated primary instance so that the disaster recovery instance can sync data with the primary instance over the Data Communication Network (DCN).
	DcnInstanceId *string `json:"DcnInstanceId,omitnil" name:"DcnInstanceId"`

	// List of parameters. Valid values: 
	// `character_set_server` (character set; required); `lower_case_table_names` (table name case sensitivity; required; 0: case-sensitive; 1: case-insensitive);
	// `innodb_page_size` (innoDB data page size; default size: 16 KB); `sync_mode` (sync mode; 0: async; 1: strong sync; 2: downgradable strong sync; default value: 2).
	InitParams []*DBParamValue `json:"InitParams,omitnil" name:"InitParams"`

	// ID of the instance to be rolled back, such as “2021-11-22 00:00:00”.
	RollbackInstanceId *string `json:"RollbackInstanceId,omitnil" name:"RollbackInstanceId"`

	// Rollback time.
	RollbackTime *string `json:"RollbackTime,omitnil" name:"RollbackTime"`
}

type CreateHourDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// AZs to deploy instance nodes. You can specify up to two AZs.
	Zones []*string `json:"Zones,omitnil" name:"Zones"`

	// Number of nodes.
	NodeCount *int64 `json:"NodeCount,omitnil" name:"NodeCount"`

	// Memory size in GB.
	Memory *int64 `json:"Memory,omitnil" name:"Memory"`

	// Storage size in GB.
	Storage *int64 `json:"Storage,omitnil" name:"Storage"`

	// Number of instances to purchase.
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// Project ID. If this parameter is not passed in, the default project will be used.
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Unique ID of the network. If this parameter is not passed in, the classic network will be used.
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// Unique ID of the subnet. If `VpcId` is specified, this parameter is required.
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// Database engine version. Valid values: `5.7`, `8.0`, `10.0`, `10.1`.
	DbVersionId *string `json:"DbVersionId,omitnil" name:"DbVersionId"`

	// Custom name of the instance.
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Security group ID. If this parameter is not passed in, no security groups will be associated when the instance is created.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`

	// Whether IPv6 is supported. Valid values: `0` (unsupported), `1` (supported).
	Ipv6Flag *int64 `json:"Ipv6Flag,omitnil" name:"Ipv6Flag"`

	// Array of tag key-value pairs.
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil" name:"ResourceTags"`

	// If you create a disaster recovery instance, you need to use this parameter to specify the region of the associated primary instance so that the disaster recovery instance can sync data with the primary instance over the Data Communication Network (DCN).
	DcnRegion *string `json:"DcnRegion,omitnil" name:"DcnRegion"`

	// If you create a disaster recovery instance, you need to use this parameter to specify the ID of the associated primary instance so that the disaster recovery instance can sync data with the primary instance over the Data Communication Network (DCN).
	DcnInstanceId *string `json:"DcnInstanceId,omitnil" name:"DcnInstanceId"`

	// List of parameters. Valid values: 
	// `character_set_server` (character set; required); `lower_case_table_names` (table name case sensitivity; required; 0: case-sensitive; 1: case-insensitive);
	// `innodb_page_size` (innoDB data page size; default size: 16 KB); `sync_mode` (sync mode; 0: async; 1: strong sync; 2: downgradable strong sync; default value: 2).
	InitParams []*DBParamValue `json:"InitParams,omitnil" name:"InitParams"`

	// ID of the instance to be rolled back, such as “2021-11-22 00:00:00”.
	RollbackInstanceId *string `json:"RollbackInstanceId,omitnil" name:"RollbackInstanceId"`

	// Rollback time.
	RollbackTime *string `json:"RollbackTime,omitnil" name:"RollbackTime"`
}

func (r *CreateHourDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHourDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zones")
	delete(f, "NodeCount")
	delete(f, "Memory")
	delete(f, "Storage")
	delete(f, "Count")
	delete(f, "ProjectId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "DbVersionId")
	delete(f, "InstanceName")
	delete(f, "SecurityGroupIds")
	delete(f, "Ipv6Flag")
	delete(f, "ResourceTags")
	delete(f, "DcnRegion")
	delete(f, "DcnInstanceId")
	delete(f, "InitParams")
	delete(f, "RollbackInstanceId")
	delete(f, "RollbackTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateHourDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHourDBInstanceResponseParams struct {
	// Order ID, which is used for calling the `DescribeOrders` API.
	//  The parameter can be used to either query order details or call the user account APIs to make another payment when this payment fails.
	DealName *string `json:"DealName,omitnil" name:"DealName"`

	// IDs of the instances you have purchased in this order. If no instance IDs are returned, you can query them with the `DescribeOrders` API. You can also use the `DescribeDBInstances` API to check whether an instance has been created successfully.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Async task ID, which can be used in the [DescribeFlow](https://www.tencentcloud.com/document/product/237/16177) API to query the async task result.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateHourDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateHourDBInstanceResponseParams `json:"Response"`
}

func (r *CreateHourDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHourDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DBAccount struct {
	// Username
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Host from which a user can log in (corresponding to the `host` field for a MySQL user; a user is uniquely identified by username and host; this parameter is in IP format and ends with % for IP range; % can be entered; if this parameter is left empty, % will be used by default).
	Host *string `json:"Host,omitnil" name:"Host"`

	// User remarks
	Description *string `json:"Description,omitnil" name:"Description"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Last updated time
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// Read-only flag. 0: no; 1: for the account's SQL requests, the replica will be used first, and if it is unavailable, the primary will be used; 2: the replica will be used first, and if it is unavailable, the operation will fail.
	ReadOnly *int64 `json:"ReadOnly,omitnil" name:"ReadOnly"`

	// This field is meaningful for read-only accounts, indicating that a replica should be selected if its delay from the primary is less than this value.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DelayThresh *int64 `json:"DelayThresh,omitnil" name:"DelayThresh"`

	// Whether to specify a replica server for read-only account. Valid values: `0` (No replica server is specified, which means that the proxy will select another available replica server to keep connection with the client if the current replica server doesn’t meet the requirement). `1` (The replica server is specified, which means that the connection will be disconnected if the specified replica server doesn’t meet the requirement.)
	SlaveConst *int64 `json:"SlaveConst,omitnil" name:"SlaveConst"`

	// Maximum number of connections. `0` indicates no limit.
	MaxUserConnections *int64 `json:"MaxUserConnections,omitnil" name:"MaxUserConnections"`
}

type DBInstance struct {
	// Instance ID, which uniquely identifies a TDSQL instance
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Customizable instance name
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Application ID of instance
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`

	// Project ID of instance
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Instance region name, such as ap-shanghai
	Region *string `json:"Region,omitnil" name:"Region"`

	// Instance AZ name, such as ap-guangzhou-1
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// VPC ID, which is 0 if the basic network is used
	VpcId *int64 `json:"VpcId,omitnil" name:"VpcId"`

	// Subnet ID, which is 0 if the basic network is used
	SubnetId *int64 `json:"SubnetId,omitnil" name:"SubnetId"`

	// Instance status. Valid values: `0` (creating), `1` (running task), `2` (running), `3` (uninitialized), `-1` (isolated), `4` (initializing), `5` (eliminating), `6` (restarting), `7` (migrating data)
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// Private IP address
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// Private network port
	Vport *int64 `json:"Vport,omitnil" name:"Vport"`

	// Domain name for public network access, which can be resolved by the public network
	WanDomain *string `json:"WanDomain,omitnil" name:"WanDomain"`

	// Public IP address, which can be accessed over the public network
	WanVip *string `json:"WanVip,omitnil" name:"WanVip"`

	// Public network port
	WanPort *int64 `json:"WanPort,omitnil" name:"WanPort"`

	// Instance creation time in the format of `2006-01-02 15:04:05`
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Last updated time of instance in the format of `2006-01-02 15:04:05`
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// Auto-renewal flag. 0: no, 1: yes
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// Instance expiration time in the format of `2006-01-02 15:04:05`
	PeriodEndTime *string `json:"PeriodEndTime,omitnil" name:"PeriodEndTime"`

	// Instance account
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// TDSQL version information
	TdsqlVersion *string `json:"TdsqlVersion,omitnil" name:"TdsqlVersion"`

	// Instance memory size in GB
	Memory *int64 `json:"Memory,omitnil" name:"Memory"`

	// Instance storage capacity in GB
	Storage *int64 `json:"Storage,omitnil" name:"Storage"`

	// VPC ID in string type
	UniqueVpcId *string `json:"UniqueVpcId,omitnil" name:"UniqueVpcId"`

	// VPC subnet ID in string type
	UniqueSubnetId *string `json:"UniqueSubnetId,omitnil" name:"UniqueSubnetId"`

	// Original ID of instance (this field is obsolete and should not be depended on)
	OriginSerialId *string `json:"OriginSerialId,omitnil" name:"OriginSerialId"`

	// Number of nodes. 2: one master and one slave, 3: one master and two slaves
	NodeCount *uint64 `json:"NodeCount,omitnil" name:"NodeCount"`

	// Whether it is a temp instance. 0: no, non-zero value: yes
	IsTmp *uint64 `json:"IsTmp,omitnil" name:"IsTmp"`

	// Dedicated cluster ID. If this parameter is empty, the instance is a general instance
	ExclusterId *string `json:"ExclusterId,omitnil" name:"ExclusterId"`

	// Numeric ID of instance (this field is obsolete and should not be depended on)
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// Product type ID
	Pid *int64 `json:"Pid,omitnil" name:"Pid"`

	// Maximum QPS value
	Qps *int64 `json:"Qps,omitnil" name:"Qps"`

	// Billing mode
	// Note: this field may return null, indicating that no valid values can be obtained.
	Paymode *string `json:"Paymode,omitnil" name:"Paymode"`

	// Async task flow ID when an async task is in progress on an instance
	// Note: this field may return null, indicating that no valid values can be obtained.
	Locker *int64 `json:"Locker,omitnil" name:"Locker"`

	// Current instance running status description
	// Note: this field may return null, indicating that no valid values can be obtained.
	StatusDesc *string `json:"StatusDesc,omitnil" name:"StatusDesc"`

	// Public network access status. 0: not enabled, 1: enabled, 2: disabled, 3: enabling
	WanStatus *int64 `json:"WanStatus,omitnil" name:"WanStatus"`

	// Whether the instance supports audit. 1: yes, 0: no
	IsAuditSupported *uint64 `json:"IsAuditSupported,omitnil" name:"IsAuditSupported"`

	// Model
	Machine *string `json:"Machine,omitnil" name:"Machine"`

	// Whether data encryption is supported. 1: yes, 0: no
	IsEncryptSupported *int64 `json:"IsEncryptSupported,omitnil" name:"IsEncryptSupported"`

	// Number of CPU cores of instance
	Cpu *int64 `json:"Cpu,omitnil" name:"Cpu"`

	// Indicates whether the instance uses IPv6
	// Note: this field may return null, indicating that no valid values can be obtained.
	Ipv6Flag *uint64 `json:"Ipv6Flag,omitnil" name:"Ipv6Flag"`

	// Private network IPv6 address
	// Note: this field may return null, indicating that no valid values can be obtained.
	Vipv6 *string `json:"Vipv6,omitnil" name:"Vipv6"`

	// Public network IPv6 address
	// Note: this field may return null, indicating that no valid values can be obtained.
	WanVipv6 *string `json:"WanVipv6,omitnil" name:"WanVipv6"`

	// Public network IPv6 port
	// Note: this field may return null, indicating that no valid values can be obtained.
	WanPortIpv6 *uint64 `json:"WanPortIpv6,omitnil" name:"WanPortIpv6"`

	// Public network IPv6 status
	// Note: this field may return null, indicating that no valid values can be obtained.
	WanStatusIpv6 *uint64 `json:"WanStatusIpv6,omitnil" name:"WanStatusIpv6"`

	// Database engine
	// Note: this field may return null, indicating that no valid values can be obtained.
	DbEngine *string `json:"DbEngine,omitnil" name:"DbEngine"`

	// Database version
	// Note: this field may return null, indicating that no valid values can be obtained.
	DbVersion *string `json:"DbVersion,omitnil" name:"DbVersion"`

	// DCN type. Valid values: 0 (null), 1 (primary instance), 2 (disaster recovery instance)
	// Note: this field may return null, indicating that no valid values can be obtained.
	DcnFlag *int64 `json:"DcnFlag,omitnil" name:"DcnFlag"`

	// DCN status. Valid values: 0 (null), 1 (creating), 2 (syncing), 3 (disconnected)
	// Note: this field may return null, indicating that no valid values can be obtained.
	DcnStatus *int64 `json:"DcnStatus,omitnil" name:"DcnStatus"`

	// The number of DCN disaster recovery instances
	// Note: this field may return null, indicating that no valid values can be obtained.
	DcnDstNum *int64 `json:"DcnDstNum,omitnil" name:"DcnDstNum"`

	// Instance type. Valid values: `1` (dedicated primary instance), `2` (primary instance), `3` (disaster recovery instance), and `4` (dedicated disaster recovery instance).
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceType *int64 `json:"InstanceType,omitnil" name:"InstanceType"`

	// Instance tag information
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil" name:"ResourceTags"`

	// Database version
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbVersionId *string `json:"DbVersionId,omitnil" name:"DbVersionId"`
}

type DBParamValue struct {
	// Parameter name
	Param *string `json:"Param,omitnil" name:"Param"`

	// Parameter value
	Value *string `json:"Value,omitnil" name:"Value"`
}

type DCNReplicaConfig struct {
	// DCN running status. Valid values: `START` (running), `STOP` (pause)
	// Note: This field may return null, indicating that no valid values can be obtained.
	RoReplicationMode *string `json:"RoReplicationMode,omitnil" name:"RoReplicationMode"`

	// Delayed replication type. Valid values: `DEFAULT` (no delay), `DUE_TIME` (specified replication time)
	// Note: This field may return null, indicating that no valid values can be obtained.
	DelayReplicationType *string `json:"DelayReplicationType,omitnil" name:"DelayReplicationType"`

	// Specified time for delayed replication
	// Note: This field may return null, indicating that no valid values can be obtained.
	DueTime *string `json:"DueTime,omitnil" name:"DueTime"`

	// The number of seconds to delay the replication
	// Note: This field may return null, indicating that no valid values can be obtained.
	ReplicationDelay *int64 `json:"ReplicationDelay,omitnil" name:"ReplicationDelay"`
}

type DCNReplicaStatus struct {
	// DCN running status. Valid values: `START` (running), `STOP` (pause).
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil" name:"Status"`

	// The current delay, which takes the delay value of the replica instance.
	Delay *int64 `json:"Delay,omitnil" name:"Delay"`
}

type Database struct {
	// Database name
	DbName *string `json:"DbName,omitnil" name:"DbName"`
}

type DatabaseFunction struct {
	// Function name
	Func *string `json:"Func,omitnil" name:"Func"`
}

type DatabasePrivilege struct {
	// Permission information
	Privileges []*string `json:"Privileges,omitnil" name:"Privileges"`

	// Database name
	Database *string `json:"Database,omitnil" name:"Database"`
}

type DatabaseProcedure struct {
	// Stored procedure name
	Proc *string `json:"Proc,omitnil" name:"Proc"`
}

type DatabaseTable struct {
	// Table name
	Table *string `json:"Table,omitnil" name:"Table"`
}

type DatabaseView struct {
	// View name
	View *string `json:"View,omitnil" name:"View"`
}

type DcnDetailItem struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Region where the instance resides
	Region *string `json:"Region,omitnil" name:"Region"`

	// Availability zone where the instance resides
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// Instance IP address
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// Instance IPv6 address
	Vipv6 *string `json:"Vipv6,omitnil" name:"Vipv6"`

	// Instance port
	Vport *int64 `json:"Vport,omitnil" name:"Vport"`

	// Instance status
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// Instance status description
	StatusDesc *string `json:"StatusDesc,omitnil" name:"StatusDesc"`

	// DCN flag. Valid values: `1` (primary), `2` (disaster recovery)
	DcnFlag *int64 `json:"DcnFlag,omitnil" name:"DcnFlag"`

	// DCN status. Valid values: `0` (none), `1` (creating), `2` (syncing), `3` (disconnected)
	DcnStatus *int64 `json:"DcnStatus,omitnil" name:"DcnStatus"`

	// Number of CPU cores of the instance
	Cpu *int64 `json:"Cpu,omitnil" name:"Cpu"`

	// Instance memory capacity in GB
	Memory *int64 `json:"Memory,omitnil" name:"Memory"`

	// Instance storage capacity in GB
	Storage *int64 `json:"Storage,omitnil" name:"Storage"`

	// Billing mode
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`

	// Creation time of the instance in the format of 2006-01-02 15:04:05
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Expiration time of the instance in the format of 2006-01-02 15:04:05
	PeriodEndTime *string `json:"PeriodEndTime,omitnil" name:"PeriodEndTime"`

	// Instance type. Valid values: `1` (dedicated primary instance), `2` (non-dedicated primary instance), `3` (non-dedicated disaster recovery instance), `4` (dedicated disaster recovery instance)
	InstanceType *int64 `json:"InstanceType,omitnil" name:"InstanceType"`

	// Configuration information of DCN replication. This field is null for a primary instance.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ReplicaConfig *DCNReplicaConfig `json:"ReplicaConfig,omitnil" name:"ReplicaConfig"`

	// DCN replication status. This field is null for the primary instance.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ReplicaStatus *DCNReplicaStatus `json:"ReplicaStatus,omitnil" name:"ReplicaStatus"`

	// Whether KMS is enabled.
	EncryptStatus *int64 `json:"EncryptStatus,omitnil" name:"EncryptStatus"`
}

type Deal struct {
	// Order number
	DealName *string `json:"DealName,omitnil" name:"DealName"`

	// Account
	OwnerUin *string `json:"OwnerUin,omitnil" name:"OwnerUin"`

	// Number of items
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// ID of the associated process, which can be used to query the process execution status.
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// The ID of the created instance, which is required only for the order that creates an instance.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Payment mode. Valid values: 0 (postpaid), 1 (prepaid)
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`
}

// Predefined struct for user
type DeleteAccountRequestParams struct {
	// Instance ID, which is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Username
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Access host allowed for user
	Host *string `json:"Host,omitnil" name:"Host"`
}

type DeleteAccountRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Username
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Access host allowed for user
	Host *string `json:"Host,omitnil" name:"Host"`
}

func (r *DeleteAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Host")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccountResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteAccountResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAccountResponseParams `json:"Response"`
}

func (r *DeleteAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountPrivilegesRequestParams struct {
	// Instance ID in the form of `tdsql-ow728lmc`, which can be obtained by querying the instance details through `DescribeDBInstances`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Login username.
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Access host allowed for a user. An account is uniquely identified by username and host.
	Host *string `json:"Host,omitnil" name:"Host"`

	// Database name. `\*` indicates that global permissions will be queried (i.e., `\*.\*`), in which case the `Type` and `Object ` parameters will be ignored.
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// Type. Valid values: table, view, proc, func, \*. If `DbName` is a specific database name and `Type` is `\*`, the permissions of the database will be queried (i.e., `db.\*`), in which case the `Object` parameter will be ignored.
	Type *string `json:"Type,omitnil" name:"Type"`

	// Type name. For example, if `Type` is `table`, `Object` indicates a specific table name; if both `DbName` and `Type` are specific names, it indicates a specific object name and cannot be `\*` or empty.
	Object *string `json:"Object,omitnil" name:"Object"`

	// If `Type` is `table` and `ColName` is `\*`, the permissions of the table will be queried; if `ColName` is a specific field name, the permissions of the corresponding field will be queried.
	ColName *string `json:"ColName,omitnil" name:"ColName"`
}

type DescribeAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the form of `tdsql-ow728lmc`, which can be obtained by querying the instance details through `DescribeDBInstances`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Login username.
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Access host allowed for a user. An account is uniquely identified by username and host.
	Host *string `json:"Host,omitnil" name:"Host"`

	// Database name. `\*` indicates that global permissions will be queried (i.e., `\*.\*`), in which case the `Type` and `Object ` parameters will be ignored.
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// Type. Valid values: table, view, proc, func, \*. If `DbName` is a specific database name and `Type` is `\*`, the permissions of the database will be queried (i.e., `db.\*`), in which case the `Object` parameter will be ignored.
	Type *string `json:"Type,omitnil" name:"Type"`

	// Type name. For example, if `Type` is `table`, `Object` indicates a specific table name; if both `DbName` and `Type` are specific names, it indicates a specific object name and cannot be `\*` or empty.
	Object *string `json:"Object,omitnil" name:"Object"`

	// If `Type` is `table` and `ColName` is `\*`, the permissions of the table will be queried; if `ColName` is a specific field name, the permissions of the corresponding field will be queried.
	ColName *string `json:"ColName,omitnil" name:"ColName"`
}

func (r *DescribeAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Host")
	delete(f, "DbName")
	delete(f, "Type")
	delete(f, "Object")
	delete(f, "ColName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountPrivilegesResponseParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Permission list.
	Privileges []*string `json:"Privileges,omitnil" name:"Privileges"`

	// Database account username
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Database account host
	Host *string `json:"Host,omitnil" name:"Host"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountPrivilegesResponseParams `json:"Response"`
}

func (r *DescribeAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsRequestParams struct {
	// Instance ID in the form of `tdsql-ow728lmc`, which can be obtained by querying the instance details through `DescribeDBInstances`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the form of `tdsql-ow728lmc`, which can be obtained by querying the instance details through `DescribeDBInstances`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsResponseParams struct {
	// Instance ID, which is passed through from the input parameters.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance user list.
	Users []*DBAccount `json:"Users,omitnil" name:"Users"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountsResponseParams `json:"Response"`
}

func (r *DescribeAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupFilesRequestParams struct {
	// Query by instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Backup type. Valid values: `Data` (data backup), `Binlog` (Binlog backup), `Errlog` (error log), `Slowlog` (slow log).
	BackupType *string `json:"BackupType,omitnil" name:"BackupType"`

	// Query by start time
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// Query by end time
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Pagination parameter
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Pagination parameter
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Sorting dimension. Valid values: `Time`, `Size`.
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// Sorting order. Valid values: `DESC`, `ASC`.
	OrderType *string `json:"OrderType,omitnil" name:"OrderType"`
}

type DescribeBackupFilesRequest struct {
	*tchttp.BaseRequest
	
	// Query by instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Backup type. Valid values: `Data` (data backup), `Binlog` (Binlog backup), `Errlog` (error log), `Slowlog` (slow log).
	BackupType *string `json:"BackupType,omitnil" name:"BackupType"`

	// Query by start time
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// Query by end time
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Pagination parameter
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Pagination parameter
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Sorting dimension. Valid values: `Time`, `Size`.
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// Sorting order. Valid values: `DESC`, `ASC`.
	OrderType *string `json:"OrderType,omitnil" name:"OrderType"`
}

func (r *DescribeBackupFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupType")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupFilesResponseParams struct {
	// List of backup files
	Files []*InstanceBackupFileItem `json:"Files,omitnil" name:"Files"`

	// Total number
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBackupFilesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupFilesResponseParams `json:"Response"`
}

func (r *DescribeBackupFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBEncryptAttributesRequestParams struct {
	// Instance ID in the format of `tdsql-ow728lmc`
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeDBEncryptAttributesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of `tdsql-ow728lmc`
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeDBEncryptAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBEncryptAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBEncryptAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBEncryptAttributesResponseParams struct {
	// Whether encryption is enabled. Valid values: `1` (enabled), `2` (disabled).
	EncryptStatus *int64 `json:"EncryptStatus,omitnil" name:"EncryptStatus"`

	// DEK key
	CipherText *string `json:"CipherText,omitnil" name:"CipherText"`

	// DEK key expiration date
	ExpireDate *string `json:"ExpireDate,omitnil" name:"ExpireDate"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDBEncryptAttributesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBEncryptAttributesResponseParams `json:"Response"`
}

func (r *DescribeDBEncryptAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBEncryptAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceDetailRequestParams struct {
	// Instance ID in the format of `tdsql-ow728lmc`
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeDBInstanceDetailRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of `tdsql-ow728lmc`
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeDBInstanceDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceDetailResponseParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Instance status
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// Current status of the instance
	StatusDesc *string `json:"StatusDesc,omitnil" name:"StatusDesc"`

	// Private IP address
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// Private port
	Vport *int64 `json:"Vport,omitnil" name:"Vport"`

	// Whether it is a temp instance. Valid values: `0` (no), non-zero value (yes).
	IsTmp *int64 `json:"IsTmp,omitnil" name:"IsTmp"`

	// Number of nodes. Valid values: `2` (1 primary-1 replica), `3` (1 primary-2 replicas).
	NodeCount *int64 `json:"NodeCount,omitnil" name:"NodeCount"`

	// Instance region name, such as ap-shanghai.
	Region *string `json:"Region,omitnil" name:"Region"`

	// Instance AZ name, such as ap-guangzhou-1.
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// VPC ID in string type
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// VPC subnet ID in string type
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// Public network status. Valid values: `0` (not enabled), `1` (enabled), `2` (disabled), `3`: (enabling), `4` (disabling).
	WanStatus *int64 `json:"WanStatus,omitnil" name:"WanStatus"`

	// Domain name for public network access, which can be resolved by the public network.
	WanDomain *string `json:"WanDomain,omitnil" name:"WanDomain"`

	// Public IP address, which can be accessed over the public network.
	WanVip *string `json:"WanVip,omitnil" name:"WanVip"`

	// Public network port
	WanPort *int64 `json:"WanPort,omitnil" name:"WanPort"`

	// Project ID of the instance
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// TDSQL version information
	TdsqlVersion *string `json:"TdsqlVersion,omitnil" name:"TdsqlVersion"`

	// Instance memory capacity in GB
	Memory *int64 `json:"Memory,omitnil" name:"Memory"`

	// Instance storage capacity in GB
	Storage *int64 `json:"Storage,omitnil" name:"Storage"`

	// Primary AZ, such as ap-shanghai-1.
	MasterZone *string `json:"MasterZone,omitnil" name:"MasterZone"`

	// List of replica AZs, such as ap-shanghai-2.
	SlaveZones []*string `json:"SlaveZones,omitnil" name:"SlaveZones"`

	// Auto-renewal flag. Valid values: `0` (no), `1` (yes).
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// Dedicated cluster ID, which is empty for a non-dedicated cluster instance.
	ExclusterId *string `json:"ExclusterId,omitnil" name:"ExclusterId"`

	// Billing mode. Valid values: `prepaid` (monthly subscription), `postpaid` (pay-as-you-go).
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// Creation time of the instance in the format of 2006-01-02 15:04:05
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Whether the instance supports audit
	IsAuditSupported *bool `json:"IsAuditSupported,omitnil" name:"IsAuditSupported"`

	// Expiration time of the instance in the format of 2006-01-02 15:04:05
	PeriodEndTime *string `json:"PeriodEndTime,omitnil" name:"PeriodEndTime"`

	// Model information
	Machine *string `json:"Machine,omitnil" name:"Machine"`

	// Storage space utilization
	StorageUsage *string `json:"StorageUsage,omitnil" name:"StorageUsage"`

	// Size of log storage space in GB
	LogStorage *int64 `json:"LogStorage,omitnil" name:"LogStorage"`

	// Whether data encryption is supported. Valid values: `1` (yes), `0`: (no).
	IsEncryptSupported *int64 `json:"IsEncryptSupported,omitnil" name:"IsEncryptSupported"`

	// Private network IPv6 address
	// Note: This field may return null, indicating that no valid values can be obtained.
	Vip6 *string `json:"Vip6,omitnil" name:"Vip6"`

	// Number of CPU cores of an instance.
	Cpu *int64 `json:"Cpu,omitnil" name:"Cpu"`

	// Product type ID
	Pid *int64 `json:"Pid,omitnil" name:"Pid"`

	// Max QPS
	Qps *int64 `json:"Qps,omitnil" name:"Qps"`

	// Whether IPv6 is supported.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Ipv6Flag *int64 `json:"Ipv6Flag,omitnil" name:"Ipv6Flag"`

	// Public IPv6 address, which can be accessed over the public network
	// Note: This field may return null, indicating that no valid values can be obtained.
	WanVipv6 *string `json:"WanVipv6,omitnil" name:"WanVipv6"`

	// Public network status. Valid values: `0` (not enabled), `1` (enabled), `2` (disabled), `3`: (enabling), `4` (disabling).
	// Note: This field may return null, indicating that no valid values can be obtained.
	WanStatusIpv6 *int64 `json:"WanStatusIpv6,omitnil" name:"WanStatusIpv6"`

	// Public network IPv6 port
	// Note: This field may return null, indicating that no valid values can be obtained.
	WanPortIpv6 *int64 `json:"WanPortIpv6,omitnil" name:"WanPortIpv6"`

	// Database engine
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbEngine *string `json:"DbEngine,omitnil" name:"DbEngine"`

	// Database version
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbVersion *string `json:"DbVersion,omitnil" name:"DbVersion"`

	// Tag information
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil" name:"ResourceTags"`

	// DCN type. Valid values: `0` (N/A), `1` (primary instance), `2` (disaster recovery read-only instance)
	// Note: This field may return null, indicating that no valid values can be obtained.
	DcnFlag *int64 `json:"DcnFlag,omitnil" name:"DcnFlag"`

	// DCN status. Valid values: `0` (N/A), `1` (creating), `2` (syncing), `3` (disconnected)
	// Note: This field may return null, indicating that no valid values can be obtained.
	DcnStatus *int64 `json:"DcnStatus,omitnil" name:"DcnStatus"`

	// Number of disaster recovery read-only instances
	// Note: This field may return null, indicating that no valid values can be obtained.
	DcnDstNum *int64 `json:"DcnDstNum,omitnil" name:"DcnDstNum"`

	// Instance type. Valid values: `1` (dedicated primary instance), `2` (non-dedicated primary instance), `3` (non-dedicated disaster recovery read-only instance), `4` (dedicated disaster recovery read-only instance)
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceType *int64 `json:"InstanceType,omitnil" name:"InstanceType"`

	// Instance node information
	// Note: This field may return null, indicating that no valid values can be obtained.
	NodesInfo []*NodeInfo `json:"NodesInfo,omitnil" name:"NodesInfo"`

	// Whether the instance supports setting the connection limit, which is not supported for kernel version 10.1.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsMaxUserConnectionsSupported *bool `json:"IsMaxUserConnectionsSupported,omitnil" name:"IsMaxUserConnectionsSupported"`

	// The displayed database version
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbVersionId *string `json:"DbVersionId,omitnil" name:"DbVersionId"`

	// Encryption status. Valid values: `0` (disabled), `1` (enabled).
	// Note: This field may return null, indicating that no valid values can be obtained.
	EncryptStatus *int64 `json:"EncryptStatus,omitnil" name:"EncryptStatus"`

	// Configuration information of DCN
	// Note: This field may return null, indicating that no valid values can be obtained.
	ReplicaConfig *DCNReplicaConfig `json:"ReplicaConfig,omitnil" name:"ReplicaConfig"`

	// Running status of DCN
	// Note: This field may return null, indicating that no valid values can be obtained.
	ReplicaStatus *DCNReplicaStatus `json:"ReplicaStatus,omitnil" name:"ReplicaStatus"`

	// Type of dedicated cluster. Valid values: `0` (public cloud), `1` (finance cage), `2` (CDC cluster).
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExclusterType *int64 `json:"ExclusterType,omitnil" name:"ExclusterType"`

	// Nearby VPC access
	// Note: This field may return null, indicating that no valid values can be obtained.
	RsAccessStrategy *int64 `json:"RsAccessStrategy,omitnil" name:"RsAccessStrategy"`

	// Unclaimed network resource
	ReservedNetResources []*ReservedNetResource `json:"ReservedNetResources,omitnil" name:"ReservedNetResources"`


	IsPhysicalReplicationSupported *bool `json:"IsPhysicalReplicationSupported,omitnil" name:"IsPhysicalReplicationSupported"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDBInstanceDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceDetailResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesRequestParams struct {
	// Queries by instance ID or IDs. Instance ID is in the format of `tdsql-ow728lmc`. Up to 100 instances can be queried in one request.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Search field name. Valid values: instancename (search by instance name), vip (search by private IP), all (search by instance ID, instance name, and private IP).
	SearchName *string `json:"SearchName,omitnil" name:"SearchName"`

	// Search keyword. Fuzzy search is supported. Multiple keywords should be separated by line breaks (`\n`).
	SearchKey *string `json:"SearchKey,omitnil" name:"SearchKey"`

	// Queries by project ID
	ProjectIds []*int64 `json:"ProjectIds,omitnil" name:"ProjectIds"`

	// Whether to search by VPC
	IsFilterVpc *bool `json:"IsFilterVpc,omitnil" name:"IsFilterVpc"`

	// VPC ID, which is valid when `IsFilterVpc` is 1
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// VPC subnet ID, which is valid when `IsFilterVpc` is 1
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// Sort by field. Valid values: projectId, createtime, instancename
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// Sorting order. Valid values: desc, asc
	OrderByType *string `json:"OrderByType,omitnil" name:"OrderByType"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Queries by `OriginSerialId`
	OriginSerialIds []*string `json:"OriginSerialIds,omitnil" name:"OriginSerialIds"`

	// Identifies whether to use the `ExclusterType` field. false: no, true: yes
	IsFilterExcluster *bool `json:"IsFilterExcluster,omitnil" name:"IsFilterExcluster"`

	// Instance cluster type. 1: non-dedicated cluster, 2: dedicated cluster, 0: all
	ExclusterType *int64 `json:"ExclusterType,omitnil" name:"ExclusterType"`

	// Filters instances by dedicated cluster ID in the format of `dbdc-4ih6uct9`
	ExclusterIds []*string `json:"ExclusterIds,omitnil" name:"ExclusterIds"`

	// Tag key used in queries
	TagKeys []*string `json:"TagKeys,omitnil" name:"TagKeys"`

	// Instance types used in filtering. Valid values: 1 (dedicated instance), 2 (primary instance), 3 (disaster recovery instance). Multiple values should be separated by commas.
	FilterInstanceType *string `json:"FilterInstanceType,omitnil" name:"FilterInstanceType"`

	// Use this filter to include instances in specific statuses
	Status []*int64 `json:"Status,omitnil" name:"Status"`

	// Use this filter to exclude instances in specific statuses
	ExcludeStatus []*int64 `json:"ExcludeStatus,omitnil" name:"ExcludeStatus"`
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Queries by instance ID or IDs. Instance ID is in the format of `tdsql-ow728lmc`. Up to 100 instances can be queried in one request.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Search field name. Valid values: instancename (search by instance name), vip (search by private IP), all (search by instance ID, instance name, and private IP).
	SearchName *string `json:"SearchName,omitnil" name:"SearchName"`

	// Search keyword. Fuzzy search is supported. Multiple keywords should be separated by line breaks (`\n`).
	SearchKey *string `json:"SearchKey,omitnil" name:"SearchKey"`

	// Queries by project ID
	ProjectIds []*int64 `json:"ProjectIds,omitnil" name:"ProjectIds"`

	// Whether to search by VPC
	IsFilterVpc *bool `json:"IsFilterVpc,omitnil" name:"IsFilterVpc"`

	// VPC ID, which is valid when `IsFilterVpc` is 1
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// VPC subnet ID, which is valid when `IsFilterVpc` is 1
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// Sort by field. Valid values: projectId, createtime, instancename
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// Sorting order. Valid values: desc, asc
	OrderByType *string `json:"OrderByType,omitnil" name:"OrderByType"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Queries by `OriginSerialId`
	OriginSerialIds []*string `json:"OriginSerialIds,omitnil" name:"OriginSerialIds"`

	// Identifies whether to use the `ExclusterType` field. false: no, true: yes
	IsFilterExcluster *bool `json:"IsFilterExcluster,omitnil" name:"IsFilterExcluster"`

	// Instance cluster type. 1: non-dedicated cluster, 2: dedicated cluster, 0: all
	ExclusterType *int64 `json:"ExclusterType,omitnil" name:"ExclusterType"`

	// Filters instances by dedicated cluster ID in the format of `dbdc-4ih6uct9`
	ExclusterIds []*string `json:"ExclusterIds,omitnil" name:"ExclusterIds"`

	// Tag key used in queries
	TagKeys []*string `json:"TagKeys,omitnil" name:"TagKeys"`

	// Instance types used in filtering. Valid values: 1 (dedicated instance), 2 (primary instance), 3 (disaster recovery instance). Multiple values should be separated by commas.
	FilterInstanceType *string `json:"FilterInstanceType,omitnil" name:"FilterInstanceType"`

	// Use this filter to include instances in specific statuses
	Status []*int64 `json:"Status,omitnil" name:"Status"`

	// Use this filter to exclude instances in specific statuses
	ExcludeStatus []*int64 `json:"ExcludeStatus,omitnil" name:"ExcludeStatus"`
}

func (r *DescribeDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "SearchName")
	delete(f, "SearchKey")
	delete(f, "ProjectIds")
	delete(f, "IsFilterVpc")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OriginSerialIds")
	delete(f, "IsFilterExcluster")
	delete(f, "ExclusterType")
	delete(f, "ExclusterIds")
	delete(f, "TagKeys")
	delete(f, "FilterInstanceType")
	delete(f, "Status")
	delete(f, "ExcludeStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesResponseParams struct {
	// Number of eligible instances
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Instance details list
	Instances []*DBInstance `json:"Instances,omitnil" name:"Instances"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstancesResponseParams `json:"Response"`
}

func (r *DescribeDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBLogFilesRequestParams struct {
	// Instance ID in the format of `tdsql-ow728lmc`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Requested log type. Valid values: 1 (binlog), 2 (cold backup), 3 (errlog), 4 (slowlog).
	Type *uint64 `json:"Type,omitnil" name:"Type"`
}

type DescribeDBLogFilesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of `tdsql-ow728lmc`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Requested log type. Valid values: 1 (binlog), 2 (cold backup), 3 (errlog), 4 (slowlog).
	Type *uint64 `json:"Type,omitnil" name:"Type"`
}

func (r *DescribeDBLogFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBLogFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBLogFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBLogFilesResponseParams struct {
	// Instance ID in the format of `tdsql-ow728lmc`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Requested log type. Valid values: 1 (binlog), 2 (cold backup), 3 (errlog), 4 (slowlog).
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// Total number of requested logs
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// Information such as `uri`, `length`, and `mtime` (modification time)
	Files []*LogFileInfo `json:"Files,omitnil" name:"Files"`

	// For an instance in a VPC, this prefix plus URI can be used as the download address
	VpcPrefix *string `json:"VpcPrefix,omitnil" name:"VpcPrefix"`

	// For an instance in a common network, this prefix plus URI can be used as the download address
	NormalPrefix *string `json:"NormalPrefix,omitnil" name:"NormalPrefix"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDBLogFilesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBLogFilesResponseParams `json:"Response"`
}

func (r *DescribeDBLogFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBLogFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBParametersRequestParams struct {
	// Instance ID in the format of `tdsql-ow728lmc`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeDBParametersRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of `tdsql-ow728lmc`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeDBParametersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBParametersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBParametersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBParametersResponseParams struct {
	// Instance ID in the format of `tdsql-ow728lmc`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Requests the current parameter values of the database
	Params []*ParamDesc `json:"Params,omitnil" name:"Params"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDBParametersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBParametersResponseParams `json:"Response"`
}

func (r *DescribeDBParametersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSecurityGroupsRequestParams struct {
	// Database engine name. Valid value: `mariadb`.
	Product *string `json:"Product,omitnil" name:"Product"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Database engine name. Valid value: `mariadb`.
	Product *string `json:"Product,omitnil" name:"Product"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeDBSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSecurityGroupsResponseParams struct {
	// Security group details
	Groups []*SecurityGroup `json:"Groups,omitnil" name:"Groups"`

	// Instance VIP
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	VIP *string `json:"VIP,omitnil" name:"VIP"`

	// Instance port
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	VPort *int64 `json:"VPort,omitnil" name:"VPort"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDBSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSecurityGroupsResponseParams `json:"Response"`
}

func (r *DescribeDBSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSlowLogsRequestParams struct {
	// Instance ID in the format of `tdsql-ow728lmc`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Data entry number starting from which to return results
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of results to be returned
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Query start time in the format of 2016-07-23 14:55:20
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// Query end time in the format of 2016-08-22 14:55:20
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Specific name of the database to be queried
	Db *string `json:"Db,omitnil" name:"Db"`

	// Sorting metric. Valid values: query_time_sum, query_count
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// Sorting order. Valid values: desc, asc
	OrderByType *string `json:"OrderByType,omitnil" name:"OrderByType"`

	// Query slow queries from either the primary or the replica. Valid values: 0 (primary), 1 (replica)
	Slave *int64 `json:"Slave,omitnil" name:"Slave"`
}

type DescribeDBSlowLogsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of `tdsql-ow728lmc`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Data entry number starting from which to return results
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of results to be returned
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Query start time in the format of 2016-07-23 14:55:20
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// Query end time in the format of 2016-08-22 14:55:20
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Specific name of the database to be queried
	Db *string `json:"Db,omitnil" name:"Db"`

	// Sorting metric. Valid values: query_time_sum, query_count
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// Sorting order. Valid values: desc, asc
	OrderByType *string `json:"OrderByType,omitnil" name:"OrderByType"`

	// Query slow queries from either the primary or the replica. Valid values: 0 (primary), 1 (replica)
	Slave *int64 `json:"Slave,omitnil" name:"Slave"`
}

func (r *DescribeDBSlowLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSlowLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Db")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Slave")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSlowLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSlowLogsResponseParams struct {
	// Slow query log data
	Data []*SlowLogData `json:"Data,omitnil" name:"Data"`

	// Total statement lock time
	LockTimeSum *float64 `json:"LockTimeSum,omitnil" name:"LockTimeSum"`

	// Total number of statement queries
	QueryCount *int64 `json:"QueryCount,omitnil" name:"QueryCount"`

	// Total number of results
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// Total statement query time
	QueryTimeSum *float64 `json:"QueryTimeSum,omitnil" name:"QueryTimeSum"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDBSlowLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSlowLogsResponseParams `json:"Response"`
}

func (r *DescribeDBSlowLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSlowLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseObjectsRequestParams struct {
	// Instance ID in the format of `dcdbt-ow7t8lmc`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Database name, which can be obtained through the `DescribeDatabases` API.
	DbName *string `json:"DbName,omitnil" name:"DbName"`
}

type DescribeDatabaseObjectsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of `dcdbt-ow7t8lmc`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Database name, which can be obtained through the `DescribeDatabases` API.
	DbName *string `json:"DbName,omitnil" name:"DbName"`
}

func (r *DescribeDatabaseObjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseObjectsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DbName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabaseObjectsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseObjectsResponseParams struct {
	// Passed through from input parameters.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Database name.
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// Table list.
	Tables []*DatabaseTable `json:"Tables,omitnil" name:"Tables"`

	// View list.
	Views []*DatabaseView `json:"Views,omitnil" name:"Views"`

	// Stored procedure list.
	Procs []*DatabaseProcedure `json:"Procs,omitnil" name:"Procs"`

	// Function list.
	Funcs []*DatabaseFunction `json:"Funcs,omitnil" name:"Funcs"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDatabaseObjectsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatabaseObjectsResponseParams `json:"Response"`
}

func (r *DescribeDatabaseObjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseObjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseTableRequestParams struct {
	// Instance ID in the format of `dcdbt-ow7t8lmc`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Database name, which can be obtained through the `DescribeDatabases` API.
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// Table name, which can be obtained through the `DescribeDatabaseObjects` API.
	Table *string `json:"Table,omitnil" name:"Table"`
}

type DescribeDatabaseTableRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of `dcdbt-ow7t8lmc`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Database name, which can be obtained through the `DescribeDatabases` API.
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// Table name, which can be obtained through the `DescribeDatabaseObjects` API.
	Table *string `json:"Table,omitnil" name:"Table"`
}

func (r *DescribeDatabaseTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DbName")
	delete(f, "Table")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabaseTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseTableResponseParams struct {
	// Instance name.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Database name.
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// Table name.
	Table *string `json:"Table,omitnil" name:"Table"`

	// Column information.
	Cols []*TableColumn `json:"Cols,omitnil" name:"Cols"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDatabaseTableResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatabaseTableResponseParams `json:"Response"`
}

func (r *DescribeDatabaseTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabasesRequestParams struct {
	// Instance ID in the format of `dcdbt-ow7t8lmc`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeDatabasesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of `dcdbt-ow7t8lmc`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeDatabasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabasesResponseParams struct {
	// The database list of this instance.
	Databases []*Database `json:"Databases,omitnil" name:"Databases"`

	// Passed through from input parameters.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatabasesResponseParams `json:"Response"`
}

func (r *DescribeDatabasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDcnDetailRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeDcnDetailRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeDcnDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDcnDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDcnDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDcnDetailResponseParams struct {
	// DCN synchronization details
	DcnDetails []*DcnDetailItem `json:"DcnDetails,omitnil" name:"DcnDetails"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDcnDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDcnDetailResponseParams `json:"Response"`
}

func (r *DescribeDcnDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDcnDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileDownloadUrlRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Unsigned file path
	FilePath *string `json:"FilePath,omitnil" name:"FilePath"`
}

type DescribeFileDownloadUrlRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Unsigned file path
	FilePath *string `json:"FilePath,omitnil" name:"FilePath"`
}

func (r *DescribeFileDownloadUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileDownloadUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FilePath")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileDownloadUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileDownloadUrlResponseParams struct {
	// Signed download URL
	PreSignedUrl *string `json:"PreSignedUrl,omitnil" name:"PreSignedUrl"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeFileDownloadUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFileDownloadUrlResponseParams `json:"Response"`
}

func (r *DescribeFileDownloadUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileDownloadUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceNodeInfoRequestParams struct {
	// Instance ID, such as tdsql-6ltok4u9
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The maximum number of results returned at a time. By default, there is no upper limit to this value, that is, all results can be returned.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset of the returned results. Default value: `0`.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeInstanceNodeInfoRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, such as tdsql-6ltok4u9
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The maximum number of results returned at a time. By default, there is no upper limit to this value, that is, all results can be returned.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset of the returned results. Default value: `0`.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeInstanceNodeInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodeInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceNodeInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceNodeInfoResponseParams struct {
	// Total number of nodes
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Node information
	NodesInfo []*NodeInfo `json:"NodesInfo,omitnil" name:"NodesInfo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstanceNodeInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceNodeInfoResponseParams `json:"Response"`
}

func (r *DescribeInstanceNodeInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodeInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogFileRetentionPeriodRequestParams struct {
	// Instance ID in the format of `tdsql-ow728lmc`
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeLogFileRetentionPeriodRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of `tdsql-ow728lmc`
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeLogFileRetentionPeriodRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogFileRetentionPeriodRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogFileRetentionPeriodRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogFileRetentionPeriodResponseParams struct {
	// Instance ID in the format of `tdsql-ow728lmc`
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Backup log retention days
	Days *uint64 `json:"Days,omitnil" name:"Days"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeLogFileRetentionPeriodResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogFileRetentionPeriodResponseParams `json:"Response"`
}

func (r *DescribeLogFileRetentionPeriodResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogFileRetentionPeriodResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrdersRequestParams struct {
	// List of long order numbers to be queried, which are returned for the APIs for creating, renewing, or scaling instances.
	DealNames []*string `json:"DealNames,omitnil" name:"DealNames"`
}

type DescribeOrdersRequest struct {
	*tchttp.BaseRequest
	
	// List of long order numbers to be queried, which are returned for the APIs for creating, renewing, or scaling instances.
	DealNames []*string `json:"DealNames,omitnil" name:"DealNames"`
}

func (r *DescribeOrdersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrdersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DealNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrdersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrdersResponseParams struct {
	// Returned number of orders.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Order information list.
	Deals []*Deal `json:"Deals,omitnil" name:"Deals"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeOrdersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrdersResponseParams `json:"Response"`
}

func (r *DescribeOrdersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrdersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePriceRequestParams struct {
	// AZ ID of the purchased instance.
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// Number of instance nodes, which can be obtained 
	//  by querying the instance specification through the `DescribeDBInstanceSpecs` API.
	NodeCount *int64 `json:"NodeCount,omitnil" name:"NodeCount"`

	// Memory size in GB, which can be obtained 
	//  by querying the instance specification through the `DescribeDBInstanceSpecs` API.
	Memory *int64 `json:"Memory,omitnil" name:"Memory"`

	// Storage capacity in GB. The maximum and minimum storage space can be obtained 
	//  by querying instance specification through the `DescribeDBInstanceSpecs` API.
	Storage *int64 `json:"Storage,omitnil" name:"Storage"`

	// Purchase period in months
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// The number of instances to be purchased. Only one instance is queried for price by default.
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// Billing type. Valid values: `postpaid` (pay-as-you-go), `prepaid` (monthly subscription).
	Paymode *string `json:"Paymode,omitnil" name:"Paymode"`

	// Price unit. Valid values:   
	// `* pent` (cent), 
	// `* microPent` (microcent).
	AmountUnit *string `json:"AmountUnit,omitnil" name:"AmountUnit"`
}

type DescribePriceRequest struct {
	*tchttp.BaseRequest
	
	// AZ ID of the purchased instance.
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// Number of instance nodes, which can be obtained 
	//  by querying the instance specification through the `DescribeDBInstanceSpecs` API.
	NodeCount *int64 `json:"NodeCount,omitnil" name:"NodeCount"`

	// Memory size in GB, which can be obtained 
	//  by querying the instance specification through the `DescribeDBInstanceSpecs` API.
	Memory *int64 `json:"Memory,omitnil" name:"Memory"`

	// Storage capacity in GB. The maximum and minimum storage space can be obtained 
	//  by querying instance specification through the `DescribeDBInstanceSpecs` API.
	Storage *int64 `json:"Storage,omitnil" name:"Storage"`

	// Purchase period in months
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// The number of instances to be purchased. Only one instance is queried for price by default.
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// Billing type. Valid values: `postpaid` (pay-as-you-go), `prepaid` (monthly subscription).
	Paymode *string `json:"Paymode,omitnil" name:"Paymode"`

	// Price unit. Valid values:   
	// `* pent` (cent), 
	// `* microPent` (microcent).
	AmountUnit *string `json:"AmountUnit,omitnil" name:"AmountUnit"`
}

func (r *DescribePriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePriceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "NodeCount")
	delete(f, "Memory")
	delete(f, "Storage")
	delete(f, "Period")
	delete(f, "Count")
	delete(f, "Paymode")
	delete(f, "AmountUnit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePriceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePriceResponseParams struct {
	// Original price  
	// * Unit: Cent (default). If the request parameter contains `AmountUnit`, see `AmountUnit` description.
	// * Currency: CNY (Chinese site), USD (international site)
	OriginalPrice *int64 `json:"OriginalPrice,omitnil" name:"OriginalPrice"`

	// The actual price may be different from the original price due to discounts. 
	// * Unit: Cent (default). If the request parameter contains `AmountUnit`, see `AmountUnit` description.
	// * Currency: CNY (Chinese site), USD (international site)
	Price *int64 `json:"Price,omitnil" name:"Price"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribePriceResponse struct {
	*tchttp.BaseResponse
	Response *DescribePriceResponseParams `json:"Response"`
}

func (r *DescribePriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectSecurityGroupsRequestParams struct {
	// Database engine name. Valid value: `mariadb`.
	Product *string `json:"Product,omitnil" name:"Product"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`
}

type DescribeProjectSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Database engine name. Valid value: `mariadb`.
	Product *string `json:"Product,omitnil" name:"Product"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`
}

func (r *DescribeProjectSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectSecurityGroupsResponseParams struct {
	// Security group details
	Groups []*SecurityGroup `json:"Groups,omitnil" name:"Groups"`

	// Total number of security groups.
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeProjectSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectSecurityGroupsResponseParams `json:"Response"`
}

func (r *DescribeProjectSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyDBInstanceRequestParams struct {
	// Instance ID in the format of “tdsqlshard-c1nl9rpv”. It is the same as the instance ID displayed in the TencentDB for MariaDB console.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DestroyDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of “tdsqlshard-c1nl9rpv”. It is the same as the instance ID displayed in the TencentDB for MariaDB console.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DestroyDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyDBInstanceResponseParams struct {
	// Instance ID, which is the same as the request parameter `InstanceId`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Async task ID, which can be used in the [DescribeFlow](https://intl.cloud.tencent.com/document/product/557/56485?from_cn_redirect=1) API to query the async task result.
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DestroyDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DestroyDBInstanceResponseParams `json:"Response"`
}

func (r *DestroyDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyHourDBInstanceRequestParams struct {
	// Instance ID in the format of tdsql-avw0207d. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DestroyHourDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of tdsql-avw0207d. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DestroyHourDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyHourDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyHourDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyHourDBInstanceResponseParams struct {
	// Async task ID, which can be used in the [DescribeFlow](https://intl.cloud.tencent.com/document/product/237/16177?from_cn_redirect=1) API to query the async task result.
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// Instance ID, which is the same as the request parameter `InstanceId`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DestroyHourDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DestroyHourDBInstanceResponseParams `json:"Response"`
}

func (r *DestroyHourDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyHourDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateSecurityGroupsRequestParams struct {
	// Database engine name. Valid value: `mariadb`.
	Product *string `json:"Product,omitnil" name:"Product"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// Instance ID list, which is an array of one or more instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Database engine name. Valid value: `mariadb`.
	Product *string `json:"Product,omitnil" name:"Product"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// Instance ID list, which is an array of one or more instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

func (r *DisassociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "SecurityGroupId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateSecurityGroupsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DisassociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateSecurityGroupsResponseParams `json:"Response"`
}

func (r *DisassociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FunctionPrivilege struct {
	// Database name
	Database *string `json:"Database,omitnil" name:"Database"`

	// Function name
	FunctionName *string `json:"FunctionName,omitnil" name:"FunctionName"`

	// Permission information
	Privileges []*string `json:"Privileges,omitnil" name:"Privileges"`
}

// Predefined struct for user
type GrantAccountPrivilegesRequestParams struct {
	// Instance ID, which is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Login username.
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Access host allowed for user. An account is uniquely identified by username and host.
	Host *string `json:"Host,omitnil" name:"Host"`

	// Database name. `\*` indicates that global permissions will be set (i.e., `\*.\*`), in which case the `Type` and `Object ` parameters will be ignored. If `DbName` is not `\*`, the input parameter `Type` is required.
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// Global permission. Valid values: `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, `REFERENCES`, `INDEX`, `ALTER`, `CREATE TEMPORARY TABLES`, `LOCK TABLES`, `EXECUTE`, `CREATE VIEW`, `SHOW VIEW`, `CREATE ROUTINE`, `ALTER ROUTINE`, `EVENT`, `TRIGGER`, `SHOW DATABASES`, `REPLICATION CLIENT`, `REPLICATION SLAVE`. 
	// Database permission. Valid values: `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, `REFERENCES`, `INDEX`, `ALTER`, `CREATE TEMPORARY TABLES`, `LOCK TABLES`, `EXECUTE`, `CREATE VIEW`, `SHOW VIEW`, `CREATE ROUTINE`, `ALTER ROUTINE`, `EVENT`, `TRIGGER`. 
	// Table/View permission. Valid values: `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, `REFERENCES`, `INDEX`, `ALTER`, `CREATE VIEW`, `SHOW VIEW`, `TRIGGER`. 
	// Stored procedure/function permission. Valid values: `ALTER ROUTINE`, `EXECUTE`. 
	// Field permission. Valid values: `INSERT`, `REFERENCES`, `SELECT`, `UPDATE`.
	Privileges []*string `json:"Privileges,omitnil" name:"Privileges"`

	// Type. Valid values: table, view, proc, func, \*. If `DbName` is a specific database name and `Type` is `\*`, the permissions of the database will be set (i.e., `db.\*`), in which case the `Object` parameter will be ignored
	Type *string `json:"Type,omitnil" name:"Type"`

	// Type name. For example, if `Type` is `table`, `Object` indicates a specific table name; if both `DbName` and `Type` are specific names, it indicates a specific object name and cannot be `\*` or empty
	Object *string `json:"Object,omitnil" name:"Object"`

	// If `Type` is `table` and `ColName` is `\*`, the permissions will be granted to the table; if `ColName` is a specific field name, the permissions will be granted to the field
	ColName *string `json:"ColName,omitnil" name:"ColName"`
}

type GrantAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Login username.
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Access host allowed for user. An account is uniquely identified by username and host.
	Host *string `json:"Host,omitnil" name:"Host"`

	// Database name. `\*` indicates that global permissions will be set (i.e., `\*.\*`), in which case the `Type` and `Object ` parameters will be ignored. If `DbName` is not `\*`, the input parameter `Type` is required.
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// Global permission. Valid values: `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, `REFERENCES`, `INDEX`, `ALTER`, `CREATE TEMPORARY TABLES`, `LOCK TABLES`, `EXECUTE`, `CREATE VIEW`, `SHOW VIEW`, `CREATE ROUTINE`, `ALTER ROUTINE`, `EVENT`, `TRIGGER`, `SHOW DATABASES`, `REPLICATION CLIENT`, `REPLICATION SLAVE`. 
	// Database permission. Valid values: `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, `REFERENCES`, `INDEX`, `ALTER`, `CREATE TEMPORARY TABLES`, `LOCK TABLES`, `EXECUTE`, `CREATE VIEW`, `SHOW VIEW`, `CREATE ROUTINE`, `ALTER ROUTINE`, `EVENT`, `TRIGGER`. 
	// Table/View permission. Valid values: `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, `REFERENCES`, `INDEX`, `ALTER`, `CREATE VIEW`, `SHOW VIEW`, `TRIGGER`. 
	// Stored procedure/function permission. Valid values: `ALTER ROUTINE`, `EXECUTE`. 
	// Field permission. Valid values: `INSERT`, `REFERENCES`, `SELECT`, `UPDATE`.
	Privileges []*string `json:"Privileges,omitnil" name:"Privileges"`

	// Type. Valid values: table, view, proc, func, \*. If `DbName` is a specific database name and `Type` is `\*`, the permissions of the database will be set (i.e., `db.\*`), in which case the `Object` parameter will be ignored
	Type *string `json:"Type,omitnil" name:"Type"`

	// Type name. For example, if `Type` is `table`, `Object` indicates a specific table name; if both `DbName` and `Type` are specific names, it indicates a specific object name and cannot be `\*` or empty
	Object *string `json:"Object,omitnil" name:"Object"`

	// If `Type` is `table` and `ColName` is `\*`, the permissions will be granted to the table; if `ColName` is a specific field name, the permissions will be granted to the field
	ColName *string `json:"ColName,omitnil" name:"ColName"`
}

func (r *GrantAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GrantAccountPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Host")
	delete(f, "DbName")
	delete(f, "Privileges")
	delete(f, "Type")
	delete(f, "Object")
	delete(f, "ColName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GrantAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GrantAccountPrivilegesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GrantAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *GrantAccountPrivilegesResponseParams `json:"Response"`
}

func (r *GrantAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GrantAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceBackupFileItem struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Instance status
	InstanceStatus *int64 `json:"InstanceStatus,omitnil" name:"InstanceStatus"`

	// Shard ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ShardId *string `json:"ShardId,omitnil" name:"ShardId"`

	// File path
	FilePath *string `json:"FilePath,omitnil" name:"FilePath"`

	// File name
	FileName *string `json:"FileName,omitnil" name:"FileName"`

	// File size
	FileSize *int64 `json:"FileSize,omitnil" name:"FileSize"`

	// Backup type. Valid values: `Data` (data backup), `Binlog` (Binlog backup), `Errlog` (error log), `Slowlog` (slow log).
	BackupType *string `json:"BackupType,omitnil" name:"BackupType"`

	// Manual backup. Valid values: `0` (no), `1` (yes).
	ManualBackup *int64 `json:"ManualBackup,omitnil" name:"ManualBackup"`

	// Backup start time
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// Backup end time
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

// Predefined struct for user
type IsolateDBInstanceRequestParams struct {
	// Instance ID in the format of `tdsql-dasjkhd`. It is the same as the instance ID displayed on the TencentDB console. You can use the `DescribeDBInstances` API to query the ID, which is the value of the `InstanceId` output parameter.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

type IsolateDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of `tdsql-dasjkhd`. It is the same as the instance ID displayed on the TencentDB console. You can use the `DescribeDBInstances` API to query the ID, which is the value of the `InstanceId` output parameter.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

func (r *IsolateDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateDBInstanceResponseParams struct {
	// IDs of isolated instances
	SuccessInstanceIds []*string `json:"SuccessInstanceIds,omitnil" name:"SuccessInstanceIds"`

	// IDs of instances failed to be isolated
	FailedInstanceIds []*string `json:"FailedInstanceIds,omitnil" name:"FailedInstanceIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type IsolateDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *IsolateDBInstanceResponseParams `json:"Response"`
}

func (r *IsolateDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateDedicatedDBInstanceRequestParams struct {
	// Instance ID in the format of `tdsql-ow728lmc`
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type IsolateDedicatedDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of `tdsql-ow728lmc`
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *IsolateDedicatedDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDedicatedDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateDedicatedDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateDedicatedDBInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type IsolateDedicatedDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *IsolateDedicatedDBInstanceResponseParams `json:"Response"`
}

func (r *IsolateDedicatedDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDedicatedDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type KillSessionRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// List of session IDs
	SessionId []*int64 `json:"SessionId,omitnil" name:"SessionId"`
}

type KillSessionRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// List of session IDs
	SessionId []*int64 `json:"SessionId,omitnil" name:"SessionId"`
}

func (r *KillSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "KillSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type KillSessionResponseParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type KillSessionResponse struct {
	*tchttp.BaseResponse
	Response *KillSessionResponseParams `json:"Response"`
}

func (r *KillSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogFileInfo struct {
	// Last modified time of log
	Mtime *uint64 `json:"Mtime,omitnil" name:"Mtime"`

	// File length
	Length *uint64 `json:"Length,omitnil" name:"Length"`

	// Uniform resource identifier (URI) used during log download
	Uri *string `json:"Uri,omitnil" name:"Uri"`

	// Filename
	FileName *string `json:"FileName,omitnil" name:"FileName"`
}

// Predefined struct for user
type ModifyAccountDescriptionRequestParams struct {
	// Instance ID, which is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Login username.
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Access host allowed for user. An account is uniquely identified by username and host.
	Host *string `json:"Host,omitnil" name:"Host"`

	// New account remarks, which can contain 0-256 characters.
	Description *string `json:"Description,omitnil" name:"Description"`
}

type ModifyAccountDescriptionRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Login username.
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Access host allowed for user. An account is uniquely identified by username and host.
	Host *string `json:"Host,omitnil" name:"Host"`

	// New account remarks, which can contain 0-256 characters.
	Description *string `json:"Description,omitnil" name:"Description"`
}

func (r *ModifyAccountDescriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountDescriptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Host")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountDescriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountDescriptionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAccountDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccountDescriptionResponseParams `json:"Response"`
}

func (r *ModifyAccountDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountDescriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountPrivilegesRequestParams struct {
	// Instance ID in the format of tdsql-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Database account, including username and host address.
	Accounts []*Account `json:"Accounts,omitnil" name:"Accounts"`

	// Global permission. Valid values of `GlobalPrivileges`: `"SELECT"`, `"INSERT"`, `"UPDATE"`, `"DELETE"`, `"CREATE"`, `"PROCESS"`, `"DROP"`, `"REFERENCES"`, `"INDEX"`, `"ALTER"`, `"SHOW DATABASES"`, `"CREATE TEMPORARY TABLES"`, `"LOCK TABLES"`, `"EXECUTE"`, `"CREATE VIEW"`, `"SHOW VIEW"`, `"CREATE ROUTINE"`, `"ALTER ROUTINE"`, `"EVENT"`, `"TRIGGER"`.
	// Note: if the parameter is left empty, no change will be made to the granted global permissions. To clear the granted global permissions, set the parameter to an empty array.
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitnil" name:"GlobalPrivileges"`

	// Database permission. Valid values of `Privileges`: `"SELECT"`, `"INSERT"`, `"UPDATE"`, `"DELETE"`, `"CREATE"`, `"DROP"`, `"REFERENCES"`, `"INDEX"`, `"ALTER"`, `"CREATE TEMPORARY TABLES"`, `"LOCK TABLES"`, `"EXECUTE"`, `"CREATE VIEW"`, `"SHOW VIEW"`, `"CREATE ROUTINE"`, `"ALTER ROUTINE"`, `"EVENT"`, `"TRIGGER"`.
	// Note: if the parameter is left empty, no change will be made to the granted database permissions. To clear the granted database permissions, set `Privileges` to an empty array.
	DatabasePrivileges []*DatabasePrivilege `json:"DatabasePrivileges,omitnil" name:"DatabasePrivileges"`

	// Database table permission. Valid values of `Privileges`: `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, `REFERENCES`, `INDEX`, `ALTER`, `CREATE VIEW`, `SHOW VIEW`, `TRIGGER`.
	// Note: if the parameter is not passed in, no change will be made to the granted table permissions. To clear the granted table permissions, set `Privileges` to an empty array.
	TablePrivileges []*TablePrivilege `json:"TablePrivileges,omitnil" name:"TablePrivileges"`

	// Column permission. Valid values of `Privileges`: `"SELECT"`, `"INSERT"`, `"UPDATE"`, `"REFERENCES"`.
	// Note: if the parameter is left empty, no change will be made to the granted column permissions. To clear the granted column permissions, set `Privileges` to an empty array.
	ColumnPrivileges []*ColumnPrivilege `json:"ColumnPrivileges,omitnil" name:"ColumnPrivileges"`

	// Database view permission. Valid values of `Privileges`: `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, `REFERENCES`, `INDEX`, `ALTER`, `CREATE VIEW`, `SHOW VIEW`, `TRIGGER`.
	// Note: if the parameter is not passed in, no change will be made to the granted view permissions. To clear the granted view permissions, set `Privileges` to an empty array.
	ViewPrivileges []*ViewPrivileges `json:"ViewPrivileges,omitnil" name:"ViewPrivileges"`

	// Database function permissions. Valid values of `Privileges`: `ALTER ROUTINE`, `EXECUTE`.
	// Note: if the parameter is not passed in, no change will be made to the granted function permissions. To clear the granted function permissions, set `Privileges` to an empty array.
	FunctionPrivileges []*FunctionPrivilege `json:"FunctionPrivileges,omitnil" name:"FunctionPrivileges"`

	// Database stored procedure permission. Valid values of `Privileges`: `ALTER ROUTINE`, `EXECUTE`.
	// Note: if the parameter is not passed in, no change will be made to the granted stored procedure permissions. To clear the granted stored procedure permissions, set `Privileges` to an empty array.
	ProcedurePrivileges []*ProcedurePrivilege `json:"ProcedurePrivileges,omitnil" name:"ProcedurePrivileges"`
}

type ModifyAccountPrivilegesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of tdsql-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Database account, including username and host address.
	Accounts []*Account `json:"Accounts,omitnil" name:"Accounts"`

	// Global permission. Valid values of `GlobalPrivileges`: `"SELECT"`, `"INSERT"`, `"UPDATE"`, `"DELETE"`, `"CREATE"`, `"PROCESS"`, `"DROP"`, `"REFERENCES"`, `"INDEX"`, `"ALTER"`, `"SHOW DATABASES"`, `"CREATE TEMPORARY TABLES"`, `"LOCK TABLES"`, `"EXECUTE"`, `"CREATE VIEW"`, `"SHOW VIEW"`, `"CREATE ROUTINE"`, `"ALTER ROUTINE"`, `"EVENT"`, `"TRIGGER"`.
	// Note: if the parameter is left empty, no change will be made to the granted global permissions. To clear the granted global permissions, set the parameter to an empty array.
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitnil" name:"GlobalPrivileges"`

	// Database permission. Valid values of `Privileges`: `"SELECT"`, `"INSERT"`, `"UPDATE"`, `"DELETE"`, `"CREATE"`, `"DROP"`, `"REFERENCES"`, `"INDEX"`, `"ALTER"`, `"CREATE TEMPORARY TABLES"`, `"LOCK TABLES"`, `"EXECUTE"`, `"CREATE VIEW"`, `"SHOW VIEW"`, `"CREATE ROUTINE"`, `"ALTER ROUTINE"`, `"EVENT"`, `"TRIGGER"`.
	// Note: if the parameter is left empty, no change will be made to the granted database permissions. To clear the granted database permissions, set `Privileges` to an empty array.
	DatabasePrivileges []*DatabasePrivilege `json:"DatabasePrivileges,omitnil" name:"DatabasePrivileges"`

	// Database table permission. Valid values of `Privileges`: `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, `REFERENCES`, `INDEX`, `ALTER`, `CREATE VIEW`, `SHOW VIEW`, `TRIGGER`.
	// Note: if the parameter is not passed in, no change will be made to the granted table permissions. To clear the granted table permissions, set `Privileges` to an empty array.
	TablePrivileges []*TablePrivilege `json:"TablePrivileges,omitnil" name:"TablePrivileges"`

	// Column permission. Valid values of `Privileges`: `"SELECT"`, `"INSERT"`, `"UPDATE"`, `"REFERENCES"`.
	// Note: if the parameter is left empty, no change will be made to the granted column permissions. To clear the granted column permissions, set `Privileges` to an empty array.
	ColumnPrivileges []*ColumnPrivilege `json:"ColumnPrivileges,omitnil" name:"ColumnPrivileges"`

	// Database view permission. Valid values of `Privileges`: `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, `REFERENCES`, `INDEX`, `ALTER`, `CREATE VIEW`, `SHOW VIEW`, `TRIGGER`.
	// Note: if the parameter is not passed in, no change will be made to the granted view permissions. To clear the granted view permissions, set `Privileges` to an empty array.
	ViewPrivileges []*ViewPrivileges `json:"ViewPrivileges,omitnil" name:"ViewPrivileges"`

	// Database function permissions. Valid values of `Privileges`: `ALTER ROUTINE`, `EXECUTE`.
	// Note: if the parameter is not passed in, no change will be made to the granted function permissions. To clear the granted function permissions, set `Privileges` to an empty array.
	FunctionPrivileges []*FunctionPrivilege `json:"FunctionPrivileges,omitnil" name:"FunctionPrivileges"`

	// Database stored procedure permission. Valid values of `Privileges`: `ALTER ROUTINE`, `EXECUTE`.
	// Note: if the parameter is not passed in, no change will be made to the granted stored procedure permissions. To clear the granted stored procedure permissions, set `Privileges` to an empty array.
	ProcedurePrivileges []*ProcedurePrivilege `json:"ProcedurePrivileges,omitnil" name:"ProcedurePrivileges"`
}

func (r *ModifyAccountPrivilegesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountPrivilegesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Accounts")
	delete(f, "GlobalPrivileges")
	delete(f, "DatabasePrivileges")
	delete(f, "TablePrivileges")
	delete(f, "ColumnPrivileges")
	delete(f, "ViewPrivileges")
	delete(f, "FunctionPrivileges")
	delete(f, "ProcedurePrivileges")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountPrivilegesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountPrivilegesResponseParams struct {
	// Async task ID, which can be used in the [DescribeFlow](https://intl.cloud.tencent.com/document/product/237/16177?from_cn_redirect=1) API to query the async task result.
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAccountPrivilegesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccountPrivilegesResponseParams `json:"Response"`
}

func (r *ModifyAccountPrivilegesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountPrivilegesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBEncryptAttributesRequestParams struct {
	// Instance ID in the format of `tdsql-ow728lmc`
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Whether to enable the data encryption (Once enabled, it can’t be disabled). Valid values: `1` (Yes), `0` (No. Default）
	EncryptEnabled *int64 `json:"EncryptEnabled,omitnil" name:"EncryptEnabled"`
}

type ModifyDBEncryptAttributesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of `tdsql-ow728lmc`
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Whether to enable the data encryption (Once enabled, it can’t be disabled). Valid values: `1` (Yes), `0` (No. Default）
	EncryptEnabled *int64 `json:"EncryptEnabled,omitnil" name:"EncryptEnabled"`
}

func (r *ModifyDBEncryptAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBEncryptAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "EncryptEnabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBEncryptAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBEncryptAttributesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDBEncryptAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBEncryptAttributesResponseParams `json:"Response"`
}

func (r *ModifyDBEncryptAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBEncryptAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstancesProjectRequestParams struct {
	// List of IDs of instances to be modified. The ID is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// ID of the project to be assigned, which can be obtained through the `DescribeProjects` API.
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`
}

type ModifyDBInstancesProjectRequest struct {
	*tchttp.BaseRequest
	
	// List of IDs of instances to be modified. The ID is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// ID of the project to be assigned, which can be obtained through the `DescribeProjects` API.
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`
}

func (r *ModifyDBInstancesProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstancesProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstancesProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstancesProjectResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDBInstancesProjectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstancesProjectResponseParams `json:"Response"`
}

func (r *ModifyDBInstancesProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstancesProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBParametersRequestParams struct {
	// Instance ID in the format of `tdsql-ow728lmc`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Parameter list. Each element is a combination of `Param` and `Value`.
	Params []*DBParamValue `json:"Params,omitnil" name:"Params"`
}

type ModifyDBParametersRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of `tdsql-ow728lmc`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Parameter list. Each element is a combination of `Param` and `Value`.
	Params []*DBParamValue `json:"Params,omitnil" name:"Params"`
}

func (r *ModifyDBParametersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBParametersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Params")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBParametersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBParametersResponseParams struct {
	// Instance ID in the format of `tdsql-ow728lmc`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Parameter modification result
	Result []*ParamModifyResult `json:"Result,omitnil" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDBParametersResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBParametersResponseParams `json:"Response"`
}

func (r *ModifyDBParametersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBSyncModeRequestParams struct {
	// ID of the instance for which to modify the sync mode. The ID is in the format of `tdsql-ow728lmc`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Sync mode. Valid values: `0` (async), `1` (strong sync), `2` (downgradable strong sync).
	SyncMode *int64 `json:"SyncMode,omitnil" name:"SyncMode"`
}

type ModifyDBSyncModeRequest struct {
	*tchttp.BaseRequest
	
	// ID of the instance for which to modify the sync mode. The ID is in the format of `tdsql-ow728lmc`.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Sync mode. Valid values: `0` (async), `1` (strong sync), `2` (downgradable strong sync).
	SyncMode *int64 `json:"SyncMode,omitnil" name:"SyncMode"`
}

func (r *ModifyDBSyncModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBSyncModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SyncMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBSyncModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBSyncModeResponseParams struct {
	// Async task ID. The task status can be queried through the `DescribeFlow` API.
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDBSyncModeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBSyncModeResponseParams `json:"Response"`
}

func (r *ModifyDBSyncModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBSyncModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceNetworkRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// VpcId, ID of the desired VPC network.
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// SubnetId, subnet ID of the desired VPC network.
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// The field is required to specify VIP.
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// The field is required to specify VIPv6.
	Vipv6 *string `json:"Vipv6,omitnil" name:"Vipv6"`

	// VIP retention period in hours. Value range: 0-168. Default value: `24`. `0` indicates that the VIP will be released immediately, but there will be 1-minute delay.
	VipReleaseDelay *uint64 `json:"VipReleaseDelay,omitnil" name:"VipReleaseDelay"`
}

type ModifyInstanceNetworkRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// VpcId, ID of the desired VPC network.
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// SubnetId, subnet ID of the desired VPC network.
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// The field is required to specify VIP.
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// The field is required to specify VIPv6.
	Vipv6 *string `json:"Vipv6,omitnil" name:"Vipv6"`

	// VIP retention period in hours. Value range: 0-168. Default value: `24`. `0` indicates that the VIP will be released immediately, but there will be 1-minute delay.
	VipReleaseDelay *uint64 `json:"VipReleaseDelay,omitnil" name:"VipReleaseDelay"`
}

func (r *ModifyInstanceNetworkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceNetworkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Vip")
	delete(f, "Vipv6")
	delete(f, "VipReleaseDelay")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceNetworkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceNetworkResponseParams struct {
	// Async task ID, which can be used to query the task status through `DescribeFlow` API.
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyInstanceNetworkResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceNetworkResponseParams `json:"Response"`
}

func (r *ModifyInstanceNetworkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceVipRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance VIP
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// IPv6 flag
	Ipv6Flag *uint64 `json:"Ipv6Flag,omitnil" name:"Ipv6Flag"`

	// VIP retention period in hours. Value range: 0-168. Default value: `24`. `0` indicates that the VIP will be released immediately, but there will be 1-minute delay.
	VipReleaseDelay *uint64 `json:"VipReleaseDelay,omitnil" name:"VipReleaseDelay"`
}

type ModifyInstanceVipRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance VIP
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// IPv6 flag
	Ipv6Flag *uint64 `json:"Ipv6Flag,omitnil" name:"Ipv6Flag"`

	// VIP retention period in hours. Value range: 0-168. Default value: `24`. `0` indicates that the VIP will be released immediately, but there will be 1-minute delay.
	VipReleaseDelay *uint64 `json:"VipReleaseDelay,omitnil" name:"VipReleaseDelay"`
}

func (r *ModifyInstanceVipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceVipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Vip")
	delete(f, "Ipv6Flag")
	delete(f, "VipReleaseDelay")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceVipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceVipResponseParams struct {
	// Async task flow ID
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyInstanceVipResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceVipResponseParams `json:"Response"`
}

func (r *ModifyInstanceVipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceVportRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance Vport
	Vport *uint64 `json:"Vport,omitnil" name:"Vport"`
}

type ModifyInstanceVportRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance Vport
	Vport *uint64 `json:"Vport,omitnil" name:"Vport"`
}

func (r *ModifyInstanceVportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceVportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Vport")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceVportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceVportResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyInstanceVportResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceVportResponseParams `json:"Response"`
}

func (r *ModifyInstanceVportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceVportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySyncTaskAttributeRequestParams struct {
	// IDs of tasks for which to modify the attributes. The IDs can be obtained by the return value of the `DescribeSyncTasks` API. Up to 100 tasks can be operated at a time.
	TaskIds []*string `json:"TaskIds,omitnil" name:"TaskIds"`

	// Task name. You can specify any name you like, but its length cannot exceed 100 characters.
	TaskName *string `json:"TaskName,omitnil" name:"TaskName"`
}

type ModifySyncTaskAttributeRequest struct {
	*tchttp.BaseRequest
	
	// IDs of tasks for which to modify the attributes. The IDs can be obtained by the return value of the `DescribeSyncTasks` API. Up to 100 tasks can be operated at a time.
	TaskIds []*string `json:"TaskIds,omitnil" name:"TaskIds"`

	// Task name. You can specify any name you like, but its length cannot exceed 100 characters.
	TaskName *string `json:"TaskName,omitnil" name:"TaskName"`
}

func (r *ModifySyncTaskAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySyncTaskAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIds")
	delete(f, "TaskName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySyncTaskAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySyncTaskAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifySyncTaskAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifySyncTaskAttributeResponseParams `json:"Response"`
}

func (r *ModifySyncTaskAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySyncTaskAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeInfo struct {
	// Node ID
	NodeId *string `json:"NodeId,omitnil" name:"NodeId"`

	// Node role. Valid values: `master`, `slave`
	Role *string `json:"Role,omitnil" name:"Role"`
}

type ParamConstraint struct {
	// Constraint type, such as `enum` and `section`.
	Type *string `json:"Type,omitnil" name:"Type"`

	// List of valid values when constraint type is `enum`
	Enum *string `json:"Enum,omitnil" name:"Enum"`

	// Range when constraint type is `section`
	// Note: This field may return null, indicating that no valid values can be obtained.
	Range *ConstraintRange `json:"Range,omitnil" name:"Range"`

	// List of valid values when constraint type is `string`
	String *string `json:"String,omitnil" name:"String"`
}

type ParamDesc struct {
	// Parameter name
	Param *string `json:"Param,omitnil" name:"Param"`

	// Current parameter value
	Value *string `json:"Value,omitnil" name:"Value"`

	// Previously set value, which is the same as `value` after the parameter takes effect.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SetValue *string `json:"SetValue,omitnil" name:"SetValue"`

	// Default value
	Default *string `json:"Default,omitnil" name:"Default"`

	// Parameter constraint
	Constraint *ParamConstraint `json:"Constraint,omitnil" name:"Constraint"`

	// Whether a value has been set. false: no, true: yes
	HaveSetValue *bool `json:"HaveSetValue,omitnil" name:"HaveSetValue"`

	// Whether restart is required. false: no;
	// true: yes.
	NeedRestart *bool `json:"NeedRestart,omitnil" name:"NeedRestart"`
}

type ParamModifyResult struct {
	// Renames parameter
	Param *string `json:"Param,omitnil" name:"Param"`

	// Result of parameter modification. 0: success; -1: failure; -2: invalid parameter value.
	Code *int64 `json:"Code,omitnil" name:"Code"`
}

type ProcedurePrivilege struct {
	// Database name
	Database *string `json:"Database,omitnil" name:"Database"`

	// Stored procedure name
	Procedure *string `json:"Procedure,omitnil" name:"Procedure"`

	// Permission information
	Privileges []*string `json:"Privileges,omitnil" name:"Privileges"`
}

type ReservedNetResource struct {
	// VPC
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// Subnet
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// Reserved private network IP under `VpcId` and `SubnetId`
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// Port under `Vip`
	Vports []*int64 `json:"Vports,omitnil" name:"Vports"`

	// Valid hours of VIP
	RecycleTime *string `json:"RecycleTime,omitnil" name:"RecycleTime"`
}

// Predefined struct for user
type ResetAccountPasswordRequestParams struct {
	// Instance ID, which is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Login username.
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Access host allowed for user. An account is uniquely identified by username and host.
	Host *string `json:"Host,omitnil" name:"Host"`

	// New password, which can contain 6-32 letters, digits, and common symbols but not semicolons, single quotation marks, and double quotation marks.
	Password *string `json:"Password,omitnil" name:"Password"`
}

type ResetAccountPasswordRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, which is in the format of `tdsql-ow728lmc` and can be obtained through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Login username.
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// Access host allowed for user. An account is uniquely identified by username and host.
	Host *string `json:"Host,omitnil" name:"Host"`

	// New password, which can contain 6-32 letters, digits, and common symbols but not semicolons, single quotation marks, and double quotation marks.
	Password *string `json:"Password,omitnil" name:"Password"`
}

func (r *ResetAccountPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAccountPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Host")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetAccountPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetAccountPasswordResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ResetAccountPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetAccountPasswordResponseParams `json:"Response"`
}

func (r *ResetAccountPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceTag struct {
	// Tag key
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}

type SecurityGroup struct {
	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Creation time in the format of yyyy-mm-dd hh:mm:ss
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// Security group name
	SecurityGroupName *string `json:"SecurityGroupName,omitnil" name:"SecurityGroupName"`

	// Security group remarks
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitnil" name:"SecurityGroupRemark"`

	// Inbound rule
	Inbound []*SecurityGroupBound `json:"Inbound,omitnil" name:"Inbound"`

	// Outbound rule
	Outbound []*SecurityGroupBound `json:"Outbound,omitnil" name:"Outbound"`
}

type SecurityGroupBound struct {
	// Policy, which can be `ACCEPT` or `DROP`
	Action *string `json:"Action,omitnil" name:"Action"`

	// Source IP or source IP range, such as 192.168.0.0/16
	CidrIp *string `json:"CidrIp,omitnil" name:"CidrIp"`

	// Port
	PortRange *string `json:"PortRange,omitnil" name:"PortRange"`

	// Network protocol. UDP and TCP are supported.
	IpProtocol *string `json:"IpProtocol,omitnil" name:"IpProtocol"`
}

type SlowLogData struct {
	// Statement checksum for querying details
	CheckSum *string `json:"CheckSum,omitnil" name:"CheckSum"`

	// Database name
	Db *string `json:"Db,omitnil" name:"Db"`

	// Abstracted SQL statement
	FingerPrint *string `json:"FingerPrint,omitnil" name:"FingerPrint"`

	// Average lock time
	LockTimeAvg *string `json:"LockTimeAvg,omitnil" name:"LockTimeAvg"`

	// Maximum lock time
	LockTimeMax *string `json:"LockTimeMax,omitnil" name:"LockTimeMax"`

	// Minimum lock time
	LockTimeMin *string `json:"LockTimeMin,omitnil" name:"LockTimeMin"`

	// Total lock time
	LockTimeSum *string `json:"LockTimeSum,omitnil" name:"LockTimeSum"`

	// Number of queries
	QueryCount *string `json:"QueryCount,omitnil" name:"QueryCount"`

	// Average query time
	QueryTimeAvg *string `json:"QueryTimeAvg,omitnil" name:"QueryTimeAvg"`

	// Maximum query time
	QueryTimeMax *string `json:"QueryTimeMax,omitnil" name:"QueryTimeMax"`

	// Minimum query time
	QueryTimeMin *string `json:"QueryTimeMin,omitnil" name:"QueryTimeMin"`

	// Total query time
	QueryTimeSum *string `json:"QueryTimeSum,omitnil" name:"QueryTimeSum"`

	// Number of scanned rows
	RowsExaminedSum *string `json:"RowsExaminedSum,omitnil" name:"RowsExaminedSum"`

	// Number of sent rows
	RowsSentSum *string `json:"RowsSentSum,omitnil" name:"RowsSentSum"`

	// Last execution time
	TsMax *string `json:"TsMax,omitnil" name:"TsMax"`

	// First execution time
	TsMin *string `json:"TsMin,omitnil" name:"TsMin"`

	// Account
	User *string `json:"User,omitnil" name:"User"`

	// Sample SQL
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExampleSql *string `json:"ExampleSql,omitnil" name:"ExampleSql"`

	// Host address of account
	Host *string `json:"Host,omitnil" name:"Host"`
}

type TableColumn struct {
	// Column name
	Col *string `json:"Col,omitnil" name:"Col"`

	// Column type
	Type *string `json:"Type,omitnil" name:"Type"`
}

type TablePrivilege struct {
	// Database name
	Database *string `json:"Database,omitnil" name:"Database"`

	// Table name
	Table *string `json:"Table,omitnil" name:"Table"`

	// Permission information
	Privileges []*string `json:"Privileges,omitnil" name:"Privileges"`
}

// Predefined struct for user
type TerminateDedicatedDBInstanceRequestParams struct {
	// Instance ID in the format of `tdsql-ow728lmc`
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type TerminateDedicatedDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of `tdsql-ow728lmc`
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *TerminateDedicatedDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateDedicatedDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateDedicatedDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateDedicatedDBInstanceResponseParams struct {
	// Async task ID
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TerminateDedicatedDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *TerminateDedicatedDBInstanceResponseParams `json:"Response"`
}

func (r *TerminateDedicatedDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateDedicatedDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDedicatedDBInstanceRequestParams struct {
	// ID of the instance to be upgraded.  It is in the form of  `tdsql-ow728lmc`, which can be obtained by querying the instance details through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Memory size in GB,  which can be obtained through the `DescribeFenceDBInstanceSpecs` API.
	Memory *int64 `json:"Memory,omitnil" name:"Memory"`

	// Storage space size in GB.  You can obtain the disk space limits by querying instance specification through the `DescribeDBInstanceSpecs` API.
	Storage *int64 `json:"Storage,omitnil" name:"Storage"`

	// Whether to retry again when missing the switch time window. Valid values: `0` (no), `1` (yes).
	SwitchAutoRetry *int64 `json:"SwitchAutoRetry,omitnil" name:"SwitchAutoRetry"`

	// Switch start time
	SwitchStartTime *string `json:"SwitchStartTime,omitnil" name:"SwitchStartTime"`

	// Switch end time
	SwitchEndTime *string `json:"SwitchEndTime,omitnil" name:"SwitchEndTime"`
}

type UpgradeDedicatedDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// ID of the instance to be upgraded.  It is in the form of  `tdsql-ow728lmc`, which can be obtained by querying the instance details through the `DescribeDBInstances` API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Memory size in GB,  which can be obtained through the `DescribeFenceDBInstanceSpecs` API.
	Memory *int64 `json:"Memory,omitnil" name:"Memory"`

	// Storage space size in GB.  You can obtain the disk space limits by querying instance specification through the `DescribeDBInstanceSpecs` API.
	Storage *int64 `json:"Storage,omitnil" name:"Storage"`

	// Whether to retry again when missing the switch time window. Valid values: `0` (no), `1` (yes).
	SwitchAutoRetry *int64 `json:"SwitchAutoRetry,omitnil" name:"SwitchAutoRetry"`

	// Switch start time
	SwitchStartTime *string `json:"SwitchStartTime,omitnil" name:"SwitchStartTime"`

	// Switch end time
	SwitchEndTime *string `json:"SwitchEndTime,omitnil" name:"SwitchEndTime"`
}

func (r *UpgradeDedicatedDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDedicatedDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Memory")
	delete(f, "Storage")
	delete(f, "SwitchAutoRetry")
	delete(f, "SwitchStartTime")
	delete(f, "SwitchEndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeDedicatedDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDedicatedDBInstanceResponseParams struct {
	// Async task ID
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpgradeDedicatedDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeDedicatedDBInstanceResponseParams `json:"Response"`
}

func (r *UpgradeDedicatedDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDedicatedDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ViewPrivileges struct {
	// Database name
	Database *string `json:"Database,omitnil" name:"Database"`

	// View name
	View *string `json:"View,omitnil" name:"View"`

	// Permission information
	Privileges []*string `json:"Privileges,omitnil" name:"Privileges"`
}