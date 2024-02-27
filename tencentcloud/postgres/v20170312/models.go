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

type AccountInfo struct {
	// Instance ID in the format of postgres-lnp6j617
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Account
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Account remarks
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Account status. 1: creating, 2: normal, 3: modifying, 4: resetting password, -1: deleting
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Account creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Account last modified time
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type AddDBInstanceToReadOnlyGroupRequestParams struct {
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// RO group ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`
}

type AddDBInstanceToReadOnlyGroupRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// RO group ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`
}

func (r *AddDBInstanceToReadOnlyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddDBInstanceToReadOnlyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "ReadOnlyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddDBInstanceToReadOnlyGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddDBInstanceToReadOnlyGroupResponseParams struct {
	// Task ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddDBInstanceToReadOnlyGroupResponse struct {
	*tchttp.BaseResponse
	Response *AddDBInstanceToReadOnlyGroupResponseParams `json:"Response"`
}

func (r *AddDBInstanceToReadOnlyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddDBInstanceToReadOnlyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AnalysisItems struct {
	// The name of the database queried by the slow query statement
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// The name of the user who executes the slow query statement
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// The slow query statement whose parameter values are abstracted
	NormalQuery *string `json:"NormalQuery,omitnil,omitempty" name:"NormalQuery"`

	// The address of the client that executes the slow query statement
	ClientAddr *string `json:"ClientAddr,omitnil,omitempty" name:"ClientAddr"`

	// The number of executions of the slow query statement during the specified period of time
	CallNum *uint64 `json:"CallNum,omitnil,omitempty" name:"CallNum"`

	// The ratio (in decimal form) of the number of executions of the slow query statement to that of all slow query statements during the specified period of time
	CallPercent *float64 `json:"CallPercent,omitnil,omitempty" name:"CallPercent"`

	// The total execution time of the slow query statement during the specified period of time
	CostTime *float64 `json:"CostTime,omitnil,omitempty" name:"CostTime"`

	// The ratio (in decimal form) of the total execution time of the slow query statement to that of all slow query statements during the specified period of time
	CostPercent *float64 `json:"CostPercent,omitnil,omitempty" name:"CostPercent"`

	// The shortest execution time (in ms) of the slow query statement during the specified period of time
	MinCostTime *float64 `json:"MinCostTime,omitnil,omitempty" name:"MinCostTime"`

	// The longest execution time (in ms) of the slow query statement during the specified period of time
	MaxCostTime *float64 `json:"MaxCostTime,omitnil,omitempty" name:"MaxCostTime"`

	// The average execution time (in ms) of the slow query statement during the specified period of time
	AvgCostTime *float64 `json:"AvgCostTime,omitnil,omitempty" name:"AvgCostTime"`

	// The timestamp when the slow query statement starts to execute for the first time during the specified period of time
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// The timestamp when the slow query statement starts to execute for the last time during the specified period of time
	LastTime *string `json:"LastTime,omitnil,omitempty" name:"LastTime"`
}

type BackupDownloadRestriction struct {
	// Type of the network restrictions for downloading backup files. Valid values: `NONE` (backups can be downloaded over both private and public networks), `INTRANET` (backups can only be downloaded over the private network), `CUSTOMIZE` (backups can be downloaded over specified VPCs or at specified IPs).
	RestrictionType *string `json:"RestrictionType,omitnil,omitempty" name:"RestrictionType"`

	// Whether VPC is allowed. Valid values: `ALLOW` (allow), `DENY` (deny).
	VpcRestrictionEffect *string `json:"VpcRestrictionEffect,omitnil,omitempty" name:"VpcRestrictionEffect"`

	// Whether it is allowed to download the VPC ID list of the backup files.
	VpcIdSet []*string `json:"VpcIdSet,omitnil,omitempty" name:"VpcIdSet"`

	// Whether IP is allowed. Valid values: `ALLOW` (allow), `DENY` (deny).
	IpRestrictionEffect *string `json:"IpRestrictionEffect,omitnil,omitempty" name:"IpRestrictionEffect"`

	// Whether it is allowed to download IP list of the backup files.
	IpSet []*string `json:"IpSet,omitnil,omitempty" name:"IpSet"`
}

type BackupPlan struct {
	// Backup cycle
	BackupPeriod *string `json:"BackupPeriod,omitnil,omitempty" name:"BackupPeriod"`

	// Retention period of basic backups
	BaseBackupRetentionPeriod *uint64 `json:"BaseBackupRetentionPeriod,omitnil,omitempty" name:"BaseBackupRetentionPeriod"`

	// The earliest time to start a backup
	MinBackupStartTime *string `json:"MinBackupStartTime,omitnil,omitempty" name:"MinBackupStartTime"`

	// The latest time to start a backup
	MaxBackupStartTime *string `json:"MaxBackupStartTime,omitnil,omitempty" name:"MaxBackupStartTime"`
}

type BackupSummary struct {
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Number of log backups of an instance
	LogBackupCount *uint64 `json:"LogBackupCount,omitnil,omitempty" name:"LogBackupCount"`

	// Size of log backups of an instance
	LogBackupSize *uint64 `json:"LogBackupSize,omitnil,omitempty" name:"LogBackupSize"`

	// Number of manually created full backups of an instance
	ManualBaseBackupCount *uint64 `json:"ManualBaseBackupCount,omitnil,omitempty" name:"ManualBaseBackupCount"`

	// Size of manually created full backups of an instance
	ManualBaseBackupSize *uint64 `json:"ManualBaseBackupSize,omitnil,omitempty" name:"ManualBaseBackupSize"`

	// Number of automatically created full backups of an instance
	AutoBaseBackupCount *uint64 `json:"AutoBaseBackupCount,omitnil,omitempty" name:"AutoBaseBackupCount"`

	// Size of automatically created full backups of an instance
	AutoBaseBackupSize *uint64 `json:"AutoBaseBackupSize,omitnil,omitempty" name:"AutoBaseBackupSize"`

	// Total number of backups
	TotalBackupCount *uint64 `json:"TotalBackupCount,omitnil,omitempty" name:"TotalBackupCount"`

	// Total backup size
	TotalBackupSize *uint64 `json:"TotalBackupSize,omitnil,omitempty" name:"TotalBackupSize"`
}

type BaseBackup struct {
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Unique ID of a backup file
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Backup file name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Backup method, including physical and logical.
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// Backup mode, including automatic and manual.
	BackupMode *string `json:"BackupMode,omitnil,omitempty" name:"BackupMode"`

	// Backup task status
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// Backup set size in bytes
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`

	// Backup start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Backup end time
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// Backup expiration time
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

type ClassInfo struct {
	// Specification ID
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// Number of CPU cores
	CPU *uint64 `json:"CPU,omitnil,omitempty" name:"CPU"`

	// Memory size in MB
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Maximum storage capacity in GB supported by this specification
	MaxStorage *uint64 `json:"MaxStorage,omitnil,omitempty" name:"MaxStorage"`

	// Minimum storage capacity in GB supported by this specification
	MinStorage *uint64 `json:"MinStorage,omitnil,omitempty" name:"MinStorage"`

	// Estimated QPS for this specification
	QPS *uint64 `json:"QPS,omitnil,omitempty" name:"QPS"`
}

// Predefined struct for user
type CloneDBInstanceRequestParams struct {
	// ID of the original instance to be cloned.
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Purchasable code, which can be obtained from the `SpecCode` field in the return value of the [DescribeClasses](https://intl.cloud.tencent.com/document/api/409/89019?from_cn_redirect=1) API.
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// Instance storage capacity in GB.
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Validity period in months. Valid values:
	// <li>Monthly subscription: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	// <li>Pay-as-you-go: `1`.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Auto-renewal flag. Valid values:
	// <li>`0`: Manual renewal.
	// <li>`1`: Automatic renewal.
	// Default value: `0`.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// VPC ID in the format of `vpc-xxxxxxx`, which can be obtained in the console or from the `unVpcId` field in the return value of the [DescribeVpcEx](https://intl.cloud.tencent.com/document/api/215/1372?from_cn_redirect=1) API.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// VPC subnet ID in the format of `subnet-xxxxxxxx`, which can be obtained in the console or from the `unSubnetId` field in the return value of the [DescribeSubnets](https://intl.cloud.tencent.com/document/api/215/15784?from_cn_redirect=1) API.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Name of the newly purchased instance, which can contain up to 60 letters, digits, or symbols (-_). If this parameter is not specified, "Unnamed" will be displayed by default.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Instance billing mode. Valid values:
	// <li>`PREPAID`: Monthly subscription.
	// <li>`POSTPAID_BY_HOUR`: Pay-as-you-go.
	// Default value: `PREPAID`.
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Security group of the instance, which can be obtained from the `sgld` field in the return value of the [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/api/215/15808?from_cn_redirect=1) API. If this parameter is not specified, the default security group will be bound.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// The information of tags to be bound with the instance, which is left empty by default. This parameter can be obtained from the `Tags` field in the return value of the [DescribeTags](https://intl.cloud.tencent.com/document/api/651/35316?from_cn_redirect=1) API.
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`

	// Deployment information of the instance node, which will display the information of each AZ when the instance node is deployed across multiple AZs.
	// The information of AZ can be obtained from the `Zone` field in the return value of the [DescribeZones](https://intl.cloud.tencent.com/document/api/409/16769?from_cn_redirect=1) API.
	DBNodeSet []*DBNode `json:"DBNodeSet,omitnil,omitempty" name:"DBNodeSet"`

	// Whether to use vouchers automatically. Valid values:
	// <li>`0`: No.
	// <li>`1`: Yes.
	// Default value: `0`.
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Voucher ID list.
	VoucherIds *string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// Campaign ID.
	ActivityId *int64 `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// Basic backup set ID.
	BackupSetId *string `json:"BackupSetId,omitnil,omitempty" name:"BackupSetId"`

	// Restoration point in time.
	RecoveryTargetTime *string `json:"RecoveryTargetTime,omitnil,omitempty" name:"RecoveryTargetTime"`

	// Primary-standby sync mode. Valid values:  
	// <li>`Semi-sync`
	// <li>`Async`
	// Default value for the primary instance: `Semi-sync`.
	// Default value for the standby instance: `Async`.
	SyncMode *string `json:"SyncMode,omitnil,omitempty" name:"SyncMode"`
}

type CloneDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// ID of the original instance to be cloned.
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Purchasable code, which can be obtained from the `SpecCode` field in the return value of the [DescribeClasses](https://intl.cloud.tencent.com/document/api/409/89019?from_cn_redirect=1) API.
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// Instance storage capacity in GB.
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Validity period in months. Valid values:
	// <li>Monthly subscription: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	// <li>Pay-as-you-go: `1`.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Auto-renewal flag. Valid values:
	// <li>`0`: Manual renewal.
	// <li>`1`: Automatic renewal.
	// Default value: `0`.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// VPC ID in the format of `vpc-xxxxxxx`, which can be obtained in the console or from the `unVpcId` field in the return value of the [DescribeVpcEx](https://intl.cloud.tencent.com/document/api/215/1372?from_cn_redirect=1) API.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// VPC subnet ID in the format of `subnet-xxxxxxxx`, which can be obtained in the console or from the `unSubnetId` field in the return value of the [DescribeSubnets](https://intl.cloud.tencent.com/document/api/215/15784?from_cn_redirect=1) API.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Name of the newly purchased instance, which can contain up to 60 letters, digits, or symbols (-_). If this parameter is not specified, "Unnamed" will be displayed by default.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Instance billing mode. Valid values:
	// <li>`PREPAID`: Monthly subscription.
	// <li>`POSTPAID_BY_HOUR`: Pay-as-you-go.
	// Default value: `PREPAID`.
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Security group of the instance, which can be obtained from the `sgld` field in the return value of the [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/api/215/15808?from_cn_redirect=1) API. If this parameter is not specified, the default security group will be bound.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// The information of tags to be bound with the instance, which is left empty by default. This parameter can be obtained from the `Tags` field in the return value of the [DescribeTags](https://intl.cloud.tencent.com/document/api/651/35316?from_cn_redirect=1) API.
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`

	// Deployment information of the instance node, which will display the information of each AZ when the instance node is deployed across multiple AZs.
	// The information of AZ can be obtained from the `Zone` field in the return value of the [DescribeZones](https://intl.cloud.tencent.com/document/api/409/16769?from_cn_redirect=1) API.
	DBNodeSet []*DBNode `json:"DBNodeSet,omitnil,omitempty" name:"DBNodeSet"`

	// Whether to use vouchers automatically. Valid values:
	// <li>`0`: No.
	// <li>`1`: Yes.
	// Default value: `0`.
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Voucher ID list.
	VoucherIds *string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// Campaign ID.
	ActivityId *int64 `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// Basic backup set ID.
	BackupSetId *string `json:"BackupSetId,omitnil,omitempty" name:"BackupSetId"`

	// Restoration point in time.
	RecoveryTargetTime *string `json:"RecoveryTargetTime,omitnil,omitempty" name:"RecoveryTargetTime"`

	// Primary-standby sync mode. Valid values:  
	// <li>`Semi-sync`
	// <li>`Async`
	// Default value for the primary instance: `Semi-sync`.
	// Default value for the standby instance: `Async`.
	SyncMode *string `json:"SyncMode,omitnil,omitempty" name:"SyncMode"`
}

func (r *CloneDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "SpecCode")
	delete(f, "Storage")
	delete(f, "Period")
	delete(f, "AutoRenewFlag")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Name")
	delete(f, "InstanceChargeType")
	delete(f, "SecurityGroupIds")
	delete(f, "ProjectId")
	delete(f, "TagList")
	delete(f, "DBNodeSet")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "ActivityId")
	delete(f, "BackupSetId")
	delete(f, "RecoveryTargetTime")
	delete(f, "SyncMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloneDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloneDBInstanceResponseParams struct {
	// Order ID.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// Bill ID.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	BillId *string `json:"BillId,omitnil,omitempty" name:"BillId"`

	// ID of the cloned instance, which will be returned only when the instance is pay-as-you-go.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloneDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CloneDBInstanceResponseParams `json:"Response"`
}

func (r *CloneDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseDBExtranetAccessRequestParams struct {
	// Instance ID in the format of postgres-6r233v55
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Whether to disable public network access over IPv6 address. Valid values: 1 (yes), 0 (no)
	IsIpv6 *int64 `json:"IsIpv6,omitnil,omitempty" name:"IsIpv6"`
}

type CloseDBExtranetAccessRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of postgres-6r233v55
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Whether to disable public network access over IPv6 address. Valid values: 1 (yes), 0 (no)
	IsIpv6 *int64 `json:"IsIpv6,omitnil,omitempty" name:"IsIpv6"`
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
	delete(f, "DBInstanceId")
	delete(f, "IsIpv6")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseDBExtranetAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseDBExtranetAccessResponseParams struct {
	// Async task flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type CloseServerlessDBExtranetAccessRequestParams struct {
	// Unique ID of an instance
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Instance name
	DBInstanceName *string `json:"DBInstanceName,omitnil,omitempty" name:"DBInstanceName"`
}

type CloseServerlessDBExtranetAccessRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of an instance
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Instance name
	DBInstanceName *string `json:"DBInstanceName,omitnil,omitempty" name:"DBInstanceName"`
}

func (r *CloseServerlessDBExtranetAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseServerlessDBExtranetAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "DBInstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseServerlessDBExtranetAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseServerlessDBExtranetAccessResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseServerlessDBExtranetAccessResponse struct {
	*tchttp.BaseResponse
	Response *CloseServerlessDBExtranetAccessResponseParams `json:"Response"`
}

func (r *CloseServerlessDBExtranetAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseServerlessDBExtranetAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBaseBackupRequestParams struct {
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

type CreateBaseBackupRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

func (r *CreateBaseBackupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBaseBackupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBaseBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBaseBackupResponseParams struct {
	// Full backup set ID
	BaseBackupId *string `json:"BaseBackupId,omitnil,omitempty" name:"BaseBackupId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateBaseBackupResponse struct {
	*tchttp.BaseResponse
	Response *CreateBaseBackupResponseParams `json:"Response"`
}

func (r *CreateBaseBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBaseBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceNetworkAccessRequestParams struct {
	// Instance ID in the format of postgres-6bwgamo3.
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Unified VPC ID.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Whether to manually assign the VIP. Valid values: `true` (manually assign), `false` (automatically assign).
	IsAssignVip *bool `json:"IsAssignVip,omitnil,omitempty" name:"IsAssignVip"`

	// Target VIP.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`
}

type CreateDBInstanceNetworkAccessRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of postgres-6bwgamo3.
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Unified VPC ID.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Whether to manually assign the VIP. Valid values: `true` (manually assign), `false` (automatically assign).
	IsAssignVip *bool `json:"IsAssignVip,omitnil,omitempty" name:"IsAssignVip"`

	// Target VIP.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`
}

func (r *CreateDBInstanceNetworkAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceNetworkAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "IsAssignVip")
	delete(f, "Vip")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBInstanceNetworkAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceNetworkAccessResponseParams struct {
	// Task ID.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDBInstanceNetworkAccessResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBInstanceNetworkAccessResponseParams `json:"Response"`
}

func (r *CreateDBInstanceNetworkAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceNetworkAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstancesRequestParams struct {
	// Purchasable specification ID, which can be obtained through the `SpecCode` field in the returned value of the `DescribeClasses` API.
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// Instance capacity size in GB.
	Storage *uint64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Number of instances purchased at a time. Value range: 1-100.
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// Length of purchase in months. Currently, only 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, and 36 are supported.
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// AZ ID, which can be obtained through the `Zone` field in the returned value of the `DescribeZones` API.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// PostgreSQL version. If it is specified, an instance running the latest kernel of PostgreSQL `DBVersion` will be created. You must pass in at least one of the following parameters: DBVersion, DBMajorVersion, DBKernelVersion.
	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`

	// Instance billing type.
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Whether to automatically use vouchers. 1: yes, 0: no. Default value: no.
	AutoVoucher *uint64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Voucher ID list (only one voucher can be specified currently).
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// VPC ID.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// VPC subnet ID.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Renewal flag. 0: normal renewal (default), 1: auto-renewal.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// Activity ID
	ActivityId *int64 `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// Instance name (which will be supported in the future)
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Whether to support IPv6 address access. Valid values: 1 (yes), 0 (no)
	NeedSupportIpv6 *uint64 `json:"NeedSupportIpv6,omitnil,omitempty" name:"NeedSupportIpv6"`

	// The information of tags to be associated with instances. This parameter is left empty by default.
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`

	// Security group ID
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// PostgreSQL major version. If it is specified, an instance running the latest kernel of PostgreSQL `DBMajorVersion` will be created. You must pass in at least one of the following parameters: DBMajorVersion, DBVersion, DBKernelVersion.
	DBMajorVersion *string `json:"DBMajorVersion,omitnil,omitempty" name:"DBMajorVersion"`

	// PostgreSQL kernel version. If it is specified, an instance running the latest kernel of PostgreSQL `DBKernelVersion` will be created. You must pass in one of the following parameters: DBKernelVersion, DBVersion, DBMajorVersion.
	DBKernelVersion *string `json:"DBKernelVersion,omitnil,omitempty" name:"DBKernelVersion"`
}

type CreateDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Purchasable specification ID, which can be obtained through the `SpecCode` field in the returned value of the `DescribeClasses` API.
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// Instance capacity size in GB.
	Storage *uint64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Number of instances purchased at a time. Value range: 1-100.
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// Length of purchase in months. Currently, only 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, and 36 are supported.
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// AZ ID, which can be obtained through the `Zone` field in the returned value of the `DescribeZones` API.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Project ID.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// PostgreSQL version. If it is specified, an instance running the latest kernel of PostgreSQL `DBVersion` will be created. You must pass in at least one of the following parameters: DBVersion, DBMajorVersion, DBKernelVersion.
	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`

	// Instance billing type.
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Whether to automatically use vouchers. 1: yes, 0: no. Default value: no.
	AutoVoucher *uint64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Voucher ID list (only one voucher can be specified currently).
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// VPC ID.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// VPC subnet ID.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Renewal flag. 0: normal renewal (default), 1: auto-renewal.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// Activity ID
	ActivityId *int64 `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// Instance name (which will be supported in the future)
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Whether to support IPv6 address access. Valid values: 1 (yes), 0 (no)
	NeedSupportIpv6 *uint64 `json:"NeedSupportIpv6,omitnil,omitempty" name:"NeedSupportIpv6"`

	// The information of tags to be associated with instances. This parameter is left empty by default.
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`

	// Security group ID
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// PostgreSQL major version. If it is specified, an instance running the latest kernel of PostgreSQL `DBMajorVersion` will be created. You must pass in at least one of the following parameters: DBMajorVersion, DBVersion, DBKernelVersion.
	DBMajorVersion *string `json:"DBMajorVersion,omitnil,omitempty" name:"DBMajorVersion"`

	// PostgreSQL kernel version. If it is specified, an instance running the latest kernel of PostgreSQL `DBKernelVersion` will be created. You must pass in one of the following parameters: DBKernelVersion, DBVersion, DBMajorVersion.
	DBKernelVersion *string `json:"DBKernelVersion,omitnil,omitempty" name:"DBKernelVersion"`
}

func (r *CreateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpecCode")
	delete(f, "Storage")
	delete(f, "InstanceCount")
	delete(f, "Period")
	delete(f, "Zone")
	delete(f, "ProjectId")
	delete(f, "DBVersion")
	delete(f, "InstanceChargeType")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "AutoRenewFlag")
	delete(f, "ActivityId")
	delete(f, "Name")
	delete(f, "NeedSupportIpv6")
	delete(f, "TagList")
	delete(f, "SecurityGroupIds")
	delete(f, "DBMajorVersion")
	delete(f, "DBKernelVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstancesResponseParams struct {
	// Order number list. Each instance corresponds to an order number.
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// Bill ID of frozen fees
	BillId *string `json:"BillId,omitnil,omitempty" name:"BillId"`

	// ID set of instances which have been created successfully. The parameter value will be returned only when the billing mode is postpaid.
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitnil,omitempty" name:"DBInstanceIdSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBInstancesResponseParams `json:"Response"`
}

func (r *CreateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstancesRequestParams struct {
	// Primary AZ of the instance in the format of `ap-guangzhou-3`. To support multiple AZs, add information of the primary and standby AZs in the `DBNodeSet.N` field.
	// The information of AZ can be obtained from the `Zone` field in the return value of the [DescribeZones](https://intl.cloud.tencent.com/document/api/409/16769?from_cn_redirect=1) API.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Purchasable code, which can be obtained from the `SpecCode` field in the return value of the [DescribeClasses](https://intl.cloud.tencent.com/document/api/409/89019?from_cn_redirect=1) API.
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// Instance storage capacity in GB
	Storage *uint64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// The number of instances to be purchased at a time. Value range: 1-10. To purchase more than 10 instances each time, you can make multiple calls.
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// Validity period in months.
	// <li>Monthly subscription: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	// <li>Pay-as-you-go: `1`.
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Instance character set. Valid values: 
	// <li> `UTF8`
	// <li> `LATIN1`
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// Username of the instance root account, which has the following rules:
	// <li>It must contain 1–16 letters , digits, or underscores
	// <li>It can't be `postgres`.
	// <li>It can't start with a digit or `pg_`.
	// <li>All rules are case-insensitive.
	AdminName *string `json:"AdminName,omitnil,omitempty" name:"AdminName"`

	// Password of the instance root account, which must contain 8-32 characters (above 12 characters preferably). It cannot begin with "/",
	// and must contain the following 4 types of characters.
	// <li>Lowercase letters: [a–z]
	// <li>Uppercase letters: [A–Z]
	// <li>Digits: 0-9
	// <li>Symbols: ()`~!@#$%^&*-+=_|{}[]:;'<>,.?/
	AdminPassword *string `json:"AdminPassword,omitnil,omitempty" name:"AdminPassword"`

	// The major PostgreSQL version number, which can be queried by the [DescribeDBVersions](https://intl.cloud.tencent.com/document/api/409/89018?from_cn_redirect=1) API. Valid values: `10`, `11`, `12`, `13`, `14`, `15`.
	// When only this parameter is specified, an instance running the latest kernel version of the latest minor version will be created based on this major version.
	// You must pass in at least one of the following parameters: `DBMajorVersion`, `DBVersion`, DBKernelVersion`. If you don't need a minor version, just pass in `DBMajorVersion`.
	DBMajorVersion *string `json:"DBMajorVersion,omitnil,omitempty" name:"DBMajorVersion"`

	// Number of the major PostgreSQL community version and minor version, which can be queried by the [DescribeDBVersions](https://intl.cloud.tencent.com/document/api/409/89018?from_cn_redirect=1) API.
	// If it is specified, an instance running the latest kernel version will be created based on the community minor version.
	// You must pass in at least one of the following parameters: `DBMajorVersion`, `DBVersion`, DBKernelVersion`.
	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`

	// PostgreSQL kernel version number, which can be queried by the [DescribeDBVersions](https://intl.cloud.tencent.com/document/api/409/89018?from_cn_redirect=1) API.
	// PostgreSQL kernel version number. If it is specified, an instance running the specified kernel version will be created. Passing in this parameter in other scenarios is not supported.This parameter is only used to specify a kernel version, which serves no other purposes.
	DBKernelVersion *string `json:"DBKernelVersion,omitnil,omitempty" name:"DBKernelVersion"`

	// Instance billing mode. Valid values:
	// <li>`PREPAID`: Monthly subscription
	// <li>`POSTPAID_BY_HOUR`: Pay-as-you-go
	// Default value: `PREPAID`.
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// VPC ID in the format of `vpc-xxxxxxx`. To obtain valid VPC IDs, you can log in to the console or call [DescribeVpcEx](https://intl.cloud.tencent.com/document/api/215/1372?from_cn_redirect=1) and look for the `unVpcId` fields in the response.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// VPC subnet ID in the format of `subnet-xxxxxxxx`, u200cwhich can be obtained in the console or from the `unSubnetId` field in the return value of the [DescribeSubnets](https://intl.cloud.tencent.com/document/api/215/15784?from_cn_redirect=1) API.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Deployment information of the instance node, which will display the information of each AZ when the instance node is deployed across multiple AZs.
	// The information of AZ can be obtained from the `Zone` field in the return value of the [DescribeZones](https://intl.cloud.tencent.com/document/api/409/16769?from_cn_redirect=1) API.
	DBNodeSet []*DBNode `json:"DBNodeSet,omitnil,omitempty" name:"DBNodeSet"`

	// Auto-renewal flag. Valid values:
	// <li>`0`: Manual renewal.
	// <li> `1`: Automatic renewal.
	// Default value: `0`.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// Whether to use vouchers automatically. Valid values:
	// <li>`0`: No.
	// <li>`1`: Yes.
	// Default value: `0`.
	AutoVoucher *uint64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Voucher ID list. Currently, you can specify only one voucher.
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Campaign ID
	ActivityId *int64 `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// Instance name, which can contain up to 60 letters, digits, hyphens, and symbols (_-). If this parameter is not specified, "Unnamed" will be displayed by default.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// The information of tags to be bound with the instance, which is left empty by default. This parameter can be obtained from the `Tags` field in the return value of the [DescribeTags](https://intl.cloud.tencent.com/document/api/651/35316?from_cn_redirect=1) API.
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`

	// Security group of the instance, which can be obtained from the `sgld` field in the return value of the [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/api/215/15808?from_cn_redirect=1) API. If this parameter is not specified, the default security group will be bound.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Whether to support TDE. Valid values:
	// <li>`0`: No.
	// <li>`1`: Yes.
	// Default value: `0`.
	// For more information, see [TDE]u200d(https://www.tencentcloud.com/document/product/409/47765).
	NeedSupportTDE *uint64 `json:"NeedSupportTDE,omitnil,omitempty" name:"NeedSupportTDE"`

	// KeyId of custom key, which is required if you select custom key encryption. It is also the unique CMK identifier.
	// For more information on creating `KeyId`, see [Enabling TDE](https://www.tencentcloud.com/document/product/409/47762).
	KMSKeyId *string `json:"KMSKeyId,omitnil,omitempty" name:"KMSKeyId"`

	// The region where the KMS service is enabled. When `KMSRegion` is left empty, the current region will be selected by default.  If the current region does not support KMS, you must select another region that does.
	// For more information on `KMSRegion`, see [Enabling TDE](https://intl.cloud.tencent.com/document/product/409/71749?from_cn_redirect=1).
	KMSRegion *string `json:"KMSRegion,omitnil,omitempty" name:"KMSRegion"`

	// Database engines. Valid values:
	// <li>`postgresql`: TencentDB for PostgreSQL
	// <li>`mssql_compatible`: MSSQL compatible-TencentDB for PostgreSQL
	// Default value: `postgresql`.
	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`

	// Configuration information of database engine in the following format:
	// {"$key1":"$value1", "$key2":"$value2"}
	// Valid values:
	// mssql_compatible engine:
	// <li>`migrationMode`: Database mode. Valid values: `single-db` (single-database mode), `multi-db` (multi-database mode). Default value: `single-db`.
	// <li>`defaultLocale`: Default locale, which can’t be modified after the initialization. Default value: `en_US`. Valid values:
	// "af_ZA", "sq_AL", "ar_DZ", "ar_BH", "ar_EG", "ar_IQ", "ar_JO", "ar_KW", "ar_LB", "ar_LY", "ar_MA", "ar_OM", "ar_QA", "ar_SA", "ar_SY", "ar_TN", "ar_AE", "ar_YE", "hy_AM", "az_Cyrl_AZ", "az_Latn_AZ", "eu_ES", "be_BY", "bg_BG", "ca_ES", "zh_HK", "zh_MO", "zh_CN", "zh_SG", "zh_TW", "hr_HR", "cs_CZ", "da_DK", "nl_BE", "nl_NL", "en_AU", "en_BZ", "en_CA", "en_IE", "en_JM", "en_NZ", "en_PH", "en_ZA", "en_TT", "en_GB", "en_US", "en_ZW", "et_EE", "fo_FO", "fa_IR", "fi_FI", "fr_BE", "fr_CA", "fr_FR", "fr_LU", "fr_MC", "fr_CH", "mk_MK", "ka_GE", "de_AT", "de_DE", "de_LI", "de_LU", "de_CH", "el_GR", "gu_IN", "he_IL", "hi_IN", "hu_HU", "is_IS", "id_ID", "it_IT", "it_CH", "ja_JP", "kn_IN", "kok_IN", "ko_KR", "ky_KG", "lv_LV", "lt_LT", "ms_BN", "ms_MY", "mr_IN", "mn_MN", "nb_NO", "nn_NO", "pl_PL", "pt_BR", "pt_PT", "pa_IN", "ro_RO", "ru_RU", "sa_IN", "sr_Cyrl_RS", "sr_Latn_RS", "sk_SK", "sl_SI", "es_AR", "es_BO", "es_CL", "es_CO", "es_CR", "es_DO", "es_EC", "es_SV", "es_GT", "es_HN", "es_MX", "es_NI", "es_PA", "es_PY","es_PE", "es_PR", "es_ES", "es_TRADITIONAL", "es_UY", "es_VE", "sw_KE", "sv_FI", "sv_SE", "tt_RU", "te_IN", "th_TH", "tr_TR", "uk_UA", "ur_IN", "ur_PK", "uz_Cyrl_UZ", "uz_Latn_UZ", "vi_VN".
	// <li>`serverCollationName`: Default collation name, which can’t be modified after the initialization. Default value: "bbf_unicode_general_ci_as". Valid values: "bbf_unicode_cp1_ci_as", "bbf_unicode_CP1250_ci_as", "bbf_unicode_CP1251_ci_as", "bbf_unicode_cp1253_ci_as", "bbf_unicode_cp1254_ci_as", "bbf_unicode_cp1255_ci_as", "bbf_unicode_cp1256_ci_as", "bbf_unicode_cp1257_ci_as", "bbf_unicode_cp1258_ci_as", "bbf_unicode_cp874_ci_as", "sql_latin1_general_cp1250_ci_as", "sql_latin1_general_cp1251_ci_as", "sql_latin1_general_cp1_ci_as", "sql_latin1_general_cp1253_ci_as", "sql_latin1_general_cp1254_ci_as", "sql_latin1_general_cp1255_ci_as","sql_latin1_general_cp1256_ci_as", "sql_latin1_general_cp1257_ci_as", "sql_latin1_general_cp1258_ci_as", "chinese_prc_ci_as", "cyrillic_general_ci_as", "finnish_swedish_ci_as", "french_ci_as", "japanese_ci_as", "korean_wansung_ci_as", "latin1_general_ci_as", "modern_spanish_ci_as", "polish_ci_as", "thai_ci_as", "traditional_spanish_ci_as", "turkish_ci_as", "ukrainian_ci_as", "vietnamese_ci_as"。
	DBEngineConfig *string `json:"DBEngineConfig,omitnil,omitempty" name:"DBEngineConfig"`

	// Primary-standby sync mode. Valid values:  
	// <li>`Semi-sync`
	// <li>`Async`
	// Default value for the primary instance: `Semi-sync`.
	// Default value for the standby instance: `Async`.
	SyncMode *string `json:"SyncMode,omitnil,omitempty" name:"SyncMode"`

	// Whether IPv6 is supported.
	// <li>`0`: No.
	// <li>`1`: Yes.
	// Default value: `0`.
	NeedSupportIpv6 *uint64 `json:"NeedSupportIpv6,omitnil,omitempty" name:"NeedSupportIpv6"`
}

type CreateInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Primary AZ of the instance in the format of `ap-guangzhou-3`. To support multiple AZs, add information of the primary and standby AZs in the `DBNodeSet.N` field.
	// The information of AZ can be obtained from the `Zone` field in the return value of the [DescribeZones](https://intl.cloud.tencent.com/document/api/409/16769?from_cn_redirect=1) API.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Purchasable code, which can be obtained from the `SpecCode` field in the return value of the [DescribeClasses](https://intl.cloud.tencent.com/document/api/409/89019?from_cn_redirect=1) API.
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// Instance storage capacity in GB
	Storage *uint64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// The number of instances to be purchased at a time. Value range: 1-10. To purchase more than 10 instances each time, you can make multiple calls.
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// Validity period in months.
	// <li>Monthly subscription: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	// <li>Pay-as-you-go: `1`.
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Instance character set. Valid values: 
	// <li> `UTF8`
	// <li> `LATIN1`
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`

	// Username of the instance root account, which has the following rules:
	// <li>It must contain 1–16 letters , digits, or underscores
	// <li>It can't be `postgres`.
	// <li>It can't start with a digit or `pg_`.
	// <li>All rules are case-insensitive.
	AdminName *string `json:"AdminName,omitnil,omitempty" name:"AdminName"`

	// Password of the instance root account, which must contain 8-32 characters (above 12 characters preferably). It cannot begin with "/",
	// and must contain the following 4 types of characters.
	// <li>Lowercase letters: [a–z]
	// <li>Uppercase letters: [A–Z]
	// <li>Digits: 0-9
	// <li>Symbols: ()`~!@#$%^&*-+=_|{}[]:;'<>,.?/
	AdminPassword *string `json:"AdminPassword,omitnil,omitempty" name:"AdminPassword"`

	// The major PostgreSQL version number, which can be queried by the [DescribeDBVersions](https://intl.cloud.tencent.com/document/api/409/89018?from_cn_redirect=1) API. Valid values: `10`, `11`, `12`, `13`, `14`, `15`.
	// When only this parameter is specified, an instance running the latest kernel version of the latest minor version will be created based on this major version.
	// You must pass in at least one of the following parameters: `DBMajorVersion`, `DBVersion`, DBKernelVersion`. If you don't need a minor version, just pass in `DBMajorVersion`.
	DBMajorVersion *string `json:"DBMajorVersion,omitnil,omitempty" name:"DBMajorVersion"`

	// Number of the major PostgreSQL community version and minor version, which can be queried by the [DescribeDBVersions](https://intl.cloud.tencent.com/document/api/409/89018?from_cn_redirect=1) API.
	// If it is specified, an instance running the latest kernel version will be created based on the community minor version.
	// You must pass in at least one of the following parameters: `DBMajorVersion`, `DBVersion`, DBKernelVersion`.
	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`

	// PostgreSQL kernel version number, which can be queried by the [DescribeDBVersions](https://intl.cloud.tencent.com/document/api/409/89018?from_cn_redirect=1) API.
	// PostgreSQL kernel version number. If it is specified, an instance running the specified kernel version will be created. Passing in this parameter in other scenarios is not supported.This parameter is only used to specify a kernel version, which serves no other purposes.
	DBKernelVersion *string `json:"DBKernelVersion,omitnil,omitempty" name:"DBKernelVersion"`

	// Instance billing mode. Valid values:
	// <li>`PREPAID`: Monthly subscription
	// <li>`POSTPAID_BY_HOUR`: Pay-as-you-go
	// Default value: `PREPAID`.
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// VPC ID in the format of `vpc-xxxxxxx`. To obtain valid VPC IDs, you can log in to the console or call [DescribeVpcEx](https://intl.cloud.tencent.com/document/api/215/1372?from_cn_redirect=1) and look for the `unVpcId` fields in the response.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// VPC subnet ID in the format of `subnet-xxxxxxxx`, u200cwhich can be obtained in the console or from the `unSubnetId` field in the return value of the [DescribeSubnets](https://intl.cloud.tencent.com/document/api/215/15784?from_cn_redirect=1) API.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Deployment information of the instance node, which will display the information of each AZ when the instance node is deployed across multiple AZs.
	// The information of AZ can be obtained from the `Zone` field in the return value of the [DescribeZones](https://intl.cloud.tencent.com/document/api/409/16769?from_cn_redirect=1) API.
	DBNodeSet []*DBNode `json:"DBNodeSet,omitnil,omitempty" name:"DBNodeSet"`

	// Auto-renewal flag. Valid values:
	// <li>`0`: Manual renewal.
	// <li> `1`: Automatic renewal.
	// Default value: `0`.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// Whether to use vouchers automatically. Valid values:
	// <li>`0`: No.
	// <li>`1`: Yes.
	// Default value: `0`.
	AutoVoucher *uint64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Voucher ID list. Currently, you can specify only one voucher.
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Campaign ID
	ActivityId *int64 `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// Instance name, which can contain up to 60 letters, digits, hyphens, and symbols (_-). If this parameter is not specified, "Unnamed" will be displayed by default.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// The information of tags to be bound with the instance, which is left empty by default. This parameter can be obtained from the `Tags` field in the return value of the [DescribeTags](https://intl.cloud.tencent.com/document/api/651/35316?from_cn_redirect=1) API.
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`

	// Security group of the instance, which can be obtained from the `sgld` field in the return value of the [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/api/215/15808?from_cn_redirect=1) API. If this parameter is not specified, the default security group will be bound.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Whether to support TDE. Valid values:
	// <li>`0`: No.
	// <li>`1`: Yes.
	// Default value: `0`.
	// For more information, see [TDE]u200d(https://www.tencentcloud.com/document/product/409/47765).
	NeedSupportTDE *uint64 `json:"NeedSupportTDE,omitnil,omitempty" name:"NeedSupportTDE"`

	// KeyId of custom key, which is required if you select custom key encryption. It is also the unique CMK identifier.
	// For more information on creating `KeyId`, see [Enabling TDE](https://www.tencentcloud.com/document/product/409/47762).
	KMSKeyId *string `json:"KMSKeyId,omitnil,omitempty" name:"KMSKeyId"`

	// The region where the KMS service is enabled. When `KMSRegion` is left empty, the current region will be selected by default.  If the current region does not support KMS, you must select another region that does.
	// For more information on `KMSRegion`, see [Enabling TDE](https://intl.cloud.tencent.com/document/product/409/71749?from_cn_redirect=1).
	KMSRegion *string `json:"KMSRegion,omitnil,omitempty" name:"KMSRegion"`

	// Database engines. Valid values:
	// <li>`postgresql`: TencentDB for PostgreSQL
	// <li>`mssql_compatible`: MSSQL compatible-TencentDB for PostgreSQL
	// Default value: `postgresql`.
	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`

	// Configuration information of database engine in the following format:
	// {"$key1":"$value1", "$key2":"$value2"}
	// Valid values:
	// mssql_compatible engine:
	// <li>`migrationMode`: Database mode. Valid values: `single-db` (single-database mode), `multi-db` (multi-database mode). Default value: `single-db`.
	// <li>`defaultLocale`: Default locale, which can’t be modified after the initialization. Default value: `en_US`. Valid values:
	// "af_ZA", "sq_AL", "ar_DZ", "ar_BH", "ar_EG", "ar_IQ", "ar_JO", "ar_KW", "ar_LB", "ar_LY", "ar_MA", "ar_OM", "ar_QA", "ar_SA", "ar_SY", "ar_TN", "ar_AE", "ar_YE", "hy_AM", "az_Cyrl_AZ", "az_Latn_AZ", "eu_ES", "be_BY", "bg_BG", "ca_ES", "zh_HK", "zh_MO", "zh_CN", "zh_SG", "zh_TW", "hr_HR", "cs_CZ", "da_DK", "nl_BE", "nl_NL", "en_AU", "en_BZ", "en_CA", "en_IE", "en_JM", "en_NZ", "en_PH", "en_ZA", "en_TT", "en_GB", "en_US", "en_ZW", "et_EE", "fo_FO", "fa_IR", "fi_FI", "fr_BE", "fr_CA", "fr_FR", "fr_LU", "fr_MC", "fr_CH", "mk_MK", "ka_GE", "de_AT", "de_DE", "de_LI", "de_LU", "de_CH", "el_GR", "gu_IN", "he_IL", "hi_IN", "hu_HU", "is_IS", "id_ID", "it_IT", "it_CH", "ja_JP", "kn_IN", "kok_IN", "ko_KR", "ky_KG", "lv_LV", "lt_LT", "ms_BN", "ms_MY", "mr_IN", "mn_MN", "nb_NO", "nn_NO", "pl_PL", "pt_BR", "pt_PT", "pa_IN", "ro_RO", "ru_RU", "sa_IN", "sr_Cyrl_RS", "sr_Latn_RS", "sk_SK", "sl_SI", "es_AR", "es_BO", "es_CL", "es_CO", "es_CR", "es_DO", "es_EC", "es_SV", "es_GT", "es_HN", "es_MX", "es_NI", "es_PA", "es_PY","es_PE", "es_PR", "es_ES", "es_TRADITIONAL", "es_UY", "es_VE", "sw_KE", "sv_FI", "sv_SE", "tt_RU", "te_IN", "th_TH", "tr_TR", "uk_UA", "ur_IN", "ur_PK", "uz_Cyrl_UZ", "uz_Latn_UZ", "vi_VN".
	// <li>`serverCollationName`: Default collation name, which can’t be modified after the initialization. Default value: "bbf_unicode_general_ci_as". Valid values: "bbf_unicode_cp1_ci_as", "bbf_unicode_CP1250_ci_as", "bbf_unicode_CP1251_ci_as", "bbf_unicode_cp1253_ci_as", "bbf_unicode_cp1254_ci_as", "bbf_unicode_cp1255_ci_as", "bbf_unicode_cp1256_ci_as", "bbf_unicode_cp1257_ci_as", "bbf_unicode_cp1258_ci_as", "bbf_unicode_cp874_ci_as", "sql_latin1_general_cp1250_ci_as", "sql_latin1_general_cp1251_ci_as", "sql_latin1_general_cp1_ci_as", "sql_latin1_general_cp1253_ci_as", "sql_latin1_general_cp1254_ci_as", "sql_latin1_general_cp1255_ci_as","sql_latin1_general_cp1256_ci_as", "sql_latin1_general_cp1257_ci_as", "sql_latin1_general_cp1258_ci_as", "chinese_prc_ci_as", "cyrillic_general_ci_as", "finnish_swedish_ci_as", "french_ci_as", "japanese_ci_as", "korean_wansung_ci_as", "latin1_general_ci_as", "modern_spanish_ci_as", "polish_ci_as", "thai_ci_as", "traditional_spanish_ci_as", "turkish_ci_as", "ukrainian_ci_as", "vietnamese_ci_as"。
	DBEngineConfig *string `json:"DBEngineConfig,omitnil,omitempty" name:"DBEngineConfig"`

	// Primary-standby sync mode. Valid values:  
	// <li>`Semi-sync`
	// <li>`Async`
	// Default value for the primary instance: `Semi-sync`.
	// Default value for the standby instance: `Async`.
	SyncMode *string `json:"SyncMode,omitnil,omitempty" name:"SyncMode"`

	// Whether IPv6 is supported.
	// <li>`0`: No.
	// <li>`1`: Yes.
	// Default value: `0`.
	NeedSupportIpv6 *uint64 `json:"NeedSupportIpv6,omitnil,omitempty" name:"NeedSupportIpv6"`
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
	delete(f, "Zone")
	delete(f, "SpecCode")
	delete(f, "Storage")
	delete(f, "InstanceCount")
	delete(f, "Period")
	delete(f, "Charset")
	delete(f, "AdminName")
	delete(f, "AdminPassword")
	delete(f, "DBMajorVersion")
	delete(f, "DBVersion")
	delete(f, "DBKernelVersion")
	delete(f, "InstanceChargeType")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "DBNodeSet")
	delete(f, "AutoRenewFlag")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "ProjectId")
	delete(f, "ActivityId")
	delete(f, "Name")
	delete(f, "TagList")
	delete(f, "SecurityGroupIds")
	delete(f, "NeedSupportTDE")
	delete(f, "KMSKeyId")
	delete(f, "KMSRegion")
	delete(f, "DBEngine")
	delete(f, "DBEngineConfig")
	delete(f, "SyncMode")
	delete(f, "NeedSupportIpv6")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstancesResponseParams struct {
	// Order number list. Each instance corresponds to an order number.
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// Bill ID of frozen fees
	BillId *string `json:"BillId,omitnil,omitempty" name:"BillId"`

	// ID set of instances which have been created successfully. The parameter value will be returned only when the pay-as-you-go billing mode is used.
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitnil,omitempty" name:"DBInstanceIdSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateParameterTemplateRequestParams struct {
	// Template name, which can contain 1-60 letters, digits, and symbols (-_./()[]()+=:@).
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// The major database version number, such as 11, 12, 13.
	DBMajorVersion *string `json:"DBMajorVersion,omitnil,omitempty" name:"DBMajorVersion"`

	// Database engine, such as postgresql, mssql_compatible.
	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`

	// Parameter template description, which can contain 1-60 letters, digits, and symbols (-_./()[]()+=:@).
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`
}

type CreateParameterTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Template name, which can contain 1-60 letters, digits, and symbols (-_./()[]()+=:@).
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// The major database version number, such as 11, 12, 13.
	DBMajorVersion *string `json:"DBMajorVersion,omitnil,omitempty" name:"DBMajorVersion"`

	// Database engine, such as postgresql, mssql_compatible.
	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`

	// Parameter template description, which can contain 1-60 letters, digits, and symbols (-_./()[]()+=:@).
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`
}

func (r *CreateParameterTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateParameterTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateName")
	delete(f, "DBMajorVersion")
	delete(f, "DBEngine")
	delete(f, "TemplateDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateParameterTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateParameterTemplateResponseParams struct {
	// Parameter template ID, which uniquely identifies a parameter template.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateParameterTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateParameterTemplateResponseParams `json:"Response"`
}

func (r *CreateParameterTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateParameterTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReadOnlyDBInstanceRequestParams struct {
	// Primary AZ of an instance, such as "ap-guangzhou-3".
	// The information of AZ can be obtained from the `Zone` field in the return value of the [DescribeZones](https://intl.cloud.tencent.com/document/api/409/16769?from_cn_redirect=1) API.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// ID of the primary instance to which the read-only instance belongs
	MasterDBInstanceId *string `json:"MasterDBInstanceId,omitnil,omitempty" name:"MasterDBInstanceId"`

	// Purchasable code, which can be obtained from the `SpecCode` field in the return value of the [DescribeClasses](https://intl.cloud.tencent.com/document/api/409/89019?from_cn_redirect=1) API.
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// Instance storage capacity in GB
	Storage *uint64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// The number of instances to be purchased at a time. Value range: 1-10. To purchase more than 10 instances each time, you can make multiple calls.
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// Validity period in months, valid values:
	// <li>Monthly subscription: `1`, `2`, `3`, 4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	// <li>Pay-as-you-go: `1`.
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// VPC ID in the format of `vpc-xxxxxxx`, which can be obtained in the console or from the `unVpcId` field in the return value of the [DescribeVpcEx](https://intl.cloud.tencent.com/document/api/215/1372?from_cn_redirect=1) API.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// VPC subnet ID in the format of `subnet-xxxxxxxx` which can be obtained in the console or from the `unSubnetId` field in the return value of the [DescribeSubnets](https://intl.cloud.tencent.com/document/api/215/15784?from_cn_redirect=1) API.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Instance billing mode. Valid values: 
	// <li>`PREPAID`: Monthly subscription
	// <li>`POSTPAID_BY_HOUR`: Pay-as-you-go
	// Default value: `PREPAID`. If the primary instance is pay-as-you-go, so is the read-only instance.
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Whether to use vouchers automatically. Valid values:
	// <li>`0`: No.
	// <li>`1`: Yes.
	// Default value: `0`.
	AutoVoucher *uint64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Voucher ID list. Currently, you can specify only one voucher.
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// Auto-renewal flag. Valid values:
	// <li>`0`: Manual renewal.
	// <li>`1`: Automatic renewal.
	// Default value: `0`.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Special offer ID
	ActivityId *int64 `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// RO group ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`

	// The information of tags to be bound with the instance, which is left empty by default. This parameter can be obtained from the `Tags` field in the return value of the [DescribeTags](https://intl.cloud.tencent.com/document/api/651/35316?from_cn_redirect=1) API.
	TagList *Tag `json:"TagList,omitnil,omitempty" name:"TagList"`

	// Security group of the instance, which can be obtained from the `sgld` field in the return value of the [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/api/215/15808?from_cn_redirect=1) API. If this parameter is not specified, the default security group will be bound.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Whether IPv6 is supported.
	// <li>`0`: No.
	// <li>`1`: Yes.
	// Default value: `0`.
	NeedSupportIpv6 *uint64 `json:"NeedSupportIpv6,omitnil,omitempty" name:"NeedSupportIpv6"`

	// Instance name (which will be supported in the future)
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// (Disused) You don’t need to specify a version, as the kernel version is as the same as that of the instance.
	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`
}

type CreateReadOnlyDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Primary AZ of an instance, such as "ap-guangzhou-3".
	// The information of AZ can be obtained from the `Zone` field in the return value of the [DescribeZones](https://intl.cloud.tencent.com/document/api/409/16769?from_cn_redirect=1) API.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// ID of the primary instance to which the read-only instance belongs
	MasterDBInstanceId *string `json:"MasterDBInstanceId,omitnil,omitempty" name:"MasterDBInstanceId"`

	// Purchasable code, which can be obtained from the `SpecCode` field in the return value of the [DescribeClasses](https://intl.cloud.tencent.com/document/api/409/89019?from_cn_redirect=1) API.
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// Instance storage capacity in GB
	Storage *uint64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// The number of instances to be purchased at a time. Value range: 1-10. To purchase more than 10 instances each time, you can make multiple calls.
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// Validity period in months, valid values:
	// <li>Monthly subscription: `1`, `2`, `3`, 4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	// <li>Pay-as-you-go: `1`.
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// VPC ID in the format of `vpc-xxxxxxx`, which can be obtained in the console or from the `unVpcId` field in the return value of the [DescribeVpcEx](https://intl.cloud.tencent.com/document/api/215/1372?from_cn_redirect=1) API.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// VPC subnet ID in the format of `subnet-xxxxxxxx` which can be obtained in the console or from the `unSubnetId` field in the return value of the [DescribeSubnets](https://intl.cloud.tencent.com/document/api/215/15784?from_cn_redirect=1) API.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Instance billing mode. Valid values: 
	// <li>`PREPAID`: Monthly subscription
	// <li>`POSTPAID_BY_HOUR`: Pay-as-you-go
	// Default value: `PREPAID`. If the primary instance is pay-as-you-go, so is the read-only instance.
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Whether to use vouchers automatically. Valid values:
	// <li>`0`: No.
	// <li>`1`: Yes.
	// Default value: `0`.
	AutoVoucher *uint64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Voucher ID list. Currently, you can specify only one voucher.
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// Auto-renewal flag. Valid values:
	// <li>`0`: Manual renewal.
	// <li>`1`: Automatic renewal.
	// Default value: `0`.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Special offer ID
	ActivityId *int64 `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// RO group ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`

	// The information of tags to be bound with the instance, which is left empty by default. This parameter can be obtained from the `Tags` field in the return value of the [DescribeTags](https://intl.cloud.tencent.com/document/api/651/35316?from_cn_redirect=1) API.
	TagList *Tag `json:"TagList,omitnil,omitempty" name:"TagList"`

	// Security group of the instance, which can be obtained from the `sgld` field in the return value of the [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/api/215/15808?from_cn_redirect=1) API. If this parameter is not specified, the default security group will be bound.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// Whether IPv6 is supported.
	// <li>`0`: No.
	// <li>`1`: Yes.
	// Default value: `0`.
	NeedSupportIpv6 *uint64 `json:"NeedSupportIpv6,omitnil,omitempty" name:"NeedSupportIpv6"`

	// Instance name (which will be supported in the future)
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// (Disused) You don’t need to specify a version, as the kernel version is as the same as that of the instance.
	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`
}

func (r *CreateReadOnlyDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReadOnlyDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "MasterDBInstanceId")
	delete(f, "SpecCode")
	delete(f, "Storage")
	delete(f, "InstanceCount")
	delete(f, "Period")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "InstanceChargeType")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "AutoRenewFlag")
	delete(f, "ProjectId")
	delete(f, "ActivityId")
	delete(f, "ReadOnlyGroupId")
	delete(f, "TagList")
	delete(f, "SecurityGroupIds")
	delete(f, "NeedSupportIpv6")
	delete(f, "Name")
	delete(f, "DBVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReadOnlyDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReadOnlyDBInstanceResponseParams struct {
	// Order number list. Each instance corresponds to an order number.
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// Bill ID of frozen fees
	BillId *string `json:"BillId,omitnil,omitempty" name:"BillId"`

	// ID set of instances which have been created successfully. The parameter value will be returned only when the pay-as-you-go billing mode is used.
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitnil,omitempty" name:"DBInstanceIdSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateReadOnlyDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateReadOnlyDBInstanceResponseParams `json:"Response"`
}

func (r *CreateReadOnlyDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReadOnlyDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReadOnlyGroupNetworkAccessRequestParams struct {
	// RO group ID in the format of pgro-4t9c6g7k.
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`

	// Unified VPC ID.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Whether to manually assign the VIP. Valid values: `true` (manually assign), `false` (automatically assign).
	IsAssignVip *bool `json:"IsAssignVip,omitnil,omitempty" name:"IsAssignVip"`

	// Target VIP.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`
}

type CreateReadOnlyGroupNetworkAccessRequest struct {
	*tchttp.BaseRequest
	
	// RO group ID in the format of pgro-4t9c6g7k.
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`

	// Unified VPC ID.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Whether to manually assign the VIP. Valid values: `true` (manually assign), `false` (automatically assign).
	IsAssignVip *bool `json:"IsAssignVip,omitnil,omitempty" name:"IsAssignVip"`

	// Target VIP.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`
}

func (r *CreateReadOnlyGroupNetworkAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReadOnlyGroupNetworkAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReadOnlyGroupId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "IsAssignVip")
	delete(f, "Vip")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReadOnlyGroupNetworkAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReadOnlyGroupNetworkAccessResponseParams struct {
	// Task ID.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateReadOnlyGroupNetworkAccessResponse struct {
	*tchttp.BaseResponse
	Response *CreateReadOnlyGroupNetworkAccessResponseParams `json:"Response"`
}

func (r *CreateReadOnlyGroupNetworkAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReadOnlyGroupNetworkAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReadOnlyGroupRequestParams struct {
	// Primary instance ID
	MasterDBInstanceId *string `json:"MasterDBInstanceId,omitnil,omitempty" name:"MasterDBInstanceId"`

	// RO group name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: `0` (no), `1` (yes).
	ReplayLagEliminate *uint64 `json:"ReplayLagEliminate,omitnil,omitempty" name:"ReplayLagEliminate"`

	// Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: `0` (no), `1` (yes).
	ReplayLatencyEliminate *uint64 `json:"ReplayLatencyEliminate,omitnil,omitempty" name:"ReplayLatencyEliminate"`

	// Delay threshold in ms
	MaxReplayLag *uint64 `json:"MaxReplayLag,omitnil,omitempty" name:"MaxReplayLag"`

	// Delayed log size threshold in MB
	MaxReplayLatency *uint64 `json:"MaxReplayLatency,omitnil,omitempty" name:"MaxReplayLatency"`

	// The minimum number of read-only replicas that must be retained in an RO group
	MinDelayEliminateReserve *uint64 `json:"MinDelayEliminateReserve,omitnil,omitempty" name:"MinDelayEliminateReserve"`

	// Security group ID
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

type CreateReadOnlyGroupRequest struct {
	*tchttp.BaseRequest
	
	// Primary instance ID
	MasterDBInstanceId *string `json:"MasterDBInstanceId,omitnil,omitempty" name:"MasterDBInstanceId"`

	// RO group name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: `0` (no), `1` (yes).
	ReplayLagEliminate *uint64 `json:"ReplayLagEliminate,omitnil,omitempty" name:"ReplayLagEliminate"`

	// Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: `0` (no), `1` (yes).
	ReplayLatencyEliminate *uint64 `json:"ReplayLatencyEliminate,omitnil,omitempty" name:"ReplayLatencyEliminate"`

	// Delay threshold in ms
	MaxReplayLag *uint64 `json:"MaxReplayLag,omitnil,omitempty" name:"MaxReplayLag"`

	// Delayed log size threshold in MB
	MaxReplayLatency *uint64 `json:"MaxReplayLatency,omitnil,omitempty" name:"MaxReplayLatency"`

	// The minimum number of read-only replicas that must be retained in an RO group
	MinDelayEliminateReserve *uint64 `json:"MinDelayEliminateReserve,omitnil,omitempty" name:"MinDelayEliminateReserve"`

	// Security group ID
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

func (r *CreateReadOnlyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReadOnlyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MasterDBInstanceId")
	delete(f, "Name")
	delete(f, "ProjectId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "ReplayLagEliminate")
	delete(f, "ReplayLatencyEliminate")
	delete(f, "MaxReplayLag")
	delete(f, "MaxReplayLatency")
	delete(f, "MinDelayEliminateReserve")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReadOnlyGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReadOnlyGroupResponseParams struct {
	// RO group ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`

	// Task ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateReadOnlyGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateReadOnlyGroupResponseParams `json:"Response"`
}

func (r *CreateReadOnlyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReadOnlyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServerlessDBInstanceRequestParams struct {
	// Availability zone ID. Only ap-shanghai-2, ap-beijing-1, and ap-guangzhou-2 are supported during the beta test.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Instance name. The value must be unique for the same account.
	DBInstanceName *string `json:"DBInstanceName,omitnil,omitempty" name:"DBInstanceName"`

	// Kernel version of a PostgreSQL instance. Currently, only 10.4 is supported.
	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`

	// Database character set of a PostgreSQL instance. Currently, only UTF-8 is supported.
	DBCharset *string `json:"DBCharset,omitnil,omitempty" name:"DBCharset"`

	// Project ID.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// VPC ID.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// VPC subnet ID.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Array of tags to be bound with the instance
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`
}

type CreateServerlessDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Availability zone ID. Only ap-shanghai-2, ap-beijing-1, and ap-guangzhou-2 are supported during the beta test.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Instance name. The value must be unique for the same account.
	DBInstanceName *string `json:"DBInstanceName,omitnil,omitempty" name:"DBInstanceName"`

	// Kernel version of a PostgreSQL instance. Currently, only 10.4 is supported.
	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`

	// Database character set of a PostgreSQL instance. Currently, only UTF-8 is supported.
	DBCharset *string `json:"DBCharset,omitnil,omitempty" name:"DBCharset"`

	// Project ID.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// VPC ID.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// VPC subnet ID.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Array of tags to be bound with the instance
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`
}

func (r *CreateServerlessDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServerlessDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "DBInstanceName")
	delete(f, "DBVersion")
	delete(f, "DBCharset")
	delete(f, "ProjectId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "TagList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateServerlessDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServerlessDBInstanceResponseParams struct {
	// Instance ID, such as "postgres-xxxxx". The value must be globally unique.
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateServerlessDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateServerlessDBInstanceResponseParams `json:"Response"`
}

func (r *CreateServerlessDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServerlessDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DBBackup struct {
	// Unique backup file ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// File generation start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// File generation end time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// File size in KB
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// Policy (0: instance backup, 1: multi-database backup)
	Strategy *int64 `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// Type (0: scheduled)
	Way *int64 `json:"Way,omitnil,omitempty" name:"Way"`

	// Backup mode (1: full)
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Status (1: creating, 2: success, 3: failure)
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// DB list
	DbList []*string `json:"DbList,omitnil,omitempty" name:"DbList"`

	// Download address on private network
	InternalAddr *string `json:"InternalAddr,omitnil,omitempty" name:"InternalAddr"`

	// Download address on public network
	ExternalAddr *string `json:"ExternalAddr,omitnil,omitempty" name:"ExternalAddr"`

	// Backup set ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SetId *string `json:"SetId,omitnil,omitempty" name:"SetId"`
}

type DBInstance struct {
	// Instance region such as ap-guangzhou, which corresponds to the`Region` field in `RegionSet`.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Instance AZ such as ap-guangzhou-3, which corresponds to the `Zone` field of `ZoneSet`.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// VPC ID in the format of `vpc-xxxxxxx`, which can be obtained in the console or from the `unVpcId` field in the return value of the [DescribeVpcEx](https://intl.cloud.tencent.com/document/api/215/1372?from_cn_redirect=1) API.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// VPC subnet ID in the format of `subnet-xxxxxxxx`, which can be obtained in the console or from the `unSubnetId` field in the return value of the [DescribeSubnets ](https://intl.cloud.tencent.com/document/api/215/15784?from_cn_redirect=1) API.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Instance name
	DBInstanceName *string `json:"DBInstanceName,omitnil,omitempty" name:"DBInstanceName"`

	// Instance status. Valid values: `applying`, `init` (to be initialized), `initing` (initializing), `running`, `limited run`, `isolating`, `isolated`, `recycling`, `recycled`, `job running`, `offline`, `migrating`, `expanding`, `waitSwitch` (waiting for switch), `switching`, `readonly`, `restarting`, `network changing`, `upgrading` (upgrading kernel version), `audit-switching` (changing audit status), `primary-switching` (primary-standby switching).
	DBInstanceStatus *string `json:"DBInstanceStatus,omitnil,omitempty" name:"DBInstanceStatus"`

	// Assigned instance memory size in GB
	DBInstanceMemory *uint64 `json:"DBInstanceMemory,omitnil,omitempty" name:"DBInstanceMemory"`

	// Assigned instance storage capacity in GB
	DBInstanceStorage *uint64 `json:"DBInstanceStorage,omitnil,omitempty" name:"DBInstanceStorage"`

	// Number of assigned CPUs
	DBInstanceCpu *uint64 `json:"DBInstanceCpu,omitnil,omitempty" name:"DBInstanceCpu"`

	// Purchasable specification ID
	DBInstanceClass *string `json:"DBInstanceClass,omitnil,omitempty" name:"DBInstanceClass"`

	// The major PostgreSQL version number, which can be queried by the [DescribeDBVersions](https://intl.cloud.tencent.com/document/api/409/89018?from_cn_redirect=1) API. Valid values: `10`, `11`, `12`, `13`, `14`, `15`.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	DBMajorVersion *string `json:"DBMajorVersion,omitnil,omitempty" name:"DBMajorVersion"`

	// Number of the major PostgreSQL community version and minor version, such as 12.4, which can be queried by the [DescribeDBVersions](https://intl.cloud.tencent.com/document/api/409/89018?from_cn_redirect=1) API.
	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`

	// PostgreSQL kernel version number (like v12.7_r1.8), which can be queried by the [DescribeDBVersions](https://intl.cloud.tencent.com/document/api/409/89018?from_cn_redirect=1) API.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	DBKernelVersion *string `json:"DBKernelVersion,omitnil,omitempty" name:"DBKernelVersion"`

	// Instance type. Valid values:
	// <li>`primary`: Primary instance
	// <li>`readonly`: Read-only instance
	// <li>`guard`: Disaster recovery instance
	// <li>`temp`: Temp instance
	DBInstanceType *string `json:"DBInstanceType,omitnil,omitempty" name:"DBInstanceType"`

	// Instance version. Valid value: `standard` (dual-server high-availability; one-primary-one-standby).
	DBInstanceVersion *string `json:"DBInstanceVersion,omitnil,omitempty" name:"DBInstanceVersion"`

	// Instance character set. Valid values:
	// <li>`UTF8`
	// <li>`LATIN1`
	DBCharset *string `json:"DBCharset,omitnil,omitempty" name:"DBCharset"`

	// Instance creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Last updated time of the instance attribute
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Instance expiration time
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// Instance isolation time
	IsolatedTime *string `json:"IsolatedTime,omitnil,omitempty" name:"IsolatedTime"`

	// Billing mode. Valid values:
	// <li>`PREPAID`: Monthly subscription
	// <li>`postpaid`: Pay-as-you-go
	PayType *string `json:"PayType,omitnil,omitempty" name:"PayType"`

	// Whether auto-renewal is enabled. Valid values:
	// <li>`0`: Manual renewal.
	// <li>`1`: Automatic renewal.
	// Default value: `0`.
	AutoRenew *uint64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// Instance network connection information
	DBInstanceNetInfo []*DBInstanceNetInfo `json:"DBInstanceNetInfo,omitnil,omitempty" name:"DBInstanceNetInfo"`

	// Machine type
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// User `AppId`
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Instance `Uid`
	Uid *uint64 `json:"Uid,omitnil,omitempty" name:"Uid"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// The information of tags associated with instances
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`

	// Primary instance information, which is returned only when the instance is read-only.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	MasterDBInstanceId *string `json:"MasterDBInstanceId,omitnil,omitempty" name:"MasterDBInstanceId"`

	// Number of read-only instances
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	ReadOnlyInstanceNum *int64 `json:"ReadOnlyInstanceNum,omitnil,omitempty" name:"ReadOnlyInstanceNum"`

	// The status of a read-only instance in a read-only group
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	StatusInReadonlyGroup *string `json:"StatusInReadonlyGroup,omitnil,omitempty" name:"StatusInReadonlyGroup"`

	// Offline time
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	OfflineTime *string `json:"OfflineTime,omitnil,omitempty" name:"OfflineTime"`

	// Instance node information
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	DBNodeSet []*DBNode `json:"DBNodeSet,omitnil,omitempty" name:"DBNodeSet"`

	// Whether the instance supports TDE. Valid values: 
	// <li>`0`: No.
	// <li>`1`: Yes.
	// Default value: `0`.
	// For more information, see [TDE](https://intl.cloud.tencent.com/document/product/409/71748?from_cn_redirect=1).
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	IsSupportTDE *int64 `json:"IsSupportTDE,omitnil,omitempty" name:"IsSupportTDE"`

	// Database engines. Valid values:
	// <li>`postgresql`: TencentDB for PostgreSQL
	// <li>`mssql_compatible`: MSSQL compatible-TencentDB for PostgreSQL
	// Default value: `postgresql`.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`

	// Configuration information of database engine in the following format:
	// {"$key1":"$value1", "$key2":"$value2"}
	// Valid values:
	// mssql_compatible engine:
	// <li>`migrationMode`: Database mode. Valid values: `single-db` (single-database mode), `multi-db` (multi-database mode). Default value: `single-db`.
	// <li>`defaultLocale`: Default locale, which can’t be modified after the initialization. Default value: `en_US`. Valid values:
	// "af_ZA", "sq_AL", "ar_DZ", "ar_BH", "ar_EG", "ar_IQ", "ar_JO", "ar_KW", "ar_LB", "ar_LY", "ar_MA", "ar_OM", "ar_QA", "ar_SA", "ar_SY", "ar_TN", "ar_AE", "ar_YE", "hy_AM", "az_Cyrl_AZ", "az_Latn_AZ", "eu_ES", "be_BY", "bg_BG", "ca_ES", "zh_HK", "zh_MO", "zh_CN", "zh_SG", "zh_TW", "hr_HR", "cs_CZ", "da_DK", "nl_BE", "nl_NL", "en_AU", "en_BZ", "en_CA", "en_IE", "en_JM", "en_NZ", "en_PH", "en_ZA", "en_TT", "en_GB", "en_US", "en_ZW", "et_EE", "fo_FO", "fa_IR", "fi_FI", "fr_BE", "fr_CA", "fr_FR", "fr_LU", "fr_MC", "fr_CH", "mk_MK", "ka_GE", "de_AT", "de_DE", "de_LI", "de_LU", "de_CH", "el_GR", "gu_IN", "he_IL", "hi_IN", "hu_HU", "is_IS", "id_ID", "it_IT", "it_CH", "ja_JP", "kn_IN", "kok_IN", "ko_KR", "ky_KG", "lv_LV", "lt_LT", "ms_BN", "ms_MY", "mr_IN", "mn_MN", "nb_NO", "nn_NO", "pl_PL", "pt_BR", "pt_PT", "pa_IN", "ro_RO", "ru_RU", "sa_IN", "sr_Cyrl_RS", "sr_Latn_RS", "sk_SK", "sl_SI", "es_AR", "es_BO", "es_CL", "es_CO", "es_CR", "es_DO", "es_EC", "es_SV", "es_GT", "es_HN", "es_MX", "es_NI", "es_PA", "es_PY","es_PE", "es_PR", "es_ES", "es_TRADITIONAL", "es_UY", "es_VE", "sw_KE", "sv_FI", "sv_SE", "tt_RU", "te_IN", "th_TH", "tr_TR", "uk_UA", "ur_IN", "ur_PK", "uz_Cyrl_UZ", "uz_Latn_UZ", "vi_VN".
	// <li>`serverCollationName`: Default collation name, which can’t be modified after the initialization. Default value: "sql_latin1_general_cp1_ci_as". Valid values: "bbf_unicode_cp1_ci_as", "bbf_unicode_CP1250_ci_as", "bbf_unicode_CP1251_ci_as", "bbf_unicode_cp1253_ci_as", "bbf_unicode_cp1254_ci_as", "bbf_unicode_cp1255_ci_as", "bbf_unicode_cp1256_ci_as", "bbf_unicode_cp1257_ci_as", "bbf_unicode_cp1258_ci_as", "bbf_unicode_cp874_ci_as", "sql_latin1_general_cp1250_ci_as", "sql_latin1_general_cp1251_ci_as", "sql_latin1_general_cp1_ci_as", "sql_latin1_general_cp1253_ci_as", "sql_latin1_general_cp1254_ci_as", "sql_latin1_general_cp1255_ci_as","sql_latin1_general_cp1256_ci_as", "sql_latin1_general_cp1257_ci_as", "sql_latin1_general_cp1258_ci_as", "chinese_prc_ci_as", "cyrillic_general_ci_as", "finnish_swedish_ci_as", "french_ci_as", "japanese_ci_as", "korean_wansung_ci_as", "latin1_general_ci_as", "modern_spanish_ci_as", "polish_ci_as", "thai_ci_as", "traditional_spanish_ci_as", "turkish_ci_as", "ukrainian_ci_as", "vietnamese_ci_as".
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	DBEngineConfig *string `json:"DBEngineConfig,omitnil,omitempty" name:"DBEngineConfig"`

	// Network access list of the instance (this field has been deprecated)
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	NetworkAccessList []*NetworkAccess `json:"NetworkAccessList,omitnil,omitempty" name:"NetworkAccessList"`

	// Whether the instance supports IPv6. Valid values:
	// <li>`0`: No.
	// <li>`1`: Yes.
	// Default value: `0`.
	SupportIpv6 *uint64 `json:"SupportIpv6,omitnil,omitempty" name:"SupportIpv6"`
}

type DBInstanceNetInfo struct {
	// DNS domain name
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// Ip
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Connection port address
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Network type. 1: inner (private network address), 2: public (public network address)
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// Network connection status. Valid values: `initing` (never enabled before), `opened` (enabled), `closed` (disabled), `opening` (enabling), `closing` (disabling)
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// VPC ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Database connection protocol type. Valid values: `postgresql`, `mssql` (MSSQL-compatible)
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProtocolType *string `json:"ProtocolType,omitnil,omitempty" name:"ProtocolType"`
}

type DBNode struct {
	// Node type. Valid values:
	// `Primary`;
	// `Standby`.
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// AZ where the node resides, such as ap-guangzhou-1.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

// Predefined struct for user
type DeleteBaseBackupRequestParams struct {
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Base backup ID
	BaseBackupId *string `json:"BaseBackupId,omitnil,omitempty" name:"BaseBackupId"`
}

type DeleteBaseBackupRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Base backup ID
	BaseBackupId *string `json:"BaseBackupId,omitnil,omitempty" name:"BaseBackupId"`
}

func (r *DeleteBaseBackupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBaseBackupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "BaseBackupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBaseBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBaseBackupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteBaseBackupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBaseBackupResponseParams `json:"Response"`
}

func (r *DeleteBaseBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBaseBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDBInstanceNetworkAccessRequestParams struct {
	// Instance ID in the format of postgres-6bwgamo3.
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Unified VPC ID. If you want to delete the classic network, set the parameter to `0`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID. If you want to delete the classic network, set the parameter to `0`.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Target VIP.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`
}

type DeleteDBInstanceNetworkAccessRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of postgres-6bwgamo3.
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Unified VPC ID. If you want to delete the classic network, set the parameter to `0`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID. If you want to delete the classic network, set the parameter to `0`.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Target VIP.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`
}

func (r *DeleteDBInstanceNetworkAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDBInstanceNetworkAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Vip")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDBInstanceNetworkAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDBInstanceNetworkAccessResponseParams struct {
	// Task ID.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDBInstanceNetworkAccessResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDBInstanceNetworkAccessResponseParams `json:"Response"`
}

func (r *DeleteDBInstanceNetworkAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDBInstanceNetworkAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLogBackupRequestParams struct {
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Log backup ID
	LogBackupId *string `json:"LogBackupId,omitnil,omitempty" name:"LogBackupId"`
}

type DeleteLogBackupRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Log backup ID
	LogBackupId *string `json:"LogBackupId,omitnil,omitempty" name:"LogBackupId"`
}

func (r *DeleteLogBackupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLogBackupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "LogBackupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLogBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLogBackupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLogBackupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLogBackupResponseParams `json:"Response"`
}

func (r *DeleteLogBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLogBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteParameterTemplateRequestParams struct {
	// Parameter template ID, which uniquely identifies the parameter template to be operated.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DeleteParameterTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Parameter template ID, which uniquely identifies the parameter template to be operated.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DeleteParameterTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteParameterTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteParameterTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteParameterTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteParameterTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteParameterTemplateResponseParams `json:"Response"`
}

func (r *DeleteParameterTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteParameterTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReadOnlyGroupNetworkAccessRequestParams struct {
	// RO group ID in the format of pgro-4t9c6g7k.
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`

	// Unified VPC ID. If you want to delete the classic network, set the parameter to `0`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID. If you want to delete the classic network, set the parameter to `0`.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Target VIP.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`
}

type DeleteReadOnlyGroupNetworkAccessRequest struct {
	*tchttp.BaseRequest
	
	// RO group ID in the format of pgro-4t9c6g7k.
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`

	// Unified VPC ID. If you want to delete the classic network, set the parameter to `0`.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID. If you want to delete the classic network, set the parameter to `0`.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Target VIP.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`
}

func (r *DeleteReadOnlyGroupNetworkAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReadOnlyGroupNetworkAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReadOnlyGroupId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Vip")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteReadOnlyGroupNetworkAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReadOnlyGroupNetworkAccessResponseParams struct {
	// Task ID.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteReadOnlyGroupNetworkAccessResponse struct {
	*tchttp.BaseResponse
	Response *DeleteReadOnlyGroupNetworkAccessResponseParams `json:"Response"`
}

func (r *DeleteReadOnlyGroupNetworkAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReadOnlyGroupNetworkAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReadOnlyGroupRequestParams struct {
	// ID of the RO group to be deleted
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`
}

type DeleteReadOnlyGroupRequest struct {
	*tchttp.BaseRequest
	
	// ID of the RO group to be deleted
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`
}

func (r *DeleteReadOnlyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReadOnlyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReadOnlyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteReadOnlyGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReadOnlyGroupResponseParams struct {
	// Task ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteReadOnlyGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteReadOnlyGroupResponseParams `json:"Response"`
}

func (r *DeleteReadOnlyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReadOnlyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServerlessDBInstanceRequestParams struct {
	// Instance name. Either instance name or instance ID (or both) must be passed in. If both are passed in, the instance ID will prevail.
	DBInstanceName *string `json:"DBInstanceName,omitnil,omitempty" name:"DBInstanceName"`

	// Instance ID. Either instance name or instance ID (or both) must be passed in. If both are passed in, the instance ID will prevail.
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

type DeleteServerlessDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance name. Either instance name or instance ID (or both) must be passed in. If both are passed in, the instance ID will prevail.
	DBInstanceName *string `json:"DBInstanceName,omitnil,omitempty" name:"DBInstanceName"`

	// Instance ID. Either instance name or instance ID (or both) must be passed in. If both are passed in, the instance ID will prevail.
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

func (r *DeleteServerlessDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServerlessDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceName")
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteServerlessDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServerlessDBInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteServerlessDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteServerlessDBInstanceResponseParams `json:"Response"`
}

func (r *DeleteServerlessDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServerlessDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsRequestParams struct {
	// Instance ID in the format of postgres-6fego161
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Number of entries returned per page. Default value: 10. Value range: 1–100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Data offset, which starts from 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Whether to sort by creation time or username. Valid values: `createTime` (sort by creation time), `name` (sort by username)
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Whether returns are sorted in ascending or descending order. Valid values: `desc` (descending), `asc` (ascending)
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of postgres-6fego161
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Number of entries returned per page. Default value: 10. Value range: 1–100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Data offset, which starts from 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Whether to sort by creation time or username. Valid values: `createTime` (sort by creation time), `name` (sort by username)
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Whether returns are sorted in ascending or descending order. Valid values: `desc` (descending), `asc` (ascending)
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
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
	delete(f, "DBInstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsResponseParams struct {
	// Number of date entries returned for this API call.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Account list details.
	Details []*AccountInfo `json:"Details,omitnil,omitempty" name:"Details"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeAvailableRecoveryTimeRequestParams struct {
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

type DescribeAvailableRecoveryTimeRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

func (r *DescribeAvailableRecoveryTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailableRecoveryTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAvailableRecoveryTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAvailableRecoveryTimeResponseParams struct {
	// The earliest restoration time (UTC+8).
	RecoveryBeginTime *string `json:"RecoveryBeginTime,omitnil,omitempty" name:"RecoveryBeginTime"`

	// The latest restoration time (UTC+8).
	RecoveryEndTime *string `json:"RecoveryEndTime,omitnil,omitempty" name:"RecoveryEndTime"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAvailableRecoveryTimeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAvailableRecoveryTimeResponseParams `json:"Response"`
}

func (r *DescribeAvailableRecoveryTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailableRecoveryTimeResponse) FromJsonString(s string) error {
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
	// Type of the network restrictions for downloading a backup file. Valid values: `NONE` (backups can be downloaded over both private and public networks), `INTRANET` (backups can only be downloaded over the private network), `CUSTOMIZE` (backups can be downloaded over specified VPCs or at specified IPs).
	RestrictionType *string `json:"RestrictionType,omitnil,omitempty" name:"RestrictionType"`

	// Whether VPC is allowed. Valid values: `ALLOW` (allow), `DENY` (deny). 
	// Note:  This field may return null, indicating that no valid values can be obtained.
	VpcRestrictionEffect *string `json:"VpcRestrictionEffect,omitnil,omitempty" name:"VpcRestrictionEffect"`

	// Whether it is allowed to download the VPC ID list of the backup files. 
	// Note:  This field may return null, indicating that no valid values can be obtained.
	VpcIdSet []*string `json:"VpcIdSet,omitnil,omitempty" name:"VpcIdSet"`

	// Whether IP is allowed. Valid values: `ALLOW` (allow), `DENY` (deny). 
	// Note: Note: This field may return null, indicating that no valid values can be obtained.
	IpRestrictionEffect *string `json:"IpRestrictionEffect,omitnil,omitempty" name:"IpRestrictionEffect"`

	// Whether it is allowed to download the IP list of the backup files. 
	// Note:  This field may return null, indicating that no valid values can be obtained.
	IpSet []*string `json:"IpSet,omitnil,omitempty" name:"IpSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeBackupDownloadURLRequestParams struct {
	// Instance ID.
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Backup type. Valid values: `LogBackup`, `BaseBackup`.
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// Unique backup ID.
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// Validity period of a URL, which is 12 hours by default.
	URLExpireTime *uint64 `json:"URLExpireTime,omitnil,omitempty" name:"URLExpireTime"`

	// Backup download restriction
	BackupDownloadRestriction *BackupDownloadRestriction `json:"BackupDownloadRestriction,omitnil,omitempty" name:"BackupDownloadRestriction"`
}

type DescribeBackupDownloadURLRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Backup type. Valid values: `LogBackup`, `BaseBackup`.
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// Unique backup ID.
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// Validity period of a URL, which is 12 hours by default.
	URLExpireTime *uint64 `json:"URLExpireTime,omitnil,omitempty" name:"URLExpireTime"`

	// Backup download restriction
	BackupDownloadRestriction *BackupDownloadRestriction `json:"BackupDownloadRestriction,omitnil,omitempty" name:"BackupDownloadRestriction"`
}

func (r *DescribeBackupDownloadURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "BackupType")
	delete(f, "BackupId")
	delete(f, "URLExpireTime")
	delete(f, "BackupDownloadRestriction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupDownloadURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDownloadURLResponseParams struct {
	// Backup download URL
	BackupDownloadURL *string `json:"BackupDownloadURL,omitnil,omitempty" name:"BackupDownloadURL"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupDownloadURLResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupDownloadURLResponseParams `json:"Response"`
}

func (r *DescribeBackupDownloadURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupOverviewRequestParams struct {

}

type DescribeBackupOverviewRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeBackupOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupOverviewResponseParams struct {
	// Total free space size in bytes
	TotalFreeSize *uint64 `json:"TotalFreeSize,omitnil,omitempty" name:"TotalFreeSize"`

	// Used free space size in bytes
	UsedFreeSize *uint64 `json:"UsedFreeSize,omitnil,omitempty" name:"UsedFreeSize"`

	// Used paid space size in bytes
	UsedBillingSize *uint64 `json:"UsedBillingSize,omitnil,omitempty" name:"UsedBillingSize"`

	// Number of log backups
	LogBackupCount *uint64 `json:"LogBackupCount,omitnil,omitempty" name:"LogBackupCount"`

	// Log backup size in bytes
	LogBackupSize *uint64 `json:"LogBackupSize,omitnil,omitempty" name:"LogBackupSize"`

	// Number of manually created full backups
	ManualBaseBackupCount *uint64 `json:"ManualBaseBackupCount,omitnil,omitempty" name:"ManualBaseBackupCount"`

	// Size of manually created full backups in bytes
	ManualBaseBackupSize *uint64 `json:"ManualBaseBackupSize,omitnil,omitempty" name:"ManualBaseBackupSize"`

	// Number of automatically created full backups
	AutoBaseBackupCount *uint64 `json:"AutoBaseBackupCount,omitnil,omitempty" name:"AutoBaseBackupCount"`

	// Size of automatically created full backups in bytes
	AutoBaseBackupSize *uint64 `json:"AutoBaseBackupSize,omitnil,omitempty" name:"AutoBaseBackupSize"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupOverviewResponseParams `json:"Response"`
}

func (r *DescribeBackupOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupPlansRequestParams struct {
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

type DescribeBackupPlansRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

func (r *DescribeBackupPlansRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupPlansRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupPlansRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupPlansResponseParams struct {
	// The set of instance backup plans
	Plans []*BackupPlan `json:"Plans,omitnil,omitempty" name:"Plans"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupPlansResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupPlansResponseParams `json:"Response"`
}

func (r *DescribeBackupPlansResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupPlansResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupSummariesRequestParams struct {
	// The maximum number of results returned per page. Value range: 1-100. Default: `10`
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Data offset, which starts from 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Filter instances using one or more criteria. Valid filter names:
	// db-instance-id: Filter by instance ID (in string format).
	// db-instance-name: Filter by instance name (in string format).
	// db-instance-ip: Filter by instance VPC IP (in string format).
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Sorting field. Valid values: `TotalBackupSize`, `LogBackupSize`, `ManualBaseBackupSize`, `AutoBaseBackupSize`.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `asc` (ascending), `desc` (descending).
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type DescribeBackupSummariesRequest struct {
	*tchttp.BaseRequest
	
	// The maximum number of results returned per page. Value range: 1-100. Default: `10`
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Data offset, which starts from 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Filter instances using one or more criteria. Valid filter names:
	// db-instance-id: Filter by instance ID (in string format).
	// db-instance-name: Filter by instance name (in string format).
	// db-instance-ip: Filter by instance VPC IP (in string format).
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Sorting field. Valid values: `TotalBackupSize`, `LogBackupSize`, `ManualBaseBackupSize`, `AutoBaseBackupSize`.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `asc` (ascending), `desc` (descending).
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

func (r *DescribeBackupSummariesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupSummariesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupSummariesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupSummariesResponseParams struct {
	// Backup statistics list.
	BackupSummarySet []*BackupSummary `json:"BackupSummarySet,omitnil,omitempty" name:"BackupSummarySet"`

	// Number of all queried backups.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupSummariesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupSummariesResponseParams `json:"Response"`
}

func (r *DescribeBackupSummariesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupSummariesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaseBackupsRequestParams struct {
	// Minimum end time of a backup in the format of `2018-01-01 00:00:00`. It is 7 days ago by default.
	MinFinishTime *string `json:"MinFinishTime,omitnil,omitempty" name:"MinFinishTime"`

	// Maximum end time of a backup in the format of `2018-01-01 00:00:00`. It is the current time by default.
	MaxFinishTime *string `json:"MaxFinishTime,omitnil,omitempty" name:"MaxFinishTime"`

	// Filter instances by using one or more filters. Valid values:  `db-instance-idFilter` (filter by instance ID in string),  `db-instance-name` (filter by instance name in string),  `db-instance-ip` (filter by instance VPC IP address in string),  `base-backup-id` (filter by backup set ID in string), 
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The maximum number of results returned per page. Value range: 1-100. Default: `10`
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Data offset, which starts from 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting field. Valid values: `StartTime`, `FinishTime`, `Size`.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `asc` (ascending), `desc` (descending).
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type DescribeBaseBackupsRequest struct {
	*tchttp.BaseRequest
	
	// Minimum end time of a backup in the format of `2018-01-01 00:00:00`. It is 7 days ago by default.
	MinFinishTime *string `json:"MinFinishTime,omitnil,omitempty" name:"MinFinishTime"`

	// Maximum end time of a backup in the format of `2018-01-01 00:00:00`. It is the current time by default.
	MaxFinishTime *string `json:"MaxFinishTime,omitnil,omitempty" name:"MaxFinishTime"`

	// Filter instances by using one or more filters. Valid values:  `db-instance-idFilter` (filter by instance ID in string),  `db-instance-name` (filter by instance name in string),  `db-instance-ip` (filter by instance VPC IP address in string),  `base-backup-id` (filter by backup set ID in string), 
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The maximum number of results returned per page. Value range: 1-100. Default: `10`
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Data offset, which starts from 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting field. Valid values: `StartTime`, `FinishTime`, `Size`.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `asc` (ascending), `desc` (descending).
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

func (r *DescribeBaseBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaseBackupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MinFinishTime")
	delete(f, "MaxFinishTime")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaseBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaseBackupsResponseParams struct {
	// Number of queried full backups
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of full backup details
	BaseBackupSet []*BaseBackup `json:"BaseBackupSet,omitnil,omitempty" name:"BaseBackupSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBaseBackupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaseBackupsResponseParams `json:"Response"`
}

func (r *DescribeBaseBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaseBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClassesRequestParams struct {
	// AZ ID, which can be obtained through the `DescribeZones` API.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Database engines. Valid values:
	// 1. `postgresql` (TencentDB for PostgreSQL)
	// 2. `mssql_compatible` (MSSQL compatible-TencentDB for PostgreSQL)
	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`

	// Major version of a database, such as 12 or 13, which can be obtained through the `DescribeDBVersions` API.
	DBMajorVersion *string `json:"DBMajorVersion,omitnil,omitempty" name:"DBMajorVersion"`
}

type DescribeClassesRequest struct {
	*tchttp.BaseRequest
	
	// AZ ID, which can be obtained through the `DescribeZones` API.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Database engines. Valid values:
	// 1. `postgresql` (TencentDB for PostgreSQL)
	// 2. `mssql_compatible` (MSSQL compatible-TencentDB for PostgreSQL)
	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`

	// Major version of a database, such as 12 or 13, which can be obtained through the `DescribeDBVersions` API.
	DBMajorVersion *string `json:"DBMajorVersion,omitnil,omitempty" name:"DBMajorVersion"`
}

func (r *DescribeClassesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClassesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "DBEngine")
	delete(f, "DBMajorVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClassesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClassesResponseParams struct {
	// List of database specifications
	ClassInfoSet []*ClassInfo `json:"ClassInfoSet,omitnil,omitempty" name:"ClassInfoSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClassesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClassesResponseParams `json:"Response"`
}

func (r *DescribeClassesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClassesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloneDBInstanceSpecRequestParams struct {
	// Instance ID.
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Basic backup set ID. Either this parameter or `RecoveryTargetTime` must be passed in. If both are passed in, only this parameter takes effect.
	BackupSetId *string `json:"BackupSetId,omitnil,omitempty" name:"BackupSetId"`

	// Restoration time (UTC+8). Either this parameter or `BackupSetId` must be passed in.
	RecoveryTargetTime *string `json:"RecoveryTargetTime,omitnil,omitempty" name:"RecoveryTargetTime"`
}

type DescribeCloneDBInstanceSpecRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Basic backup set ID. Either this parameter or `RecoveryTargetTime` must be passed in. If both are passed in, only this parameter takes effect.
	BackupSetId *string `json:"BackupSetId,omitnil,omitempty" name:"BackupSetId"`

	// Restoration time (UTC+8). Either this parameter or `BackupSetId` must be passed in.
	RecoveryTargetTime *string `json:"RecoveryTargetTime,omitnil,omitempty" name:"RecoveryTargetTime"`
}

func (r *DescribeCloneDBInstanceSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloneDBInstanceSpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "BackupSetId")
	delete(f, "RecoveryTargetTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloneDBInstanceSpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloneDBInstanceSpecResponseParams struct {
	// Code of the minimum specification available for purchase.
	MinSpecCode *string `json:"MinSpecCode,omitnil,omitempty" name:"MinSpecCode"`

	// The minimum disk capacity in GB available for purchase.
	MinStorage *int64 `json:"MinStorage,omitnil,omitempty" name:"MinStorage"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloneDBInstanceSpecResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloneDBInstanceSpecResponseParams `json:"Response"`
}

func (r *DescribeCloneDBInstanceSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloneDBInstanceSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBBackupsRequestParams struct {
	// Instance ID in the format of postgres-4wdeb0zv.
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Backup mode (1: full). Currently, only full backup is supported. The value is 1.
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Query start time in the format of 2018-06-10 17:06:38, which cannot be more than 7 days ago
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Query end time in the format of 2018-06-10 17:06:38
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Number of entries to be returned per page for backup list. Default value: 20. Minimum value: 1. Maximum value: 100. (If this parameter is left empty or 0, the default value will be used)
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page number for data return in paged query. Pagination starts from 0. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeDBBackupsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of postgres-4wdeb0zv.
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Backup mode (1: full). Currently, only full backup is supported. The value is 1.
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// Query start time in the format of 2018-06-10 17:06:38, which cannot be more than 7 days ago
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Query end time in the format of 2018-06-10 17:06:38
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Number of entries to be returned per page for backup list. Default value: 20. Minimum value: 1. Maximum value: 100. (If this parameter is left empty or 0, the default value will be used)
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page number for data return in paged query. Pagination starts from 0. Default value: 0.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeDBBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBBackupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "Type")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBBackupsResponseParams struct {
	// Number of backup files in the returned backup list
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Backup list
	BackupList []*DBBackup `json:"BackupList,omitnil,omitempty" name:"BackupList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBBackupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBBackupsResponseParams `json:"Response"`
}

func (r *DescribeDBBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBErrlogsRequestParams struct {
	// Instance ID	
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// u200cu200cu200cQuery start time in the format of 2018-01-01 00:00:00. The log is retained for seven days by default, so the start time must fall within the retention period.	
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// u200cu200cu200cu200cQuery end time in the format of 2018-01-01 00:00:00	
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Database name
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// Keywords used for search
	SearchKeys []*string `json:"SearchKeys,omitnil,omitempty" name:"SearchKeys"`

	// Number of results returned per page. Value range: 1-100. Default value: `50`.	
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Data offset, which starts from 0. Default value: `0`.	
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeDBErrlogsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID	
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// u200cu200cu200cQuery start time in the format of 2018-01-01 00:00:00. The log is retained for seven days by default, so the start time must fall within the retention period.	
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// u200cu200cu200cu200cQuery end time in the format of 2018-01-01 00:00:00	
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Database name
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// Keywords used for search
	SearchKeys []*string `json:"SearchKeys,omitnil,omitempty" name:"SearchKeys"`

	// Number of results returned per page. Value range: 1-100. Default value: `50`.	
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Data offset, which starts from 0. Default value: `0`.	
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeDBErrlogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBErrlogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "DatabaseName")
	delete(f, "SearchKeys")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBErrlogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBErrlogsResponseParams struct {
	// Number of logs returned in a single query. Maximum value: `10000`.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Detailed sets of error logs
	Details []*ErrLogDetail `json:"Details,omitnil,omitempty" name:"Details"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBErrlogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBErrlogsResponseParams `json:"Response"`
}

func (r *DescribeDBErrlogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBErrlogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceAttributeRequestParams struct {
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

type DescribeDBInstanceAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

func (r *DescribeDBInstanceAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceAttributeResponseParams struct {
	// Instance details.
	DBInstance *DBInstance `json:"DBInstance,omitnil,omitempty" name:"DBInstance"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstanceAttributeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceAttributeResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceHAConfigRequestParams struct {
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

type DescribeDBInstanceHAConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

func (r *DescribeDBInstanceHAConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceHAConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceHAConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceHAConfigResponseParams struct {
	// Primary-standby sync mode. Valid values:
	// <li>`Semi-sync`
	// <li>`Async`
	SyncMode *string `json:"SyncMode,omitnil,omitempty" name:"SyncMode"`

	// Maximum data lag for high-availability standby server. The standby node can be promoted to the primary node when its data lag and the delay time are both less than the value of `MaxStandbyLatency` and `MaxStandbyLag` respectively.
	// <li>Unit: byte
	// <li>Value range: 1073741824-322122547200
	MaxStandbyLatency *uint64 `json:"MaxStandbyLatency,omitnil,omitempty" name:"MaxStandbyLatency"`

	// The maximum delay for high-availability standby server The standby node can be promoted to the primary node when its data lag and the delay time are both less than or equals to the value of `MaxStandbyLatency` and `MaxStandbyLag` respectively.
	// <li>Unit: s
	// <li>Value range: 5-10
	MaxStandbyLag *uint64 `json:"MaxStandbyLag,omitnil,omitempty" name:"MaxStandbyLag"`

	// Maximum data sync lag for u200du200dstandby server. If data lag of the standby node and the delay time are both less than or equals to the values of `MaxSyncStandbyLatency` and `MaxSyncStandbyLag` respectively, the standby server adopts semi-sync replication; if not, it adopts async replication.
	// This value is only valid for the instance with `SyncMode` set to `Semi-sync`.
	// This field returns null for async instance
	// and semi-sync (non-downgradable to async) instance.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	MaxSyncStandbyLatency *uint64 `json:"MaxSyncStandbyLatency,omitnil,omitempty" name:"MaxSyncStandbyLatency"`

	// Maximum sync delay time for u200dstandby server. If the delay time for u200dstandby server and the data lag are both less than or equals to the values of `MaxSyncStandbyLag` and `MaxSyncStandbyLatency` respectively, the standby server adopts sync replication mode; if not, it adopts async replication.
	// This value is only valid for the instance with `SyncMode` set to `Semi-sync`.
	// This field will not return for async instance
	// and semi-sync (non-downgradable to async) instance.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	MaxSyncStandbyLag *uint64 `json:"MaxSyncStandbyLag,omitnil,omitempty" name:"MaxSyncStandbyLag"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstanceHAConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceHAConfigResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceHAConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceHAConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceParametersRequestParams struct {
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Name of the parameter to be queried. If `ParamName` is left empty or not passed in, the list of all parameters will be returned.
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`
}

type DescribeDBInstanceParametersRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Name of the parameter to be queried. If `ParamName` is left empty or not passed in, the list of all parameters will be returned.
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`
}

func (r *DescribeDBInstanceParametersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceParametersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "ParamName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceParametersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceParametersResponseParams struct {
	// Total number of the parameters in the returned list
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Details of the returned parameter list
	Detail []*ParamInfo `json:"Detail,omitnil,omitempty" name:"Detail"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstanceParametersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceParametersResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceParametersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceSecurityGroupsRequestParams struct {
	// Instance ID. Either this parameter or `ReadOnlyGroupId` must be passed in. If both parameters are passed in, `ReadOnlyGroupId` will be ignored.
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// RO group ID. Either this parameter or `DBInstanceId` must be passed in. To query the security groups associated with the RO groups, only pass in `ReadOnlyGroupId`.
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`
}

type DescribeDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID. Either this parameter or `ReadOnlyGroupId` must be passed in. If both parameters are passed in, `ReadOnlyGroupId` will be ignored.
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// RO group ID. Either this parameter or `DBInstanceId` must be passed in. To query the security groups associated with the RO groups, only pass in `ReadOnlyGroupId`.
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`
}

func (r *DescribeDBInstanceSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "ReadOnlyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceSecurityGroupsResponseParams struct {
	// Information of security groups in array
	SecurityGroupSet []*SecurityGroup `json:"SecurityGroupSet,omitnil,omitempty" name:"SecurityGroupSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstanceSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceSecurityGroupsResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesRequestParams struct {
	// Filter instances using one or more criteria. Valid filter names:
	// db-instance-id: filter by instance ID (in string format)
	// db-instance-name: filter by instance name (in string format)
	// db-project-id: filter by project ID (in integer format)
	// db-pay-mode: filter by billing mode (in string format)
	// db-tag-key: filter by tag key (in string format)
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The maximum number of results returned per page. Value range: 1-100. Default: `10`
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Data offset, which starts from 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting metric, such as instance name or creation time. Valid values: DBInstanceId, CreateTime, Name, EndTime
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `asc` (ascending), `desc` (descending)
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Filter instances using one or more criteria. Valid filter names:
	// db-instance-id: filter by instance ID (in string format)
	// db-instance-name: filter by instance name (in string format)
	// db-project-id: filter by project ID (in integer format)
	// db-pay-mode: filter by billing mode (in string format)
	// db-tag-key: filter by tag key (in string format)
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The maximum number of results returned per page. Value range: 1-100. Default: `10`
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Data offset, which starts from 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting metric, such as instance name or creation time. Valid values: DBInstanceId, CreateTime, Name, EndTime
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `asc` (ascending), `desc` (descending)
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
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
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesResponseParams struct {
	// Number of instances found.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Instance details set.
	DBInstanceSet []*DBInstance `json:"DBInstanceSet,omitnil,omitempty" name:"DBInstanceSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeDBSlowlogsRequestParams struct {
	// Instance ID in the format of postgres-lnp6j617
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Query start time in the format of 2018-06-10 17:06:38, which cannot be more than 7 days ago
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Query end time in the format of 2018-06-10 17:06:38
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Database name
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// Metric for sorting. Valid values: `sum_calls` (total number of calls), `sum_cost_time` (total time consumed)
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. desc: descending, asc: ascending
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Number of entries returned per page. Value range: 1-100. Default value: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page number for data return in paged query. Pagination starts from 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeDBSlowlogsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of postgres-lnp6j617
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Query start time in the format of 2018-06-10 17:06:38, which cannot be more than 7 days ago
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Query end time in the format of 2018-06-10 17:06:38
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Database name
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// Metric for sorting. Valid values: `sum_calls` (total number of calls), `sum_cost_time` (total time consumed)
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. desc: descending, asc: ascending
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Number of entries returned per page. Value range: 1-100. Default value: 20.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Page number for data return in paged query. Pagination starts from 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeDBSlowlogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSlowlogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "DatabaseName")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSlowlogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSlowlogsResponseParams struct {
	// Number of date entries returned this time
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Slow query log details
	Detail *SlowlogDetail `json:"Detail,omitnil,omitempty" name:"Detail"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBSlowlogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSlowlogsResponseParams `json:"Response"`
}

func (r *DescribeDBSlowlogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSlowlogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBVersionsRequestParams struct {

}

type DescribeDBVersionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDBVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBVersionsResponseParams struct {
	// List of database versions
	VersionSet []*Version `json:"VersionSet,omitnil,omitempty" name:"VersionSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBVersionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBVersionsResponseParams `json:"Response"`
}

func (r *DescribeDBVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBXlogsRequestParams struct {
	// Instance ID in the format of postgres-4wdeb0zv.
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Query start time in the format of 2018-06-10 17:06:38, which cannot be more than 7 days ago
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Query end time in the format of 2018-06-10 17:06:38
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Page number for data return in paged query. Pagination starts from 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries returned per page in paged query. Value range: 1-100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDBXlogsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of postgres-4wdeb0zv.
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Query start time in the format of 2018-06-10 17:06:38, which cannot be more than 7 days ago
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Query end time in the format of 2018-06-10 17:06:38
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Page number for data return in paged query. Pagination starts from 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Number of entries returned per page in paged query. Value range: 1-100.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeDBXlogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBXlogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBXlogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBXlogsResponseParams struct {
	// Number of date entries returned this time.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Xlog list
	XlogList []*Xlog `json:"XlogList,omitnil,omitempty" name:"XlogList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBXlogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBXlogsResponseParams `json:"Response"`
}

func (r *DescribeDBXlogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBXlogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabasesRequestParams struct {
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

type DescribeDatabasesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
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
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabasesResponseParams struct {
	// Database information
	Items []*string `json:"Items,omitnil,omitempty" name:"Items"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeDefaultParametersRequestParams struct {
	// The major database version number, such as 11, 12, 13.
	DBMajorVersion *string `json:"DBMajorVersion,omitnil,omitempty" name:"DBMajorVersion"`

	// Database engine, such as postgresql, mssql_compatible.
	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`
}

type DescribeDefaultParametersRequest struct {
	*tchttp.BaseRequest
	
	// The major database version number, such as 11, 12, 13.
	DBMajorVersion *string `json:"DBMajorVersion,omitnil,omitempty" name:"DBMajorVersion"`

	// Database engine, such as postgresql, mssql_compatible.
	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`
}

func (r *DescribeDefaultParametersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefaultParametersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBMajorVersion")
	delete(f, "DBEngine")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDefaultParametersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDefaultParametersResponseParams struct {
	// Number of parameters
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Parameter information
	// Note: This field may return null, indicating that no valid values can be obtained.
	ParamInfoSet []*ParamInfo `json:"ParamInfoSet,omitnil,omitempty" name:"ParamInfoSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDefaultParametersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDefaultParametersResponseParams `json:"Response"`
}

func (r *DescribeDefaultParametersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefaultParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEncryptionKeysRequestParams struct {
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

type DescribeEncryptionKeysRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

func (r *DescribeEncryptionKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEncryptionKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEncryptionKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEncryptionKeysResponseParams struct {
	// Instance key list
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	EncryptionKeys []*EncryptionKey `json:"EncryptionKeys,omitnil,omitempty" name:"EncryptionKeys"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEncryptionKeysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEncryptionKeysResponseParams `json:"Response"`
}

func (r *DescribeEncryptionKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEncryptionKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogBackupsRequestParams struct {
	// Minimum end time of a backup in the format of `2018-01-01 00:00:00`. It is 7 days ago by default.
	MinFinishTime *string `json:"MinFinishTime,omitnil,omitempty" name:"MinFinishTime"`

	// Maximum end time of a backup in the format of `2018-01-01 00:00:00`. It is the current time by default.
	MaxFinishTime *string `json:"MaxFinishTime,omitnil,omitempty" name:"MaxFinishTime"`

	// Filter instances using one or more criteria. Valid filter names:
	// db-instance-id: Filter by instance ID (in string format).
	// db-instance-name: Filter by instance name (in string format).
	// db-instance-ip: Filter by instance VPC IP (in string format).
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The maximum number of results returned per page. Value range: 1-100. Default: `10`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Data offset, which starts from 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting field. Valid values: `StartTime`, `FinishTime`, `Size`.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `asc` (ascending), `desc` (descending).
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type DescribeLogBackupsRequest struct {
	*tchttp.BaseRequest
	
	// Minimum end time of a backup in the format of `2018-01-01 00:00:00`. It is 7 days ago by default.
	MinFinishTime *string `json:"MinFinishTime,omitnil,omitempty" name:"MinFinishTime"`

	// Maximum end time of a backup in the format of `2018-01-01 00:00:00`. It is the current time by default.
	MaxFinishTime *string `json:"MaxFinishTime,omitnil,omitempty" name:"MaxFinishTime"`

	// Filter instances using one or more criteria. Valid filter names:
	// db-instance-id: Filter by instance ID (in string format).
	// db-instance-name: Filter by instance name (in string format).
	// db-instance-ip: Filter by instance VPC IP (in string format).
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The maximum number of results returned per page. Value range: 1-100. Default: `10`.
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Data offset, which starts from 0.
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting field. Valid values: `StartTime`, `FinishTime`, `Size`.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `asc` (ascending), `desc` (descending).
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

func (r *DescribeLogBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogBackupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MinFinishTime")
	delete(f, "MaxFinishTime")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogBackupsResponseParams struct {
	// Number of queried log backups
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of log backup details
	LogBackupSet []*LogBackup `json:"LogBackupSet,omitnil,omitempty" name:"LogBackupSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLogBackupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogBackupsResponseParams `json:"Response"`
}

func (r *DescribeLogBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrdersRequestParams struct {
	// Order name set
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`
}

type DescribeOrdersRequest struct {
	*tchttp.BaseRequest
	
	// Order name set
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`
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
	// Number of orders
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Order array
	Deals []*PgDeal `json:"Deals,omitnil,omitempty" name:"Deals"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeParameterTemplateAttributesRequestParams struct {
	// Parameter template ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DescribeParameterTemplateAttributesRequest struct {
	*tchttp.BaseRequest
	
	// Parameter template ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DescribeParameterTemplateAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParameterTemplateAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeParameterTemplateAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParameterTemplateAttributesResponseParams struct {
	// Parameter template ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Number of parameters contained in the parameter template
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Parameter information contained in the parameter template
	// Note: This field may return null, indicating that no valid values can be obtained.
	ParamInfoSet []*ParamInfo `json:"ParamInfoSet,omitnil,omitempty" name:"ParamInfoSet"`

	// Parameter template name
	// Note: This field may return null, indicating that no valid values can be obtained.
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// Database version applicable to a parameter template
	// Note: This field may return null, indicating that no valid values can be obtained.
	DBMajorVersion *string `json:"DBMajorVersion,omitnil,omitempty" name:"DBMajorVersion"`

	// Database engine applicable to a parameter template
	// Note: This field may return null, indicating that no valid values can be obtained.
	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`

	// Parameter template description
	// Note: This field may return null, indicating that no valid values can be obtained.
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeParameterTemplateAttributesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeParameterTemplateAttributesResponseParams `json:"Response"`
}

func (r *DescribeParameterTemplateAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParameterTemplateAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParameterTemplatesRequestParams struct {
	// Filter conditions. Valid values: `TemplateName`, `TemplateId`, `DBMajorVersion`, `DBEngine`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The maximum number of results returned per page. Value range: 0-100. Default: `20`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Data offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting metric. Valid values: `CreateTime`, `TemplateName`, `DBMajorVersion`.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `asc` (ascending order),`desc` (descending order).
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type DescribeParameterTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// Filter conditions. Valid values: `TemplateName`, `TemplateId`, `DBMajorVersion`, `DBEngine`.
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The maximum number of results returned per page. Value range: 0-100. Default: `20`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Data offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting metric. Valid values: `CreateTime`, `TemplateName`, `DBMajorVersion`.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `asc` (ascending order),`desc` (descending order).
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

func (r *DescribeParameterTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParameterTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeParameterTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParameterTemplatesResponseParams struct {
	// The total number of eligible parameter templates
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Parameter template list
	ParameterTemplateSet []*ParameterTemplate `json:"ParameterTemplateSet,omitnil,omitempty" name:"ParameterTemplateSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeParameterTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeParameterTemplatesResponseParams `json:"Response"`
}

func (r *DescribeParameterTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParameterTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamsEventRequestParams struct {
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

type DescribeParamsEventRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

func (r *DescribeParamsEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamsEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeParamsEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamsEventResponseParams struct {
	// Total number of modified parameters
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Details of parameter modification events
	EventItems []*EventItem `json:"EventItems,omitnil,omitempty" name:"EventItems"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeParamsEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeParamsEventResponseParams `json:"Response"`
}

func (r *DescribeParamsEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamsEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductConfigRequestParams struct {
	// AZ name
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Database engines. Valid values:
	// 1. `postgresql` (TencentDB for PostgreSQL)
	// 2. `mssql_compatible` (MSSQL compatible-TencentDB for PostgreSQL)
	// Default value: `postgresql`
	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`
}

type DescribeProductConfigRequest struct {
	*tchttp.BaseRequest
	
	// AZ name
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Database engines. Valid values:
	// 1. `postgresql` (TencentDB for PostgreSQL)
	// 2. `mssql_compatible` (MSSQL compatible-TencentDB for PostgreSQL)
	// Default value: `postgresql`
	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`
}

func (r *DescribeProductConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "DBEngine")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductConfigResponseParams struct {
	// Purchasable specification list.
	SpecInfoList []*SpecInfo `json:"SpecInfoList,omitnil,omitempty" name:"SpecInfoList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProductConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProductConfigResponseParams `json:"Response"`
}

func (r *DescribeProductConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReadOnlyGroupsRequestParams struct {
	// Filter instances by using one or more filters. Valid values:  `db-master-instance-id` (filter by the primary instance ID in string), `read-only-group-id` (filter by the read-only group ID in string),
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The number of results per page. Default value: 10.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Specify which page is displayed. Default value: 1 (the first page).
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Sorting criterion. Valid values: `ROGroupId`, `CreateTime`, `Name`.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `desc`, `asc`.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type DescribeReadOnlyGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Filter instances by using one or more filters. Valid values:  `db-master-instance-id` (filter by the primary instance ID in string), `read-only-group-id` (filter by the read-only group ID in string),
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// The number of results per page. Default value: 10.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Specify which page is displayed. Default value: 1 (the first page).
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// Sorting criterion. Valid values: `ROGroupId`, `CreateTime`, `Name`.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `desc`, `asc`.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

func (r *DescribeReadOnlyGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReadOnlyGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReadOnlyGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReadOnlyGroupsResponseParams struct {
	// RO group list
	ReadOnlyGroupList []*ReadOnlyGroup `json:"ReadOnlyGroupList,omitnil,omitempty" name:"ReadOnlyGroupList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReadOnlyGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReadOnlyGroupsResponseParams `json:"Response"`
}

func (r *DescribeReadOnlyGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReadOnlyGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionsRequestParams struct {

}

type DescribeRegionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionsResponseParams struct {
	// Number of returned results.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Region information set.
	RegionSet []*RegionInfo `json:"RegionSet,omitnil,omitempty" name:"RegionSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRegionsResponseParams `json:"Response"`
}

func (r *DescribeRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServerlessDBInstancesRequestParams struct {
	// Query conditions
	Filter []*Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// The number of queries
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The offset value
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting metric. Currently, only "CreateTime" (instance creation time) is supported.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Ascending and descending are supported.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type DescribeServerlessDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Query conditions
	Filter []*Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// The number of queries
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// The offset value
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting metric. Currently, only "CreateTime" (instance creation time) is supported.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Ascending and descending are supported.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

func (r *DescribeServerlessDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServerlessDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServerlessDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServerlessDBInstancesResponseParams struct {
	// The number of query results
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Query results
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DBInstanceSet []*ServerlessDBInstance `json:"DBInstanceSet,omitnil,omitempty" name:"DBInstanceSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeServerlessDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServerlessDBInstancesResponseParams `json:"Response"`
}

func (r *DescribeServerlessDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServerlessDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowQueryAnalysisRequestParams struct {
	// Instance ID.
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Start timestamp of the query range in the format of "YYYY-MM-DD HH:mm:ss". The log is retained for seven days by default, so the start timestamp must fall within the retention period.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End timestamp of the query range in the format of "YYYY-MM-DD HH:mm:ss".
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Filter by database name. This parameter is optional.
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// Sort by field. Valid values: `CallNum`, `CostTime`, `AvgCostTime`. Default value: `CallNum`.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `asc` (ascending), `desc` (descending). Default value: `desc`.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Number of entries per page. Value range: [1,100]. Default value: `50`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset. Value range: [0,INF). Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeSlowQueryAnalysisRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Start timestamp of the query range in the format of "YYYY-MM-DD HH:mm:ss". The log is retained for seven days by default, so the start timestamp must fall within the retention period.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End timestamp of the query range in the format of "YYYY-MM-DD HH:mm:ss".
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Filter by database name. This parameter is optional.
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// Sort by field. Valid values: `CallNum`, `CostTime`, `AvgCostTime`. Default value: `CallNum`.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Sorting order. Valid values: `asc` (ascending), `desc` (descending). Default value: `desc`.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Number of entries per page. Value range: [1,100]. Default value: `50`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset. Value range: [0,INF). Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeSlowQueryAnalysisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowQueryAnalysisRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "DatabaseName")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowQueryAnalysisRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowQueryAnalysisResponseParams struct {
	// The total number of query results.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Detailed analysis.
	Detail *Detail `json:"Detail,omitnil,omitempty" name:"Detail"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowQueryAnalysisResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowQueryAnalysisResponseParams `json:"Response"`
}

func (r *DescribeSlowQueryAnalysisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowQueryAnalysisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowQueryListRequestParams struct {
	// Instance ID.
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Start timestamp of the query range in the format of "YYYY-MM-DD HH:mm:ss". The log is retained for seven days by default, so the start timestamp must fall within the retention period.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End timestamp of the query range in the format of "YYYY-MM-DD HH:mm:ss".
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Filter by database name. This parameter is optional.
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// Sorting order. Valid values: `asc` (ascending), `desc` (descending). Default value: `desc`.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Sort by field. Valid values: `SessionStartTime` (default), `Duration`.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Number of entries per page. Value range: [1,100]. Default value: `20`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset. Value range: [0,INF). Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeSlowQueryListRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Start timestamp of the query range in the format of "YYYY-MM-DD HH:mm:ss". The log is retained for seven days by default, so the start timestamp must fall within the retention period.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End timestamp of the query range in the format of "YYYY-MM-DD HH:mm:ss".
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Filter by database name. This parameter is optional.
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// Sorting order. Valid values: `asc` (ascending), `desc` (descending). Default value: `desc`.
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// Sort by field. Valid values: `SessionStartTime` (default), `Duration`.
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Number of entries per page. Value range: [1,100]. Default value: `20`.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pagination offset. Value range: [0,INF). Default value: `0`.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeSlowQueryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowQueryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "DatabaseName")
	delete(f, "OrderByType")
	delete(f, "OrderBy")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowQueryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowQueryListResponseParams struct {
	// The total number of slow query statements during the specified period of time.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Analysis of the execution time of slow query statements by classifying them to different time ranges. These slow query statements fall within the query range you specified in the request parameters.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DurationAnalysis []*DurationAnalysis `json:"DurationAnalysis,omitnil,omitempty" name:"DurationAnalysis"`

	// The list of slow query details during the specified period of time.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	RawSlowQueryList []*RawSlowQuery `json:"RawSlowQueryList,omitnil,omitempty" name:"RawSlowQueryList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowQueryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowQueryListResponseParams `json:"Response"`
}

func (r *DescribeSlowQueryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowQueryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZonesRequestParams struct {

}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZonesResponseParams struct {
	// Number of returned results.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// AZ information set.
	ZoneSet []*ZoneInfo `json:"ZoneSet,omitnil,omitempty" name:"ZoneSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeZonesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZonesResponseParams `json:"Response"`
}

func (r *DescribeZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyDBInstanceRequestParams struct {
	// The ID of the instance to be eliminated
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

type DestroyDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the instance to be eliminated
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
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
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyDBInstanceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type Detail struct {
	// The total execution time (in ms) of all slow query statements during the specified period of time
	TotalTime *float64 `json:"TotalTime,omitnil,omitempty" name:"TotalTime"`

	// The total number of all slow query statements during the specified period of time
	TotalCallNum *uint64 `json:"TotalCallNum,omitnil,omitempty" name:"TotalCallNum"`

	// The statistical analysis list of slow queries
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AnalysisItems []*AnalysisItems `json:"AnalysisItems,omitnil,omitempty" name:"AnalysisItems"`
}

// Predefined struct for user
type DisIsolateDBInstancesRequestParams struct {
	// Instance ID list. Currently, you can't remove multiple instances from isolation in batches. Only one instance ID can be passed in here.
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitnil,omitempty" name:"DBInstanceIdSet"`

	// Validity period in months
	// <li>Monthly subscription: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	// <li>Pay-as-you-go: `1`.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Whether to use vouchers. Valid values:
	// <li>`true`: Yes.
	// u200c<li>`false`: No.
	// Default value: `false`.
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Voucher ID list
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`
}

type DisIsolateDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID list. Currently, you can't remove multiple instances from isolation in batches. Only one instance ID can be passed in here.
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitnil,omitempty" name:"DBInstanceIdSet"`

	// Validity period in months
	// <li>Monthly subscription: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	// <li>Pay-as-you-go: `1`.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Whether to use vouchers. Valid values:
	// <li>`true`: Yes.
	// u200c<li>`false`: No.
	// Default value: `false`.
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Voucher ID list
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`
}

func (r *DisIsolateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisIsolateDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceIdSet")
	delete(f, "Period")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisIsolateDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisIsolateDBInstancesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisIsolateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DisIsolateDBInstancesResponseParams `json:"Response"`
}

func (r *DisIsolateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisIsolateDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DurationAnalysis struct {
	// Time range
	TimeSegment *string `json:"TimeSegment,omitnil,omitempty" name:"TimeSegment"`

	// The number of slow query statements whose execution time falls within the time range
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type EncryptionKey struct {
	// Encrypted KeyId of KMS instance
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// Encryption key alias of KMS instance 
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	KeyAlias *string `json:"KeyAlias,omitnil,omitempty" name:"KeyAlias"`

	// Instance DEK ciphertext
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	DEKCipherTextBlob *string `json:"DEKCipherTextBlob,omitnil,omitempty" name:"DEKCipherTextBlob"`

	// Whether the key is enabled. Valid values: `1` (yes), `0` (no)
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	IsEnabled *int64 `json:"IsEnabled,omitnil,omitempty" name:"IsEnabled"`

	// Region where KMS key resides
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	KeyRegion *string `json:"KeyRegion,omitnil,omitempty" name:"KeyRegion"`

	// DEK key creation time
	// Note: This field may return `null`, indicating that no valid value can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type ErrLogDetail struct {
	// Username
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Error occurrence time
	ErrTime *string `json:"ErrTime,omitnil,omitempty" name:"ErrTime"`

	// Error message
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`
}

type EventInfo struct {
	// Parameter name
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// Original parameter value
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	OldValue *string `json:"OldValue,omitnil,omitempty" name:"OldValue"`

	// New parameter value in this modification event
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	NewValue *string `json:"NewValue,omitnil,omitempty" name:"NewValue"`

	// Start time of parameter modification
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// Start time when the modified parameter takes effect
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	EffectiveTime *string `json:"EffectiveTime,omitnil,omitempty" name:"EffectiveTime"`

	// Modification status
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// Operator (generally, the value is the UIN of a sub-user)
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// Event log
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	EventLog *string `json:"EventLog,omitnil,omitempty" name:"EventLog"`
}

type EventItem struct {
	// Parameter name
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// The number of modification events
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	EventCount *int64 `json:"EventCount,omitnil,omitempty" name:"EventCount"`

	// Modification event details
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	EventDetail []*EventInfo `json:"EventDetail,omitnil,omitempty" name:"EventDetail"`
}

type Filter struct {
	// Filter name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// One or more filter values.
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

// Predefined struct for user
type InitDBInstancesRequestParams struct {
	// Instance ID set.
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitnil,omitempty" name:"DBInstanceIdSet"`

	// Instance admin account username.
	AdminName *string `json:"AdminName,omitnil,omitempty" name:"AdminName"`

	// Password corresponding to instance root account username.
	AdminPassword *string `json:"AdminPassword,omitnil,omitempty" name:"AdminPassword"`

	// Instance character set. Valid values: UTF8, LATIN1.
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`
}

type InitDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID set.
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitnil,omitempty" name:"DBInstanceIdSet"`

	// Instance admin account username.
	AdminName *string `json:"AdminName,omitnil,omitempty" name:"AdminName"`

	// Password corresponding to instance root account username.
	AdminPassword *string `json:"AdminPassword,omitnil,omitempty" name:"AdminPassword"`

	// Instance character set. Valid values: UTF8, LATIN1.
	Charset *string `json:"Charset,omitnil,omitempty" name:"Charset"`
}

func (r *InitDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InitDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceIdSet")
	delete(f, "AdminName")
	delete(f, "AdminPassword")
	delete(f, "Charset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InitDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InitDBInstancesResponseParams struct {
	// Instance ID set.
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitnil,omitempty" name:"DBInstanceIdSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InitDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *InitDBInstancesResponseParams `json:"Response"`
}

func (r *InitDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InitDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceCreateDBInstancesRequestParams struct {
	// AZ ID, which can be obtained through the `Zone` field in the returned value of the `DescribeZones` API.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Specification ID, which can be obtained through the `SpecCode` field in the returned value of the `DescribeClasses` API.
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// Storage capacity size in GB.
	Storage *uint64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Number of instances. Maximum value: 100. If you need to create more instances at a time, please contact customer service.
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// Length of purchase in months. Currently, only 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, and 36 are supported.
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// [Disused] Billing ID, which can be obtained through the `Pid` field in the returned value of the `DescribeProductConfig` API.
	Pid *uint64 `json:"Pid,omitnil,omitempty" name:"Pid"`

	// Instance billing type. Valid value: POSTPAID_BY_HOUR (pay-as-you-go)
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Instance type. Default value: `primary`. Valid values:
	// `primary` (dual-server high-availability, one-primary-one-standby)
	// `readonly` (read-only instance)
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`


	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`
}

type InquiryPriceCreateDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// AZ ID, which can be obtained through the `Zone` field in the returned value of the `DescribeZones` API.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Specification ID, which can be obtained through the `SpecCode` field in the returned value of the `DescribeClasses` API.
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// Storage capacity size in GB.
	Storage *uint64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Number of instances. Maximum value: 100. If you need to create more instances at a time, please contact customer service.
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// Length of purchase in months. Currently, only 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, and 36 are supported.
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// [Disused] Billing ID, which can be obtained through the `Pid` field in the returned value of the `DescribeProductConfig` API.
	Pid *uint64 `json:"Pid,omitnil,omitempty" name:"Pid"`

	// Instance billing type. Valid value: POSTPAID_BY_HOUR (pay-as-you-go)
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Instance type. Default value: `primary`. Valid values:
	// `primary` (dual-server high-availability, one-primary-one-standby)
	// `readonly` (read-only instance)
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`
}

func (r *InquiryPriceCreateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "SpecCode")
	delete(f, "Storage")
	delete(f, "InstanceCount")
	delete(f, "Period")
	delete(f, "Pid")
	delete(f, "InstanceChargeType")
	delete(f, "InstanceType")
	delete(f, "DBEngine")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceCreateDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceCreateDBInstancesResponseParams struct {
	// Published price in US Cent
	OriginalPrice *uint64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// Discounted total amount in US Cent
	Price *uint64 `json:"Price,omitnil,omitempty" name:"Price"`

	// Currency, such as USD.
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceCreateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceCreateDBInstancesResponseParams `json:"Response"`
}

func (r *InquiryPriceCreateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRenewDBInstanceRequestParams struct {
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Renewal duration in months. Maximum value: 48
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`
}

type InquiryPriceRenewDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Renewal duration in months. Maximum value: 48
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`
}

func (r *InquiryPriceRenewDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRenewDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceRenewDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRenewDBInstanceResponseParams struct {
	// Published price in cents. For example, 24650 indicates 246.5 USD.
	OriginalPrice *int64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// Discounted total amount. For example, 24650 indicates 246.5 USD.
	Price *int64 `json:"Price,omitnil,omitempty" name:"Price"`

	// Currency, such as USD.
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceRenewDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceRenewDBInstanceResponseParams `json:"Response"`
}

func (r *InquiryPriceRenewDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRenewDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceUpgradeDBInstanceRequestParams struct {
	// Instance disk size in GB
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Instance memory size in GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Instance ID in the format of postgres-hez4fh0v
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Instance billing type. Valid value: `POSTPAID_BY_HOUR` (pay-as-you-go hourly)
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`
}

type InquiryPriceUpgradeDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance disk size in GB
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Instance memory size in GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Instance ID in the format of postgres-hez4fh0v
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Instance billing type. Valid value: `POSTPAID_BY_HOUR` (pay-as-you-go hourly)
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`
}

func (r *InquiryPriceUpgradeDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceUpgradeDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Storage")
	delete(f, "Memory")
	delete(f, "DBInstanceId")
	delete(f, "InstanceChargeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceUpgradeDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceUpgradeDBInstanceResponseParams struct {
	// Total cost before discount.
	OriginalPrice *int64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// Discounted total amount
	Price *int64 `json:"Price,omitnil,omitempty" name:"Price"`

	// Currency, such as USD.
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceUpgradeDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceUpgradeDBInstanceResponseParams `json:"Response"`
}

func (r *InquiryPriceUpgradeDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceUpgradeDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateDBInstancesRequestParams struct {
	// List of instance IDs. Note that currently you cannot isolate multiple instances at the same time. Only one instance ID can be passed in here.
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitnil,omitempty" name:"DBInstanceIdSet"`
}

type IsolateDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// List of instance IDs. Note that currently you cannot isolate multiple instances at the same time. Only one instance ID can be passed in here.
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitnil,omitempty" name:"DBInstanceIdSet"`
}

func (r *IsolateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateDBInstancesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type IsolateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *IsolateDBInstancesResponseParams `json:"Response"`
}

func (r *IsolateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogBackup struct {
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Unique ID of a backup file
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Backup file name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Backup method, including physical and logical.
	BackupMethod *string `json:"BackupMethod,omitnil,omitempty" name:"BackupMethod"`

	// Backup mode, including automatic and manual.
	BackupMode *string `json:"BackupMode,omitnil,omitempty" name:"BackupMode"`

	// Backup task status
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// Backup set size in bytes
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`

	// Backup start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Backup end time
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// Backup expiration time
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

// Predefined struct for user
type ModifyAccountRemarkRequestParams struct {
	// Instance ID in the format of postgres-4wdeb0zv
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Instance username
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// New remarks corresponding to user `UserName`
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ModifyAccountRemarkRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of postgres-4wdeb0zv
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Instance username
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// New remarks corresponding to user `UserName`
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *ModifyAccountRemarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountRemarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "UserName")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountRemarkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountRemarkResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAccountRemarkResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccountRemarkResponseParams `json:"Response"`
}

func (r *ModifyAccountRemarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountRemarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupDownloadRestrictionRequestParams struct {
	// Type of the network restrictions for downloading a backup file. Valid values: `NONE` (backups can be downloaded over both private and public networks), `INTRANET` (backups can only be downloaded over the private network), `CUSTOMIZE` (backups can be downloaded over specified VPCs or at specified IPs).
	RestrictionType *string `json:"RestrictionType,omitnil,omitempty" name:"RestrictionType"`

	// Whether VPC is allowed. Valid values: `ALLOW` (allow), `DENY` (deny).
	VpcRestrictionEffect *string `json:"VpcRestrictionEffect,omitnil,omitempty" name:"VpcRestrictionEffect"`

	// Whether it is allowed to download the VPC ID list of the backup files.
	VpcIdSet []*string `json:"VpcIdSet,omitnil,omitempty" name:"VpcIdSet"`

	// Whether IP is allowed. Valid values: `ALLOW` (allow), `DENY` (deny).
	IpRestrictionEffect *string `json:"IpRestrictionEffect,omitnil,omitempty" name:"IpRestrictionEffect"`

	// Whether it is allowed to download the IP list of the backup files.
	IpSet []*string `json:"IpSet,omitnil,omitempty" name:"IpSet"`
}

type ModifyBackupDownloadRestrictionRequest struct {
	*tchttp.BaseRequest
	
	// Type of the network restrictions for downloading a backup file. Valid values: `NONE` (backups can be downloaded over both private and public networks), `INTRANET` (backups can only be downloaded over the private network), `CUSTOMIZE` (backups can be downloaded over specified VPCs or at specified IPs).
	RestrictionType *string `json:"RestrictionType,omitnil,omitempty" name:"RestrictionType"`

	// Whether VPC is allowed. Valid values: `ALLOW` (allow), `DENY` (deny).
	VpcRestrictionEffect *string `json:"VpcRestrictionEffect,omitnil,omitempty" name:"VpcRestrictionEffect"`

	// Whether it is allowed to download the VPC ID list of the backup files.
	VpcIdSet []*string `json:"VpcIdSet,omitnil,omitempty" name:"VpcIdSet"`

	// Whether IP is allowed. Valid values: `ALLOW` (allow), `DENY` (deny).
	IpRestrictionEffect *string `json:"IpRestrictionEffect,omitnil,omitempty" name:"IpRestrictionEffect"`

	// Whether it is allowed to download the IP list of the backup files.
	IpSet []*string `json:"IpSet,omitnil,omitempty" name:"IpSet"`
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
	delete(f, "RestrictionType")
	delete(f, "VpcRestrictionEffect")
	delete(f, "VpcIdSet")
	delete(f, "IpRestrictionEffect")
	delete(f, "IpSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupDownloadRestrictionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupDownloadRestrictionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyBackupPlanRequestParams struct {
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// The earliest time to start a backup
	MinBackupStartTime *string `json:"MinBackupStartTime,omitnil,omitempty" name:"MinBackupStartTime"`

	// The latest time to start a backup
	MaxBackupStartTime *string `json:"MaxBackupStartTime,omitnil,omitempty" name:"MaxBackupStartTime"`

	// Backup retention period in days. Value range: 7-1830
	BaseBackupRetentionPeriod *uint64 `json:"BaseBackupRetentionPeriod,omitnil,omitempty" name:"BaseBackupRetentionPeriod"`

	// Backup cycle, which means on which days each week the instance will be backed up. The parameter value should be the lowercase names of the days of the week.
	BackupPeriod []*string `json:"BackupPeriod,omitnil,omitempty" name:"BackupPeriod"`
}

type ModifyBackupPlanRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// The earliest time to start a backup
	MinBackupStartTime *string `json:"MinBackupStartTime,omitnil,omitempty" name:"MinBackupStartTime"`

	// The latest time to start a backup
	MaxBackupStartTime *string `json:"MaxBackupStartTime,omitnil,omitempty" name:"MaxBackupStartTime"`

	// Backup retention period in days. Value range: 7-1830
	BaseBackupRetentionPeriod *uint64 `json:"BaseBackupRetentionPeriod,omitnil,omitempty" name:"BaseBackupRetentionPeriod"`

	// Backup cycle, which means on which days each week the instance will be backed up. The parameter value should be the lowercase names of the days of the week.
	BackupPeriod []*string `json:"BackupPeriod,omitnil,omitempty" name:"BackupPeriod"`
}

func (r *ModifyBackupPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "MinBackupStartTime")
	delete(f, "MaxBackupStartTime")
	delete(f, "BaseBackupRetentionPeriod")
	delete(f, "BackupPeriod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupPlanResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBackupPlanResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBackupPlanResponseParams `json:"Response"`
}

func (r *ModifyBackupPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBaseBackupExpireTimeRequestParams struct {
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Base backup ID
	BaseBackupId *string `json:"BaseBackupId,omitnil,omitempty" name:"BaseBackupId"`

	// New expiration time
	NewExpireTime *string `json:"NewExpireTime,omitnil,omitempty" name:"NewExpireTime"`
}

type ModifyBaseBackupExpireTimeRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Base backup ID
	BaseBackupId *string `json:"BaseBackupId,omitnil,omitempty" name:"BaseBackupId"`

	// New expiration time
	NewExpireTime *string `json:"NewExpireTime,omitnil,omitempty" name:"NewExpireTime"`
}

func (r *ModifyBaseBackupExpireTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBaseBackupExpireTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "BaseBackupId")
	delete(f, "NewExpireTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBaseBackupExpireTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBaseBackupExpireTimeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBaseBackupExpireTimeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBaseBackupExpireTimeResponseParams `json:"Response"`
}

func (r *ModifyBaseBackupExpireTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBaseBackupExpireTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceChargeTypeRequestParams struct {
	// Instance ID in the format of `postgres-6fego161`
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Instance billing mode. Valid values:
	// <li>`PREPAID`: Monthly subscription.
	// <li>`POSTPAID_BY_HOUR`: Pay-as-you-go.
	// Default value: `PREPAID`.
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Validity period in months
	// <li>Monthly subscription: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	// <li>Pay-as-you-go: `1`.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Auto-renewal flag. Valid values:
	// <li>`0`: Manual renewal.
	// <li>`1`: Automatic renewal.
	// Default value: `0`.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// Whether to use vouchers automatically. Valid values:
	// <li>`0`: No.
	// <li>`1`: Yes.
	// Default value: `0`.
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`
}

type ModifyDBInstanceChargeTypeRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of `postgres-6fego161`
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Instance billing mode. Valid values:
	// <li>`PREPAID`: Monthly subscription.
	// <li>`POSTPAID_BY_HOUR`: Pay-as-you-go.
	// Default value: `PREPAID`.
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// Validity period in months
	// <li>Monthly subscription: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	// <li>Pay-as-you-go: `1`.
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Auto-renewal flag. Valid values:
	// <li>`0`: Manual renewal.
	// <li>`1`: Automatic renewal.
	// Default value: `0`.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// Whether to use vouchers automatically. Valid values:
	// <li>`0`: No.
	// <li>`1`: Yes.
	// Default value: `0`.
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`
}

func (r *ModifyDBInstanceChargeTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceChargeTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "InstanceChargeType")
	delete(f, "Period")
	delete(f, "AutoRenewFlag")
	delete(f, "AutoVoucher")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceChargeTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceChargeTypeResponseParams struct {
	// Order name
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceChargeTypeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceChargeTypeResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceChargeTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceChargeTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceDeploymentRequestParams struct {
	// Instance ID.
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Deployment information of the instance node, which will display the information of each AZ when the instance node is deployed across multiple AZs.
	// The information of AZ can be obtained from the `Zone` field in the returned value of the [DescribeZones](https://intl.cloud.tencent.com/document/api/409/16769?from_cn_redirect=1) API.
	DBNodeSet []*DBNode `json:"DBNodeSet,omitnil,omitempty" name:"DBNodeSet"`

	// Switch time for the specified instance after configuration modification.
	// <li>`0`: Switch now. 
	// <li>`1`: Switch at the specified time.
	// <li>`2`: Switch in the maintenance time.
	// Default value: `0`. 
	SwitchTag *int64 `json:"SwitchTag,omitnil,omitempty" name:"SwitchTag"`

	// Switch start time in the format of `HH:MM:SS`, such as 01:00:00. When `SwitchTag` is 0 or 2, this parameter becomes invalid.
	SwitchStartTime *string `json:"SwitchStartTime,omitnil,omitempty" name:"SwitchStartTime"`

	// Switch end time in the format of `HH:MM:SS`, such as 01:30:00. When `SwitchTag` is 0 or 2, this parameter becomes invalid.
	SwitchEndTime *string `json:"SwitchEndTime,omitnil,omitempty" name:"SwitchEndTime"`
}

type ModifyDBInstanceDeploymentRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Deployment information of the instance node, which will display the information of each AZ when the instance node is deployed across multiple AZs.
	// The information of AZ can be obtained from the `Zone` field in the returned value of the [DescribeZones](https://intl.cloud.tencent.com/document/api/409/16769?from_cn_redirect=1) API.
	DBNodeSet []*DBNode `json:"DBNodeSet,omitnil,omitempty" name:"DBNodeSet"`

	// Switch time for the specified instance after configuration modification.
	// <li>`0`: Switch now. 
	// <li>`1`: Switch at the specified time.
	// <li>`2`: Switch in the maintenance time.
	// Default value: `0`. 
	SwitchTag *int64 `json:"SwitchTag,omitnil,omitempty" name:"SwitchTag"`

	// Switch start time in the format of `HH:MM:SS`, such as 01:00:00. When `SwitchTag` is 0 or 2, this parameter becomes invalid.
	SwitchStartTime *string `json:"SwitchStartTime,omitnil,omitempty" name:"SwitchStartTime"`

	// Switch end time in the format of `HH:MM:SS`, such as 01:30:00. When `SwitchTag` is 0 or 2, this parameter becomes invalid.
	SwitchEndTime *string `json:"SwitchEndTime,omitnil,omitempty" name:"SwitchEndTime"`
}

func (r *ModifyDBInstanceDeploymentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceDeploymentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "DBNodeSet")
	delete(f, "SwitchTag")
	delete(f, "SwitchStartTime")
	delete(f, "SwitchEndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceDeploymentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceDeploymentResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceDeploymentResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceDeploymentResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceDeploymentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceDeploymentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceHAConfigRequestParams struct {
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Primary-standby sync mode. Valid values:
	// <li>`Semi-sync`
	// <li>`Async`
	SyncMode *string `json:"SyncMode,omitnil,omitempty" name:"SyncMode"`

	// u200cMaximum data lag for high-availability standby server. The standby node can be promoted to the primary node when its data lag and the delay time are both less than the value of `MaxStandbyLatency` and `MaxStandbyLag` respectively.
	// <li>Unit: byte
	// <li>Value range: 1073741824-322122547200
	MaxStandbyLatency *uint64 `json:"MaxStandbyLatency,omitnil,omitempty" name:"MaxStandbyLatency"`

	// The maximum delay for high-availability standby server The standby node can be promoted to the primary node when its data lag and the delay time are both less or equals to the value of `MaxStandbyLatency` and `MaxStandbyLag` respectively.
	// <li>Unit: s
	// <li>Value range: 5-10
	MaxStandbyLag *uint64 `json:"MaxStandbyLag,omitnil,omitempty" name:"MaxStandbyLag"`

	// Maximum data sync lag for u200dstandby server. If data lag of the standby node and the delay ime are both less than or equals to the values of `MaxSyncStandbyLatency` and `MaxSyncStandbyLag`, the standby server adopts semi-sync replication; if not, it adopts async replication.
	// This value is only valid for the instance with `SyncMode` set to `Semi-sync`.
	// When the semi-sync replication mode of the instance is not allowed to downgrade to async replication, `MaxSyncStandbyLatency` and `MaxSyncStandbyLag` are not required.
	// When the semi-sync instance is allowed to downgrade to async replication, `MaxSyncStandbyLatency` is required and `MaxSyncStandbyLag` must be left empty for PostgreSQL 9; `MaxSyncStandbyLatency` and MaxSyncStandbyLag` are required for PostgreSQL 10 and later.
	MaxSyncStandbyLatency *uint64 `json:"MaxSyncStandbyLatency,omitnil,omitempty" name:"MaxSyncStandbyLatency"`

	// Maximum delay for u200dsync u200dstandby server. If the delay time for u200dstandby server and the data lag are both less than or equals to the value of `MaxSyncStandbyLag` and `MaxSyncStandbyLatency` respectively, the standby server adopts sync replication mode; if not, it adopts async replication.
	// This value is only valid for the instance with `SyncMode` set to `Semi-sync`.
	// When the semi-sync replication mode of the instance is not allowed to downgrade to async replication, `MaxSyncStandbyLatency` and `MaxSyncStandbyLag` are not required.
	// When the semi-sync instance is allowed to downgrade to async replication, `MaxSyncStandbyLatency` is required and `MaxSyncStandbyLag` must be left empty for PostgreSQL 9; `MaxSyncStandbyLatency` and MaxSyncStandbyLag` are required for PostgreSQL 10 and later.
	MaxSyncStandbyLag *uint64 `json:"MaxSyncStandbyLag,omitnil,omitempty" name:"MaxSyncStandbyLag"`
}

type ModifyDBInstanceHAConfigRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Primary-standby sync mode. Valid values:
	// <li>`Semi-sync`
	// <li>`Async`
	SyncMode *string `json:"SyncMode,omitnil,omitempty" name:"SyncMode"`

	// u200cMaximum data lag for high-availability standby server. The standby node can be promoted to the primary node when its data lag and the delay time are both less than the value of `MaxStandbyLatency` and `MaxStandbyLag` respectively.
	// <li>Unit: byte
	// <li>Value range: 1073741824-322122547200
	MaxStandbyLatency *uint64 `json:"MaxStandbyLatency,omitnil,omitempty" name:"MaxStandbyLatency"`

	// The maximum delay for high-availability standby server The standby node can be promoted to the primary node when its data lag and the delay time are both less or equals to the value of `MaxStandbyLatency` and `MaxStandbyLag` respectively.
	// <li>Unit: s
	// <li>Value range: 5-10
	MaxStandbyLag *uint64 `json:"MaxStandbyLag,omitnil,omitempty" name:"MaxStandbyLag"`

	// Maximum data sync lag for u200dstandby server. If data lag of the standby node and the delay ime are both less than or equals to the values of `MaxSyncStandbyLatency` and `MaxSyncStandbyLag`, the standby server adopts semi-sync replication; if not, it adopts async replication.
	// This value is only valid for the instance with `SyncMode` set to `Semi-sync`.
	// When the semi-sync replication mode of the instance is not allowed to downgrade to async replication, `MaxSyncStandbyLatency` and `MaxSyncStandbyLag` are not required.
	// When the semi-sync instance is allowed to downgrade to async replication, `MaxSyncStandbyLatency` is required and `MaxSyncStandbyLag` must be left empty for PostgreSQL 9; `MaxSyncStandbyLatency` and MaxSyncStandbyLag` are required for PostgreSQL 10 and later.
	MaxSyncStandbyLatency *uint64 `json:"MaxSyncStandbyLatency,omitnil,omitempty" name:"MaxSyncStandbyLatency"`

	// Maximum delay for u200dsync u200dstandby server. If the delay time for u200dstandby server and the data lag are both less than or equals to the value of `MaxSyncStandbyLag` and `MaxSyncStandbyLatency` respectively, the standby server adopts sync replication mode; if not, it adopts async replication.
	// This value is only valid for the instance with `SyncMode` set to `Semi-sync`.
	// When the semi-sync replication mode of the instance is not allowed to downgrade to async replication, `MaxSyncStandbyLatency` and `MaxSyncStandbyLag` are not required.
	// When the semi-sync instance is allowed to downgrade to async replication, `MaxSyncStandbyLatency` is required and `MaxSyncStandbyLag` must be left empty for PostgreSQL 9; `MaxSyncStandbyLatency` and MaxSyncStandbyLag` are required for PostgreSQL 10 and later.
	MaxSyncStandbyLag *uint64 `json:"MaxSyncStandbyLag,omitnil,omitempty" name:"MaxSyncStandbyLag"`
}

func (r *ModifyDBInstanceHAConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceHAConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "SyncMode")
	delete(f, "MaxStandbyLatency")
	delete(f, "MaxStandbyLag")
	delete(f, "MaxSyncStandbyLatency")
	delete(f, "MaxSyncStandbyLag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceHAConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceHAConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceHAConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceHAConfigResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceHAConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceHAConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceNameRequestParams struct {
	// Database instance ID in the format of postgres-6fego161
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Instance name, which can contain up to 60 letters, digits, hyphens, and symbols (_-). If this parameter is not specified, "Unnamed" will be displayed by default.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type ModifyDBInstanceNameRequest struct {
	*tchttp.BaseRequest
	
	// Database instance ID in the format of postgres-6fego161
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Instance name, which can contain up to 60 letters, digits, hyphens, and symbols (_-). If this parameter is not specified, "Unnamed" will be displayed by default.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

func (r *ModifyDBInstanceNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceNameResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceNameResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceParametersRequestParams struct {
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Parameters to be modified and their new values
	ParamList []*ParamEntry `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

type ModifyDBInstanceParametersRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Parameters to be modified and their new values
	ParamList []*ParamEntry `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

func (r *ModifyDBInstanceParametersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceParametersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "ParamList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceParametersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceParametersResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceParametersResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceParametersResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceParametersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceReadOnlyGroupRequestParams struct {
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// ID of the RO group to which the read-only replica belongs
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`

	// ID of the new RO group into which the read-only replica will move
	NewReadOnlyGroupId *string `json:"NewReadOnlyGroupId,omitnil,omitempty" name:"NewReadOnlyGroupId"`
}

type ModifyDBInstanceReadOnlyGroupRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// ID of the RO group to which the read-only replica belongs
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`

	// ID of the new RO group into which the read-only replica will move
	NewReadOnlyGroupId *string `json:"NewReadOnlyGroupId,omitnil,omitempty" name:"NewReadOnlyGroupId"`
}

func (r *ModifyDBInstanceReadOnlyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceReadOnlyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "ReadOnlyGroupId")
	delete(f, "NewReadOnlyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceReadOnlyGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceReadOnlyGroupResponseParams struct {
	// Task ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceReadOnlyGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceReadOnlyGroupResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceReadOnlyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceReadOnlyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupsRequestParams struct {
	// The list of security groups to be associated with the instance or RO groups.
	// Information of security groups can be obtained from the `sgld` field in the returned value of the [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/api/215/15808?from_cn_redirect=1) API.
	SecurityGroupIdSet []*string `json:"SecurityGroupIdSet,omitnil,omitempty" name:"SecurityGroupIdSet"`

	// Instance ID. Either this parameter or `ReadOnlyGroupId` must be passed in. If both parameters are passed in, `ReadOnlyGroupId` will be ignored.
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// RO group ID. Either this parameter or `DBInstanceId` must be passed in. To modify  the security groups associated with the RO groups, only pass in `ReadOnlyGroupId`.
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// The list of security groups to be associated with the instance or RO groups.
	// Information of security groups can be obtained from the `sgld` field in the returned value of the [DescribeSecurityGroups](https://intl.cloud.tencent.com/document/api/215/15808?from_cn_redirect=1) API.
	SecurityGroupIdSet []*string `json:"SecurityGroupIdSet,omitnil,omitempty" name:"SecurityGroupIdSet"`

	// Instance ID. Either this parameter or `ReadOnlyGroupId` must be passed in. If both parameters are passed in, `ReadOnlyGroupId` will be ignored.
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// RO group ID. Either this parameter or `DBInstanceId` must be passed in. To modify  the security groups associated with the RO groups, only pass in `ReadOnlyGroupId`.
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`
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
	delete(f, "SecurityGroupIdSet")
	delete(f, "DBInstanceId")
	delete(f, "ReadOnlyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyDBInstanceSpecRequestParams struct {
	// Instance ID in the format of postgres-6bwgamo3.
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Instance memory size in GiB after modification.
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Instance disk size in GiB after modification.
	Storage *uint64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Whether to use vouchers automatically. Valid values:
	// <li>`0`: No.
	// <li>`1`: Yes.
	// Default value: `0`.
	AutoVoucher *uint64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Voucher ID list. Currently, you can specify only one voucher.
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// Campaign ID.
	ActivityId *uint64 `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// Switch time for the specified instance after configuration modification.
	// <li>`0`: Switch now. 
	// <li>`1`: Switch at the specified time.
	// <li>`2`: Switch in the maintenance time.
	// Default value: `0`. 
	SwitchTag *uint64 `json:"SwitchTag,omitnil,omitempty" name:"SwitchTag"`

	// Switch start time in the format of `HH:MM:SS`, such as 01:00:00. When `SwitchTag` is 0 or 2, this parameter becomes invalid.
	SwitchStartTime *string `json:"SwitchStartTime,omitnil,omitempty" name:"SwitchStartTime"`

	// Switch end time in the format of `HH:MM:SS`, such as 01:30:00. When `SwitchTag` is 0 or 2, this parameter becomes invalid.
	SwitchEndTime *string `json:"SwitchEndTime,omitnil,omitempty" name:"SwitchEndTime"`
}

type ModifyDBInstanceSpecRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of postgres-6bwgamo3.
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Instance memory size in GiB after modification.
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Instance disk size in GiB after modification.
	Storage *uint64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Whether to use vouchers automatically. Valid values:
	// <li>`0`: No.
	// <li>`1`: Yes.
	// Default value: `0`.
	AutoVoucher *uint64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Voucher ID list. Currently, you can specify only one voucher.
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// Campaign ID.
	ActivityId *uint64 `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// Switch time for the specified instance after configuration modification.
	// <li>`0`: Switch now. 
	// <li>`1`: Switch at the specified time.
	// <li>`2`: Switch in the maintenance time.
	// Default value: `0`. 
	SwitchTag *uint64 `json:"SwitchTag,omitnil,omitempty" name:"SwitchTag"`

	// Switch start time in the format of `HH:MM:SS`, such as 01:00:00. When `SwitchTag` is 0 or 2, this parameter becomes invalid.
	SwitchStartTime *string `json:"SwitchStartTime,omitnil,omitempty" name:"SwitchStartTime"`

	// Switch end time in the format of `HH:MM:SS`, such as 01:30:00. When `SwitchTag` is 0 or 2, this parameter becomes invalid.
	SwitchEndTime *string `json:"SwitchEndTime,omitnil,omitempty" name:"SwitchEndTime"`
}

func (r *ModifyDBInstanceSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "Memory")
	delete(f, "Storage")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "ActivityId")
	delete(f, "SwitchTag")
	delete(f, "SwitchStartTime")
	delete(f, "SwitchEndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceSpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSpecResponseParams struct {
	// Order ID.
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// Bill ID of frozen fees.
	BillId *string `json:"BillId,omitnil,omitempty" name:"BillId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceSpecResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceSpecResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstancesProjectRequestParams struct {
	// List of instance IDs. Note that currently you cannot manipulate multiple instances at the same time. Only one instance ID can be passed in here.
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitnil,omitempty" name:"DBInstanceIdSet"`

	// ID of the new project
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type ModifyDBInstancesProjectRequest struct {
	*tchttp.BaseRequest
	
	// List of instance IDs. Note that currently you cannot manipulate multiple instances at the same time. Only one instance ID can be passed in here.
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitnil,omitempty" name:"DBInstanceIdSet"`

	// ID of the new project
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	delete(f, "DBInstanceIdSet")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstancesProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstancesProjectResponseParams struct {
	// Number of successfully transferred instances
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyParameterTemplateRequestParams struct {
	// Parameter template ID, which uniquely identifies a parameter template and cannot be modified.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Parameter template name, which can contain 1-60 letters, digits, and symbols (-_./()[]()+=:@). If this field is empty, the original parameter template name will be used.
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// Parameter template description, which can contain 1-60 letters, digits, and symbols (-_./()[]()+=:@). If this parameter is not passed in, the original parameter template description will be used.
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`

	// The set of parameters to be modified or added. A parameter cannot be put to `ModifyParamEntrySet` and `DeleteParamSet` at the same time, that is, it cannot be modified/added and deleted at the same time.
	ModifyParamEntrySet []*ParamEntry `json:"ModifyParamEntrySet,omitnil,omitempty" name:"ModifyParamEntrySet"`

	// The set of parameters to be deleted in the template. A parameter cannot be put to `ModifyParamEntrySet` and `DeleteParamSet` at the same time, that is, it cannot be modified/added and deleted at the same time.
	DeleteParamSet []*string `json:"DeleteParamSet,omitnil,omitempty" name:"DeleteParamSet"`
}

type ModifyParameterTemplateRequest struct {
	*tchttp.BaseRequest
	
	// Parameter template ID, which uniquely identifies a parameter template and cannot be modified.
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Parameter template name, which can contain 1-60 letters, digits, and symbols (-_./()[]()+=:@). If this field is empty, the original parameter template name will be used.
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// Parameter template description, which can contain 1-60 letters, digits, and symbols (-_./()[]()+=:@). If this parameter is not passed in, the original parameter template description will be used.
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`

	// The set of parameters to be modified or added. A parameter cannot be put to `ModifyParamEntrySet` and `DeleteParamSet` at the same time, that is, it cannot be modified/added and deleted at the same time.
	ModifyParamEntrySet []*ParamEntry `json:"ModifyParamEntrySet,omitnil,omitempty" name:"ModifyParamEntrySet"`

	// The set of parameters to be deleted in the template. A parameter cannot be put to `ModifyParamEntrySet` and `DeleteParamSet` at the same time, that is, it cannot be modified/added and deleted at the same time.
	DeleteParamSet []*string `json:"DeleteParamSet,omitnil,omitempty" name:"DeleteParamSet"`
}

func (r *ModifyParameterTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyParameterTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "TemplateName")
	delete(f, "TemplateDescription")
	delete(f, "ModifyParamEntrySet")
	delete(f, "DeleteParamSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyParameterTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyParameterTemplateResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyParameterTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyParameterTemplateResponseParams `json:"Response"`
}

func (r *ModifyParameterTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyParameterTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyReadOnlyGroupConfigRequestParams struct {
	// RO group ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`

	// RO group name
	ReadOnlyGroupName *string `json:"ReadOnlyGroupName,omitnil,omitempty" name:"ReadOnlyGroupName"`

	// Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: `0` (no), `1` (yes).
	ReplayLagEliminate *uint64 `json:"ReplayLagEliminate,omitnil,omitempty" name:"ReplayLagEliminate"`

	// Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: `0` (no), `1` (yes).
	ReplayLatencyEliminate *uint64 `json:"ReplayLatencyEliminate,omitnil,omitempty" name:"ReplayLatencyEliminate"`

	// Delayed log size threshold in MB
	MaxReplayLatency *uint64 `json:"MaxReplayLatency,omitnil,omitempty" name:"MaxReplayLatency"`

	// Delay threshold in ms
	MaxReplayLag *uint64 `json:"MaxReplayLag,omitnil,omitempty" name:"MaxReplayLag"`

	// Whether to enable automatic load balancing. Valid values: `0` (disable), `1` (enable).
	Rebalance *uint64 `json:"Rebalance,omitnil,omitempty" name:"Rebalance"`

	// The minimum number of read-only replicas that must be retained in an RO group
	MinDelayEliminateReserve *uint64 `json:"MinDelayEliminateReserve,omitnil,omitempty" name:"MinDelayEliminateReserve"`
}

type ModifyReadOnlyGroupConfigRequest struct {
	*tchttp.BaseRequest
	
	// RO group ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`

	// RO group name
	ReadOnlyGroupName *string `json:"ReadOnlyGroupName,omitnil,omitempty" name:"ReadOnlyGroupName"`

	// Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: `0` (no), `1` (yes).
	ReplayLagEliminate *uint64 `json:"ReplayLagEliminate,omitnil,omitempty" name:"ReplayLagEliminate"`

	// Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: `0` (no), `1` (yes).
	ReplayLatencyEliminate *uint64 `json:"ReplayLatencyEliminate,omitnil,omitempty" name:"ReplayLatencyEliminate"`

	// Delayed log size threshold in MB
	MaxReplayLatency *uint64 `json:"MaxReplayLatency,omitnil,omitempty" name:"MaxReplayLatency"`

	// Delay threshold in ms
	MaxReplayLag *uint64 `json:"MaxReplayLag,omitnil,omitempty" name:"MaxReplayLag"`

	// Whether to enable automatic load balancing. Valid values: `0` (disable), `1` (enable).
	Rebalance *uint64 `json:"Rebalance,omitnil,omitempty" name:"Rebalance"`

	// The minimum number of read-only replicas that must be retained in an RO group
	MinDelayEliminateReserve *uint64 `json:"MinDelayEliminateReserve,omitnil,omitempty" name:"MinDelayEliminateReserve"`
}

func (r *ModifyReadOnlyGroupConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyReadOnlyGroupConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReadOnlyGroupId")
	delete(f, "ReadOnlyGroupName")
	delete(f, "ReplayLagEliminate")
	delete(f, "ReplayLatencyEliminate")
	delete(f, "MaxReplayLatency")
	delete(f, "MaxReplayLag")
	delete(f, "Rebalance")
	delete(f, "MinDelayEliminateReserve")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyReadOnlyGroupConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyReadOnlyGroupConfigResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyReadOnlyGroupConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyReadOnlyGroupConfigResponseParams `json:"Response"`
}

func (r *ModifyReadOnlyGroupConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyReadOnlyGroupConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySwitchTimePeriodRequestParams struct {
	// The ID of the instance waiting for a switch
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Valid value: `0` (switch immediately)
	SwitchTag *uint64 `json:"SwitchTag,omitnil,omitempty" name:"SwitchTag"`
}

type ModifySwitchTimePeriodRequest struct {
	*tchttp.BaseRequest
	
	// The ID of the instance waiting for a switch
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Valid value: `0` (switch immediately)
	SwitchTag *uint64 `json:"SwitchTag,omitnil,omitempty" name:"SwitchTag"`
}

func (r *ModifySwitchTimePeriodRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySwitchTimePeriodRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "SwitchTag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySwitchTimePeriodRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySwitchTimePeriodResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySwitchTimePeriodResponse struct {
	*tchttp.BaseResponse
	Response *ModifySwitchTimePeriodResponseParams `json:"Response"`
}

func (r *ModifySwitchTimePeriodResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySwitchTimePeriodResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetworkAccess struct {
	// Network resource ID, instance ID, or RO group ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// Resource type. Valid values: `1` (instance), `2` (RO group)
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ResourceType *uint64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// VPC ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// IPv4 address
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// IPv6 address
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Vip6 *string `json:"Vip6,omitnil,omitempty" name:"Vip6"`

	// Access port
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// Subnet ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Network status. Valid values: `1` (applying), `2` (in use), `3` (deleting), `4` (deleted)
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	VpcStatus *int64 `json:"VpcStatus,omitnil,omitempty" name:"VpcStatus"`
}

type NormalQueryItem struct {
	// Username
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Number of calls
	Calls *int64 `json:"Calls,omitnil,omitempty" name:"Calls"`

	// Granularity
	CallsGrids []*int64 `json:"CallsGrids,omitnil,omitempty" name:"CallsGrids"`

	// Total time consumed
	CostTime *float64 `json:"CostTime,omitnil,omitempty" name:"CostTime"`

	// Number of affected rows
	Rows *int64 `json:"Rows,omitnil,omitempty" name:"Rows"`

	// Minimum time consumed
	MinCostTime *float64 `json:"MinCostTime,omitnil,omitempty" name:"MinCostTime"`

	// Maximum time consumed
	MaxCostTime *float64 `json:"MaxCostTime,omitnil,omitempty" name:"MaxCostTime"`

	// Time of the earliest slow SQL statement
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// Time of the latest slow SQL statement
	LastTime *string `json:"LastTime,omitnil,omitempty" name:"LastTime"`

	// Shared memory blocks for reads
	SharedReadBlks *int64 `json:"SharedReadBlks,omitnil,omitempty" name:"SharedReadBlks"`

	// Shared memory blocks for writes
	SharedWriteBlks *int64 `json:"SharedWriteBlks,omitnil,omitempty" name:"SharedWriteBlks"`

	// Total IO read time
	ReadCostTime *int64 `json:"ReadCostTime,omitnil,omitempty" name:"ReadCostTime"`

	// Total IO write time
	WriteCostTime *int64 `json:"WriteCostTime,omitnil,omitempty" name:"WriteCostTime"`

	// Database name
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// Slow SQL statement after desensitization
	NormalQuery *string `json:"NormalQuery,omitnil,omitempty" name:"NormalQuery"`
}

// Predefined struct for user
type OpenDBExtranetAccessRequestParams struct {
	// Instance ID in the format of postgres-hez4fh0v
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Whether to enable public network access over IPv6 address. Valid values: 1 (yes), 0 (no)
	IsIpv6 *int64 `json:"IsIpv6,omitnil,omitempty" name:"IsIpv6"`
}

type OpenDBExtranetAccessRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of postgres-hez4fh0v
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Whether to enable public network access over IPv6 address. Valid values: 1 (yes), 0 (no)
	IsIpv6 *int64 `json:"IsIpv6,omitnil,omitempty" name:"IsIpv6"`
}

func (r *OpenDBExtranetAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenDBExtranetAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "IsIpv6")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenDBExtranetAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenDBExtranetAccessResponseParams struct {
	// Async task flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenDBExtranetAccessResponse struct {
	*tchttp.BaseResponse
	Response *OpenDBExtranetAccessResponseParams `json:"Response"`
}

func (r *OpenDBExtranetAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenDBExtranetAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenServerlessDBExtranetAccessRequestParams struct {
	// Unique ID of an instance
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Instance name
	DBInstanceName *string `json:"DBInstanceName,omitnil,omitempty" name:"DBInstanceName"`
}

type OpenServerlessDBExtranetAccessRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID of an instance
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Instance name
	DBInstanceName *string `json:"DBInstanceName,omitnil,omitempty" name:"DBInstanceName"`
}

func (r *OpenServerlessDBExtranetAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenServerlessDBExtranetAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "DBInstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenServerlessDBExtranetAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenServerlessDBExtranetAccessResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenServerlessDBExtranetAccessResponse struct {
	*tchttp.BaseResponse
	Response *OpenServerlessDBExtranetAccessResponseParams `json:"Response"`
}

func (r *OpenServerlessDBExtranetAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenServerlessDBExtranetAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParamEntry struct {
	// Parameter name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// The new value to which the parameter will be modified. When this parameter is used as an input parameter, its value must be a string, such as `0.1` (decimal), `1000` (integer), and `replica` (enum).
	ExpectedValue *string `json:"ExpectedValue,omitnil,omitempty" name:"ExpectedValue"`
}

type ParamInfo struct {
	// Parameter ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// Parameter name
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Value type of the parameter. Valid values: `integer`, `real` (floating-point), `bool`, `enum`, `mutil_enum` (this type of parameter can be set to multiple enumerated values).
	// For an `integer` or `real` parameter, the `Min` field represents the minimum value and the `Max` field the maximum value. 
	// For a `bool` parameter, the valid values include `true` and `false`; 
	// For an `enum` or `mutil_enum` parameter, the `EnumValue` field represents the valid values.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ParamValueType *string `json:"ParamValueType,omitnil,omitempty" name:"ParamValueType"`

	// Unit of the parameter value. If the parameter has no unit, this field will return null.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// Default value of the parameter, which is returned as a string
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// Current value of the parameter, which is returned as a string
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// The maximum value of the `integer` or `real` parameter
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Max *float64 `json:"Max,omitnil,omitempty" name:"Max"`

	// Value range of the enum parameter
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	EnumValue []*string `json:"EnumValue,omitnil,omitempty" name:"EnumValue"`

	// The minimum value of the `integer` or `real` parameter
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Min *float64 `json:"Min,omitnil,omitempty" name:"Min"`

	// Parameter description in Chinese
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ParamDescriptionCH *string `json:"ParamDescriptionCH,omitnil,omitempty" name:"ParamDescriptionCH"`

	// Parameter description in English
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ParamDescriptionEN *string `json:"ParamDescriptionEN,omitnil,omitempty" name:"ParamDescriptionEN"`

	// Whether to restart the instance for the modified parameter to take effect. Valid values: `true` (yes), `false` (no)
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	NeedReboot *bool `json:"NeedReboot,omitnil,omitempty" name:"NeedReboot"`

	// Parameter category in Chinese
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ClassificationCN *string `json:"ClassificationCN,omitnil,omitempty" name:"ClassificationCN"`

	// Parameter category in English
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ClassificationEN *string `json:"ClassificationEN,omitnil,omitempty" name:"ClassificationEN"`

	// Whether the parameter is related to specifications. Valid values: `true` (yes), `false` (no)
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SpecRelated *bool `json:"SpecRelated,omitnil,omitempty" name:"SpecRelated"`

	// Whether it is a key parameter. Valid values: `true` (yes, and modifying it may affect instance performance), `false` (no)
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Advanced *bool `json:"Advanced,omitnil,omitempty" name:"Advanced"`

	// The last modified time of the parameter
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	LastModifyTime *string `json:"LastModifyTime,omitnil,omitempty" name:"LastModifyTime"`

	// Primary-standby constraint. Valid values: `0` (no constraint), `1` (The parameter value of the standby server must be greater than that of the primary server), `2` (The parameter value of the primary server must be greater than that of the standby server.)
	// Note: This field may return null, indicating that no valid values can be obtained.
	StandbyRelated *int64 `json:"StandbyRelated,omitnil,omitempty" name:"StandbyRelated"`

	// Associated parameter version information, which refers to the detailed parameter information of the kernel version.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VersionRelationSet []*ParamVersionRelation `json:"VersionRelationSet,omitnil,omitempty" name:"VersionRelationSet"`

	// Associated parameter specification information, which refers to the detailed parameter information of the specifications.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SpecRelationSet []*ParamSpecRelation `json:"SpecRelationSet,omitnil,omitempty" name:"SpecRelationSet"`
}

type ParamSpecRelation struct {
	// Parameter name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// The specification that corresponds to the parameter information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Memory *string `json:"Memory,omitnil,omitempty" name:"Memory"`

	// The default parameter value under this specification
	// Note: This field may return null, indicating that no valid values can be obtained.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// Unit of the parameter value. If the parameter has no unit, this field will return null.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// The maximum value of the `integer` or `real` parameter
	// Note: This field may return null, indicating that no valid values can be obtained.
	Max *float64 `json:"Max,omitnil,omitempty" name:"Max"`

	// The minimum value of the `integer` or `real` parameter
	// Note: This field may return null, indicating that no valid values can be obtained.
	Min *float64 `json:"Min,omitnil,omitempty" name:"Min"`

	// Value range of the enum parameter
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnumValue []*string `json:"EnumValue,omitnil,omitempty" name:"EnumValue"`
}

type ParamVersionRelation struct {
	// Parameter name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// The kernel version that corresponds to the parameter information
	// Note: This field may return null, indicating that no valid values can be obtained.
	DBKernelVersion *string `json:"DBKernelVersion,omitnil,omitempty" name:"DBKernelVersion"`

	// Default parameter value under the kernel version and specification of the instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// Unit of the parameter value. If the parameter has no unit, this field will return null.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// The maximum value of the `integer` or `real` parameter
	// Note: This field may return null, indicating that no valid values can be obtained.
	Max *float64 `json:"Max,omitnil,omitempty" name:"Max"`

	// The minimum value of the `integer` or `real` parameter
	// Note: This field may return null, indicating that no valid values can be obtained.
	Min *float64 `json:"Min,omitnil,omitempty" name:"Min"`

	// Value range of the enum parameter
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnumValue []*string `json:"EnumValue,omitnil,omitempty" name:"EnumValue"`
}

type ParameterTemplate struct {
	// Parameter template ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// Parameter template name
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// Database version applicable to a parameter template
	DBMajorVersion *string `json:"DBMajorVersion,omitnil,omitempty" name:"DBMajorVersion"`

	// Database engine applicable to a parameter template
	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`

	// Parameter template description
	TemplateDescription *string `json:"TemplateDescription,omitnil,omitempty" name:"TemplateDescription"`
}

type PgDeal struct {
	// Order name
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// User
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// Number of instances involved in order
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Billing mode. 0: pay-as-you-go
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Async task flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Instance ID array
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitnil,omitempty" name:"DBInstanceIdSet"`
}

type PolicyRule struct {
	// Policy, Valid values: `ACCEPT`, `DROP`.
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// Source or destination IP or IP range, such as 172.16.0.0/12.
	CidrIp *string `json:"CidrIp,omitnil,omitempty" name:"CidrIp"`

	// Port
	PortRange *string `json:"PortRange,omitnil,omitempty" name:"PortRange"`

	// Network protocol. UDP and TCP are supported.
	IpProtocol *string `json:"IpProtocol,omitnil,omitempty" name:"IpProtocol"`

	// The rule description
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type RawSlowQuery struct {
	// Slow query statement
	RawQuery *string `json:"RawQuery,omitnil,omitempty" name:"RawQuery"`

	// The database queried by the slow query statement
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// The execution time of the slow query statement
	Duration *float64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// The client that executes the slow query statement
	ClientAddr *string `json:"ClientAddr,omitnil,omitempty" name:"ClientAddr"`

	// The name of the user who executes the slow query statement
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// The time when the slow query statement starts to execute
	SessionStartTime *string `json:"SessionStartTime,omitnil,omitempty" name:"SessionStartTime"`
}

type ReadOnlyGroup struct {
	// RO group identifier
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`

	// RO group name
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ReadOnlyGroupName *string `json:"ReadOnlyGroupName,omitnil,omitempty" name:"ReadOnlyGroupName"`

	// Project ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Primary instance ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MasterDBInstanceId *string `json:"MasterDBInstanceId,omitnil,omitempty" name:"MasterDBInstanceId"`

	// The minimum number of read-only replicas that must be retained in an RO group
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MinDelayEliminateReserve *int64 `json:"MinDelayEliminateReserve,omitnil,omitempty" name:"MinDelayEliminateReserve"`

	// Delayed log size threshold
	MaxReplayLatency *int64 `json:"MaxReplayLatency,omitnil,omitempty" name:"MaxReplayLatency"`

	// Whether to remove a read-only replica from an RO group if the sync log size difference between the read-only replica and the primary instance exceeds the threshold. Valid values: `0` (no), `1` (yes).
	ReplayLatencyEliminate *int64 `json:"ReplayLatencyEliminate,omitnil,omitempty" name:"ReplayLatencyEliminate"`

	// Delay threshold
	MaxReplayLag *float64 `json:"MaxReplayLag,omitnil,omitempty" name:"MaxReplayLag"`

	// Whether to remove a read-only replica from an RO group if the delay between the read-only replica and the primary instance exceeds the threshold. Valid values: `0` (no), `1` (yes).
	ReplayLagEliminate *int64 `json:"ReplayLagEliminate,omitnil,omitempty" name:"ReplayLagEliminate"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Region ID
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Availability zone ID
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Status
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Instance details
	ReadOnlyDBInstanceList []*DBInstance `json:"ReadOnlyDBInstanceList,omitnil,omitempty" name:"ReadOnlyDBInstanceList"`

	// Whether to enable automatic load balancing
	Rebalance *int64 `json:"Rebalance,omitnil,omitempty" name:"Rebalance"`

	// Network information
	DBInstanceNetInfo []*DBInstanceNetInfo `json:"DBInstanceNetInfo,omitnil,omitempty" name:"DBInstanceNetInfo"`

	// Network access list of the RO group (this field has been deprecated)
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	NetworkAccessList []*NetworkAccess `json:"NetworkAccessList,omitnil,omitempty" name:"NetworkAccessList"`
}

// Predefined struct for user
type RebalanceReadOnlyGroupRequestParams struct {
	// RO group ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`
}

type RebalanceReadOnlyGroupRequest struct {
	*tchttp.BaseRequest
	
	// RO group ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`
}

func (r *RebalanceReadOnlyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RebalanceReadOnlyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReadOnlyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RebalanceReadOnlyGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RebalanceReadOnlyGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RebalanceReadOnlyGroupResponse struct {
	*tchttp.BaseResponse
	Response *RebalanceReadOnlyGroupResponseParams `json:"Response"`
}

func (r *RebalanceReadOnlyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RebalanceReadOnlyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {
	// Region abbreviation
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Region name
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// Region number
	RegionId *uint64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Availability status. UNAVAILABLE: unavailable, AVAILABLE: available
	RegionState *string `json:"RegionState,omitnil,omitempty" name:"RegionState"`

	// Whether the resource can be purchased in this region. Valid values: `0` (no), `1` (yes).
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SupportInternational *uint64 `json:"SupportInternational,omitnil,omitempty" name:"SupportInternational"`
}

// Predefined struct for user
type RemoveDBInstanceFromReadOnlyGroupRequestParams struct {
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// RO group ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`
}

type RemoveDBInstanceFromReadOnlyGroupRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// RO group ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitnil,omitempty" name:"ReadOnlyGroupId"`
}

func (r *RemoveDBInstanceFromReadOnlyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveDBInstanceFromReadOnlyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "ReadOnlyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveDBInstanceFromReadOnlyGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveDBInstanceFromReadOnlyGroupResponseParams struct {
	// Task ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveDBInstanceFromReadOnlyGroupResponse struct {
	*tchttp.BaseResponse
	Response *RemoveDBInstanceFromReadOnlyGroupResponseParams `json:"Response"`
}

func (r *RemoveDBInstanceFromReadOnlyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveDBInstanceFromReadOnlyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewInstanceRequestParams struct {
	// Instance ID in the format of `postgres-6fego161`
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Renewal duration in months
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Whether to automatically use vouchers. 1: yes, 0: no. Default value: 0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Voucher ID list (only one voucher can be specified currently)
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`
}

type RenewInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of `postgres-6fego161`
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Renewal duration in months
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// Whether to automatically use vouchers. 1: yes, 0: no. Default value: 0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Voucher ID list (only one voucher can be specified currently)
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`
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
	delete(f, "DBInstanceId")
	delete(f, "Period")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewInstanceResponseParams struct {
	// Order name
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type ResetAccountPasswordRequestParams struct {
	// Instance ID in the format of postgres-4wdeb0zv
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Instance account name
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// New password corresponding to `UserName` account
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type ResetAccountPasswordRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of postgres-4wdeb0zv
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Instance account name
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// New password corresponding to `UserName` account
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
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
	delete(f, "DBInstanceId")
	delete(f, "UserName")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetAccountPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetAccountPasswordResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type RestartDBInstanceRequestParams struct {
	// Instance ID in the format of postgres-6r233v55
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

type RestartDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID in the format of postgres-6r233v55
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`
}

func (r *RestartDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartDBInstanceResponseParams struct {
	// Async flow ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RestartDBInstanceResponseParams `json:"Response"`
}

func (r *RestartDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroup struct {
	// Project ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Inbound rule
	Inbound []*PolicyRule `json:"Inbound,omitnil,omitempty" name:"Inbound"`

	// Outbound rule
	Outbound []*PolicyRule `json:"Outbound,omitnil,omitempty" name:"Outbound"`

	// Security group ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// Security group name
	SecurityGroupName *string `json:"SecurityGroupName,omitnil,omitempty" name:"SecurityGroupName"`

	// Security group remarks
	SecurityGroupDescription *string `json:"SecurityGroupDescription,omitnil,omitempty" name:"SecurityGroupDescription"`
}

type ServerlessDBAccount struct {
	// Username
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DBUser *string `json:"DBUser,omitnil,omitempty" name:"DBUser"`

	// Password
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DBPassword *string `json:"DBPassword,omitnil,omitempty" name:"DBPassword"`

	// The maximum number of connections
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DBConnLimit *int64 `json:"DBConnLimit,omitnil,omitempty" name:"DBConnLimit"`
}

type ServerlessDBInstance struct {
	// Instance ID, which is the unique identifier
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Instance name
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DBInstanceName *string `json:"DBInstanceName,omitnil,omitempty" name:"DBInstanceName"`

	// Instance status
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DBInstanceStatus *string `json:"DBInstanceStatus,omitnil,omitempty" name:"DBInstanceStatus"`

	// Region
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Availability zone
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Project ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// VPC ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Character set
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DBCharset *string `json:"DBCharset,omitnil,omitempty" name:"DBCharset"`

	// Database version
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`

	// Creation time
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Instance network information
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DBInstanceNetInfo []*ServerlessDBInstanceNetInfo `json:"DBInstanceNetInfo,omitnil,omitempty" name:"DBInstanceNetInfo"`

	// Instance account information
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DBAccountSet []*ServerlessDBAccount `json:"DBAccountSet,omitnil,omitempty" name:"DBAccountSet"`

	// Information of the databases in an instance
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DBDatabaseList []*string `json:"DBDatabaseList,omitnil,omitempty" name:"DBDatabaseList"`

	// The array of tags bound to an instance
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`

	// Database kernel version
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DBKernelVersion *string `json:"DBKernelVersion,omitnil,omitempty" name:"DBKernelVersion"`

	// Database major version number
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	DBMajorVersion *string `json:"DBMajorVersion,omitnil,omitempty" name:"DBMajorVersion"`
}

type ServerlessDBInstanceNetInfo struct {
	// Address
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// IP address
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Port number
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Status
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Network type
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`
}

// Predefined struct for user
type SetAutoRenewFlagRequestParams struct {
	// List of instance IDs. Note that currently you cannot manipulate multiple instances at the same time. Only one instance ID can be passed in here.
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitnil,omitempty" name:"DBInstanceIdSet"`

	// Renewal flag. 0: normal renewal, 1: auto-renewal, 2: no renewal upon expiration
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`
}

type SetAutoRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// List of instance IDs. Note that currently you cannot manipulate multiple instances at the same time. Only one instance ID can be passed in here.
	DBInstanceIdSet []*string `json:"DBInstanceIdSet,omitnil,omitempty" name:"DBInstanceIdSet"`

	// Renewal flag. 0: normal renewal, 1: auto-renewal, 2: no renewal upon expiration
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`
}

func (r *SetAutoRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetAutoRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceIdSet")
	delete(f, "AutoRenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetAutoRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetAutoRenewFlagResponseParams struct {
	// Number of successfully set instances
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetAutoRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *SetAutoRenewFlagResponseParams `json:"Response"`
}

func (r *SetAutoRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetAutoRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SlowlogDetail struct {
	// Total time consumed
	TotalTime *float64 `json:"TotalTime,omitnil,omitempty" name:"TotalTime"`

	// Total number of calls
	TotalCalls *int64 `json:"TotalCalls,omitnil,omitempty" name:"TotalCalls"`

	// List of slow SQL statements after desensitization
	NormalQueries []*NormalQueryItem `json:"NormalQueries,omitnil,omitempty" name:"NormalQueries"`
}

type SpecInfo struct {
	// Region abbreviation, which corresponds to the `Region` field of `RegionSet`
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// AZ abbreviate, which corresponds to the `Zone` field of `ZoneSet`
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Specification details list
	SpecItemInfoList []*SpecItemInfo `json:"SpecItemInfoList,omitnil,omitempty" name:"SpecItemInfoList"`

	// Regions where KMS is supported
	// Note: This field may return `null`, indicating that no valid value was found.
	SupportKMSRegions []*string `json:"SupportKMSRegions,omitnil,omitempty" name:"SupportKMSRegions"`
}

type SpecItemInfo struct {
	// Specification ID
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// PostgerSQL version number
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// Full version name corresponding to kernel number
	VersionName *string `json:"VersionName,omitnil,omitempty" name:"VersionName"`

	// Number of CPU cores
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory size in MB
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Maximum storage capacity in GB supported by this specification
	MaxStorage *uint64 `json:"MaxStorage,omitnil,omitempty" name:"MaxStorage"`

	// Minimum storage capacity in GB supported by this specification
	MinStorage *uint64 `json:"MinStorage,omitnil,omitempty" name:"MinStorage"`

	// Estimated QPS for this specification
	Qps *uint64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// (Disused)
	Pid *uint64 `json:"Pid,omitnil,omitempty" name:"Pid"`

	// Machine type
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// PostgreSQL major version number
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MajorVersion *string `json:"MajorVersion,omitnil,omitempty" name:"MajorVersion"`

	// PostgreSQL kernel version number
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	KernelVersion *string `json:"KernelVersion,omitnil,omitempty" name:"KernelVersion"`

	// Whether TDE data encryption is supported. Valid values: 0 (no), 1 (yes)
	// Note: This field may return `null`, indicating that no valid value was found.
	IsSupportTDE *int64 `json:"IsSupportTDE,omitnil,omitempty" name:"IsSupportTDE"`
}

// Predefined struct for user
type SwitchDBInstancePrimaryRequestParams struct {
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Whether to perform forced switch. As long as the standby node can be accessed, the switch will be performed regardless of the primary-standby sync delay. You can switch immediately only when `SwitchTag` is `0.
	// <li>Default: `false`.
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`

	// Switch time for the specified instance after configuration modification.
	// <li>`0`: Switch now. 
	// <li>`1`: Switch at the specified time.
	// <li>`2`: Switch in the maintenance time.
	// <li>Default value: `0`. 
	SwitchTag *uint64 `json:"SwitchTag,omitnil,omitempty" name:"SwitchTag"`

	// The earliest time to start a switch in the format of "HH:MM:SS", such as "01:00:00". This parameter is invalid when `SwitchTag` is `0` or `2`.
	SwitchStartTime *string `json:"SwitchStartTime,omitnil,omitempty" name:"SwitchStartTime"`

	// The latest time to start a switch in the format of "HH:MM:SS", such as "01:30:00". This parameter is invalid when `SwitchTag` is `0` or `2`. The difference between `SwitchStartTime` and `SwitchEndTime` cannot be less than 30 minutes.
	SwitchEndTime *string `json:"SwitchEndTime,omitnil,omitempty" name:"SwitchEndTime"`
}

type SwitchDBInstancePrimaryRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Whether to perform forced switch. As long as the standby node can be accessed, the switch will be performed regardless of the primary-standby sync delay. You can switch immediately only when `SwitchTag` is `0.
	// <li>Default: `false`.
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`

	// Switch time for the specified instance after configuration modification.
	// <li>`0`: Switch now. 
	// <li>`1`: Switch at the specified time.
	// <li>`2`: Switch in the maintenance time.
	// <li>Default value: `0`. 
	SwitchTag *uint64 `json:"SwitchTag,omitnil,omitempty" name:"SwitchTag"`

	// The earliest time to start a switch in the format of "HH:MM:SS", such as "01:00:00". This parameter is invalid when `SwitchTag` is `0` or `2`.
	SwitchStartTime *string `json:"SwitchStartTime,omitnil,omitempty" name:"SwitchStartTime"`

	// The latest time to start a switch in the format of "HH:MM:SS", such as "01:30:00". This parameter is invalid when `SwitchTag` is `0` or `2`. The difference between `SwitchStartTime` and `SwitchEndTime` cannot be less than 30 minutes.
	SwitchEndTime *string `json:"SwitchEndTime,omitnil,omitempty" name:"SwitchEndTime"`
}

func (r *SwitchDBInstancePrimaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDBInstancePrimaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "Force")
	delete(f, "SwitchTag")
	delete(f, "SwitchStartTime")
	delete(f, "SwitchEndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchDBInstancePrimaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchDBInstancePrimaryResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SwitchDBInstancePrimaryResponse struct {
	*tchttp.BaseResponse
	Response *SwitchDBInstancePrimaryResponseParams `json:"Response"`
}

func (r *SwitchDBInstancePrimaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDBInstancePrimaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// Tag key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

// Predefined struct for user
type UpgradeDBInstanceKernelVersionRequestParams struct {
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Target kernel version, which can be obtained in the `AvailableUpgradeTarget` field in the returned value of the [DescribeDBVersions](https://intl.cloud.tencent.com/document/api/409/89018?from_cn_redirect=1) API.
	TargetDBKernelVersion *string `json:"TargetDBKernelVersion,omitnil,omitempty" name:"TargetDBKernelVersion"`

	// Switch time after the kernel version upgrade for the specified instance. Valid values:
	// <li>`0`: Switch now.
	// <li>`1`: Switch at the specified time.
	// <li>`2`: Switch in the maintenance time.
	// Default value: `0`. 
	SwitchTag *uint64 `json:"SwitchTag,omitnil,omitempty" name:"SwitchTag"`

	// Switch start time in the format of `HH:MM:SS`, such as 01:00:00. When `SwitchTag` is `0` or `2`, this parameter is invalid.
	SwitchStartTime *string `json:"SwitchStartTime,omitnil,omitempty" name:"SwitchStartTime"`

	// Switch end time in the format of `HH:MM:SS`, such as 01:30:00. When `SwitchTag` is `0` or `2`, this parameter is invalid. The difference between `SwitchStartTime` and `SwitchEndTime` cannot be less than 30 minutes.
	SwitchEndTime *string `json:"SwitchEndTime,omitnil,omitempty" name:"SwitchEndTime"`

	// Whether to perform a pre-check on the current operation of upgrading the instance kernel version. Valid values:
	// u200c<li>u200c`true`: Performs a pre-check without upgrading the kernel version. Check items include request parameters, kernel version compatibility, and instance parameters.
	// u200cu200c<li>`false`: Sends a normal request and upgrades the kernel version directly after the check is passed.
	// Default value: `false`.
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type UpgradeDBInstanceKernelVersionRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Target kernel version, which can be obtained in the `AvailableUpgradeTarget` field in the returned value of the [DescribeDBVersions](https://intl.cloud.tencent.com/document/api/409/89018?from_cn_redirect=1) API.
	TargetDBKernelVersion *string `json:"TargetDBKernelVersion,omitnil,omitempty" name:"TargetDBKernelVersion"`

	// Switch time after the kernel version upgrade for the specified instance. Valid values:
	// <li>`0`: Switch now.
	// <li>`1`: Switch at the specified time.
	// <li>`2`: Switch in the maintenance time.
	// Default value: `0`. 
	SwitchTag *uint64 `json:"SwitchTag,omitnil,omitempty" name:"SwitchTag"`

	// Switch start time in the format of `HH:MM:SS`, such as 01:00:00. When `SwitchTag` is `0` or `2`, this parameter is invalid.
	SwitchStartTime *string `json:"SwitchStartTime,omitnil,omitempty" name:"SwitchStartTime"`

	// Switch end time in the format of `HH:MM:SS`, such as 01:30:00. When `SwitchTag` is `0` or `2`, this parameter is invalid. The difference between `SwitchStartTime` and `SwitchEndTime` cannot be less than 30 minutes.
	SwitchEndTime *string `json:"SwitchEndTime,omitnil,omitempty" name:"SwitchEndTime"`

	// Whether to perform a pre-check on the current operation of upgrading the instance kernel version. Valid values:
	// u200c<li>u200c`true`: Performs a pre-check without upgrading the kernel version. Check items include request parameters, kernel version compatibility, and instance parameters.
	// u200cu200c<li>`false`: Sends a normal request and upgrades the kernel version directly after the check is passed.
	// Default value: `false`.
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *UpgradeDBInstanceKernelVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDBInstanceKernelVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBInstanceId")
	delete(f, "TargetDBKernelVersion")
	delete(f, "SwitchTag")
	delete(f, "SwitchStartTime")
	delete(f, "SwitchEndTime")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeDBInstanceKernelVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDBInstanceKernelVersionResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeDBInstanceKernelVersionResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeDBInstanceKernelVersionResponseParams `json:"Response"`
}

func (r *UpgradeDBInstanceKernelVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDBInstanceKernelVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDBInstanceRequestParams struct {
	// Instance memory size in GB after upgrade
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Instance disk size in GB after upgrade
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Instance ID in the format of postgres-lnp6j617
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Whether to automatically use vouchers. 1: yes, 0: no. Default value: no
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Voucher ID list (only one voucher can be specified currently)
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// Activity ID
	ActivityId *int64 `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// Switch time after instance configurations are modified. Valid values: `0` (switch immediately), `1` (switch at specified time). Default value: `0`
	SwitchTag *int64 `json:"SwitchTag,omitnil,omitempty" name:"SwitchTag"`

	// The earliest time to start a switch
	SwitchStartTime *string `json:"SwitchStartTime,omitnil,omitempty" name:"SwitchStartTime"`

	// The latest time to start a switch
	SwitchEndTime *string `json:"SwitchEndTime,omitnil,omitempty" name:"SwitchEndTime"`
}

type UpgradeDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance memory size in GB after upgrade
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Instance disk size in GB after upgrade
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// Instance ID in the format of postgres-lnp6j617
	DBInstanceId *string `json:"DBInstanceId,omitnil,omitempty" name:"DBInstanceId"`

	// Whether to automatically use vouchers. 1: yes, 0: no. Default value: no
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// Voucher ID list (only one voucher can be specified currently)
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// Activity ID
	ActivityId *int64 `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// Switch time after instance configurations are modified. Valid values: `0` (switch immediately), `1` (switch at specified time). Default value: `0`
	SwitchTag *int64 `json:"SwitchTag,omitnil,omitempty" name:"SwitchTag"`

	// The earliest time to start a switch
	SwitchStartTime *string `json:"SwitchStartTime,omitnil,omitempty" name:"SwitchStartTime"`

	// The latest time to start a switch
	SwitchEndTime *string `json:"SwitchEndTime,omitnil,omitempty" name:"SwitchEndTime"`
}

func (r *UpgradeDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Memory")
	delete(f, "Storage")
	delete(f, "DBInstanceId")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "ActivityId")
	delete(f, "SwitchTag")
	delete(f, "SwitchStartTime")
	delete(f, "SwitchEndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDBInstanceResponseParams struct {
	// Transaction name.
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// Bill ID of frozen fees
	BillId *string `json:"BillId,omitnil,omitempty" name:"BillId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeDBInstanceResponseParams `json:"Response"`
}

func (r *UpgradeDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Version struct {
	// Database engines. Valid values:
	// 1. `postgresql` (TencentDB for PostgreSQL)
	// 2. `mssql_compatible` (MSSQL compatible-TencentDB for PostgreSQL)
	DBEngine *string `json:"DBEngine,omitnil,omitempty" name:"DBEngine"`

	// Database version, such as 12.4.
	DBVersion *string `json:"DBVersion,omitnil,omitempty" name:"DBVersion"`

	// Database major version, such as 12.
	DBMajorVersion *string `json:"DBMajorVersion,omitnil,omitempty" name:"DBMajorVersion"`

	// Database kernel version, such as v12.4_r1.3.
	DBKernelVersion *string `json:"DBKernelVersion,omitnil,omitempty" name:"DBKernelVersion"`

	// List of features supported by the database kernel, such as:
	// TDE: Supports data encryption.
	SupportedFeatureNames []*string `json:"SupportedFeatureNames,omitnil,omitempty" name:"SupportedFeatureNames"`

	// Database version status. Valid values:
	// `AVAILABLE`.
	// `DEPRECATED`.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// List of versions to which this database version (`DBKernelVersion`) can be upgraded.
	AvailableUpgradeTarget []*string `json:"AvailableUpgradeTarget,omitnil,omitempty" name:"AvailableUpgradeTarget"`
}

type Xlog struct {
	// Unique backup file ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// File generation start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// File generation end time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Download address on private network
	InternalAddr *string `json:"InternalAddr,omitnil,omitempty" name:"InternalAddr"`

	// Download address on public network
	ExternalAddr *string `json:"ExternalAddr,omitnil,omitempty" name:"ExternalAddr"`

	// Backup file size
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`
}

type ZoneInfo struct {
	// AZ abbreviation
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// AZ name
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// AZ number
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Availability status. Valid values:
	// `UNAVAILABLE`.
	// `AVAILABLE`.
	// `SELLOUT`.
	// `SUPPORTMODIFYONLY` (supports configuration adjustment).
	ZoneState *string `json:"ZoneState,omitnil,omitempty" name:"ZoneState"`

	// Whether the AZ supports IPv6 address access
	ZoneSupportIpv6 *uint64 `json:"ZoneSupportIpv6,omitnil,omitempty" name:"ZoneSupportIpv6"`

	// AZs that can be used as standby when this AZ is primary
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	StandbyZoneSet []*string `json:"StandbyZoneSet,omitnil,omitempty" name:"StandbyZoneSet"`
}