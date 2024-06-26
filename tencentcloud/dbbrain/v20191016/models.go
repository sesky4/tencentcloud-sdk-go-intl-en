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

package v20191016

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

// Predefined struct for user
type AddUserContactRequestParams struct {
	// Contact name, which needs to be unique and can contain 2-60 characters, supporting uppercase and lowercase letters, numbers, and underline “_”. It cannot start with “_”.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Email address, which can contain uppercase and lowercase letters, numbers, and underline “_”, and cannot start with “_”.
	ContactInfo *string `json:"ContactInfo,omitnil" name:"ContactInfo"`

	// Service type, which is fixed to “mysql”.
	Product *string `json:"Product,omitnil" name:"Product"`
}

type AddUserContactRequest struct {
	*tchttp.BaseRequest
	
	// Contact name, which needs to be unique and can contain 2-60 characters, supporting uppercase and lowercase letters, numbers, and underline “_”. It cannot start with “_”.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Email address, which can contain uppercase and lowercase letters, numbers, and underline “_”, and cannot start with “_”.
	ContactInfo *string `json:"ContactInfo,omitnil" name:"ContactInfo"`

	// Service type, which is fixed to “mysql”.
	Product *string `json:"Product,omitnil" name:"Product"`
}

func (r *AddUserContactRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUserContactRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "ContactInfo")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddUserContactRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddUserContactResponseParams struct {
	// The successfully added contact ID
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddUserContactResponse struct {
	*tchttp.BaseResponse
	Response *AddUserContactResponseParams `json:"Response"`
}

func (r *AddUserContactResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUserContactResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ContactItem struct {
	// Contact ID.
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// Contact name.
	Name *string `json:"Name,omitnil" name:"Name"`

	// The email address of the contact.
	Mail *string `json:"Mail,omitnil" name:"Mail"`
}

// Predefined struct for user
type CreateDBDiagReportTaskRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Start time, such as `2020-11-08T14:00:00+08:00`.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End time, such as `2020-11-09T14:00:00+08:00`.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Whether to send an email. Valid values: 0 - Yes, 1 - No.
	SendMailFlag *int64 `json:"SendMailFlag,omitnil" name:"SendMailFlag"`

	// An array of contact IDs to receive the email.
	ContactPerson []*int64 `json:"ContactPerson,omitnil" name:"ContactPerson"`

	// An array of contact group IDs to receive the email.
	ContactGroup []*int64 `json:"ContactGroup,omitnil" name:"ContactGroup"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TencentDB for CynosDB (compatible with MySQL)). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`
}

type CreateDBDiagReportTaskRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Start time, such as `2020-11-08T14:00:00+08:00`.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End time, such as `2020-11-09T14:00:00+08:00`.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Whether to send an email. Valid values: 0 - Yes, 1 - No.
	SendMailFlag *int64 `json:"SendMailFlag,omitnil" name:"SendMailFlag"`

	// An array of contact IDs to receive the email.
	ContactPerson []*int64 `json:"ContactPerson,omitnil" name:"ContactPerson"`

	// An array of contact group IDs to receive the email.
	ContactGroup []*int64 `json:"ContactGroup,omitnil" name:"ContactGroup"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TencentDB for CynosDB (compatible with MySQL)). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`
}

func (r *CreateDBDiagReportTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBDiagReportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SendMailFlag")
	delete(f, "ContactPerson")
	delete(f, "ContactGroup")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBDiagReportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBDiagReportTaskResponseParams struct {
	// ID of an async task request, which can be used to query the execution result of an async task.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil" name:"AsyncRequestId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDBDiagReportTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBDiagReportTaskResponseParams `json:"Response"`
}

func (r *CreateDBDiagReportTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBDiagReportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBDiagReportUrlRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The health report task ID, which can be queried through `DescribeDBDiagReportTasks`.
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil" name:"AsyncRequestId"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`
}

type CreateDBDiagReportUrlRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The health report task ID, which can be queried through `DescribeDBDiagReportTasks`.
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil" name:"AsyncRequestId"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`
}

func (r *CreateDBDiagReportUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBDiagReportUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AsyncRequestId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBDiagReportUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBDiagReportUrlResponseParams struct {
	// The URL of the health report.
	ReportUrl *string `json:"ReportUrl,omitnil" name:"ReportUrl"`

	// The expiration timestamp of the health report URL (in seconds).
	ExpireTime *int64 `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDBDiagReportUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBDiagReportUrlResponseParams `json:"Response"`
}

func (r *CreateDBDiagReportUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBDiagReportUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMailProfileRequestParams struct {
	// Email configurations
	ProfileInfo *ProfileInfo `json:"ProfileInfo,omitnil" name:"ProfileInfo"`

	// Configuration level. Valid values: "User" (user-level), "Instance" (instance-level). For database inspection report, it should be `User`; and for scheduled task reports, it should be `Instance`.
	ProfileLevel *string `json:"ProfileLevel,omitnil" name:"ProfileLevel"`

	// Configuration name, which needs to be unique. For database inspection reports, this name can be customize as needed. For scheduled task reports, the name should be in the format of "scheduler_" + {instanceId}, such as "schduler_cdb-test".
	ProfileName *string `json:"ProfileName,omitnil" name:"ProfileName"`

	// Configuration type. Valid values: "dbScan_mail_configuration" (email configuration of database inspection report), "scheduler_mail_configuration" (email configuration of scheduled task report).
	ProfileType *string `json:"ProfileType,omitnil" name:"ProfileType"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TencentDB for CynosDB (compatible with MySQL)).
	Product *string `json:"Product,omitnil" name:"Product"`

	// Instance ID bound with the configuration, which is set when the configuration level is "Instance". Only one instance can be bound at a time. When the configuration level is “User”, leave this parameter empty.
	BindInstanceIds []*string `json:"BindInstanceIds,omitnil" name:"BindInstanceIds"`
}

