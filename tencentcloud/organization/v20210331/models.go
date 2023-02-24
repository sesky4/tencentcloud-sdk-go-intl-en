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

package v20210331

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

// Predefined struct for user
type AddOrganizationNodeRequestParams struct {
	// Parent node ID, which can be obtained through the `DescribeOrganizationNodes` API.
	ParentNodeId *uint64 `json:"ParentNodeId,omitempty" name:"ParentNodeId"`

	// Node name, which can contain up to 40 letters, digits, and symbols `+@&._[]-`.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Remarks.
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type AddOrganizationNodeRequest struct {
	*tchttp.BaseRequest
	
	// Parent node ID, which can be obtained through the `DescribeOrganizationNodes` API.
	ParentNodeId *uint64 `json:"ParentNodeId,omitempty" name:"ParentNodeId"`

	// Node name, which can contain up to 40 letters, digits, and symbols `+@&._[]-`.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Remarks.
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *AddOrganizationNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddOrganizationNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ParentNodeId")
	delete(f, "Name")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddOrganizationNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddOrganizationNodeResponseParams struct {
	// Node ID.
	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddOrganizationNodeResponse struct {
	*tchttp.BaseResponse
	Response *AddOrganizationNodeResponseParams `json:"Response"`
}

func (r *AddOrganizationNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddOrganizationNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindOrganizationMemberAuthAccountRequestParams struct {
	// Member UIN.
	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`

	// Policy ID, which can be obtained through the `DescribeOrganizationMemberPolicies` API.
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// List of sub-account UINs of the organization admin, which can contain up to five UINs.
	OrgSubAccountUins []*int64 `json:"OrgSubAccountUins,omitempty" name:"OrgSubAccountUins"`
}

type BindOrganizationMemberAuthAccountRequest struct {
	*tchttp.BaseRequest
	
	// Member UIN.
	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`

	// Policy ID, which can be obtained through the `DescribeOrganizationMemberPolicies` API.
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// List of sub-account UINs of the organization admin, which can contain up to five UINs.
	OrgSubAccountUins []*int64 `json:"OrgSubAccountUins,omitempty" name:"OrgSubAccountUins"`
}

func (r *BindOrganizationMemberAuthAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindOrganizationMemberAuthAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	delete(f, "PolicyId")
	delete(f, "OrgSubAccountUins")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindOrganizationMemberAuthAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindOrganizationMemberAuthAccountResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindOrganizationMemberAuthAccountResponse struct {
	*tchttp.BaseResponse
	Response *BindOrganizationMemberAuthAccountResponseParams `json:"Response"`
}

func (r *BindOrganizationMemberAuthAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindOrganizationMemberAuthAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelOrganizationMemberAuthAccountRequestParams struct {
	// Member UIN.
	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`

	// Policy ID.
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Organization sub-account UIN.
	OrgSubAccountUin *int64 `json:"OrgSubAccountUin,omitempty" name:"OrgSubAccountUin"`
}

type CancelOrganizationMemberAuthAccountRequest struct {
	*tchttp.BaseRequest
	
	// Member UIN.
	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`

	// Policy ID.
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Organization sub-account UIN.
	OrgSubAccountUin *int64 `json:"OrgSubAccountUin,omitempty" name:"OrgSubAccountUin"`
}

func (r *CancelOrganizationMemberAuthAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelOrganizationMemberAuthAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	delete(f, "PolicyId")
	delete(f, "OrgSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelOrganizationMemberAuthAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelOrganizationMemberAuthAccountResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CancelOrganizationMemberAuthAccountResponse struct {
	*tchttp.BaseResponse
	Response *CancelOrganizationMemberAuthAccountResponseParams `json:"Response"`
}

func (r *CancelOrganizationMemberAuthAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelOrganizationMemberAuthAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationMemberPolicyRequestParams struct {
	// Member UIN.
	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`

	// Policy name, which can contain up to 128 letters, digits, and symbols `+=,.@_-`.
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// Member access identity ID, which can be obtained through the `DescribeOrganizationMemberAuthIdentities` API.
	IdentityId *int64 `json:"IdentityId,omitempty" name:"IdentityId"`

	// Description.
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateOrganizationMemberPolicyRequest struct {
	*tchttp.BaseRequest
	
	// Member UIN.
	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`

	// Policy name, which can contain up to 128 letters, digits, and symbols `+=,.@_-`.
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// Member access identity ID, which can be obtained through the `DescribeOrganizationMemberAuthIdentities` API.
	IdentityId *int64 `json:"IdentityId,omitempty" name:"IdentityId"`

	// Description.
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateOrganizationMemberPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationMemberPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	delete(f, "PolicyName")
	delete(f, "IdentityId")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrganizationMemberPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationMemberPolicyResponseParams struct {
	// Policy ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateOrganizationMemberPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrganizationMemberPolicyResponseParams `json:"Response"`
}

func (r *CreateOrganizationMemberPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationMemberPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationMemberRequestParams struct {
	// Member name, which can contain up to 25 letters, digits, and symbols `+@&._[]-:,`.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Relationship policy. Valid value: `Financial`.
	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`

	// List of member financial permission IDs. `7` indicates paying, which is the default value.
	PermissionIds []*uint64 `json:"PermissionIds,omitempty" name:"PermissionIds"`

	// ID of the node of the member's department, which can be obtained through the `DescribeOrganizationNodes` API.
	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`

	// Account name, which can contain up to 25 letters, digits, and symbols `+@&._[]-:,`.
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// Remarks.
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Member creation record ID, which is required during retry upon creation exception.
	RecordId *int64 `json:"RecordId,omitempty" name:"RecordId"`

	// Payer UIN, which is required during paying for a member.
	PayUin *string `json:"PayUin,omitempty" name:"PayUin"`

	// List of member access identity IDs, which can be obtained through the `ListOrganizationIdentity` API. `1` indicates supported, which is the default value.
	IdentityRoleID []*uint64 `json:"IdentityRoleID,omitempty" name:"IdentityRoleID"`

	// Verified entity relationship ID, which is required during creating members for different entities.
	AuthRelationId *int64 `json:"AuthRelationId,omitempty" name:"AuthRelationId"`
}

type CreateOrganizationMemberRequest struct {
	*tchttp.BaseRequest
	
	// Member name, which can contain up to 25 letters, digits, and symbols `+@&._[]-:,`.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Relationship policy. Valid value: `Financial`.
	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`

	// List of member financial permission IDs. `7` indicates paying, which is the default value.
	PermissionIds []*uint64 `json:"PermissionIds,omitempty" name:"PermissionIds"`

	// ID of the node of the member's department, which can be obtained through the `DescribeOrganizationNodes` API.
	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`

	// Account name, which can contain up to 25 letters, digits, and symbols `+@&._[]-:,`.
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// Remarks.
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Member creation record ID, which is required during retry upon creation exception.
	RecordId *int64 `json:"RecordId,omitempty" name:"RecordId"`

	// Payer UIN, which is required during paying for a member.
	PayUin *string `json:"PayUin,omitempty" name:"PayUin"`

	// List of member access identity IDs, which can be obtained through the `ListOrganizationIdentity` API. `1` indicates supported, which is the default value.
	IdentityRoleID []*uint64 `json:"IdentityRoleID,omitempty" name:"IdentityRoleID"`

	// Verified entity relationship ID, which is required during creating members for different entities.
	AuthRelationId *int64 `json:"AuthRelationId,omitempty" name:"AuthRelationId"`
}

func (r *CreateOrganizationMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "PolicyType")
	delete(f, "PermissionIds")
	delete(f, "NodeId")
	delete(f, "AccountName")
	delete(f, "Remark")
	delete(f, "RecordId")
	delete(f, "PayUin")
	delete(f, "IdentityRoleID")
	delete(f, "AuthRelationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrganizationMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationMemberResponseParams struct {
	// Member UIN.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Uin *int64 `json:"Uin,omitempty" name:"Uin"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateOrganizationMemberResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrganizationMemberResponseParams `json:"Response"`
}

func (r *CreateOrganizationMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationMembersRequestParams struct {
	// List of UINs of the members to be deleted.
	MemberUin []*int64 `json:"MemberUin,omitempty" name:"MemberUin"`
}

type DeleteOrganizationMembersRequest struct {
	*tchttp.BaseRequest
	
	// List of UINs of the members to be deleted.
	MemberUin []*int64 `json:"MemberUin,omitempty" name:"MemberUin"`
}

func (r *DeleteOrganizationMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOrganizationMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationMembersResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteOrganizationMembersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOrganizationMembersResponseParams `json:"Response"`
}

func (r *DeleteOrganizationMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationNodesRequestParams struct {
	// List of node IDs.
	NodeId []*int64 `json:"NodeId,omitempty" name:"NodeId"`
}

type DeleteOrganizationNodesRequest struct {
	*tchttp.BaseRequest
	
	// List of node IDs.
	NodeId []*int64 `json:"NodeId,omitempty" name:"NodeId"`
}

func (r *DeleteOrganizationNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NodeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOrganizationNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationNodesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteOrganizationNodesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOrganizationNodesResponseParams `json:"Response"`
}

func (r *DeleteOrganizationNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMemberAuthAccountsRequestParams struct {
	// Offset.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of returned results.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Member UIN.
	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`

	// Policy ID.
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

type DescribeOrganizationMemberAuthAccountsRequest struct {
	*tchttp.BaseRequest
	
	// Offset.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of returned results.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Member UIN.
	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`

	// Policy ID.
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DescribeOrganizationMemberAuthAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberAuthAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "MemberUin")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationMemberAuthAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMemberAuthAccountsResponseParams struct {
	// List
	// Note: This field may return null, indicating that no valid values can be obtained.
	Items []*OrgMemberAuthAccount `json:"Items,omitempty" name:"Items"`

	// Total number
	// Note: This field may return null, indicating that no valid values can be obtained.
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOrganizationMemberAuthAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationMemberAuthAccountsResponseParams `json:"Response"`
}

func (r *DescribeOrganizationMemberAuthAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberAuthAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMemberAuthIdentitiesRequestParams struct {
	// Offset.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of returned results. Maximum value: `50`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Organization member UIN.
	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
}

type DescribeOrganizationMemberAuthIdentitiesRequest struct {
	*tchttp.BaseRequest
	
	// Offset.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of returned results. Maximum value: `50`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Organization member UIN.
	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`
}

func (r *DescribeOrganizationMemberAuthIdentitiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberAuthIdentitiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "MemberUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationMemberAuthIdentitiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMemberAuthIdentitiesResponseParams struct {
	// List.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Items []*OrgMemberAuthIdentity `json:"Items,omitempty" name:"Items"`

	// Total number.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOrganizationMemberAuthIdentitiesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationMemberAuthIdentitiesResponseParams `json:"Response"`
}

func (r *DescribeOrganizationMemberAuthIdentitiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberAuthIdentitiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMemberPoliciesRequestParams struct {
	// Offset.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of returned results. Maximum value: `50`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Member UIN.
	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`

	// Search keyword, which can be the policy name or description.
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

type DescribeOrganizationMemberPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// Offset.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of returned results. Maximum value: `50`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Member UIN.
	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`

	// Search keyword, which can be the policy name or description.
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *DescribeOrganizationMemberPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "MemberUin")
	delete(f, "SearchKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationMemberPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMemberPoliciesResponseParams struct {
	// List.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Items []*OrgMemberPolicy `json:"Items,omitempty" name:"Items"`

	// Total number.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOrganizationMemberPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationMemberPoliciesResponseParams `json:"Response"`
}

func (r *DescribeOrganizationMemberPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMembersRequestParams struct {
	// Offset.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of returned results. Maximum value: `50`.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Valid values: `en` (Tencent Cloud International); `zh` (Tencent Cloud).
	Lang *string `json:"Lang,omitempty" name:"Lang"`

	// Search by member name or ID.
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// Entity name.
	AuthName *string `json:"AuthName,omitempty" name:"AuthName"`

	// Abbreviation of the trusted service, which is required during querying the trusted service admin.
	Product *string `json:"Product,omitempty" name:"Product"`
}

type DescribeOrganizationMembersRequest struct {
	*tchttp.BaseRequest
	
	// Offset.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of returned results. Maximum value: `50`.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Valid values: `en` (Tencent Cloud International); `zh` (Tencent Cloud).
	Lang *string `json:"Lang,omitempty" name:"Lang"`

	// Search by member name or ID.
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// Entity name.
	AuthName *string `json:"AuthName,omitempty" name:"AuthName"`

	// Abbreviation of the trusted service, which is required during querying the trusted service admin.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeOrganizationMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Lang")
	delete(f, "SearchKey")
	delete(f, "AuthName")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMembersResponseParams struct {
	// Member list.
	Items []*OrgMember `json:"Items,omitempty" name:"Items"`

	// Total number.
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOrganizationMembersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationMembersResponseParams `json:"Response"`
}

func (r *DescribeOrganizationMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationNodesRequestParams struct {
	// Maximum number of returned results. Maximum value: `50`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeOrganizationNodesRequest struct {
	*tchttp.BaseRequest
	
	// Maximum number of returned results. Maximum value: `50`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeOrganizationNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationNodesResponseParams struct {
	// Total number.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// List details.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Items []*OrgNode `json:"Items,omitempty" name:"Items"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOrganizationNodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationNodesResponseParams `json:"Response"`
}

func (r *DescribeOrganizationNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationRequestParams struct {
	// Valid values: `en` (Tencent Cloud International); `zh` (Tencent Cloud).
	Lang *string `json:"Lang,omitempty" name:"Lang"`

	// Abbreviation of the trusted service, which is required during querying the trusted service admin.
	Product *string `json:"Product,omitempty" name:"Product"`
}

type DescribeOrganizationRequest struct {
	*tchttp.BaseRequest
	
	// Valid values: `en` (Tencent Cloud International); `zh` (Tencent Cloud).
	Lang *string `json:"Lang,omitempty" name:"Lang"`

	// Abbreviation of the trusted service, which is required during querying the trusted service admin.
	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Lang")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationResponseParams struct {
	// Organization ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrgId *int64 `json:"OrgId,omitempty" name:"OrgId"`

	// Creator UIN.
	// Note: This field may return null, indicating that no valid values can be obtained.
	HostUin *int64 `json:"HostUin,omitempty" name:"HostUin"`

	// Creator name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NickName *string `json:"NickName,omitempty" name:"NickName"`

	// Organization type.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrgType *int64 `json:"OrgType,omitempty" name:"OrgType"`

	// Whether the member is the organization admin. Valid values: `true` (yes); `false` (no).
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsManager *bool `json:"IsManager,omitempty" name:"IsManager"`

	// Policy type. Valid values: `Financial` (finance management).
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrgPolicyType *string `json:"OrgPolicyType,omitempty" name:"OrgPolicyType"`

	// Policy name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrgPolicyName *string `json:"OrgPolicyName,omitempty" name:"OrgPolicyName"`

	// List of member financial permissions.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrgPermission []*OrgPermission `json:"OrgPermission,omitempty" name:"OrgPermission"`

	// Organization root node ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RootNodeId *int64 `json:"RootNodeId,omitempty" name:"RootNodeId"`

	// Organization creation time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Member joining time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	JoinTime *string `json:"JoinTime,omitempty" name:"JoinTime"`

	// Whether the member is allowed to leave. Valid values: `Allow`, `Denied`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsAllowQuit *string `json:"IsAllowQuit,omitempty" name:"IsAllowQuit"`

	// Payer UIN.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayUin *string `json:"PayUin,omitempty" name:"PayUin"`

	// Payer name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayName *string `json:"PayName,omitempty" name:"PayName"`

	// Whether the member is the trusted service admin. Valid values: `true` (yes); `false` (no).
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsAssignManager *bool `json:"IsAssignManager,omitempty" name:"IsAssignManager"`

	// Whether the member is the verified entity admin. Valid values: `true` (yes); `false` (no).
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsAuthManager *bool `json:"IsAuthManager,omitempty" name:"IsAuthManager"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationResponseParams `json:"Response"`
}

func (r *DescribeOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IdentityPolicy struct {
	// Policy ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Policy name
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
}

// Predefined struct for user
type ListOrganizationIdentityRequestParams struct {
	// Offset.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of returned results. Maximum value: `50`.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Search by name.
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// Search by identity ID.
	IdentityId *uint64 `json:"IdentityId,omitempty" name:"IdentityId"`
}

type ListOrganizationIdentityRequest struct {
	*tchttp.BaseRequest
	
	// Offset.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Maximum number of returned results. Maximum value: `50`.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Search by name.
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// Search by identity ID.
	IdentityId *uint64 `json:"IdentityId,omitempty" name:"IdentityId"`
}

func (r *ListOrganizationIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrganizationIdentityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchKey")
	delete(f, "IdentityId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListOrganizationIdentityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOrganizationIdentityResponseParams struct {
	// Total number.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// Item details.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Items []*OrgIdentity `json:"Items,omitempty" name:"Items"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListOrganizationIdentityResponse struct {
	*tchttp.BaseResponse
	Response *ListOrganizationIdentityResponseParams `json:"Response"`
}

func (r *ListOrganizationIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrganizationIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MemberIdentity struct {
	// Identity ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityId *int64 `json:"IdentityId,omitempty" name:"IdentityId"`

	// Identity name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityAliasName *string `json:"IdentityAliasName,omitempty" name:"IdentityAliasName"`
}

// Predefined struct for user
type MoveOrganizationNodeMembersRequestParams struct {
	// Organization node ID.
	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`

	// Member UIN list.
	MemberUin []*int64 `json:"MemberUin,omitempty" name:"MemberUin"`
}

type MoveOrganizationNodeMembersRequest struct {
	*tchttp.BaseRequest
	
	// Organization node ID.
	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`

	// Member UIN list.
	MemberUin []*int64 `json:"MemberUin,omitempty" name:"MemberUin"`
}

func (r *MoveOrganizationNodeMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MoveOrganizationNodeMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NodeId")
	delete(f, "MemberUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MoveOrganizationNodeMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MoveOrganizationNodeMembersResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type MoveOrganizationNodeMembersResponse struct {
	*tchttp.BaseResponse
	Response *MoveOrganizationNodeMembersResponseParams `json:"Response"`
}

func (r *MoveOrganizationNodeMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MoveOrganizationNodeMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrgIdentity struct {
	// Identity ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityId *int64 `json:"IdentityId,omitempty" name:"IdentityId"`

	// Identity name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityAliasName *string `json:"IdentityAliasName,omitempty" name:"IdentityAliasName"`

	// Description.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Identity policy.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityPolicy []*IdentityPolicy `json:"IdentityPolicy,omitempty" name:"IdentityPolicy"`

	// Identity type. Valid values: `1` (preset); `2` (custom).
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityType *uint64 `json:"IdentityType,omitempty" name:"IdentityType"`

	// Update time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type OrgMember struct {
	// Member UIN
	// Note: This field may return null, indicating that no valid values can be obtained.
	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`

	// Member name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Member type. Valid values: `Invite` (invited); `Create` (created).
	// Note: This field may return null, indicating that no valid values can be obtained.
	MemberType *string `json:"MemberType,omitempty" name:"MemberType"`

	// Relationship policy type
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrgPolicyType *string `json:"OrgPolicyType,omitempty" name:"OrgPolicyType"`

	// Relationship policy name
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrgPolicyName *string `json:"OrgPolicyName,omitempty" name:"OrgPolicyName"`

	// Relationship policy permission
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrgPermission []*OrgPermission `json:"OrgPermission,omitempty" name:"OrgPermission"`

	// Node ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`

	// Node name
	// Note: This field may return null, indicating that no valid values can be obtained.
	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`

	// Remarks
	// Note: This field may return null, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Creation time
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Update time
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Whether the member is allowed to leave. Valid values: `Allow`, `Denied`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsAllowQuit *string `json:"IsAllowQuit,omitempty" name:"IsAllowQuit"`

	// Payer UIN
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayUin *string `json:"PayUin,omitempty" name:"PayUin"`

	// Payer name
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayName *string `json:"PayName,omitempty" name:"PayName"`

	// Management identity
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrgIdentity []*MemberIdentity `json:"OrgIdentity,omitempty" name:"OrgIdentity"`

	// Security information binding status. Valid values: `Unbound`, `Valid`, `Success`, `Failed`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BindStatus *string `json:"BindStatus,omitempty" name:"BindStatus"`

	// Member permission status. Valid values: `Confirmed`, `UnConfirmed`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PermissionStatus *string `json:"PermissionStatus,omitempty" name:"PermissionStatus"`
}

type OrgMemberAuthAccount struct {
	// Organization sub-account UIN.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrgSubAccountUin *int64 `json:"OrgSubAccountUin,omitempty" name:"OrgSubAccountUin"`

	// Policy ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Policy name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// Identity ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityId *int64 `json:"IdentityId,omitempty" name:"IdentityId"`

	// Identity role name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityRoleName *string `json:"IdentityRoleName,omitempty" name:"IdentityRoleName"`

	// Identity role alias.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityRoleAliasName *string `json:"IdentityRoleAliasName,omitempty" name:"IdentityRoleAliasName"`

	// Creation time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Update time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Sub-account name
	// Note: This field may return null, indicating that no valid values can be obtained.
	OrgSubAccountName *string `json:"OrgSubAccountName,omitempty" name:"OrgSubAccountName"`
}

type OrgMemberAuthIdentity struct {
	// Identity ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityId *int64 `json:"IdentityId,omitempty" name:"IdentityId"`

	// Identity role name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityRoleName *string `json:"IdentityRoleName,omitempty" name:"IdentityRoleName"`

	// Identity role alias.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityRoleAliasName *string `json:"IdentityRoleAliasName,omitempty" name:"IdentityRoleAliasName"`

	// Description
	// Note: This field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Creation time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Update time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type OrgMemberPolicy struct {
	// Policy ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// Policy name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// Identity ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityId *int64 `json:"IdentityId,omitempty" name:"IdentityId"`

	// Identity role name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityRoleName *string `json:"IdentityRoleName,omitempty" name:"IdentityRoleName"`

	// Identity role alias.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityRoleAliasName *string `json:"IdentityRoleAliasName,omitempty" name:"IdentityRoleAliasName"`

	// Description.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Creation time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Update time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type OrgNode struct {
	// Organization node ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`

	// Name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Parent node ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ParentNodeId *int64 `json:"ParentNodeId,omitempty" name:"ParentNodeId"`

	// Remarks
	// Note: This field may return null, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Creation time
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Update time
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type OrgPermission struct {
	// Permission ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// Permission name
	Name *string `json:"Name,omitempty" name:"Name"`
}

// Predefined struct for user
type UpdateOrganizationNodeRequestParams struct {
	// Node ID.
	NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`

	// Node name, which can contain up to 40 letters, digits, and symbols `+@&._[]-`.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Remarks.
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type UpdateOrganizationNodeRequest struct {
	*tchttp.BaseRequest
	
	// Node ID.
	NodeId *uint64 `json:"NodeId,omitempty" name:"NodeId"`

	// Node name, which can contain up to 40 letters, digits, and symbols `+@&._[]-`.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Remarks.
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *UpdateOrganizationNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NodeId")
	delete(f, "Name")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateOrganizationNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOrganizationNodeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateOrganizationNodeResponse struct {
	*tchttp.BaseResponse
	Response *UpdateOrganizationNodeResponseParams `json:"Response"`
}

func (r *UpdateOrganizationNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}