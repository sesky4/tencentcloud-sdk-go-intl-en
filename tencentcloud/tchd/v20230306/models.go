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

package v20230306

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

// Predefined struct for user
type DescribeEventsRequestParams struct {
	// event occurrence date
	EventDate *string `json:"EventDate,omitnil,omitempty" name:"EventDate"`

	// Query by Product ID(s). Product ID examples: cvm, lb, cdb, cdn, crs.
	ProductIds []*string `json:"ProductIds,omitnil,omitempty" name:"ProductIds"`

	//  1. Query by Region ID(s). Region ID examples: ap-guangzhou、ap-shanghai、ap-singapore.
	// 2. The region ID for non-region-specific products should be set to non-regional.
	RegionIds []*string `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`
}

type DescribeEventsRequest struct {
	*tchttp.BaseRequest
	
	// event occurrence date
	EventDate *string `json:"EventDate,omitnil,omitempty" name:"EventDate"`

	// Query by Product ID(s). Product ID examples: cvm, lb, cdb, cdn, crs.
	ProductIds []*string `json:"ProductIds,omitnil,omitempty" name:"ProductIds"`

	//  1. Query by Region ID(s). Region ID examples: ap-guangzhou、ap-shanghai、ap-singapore.
	// 2. The region ID for non-region-specific products should be set to non-regional.
	RegionIds []*string `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`
}

func (r *DescribeEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventDate")
	delete(f, "ProductIds")
	delete(f, "RegionIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEventsResponseParams struct {
	// Detailed event information.
	Data *ProductEventList `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEventsResponseParams `json:"Response"`
}

func (r *DescribeEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EventDetail struct {
	// Product ID.
	ProductId *string `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// Product name.
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// Region ID.
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Region name.
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// Event start time.
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Event end time. If the event is still ongoing and has not ended, the end time will be empty.
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Current status: Normally, Informational, Degradation.
	CurrentStatus *string `json:"CurrentStatus,omitnil,omitempty" name:"CurrentStatus"`
}

type ProductEventList struct {
	// Detailed event information.
	// Note: this field may return null, indicating that no valid value is obtained.
	EventList []*EventDetail `json:"EventList,omitnil,omitempty" name:"EventList"`
}