type CreateMailProfileRequest struct {
	*tchttp.BaseRequest
	
	// Email configurations
	ProfileInfo *ProfileInfo `json:"ProfileInfo,omitnil" name:"ProfileInfo"`

	// Configuration level. Valid values: "User" (user-level), "Instance" (instance-level). For database inspection report, it should be `User`; and for scheduled task reports, it should be `Instance`.
	ProfileLevel *string `json:"ProfileLevel,omitnil" name:"ProfileLevel"`

	// Configuration name, which needs to be unique. For database inspection reports, this name can be customize as needed. For scheduled task reports, the name should be in the format of "scheduler_" + {instanceId}, such as "schduler_cdb-test".
	ProfileName *string `json:"ProfileName,omitnil" name:"ProfileName"`

	// Configuration type. Valid values: "dbScan_mail_configuration" (email configuration of database inspection report), "scheduler_mail_configuration" (email configuration of scheduled task report).
	ProfileType *string `json:"ProfileType,omitnil" name:"ProfileType"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TencentDB for CynosDB (compatible with MySQL)).
	Product *string `json:"Product,omitnil" name:"Product"`

	// Instance ID bound with the configuration, which is set when the configuration level is "Instance". Only one instance can be bound at a time. When the configuration level is “User”, leave this parameter empty.
	BindInstanceIds []*string `json:"BindInstanceIds,omitnil" name:"BindInstanceIds"`
}

func (r *CreateMailProfileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMailProfileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProfileInfo")
	delete(f, "ProfileLevel")
	delete(f, "ProfileName")
	delete(f, "ProfileType")
	delete(f, "Product")
	delete(f, "BindInstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMailProfileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMailProfileResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateMailProfileResponse struct {
	*tchttp.BaseResponse
	Response *CreateMailProfileResponseParams `json:"Response"`
}

func (r *CreateMailProfileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMailProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSchedulerMailProfileRequestParams struct {
	// Value range: 1-7, representing Monday to Sunday respectively.
	WeekConfiguration []*int64 `json:"WeekConfiguration,omitnil" name:"WeekConfiguration"`

	// Email configurations
	ProfileInfo *ProfileInfo `json:"ProfileInfo,omitnil" name:"ProfileInfo"`

	// Configuration name, which needs to be unique. For scheduled task reports, the name should be in the format of "scheduler_" + {instanceId}, such as "schduler_cdb-test".
	ProfileName *string `json:"ProfileName,omitnil" name:"ProfileName"`

	// Configure the instance ID that you need to generate the health report.
	BindInstanceId *string `json:"BindInstanceId,omitnil" name:"BindInstanceId"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`
}

type CreateSchedulerMailProfileRequest struct {
	*tchttp.BaseRequest
	
	// Value range: 1-7, representing Monday to Sunday respectively.
	WeekConfiguration []*int64 `json:"WeekConfiguration,omitnil" name:"WeekConfiguration"`

	// Email configurations
	ProfileInfo *ProfileInfo `json:"ProfileInfo,omitnil" name:"ProfileInfo"`

	// Configuration name, which needs to be unique. For scheduled task reports, the name should be in the format of "scheduler_" + {instanceId}, such as "schduler_cdb-test".
	ProfileName *string `json:"ProfileName,omitnil" name:"ProfileName"`

	// Configure the instance ID that you need to generate the health report.
	BindInstanceId *string `json:"BindInstanceId,omitnil" name:"BindInstanceId"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`
}

func (r *CreateSchedulerMailProfileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSchedulerMailProfileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WeekConfiguration")
	delete(f, "ProfileInfo")
	delete(f, "ProfileName")
	delete(f, "BindInstanceId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSchedulerMailProfileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSchedulerMailProfileResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateSchedulerMailProfileResponse struct {
	*tchttp.BaseResponse
	Response *CreateSchedulerMailProfileResponseParams `json:"Response"`
}

func (r *CreateSchedulerMailProfileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSchedulerMailProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllUserContactRequestParams struct {
	// Service type, which is fixed to “mysql”.
	Product *string `json:"Product,omitnil" name:"Product"`

	// An array of contact name. Fuzzy search is supported.
	Names []*string `json:"Names,omitnil" name:"Names"`
}

type DescribeAllUserContactRequest struct {
	*tchttp.BaseRequest
	
	// Service type, which is fixed to “mysql”.
	Product *string `json:"Product,omitnil" name:"Product"`

	// An array of contact name. Fuzzy search is supported.
	Names []*string `json:"Names,omitnil" name:"Names"`
}

func (r *DescribeAllUserContactRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllUserContactRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "Names")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllUserContactRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllUserContactResponseParams struct {
	// Total number of contacts.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Contact information.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Contacts []*ContactItem `json:"Contacts,omitnil" name:"Contacts"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAllUserContactResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllUserContactResponseParams `json:"Response"`
}

func (r *DescribeAllUserContactResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllUserContactResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllUserGroupRequestParams struct {
	// Service type, which is fixed to “mysql”.
	Product *string `json:"Product,omitnil" name:"Product"`

	// An array of contact group name. Fuzzy search is supported.
	Names []*string `json:"Names,omitnil" name:"Names"`
}

type DescribeAllUserGroupRequest struct {
	*tchttp.BaseRequest
	
	// Service type, which is fixed to “mysql”.
	Product *string `json:"Product,omitnil" name:"Product"`

	// An array of contact group name. Fuzzy search is supported.
	Names []*string `json:"Names,omitnil" name:"Names"`
}

func (r *DescribeAllUserGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllUserGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "Names")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllUserGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllUserGroupResponseParams struct {
	// Total number of contact groups.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Contact group information.
	// Note: this field may return `null`, indicating that no valid value is obtained.
	Groups []*GroupItem `json:"Groups,omitnil" name:"Groups"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAllUserGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllUserGroupResponseParams `json:"Response"`
}

func (r *DescribeAllUserGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagEventRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Event ID, which can be obtained through the `DescribeDBDiagHistory` API.
	EventId *int64 `json:"EventId,omitnil" name:"EventId"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TencentDB for CynosDB (compatible with MySQL)). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`
}

type DescribeDBDiagEventRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Event ID, which can be obtained through the `DescribeDBDiagHistory` API.
	EventId *int64 `json:"EventId,omitnil" name:"EventId"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TencentDB for CynosDB (compatible with MySQL)). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`
}

func (r *DescribeDBDiagEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "EventId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBDiagEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagEventResponseParams struct {
	// Diagnosis item.
	DiagItem *string `json:"DiagItem,omitnil" name:"DiagItem"`

	// Diagnosis type.
	DiagType *string `json:"DiagType,omitnil" name:"DiagType"`

	// Event ID.
	EventId *int64 `json:"EventId,omitnil" name:"EventId"`

	// Event details.
	Explanation *string `json:"Explanation,omitnil" name:"Explanation"`

	// Summary.
	Outline *string `json:"Outline,omitnil" name:"Outline"`

	// Problem found.
	Problem *string `json:"Problem,omitnil" name:"Problem"`

	// Severity, which can be divided into 5 levels: 1: fatal, 2: severe, 3: warning, 4: notice, 5: healthy.
	Severity *int64 `json:"Severity,omitnil" name:"Severity"`

	// Start time
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// Suggestion.
	Suggestions *string `json:"Suggestions,omitnil" name:"Suggestions"`

	// Reserved field.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// End time.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDBDiagEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBDiagEventResponseParams `json:"Response"`
}

func (r *DescribeDBDiagEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagHistoryRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Start time, such as "2019-09-10 12:13:14".
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End time, such as "2019-09-11 12:13:14". The interval between the end time and the start time can be up to 2 days.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TencentDB for CynosDB (compatible with MySQL)). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`
}

type DescribeDBDiagHistoryRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Start time, such as "2019-09-10 12:13:14".
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End time, such as "2019-09-11 12:13:14". The interval between the end time and the start time can be up to 2 days.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TencentDB for CynosDB (compatible with MySQL)). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`
}

func (r *DescribeDBDiagHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBDiagHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagHistoryResponseParams struct {
	// Event description.
	Events []*DiagHistoryEventItem `json:"Events,omitnil" name:"Events"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDBDiagHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBDiagHistoryResponseParams `json:"Response"`
}

