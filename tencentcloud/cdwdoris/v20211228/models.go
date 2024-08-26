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

package v20211228

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AccountDetailInfo struct {
	// Username
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Host name or IP address, which indicates the host to which the user belongs.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// User description information
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserDescription *string `json:"UserDescription,omitnil,omitempty" name:"UserDescription"`
}

type AttachCBSSpec struct {
	// Node disk type, such as CLOUD_SSD"\"CLOUD_PREMIUM
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Disk capacity, in GB
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// Total number of disks
	DiskCount *int64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`

	// Description
	DiskDesc *string `json:"DiskDesc,omitnil,omitempty" name:"DiskDesc"`
}

type BackUpJobDisplay struct {
	// Backup instance ID
	JobId *int64 `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Backup instance name
	Snapshot *string `json:"Snapshot,omitnil,omitempty" name:"Snapshot"`

	// Backup data volume
	BackUpSize *int64 `json:"BackUpSize,omitnil,omitempty" name:"BackUpSize"`

	// Backup single replica data volume
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackUpSingleSize *int64 `json:"BackUpSingleSize,omitnil,omitempty" name:"BackUpSingleSize"`

	// Instance creation time
	BackUpTime *string `json:"BackUpTime,omitnil,omitempty" name:"BackUpTime"`

	// Instance expiration time
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// Instance status
	JobStatus *string `json:"JobStatus,omitnil,omitempty" name:"JobStatus"`

	// 0: default; 1: one-time backup for the remote Doris
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupType *int64 `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// 0: default; 1: immediate backup; 2: migration
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupTimeType *int64 `json:"BackupTimeType,omitnil,omitempty" name:"BackupTimeType"`

	// Connection information of the remote Doris
	// Note: This field may return null, indicating that no valid values can be obtained.
	DorisSourceInfo *DorisSourceInfo `json:"DorisSourceInfo,omitnil,omitempty" name:"DorisSourceInfo"`

	// The value corresponding to the instance status
	// Note: This field may return null, indicating that no valid values can be obtained.
	JobStatusNum *int64 `json:"JobStatusNum,omitnil,omitempty" name:"JobStatusNum"`

	// Information about cos in the backup instance	
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupCosInfo *BackupCosInfo `json:"BackupCosInfo,omitnil,omitempty" name:"BackupCosInfo"`
}

type BackupCosInfo struct {
	// The cos bucket where the backup file is located.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CosBucket *string `json:"CosBucket,omitnil,omitempty" name:"CosBucket"`

	// The full cos path where the backup file is located.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CosPath *string `json:"CosPath,omitnil,omitempty" name:"CosPath"`

	// Backup file name
	// Note: This field may return null, indicating that no valid values can be obtained.
	SnapShotPath *string `json:"SnapShotPath,omitnil,omitempty" name:"SnapShotPath"`
}

type BackupStatus struct {
	// Backup task ID
	JobId *int64 `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Snapshot name
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// Database name
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Status
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// Backup object
	BackupObjects *string `json:"BackupObjects,omitnil,omitempty" name:"BackupObjects"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Snapshot end time
	SnapshotFinishedTime *string `json:"SnapshotFinishedTime,omitnil,omitempty" name:"SnapshotFinishedTime"`

	// Upload end time
	UploadFinishedTime *string `json:"UploadFinishedTime,omitnil,omitempty" name:"UploadFinishedTime"`

	// End time
	FinishedTime *string `json:"FinishedTime,omitnil,omitempty" name:"FinishedTime"`

	// Unfinished tasks
	UnfinishedTasks *string `json:"UnfinishedTasks,omitnil,omitempty" name:"UnfinishedTasks"`

	// Progress
	Progress *string `json:"Progress,omitnil,omitempty" name:"Progress"`

	// Error message
	TaskErrMsg *string `json:"TaskErrMsg,omitnil,omitempty" name:"TaskErrMsg"`

	// Status
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Timeout information
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// Backup instance ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupJobId *int64 `json:"BackupJobId,omitnil,omitempty" name:"BackupJobId"`

	// The ID of the snapshoit corresponding to the instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type BackupTableContent struct {
	// Database
	// Note: This field may return null, indicating that no valid values can be obtained.
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Table
	// Note: This field may return null, indicating that no valid values can be obtained.
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// Total number of bytes in the table
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalBytes *int64 `json:"TotalBytes,omitnil,omitempty" name:"TotalBytes"`

	// Size of a single replica
	// Note: This field may return null, indicating that no valid values can be obtained.
	SingleReplicaBytes *string `json:"SingleReplicaBytes,omitnil,omitempty" name:"SingleReplicaBytes"`

	// Backup status
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupStatus *int64 `json:"BackupStatus,omitnil,omitempty" name:"BackupStatus"`

	// Error message of the backup
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupErrorMsg *string `json:"BackupErrorMsg,omitnil,omitempty" name:"BackupErrorMsg"`

	// Whether to bind the cold storage policy to the database and table
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsOpenCoolDown *bool `json:"IsOpenCoolDown,omitnil,omitempty" name:"IsOpenCoolDown"`
}

type BindUser struct {
	// Username
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Host information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`
}

// Predefined struct for user
type CancelBackupJobRequestParams struct {
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup instance ID to be canceled
	BackUpJobId *int64 `json:"BackUpJobId,omitnil,omitempty" name:"BackUpJobId"`
}

type CancelBackupJobRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Backup instance ID to be canceled
	BackUpJobId *int64 `json:"BackUpJobId,omitnil,omitempty" name:"BackUpJobId"`
}

