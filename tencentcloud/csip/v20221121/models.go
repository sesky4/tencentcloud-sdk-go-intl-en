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

package v20221121

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

// Predefined struct for user
type AddNewBindRoleUserRequestParams struct {

}

type AddNewBindRoleUserRequest struct {
	*tchttp.BaseRequest
	
}

func (r *AddNewBindRoleUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddNewBindRoleUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddNewBindRoleUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddNewBindRoleUserResponseParams struct {
	// `0`: successful. Other values: failed.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddNewBindRoleUserResponse struct {
	*tchttp.BaseResponse
	Response *AddNewBindRoleUserResponseParams `json:"Response"`
}

func (r *AddNewBindRoleUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddNewBindRoleUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetBaseInfoResponse struct {

	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// vpc-name
	// Note: This field may return·null, indicating that no valid values can be obtained.
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// Asset name
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// Operating system
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// Public IP
	// Note: This field may return·null, indicating that no valid values can be obtained.
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// Private IP
	// Note: This field may return·null, indicating that no valid values can be obtained.
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// Region
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Asset type
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// Asset ID
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// Total number of accounts
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AccountNum *uint64 `json:"AccountNum,omitnil,omitempty" name:"AccountNum"`

	// Number of ports
	// Note: This field may return·null, indicating that no valid values can be obtained.
	PortNum *uint64 `json:"PortNum,omitnil,omitempty" name:"PortNum"`

	// Number of processes
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ProcessNum *uint64 `json:"ProcessNum,omitnil,omitempty" name:"ProcessNum"`

	// Numbernumb of software applications
	// Note: This field may return·null, indicating that no valid values can be obtained.
	SoftApplicationNum *uint64 `json:"SoftApplicationNum,omitnil,omitempty" name:"SoftApplicationNum"`

	// Number of databases
	// Note: This field may return·null, indicating that no valid values can be obtained.
	DatabaseNum *uint64 `json:"DatabaseNum,omitnil,omitempty" name:"DatabaseNum"`

	// Number of web applications
	// Note: This field may return·null, indicating that no valid values can be obtained.
	WebApplicationNum *uint64 `json:"WebApplicationNum,omitnil,omitempty" name:"WebApplicationNum"`

	// Number of services
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ServiceNum *uint64 `json:"ServiceNum,omitnil,omitempty" name:"ServiceNum"`

	// Number of web frameworks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	WebFrameworkNum *uint64 `json:"WebFrameworkNum,omitnil,omitempty" name:"WebFrameworkNum"`

	// Number of websites
	// Note: This field may return·null, indicating that no valid values can be obtained.
	WebSiteNum *uint64 `json:"WebSiteNum,omitnil,omitempty" name:"WebSiteNum"`

	// Number of JAR packages
	// Note: This field may return·null, indicating that no valid values can be obtained.
	JarPackageNum *uint64 `json:"JarPackageNum,omitnil,omitempty" name:"JarPackageNum"`

	// Number of enabled services
	// Note: This field may return·null, indicating that no valid values can be obtained.
	StartServiceNum *uint64 `json:"StartServiceNum,omitnil,omitempty" name:"StartServiceNum"`

	// Number of scheduled tasks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ScheduledTaskNum *uint64 `json:"ScheduledTaskNum,omitnil,omitempty" name:"ScheduledTaskNum"`

	// Number of environment variables
	// Note: This field may return·null, indicating that no valid values can be obtained.
	EnvironmentVariableNum *uint64 `json:"EnvironmentVariableNum,omitnil,omitempty" name:"EnvironmentVariableNum"`

	// Number of kernel modules
	// Note: This field may return·null, indicating that no valid values can be obtained.
	KernelModuleNum *uint64 `json:"KernelModuleNum,omitnil,omitempty" name:"KernelModuleNum"`

	// Number of system installation packages
	// Note: This field may return·null, indicating that no valid values can be obtained.
	SystemInstallationPackageNum *uint64 `json:"SystemInstallationPackageNum,omitnil,omitempty" name:"SystemInstallationPackageNum"`

	// Remaining service validity in days
	// Note: This field may return·null, indicating that no valid values can be obtained.
	SurplusProtectDay *uint64 `json:"SurplusProtectDay,omitnil,omitempty" name:"SurplusProtectDay"`

	// Whether the CWPP agent is installed. Values: `1` (installed) and `0` (not installed)
	// Note: This field may return·null, indicating that no valid values can be obtained.
	CWPStatus *uint64 `json:"CWPStatus,omitnil,omitempty" name:"CWPStatus"`

	// Tags
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Protection level
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ProtectLevel *string `json:"ProtectLevel,omitnil,omitempty" name:"ProtectLevel"`

	// Usage of CWPP service in days
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ProtectedDay *uint64 `json:"ProtectedDay,omitnil,omitempty" name:"ProtectedDay"`
}

type AssetClusterPod struct {
	// Tenant ID
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Tenant UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Tenant name
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// Region
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Pod ID
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// Pod name
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// Creation time of the pod
	// Note: This field may return·null, indicating that no valid values can be obtained.
	InstanceCreateTime *string `json:"InstanceCreateTime,omitnil,omitempty" name:"InstanceCreateTime"`

	// Namespace
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Status
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Cluster ID
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Cluster name
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// Server ID
	// Note: This field may return·null, indicating that no valid values can be obtained.
	MachineId *string `json:"MachineId,omitnil,omitempty" name:"MachineId"`

	// Server name
	// Note: This field may return·null, indicating that no valid values can be obtained.
	MachineName *string `json:"MachineName,omitnil,omitempty" name:"MachineName"`

	// Pod IP
	// Note: This field may return·null, indicating that no valid values can be obtained.
	PodIp *string `json:"PodIp,omitnil,omitempty" name:"PodIp"`

	// Number of associated services
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ServiceCount *int64 `json:"ServiceCount,omitnil,omitempty" name:"ServiceCount"`

	// Number of associated containers
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ContainerCount *int64 `json:"ContainerCount,omitnil,omitempty" name:"ContainerCount"`

	// Public IP
	// Note: This field may return·null, indicating that no valid values can be obtained.
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// Private IP
	// Note: This field may return·null, indicating that no valid values can be obtained.
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// Whether it's a critical asset. Values: `1` (critical asset), `0` (non-critical asset)
	// Note: This field may return·null, indicating that no valid values can be obtained.
	IsCore *int64 `json:"IsCore,omitnil,omitempty" name:"IsCore"`

	// Whether it's a newly-added asset. Values: `1` (Yes), `0` (No)
	// Note: This field may return·null, indicating that no valid values can be obtained.
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil,omitempty" name:"IsNewAsset"`
}

type AssetInfoDetail struct {
	// AppID of the user
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AppID *string `json:"AppID,omitnil,omitempty" name:"AppID"`

	// CVE number
	// Note: This field may return·null, indicating that no valid values can be obtained.
	CVEId *string `json:"CVEId,omitnil,omitempty" name:"CVEId"`

	// Whether the asset is scanned. Values: `0`: (default) Not scanned; `1`: Scanning; `2`: Scan completed; `3`: Error while scanning
	// Note: This field may return·null, indicating that no valid values can be obtained.
	IsScan *int64 `json:"IsScan,omitnil,omitempty" name:"IsScan"`

	// Number of affected assets
	// Note: This field may return·null, indicating that no valid values can be obtained.
	InfluenceAsset *int64 `json:"InfluenceAsset,omitnil,omitempty" name:"InfluenceAsset"`

	// Number of not fixed assets
	// Note: This field may return·null, indicating that no valid values can be obtained.
	NotRepairAsset *int64 `json:"NotRepairAsset,omitnil,omitempty" name:"NotRepairAsset"`

	// Number of not protected assets
	// Note: This field may return·null, indicating that no valid values can be obtained.
	NotProtectAsset *int64 `json:"NotProtectAsset,omitnil,omitempty" name:"NotProtectAsset"`

	// Task ID
	// Note: This field may return·null, indicating that no valid values can be obtained.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task progress in terms of percentage
	// Note: This field may return·null, indicating that no valid values can be obtained.
	TaskPercent *int64 `json:"TaskPercent,omitnil,omitempty" name:"TaskPercent"`

	// Task creation time
	// Note: This field may return·null, indicating that no valid values can be obtained.
	TaskTime *int64 `json:"TaskTime,omitnil,omitempty" name:"TaskTime"`

	// Scan start time
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ScanTime *string `json:"ScanTime,omitnil,omitempty" name:"ScanTime"`
}

type AssetTag struct {
	// Tag key. It supports alphanumeric characters and underscores (_).
	// Note: This field may return·null, indicating that no valid values can be obtained.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value. It supports alphanumeric characters and underscores (_).
	// Note: This field may return·null, indicating that no valid values can be obtained.
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type AssetViewCFGRisk struct {
	// The unique ID.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Configuration name
	CFGName *string `json:"CFGName,omitnil,omitempty" name:"CFGName"`

	// Check type
	CheckType *string `json:"CheckType,omitnil,omitempty" name:"CheckType"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Instance type
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Affected assets
	AffectAsset *string `json:"AffectAsset,omitnil,omitempty" name:"AffectAsset"`

	// Risk level
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// First detected
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// Last detected 
	RecentTime *string `json:"RecentTime,omitnil,omitempty" name:"RecentTime"`

	// Source of the task
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// Status
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// u200c-
	CFGSTD *string `json:"CFGSTD,omitnil,omitempty" name:"CFGSTD"`

	// Configuration details.
	CFGDescribe *string `json:"CFGDescribe,omitnil,omitempty" name:"CFGDescribe"`

	// Fix suggestion
	CFGFix *string `json:"CFGFix,omitnil,omitempty" name:"CFGFix"`

	// URL of the help documentation
	CFGHelpURL *string `json:"CFGHelpURL,omitnil,omitempty" name:"CFGHelpURL"`

	// Data entry key
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// User AppId
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User name.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// User UIN
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`
}

type AssetViewPortRisk struct {
	// Port
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Affected assets
	AffectAsset *string `json:"AffectAsset,omitnil,omitempty" name:"AffectAsset"`

	// Risk level
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// Asset type
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Network protocol
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Components
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// Service
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// Last detected 
	RecentTime *string `json:"RecentTime,omitnil,omitempty" name:"RecentTime"`

	// First detected
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// Suggested action. `0`: Keep as it is; `1`: Block access requests; `2`: Block the port
	Suggestion *uint64 `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// Status of the risk. `0`: Not handled, `1`: Handled; `2`: Ignored
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Unique ID of the asset
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Frontend index
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// User `appid`
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User name.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// User `uin`
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Source of the task
	From *string `json:"From,omitnil,omitempty" name:"From"`
}

type AssetViewVULRisk struct {
	// Affected assets
	AffectAsset *string `json:"AffectAsset,omitnil,omitempty" name:"AffectAsset"`

	// Risk level
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// Asset type
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Components
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// Service
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// Last detected 
	RecentTime *string `json:"RecentTime,omitnil,omitempty" name:"RecentTime"`

	// First detected
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// Status of the risk. `0`: Not handled, `1`: Handled; `2`: Ignored
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Unique ID of the asset
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Frontend index
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// User `appid`
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User name.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// User `uin`
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Vulnerability type
	VULType *string `json:"VULType,omitnil,omitempty" name:"VULType"`

	// Port
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// Description
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// Components affected by the vulnerability 
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// Reference information about the vulnerability
	References *string `json:"References,omitnil,omitempty" name:"References"`

	// Version
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// Vulnerability URL
	VULURL *string `json:"VULURL,omitnil,omitempty" name:"VULURL"`

	// Vulnerability name
	VULName *string `json:"VULName,omitnil,omitempty" name:"VULName"`

	// CVE number
	CVE *string `json:"CVE,omitnil,omitempty" name:"CVE"`

	// Fix suggestion
	Fix *string `json:"Fix,omitnil,omitempty" name:"Fix"`

	// POC ID
	POCId *string `json:"POCId,omitnil,omitempty" name:"POCId"`

	// Source of the task
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// CWPP edition
	CWPVersion *int64 `json:"CWPVersion,omitnil,omitempty" name:"CWPVersion"`

	// Whether it can be fixed 
	IsSupportRepair *bool `json:"IsSupportRepair,omitnil,omitempty" name:"IsSupportRepair"`

	// Whether it can be detected
	IsSupportDetect *bool `json:"IsSupportDetect,omitnil,omitempty" name:"IsSupportDetect"`

	// Instance UUID
	InstanceUUID *string `json:"InstanceUUID,omitnil,omitempty" name:"InstanceUUID"`

	// Pay load
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// Whether it's an emergency vulnerability. Values: `1` (emergency vulnerability); `0` (non-emergency vulnerability
	// Note: This field may return·null, indicating that no valid values can be obtained.
	EMGCVulType *int64 `json:"EMGCVulType,omitnil,omitempty" name:"EMGCVulType"`
}

type AssetViewWeakPassRisk struct {
	// Affected assets
	AffectAsset *string `json:"AffectAsset,omitnil,omitempty" name:"AffectAsset"`

	// Risk level
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// Asset type
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Components
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// Service
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// Last detected 
	RecentTime *string `json:"RecentTime,omitnil,omitempty" name:"RecentTime"`

	// First detected
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// Status of the risk. `0`: Not handled, `1`: Handled; `2`: Ignored
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Unique ID of the asset
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Frontend index
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// User AppId
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User name.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// User `uin`
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Weak password type
	PasswordType *string `json:"PasswordType,omitnil,omitempty" name:"PasswordType"`

	// Source of the task
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// Vulnerability type
	VULType *string `json:"VULType,omitnil,omitempty" name:"VULType"`

	// Vulnerability URL
	VULURL *string `json:"VULURL,omitnil,omitempty" name:"VULURL"`

	// Fix suggestion
	Fix *string `json:"Fix,omitnil,omitempty" name:"Fix"`

	// Pay load
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`
}

type BugInfoDetail struct {
	// Vulnerability ID
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// POC ID of the vulnerability
	// Note: This field may return·null, indicating that no valid values can be obtained.
	PatchId *string `json:"PatchId,omitnil,omitempty" name:"PatchId"`

	// Vulnerability name
	// Note: This field may return·null, indicating that no valid values can be obtained.
	VULName *string `json:"VULName,omitnil,omitempty" name:"VULName"`

	// Vulnerability severity: `high`, `middle`, `low`, `info`
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// CVSS score
	// Note: This field may return·null, indicating that no valid values can be obtained.
	CVSSScore *string `json:"CVSSScore,omitnil,omitempty" name:"CVSSScore"`

	// CVE number
	// Note: This field may return·null, indicating that no valid values can be obtained.
	CVEId *string `json:"CVEId,omitnil,omitempty" name:"CVEId"`

	// Vulnerability tag
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Vulnerability category: `1`: Web application vulnerabilities, `2`: System component vulnerabilities, `3`: Configuration risks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	VULCategory *uint64 `json:"VULCategory,omitnil,omitempty" name:"VULCategory"`

	// Operating systems affected by the vulnerability 
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ImpactOs *string `json:"ImpactOs,omitnil,omitempty" name:"ImpactOs"`

	// Components affected by the vulnerability 
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ImpactCOMPENT *string `json:"ImpactCOMPENT,omitnil,omitempty" name:"ImpactCOMPENT"`

	// Versions affected by the vulnerability
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ImpactVersion *string `json:"ImpactVersion,omitnil,omitempty" name:"ImpactVersion"`

	// Reference information of the vulnerability
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Reference *string `json:"Reference,omitnil,omitempty" name:"Reference"`

	// Vulnerability description
	// Note: This field may return·null, indicating that no valid values can be obtained.
	VULDescribe *string `json:"VULDescribe,omitnil,omitempty" name:"VULDescribe"`

	// Fix suggestion
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Fix *string `json:"Fix,omitnil,omitempty" name:"Fix"`

	// Product support status. The real-time status is returned.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ProSupport *uint64 `json:"ProSupport,omitnil,omitempty" name:"ProSupport"`

	// Specify whether the vulnerability is published as an emergency vulnerability. `1`: Published as an emergency vulnerability; `0`: Not an emergency vulnerability.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	IsPublish *uint64 `json:"IsPublish,omitnil,omitempty" name:"IsPublish"`

	// Disclosure time of the vulnerability. 
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ReleaseTime *string `json:"ReleaseTime,omitnil,omitempty" name:"ReleaseTime"`

	// The time when the vulnerability is added to the vulnerability database.
	// Note: u200dThis field may return `null`, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// The last update time of the vulnerability in the database
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Sub-category of the vulnerability
	// Note: This field may return·null, indicating that no valid values can be obtained.
	SubCategory *string `json:"SubCategory,omitnil,omitempty" name:"SubCategory"`
}

type CVMAssetVO struct {
	// Asset ID
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// Asset name
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// Asset type
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// Region
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Protection status
	// Note: This field may return·null, indicating that no valid values can be obtained.
	CWPStatus *uint64 `json:"CWPStatus,omitnil,omitempty" name:"CWPStatus"`

	// Asset creation time
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetCreateTime *string `json:"AssetCreateTime,omitnil,omitempty" name:"AssetCreateTime"`

	// Public IP
	// Note: This field may return·null, indicating that no valid values can be obtained.
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// Private IP
	// Note: This field may return·null, indicating that no valid values can be obtained.
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`


	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// VPC name
	// Note: This field may return·null, indicating that no valid values can be obtained.
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// App ID
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User `uin`
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// User name
	// Note: This field may return·null, indicating that no valid values can be obtained.
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// Availability zone
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AvailableArea *string `json:"AvailableArea,omitnil,omitempty" name:"AvailableArea"`

	// Whether it's a critical asset
	// Note: This field may return·null, indicating that no valid values can be obtained.
	IsCore *uint64 `json:"IsCore,omitnil,omitempty" name:"IsCore"`

	// Subnet ID
	// Note: This field may return·null, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Subnet name
	// Note: This field may return·null, indicating that no valid values can be obtained.
	SubnetName *string `json:"SubnetName,omitnil,omitempty" name:"SubnetName"`

	// UUID of the instance
	// Note: This field may return·null, indicating that no valid values can be obtained.
	InstanceUuid *string `json:"InstanceUuid,omitnil,omitempty" name:"InstanceUuid"`

	// QUuid of the instance
	// Note: This field may return·null, indicating that no valid values can be obtained.
	InstanceQUuid *string `json:"InstanceQUuid,omitnil,omitempty" name:"InstanceQUuid"`

	// OS name
	// Note: This field may return·null, indicating that no valid values can be obtained.
	OsName *string `json:"OsName,omitnil,omitempty" name:"OsName"`

	// Number of partitions
	// Note: This field may return·null, indicating that no valid values can be obtained.
	PartitionCount *uint64 `json:"PartitionCount,omitnil,omitempty" name:"PartitionCount"`

	// CPU information
	// Note: This field may return·null, indicating that no valid values can be obtained.
	CPUInfo *string `json:"CPUInfo,omitnil,omitempty" name:"CPUInfo"`

	// CPU size
	// Note: This field may return·null, indicating that no valid values can be obtained.
	CPUSize *uint64 `json:"CPUSize,omitnil,omitempty" name:"CPUSize"`

	// CPU load
	// Note: This field may return·null, indicating that no valid values can be obtained.
	CPULoad *string `json:"CPULoad,omitnil,omitempty" name:"CPULoad"`

	// Memory size
	// Note: This field may return·null, indicating that no valid values can be obtained.
	MemorySize *string `json:"MemorySize,omitnil,omitempty" name:"MemorySize"`

	// Memory load
	// Note: This field may return·null, indicating that no valid values can be obtained.
	MemoryLoad *string `json:"MemoryLoad,omitnil,omitempty" name:"MemoryLoad"`

	// Disk size.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	DiskSize *string `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// Disk load
	// Note: This field may return·null, indicating that no valid values can be obtained.
	DiskLoad *string `json:"DiskLoad,omitnil,omitempty" name:"DiskLoad"`

	// Number of accounts
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AccountCount *string `json:"AccountCount,omitnil,omitempty" name:"AccountCount"`

	// Number of processes
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ProcessCount *string `json:"ProcessCount,omitnil,omitempty" name:"ProcessCount"`

	// Number of applications
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AppCount *string `json:"AppCount,omitnil,omitempty" name:"AppCount"`

	// Number of listened ports.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	PortCount *uint64 `json:"PortCount,omitnil,omitempty" name:"PortCount"`

	// Number of network attacks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Attack *uint64 `json:"Attack,omitnil,omitempty" name:"Attack"`

	// Number of network access requests
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Access *uint64 `json:"Access,omitnil,omitempty" name:"Access"`

	// Number of blocked attacks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Intercept *uint64 `json:"Intercept,omitnil,omitempty" name:"Intercept"`

	// Inbound peak bandwidth
	// Note: This field may return·null, indicating that no valid values can be obtained.
	InBandwidth *string `json:"InBandwidth,omitnil,omitempty" name:"InBandwidth"`

	// OutInbound peak bandwidth
	// Note: This field may return·null, indicating that no valid values can be obtained.
	OutBandwidth *string `json:"OutBandwidth,omitnil,omitempty" name:"OutBandwidth"`

	// Total inbound traffic
	// Note: This field may return·null, indicating that no valid values can be obtained.
	InFlow *string `json:"InFlow,omitnil,omitempty" name:"InFlow"`

	// Total outbound traffic
	// Note: This field may return·null, indicating that no valid values can be obtained.
	OutFlow *string `json:"OutFlow,omitnil,omitempty" name:"OutFlow"`

	// Last scan time
	// Note: This field may return·null, indicating that no valid values can be obtained.
	LastScanTime *string `json:"LastScanTime,omitnil,omitempty" name:"LastScanTime"`

	// Proactive malicious outgoing requests
	// Note: This field may return·null, indicating that no valid values can be obtained.
	NetWorkOut *uint64 `json:"NetWorkOut,omitnil,omitempty" name:"NetWorkOut"`

	// Port risks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	PortRisk *uint64 `json:"PortRisk,omitnil,omitempty" name:"PortRisk"`

	// Vulnerabilities
	// Note: This field may return·null, indicating that no valid values can be obtained.
	VulnerabilityRisk *uint64 `json:"VulnerabilityRisk,omitnil,omitempty" name:"VulnerabilityRisk"`

	// Configuraiton risks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ConfigurationRisk *uint64 `json:"ConfigurationRisk,omitnil,omitempty" name:"ConfigurationRisk"`

	// Number of scan tasks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ScanTask *uint64 `json:"ScanTask,omitnil,omitempty" name:"ScanTask"`

	// Tags
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Member ID
	// Note: This field may return·null, indicating that no valid values can be obtained.
	MemberId *string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Full name of the operating system
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// Risk exposure
	// Note: This field may return·null, indicating that no valid values can be obtained.
	RiskExposure *int64 `json:"RiskExposure,omitnil,omitempty" name:"RiskExposure"`

	// BAS toolkit status. `0`: Not installed; `1`: Installed; `2`: Offline.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	BASAgentStatus *int64 `json:"BASAgentStatus,omitnil,omitempty" name:"BASAgentStatus"`

	// `1`: New asset; `0`: Not a new asset
	// Note: This field may return·null, indicating that no valid values can be obtained.
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil,omitempty" name:"IsNewAsset"`
}

type ClbListenerListInfo struct {
	// Listener ID
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// The listener name.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ListenerName *string `json:"ListenerName,omitnil,omitempty" name:"ListenerName"`

	// Load balancer ID
	// Note: This field may return·null, indicating that no valid values can be obtained.
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB instance name
	// Note: This field may return·null, indicating that no valid values can be obtained.
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// Network protocol
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Region
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// CLB instance IP
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// Port
	// Note: This field may return·null, indicating that no valid values can be obtained.
	VPort *int64 `json:"VPort,omitnil,omitempty" name:"VPort"`

	// Availability zone
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// VPC ID
	// Note: This field may return·null, indicating that no valid values can be obtained.
	NumericalVpcId *int64 `json:"NumericalVpcId,omitnil,omitempty" name:"NumericalVpcId"`

	// CLB instance type
	// Note: This field may return·null, indicating that no valid values can be obtained.
	LoadBalancerType *string `json:"LoadBalancerType,omitnil,omitempty" name:"LoadBalancerType"`

	// Listener domain name
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Load balancer domain name
	// Note: This field may return·null, indicating that no valid values can be obtained.
	LoadBalancerDomain *string `json:"LoadBalancerDomain,omitnil,omitempty" name:"LoadBalancerDomain"`
}

// Predefined struct for user
type CreateDomainAndIpRequestParams struct {
	// Public IP/domain name
	Content []*string `json:"Content,omitnil,omitempty" name:"Content"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateDomainAndIpRequest struct {
	*tchttp.BaseRequest
	
	// Public IP/domain name
	Content []*string `json:"Content,omitnil,omitempty" name:"Content"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateDomainAndIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainAndIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Content")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDomainAndIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDomainAndIpResponseParams struct {
	// Number of created assets
	Data *int64 `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDomainAndIpResponse struct {
	*tchttp.BaseResponse
	Response *CreateDomainAndIpResponseParams `json:"Response"`
}

func (r *CreateDomainAndIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainAndIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRiskCenterScanTaskRequestParams struct {
	// Task name
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Values: `0` (Scan all); `1` (Scan specific assets); `2` (Scan all expect the specified assets); `3` (Custom assets). When `ScanAssetType=1/2`, `Assets` is required. When `ScanAssetType=3`, `SelfDefiningAssets` is required. 
	ScanAssetType *int64 `json:"ScanAssetType,omitnil,omitempty" name:"ScanAssetType"`

	// Project to scan: port/poc/weakpass/webcontent/configrisk/exposedserver
	ScanItem []*string `json:"ScanItem,omitnil,omitempty" name:"ScanItem"`

	// Task type. `0`: Scheduled task, `1`: Scan immediately; `2`: Scanned at the specified time; `3`: Custom. When ScanPlanType=0,2,3, `ScanPlanContent` is required.
	ScanPlanType *int64 `json:"ScanPlanType,omitnil,omitempty" name:"ScanPlanType"`

	// List of assets to scan
	Assets []*TaskAssetObject `json:"Assets,omitnil,omitempty" name:"Assets"`

	// Details of a scheduled scan task
	ScanPlanContent *string `json:"ScanPlanContent,omitnil,omitempty" name:"ScanPlanContent"`

	// IP/Domain name/URL
	SelfDefiningAssets []*string `json:"SelfDefiningAssets,omitnil,omitempty" name:"SelfDefiningAssets"`

	// Request source. Values: `vss` (Vulnerability Scan Service), `csip` (Cloud Security Center). It defaults to `vss`.
	ScanFrom *string `json:"ScanFrom,omitnil,omitempty" name:"ScanFrom"`

	// Advanced settings
	TaskAdvanceCFG *TaskAdvanceCFG `json:"TaskAdvanceCFG,omitnil,omitempty" name:"TaskAdvanceCFG"`

	// Scan task mode: `0` (Standard), `1` (Quick), `2` (Advanced). Default: `0`
	TaskMode *int64 `json:"TaskMode,omitnil,omitempty" name:"TaskMode"`

	// Asset tags
	Tags *AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateRiskCenterScanTaskRequest struct {
	*tchttp.BaseRequest
	
	// Task name
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Values: `0` (Scan all); `1` (Scan specific assets); `2` (Scan all expect the specified assets); `3` (Custom assets). When `ScanAssetType=1/2`, `Assets` is required. When `ScanAssetType=3`, `SelfDefiningAssets` is required. 
	ScanAssetType *int64 `json:"ScanAssetType,omitnil,omitempty" name:"ScanAssetType"`

	// Project to scan: port/poc/weakpass/webcontent/configrisk/exposedserver
	ScanItem []*string `json:"ScanItem,omitnil,omitempty" name:"ScanItem"`

	// Task type. `0`: Scheduled task, `1`: Scan immediately; `2`: Scanned at the specified time; `3`: Custom. When ScanPlanType=0,2,3, `ScanPlanContent` is required.
	ScanPlanType *int64 `json:"ScanPlanType,omitnil,omitempty" name:"ScanPlanType"`

	// List of assets to scan
	Assets []*TaskAssetObject `json:"Assets,omitnil,omitempty" name:"Assets"`

	// Details of a scheduled scan task
	ScanPlanContent *string `json:"ScanPlanContent,omitnil,omitempty" name:"ScanPlanContent"`

	// IP/Domain name/URL
	SelfDefiningAssets []*string `json:"SelfDefiningAssets,omitnil,omitempty" name:"SelfDefiningAssets"`

	// Request source. Values: `vss` (Vulnerability Scan Service), `csip` (Cloud Security Center). It defaults to `vss`.
	ScanFrom *string `json:"ScanFrom,omitnil,omitempty" name:"ScanFrom"`

	// Advanced settings
	TaskAdvanceCFG *TaskAdvanceCFG `json:"TaskAdvanceCFG,omitnil,omitempty" name:"TaskAdvanceCFG"`

	// Scan task mode: `0` (Standard), `1` (Quick), `2` (Advanced). Default: `0`
	TaskMode *int64 `json:"TaskMode,omitnil,omitempty" name:"TaskMode"`

	// Asset tags
	Tags *AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateRiskCenterScanTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRiskCenterScanTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskName")
	delete(f, "ScanAssetType")
	delete(f, "ScanItem")
	delete(f, "ScanPlanType")
	delete(f, "Assets")
	delete(f, "ScanPlanContent")
	delete(f, "SelfDefiningAssets")
	delete(f, "ScanFrom")
	delete(f, "TaskAdvanceCFG")
	delete(f, "TaskMode")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRiskCenterScanTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRiskCenterScanTaskResponseParams struct {
	// Task ID.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// `0`: Task created successfully. `-1`: There are unauthorized assets. 
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// List of unauthorized assets
	UnAuthAsset []*string `json:"UnAuthAsset,omitnil,omitempty" name:"UnAuthAsset"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRiskCenterScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateRiskCenterScanTaskResponseParams `json:"Response"`
}

func (r *CreateRiskCenterScanTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRiskCenterScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DBAssetVO struct {
	// Asset ID
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// Asset name
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// Asset type
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`


	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// VPC tags
	// Note: This field may return·null, indicating that no valid values can be obtained.
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// Region
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Domain name
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Asset creation time
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetCreateTime *string `json:"AssetCreateTime,omitnil,omitempty" name:"AssetCreateTime"`

	// Last scan time
	// Note: This field may return·null, indicating that no valid values can be obtained.
	LastScanTime *string `json:"LastScanTime,omitnil,omitempty" name:"LastScanTime"`

	// Configuration risks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ConfigurationRisk *uint64 `json:"ConfigurationRisk,omitnil,omitempty" name:"ConfigurationRisk"`

	// Network attacks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Attack *uint64 `json:"Attack,omitnil,omitempty" name:"Attack"`


	Access *uint64 `json:"Access,omitnil,omitempty" name:"Access"`

	// Scan tasks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ScanTask *uint64 `json:"ScanTask,omitnil,omitempty" name:"ScanTask"`

	// User `appid`
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User UIN
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// User name
	// Note: This field may return·null, indicating that no valid values can be obtained.
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// Port
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Tags
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// Private IP
	// Note: This field may return·null, indicating that no valid values can be obtained.
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// Public IP
	// Note: This field may return·null, indicating that no valid values can be obtained.
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// Status
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Whether it's a critical asset
	// Note: This field may return·null, indicating that no valid values can be obtained.
	IsCore *uint64 `json:"IsCore,omitnil,omitempty" name:"IsCore"`

	// Whether it's a newly-added asset. Values: `1` (Yes), `0` (No)
	// Note: This field may return·null, indicating that no valid values can be obtained.
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil,omitempty" name:"IsNewAsset"`
}

type DataSearchBug struct {
	// Query status code
	StateCode *string `json:"StateCode,omitnil,omitempty" name:"StateCode"`

	//  
	// Note: This field may return·null, indicating that no valid values can be obtained.
	DataBug []*BugInfoDetail `json:"DataBug,omitnil,omitempty" name:"DataBug"`

	// None
	// Note: This field may return·null, indicating that no valid values can be obtained.
	DataAsset []*AssetInfoDetail `json:"DataAsset,omitnil,omitempty" name:"DataAsset"`

	// `true`: Support vulnerability scan; `false`: Do not support vulnerability scan
	// Note: This field may return·null, indicating that no valid values can be obtained.
	VSSScan *bool `json:"VSSScan,omitnil,omitempty" name:"VSSScan"`

	// `0`: Do not support; `1`: Support
	// Note: This field may return·null, indicating that no valid values can be obtained.
	CWPScan *string `json:"CWPScan,omitnil,omitempty" name:"CWPScan"`

	// `1`: Support virtual patches; `0` or null: Do not support
	// Note: This field may return·null, indicating that no valid values can be obtained.
	CFWPatch *string `json:"CFWPatch,omitnil,omitempty" name:"CFWPatch"`

	// `0`: Do not support; `1`: Support
	// Note: This field may return·null, indicating that no valid values can be obtained.
	WafPatch *int64 `json:"WafPatch,omitnil,omitempty" name:"WafPatch"`

	// `0`: Do not support; `1`: Support
	// Note: This field may return·null, indicating that no valid values can be obtained.
	CWPFix *int64 `json:"CWPFix,omitnil,omitempty" name:"CWPFix"`
}

type DbAssetInfo struct {
	// CFW status
	// Note: This field may return·null, indicating that no valid values can be obtained.
	CFWStatus *uint64 `json:"CFWStatus,omitnil,omitempty" name:"CFWStatus"`

	// Asset ID
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// VPC information
	// Note: This field may return·null, indicating that no valid values can be obtained.
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// Asset type
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// Public IP
	// Note: This field may return·null, indicating that no valid values can be obtained.
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// Private IP
	// Note: This field may return·null, indicating that no valid values can be obtained.
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// Region
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`


	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Asset name
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// CFW edition
	// Note: This field may return·null, indicating that no valid values can be obtained.
	CFWProtectLevel *uint64 `json:"CFWProtectLevel,omitnil,omitempty" name:"CFWProtectLevel"`

	// Tag information
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`
}

// Predefined struct for user
type DeleteDomainAndIpRequestParams struct {
	// u200c-
	Content []*PublicIpDomainListKey `json:"Content,omitnil,omitempty" name:"Content"`

	// Whether to retain the path configuration. `1`: Retain; Others: Do not retain. It defaults to do not retain if not specified.
	RetainPath *int64 `json:"RetainPath,omitnil,omitempty" name:"RetainPath"`

	// Whether to ignore this asset in the future. `1`: Ignore; Others: Do not ignore. It defaults to ignore if not specified.
	IgnoreAsset *int64 `json:"IgnoreAsset,omitnil,omitempty" name:"IgnoreAsset"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Deletion mode. Values: `ALL` (delete all). If it's not specified, `Content` is required.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DeleteDomainAndIpRequest struct {
	*tchttp.BaseRequest
	
	// u200c-
	Content []*PublicIpDomainListKey `json:"Content,omitnil,omitempty" name:"Content"`

	// Whether to retain the path configuration. `1`: Retain; Others: Do not retain. It defaults to do not retain if not specified.
	RetainPath *int64 `json:"RetainPath,omitnil,omitempty" name:"RetainPath"`

	// Whether to ignore this asset in the future. `1`: Ignore; Others: Do not ignore. It defaults to ignore if not specified.
	IgnoreAsset *int64 `json:"IgnoreAsset,omitnil,omitempty" name:"IgnoreAsset"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Deletion mode. Values: `ALL` (delete all). If it's not specified, `Content` is required.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DeleteDomainAndIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainAndIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Content")
	delete(f, "RetainPath")
	delete(f, "IgnoreAsset")
	delete(f, "Tags")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDomainAndIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDomainAndIpResponseParams struct {
	// Number of deleted assets
	Data *int64 `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDomainAndIpResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDomainAndIpResponseParams `json:"Response"`
}

func (r *DeleteDomainAndIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainAndIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRiskScanTaskRequestParams struct {
	// List of task IDs
	TaskIdList []*TaskIdListKey `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`
}

type DeleteRiskScanTaskRequest struct {
	*tchttp.BaseRequest
	
	// List of task IDs
	TaskIdList []*TaskIdListKey `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`
}

func (r *DeleteRiskScanTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRiskScanTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRiskScanTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRiskScanTaskResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRiskScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRiskScanTaskResponseParams `json:"Response"`
}

func (r *DeleteRiskScanTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRiskScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCVMAssetInfoRequestParams struct {
	// u200c-
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`
}

type DescribeCVMAssetInfoRequest struct {
	*tchttp.BaseRequest
	
	// u200c-
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`
}

func (r *DescribeCVMAssetInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCVMAssetInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCVMAssetInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCVMAssetInfoResponseParams struct {
	// u200c-
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Data *AssetBaseInfoResponse `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCVMAssetInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCVMAssetInfoResponseParams `json:"Response"`
}

func (r *DescribeCVMAssetInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCVMAssetInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCVMAssetsRequestParams struct {
	// u200c-
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeCVMAssetsRequest struct {
	*tchttp.BaseRequest
	
	// u200c-
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeCVMAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCVMAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCVMAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCVMAssetsResponseParams struct {
	// u200c-
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// u200c-
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Data []*CVMAssetVO `json:"Data,omitnil,omitempty" name:"Data"`

	// List of regions
	// Note: This field may return·null, indicating that no valid values can be obtained.
	RegionList []*FilterDataObject `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// Protection status
	// Note: This field may return·null, indicating that no valid values can be obtained.
	DefenseStatusList []*FilterDataObject `json:"DefenseStatusList,omitnil,omitempty" name:"DefenseStatusList"`

	// List of VPCs
	// Note: This field may return·null, indicating that no valid values can be obtained.
	VpcList []*FilterDataObject `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// List of asset types
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetTypeList []*FilterDataObject `json:"AssetTypeList,omitnil,omitempty" name:"AssetTypeList"`

	// List of operating systems
	// Note: This field may return·null, indicating that no valid values can be obtained.
	SystemTypeList []*FilterDataObject `json:"SystemTypeList,omitnil,omitempty" name:"SystemTypeList"`

	// List of IP types
	// Note: This field may return·null, indicating that no valid values can be obtained.
	IpTypeList []*FilterDataObject `json:"IpTypeList,omitnil,omitempty" name:"IpTypeList"`

	// List of AppIds
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil,omitempty" name:"AppIdList"`

	// List of availability zones
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ZoneList []*FilterDataObject `json:"ZoneList,omitnil,omitempty" name:"ZoneList"`

	// List of operating systems
	// Note: This field may return·null, indicating that no valid values can be obtained.
	OsList []*FilterDataObject `json:"OsList,omitnil,omitempty" name:"OsList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCVMAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCVMAssetsResponseParams `json:"Response"`
}

func (r *DescribeCVMAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCVMAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterPodAssetsRequestParams struct {
	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeClusterPodAssetsRequest struct {
	*tchttp.BaseRequest
	
	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeClusterPodAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterPodAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterPodAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterPodAssetsResponseParams struct {
	// Data list
	Data []*AssetClusterPod `json:"Data,omitnil,omitempty" name:"Data"`

	// Total number of results
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of cluster pod status
	PodStatusList []*FilterDataObject `json:"PodStatusList,omitnil,omitempty" name:"PodStatusList"`

	// List of namespaces
	NamespaceList []*FilterDataObject `json:"NamespaceList,omitnil,omitempty" name:"NamespaceList"`

	// List of regions
	RegionList []*FilterDataObject `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// List of users (AppId)
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil,omitempty" name:"AppIdList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterPodAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterPodAssetsResponseParams `json:"Response"`
}

func (r *DescribeClusterPodAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterPodAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDbAssetInfoRequestParams struct {
	// Asset ID
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`
}

type DescribeDbAssetInfoRequest struct {
	*tchttp.BaseRequest
	
	// Asset ID
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`
}

func (r *DescribeDbAssetInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDbAssetInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AssetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDbAssetInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDbAssetInfoResponseParams struct {
	// Details of a database asset. 
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Data *DbAssetInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDbAssetInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDbAssetInfoResponseParams `json:"Response"`
}

func (r *DescribeDbAssetInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDbAssetInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDbAssetsRequestParams struct {
	// u200c-
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset types. Values: MYSQL/MARIADB/REDIS/MONGODB/POSTGRES/CTS/ES/KAFKA/COS/CBS/CFS
	AssetTypes []*string `json:"AssetTypes,omitnil,omitempty" name:"AssetTypes"`
}

type DescribeDbAssetsRequest struct {
	*tchttp.BaseRequest
	
	// u200c-
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset types. Values: MYSQL/MARIADB/REDIS/MONGODB/POSTGRES/CTS/ES/KAFKA/COS/CBS/CFS
	AssetTypes []*string `json:"AssetTypes,omitnil,omitempty" name:"AssetTypes"`
}

func (r *DescribeDbAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDbAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "AssetTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDbAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDbAssetsResponseParams struct {
	// Total number of results
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Total of assets
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Data []*DBAssetVO `json:"Data,omitnil,omitempty" name:"Data"`

	// List of regions
	// Note: This field may return·null, indicating that no valid values can be obtained.
	RegionList []*FilterDataObject `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// List of asset types
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetTypeList []*FilterDataObject `json:"AssetTypeList,omitnil,omitempty" name:"AssetTypeList"`

	// List of VPCs
	// Note: This field may return·null, indicating that no valid values can be obtained.
	VpcList []*FilterDataObject `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// List of users (AppId)
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil,omitempty" name:"AppIdList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDbAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDbAssetsResponseParams `json:"Response"`
}

func (r *DescribeDbAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDbAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainAssetsRequestParams struct {
	// u200c-
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// CSC tags of the asset
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeDomainAssetsRequest struct {
	*tchttp.BaseRequest
	
	// u200c-
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// CSC tags of the asset
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeDomainAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainAssetsResponseParams struct {
	// u200c-
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// u200c-
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Data []*DomainAssetVO `json:"Data,omitnil,omitempty" name:"Data"`

	// List of WAF protection status
	// Note: This field may return·null, indicating that no valid values can be obtained.
	DefenseStatusList []*FilterDataObject `json:"DefenseStatusList,omitnil,omitempty" name:"DefenseStatusList"`

	// List of asset locations
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetLocationList []*FilterDataObject `json:"AssetLocationList,omitnil,omitempty" name:"AssetLocationList"`

	// List of asset types
	// Note: This field may return·null, indicating that no valid values can be obtained.
	SourceTypeList []*FilterDataObject `json:"SourceTypeList,omitnil,omitempty" name:"SourceTypeList"`

	// List of regions
	// Note: This field may return·null, indicating that no valid values can be obtained.
	RegionList []*FilterDataObject `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDomainAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainAssetsResponseParams `json:"Response"`
}

func (r *DescribeDomainAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListenerListRequestParams struct {
	// u200c-
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeListenerListRequest struct {
	*tchttp.BaseRequest
	
	// u200c-
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeListenerListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListenerListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListenerListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListenerListResponseParams struct {
	// Total number of results
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// List of listeners
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Data []*ClbListenerListInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeListenerListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListenerListResponseParams `json:"Response"`
}

func (r *DescribeListenerListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListenerListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublicIpAssetsRequestParams struct {
	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// CSC tags of the asset
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribePublicIpAssetsRequest struct {
	*tchttp.BaseRequest
	
	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// CSC tags of the asset
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribePublicIpAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicIpAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePublicIpAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublicIpAssetsResponseParams struct {
	// Data list
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Data []*IpAssetListVO `json:"Data,omitnil,omitempty" name:"Data"`

	// Total number of results
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// List of asset locations
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetLocationList []*FilterDataObject `json:"AssetLocationList,omitnil,omitempty" name:"AssetLocationList"`

	// List of IP types
	// Note: This field may return·null, indicating that no valid values can be obtained.
	IpTypeList []*FilterDataObject `json:"IpTypeList,omitnil,omitempty" name:"IpTypeList"`

	// List of regions
	// Note: This field may return·null, indicating that no valid values can be obtained.
	RegionList []*FilterDataObject `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// List of protection status
	// Note: This field may return·null, indicating that no valid values can be obtained.
	DefenseStatusList []*FilterDataObject `json:"DefenseStatusList,omitnil,omitempty" name:"DefenseStatusList"`

	// List of asset types
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetTypeList []*FilterDataObject `json:"AssetTypeList,omitnil,omitempty" name:"AssetTypeList"`

	// List of AppIds 
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil,omitempty" name:"AppIdList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePublicIpAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePublicIpAssetsResponseParams `json:"Response"`
}

func (r *DescribePublicIpAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicIpAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterAssetViewCFGRiskListRequestParams struct {
	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeRiskCenterAssetViewCFGRiskListRequest struct {
	*tchttp.BaseRequest
	
	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeRiskCenterAssetViewCFGRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterAssetViewCFGRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskCenterAssetViewCFGRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterAssetViewCFGRiskListResponseParams struct {
	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of configuration risks
	Data []*AssetViewCFGRisk `json:"Data,omitnil,omitempty" name:"Data"`

	// List of risk handling status
	StatusLists []*FilterDataObject `json:"StatusLists,omitnil,omitempty" name:"StatusLists"`

	// List of risk levels
	LevelLists []*FilterDataObject `json:"LevelLists,omitnil,omitempty" name:"LevelLists"`

	// List of configuration names
	CFGNameLists []*FilterDataObject `json:"CFGNameLists,omitnil,omitempty" name:"CFGNameLists"`

	// List of check types
	CheckTypeLists []*FilterDataObject `json:"CheckTypeLists,omitnil,omitempty" name:"CheckTypeLists"`

	// List of asset types
	InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitnil,omitempty" name:"InstanceTypeLists"`

	// List of check source
	FromLists []*FilterDataObject `json:"FromLists,omitnil,omitempty" name:"FromLists"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskCenterAssetViewCFGRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskCenterAssetViewCFGRiskListResponseParams `json:"Response"`
}

func (r *DescribeRiskCenterAssetViewCFGRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterAssetViewCFGRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterAssetViewPortRiskListRequestParams struct {
	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeRiskCenterAssetViewPortRiskListRequest struct {
	*tchttp.BaseRequest
	
	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeRiskCenterAssetViewPortRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterAssetViewPortRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskCenterAssetViewPortRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterAssetViewPortRiskListResponseParams struct {
	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of configuration risks
	Data []*AssetViewPortRisk `json:"Data,omitnil,omitempty" name:"Data"`

	// List of risk handling status
	StatusLists []*FilterDataObject `json:"StatusLists,omitnil,omitempty" name:"StatusLists"`

	// List of risk levels
	LevelLists []*FilterDataObject `json:"LevelLists,omitnil,omitempty" name:"LevelLists"`

	// List of fix suggestions 
	SuggestionLists []*FilterDataObject `json:"SuggestionLists,omitnil,omitempty" name:"SuggestionLists"`

	// List of asset types
	InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitnil,omitempty" name:"InstanceTypeLists"`

	// List of check source
	FromLists []*FilterDataObject `json:"FromLists,omitnil,omitempty" name:"FromLists"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskCenterAssetViewPortRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskCenterAssetViewPortRiskListResponseParams `json:"Response"`
}

func (r *DescribeRiskCenterAssetViewPortRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterAssetViewPortRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterAssetViewVULRiskListRequestParams struct {
	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeRiskCenterAssetViewVULRiskListRequest struct {
	*tchttp.BaseRequest
	
	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeRiskCenterAssetViewVULRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterAssetViewVULRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskCenterAssetViewVULRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterAssetViewVULRiskListResponseParams struct {
	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of vulnerabilities
	Data []*AssetViewVULRisk `json:"Data,omitnil,omitempty" name:"Data"`

	// List of risk handling status
	StatusLists []*FilterDataObject `json:"StatusLists,omitnil,omitempty" name:"StatusLists"`

	// List of risk levels
	LevelLists []*FilterDataObject `json:"LevelLists,omitnil,omitempty" name:"LevelLists"`

	// List of check source
	FromLists []*FilterDataObject `json:"FromLists,omitnil,omitempty" name:"FromLists"`

	// List of vulnerability types
	VULTypeLists []*FilterDataObject `json:"VULTypeLists,omitnil,omitempty" name:"VULTypeLists"`

	// List of asset types
	InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitnil,omitempty" name:"InstanceTypeLists"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskCenterAssetViewVULRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskCenterAssetViewVULRiskListResponseParams `json:"Response"`
}

func (r *DescribeRiskCenterAssetViewVULRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterAssetViewVULRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterAssetViewWeakPasswordRiskListRequestParams struct {
	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeRiskCenterAssetViewWeakPasswordRiskListRequest struct {
	*tchttp.BaseRequest
	
	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeRiskCenterAssetViewWeakPasswordRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterAssetViewWeakPasswordRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskCenterAssetViewWeakPasswordRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterAssetViewWeakPasswordRiskListResponseParams struct {
	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of risks
	Data []*AssetViewWeakPassRisk `json:"Data,omitnil,omitempty" name:"Data"`

	// List of risk handling status
	StatusLists []*FilterDataObject `json:"StatusLists,omitnil,omitempty" name:"StatusLists"`

	// List of risk levels
	LevelLists []*FilterDataObject `json:"LevelLists,omitnil,omitempty" name:"LevelLists"`

	// List of check source
	FromLists []*FilterDataObject `json:"FromLists,omitnil,omitempty" name:"FromLists"`

	// List of asset types
	InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitnil,omitempty" name:"InstanceTypeLists"`

	// List of weak password types
	PasswordTypeLists []*FilterDataObject `json:"PasswordTypeLists,omitnil,omitempty" name:"PasswordTypeLists"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskCenterAssetViewWeakPasswordRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskCenterAssetViewWeakPasswordRiskListResponseParams `json:"Response"`
}

func (r *DescribeRiskCenterAssetViewWeakPasswordRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterAssetViewWeakPasswordRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterPortViewPortRiskListRequestParams struct {
	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeRiskCenterPortViewPortRiskListRequest struct {
	*tchttp.BaseRequest
	
	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeRiskCenterPortViewPortRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterPortViewPortRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskCenterPortViewPortRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterPortViewPortRiskListResponseParams struct {
	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of port risks by assets
	Data []*PortViewPortRisk `json:"Data,omitnil,omitempty" name:"Data"`

	// List of risk levels
	LevelLists []*FilterDataObject `json:"LevelLists,omitnil,omitempty" name:"LevelLists"`

	// List of suggestions
	SuggestionLists []*FilterDataObject `json:"SuggestionLists,omitnil,omitempty" name:"SuggestionLists"`

	// List of check source
	FromLists []*FilterDataObject `json:"FromLists,omitnil,omitempty" name:"FromLists"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskCenterPortViewPortRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskCenterPortViewPortRiskListResponseParams `json:"Response"`
}

func (r *DescribeRiskCenterPortViewPortRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterPortViewPortRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterServerRiskListRequestParams struct {
	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeRiskCenterServerRiskListRequest struct {
	*tchttp.BaseRequest
	
	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeRiskCenterServerRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterServerRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskCenterServerRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterServerRiskListResponseParams struct {
	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of services in risk
	Data []*ServerRisk `json:"Data,omitnil,omitempty" name:"Data"`

	// List of asset types
	InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitnil,omitempty" name:"InstanceTypeLists"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskCenterServerRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskCenterServerRiskListResponseParams `json:"Response"`
}

func (r *DescribeRiskCenterServerRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterServerRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterVULViewVULRiskListRequestParams struct {
	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeRiskCenterVULViewVULRiskListRequest struct {
	*tchttp.BaseRequest
	
	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeRiskCenterVULViewVULRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterVULViewVULRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskCenterVULViewVULRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterVULViewVULRiskListResponseParams struct {
	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of vulnerabilities
	Data []*VULViewVULRisk `json:"Data,omitnil,omitempty" name:"Data"`

	// List of risk levels
	LevelLists []*FilterDataObject `json:"LevelLists,omitnil,omitempty" name:"LevelLists"`

	// List of check source
	FromLists []*FilterDataObject `json:"FromLists,omitnil,omitempty" name:"FromLists"`

	// List of vulnerability types
	VULTypeLists []*FilterDataObject `json:"VULTypeLists,omitnil,omitempty" name:"VULTypeLists"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskCenterVULViewVULRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskCenterVULViewVULRiskListResponseParams `json:"Response"`
}

func (r *DescribeRiskCenterVULViewVULRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterVULViewVULRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterWebsiteRiskListRequestParams struct {
	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeRiskCenterWebsiteRiskListRequest struct {
	*tchttp.BaseRequest
	
	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Asset tags
	Tags []*AssetTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeRiskCenterWebsiteRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterWebsiteRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskCenterWebsiteRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCenterWebsiteRiskListResponseParams struct {
	// Total number of entries
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of content risks
	Data []*WebsiteRisk `json:"Data,omitnil,omitempty" name:"Data"`

	// List of risk handling status
	StatusLists []*FilterDataObject `json:"StatusLists,omitnil,omitempty" name:"StatusLists"`

	// List of risk levels
	LevelLists []*FilterDataObject `json:"LevelLists,omitnil,omitempty" name:"LevelLists"`

	// List of asset types
	InstanceTypeLists []*FilterDataObject `json:"InstanceTypeLists,omitnil,omitempty" name:"InstanceTypeLists"`

	// List of risk types
	DetectEngineLists []*FilterDataObject `json:"DetectEngineLists,omitnil,omitempty" name:"DetectEngineLists"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskCenterWebsiteRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskCenterWebsiteRiskListResponseParams `json:"Response"`
}

func (r *DescribeRiskCenterWebsiteRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCenterWebsiteRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanReportListRequestParams struct {
	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeScanReportListRequest struct {
	*tchttp.BaseRequest
	
	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeScanReportListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanReportListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanReportListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanReportListResponseParams struct {
	// Total number of entries
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of scan reports
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Data []*ScanTaskInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// List of account UINs
	UINList []*string `json:"UINList,omitnil,omitempty" name:"UINList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeScanReportListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScanReportListResponseParams `json:"Response"`
}

func (r *DescribeScanReportListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanReportListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanTaskListRequestParams struct {
	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Tags
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeScanTaskListRequest struct {
	*tchttp.BaseRequest
	
	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// Tags
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeScanTaskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanTaskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanTaskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanTaskListResponseParams struct {
	// Total number of entries
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of scan tasks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Data []*ScanTaskInfoList `json:"Data,omitnil,omitempty" name:"Data"`

	// List of account UINs
	// Note: This field may return·null, indicating that no valid values can be obtained.
	UINList []*string `json:"UINList,omitnil,omitempty" name:"UINList"`

	// List of task modes
	// Note: This field may return·null, indicating that no valid values can be obtained.
	TaskModeList []*FilterDataObject `json:"TaskModeList,omitnil,omitempty" name:"TaskModeList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeScanTaskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScanTaskListResponseParams `json:"Response"`
}

func (r *DescribeScanTaskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSearchBugInfoRequestParams struct {
	// Type of the query action. `1`: Query emergency vulnerabilities; `2`: Query all vulnerabilities; `3`: Query a specific vulnerability. When `Id=3`, `CVEId` is required. 
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// CVE number of the vulnerability. It's required when `Id=3`.
	CVEId *string `json:"CVEId,omitnil,omitempty" name:"CVEId"`
}

type DescribeSearchBugInfoRequest struct {
	*tchttp.BaseRequest
	
	// Type of the query action. `1`: Query emergency vulnerabilities; `2`: Query all vulnerabilities; `3`: Query a specific vulnerability. When `Id=3`, `CVEId` is required. 
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// CVE number of the vulnerability. It's required when `Id=3`.
	CVEId *string `json:"CVEId,omitnil,omitempty" name:"CVEId"`
}

func (r *DescribeSearchBugInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSearchBugInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "CVEId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSearchBugInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSearchBugInfoResponseParams struct {
	// Vulnerability and asset information
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Data *DataSearchBug `json:"Data,omitnil,omitempty" name:"Data"`

	// Status code. Valid values: 0: successful; others: failed.
	ReturnCode *int64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// Status message. Valid values: success: successful query; fail: failed query.
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSearchBugInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSearchBugInfoResponseParams `json:"Response"`
}

func (r *DescribeSearchBugInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSearchBugInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubnetAssetsRequestParams struct {
	// Filter parameters
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeSubnetAssetsRequest struct {
	*tchttp.BaseRequest
	
	// Filter parameters
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeSubnetAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubnetAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubnetAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubnetAssetsResponseParams struct {
	// Data list
	Data []*SubnetAsset `json:"Data,omitnil,omitempty" name:"Data"`

	// Total number of results
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of regions
	RegionList []*FilterDataObject `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// List of VPCs
	VpcList []*FilterDataObject `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// List of AppIds
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil,omitempty" name:"AppIdList"`

	// List of availability zones
	ZoneList []*FilterDataObject `json:"ZoneList,omitnil,omitempty" name:"ZoneList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSubnetAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubnetAssetsResponseParams `json:"Response"`
}

func (r *DescribeSubnetAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubnetAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskLogListRequestParams struct {
	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeTaskLogListRequest struct {
	*tchttp.BaseRequest
	
	// Filter conditions
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeTaskLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskLogListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskLogListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskLogListResponseParams struct {
	// Total number of entries
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of reports
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Data []*TaskLogInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// Number of reports pending viewed
	// Note: This field may return·null, indicating that no valid values can be obtained.
	NotViewNumber *int64 `json:"NotViewNumber,omitnil,omitempty" name:"NotViewNumber"`

	// Number of report templates
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ReportTemplateNumber *int64 `json:"ReportTemplateNumber,omitnil,omitempty" name:"ReportTemplateNumber"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskLogListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskLogListResponseParams `json:"Response"`
}

func (r *DescribeTaskLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskLogURLRequestParams struct {
	// Type of the task. `0`: Preview; `1`: Download
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// List of task report IDs
	ReportItemKeyList []*ReportItemKey `json:"ReportItemKeyList,omitnil,omitempty" name:"ReportItemKeyList"`

	// List of task IDs in the report
	ReportTaskIdList []*ReportTaskIdList `json:"ReportTaskIdList,omitnil,omitempty" name:"ReportTaskIdList"`
}

type DescribeTaskLogURLRequest struct {
	*tchttp.BaseRequest
	
	// Type of the task. `0`: Preview; `1`: Download
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// List of task report IDs
	ReportItemKeyList []*ReportItemKey `json:"ReportItemKeyList,omitnil,omitempty" name:"ReportItemKeyList"`

	// List of task IDs in the report
	ReportTaskIdList []*ReportTaskIdList `json:"ReportTaskIdList,omitnil,omitempty" name:"ReportTaskIdList"`
}

func (r *DescribeTaskLogURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskLogURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "ReportItemKeyList")
	delete(f, "ReportTaskIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskLogURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskLogURLResponseParams struct {
	// Temp download URL of the report
	Data []*TaskLogURL `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskLogURLResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskLogURLResponseParams `json:"Response"`
}

func (r *DescribeTaskLogURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskLogURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVULRiskAdvanceCFGListRequestParams struct {
	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Filter conditions.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeVULRiskAdvanceCFGListRequest struct {
	*tchttp.BaseRequest
	
	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Filter conditions.
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeVULRiskAdvanceCFGListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVULRiskAdvanceCFGListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVULRiskAdvanceCFGListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVULRiskAdvanceCFGListResponseParams struct {
	// List of configuration items
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Data []*VULRiskAdvanceCFGList `json:"Data,omitnil,omitempty" name:"Data"`

	// Total number of results
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of risk levels
	// Note: This field may return·null, indicating that no valid values can be obtained.
	RiskLevelLists []*FilterDataObject `json:"RiskLevelLists,omitnil,omitempty" name:"RiskLevelLists"`

	// List of vulnerabilities types
	// Note: This field may return·null, indicating that no valid values can be obtained.
	VULTypeLists []*FilterDataObject `json:"VULTypeLists,omitnil,omitempty" name:"VULTypeLists"`

	// List of check source
	// Note: This field may return·null, indicating that no valid values can be obtained.
	CheckFromLists []*FilterDataObject `json:"CheckFromLists,omitnil,omitempty" name:"CheckFromLists"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVULRiskAdvanceCFGListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVULRiskAdvanceCFGListResponseParams `json:"Response"`
}

func (r *DescribeVULRiskAdvanceCFGListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVULRiskAdvanceCFGListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcAssetsRequestParams struct {
	// Filter parameters
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeVpcAssetsRequest struct {
	*tchttp.BaseRequest
	
	// Filter parameters
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeVpcAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcAssetsResponseParams struct {
	// Data list
	Data []*Vpc `json:"Data,omitnil,omitempty" name:"Data"`

	// Total number of results
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// List of VPCs
	VpcList []*FilterDataObject `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// List of regions
	RegionList []*FilterDataObject `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// List of AppIds
	AppIdList []*FilterDataObject `json:"AppIdList,omitnil,omitempty" name:"AppIdList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVpcAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcAssetsResponseParams `json:"Response"`
}

func (r *DescribeVpcAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainAssetVO struct {
	// Asset ID
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetId []*string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// Asset name
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetName []*string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// Asset type
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetType []*string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// Region
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Region []*string `json:"Region,omitnil,omitempty" name:"Region"`

	// WAF status
	// Note: This field may return·null, indicating that no valid values can be obtained.
	WAFStatus *uint64 `json:"WAFStatus,omitnil,omitempty" name:"WAFStatus"`

	// Asset creation time
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetCreateTime *string `json:"AssetCreateTime,omitnil,omitempty" name:"AssetCreateTime"`

	// Appid
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Account ID
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Account name
	// Note: This field may return·null, indicating that no valid values can be obtained.
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// Whether it's a critical asset
	// Note: This field may return·null, indicating that no valid values can be obtained.
	IsCore *uint64 `json:"IsCore,omitnil,omitempty" name:"IsCore"`

	// Whether it's a cloud asset
	// Note: This field may return·null, indicating that no valid values can be obtained.
	IsCloud *uint64 `json:"IsCloud,omitnil,omitempty" name:"IsCloud"`

	// Network attacks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Attack *uint64 `json:"Attack,omitnil,omitempty" name:"Attack"`

	// Network access
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Access *uint64 `json:"Access,omitnil,omitempty" name:"Access"`

	// Number of blocked attacks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Intercept *uint64 `json:"Intercept,omitnil,omitempty" name:"Intercept"`

	// Inbound peak bandwidth
	// Note: This field may return·null, indicating that no valid values can be obtained.
	InBandwidth *string `json:"InBandwidth,omitnil,omitempty" name:"InBandwidth"`

	// Outbound peak bandwidth
	// Note: This field may return·null, indicating that no valid values can be obtained.
	OutBandwidth *string `json:"OutBandwidth,omitnil,omitempty" name:"OutBandwidth"`

	// Total inbound traffic
	// Note: This field may return·null, indicating that no valid values can be obtained.
	InFlow *string `json:"InFlow,omitnil,omitempty" name:"InFlow"`

	// Total outbound traffic
	// Note: This field may return·null, indicating that no valid values can be obtained.
	OutFlow *string `json:"OutFlow,omitnil,omitempty" name:"OutFlow"`

	// Last scan time
	// Note: This field may return·null, indicating that no valid values can be obtained.
	LastScanTime *string `json:"LastScanTime,omitnil,omitempty" name:"LastScanTime"`

	// Port risks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	PortRisk *uint64 `json:"PortRisk,omitnil,omitempty" name:"PortRisk"`

	// Vulnerabilities
	// Note: This field may return·null, indicating that no valid values can be obtained.
	VulnerabilityRisk *uint64 `json:"VulnerabilityRisk,omitnil,omitempty" name:"VulnerabilityRisk"`

	// Configuration risks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ConfigurationRisk *uint64 `json:"ConfigurationRisk,omitnil,omitempty" name:"ConfigurationRisk"`

	// Scan tasks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ScanTask *uint64 `json:"ScanTask,omitnil,omitempty" name:"ScanTask"`

	// Domain name
	// Note: This field may return·null, indicating that no valid values can be obtained.
	SubDomain *string `json:"SubDomain,omitnil,omitempty" name:"SubDomain"`

	// Resolved IP
	// Note: This field may return·null, indicating that no valid values can be obtained.
	SeverIp []*string `json:"SeverIp,omitnil,omitempty" name:"SeverIp"`

	// Bot access requests
	// Note: This field may return·null, indicating that no valid values can be obtained.
	BotCount *uint64 `json:"BotCount,omitnil,omitempty" name:"BotCount"`

	// Weak password risks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	WeakPassword *uint64 `json:"WeakPassword,omitnil,omitempty" name:"WeakPassword"`

	// Content risks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	WebContentRisk *uint64 `json:"WebContentRisk,omitnil,omitempty" name:"WebContentRisk"`

	// Tags
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// The type of associated instances.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// Member ID
	// Note: This field may return·null, indicating that no valid values can be obtained.
	MemberId *string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// CC attacks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	CCAttack *int64 `json:"CCAttack,omitnil,omitempty" name:"CCAttack"`

	// Web attack
	// Note: This field may return·null, indicating that no valid values can be obtained.
	WebAttack *int64 `json:"WebAttack,omitnil,omitempty" name:"WebAttack"`

	// Services exposed to risks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ServiceRisk *uint64 `json:"ServiceRisk,omitnil,omitempty" name:"ServiceRisk"`

	// Whether it's a newly-added asset. Values: `1` (Yes), `0` (No)
	// Note: This field may return·null, indicating that no valid values can be obtained.
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil,omitempty" name:"IsNewAsset"`

	// Random third-level domain name of the asset pending ownership verification
	// Note: This field may return·null, indicating that no valid values can be obtained.
	VerifyDomain *string `json:"VerifyDomain,omitnil,omitempty" name:"VerifyDomain"`

	// TXT record of the asset pending ownership verification
	// Note: This field may return·null, indicating that no valid values can be obtained.
	VerifyTXTRecord *string `json:"VerifyTXTRecord,omitnil,omitempty" name:"VerifyTXTRecord"`

	// Ownership verification status of the asset. `0`: Pending verification; `1`: Verified; `2`: Verifying; `3`: TXT record verification failed; `4`: Human verification failed.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	VerifyStatus *int64 `json:"VerifyStatus,omitnil,omitempty" name:"VerifyStatus"`

	// u200cNumber of bot attacks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	BotAccessCount *int64 `json:"BotAccessCount,omitnil,omitempty" name:"BotAccessCount"`
}

type Filter struct {
	// Max number of returned results
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Query offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Sorting order. Values: `asc` (ascending), `desc` (descending).
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// Specify the field used for sorting
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// Filtered columns and content
	Filters []*WhereFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Start time of the query period. 
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of the query period.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type FilterDataObject struct {
	// Filter value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// Filter name
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type IpAssetListVO struct {
	// Asset ID
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// Asset name
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// Asset type
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// Region
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// CFW status
	// Note: This field may return·null, indicating that no valid values can be obtained.
	CFWStatus *uint64 `json:"CFWStatus,omitnil,omitempty" name:"CFWStatus"`

	// Asset creation time
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetCreateTime *string `json:"AssetCreateTime,omitnil,omitempty" name:"AssetCreateTime"`

	// Public IP
	// Note: This field may return·null, indicating that no valid values can be obtained.
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// Public IP type
	// Note: This field may return·null, indicating that no valid values can be obtained.
	PublicIpType *uint64 `json:"PublicIpType,omitnil,omitempty" name:"PublicIpType"`


	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// VPC name
	// Note: This field may return·null, indicating that no valid values can be obtained.
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// appid
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User `uin`
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Name
	// Note: This field may return·null, indicating that no valid values can be obtained.
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// Whether it's a critical asset
	// Note: This field may return·null, indicating that no valid values can be obtained.
	IsCore *uint64 `json:"IsCore,omitnil,omitempty" name:"IsCore"`

	// Whether it's a cloud asset
	// Note: This field may return·null, indicating that no valid values can be obtained.
	IsCloud *uint64 `json:"IsCloud,omitnil,omitempty" name:"IsCloud"`

	// Number of network attacks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Attack *uint64 `json:"Attack,omitnil,omitempty" name:"Attack"`

	// Number of network access requests
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Access *uint64 `json:"Access,omitnil,omitempty" name:"Access"`

	// Number of blocked attacks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Intercept *uint64 `json:"Intercept,omitnil,omitempty" name:"Intercept"`

	// Inbound bandwidth
	// Note: This field may return·null, indicating that no valid values can be obtained.
	InBandwidth *string `json:"InBandwidth,omitnil,omitempty" name:"InBandwidth"`

	// Outbound bandwidthtraffic peak bandwidth (bps)
	// Note: This field may return·null, indicating that no valid values can be obtained.
	OutBandwidth *string `json:"OutBandwidth,omitnil,omitempty" name:"OutBandwidth"`

	// Inbound traffic
	// Note: This field may return·null, indicating that no valid values can be obtained.
	InFlow *string `json:"InFlow,omitnil,omitempty" name:"InFlow"`

	// Outbound traffic
	// Note: This field may return·null, indicating that no valid values can be obtained.
	OutFlow *string `json:"OutFlow,omitnil,omitempty" name:"OutFlow"`

	// Last scan time
	// Note: This field may return·null, indicating that no valid values can be obtained.
	LastScanTime *string `json:"LastScanTime,omitnil,omitempty" name:"LastScanTime"`

	// Port risks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	PortRisk *uint64 `json:"PortRisk,omitnil,omitempty" name:"PortRisk"`

	// Vulnerabilities
	// Note: This field may return·null, indicating that no valid values can be obtained.
	VulnerabilityRisk *uint64 `json:"VulnerabilityRisk,omitnil,omitempty" name:"VulnerabilityRisk"`

	// Configuration risks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ConfigurationRisk *uint64 `json:"ConfigurationRisk,omitnil,omitempty" name:"ConfigurationRisk"`

	// Scan tasks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ScanTask *uint64 `json:"ScanTask,omitnil,omitempty" name:"ScanTask"`

	// Weak passwords
	// Note: This field may return·null, indicating that no valid values can be obtained.
	WeakPassword *uint64 `json:"WeakPassword,omitnil,omitempty" name:"WeakPassword"`

	// Content risks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	WebContentRisk *uint64 `json:"WebContentRisk,omitnil,omitempty" name:"WebContentRisk"`

	// Tags
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// EIP ID
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AddressId *string `json:"AddressId,omitnil,omitempty" name:"AddressId"`

	// Member ID
	// Note: This field may return·null, indicating that no valid values can be obtained.
	MemberId *string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// Risk exposure
	// Note: This field may return·null, indicating that no valid values can be obtained.
	RiskExposure *int64 `json:"RiskExposure,omitnil,omitempty" name:"RiskExposure"`

	// Whether it's a newly-added asset. Values: `1` (Yes), `0` (No)
	// Note: This field may return·null, indicating that no valid values can be obtained.
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil,omitempty" name:"IsNewAsset"`

	// Asset ownership verification status. `0`: Pending verification; `1`: Verified; `2`: Verifying; `3` and above: Verification failed
	// Note: This field may return·null, indicating that no valid values can be obtained.
	VerifyStatus *int64 `json:"VerifyStatus,omitnil,omitempty" name:"VerifyStatus"`
}

// Predefined struct for user
type ModifyRiskCenterRiskStatusRequestParams struct {
	// Data of risk assets
	RiskStatusKeys []*RiskCenterStatusKey `json:"RiskStatusKeys,omitnil,omitempty" name:"RiskStatusKeys"`

	// Specify how you want to change the risk status. `1`: Change to Handled, `2`: Change to Ignored; `3`: Remove from Handled; `4`: Remove from Ignored
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Risk type. `0`: Port risk; `1`: Vulnerability; `2`: Weak password; `3`: Website content risk; `4`: Configuration risk; `5`: Risk services
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type ModifyRiskCenterRiskStatusRequest struct {
	*tchttp.BaseRequest
	
	// Data of risk assets
	RiskStatusKeys []*RiskCenterStatusKey `json:"RiskStatusKeys,omitnil,omitempty" name:"RiskStatusKeys"`

	// Specify how you want to change the risk status. `1`: Change to Handled, `2`: Change to Ignored; `3`: Remove from Handled; `4`: Remove from Ignored
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Risk type. `0`: Port risk; `1`: Vulnerability; `2`: Weak password; `3`: Website content risk; `4`: Configuration risk; `5`: Risk services
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *ModifyRiskCenterRiskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRiskCenterRiskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RiskStatusKeys")
	delete(f, "Status")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRiskCenterRiskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRiskCenterRiskStatusResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRiskCenterRiskStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRiskCenterRiskStatusResponseParams `json:"Response"`
}

func (r *ModifyRiskCenterRiskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRiskCenterRiskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PortViewPortRisk struct {
	// Affected assets
	NoHandleCount *int64 `json:"NoHandleCount,omitnil,omitempty" name:"NoHandleCount"`

	// Risk level
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// Network protocol
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Components
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// Port
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Last detected 
	RecentTime *string `json:"RecentTime,omitnil,omitempty" name:"RecentTime"`

	// First detected
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// Suggested action. `0`: Keep as it is; `1`: Block access requests; `2`: Block the port
	Suggestion *uint64 `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// Status of the risk. `0`: Not handled, `1`: Handled; `2`: Ignored
	AffectAssetCount *string `json:"AffectAssetCount,omitnil,omitempty" name:"AffectAssetCount"`

	// Unique ID of the asset
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Asset sub-category
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// Data entry key
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// User AppId
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User name.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// User `uin`
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Service
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`
}

type PublicIpDomainListKey struct {
	// IP/Domain
	Asset *string `json:"Asset,omitnil,omitempty" name:"Asset"`
}

type ReportItemKey struct {
	// List of report IDs.
	TaskLogList []*string `json:"TaskLogList,omitnil,omitempty" name:"TaskLogList"`
}

type ReportTaskIdList struct {
	// List of task IDs
	TaskIdList []*string `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`

	// User AppId
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

type RiskCenterStatusKey struct {
	// Risk ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// User AppId
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// Public IP/domain name
	PublicIPDomain *string `json:"PublicIPDomain,omitnil,omitempty" name:"PublicIPDomain"`

	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type ScanTaskInfo struct {
	// Task ID Id
	// Note: This field may return·null, indicating that no valid values can be obtained.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task name
	// Note: This field may return·null, indicating that no valid values can be obtained.
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Task status. `1`: Pending start; `2`: Scanning; `3`: Failed; `4`: Completed
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Task progress
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// Displayed time point
	// Note: This field may return·null, indicating that no valid values can be obtained.
	TaskTime *string `json:"TaskTime,omitnil,omitempty" name:"TaskTime"`

	// Report ID
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ReportId *string `json:"ReportId,omitnil,omitempty" name:"ReportId"`

	// Report name
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ReportName *string `json:"ReportName,omitnil,omitempty" name:"ReportName"`

	// Task type. `0`: Scheduled task, `1`: Scan immediately; `2`: Scanned at the specified time; `3`: Custom. 
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ScanPlan *int64 `json:"ScanPlan,omitnil,omitempty" name:"ScanPlan"`

	// Number of associated assets
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetCount *int64 `json:"AssetCount,omitnil,omitempty" name:"AssetCount"`

	// User AppId
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User UIN
	// Note: This field may return·null, indicating that no valid values can be obtained.
	UIN *string `json:"UIN,omitnil,omitempty" name:"UIN"`

	// User name.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`
}

type ScanTaskInfoList struct {
	// Task name
	// Note: This field may return·null, indicating that no valid values can be obtained.
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Start time of the task
	// Note: This field may return·null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Task end time
	// Note: This field may return·null, indicating that no valid values can be obtained.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// CRON format
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ScanPlanContent *string `json:"ScanPlanContent,omitnil,omitempty" name:"ScanPlanContent"`

	// Task type. `0`: Scheduled task, `1`: Scan immediately; `2`: Scanned at the specified time; `3`: Custom.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// Creation time
	// Note: u200dThis field may return `null`, indicating that no valid values can be obtained.
	InsertTime *string `json:"InsertTime,omitnil,omitempty" name:"InsertTime"`

	// Task ID
	// Note: This field may return·null, indicating that no valid values can be obtained.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Custom list of assets to scan
	// Note: This field may return·null, indicating that no valid values can be obtained.
	SelfDefiningAssets []*string `json:"SelfDefiningAssets,omitnil,omitempty" name:"SelfDefiningAssets"`

	// Estimated period to complete the task
	// Note: This field may return·null, indicating that no valid values can be obtained.
	PredictTime *int64 `json:"PredictTime,omitnil,omitempty" name:"PredictTime"`

	// Estimated completion time of the task
	// Note: This field may return·null, indicating that no valid values can be obtained.
	PredictEndTime *string `json:"PredictEndTime,omitnil,omitempty" name:"PredictEndTime"`

	// Number of reports
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ReportNumber *int64 `json:"ReportNumber,omitnil,omitempty" name:"ReportNumber"`

	// Number of assets
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetNumber *int64 `json:"AssetNumber,omitnil,omitempty" name:"AssetNumber"`

	// Scanning status. `0`: (default) Not scanned; `1`: Scanning; `2`: Scan completed; `3`: Error while scanning; `4`: Scanning stopped
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ScanStatus *int64 `json:"ScanStatus,omitnil,omitempty" name:"ScanStatus"`

	// Task progress
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Percent *float64 `json:"Percent,omitnil,omitempty" name:"Percent"`

	// port/poc/weakpass/webcontent/configrisk
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ScanItem *string `json:"ScanItem,omitnil,omitempty" name:"ScanItem"`

	// Values: `0` (Scan all); `1` (Scan specific assets); `2` (Scan all expect the specified assets); `3` (Custom assets).
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ScanAssetType *int64 `json:"ScanAssetType,omitnil,omitempty" name:"ScanAssetType"`

	// VSS subtask ID
	// Note: This field may return·null, indicating that no valid values can be obtained.
	VSSTaskId *string `json:"VSSTaskId,omitnil,omitempty" name:"VSSTaskId"`

	// CSPM subtask ID
	// Note: This field may return·null, indicating that no valid values can be obtained.
	CSPMTaskId *string `json:"CSPMTaskId,omitnil,omitempty" name:"CSPMTaskId"`

	// CWPP vulnerability scan task IDHost missed scan subtask id
	// Note: This field may return·null, indicating that no valid values can be obtained.
	CWPPOCId *string `json:"CWPPOCId,omitnil,omitempty" name:"CWPPOCId"`

	// CWPP baseline check task ID
	// Note: This field may return·null, indicating that no valid values can be obtained.
	CWPBlId *string `json:"CWPBlId,omitnil,omitempty" name:"CWPBlId"`

	// VSS task progess
	// Note: This field may return·null, indicating that no valid values can be obtained.
	VSSTaskProcess *int64 `json:"VSSTaskProcess,omitnil,omitempty" name:"VSSTaskProcess"`

	// CSPM task progress
	// Note: This field may return·null, indicating that no valid values can be obtained.
	CSPMTaskProcess *uint64 `json:"CSPMTaskProcess,omitnil,omitempty" name:"CSPMTaskProcess"`

	// CWPP vulnerability scan task progress
	// Note: This field may return·null, indicating that no valid values can be obtained.
	CWPPOCProcess *int64 `json:"CWPPOCProcess,omitnil,omitempty" name:"CWPPOCProcess"`

	// CWPP baseline check task progress
	// Note: This field may return·null, indicating that no valid values can be obtained.
	CWPBlProcess *uint64 `json:"CWPBlProcess,omitnil,omitempty" name:"CWPBlProcess"`


	ErrorCode *int64 `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// Exception information
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ErrorInfo *string `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// Day of the month to start the scheduled task
	// Note: This field may return·null, indicating that no valid values can be obtained.
	StartDay *int64 `json:"StartDay,omitnil,omitempty" name:"StartDay"`

	// Scan frequency in days. `1`: Every day; `7`: Every seven days; `30`: Every 30 days; `0`: Scan once only
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Frequency *int64 `json:"Frequency,omitnil,omitempty" name:"Frequency"`

	// Number of completed tasks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	CompleteNumber *int64 `json:"CompleteNumber,omitnil,omitempty" name:"CompleteNumber"`

	// Number of scanned assets
	// Note: This field may return·null, indicating that no valid values can be obtained.
	CompleteAssetNumber *int64 `json:"CompleteAssetNumber,omitnil,omitempty" name:"CompleteAssetNumber"`

	// Number of risks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	RiskCount *int64 `json:"RiskCount,omitnil,omitempty" name:"RiskCount"`

	// Assets
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Assets []*TaskAssetObject `json:"Assets,omitnil,omitempty" name:"Assets"`

	// User `Appid`
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User UIN
	// Note: This field may return·null, indicating that no valid values can be obtained.
	UIN *string `json:"UIN,omitnil,omitempty" name:"UIN"`

	// User name.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Scan task mode: `0` (Standard), `1` (Quick), `2` (Advanced). 
	// Note: This field may return·null, indicating that no valid values can be obtained.
	TaskMode *int64 `json:"TaskMode,omitnil,omitempty" name:"TaskMode"`

	// Source of scanning request
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ScanFrom *string `json:"ScanFrom,omitnil,omitempty" name:"ScanFrom"`

	// Whether it's a limited-time free health check. `0`: No; `1`: Yes
	// Note: This field may return·null, indicating that no valid values can be obtained.
	IsFree *int64 `json:"IsFree,omitnil,omitempty" name:"IsFree"`

	// Whether the user is authorized to delete this task. `1` :Yes; `0`: No. It's available for multi-account management.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	IsDelete *int64 `json:"IsDelete,omitnil,omitempty" name:"IsDelete"`

	// Source of the task. `0`: Default, `1`: Assistant; `2`: Health check
	// Note: This field may return·null, indicating that no valid values can be obtained.
	SourceType *int64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`
}

type ServerRisk struct {
	// Service tag
	ServiceTag *string `json:"ServiceTag,omitnil,omitempty" name:"ServiceTag"`

	// Port
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Affected assets
	AffectAsset *string `json:"AffectAsset,omitnil,omitempty" name:"AffectAsset"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Asset type
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Risk level
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// Network protocol
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// Components
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// Service
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// Last detected 
	RecentTime *string `json:"RecentTime,omitnil,omitempty" name:"RecentTime"`

	// First detected
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// Risk Details
	// Note: This field may return·null, indicating that no valid values can be obtained.
	RiskDetails *string `json:"RiskDetails,omitnil,omitempty" name:"RiskDetails"`

	// Handling suggestion
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// Status of the risk. `0`: Not handled, `1`: Handled; `2`: Ignored
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Unique ID of the asset
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// User `appid`
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User name.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// User `uin`
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Service snapshot
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ServiceSnapshot *string `json:"ServiceSnapshot,omitnil,omitempty" name:"ServiceSnapshot"`

	// Service access URL.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Data entry key
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// List of risks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	RiskList []*ServerRiskSuggestion `json:"RiskList,omitnil,omitempty" name:"RiskList"`

	// List of fix suggestions 
	// Note: This field may return·null, indicating that no valid values can be obtained.
	SuggestionList []*ServerRiskSuggestion `json:"SuggestionList,omitnil,omitempty" name:"SuggestionList"`

	// HTTP response code
	// Note: This field may return·null, indicating that no valid values can be obtained.
	StatusCode *string `json:"StatusCode,omitnil,omitempty" name:"StatusCode"`
}

type ServerRiskSuggestion struct {
	// Risk title
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// Risk details
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Body *string `json:"Body,omitnil,omitempty" name:"Body"`
}

// Predefined struct for user
type StopRiskCenterTaskRequestParams struct {
	// List of task IDs
	TaskIdList []*TaskIdListKey `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`
}

type StopRiskCenterTaskRequest struct {
	*tchttp.BaseRequest
	
	// List of task IDs
	TaskIdList []*TaskIdListKey `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`
}

func (r *StopRiskCenterTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopRiskCenterTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopRiskCenterTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopRiskCenterTaskResponseParams struct {
	// `0`: Operation succeeded; Others: failed
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopRiskCenterTaskResponse struct {
	*tchttp.BaseResponse
	Response *StopRiskCenterTaskResponseParams `json:"Response"`
}

func (r *StopRiskCenterTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopRiskCenterTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubnetAsset struct {
	// appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Asset ID
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// Asset name
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// Region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// VPC name
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// Tags
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// User name
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// CIDR block
	CIDR *string `json:"CIDR,omitnil,omitempty" name:"CIDR"`

	// Availability zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Number of CVMs
	CVM *int64 `json:"CVM,omitnil,omitempty" name:"CVM"`

	// Number of available IPs
	AvailableIp *int64 `json:"AvailableIp,omitnil,omitempty" name:"AvailableIp"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Configuration risks
	ConfigureRisk *int64 `json:"ConfigureRisk,omitnil,omitempty" name:"ConfigureRisk"`

	// Number of tasks.
	ScanTask *int64 `json:"ScanTask,omitnil,omitempty" name:"ScanTask"`

	// Last scan time
	LastScanTime *string `json:"LastScanTime,omitnil,omitempty" name:"LastScanTime"`

	// Whether it's a critical asset
	// Note: This field may return·null, indicating that no valid values can be obtained.
	IsCore *uint64 `json:"IsCore,omitnil,omitempty" name:"IsCore"`

	// Whether it's a newly-added asset. Values: `1` (Yes), `0` (No)
	// Note: This field may return·null, indicating that no valid values can be obtained.
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil,omitempty" name:"IsNewAsset"`
}

type Tag struct {
	// Tag name.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Tag value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Tags struct {
	// None
	// Note: This field may return·null, indicating that no valid values can be obtained.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// None
	// Note: This field may return·null, indicating that no valid values can be obtained.
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TaskAdvanceCFG struct {
	// Advanced vulnerability scan configuration
	VulRisk []*TaskCenterVulRiskInputParam `json:"VulRisk,omitnil,omitempty" name:"VulRisk"`

	// Advanced weak password check configuration
	WeakPwdRisk []*TaskCenterWeakPwdRiskInputParam `json:"WeakPwdRisk,omitnil,omitempty" name:"WeakPwdRisk"`

	// Advanced configuration risk scan configuration
	CFGRisk []*TaskCenterCFGRiskInputParam `json:"CFGRisk,omitnil,omitempty" name:"CFGRisk"`
}

type TaskAssetObject struct {
	// Asset name
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// 	Asset category
	// Note: This field may return·null, indicating that no valid values can be obtained.
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Asset sub-category
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// IP, domain name, asset ID, database ID, and more
	Asset *string `json:"Asset,omitnil,omitempty" name:"Asset"`

	// Region
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// The ID specific for an asset synched from another cloud platform
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Arn *string `json:"Arn,omitnil,omitempty" name:"Arn"`
}

type TaskCenterCFGRiskInputParam struct {
	// Check item ID
	ItemId *string `json:"ItemId,omitnil,omitempty" name:"ItemId"`

	// Whether to enable. `0`: no, `1`: yes.
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// Resource type
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`
}

type TaskCenterVulRiskInputParam struct {
	// Risk ID
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`

	// Whether to enable. `0`: no, `1`: yes.
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type TaskCenterWeakPwdRiskInputParam struct {
	// Check item ID
	CheckItemId *int64 `json:"CheckItemId,omitnil,omitempty" name:"CheckItemId"`

	// Whether to enable. `0`: no, `1`: yes.
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type TaskIdListKey struct {
	// Task ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type TaskLogInfo struct {
	// Report name
	// Note: This field may return·null, indicating that no valid values can be obtained.
	TaskLogName *string `json:"TaskLogName,omitnil,omitempty" name:"TaskLogName"`

	// Report ID.
	TaskLogId *string `json:"TaskLogId,omitnil,omitempty" name:"TaskLogId"`

	// Number of associated assets
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AssetsNumber *int64 `json:"AssetsNumber,omitnil,omitempty" name:"AssetsNumber"`

	// Number of risks
	// Note: This field may return·null, indicating that no valid values can be obtained.
	RiskNumber *int64 `json:"RiskNumber,omitnil,omitempty" name:"RiskNumber"`

	// Report generation time
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// Task status. `0`: Initial value; `1`: Scanning; `2`: Completed; `3`: Failed; `4`: Stopped; `5`: Paused; `6`: Retried
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Name of the associated task
	// Note: This field may return·null, indicating that no valid values can be obtained.
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// Scan start time
	// Note: This field may return·null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Scan task ID
	// Note: This field may return·null, indicating that no valid values can be obtained.
	TaskCenterTaskId *string `json:"TaskCenterTaskId,omitnil,omitempty" name:"TaskCenterTaskId"`

	// User AppId
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User UIN
	// Note: This field may return·null, indicating that no valid values can be obtained.
	UIN *string `json:"UIN,omitnil,omitempty" name:"UIN"`

	// User name.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Report type: `1`: Health check report; `2`: Daily report; `3`: Weekly report; `4`: Monthly report
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ReportType *int64 `json:"ReportType,omitnil,omitempty" name:"ReportType"`

	// Report template ID
	// Note: This field may return·null, indicating that no valid values can be obtained.
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type TaskLogURL struct {
	// Temp download URL for the report
	// Note: This field may return·null, indicating that no valid values can be obtained.
	URL *string `json:"URL,omitnil,omitempty" name:"URL"`

	// Task report ID
	// Note: This field may return·null, indicating that no valid values can be obtained.
	LogId *string `json:"LogId,omitnil,omitempty" name:"LogId"`

	// Task report name
	// Note: This field may return·null, indicating that no valid values can be obtained.
	TaskLogName *string `json:"TaskLogName,omitnil,omitempty" name:"TaskLogName"`

	// APP ID
	// Note: This field may return·null, indicating that no valid values can be obtained.
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

type VULRiskAdvanceCFGList struct {
	// Risk ID
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`

	// Vulnerability name
	VULName *string `json:"VULName,omitnil,omitempty" name:"VULName"`

	// Risk level
	RiskLevel *string `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// Source of the check task
	CheckFrom *string `json:"CheckFrom,omitnil,omitempty" name:"CheckFrom"`

	// Whether it's enabled. `1`: Enable; `0`: Disabled
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// Risk type.
	VULType *string `json:"VULType,omitnil,omitempty" name:"VULType"`

	// Affected versions
	ImpactVersion *string `json:"ImpactVersion,omitnil,omitempty" name:"ImpactVersion"`

	// CVE number
	// Note: This field may return·null, indicating that no valid values can be obtained.
	CVE *string `json:"CVE,omitnil,omitempty" name:"CVE"`

	// Vulnerability tag
	VULTag []*string `json:"VULTag,omitnil,omitempty" name:"VULTag"`

	// Fix solution
	// Note: This field may return·null, indicating that no valid values can be obtained.
	FixMethod []*string `json:"FixMethod,omitnil,omitempty" name:"FixMethod"`

	// Disclosure time
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ReleaseTime *string `json:"ReleaseTime,omitnil,omitempty" name:"ReleaseTime"`

	// Whether it's an emergency vulnerability. Values: `1` (emergency vulnerability); `0` (non-emergency vulnerability
	// Note: This field may return·null, indicating that no valid values can be obtained.
	EMGCVulType *int64 `json:"EMGCVulType,omitnil,omitempty" name:"EMGCVulType"`

	// Vulnerability description
	// Note: This field may return·null, indicating that no valid values can be obtained.
	VULDescribe *string `json:"VULDescribe,omitnil,omitempty" name:"VULDescribe"`

	// Affected components
	// Note: This field may return·null, indicating that no valid values can be obtained.
	ImpactComponent *string `json:"ImpactComponent,omitnil,omitempty" name:"ImpactComponent"`
}

type VULViewVULRisk struct {
	// Port
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// Affected assets
	NoHandleCount *int64 `json:"NoHandleCount,omitnil,omitempty" name:"NoHandleCount"`

	// Risk level
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// Components
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// Last detected 
	RecentTime *string `json:"RecentTime,omitnil,omitempty" name:"RecentTime"`

	// First detected
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// Status of the risk. `0`: Not handled, `1`: Handled; `2`: Ignored
	AffectAssetCount *uint64 `json:"AffectAssetCount,omitnil,omitempty" name:"AffectAssetCount"`

	// Unique ID of the asset
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Asset sub-category
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// Frontend index
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// Vulnerability type
	VULType *string `json:"VULType,omitnil,omitempty" name:"VULType"`

	// Vulnerability name
	VULName *string `json:"VULName,omitnil,omitempty" name:"VULName"`

	// CVE number
	CVE *string `json:"CVE,omitnil,omitempty" name:"CVE"`

	// Description
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// Pay load
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// Affected components
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// Reference information of the vulnerability
	References *string `json:"References,omitnil,omitempty" name:"References"`

	// Affected version
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// Vulnerability URL
	VULURL *string `json:"VULURL,omitnil,omitempty" name:"VULURL"`

	// User name.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// User `appid`
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User `uin`
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Fix suggestion
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Fix *string `json:"Fix,omitnil,omitempty" name:"Fix"`

	// Whether it's an emergency vulnerability. Values: `1` (emergency vulnerability); `0` (non-emergency vulnerability
	// Note: This field may return·null, indicating that no valid values can be obtained.
	EMGCVulType *int64 `json:"EMGCVulType,omitnil,omitempty" name:"EMGCVulType"`
}

type Vpc struct {
	// Subnet (32-bit mask)
	Subnet *uint64 `json:"Subnet,omitnil,omitempty" name:"Subnet"`

	// Connected VPC (32-bit mask)
	ConnectedVpc *uint64 `json:"ConnectedVpc,omitnil,omitempty" name:"ConnectedVpc"`

	// Asset ID
	AssetId *string `json:"AssetId,omitnil,omitempty" name:"AssetId"`

	// Region
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// CVM (only 32-bit)
	CVM *uint64 `json:"CVM,omitnil,omitempty" name:"CVM"`

	// Tags
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Tag []*Tag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// DNS domain
	// Note: This field may return·null, indicating that no valid values can be obtained.
	DNS []*string `json:"DNS,omitnil,omitempty" name:"DNS"`

	// Asset name
	AssetName *string `json:"AssetName,omitnil,omitempty" name:"AssetName"`

	// CIDR block
	CIDR *string `json:"CIDR,omitnil,omitempty" name:"CIDR"`

	// Asset creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// User name
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// Whether it's a newly-added asset. Values: `1` (Yes), `0` (No)
	// Note: This field may return·null, indicating that no valid values can be obtained.
	IsNewAsset *uint64 `json:"IsNewAsset,omitnil,omitempty" name:"IsNewAsset"`

	// Whether it's a critical asset. `1`: Yes; `2`: No
	// Note: This field may return·null, indicating that no valid values can be obtained.
	IsCore *uint64 `json:"IsCore,omitnil,omitempty" name:"IsCore"`
}

type WebsiteRisk struct {
	// Affected assets
	AffectAsset *string `json:"AffectAsset,omitnil,omitempty" name:"AffectAsset"`

	// Risk level
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// Last detected
	RecentTime *string `json:"RecentTime,omitnil,omitempty" name:"RecentTime"`

	// First detected
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// Status of the risk. `0`: Not handled, `1`: Handled; `2`: Ignored
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Unique ID of the asset
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Frontend index
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// User `appid`
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// User name.
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Nick *string `json:"Nick,omitnil,omitempty" name:"Nick"`

	// User `uin`
	// Note: This field may return·null, indicating that no valid values can be obtained.
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// URL of the risk
	URL *string `json:"URL,omitnil,omitempty" name:"URL"`

	// URL of the risk file
	URLPath *string `json:"URLPath,omitnil,omitempty" name:"URLPath"`

	// Instance type
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Check type.
	DetectEngine *string `json:"DetectEngine,omitnil,omitempty" name:"DetectEngine"`

	// Result description.
	ResultDescribe *string `json:"ResultDescribe,omitnil,omitempty" name:"ResultDescribe"`

	// Source URL
	SourceURL *string `json:"SourceURL,omitnil,omitempty" name:"SourceURL"`

	// Source file URL
	SourceURLPath *string `json:"SourceURLPath,omitnil,omitempty" name:"SourceURLPath"`
}

type WhereFilter struct {
	// Filter item
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Filter value
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	//  
	// `1`: =; `2`: >; `3`: <; `4`: ≥; `5`: ≤; `6`: ≠;
	// `7`: Exact match; `9`: Fuzzy match; `13`: Non-fuzzy match; `14`: AND
	OperatorType *int64 `json:"OperatorType,omitnil,omitempty" name:"OperatorType"`
}