func (r *DescribeDBDiagHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagReportTasksRequestParams struct {
	// Start time of the first task in the format of yyyy-MM-dd HH:mm:ss, such as 2019-09-10 12:13:14. It is used for queries by time range.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End time of the last task in the format of yyyy-MM-dd HH:mm:ss, such as 2019-09-10 12:13:14. It is used for queries by time range.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Instance ID array, which is used to filter the task list of a specified instance.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Source that triggers the task. Valid values: `DAILY_INSPECTION` (instance inspection), `SCHEDULED` (timed generation), and `MANUAL` (manual trigger).
	Sources []*string `json:"Sources,omitnil" name:"Sources"`

	// Health level. Valid values: `HEALTH` (healthy), `SUB_HEALTH` (suboptimal), `RISK` (risky), and `HIGH_RISK` (critical).
	HealthLevels *string `json:"HealthLevels,omitnil" name:"HealthLevels"`

	// The task status. Valid values: `created` (create), `chosen` (to be executed), `running` (being executed), `failed` (failed), and `finished` (completed).
	TaskStatuses *string `json:"TaskStatuses,omitnil" name:"TaskStatuses"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`
}

type DescribeDBDiagReportTasksRequest struct {
	*tchttp.BaseRequest
	
	// Start time of the first task in the format of yyyy-MM-dd HH:mm:ss, such as 2019-09-10 12:13:14. It is used for queries by time range.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End time of the last task in the format of yyyy-MM-dd HH:mm:ss, such as 2019-09-10 12:13:14. It is used for queries by time range.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Instance ID array, which is used to filter the task list of a specified instance.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Source that triggers the task. Valid values: `DAILY_INSPECTION` (instance inspection), `SCHEDULED` (timed generation), and `MANUAL` (manual trigger).
	Sources []*string `json:"Sources,omitnil" name:"Sources"`

	// Health level. Valid values: `HEALTH` (healthy), `SUB_HEALTH` (suboptimal), `RISK` (risky), and `HIGH_RISK` (critical).
	HealthLevels *string `json:"HealthLevels,omitnil" name:"HealthLevels"`

	// The task status. Valid values: `created` (create), `chosen` (to be executed), `running` (being executed), `failed` (failed), and `finished` (completed).
	TaskStatuses *string `json:"TaskStatuses,omitnil" name:"TaskStatuses"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of returned results. Default value: 20.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`
}

func (r *DescribeDBDiagReportTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagReportTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "InstanceIds")
	delete(f, "Sources")
	delete(f, "HealthLevels")
	delete(f, "TaskStatuses")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBDiagReportTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagReportTasksResponseParams struct {
	// Total number of tasks.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Task list.
	Tasks []*HealthReportTask `json:"Tasks,omitnil" name:"Tasks"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDBDiagReportTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBDiagReportTasksResponseParams `json:"Response"`
}

func (r *DescribeDBDiagReportTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagReportTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSpaceStatusRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Query period in days. The end date is the current date and the query period is 7 days by default.
	RangeDays *int64 `json:"RangeDays,omitnil" name:"RangeDays"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TencentDB for CynosDB (compatible with MySQL)). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`
}

type DescribeDBSpaceStatusRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Query period in days. The end date is the current date and the query period is 7 days by default.
	RangeDays *int64 `json:"RangeDays,omitnil" name:"RangeDays"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TencentDB for CynosDB (compatible with MySQL)). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`
}

func (r *DescribeDBSpaceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSpaceStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RangeDays")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSpaceStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSpaceStatusResponseParams struct {
	// Disk usage growth in MB.
	Growth *int64 `json:"Growth,omitnil" name:"Growth"`

	// Available disk space in MB.
	Remain *int64 `json:"Remain,omitnil" name:"Remain"`

	// Total disk space in MB.
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// Estimated number of available days.
	AvailableDays *int64 `json:"AvailableDays,omitnil" name:"AvailableDays"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDBSpaceStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSpaceStatusResponseParams `json:"Response"`
}

func (r *DescribeDBSpaceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSpaceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiagDBInstancesRequestParams struct {
	// Whether it is an instance supported by DBbrain. It is fixed to “true”.
	IsSupported *bool `json:"IsSupported,omitnil" name:"IsSupported"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`

	// Pagination parameter indicating the offset.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Pagination parameter indicating the number of entries for each page.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Query by instance name.
	InstanceNames []*string `json:"InstanceNames,omitnil" name:"InstanceNames"`

	// Query by instance ID.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Query by region.
	Regions []*string `json:"Regions,omitnil" name:"Regions"`
}

type DescribeDiagDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Whether it is an instance supported by DBbrain. It is fixed to “true”.
	IsSupported *bool `json:"IsSupported,omitnil" name:"IsSupported"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`

	// Pagination parameter indicating the offset.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Pagination parameter indicating the number of entries for each page.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Query by instance name.
	InstanceNames []*string `json:"InstanceNames,omitnil" name:"InstanceNames"`

	// Query by instance ID.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// Query by region.
	Regions []*string `json:"Regions,omitnil" name:"Regions"`
}

func (r *DescribeDiagDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiagDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IsSupported")
	delete(f, "Product")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstanceNames")
	delete(f, "InstanceIds")
	delete(f, "Regions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDiagDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiagDBInstancesResponseParams struct {
	// Total Number of Instances
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Status of all instance inspection. 0: all instance inspection enabled, 1: all instance inspection disabled
	DbScanStatus *int64 `json:"DbScanStatus,omitnil" name:"DbScanStatus"`

	// Instance related information
	Items []*InstanceInfo `json:"Items,omitnil" name:"Items"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDiagDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDiagDBInstancesResponseParams `json:"Response"`
}

func (r *DescribeDiagDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiagDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHealthScoreRequestParams struct {
	// The instance ID that needs to obtain the health score
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Time to obtain the health score
	Time *string `json:"Time,omitnil" name:"Time"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL (compatible with MySQL)). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`
}

type DescribeHealthScoreRequest struct {
	*tchttp.BaseRequest
	
	// The instance ID that needs to obtain the health score
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Time to obtain the health score
	Time *string `json:"Time,omitnil" name:"Time"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL (compatible with MySQL)). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`
}

func (r *DescribeHealthScoreRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHealthScoreRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Time")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHealthScoreRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHealthScoreResponseParams struct {
	// Health score and deduction for exceptions
	Data *HealthScoreInfo `json:"Data,omitnil" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeHealthScoreResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHealthScoreResponseParams `json:"Response"`
}