func (r *CancelBackupJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelBackupJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackUpJobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelBackupJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelBackupJobResponseParams struct {
	// Error message
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelBackupJobResponse struct {
	*tchttp.BaseResponse
	Response *CancelBackupJobResponseParams `json:"Response"`
}

func (r *CancelBackupJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelBackupJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChargeProperties struct {
	// Billing type: PREPAID for prepayment, and POSTPAID_BY_HOUR for postpayment.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// Whether to automatically renew. 1 means automatic renewal is enabled.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// Billing duration
	// Note: This field may return null, indicating that no valid values can be obtained.
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// Billing time unit, and "m" means month, etc.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`
}

type ClusterConfigsHistory struct {
	// Configuration file's name
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Modified configuration file content; base64 code
	NewConfValue *string `json:"NewConfValue,omitnil,omitempty" name:"NewConfValue"`

	// Configuration file content before modification; base64 code
	OldConfValue *string `json:"OldConfValue,omitnil,omitempty" name:"OldConfValue"`

	// Reason for modification
	// Note: This field may return null, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Modification time
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// Modify sub-account ID
	UserUin *string `json:"UserUin,omitnil,omitempty" name:"UserUin"`
}

type ClusterConfigsInfoFromEMR struct {
	// Configuration file's name
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Related attribute information corresponding to the configuration files
	FileConf *string `json:"FileConf,omitnil,omitempty" name:"FileConf"`

	// Other attribute information corresponding to the configuration files
	KeyConf *string `json:"KeyConf,omitnil,omitempty" name:"KeyConf"`

	// Contents of the configuration files, base64 encoded
	OriParam *string `json:"OriParam,omitnil,omitempty" name:"OriParam"`

	// This is used to indicate whether the current configuration file has been modified without a restart, and reminds the user that a restart is needed.
	NeedRestart *int64 `json:"NeedRestart,omitnil,omitempty" name:"NeedRestart"`

	// Configuration file path
	// Note: This field may return null, indicating that no valid values can be obtained.
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`

	// kv value of a configuration file
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: FileKeyValues is deprecated.
	FileKeyValues *string `json:"FileKeyValues,omitnil,omitempty" name:"FileKeyValues"`

	// kv value of a configuration file
	// Note: This field may return null, indicating that no valid values can be obtained.
	FileKeyValuesNew []*ConfigKeyValue `json:"FileKeyValuesNew,omitnil,omitempty" name:"FileKeyValuesNew"`
}

type Column struct {
	// Column name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Column type
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Aggregation type: When the table is an aggregation model (AGG_KEY), the column with the aggregation type is set as the metric column, and other columns are dimension columns. Aggregation type: ●SUM: sum; the values of multiple rows are accumulated. ●REPLACE: replacement; the values in the next batch of data will replace the values in the previously imported rows. ●MAX: retain the maximum value.
	//  ●MIN: retain the minimum value. ●REPLACE_IF_NOT_NULL: non-null values replacement. The difference from REPLACE is that null values are not replaced. ●HLL_UNION: aggregation method for HLL type columns, which is aggregated by HyperLogLog algorithm. ●BITMAP_UNION: aggregation method for BIMTAP type columns; bitmap union aggregation.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AggType *string `json:"AggType,omitnil,omitempty" name:"AggType"`

	// Whether the column value is allowed to be Null
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsNull *bool `json:"IsNull,omitnil,omitempty" name:"IsNull"`

	// Whether it is a Key column. The meaning of different data models:
	// ●DUP_KEY: The column specified afterwards is the sorting column.
	// ●AGG_KEY: The column specified afterwards is the dimension column.
	// ●UNI_KEY: The column specified afterward is the primary key column.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsKey *bool `json:"IsKey,omitnil,omitempty" name:"IsKey"`

	// Column's default value
	// Note: This field may return null, indicating that no valid values can be obtained.
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// Whether it is a partition column. The partition column must be a Key column.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsPartition *bool `json:"IsPartition,omitnil,omitempty" name:"IsPartition"`

	// Whether it is a bucket column. The bucket column of the aggregation model and primary key model must be Key columns, while the bucket column of the detail model can be any column.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsDistribution *bool `json:"IsDistribution,omitnil,omitempty" name:"IsDistribution"`

	// Whether it is an auto-increment column. Supported by TCHouse-D 2.1 version and later.
	// Limit:
	// 1. Only DUP_KEY and UNI_KEY model tables can contain auto-increment columns.
	// 2. A table can contain at most one auto-increment column.
	// 3. The type of the auto-increment column must be BIGINT type and cannot be null.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AutoInc *bool `json:"AutoInc,omitnil,omitempty" name:"AutoInc"`

	// Column description
	// Note: This field may return null, indicating that no valid values can be obtained.
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type ConfigKeyValue struct {
	// key
	// Note: This field may return null, indicating that no valid values can be obtained.
	KeyName *string `json:"KeyName,omitnil,omitempty" name:"KeyName"`

	// Value
	// Note: This field may return null, indicating that no valid values can be obtained.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// Notes
	// Note: This field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 1 indicates read-only, 2 indicates editable but undeletable, and 3 indicates deletable.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Display *int64 `json:"Display,omitnil,omitempty" name:"Display"`

	// 0 means not supported, and 1 means hot update is supported.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SupportHotUpdate *int64 `json:"SupportHotUpdate,omitnil,omitempty" name:"SupportHotUpdate"`
}

// Predefined struct for user
type CopyTableDatasRequestParams struct {
	// Resource ID, which is the TCHouse-D resource ID used for table creation.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Name of the database where the source table is located
	CopiedFromDb *string `json:"CopiedFromDb,omitnil,omitempty" name:"CopiedFromDb"`

	// Source table name
	CopiedFromTable *string `json:"CopiedFromTable,omitnil,omitempty" name:"CopiedFromTable"`

	// Name of the database where the target table is located
	CopyToDb *string `json:"CopyToDb,omitnil,omitempty" name:"CopyToDb"`

	// Target table name. If the table already exists, the structure of the source table and target table should be the same.
	CopyToTable *string `json:"CopyToTable,omitnil,omitempty" name:"CopyToTable"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`

	// Whether the data in the target table is overwritten. The default value is False.
	IsDataOverwrite *bool `json:"IsDataOverwrite,omitnil,omitempty" name:"IsDataOverwrite"`
}

type CopyTableDatasRequest struct {
	*tchttp.BaseRequest
	
	// Resource ID, which is the TCHouse-D resource ID used for table creation.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Name of the database where the source table is located
	CopiedFromDb *string `json:"CopiedFromDb,omitnil,omitempty" name:"CopiedFromDb"`

	// Source table name
	CopiedFromTable *string `json:"CopiedFromTable,omitnil,omitempty" name:"CopiedFromTable"`

	// Name of the database where the target table is located
	CopyToDb *string `json:"CopyToDb,omitnil,omitempty" name:"CopyToDb"`

	// Target table name. If the table already exists, the structure of the source table and target table should be the same.
	CopyToTable *string `json:"CopyToTable,omitnil,omitempty" name:"CopyToTable"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`

	// Whether the data in the target table is overwritten. The default value is False.
	IsDataOverwrite *bool `json:"IsDataOverwrite,omitnil,omitempty" name:"IsDataOverwrite"`
}

func (r *CopyTableDatasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyTableDatasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "CopiedFromDb")
	delete(f, "CopiedFromTable")
	delete(f, "CopyToDb")
	delete(f, "CopyToTable")
	delete(f, "UserName")
	delete(f, "PassWord")
	delete(f, "IsDataOverwrite")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CopyTableDatasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CopyTableDatasResponseParams struct {
	// Error message
	// Note: This field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CopyTableDatasResponse struct {
	*tchttp.BaseResponse
	Response *CopyTableDatasResponseParams `json:"Response"`
}

func (r *CopyTableDatasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyTableDatasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CosSourceInfo struct {
	// ID in cos authentication
	// Note: This field may return null, indicating that no valid values can be obtained.
	SecretId *string `json:"SecretId,omitnil,omitempty" name:"SecretId"`

	// Key in cos authentication
	// Note: This field may return null, indicating that no valid values can be obtained.
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`

	// Path in cos authentication
	// Note: This field may return null, indicating that no valid values can be obtained.
	CosPath *string `json:"CosPath,omitnil,omitempty" name:"CosPath"`
}

// Predefined struct for user
type CreateBackUpScheduleRequestParams struct {
	// Required to be uploaded when editing
	ScheduleId *int64 `json:"ScheduleId,omitnil,omitempty" name:"ScheduleId"`

	// Selected weeks, separated by commas.
	// Discarded: Use ScheduleInfo.
	WeekDays *string `json:"WeekDays,omitnil,omitempty" name:"WeekDays"`

	// Hour for executing the backup taskDiscarded: Use ScheduleInfo.
	ExecuteHour *int64 `json:"ExecuteHour,omitnil,omitempty" name:"ExecuteHour"`

	// Backup table list
	BackUpTables []*BackupTableContent `json:"BackUpTables,omitnil,omitempty" name:"BackUpTables"`

	// 0: default; 1: one-time backup for the remote Doris
	BackupType *int64 `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// Connection information of the remote doris cluster
	DorisSourceInfo *DorisSourceInfo `json:"DorisSourceInfo,omitnil,omitempty" name:"DorisSourceInfo"`

	// 0: default; 1: one-time backup; 2: remote backup
	BackupTimeType *int64 `json:"BackupTimeType,omitnil,omitempty" name:"BackupTimeType"`

	// 0: default; 1: immediate recovery after the backup is completed.
	RestoreType *int64 `json:"RestoreType,omitnil,omitempty" name:"RestoreType"`

	// 0: default; 1: connecting to COS using a custom key.
	AuthType *int64 `json:"AuthType,omitnil,omitempty" name:"AuthType"`

	// Cos certification information
	CosSourceInfo *CosSourceInfo `json:"CosSourceInfo,omitnil,omitempty" name:"CosSourceInfo"`
}

type CreateBackUpScheduleRequest struct {
	*tchttp.BaseRequest
	
	// Required to be uploaded when editing
	ScheduleId *int64 `json:"ScheduleId,omitnil,omitempty" name:"ScheduleId"`

	// Selected weeks, separated by commas.
	// Discarded: Use ScheduleInfo.
	WeekDays *string `json:"WeekDays,omitnil,omitempty" name:"WeekDays"`

	// Hour for executing the backup taskDiscarded: Use ScheduleInfo.
	ExecuteHour *int64 `json:"ExecuteHour,omitnil,omitempty" name:"ExecuteHour"`

	// Backup table list
	BackUpTables []*BackupTableContent `json:"BackUpTables,omitnil,omitempty" name:"BackUpTables"`

	// 0: default; 1: one-time backup for the remote Doris
	BackupType *int64 `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// Connection information of the remote doris cluster
	DorisSourceInfo *DorisSourceInfo `json:"DorisSourceInfo,omitnil,omitempty" name:"DorisSourceInfo"`

	// 0: default; 1: one-time backup; 2: remote backup
	BackupTimeType *int64 `json:"BackupTimeType,omitnil,omitempty" name:"BackupTimeType"`

	// 0: default; 1: immediate recovery after the backup is completed.
	RestoreType *int64 `json:"RestoreType,omitnil,omitempty" name:"RestoreType"`

	// 0: default; 1: connecting to COS using a custom key.
	AuthType *int64 `json:"AuthType,omitnil,omitempty" name:"AuthType"`

	// Cos certification information
	CosSourceInfo *CosSourceInfo `json:"CosSourceInfo,omitnil,omitempty" name:"CosSourceInfo"`
}

func (r *CreateBackUpScheduleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackUpScheduleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScheduleId")
	delete(f, "WeekDays")
	delete(f, "ExecuteHour")
	delete(f, "BackUpTables")
	delete(f, "BackupType")
	delete(f, "DorisSourceInfo")
	delete(f, "BackupTimeType")
	delete(f, "RestoreType")
	delete(f, "AuthType")
	delete(f, "CosSourceInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBackUpScheduleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackUpScheduleResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateBackUpScheduleResponse struct {
	*tchttp.BaseResponse
	Response *CreateBackUpScheduleResponseParams `json:"Response"`
}

func (r *CreateBackUpScheduleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackUpScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDatabaseRequestParams struct {
	// Resource ID, which is the TCHouse-D resource ID used for table creation.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Name of database to be created
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`

	// Database attributes. For keys with the same attributes, the priority of the table attribute is higher than that of the database attribute.
	Properties []*Property `json:"Properties,omitnil,omitempty" name:"Properties"`
}

type CreateDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// Resource ID, which is the TCHouse-D resource ID used for table creation.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Name of database to be created
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`

	// Database attributes. For keys with the same attributes, the priority of the table attribute is higher than that of the database attribute.
	Properties []*Property `json:"Properties,omitnil,omitempty" name:"Properties"`
}

func (r *CreateDatabaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DbName")
	delete(f, "UserName")
	delete(f, "PassWord")
	delete(f, "Properties")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDatabaseResponseParams struct {
	// Error message
	// Note: This field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDatabaseResponse struct {
	*tchttp.BaseResponse
	Response *CreateDatabaseResponseParams `json:"Response"`
}

func (r *CreateDatabaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceNewRequestParams struct {
	// Availability zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// FE specifications
	FeSpec *CreateInstanceSpec `json:"FeSpec,omitnil,omitempty" name:"FeSpec"`

	// BE specifications.
	BeSpec *CreateInstanceSpec `json:"BeSpec,omitnil,omitempty" name:"BeSpec"`

	// Whether it is highly available.
	HaFlag *bool `json:"HaFlag,omitnil,omitempty" name:"HaFlag"`

	// User VPCID
	UserVPCId *string `json:"UserVPCId,omitnil,omitempty" name:"UserVPCId"`

	// User subnet ID
	UserSubnetId *string `json:"UserSubnetId,omitnil,omitempty" name:"UserSubnetId"`

	// Product version number
	ProductVersion *string `json:"ProductVersion,omitnil,omitempty" name:"ProductVersion"`

	// Payment type
	ChargeProperties *ChargeProperties `json:"ChargeProperties,omitnil,omitempty" name:"ChargeProperties"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Database password
	DorisUserPwd *string `json:"DorisUserPwd,omitnil,omitempty" name:"DorisUserPwd"`

	// Tag list
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// High availability type:
	// 0 indicates non-high availability (only one FE, FeSpec.CreateInstanceSpec.Count=1),
	// 1 indicates read high availability (at least 3 FEs must be deployed, FeSpec.CreateInstanceSpec.Count>=3, and it must be an odd number),
	// 2 indicates read and write high availability (at least 5 FEs must be deployed, FeSpec.CreateInstanceSpec.Count>=5, and it must be an odd number).
	HaType *int64 `json:"HaType,omitnil,omitempty" name:"HaType"`

	// Whether the table name is case sensitive, 0 refers to sensitive, 1 refers to insensitive, compared in lowercase; 2 refers to insensitive, and the table name is changed to lowercase for storage.
	CaseSensitive *int64 `json:"CaseSensitive,omitnil,omitempty" name:"CaseSensitive"`

	// Whether to enable multi-availability zone.
	EnableMultiZones *bool `json:"EnableMultiZones,omitnil,omitempty" name:"EnableMultiZones"`

	// After the Multi-AZ is enabled, all user's Availability Zones and Subnets information are shown.
	UserMultiZoneInfos *NetworkInfo `json:"UserMultiZoneInfos,omitnil,omitempty" name:"UserMultiZoneInfos"`
}

type CreateInstanceNewRequest struct {
	*tchttp.BaseRequest
	
	// Availability zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// FE specifications
	FeSpec *CreateInstanceSpec `json:"FeSpec,omitnil,omitempty" name:"FeSpec"`

	// BE specifications.
	BeSpec *CreateInstanceSpec `json:"BeSpec,omitnil,omitempty" name:"BeSpec"`

	// Whether it is highly available.
	HaFlag *bool `json:"HaFlag,omitnil,omitempty" name:"HaFlag"`

	// User VPCID
	UserVPCId *string `json:"UserVPCId,omitnil,omitempty" name:"UserVPCId"`

	// User subnet ID
	UserSubnetId *string `json:"UserSubnetId,omitnil,omitempty" name:"UserSubnetId"`

	// Product version number
	ProductVersion *string `json:"ProductVersion,omitnil,omitempty" name:"ProductVersion"`

	// Payment type
	ChargeProperties *ChargeProperties `json:"ChargeProperties,omitnil,omitempty" name:"ChargeProperties"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Database password
	DorisUserPwd *string `json:"DorisUserPwd,omitnil,omitempty" name:"DorisUserPwd"`

	// Tag list
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// High availability type:
	// 0 indicates non-high availability (only one FE, FeSpec.CreateInstanceSpec.Count=1),
	// 1 indicates read high availability (at least 3 FEs must be deployed, FeSpec.CreateInstanceSpec.Count>=3, and it must be an odd number),
	// 2 indicates read and write high availability (at least 5 FEs must be deployed, FeSpec.CreateInstanceSpec.Count>=5, and it must be an odd number).
	HaType *int64 `json:"HaType,omitnil,omitempty" name:"HaType"`

	// Whether the table name is case sensitive, 0 refers to sensitive, 1 refers to insensitive, compared in lowercase; 2 refers to insensitive, and the table name is changed to lowercase for storage.
	CaseSensitive *int64 `json:"CaseSensitive,omitnil,omitempty" name:"CaseSensitive"`

	// Whether to enable multi-availability zone.
	EnableMultiZones *bool `json:"EnableMultiZones,omitnil,omitempty" name:"EnableMultiZones"`

	// After the Multi-AZ is enabled, all user's Availability Zones and Subnets information are shown.
	UserMultiZoneInfos *NetworkInfo `json:"UserMultiZoneInfos,omitnil,omitempty" name:"UserMultiZoneInfos"`
}

func (r *CreateInstanceNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "FeSpec")
	delete(f, "BeSpec")
	delete(f, "HaFlag")
	delete(f, "UserVPCId")
	delete(f, "UserSubnetId")
	delete(f, "ProductVersion")
	delete(f, "ChargeProperties")
	delete(f, "InstanceName")
	delete(f, "DorisUserPwd")
	delete(f, "Tags")
	delete(f, "HaType")
	delete(f, "CaseSensitive")
	delete(f, "EnableMultiZones")
	delete(f, "UserMultiZoneInfos")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceNewResponseParams struct {
	// Process ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Error message
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInstanceNewResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceNewResponseParams `json:"Response"`
}

func (r *CreateInstanceNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceSpec struct {
	// Specification name
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// Quantities
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// Cloud disk size
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`
}

// Predefined struct for user
type CreateTableRequestParams struct {
	// Resource ID, which is the TCHouse-D resource ID used for table creation.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The database where the table is located; if it does not exist, create one.
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Name of the table to be created
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// Table data model: 
	// AGG_KEY: aggregation model; 
	// UNI_KEY: primary key model; 
	// DUP_KEY: detail model
	KeysType *string `json:"KeysType,omitnil,omitempty" name:"KeysType"`

	// Column information of the table
	Columns []*Column `json:"Columns,omitnil,omitempty" name:"Columns"`

	// Bucket information
	Distribution *Distribution `json:"Distribution,omitnil,omitempty" name:"Distribution"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`

	// Index information. The inverted index and N-Gram index can be configured through this parameter. The Prefix index is related to the sort key and key column, and do not require additional configuration. Configure bloom_filter_columns in the table attribute when BloomFilter index is required.
	IndexInfos []*IndexInfo `json:"IndexInfos,omitnil,omitempty" name:"IndexInfos"`

	// Partition information
	Partition *Partition `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Table description
	TableComment *string `json:"TableComment,omitnil,omitempty" name:"TableComment"`

	// Table attribute
	Properties []*Property `json:"Properties,omitnil,omitempty" name:"Properties"`
}

type CreateTableRequest struct {
	*tchttp.BaseRequest
	
	// Resource ID, which is the TCHouse-D resource ID used for table creation.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The database where the table is located; if it does not exist, create one.
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Name of the table to be created
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// Table data model: 
	// AGG_KEY: aggregation model; 
	// UNI_KEY: primary key model; 
	// DUP_KEY: detail model
	KeysType *string `json:"KeysType,omitnil,omitempty" name:"KeysType"`

	// Column information of the table
	Columns []*Column `json:"Columns,omitnil,omitempty" name:"Columns"`

	// Bucket information
	Distribution *Distribution `json:"Distribution,omitnil,omitempty" name:"Distribution"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`

	// Index information. The inverted index and N-Gram index can be configured through this parameter. The Prefix index is related to the sort key and key column, and do not require additional configuration. Configure bloom_filter_columns in the table attribute when BloomFilter index is required.
	IndexInfos []*IndexInfo `json:"IndexInfos,omitnil,omitempty" name:"IndexInfos"`

	// Partition information
	Partition *Partition `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Table description
	TableComment *string `json:"TableComment,omitnil,omitempty" name:"TableComment"`

	// Table attribute
	Properties []*Property `json:"Properties,omitnil,omitempty" name:"Properties"`
}

func (r *CreateTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DbName")
	delete(f, "TableName")
	delete(f, "KeysType")
	delete(f, "Columns")
	delete(f, "Distribution")
	delete(f, "UserName")
	delete(f, "PassWord")
	delete(f, "IndexInfos")
	delete(f, "Partition")
	delete(f, "TableComment")
	delete(f, "Properties")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTableResponseParams struct {
	// Error message
	// Note: This field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTableResponse struct {
	*tchttp.BaseResponse
	Response *CreateTableResponseParams `json:"Response"`
}

func (r *CreateTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkloadGroupRequestParams struct {
	// Cluster ID	
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Resource group configuration
	WorkloadGroup *WorkloadGroupConfig `json:"WorkloadGroup,omitnil,omitempty" name:"WorkloadGroup"`
}

type CreateWorkloadGroupRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID	
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Resource group configuration
	WorkloadGroup *WorkloadGroupConfig `json:"WorkloadGroup,omitnil,omitempty" name:"WorkloadGroup"`
}

func (r *CreateWorkloadGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkloadGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "WorkloadGroup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkloadGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkloadGroupResponseParams struct {
	// Error message
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateWorkloadGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateWorkloadGroupResponseParams `json:"Response"`
}

func (r *CreateWorkloadGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkloadGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataBaseAuditRecord struct {
	// Query user
	// Note: This field may return null, indicating that no valid values can be obtained.
	OsUser *string `json:"OsUser,omitnil,omitempty" name:"OsUser"`

	// Query ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	InitialQueryId *string `json:"InitialQueryId,omitnil,omitempty" name:"InitialQueryId"`

	// SQL statement
	// Note: This field may return null, indicating that no valid values can be obtained.
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// Start time
	// Note: This field may return null, indicating that no valid values can be obtained.
	QueryStartTime *string `json:"QueryStartTime,omitnil,omitempty" name:"QueryStartTime"`

	// Execution duration
	// Note: This field may return null, indicating that no valid values can be obtained.
	DurationMs *int64 `json:"DurationMs,omitnil,omitempty" name:"DurationMs"`

	// The number of read rows
	// Note: This field may return null, indicating that no valid values can be obtained.
	ReadRows *int64 `json:"ReadRows,omitnil,omitempty" name:"ReadRows"`

	// Total number of read bytes
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResultRows *int64 `json:"ResultRows,omitnil,omitempty" name:"ResultRows"`

	// Result bytes
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResultBytes *uint64 `json:"ResultBytes,omitnil,omitempty" name:"ResultBytes"`

	// Memory
	// Note: This field may return null, indicating that no valid values can be obtained.
	MemoryUsage *int64 `json:"MemoryUsage,omitnil,omitempty" name:"MemoryUsage"`

	// Initial query IP
	// Note: This field may return null, indicating that no valid values can be obtained.
	InitialAddress *string `json:"InitialAddress,omitnil,omitempty" name:"InitialAddress"`

	// Database
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// SQL type
	// Note: This field may return null, indicating that no valid values can be obtained.
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// Catalog name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`
}

type DatabasePermissions struct {
	// Database name
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// Permission name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Permissions []*string `json:"Permissions,omitnil,omitempty" name:"Permissions"`
}

type DbInfo struct {
	// Database name
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Database attribute
	// Note: This field may return null, indicating that no valid values can be obtained.
	Properties []*Property `json:"Properties,omitnil,omitempty" name:"Properties"`

	// Metadata address (Available when the data source is Hive or DLC.)
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`
}

// Predefined struct for user
type DeleteBackUpDataRequestParams struct {
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Task ID
	BackUpJobId *int64 `json:"BackUpJobId,omitnil,omitempty" name:"BackUpJobId"`

	// Whether to delete all data
	IsDeleteAll *bool `json:"IsDeleteAll,omitnil,omitempty" name:"IsDeleteAll"`
}

type DeleteBackUpDataRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Task ID
	BackUpJobId *int64 `json:"BackUpJobId,omitnil,omitempty" name:"BackUpJobId"`

	// Whether to delete all data
	IsDeleteAll *bool `json:"IsDeleteAll,omitnil,omitempty" name:"IsDeleteAll"`
}

func (r *DeleteBackUpDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBackUpDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackUpJobId")
	delete(f, "IsDeleteAll")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBackUpDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBackUpDataResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteBackUpDataResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBackUpDataResponseParams `json:"Response"`
}

func (r *DeleteBackUpDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBackUpDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTableRequestParams struct {
	// Resource ID, which is the TCHouse-D resource ID used for table creation.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The name of the database where the table belongs needs to be deleted
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Table name to be deleted
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`

	// True: The system will not check whether there are unfinished transactions in the table. The table will be deleted directly and cannot be recovered. False: The deleted table can be recovered within a period of time (default value).
	IsForce *bool `json:"IsForce,omitnil,omitempty" name:"IsForce"`
}

type DeleteTableRequest struct {
	*tchttp.BaseRequest
	
	// Resource ID, which is the TCHouse-D resource ID used for table creation.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The name of the database where the table belongs needs to be deleted
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Table name to be deleted
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`

	// True: The system will not check whether there are unfinished transactions in the table. The table will be deleted directly and cannot be recovered. False: The deleted table can be recovered within a period of time (default value).
	IsForce *bool `json:"IsForce,omitnil,omitempty" name:"IsForce"`
}

func (r *DeleteTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DbName")
	delete(f, "TableName")
	delete(f, "UserName")
	delete(f, "PassWord")
	delete(f, "IsForce")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTableResponseParams struct {
	// Error message
	// Note: This field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTableResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTableResponseParams `json:"Response"`
}

func (r *DeleteTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWorkloadGroupRequestParams struct {
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Resource group name to be deleted
	WorkloadGroupName *string `json:"WorkloadGroupName,omitnil,omitempty" name:"WorkloadGroupName"`
}

type DeleteWorkloadGroupRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Resource group name to be deleted
	WorkloadGroupName *string `json:"WorkloadGroupName,omitnil,omitempty" name:"WorkloadGroupName"`
}

func (r *DeleteWorkloadGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkloadGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "WorkloadGroupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWorkloadGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWorkloadGroupResponseParams struct {
	// Error message
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteWorkloadGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWorkloadGroupResponseParams `json:"Response"`
}

func (r *DeleteWorkloadGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkloadGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAreaRegionRequestParams struct {
	// Whether it is an international site
	IsInternationalSite *bool `json:"IsInternationalSite,omitnil,omitempty" name:"IsInternationalSite"`
}

type DescribeAreaRegionRequest struct {
	*tchttp.BaseRequest
	
	// Whether it is an international site
	IsInternationalSite *bool `json:"IsInternationalSite,omitnil,omitempty" name:"IsInternationalSite"`
}

func (r *DescribeAreaRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAreaRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IsInternationalSite")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAreaRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAreaRegionResponseParams struct {
	// Region list
	Items []*RegionAreaInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// Front-end rule description
	// Note: This field may return null, indicating that no valid values can be obtained.
	FrontEndRules []*FrontEndRule `json:"FrontEndRules,omitnil,omitempty" name:"FrontEndRules"`

	// Return available allowlist names
	// Note: This field may return null, indicating that no valid values can be obtained.
	AvailableWhiteListNames []*string `json:"AvailableWhiteListNames,omitnil,omitempty" name:"AvailableWhiteListNames"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAreaRegionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAreaRegionResponseParams `json:"Response"`
}

func (r *DescribeAreaRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAreaRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackUpJobDetailRequestParams struct {
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Task ID
	BackUpJobId *int64 `json:"BackUpJobId,omitnil,omitempty" name:"BackUpJobId"`
}

type DescribeBackUpJobDetailRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Task ID
	BackUpJobId *int64 `json:"BackUpJobId,omitnil,omitempty" name:"BackUpJobId"`
}

func (r *DescribeBackUpJobDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackUpJobDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackUpJobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackUpJobDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackUpJobDetailResponseParams struct {
	// Backup table details
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableContents []*BackupTableContent `json:"TableContents,omitnil,omitempty" name:"TableContents"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackUpJobDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackUpJobDetailResponseParams `json:"Response"`
}

func (r *DescribeBackUpJobDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackUpJobDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackUpJobRequestParams struct {
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Pagination size
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Page number
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// Start time
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// String type of jobid
	JobIdFiltersStr *string `json:"JobIdFiltersStr,omitnil,omitempty" name:"JobIdFiltersStr"`
}

type DescribeBackUpJobRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Pagination size
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Page number
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// Start time
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// String type of jobid
	JobIdFiltersStr *string `json:"JobIdFiltersStr,omitnil,omitempty" name:"JobIdFiltersStr"`
}

func (r *DescribeBackUpJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackUpJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PageSize")
	delete(f, "PageNum")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "JobIdFiltersStr")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackUpJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackUpJobResponseParams struct {
	// Task list
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackUpJobs []*BackUpJobDisplay `json:"BackUpJobs,omitnil,omitempty" name:"BackUpJobs"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackUpJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackUpJobResponseParams `json:"Response"`
}

func (r *DescribeBackUpJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackUpJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackUpSchedulesRequestParams struct {

}

type DescribeBackUpSchedulesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeBackUpSchedulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackUpSchedulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackUpSchedulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackUpSchedulesResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackUpSchedulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackUpSchedulesResponseParams `json:"Response"`
}

func (r *DescribeBackUpSchedulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackUpSchedulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackUpTablesRequestParams struct {
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// It is 0 by default. It is 1 when a one-time backup of the remote doris is performed. It is 2 when one-time COS recovery is performed.
	BackupType *int64 `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// Connection information of the remote doris cluster
	DorisSourceInfo *DorisSourceInfo `json:"DorisSourceInfo,omitnil,omitempty" name:"DorisSourceInfo"`

	// COS information
	CosSourceInfo *CosSourceInfo `json:"CosSourceInfo,omitnil,omitempty" name:"CosSourceInfo"`
}

type DescribeBackUpTablesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// It is 0 by default. It is 1 when a one-time backup of the remote doris is performed. It is 2 when one-time COS recovery is performed.
	BackupType *int64 `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// Connection information of the remote doris cluster
	DorisSourceInfo *DorisSourceInfo `json:"DorisSourceInfo,omitnil,omitempty" name:"DorisSourceInfo"`

	// COS information
	CosSourceInfo *CosSourceInfo `json:"CosSourceInfo,omitnil,omitempty" name:"CosSourceInfo"`
}

func (r *DescribeBackUpTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackUpTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupType")
	delete(f, "DorisSourceInfo")
	delete(f, "CosSourceInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackUpTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackUpTablesResponseParams struct {
	// List of tables available for backup
	AvailableTables []*BackupTableContent `json:"AvailableTables,omitnil,omitempty" name:"AvailableTables"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackUpTablesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackUpTablesResponseParams `json:"Response"`
}

func (r *DescribeBackUpTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackUpTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackUpTaskDetailRequestParams struct {
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Task ID
	BackUpJobId *int64 `json:"BackUpJobId,omitnil,omitempty" name:"BackUpJobId"`
}

type DescribeBackUpTaskDetailRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Task ID
	BackUpJobId *int64 `json:"BackUpJobId,omitnil,omitempty" name:"BackUpJobId"`
}

func (r *DescribeBackUpTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackUpTaskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackUpJobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackUpTaskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackUpTaskDetailResponseParams struct {
	// Progress details of the backup task
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupStatus []*BackupStatus `json:"BackupStatus,omitnil,omitempty" name:"BackupStatus"`

	// Error message
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackUpTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackUpTaskDetailResponseParams `json:"Response"`
}

func (r *DescribeBackUpTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackUpTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterConfigsHistoryRequestParams struct {
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Pagination parameters. The first page is 0, and the second page is 10.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Pagination parameters. The pagination step length is 10 by default.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Start of the time range for configuration modification history
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End of the time range for configuration modification history
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Configuration file name array to be queried. If it is empty, all historical records will be queried. Currently supported configuration file names are as follows:
	// apache_hdfs_broker.conf; be.conf; fe.conf; core-site.xml; hdfs-site.xml; odbcinst.ini
	ConfigFileNames []*string `json:"ConfigFileNames,omitnil,omitempty" name:"ConfigFileNames"`
}

type DescribeClusterConfigsHistoryRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Pagination parameters. The first page is 0, and the second page is 10.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Pagination parameters. The pagination step length is 10 by default.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Start of the time range for configuration modification history
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End of the time range for configuration modification history
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Configuration file name array to be queried. If it is empty, all historical records will be queried. Currently supported configuration file names are as follows:
	// apache_hdfs_broker.conf; be.conf; fe.conf; core-site.xml; hdfs-site.xml; odbcinst.ini
	ConfigFileNames []*string `json:"ConfigFileNames,omitnil,omitempty" name:"ConfigFileNames"`
}

func (r *DescribeClusterConfigsHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterConfigsHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ConfigFileNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterConfigsHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterConfigsHistoryResponseParams struct {
	// Total number of instances
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Modification history of the configuration file
	ClusterConfHistory []*ClusterConfigsHistory `json:"ClusterConfHistory,omitnil,omitempty" name:"ClusterConfHistory"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterConfigsHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterConfigsHistoryResponseParams `json:"Response"`
}

func (r *DescribeClusterConfigsHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterConfigsHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterConfigsRequestParams struct {
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 0 indicates public cloud query, and 1 Qinge query. Qinge query shows all that needs to be displayed.
	ConfigType *int64 `json:"ConfigType,omitnil,omitempty" name:"ConfigType"`

	// Search for files with fuzzy keywords
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 0 indicates cluster dimension and 1 node dimension
	ClusterConfigType *int64 `json:"ClusterConfigType,omitnil,omitempty" name:"ClusterConfigType"`

	// eth0's IP address
	IPAddress *string `json:"IPAddress,omitnil,omitempty" name:"IPAddress"`
}

type DescribeClusterConfigsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 0 indicates public cloud query, and 1 Qinge query. Qinge query shows all that needs to be displayed.
	ConfigType *int64 `json:"ConfigType,omitnil,omitempty" name:"ConfigType"`

	// Search for files with fuzzy keywords
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 0 indicates cluster dimension and 1 node dimension
	ClusterConfigType *int64 `json:"ClusterConfigType,omitnil,omitempty" name:"ClusterConfigType"`

	// eth0's IP address
	IPAddress *string `json:"IPAddress,omitnil,omitempty" name:"IPAddress"`
}

func (r *DescribeClusterConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ConfigType")
	delete(f, "FileName")
	delete(f, "ClusterConfigType")
	delete(f, "IPAddress")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterConfigsResponseParams struct {
	// Return information about the instance's configuration file.
	ClusterConfList []*ClusterConfigsInfoFromEMR `json:"ClusterConfList,omitnil,omitempty" name:"ClusterConfList"`

	// Return the current kernel version. If it does not exist, a null character string is returned.
	BuildVersion *string `json:"BuildVersion,omitnil,omitempty" name:"BuildVersion"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterConfigsResponseParams `json:"Response"`
}

func (r *DescribeClusterConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseAuditDownloadRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Paging
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Paging
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// Sort parameters
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// User
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Database
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// SQL type
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// SQL statement
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// Users (multiple selections)
	Users []*string `json:"Users,omitnil,omitempty" name:"Users"`

	// Databases (multiple selections)
	DbNames []*string `json:"DbNames,omitnil,omitempty" name:"DbNames"`

	// SQL types (multiple selections)
	SqlTypes []*string `json:"SqlTypes,omitnil,omitempty" name:"SqlTypes"`

	// Catalog names (multiple selections)
	Catalogs []*string `json:"Catalogs,omitnil,omitempty" name:"Catalogs"`
}

type DescribeDatabaseAuditDownloadRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Paging
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Paging
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// Sort parameters
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// User
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Database
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// SQL type
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// SQL statement
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// Users (multiple selections)
	Users []*string `json:"Users,omitnil,omitempty" name:"Users"`

	// Databases (multiple selections)
	DbNames []*string `json:"DbNames,omitnil,omitempty" name:"DbNames"`

	// SQL types (multiple selections)
	SqlTypes []*string `json:"SqlTypes,omitnil,omitempty" name:"SqlTypes"`

	// Catalog names (multiple selections)
	Catalogs []*string `json:"Catalogs,omitnil,omitempty" name:"Catalogs"`
}

func (r *DescribeDatabaseAuditDownloadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseAuditDownloadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "PageNum")
	delete(f, "OrderType")
	delete(f, "User")
	delete(f, "DbName")
	delete(f, "SqlType")
	delete(f, "Sql")
	delete(f, "Users")
	delete(f, "DbNames")
	delete(f, "SqlTypes")
	delete(f, "Catalogs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabaseAuditDownloadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseAuditDownloadResponseParams struct {
	// The cos address of the log
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatabaseAuditDownloadResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatabaseAuditDownloadResponseParams `json:"Response"`
}

func (r *DescribeDatabaseAuditDownloadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseAuditDownloadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseAuditRecordsRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Paging
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Paging
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// Sort parameters
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// User
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Database
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// SQL type
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// SQL statement
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// Users (multiple selections)
	Users []*string `json:"Users,omitnil,omitempty" name:"Users"`

	// Databases (multiple selections)
	DbNames []*string `json:"DbNames,omitnil,omitempty" name:"DbNames"`

	// SQL types (multiple selections)
	SqlTypes []*string `json:"SqlTypes,omitnil,omitempty" name:"SqlTypes"`

	// Catalog names (multiple selections)
	Catalogs []*string `json:"Catalogs,omitnil,omitempty" name:"Catalogs"`
}

type DescribeDatabaseAuditRecordsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Paging
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Paging
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// Sort parameters
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// User
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Database
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// SQL type
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// SQL statement
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// Users (multiple selections)
	Users []*string `json:"Users,omitnil,omitempty" name:"Users"`

	// Databases (multiple selections)
	DbNames []*string `json:"DbNames,omitnil,omitempty" name:"DbNames"`

	// SQL types (multiple selections)
	SqlTypes []*string `json:"SqlTypes,omitnil,omitempty" name:"SqlTypes"`

	// Catalog names (multiple selections)
	Catalogs []*string `json:"Catalogs,omitnil,omitempty" name:"Catalogs"`
}

func (r *DescribeDatabaseAuditRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseAuditRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "PageNum")
	delete(f, "OrderType")
	delete(f, "User")
	delete(f, "DbName")
	delete(f, "SqlType")
	delete(f, "Sql")
	delete(f, "Users")
	delete(f, "DbNames")
	delete(f, "SqlTypes")
	delete(f, "Catalogs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabaseAuditRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseAuditRecordsResponseParams struct {
	// Total
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Record list
	SlowQueryRecords *DataBaseAuditRecord `json:"SlowQueryRecords,omitnil,omitempty" name:"SlowQueryRecords"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatabaseAuditRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatabaseAuditRecordsResponseParams `json:"Response"`
}

func (r *DescribeDatabaseAuditRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseAuditRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseRequestParams struct {
	// Resource ID, which is the TCHouse-D resource ID used for table creation.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`

	// Query the data source where the database is located. If it is not filled in, the internal data source (internal) will be used by default.
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// The database information to be queried. If this parameter and FilterDbName are not filled in, all database information of the current data source will be queried by default.
	DbNames []*string `json:"DbNames,omitnil,omitempty" name:"DbNames"`

	// Display the filtered database information. For example, %infor indicates database names ending with infor. This parameter only supports scenarios where CatalogName is internal.
	FilterDbName *string `json:"FilterDbName,omitnil,omitempty" name:"FilterDbName"`
}

type DescribeDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// Resource ID, which is the TCHouse-D resource ID used for table creation.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`

	// Query the data source where the database is located. If it is not filled in, the internal data source (internal) will be used by default.
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// The database information to be queried. If this parameter and FilterDbName are not filled in, all database information of the current data source will be queried by default.
	DbNames []*string `json:"DbNames,omitnil,omitempty" name:"DbNames"`

	// Display the filtered database information. For example, %infor indicates database names ending with infor. This parameter only supports scenarios where CatalogName is internal.
	FilterDbName *string `json:"FilterDbName,omitnil,omitempty" name:"FilterDbName"`
}

func (r *DescribeDatabaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "PassWord")
	delete(f, "CatalogName")
	delete(f, "DbNames")
	delete(f, "FilterDbName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseResponseParams struct {
	// Database information
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbInfos []*DbInfo `json:"DbInfos,omitnil,omitempty" name:"DbInfos"`

	// Error message
	// Note: This field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatabaseResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatabaseResponseParams `json:"Response"`
}

func (r *DescribeDatabaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceNodesInfoRequestParams struct {
	// Cluster ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

type DescribeInstanceNodesInfoRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

func (r *DescribeInstanceNodesInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodesInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceNodesInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceNodesInfoResponseParams struct {
	// Be node
	// Note: This field may return null, indicating that no valid values can be obtained.
	BeNodes []*string `json:"BeNodes,omitnil,omitempty" name:"BeNodes"`

	// Fe node
	// Note: This field may return null, indicating that no valid values can be obtained.
	FeNodes []*string `json:"FeNodes,omitnil,omitempty" name:"FeNodes"`

	// Fe master node
	FeMaster *string `json:"FeMaster,omitnil,omitempty" name:"FeMaster"`

	// Be node information
	// Note: This field may return null, indicating that no valid values can be obtained.
	BeNodeInfos []*NodeInfo `json:"BeNodeInfos,omitnil,omitempty" name:"BeNodeInfos"`

	// Fe node information
	// Note: This field may return null, indicating that no valid values can be obtained.
	FeNodeInfos []*NodeInfo `json:"FeNodeInfos,omitnil,omitempty" name:"FeNodeInfos"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceNodesInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceNodesInfoResponseParams `json:"Response"`
}

func (r *DescribeInstanceNodesInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodesInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceNodesRequestParams struct {
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Cluster role type, defaulted as "data node".
	NodeRole *string `json:"NodeRole,omitnil,omitempty" name:"NodeRole"`

	// Pagination parameters. The first page is 0, and the second page is 10.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Pagination parameters. The pagination step length is 10 by default.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Display policy, and all items are displayed when All is selected.
	DisplayPolicy *string `json:"DisplayPolicy,omitnil,omitempty" name:"DisplayPolicy"`
}

type DescribeInstanceNodesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Cluster role type, defaulted as "data node".
	NodeRole *string `json:"NodeRole,omitnil,omitempty" name:"NodeRole"`

	// Pagination parameters. The first page is 0, and the second page is 10.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Pagination parameters. The pagination step length is 10 by default.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Display policy, and all items are displayed when All is selected.
	DisplayPolicy *string `json:"DisplayPolicy,omitnil,omitempty" name:"DisplayPolicy"`
}

func (r *DescribeInstanceNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NodeRole")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "DisplayPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceNodesResponseParams struct {
	// Total number
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Total number of instance nodes
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceNodesList []*InstanceNode `json:"InstanceNodesList,omitnil,omitempty" name:"InstanceNodesList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceNodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceNodesResponseParams `json:"Response"`
}

func (r *DescribeInstanceNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceNodesRoleRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Filter IP addresses
	IpFilter *string `json:"IpFilter,omitnil,omitempty" name:"IpFilter"`
}

type DescribeInstanceNodesRoleRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Filter IP addresses
	IpFilter *string `json:"IpFilter,omitnil,omitempty" name:"IpFilter"`
}

func (r *DescribeInstanceNodesRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodesRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IpFilter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceNodesRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceNodesRoleResponseParams struct {
	// Error code
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// Total number of nodes
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// None
	NodeInfos []*NodeInfos `json:"NodeInfos,omitnil,omitempty" name:"NodeInfos"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceNodesRoleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceNodesRoleResponseParams `json:"Response"`
}

func (r *DescribeInstanceNodesRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodesRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceOperationHistoryRequestParams struct {
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Page number, which is 1 by default.
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// Number of records per page, which is 10 by default.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`
}

type DescribeInstanceOperationHistoryRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Page number, which is 1 by default.
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// Number of records per page, which is 10 by default.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`
}

func (r *DescribeInstanceOperationHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceOperationHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PageNum")
	delete(f, "PageSize")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "UserName")
	delete(f, "PassWord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceOperationHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceOperationHistoryResponseParams struct {
	// Total number of operation records
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Specific data of operation records
	Operations []*InstanceOperation `json:"Operations,omitnil,omitempty" name:"Operations"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceOperationHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceOperationHistoryResponseParams `json:"Response"`
}

func (r *DescribeInstanceOperationHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceOperationHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceOperationsRequestParams struct {

	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`


	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`


	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`


	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`


	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeInstanceOperationsRequest struct {
	*tchttp.BaseRequest
	
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeInstanceOperationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceOperationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceOperationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceOperationsResponseParams struct {

	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Note: This field may return null, indicating that no valid values can be obtained.
	Operations []*InstanceOperation `json:"Operations,omitnil,omitempty" name:"Operations"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceOperationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceOperationsResponseParams `json:"Response"`
}

func (r *DescribeInstanceOperationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceOperationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceRequestParams struct {
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceResponseParams struct {
	// Instance description information
	InstanceInfo *InstanceInfo `json:"InstanceInfo,omitnil,omitempty" name:"InstanceInfo"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceResponseParams `json:"Response"`
}

func (r *DescribeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceStateRequestParams struct {
	// Cluster instance name
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceStateRequest struct {
	*tchttp.BaseRequest
	
	// Cluster instance name
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceStateResponseParams struct {
	// Cluster status. Example: Serving
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// Creation time of cluster operation
	// Note: This field may return null, indicating that no valid values can be obtained.
	FlowCreateTime *string `json:"FlowCreateTime,omitnil,omitempty" name:"FlowCreateTime"`

	// Cluster operation name
	// Note: This field may return null, indicating that no valid values can be obtained.
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// Cluster operation progress
	// Note: This field may return null, indicating that no valid values can be obtained.
	FlowProgress *float64 `json:"FlowProgress,omitnil,omitempty" name:"FlowProgress"`

	// Cluster status description. Example: running
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceStateDesc *string `json:"InstanceStateDesc,omitnil,omitempty" name:"InstanceStateDesc"`

	// Cluster process error messages, such as "Creation failed, insufficient resources"
	// Note: This field may return null, indicating that no valid values can be obtained.
	FlowMsg *string `json:"FlowMsg,omitnil,omitempty" name:"FlowMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceStateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceStateResponseParams `json:"Response"`
}

func (r *DescribeInstanceStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceUsedSubnetsRequestParams struct {
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceUsedSubnetsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceUsedSubnetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceUsedSubnetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceUsedSubnetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceUsedSubnetsResponseParams struct {
	// VPC information used by the cluster
	// Note: This field may return null, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet information used by the cluster
	// Note: This field may return null, indicating that no valid values can be obtained.
	UsedSubnets []*string `json:"UsedSubnets,omitnil,omitempty" name:"UsedSubnets"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceUsedSubnetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceUsedSubnetsResponseParams `json:"Response"`
}

func (r *DescribeInstanceUsedSubnetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceUsedSubnetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesHealthStateRequestParams struct {
	// Cluster ID
	//
	// Deprecated: InstanceID is deprecated.
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// "" or a cluster ID
	Input *string `json:"Input,omitnil,omitempty" name:"Input"`
}

type DescribeInstancesHealthStateRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// "" or a cluster ID
	Input *string `json:"Input,omitnil,omitempty" name:"Input"`
}

func (r *DescribeInstancesHealthStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesHealthStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "Input")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesHealthStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesHealthStateResponseParams struct {
	// Output parameter
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstancesHealthStateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesHealthStateResponseParams `json:"Response"`
}

func (r *DescribeInstancesHealthStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesHealthStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// The name of the cluster ID for the search
	SearchInstanceId *string `json:"SearchInstanceId,omitnil,omitempty" name:"SearchInstanceId"`

	// The cluster name for the search
	SearchInstanceName *string `json:"SearchInstanceName,omitnil,omitempty" name:"SearchInstanceName"`

	// Pagination parameters. The first page is 0, and the second page is 10.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Pagination parameters. The pagination step length is 10 by default.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Search tag list
	SearchTags []*SearchTags `json:"SearchTags,omitnil,omitempty" name:"SearchTags"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// The name of the cluster ID for the search
	SearchInstanceId *string `json:"SearchInstanceId,omitnil,omitempty" name:"SearchInstanceId"`

	// The cluster name for the search
	SearchInstanceName *string `json:"SearchInstanceName,omitnil,omitempty" name:"SearchInstanceName"`

	// Pagination parameters. The first page is 0, and the second page is 10.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Pagination parameters. The pagination step length is 10 by default.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Search tag list
	SearchTags []*SearchTags `json:"SearchTags,omitnil,omitempty" name:"SearchTags"`
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
	delete(f, "SearchInstanceId")
	delete(f, "SearchInstanceName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// Total Number of Instances
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Quantities of instances array
	InstancesList []*InstanceInfo `json:"InstancesList,omitnil,omitempty" name:"InstancesList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeQueryAnalyseRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`

	// Start time of operation period
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of operation period
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// SQL fragments (fuzzy query supported)
	SQLFragment *string `json:"SQLFragment,omitnil,omitempty" name:"SQLFragment"`

	// Catalog filter condition
	CatalogFilter *string `json:"CatalogFilter,omitnil,omitempty" name:"CatalogFilter"`

	// Database name filter condition
	DatabaseFilter *string `json:"DatabaseFilter,omitnil,omitempty" name:"DatabaseFilter"`

	// SQL type filter criteria
	SQLTypeFilter *string `json:"SQLTypeFilter,omitnil,omitempty" name:"SQLTypeFilter"`

	// Sorting field
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// Sorting order: asc (ascending) or desc (descending)
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`
}

type DescribeQueryAnalyseRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`

	// Start time of operation period
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time of operation period
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// SQL fragments (fuzzy query supported)
	SQLFragment *string `json:"SQLFragment,omitnil,omitempty" name:"SQLFragment"`

	// Catalog filter condition
	CatalogFilter *string `json:"CatalogFilter,omitnil,omitempty" name:"CatalogFilter"`

	// Database name filter condition
	DatabaseFilter *string `json:"DatabaseFilter,omitnil,omitempty" name:"DatabaseFilter"`

	// SQL type filter criteria
	SQLTypeFilter *string `json:"SQLTypeFilter,omitnil,omitempty" name:"SQLTypeFilter"`

	// Sorting field
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// Sorting order: asc (ascending) or desc (descending)
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`
}

func (r *DescribeQueryAnalyseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQueryAnalyseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "PassWord")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SQLFragment")
	delete(f, "CatalogFilter")
	delete(f, "DatabaseFilter")
	delete(f, "SQLTypeFilter")
	delete(f, "SortField")
	delete(f, "SortOrder")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeQueryAnalyseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQueryAnalyseResponseParams struct {
	// Query details
	QueryDetails []*QueryDetails `json:"QueryDetails,omitnil,omitempty" name:"QueryDetails"`

	// Total number of records
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Current page
	CurrentPage *uint64 `json:"CurrentPage,omitnil,omitempty" name:"CurrentPage"`

	// Number of records per page
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Total pages
	TotalPages *uint64 `json:"TotalPages,omitnil,omitempty" name:"TotalPages"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeQueryAnalyseResponse struct {
	*tchttp.BaseResponse
	Response *DescribeQueryAnalyseResponseParams `json:"Response"`
}

func (r *DescribeQueryAnalyseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQueryAnalyseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRestoreTaskDetailRequestParams struct {
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Task ID
	BackUpJobId *int64 `json:"BackUpJobId,omitnil,omitempty" name:"BackUpJobId"`
}

type DescribeRestoreTaskDetailRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Task ID
	BackUpJobId *int64 `json:"BackUpJobId,omitnil,omitempty" name:"BackUpJobId"`
}

func (r *DescribeRestoreTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRestoreTaskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackUpJobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRestoreTaskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRestoreTaskDetailResponseParams struct {
	// Progress details of the recovery tasks
	// Note: This field may return null, indicating that no valid values can be obtained.
	RestoreStatus []*RestoreStatus `json:"RestoreStatus,omitnil,omitempty" name:"RestoreStatus"`

	// Error message
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRestoreTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRestoreTaskDetailResponseParams `json:"Response"`
}

func (r *DescribeRestoreTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRestoreTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowQueryRecordsDownloadRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Slow log time
	QueryDurationMs *int64 `json:"QueryDurationMs,omitnil,omitempty" name:"QueryDurationMs"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Sort parameters
	DurationMs *string `json:"DurationMs,omitnil,omitempty" name:"DurationMs"`

	// Query SQL
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// Sort parameters
	ReadRows *string `json:"ReadRows,omitnil,omitempty" name:"ReadRows"`

	// Sort parameters
	ResultBytes *string `json:"ResultBytes,omitnil,omitempty" name:"ResultBytes"`

	// Sort parameters
	MemoryUsage *string `json:"MemoryUsage,omitnil,omitempty" name:"MemoryUsage"`

	// IsQuery condition
	IsQuery *int64 `json:"IsQuery,omitnil,omitempty" name:"IsQuery"`

	// Database name
	DbName []*string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// catalog name
	CatalogName []*string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`
}

type DescribeSlowQueryRecordsDownloadRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Slow log time
	QueryDurationMs *int64 `json:"QueryDurationMs,omitnil,omitempty" name:"QueryDurationMs"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Sort parameters
	DurationMs *string `json:"DurationMs,omitnil,omitempty" name:"DurationMs"`

	// Query SQL
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// Sort parameters
	ReadRows *string `json:"ReadRows,omitnil,omitempty" name:"ReadRows"`

	// Sort parameters
	ResultBytes *string `json:"ResultBytes,omitnil,omitempty" name:"ResultBytes"`

	// Sort parameters
	MemoryUsage *string `json:"MemoryUsage,omitnil,omitempty" name:"MemoryUsage"`

	// IsQuery condition
	IsQuery *int64 `json:"IsQuery,omitnil,omitempty" name:"IsQuery"`

	// Database name
	DbName []*string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// catalog name
	CatalogName []*string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`
}

func (r *DescribeSlowQueryRecordsDownloadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowQueryRecordsDownloadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "QueryDurationMs")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "DurationMs")
	delete(f, "Sql")
	delete(f, "ReadRows")
	delete(f, "ResultBytes")
	delete(f, "MemoryUsage")
	delete(f, "IsQuery")
	delete(f, "DbName")
	delete(f, "CatalogName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowQueryRecordsDownloadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowQueryRecordsDownloadResponseParams struct {
	// cos address
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowQueryRecordsDownloadResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowQueryRecordsDownloadResponseParams `json:"Response"`
}

func (r *DescribeSlowQueryRecordsDownloadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowQueryRecordsDownloadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowQueryRecordsRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Slow log time
	QueryDurationMs *int64 `json:"QueryDurationMs,omitnil,omitempty" name:"QueryDurationMs"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Paging
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Paging
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// Sort parameters
	DurationMs *string `json:"DurationMs,omitnil,omitempty" name:"DurationMs"`

	// Database name
	DbName []*string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Whether it is a query. 0 indicates no, and 1 indicates yes.
	IsQuery *int64 `json:"IsQuery,omitnil,omitempty" name:"IsQuery"`

	// catalog name
	CatalogName []*string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// SQL name
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// ReadRows sort field
	ReadRows *string `json:"ReadRows,omitnil,omitempty" name:"ReadRows"`

	// ResultBytes sort field
	ResultBytes *string `json:"ResultBytes,omitnil,omitempty" name:"ResultBytes"`

	// MemoryUsage sort field
	MemoryUsage *string `json:"MemoryUsage,omitnil,omitempty" name:"MemoryUsage"`
}

type DescribeSlowQueryRecordsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Slow log time
	QueryDurationMs *int64 `json:"QueryDurationMs,omitnil,omitempty" name:"QueryDurationMs"`

	// Start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Paging
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Paging
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// Sort parameters
	DurationMs *string `json:"DurationMs,omitnil,omitempty" name:"DurationMs"`

	// Database name
	DbName []*string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Whether it is a query. 0 indicates no, and 1 indicates yes.
	IsQuery *int64 `json:"IsQuery,omitnil,omitempty" name:"IsQuery"`

	// catalog name
	CatalogName []*string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// SQL name
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// ReadRows sort field
	ReadRows *string `json:"ReadRows,omitnil,omitempty" name:"ReadRows"`

	// ResultBytes sort field
	ResultBytes *string `json:"ResultBytes,omitnil,omitempty" name:"ResultBytes"`

	// MemoryUsage sort field
	MemoryUsage *string `json:"MemoryUsage,omitnil,omitempty" name:"MemoryUsage"`
}

func (r *DescribeSlowQueryRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowQueryRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "QueryDurationMs")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "PageNum")
	delete(f, "DurationMs")
	delete(f, "DbName")
	delete(f, "IsQuery")
	delete(f, "CatalogName")
	delete(f, "Sql")
	delete(f, "ReadRows")
	delete(f, "ResultBytes")
	delete(f, "MemoryUsage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowQueryRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowQueryRecordsResponseParams struct {
	// Total
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Record list
	SlowQueryRecords []*SlowQueryRecord `json:"SlowQueryRecords,omitnil,omitempty" name:"SlowQueryRecords"`

	// All database names
	// Note: This field may return null, indicating that no valid values can be obtained.
	DBNameList []*string `json:"DBNameList,omitnil,omitempty" name:"DBNameList"`

	// All catalog names
	// Note: This field may return null, indicating that no valid values can be obtained.
	CatalogNameList []*string `json:"CatalogNameList,omitnil,omitempty" name:"CatalogNameList"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowQueryRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowQueryRecordsResponseParams `json:"Response"`
}

func (r *DescribeSlowQueryRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowQueryRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpecRequestParams struct {
	// Region information, such as ap-guangzhou-1.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Billing type. PREPAID: annual/monthly package; POSTPAID_BY_HOUR: pay-as-you-go
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Multi-availability zone
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// Model name
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`
}

type DescribeSpecRequest struct {
	*tchttp.BaseRequest
	
	// Region information, such as ap-guangzhou-1.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Billing type. PREPAID: annual/monthly package; POSTPAID_BY_HOUR: pay-as-you-go
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Multi-availability zone
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// Model name
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`
}

func (r *DescribeSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "PayMode")
	delete(f, "Zones")
	delete(f, "SpecName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpecResponseParams struct {
	// Zookeeper node specification description
	MasterSpec []*ResourceSpec `json:"MasterSpec,omitnil,omitempty" name:"MasterSpec"`

	// Data node specification description
	CoreSpec []*ResourceSpec `json:"CoreSpec,omitnil,omitempty" name:"CoreSpec"`

	// Cloud disk list
	// Note: This field may return null, indicating that no valid values can be obtained.
	AttachCBSSpec []*DiskSpec `json:"AttachCBSSpec,omitnil,omitempty" name:"AttachCBSSpec"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSpecResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpecResponseParams `json:"Response"`
}

func (r *DescribeSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSqlApisRequestParams struct {
	// The IP address of the user link
	WhiteHost *string `json:"WhiteHost,omitnil,omitempty" name:"WhiteHost"`

	// catalog name
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// Catalog collection
	Catalogs []*string `json:"Catalogs,omitnil,omitempty" name:"Catalogs"`
}

type DescribeSqlApisRequest struct {
	*tchttp.BaseRequest
	
	// The IP address of the user link
	WhiteHost *string `json:"WhiteHost,omitnil,omitempty" name:"WhiteHost"`

	// catalog name
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// Catalog collection
	Catalogs []*string `json:"Catalogs,omitnil,omitempty" name:"Catalogs"`
}

func (r *DescribeSqlApisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSqlApisRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WhiteHost")
	delete(f, "Catalog")
	delete(f, "Catalogs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSqlApisRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSqlApisResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSqlApisResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSqlApisResponseParams `json:"Response"`
}

func (r *DescribeSqlApisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSqlApisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableListRequestParams struct {
	// Resource ID, which is the TCHouse-D resource ID used for table creation.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database for obtaining the list of tables
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`

	// Query the data source where the database is located. If it is not filled in, the internal data source (internal) will be used by default.
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`
}

type DescribeTableListRequest struct {
	*tchttp.BaseRequest
	
	// Resource ID, which is the TCHouse-D resource ID used for table creation.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Database for obtaining the list of tables
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`

	// Query the data source where the database is located. If it is not filled in, the internal data source (internal) will be used by default.
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`
}

func (r *DescribeTableListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DbName")
	delete(f, "UserName")
	delete(f, "PassWord")
	delete(f, "CatalogName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTableListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableListResponseParams struct {
	// List of table names
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableNames []*string `json:"TableNames,omitnil,omitempty" name:"TableNames"`

	// Error message
	// Note: This field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTableListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTableListResponseParams `json:"Response"`
}

func (r *DescribeTableListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableRequestParams struct {
	// Resource ID, which is the TCHouse-D resource ID used for table creation.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Name of database where the table is located
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Table name (Currently only internal tables are supported.)
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`
}

type DescribeTableRequest struct {
	*tchttp.BaseRequest
	
	// Resource ID, which is the TCHouse-D resource ID used for table creation.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Name of database where the table is located
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Table name (Currently only internal tables are supported.)
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`
}

func (r *DescribeTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DbName")
	delete(f, "TableName")
	delete(f, "UserName")
	delete(f, "PassWord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableResponseParams struct {
	// Table data model: 
	// AGG_KEY: aggregation model; 
	// UNI_KEY: primary key model; 
	// DUP_KEY: detail model
	KeysType *string `json:"KeysType,omitnil,omitempty" name:"KeysType"`

	// Table column information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Columns []*Column `json:"Columns,omitnil,omitempty" name:"Columns"`

	// Index information. The inverted index and N-Gram index can be viewed through this parameter.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IndexInfos []*IndexInfo `json:"IndexInfos,omitnil,omitempty" name:"IndexInfos"`

	// Partition information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Partition *Partition `json:"Partition,omitnil,omitempty" name:"Partition"`

	// Bucket information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Distribution *Distribution `json:"Distribution,omitnil,omitempty" name:"Distribution"`

	// Table description
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableComment *string `json:"TableComment,omitnil,omitempty" name:"TableComment"`

	// Table attributes
	// Note: This field may return null, indicating that no valid values can be obtained.
	Properties []*Property `json:"Properties,omitnil,omitempty" name:"Properties"`

	// Error message
	// Note: This field may return null, indicating that no valid values can be obtained.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTableResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTableResponseParams `json:"Response"`
}

func (r *DescribeTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserBindWorkloadGroupRequestParams struct {
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeUserBindWorkloadGroupRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeUserBindWorkloadGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserBindWorkloadGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserBindWorkloadGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserBindWorkloadGroupResponseParams struct {
	// Resource group information bound to the user
	UserBindInfos []*UserWorkloadGroup `json:"UserBindInfos,omitnil,omitempty" name:"UserBindInfos"`

	// Error message
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserBindWorkloadGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserBindWorkloadGroupResponseParams `json:"Response"`
}

func (r *DescribeUserBindWorkloadGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserBindWorkloadGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserPolicyRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`

	// You can specify one IP address or specify a percent sign (%) to indicate no limit.
	WhiteHost *string `json:"WhiteHost,omitnil,omitempty" name:"WhiteHost"`
}

type DescribeUserPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`

	// You can specify one IP address or specify a percent sign (%) to indicate no limit.
	WhiteHost *string `json:"WhiteHost,omitnil,omitempty" name:"WhiteHost"`
}

func (r *DescribeUserPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "PassWord")
	delete(f, "WhiteHost")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserPolicyResponseParams struct {
	// Account details
	AccountInfo *AccountDetailInfo `json:"AccountInfo,omitnil,omitempty" name:"AccountInfo"`

	// Permission configuration information associated with different hosts
	Permissions []*PermissionHostInfo `json:"Permissions,omitnil,omitempty" name:"Permissions"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserPolicyResponseParams `json:"Response"`
}

func (r *DescribeUserPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkloadGroupRequestParams struct {
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeWorkloadGroupRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeWorkloadGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkloadGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkloadGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkloadGroupResponseParams struct {
	// Resource group information
	WorkloadGroups []*WorkloadGroupConfig `json:"WorkloadGroups,omitnil,omitempty" name:"WorkloadGroups"`

	// Whether to enable the resource group: open and close
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Error message
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWorkloadGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkloadGroupResponseParams `json:"Response"`
}

func (r *DescribeWorkloadGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkloadGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyInstanceRequestParams struct {
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DestroyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DestroyInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyInstanceResponseParams struct {
	// Process ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Error message
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DestroyInstanceResponseParams `json:"Response"`
}

func (r *DestroyInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiskSpec struct {
	// Disk type, such as CLOUD_SSD and LOCAL_SSD
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Disk type description, such as cloud SSD and local SSD
	DiskDesc *string `json:"DiskDesc,omitnil,omitempty" name:"DiskDesc"`

	// Minimum disk size, in GB
	MinDiskSize *int64 `json:"MinDiskSize,omitnil,omitempty" name:"MinDiskSize"`

	// Maximum disk size, in GB
	MaxDiskSize *int64 `json:"MaxDiskSize,omitnil,omitempty" name:"MaxDiskSize"`

	// Number of disks
	DiskCount *int64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`
}

type Distribution struct {
	// Bucket type:
	// ●Hash: hash bucket
	// ●Random: random number bucket
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	DistributionType *string `json:"DistributionType,omitnil,omitempty" name:"DistributionType"`

	// Number of buckets
	// Note: This field may return null, indicating that no valid values can be obtained.
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type DorisSourceInfo struct {
	// The IP address of fe in the Doris cluster
	// Note: This field may return null, indicating that no valid values can be obtained.
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// The fe port number of the Doris cluster
	// Note: This field may return null, indicating that no valid values can be obtained.
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Account of the Doris cluster
	// Note: This field may return null, indicating that no valid values can be obtained.
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// Password of the Doris cluster
	// Note: This field may return null, indicating that no valid values can be obtained.
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

// Predefined struct for user
type ExecuteParametrizedQueryRequestParams struct {
	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// SQL query statement
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// Query parameter array.
	QueryParameter []*PropertiesMap `json:"QueryParameter,omitnil,omitempty" name:"QueryParameter"`

	// Page number, which is 1 by default.
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// Number of records per page, which is 10 by default.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`
}

type ExecuteParametrizedQueryRequest struct {
	*tchttp.BaseRequest
	
	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// SQL query statement
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// Query parameter array.
	QueryParameter []*PropertiesMap `json:"QueryParameter,omitnil,omitempty" name:"QueryParameter"`

	// Page number, which is 1 by default.
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// Number of records per page, which is 10 by default.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`
}

func (r *ExecuteParametrizedQueryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteParametrizedQueryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Database")
	delete(f, "Sql")
	delete(f, "QueryParameter")
	delete(f, "PageNum")
	delete(f, "PageSize")
	delete(f, "UserName")
	delete(f, "PassWord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExecuteParametrizedQueryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExecuteParametrizedQueryResponseParams struct {
	// Total records of query results
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Field name array of query results
	Fields []*string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// Field type array of query results
	FieldTypes []*string `json:"FieldTypes,omitnil,omitempty" name:"FieldTypes"`

	// Returned data record array. Each element is an array, containing the value of the corresponding field.
	Rows []*Rows `json:"Rows,omitnil,omitempty" name:"Rows"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExecuteParametrizedQueryResponse struct {
	*tchttp.BaseResponse
	Response *ExecuteParametrizedQueryResponseParams `json:"Response"`
}

func (r *ExecuteParametrizedQueryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteParametrizedQueryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExecuteSelectQueryRequestParams struct {
	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// SQL query statements only support select statements.
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// Page number, which is 1 by default.
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// Number of records per page, which is 10 by default.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`
}

type ExecuteSelectQueryRequest struct {
	*tchttp.BaseRequest
	
	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// SQL query statements only support select statements.
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// Page number, which is 1 by default.
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// Number of records per page, which is 10 by default.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`
}

func (r *ExecuteSelectQueryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteSelectQueryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Database")
	delete(f, "Query")
	delete(f, "PageNum")
	delete(f, "PageSize")
	delete(f, "UserName")
	delete(f, "PassWord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExecuteSelectQueryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExecuteSelectQueryResponseParams struct {
	// Total records of query results
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Field name array of query results
	Fields []*string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// Field type array of query results
	FieldTypes []*string `json:"FieldTypes,omitnil,omitempty" name:"FieldTypes"`

	// Returned data record array. Each element is an array, containing the value of the corresponding field.
	Rows []*Rows `json:"Rows,omitnil,omitempty" name:"Rows"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExecuteSelectQueryResponse struct {
	*tchttp.BaseResponse
	Response *ExecuteSelectQueryResponseParams `json:"Response"`
}

func (r *ExecuteSelectQueryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteSelectQueryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FrontEndRule struct {
	// ID sequence
	// Note: This field may return null, indicating that no valid values can be obtained.
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// Rule name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Detailed rules
	// Note: This field may return null, indicating that no valid values can be obtained.
	Rule *string `json:"Rule,omitnil,omitempty" name:"Rule"`
}

type IndexInfo struct {
	// Index name
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdxName *string `json:"IdxName,omitnil,omitempty" name:"IdxName"`

	// Column name for creating the index
	// Note: This field may return null, indicating that no valid values can be obtained.
	ColumnName *string `json:"ColumnName,omitnil,omitempty" name:"ColumnName"`

	// Index type:
	// INVERTED: inverted index
	// NGRAM_BF: N-Gram index
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdxType *string `json:"IdxType,omitnil,omitempty" name:"IdxType"`

	// Index attributes
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdxProperties []*Property `json:"IdxProperties,omitnil,omitempty" name:"IdxProperties"`

	// Index description
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdxComment *string `json:"IdxComment,omitnil,omitempty" name:"IdxComment"`
}

// Predefined struct for user
type InsertDatasToTableRequestParams struct {
	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Table name
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// Whether to use the strict mode
	Strict *bool `json:"Strict,omitnil,omitempty" name:"Strict"`

	// Maximum filtration ratio, ranging from 0 to 1.0
	MaxFilterRatio *float64 `json:"MaxFilterRatio,omitnil,omitempty" name:"MaxFilterRatio"`

	// Array of column names
	Columns []*string `json:"Columns,omitnil,omitempty" name:"Columns"`

	// Data line
	Rows []*Rows `json:"Rows,omitnil,omitempty" name:"Rows"`

	// Tags for inserting data
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`

	// Column type
	ColumnTypes *string `json:"ColumnTypes,omitnil,omitempty" name:"ColumnTypes"`
}

type InsertDatasToTableRequest struct {
	*tchttp.BaseRequest
	
	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Table name
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// Whether to use the strict mode
	Strict *bool `json:"Strict,omitnil,omitempty" name:"Strict"`

	// Maximum filtration ratio, ranging from 0 to 1.0
	MaxFilterRatio *float64 `json:"MaxFilterRatio,omitnil,omitempty" name:"MaxFilterRatio"`

	// Array of column names
	Columns []*string `json:"Columns,omitnil,omitempty" name:"Columns"`

	// Data line
	Rows []*Rows `json:"Rows,omitnil,omitempty" name:"Rows"`

	// Tags for inserting data
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`

	// Column type
	ColumnTypes *string `json:"ColumnTypes,omitnil,omitempty" name:"ColumnTypes"`
}

func (r *InsertDatasToTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InsertDatasToTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Database")
	delete(f, "Table")
	delete(f, "Strict")
	delete(f, "MaxFilterRatio")
	delete(f, "Columns")
	delete(f, "Rows")
	delete(f, "Label")
	delete(f, "UserName")
	delete(f, "PassWord")
	delete(f, "ColumnTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InsertDatasToTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InsertDatasToTableResponseParams struct {
	// Whether the insertion operation is successful
	Success *bool `json:"Success,omitnil,omitempty" name:"Success"`

	// Message description of the operation result
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Number of inserted data rows
	InsertCount *string `json:"InsertCount,omitnil,omitempty" name:"InsertCount"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InsertDatasToTableResponse struct {
	*tchttp.BaseResponse
	Response *InsertDatasToTableResponseParams `json:"Response"`
}

func (r *InsertDatasToTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InsertDatasToTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceConfigItem struct {
	// key
	ConfKey *string `json:"ConfKey,omitnil,omitempty" name:"ConfKey"`

	// value
	ConfValue *string `json:"ConfValue,omitnil,omitempty" name:"ConfValue"`
}

type InstanceInfo struct {
	// Cluster instance ID, "cdw-xxxx" string type
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Cluster instance name
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Status,
	// Init is being created. Serving is running. 
	// Deleted indicates the cluster has been terminated. Deleting indicates the cluster is being terminated.
	// Modify indicates the cluster is being changed.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Version
	// Note: This field may return null, indicating that no valid values can be obtained.
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// Region, ap-guangzhou
	// Note: This field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Availability zone, ap-guangzhou-3
	// Note: This field may return null, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// VPC name
	// Note: This field may return null, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Subnet name
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Payment type: hour and prepay
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Creation time
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Expiration time
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// Data node description information
	// Note: This field may return null, indicating that no valid values can be obtained.
	MasterSummary *NodesSummary `json:"MasterSummary,omitnil,omitempty" name:"MasterSummary"`

	// Zookeeper node description information
	// Note: This field may return null, indicating that no valid values can be obtained.
	CoreSummary *NodesSummary `json:"CoreSummary,omitnil,omitempty" name:"CoreSummary"`

	// High availability, being true or false
	// Note: This field may return null, indicating that no valid values can be obtained.
	HA *string `json:"HA,omitnil,omitempty" name:"HA"`

	// High availability type:
	// 0: non-high availability
	// 1: read high availability
	// 2: read-write high availability
	// Note: This field may return null, indicating that no valid values can be obtained.
	HaType *int64 `json:"HaType,omitnil,omitempty" name:"HaType"`

	// Access address. Example: 10.0.0.1:9000
	// Note: This field may return null, indicating that no valid values can be obtained.
	AccessInfo *string `json:"AccessInfo,omitnil,omitempty" name:"AccessInfo"`

	// Record ID, in numerical type
	// Note: This field may return null, indicating that no valid values can be obtained.
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// Region ID, indicating the region
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Note about availability zone, such as Guangzhou Zone 2
	// Note: This field may return null, indicating that no valid values can be obtained.
	ZoneDesc *string `json:"ZoneDesc,omitnil,omitempty" name:"ZoneDesc"`

	// Error process description information
	// Note: This field may return null, indicating that no valid values can be obtained.
	FlowMsg *string `json:"FlowMsg,omitnil,omitempty" name:"FlowMsg"`

	// Status description, such as "running"
	// Note: This field may return null, indicating that no valid values can be obtained.
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// Automatic renewal marker
	// Note: This field may return null, indicating that no valid values can be obtained.
	RenewFlag *bool `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// Tag list
	// Note: This field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// Monitoring Information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Monitor *string `json:"Monitor,omitnil,omitempty" name:"Monitor"`

	// Whether to enable logs.
	// Note: This field may return null, indicating that no valid values can be obtained.
	HasClsTopic *bool `json:"HasClsTopic,omitnil,omitempty" name:"HasClsTopic"`

	// Log Topic ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClsTopicId *string `json:"ClsTopicId,omitnil,omitempty" name:"ClsTopicId"`

	// Logset ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClsLogSetId *string `json:"ClsLogSetId,omitnil,omitempty" name:"ClsLogSetId"`

	// Whether to support XML configuration management.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnableXMLConfig *int64 `json:"EnableXMLConfig,omitnil,omitempty" name:"EnableXMLConfig"`

	// Region
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegionDesc *string `json:"RegionDesc,omitnil,omitempty" name:"RegionDesc"`

	// Elastic network interface address
	// Note: This field may return null, indicating that no valid values can be obtained.
	Eip *string `json:"Eip,omitnil,omitempty" name:"Eip"`

	// Cold and hot stratification coefficient
	// Note: This field may return null, indicating that no valid values can be obtained.
	CosMoveFactor *int64 `json:"CosMoveFactor,omitnil,omitempty" name:"CosMoveFactor"`

	// external/local/yunti
	// Note: This field may return null, indicating that no valid values can be obtained.
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// COS bucket
	// Note: This field may return null, indicating that no valid values can be obtained.
	CosBucketName *string `json:"CosBucketName,omitnil,omitempty" name:"CosBucketName"`

	// cbs
	// Note: This field may return null, indicating that no valid values can be obtained.
	CanAttachCbs *bool `json:"CanAttachCbs,omitnil,omitempty" name:"CanAttachCbs"`

	// Minor versions
	// Note: This field may return null, indicating that no valid values can be obtained.
	BuildVersion *string `json:"BuildVersion,omitnil,omitempty" name:"BuildVersion"`

	// Component Information
	// Note: The return type here is map[string]struct, not the string type displayed. You can refer to "Sample Value" to parse the data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Components *string `json:"Components,omitnil,omitempty" name:"Components"`

	// Determine whether the audit log table has a catalog field.
	// Note: This field may return null, indicating that no valid values can be obtained.
	//
	// Deprecated: IfExistCatalog is deprecated.
	IfExistCatalog *int64 `json:"IfExistCatalog,omitnil,omitempty" name:"IfExistCatalog"`

	// Page features, used to block some page entrances on the front end.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Characteristic []*string `json:"Characteristic,omitnil,omitempty" name:"Characteristic"`

	// Timeout period, in seconds
	// Note: This field may return null, indicating that no valid values can be obtained.
	RestartTimeout *string `json:"RestartTimeout,omitnil,omitempty" name:"RestartTimeout"`

	// The timeout time for the graceful restart of the kernel. If it is -1, it means it is not set.
	// Note: This field may return null, indicating that no valid values can be obtained.
	GraceShutdownWaitSeconds *string `json:"GraceShutdownWaitSeconds,omitnil,omitempty" name:"GraceShutdownWaitSeconds"`

	// Whether the table name is case sensitive, 0 refers to sensitive, 1 refers to insensitive, compared in lowercase; 2 refers to insensitive, and the table name is changed to lowercase for storage.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CaseSensitive *int64 `json:"CaseSensitive,omitnil,omitempty" name:"CaseSensitive"`

	// Whether users can bind security groups.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsWhiteSGs *bool `json:"IsWhiteSGs,omitnil,omitempty" name:"IsWhiteSGs"`

	// Bound security group information
	// Note: This field may return null, indicating that no valid values can be obtained.
	BindSGs []*string `json:"BindSGs,omitnil,omitempty" name:"BindSGs"`

	// Whether it is a multi-AZ.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnableMultiZones *bool `json:"EnableMultiZones,omitnil,omitempty" name:"EnableMultiZones"`

	// User availability zone and subnet information
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserNetworkInfos *string `json:"UserNetworkInfos,omitnil,omitempty" name:"UserNetworkInfos"`

	// Whether to enable hot and cold stratification. 0 refers to disabled, and 1 refers to enabled.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnableCoolDown *int64 `json:"EnableCoolDown,omitnil,omitempty" name:"EnableCoolDown"`

	// COS buckets are used for hot and cold stratification
	// Note: This field may return null, indicating that no valid values can be obtained.
	CoolDownBucket *string `json:"CoolDownBucket,omitnil,omitempty" name:"CoolDownBucket"`
}

type InstanceNode struct {
	// IP address
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Model, such as S1
	Spec *string `json:"Spec,omitnil,omitempty" name:"Spec"`

	// Number of CPU cores
	Core *int64 `json:"Core,omitnil,omitempty" name:"Core"`

	// Memory size
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Disk type
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Disk size
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// The name of the clickhouse cluster to which it belongs.
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// Status
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// rip
	// Note: This field may return null, indicating that no valid values can be obtained.
	Rip *string `json:"Rip,omitnil,omitempty" name:"Rip"`

	// FE node role
	// Note: This field may return null, indicating that no valid values can be obtained.
	FeRole *string `json:"FeRole,omitnil,omitempty" name:"FeRole"`

	// UUID
	// Note: This field may return null, indicating that no valid values can be obtained.
	UUID *string `json:"UUID,omitnil,omitempty" name:"UUID"`
}

type InstanceOperation struct {
	// Operation name, such as create_instance, and scaleout_instance
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Operation result. Success indicates success; Fail indicates failure.
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// Operation name description, such as create, and modify the cluster name
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// Operation level, such as Critical, Normal
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// Operation level description, such as high risk, and normal
	LevelDesc *string `json:"LevelDesc,omitnil,omitempty" name:"LevelDesc"`

	// Operation start time
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Operation end time
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Operation result description, such as Success and Fail.
	ResultDesc *string `json:"ResultDesc,omitnil,omitempty" name:"ResultDesc"`

	// Operation user ID
	OperateUin *string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`

	// The jobid corresponding to the operation
	JobId *int64 `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Operation details
	OperationDetail *string `json:"OperationDetail,omitnil,omitempty" name:"OperationDetail"`
}

type ListInfo struct {
	// Partition name
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	PartitionName *string `json:"PartitionName,omitnil,omitempty" name:"PartitionName"`

	// Enumeration values of each partition column
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnumValues []*string `json:"EnumValues,omitnil,omitempty" name:"EnumValues"`
}

// Predefined struct for user
type ModifyDatabaseTableAccessRequestParams struct {
	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Table name. If it is null, it indicates that the entire database is authorized.
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// Permission list
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`

	// Role name, if authorized to the role
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// Operation type: GRANT or REVOKE
	GrantOrRevoke *string `json:"GrantOrRevoke,omitnil,omitempty" name:"GrantOrRevoke"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`
}

type ModifyDatabaseTableAccessRequest struct {
	*tchttp.BaseRequest
	
	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Table name. If it is null, it indicates that the entire database is authorized.
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// Permission list
	Privileges []*string `json:"Privileges,omitnil,omitempty" name:"Privileges"`

	// Role name, if authorized to the role
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// Operation type: GRANT or REVOKE
	GrantOrRevoke *string `json:"GrantOrRevoke,omitnil,omitempty" name:"GrantOrRevoke"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`
}

func (r *ModifyDatabaseTableAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatabaseTableAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Database")
	delete(f, "Table")
	delete(f, "Privileges")
	delete(f, "Role")
	delete(f, "GrantOrRevoke")
	delete(f, "UserName")
	delete(f, "PassWord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDatabaseTableAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatabaseTableAccessResponseParams struct {
	// Whether the operation is successful
	Success *bool `json:"Success,omitnil,omitempty" name:"Success"`

	// Message description of the operation result
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDatabaseTableAccessResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDatabaseTableAccessResponseParams `json:"Response"`
}

func (r *ModifyDatabaseTableAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatabaseTableAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceKeyValConfigsRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// File name
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Add configuration list
	AddItems []*InstanceConfigItem `json:"AddItems,omitnil,omitempty" name:"AddItems"`

	// Update configuration list
	UpdateItems []*InstanceConfigItem `json:"UpdateItems,omitnil,omitempty" name:"UpdateItems"`

	// Delete configuration list
	DelItems []*InstanceConfigItem `json:"DelItems,omitnil,omitempty" name:"DelItems"`

	// Remarks (within 50 words)
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Hot update list
	HotUpdateItems []*InstanceConfigItem `json:"HotUpdateItems,omitnil,omitempty" name:"HotUpdateItems"`

	// Delete configuration list
	DeleteItems *InstanceConfigItem `json:"DeleteItems,omitnil,omitempty" name:"DeleteItems"`

	// IP address
	IPAddress *string `json:"IPAddress,omitnil,omitempty" name:"IPAddress"`
}

type ModifyInstanceKeyValConfigsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// File name
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// Add configuration list
	AddItems []*InstanceConfigItem `json:"AddItems,omitnil,omitempty" name:"AddItems"`

	// Update configuration list
	UpdateItems []*InstanceConfigItem `json:"UpdateItems,omitnil,omitempty" name:"UpdateItems"`

	// Delete configuration list
	DelItems []*InstanceConfigItem `json:"DelItems,omitnil,omitempty" name:"DelItems"`

	// Remarks (within 50 words)
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Hot update list
	HotUpdateItems []*InstanceConfigItem `json:"HotUpdateItems,omitnil,omitempty" name:"HotUpdateItems"`

	// Delete configuration list
	DeleteItems *InstanceConfigItem `json:"DeleteItems,omitnil,omitempty" name:"DeleteItems"`

	// IP address
	IPAddress *string `json:"IPAddress,omitnil,omitempty" name:"IPAddress"`
}

func (r *ModifyInstanceKeyValConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceKeyValConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FileName")
	delete(f, "AddItems")
	delete(f, "UpdateItems")
	delete(f, "DelItems")
	delete(f, "Message")
	delete(f, "HotUpdateItems")
	delete(f, "DeleteItems")
	delete(f, "IPAddress")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceKeyValConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceKeyValConfigsResponseParams struct {
	// Error message
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceKeyValConfigsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceKeyValConfigsResponseParams `json:"Response"`
}

func (r *ModifyInstanceKeyValConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceKeyValConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Newly modified instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type ModifyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Newly modified instance name
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
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
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyNodeStatusRequestParams struct {
	// Cluster ID, such as cdwch-xxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Node information
	NodeInfos []*NodeInfos `json:"NodeInfos,omitnil,omitempty" name:"NodeInfos"`

	// Node operation
	OperationCode *string `json:"OperationCode,omitnil,omitempty" name:"OperationCode"`

	// Timeout period (s)
	RestartTimeOut *string `json:"RestartTimeOut,omitnil,omitempty" name:"RestartTimeOut"`
}

type ModifyNodeStatusRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID, such as cdwch-xxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Node information
	NodeInfos []*NodeInfos `json:"NodeInfos,omitnil,omitempty" name:"NodeInfos"`

	// Node operation
	OperationCode *string `json:"OperationCode,omitnil,omitempty" name:"OperationCode"`

	// Timeout period (s)
	RestartTimeOut *string `json:"RestartTimeOut,omitnil,omitempty" name:"RestartTimeOut"`
}

func (r *ModifyNodeStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNodeStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NodeInfos")
	delete(f, "OperationCode")
	delete(f, "RestartTimeOut")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNodeStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNodeStatusResponseParams struct {
	// Process related information
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Error message
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNodeStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNodeStatusResponseParams `json:"Response"`
}

func (r *ModifyNodeStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNodeStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityGroupsRequestParams struct {
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Security group information before modification
	OldSecurityGroupIds []*string `json:"OldSecurityGroupIds,omitnil,omitempty" name:"OldSecurityGroupIds"`

	// Modified security group information
	ModifySecurityGroupIds []*string `json:"ModifySecurityGroupIds,omitnil,omitempty" name:"ModifySecurityGroupIds"`
}

type ModifySecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Security group information before modification
	OldSecurityGroupIds []*string `json:"OldSecurityGroupIds,omitnil,omitempty" name:"OldSecurityGroupIds"`

	// Modified security group information
	ModifySecurityGroupIds []*string `json:"ModifySecurityGroupIds,omitnil,omitempty" name:"ModifySecurityGroupIds"`
}

func (r *ModifySecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "OldSecurityGroupIds")
	delete(f, "ModifySecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityGroupsResponseParams struct {
	// Error message
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecurityGroupsResponseParams `json:"Response"`
}

func (r *ModifySecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserBindWorkloadGroupRequestParams struct {
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The user information of the resource group needs to be bound. If an account has information of multiple hosts, all information needs to be uploaded.
	BindUsers []*BindUser `json:"BindUsers,omitnil,omitempty" name:"BindUsers"`

	// Name of the resource group bound before modification
	OldWorkloadGroupName *string `json:"OldWorkloadGroupName,omitnil,omitempty" name:"OldWorkloadGroupName"`

	// Name of the resource group bound after modification
	NewWorkloadGroupName *string `json:"NewWorkloadGroupName,omitnil,omitempty" name:"NewWorkloadGroupName"`
}

type ModifyUserBindWorkloadGroupRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The user information of the resource group needs to be bound. If an account has information of multiple hosts, all information needs to be uploaded.
	BindUsers []*BindUser `json:"BindUsers,omitnil,omitempty" name:"BindUsers"`

	// Name of the resource group bound before modification
	OldWorkloadGroupName *string `json:"OldWorkloadGroupName,omitnil,omitempty" name:"OldWorkloadGroupName"`

	// Name of the resource group bound after modification
	NewWorkloadGroupName *string `json:"NewWorkloadGroupName,omitnil,omitempty" name:"NewWorkloadGroupName"`
}

func (r *ModifyUserBindWorkloadGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserBindWorkloadGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BindUsers")
	delete(f, "OldWorkloadGroupName")
	delete(f, "NewWorkloadGroupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserBindWorkloadGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserBindWorkloadGroupResponseParams struct {
	// Error message
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyUserBindWorkloadGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserBindWorkloadGroupResponseParams `json:"Response"`
}

func (r *ModifyUserBindWorkloadGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserBindWorkloadGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserPrivilegesV3RequestParams struct {
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Username
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// User permission
	UserPrivileges *UpdateUserPrivileges `json:"UserPrivileges,omitnil,omitempty" name:"UserPrivileges"`

	// The IP address of the user link	
	WhiteHost *string `json:"WhiteHost,omitnil,omitempty" name:"WhiteHost"`
}

type ModifyUserPrivilegesV3Request struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Username
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// User permission
	UserPrivileges *UpdateUserPrivileges `json:"UserPrivileges,omitnil,omitempty" name:"UserPrivileges"`

	// The IP address of the user link	
	WhiteHost *string `json:"WhiteHost,omitnil,omitempty" name:"WhiteHost"`
}

func (r *ModifyUserPrivilegesV3Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserPrivilegesV3Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "UserPrivileges")
	delete(f, "WhiteHost")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserPrivilegesV3Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserPrivilegesV3ResponseParams struct {
	// Error message; null means no error.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyUserPrivilegesV3Response struct {
	*tchttp.BaseResponse
	Response *ModifyUserPrivilegesV3ResponseParams `json:"Response"`
}

func (r *ModifyUserPrivilegesV3Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserPrivilegesV3Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWorkloadGroupRequestParams struct {
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Modified resource group information
	WorkloadGroup *WorkloadGroupConfig `json:"WorkloadGroup,omitnil,omitempty" name:"WorkloadGroup"`
}

type ModifyWorkloadGroupRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Modified resource group information
	WorkloadGroup *WorkloadGroupConfig `json:"WorkloadGroup,omitnil,omitempty" name:"WorkloadGroup"`
}

func (r *ModifyWorkloadGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkloadGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "WorkloadGroup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWorkloadGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWorkloadGroupResponseParams struct {
	// Error message	
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyWorkloadGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWorkloadGroupResponseParams `json:"Response"`
}

func (r *ModifyWorkloadGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkloadGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWorkloadGroupStatusRequestParams struct {
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Whether to enable resource group: open and close
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`
}

type ModifyWorkloadGroupStatusRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Whether to enable resource group: open and close
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`
}

func (r *ModifyWorkloadGroupStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkloadGroupStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "OperationType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWorkloadGroupStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWorkloadGroupStatusResponseParams struct {
	// Error message	
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyWorkloadGroupStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWorkloadGroupStatusResponseParams `json:"Response"`
}

func (r *ModifyWorkloadGroupStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkloadGroupStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetworkInfo struct {
	// Availability zone
	// Note: This field may return null, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// Subnet ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// The number of available IP addresses in the current subnet
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubnetIpNum *int64 `json:"SubnetIpNum,omitnil,omitempty" name:"SubnetIpNum"`
}

type NodeInfo struct {
	// User IP
	// Note: This field may return null, indicating that no valid values can be obtained.
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Node status
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Node role name
	// Note: This field may return null, indicating that no valid values can be obtained.
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// Component name
	// Note: This field may return null, indicating that no valid values can be obtained.
	ComponentName *string `json:"ComponentName,omitnil,omitempty" name:"ComponentName"`

	// Node role
	// Note: This field may return null, indicating that no valid values can be obtained.
	NodeRole *string `json:"NodeRole,omitnil,omitempty" name:"NodeRole"`

	// The time when the node was last restarted
	// Note: This field may return null, indicating that no valid values can be obtained.
	LastRestartTime *string `json:"LastRestartTime,omitnil,omitempty" name:"LastRestartTime"`

	// The availability zone where the node is located
	// Note: This field may return null, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type NodeInfos struct {
	// Node name
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// Node status
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Node IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Node role
	NodeRole *string `json:"NodeRole,omitnil,omitempty" name:"NodeRole"`

	// Component name
	ComponentName *string `json:"ComponentName,omitnil,omitempty" name:"ComponentName"`

	// Last restart time
	LastRestartTime *string `json:"LastRestartTime,omitnil,omitempty" name:"LastRestartTime"`
}

type NodesSummary struct {
	// Model, such as S1
	Spec *string `json:"Spec,omitnil,omitempty" name:"Spec"`

	// Number of nodes
	NodeSize *int64 `json:"NodeSize,omitnil,omitempty" name:"NodeSize"`

	// Number of CPU cores, in counts
	Core *int64 `json:"Core,omitnil,omitempty" name:"Core"`

	// Memory size, in GB
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Disk size, in GB
	Disk *int64 `json:"Disk,omitnil,omitempty" name:"Disk"`

	// Disk type
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// Disk description
	DiskDesc *string `json:"DiskDesc,omitnil,omitempty" name:"DiskDesc"`

	// Information of mounted cloud disks
	// Note: This field may return null, indicating that no valid values can be obtained.
	AttachCBSSpec *AttachCBSSpec `json:"AttachCBSSpec,omitnil,omitempty" name:"AttachCBSSpec"`

	// Sub-product name
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubProductType *string `json:"SubProductType,omitnil,omitempty" name:"SubProductType"`

	// Specified cores
	// Note: This field may return null, indicating that no valid values can be obtained.
	SpecCore *int64 `json:"SpecCore,omitnil,omitempty" name:"SpecCore"`

	// Specified memory
	// Note: This field may return null, indicating that no valid values can be obtained.
	SpecMemory *int64 `json:"SpecMemory,omitnil,omitempty" name:"SpecMemory"`

	// Disk size
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskCount *int64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`

	// Whether it is encrypted.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Encrypt *int64 `json:"Encrypt,omitnil,omitempty" name:"Encrypt"`

	// Maximum disk
	// Note: This field may return null, indicating that no valid values can be obtained.
	MaxDiskSize *int64 `json:"MaxDiskSize,omitnil,omitempty" name:"MaxDiskSize"`
}

type Partition struct {
	// Partition type:
	// ●Range: The partition column is usually of time or integer type. Four writing methods are supported.
	// ●List: The partition value is an enumeration value.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	PartitionType *string `json:"PartitionType,omitnil,omitempty" name:"PartitionType"`

	// Whether to automatically partition
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	AutoPartition *bool `json:"AutoPartition,omitnil,omitempty" name:"AutoPartition"`

	// Partition information when the partition type is Range
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	RangeInfos []*RangeInfo `json:"RangeInfos,omitnil,omitempty" name:"RangeInfos"`

	// Partition information when the partition type is List
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	ListInfos []*ListInfo `json:"ListInfos,omitnil,omitempty" name:"ListInfos"`
}

type PermissionHostInfo struct {
	// A list of user permissions in the global scope, which can be applied to all databases and tables.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	GlobalPermissions []*string `json:"GlobalPermissions,omitnil,omitempty" name:"GlobalPermissions"`

	// The key is the database name, and the value is the permission list of the user on the database.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DatabasePermissions []*DatabasePermissions `json:"DatabasePermissions,omitnil,omitempty" name:"DatabasePermissions"`

	// The key is the full name of the table, and the value is the permission list of the user on the table.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	TablePermissions []*TablePermissions `json:"TablePermissions,omitnil,omitempty" name:"TablePermissions"`
}

type PropertiesMap struct {
	// key
	// Note: This field may return null, indicating that no valid values can be obtained.
	PropertyKey *string `json:"PropertyKey,omitnil,omitempty" name:"PropertyKey"`

	// value
	// Note: This field may return null, indicating that no valid values can be obtained.
	PropertyValue *string `json:"PropertyValue,omitnil,omitempty" name:"PropertyValue"`
}

type Property struct {
	// Attribute key
	// Note: This field may return null, indicating that no valid values can be obtained.
	PropertyKey *string `json:"PropertyKey,omitnil,omitempty" name:"PropertyKey"`

	// Attribute value
	// Note: This field may return null, indicating that no valid values can be obtained.
	PropertyValue *string `json:"PropertyValue,omitnil,omitempty" name:"PropertyValue"`
}

type QueryDetails struct {
	// Initiating User
	// Note: This field may return null, indicating that no valid values can be obtained.
	Initiator *string `json:"Initiator,omitnil,omitempty" name:"Initiator"`

	// Access source address
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	SourceAddress *string `json:"SourceAddress,omitnil,omitempty" name:"SourceAddress"`

	// Initial request ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	InitialRequestId *string `json:"InitialRequestId,omitnil,omitempty" name:"InitialRequestId"`

	// catalog name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// Database name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// SQL type
	// Note: This field may return null, indicating that no valid values can be obtained.
	SQLType *string `json:"SQLType,omitnil,omitempty" name:"SQLType"`

	// SQL statement
	// Note: This field may return null, indicating that no valid values can be obtained.
	SQLStatement *string `json:"SQLStatement,omitnil,omitempty" name:"SQLStatement"`

	// Execution start time
	// Note: This field may return null, indicating that no valid values can be obtained.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Runtime (s)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// The number of read rows
	// Note: This field may return null, indicating that no valid values can be obtained.
	RowsRead *uint64 `json:"RowsRead,omitnil,omitempty" name:"RowsRead"`

	// Read data volume (MB)
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataRead *float64 `json:"DataRead,omitnil,omitempty" name:"DataRead"`

	// Memory usage (MB)
	// Note: This field may return null, indicating that no valid values can be obtained.
	MemoryUsage *float64 `json:"MemoryUsage,omitnil,omitempty" name:"MemoryUsage"`
}

// Predefined struct for user
type QueryTableDataRequestParams struct {
	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Table name
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// Array of fields to be queried
	SelectedFields []*string `json:"SelectedFields,omitnil,omitempty" name:"SelectedFields"`

	// Page number, which is 1 by default.
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// Number of records per page, which is 10 by default.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`
}

type QueryTableDataRequest struct {
	*tchttp.BaseRequest
	
	// Database name
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// Table name
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// Array of fields to be queried
	SelectedFields []*string `json:"SelectedFields,omitnil,omitempty" name:"SelectedFields"`

	// Page number, which is 1 by default.
	PageNum *uint64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// Number of records per page, which is 10 by default.
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`
}

func (r *QueryTableDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryTableDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Database")
	delete(f, "Table")
	delete(f, "SelectedFields")
	delete(f, "PageNum")
	delete(f, "PageSize")
	delete(f, "UserName")
	delete(f, "PassWord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryTableDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryTableDataResponseParams struct {
	// Total records of query results
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Field name array of query results
	Fields []*string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// Field type array of query results
	FieldTypes []*string `json:"FieldTypes,omitnil,omitempty" name:"FieldTypes"`

	// Returned data record array. Each element is an array, containing the value of the corresponding field.
	Rows []*Rows `json:"Rows,omitnil,omitempty" name:"Rows"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryTableDataResponse struct {
	*tchttp.BaseResponse
	Response *QueryTableDataResponseParams `json:"Response"`
}

func (r *QueryTableDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryTableDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RangeInfo struct {
	// Range partition type:
	// ●FIXED: Define the left closed and right open interval of the partition.
	// ●LESS THAN: Only define the upper bound of the partition.
	// ●BATCH RANGE: Batch create RANGE partitions of numeric and time types, define the left closed and right open intervals of the partitions, and set the step size.
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	RangeType *string `json:"RangeType,omitnil,omitempty" name:"RangeType"`

	// Partition name
	// Note: This field may return null, indicating that no valid values can be obtained.
	PartitionName *string `json:"PartitionName,omitnil,omitempty" name:"PartitionName"`

	// The left-closed interval of each partition column when RangeType is FIXED or BATCH RANGE
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Left *string `json:"Left,omitnil,omitempty" name:"Left"`

	// The right open interval of each partition column when RangeType is FIXED or BATCH RANGE
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Right *string `json:"Right,omitnil,omitempty" name:"Right"`

	// The upper bound of each partition column when RangeType is LESS THAN
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Max *string `json:"Max,omitnil,omitempty" name:"Max"`

	// RangeType is the step size of BATCH RANGE
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	StepLength *int64 `json:"StepLength,omitnil,omitempty" name:"StepLength"`

	// Fill it in when RangeType is BATCH RANGE or automatic partitioning. It indicates the step size unit when the partition column is of time type.
	// ●YEAR: year
	// ●MONTH: month
	// ●WEEK: week
	// ●DAY: day
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`
}

// Predefined struct for user
type RecoverBackUpJobRequestParams struct {
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Task ID
	BackUpJobId *int64 `json:"BackUpJobId,omitnil,omitempty" name:"BackUpJobId"`

	// Number of new table replicas recovered
	ReplicationNum *int64 `json:"ReplicationNum,omitnil,omitempty" name:"ReplicationNum"`

	// Whether to retain the configuration in the source table during recovery. 1 indicates that the configurations in the source table are retained.
	ReserveSourceConfig *int64 `json:"ReserveSourceConfig,omitnil,omitempty" name:"ReserveSourceConfig"`

	// 0: default; 1: cos recovery
	RecoverType *int64 `json:"RecoverType,omitnil,omitempty" name:"RecoverType"`

	// CosSourceInfo object
	CosSourceInfo *CosSourceInfo `json:"CosSourceInfo,omitnil,omitempty" name:"CosSourceInfo"`

	// 0: default; 1: regular execution
	ScheduleType *int64 `json:"ScheduleType,omitnil,omitempty" name:"ScheduleType"`

	// YY-MM-DD Hour : Minute : Second
	NextTime *string `json:"NextTime,omitnil,omitempty" name:"NextTime"`

	// Scheduling name
	ScheduleName *string `json:"ScheduleName,omitnil,omitempty" name:"ScheduleName"`

	// create update
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// Recovery granularity: All, Database, and Table
	RecoverScope *string `json:"RecoverScope,omitnil,omitempty" name:"RecoverScope"`

	// Recover database: If you back up by database, this field is required. Use commas to separate databases.
	RecoverDatabase *string `json:"RecoverDatabase,omitnil,omitempty" name:"RecoverDatabase"`
}

type RecoverBackUpJobRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Task ID
	BackUpJobId *int64 `json:"BackUpJobId,omitnil,omitempty" name:"BackUpJobId"`

	// Number of new table replicas recovered
	ReplicationNum *int64 `json:"ReplicationNum,omitnil,omitempty" name:"ReplicationNum"`

	// Whether to retain the configuration in the source table during recovery. 1 indicates that the configurations in the source table are retained.
	ReserveSourceConfig *int64 `json:"ReserveSourceConfig,omitnil,omitempty" name:"ReserveSourceConfig"`

	// 0: default; 1: cos recovery
	RecoverType *int64 `json:"RecoverType,omitnil,omitempty" name:"RecoverType"`

	// CosSourceInfo object
	CosSourceInfo *CosSourceInfo `json:"CosSourceInfo,omitnil,omitempty" name:"CosSourceInfo"`

	// 0: default; 1: regular execution
	ScheduleType *int64 `json:"ScheduleType,omitnil,omitempty" name:"ScheduleType"`

	// YY-MM-DD Hour : Minute : Second
	NextTime *string `json:"NextTime,omitnil,omitempty" name:"NextTime"`

	// Scheduling name
	ScheduleName *string `json:"ScheduleName,omitnil,omitempty" name:"ScheduleName"`

	// create update
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// Recovery granularity: All, Database, and Table
	RecoverScope *string `json:"RecoverScope,omitnil,omitempty" name:"RecoverScope"`

	// Recover database: If you back up by database, this field is required. Use commas to separate databases.
	RecoverDatabase *string `json:"RecoverDatabase,omitnil,omitempty" name:"RecoverDatabase"`
}

func (r *RecoverBackUpJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverBackUpJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackUpJobId")
	delete(f, "ReplicationNum")
	delete(f, "ReserveSourceConfig")
	delete(f, "RecoverType")
	delete(f, "CosSourceInfo")
	delete(f, "ScheduleType")
	delete(f, "NextTime")
	delete(f, "ScheduleName")
	delete(f, "OperationType")
	delete(f, "RecoverScope")
	delete(f, "RecoverDatabase")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecoverBackUpJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecoverBackUpJobResponseParams struct {
	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RecoverBackUpJobResponse struct {
	*tchttp.BaseResponse
	Response *RecoverBackUpJobResponseParams `json:"Response"`
}

func (r *RecoverBackUpJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverBackUpJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReduceInstanceRequestParams struct {
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Node list
	DelHosts []*string `json:"DelHosts,omitnil,omitempty" name:"DelHosts"`

	// Role (MATER/CORE), MASTER corresponds to FE, CORE corresponds to BE.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// High availability cluster type after scale-in. 0: non-high availability; 1: read high availability; 2: read-write high availability
	HaType *int64 `json:"HaType,omitnil,omitempty" name:"HaType"`
}

type ReduceInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Node list
	DelHosts []*string `json:"DelHosts,omitnil,omitempty" name:"DelHosts"`

	// Role (MATER/CORE), MASTER corresponds to FE, CORE corresponds to BE.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// High availability cluster type after scale-in. 0: non-high availability; 1: read high availability; 2: read-write high availability
	HaType *int64 `json:"HaType,omitnil,omitempty" name:"HaType"`
}

func (r *ReduceInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReduceInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DelHosts")
	delete(f, "Type")
	delete(f, "HaType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReduceInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReduceInstanceResponseParams struct {
	// Process ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Error message
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReduceInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ReduceInstanceResponseParams `json:"Response"`
}

func (r *ReduceInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReduceInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionAreaInfo struct {
	// Region category information, such as south_china, east_china, etc.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Description of the corresponding Name, such as South China, East China, etc.
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// Specific region list 1
	Regions []*RegionInfo `json:"Regions,omitnil,omitempty" name:"Regions"`
}

type RegionInfo struct {
	// Region name, such as ap-guangzhou
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Region description, such as Guangzhou
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// Unique marker of region
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// List of all availability zones in the region
	// Note: This field may return null, indicating that no valid values can be obtained.
	Zones []*ZoneInfo `json:"Zones,omitnil,omitempty" name:"Zones"`

	// Number of clusters in the region
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 0 indicates the international site; 1 indicates not.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsInternationalSite *uint64 `json:"IsInternationalSite,omitnil,omitempty" name:"IsInternationalSite"`

	// Bucket
	// Note: This field may return null, indicating that no valid values can be obtained.
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`
}

// Predefined struct for user
type ResizeDiskRequestParams struct {
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Role (MATER/CORE), MASTER corresponds to FE, CORE corresponds to BE.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Cloud disk size
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`
}

type ResizeDiskRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Role (MATER/CORE), MASTER corresponds to FE, CORE corresponds to BE.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Cloud disk size
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`
}

func (r *ResizeDiskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResizeDiskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Type")
	delete(f, "DiskSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResizeDiskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResizeDiskResponseParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Process ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Error message
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResizeDiskResponse struct {
	*tchttp.BaseResponse
	Response *ResizeDiskResponseParams `json:"Response"`
}

func (r *ResizeDiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResizeDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceSpec struct {
	// Specification name, such as SCH1
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Number of CPU cores
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Memory size, in GB
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// Classification markers, STANDARD/BIGDATA/HIGHIO respectively represent standard type/big data type/high IO.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// System disk description information
	// Note: This field may return null, indicating that no valid values can be obtained.
	SystemDisk *DiskSpec `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// Data disk description information
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataDisk *DiskSpec `json:"DataDisk,omitnil,omitempty" name:"DataDisk"`

	// Limit of the maximum number of nodes
	// Note: This field may return null, indicating that no valid values can be obtained.
	MaxNodeSize *int64 `json:"MaxNodeSize,omitnil,omitempty" name:"MaxNodeSize"`

	// Whether it is available. False indicates sell-out.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Available *bool `json:"Available,omitnil,omitempty" name:"Available"`

	// Specification description information
	// Note: This field may return null, indicating that no valid values can be obtained.
	ComputeSpecDesc *string `json:"ComputeSpecDesc,omitnil,omitempty" name:"ComputeSpecDesc"`

	// CVM inventory
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceQuota *int64 `json:"InstanceQuota,omitnil,omitempty" name:"InstanceQuota"`
}

// Predefined struct for user
type RestartClusterForConfigsRequestParams struct {
	// Cluster ID, such as cdwch-xxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Configuration file's name
	ConfigName *string `json:"ConfigName,omitnil,omitempty" name:"ConfigName"`

	// grace_restart is an elegant scrolling restart. If this parameter is not filled in, restart now by default.
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`
}

type RestartClusterForConfigsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID, such as cdwch-xxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Configuration file's name
	ConfigName *string `json:"ConfigName,omitnil,omitempty" name:"ConfigName"`

	// grace_restart is an elegant scrolling restart. If this parameter is not filled in, restart now by default.
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`
}

func (r *RestartClusterForConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartClusterForConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ConfigName")
	delete(f, "OperationType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartClusterForConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartClusterForConfigsResponseParams struct {
	// Process related information
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Error message
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartClusterForConfigsResponse struct {
	*tchttp.BaseResponse
	Response *RestartClusterForConfigsResponseParams `json:"Response"`
}

func (r *RestartClusterForConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartClusterForConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartClusterForNodeRequestParams struct {
	// Cluster ID, such as cdwch-xxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Configuration file's name
	ConfigName *string `json:"ConfigName,omitnil,omitempty" name:"ConfigName"`

	// Each batch of restarts
	BatchSize *int64 `json:"BatchSize,omitnil,omitempty" name:"BatchSize"`

	// Restart node
	NodeList []*string `json:"NodeList,omitnil,omitempty" name:"NodeList"`

	// False means non-rolling restart, and true means rolling restart.
	RollingRestart *bool `json:"RollingRestart,omitnil,omitempty" name:"RollingRestart"`
}

type RestartClusterForNodeRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID, such as cdwch-xxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Configuration file's name
	ConfigName *string `json:"ConfigName,omitnil,omitempty" name:"ConfigName"`

	// Each batch of restarts
	BatchSize *int64 `json:"BatchSize,omitnil,omitempty" name:"BatchSize"`

	// Restart node
	NodeList []*string `json:"NodeList,omitnil,omitempty" name:"NodeList"`

	// False means non-rolling restart, and true means rolling restart.
	RollingRestart *bool `json:"RollingRestart,omitnil,omitempty" name:"RollingRestart"`
}

func (r *RestartClusterForNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartClusterForNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ConfigName")
	delete(f, "BatchSize")
	delete(f, "NodeList")
	delete(f, "RollingRestart")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartClusterForNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartClusterForNodeResponseParams struct {
	// Process related information
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Error message
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartClusterForNodeResponse struct {
	*tchttp.BaseResponse
	Response *RestartClusterForNodeResponseParams `json:"Response"`
}

func (r *RestartClusterForNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartClusterForNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestoreStatus struct {
	// Recover the task ID
	JobId *int64 `json:"JobId,omitnil,omitempty" name:"JobId"`

	// Recover the task tag
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// Recover the task timestamp
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// Recover the database where the task is located
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Recover the task status
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// Whether to allow import during recovery
	AllowLoad *bool `json:"AllowLoad,omitnil,omitempty" name:"AllowLoad"`

	// Number of replicas
	ReplicationNum *string `json:"ReplicationNum,omitnil,omitempty" name:"ReplicationNum"`

	// Number of replicas
	ReplicaAllocation *string `json:"ReplicaAllocation,omitnil,omitempty" name:"ReplicaAllocation"`

	// Recover object
	RestoreObjects *string `json:"RestoreObjects,omitnil,omitempty" name:"RestoreObjects"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Metadata preparation time
	MetaPreparedTime *string `json:"MetaPreparedTime,omitnil,omitempty" name:"MetaPreparedTime"`

	// Snapshot end time
	SnapshotFinishedTime *string `json:"SnapshotFinishedTime,omitnil,omitempty" name:"SnapshotFinishedTime"`

	// Download end time
	DownloadFinishedTime *string `json:"DownloadFinishedTime,omitnil,omitempty" name:"DownloadFinishedTime"`

	// End time
	FinishedTime *string `json:"FinishedTime,omitnil,omitempty" name:"FinishedTime"`

	// Unfinished tasks
	UnfinishedTasks *string `json:"UnfinishedTasks,omitnil,omitempty" name:"UnfinishedTasks"`

	// Progress
	Progress *string `json:"Progress,omitnil,omitempty" name:"Progress"`

	// Error message
	TaskErrMsg *string `json:"TaskErrMsg,omitnil,omitempty" name:"TaskErrMsg"`

	// Status
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Operation timeout period
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// Whether to maintain the number of replicas in the source table
	ReserveReplica *bool `json:"ReserveReplica,omitnil,omitempty" name:"ReserveReplica"`

	// Whether to maintain dynamic partitions in the source table
	ReserveDynamicPartitionEnable *bool `json:"ReserveDynamicPartitionEnable,omitnil,omitempty" name:"ReserveDynamicPartitionEnable"`

	// Backup instance ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	BackupJobId *int64 `json:"BackupJobId,omitnil,omitempty" name:"BackupJobId"`

	// ID of the snapshot corresponding to the instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type Rows struct {
	// A row of data
	// Note: This field may return null, indicating that no valid values can be obtained.
	DataRow []*string `json:"DataRow,omitnil,omitempty" name:"DataRow"`
}

// Predefined struct for user
type ScaleOutInstanceRequestParams struct {
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Role (MATER/CORE), MASTER corresponds to FE, CORE corresponds to BE.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Number of nodes
	NodeCount *uint64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// Cluster high availability type after scaled out: 0 indicates non-high availability, 1 indicates read high availability, and 2 indicates read-write high availability.
	HaType *int64 `json:"HaType,omitnil,omitempty" name:"HaType"`
}

type ScaleOutInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Role (MATER/CORE), MASTER corresponds to FE, CORE corresponds to BE.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Number of nodes
	NodeCount *uint64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// Cluster high availability type after scaled out: 0 indicates non-high availability, 1 indicates read high availability, and 2 indicates read-write high availability.
	HaType *int64 `json:"HaType,omitnil,omitempty" name:"HaType"`
}

func (r *ScaleOutInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleOutInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Type")
	delete(f, "NodeCount")
	delete(f, "HaType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScaleOutInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleOutInstanceResponseParams struct {
	// Process ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Error message
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ScaleOutInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ScaleOutInstanceResponseParams `json:"Response"`
}

func (r *ScaleOutInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleOutInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleUpInstanceRequestParams struct {
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Node specifications
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// Role (MATER/CORE). MASTER corresponds to FE, and CORE corresponds to BE.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type ScaleUpInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Node specifications
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// Role (MATER/CORE). MASTER corresponds to FE, and CORE corresponds to BE.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *ScaleUpInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleUpInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SpecName")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScaleUpInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleUpInstanceResponseParams struct {
	// Process ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Error message
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ScaleUpInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ScaleUpInstanceResponseParams `json:"Response"`
}

func (r *ScaleUpInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleUpInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchTags struct {
	// Tag key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`

	// 1 means only the tag key is entered without a value, and 0 means both the key and the value are entered.
	AllValue *int64 `json:"AllValue,omitnil,omitempty" name:"AllValue"`
}

type SlowQueryRecord struct {
	// User query 
	OsUser *string `json:"OsUser,omitnil,omitempty" name:"OsUser"`

	// ID query
	InitialQueryId *string `json:"InitialQueryId,omitnil,omitempty" name:"InitialQueryId"`

	// SQL statement
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// Start time
	QueryStartTime *string `json:"QueryStartTime,omitnil,omitempty" name:"QueryStartTime"`

	// Execution duration
	DurationMs *int64 `json:"DurationMs,omitnil,omitempty" name:"DurationMs"`

	// The number of read rows
	ReadRows *int64 `json:"ReadRows,omitnil,omitempty" name:"ReadRows"`

	// Total number of read bytes
	ResultRows *int64 `json:"ResultRows,omitnil,omitempty" name:"ResultRows"`

	// Result bytes
	ResultBytes *uint64 `json:"ResultBytes,omitnil,omitempty" name:"ResultBytes"`

	// Memory
	MemoryUsage *int64 `json:"MemoryUsage,omitnil,omitempty" name:"MemoryUsage"`

	// Initial query IP
	InitialAddress *string `json:"InitialAddress,omitnil,omitempty" name:"InitialAddress"`

	// Database name
	// Note: This field may return null, indicating that no valid values can be obtained.
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Whether it is a query. 0 indicates no, and 1 indicates query statement.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsQuery *int64 `json:"IsQuery,omitnil,omitempty" name:"IsQuery"`

	// MB format of ResultBytes
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResultBytesMB *float64 `json:"ResultBytesMB,omitnil,omitempty" name:"ResultBytesMB"`

	// MemoryUsage, in MB
	// Note: This field may return null, indicating that no valid values can be obtained.
	MemoryUsageMB *float64 `json:"MemoryUsageMB,omitnil,omitempty" name:"MemoryUsageMB"`

	// DurationMs, in seconds
	// Note: This field may return null, indicating that no valid values can be obtained.
	DurationSec *float64 `json:"DurationSec,omitnil,omitempty" name:"DurationSec"`
}

type TablePermissions struct {
	// Full name of the table
	// Note: This field may return null, indicating that no valid values can be obtained.
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// Table permission
	// Note: This field may return null, indicating that no valid values can be obtained.
	Permissions []*string `json:"Permissions,omitnil,omitempty" name:"Permissions"`
}

type Tag struct {
	// Tag key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

// Predefined struct for user
type UpdateDatabaseRequestParams struct {
	// The database name to be modified
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Modify the operation type, such as SET_QUOTA, RENAME, SET_REPLICA_QUOTA, and SET_PROPERTIES. Modify the operation type, such as SET_QUOTA, RENAME, SET_REPLICA_QUOTA, and SET_PROPERTIES.
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// Quota value, which is used to set the quota of data volume or replicas.
	Quota *string `json:"Quota,omitnil,omitempty" name:"Quota"`

	// New database name, used for renaming operation.
	NewDbName *string `json:"NewDbName,omitnil,omitempty" name:"NewDbName"`

	// Attribute key-value pair to be set
	Properties []*PropertiesMap `json:"Properties,omitnil,omitempty" name:"Properties"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`
}

type UpdateDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// The database name to be modified
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Modify the operation type, such as SET_QUOTA, RENAME, SET_REPLICA_QUOTA, and SET_PROPERTIES. Modify the operation type, such as SET_QUOTA, RENAME, SET_REPLICA_QUOTA, and SET_PROPERTIES.
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// Quota value, which is used to set the quota of data volume or replicas.
	Quota *string `json:"Quota,omitnil,omitempty" name:"Quota"`

	// New database name, used for renaming operation.
	NewDbName *string `json:"NewDbName,omitnil,omitempty" name:"NewDbName"`

	// Attribute key-value pair to be set
	Properties []*PropertiesMap `json:"Properties,omitnil,omitempty" name:"Properties"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`
}

func (r *UpdateDatabaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DbName")
	delete(f, "Operation")
	delete(f, "Quota")
	delete(f, "NewDbName")
	delete(f, "Properties")
	delete(f, "UserName")
	delete(f, "PassWord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDatabaseResponseParams struct {
	// Whether the operation is successful
	Success *bool `json:"Success,omitnil,omitempty" name:"Success"`

	// Message description of the operation result
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateDatabaseResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDatabaseResponseParams `json:"Response"`
}

func (r *UpdateDatabaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateTableSchemaRequestParams struct {
	// Resource ID, which is the TCHouse-D resource ID used for table creation.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`

	// Database name
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Table name
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// Column
	Columns []*Column `json:"Columns,omitnil,omitempty" name:"Columns"`

	// Index information. The inverted index and N-Gram index can be configured through this parameter. The Prefix index is related to the sort key and key column, and do not require additional configuration. Configure bloom_filter_columns in the table attribute when BloomFilter index is required.
	IndexInfos []*IndexInfo `json:"IndexInfos,omitnil,omitempty" name:"IndexInfos"`

	// Bucket information
	Distribution *Distribution `json:"Distribution,omitnil,omitempty" name:"Distribution"`

	// Table description
	TableComment *string `json:"TableComment,omitnil,omitempty" name:"TableComment"`

	// Table attribute
	Properties []*Property `json:"Properties,omitnil,omitempty" name:"Properties"`
}

type UpdateTableSchemaRequest struct {
	*tchttp.BaseRequest
	
	// Resource ID, which is the TCHouse-D resource ID used for table creation.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Use the user who has corresponding permissions for operations. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Password corresponding to the user. If the TCHouse-D cluster uses a kernel account registered by a CAM user, you do not need to fill it in.
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`

	// Database name
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Table name
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// Column
	Columns []*Column `json:"Columns,omitnil,omitempty" name:"Columns"`

	// Index information. The inverted index and N-Gram index can be configured through this parameter. The Prefix index is related to the sort key and key column, and do not require additional configuration. Configure bloom_filter_columns in the table attribute when BloomFilter index is required.
	IndexInfos []*IndexInfo `json:"IndexInfos,omitnil,omitempty" name:"IndexInfos"`

	// Bucket information
	Distribution *Distribution `json:"Distribution,omitnil,omitempty" name:"Distribution"`

	// Table description
	TableComment *string `json:"TableComment,omitnil,omitempty" name:"TableComment"`

	// Table attribute
	Properties []*Property `json:"Properties,omitnil,omitempty" name:"Properties"`
}

func (r *UpdateTableSchemaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateTableSchemaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "PassWord")
	delete(f, "DbName")
	delete(f, "TableName")
	delete(f, "Columns")
	delete(f, "IndexInfos")
	delete(f, "Distribution")
	delete(f, "TableComment")
	delete(f, "Properties")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateTableSchemaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateTableSchemaResponseParams struct {
	// Error message
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateTableSchemaResponse struct {
	*tchttp.BaseResponse
	Response *UpdateTableSchemaResponseParams `json:"Response"`
}

func (r *UpdateTableSchemaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateTableSchemaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateUserPrivileges struct {
	// Whether to set catalog permission
	IsSetGlobalCatalog *bool `json:"IsSetGlobalCatalog,omitnil,omitempty" name:"IsSetGlobalCatalog"`
}

type UserWorkloadGroup struct {
	// test
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// normal
	// Note: This field may return null, indicating that no valid values can be obtained.
	WorkloadGroupName *string `json:"WorkloadGroupName,omitnil,omitempty" name:"WorkloadGroupName"`
}

type WorkloadGroupConfig struct {
	// Resource group name
	// Note: This field may return null, indicating that no valid values can be obtained.
	WorkloadGroupName *string `json:"WorkloadGroupName,omitnil,omitempty" name:"WorkloadGroupName"`

	// CPU weight
	// Note: This field may return null, indicating that no valid values can be obtained.
	CpuShare *int64 `json:"CpuShare,omitnil,omitempty" name:"CpuShare"`

	// Memory limit. The sum of memory limit values of all resource groups should be less than or equal to 100.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MemoryLimit *int64 `json:"MemoryLimit,omitnil,omitempty" name:"MemoryLimit"`

	// Whether to allow over-allocation
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnableMemoryOverCommit *bool `json:"EnableMemoryOverCommit,omitnil,omitempty" name:"EnableMemoryOverCommit"`

	// CPU hard limit
	// Note: This field may return null, indicating that no valid values can be obtained.
	CpuHardLimit *string `json:"CpuHardLimit,omitnil,omitempty" name:"CpuHardLimit"`
}

type ZoneInfo struct {
	// Availability zone name, such as ap-guangzhou-1
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Availability zone description, such as Guangzhou region 1
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// Unique tag of the availability zone
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Encryptid
	// Note: This field may return null, indicating that no valid values can be obtained.
	Encrypt *int64 `json:"Encrypt,omitnil,omitempty" name:"Encrypt"`
}