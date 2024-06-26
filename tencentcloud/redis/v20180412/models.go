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

package v20180412

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type Account struct {
	// Instance ID 
	// Note:  This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Account name 
	// Note:  This field may return null, indicating that no valid values can be obtained.
	AccountName *string `json:"AccountName,omitnil" name:"AccountName"`

	// Account description information 
	// Note:  This field may return null, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Read/write permission policy. Valid values: `r` (read-only),  `w` (write-only),  `rw`  (read/write). 
	// Note:  This field may return null, indicating that no valid values can be obtained.
	Privilege *string `json:"Privilege,omitnil" name:"Privilege"`

	// Read-only routing policy. Valid values: `master` (master node),  `replication`  (replica node). 
	// Note:  This field may return null, indicating that no valid values can be obtained.
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitnil" name:"ReadonlyPolicy"`

	// Sub-account status. Valid values:  `1` (being changed),  `2` (valid). `4` (deleted). 
	// Note:  This field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

// Predefined struct for user
type AllocateWanAddressRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type AllocateWanAddressRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *AllocateWanAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AllocateWanAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AllocateWanAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AllocateWanAddressResponseParams struct {
	// Async task ID
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// Status of enabling public network access
	WanStatus *string `json:"WanStatus,omitnil" name:"WanStatus"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AllocateWanAddressResponse struct {
	*tchttp.BaseResponse
	Response *AllocateWanAddressResponseParams `json:"Response"`
}

func (r *AllocateWanAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AllocateWanAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyParamsTemplateRequestParams struct {
	// List of instance IDs
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// ID of the parameter template to be applied
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`
}

type ApplyParamsTemplateRequest struct {
	*tchttp.BaseRequest
	
	// List of instance IDs
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// ID of the parameter template to be applied
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`
}

func (r *ApplyParamsTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyParamsTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyParamsTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyParamsTemplateResponseParams struct {
	// Task ID
	TaskIds []*int64 `json:"TaskIds,omitnil" name:"TaskIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ApplyParamsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ApplyParamsTemplateResponseParams `json:"Response"`
}

func (r *ApplyParamsTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyParamsTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateSecurityGroupsRequestParams struct {
	// Database engine name, which is `redis` for this API.
	Product *string `json:"Product,omitnil" name:"Product"`

	// ID of the security group to be associated in the format of sg-efil73jd.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// ID(s) of the instance(s) to be associated in the format of ins-lesecurk. You can specify multiple instances.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Database engine name, which is `redis` for this API.
	Product *string `json:"Product,omitnil" name:"Product"`

	// ID of the security group to be associated in the format of sg-efil73jd.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// ID(s) of the instance(s) to be associated in the format of ins-lesecurk. You can specify multiple instances.
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

type BackupDownloadInfo struct {
	// Backup file name
	FileName *string `json:"FileName,omitnil" name:"FileName"`

	// Backup file size in bytes. If the parameter value is `0`, the backup file size is unknown.
	FileSize *uint64 `json:"FileSize,omitnil" name:"FileSize"`

	// Address (valid for six hours) used to download the backup file over the public network
	DownloadUrl *string `json:"DownloadUrl,omitnil" name:"DownloadUrl"`

	// Address (valid for six hours) used to download the backup file over the private network
	InnerDownloadUrl *string `json:"InnerDownloadUrl,omitnil" name:"InnerDownloadUrl"`
}

type BackupLimitVpcItem struct {
	// The region of the VPC that corresponds to the download address of the backup file
	Region *string `json:"Region,omitnil" name:"Region"`

	// The list of VPCs that correspond to the download addresses of the backup files
	VpcList []*string `json:"VpcList,omitnil" name:"VpcList"`
}

type BigKeyInfo struct {
	// Database
	DB *int64 `json:"DB,omitnil" name:"DB"`

	// Big key
	Key *string `json:"Key,omitnil" name:"Key"`

	// Type
	Type *string `json:"Type,omitnil" name:"Type"`

	// Size
	Size *int64 `json:"Size,omitnil" name:"Size"`

	// Data timestamp
	Updatetime *int64 `json:"Updatetime,omitnil" name:"Updatetime"`
}

type BigKeyTypeInfo struct {
	// Type
	Type *string `json:"Type,omitnil" name:"Type"`

	// Count
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// Size
	Size *int64 `json:"Size,omitnil" name:"Size"`

	// Timestamp
	Updatetime *int64 `json:"Updatetime,omitnil" name:"Updatetime"`
}

// Predefined struct for user
type ChangeInstanceRoleRequestParams struct {
	// Replication group ID
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance role. Valid values: `rw` (read-write), `r`( read-only).
	InstanceRole *string `json:"InstanceRole,omitnil" name:"InstanceRole"`
}

type ChangeInstanceRoleRequest struct {
	*tchttp.BaseRequest
	
	// Replication group ID
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance role. Valid values: `rw` (read-write), `r`( read-only).
	InstanceRole *string `json:"InstanceRole,omitnil" name:"InstanceRole"`
}

func (r *ChangeInstanceRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeInstanceRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "InstanceId")
	delete(f, "InstanceRole")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChangeInstanceRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeInstanceRoleResponseParams struct {
	// Async task ID
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ChangeInstanceRoleResponse struct {
	*tchttp.BaseResponse
	Response *ChangeInstanceRoleResponseParams `json:"Response"`
}

func (r *ChangeInstanceRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeInstanceRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeMasterInstanceRequestParams struct {
	// Replication group ID, such as `crs-rpl-m3zt****`. It is the unique identifier automatically assigned by the system when creating a replication group. Log in to the [TencentDB for Redis console](https://console.cloud.tencent.com/redis/replication), and get it in the global replication list.
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// ID of the read-only instance to be promoted to the master instance, such as `crs-xjhsdj****`. Log in to the the [TencentDB for Redis console](https://console.cloud.tencent.com/redis), and copy it in the instance list.
	// 
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Whether to promote the read-only instance to the master instance forcibly. Valid values:
	// - `true`: Yes
	// - `false`: No
	ForceSwitch *bool `json:"ForceSwitch,omitnil" name:"ForceSwitch"`
}

type ChangeMasterInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Replication group ID, such as `crs-rpl-m3zt****`. It is the unique identifier automatically assigned by the system when creating a replication group. Log in to the [TencentDB for Redis console](https://console.cloud.tencent.com/redis/replication), and get it in the global replication list.
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// ID of the read-only instance to be promoted to the master instance, such as `crs-xjhsdj****`. Log in to the the [TencentDB for Redis console](https://console.cloud.tencent.com/redis), and copy it in the instance list.
	// 
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Whether to promote the read-only instance to the master instance forcibly. Valid values:
	// - `true`: Yes
	// - `false`: No
	ForceSwitch *bool `json:"ForceSwitch,omitnil" name:"ForceSwitch"`
}

func (r *ChangeMasterInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeMasterInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "InstanceId")
	delete(f, "ForceSwitch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChangeMasterInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeMasterInstanceResponseParams struct {
	// Async task ID
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ChangeMasterInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ChangeMasterInstanceResponseParams `json:"Response"`
}

func (r *ChangeMasterInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeMasterInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeReplicaToMasterRequestParams struct {
	// ID of the specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// ID of the replica node group. You can get more ID information of the multi-AZ replica node group though the [DescribeInstanceZoneInfo](https://intl.cloud.tencent.com/document/product/239/50312?from_cn_redirect=1) API.  This parameter is not required for a single-AZ replica node group.
	GroupId *int64 `json:"GroupId,omitnil" name:"GroupId"`
}

type ChangeReplicaToMasterRequest struct {
	*tchttp.BaseRequest
	
	// ID of the specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// ID of the replica node group. You can get more ID information of the multi-AZ replica node group though the [DescribeInstanceZoneInfo](https://intl.cloud.tencent.com/document/product/239/50312?from_cn_redirect=1) API.  This parameter is not required for a single-AZ replica node group.
	GroupId *int64 `json:"GroupId,omitnil" name:"GroupId"`
}

func (r *ChangeReplicaToMasterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeReplicaToMasterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChangeReplicaToMasterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeReplicaToMasterResponseParams struct {
	// Async task ID
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ChangeReplicaToMasterResponse struct {
	*tchttp.BaseResponse
	Response *ChangeReplicaToMasterResponseParams `json:"Response"`
}

func (r *ChangeReplicaToMasterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeReplicaToMasterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CleanUpInstanceRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type CleanUpInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *CleanUpInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CleanUpInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CleanUpInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CleanUpInstanceResponseParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CleanUpInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CleanUpInstanceResponseParams `json:"Response"`
}

func (r *CleanUpInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CleanUpInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ClearInstanceRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Redis instance password (this parameter is required for password-enabled instances but not for password-free instances)
	Password *string `json:"Password,omitnil" name:"Password"`
}

type ClearInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Redis instance password (this parameter is required for password-enabled instances but not for password-free instances)
	Password *string `json:"Password,omitnil" name:"Password"`
}

func (r *ClearInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ClearInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ClearInstanceResponseParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ClearInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ClearInstanceResponseParams `json:"Response"`
}

func (r *ClearInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloneInstancesRequestParams struct {
	// The ID of the source instance to be cloned, such as "crs-xjhsdj****". Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The number of clone instances at a time
	// - The maximum number of monthly subscribed instances is 100 for each purchase.
	// - The maximum number of pay-as-you-go instances is 30 for each purchase.
	GoodsNum *uint64 `json:"GoodsNum,omitnil" name:"GoodsNum"`

	// ID of the AZ where the clone instance resides. For more information, see [Regions and AZs](https://intl.cloud.tencent.com/document/product/239/4106?from_cn_redirect=1).
	ZoneId *uint64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// Billing mode. Valid values: <ul><li>`0` (Pay-as-you-go) </li><li>`1` (Monthly subscription) </li></ul>
	BillingMode *int64 `json:"BillingMode,omitnil" name:"BillingMode"`

	// Purchase duration of an instance. <ul><li>Unit: Month</li><li>Valid values: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`, `48`, `60` (for monthly subscription mode).</li><li> Valid value: `1` (for pay-as-you-go mode).</li></ul>
	Period *uint64 `json:"Period,omitnil" name:"Period"`

	// Security group ID, which can be obtained on the <b>Security Group</b> page in the console.
	SecurityGroupIdList []*string `json:"SecurityGroupIdList,omitnil" name:"SecurityGroupIdList"`

	// Backup ID of the clone instance, which can be obtained through the [DescribeInstanceBackups](https://intl.cloud.tencent.com/document/product/239/20011?from_cn_redirect=1) API.
	BackupId *string `json:"BackupId,omitnil" name:"BackupId"`

	// Whether the clone instance supports password-free access. Valid values: <ul><li>`true` (Yes)</li><li>`false` (No. When SSL or public network is enabled). Default value: `false`.</li></ul>
	NoAuth *bool `json:"NoAuth,omitnil" name:"NoAuth"`

	// The VPC ID of the clone instance. If this parameter is not passed in, the classic network will be selected by default.
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// The VPC subnet ID to which the clone instance belongs, which is not required for the classic network.
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// Name of the clone instance. <br>Enter up to 60 letters, digits, hyphens, and underscores.</br>
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// The access password of the clone instance. <ul><li>When the input parameter <b>NoAuth</b> is <b>true</b>, this parameter is not required. </li><li>When the instance is Redis 2.8, 4.0, or 5.0, the password must contain 8–30 characters in at least two of the following types: lowercase letters, uppercase letters, digits, and special characters `()`~!@#$%^&*-+=_|{}[]:;<>,.?/` and cannot start with `/`.</li><li>When the instance is CKV 3.2, the password must and can only contain 8–30 letters and digits.</li></ul>
	Password *string `json:"Password,omitnil" name:"Password"`

	// The auto-renewal flag. Valid values <ul><li>`0`: Manual renewal (default). </li><li>`1`: Auto-renewal. </li><li>`2`: Not auto-renewal (set by user).</ul>
	AutoRenew *uint64 `json:"AutoRenew,omitnil" name:"AutoRenew"`

	// Customized port. Valid range: 1024-65535. Default value: `6379`.
	VPort *uint64 `json:"VPort,omitnil" name:"VPort"`

	// Node information of an instance. <ul><li>Currently supported type and AZ information of a node to be configured (master node or replica node) For more information, see [RedisNodeInfo](https://intl.cloud.tencent.com/document/product/239/20022?from_cn_redirect=1#RedisNodeInfo).</li><li>This parameter is not required for single-AZ deployment.</li></ul>
	NodeSet []*RedisNodeInfo `json:"NodeSet,omitnil" name:"NodeSet"`

	// Project ID. Log in to the [Redis console](https://console.cloud.tencent.com/redis#/), and find the project ID in <b>Account Center</b> > <b>Project Management</b> in the top-right corner.
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Tag to be bound for the clone instance
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil" name:"ResourceTags"`

	// The parameter template ID associated with the clone instance
	// - If this parameter is not configured, the system will automatically adapt the corresponding default template based on the selected compatible version and architecture.
	// - You can query the parameter template list of the instance to get the template ID through the [DescribeParamTemplates](https://intl.cloud.tencent.com/document/product/239/58750?from_cn_redirect=1) API.
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// The alarm policy ID of the instance to be cloned. Log in to the [Tencent Cloud Observable Platform console](https://console.cloud.tencent.com/monitor/alarm2/policy), and get the policy ID in <b>Alarm Management</b> > <b>Policy Management</b>.
	AlarmPolicyList []*string `json:"AlarmPolicyList,omitnil" name:"AlarmPolicyList"`
}

type CloneInstancesRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the source instance to be cloned, such as "crs-xjhsdj****". Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The number of clone instances at a time
	// - The maximum number of monthly subscribed instances is 100 for each purchase.
	// - The maximum number of pay-as-you-go instances is 30 for each purchase.
	GoodsNum *uint64 `json:"GoodsNum,omitnil" name:"GoodsNum"`

	// ID of the AZ where the clone instance resides. For more information, see [Regions and AZs](https://intl.cloud.tencent.com/document/product/239/4106?from_cn_redirect=1).
	ZoneId *uint64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// Billing mode. Valid values: <ul><li>`0` (Pay-as-you-go) </li><li>`1` (Monthly subscription) </li></ul>
	BillingMode *int64 `json:"BillingMode,omitnil" name:"BillingMode"`

	// Purchase duration of an instance. <ul><li>Unit: Month</li><li>Valid values: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`, `48`, `60` (for monthly subscription mode).</li><li> Valid value: `1` (for pay-as-you-go mode).</li></ul>
	Period *uint64 `json:"Period,omitnil" name:"Period"`

	// Security group ID, which can be obtained on the <b>Security Group</b> page in the console.
	SecurityGroupIdList []*string `json:"SecurityGroupIdList,omitnil" name:"SecurityGroupIdList"`

	// Backup ID of the clone instance, which can be obtained through the [DescribeInstanceBackups](https://intl.cloud.tencent.com/document/product/239/20011?from_cn_redirect=1) API.
	BackupId *string `json:"BackupId,omitnil" name:"BackupId"`

	// Whether the clone instance supports password-free access. Valid values: <ul><li>`true` (Yes)</li><li>`false` (No. When SSL or public network is enabled). Default value: `false`.</li></ul>
	NoAuth *bool `json:"NoAuth,omitnil" name:"NoAuth"`

	// The VPC ID of the clone instance. If this parameter is not passed in, the classic network will be selected by default.
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// The VPC subnet ID to which the clone instance belongs, which is not required for the classic network.
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// Name of the clone instance. <br>Enter up to 60 letters, digits, hyphens, and underscores.</br>
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// The access password of the clone instance. <ul><li>When the input parameter <b>NoAuth</b> is <b>true</b>, this parameter is not required. </li><li>When the instance is Redis 2.8, 4.0, or 5.0, the password must contain 8–30 characters in at least two of the following types: lowercase letters, uppercase letters, digits, and special characters `()`~!@#$%^&*-+=_|{}[]:;<>,.?/` and cannot start with `/`.</li><li>When the instance is CKV 3.2, the password must and can only contain 8–30 letters and digits.</li></ul>
	Password *string `json:"Password,omitnil" name:"Password"`

	// The auto-renewal flag. Valid values <ul><li>`0`: Manual renewal (default). </li><li>`1`: Auto-renewal. </li><li>`2`: Not auto-renewal (set by user).</ul>
	AutoRenew *uint64 `json:"AutoRenew,omitnil" name:"AutoRenew"`

	// Customized port. Valid range: 1024-65535. Default value: `6379`.
	VPort *uint64 `json:"VPort,omitnil" name:"VPort"`

	// Node information of an instance. <ul><li>Currently supported type and AZ information of a node to be configured (master node or replica node) For more information, see [RedisNodeInfo](https://intl.cloud.tencent.com/document/product/239/20022?from_cn_redirect=1#RedisNodeInfo).</li><li>This parameter is not required for single-AZ deployment.</li></ul>
	NodeSet []*RedisNodeInfo `json:"NodeSet,omitnil" name:"NodeSet"`

	// Project ID. Log in to the [Redis console](https://console.cloud.tencent.com/redis#/), and find the project ID in <b>Account Center</b> > <b>Project Management</b> in the top-right corner.
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Tag to be bound for the clone instance
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil" name:"ResourceTags"`

	// The parameter template ID associated with the clone instance
	// - If this parameter is not configured, the system will automatically adapt the corresponding default template based on the selected compatible version and architecture.
	// - You can query the parameter template list of the instance to get the template ID through the [DescribeParamTemplates](https://intl.cloud.tencent.com/document/product/239/58750?from_cn_redirect=1) API.
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// The alarm policy ID of the instance to be cloned. Log in to the [Tencent Cloud Observable Platform console](https://console.cloud.tencent.com/monitor/alarm2/policy), and get the policy ID in <b>Alarm Management</b> > <b>Policy Management</b>.
	AlarmPolicyList []*string `json:"AlarmPolicyList,omitnil" name:"AlarmPolicyList"`
}

func (r *CloneInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GoodsNum")
	delete(f, "ZoneId")
	delete(f, "BillingMode")
	delete(f, "Period")
	delete(f, "SecurityGroupIdList")
	delete(f, "BackupId")
	delete(f, "NoAuth")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "InstanceName")
	delete(f, "Password")
	delete(f, "AutoRenew")
	delete(f, "VPort")
	delete(f, "NodeSet")
	delete(f, "ProjectId")
	delete(f, "ResourceTags")
	delete(f, "TemplateId")
	delete(f, "AlarmPolicyList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloneInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloneInstancesResponseParams struct {
	// Request task ID
	DealId *string `json:"DealId,omitnil" name:"DealId"`

	// Clone instance ID
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CloneInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CloneInstancesResponseParams `json:"Response"`
}

func (r *CloneInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseSSLRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type CloseSSLRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *CloseSSLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseSSLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseSSLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseSSLResponseParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CloseSSLResponse struct {
	*tchttp.BaseResponse
	Response *CloseSSLResponseParams `json:"Response"`
}

func (r *CloseSSLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseSSLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CommandTake struct {
	// Command
	Cmd *string `json:"Cmd,omitnil" name:"Cmd"`

	// Duration
	Took *int64 `json:"Took,omitnil" name:"Took"`
}

// Predefined struct for user
type CreateInstanceAccountRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Sub-account name
	AccountName *string `json:"AccountName,omitnil" name:"AccountName"`

	// 1. The password must contain 8–30 characters. A password of 12 or more characters is recommended.
	// 2. It cannot start with a slash (/).
	// 3. It must contain characters in at least two of the following types:
	//     a. Lowercase letters (a–z)
	//     b. Uppercase letters (A–Z)
	//     c. Digits (0–9)
	//     d. ()`~!@#$%^&*-+=_|{}[]:;<>,.?/
	AccountPassword *string `json:"AccountPassword,omitnil" name:"AccountPassword"`

	// Routing policy. Valid values: master (master node); replication (replica node)
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitnil" name:"ReadonlyPolicy"`

	// Read/Write policy. Valid values: r (read-only); rw (read/write).
	Privilege *string `json:"Privilege,omitnil" name:"Privilege"`

	// Sub-account description information
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

type CreateInstanceAccountRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Sub-account name
	AccountName *string `json:"AccountName,omitnil" name:"AccountName"`

	// 1. The password must contain 8–30 characters. A password of 12 or more characters is recommended.
	// 2. It cannot start with a slash (/).
	// 3. It must contain characters in at least two of the following types:
	//     a. Lowercase letters (a–z)
	//     b. Uppercase letters (A–Z)
	//     c. Digits (0–9)
	//     d. ()`~!@#$%^&*-+=_|{}[]:;<>,.?/
	AccountPassword *string `json:"AccountPassword,omitnil" name:"AccountPassword"`

	// Routing policy. Valid values: master (master node); replication (replica node)
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitnil" name:"ReadonlyPolicy"`

	// Read/Write policy. Valid values: r (read-only); rw (read/write).
	Privilege *string `json:"Privilege,omitnil" name:"Privilege"`

	// Sub-account description information
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

func (r *CreateInstanceAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AccountName")
	delete(f, "AccountPassword")
	delete(f, "ReadonlyPolicy")
	delete(f, "Privilege")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceAccountResponseParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateInstanceAccountResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceAccountResponseParams `json:"Response"`
}

func (r *CreateInstanceAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstancesRequestParams struct {
	// Instance type
	// <ul><li>`2`: Redis 2.8 Memory Edition (Standard Architecture). </li><li>`3`: CKV 3.2 Memory Edition (Standard Architecture). </li><li>`4`: CKV 3.2 Memory Edition (Cluster Architecture). </li><li>`6`: Redis 4.0 Memory Edition (Standard Architecture). </li><li>`7`: Redis 4.0 Memory Edition (Cluster Architecture). </li><li>`8`: Redis 5.0 Memory Edition (Standard Architecture). </li><li>`9`: Redis 5.0 Memory Edition (Cluster Architecture). </li><li>`15`: Redis 6.2 Memory Edition (Standard Architecture). </li><li>`16`: Redis 6.2 Memory Edition (Cluster Architecture).</li></ul>
	TypeId *uint64 `json:"TypeId,omitnil" name:"TypeId"`

	// Memory capacity in MB, which must be an integer multiple of 1024. For specific specifications, query the sales specifications in all regions through the [DescribeProductInfo](https://intl.cloud.tencent.com/document/api/239/30600?from_cn_redirect=1) API.
	// - When **TypeId** is a standard architecture, **MemSize** is the total memory capacity of the instance;
	// - When **TypeId** is a cluster architecture, **MemSize** is the single-shard memory capacity.
	MemSize *uint64 `json:"MemSize,omitnil" name:"MemSize"`

	// The number of instances for each purchase. For details, query the sales specifications in all regions through the [DescribeProductInfo](https://intl.cloud.tencent.com/document/api/239/30600?from_cn_redirect=1) API.
	GoodsNum *uint64 `json:"GoodsNum,omitnil" name:"GoodsNum"`

	// The purchase duration of an instance
	// - If `BillingMode` is `1`, that is, when the billing mode is monthly subscription, you need to set this parameter to specify the duration of the purchased instance. Unit: month. Value range: [1,2,3,4,5,6,7,8,9,10,11,12,24,36].
	// - If `BillingMode` is `0`, that is, when the billing mode is pay-as-you-go, you need to set this parameter to `1`.
	Period *uint64 `json:"Period,omitnil" name:"Period"`

	// Billing mode. 0: pay-as-you-go
	BillingMode *int64 `json:"BillingMode,omitnil" name:"BillingMode"`

	// ID of the AZ where the instance resides. For more information, see [Regions and AZs](https://intl.cloud.tencent.com/document/product/239/4106?from_cn_redirect=1).
	ZoneId *uint64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// Instance access password
	// - When the input parameter `NoAuth` is `true`, it means that the instance access is set to be password-free, and the `Password` field does not need to be configured; otherwise, `Password` is a required parameter.
	// - When the instance type `TypeId` is Redis 2.8 Memory Edition (Standard Architecture), Redis 4.0, 5.0, 6.0 (regardless of architecture), the password must contains 8-30 characters in at least two of the following types: lowercase letters, uppercase letters, digits, and symbols (()`~!@#$%^&*-+=_|{}[]:;<>,.?/), and it cannot start with a slash (/).
	// - When the instance type **TypeId** is CKV 3.2 Memory Edition (regardless of architecture), the password contains 8-30 letters and digits and excludes other characters.
	Password *string `json:"Password,omitnil" name:"Password"`

	// VPC ID. If this parameter is not passed in, the classic network will be selected by default. You can query the specific VPC ID in the [VPC console](https://console.cloud.tencent.com/vpc).
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// VPC subnet ID. This parameter is not required for the classic network. You can get the specific subnet ID by querying the subnet list in the [VPC console](https://console.cloud.tencent.com/vpc).
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// Project ID. Log in to the [Redis console](https://console.cloud.tencent.com/redis#/), go to the account information menu in the top-right corner, and select **Project Management** to query the project ID.
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Auto-renewal flag
	// - `0`: Manual renewal (default).
	// - `1`: Auto-renewal.
	// - `2`: Not auto-renewal (set by user).
	AutoRenew *uint64 `json:"AutoRenew,omitnil" name:"AutoRenew"`

	// Array of security group IDs. Get the security group ID of the instance through the [DescribeInstanceSecurityGroup](https://intl.cloud.tencent.com/document/product/239/34447?from_cn_redirect=1) API.
	SecurityGroupIdList []*string `json:"SecurityGroupIdList,omitnil" name:"SecurityGroupIdList"`

	// User-defined network port. Default value: `6379`. Range: [1024,65535].
	VPort *uint64 `json:"VPort,omitnil" name:"VPort"`

	// Quantity of instance shards
	// - This parameter is not required for instances of Standard Edition.
	// - For instances of Cluster Edition, the range of shard quantity is [1, 3, 5, 8, 12, 16, 24, 32, 40, 48, 64, 80, 96, 128].
	RedisShardNum *int64 `json:"RedisShardNum,omitnil" name:"RedisShardNum"`

	// Quantity of instance replicas
	// - For Redis Memory Edition 4.0, 5.0, 6.2 (regardless of architecture), the range of replica quantity is [1,5].
	// - For Redis 2.8 Standard Edition and CKV Standard Edition, the replica quantity is `1`.
	RedisReplicasNum *int64 `json:"RedisReplicasNum,omitnil" name:"RedisReplicasNum"`

	// Whether to support read-only replicas.
	// - Redis 2.8 Standard Edition and CKV Standard Edition don’t support read-only replicas.
	// - If read-only replicas are enabled, read/write separation will be automatically enabled for an instance, with write requests routed to the master node and read requests to the replica node.
	// - To enable read-only replicas, we recommend that you create two or more replicas.
	ReplicasReadonly *bool `json:"ReplicasReadonly,omitnil" name:"ReplicasReadonly"`

	// Instance name, which can contain up to 60 letters, digits, hyphens, and underscores.
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Whether to support password-free access for an instance
	// - `true`: The instance access is password-free.
	// - `false`: The instance access is password-enabled. Default value: `false`. Only instances in a VPC support the password-free access.
	NoAuth *bool `json:"NoAuth,omitnil" name:"NoAuth"`

	// The node information of the instance, including node ID, type, and AZ. For more information, see [RedisNodeInfo](https://intl.cloud.tencent.com/document/product/239/20022?from_cn_redirect=1).
	// Node information of an instance. Currently, information about the node type (master or replica) and node AZ can be passed in. This parameter is not required for instances deployed in a single AZ.
	NodeSet []*RedisNodeInfo `json:"NodeSet,omitnil" name:"NodeSet"`

	// The tag for an instance
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil" name:"ResourceTags"`

	// Name of the AZ where the instance resides. For more information, see [Regions and AZs](https://intl.cloud.tencent.com/document/product/239/4106?from_cn_redirect=1).
	ZoneName *string `json:"ZoneName,omitnil" name:"ZoneName"`

	// The parameter template ID associated with the instance
	// - If this parameter is not configured, the system will automatically adapt the corresponding default template based on the selected compatible version and architecture.
	// - Query the list of parameter templates of an instance to get the template ID through the [DescribeParamTemplates](https://intl.cloud.tencent.com/document/product/239/58750?from_cn_redirect=1) API.
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// An internal parameter used to indicate whether to check when creating an instance.
	// - `false`: Default value. Send a normal request and create an instance if all the requirements are met.
	// - `true`: Send a check request and create no instance.
	DryRun *bool `json:"DryRun,omitnil" name:"DryRun"`

	// The product edition of the instance
	// - `local`: Local Disk Edition.
	// - `cloud`: Cloud Disk Edition.
	// - `cdc`: Dedicated Cluster Edition. Default value: `local`.
	ProductVersion *string `json:"ProductVersion,omitnil" name:"ProductVersion"`

	// Exclusive cluster ID. When `ProductVersion` is set to `cdc`, this parameter is required.
	RedisClusterId *string `json:"RedisClusterId,omitnil" name:"RedisClusterId"`
}

type CreateInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Instance type
	// <ul><li>`2`: Redis 2.8 Memory Edition (Standard Architecture). </li><li>`3`: CKV 3.2 Memory Edition (Standard Architecture). </li><li>`4`: CKV 3.2 Memory Edition (Cluster Architecture). </li><li>`6`: Redis 4.0 Memory Edition (Standard Architecture). </li><li>`7`: Redis 4.0 Memory Edition (Cluster Architecture). </li><li>`8`: Redis 5.0 Memory Edition (Standard Architecture). </li><li>`9`: Redis 5.0 Memory Edition (Cluster Architecture). </li><li>`15`: Redis 6.2 Memory Edition (Standard Architecture). </li><li>`16`: Redis 6.2 Memory Edition (Cluster Architecture).</li></ul>
	TypeId *uint64 `json:"TypeId,omitnil" name:"TypeId"`

	// Memory capacity in MB, which must be an integer multiple of 1024. For specific specifications, query the sales specifications in all regions through the [DescribeProductInfo](https://intl.cloud.tencent.com/document/api/239/30600?from_cn_redirect=1) API.
	// - When **TypeId** is a standard architecture, **MemSize** is the total memory capacity of the instance;
	// - When **TypeId** is a cluster architecture, **MemSize** is the single-shard memory capacity.
	MemSize *uint64 `json:"MemSize,omitnil" name:"MemSize"`

	// The number of instances for each purchase. For details, query the sales specifications in all regions through the [DescribeProductInfo](https://intl.cloud.tencent.com/document/api/239/30600?from_cn_redirect=1) API.
	GoodsNum *uint64 `json:"GoodsNum,omitnil" name:"GoodsNum"`

	// The purchase duration of an instance
	// - If `BillingMode` is `1`, that is, when the billing mode is monthly subscription, you need to set this parameter to specify the duration of the purchased instance. Unit: month. Value range: [1,2,3,4,5,6,7,8,9,10,11,12,24,36].
	// - If `BillingMode` is `0`, that is, when the billing mode is pay-as-you-go, you need to set this parameter to `1`.
	Period *uint64 `json:"Period,omitnil" name:"Period"`

	// Billing mode. 0: pay-as-you-go
	BillingMode *int64 `json:"BillingMode,omitnil" name:"BillingMode"`

	// ID of the AZ where the instance resides. For more information, see [Regions and AZs](https://intl.cloud.tencent.com/document/product/239/4106?from_cn_redirect=1).
	ZoneId *uint64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// Instance access password
	// - When the input parameter `NoAuth` is `true`, it means that the instance access is set to be password-free, and the `Password` field does not need to be configured; otherwise, `Password` is a required parameter.
	// - When the instance type `TypeId` is Redis 2.8 Memory Edition (Standard Architecture), Redis 4.0, 5.0, 6.0 (regardless of architecture), the password must contains 8-30 characters in at least two of the following types: lowercase letters, uppercase letters, digits, and symbols (()`~!@#$%^&*-+=_|{}[]:;<>,.?/), and it cannot start with a slash (/).
	// - When the instance type **TypeId** is CKV 3.2 Memory Edition (regardless of architecture), the password contains 8-30 letters and digits and excludes other characters.
	Password *string `json:"Password,omitnil" name:"Password"`

	// VPC ID. If this parameter is not passed in, the classic network will be selected by default. You can query the specific VPC ID in the [VPC console](https://console.cloud.tencent.com/vpc).
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// VPC subnet ID. This parameter is not required for the classic network. You can get the specific subnet ID by querying the subnet list in the [VPC console](https://console.cloud.tencent.com/vpc).
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// Project ID. Log in to the [Redis console](https://console.cloud.tencent.com/redis#/), go to the account information menu in the top-right corner, and select **Project Management** to query the project ID.
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Auto-renewal flag
	// - `0`: Manual renewal (default).
	// - `1`: Auto-renewal.
	// - `2`: Not auto-renewal (set by user).
	AutoRenew *uint64 `json:"AutoRenew,omitnil" name:"AutoRenew"`

	// Array of security group IDs. Get the security group ID of the instance through the [DescribeInstanceSecurityGroup](https://intl.cloud.tencent.com/document/product/239/34447?from_cn_redirect=1) API.
	SecurityGroupIdList []*string `json:"SecurityGroupIdList,omitnil" name:"SecurityGroupIdList"`

	// User-defined network port. Default value: `6379`. Range: [1024,65535].
	VPort *uint64 `json:"VPort,omitnil" name:"VPort"`

	// Quantity of instance shards
	// - This parameter is not required for instances of Standard Edition.
	// - For instances of Cluster Edition, the range of shard quantity is [1, 3, 5, 8, 12, 16, 24, 32, 40, 48, 64, 80, 96, 128].
	RedisShardNum *int64 `json:"RedisShardNum,omitnil" name:"RedisShardNum"`

	// Quantity of instance replicas
	// - For Redis Memory Edition 4.0, 5.0, 6.2 (regardless of architecture), the range of replica quantity is [1,5].
	// - For Redis 2.8 Standard Edition and CKV Standard Edition, the replica quantity is `1`.
	RedisReplicasNum *int64 `json:"RedisReplicasNum,omitnil" name:"RedisReplicasNum"`

	// Whether to support read-only replicas.
	// - Redis 2.8 Standard Edition and CKV Standard Edition don’t support read-only replicas.
	// - If read-only replicas are enabled, read/write separation will be automatically enabled for an instance, with write requests routed to the master node and read requests to the replica node.
	// - To enable read-only replicas, we recommend that you create two or more replicas.
	ReplicasReadonly *bool `json:"ReplicasReadonly,omitnil" name:"ReplicasReadonly"`

	// Instance name, which can contain up to 60 letters, digits, hyphens, and underscores.
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Whether to support password-free access for an instance
	// - `true`: The instance access is password-free.
	// - `false`: The instance access is password-enabled. Default value: `false`. Only instances in a VPC support the password-free access.
	NoAuth *bool `json:"NoAuth,omitnil" name:"NoAuth"`

	// The node information of the instance, including node ID, type, and AZ. For more information, see [RedisNodeInfo](https://intl.cloud.tencent.com/document/product/239/20022?from_cn_redirect=1).
	// Node information of an instance. Currently, information about the node type (master or replica) and node AZ can be passed in. This parameter is not required for instances deployed in a single AZ.
	NodeSet []*RedisNodeInfo `json:"NodeSet,omitnil" name:"NodeSet"`

	// The tag for an instance
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil" name:"ResourceTags"`

	// Name of the AZ where the instance resides. For more information, see [Regions and AZs](https://intl.cloud.tencent.com/document/product/239/4106?from_cn_redirect=1).
	ZoneName *string `json:"ZoneName,omitnil" name:"ZoneName"`

	// The parameter template ID associated with the instance
	// - If this parameter is not configured, the system will automatically adapt the corresponding default template based on the selected compatible version and architecture.
	// - Query the list of parameter templates of an instance to get the template ID through the [DescribeParamTemplates](https://intl.cloud.tencent.com/document/product/239/58750?from_cn_redirect=1) API.
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// An internal parameter used to indicate whether to check when creating an instance.
	// - `false`: Default value. Send a normal request and create an instance if all the requirements are met.
	// - `true`: Send a check request and create no instance.
	DryRun *bool `json:"DryRun,omitnil" name:"DryRun"`

	// The product edition of the instance
	// - `local`: Local Disk Edition.
	// - `cloud`: Cloud Disk Edition.
	// - `cdc`: Dedicated Cluster Edition. Default value: `local`.
	ProductVersion *string `json:"ProductVersion,omitnil" name:"ProductVersion"`

	// Exclusive cluster ID. When `ProductVersion` is set to `cdc`, this parameter is required.
	RedisClusterId *string `json:"RedisClusterId,omitnil" name:"RedisClusterId"`
}

func (r *CreateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TypeId")
	delete(f, "MemSize")
	delete(f, "GoodsNum")
	delete(f, "Period")
	delete(f, "BillingMode")
	delete(f, "ZoneId")
	delete(f, "Password")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "ProjectId")
	delete(f, "AutoRenew")
	delete(f, "SecurityGroupIdList")
	delete(f, "VPort")
	delete(f, "RedisShardNum")
	delete(f, "RedisReplicasNum")
	delete(f, "ReplicasReadonly")
	delete(f, "InstanceName")
	delete(f, "NoAuth")
	delete(f, "NodeSet")
	delete(f, "ResourceTags")
	delete(f, "ZoneName")
	delete(f, "TemplateId")
	delete(f, "DryRun")
	delete(f, "ProductVersion")
	delete(f, "RedisClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstancesResponseParams struct {
	// Transaction ID
	DealId *string `json:"DealId,omitnil" name:"DealId"`

	// Instance ID
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstancesResponseParams `json:"Response"`
}

func (r *CreateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateParamTemplateRequestParams struct {
	// Parameter template name.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Parameter template description.
	Description *string `json:"Description,omitnil" name:"Description"`

	// Instance type. Valid values: `1` (Redis 2.8 Memory Edition in cluster architecture), `2` (Redis 2.8 Memory Edition in standard architecture), `3` (CKV 3.2 Memory Edition in standard architecture), `4` (CKV 3.2 Memory Edition in cluster architecture), `5` (Redis 2.8 Memory Edition in standalone architecture), `6` (Redis 4.0 Memory Edition in standard architecture), `7` (Redis 4.0 Memory Edition in cluster architecture), `8` (Redis 5.0 Memory Edition in standard architecture), `9` (Redis 5.0 Memory Edition in cluster architecture). If `TempateId` is specified, this parameter can be left blank; otherwise, it is required.
	ProductType *uint64 `json:"ProductType,omitnil" name:"ProductType"`

	// ID of the source parameter template.
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// Parameter list.
	ParamList []*InstanceParam `json:"ParamList,omitnil" name:"ParamList"`
}

type CreateParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Parameter template name.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Parameter template description.
	Description *string `json:"Description,omitnil" name:"Description"`

	// Instance type. Valid values: `1` (Redis 2.8 Memory Edition in cluster architecture), `2` (Redis 2.8 Memory Edition in standard architecture), `3` (CKV 3.2 Memory Edition in standard architecture), `4` (CKV 3.2 Memory Edition in cluster architecture), `5` (Redis 2.8 Memory Edition in standalone architecture), `6` (Redis 4.0 Memory Edition in standard architecture), `7` (Redis 4.0 Memory Edition in cluster architecture), `8` (Redis 5.0 Memory Edition in standard architecture), `9` (Redis 5.0 Memory Edition in cluster architecture). If `TempateId` is specified, this parameter can be left blank; otherwise, it is required.
	ProductType *uint64 `json:"ProductType,omitnil" name:"ProductType"`

	// ID of the source parameter template.
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// Parameter list.
	ParamList []*InstanceParam `json:"ParamList,omitnil" name:"ParamList"`
}

func (r *CreateParamTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateParamTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "ProductType")
	delete(f, "TemplateId")
	delete(f, "ParamList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateParamTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateParamTemplateResponseParams struct {
	// Parameter template ID.
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateParamTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateParamTemplateResponseParams `json:"Response"`
}

func (r *CreateParamTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateParamTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DelayDistribution struct {
	// The delay distribution. The mapping between delay range and `Ladder` value is as follows:  - `1`: [0ms,1ms]. - `5`: [1ms,5ms]. - `10`: [5ms,10ms]. - `50`: [10ms,50ms]. - `200`:  [50ms,200ms]. - `-1`: [200ms,∞].
	Ladder *int64 `json:"Ladder,omitnil" name:"Ladder"`

	// The number of commands with delay falling within the current delay range -
	Size *int64 `json:"Size,omitnil" name:"Size"`

	// Modification time
	Updatetime *int64 `json:"Updatetime,omitnil" name:"Updatetime"`
}

// Predefined struct for user
type DeleteInstanceAccountRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Sub-account name
	AccountName *string `json:"AccountName,omitnil" name:"AccountName"`
}

type DeleteInstanceAccountRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Sub-account name
	AccountName *string `json:"AccountName,omitnil" name:"AccountName"`
}

func (r *DeleteInstanceAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AccountName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteInstanceAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceAccountResponseParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteInstanceAccountResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInstanceAccountResponseParams `json:"Response"`
}

func (r *DeleteInstanceAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteParamTemplateRequestParams struct {
	// Parameter template ID.
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`
}

type DeleteParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Parameter template ID.
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`
}

func (r *DeleteParamTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteParamTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteParamTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteParamTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteParamTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteParamTemplateResponseParams `json:"Response"`
}

func (r *DeleteParamTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteParamTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoBackupConfigRequestParams struct {
	// ID of the specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeAutoBackupConfigRequest struct {
	*tchttp.BaseRequest
	
	// ID of the specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeAutoBackupConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoBackupConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoBackupConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoBackupConfigResponseParams struct {
	// This parameter is retained due to compatibility and can be ignored.
	AutoBackupType *int64 `json:"AutoBackupType,omitnil" name:"AutoBackupType"`

	// Backup cycle, which will be daily by default. Valid values: `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`.
	WeekDays []*string `json:"WeekDays,omitnil" name:"WeekDays"`

	// Time period for backup task initialization
	TimePeriod *string `json:"TimePeriod,omitnil" name:"TimePeriod"`

	// Retention time of full backup files in days.  Default value: `7`.  To retain the files for more days, [submit a ticket](https://console.cloud.tencent.com/workorder/category) for application.
	BackupStorageDays *int64 `json:"BackupStorageDays,omitnil" name:"BackupStorageDays"`

	// This parameter has been disused.
	BinlogStorageDays *int64 `json:"BinlogStorageDays,omitnil" name:"BinlogStorageDays"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAutoBackupConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoBackupConfigResponseParams `json:"Response"`
}

func (r *DescribeAutoBackupConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoBackupConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDownloadRestrictionRequestParams struct {

}

type DescribeBackupDownloadRestrictionRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeBackupDownloadRestrictionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadRestrictionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupDownloadRestrictionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDownloadRestrictionResponseParams struct {
	// Type of the network restrictions for downloading backup files. Valid values:
	// 
	// - `NoLimit`: Backup files can be downloaded over both public and private networks.
	// - `LimitOnlyIntranet`: Backup files can be downloaded only at private network addresses auto-assigned by Tencent Cloud.
	// - `Customize`: Backup files can be downloaded only in the customized VPC.
	LimitType *string `json:"LimitType,omitnil" name:"LimitType"`

	// Only `In` can be passed in for this parameter, indicating that backup files can be downloaded in the custom `LimitVpc`.
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitnil" name:"VpcComparisonSymbol"`

	// Whether backups can be downloaded at the custom `LimitIp` address.
	// 
	// - `In`: Download is allowed for the custom IP.
	// - `NotIn`: Download is not allowed for the custom IP.
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitnil" name:"IpComparisonSymbol"`

	// VPC ID of the custom backup file download address, which will be displayed if `LimitType` is `Customize`.
	LimitVpc []*BackupLimitVpcItem `json:"LimitVpc,omitnil" name:"LimitVpc"`

	// VPC ID of the custom backup file download address, which will be displayed if `LimitType` is `Customize`.
	LimitIp []*string `json:"LimitIp,omitnil" name:"LimitIp"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBackupDownloadRestrictionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupDownloadRestrictionResponseParams `json:"Response"`
}

func (r *DescribeBackupDownloadRestrictionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadRestrictionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupUrlRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Backup ID, which can be obtained through the `RedisBackupSet` parameter returned by the [DescribeInstanceBackups](https://intl.cloud.tencent.com/document/product/239/20011?from_cn_redirect=1) API.
	BackupId *string `json:"BackupId,omitnil" name:"BackupId"`

	// Type of the network restriction for downloading backup files. If this parameter is not configured, the user-defined configuration will be used.
	// 
	// - `NoLimit`: Backup files can be downloaded over both public and private networks.
	// - `LimitOnlyIntranet`: Backup files can be downloaded only at private network addresses auto-assigned by Tencent Cloud.
	// - `Customize`: Backup files can be downloaded only in the customized VPC.
	LimitType *string `json:"LimitType,omitnil" name:"LimitType"`

	// Only `In` can be passed in for this parameter, indicating that backup files can be downloaded in the custom `LimitVpc`.
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitnil" name:"VpcComparisonSymbol"`

	// Whether backups can be downloaded at the custom `LimitIp` address.
	// 
	// - `In` (default value): Download is allowed for the custom IP.
	// - `NotIn`: Download is not allowed for the custom IP.
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitnil" name:"IpComparisonSymbol"`

	// VPC ID of the custom backup file download address, which is required if `LimitType` is `Customize`.
	LimitVpc []*BackupLimitVpcItem `json:"LimitVpc,omitnil" name:"LimitVpc"`

	// VPC IP of the custom backup file download address, which is required if `LimitType` is `Customize`.
	LimitIp []*string `json:"LimitIp,omitnil" name:"LimitIp"`
}

type DescribeBackupUrlRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Backup ID, which can be obtained through the `RedisBackupSet` parameter returned by the [DescribeInstanceBackups](https://intl.cloud.tencent.com/document/product/239/20011?from_cn_redirect=1) API.
	BackupId *string `json:"BackupId,omitnil" name:"BackupId"`

	// Type of the network restriction for downloading backup files. If this parameter is not configured, the user-defined configuration will be used.
	// 
	// - `NoLimit`: Backup files can be downloaded over both public and private networks.
	// - `LimitOnlyIntranet`: Backup files can be downloaded only at private network addresses auto-assigned by Tencent Cloud.
	// - `Customize`: Backup files can be downloaded only in the customized VPC.
	LimitType *string `json:"LimitType,omitnil" name:"LimitType"`

	// Only `In` can be passed in for this parameter, indicating that backup files can be downloaded in the custom `LimitVpc`.
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitnil" name:"VpcComparisonSymbol"`

	// Whether backups can be downloaded at the custom `LimitIp` address.
	// 
	// - `In` (default value): Download is allowed for the custom IP.
	// - `NotIn`: Download is not allowed for the custom IP.
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitnil" name:"IpComparisonSymbol"`

	// VPC ID of the custom backup file download address, which is required if `LimitType` is `Customize`.
	LimitVpc []*BackupLimitVpcItem `json:"LimitVpc,omitnil" name:"LimitVpc"`

	// VPC IP of the custom backup file download address, which is required if `LimitType` is `Customize`.
	LimitIp []*string `json:"LimitIp,omitnil" name:"LimitIp"`
}

func (r *DescribeBackupUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupId")
	delete(f, "LimitType")
	delete(f, "VpcComparisonSymbol")
	delete(f, "IpComparisonSymbol")
	delete(f, "LimitVpc")
	delete(f, "LimitIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupUrlResponseParams struct {
	// Public network download address (valid for six hours). This field will be disused soon.
	DownloadUrl []*string `json:"DownloadUrl,omitnil" name:"DownloadUrl"`

	// Private network download address (valid for six hours). This field will be disused soon.
	InnerDownloadUrl []*string `json:"InnerDownloadUrl,omitnil" name:"InnerDownloadUrl"`

	// Filename. This field will be disused soon.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Filenames []*string `json:"Filenames,omitnil" name:"Filenames"`

	// List of backup file information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupInfos []*BackupDownloadInfo `json:"BackupInfos,omitnil" name:"BackupInfos"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBackupUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupUrlResponseParams `json:"Response"`
}

func (r *DescribeBackupUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBandwidthRangeRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeBandwidthRangeRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeBandwidthRangeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBandwidthRangeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBandwidthRangeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBandwidthRangeResponseParams struct {
	// Standard bandwidth, which is the bandwidth allocated by the system to each node when an instance is purchased.
	BaseBandwidth *int64 `json:"BaseBandwidth,omitnil" name:"BaseBandwidth"`

	// The additional bandwidth of the instance. If the standard bandwidth does not meet your needs, you can increase the bandwidth on your own. <ul><li>If read-only replica is enabled, the total instance bandwidth = additional bandwidth * shard quantity + standard bandwidth * shard quantity * Max ([read-only replica quantity, 1]). The shard quantity in the standard architecture is 1. </li><li>If read-only replica is not enabled, the total instance bandwidth = additional bandwidth * shard quantity + standard bandwidth * shard quantity. The shard quantity in the standard architecture is 1.</li></ul>
	AddBandwidth *int64 `json:"AddBandwidth,omitnil" name:"AddBandwidth"`

	// The lower limit for additional bandwidth
	MinAddBandwidth *int64 `json:"MinAddBandwidth,omitnil" name:"MinAddBandwidth"`

	// The upper limit for additional bandwidth
	MaxAddBandwidth *int64 `json:"MaxAddBandwidth,omitnil" name:"MaxAddBandwidth"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBandwidthRangeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBandwidthRangeResponseParams `json:"Response"`
}

func (r *DescribeBandwidthRangeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBandwidthRangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCommonDBInstancesRequestParams struct {
	// List of VPC IDs
	VpcIds []*int64 `json:"VpcIds,omitnil" name:"VpcIds"`

	// List of subnet IDs
	SubnetIds []*int64 `json:"SubnetIds,omitnil" name:"SubnetIds"`

	// List of billing modes. 0: monthly subscription; 1: pay-as-you-go
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`

	// List of instance IDs
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// List of instance names
	InstanceNames []*string `json:"InstanceNames,omitnil" name:"InstanceNames"`

	// List of instance status
	Status []*string `json:"Status,omitnil" name:"Status"`

	// Sorting field
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// Sorting order
	OrderByType *string `json:"OrderByType,omitnil" name:"OrderByType"`

	// List of instance VIPs
	Vips []*string `json:"Vips,omitnil" name:"Vips"`

	// List of VPC IDs
	UniqVpcIds []*string `json:"UniqVpcIds,omitnil" name:"UniqVpcIds"`

	// List of unique subnet IDs
	UniqSubnetIds []*string `json:"UniqSubnetIds,omitnil" name:"UniqSubnetIds"`

	// Quantity limit. Recommended default value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeCommonDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// List of VPC IDs
	VpcIds []*int64 `json:"VpcIds,omitnil" name:"VpcIds"`

	// List of subnet IDs
	SubnetIds []*int64 `json:"SubnetIds,omitnil" name:"SubnetIds"`

	// List of billing modes. 0: monthly subscription; 1: pay-as-you-go
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`

	// List of instance IDs
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// List of instance names
	InstanceNames []*string `json:"InstanceNames,omitnil" name:"InstanceNames"`

	// List of instance status
	Status []*string `json:"Status,omitnil" name:"Status"`

	// Sorting field
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// Sorting order
	OrderByType *string `json:"OrderByType,omitnil" name:"OrderByType"`

	// List of instance VIPs
	Vips []*string `json:"Vips,omitnil" name:"Vips"`

	// List of VPC IDs
	UniqVpcIds []*string `json:"UniqVpcIds,omitnil" name:"UniqVpcIds"`

	// List of unique subnet IDs
	UniqSubnetIds []*string `json:"UniqSubnetIds,omitnil" name:"UniqSubnetIds"`

	// Quantity limit. Recommended default value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeCommonDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCommonDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcIds")
	delete(f, "SubnetIds")
	delete(f, "PayMode")
	delete(f, "InstanceIds")
	delete(f, "InstanceNames")
	delete(f, "Status")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Vips")
	delete(f, "UniqVpcIds")
	delete(f, "UniqSubnetIds")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCommonDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCommonDBInstancesResponseParams struct {
	// Number of instances
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Instance information
	InstanceDetails []*RedisCommonInstanceList `json:"InstanceDetails,omitnil" name:"InstanceDetails"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCommonDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCommonDBInstancesResponseParams `json:"Response"`
}

func (r *DescribeCommonDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCommonDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSecurityGroupsRequestParams struct {
	// Database engine name, which is `redis` for this API.
	Product *string `json:"Product,omitnil" name:"Product"`

	// ID of the specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Database engine name, which is `redis` for this API.
	Product *string `json:"Product,omitnil" name:"Product"`

	// ID of the specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
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
	// Security group rules
	Groups []*SecurityGroup `json:"Groups,omitnil" name:"Groups"`

	// Private IPv4 address of an instance
	VIP *string `json:"VIP,omitnil" name:"VIP"`

	// Private network port
	VPort *string `json:"VPort,omitnil" name:"VPort"`

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
type DescribeInstanceAccountRequestParams struct {
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Number of entries per page
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Pagination offset,  which is an integral multiple of `Limit`.  Calculation formula:  `offset` = `limit` * (page number - 1).
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeInstanceAccountRequest struct {
	*tchttp.BaseRequest
	
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Number of entries per page
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Pagination offset,  which is an integral multiple of `Limit`.  Calculation formula:  `offset` = `limit` * (page number - 1).
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeInstanceAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceAccountResponseParams struct {
	// Account details 
	// Note:  This field may return null, indicating that no valid values can be obtained.
	Accounts []*Account `json:"Accounts,omitnil" name:"Accounts"`

	// Number of accounts 
	// Note:  This field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstanceAccountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceAccountResponseParams `json:"Response"`
}

func (r *DescribeInstanceAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceBackupsRequestParams struct {
	// Number of backups returned per page. Default value: `20`. Maximum value: `100`.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Pagination offset, which is an integral multiple of `Limit`. `offset` = `limit` * (page number - 1).
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// ID of the instance to be operated on, which can be obtained through the `InstanceId` field in the return value of the `DescribeInstance` API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Start time in the format of yyyy-MM-dd HH:mm:ss, such as 2017-02-08 16:46:34. This parameter is used to query the list of instance backups started during the [beginTime, endTime] range.
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// End time in the format of yyyy-MM-dd HH:mm:ss, such as 2017-02-08 19:09:26. This parameter is used to query the list of instance backups started during the [beginTime, endTime] range.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Backup task status:
	// `1`: The backup is in the process.
	// `2`: The backup is normal.
	// `3`: The backup is being converted to an RDB file.
	// `4`: Conversion to RDB has been completed.
	// `-1`: The backup expired.
	// `-2`: The backup has been deleted.
	Status []*int64 `json:"Status,omitnil" name:"Status"`

	// Instance name, which can be fuzzily searched.
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`
}

type DescribeInstanceBackupsRequest struct {
	*tchttp.BaseRequest
	
	// Number of backups returned per page. Default value: `20`. Maximum value: `100`.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Pagination offset, which is an integral multiple of `Limit`. `offset` = `limit` * (page number - 1).
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// ID of the instance to be operated on, which can be obtained through the `InstanceId` field in the return value of the `DescribeInstance` API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Start time in the format of yyyy-MM-dd HH:mm:ss, such as 2017-02-08 16:46:34. This parameter is used to query the list of instance backups started during the [beginTime, endTime] range.
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// End time in the format of yyyy-MM-dd HH:mm:ss, such as 2017-02-08 19:09:26. This parameter is used to query the list of instance backups started during the [beginTime, endTime] range.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Backup task status:
	// `1`: The backup is in the process.
	// `2`: The backup is normal.
	// `3`: The backup is being converted to an RDB file.
	// `4`: Conversion to RDB has been completed.
	// `-1`: The backup expired.
	// `-2`: The backup has been deleted.
	Status []*int64 `json:"Status,omitnil" name:"Status"`

	// Instance name, which can be fuzzily searched.
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`
}

func (r *DescribeInstanceBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceBackupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "InstanceId")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Status")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceBackupsResponseParams struct {
	// Total number of backups.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Array of instance backups.
	BackupSet []*RedisBackupSet `json:"BackupSet,omitnil" name:"BackupSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstanceBackupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceBackupsResponseParams `json:"Response"`
}

func (r *DescribeInstanceBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceDTSInfoRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeInstanceDTSInfoRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeInstanceDTSInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceDTSInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceDTSInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceDTSInfoResponseParams struct {
	// DTS task ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// DTS task name
	// Note: This field may return null, indicating that no valid values can be obtained.
	JobName *string `json:"JobName,omitnil" name:"JobName"`

	// Task status. Valid values: 1 (Creating), 3 (Checking), 4 (CheckPass), 5 (CheckNotPass), 7 (Running), 8 (ReadyComplete), 9 (Success), 10 (Failed), 11 (Stopping), 12 (Completing)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// Status description
	// Note: This field may return null, indicating that no valid values can be obtained.
	StatusDesc *string `json:"StatusDesc,omitnil" name:"StatusDesc"`

	// Sync latency in bytes
	// Note: This field may return null, indicating that no valid values can be obtained.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Disconnection time
	// Note: This field may return null, indicating that no valid values can be obtained.
	CutDownTime *string `json:"CutDownTime,omitnil" name:"CutDownTime"`

	// Source instance information
	// Note: This field may return null, indicating that no valid values can be obtained.
	SrcInfo *DescribeInstanceDTSInstanceInfo `json:"SrcInfo,omitnil" name:"SrcInfo"`

	// Target instance information
	// Note: This field may return null, indicating that no valid values can be obtained.
	DstInfo *DescribeInstanceDTSInstanceInfo `json:"DstInfo,omitnil" name:"DstInfo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstanceDTSInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceDTSInfoResponseParams `json:"Response"`
}

func (r *DescribeInstanceDTSInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceDTSInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceDTSInstanceInfo struct {
	// Region ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegionId *int64 `json:"RegionId,omitnil" name:"RegionId"`

	// Instance ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Repository ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	SetId *int64 `json:"SetId,omitnil" name:"SetId"`

	// AZ ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneId *int64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// Instance type
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// Instance name
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Instance access address
	// Note: This field may return null, indicating that no valid values can be obtained.
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// Status
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

// Predefined struct for user
type DescribeInstanceDealDetailRequestParams struct {
	// Array of order transaction IDs, i.e., the `DealId` output parameter of the [CreateInstances](https://intl.cloud.tencent.com/document/api/239/20026?from_cn_redirect=1) API.
	DealIds []*string `json:"DealIds,omitnil" name:"DealIds"`
}

type DescribeInstanceDealDetailRequest struct {
	*tchttp.BaseRequest
	
	// Array of order transaction IDs, i.e., the `DealId` output parameter of the [CreateInstances](https://intl.cloud.tencent.com/document/api/239/20026?from_cn_redirect=1) API.
	DealIds []*string `json:"DealIds,omitnil" name:"DealIds"`
}

func (r *DescribeInstanceDealDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceDealDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DealIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceDealDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceDealDetailResponseParams struct {
	// Order details
	DealDetails []*TradeDealDetail `json:"DealDetails,omitnil" name:"DealDetails"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstanceDealDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceDealDetailResponseParams `json:"Response"`
}

func (r *DescribeInstanceDealDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceDealDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorBigKeyRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Request type. 1: string type; 2: all types
	ReqType *int64 `json:"ReqType,omitnil" name:"ReqType"`

	// Time, such as "20190219"
	Date *string `json:"Date,omitnil" name:"Date"`
}

type DescribeInstanceMonitorBigKeyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Request type. 1: string type; 2: all types
	ReqType *int64 `json:"ReqType,omitnil" name:"ReqType"`

	// Time, such as "20190219"
	Date *string `json:"Date,omitnil" name:"Date"`
}

func (r *DescribeInstanceMonitorBigKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorBigKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ReqType")
	delete(f, "Date")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorBigKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorBigKeyResponseParams struct {
	// Big key details
	Data []*BigKeyInfo `json:"Data,omitnil" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstanceMonitorBigKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceMonitorBigKeyResponseParams `json:"Response"`
}

func (r *DescribeInstanceMonitorBigKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorBigKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorBigKeySizeDistRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Time, such as "20190219"
	Date *string `json:"Date,omitnil" name:"Date"`
}

type DescribeInstanceMonitorBigKeySizeDistRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Time, such as "20190219"
	Date *string `json:"Date,omitnil" name:"Date"`
}

func (r *DescribeInstanceMonitorBigKeySizeDistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorBigKeySizeDistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Date")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorBigKeySizeDistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorBigKeySizeDistResponseParams struct {
	// Big key size distribution details
	Data []*DelayDistribution `json:"Data,omitnil" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstanceMonitorBigKeySizeDistResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceMonitorBigKeySizeDistResponseParams `json:"Response"`
}

func (r *DescribeInstanceMonitorBigKeySizeDistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorBigKeySizeDistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorBigKeyTypeDistRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Time, such as "20190219"
	Date *string `json:"Date,omitnil" name:"Date"`
}

type DescribeInstanceMonitorBigKeyTypeDistRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Time, such as "20190219"
	Date *string `json:"Date,omitnil" name:"Date"`
}

func (r *DescribeInstanceMonitorBigKeyTypeDistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorBigKeyTypeDistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Date")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorBigKeyTypeDistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorBigKeyTypeDistResponseParams struct {
	// Big key type distribution details
	Data []*BigKeyTypeInfo `json:"Data,omitnil" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstanceMonitorBigKeyTypeDistResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceMonitorBigKeyTypeDistResponseParams `json:"Response"`
}

func (r *DescribeInstanceMonitorBigKeyTypeDistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorBigKeyTypeDistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorHotKeyRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Time span. 1: real time; 2: past 30 minutes; 3: past 6 hours; 4: past 24 hours
	SpanType *int64 `json:"SpanType,omitnil" name:"SpanType"`
}

type DescribeInstanceMonitorHotKeyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Time span. 1: real time; 2: past 30 minutes; 3: past 6 hours; 4: past 24 hours
	SpanType *int64 `json:"SpanType,omitnil" name:"SpanType"`
}

func (r *DescribeInstanceMonitorHotKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorHotKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SpanType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorHotKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorHotKeyResponseParams struct {
	// Hot key details
	Data []*HotKeyInfo `json:"Data,omitnil" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstanceMonitorHotKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceMonitorHotKeyResponseParams `json:"Response"`
}

func (r *DescribeInstanceMonitorHotKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorHotKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorSIPRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeInstanceMonitorSIPRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeInstanceMonitorSIPRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorSIPRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorSIPRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorSIPResponseParams struct {
	// Access source information
	Data []*SourceInfo `json:"Data,omitnil" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstanceMonitorSIPResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceMonitorSIPResponseParams `json:"Response"`
}

func (r *DescribeInstanceMonitorSIPResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorSIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorTookDistRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Time, such as "20190219"
	Date *string `json:"Date,omitnil" name:"Date"`

	// Time span. 1: real time; 2: last 30 minutes; 3: last 6 hours; 4: last 24 hours
	SpanType *int64 `json:"SpanType,omitnil" name:"SpanType"`
}

type DescribeInstanceMonitorTookDistRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Time, such as "20190219"
	Date *string `json:"Date,omitnil" name:"Date"`

	// Time span. 1: real time; 2: last 30 minutes; 3: last 6 hours; 4: last 24 hours
	SpanType *int64 `json:"SpanType,omitnil" name:"SpanType"`
}

func (r *DescribeInstanceMonitorTookDistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorTookDistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Date")
	delete(f, "SpanType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorTookDistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorTookDistResponseParams struct {
	// Latency distribution information
	Data []*DelayDistribution `json:"Data,omitnil" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstanceMonitorTookDistResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceMonitorTookDistResponseParams `json:"Response"`
}

func (r *DescribeInstanceMonitorTookDistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorTookDistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorTopNCmdRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Time span. 1: real time; 2: last 30 minutes; 3: last 6 hours; 4: last 24 hours
	SpanType *int64 `json:"SpanType,omitnil" name:"SpanType"`
}

type DescribeInstanceMonitorTopNCmdRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Time span. 1: real time; 2: last 30 minutes; 3: last 6 hours; 4: last 24 hours
	SpanType *int64 `json:"SpanType,omitnil" name:"SpanType"`
}

func (r *DescribeInstanceMonitorTopNCmdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorTopNCmdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SpanType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorTopNCmdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorTopNCmdResponseParams struct {
	// Access command information
	Data []*SourceCommand `json:"Data,omitnil" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstanceMonitorTopNCmdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceMonitorTopNCmdResponseParams `json:"Response"`
}

func (r *DescribeInstanceMonitorTopNCmdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorTopNCmdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorTopNCmdTookRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Time span. 1: real time; 2: last 30 minutes; 3: last 6 hours; 4: last 24 hours
	SpanType *int64 `json:"SpanType,omitnil" name:"SpanType"`
}

type DescribeInstanceMonitorTopNCmdTookRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Time span. 1: real time; 2: last 30 minutes; 3: last 6 hours; 4: last 24 hours
	SpanType *int64 `json:"SpanType,omitnil" name:"SpanType"`
}

func (r *DescribeInstanceMonitorTopNCmdTookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorTopNCmdTookRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SpanType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorTopNCmdTookRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorTopNCmdTookResponseParams struct {
	// Duration details
	Data []*CommandTake `json:"Data,omitnil" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstanceMonitorTopNCmdTookResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceMonitorTopNCmdTookResponseParams `json:"Response"`
}

func (r *DescribeInstanceMonitorTopNCmdTookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorTopNCmdTookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceNodeInfoRequestParams struct {
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// List size Size of node information returned per page.  Default value: `20`. Maximum value: `1000`.  This field has been disused.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Pagination offset, which is an integral multiple of `Limit`. Calculation formula:  `offset` = `limit` * (page number - 1). This field has been disused.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeInstanceNodeInfoRequest struct {
	*tchttp.BaseRequest
	
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// List size Size of node information returned per page.  Default value: `20`. Maximum value: `1000`.  This field has been disused.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Pagination offset, which is an integral multiple of `Limit`. Calculation formula:  `offset` = `limit` * (page number - 1). This field has been disused.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
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
	// The number of proxy nodes
	ProxyCount *int64 `json:"ProxyCount,omitnil" name:"ProxyCount"`

	// Proxy node information 
	// Note:  This field may return null, indicating that no valid values can be obtained.
	Proxy []*ProxyNodes `json:"Proxy,omitnil" name:"Proxy"`

	// The number of Redis nodes
	RedisCount *int64 `json:"RedisCount,omitnil" name:"RedisCount"`

	// Redis node information 
	// Note:  This field may return null, indicating that no valid values can be obtained.
	Redis []*RedisNodes `json:"Redis,omitnil" name:"Redis"`

	// This parameter has been disused.
	TendisCount *int64 `json:"TendisCount,omitnil" name:"TendisCount"`

	// This parameter has been disused. 
	// Note:  This field may return null, indicating that no valid values can be obtained.
	Tendis []*TendisNodes `json:"Tendis,omitnil" name:"Tendis"`

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
type DescribeInstanceParamRecordsRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Maximum number of results returned per page
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset, which is an integral multiple of `Limit`.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeInstanceParamRecordsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Maximum number of results returned per page
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Offset, which is an integral multiple of `Limit`.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeInstanceParamRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceParamRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamRecordsResponseParams struct {
	// Total number of modifications.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Information of modifications.
	InstanceParamHistory []*InstanceParamHistory `json:"InstanceParamHistory,omitnil" name:"InstanceParamHistory"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstanceParamRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceParamRecordsResponseParams `json:"Response"`
}

func (r *DescribeInstanceParamRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamsRequestParams struct {
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeInstanceParamsRequest struct {
	*tchttp.BaseRequest
	
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeInstanceParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceParamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamsResponseParams struct {
	// Total number of the parameter lists
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Instance parameter in Enum type
	InstanceEnumParam []*InstanceEnumParam `json:"InstanceEnumParam,omitnil" name:"InstanceEnumParam"`

	// Instance parameter in Integer type
	InstanceIntegerParam []*InstanceIntegerParam `json:"InstanceIntegerParam,omitnil" name:"InstanceIntegerParam"`

	// Instance parameter in Char type
	InstanceTextParam []*InstanceTextParam `json:"InstanceTextParam,omitnil" name:"InstanceTextParam"`

	// Instance parameter in Multi type
	InstanceMultiParam []*InstanceMultiParam `json:"InstanceMultiParam,omitnil" name:"InstanceMultiParam"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstanceParamsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceParamsResponseParams `json:"Response"`
}

func (r *DescribeInstanceParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSecurityGroupRequestParams struct {
	// List of instance IDs,  such as "crs-f2ho5rsz\n".
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

type DescribeInstanceSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
	// List of instance IDs,  such as "crs-f2ho5rsz\n".
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

func (r *DescribeInstanceSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSecurityGroupResponseParams struct {
	// Security group information of an instance
	InstanceSecurityGroupsDetail []*InstanceSecurityGroupDetail `json:"InstanceSecurityGroupsDetail,omitnil" name:"InstanceSecurityGroupsDetail"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstanceSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceSecurityGroupResponseParams `json:"Response"`
}

func (r *DescribeInstanceSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceShardsRequestParams struct {
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Whether to filter out the replica node information. Valid values: `true` (yes),  `false` (no).
	FilterSlave *bool `json:"FilterSlave,omitnil" name:"FilterSlave"`
}

type DescribeInstanceShardsRequest struct {
	*tchttp.BaseRequest
	
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Whether to filter out the replica node information. Valid values: `true` (yes),  `false` (no).
	FilterSlave *bool `json:"FilterSlave,omitnil" name:"FilterSlave"`
}

func (r *DescribeInstanceShardsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceShardsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FilterSlave")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceShardsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceShardsResponseParams struct {
	// List information of the instance shards, which includes  node information, node ID, key count, used capacity, and capacity slope.
	InstanceShards []*InstanceClusterShard `json:"InstanceShards,omitnil" name:"InstanceShards"`

	// Number of instance shard nodes
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstanceShardsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceShardsResponseParams `json:"Response"`
}

func (r *DescribeInstanceShardsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceShardsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceZoneInfoRequestParams struct {
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeInstanceZoneInfoRequest struct {
	*tchttp.BaseRequest
	
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeInstanceZoneInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceZoneInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceZoneInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceZoneInfoResponseParams struct {
	// Number of instance node groups
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// List of instance node groups
	ReplicaGroups []*ReplicaGroup `json:"ReplicaGroups,omitnil" name:"ReplicaGroups"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstanceZoneInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceZoneInfoResponseParams `json:"Response"`
}

func (r *DescribeInstanceZoneInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceZoneInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// Number of instances returned per page. Default value: `20`. Maximum value: `1000`.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Pagination offset, which is an integral multiple of `Limit`. Calculation formula:  `offset` = `limit` * (page number - 1).
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	// 
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance list sorting criteria. The enumerated values are as listed below:  <ul><li>`projectId`:  Project ID.  </li><li>`createtime`:  Instance creation time.  </li><li>`instancename`:  Instance name.  </li><li>`type`:  Instance type. </li><li>`curDeadline`:  Instance expiration time. </li></ul>
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// Instance sorting order. <ul><li>`1`: Descending. </li><li>`0`: Ascending. Default value: `1`.</li></ul>
	OrderType *int64 `json:"OrderType,omitnil" name:"OrderType"`

	// Array of VPC IDs such as 47525. If this parameter is not passed in or the array is empty, the classic network will be selected by default. This parameter is retained and can be ignored. It is set based on `UniqVpcIds` parameter format.
	VpcIds []*string `json:"VpcIds,omitnil" name:"VpcIds"`

	// Array of VPC subnet IDs such as 56854. This parameter is retained and can be ignored. It is set based on `UniqSubnetIds` parameter format.
	SubnetIds []*string `json:"SubnetIds,omitnil" name:"SubnetIds"`

	// Keywords for fuzzy query. which can be used to fuzzy query an instance by its ID or name.
	SearchKey *string `json:"SearchKey,omitnil" name:"SearchKey"`

	// Array of project IDs
	ProjectIds []*int64 `json:"ProjectIds,omitnil" name:"ProjectIds"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Array of VPC IDs such as vpc-sad23jfdfk. If this parameter is not passed in or or the array is empty, the classic network will be selected by default.
	UniqVpcIds []*string `json:"UniqVpcIds,omitnil" name:"UniqVpcIds"`

	// Array of VPC subnet IDs such as subnet-fdj24n34j2
	UniqSubnetIds []*string `json:"UniqSubnetIds,omitnil" name:"UniqSubnetIds"`

	// Array of region IDs (disused). The corresponding region can be queried through the common parameter `Region`.
	RegionIds []*int64 `json:"RegionIds,omitnil" name:"RegionIds"`

	// Instance status. <ul><li>`0`: Uninitialized. </li><li>`1`: In the process. </li><li>`2`: Running. </li><li>`-2`: Isolated. </li><li>`-3`: To be deleted. </li></ul>
	Status []*int64 `json:"Status,omitnil" name:"Status"`

	// Instance architecture. <ul><li>`1`: Standalone edition. </li><li>`2`: Master-replica edition. </li><li>`3`: Cluster edition. </li></ul>
	TypeVersion *int64 `json:"TypeVersion,omitnil" name:"TypeVersion"`

	// Storage engine information. Valid values: `Redis-2.8`, `Redis-4.0`, `Redis-5.0`, `Redis-6.0` or `CKV`.
	EngineName *string `json:"EngineName,omitnil" name:"EngineName"`

	// Renewal mode. Valid values:  <ul><li>`0`:  Manual renewal </li><li>`1`:  Auto-renewal </li><li>`2`:  No renewal upon expiration </ul>
	AutoRenew []*int64 `json:"AutoRenew,omitnil" name:"AutoRenew"`

	// Billing mode. Only pay-as-you-go billing is supported.
	BillingMode *string `json:"BillingMode,omitnil" name:"BillingMode"`

	// Instance type. Valid values:  - `2`: Redis 2.8 Memory Edition (Standard Architecture). - `3`: CKV 3.2 Memory Edition (Standard Architecture). - `4`: CKV 3.2 Memory Edition (Cluster Architecture). - `5`: Redis 2.8 Memory Edition (Standalone). - `6`: Redis 4.0 Memory Edition (Standard Architecture). - `7`: Redis 4.0 Memory Edition (Cluster Architecture). - `8`: Redis 5.0 Memory Edition (Standard Architecture). - `9`: Redis 5.0 Memory Edition (Cluster Architecture). - `15`: Redis 6.2 Memory Edition (Standard Architecture). - `16`: Redis 6.2 Memory Edition (Cluster Architecture).
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// Array of the search keywords, which can query the instance by its ID, name, IP address.
	SearchKeys []*string `json:"SearchKeys,omitnil" name:"SearchKeys"`

	// Internal parameter, which can be ignored.
	TypeList []*int64 `json:"TypeList,omitnil" name:"TypeList"`

	// Internal parameter, which can be ignored.
	MonitorVersion *string `json:"MonitorVersion,omitnil" name:"MonitorVersion"`

	// Resources filter by tag key and value. If this parameter is not specified or is an empty array, resources will not be filtered.
	InstanceTags []*InstanceTagInfo `json:"InstanceTags,omitnil" name:"InstanceTags"`

	// Resources filter by tag key. If this parameter is not specified or is an empty array, resources will not be filtered.
	TagKeys []*string `json:"TagKeys,omitnil" name:"TagKeys"`

	// Instance product version.  If this parameter is not passed in or the array is empty, the instances will not be filtered based this parameter by default.  <ul><li>`local`:  Local disk edition.  </li><li>`cdc`:  Dedicated cluster edition.  </li></ul>
	ProductVersions []*string `json:"ProductVersions,omitnil" name:"ProductVersions"`

	// Batch query of the specified instances ID. The number of results returned is based on `Limit`.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// AZ deployment mode. <ul><li>`singleaz`: Single-AZ. </li><li>`1`: Multi-AZ. </li></ul>
	AzMode *string `json:"AzMode,omitnil" name:"AzMode"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Number of instances returned per page. Default value: `20`. Maximum value: `1000`.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Pagination offset, which is an integral multiple of `Limit`. Calculation formula:  `offset` = `limit` * (page number - 1).
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	// 
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance list sorting criteria. The enumerated values are as listed below:  <ul><li>`projectId`:  Project ID.  </li><li>`createtime`:  Instance creation time.  </li><li>`instancename`:  Instance name.  </li><li>`type`:  Instance type. </li><li>`curDeadline`:  Instance expiration time. </li></ul>
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// Instance sorting order. <ul><li>`1`: Descending. </li><li>`0`: Ascending. Default value: `1`.</li></ul>
	OrderType *int64 `json:"OrderType,omitnil" name:"OrderType"`

	// Array of VPC IDs such as 47525. If this parameter is not passed in or the array is empty, the classic network will be selected by default. This parameter is retained and can be ignored. It is set based on `UniqVpcIds` parameter format.
	VpcIds []*string `json:"VpcIds,omitnil" name:"VpcIds"`

	// Array of VPC subnet IDs such as 56854. This parameter is retained and can be ignored. It is set based on `UniqSubnetIds` parameter format.
	SubnetIds []*string `json:"SubnetIds,omitnil" name:"SubnetIds"`

	// Keywords for fuzzy query. which can be used to fuzzy query an instance by its ID or name.
	SearchKey *string `json:"SearchKey,omitnil" name:"SearchKey"`

	// Array of project IDs
	ProjectIds []*int64 `json:"ProjectIds,omitnil" name:"ProjectIds"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Array of VPC IDs such as vpc-sad23jfdfk. If this parameter is not passed in or or the array is empty, the classic network will be selected by default.
	UniqVpcIds []*string `json:"UniqVpcIds,omitnil" name:"UniqVpcIds"`

	// Array of VPC subnet IDs such as subnet-fdj24n34j2
	UniqSubnetIds []*string `json:"UniqSubnetIds,omitnil" name:"UniqSubnetIds"`

	// Array of region IDs (disused). The corresponding region can be queried through the common parameter `Region`.
	RegionIds []*int64 `json:"RegionIds,omitnil" name:"RegionIds"`

	// Instance status. <ul><li>`0`: Uninitialized. </li><li>`1`: In the process. </li><li>`2`: Running. </li><li>`-2`: Isolated. </li><li>`-3`: To be deleted. </li></ul>
	Status []*int64 `json:"Status,omitnil" name:"Status"`

	// Instance architecture. <ul><li>`1`: Standalone edition. </li><li>`2`: Master-replica edition. </li><li>`3`: Cluster edition. </li></ul>
	TypeVersion *int64 `json:"TypeVersion,omitnil" name:"TypeVersion"`

	// Storage engine information. Valid values: `Redis-2.8`, `Redis-4.0`, `Redis-5.0`, `Redis-6.0` or `CKV`.
	EngineName *string `json:"EngineName,omitnil" name:"EngineName"`

	// Renewal mode. Valid values:  <ul><li>`0`:  Manual renewal </li><li>`1`:  Auto-renewal </li><li>`2`:  No renewal upon expiration </ul>
	AutoRenew []*int64 `json:"AutoRenew,omitnil" name:"AutoRenew"`

	// Billing mode. Only pay-as-you-go billing is supported.
	BillingMode *string `json:"BillingMode,omitnil" name:"BillingMode"`

	// Instance type. Valid values:  - `2`: Redis 2.8 Memory Edition (Standard Architecture). - `3`: CKV 3.2 Memory Edition (Standard Architecture). - `4`: CKV 3.2 Memory Edition (Cluster Architecture). - `5`: Redis 2.8 Memory Edition (Standalone). - `6`: Redis 4.0 Memory Edition (Standard Architecture). - `7`: Redis 4.0 Memory Edition (Cluster Architecture). - `8`: Redis 5.0 Memory Edition (Standard Architecture). - `9`: Redis 5.0 Memory Edition (Cluster Architecture). - `15`: Redis 6.2 Memory Edition (Standard Architecture). - `16`: Redis 6.2 Memory Edition (Cluster Architecture).
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// Array of the search keywords, which can query the instance by its ID, name, IP address.
	SearchKeys []*string `json:"SearchKeys,omitnil" name:"SearchKeys"`

	// Internal parameter, which can be ignored.
	TypeList []*int64 `json:"TypeList,omitnil" name:"TypeList"`

	// Internal parameter, which can be ignored.
	MonitorVersion *string `json:"MonitorVersion,omitnil" name:"MonitorVersion"`

	// Resources filter by tag key and value. If this parameter is not specified or is an empty array, resources will not be filtered.
	InstanceTags []*InstanceTagInfo `json:"InstanceTags,omitnil" name:"InstanceTags"`

	// Resources filter by tag key. If this parameter is not specified or is an empty array, resources will not be filtered.
	TagKeys []*string `json:"TagKeys,omitnil" name:"TagKeys"`

	// Instance product version.  If this parameter is not passed in or the array is empty, the instances will not be filtered based this parameter by default.  <ul><li>`local`:  Local disk edition.  </li><li>`cdc`:  Dedicated cluster edition.  </li></ul>
	ProductVersions []*string `json:"ProductVersions,omitnil" name:"ProductVersions"`

	// Batch query of the specified instances ID. The number of results returned is based on `Limit`.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// AZ deployment mode. <ul><li>`singleaz`: Single-AZ. </li><li>`1`: Multi-AZ. </li></ul>
	AzMode *string `json:"AzMode,omitnil" name:"AzMode"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "InstanceId")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "VpcIds")
	delete(f, "SubnetIds")
	delete(f, "SearchKey")
	delete(f, "ProjectIds")
	delete(f, "InstanceName")
	delete(f, "UniqVpcIds")
	delete(f, "UniqSubnetIds")
	delete(f, "RegionIds")
	delete(f, "Status")
	delete(f, "TypeVersion")
	delete(f, "EngineName")
	delete(f, "AutoRenew")
	delete(f, "BillingMode")
	delete(f, "Type")
	delete(f, "SearchKeys")
	delete(f, "TypeList")
	delete(f, "MonitorVersion")
	delete(f, "InstanceTags")
	delete(f, "TagKeys")
	delete(f, "ProductVersions")
	delete(f, "InstanceIds")
	delete(f, "AzMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// Total number of instances
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// List of instance details
	InstanceSet []*InstanceSet `json:"InstanceSet,omitnil" name:"InstanceSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesResponseParams `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMaintenanceWindowRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeMaintenanceWindowRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeMaintenanceWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaintenanceWindowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMaintenanceWindowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMaintenanceWindowResponseParams struct {
	// Start time of the maintenance window, such as 17:00.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End time of the maintenance window, such as 19:00.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMaintenanceWindowResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMaintenanceWindowResponseParams `json:"Response"`
}

func (r *DescribeMaintenanceWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaintenanceWindowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplateInfoRequestParams struct {
	// The parameter template ID for query. Get parameter template list information through the [DescribeParamTemplates](https://intl.cloud.tencent.com/document/product/239/58750?from_cn_redirect=1) API.
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`
}

type DescribeParamTemplateInfoRequest struct {
	*tchttp.BaseRequest
	
	// The parameter template ID for query. Get parameter template list information through the [DescribeParamTemplates](https://intl.cloud.tencent.com/document/product/239/58750?from_cn_redirect=1) API.
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`
}

func (r *DescribeParamTemplateInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplateInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeParamTemplateInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplateInfoResponseParams struct {
	// Quantity of parameters in the parameter template
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Parameter template ID.
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// Parameter template name.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Product type
	// - `2`: Redis 2.8 Memory Edition (Standard Architecture).
	// - `3`: CKV 3.2 Memory Edition (Standard Architecture).
	// - `4`: CKV 3.2 Memory Edition (Cluster Architecture).
	// - `5`: Redis 2.8 Memory Edition (Standalone).
	// - `6`: Redis 4.0 Memory Edition (Standard Architecture).
	// - `7`: Redis 4.0 Memory Edition (Cluster Architecture).
	// - `8`: Redis 5.0 Memory Edition (Standard Architecture).
	// - `9`: Redis 5.0 Memory Edition (Cluster Architecture).
	// - `15`: Redis 6.2 Memory Edition (Standard Architecture).
	// - `16`: Redis 6.2 Memory Edition (Cluster Architecture).
	ProductType *uint64 `json:"ProductType,omitnil" name:"ProductType"`

	// Parameter template description
	Description *string `json:"Description,omitnil" name:"Description"`

	// Parameter details, including parameter name, current value, default value, maximum value, minimum value, enumeration value and other information.
	Items []*ParameterDetail `json:"Items,omitnil" name:"Items"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeParamTemplateInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeParamTemplateInfoResponseParams `json:"Response"`
}

func (r *DescribeParamTemplateInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplateInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplatesRequestParams struct {
	// Array of instance types. Valid values: `1` (Redis 2.8 Memory Edition in cluster architecture), `2` (Redis 2.8 Memory Edition in standard architecture), `3` (CKV 3.2 Memory Edition in standard architecture), `4` (CKV 3.2 Memory Edition in cluster architecture), `5` (Redis 2.8 Memory Edition in standalone architecture), `6` (Redis 4.0 Memory Edition in standard architecture), `7` (Redis 4.0 Memory Edition in cluster architecture), `8` (Redis 5.0 Memory Edition in standard architecture), `9` (Redis 5.0 Memory Edition in cluster architecture).
	ProductTypes []*int64 `json:"ProductTypes,omitnil" name:"ProductTypes"`

	// Array of template names.
	TemplateNames []*string `json:"TemplateNames,omitnil" name:"TemplateNames"`

	// Array of template IDs.
	TemplateIds []*string `json:"TemplateIds,omitnil" name:"TemplateIds"`
}

type DescribeParamTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Array of instance types. Valid values: `1` (Redis 2.8 Memory Edition in cluster architecture), `2` (Redis 2.8 Memory Edition in standard architecture), `3` (CKV 3.2 Memory Edition in standard architecture), `4` (CKV 3.2 Memory Edition in cluster architecture), `5` (Redis 2.8 Memory Edition in standalone architecture), `6` (Redis 4.0 Memory Edition in standard architecture), `7` (Redis 4.0 Memory Edition in cluster architecture), `8` (Redis 5.0 Memory Edition in standard architecture), `9` (Redis 5.0 Memory Edition in cluster architecture).
	ProductTypes []*int64 `json:"ProductTypes,omitnil" name:"ProductTypes"`

	// Array of template names.
	TemplateNames []*string `json:"TemplateNames,omitnil" name:"TemplateNames"`

	// Array of template IDs.
	TemplateIds []*string `json:"TemplateIds,omitnil" name:"TemplateIds"`
}

func (r *DescribeParamTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductTypes")
	delete(f, "TemplateNames")
	delete(f, "TemplateIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeParamTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplatesResponseParams struct {
	// Number of parameter templates of the user.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Parameter template details.
	Items []*ParamTemplateInfo `json:"Items,omitnil" name:"Items"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeParamTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeParamTemplatesResponseParams `json:"Response"`
}

func (r *DescribeParamTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductInfoRequestParams struct {

}

type DescribeProductInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeProductInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductInfoResponseParams struct {
	// Sale information of a region.
	RegionSet []*RegionConf `json:"RegionSet,omitnil" name:"RegionSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeProductInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProductInfoResponseParams `json:"Response"`
}

func (r *DescribeProductInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectSecurityGroupRequestParams struct {
	// 0: default project; -1: all projects; >0: specified project
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`
}

type DescribeProjectSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
	// 0: default project; -1: all projects; >0: specified project
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`
}

func (r *DescribeProjectSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "SecurityGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectSecurityGroupResponseParams struct {
	// Security group of the project
	SecurityGroupDetails []*SecurityGroupDetail `json:"SecurityGroupDetails,omitnil" name:"SecurityGroupDetails"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeProjectSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectSecurityGroupResponseParams `json:"Response"`
}

func (r *DescribeProjectSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectSecurityGroupsRequestParams struct {
	// Database engine name, which is `redis` for this API.
	Product *string `json:"Product,omitnil" name:"Product"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Offset, which is an integral multiple of `Limit`.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// The number of security groups to be pulled. Default value: `20`.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Search criteria. You can enter a security group ID or name.
	SearchKey *string `json:"SearchKey,omitnil" name:"SearchKey"`
}

type DescribeProjectSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Database engine name, which is `redis` for this API.
	Product *string `json:"Product,omitnil" name:"Product"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Offset, which is an integral multiple of `Limit`.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// The number of security groups to be pulled. Default value: `20`.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Search criteria. You can enter a security group ID or name.
	SearchKey *string `json:"SearchKey,omitnil" name:"SearchKey"`
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
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectSecurityGroupsResponseParams struct {
	// Security group rules.
	Groups []*SecurityGroup `json:"Groups,omitnil" name:"Groups"`

	// Total number of eligible security groups.
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
type DescribeProxySlowLogRequestParams struct {
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Start time of slow query
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// End time of slow query
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Slow query threshold  in milliseconds
	MinQueryTime *int64 `json:"MinQueryTime,omitnil" name:"MinQueryTime"`

	// Number of results per page.  Default value: `20`. Value range: [20,1000].
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Offset, which is an integral multiple of `Limit`.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeProxySlowLogRequest struct {
	*tchttp.BaseRequest
	
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Start time of slow query
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// End time of slow query
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Slow query threshold  in milliseconds
	MinQueryTime *int64 `json:"MinQueryTime,omitnil" name:"MinQueryTime"`

	// Number of results per page.  Default value: `20`. Value range: [20,1000].
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Offset, which is an integral multiple of `Limit`.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeProxySlowLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxySlowLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "MinQueryTime")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxySlowLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxySlowLogResponseParams struct {
	// Total number of slow queries
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Slow query details
	InstanceProxySlowLogDetail []*InstanceProxySlowlogDetail `json:"InstanceProxySlowLogDetail,omitnil" name:"InstanceProxySlowLogDetail"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeProxySlowLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProxySlowLogResponseParams `json:"Response"`
}

func (r *DescribeProxySlowLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxySlowLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReplicationGroupRequestParams struct {
	// Number of instances returned per page. Default value: `20`.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Pagination offset, which is an integral multiple of `Limit`. `offset` = `limit` * (page number - 1).
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// ID of the specified replication group, such as `crs-rpl-m3zt****`. Log in to the [TencentDB for Redis console](https://console.cloud.tencent.com/redis/replication), and get it in the global replication group list.
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// Keyword for fuzzy query, which can be a replication group ID or name. Log in to the [TencentDB for Redis console](https://console.cloud.tencent.com/redis/replication), and get them in the global replication group list.
	SearchKey *string `json:"SearchKey,omitnil" name:"SearchKey"`
}

type DescribeReplicationGroupRequest struct {
	*tchttp.BaseRequest
	
	// Number of instances returned per page. Default value: `20`.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Pagination offset, which is an integral multiple of `Limit`. `offset` = `limit` * (page number - 1).
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// ID of the specified replication group, such as `crs-rpl-m3zt****`. Log in to the [TencentDB for Redis console](https://console.cloud.tencent.com/redis/replication), and get it in the global replication group list.
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// Keyword for fuzzy query, which can be a replication group ID or name. Log in to the [TencentDB for Redis console](https://console.cloud.tencent.com/redis/replication), and get them in the global replication group list.
	SearchKey *string `json:"SearchKey,omitnil" name:"SearchKey"`
}

func (r *DescribeReplicationGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "GroupId")
	delete(f, "SearchKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReplicationGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReplicationGroupResponseParams struct {
	// Number of replication groups
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Replication group information
	Groups []*Groups `json:"Groups,omitnil" name:"Groups"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeReplicationGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReplicationGroupResponseParams `json:"Response"`
}

func (r *DescribeReplicationGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSSLStatusRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeSSLStatusRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeSSLStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSSLStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSSLStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSSLStatusResponseParams struct {
	// Download address for SSL certificate
	CertDownloadUrl *string `json:"CertDownloadUrl,omitnil" name:"CertDownloadUrl"`

	// Expiration time of the certificate download address
	UrlExpiredTime *string `json:"UrlExpiredTime,omitnil" name:"UrlExpiredTime"`

	// Whether the SSL is enabled for the identified instance.
	// - `true`: Enabled
	// - `false`: Disabled
	SSLConfig *bool `json:"SSLConfig,omitnil" name:"SSLConfig"`

	// Whether SSL is supported for the identified instance.
	// -`true`: Supported
	// -`false`: Not supported
	FeatureSupport *bool `json:"FeatureSupport,omitnil" name:"FeatureSupport"`

	// Status of SSL configuration
	// - `1`: Configuring
	// - `2`: Configured successfully
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSSLStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSSLStatusResponseParams `json:"Response"`
}

func (r *DescribeSSLStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSSLStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogRequestParams struct {
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Start time for prequerying a slow log
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// End time for prequerying a slow log
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// The average execution time threshold of slow query  in microseconds
	MinQueryTime *int64 `json:"MinQueryTime,omitnil" name:"MinQueryTime"`

	// Number of slow queries displayed per page. Default value: `20`. Value range:  [20,1000].
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Slow query offset, which is an integral multiple of `Limit`. Calculation formula:  `offset` = `limit` * (page number - 1).
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Node role. <ul><li>`Master`: Master node</li><li>`Slave`: Replica node</li></ul>
	Role *string `json:"Role,omitnil" name:"Role"`
}

type DescribeSlowLogRequest struct {
	*tchttp.BaseRequest
	
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Start time for prequerying a slow log
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// End time for prequerying a slow log
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// The average execution time threshold of slow query  in microseconds
	MinQueryTime *int64 `json:"MinQueryTime,omitnil" name:"MinQueryTime"`

	// Number of slow queries displayed per page. Default value: `20`. Value range:  [20,1000].
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Slow query offset, which is an integral multiple of `Limit`. Calculation formula:  `offset` = `limit` * (page number - 1).
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Node role. <ul><li>`Master`: Master node</li><li>`Slave`: Replica node</li></ul>
	Role *string `json:"Role,omitnil" name:"Role"`
}

func (r *DescribeSlowLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "MinQueryTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Role")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogResponseParams struct {
	// Total number of slow queries
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Slow query details
	InstanceSlowlogDetail []*InstanceSlowlogDetail `json:"InstanceSlowlogDetail,omitnil" name:"InstanceSlowlogDetail"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSlowLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogResponseParams `json:"Response"`
}

func (r *DescribeSlowLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskInfoRequestParams struct {
	// Task ID
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`
}

type DescribeTaskInfoRequest struct {
	*tchttp.BaseRequest
	
	// Task ID
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`
}

func (r *DescribeTaskInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskInfoResponseParams struct {
	// Task status. Valid values: 
	// - `preparing`: To be run
	// - `running`: Running
	// - `succeed`: Succeeded
	// - `failed`: Failed
	// - `Error`: Error occurred while running
	Status *string `json:"Status,omitnil" name:"Status"`

	// Task start time
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// Task type, including `Create`, `Configure`, `Disable Instance`, `Clear Instance`, `Reset Password`, `Upgrade Version`, `Back up Instance`, `Modify Network`, `Migrate to New AZ` and `Promote to Master`.
	TaskType *string `json:"TaskType,omitnil" name:"TaskType"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Message returned by task execution, which will be an error message when execution fails or be empty when the status is `running `or `succeed-`.
	TaskMessage *string `json:"TaskMessage,omitnil" name:"TaskMessage"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTaskInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskInfoResponseParams `json:"Response"`
}

func (r *DescribeTaskInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskListRequestParams struct {
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Number of taskss returned per page.  Default value: `20`. Maximum value: `100`.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Pagination offset, which is an integral multiple of `Limit`. Calculation formula:  `offset` = `limit` * (page number - 1).
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Project ID Log in to the [Redis console](https://console.cloud.tencent.com/redis#/), go to the account information menu in the top-right corner, and select **Project Management** to query the project ID.
	ProjectIds []*int64 `json:"ProjectIds,omitnil" name:"ProjectIds"`

	// Task type. Valid values:  - `FLOW_CREATE`: Create an instance. - `FLOW_MODIFYCONNECTIONCONFIG`: Adjust the number of bandwidth connections. - `FLOW_MODIFYINSTANCEPASSWORDFREE`: Modify the process of password-free access. - `FLOW_CLEARNETWORK`: Returning VPC - `FLOW_SETPWD`: Set the access password. - `FLOW_EXPORSHR`: Expand or reduce the capacity. - `FLOW_UpgradeArch`: Upgrade the instance architecture. - `FLOW_MODIFYINSTANCEPARAMS`: Modify the instance parameters. - `FLOW_MODIFYINSTACEREADONLY`: Modify read-only process. - `FLOW_CLOSE`: Disable the instance. - `FLOW_DELETE`: Delete the instance. - `FLOW_OPEN_WAN`: Enable the public network. - `FLOW_FLOW_CLEAN`: Clear the instance. - `FLOW_MODIFYINSTANCEACCOUNT`: Modify the instance account. - `FLOW_ENABLEINSTANCE_REPLICATE`: Enable the replica read-only feature. - `FLOW_DISABLEINSTANCE_REPLICATE`: Disable the replica read-only feature. - `FLOW_SWITCHINSTANCEVIP`: Swap the VIPs of instances. - FLOW_CHANGE_REPLICA_TO_MSTER: Promote the replica node to the mater node. - `FLOW_BACKUPINSTANCE`: Back up an instance.
	TaskTypes []*string `json:"TaskTypes,omitnil" name:"TaskTypes"`

	// Start time for executing a task,  in the format of  “2020-10-12 00:00:00”.
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// End time for executing a task,  in the format of  “2021-12-30 20:59:35”.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// This parameter is only for internal use and can be ignored.
	TaskStatus []*int64 `json:"TaskStatus,omitnil" name:"TaskStatus"`

	// Task execution status. Valid values: - `0` (initilized) - `1` (executing) - `2` (completed) - `4` (failed)
	Result []*int64 `json:"Result,omitnil" name:"Result"`

	// The field `OperatorUin` has been disused and replaced by `OperateUin`.
	OperatorUin []*int64 `json:"OperatorUin,omitnil" name:"OperatorUin"`

	// Operator account ID or UIN
	OperateUin []*string `json:"OperateUin,omitnil" name:"OperateUin"`
}

type DescribeTaskListRequest struct {
	*tchttp.BaseRequest
	
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Number of taskss returned per page.  Default value: `20`. Maximum value: `100`.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Pagination offset, which is an integral multiple of `Limit`. Calculation formula:  `offset` = `limit` * (page number - 1).
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Project ID Log in to the [Redis console](https://console.cloud.tencent.com/redis#/), go to the account information menu in the top-right corner, and select **Project Management** to query the project ID.
	ProjectIds []*int64 `json:"ProjectIds,omitnil" name:"ProjectIds"`

	// Task type. Valid values:  - `FLOW_CREATE`: Create an instance. - `FLOW_MODIFYCONNECTIONCONFIG`: Adjust the number of bandwidth connections. - `FLOW_MODIFYINSTANCEPASSWORDFREE`: Modify the process of password-free access. - `FLOW_CLEARNETWORK`: Returning VPC - `FLOW_SETPWD`: Set the access password. - `FLOW_EXPORSHR`: Expand or reduce the capacity. - `FLOW_UpgradeArch`: Upgrade the instance architecture. - `FLOW_MODIFYINSTANCEPARAMS`: Modify the instance parameters. - `FLOW_MODIFYINSTACEREADONLY`: Modify read-only process. - `FLOW_CLOSE`: Disable the instance. - `FLOW_DELETE`: Delete the instance. - `FLOW_OPEN_WAN`: Enable the public network. - `FLOW_FLOW_CLEAN`: Clear the instance. - `FLOW_MODIFYINSTANCEACCOUNT`: Modify the instance account. - `FLOW_ENABLEINSTANCE_REPLICATE`: Enable the replica read-only feature. - `FLOW_DISABLEINSTANCE_REPLICATE`: Disable the replica read-only feature. - `FLOW_SWITCHINSTANCEVIP`: Swap the VIPs of instances. - FLOW_CHANGE_REPLICA_TO_MSTER: Promote the replica node to the mater node. - `FLOW_BACKUPINSTANCE`: Back up an instance.
	TaskTypes []*string `json:"TaskTypes,omitnil" name:"TaskTypes"`

	// Start time for executing a task,  in the format of  “2020-10-12 00:00:00”.
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// End time for executing a task,  in the format of  “2021-12-30 20:59:35”.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// This parameter is only for internal use and can be ignored.
	TaskStatus []*int64 `json:"TaskStatus,omitnil" name:"TaskStatus"`

	// Task execution status. Valid values: - `0` (initilized) - `1` (executing) - `2` (completed) - `4` (failed)
	Result []*int64 `json:"Result,omitnil" name:"Result"`

	// The field `OperatorUin` has been disused and replaced by `OperateUin`.
	OperatorUin []*int64 `json:"OperatorUin,omitnil" name:"OperatorUin"`

	// Operator account ID or UIN
	OperateUin []*string `json:"OperateUin,omitnil" name:"OperateUin"`
}

func (r *DescribeTaskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ProjectIds")
	delete(f, "TaskTypes")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "TaskStatus")
	delete(f, "Result")
	delete(f, "OperatorUin")
	delete(f, "OperateUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskListResponseParams struct {
	// Total number of tasks
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Task details
	Tasks []*TaskInfoDetail `json:"Tasks,omitnil" name:"Tasks"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTaskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskListResponseParams `json:"Response"`
}

func (r *DescribeTaskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTendisSlowLogRequestParams struct {
	// Instance ID in the format of crs-ngvou0i1
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Start time in the format of 2019-09-08 12:12:41
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// End time in the format of 2019-09-09 12:12:41
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Slow query threshold in ms
	MinQueryTime *int64 `json:"MinQueryTime,omitnil" name:"MinQueryTime"`

	// Maximum number of results returned per page. Default value: 20.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Offset, which is an integral multiple of `Limit`.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeTendisSlowLogRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of crs-ngvou0i1
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Start time in the format of 2019-09-08 12:12:41
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// End time in the format of 2019-09-09 12:12:41
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Slow query threshold in ms
	MinQueryTime *int64 `json:"MinQueryTime,omitnil" name:"MinQueryTime"`

	// Maximum number of results returned per page. Default value: 20.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Offset, which is an integral multiple of `Limit`.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeTendisSlowLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTendisSlowLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "MinQueryTime")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTendisSlowLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTendisSlowLogResponseParams struct {
	// Total number of slow queries
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Slow query details
	TendisSlowLogDetail []*TendisSlowLogDetail `json:"TendisSlowLogDetail,omitnil" name:"TendisSlowLogDetail"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTendisSlowLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTendisSlowLogResponseParams `json:"Response"`
}

func (r *DescribeTendisSlowLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTendisSlowLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyPostpaidInstanceRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DestroyPostpaidInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DestroyPostpaidInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyPostpaidInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyPostpaidInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyPostpaidInstanceResponseParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DestroyPostpaidInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DestroyPostpaidInstanceResponseParams `json:"Response"`
}

func (r *DestroyPostpaidInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyPostpaidInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyPrepaidInstanceRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DestroyPrepaidInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DestroyPrepaidInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyPrepaidInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyPrepaidInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyPrepaidInstanceResponseParams struct {
	// Order ID
	DealId *string `json:"DealId,omitnil" name:"DealId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DestroyPrepaidInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DestroyPrepaidInstanceResponseParams `json:"Response"`
}

func (r *DestroyPrepaidInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyPrepaidInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableReplicaReadonlyRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DisableReplicaReadonlyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DisableReplicaReadonlyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableReplicaReadonlyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableReplicaReadonlyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableReplicaReadonlyResponseParams struct {
	// Task ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DisableReplicaReadonlyResponse struct {
	*tchttp.BaseResponse
	Response *DisableReplicaReadonlyResponseParams `json:"Response"`
}

func (r *DisableReplicaReadonlyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableReplicaReadonlyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateSecurityGroupsRequestParams struct {
	// Database engine name, which is `redis` for this API.
	Product *string `json:"Product,omitnil" name:"Product"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// List of instance IDs, which is an array of one or more instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Database engine name, which is `redis` for this API.
	Product *string `json:"Product,omitnil" name:"Product"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// List of instance IDs, which is an array of one or more instance IDs.
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

// Predefined struct for user
type EnableReplicaReadonlyRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Account routing policy. If `master` or `replication` is entered, it means to route to the master or replica node; if this parameter is left empty, it means to write into the master node and read from the replica node by default.
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitnil" name:"ReadonlyPolicy"`
}

type EnableReplicaReadonlyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Account routing policy. If `master` or `replication` is entered, it means to route to the master or replica node; if this parameter is left empty, it means to write into the master node and read from the replica node by default.
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitnil" name:"ReadonlyPolicy"`
}

func (r *EnableReplicaReadonlyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableReplicaReadonlyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ReadonlyPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableReplicaReadonlyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableReplicaReadonlyResponseParams struct {
	// Valid values: `ERROR`, `OK`. This field has been disused.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil" name:"Status"`

	// Task ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type EnableReplicaReadonlyResponse struct {
	*tchttp.BaseResponse
	Response *EnableReplicaReadonlyResponseParams `json:"Response"`
}

func (r *EnableReplicaReadonlyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableReplicaReadonlyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Groups struct {
	// User APPID, which is the unique application ID that matches an account. Some Tencent Cloud products use this APPID.
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`

	// Region ID. Valid values:
	// - `1`: Guangzhou 
	// - `4`: Shanghai 
	// - `5`: Hong Kong (China) 
	// - `6`: Toronto 
	// - `7`: Shanghai Finance 
	// - `8`: Beijing 
	// - `9`: Singapore
	// - `11`: Shenzhen Finance
	// - `15`: Silicon Valley (West US)
	// - `16`: Chengdu 
	// - `17`: Germany 
	// - `18`: South Korea 
	// - `19`: Chongqing 
	// - `21`: India 
	// - `22`: Virginia (East US)
	// - `23`: Thailand 
	// - `25`: Japan
	RegionId *int64 `json:"RegionId,omitnil" name:"RegionId"`

	// Replication group ID in the format of "crs-rpl-deind****"
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// Replication group name
	// Note: This field may return null, indicating that no valid values can be obtained.
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// Status of replication group
	// - `37`: Associating replication group
	// - `38`: Reconnecting to replication group
	// - `51`: Disassociating replication group
	// - `52`: Switching with master instance in replication group
	// - `53`: Modifying the roles
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// Number of replication groups
	InstanceCount *int64 `json:"InstanceCount,omitnil" name:"InstanceCount"`

	// Instance information in replication groups
	// Note: This field may return null, indicating that no valid values can be obtained.
	Instances []*Instances `json:"Instances,omitnil" name:"Instances"`

	// Remarks
	// Note: This field may return null, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

type HotKeyInfo struct {
	// Hot key
	Key *string `json:"Key,omitnil" name:"Key"`

	// Type
	Type *string `json:"Type,omitnil" name:"Type"`

	// Count
	Count *int64 `json:"Count,omitnil" name:"Count"`
}

type Inbound struct {
	// Policy. Valid values: ACCEPT, DROP.
	Action *string `json:"Action,omitnil" name:"Action"`

	// All the addresses that the address group ID represents.
	AddressModule *string `json:"AddressModule,omitnil" name:"AddressModule"`

	// Source IP or IP address range, such as 192.168.0.0/16.
	CidrIp *string `json:"CidrIp,omitnil" name:"CidrIp"`

	// Description.
	Desc *string `json:"Desc,omitnil" name:"Desc"`

	// Network protocol, such as UDP and TCP.
	IpProtocol *string `json:"IpProtocol,omitnil" name:"IpProtocol"`

	// Port.
	PortRange *string `json:"PortRange,omitnil" name:"PortRange"`

	// All the protocols and ports that the service group ID represents.
	ServiceModule *string `json:"ServiceModule,omitnil" name:"ServiceModule"`

	// All the addresses that the security group ID represents.
	Id *string `json:"Id,omitnil" name:"Id"`
}

// Predefined struct for user
type InquiryPriceCreateInstanceRequestParams struct {
	// Instance type. Valid values: `2` (Redis 2.8 memory edition in standard architecture), `3` (CKV 3.2 memory edition in standard architecture), `4` (CKV 3.2 memory edition in cluster architecture), `6` (Redis 4.0 memory edition in standard architecture), `7` (Redis 4.0 memory edition in cluster architecture), `8` (Redis 5.0 memory edition in standard architecture), `9` (Redis 5.0 memory edition in cluster architecture).
	TypeId *uint64 `json:"TypeId,omitnil" name:"TypeId"`

	// Memory capacity in MB, which must be a multiple of 1,024. It is subject to the purchasable specifications returned by the [DescribeProductInfo API](https://intl.cloud.tencent.com/document/api/239/30600?from_cn_redirect=1).
	// If `TypeId` indicates the standard architecture, `MemSize` indicates the total memory capacity of an instance; if `TypeId` indicates the cluster architecture, `MemSize` indicates the memory capacity per shard.
	MemSize *uint64 `json:"MemSize,omitnil" name:"MemSize"`

	// Number of instances. The actual quantity purchasable at a time is subject to the specifications returned by the [DescribeProductInfo API](https://intl.cloud.tencent.com/document/api/239/30600?from_cn_redirect=1).
	GoodsNum *uint64 `json:"GoodsNum,omitnil" name:"GoodsNum"`

	// Length of purchase in months, which is required when creating a monthly-subscribed instance. Value range: [1,2,3,4,5,6,7,8,9,10,11,12,24,36]. For pay-as-you-go instances, set the parameter to `1`.
	Period *uint64 `json:"Period,omitnil" name:"Period"`

	// Billing mode. Valid values: `0` (pay-as-you-go), `1` (monthly subscription).
	BillingMode *int64 `json:"BillingMode,omitnil" name:"BillingMode"`

	// ID of the AZ where the instance resides. For more information, see [Regions and AZs](https://intl.cloud.tencent.com/document/product/239/4106?from_cn_redirect=1).
	ZoneId *uint64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// Instance shard quantity. This field is not required by Redis 2.8 standard architecture, CKV standard architecture, Redis 2.8 standalone edition, and Redis 4.0 standard architecture.
	RedisShardNum *int64 `json:"RedisShardNum,omitnil" name:"RedisShardNum"`

	// Instance replica quantity. This field is not required by Redis 2.8 standard architecture, CKV standard architecture, and Redis 2.8 standalone edition.
	RedisReplicasNum *int64 `json:"RedisReplicasNum,omitnil" name:"RedisReplicasNum"`

	// Whether to support read-only replicas. This field is not required by Redis 2.8 standard architecture, CKV standard architecture, and Redis 2.8 standalone edition.
	ReplicasReadonly *bool `json:"ReplicasReadonly,omitnil" name:"ReplicasReadonly"`

	// Name of the AZ where the instance resides. For more information, see [Regions and AZs](https://intl.cloud.tencent.com/document/product/239/4106?from_cn_redirect=1).
	ZoneName *string `json:"ZoneName,omitnil" name:"ZoneName"`

	// Valid values: `local` (local disk edition), `cloud` (cloud disk edition), `cdc` (dedicated cluster edition). Default value: `local` (local disk edition)
	ProductVersion *string `json:"ProductVersion,omitnil" name:"ProductVersion"`
}

type InquiryPriceCreateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance type. Valid values: `2` (Redis 2.8 memory edition in standard architecture), `3` (CKV 3.2 memory edition in standard architecture), `4` (CKV 3.2 memory edition in cluster architecture), `6` (Redis 4.0 memory edition in standard architecture), `7` (Redis 4.0 memory edition in cluster architecture), `8` (Redis 5.0 memory edition in standard architecture), `9` (Redis 5.0 memory edition in cluster architecture).
	TypeId *uint64 `json:"TypeId,omitnil" name:"TypeId"`

	// Memory capacity in MB, which must be a multiple of 1,024. It is subject to the purchasable specifications returned by the [DescribeProductInfo API](https://intl.cloud.tencent.com/document/api/239/30600?from_cn_redirect=1).
	// If `TypeId` indicates the standard architecture, `MemSize` indicates the total memory capacity of an instance; if `TypeId` indicates the cluster architecture, `MemSize` indicates the memory capacity per shard.
	MemSize *uint64 `json:"MemSize,omitnil" name:"MemSize"`

	// Number of instances. The actual quantity purchasable at a time is subject to the specifications returned by the [DescribeProductInfo API](https://intl.cloud.tencent.com/document/api/239/30600?from_cn_redirect=1).
	GoodsNum *uint64 `json:"GoodsNum,omitnil" name:"GoodsNum"`

	// Length of purchase in months, which is required when creating a monthly-subscribed instance. Value range: [1,2,3,4,5,6,7,8,9,10,11,12,24,36]. For pay-as-you-go instances, set the parameter to `1`.
	Period *uint64 `json:"Period,omitnil" name:"Period"`

	// Billing mode. Valid values: `0` (pay-as-you-go), `1` (monthly subscription).
	BillingMode *int64 `json:"BillingMode,omitnil" name:"BillingMode"`

	// ID of the AZ where the instance resides. For more information, see [Regions and AZs](https://intl.cloud.tencent.com/document/product/239/4106?from_cn_redirect=1).
	ZoneId *uint64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// Instance shard quantity. This field is not required by Redis 2.8 standard architecture, CKV standard architecture, Redis 2.8 standalone edition, and Redis 4.0 standard architecture.
	RedisShardNum *int64 `json:"RedisShardNum,omitnil" name:"RedisShardNum"`

	// Instance replica quantity. This field is not required by Redis 2.8 standard architecture, CKV standard architecture, and Redis 2.8 standalone edition.
	RedisReplicasNum *int64 `json:"RedisReplicasNum,omitnil" name:"RedisReplicasNum"`

	// Whether to support read-only replicas. This field is not required by Redis 2.8 standard architecture, CKV standard architecture, and Redis 2.8 standalone edition.
	ReplicasReadonly *bool `json:"ReplicasReadonly,omitnil" name:"ReplicasReadonly"`

	// Name of the AZ where the instance resides. For more information, see [Regions and AZs](https://intl.cloud.tencent.com/document/product/239/4106?from_cn_redirect=1).
	ZoneName *string `json:"ZoneName,omitnil" name:"ZoneName"`

	// Valid values: `local` (local disk edition), `cloud` (cloud disk edition), `cdc` (dedicated cluster edition). Default value: `local` (local disk edition)
	ProductVersion *string `json:"ProductVersion,omitnil" name:"ProductVersion"`
}

func (r *InquiryPriceCreateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TypeId")
	delete(f, "MemSize")
	delete(f, "GoodsNum")
	delete(f, "Period")
	delete(f, "BillingMode")
	delete(f, "ZoneId")
	delete(f, "RedisShardNum")
	delete(f, "RedisReplicasNum")
	delete(f, "ReplicasReadonly")
	delete(f, "ZoneName")
	delete(f, "ProductVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceCreateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceCreateInstanceResponseParams struct {
	// Price. Unit: USD (accurate down to the cent)
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Price *float64 `json:"Price,omitnil" name:"Price"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type InquiryPriceCreateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceCreateInstanceResponseParams `json:"Response"`
}

func (r *InquiryPriceCreateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceUpgradeInstanceRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Shard size in MB.
	MemSize *uint64 `json:"MemSize,omitnil" name:"MemSize"`

	// Number of shards. This parameter can be left blank for Redis 2.8 in standard architecture, CKV in standard architecture, and Redis 2.8 in standalone architecture.
	RedisShardNum *uint64 `json:"RedisShardNum,omitnil" name:"RedisShardNum"`

	// Number of replicas. This parameter can be left blank for Redis 2.8 in standard architecture, CKV in standard architecture, and Redis 2.8 in standalone architecture.
	RedisReplicasNum *uint64 `json:"RedisReplicasNum,omitnil" name:"RedisReplicasNum"`
}

type InquiryPriceUpgradeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Shard size in MB.
	MemSize *uint64 `json:"MemSize,omitnil" name:"MemSize"`

	// Number of shards. This parameter can be left blank for Redis 2.8 in standard architecture, CKV in standard architecture, and Redis 2.8 in standalone architecture.
	RedisShardNum *uint64 `json:"RedisShardNum,omitnil" name:"RedisShardNum"`

	// Number of replicas. This parameter can be left blank for Redis 2.8 in standard architecture, CKV in standard architecture, and Redis 2.8 in standalone architecture.
	RedisReplicasNum *uint64 `json:"RedisReplicasNum,omitnil" name:"RedisReplicasNum"`
}

func (r *InquiryPriceUpgradeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceUpgradeInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "MemSize")
	delete(f, "RedisShardNum")
	delete(f, "RedisReplicasNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceUpgradeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceUpgradeInstanceResponseParams struct {
	// Price. Unit: USD
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Price *float64 `json:"Price,omitnil" name:"Price"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type InquiryPriceUpgradeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceUpgradeInstanceResponseParams `json:"Response"`
}

func (r *InquiryPriceUpgradeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceUpgradeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceClusterNode struct {
	// Node name
	Name *string `json:"Name,omitnil" name:"Name"`

	// ID of the runtime node of an instance
	RunId *string `json:"RunId,omitnil" name:"RunId"`

	// Cluster role. Valid values:  - `0` (master) - `1` (replica)
	Role *int64 `json:"Role,omitnil" name:"Role"`

	// Node status. Valid values:  - `0` (read/write) - `1` (read) - `2` (backup)
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// Service status. Valid values: `0` (down), `1` (on).
	Connected *int64 `json:"Connected,omitnil" name:"Connected"`

	// Node creation time
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Node elimination time
	DownTime *string `json:"DownTime,omitnil" name:"DownTime"`

	// Node slot distribution range
	Slots *string `json:"Slots,omitnil" name:"Slots"`

	// Distribution of node keys
	Keys *int64 `json:"Keys,omitnil" name:"Keys"`

	// Node QPS Number of executions per second on sharded nodes Unit: Counts/sec
	Qps *int64 `json:"Qps,omitnil" name:"Qps"`

	// QPS slope of a node
	QpsSlope *float64 `json:"QpsSlope,omitnil" name:"QpsSlope"`

	// Node storage
	Storage *int64 `json:"Storage,omitnil" name:"Storage"`

	// Node storage slope
	StorageSlope *float64 `json:"StorageSlope,omitnil" name:"StorageSlope"`
}

type InstanceClusterShard struct {
	// The name of a shard node
	ShardName *string `json:"ShardName,omitnil" name:"ShardName"`

	// The serial number of a shard node
	ShardId *string `json:"ShardId,omitnil" name:"ShardId"`

	// The role of a shard node
	// - `0`: Master node.
	// - `1`: Replica node.
	Role *int64 `json:"Role,omitnil" name:"Role"`

	// Number of keys
	Keys *int64 `json:"Keys,omitnil" name:"Keys"`

	// Slot information
	Slots *string `json:"Slots,omitnil" name:"Slots"`

	// Used Capacity
	Storage *int64 `json:"Storage,omitnil" name:"Storage"`

	// Capacity slope
	StorageSlope *float64 `json:"StorageSlope,omitnil" name:"StorageSlope"`

	// Instance runtime node ID
	Runid *string `json:"Runid,omitnil" name:"Runid"`

	// Service status
	// - `0`: Down.
	// - `1`: On.
	Connected *int64 `json:"Connected,omitnil" name:"Connected"`
}

type InstanceEnumParam struct {
	// Parameter name
	ParamName *string `json:"ParamName,omitnil" name:"ParamName"`

	// Parameter type, such as `Enum`.
	ValueType *string `json:"ValueType,omitnil" name:"ValueType"`

	// Whether to restart the database after modifying the parameters. Valid values: - `true` (required) - `false` (not required)
	NeedRestart *string `json:"NeedRestart,omitnil" name:"NeedRestart"`

	// Default value of the parameter
	DefaultValue *string `json:"DefaultValue,omitnil" name:"DefaultValue"`

	// Current value
	CurrentValue *string `json:"CurrentValue,omitnil" name:"CurrentValue"`

	// Description
	Tips *string `json:"Tips,omitnil" name:"Tips"`

	// Acceptable values for the parameter
	EnumValue []*string `json:"EnumValue,omitnil" name:"EnumValue"`

	// Parameter modification status. Valid values: - `1` (modifying) - `2` (modified)
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

type InstanceIntegerParam struct {
	// Parameter name
	ParamName *string `json:"ParamName,omitnil" name:"ParamName"`

	// Parameter type: Integer
	ValueType *string `json:"ValueType,omitnil" name:"ValueType"`

	// Whether restart is required after a modification is made. Valid values: true, false
	NeedRestart *string `json:"NeedRestart,omitnil" name:"NeedRestart"`

	// Default value of the parameter
	DefaultValue *string `json:"DefaultValue,omitnil" name:"DefaultValue"`

	// Current value
	CurrentValue *string `json:"CurrentValue,omitnil" name:"CurrentValue"`

	// Parameter description
	Tips *string `json:"Tips,omitnil" name:"Tips"`

	// Minimum value of the parameter
	Min *string `json:"Min,omitnil" name:"Min"`

	// Maximum value of the parameter
	Max *string `json:"Max,omitnil" name:"Max"`

	// Parameter status. 1: modifying; 2: modified
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// Parameter unit
	// Note: This field may return null, indicating that no valid values can be obtained.
	Unit *string `json:"Unit,omitnil" name:"Unit"`
}

type InstanceMultiParam struct {
	// Parameter name
	ParamName *string `json:"ParamName,omitnil" name:"ParamName"`

	// Parameter Type such as  `MULTI`
	ValueType *string `json:"ValueType,omitnil" name:"ValueType"`

	// Whether to restart the database after modifying the parameter. Valid values:  - `true` (required) - `false` (not required)
	NeedRestart *string `json:"NeedRestart,omitnil" name:"NeedRestart"`

	// Default value of the parameter
	DefaultValue *string `json:"DefaultValue,omitnil" name:"DefaultValue"`

	// Current value
	CurrentValue *string `json:"CurrentValue,omitnil" name:"CurrentValue"`

	// Description
	Tips *string `json:"Tips,omitnil" name:"Tips"`

	// Description
	EnumValue []*string `json:"EnumValue,omitnil" name:"EnumValue"`

	// Parameter modification status. Valid values: - `1` (modifying) - `2` (modified)
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

type InstanceNode struct {
	// Instance ID
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// Node details
	InstanceClusterNode []*InstanceClusterNode `json:"InstanceClusterNode,omitnil" name:"InstanceClusterNode"`
}

type InstanceParam struct {
	// Parameter name, such as “timeout”. For supported custom parameters, see <a href="https://www.tencentcloud.com/document/product/239/39796">Setting Instance Parameters</a>
	Key *string `json:"Key,omitnil" name:"Key"`

	// Current parameter value. For example, if you set the current value of “timeout” to 120 (in seconds), the client connections that remain idle longer than 120 seconds will be closed. For more information on parameter values, see <a href="https://www.tencentcloud.com/document/product/239/39796">Setting Instance Parameters</a>
	Value *string `json:"Value,omitnil" name:"Value"`
}

type InstanceParamHistory struct {
	// Parameter name
	ParamName *string `json:"ParamName,omitnil" name:"ParamName"`

	// The value of the parameter before modification
	PreValue *string `json:"PreValue,omitnil" name:"PreValue"`

	// The value of the parameter after modification
	NewValue *string `json:"NewValue,omitnil" name:"NewValue"`

	// Parameter configuration status
	// - `1`: The parameter configuration is being modified.
	// - `2`: The parameter configuration has been modified successfully.
	// - `3`: Failed to modify the parameter configuration.
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// Modification time
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`
}

type InstanceProxySlowlogDetail struct {
	// Slow query duration in milliseconds
	Duration *int64 `json:"Duration,omitnil" name:"Duration"`

	// Client address
	Client *string `json:"Client,omitnil" name:"Client"`

	// Slow query command
	Command *string `json:"Command,omitnil" name:"Command"`

	// Detailed command line information of slow query
	CommandLine *string `json:"CommandLine,omitnil" name:"CommandLine"`

	// Execution time
	ExecuteTime *string `json:"ExecuteTime,omitnil" name:"ExecuteTime"`
}

type InstanceSecurityGroupDetail struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Security group information, which includes  security group ID, name, outbound and inbound rules.
	SecurityGroupDetails []*SecurityGroupDetail `json:"SecurityGroupDetails,omitnil" name:"SecurityGroupDetails"`
}

type InstanceSet struct {
	// Instance name
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// User APPID, which is the unique application ID that matches an account. Some Tencent Cloud products use this APPID.
	Appid *int64 `json:"Appid,omitnil" name:"Appid"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Region ID. <ul><li>`1`: Guangzhou. </li><li>`4`: Shanghai. </li><li>`5`: Hong Kong (China). </li><li>`6`: Toronto. </li> <li>`7`: Shanghai Finance. </li> <li>`8`: Beijing. </li> <li>`9`: Singapore. </li> <li>`11`: Shenzhen Finance. </li> <li>`15`: West US (Silicon Valley). </li><li>`16`: Chengdu. </li><li>`17`: Frankfurt. </li><li>`18`: Seoul. </li><li>`19`: Chongqing. </li><li>`21`: Mumbai. </li><li>`22`: East US (Virginia). </li><li>`23`: Bangkok. </li><li>`24`: Moscow. </li><li>`25`: Tokyo. </li></ul>
	RegionId *int64 `json:"RegionId,omitnil" name:"RegionId"`

	// Region ID
	ZoneId *int64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// VPC ID, such as `75101`.
	VpcId *int64 `json:"VpcId,omitnil" name:"VpcId"`

	// Subnet ID, such as `46315`.
	SubnetId *int64 `json:"SubnetId,omitnil" name:"SubnetId"`

	// Current instance status. <ul><li>`0`: To be initialized. </li><li>`1`: In the process. </li><li>`2`: Running. </li><li>`-2`: Isolated. </li><li>`-3`: To be deleted. </li></ul>
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// Instance VIP
	WanIp *string `json:"WanIp,omitnil" name:"WanIp"`

	// Port number of an instance
	Port *int64 `json:"Port,omitnil" name:"Port"`

	// Instance creation time in the format of "2020-01-15 10:20:00"
	Createtime *string `json:"Createtime,omitnil" name:"Createtime"`

	// Instance memory capacity in MB (1 MB = 1024 KB)
	Size *float64 `json:"Size,omitnil" name:"Size"`

	// This field has been disused. You can use the TCOP’s [GetMonitorData](https://intl.cloud.tencent.com/document/product/248/31014?from_cn_redirect=1) API to query the capacity used by the instance.
	SizeUsed *float64 `json:"SizeUsed,omitnil" name:"SizeUsed"`

	// Instance type
	// - `2`: Redis 2.8 Memory Edition (Standard Architecture).
	// - `3`: CKV 3.2 Memory Edition (Standard Architecture).
	// - `4`: CKV 3.2 Memory Edition (Cluster Architecture).
	// - `5`: Redis 2.8 Memory Edition (Standalone).
	// - `6`: Redis 4.0 Memory Edition (Standard Architecture).
	// - `7`: Redis 4.0 Memory Edition (Cluster Architecture).
	// - `8`: Redis 5.0 Memory Edition (Standard Architecture).
	// - `9`: Redis 5.0 Memory Edition (Cluster Architecture).
	// - `15`: Redis 6.2 Memory Edition (Standard Architecture).
	// - `16`: Redis 6.2 Memory Edition (Cluster Architecture).
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// Whether to set the auto-renewal flag for an instance. <ul><li>`1`: Auto-renewal set. </li><li>`0`: Auto-renewal not set.</li></ul>
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// The time when a monthly subscribed instance expires
	DeadlineTime *string `json:"DeadlineTime,omitnil" name:"DeadlineTime"`

	// Engine: Redis community edition, Tencent Cloud CKV
	Engine *string `json:"Engine,omitnil" name:"Engine"`

	// Product type. <ul><li>`standalone`: Standard edition. </li><li>`cluster`: Cluster edition. </li></ul>
	ProductType *string `json:"ProductType,omitnil" name:"ProductType"`

	// VPC ID, such as vpc-fk33jsf43kgv.
	UniqVpcId *string `json:"UniqVpcId,omitnil" name:"UniqVpcId"`

	// VPC subnet ID, such as subnet-fd3j6l35mm0.
	UniqSubnetId *string `json:"UniqSubnetId,omitnil" name:"UniqSubnetId"`

	// Billing mode. Only pay-as-you-go billing is supported.
	BillingMode *int64 `json:"BillingMode,omitnil" name:"BillingMode"`

	// Description of an instance status, such as "Running".
	InstanceTitle *string `json:"InstanceTitle,omitnil" name:"InstanceTitle"`

	// The default termination time for isolated instances in the format of "2020-02-15 10:20:00". By default, a pay-as-you-go instance will be terminated after two hours of isolation, and a monthly subscribed instance will be terminated after seven days by default.
	OfflineTime *string `json:"OfflineTime,omitnil" name:"OfflineTime"`

	// Sub-status returned for an instance in process
	SubStatus *int64 `json:"SubStatus,omitnil" name:"SubStatus"`

	// Anti-affinity tag
	Tags []*string `json:"Tags,omitnil" name:"Tags"`

	// Instance node information
	InstanceNode []*InstanceNode `json:"InstanceNode,omitnil" name:"InstanceNode"`

	// Shard size
	RedisShardSize *int64 `json:"RedisShardSize,omitnil" name:"RedisShardSize"`

	// Number of shards
	RedisShardNum *int64 `json:"RedisShardNum,omitnil" name:"RedisShardNum"`

	// Number of replicas
	RedisReplicasNum *int64 `json:"RedisReplicasNum,omitnil" name:"RedisReplicasNum"`

	// Billing ID
	PriceId *int64 `json:"PriceId,omitnil" name:"PriceId"`

	// The time when an instance start to be isolated
	CloseTime *string `json:"CloseTime,omitnil" name:"CloseTime"`

	// Read weight of a replica node
	SlaveReadWeight *int64 `json:"SlaveReadWeight,omitnil" name:"SlaveReadWeight"`

	// Instance tag information
	// Note: This field may return null, indicating that no valid value can be obtained.
	InstanceTags []*InstanceTagInfo `json:"InstanceTags,omitnil" name:"InstanceTags"`

	// Project name
	// Note: This field may return null, indicating that no valid value can be obtained.
	ProjectName *string `json:"ProjectName,omitnil" name:"ProjectName"`

	// Whether an instance is password-free. <ul><li>`true`: Yes. </li><li>`false`: No. </li></ul>
	// Note: This field may return null, indicating that no valid value can be obtained.
	NoAuth *bool `json:"NoAuth,omitnil" name:"NoAuth"`

	// Number of client connections
	// Note: This field may return null, indicating that no valid value can be obtained.
	ClientLimit *int64 `json:"ClientLimit,omitnil" name:"ClientLimit"`

	// DTS status (internal parameter, which can be ignored)
	// Note: This field may return null, indicating that no valid value can be obtained.
	DtsStatus *int64 `json:"DtsStatus,omitnil" name:"DtsStatus"`

	// Upper shard bandwidth limit in MB
	// Note: This field may return null, indicating that no valid value can be obtained.
	NetLimit *int64 `json:"NetLimit,omitnil" name:"NetLimit"`

	// Password-free instance flag (internal parameter, which can be ignored)
	// Note: This field may return null, indicating that no valid value can be obtained.
	PasswordFree *int64 `json:"PasswordFree,omitnil" name:"PasswordFree"`

	// Internal parameter, which can be ignored.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Vip6 *string `json:"Vip6,omitnil" name:"Vip6"`

	// Read-only instance flag (internal parameter, which can be ignored)
	// Note: This field may return null, indicating that no valid value can be obtained.
	ReadOnly *int64 `json:"ReadOnly,omitnil" name:"ReadOnly"`

	// Internal parameter, which can be ignored.
	// Note: This field may return null, indicating that no valid value can be obtained.
	RemainBandwidthDuration *string `json:"RemainBandwidthDuration,omitnil" name:"RemainBandwidthDuration"`

	// This parameter can be ignored for Redis instance.
	// Note: This field may return null, indicating that no valid value can be obtained.
	DiskSize *int64 `json:"DiskSize,omitnil" name:"DiskSize"`

	// Monitoring granularity. <ul><li>`1m`: Monitoring at one-minute granularity. This granularity has been disused. For more information, see [1-Minute Granularity Will Be Disused](https://www.tencentcloud.com/document/product/239/50440).</li><li>`5s`: Monitoring at five-second granularity.</li></ul>
	// Note: This field may return null, indicating that no valid values can be obtained.
	MonitorVersion *string `json:"MonitorVersion,omitnil" name:"MonitorVersion"`

	// The minimum number of max client connections
	// Note: This field may return null, indicating that no valid value can be obtained.
	ClientLimitMin *int64 `json:"ClientLimitMin,omitnil" name:"ClientLimitMin"`

	// The maximum number of max client connections
	// Note: This field may return null, indicating that no valid value can be obtained.
	ClientLimitMax *int64 `json:"ClientLimitMax,omitnil" name:"ClientLimitMax"`

	// Instance node details
	// Note: This field may return null, indicating that no valid value can be obtained.
	NodeSet []*RedisNodeInfo `json:"NodeSet,omitnil" name:"NodeSet"`

	// Information of the region where the instance is deployed, such as `ap-guangzhou`.
	// Note: This field may return null, indicating that no valid value can be obtained.
	Region *string `json:"Region,omitnil" name:"Region"`

	// Public IP
	// Note: This field may return null, indicating that no valid value can be obtained.
	WanAddress *string `json:"WanAddress,omitnil" name:"WanAddress"`

	// Polaris service address, which is for internal use.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PolarisServer *string `json:"PolarisServer,omitnil" name:"PolarisServer"`

	// The current proxy version of an instance
	// Note: This field may return null, indicating that no valid value can be obtained.
	CurrentProxyVersion *string `json:"CurrentProxyVersion,omitnil" name:"CurrentProxyVersion"`

	// The current cache minor version of an instance
	// Note: This field may return null, indicating that no valid value can be obtained.
	CurrentRedisVersion *string `json:"CurrentRedisVersion,omitnil" name:"CurrentRedisVersion"`

	// Proxy version, which can be upgraded for the instance
	// Note: This field may return null, indicating that no valid value can be obtained.
	UpgradeProxyVersion *string `json:"UpgradeProxyVersion,omitnil" name:"UpgradeProxyVersion"`

	// Cache minor version, which can be upgraded for the instance
	// Note: This field may return null, indicating that no valid value can be obtained.
	UpgradeRedisVersion *string `json:"UpgradeRedisVersion,omitnil" name:"UpgradeRedisVersion"`
}

type InstanceSlowlogDetail struct {
	// Slow log duration
	Duration *int64 `json:"Duration,omitnil" name:"Duration"`

	// Client address
	Client *string `json:"Client,omitnil" name:"Client"`

	// Command
	Command *string `json:"Command,omitnil" name:"Command"`

	// Command line details
	CommandLine *string `json:"CommandLine,omitnil" name:"CommandLine"`

	// Execution duration
	ExecuteTime *string `json:"ExecuteTime,omitnil" name:"ExecuteTime"`

	// Node ID
	Node *string `json:"Node,omitnil" name:"Node"`
}

type InstanceTagInfo struct {
	// Tag key
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}

type InstanceTextParam struct {
	// Parameter name
	ParamName *string `json:"ParamName,omitnil" name:"ParamName"`

	// Parameter type such as  `Text`.
	ValueType *string `json:"ValueType,omitnil" name:"ValueType"`

	// Whether to restart the database after modifying the parameter. Valid values:  - `true` (required) - `false` (not required)
	NeedRestart *string `json:"NeedRestart,omitnil" name:"NeedRestart"`

	// Default value of the parameter
	DefaultValue *string `json:"DefaultValue,omitnil" name:"DefaultValue"`

	// Current value
	CurrentValue *string `json:"CurrentValue,omitnil" name:"CurrentValue"`

	// Description
	Tips *string `json:"Tips,omitnil" name:"Tips"`

	// Acceptable values of the parameter
	TextValue []*string `json:"TextValue,omitnil" name:"TextValue"`

	// Parameter modification status. Valid values: - `1` (modifying) - `2` (modified)
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

type Instances struct {
	// User APPID, which is the unique application ID that matches an account. Some Tencent Cloud products use this APPID.
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Region ID. <ul><li>`1`: Guangzhou. </li><li>`4`: Shanghai. </li><li>`5`: Hong Kong (China). </li> <li>`6`: Toronto. </li> <li>`7`: Shanghai Finance. </li> <li>`8`: Beijing. </li> <li>`9`: Singapore. </li> <li>`11`: Shenzhen Finance. </li> <li>`15`: West US (Silicon Valley). </li> </ul>
	RegionId *uint64 `json:"RegionId,omitnil" name:"RegionId"`

	// Region ID
	ZoneId *uint64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// Number of replicas
	RedisReplicasNum *uint64 `json:"RedisReplicasNum,omitnil" name:"RedisReplicasNum"`

	// Number of shards
	RedisShardNum *int64 `json:"RedisShardNum,omitnil" name:"RedisShardNum"`

	// Shard memory size.
	RedisShardSize *int64 `json:"RedisShardSize,omitnil" name:"RedisShardSize"`

	// Instance disk size
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskSize *int64 `json:"DiskSize,omitnil" name:"DiskSize"`

	// Engine: Redis Community Edition, Tencent Cloud CKV.
	Engine *string `json:"Engine,omitnil" name:"Engine"`

	// Read-write permission of the instance. <ul><li>`rw`: Read/Write. </li><li>`r`: Read-only. </li></ul>
	Role *string `json:"Role,omitnil" name:"Role"`

	// Instance VIP
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// Internal parameter, which can be ignored.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Vip6 *string `json:"Vip6,omitnil" name:"Vip6"`

	// VPC ID, such as `75101`.
	VpcID *int64 `json:"VpcID,omitnil" name:"VpcID"`

	// Instance port
	VPort *int64 `json:"VPort,omitnil" name:"VPort"`

	// Instance status. <ul><li>`0`: Uninitialized. </li><li>`1`: In the process. </li><li>`2`: Running. </li><li>`-2`: Isolated. </li><li>`-3`: To be deleted. </li></ul>
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// Repository ID
	GrocerySysId *int64 `json:"GrocerySysId,omitnil" name:"GrocerySysId"`

	// Instance type
	// - `2`: Redis 2.8 Memory Edition (Standard Architecture).
	// - `3`: CKV 3.2 Memory Edition (Standard Architecture).
	// - `4`: CKV 3.2 Memory Edition (Cluster Architecture)
	// - `5`: Redis 2.8 Memory Edition (Standalone)
	// - `6`: Redis 4.0 Memory Edition (Standard Architecture).
	// - `7`: Redis 4.0 Memory Edition (Cluster Architecture)
	// - `8`: Redis 5.0 Memory Edition (Standard Architecture).
	// - `9`: Redis 5.0 Memory Edition (Cluster Architecture)
	// - `15`: Redis 6.2 Memory Edition (Standard Architecture).
	// - `16`: Redis 6.2 Memory Edition (Cluster Architecture)
	ProductType *int64 `json:"ProductType,omitnil" name:"ProductType"`

	// The time when the instance was added to the replication group.
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// The time when instances in the replication group were updated.
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

// Predefined struct for user
type KillMasterGroupRequestParams struct {
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// A parameter used to configure the access password for a specified instance. If password-free authentication is enabled, this parameter will not be required. Required password strength. - It must contains 8-30 characters. We recommend that you use a password of more than 12 characters. - It must contain at least two of the following types: lowercase letters, uppercase letters, digits, and symbols (()`~!@#$%^&*-+=_|{}[]:;<>,.?/), and it cannot start with a slash (/).
	Password *string `json:"Password,omitnil" name:"Password"`

	// Shard ID of a sharded cluster
	ShardIds []*int64 `json:"ShardIds,omitnil" name:"ShardIds"`
}

type KillMasterGroupRequest struct {
	*tchttp.BaseRequest
	
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// A parameter used to configure the access password for a specified instance. If password-free authentication is enabled, this parameter will not be required. Required password strength. - It must contains 8-30 characters. We recommend that you use a password of more than 12 characters. - It must contain at least two of the following types: lowercase letters, uppercase letters, digits, and symbols (()`~!@#$%^&*-+=_|{}[]:;<>,.?/), and it cannot start with a slash (/).
	Password *string `json:"Password,omitnil" name:"Password"`

	// Shard ID of a sharded cluster
	ShardIds []*int64 `json:"ShardIds,omitnil" name:"ShardIds"`
}

func (r *KillMasterGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillMasterGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Password")
	delete(f, "ShardIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "KillMasterGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type KillMasterGroupResponseParams struct {
	// Async task ID
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type KillMasterGroupResponse struct {
	*tchttp.BaseResponse
	Response *KillMasterGroupResponseParams `json:"Response"`
}

func (r *KillMasterGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillMasterGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ManualBackupInstanceRequestParams struct {
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Remarks for manual backup task
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Retention period of backup data in days.  Default value: 7 days.  Value range: [0,1825].  If the value exceeds 7 days, [submit a ticket](https://console.cloud.tencent.com/workorder/category) for application. - If this parameter is not configured, it will set to be the same as the period of automatic backup retention. - If automatic backup is not set, the default value will be 7 days.
	StorageDays *int64 `json:"StorageDays,omitnil" name:"StorageDays"`
}

type ManualBackupInstanceRequest struct {
	*tchttp.BaseRequest
	
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Remarks for manual backup task
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Retention period of backup data in days.  Default value: 7 days.  Value range: [0,1825].  If the value exceeds 7 days, [submit a ticket](https://console.cloud.tencent.com/workorder/category) for application. - If this parameter is not configured, it will set to be the same as the period of automatic backup retention. - If automatic backup is not set, the default value will be 7 days.
	StorageDays *int64 `json:"StorageDays,omitnil" name:"StorageDays"`
}

func (r *ManualBackupInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManualBackupInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Remark")
	delete(f, "StorageDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ManualBackupInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ManualBackupInstanceResponseParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ManualBackupInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ManualBackupInstanceResponseParams `json:"Response"`
}

func (r *ManualBackupInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManualBackupInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModfiyInstancePasswordRequestParams struct {
	// Instance ID, such as "crs-xjhsdj****". Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Old password of an instance
	OldPassword *string `json:"OldPassword,omitnil" name:"OldPassword"`

	// New instance password, which has the following requirements:
	// - It must contain 8-30 characters, preferably 12 or more.
	// - It cannot start with a slash (/)
	// - It must contain two of the following three types: lowercase letters, uppercase letters, and symbols (()~!@#$%^&*-+=_|{}[]:;<>,.?/)
	Password *string `json:"Password,omitnil" name:"Password"`
}

type ModfiyInstancePasswordRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID, such as "crs-xjhsdj****". Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Old password of an instance
	OldPassword *string `json:"OldPassword,omitnil" name:"OldPassword"`

	// New instance password, which has the following requirements:
	// - It must contain 8-30 characters, preferably 12 or more.
	// - It cannot start with a slash (/)
	// - It must contain two of the following three types: lowercase letters, uppercase letters, and symbols (()~!@#$%^&*-+=_|{}[]:;<>,.?/)
	Password *string `json:"Password,omitnil" name:"Password"`
}

func (r *ModfiyInstancePasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModfiyInstancePasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "OldPassword")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModfiyInstancePasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModfiyInstancePasswordResponseParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModfiyInstancePasswordResponse struct {
	*tchttp.BaseResponse
	Response *ModfiyInstancePasswordResponseParams `json:"Response"`
}

func (r *ModfiyInstancePasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModfiyInstancePasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoBackupConfigRequestParams struct {
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Automatic backup cycle. Valid values: `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`. This parameter currently cannot be modified.
	WeekDays []*string `json:"WeekDays,omitnil" name:"WeekDays"`

	// Automatic backup time in the format of 00:00-01:00, 01:00-02:00... 23:00-00:00.
	TimePeriod *string `json:"TimePeriod,omitnil" name:"TimePeriod"`

	// Automatic backup type.  Valid value:  `1` (scheduled backup).
	AutoBackupType *int64 `json:"AutoBackupType,omitnil" name:"AutoBackupType"`
}

type ModifyAutoBackupConfigRequest struct {
	*tchttp.BaseRequest
	
	// ID of a specified instance,  such as  "crs-xjhsdj****" Log in to the [Redis console](https://console.cloud.tencent.com/redis) and copy the instance ID in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Automatic backup cycle. Valid values: `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`. This parameter currently cannot be modified.
	WeekDays []*string `json:"WeekDays,omitnil" name:"WeekDays"`

	// Automatic backup time in the format of 00:00-01:00, 01:00-02:00... 23:00-00:00.
	TimePeriod *string `json:"TimePeriod,omitnil" name:"TimePeriod"`

	// Automatic backup type.  Valid value:  `1` (scheduled backup).
	AutoBackupType *int64 `json:"AutoBackupType,omitnil" name:"AutoBackupType"`
}

func (r *ModifyAutoBackupConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoBackupConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "WeekDays")
	delete(f, "TimePeriod")
	delete(f, "AutoBackupType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAutoBackupConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoBackupConfigResponseParams struct {
	// Automatic backup type.  Valid value:  `1` (scheduled backup).
	AutoBackupType *int64 `json:"AutoBackupType,omitnil" name:"AutoBackupType"`

	// Automatic backup cycle. Valid values: `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`.
	WeekDays []*string `json:"WeekDays,omitnil" name:"WeekDays"`

	// Time period for automatic scheduled backup  in the format of  “00:00-01:00, 01:00-02:00...... 23:00-00:00”.
	TimePeriod *string `json:"TimePeriod,omitnil" name:"TimePeriod"`

	// Retention time of full backup files in days
	BackupStorageDays *int64 `json:"BackupStorageDays,omitnil" name:"BackupStorageDays"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAutoBackupConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAutoBackupConfigResponseParams `json:"Response"`
}

func (r *ModifyAutoBackupConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoBackupConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupDownloadRestrictionRequestParams struct {
	// Type of the network restrictions for downloading backup files. Valid values:
	// 
	// - `NoLimit`: Backup files can be downloaded over both public and private networks.
	// - `LimitOnlyIntranet`: Backup files can be downloaded only at private network addresses auto-assigned by Tencent Cloud.
	// - `Customize`: Backup files can be downloaded only in the customized VPC.
	LimitType *string `json:"LimitType,omitnil" name:"LimitType"`

	// Only `In` can be passed in for this parameter, indicating that backup files can be downloaded in the custom `LimitVpc`.
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitnil" name:"VpcComparisonSymbol"`

	// Whether backups can be downloaded at the custom `LimitIp` address.
	// 
	// - `In`: Download is allowed for the custom IP.
	// - `NotIn`: Download is not allowed for the custom IP.
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitnil" name:"IpComparisonSymbol"`

	// VPC ID of the custom backup file download address, which is required if `LimitType` is `Customize`.
	LimitVpc []*BackupLimitVpcItem `json:"LimitVpc,omitnil" name:"LimitVpc"`

	// VPC IP of the custom backup file download address, which is required if `LimitType` is `Customize`.
	LimitIp []*string `json:"LimitIp,omitnil" name:"LimitIp"`
}

type ModifyBackupDownloadRestrictionRequest struct {
	*tchttp.BaseRequest
	
	// Type of the network restrictions for downloading backup files. Valid values:
	// 
	// - `NoLimit`: Backup files can be downloaded over both public and private networks.
	// - `LimitOnlyIntranet`: Backup files can be downloaded only at private network addresses auto-assigned by Tencent Cloud.
	// - `Customize`: Backup files can be downloaded only in the customized VPC.
	LimitType *string `json:"LimitType,omitnil" name:"LimitType"`

	// Only `In` can be passed in for this parameter, indicating that backup files can be downloaded in the custom `LimitVpc`.
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitnil" name:"VpcComparisonSymbol"`

	// Whether backups can be downloaded at the custom `LimitIp` address.
	// 
	// - `In`: Download is allowed for the custom IP.
	// - `NotIn`: Download is not allowed for the custom IP.
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitnil" name:"IpComparisonSymbol"`

	// VPC ID of the custom backup file download address, which is required if `LimitType` is `Customize`.
	LimitVpc []*BackupLimitVpcItem `json:"LimitVpc,omitnil" name:"LimitVpc"`

	// VPC IP of the custom backup file download address, which is required if `LimitType` is `Customize`.
	LimitIp []*string `json:"LimitIp,omitnil" name:"LimitIp"`
}

func (r *ModifyBackupDownloadRestrictionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupDownloadRestrictionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LimitType")
	delete(f, "VpcComparisonSymbol")
	delete(f, "IpComparisonSymbol")
	delete(f, "LimitVpc")
	delete(f, "LimitIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupDownloadRestrictionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupDownloadRestrictionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyBackupDownloadRestrictionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBackupDownloadRestrictionResponseParams `json:"Response"`
}

func (r *ModifyBackupDownloadRestrictionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupDownloadRestrictionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupsRequestParams struct {
	// Database engine name, which is `redis` for this API.
	Product *string `json:"Product,omitnil" name:"Product"`

	// List of IDs of security groups to be modified, which is an array of one or more security group IDs.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`

	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Database engine name, which is `redis` for this API.
	Product *string `json:"Product,omitnil" name:"Product"`

	// List of IDs of security groups to be modified, which is an array of one or more security group IDs.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`

	// Instance ID in the format of cdb-c1nl9rpv or cdbro-c1nl9rpv. It is the same as the instance ID displayed in the TencentDB console.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *ModifyDBInstanceSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "SecurityGroupIds")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDBInstanceSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceSecurityGroupsResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceAccountRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Sub-account name. If the root account is to be modified, enter `root`.
	AccountName *string `json:"AccountName,omitnil" name:"AccountName"`

	// Sub-account password
	AccountPassword *string `json:"AccountPassword,omitnil" name:"AccountPassword"`

	// Sub-account description information
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Routing policy. Valid values: master (master node); replication (replica node)
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitnil" name:"ReadonlyPolicy"`

	// Sub-account read/write policy. Valid values: r (read-only); w (write-only); rw (read/write).
	Privilege *string `json:"Privilege,omitnil" name:"Privilege"`

	// true: make the root account password-free. This is applicable to root accounts only. Sub-accounts cannot be made password-free.
	NoAuth *bool `json:"NoAuth,omitnil" name:"NoAuth"`
}

type ModifyInstanceAccountRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Sub-account name. If the root account is to be modified, enter `root`.
	AccountName *string `json:"AccountName,omitnil" name:"AccountName"`

	// Sub-account password
	AccountPassword *string `json:"AccountPassword,omitnil" name:"AccountPassword"`

	// Sub-account description information
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Routing policy. Valid values: master (master node); replication (replica node)
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitnil" name:"ReadonlyPolicy"`

	// Sub-account read/write policy. Valid values: r (read-only); w (write-only); rw (read/write).
	Privilege *string `json:"Privilege,omitnil" name:"Privilege"`

	// true: make the root account password-free. This is applicable to root accounts only. Sub-accounts cannot be made password-free.
	NoAuth *bool `json:"NoAuth,omitnil" name:"NoAuth"`
}

func (r *ModifyInstanceAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AccountName")
	delete(f, "AccountPassword")
	delete(f, "Remark")
	delete(f, "ReadonlyPolicy")
	delete(f, "Privilege")
	delete(f, "NoAuth")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceAccountResponseParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyInstanceAccountResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceAccountResponseParams `json:"Response"`
}

func (r *ModifyInstanceAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceParamsRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// List of instance parameters modified
	InstanceParams []*InstanceParam `json:"InstanceParams,omitnil" name:"InstanceParams"`
}

type ModifyInstanceParamsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// List of instance parameters modified
	InstanceParams []*InstanceParam `json:"InstanceParams,omitnil" name:"InstanceParams"`
}

func (r *ModifyInstanceParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceParamsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceParamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceParamsResponseParams struct {
	// Whether the parameter is modified successfully. <br><li>`True`: Yes<br><li>`False`: No<br>
	Changed *bool `json:"Changed,omitnil" name:"Changed"`

	// ID of the task
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyInstanceParamsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceParamsResponseParams `json:"Response"`
}

func (r *ModifyInstanceParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceReadOnlyRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance input mode. Valid values: `0` (read/write), `1` (read-only)
	InputMode *string `json:"InputMode,omitnil" name:"InputMode"`
}

type ModifyInstanceReadOnlyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance input mode. Valid values: `0` (read/write), `1` (read-only)
	InputMode *string `json:"InputMode,omitnil" name:"InputMode"`
}

func (r *ModifyInstanceReadOnlyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceReadOnlyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InputMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceReadOnlyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceReadOnlyResponseParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyInstanceReadOnlyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceReadOnlyResponseParams `json:"Response"`
}

func (r *ModifyInstanceReadOnlyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceReadOnlyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceRequestParams struct {
	// Instance modification type. rename: rename an instance; modifyProject: modify the project of an instance; modifyAutoRenew: modify the auto-renewal flag of an instance.
	Operation *string `json:"Operation,omitnil" name:"Operation"`

	// Instance ID
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// New name of the instance
	InstanceNames []*string `json:"InstanceNames,omitnil" name:"InstanceNames"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Auto-renewal flag. 0: default status (manual renewal); 1: auto-renewal enabled; 2: auto-renewal disabled
	AutoRenews []*int64 `json:"AutoRenews,omitnil" name:"AutoRenews"`

	// Disused
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Disused
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Disused
	AutoRenew *int64 `json:"AutoRenew,omitnil" name:"AutoRenew"`
}

type ModifyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance modification type. rename: rename an instance; modifyProject: modify the project of an instance; modifyAutoRenew: modify the auto-renewal flag of an instance.
	Operation *string `json:"Operation,omitnil" name:"Operation"`

	// Instance ID
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// New name of the instance
	InstanceNames []*string `json:"InstanceNames,omitnil" name:"InstanceNames"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Auto-renewal flag. 0: default status (manual renewal); 1: auto-renewal enabled; 2: auto-renewal disabled
	AutoRenews []*int64 `json:"AutoRenews,omitnil" name:"AutoRenews"`

	// Disused
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Disused
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Disused
	AutoRenew *int64 `json:"AutoRenew,omitnil" name:"AutoRenew"`
}

func (r *ModifyInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operation")
	delete(f, "InstanceIds")
	delete(f, "InstanceNames")
	delete(f, "ProjectId")
	delete(f, "AutoRenews")
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	delete(f, "AutoRenew")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceResponseParams `json:"Response"`
}

func (r *ModifyInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMaintenanceWindowRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Maintenance start time, such as 17:00
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// Maintenance end time, such as 19:00
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

type ModifyMaintenanceWindowRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Maintenance start time, such as 17:00
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// Maintenance end time, such as 19:00
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

func (r *ModifyMaintenanceWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMaintenanceWindowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMaintenanceWindowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMaintenanceWindowResponseParams struct {
	// Modification status. Valid values: success, failed.
	Status *string `json:"Status,omitnil" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyMaintenanceWindowResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMaintenanceWindowResponseParams `json:"Response"`
}

func (r *ModifyMaintenanceWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMaintenanceWindowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetworkConfigRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Network change type. Valid values:
	// - `changeVip`: VPC change, including the private IPv4 address and port.
	// - `changeVpc`: Subnet change.
	// - `changeBaseToVpc`: Change from classic network to VPC.
	// - `changeVPort`: Port change.
	Operation *string `json:"Operation,omitnil" name:"Operation"`

	// Private IPv4 address of the instance, which is required if `Operation` is `changeVip`.
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// VPC ID after the change, which is required if `Operation` is `changeVpc` or `changeBaseToVpc`.
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// Subnet ID after the change, which is required if `Operation` is `changeVpc` or `changeBaseToVpc`.
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// Retention period of the original private IPv4 address
	// - Unit: Days.
	// - Valid values: `0`, `1`, `2`, `3`, `7`, `15`.
	// 
	// **Note**: You can set the retention period of the original address only in the latest SDK. In earlier SDKs, the original address is released immediately. To view the SDK version, go to [SDK Center](https://intl.cloud.tencent.com/document/sdk?from_cn_redirect=1).
	Recycle *int64 `json:"Recycle,omitnil" name:"Recycle"`

	// Network port after the change, which is required if `Operation` is `changeVPort` or `changeVip`. Value range: [1024,65535].
	VPort *int64 `json:"VPort,omitnil" name:"VPort"`
}

type ModifyNetworkConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Network change type. Valid values:
	// - `changeVip`: VPC change, including the private IPv4 address and port.
	// - `changeVpc`: Subnet change.
	// - `changeBaseToVpc`: Change from classic network to VPC.
	// - `changeVPort`: Port change.
	Operation *string `json:"Operation,omitnil" name:"Operation"`

	// Private IPv4 address of the instance, which is required if `Operation` is `changeVip`.
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// VPC ID after the change, which is required if `Operation` is `changeVpc` or `changeBaseToVpc`.
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// Subnet ID after the change, which is required if `Operation` is `changeVpc` or `changeBaseToVpc`.
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// Retention period of the original private IPv4 address
	// - Unit: Days.
	// - Valid values: `0`, `1`, `2`, `3`, `7`, `15`.
	// 
	// **Note**: You can set the retention period of the original address only in the latest SDK. In earlier SDKs, the original address is released immediately. To view the SDK version, go to [SDK Center](https://intl.cloud.tencent.com/document/sdk?from_cn_redirect=1).
	Recycle *int64 `json:"Recycle,omitnil" name:"Recycle"`

	// Network port after the change, which is required if `Operation` is `changeVPort` or `changeVip`. Value range: [1024,65535].
	VPort *int64 `json:"VPort,omitnil" name:"VPort"`
}

func (r *ModifyNetworkConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetworkConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Operation")
	delete(f, "Vip")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Recycle")
	delete(f, "VPort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNetworkConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetworkConfigResponseParams struct {
	// Execution status. Ignore this parameter.
	Status *bool `json:"Status,omitnil" name:"Status"`

	// New subnet ID of the instance
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// New VPC ID of the instance
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// New private IPv4 address of the instance
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// Task ID, which can be used to query the task execution status through the `DescribeTaskInfo` API.
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyNetworkConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNetworkConfigResponseParams `json:"Response"`
}

func (r *ModifyNetworkConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetworkConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyParamTemplateRequestParams struct {
	// ID of the source parameter template.
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// New name after the parameter template is modified.
	Name *string `json:"Name,omitnil" name:"Name"`

	// New description after the parameter template is modified.
	Description *string `json:"Description,omitnil" name:"Description"`

	// New parameter list after the parameter template is modified.
	ParamList []*InstanceParam `json:"ParamList,omitnil" name:"ParamList"`
}

type ModifyParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// ID of the source parameter template.
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// New name after the parameter template is modified.
	Name *string `json:"Name,omitnil" name:"Name"`

	// New description after the parameter template is modified.
	Description *string `json:"Description,omitnil" name:"Description"`

	// New parameter list after the parameter template is modified.
	ParamList []*InstanceParam `json:"ParamList,omitnil" name:"ParamList"`
}

func (r *ModifyParamTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyParamTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "ParamList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyParamTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyParamTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyParamTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyParamTemplateResponseParams `json:"Response"`
}

func (r *ModifyParamTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyParamTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenSSLRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type OpenSSLRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *OpenSSLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenSSLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenSSLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenSSLResponseParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type OpenSSLResponse struct {
	*tchttp.BaseResponse
	Response *OpenSSLResponseParams `json:"Response"`
}

func (r *OpenSSLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenSSLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Outbound struct {
	// Policy. Valid values: ACCEPT, DROP.
	Action *string `json:"Action,omitnil" name:"Action"`

	// All the addresses that the address group ID represents.
	AddressModule *string `json:"AddressModule,omitnil" name:"AddressModule"`

	// Source IP or IP address range, such as 192.168.0.0/16.
	CidrIp *string `json:"CidrIp,omitnil" name:"CidrIp"`

	// Description.
	Desc *string `json:"Desc,omitnil" name:"Desc"`

	// Network protocol, such as UDP and TCP.
	IpProtocol *string `json:"IpProtocol,omitnil" name:"IpProtocol"`

	// Port.
	PortRange *string `json:"PortRange,omitnil" name:"PortRange"`

	// All the protocols and ports that the service group ID represents.
	ServiceModule *string `json:"ServiceModule,omitnil" name:"ServiceModule"`

	// All the addresses that the security group ID represents.
	Id *string `json:"Id,omitnil" name:"Id"`
}

type ParamTemplateInfo struct {
	// Parameter template ID
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// Parameter template name
	Name *string `json:"Name,omitnil" name:"Name"`

	// Parameter template description
	Description *string `json:"Description,omitnil" name:"Description"`

	// Instance type
	// - `2`: Redis 2.8 Memory Edition (Standard Architecture).
	// - `3`: CKV 3.2 Memory Edition (Standard Architecture).
	// - `4`: CKV 3.2 Memory Edition (Cluster Architecture).
	// - `5`: Redis 2.8 Memory Edition (Standalone).
	// - `6`: Redis 4.0 Memory Edition (Standard Architecture).
	// - `7`: Redis 4.0 Memory Edition (Cluster Architecture).
	// - `8`: Redis 5.0 Memory Edition (Standard Architecture).
	// - `9`: Redis 5.0 Memory Edition (Cluster Architecture).
	// - `15`: Redis 6.2 Memory Edition (Standard Architecture).
	// - `16`: Redis 6.2 Memory Edition (Cluster Architecture).
	ProductType *uint64 `json:"ProductType,omitnil" name:"ProductType"`
}

type ParameterDetail struct {
	// Parameter name
	Name *string `json:"Name,omitnil" name:"Name"`

	// Parameter Type
	ParamType *string `json:"ParamType,omitnil" name:"ParamType"`

	// Default value of the parameter
	Default *string `json:"Default,omitnil" name:"Default"`

	// Parameter description
	Description *string `json:"Description,omitnil" name:"Description"`

	// Current value of the parameter
	CurrentValue *string `json:"CurrentValue,omitnil" name:"CurrentValue"`

	// Whether to restart the database for the modified parameters to take effect
	// - `0`: No restart.
	// - `1`: Restart required.
	NeedReboot *int64 `json:"NeedReboot,omitnil" name:"NeedReboot"`

	// Maximum value of the parameter
	// Note: This field may return null, indicating that no valid values can be obtained.
	Max *string `json:"Max,omitnil" name:"Max"`

	// Minimum value of the parameter
	// Note: This field may return null, indicating that no valid values can be obtained.
	Min *string `json:"Min,omitnil" name:"Min"`

	// Enumerated values of the parameter. It is null if the parameter is non-enumerated
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnumValue []*string `json:"EnumValue,omitnil" name:"EnumValue"`
}

type ProductConf struct {
	// Product type
	// - `2`: Redis 2.8 Memory Edition (Standard Architecture).
	// - `3`: CKV 3.2 Memory Edition (Standard Architecture).
	// - `4`: CKV 3.2 Memory Edition (Cluster Architecture).
	// - `5`: Redis 2.8 Memory Edition (Standalone).
	// - `6`: Redis 4.0 Memory Edition (Standard Architecture).
	// - `7`: Redis 4.0 Memory Edition (Cluster Architecture).
	// - `8`: Redis 5.0 Memory Edition (Standard Architecture).
	// - `9`: Redis 5.0 Memory Edition (Cluster Architecture).
	// - `15`: Redis 6.2 Memory Edition (Standard Architecture).
	// - `16`: Redis 6.2 Memory Edition (Cluster Architecture).
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// Product names, including Redis Master-Replica Edition, Redis Standalone Edition, Redis 4.0 Cluster Edition, CKV Master-Replica Edition, and CKV Standalone Edition.
	TypeName *string `json:"TypeName,omitnil" name:"TypeName"`

	// Minimum purchasable quantity
	MinBuyNum *int64 `json:"MinBuyNum,omitnil" name:"MinBuyNum"`

	// Maximum purchasable quantity
	MaxBuyNum *int64 `json:"MaxBuyNum,omitnil" name:"MaxBuyNum"`

	// Whether a product is sold out
	// - `true`: Sold out.
	// - `false`: Not sold out.
	Saleout *bool `json:"Saleout,omitnil" name:"Saleout"`

	// Product engines, including Tencent Cloud CKV and Redis Community Edition.
	Engine *string `json:"Engine,omitnil" name:"Engine"`

	// Compatible versions, including Redis 2.8, 3.2, 4.0, 5.0, and 6.2.
	Version *string `json:"Version,omitnil" name:"Version"`

	// Total capacity in GB
	TotalSize []*string `json:"TotalSize,omitnil" name:"TotalSize"`

	// Shard size in GB
	ShardSize []*string `json:"ShardSize,omitnil" name:"ShardSize"`

	// Quantity of replicas
	ReplicaNum []*string `json:"ReplicaNum,omitnil" name:"ReplicaNum"`

	// Quantity of shards
	ShardNum []*string `json:"ShardNum,omitnil" name:"ShardNum"`

	// Supported billing modes
	// - `1`: Monthly subscription.
	// - `0`: Pay-as-you-go.
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// Whether to support read-only replicas
	// - `true`: Supported.
	// -`false`: Not supported.
	EnableRepicaReadOnly *bool `json:"EnableRepicaReadOnly,omitnil" name:"EnableRepicaReadOnly"`
}

type ProxyNodes struct {
	// Node ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	NodeId *string `json:"NodeId,omitnil" name:"NodeId"`

	// AZ ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneId *int64 `json:"ZoneId,omitnil" name:"ZoneId"`
}

type RedisBackupSet struct {
	// Backup start time
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// Backup task ID
	BackupId *string `json:"BackupId,omitnil" name:"BackupId"`

	// Backup type. Valid values:  `1` (Automatic backup in the early morning initiated by the system.) `0`: Manual backup initiated by the user.
	BackupType *string `json:"BackupType,omitnil" name:"BackupType"`

	// Backup status. Valid values:  - `1`: The backup is locked by another process. - `2`: The backup is normal and not locked by any process. - `-1`: The backup is expired. - `3`: The backup is being exported. - `4`: Exported the backup successfully.
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// Backup remarks
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Whether the backup is locked. Valid values:  - `0` (no) - `1` (yes)
	Locked *int64 `json:"Locked,omitnil" name:"Locked"`

	// Internal field, which can be ignored.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupSize *int64 `json:"BackupSize,omitnil" name:"BackupSize"`

	// Internal field, which can be ignored.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FullBackup *int64 `json:"FullBackup,omitnil" name:"FullBackup"`

	// Internal field, which can be ignored.
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceType *int64 `json:"InstanceType,omitnil" name:"InstanceType"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// The region where the local backup resides.
	Region *string `json:"Region,omitnil" name:"Region"`

	// Backup end time
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Backup file type
	FileType *string `json:"FileType,omitnil" name:"FileType"`

	// Backup file expiration time
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`
}

type RedisCommonInstanceList struct {
	// Instance name
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// User APPID, which is the unique application ID that matches an account. Some Tencent Cloud products use this APPID.
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`

	// Project ID of the instance
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Instance region
	Region *string `json:"Region,omitnil" name:"Region"`

	// Instance AZ
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// Instance VPC ID
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// VPC subnet ID
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// Instance status information
	// - `1`: Task running.
	// - `2`: Instance running.
	// - `-2`: Instance isolated.
	// - `-3`: Instance being eliminated.
	// - `-4`: Instance eliminated.
	Status *string `json:"Status,omitnil" name:"Status"`

	// Private network IP address of an instance
	Vips []*string `json:"Vips,omitnil" name:"Vips"`

	// Instance network port
	Vport *int64 `json:"Vport,omitnil" name:"Vport"`

	// Instance creation time
	Createtime *string `json:"Createtime,omitnil" name:"Createtime"`

	// Billing type
	// - `0`: Pay-as-you-go.
	// - `1`: Monthly subscription.
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`

	// Network Type
	// - `0`: Classic network.
	// - `1`: VPC.
	NetType *int64 `json:"NetType,omitnil" name:"NetType"`
}

type RedisNode struct {
	// Number of keys on Redis nodes
	Keys *int64 `json:"Keys,omitnil" name:"Keys"`

	// Slot distribution range for Redis node.  Value range:  0-5460.
	Slot *string `json:"Slot,omitnil" name:"Slot"`

	// Node sequence ID
	NodeId *string `json:"NodeId,omitnil" name:"NodeId"`

	// Node status
	Status *string `json:"Status,omitnil" name:"Status"`

	// Node role
	Role *string `json:"Role,omitnil" name:"Role"`
}

type RedisNodeInfo struct {
	// Node type. <ul><li>`0`: Master node.</li><li>`1`: Replica node.</li></ul>
	NodeType *int64 `json:"NodeType,omitnil" name:"NodeType"`

	// ID of the master or replica node <ul><li>This parameter is optional when the [CreateInstances](https://intl.cloud.tencent.com/document/product/239/20026?from_cn_redirect=1) API is used to create a TencentDB for Redis instance, but it is required when the [UpgradeInstance](https://intl.cloud.tencent.com/document/product/239/20013?from_cn_redirect=1) API is used to adjust the configuration of an instance by deleting a replica.  </li><li>You can use the [DescribeInstances](https://intl.cloud.tencent.com/document/product/239/20018?from_cn_redirect=1) API to get the node ID of integer type. </li></ul> </li></ul>
	NodeId *int64 `json:"NodeId,omitnil" name:"NodeId"`

	// ID of the AZ of the master or replica node
	ZoneId *uint64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// Name of the AZ of the master or replica node
	ZoneName *string `json:"ZoneName,omitnil" name:"ZoneName"`
}

type RedisNodes struct {
	// Node ID
	NodeId *string `json:"NodeId,omitnil" name:"NodeId"`

	// Node role
	NodeRole *string `json:"NodeRole,omitnil" name:"NodeRole"`

	// Shard ID
	ClusterId *int64 `json:"ClusterId,omitnil" name:"ClusterId"`

	// AZ ID
	ZoneId *int64 `json:"ZoneId,omitnil" name:"ZoneId"`
}

type RegionConf struct {
	// Region ID
	RegionId *string `json:"RegionId,omitnil" name:"RegionId"`

	// Region name
	RegionName *string `json:"RegionName,omitnil" name:"RegionName"`

	// Region abbreviation
	RegionShortName *string `json:"RegionShortName,omitnil" name:"RegionShortName"`

	// Name of the area where a region is located
	Area *string `json:"Area,omitnil" name:"Area"`

	// AZ information
	ZoneSet []*ZoneCapacityConf `json:"ZoneSet,omitnil" name:"ZoneSet"`
}

// Predefined struct for user
type ReleaseWanAddressRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type ReleaseWanAddressRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *ReleaseWanAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseWanAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReleaseWanAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseWanAddressResponseParams struct {
	// Async task ID
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// Status of disabling public network access
	WanStatus *string `json:"WanStatus,omitnil" name:"WanStatus"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ReleaseWanAddressResponse struct {
	*tchttp.BaseResponse
	Response *ReleaseWanAddressResponseParams `json:"Response"`
}

func (r *ReleaseWanAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseWanAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveReplicationInstanceRequestParams struct {
	// Replication group ID
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Data sync type. Valid values: `true` (strong sync is required), `false` (strong sync is not required, only the master instance can be deleted).
	SyncType *bool `json:"SyncType,omitnil" name:"SyncType"`
}

type RemoveReplicationInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Replication group ID
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Data sync type. Valid values: `true` (strong sync is required), `false` (strong sync is not required, only the master instance can be deleted).
	SyncType *bool `json:"SyncType,omitnil" name:"SyncType"`
}

func (r *RemoveReplicationInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveReplicationInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "InstanceId")
	delete(f, "SyncType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveReplicationInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveReplicationInstanceResponseParams struct {
	// Async task ID
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RemoveReplicationInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RemoveReplicationInstanceResponseParams `json:"Response"`
}

func (r *RemoveReplicationInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveReplicationInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewInstanceRequestParams struct {
	// Validity period in months
	Period *uint64 `json:"Period,omitnil" name:"Period"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The parameter used to determine whether to modify the billing mode. <ul><li>If you want to change the billing mode from pay-as-you-go to monthly subscription, specify this parameter as <b>prepaid</b>. </li><li>If the current instance is monthly subscribed, this parameter is not required. </li></ul>
	ModifyPayMode *string `json:"ModifyPayMode,omitnil" name:"ModifyPayMode"`
}

type RenewInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Validity period in months
	Period *uint64 `json:"Period,omitnil" name:"Period"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The parameter used to determine whether to modify the billing mode. <ul><li>If you want to change the billing mode from pay-as-you-go to monthly subscription, specify this parameter as <b>prepaid</b>. </li><li>If the current instance is monthly subscribed, this parameter is not required. </li></ul>
	ModifyPayMode *string `json:"ModifyPayMode,omitnil" name:"ModifyPayMode"`
}

func (r *RenewInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Period")
	delete(f, "InstanceId")
	delete(f, "ModifyPayMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewInstanceResponseParams struct {
	// Transaction ID
	DealId *string `json:"DealId,omitnil" name:"DealId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RenewInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RenewInstanceResponseParams `json:"Response"`
}

func (r *RenewInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplicaGroup struct {
	// Node group ID
	GroupId *int64 `json:"GroupId,omitnil" name:"GroupId"`

	// Node group name, which is empty for the master node
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// Node AZ ID, such as ap-guangzhou-1
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// Node group type. Valid values: master (master node group); replica (replica node group)
	Role *string `json:"Role,omitnil" name:"Role"`

	// List of nodes in the node group
	RedisNodes []*RedisNode `json:"RedisNodes,omitnil" name:"RedisNodes"`
}

// Predefined struct for user
type ResetPasswordRequestParams struct {
	// Redis instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Password reset (this parameter can be left blank when switching to password-free instance mode and is required in other cases)
	Password *string `json:"Password,omitnil" name:"Password"`

	// Whether to switch to password-free instance mode. false: switch to password-enabled instance mode; true: switch to password-free instance mode. Default value: false.
	NoAuth *bool `json:"NoAuth,omitnil" name:"NoAuth"`
}

type ResetPasswordRequest struct {
	*tchttp.BaseRequest
	
	// Redis instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Password reset (this parameter can be left blank when switching to password-free instance mode and is required in other cases)
	Password *string `json:"Password,omitnil" name:"Password"`

	// Whether to switch to password-free instance mode. false: switch to password-enabled instance mode; true: switch to password-free instance mode. Default value: false.
	NoAuth *bool `json:"NoAuth,omitnil" name:"NoAuth"`
}

func (r *ResetPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Password")
	delete(f, "NoAuth")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetPasswordResponseParams struct {
	// Task ID (this parameter is the task ID when changing a password. If you are switching between password-free and password-enabled instance mode, you can leave this parameter alone)
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ResetPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetPasswordResponseParams `json:"Response"`
}

func (r *ResetPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceTag struct {
	// Tag key
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// The value corresponding to the tag key
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}

// Predefined struct for user
type RestoreInstanceRequestParams struct {
	// ID of the instance to be operated on, which can be obtained through the `InstanceId` field in the return value of the `DescribeInstances` API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Backup ID, which can be obtained through the `backupId` field in the return value of the `GetRedisBackupList` API.
	BackupId *string `json:"BackupId,omitnil" name:"BackupId"`

	// Instance password, which needs to be validated during instance restoration (this parameter is not required for password-free instances)
	Password *string `json:"Password,omitnil" name:"Password"`
}

type RestoreInstanceRequest struct {
	*tchttp.BaseRequest
	
	// ID of the instance to be operated on, which can be obtained through the `InstanceId` field in the return value of the `DescribeInstances` API.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Backup ID, which can be obtained through the `backupId` field in the return value of the `GetRedisBackupList` API.
	BackupId *string `json:"BackupId,omitnil" name:"BackupId"`

	// Instance password, which needs to be validated during instance restoration (this parameter is not required for password-free instances)
	Password *string `json:"Password,omitnil" name:"Password"`
}

func (r *RestoreInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupId")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestoreInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestoreInstanceResponseParams struct {
	// Task ID, which can be used to query the task execution status through the `DescribeTaskInfo` API.
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RestoreInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RestoreInstanceResponseParams `json:"Response"`
}

func (r *RestoreInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroup struct {
	// Creation time in the format of yyyy-mm-dd hh:mm:ss.
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Project ID.
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Security group ID.
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// Security group name.
	SecurityGroupName *string `json:"SecurityGroupName,omitnil" name:"SecurityGroupName"`

	// Security group remarks.
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitnil" name:"SecurityGroupRemark"`

	// Outbound rule.
	Outbound []*Outbound `json:"Outbound,omitnil" name:"Outbound"`

	// Inbound rule.
	Inbound []*Inbound `json:"Inbound,omitnil" name:"Inbound"`
}

type SecurityGroupDetail struct {
	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Security group creation time
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// Security group name
	SecurityGroupName *string `json:"SecurityGroupName,omitnil" name:"SecurityGroupName"`

	// Security group remarks
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitnil" name:"SecurityGroupRemark"`

	// Inbound rules of the security group, which control the access source to the database.
	InboundRule []*SecurityGroupsInboundAndOutbound `json:"InboundRule,omitnil" name:"InboundRule"`

	// Security group outbound rule
	OutboundRule []*SecurityGroupsInboundAndOutbound `json:"OutboundRule,omitnil" name:"OutboundRule"`
}

type SecurityGroupsInboundAndOutbound struct {
	// Identify whether the IP and port for accessing the database are allowed
	Action *string `json:"Action,omitnil" name:"Action"`

	// IP address for accessing the database
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Port number
	Port *string `json:"Port,omitnil" name:"Port"`

	// Protocol type
	Proto *string `json:"Proto,omitnil" name:"Proto"`
}

type SourceCommand struct {
	// Command
	Cmd *string `json:"Cmd,omitnil" name:"Cmd"`

	// Number of executions
	Count *int64 `json:"Count,omitnil" name:"Count"`
}

type SourceInfo struct {
	// Source IP
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// Number of connections
	Conn *int64 `json:"Conn,omitnil" name:"Conn"`

	// Command
	Cmd *int64 `json:"Cmd,omitnil" name:"Cmd"`
}

// Predefined struct for user
type StartupInstanceRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type StartupInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *StartupInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartupInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartupInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartupInstanceResponseParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StartupInstanceResponse struct {
	*tchttp.BaseResponse
	Response *StartupInstanceResponseParams `json:"Response"`
}

func (r *StartupInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartupInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchInstanceVipRequestParams struct {
	// Source instance ID
	SrcInstanceId *string `json:"SrcInstanceId,omitnil" name:"SrcInstanceId"`

	// Target instance ID
	DstInstanceId *string `json:"DstInstanceId,omitnil" name:"DstInstanceId"`

	// The time that lapses in seconds since DTS is disconnected between the source instance and the target instance. If the DTS disconnection time period is greater than TimeDelay, the VIP will not be switched. We recommend you set an acceptable value based on the actual business conditions.
	TimeDelay *int64 `json:"TimeDelay,omitnil" name:"TimeDelay"`

	// Whether to force the switch when DTS is disconnected. 1: yes; 0: no.
	ForceSwitch *int64 `json:"ForceSwitch,omitnil" name:"ForceSwitch"`

	// now: switch now; syncComplete: switch after sync is completed
	SwitchTime *string `json:"SwitchTime,omitnil" name:"SwitchTime"`
}

type SwitchInstanceVipRequest struct {
	*tchttp.BaseRequest
	
	// Source instance ID
	SrcInstanceId *string `json:"SrcInstanceId,omitnil" name:"SrcInstanceId"`

	// Target instance ID
	DstInstanceId *string `json:"DstInstanceId,omitnil" name:"DstInstanceId"`

	// The time that lapses in seconds since DTS is disconnected between the source instance and the target instance. If the DTS disconnection time period is greater than TimeDelay, the VIP will not be switched. We recommend you set an acceptable value based on the actual business conditions.
	TimeDelay *int64 `json:"TimeDelay,omitnil" name:"TimeDelay"`

	// Whether to force the switch when DTS is disconnected. 1: yes; 0: no.
	ForceSwitch *int64 `json:"ForceSwitch,omitnil" name:"ForceSwitch"`

	// now: switch now; syncComplete: switch after sync is completed
	SwitchTime *string `json:"SwitchTime,omitnil" name:"SwitchTime"`
}

func (r *SwitchInstanceVipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchInstanceVipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SrcInstanceId")
	delete(f, "DstInstanceId")
	delete(f, "TimeDelay")
	delete(f, "ForceSwitch")
	delete(f, "SwitchTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchInstanceVipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchInstanceVipResponseParams struct {
	// Task ID
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SwitchInstanceVipResponse struct {
	*tchttp.BaseResponse
	Response *SwitchInstanceVipResponseParams `json:"Response"`
}

func (r *SwitchInstanceVipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchInstanceVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchProxyRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance ProxyID
	ProxyID *string `json:"ProxyID,omitnil" name:"ProxyID"`
}

type SwitchProxyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance ProxyID
	ProxyID *string `json:"ProxyID,omitnil" name:"ProxyID"`
}

func (r *SwitchProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ProxyID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchProxyResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SwitchProxyResponse struct {
	*tchttp.BaseResponse
	Response *SwitchProxyResponseParams `json:"Response"`
}

func (r *SwitchProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TaskInfoDetail struct {
	// Task ID 
	// Note:  This field may return null, indicating that no valid values can be obtained.
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// Task start time 
	// Note:  This field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// Task type. Valid values:  - `FLOW_CREATE`: Create an instance. - `FLOW_MODIFYCONNECTIONCONFIG`: Adjust the number of bandwidth connections. - `FLOW_MODIFYINSTANCEPASSWORDFREE`: Modify the process of password-free access. - `FLOW_CLEARNETWORK`: Returning VPC - `FLOW_SETPWD`: Set the access password. - `FLOW_EXPORSHR`: Expand or reduce the capacity. - `FLOW_UpgradeArch`: Upgrade the instance architecture. - `FLOW_MODIFYINSTANCEPARAMS`: Modify the instance parameters. - `FLOW_MODIFYINSTACEREADONLY`: Modify read-only process. - `FLOW_CLOSE`: Disable the instance. - `FLOW_DELETE`: Delete the instance. - `FLOW_OPEN_WAN`: Enable the public network. - `FLOW_FLOW_CLEAN`: Clear the instance. - `FLOW_MODIFYINSTANCEACCOUNT`: Modify the instance account. - `FLOW_ENABLEINSTANCE_REPLICATE`: Enable the replica read-only feature. - `FLOW_DISABLEINSTANCE_REPLICATE`: Disable the replica read-only feature. - `FLOW_SWITCHINSTANCEVIP`: Swap the VIPs of instances. - FLOW_CHANGE_REPLICA_TO_MSTER: Promote the replica node to the mater node. Backup instance 
	// Note:  This field may return null, indicating that no valid values can be obtained.
	TaskType *string `json:"TaskType,omitnil" name:"TaskType"`

	// Instance name 
	// Note:  This field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Instance ID 
	// Note:  This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Project ID 
	// Note:  This field may return null, indicating that no valid values can be obtained.
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Task progress 
	// Note:  This field may return null, indicating that no valid values can be obtained.
	Progress *float64 `json:"Progress,omitnil" name:"Progress"`

	// Task end time 
	// Note:  This field may return null, indicating that no valid values can be obtained.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Task execution status. Valid values: - `0` (initilized) - `1` (executing) - `2` (completed) - `4` (failed) 
	// Note:  This field may return null, indicating that no valid values can be obtained.
	Result *int64 `json:"Result,omitnil" name:"Result"`
}

type TendisNodes struct {
	// Node ID
	NodeId *string `json:"NodeId,omitnil" name:"NodeId"`

	// Node role
	NodeRole *string `json:"NodeRole,omitnil" name:"NodeRole"`
}

type TendisSlowLogDetail struct {
	// Execution time
	ExecuteTime *string `json:"ExecuteTime,omitnil" name:"ExecuteTime"`

	// Duration of the slow query (ms)
	Duration *int64 `json:"Duration,omitnil" name:"Duration"`

	// Command
	Command *string `json:"Command,omitnil" name:"Command"`

	// Command line details
	CommandLine *string `json:"CommandLine,omitnil" name:"CommandLine"`

	// Node ID
	Node *string `json:"Node,omitnil" name:"Node"`
}

type TradeDealDetail struct {
	// Order ID, which is used when a TencentCloud API is called.
	DealId *string `json:"DealId,omitnil" name:"DealId"`

	// Long order ID, which is used when an order issue is submitted for assistance.
	DealName *string `json:"DealName,omitnil" name:"DealName"`

	// AZ ID
	ZoneId *int64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// Number of instances associated with the order
	GoodsNum *int64 `json:"GoodsNum,omitnil" name:"GoodsNum"`

	// Creator UIN
	Creater *string `json:"Creater,omitnil" name:"Creater"`

	// Order creation time
	CreatTime *string `json:"CreatTime,omitnil" name:"CreatTime"`

	// Order timeout period
	OverdueTime *string `json:"OverdueTime,omitnil" name:"OverdueTime"`

	// Order completion time
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Order status. 1: unpaid; 2: paid but not delivered; 3: In delivery; 4: successfully delivered; 5: delivery failed; 6: refunded; 7: order closed; 8: order expired; 9: order invalidated; 10: product invalidated; 11: requested payment rejected; 12: paying
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// Order status description
	Description *string `json:"Description,omitnil" name:"Description"`

	// Actual total price of the order in 0.01 CNY
	Price *int64 `json:"Price,omitnil" name:"Price"`

	// Instance ID
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

// Predefined struct for user
type UpgradeInstanceRequestParams struct {
	// The ID of instance to be modified.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// New memory size of an instance shard. <ul><li>Unit: MB. </li><li>You can only modify one of the three parameters at a time: `MemSize`, `RedisShardNum`, and `RedisReplicasNum`. To modify one of them, you need to enter the other two, which are consistent with the original configuration specifications of the instance. </li><li>In case of capacity reduction, the new specification must be at least 1.3 times the used capacity; otherwise, the operation will fail.</li></ul>
	MemSize *uint64 `json:"MemSize,omitnil" name:"MemSize"`

	// New number of instance shards. <ul><li>This parameter is not required for standard architecture instances, but for cluster architecture instances. </li><li>For cluster architecture, you can only modify one of the three parameters at a time: `MemSize`, `RedisShardNum`, and `RedisReplicasNum`. To modify one of them, you need to enter the other two, which are consistent with the original configuration specifications of the instance. </li></ul>
	RedisShardNum *uint64 `json:"RedisShardNum,omitnil" name:"RedisShardNum"`

	// New replica quantity. <ul><li>You can only modify one of the three parameters at a time: `MemSize`, `RedisShardNum`, and `RedisReplicasNum`. To modify one of them, you need to enter the other two, which are consistent with the original configuration specifications of the instance. </li></ul>To modify the number of replicas in a multi-AZ instance, `NodeSet` must be passed in.</li></ul>
	RedisReplicasNum *uint64 `json:"RedisReplicasNum,omitnil" name:"RedisReplicasNum"`

	// Additional information for adding replicas for multi-AZ instances, including replica AZ and type (`NodeType` is `1`). This parameter is not required for single-AZ instances.
	NodeSet []*RedisNodeInfo `json:"NodeSet,omitnil" name:"NodeSet"`
}

type UpgradeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// The ID of instance to be modified.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// New memory size of an instance shard. <ul><li>Unit: MB. </li><li>You can only modify one of the three parameters at a time: `MemSize`, `RedisShardNum`, and `RedisReplicasNum`. To modify one of them, you need to enter the other two, which are consistent with the original configuration specifications of the instance. </li><li>In case of capacity reduction, the new specification must be at least 1.3 times the used capacity; otherwise, the operation will fail.</li></ul>
	MemSize *uint64 `json:"MemSize,omitnil" name:"MemSize"`

	// New number of instance shards. <ul><li>This parameter is not required for standard architecture instances, but for cluster architecture instances. </li><li>For cluster architecture, you can only modify one of the three parameters at a time: `MemSize`, `RedisShardNum`, and `RedisReplicasNum`. To modify one of them, you need to enter the other two, which are consistent with the original configuration specifications of the instance. </li></ul>
	RedisShardNum *uint64 `json:"RedisShardNum,omitnil" name:"RedisShardNum"`

	// New replica quantity. <ul><li>You can only modify one of the three parameters at a time: `MemSize`, `RedisShardNum`, and `RedisReplicasNum`. To modify one of them, you need to enter the other two, which are consistent with the original configuration specifications of the instance. </li></ul>To modify the number of replicas in a multi-AZ instance, `NodeSet` must be passed in.</li></ul>
	RedisReplicasNum *uint64 `json:"RedisReplicasNum,omitnil" name:"RedisReplicasNum"`

	// Additional information for adding replicas for multi-AZ instances, including replica AZ and type (`NodeType` is `1`). This parameter is not required for single-AZ instances.
	NodeSet []*RedisNodeInfo `json:"NodeSet,omitnil" name:"NodeSet"`
}

func (r *UpgradeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "MemSize")
	delete(f, "RedisShardNum")
	delete(f, "RedisReplicasNum")
	delete(f, "NodeSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeInstanceResponseParams struct {
	// Order ID
	DealId *string `json:"DealId,omitnil" name:"DealId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpgradeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeInstanceResponseParams `json:"Response"`
}

func (r *UpgradeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeInstanceVersionRequestParams struct {
	// Target instance type after the change, which is the same as the `TypeId` of the [CreateInstances](https://intl.cloud.tencent.com/document/api/239/20026?from_cn_redirect=1) API.
	// - For Redis 4.0 or later, a standard architecture instance can be upgraded to a cluster architecture instance on the same version; for example, you can upgrade from Redis 4.0 Standard Architecture to Redis 4.0 Cluster Architecture.
	// - Cross-version architecture upgrade is not supported; for example, you cannot upgrade from Redis 4.0 Standard Architecture to Redis 5.0 Cluster Architecture.
	// - The architecture of Redis 2.8 cannot be upgraded.
	// - Cluster architecture cannot be downgraded to standard architecture.
	TargetInstanceType *string `json:"TargetInstanceType,omitnil" name:"TargetInstanceType"`

	// Switch time. Valid values:
	// - `1`: Switch in the maintenance time.
	// - `2` (default value): Switch now.
	SwitchOption *int64 `json:"SwitchOption,omitnil" name:"SwitchOption"`

	// ID of the specified instance, such as `crs-xjhsdj****`. Log in to the [Redis console](https://console.cloud.tencent.com/redis#/), and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type UpgradeInstanceVersionRequest struct {
	*tchttp.BaseRequest
	
	// Target instance type after the change, which is the same as the `TypeId` of the [CreateInstances](https://intl.cloud.tencent.com/document/api/239/20026?from_cn_redirect=1) API.
	// - For Redis 4.0 or later, a standard architecture instance can be upgraded to a cluster architecture instance on the same version; for example, you can upgrade from Redis 4.0 Standard Architecture to Redis 4.0 Cluster Architecture.
	// - Cross-version architecture upgrade is not supported; for example, you cannot upgrade from Redis 4.0 Standard Architecture to Redis 5.0 Cluster Architecture.
	// - The architecture of Redis 2.8 cannot be upgraded.
	// - Cluster architecture cannot be downgraded to standard architecture.
	TargetInstanceType *string `json:"TargetInstanceType,omitnil" name:"TargetInstanceType"`

	// Switch time. Valid values:
	// - `1`: Switch in the maintenance time.
	// - `2` (default value): Switch now.
	SwitchOption *int64 `json:"SwitchOption,omitnil" name:"SwitchOption"`

	// ID of the specified instance, such as `crs-xjhsdj****`. Log in to the [Redis console](https://console.cloud.tencent.com/redis#/), and copy it in the instance list.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *UpgradeInstanceVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeInstanceVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetInstanceType")
	delete(f, "SwitchOption")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeInstanceVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeInstanceVersionResponseParams struct {
	// Order ID
	DealId *string `json:"DealId,omitnil" name:"DealId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpgradeInstanceVersionResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeInstanceVersionResponseParams `json:"Response"`
}

func (r *UpgradeInstanceVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeInstanceVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeProxyVersionRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The current proxy version
	CurrentProxyVersion *string `json:"CurrentProxyVersion,omitnil" name:"CurrentProxyVersion"`

	// Upgradeable redis version
	UpgradeProxyVersion *string `json:"UpgradeProxyVersion,omitnil" name:"UpgradeProxyVersion"`

	// `1` (upgrade immediately), `0` (upgrade during maintenance time)
	InstanceTypeUpgradeNow *int64 `json:"InstanceTypeUpgradeNow,omitnil" name:"InstanceTypeUpgradeNow"`
}

type UpgradeProxyVersionRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The current proxy version
	CurrentProxyVersion *string `json:"CurrentProxyVersion,omitnil" name:"CurrentProxyVersion"`

	// Upgradeable redis version
	UpgradeProxyVersion *string `json:"UpgradeProxyVersion,omitnil" name:"UpgradeProxyVersion"`

	// `1` (upgrade immediately), `0` (upgrade during maintenance time)
	InstanceTypeUpgradeNow *int64 `json:"InstanceTypeUpgradeNow,omitnil" name:"InstanceTypeUpgradeNow"`
}

func (r *UpgradeProxyVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeProxyVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "CurrentProxyVersion")
	delete(f, "UpgradeProxyVersion")
	delete(f, "InstanceTypeUpgradeNow")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeProxyVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeProxyVersionResponseParams struct {
	// Async task ID
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpgradeProxyVersionResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeProxyVersionResponseParams `json:"Response"`
}

func (r *UpgradeProxyVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeProxyVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeSmallVersionRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The current redis version
	CurrentRedisVersion *string `json:"CurrentRedisVersion,omitnil" name:"CurrentRedisVersion"`

	// Upgradeable redis version
	UpgradeRedisVersion *string `json:"UpgradeRedisVersion,omitnil" name:"UpgradeRedisVersion"`

	// `1` (upgrade immediately), `0` (upgrade during maintenance time)
	InstanceTypeUpgradeNow *int64 `json:"InstanceTypeUpgradeNow,omitnil" name:"InstanceTypeUpgradeNow"`
}

type UpgradeSmallVersionRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The current redis version
	CurrentRedisVersion *string `json:"CurrentRedisVersion,omitnil" name:"CurrentRedisVersion"`

	// Upgradeable redis version
	UpgradeRedisVersion *string `json:"UpgradeRedisVersion,omitnil" name:"UpgradeRedisVersion"`

	// `1` (upgrade immediately), `0` (upgrade during maintenance time)
	InstanceTypeUpgradeNow *int64 `json:"InstanceTypeUpgradeNow,omitnil" name:"InstanceTypeUpgradeNow"`
}

func (r *UpgradeSmallVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeSmallVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "CurrentRedisVersion")
	delete(f, "UpgradeRedisVersion")
	delete(f, "InstanceTypeUpgradeNow")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeSmallVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeSmallVersionResponseParams struct {
	// Async task ID
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpgradeSmallVersionResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeSmallVersionResponseParams `json:"Response"`
}

func (r *UpgradeSmallVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeSmallVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeVersionToMultiAvailabilityZonesRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Whether to support “Reading Local Nodes Only” feature after upgrading to multi-AZ deployment.
	// ul><li>`true`: The “Read Local Nodes Only” feature is supported. During the upgrade, you need to upgrade the proxy version and Redis kernel minor version simultaneously, which will involve data migration and may take hours to complete. </li><li>`false`: The “Read Local Nodes Only” feature is not supported. Upgrading to multi-AZ deployment will involve metadata migration only without affecting the service, which generally take less than three minutes to complete.</li></ul>
	UpgradeProxyAndRedisServer *bool `json:"UpgradeProxyAndRedisServer,omitnil" name:"UpgradeProxyAndRedisServer"`
}

type UpgradeVersionToMultiAvailabilityZonesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Whether to support “Reading Local Nodes Only” feature after upgrading to multi-AZ deployment.
	// ul><li>`true`: The “Read Local Nodes Only” feature is supported. During the upgrade, you need to upgrade the proxy version and Redis kernel minor version simultaneously, which will involve data migration and may take hours to complete. </li><li>`false`: The “Read Local Nodes Only” feature is not supported. Upgrading to multi-AZ deployment will involve metadata migration only without affecting the service, which generally take less than three minutes to complete.</li></ul>
	UpgradeProxyAndRedisServer *bool `json:"UpgradeProxyAndRedisServer,omitnil" name:"UpgradeProxyAndRedisServer"`
}

func (r *UpgradeVersionToMultiAvailabilityZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeVersionToMultiAvailabilityZonesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UpgradeProxyAndRedisServer")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeVersionToMultiAvailabilityZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeVersionToMultiAvailabilityZonesResponseParams struct {
	// Task ID
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpgradeVersionToMultiAvailabilityZonesResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeVersionToMultiAvailabilityZonesResponseParams `json:"Response"`
}

func (r *UpgradeVersionToMultiAvailabilityZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeVersionToMultiAvailabilityZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ZoneCapacityConf struct {
	// AZ ID, such as ap-guangzhou-3
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// AZ name
	ZoneName *string `json:"ZoneName,omitnil" name:"ZoneName"`

	// Whether a product is sold out in an AZ
	IsSaleout *bool `json:"IsSaleout,omitnil" name:"IsSaleout"`

	// Whether it is a default AZ
	IsDefault *bool `json:"IsDefault,omitnil" name:"IsDefault"`

	// Network type. basenet: basic network; vpcnet: VPC
	NetWorkType []*string `json:"NetWorkType,omitnil" name:"NetWorkType"`

	// Information of an AZ, such as product specifications in it
	ProductSet []*ProductConf `json:"ProductSet,omitnil" name:"ProductSet"`

	// AZ ID, such as 100003
	OldZoneId *int64 `json:"OldZoneId,omitnil" name:"OldZoneId"`
}