func (r *DescribeHealthScoreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHealthScoreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMailProfileRequestParams struct {
	// Configuration type. Valid values: "dbScan_mail_configuration" (email configuration of database inspection report), "scheduler_mail_configuration" (email configuration of scheduled task report).
	ProfileType *string `json:"ProfileType,omitnil" name:"ProfileType"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`

	// Pagination offset
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// The number of results per page in paginated queries. Maximum value: 50
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Query by the name of email configuration. The name of the regularly sent email configuration should be in the format of "scheduler_"+{instanceId}.
	ProfileName *string `json:"ProfileName,omitnil" name:"ProfileName"`
}

type DescribeMailProfileRequest struct {
	*tchttp.BaseRequest
	
	// Configuration type. Valid values: "dbScan_mail_configuration" (email configuration of database inspection report), "scheduler_mail_configuration" (email configuration of scheduled task report).
	ProfileType *string `json:"ProfileType,omitnil" name:"ProfileType"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`

	// Pagination offset
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// The number of results per page in paginated queries. Maximum value: 50
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Query by the name of email configuration. The name of the regularly sent email configuration should be in the format of "scheduler_"+{instanceId}.
	ProfileName *string `json:"ProfileName,omitnil" name:"ProfileName"`
}

func (r *DescribeMailProfileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMailProfileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProfileType")
	delete(f, "Product")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ProfileName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMailProfileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMailProfileResponseParams struct {
	// Email configuration details
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ProfileList []*UserProfile `json:"ProfileList,omitnil" name:"ProfileList"`

	// Total number of email templates
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMailProfileResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMailProfileResponseParams `json:"Response"`
}

func (r *DescribeMailProfileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMailProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogTimeSeriesStatsRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Start time, such as "2019-09-10 12:13:14".
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End time, such as "2019-09-10 12:13:14". The interval between the end time and the start time can be up to 7 days.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TencentDB for CynosDB (compatible with MySQL)). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`
}

type DescribeSlowLogTimeSeriesStatsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Start time, such as "2019-09-10 12:13:14".
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End time, such as "2019-09-10 12:13:14". The interval between the end time and the start time can be up to 7 days.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TencentDB for CynosDB (compatible with MySQL)). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`
}

func (r *DescribeSlowLogTimeSeriesStatsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogTimeSeriesStatsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogTimeSeriesStatsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogTimeSeriesStatsResponseParams struct {
	// Time range in seconds in histogram.
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// Number of slow logs in specified time range.
	TimeSeries []*TimeSlice `json:"TimeSeries,omitnil" name:"TimeSeries"`

	// Instance CPU utilization monitoring data in specified time range.
	SeriesData *MonitorMetricSeriesData `json:"SeriesData,omitnil" name:"SeriesData"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSlowLogTimeSeriesStatsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogTimeSeriesStatsResponseParams `json:"Response"`
}

func (r *DescribeSlowLogTimeSeriesStatsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogTimeSeriesStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogTopSqlsRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Start time, such as "2019-09-10 12:13:14".
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End time, such as "2019-09-10 12:13:14". The interval between the end time and the start time can be up to 7 days.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Sorting key. Valid values: QueryTime, ExecTimes, RowsSent, LockTime, RowsExamined.
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// Sorting order. Valid values: ASC (ascending), DESC (descending).
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Database name array.
	SchemaList []*SchemaItem `json:"SchemaList,omitnil" name:"SchemaList"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TencentDB for CynosDB (compatible with MySQL)). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`
}

type DescribeSlowLogTopSqlsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Start time, such as "2019-09-10 12:13:14".
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End time, such as "2019-09-10 12:13:14". The interval between the end time and the start time can be up to 7 days.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Sorting key. Valid values: QueryTime, ExecTimes, RowsSent, LockTime, RowsExamined.
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// Sorting order. Valid values: ASC (ascending), DESC (descending).
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// Number of results to be returned. Default value: 20. Maximum value: 100.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Offset. Default value: 0.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Database name array.
	SchemaList []*SchemaItem `json:"SchemaList,omitnil" name:"SchemaList"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TencentDB for CynosDB (compatible with MySQL)). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`
}

func (r *DescribeSlowLogTopSqlsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogTopSqlsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SortBy")
	delete(f, "OrderBy")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SchemaList")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogTopSqlsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogTopSqlsResponseParams struct {
	// Number of eligible entries.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// List of top slow SQL statements
	Rows []*SlowLogTopSqlItem `json:"Rows,omitnil" name:"Rows"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSlowLogTopSqlsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogTopSqlsResponseParams `json:"Response"`
}

func (r *DescribeSlowLogTopSqlsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogTopSqlsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogUserHostStatsRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Start time of the time range in the format of yyyy-MM-dd HH:mm:ss, such as 2019-09-10 12:13:14.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End time of the time range in the format of yyyy-MM-dd HH:mm:ss, such as 2019-09-10 12:13:14.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`
}

type DescribeSlowLogUserHostStatsRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Start time of the time range in the format of yyyy-MM-dd HH:mm:ss, such as 2019-09-10 12:13:14.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End time of the time range in the format of yyyy-MM-dd HH:mm:ss, such as 2019-09-10 12:13:14.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TDSQL-C for MySQL). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`
}

func (r *DescribeSlowLogUserHostStatsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogUserHostStatsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogUserHostStatsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogUserHostStatsResponseParams struct {
	// Total number of source addresses.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Detailed list of the proportion of slow logs from each source address.
	Items []*SlowLogHost `json:"Items,omitnil" name:"Items"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSlowLogUserHostStatsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogUserHostStatsResponseParams `json:"Response"`
}

func (r *DescribeSlowLogUserHostStatsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogUserHostStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceSchemaTimeSeriesRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Number of returned top databases. Maximum value: 100. Default value: 20.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Field used to sort top tables. Valid values: `DataLength`, `IndexLength`, `TotalLength`, `DataFree`, `FragRatio`, `TableRows`, and `PhysicalFileSize` (supported only by TencentDB for MySQL instances). For TencentDB for MySQL instances, the default value is `PhysicalFileSize`; for other database instances, the default value is `TotalLength`.
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// Start date. It can be as early as 29 days before the current date, and defaults to 6 days before the end date.
	StartDate *string `json:"StartDate,omitnil" name:"StartDate"`

	// End date. It can be as early as 29 days before the current date, and defaults to the current date.
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TencentDB for CynosDB (compatible with MySQL)). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`
}

type DescribeTopSpaceSchemaTimeSeriesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Number of returned top databases. Maximum value: 100. Default value: 20.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Field used to sort top tables. Valid values: `DataLength`, `IndexLength`, `TotalLength`, `DataFree`, `FragRatio`, `TableRows`, and `PhysicalFileSize` (supported only by TencentDB for MySQL instances). For TencentDB for MySQL instances, the default value is `PhysicalFileSize`; for other database instances, the default value is `TotalLength`.
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// Start date. It can be as early as 29 days before the current date, and defaults to 6 days before the end date.
	StartDate *string `json:"StartDate,omitnil" name:"StartDate"`

	// End date. It can be as early as 29 days before the current date, and defaults to the current date.
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TencentDB for CynosDB (compatible with MySQL)). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`
}

func (r *DescribeTopSpaceSchemaTimeSeriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceSchemaTimeSeriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "SortBy")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopSpaceSchemaTimeSeriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceSchemaTimeSeriesResponseParams struct {
	// Time series list of the returned space statistics of top databases.
	TopSpaceSchemaTimeSeries []*SchemaSpaceTimeSeries `json:"TopSpaceSchemaTimeSeries,omitnil" name:"TopSpaceSchemaTimeSeries"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTopSpaceSchemaTimeSeriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopSpaceSchemaTimeSeriesResponseParams `json:"Response"`
}

func (r *DescribeTopSpaceSchemaTimeSeriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceSchemaTimeSeriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceSchemasRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Number of returned top databases. Maximum value: 100. Default value: 20.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Field used to sort top tables. Valid values: `DataLength`, `IndexLength`, `TotalLength`, `DataFree`, `FragRatio`, `TableRows`, and `PhysicalFileSize` (supported only by TencentDB for MySQL instances). For TencentDB for MySQL instances, the default value is `PhysicalFileSize`; for other database instances, the default value is `TotalLength`.
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TencentDB for CynosDB (compatible with MySQL)). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`
}

type DescribeTopSpaceSchemasRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Number of returned top databases. Maximum value: 100. Default value: 20.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Field used to sort top tables. Valid values: `DataLength`, `IndexLength`, `TotalLength`, `DataFree`, `FragRatio`, `TableRows`, and `PhysicalFileSize` (supported only by TencentDB for MySQL instances). For TencentDB for MySQL instances, the default value is `PhysicalFileSize`; for other database instances, the default value is `TotalLength`.
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TencentDB for CynosDB (compatible with MySQL)). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`
}

func (r *DescribeTopSpaceSchemasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceSchemasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "SortBy")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopSpaceSchemasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceSchemasResponseParams struct {
	// List of the returned space statistics of top databases.
	TopSpaceSchemas []*SchemaSpaceData `json:"TopSpaceSchemas,omitnil" name:"TopSpaceSchemas"`

	// Timestamp (in seconds) of database space data collect points
	Timestamp *int64 `json:"Timestamp,omitnil" name:"Timestamp"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTopSpaceSchemasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopSpaceSchemasResponseParams `json:"Response"`
}

func (r *DescribeTopSpaceSchemasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceSchemasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceTableTimeSeriesRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Number of returned top tables. Maximum value: 100. Default value: 20.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Field used to sort top tables. Valid values: DataLength, IndexLength, TotalLength, DataFree, FragRatio, TableRows, PhysicalFileSize. Default value: PhysicalFileSize.
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// Start date. It can be as early as 29 days before the current date, and defaults to 6 days before the end date.
	StartDate *string `json:"StartDate,omitnil" name:"StartDate"`

	// End date. It can be as early as 29 days before the current date, and defaults to the current date.
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TencentDB for CynosDB (compatible with MySQL)). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`
}

type DescribeTopSpaceTableTimeSeriesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Number of returned top tables. Maximum value: 100. Default value: 20.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Field used to sort top tables. Valid values: DataLength, IndexLength, TotalLength, DataFree, FragRatio, TableRows, PhysicalFileSize. Default value: PhysicalFileSize.
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// Start date. It can be as early as 29 days before the current date, and defaults to 6 days before the end date.
	StartDate *string `json:"StartDate,omitnil" name:"StartDate"`

	// End date. It can be as early as 29 days before the current date, and defaults to the current date.
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TencentDB for CynosDB (compatible with MySQL)). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`
}

func (r *DescribeTopSpaceTableTimeSeriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceTableTimeSeriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "SortBy")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopSpaceTableTimeSeriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceTableTimeSeriesResponseParams struct {
	// Time series list of the returned space statistics of top tables.
	TopSpaceTableTimeSeries []*TableSpaceTimeSeries `json:"TopSpaceTableTimeSeries,omitnil" name:"TopSpaceTableTimeSeries"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTopSpaceTableTimeSeriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopSpaceTableTimeSeriesResponseParams `json:"Response"`
}

func (r *DescribeTopSpaceTableTimeSeriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceTableTimeSeriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceTablesRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Number of returned top tables. Maximum value: 100. Default value: 20.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Field used to sort top tables. Valid values: `DataLength`, `IndexLength`, `TotalLength`, `DataFree`, `FragRatio`, `TableRows`, and `PhysicalFileSize` (only supported by TencentDB for MySQL instances). For TencentDB for MySQL instances, the default value is PhysicalFileSize; for other database instances, the default value is `TotalLength`.
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TencentDB for CynosDB (compatible with MySQL)). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`
}

type DescribeTopSpaceTablesRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Number of returned top tables. Maximum value: 100. Default value: 20.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Field used to sort top tables. Valid values: `DataLength`, `IndexLength`, `TotalLength`, `DataFree`, `FragRatio`, `TableRows`, and `PhysicalFileSize` (only supported by TencentDB for MySQL instances). For TencentDB for MySQL instances, the default value is PhysicalFileSize; for other database instances, the default value is `TotalLength`.
	SortBy *string `json:"SortBy,omitnil" name:"SortBy"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TencentDB for CynosDB (compatible with MySQL)). Default value: `mysql`.
	Product *string `json:"Product,omitnil" name:"Product"`
}

func (r *DescribeTopSpaceTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "SortBy")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopSpaceTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceTablesResponseParams struct {
	// List of the returned space statistics of top tables.
	TopSpaceTables []*TableSpaceData `json:"TopSpaceTables,omitnil" name:"TopSpaceTables"`

	// Timestamp (in seconds) of tablespace data collect points
	Timestamp *int64 `json:"Timestamp,omitnil" name:"Timestamp"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTopSpaceTablesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopSpaceTablesResponseParams `json:"Response"`
}

func (r *DescribeTopSpaceTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserSqlAdviceRequestParams struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// SQL statement.
	SqlText *string `json:"SqlText,omitnil" name:"SqlText"`

	// Database name.
	Schema *string `json:"Schema,omitnil" name:"Schema"`
}

type DescribeUserSqlAdviceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// SQL statement.
	SqlText *string `json:"SqlText,omitnil" name:"SqlText"`

	// Database name.
	Schema *string `json:"Schema,omitnil" name:"Schema"`
}

func (r *DescribeUserSqlAdviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserSqlAdviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SqlText")
	delete(f, "Schema")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserSqlAdviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserSqlAdviceResponseParams struct {
	// SQL statement optimization suggestions, which can be parsed into JSON arrays.
	Advices *string `json:"Advices,omitnil" name:"Advices"`

	// Notes of SQL statement optimization suggestions, which can be parsed into String arrays.
	Comments *string `json:"Comments,omitnil" name:"Comments"`

	// SQL statement.
	SqlText *string `json:"SqlText,omitnil" name:"SqlText"`

	// Database name.
	Schema *string `json:"Schema,omitnil" name:"Schema"`

	// DDL information of related tables, which can be parsed into JSON arrays.
	Tables *string `json:"Tables,omitnil" name:"Tables"`

	// SQL execution plan, which can be parsed into JSON.
	SqlPlan *string `json:"SqlPlan,omitnil" name:"SqlPlan"`

	// Cost saving details after SQL statement optimization, which can be parsed into JSON.
	Cost *string `json:"Cost,omitnil" name:"Cost"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeUserSqlAdviceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserSqlAdviceResponseParams `json:"Response"`
}

func (r *DescribeUserSqlAdviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserSqlAdviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiagHistoryEventItem struct {
	// Diagnosis type.
	DiagType *string `json:"DiagType,omitnil" name:"DiagType"`

	// End time.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Start time.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// Event ID.
	EventId *int64 `json:"EventId,omitnil" name:"EventId"`

	// Severity, which can be divided into 5 levels: 1: fatal, 2: severe, 3: warning, 4: notice, 5: healthy.
	Severity *int64 `json:"Severity,omitnil" name:"Severity"`

	// Summary.
	Outline *string `json:"Outline,omitnil" name:"Outline"`

	// Diagnosis item.
	DiagItem *string `json:"DiagItem,omitnil" name:"DiagItem"`

	// Instance ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Reserved field
	// Note: this field may return null, indicating that no valid values can be obtained.
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// Region
	// Note: this field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil" name:"Region"`
}

type EventInfo struct {
	// Event ID
	EventId *int64 `json:"EventId,omitnil" name:"EventId"`

	// Diagnosis type
	DiagType *string `json:"DiagType,omitnil" name:"DiagType"`

	// Start time
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End time
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Summary
	Outline *string `json:"Outline,omitnil" name:"Outline"`

	// Severity, which can be divided into 5 levels: 1: fatal, 2: severe, 3: warning, 4: notice, 5: healthy.
	Severity *int64 `json:"Severity,omitnil" name:"Severity"`

	// Deduction
	ScoreLost *int64 `json:"ScoreLost,omitnil" name:"ScoreLost"`

	// Reserved field
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// The number of alarms
	Count *int64 `json:"Count,omitnil" name:"Count"`
}

type GroupItem struct {
	// Group ID.
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// Group name.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Number of group members.
	MemberCount *int64 `json:"MemberCount,omitnil" name:"MemberCount"`
}

type HealthReportTask struct {
	// Async task request ID.
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil" name:"AsyncRequestId"`

	// Source that triggers the task. Valid values: `DAILY_INSPECTION` (instance inspection), `SCHEDULED` (timed generation), and `MANUAL` (manual trigger).
	Source *string `json:"Source,omitnil" name:"Source"`

	// Task progress in %.
	Progress *int64 `json:"Progress,omitnil" name:"Progress"`

	// Task creation time.
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Task start time.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// Task end time.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Basic information about the instance to which the task belongs.
	InstanceInfo *InstanceBasicInfo `json:"InstanceInfo,omitnil" name:"InstanceInfo"`

	// Health information in a health report.
	HealthStatus *HealthStatus `json:"HealthStatus,omitnil" name:"HealthStatus"`
}

type HealthScoreInfo struct {
	// Exception details
	IssueTypes []*IssueTypeInfo `json:"IssueTypes,omitnil" name:"IssueTypes"`

	// Total number of the exceptions
	EventsTotalCount *int64 `json:"EventsTotalCount,omitnil" name:"EventsTotalCount"`

	// Health score
	HealthScore *int64 `json:"HealthScore,omitnil" name:"HealthScore"`

	// Health level, such as "HEALTH", "SUB_HEALTH", "RISK", "HIGH_RISK".
	HealthLevel *string `json:"HealthLevel,omitnil" name:"HealthLevel"`
}

type HealthStatus struct {
	// Health score out of 100 points.
	HealthScore *int64 `json:"HealthScore,omitnil" name:"HealthScore"`

	// Health level. Valid values: `HEALTH` (healthy), `SUB_HEALTH` (suboptimal), `RISK` (risky), and `HIGH_RISK` (critical).
	HealthLevel *string `json:"HealthLevel,omitnil" name:"HealthLevel"`

	// Total scores deducted.
	ScoreLost *int64 `json:"ScoreLost,omitnil" name:"ScoreLost"`

	// Deduction details.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	ScoreDetails []*ScoreDetail `json:"ScoreDetails,omitnil" name:"ScoreDetails"`
}

type InstanceBasicInfo struct {
	// Instance ID.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance name.
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Private IP of the instance.
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// Private network port of the instance.
	Vport *int64 `json:"Vport,omitnil" name:"Vport"`

	// Instance product.
	Product *string `json:"Product,omitnil" name:"Product"`

	// Instance engine version.
	EngineVersion *string `json:"EngineVersion,omitnil" name:"EngineVersion"`
}

type InstanceConfs struct {
	// Whether to enable database inspection. Valid values: Yes/No.
	DailyInspection *string `json:"DailyInspection,omitnil" name:"DailyInspection"`

	// Whether to enable instance overview. Valid values: Yes/No.
	OverviewDisplay *string `json:"OverviewDisplay,omitnil" name:"OverviewDisplay"`
}

type InstanceInfo struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// The region where the instance belongs
	Region *string `json:"Region,omitnil" name:"Region"`

	// Health score
	HealthScore *int64 `json:"HealthScore,omitnil" name:"HealthScore"`

	// Service
	Product *string `json:"Product,omitnil" name:"Product"`

	// Number of exceptions
	EventCount *int64 `json:"EventCount,omitnil" name:"EventCount"`

	// Instance type. Valid values: 1: MASTER, 2: DR, 3: RO, 4: SDR
	InstanceType *int64 `json:"InstanceType,omitnil" name:"InstanceType"`

	// Number of cores
	Cpu *int64 `json:"Cpu,omitnil" name:"Cpu"`

	// Memory in MB
	Memory *int64 `json:"Memory,omitnil" name:"Memory"`

	// Disk storage in GB
	Volume *int64 `json:"Volume,omitnil" name:"Volume"`

	// Database version
	EngineVersion *string `json:"EngineVersion,omitnil" name:"EngineVersion"`

	// Private network address
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// Private network port
	Vport *int64 `json:"Vport,omitnil" name:"Vport"`

	// Access source
	Source *string `json:"Source,omitnil" name:"Source"`

	// Group ID
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// Group name
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// Instance status. Valid values: 0: Delivering, 1: Running, 4: Terminating, 5: Isolated
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// Subnet unified ID
	UniqSubnetId *string `json:"UniqSubnetId,omitnil" name:"UniqSubnetId"`

	// cdb (TencentDB instance) type
	DeployMode *string `json:"DeployMode,omitnil" name:"DeployMode"`

	// cdb (TencentDB instance) initialization flag. Valid values: 0: not initialized, 1: initialized
	InitFlag *int64 `json:"InitFlag,omitnil" name:"InitFlag"`

	// Task status
	TaskStatus *int64 `json:"TaskStatus,omitnil" name:"TaskStatus"`

	// Unified VPC ID
	UniqVpcId *string `json:"UniqVpcId,omitnil" name:"UniqVpcId"`

	// Instance inspection/overview status
	InstanceConf *InstanceConfs `json:"InstanceConf,omitnil" name:"InstanceConf"`

	// Resource expiration time
	DeadlineTime *string `json:"DeadlineTime,omitnil" name:"DeadlineTime"`

	// Whether it is an instance supported by DBbrain.
	IsSupported *bool `json:"IsSupported,omitnil" name:"IsSupported"`

	// The status of instance security audit log. ON: enabled, OFF: disabled.
	SecAuditStatus *string `json:"SecAuditStatus,omitnil" name:"SecAuditStatus"`

	// The status of instance audit log. ALL_AUDIT: full audit is enabled, RULE_AUDIT: rule audit is enabled, UNBOUND: audit is disabled.
	AuditPolicyStatus *string `json:"AuditPolicyStatus,omitnil" name:"AuditPolicyStatus"`

	// The running status of instance audit log. normal: running, paused: suspension due to arrears
	AuditRunningStatus *string `json:"AuditRunningStatus,omitnil" name:"AuditRunningStatus"`
}

type IssueTypeInfo struct {
	// Metric categories: AVAILABILITY, MAINTAINABILITY, PERFORMANCE, and RELIABILITY
	IssueType *string `json:"IssueType,omitnil" name:"IssueType"`

	// Exception
	Events []*EventInfo `json:"Events,omitnil" name:"Events"`

	// Total number of the exceptions
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`
}

type MailConfiguration struct {
	// Whether to enable email sending. Valid values: 0 (No), 1 (Yes).
	SendMail *int64 `json:"SendMail,omitnil" name:"SendMail"`

	// Region configuration, such as "ap-guangzhou", "ap-shanghai". For the inspection email sending template, configure the region where you need to send the inspection email. For the subscription email sending template, configure the region to which the current subscribed instance belongs.
	Region []*string `json:"Region,omitnil" name:"Region"`

	// Sending a report with the specified health level, such as "HEALTH", "SUB_HEALTH", "RISK", "HIGH_RISK".
	HealthStatus []*string `json:"HealthStatus,omitnil" name:"HealthStatus"`

	// Contact ID. Either `ContactGroup` or `ContactID` should be passed in.
	ContactPerson []*int64 `json:"ContactPerson,omitnil" name:"ContactPerson"`

	// Contact group ID. Either `ContactGroup` or `ContactID` should be passed in.
	ContactGroup []*int64 `json:"ContactGroup,omitnil" name:"ContactGroup"`
}

// Predefined struct for user
type ModifyDiagDBInstanceConfRequestParams struct {
	// Whether to enable inspection
	InstanceConfs *InstanceConfs `json:"InstanceConfs,omitnil" name:"InstanceConfs"`

	// Target regions of the request. If the value is `All`, it is applied to all regions.
	Regions *string `json:"Regions,omitnil" name:"Regions"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TencentDB for CynosDB (compatible with MySQL)).
	Product *string `json:"Product,omitnil" name:"Product"`

	// ID of the instance to modify.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

type ModifyDiagDBInstanceConfRequest struct {
	*tchttp.BaseRequest
	
	// Whether to enable inspection
	InstanceConfs *InstanceConfs `json:"InstanceConfs,omitnil" name:"InstanceConfs"`

	// Target regions of the request. If the value is `All`, it is applied to all regions.
	Regions *string `json:"Regions,omitnil" name:"Regions"`

	// Service type. Valid values: `mysql` (TencentDB for MySQL), `cynosdb` (TencentDB for CynosDB (compatible with MySQL)).
	Product *string `json:"Product,omitnil" name:"Product"`

	// ID of the instance to modify.
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

func (r *ModifyDiagDBInstanceConfRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDiagDBInstanceConfRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceConfs")
	delete(f, "Regions")
	delete(f, "Product")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDiagDBInstanceConfRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDiagDBInstanceConfResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDiagDBInstanceConfResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDiagDBInstanceConfResponseParams `json:"Response"`
}

func (r *ModifyDiagDBInstanceConfResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDiagDBInstanceConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorFloatMetric struct {
	// Metric name.
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// Metric unit.
	Unit *string `json:"Unit,omitnil" name:"Unit"`

	// Metric value.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Values []*float64 `json:"Values,omitnil" name:"Values"`
}

type MonitorFloatMetricSeriesData struct {
	// Monitoring metric.
	Series []*MonitorFloatMetric `json:"Series,omitnil" name:"Series"`

	// Timestamp corresponding to monitoring metric.
	Timestamp []*int64 `json:"Timestamp,omitnil" name:"Timestamp"`
}

type MonitorMetric struct {
	// Metric name.
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// Metric unit.
	Unit *string `json:"Unit,omitnil" name:"Unit"`

	// Metric value.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Values []*int64 `json:"Values,omitnil" name:"Values"`
}

type MonitorMetricSeriesData struct {
	// Monitoring metric.
	Series []*MonitorMetric `json:"Series,omitnil" name:"Series"`

	// Timestamp corresponding to monitoring metric.
	Timestamp []*int64 `json:"Timestamp,omitnil" name:"Timestamp"`
}

type ProfileInfo struct {
	// Language of the email, such as `en`.
	Language *string `json:"Language,omitnil" name:"Language"`

	// The content of email template.
	MailConfiguration *MailConfiguration `json:"MailConfiguration,omitnil" name:"MailConfiguration"`
}

type SchemaItem struct {
	// Database name
	Schema *string `json:"Schema,omitnil" name:"Schema"`
}

type SchemaSpaceData struct {
	// Database name.
	TableSchema *string `json:"TableSchema,omitnil" name:"TableSchema"`

	// Data space in MB.
	DataLength *float64 `json:"DataLength,omitnil" name:"DataLength"`

	// Index space in MB.
	IndexLength *float64 `json:"IndexLength,omitnil" name:"IndexLength"`

	// Fragmented space in MB.
	DataFree *float64 `json:"DataFree,omitnil" name:"DataFree"`

	// Total space usage in MB.
	TotalLength *float64 `json:"TotalLength,omitnil" name:"TotalLength"`

	// Fragmented rate (%).
	FragRatio *float64 `json:"FragRatio,omitnil" name:"FragRatio"`

	// Number of rows.
	TableRows *int64 `json:"TableRows,omitnil" name:"TableRows"`

	// The total size of the independent physical files corresponding to all the database tables (MB).
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	PhysicalFileSize *float64 `json:"PhysicalFileSize,omitnil" name:"PhysicalFileSize"`
}

type SchemaSpaceTimeSeries struct {
	// Database name
	TableSchema *string `json:"TableSchema,omitnil" name:"TableSchema"`

	// Monitoring metric data in a unit of time interval.
	SeriesData *MonitorMetricSeriesData `json:"SeriesData,omitnil" name:"SeriesData"`
}

type ScoreDetail struct {
	// Deduction item types. Valid values: availability, maintainability, performance, and reliability.
	IssueType *string `json:"IssueType,omitnil" name:"IssueType"`

	// Total scores deducted.
	ScoreLost *int64 `json:"ScoreLost,omitnil" name:"ScoreLost"`

	// Upper limit of the deducted scores.
	ScoreLostMax *int64 `json:"ScoreLostMax,omitnil" name:"ScoreLostMax"`

	// Deduction item list.
	// Note: `null` may be returned for this field, indicating that no valid values can be obtained.
	Items []*ScoreItem `json:"Items,omitnil" name:"Items"`
}

type ScoreItem struct {
	// Exception diagnosis item name.
	DiagItem *string `json:"DiagItem,omitnil" name:"DiagItem"`

	// Diagnosis item types. Valid values: availability, maintainability, performance, and reliability.
	IssueType *string `json:"IssueType,omitnil" name:"IssueType"`

	// Health level. Valid values: information, reminder, alarm, serious, fatal.
	TopSeverity *string `json:"TopSeverity,omitnil" name:"TopSeverity"`

	// Number of occurrences of this exception diagnosis item.
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// Scores deducted.
	ScoreLost *int64 `json:"ScoreLost,omitnil" name:"ScoreLost"`
}

type SlowLogHost struct {
	// Source addresses.
	UserHost *string `json:"UserHost,omitnil" name:"UserHost"`

	// The proportion (in %) of slow logs from this source address to the total number of slow logs
	Ratio *float64 `json:"Ratio,omitnil" name:"Ratio"`

	// Number of slow logs from this source address.
	Count *int64 `json:"Count,omitnil" name:"Count"`
}

type SlowLogTopSqlItem struct {
	// Total SQL lock wait time
	LockTime *float64 `json:"LockTime,omitnil" name:"LockTime"`

	// Maximum lock wait time
	LockTimeMax *float64 `json:"LockTimeMax,omitnil" name:"LockTimeMax"`

	// Minimum lock wait time
	LockTimeMin *float64 `json:"LockTimeMin,omitnil" name:"LockTimeMin"`

	// Total number of scanned rows
	RowsExamined *int64 `json:"RowsExamined,omitnil" name:"RowsExamined"`

	// Maximum number of scanned rows
	RowsExaminedMax *int64 `json:"RowsExaminedMax,omitnil" name:"RowsExaminedMax"`

	// Minimum number of scanned rows
	RowsExaminedMin *int64 `json:"RowsExaminedMin,omitnil" name:"RowsExaminedMin"`

	// Total duration
	QueryTime *float64 `json:"QueryTime,omitnil" name:"QueryTime"`

	// Maximum execution time
	QueryTimeMax *float64 `json:"QueryTimeMax,omitnil" name:"QueryTimeMax"`

	// Minimum execution time
	QueryTimeMin *float64 `json:"QueryTimeMin,omitnil" name:"QueryTimeMin"`

	// Total number of returned rows
	RowsSent *int64 `json:"RowsSent,omitnil" name:"RowsSent"`

	// Maximum number of returned rows
	RowsSentMax *int64 `json:"RowsSentMax,omitnil" name:"RowsSentMax"`

	// Minimum number of returned rows
	RowsSentMin *int64 `json:"RowsSentMin,omitnil" name:"RowsSentMin"`

	// Number of executions
	ExecTimes *int64 `json:"ExecTimes,omitnil" name:"ExecTimes"`

	// SQL template
	SqlTemplate *string `json:"SqlTemplate,omitnil" name:"SqlTemplate"`

	// SQL with parameter (random)
	SqlText *string `json:"SqlText,omitnil" name:"SqlText"`

	// Database name
	Schema *string `json:"Schema,omitnil" name:"Schema"`

	// Ratio of total duration
	QueryTimeRatio *float64 `json:"QueryTimeRatio,omitnil" name:"QueryTimeRatio"`

	// Ratio of total SQL lock wait time
	LockTimeRatio *float64 `json:"LockTimeRatio,omitnil" name:"LockTimeRatio"`

	// Ratio of total number of scanned rows
	RowsExaminedRatio *float64 `json:"RowsExaminedRatio,omitnil" name:"RowsExaminedRatio"`

	// Ratio of total number of returned rows
	RowsSentRatio *float64 `json:"RowsSentRatio,omitnil" name:"RowsSentRatio"`

	// Average execution time
	QueryTimeAvg *float64 `json:"QueryTimeAvg,omitnil" name:"QueryTimeAvg"`

	// Average number of rows returned
	RowsSentAvg *float64 `json:"RowsSentAvg,omitnil" name:"RowsSentAvg"`

	// Average lock wait time
	LockTimeAvg *float64 `json:"LockTimeAvg,omitnil" name:"LockTimeAvg"`

	// Average number of rows scanned
	RowsExaminedAvg *float64 `json:"RowsExaminedAvg,omitnil" name:"RowsExaminedAvg"`
}

type TableSpaceData struct {
	// Table name.
	TableName *string `json:"TableName,omitnil" name:"TableName"`

	// Database name.
	TableSchema *string `json:"TableSchema,omitnil" name:"TableSchema"`

	// Database table storage engine.
	Engine *string `json:"Engine,omitnil" name:"Engine"`

	// Data space in MB.
	DataLength *float64 `json:"DataLength,omitnil" name:"DataLength"`

	// Index space in MB.
	IndexLength *float64 `json:"IndexLength,omitnil" name:"IndexLength"`

	// Fragmented space in MB.
	DataFree *float64 `json:"DataFree,omitnil" name:"DataFree"`

	// Total space usage in MB.
	TotalLength *float64 `json:"TotalLength,omitnil" name:"TotalLength"`

	// Fragmented rate (%).
	FragRatio *float64 `json:"FragRatio,omitnil" name:"FragRatio"`

	// Number of rows.
	TableRows *int64 `json:"TableRows,omitnil" name:"TableRows"`

	// Size in MB of the physical file exclusive to a table.
	PhysicalFileSize *float64 `json:"PhysicalFileSize,omitnil" name:"PhysicalFileSize"`
}

type TableSpaceTimeSeries struct {
	// Table name.
	TableName *string `json:"TableName,omitnil" name:"TableName"`

	// Database name.
	TableSchema *string `json:"TableSchema,omitnil" name:"TableSchema"`

	// Database table storage engine.
	Engine *string `json:"Engine,omitnil" name:"Engine"`

	// Monitoring metric data in a unit of time interval.
	SeriesData *MonitorFloatMetricSeriesData `json:"SeriesData,omitnil" name:"SeriesData"`
}

type TimeSlice struct {
	// Total number
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// Statistics start time
	Timestamp *int64 `json:"Timestamp,omitnil" name:"Timestamp"`
}

type UserProfile struct {
	// Configured ID
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ProfileId *string `json:"ProfileId,omitnil" name:"ProfileId"`

	// Configuration type
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ProfileType *string `json:"ProfileType,omitnil" name:"ProfileType"`

	// Configuration level. Valid values: “User” or “Instance”
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ProfileLevel *string `json:"ProfileLevel,omitnil" name:"ProfileLevel"`

	// Configuration name
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ProfileName *string `json:"ProfileName,omitnil" name:"ProfileName"`

	// Configuration details
	ProfileInfo *ProfileInfo `json:"ProfileInfo,omitnil" name:"ProfileInfo"`
}