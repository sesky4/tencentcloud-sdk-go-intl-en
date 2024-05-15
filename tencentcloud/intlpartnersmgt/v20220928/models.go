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

package v20220928

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type ActionSummaryOverviewItem struct {
	// Transaction type code
	// Note: This field may return null, indicating that no valid values can be obtained.
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Transaction type name
	// Note: This field may return null, indicating that no valid values can be obtained.
	ActionTypeName *string `json:"ActionTypeName,omitnil,omitempty" name:"ActionTypeName"`

	// The actual total consumption amount accurate down to eight decimal places
	// Note: This field may return null, indicating that no valid values can be obtained.
	OriginalCost *string `json:"OriginalCost,omitnil,omitempty" name:"OriginalCost"`

	// The deducted voucher amount accurate down to eight decimal places
	// Note: This field may return null, indicating that no valid values can be obtained.
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Total consumption amount accurate down to eight decimal places
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`
}

// Predefined struct for user
type AllocateCustomerCreditRequestParams struct {
	// Specific value of the credit allocated to the customer
	AddedCredit *float64 `json:"AddedCredit,omitnil,omitempty" name:"AddedCredit"`

	// Customer UIN
	ClientUin *uint64 `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`
}

type AllocateCustomerCreditRequest struct {
	*tchttp.BaseRequest
	
	// Specific value of the credit allocated to the customer
	AddedCredit *float64 `json:"AddedCredit,omitnil,omitempty" name:"AddedCredit"`

	// Customer UIN
	ClientUin *uint64 `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`
}

func (r *AllocateCustomerCreditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AllocateCustomerCreditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddedCredit")
	delete(f, "ClientUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AllocateCustomerCreditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AllocateCustomerCreditResponseParams struct {
	// The updated total credit
	TotalCredit *float64 `json:"TotalCredit,omitnil,omitempty" name:"TotalCredit"`

	// The updated available credit
	RemainingCredit *float64 `json:"RemainingCredit,omitnil,omitempty" name:"RemainingCredit"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AllocateCustomerCreditResponse struct {
	*tchttp.BaseResponse
	Response *AllocateCustomerCreditResponseParams `json:"Response"`
}

func (r *AllocateCustomerCreditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AllocateCustomerCreditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BillDetailData struct {
	// Reseller account
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayerAccountId *int64 `json:"PayerAccountId,omitnil,omitempty" name:"PayerAccountId"`

	// Customer account
	// Note: This field may return null, indicating that no valid values can be obtained.
	OwnerAccountId *int64 `json:"OwnerAccountId,omitnil,omitempty" name:"OwnerAccountId"`

	// Operator account
	// Note: This field may return null, indicating that no valid values can be obtained.
	OperatorAccountId *int64 `json:"OperatorAccountId,omitnil,omitempty" name:"OperatorAccountId"`

	// Product name
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// Billing mode
	// `Monthly subscription` (Monthly subscription)
	// `Pay-As-You-Go resources` (Pay-as-you-go)
	// `Standard RI` (Reserved instance)
	// Note: This field may return null, indicating that no valid values can be obtained.
	BillingMode *string `json:"BillingMode,omitnil,omitempty" name:"BillingMode"`

	// Project name
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// Resource region
	// Note: This field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Resource AZ
	// Note: This field may return null, indicating that no valid values can be obtained.
	AvailabilityZone *string `json:"AvailabilityZone,omitnil,omitempty" name:"AvailabilityZone"`

	// Instance ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Subproduct name
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubProductName *string `json:"SubProductName,omitnil,omitempty" name:"SubProductName"`

	// Settlement type
	// Note: This field may return null, indicating that no valid values can be obtained.
	TransactionType *string `json:"TransactionType,omitnil,omitempty" name:"TransactionType"`

	// Transaction ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	TransactionId *string `json:"TransactionId,omitnil,omitempty" name:"TransactionId"`

	// Settlement time
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	TransactionTime *string `json:"TransactionTime,omitnil,omitempty" name:"TransactionTime"`

	// Start time of resource use
	// Note: This field may return null, indicating that no valid values can be obtained.
	UsageStartTime *string `json:"UsageStartTime,omitnil,omitempty" name:"UsageStartTime"`

	// End time of resource use
	// Note: This field may return null, indicating that no valid values can be obtained.
	UsageEndTime *string `json:"UsageEndTime,omitnil,omitempty" name:"UsageEndTime"`

	// Component
	// Note: This field may return null, indicating that no valid values can be obtained.
	ComponentType *string `json:"ComponentType,omitnil,omitempty" name:"ComponentType"`

	// Component name
	// Note: This field may return null, indicating that no valid values can be obtained.
	ComponentName *string `json:"ComponentName,omitnil,omitempty" name:"ComponentName"`

	// Component list price
	// Note: This field may return null, indicating that no valid values can be obtained.
	ComponentListPrice *string `json:"ComponentListPrice,omitnil,omitempty" name:"ComponentListPrice"`

	// Price unit
	// Note: This field may return null, indicating that no valid values can be obtained.
	ComponentPriceMeasurementUnit *string `json:"ComponentPriceMeasurementUnit,omitnil,omitempty" name:"ComponentPriceMeasurementUnit"`

	// Component usage
	// Note: This field may return null, indicating that no valid values can be obtained.
	ComponentUsage *string `json:"ComponentUsage,omitnil,omitempty" name:"ComponentUsage"`

	// Component usage unit
	// Note: This field may return null, indicating that no valid values can be obtained.
	ComponentUsageUnit *string `json:"ComponentUsageUnit,omitnil,omitempty" name:"ComponentUsageUnit"`

	// Resource usage duration
	// Note: This field may return null, indicating that no valid values can be obtained.
	UsageDuration *string `json:"UsageDuration,omitnil,omitempty" name:"UsageDuration"`

	// Duration unit
	// Note: This field may return null, indicating that no valid values can be obtained.
	DurationUnit *string `json:"DurationUnit,omitnil,omitempty" name:"DurationUnit"`

	// Original cost
	// Original cost = component list price * component usage * usage duration
	// Note: This field may return null, indicating that no valid values can be obtained.
	OriginalCost *string `json:"OriginalCost,omitnil,omitempty" name:"OriginalCost"`

	// Discount, which defaults to `1`, indicating there is no discount.
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiscountRate *string `json:"DiscountRate,omitnil,omitempty" name:"DiscountRate"`

	// Currency
	// Note: This field may return null, indicating that no valid values can be obtained.
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// Discounted total
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalAmountAfterDiscount *string `json:"TotalAmountAfterDiscount,omitnil,omitempty" name:"TotalAmountAfterDiscount"`

	// Voucher deduction
	// Note: This field may return null, indicating that no valid values can be obtained.
	VoucherDeduction *string `json:"VoucherDeduction,omitnil,omitempty" name:"VoucherDeduction"`

	// Total cost = discounted total - voucher deduction
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// ID
	// Note: The return value may be null, indicating that no valid data can be obtained.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type BusinessInfo struct {
	// ProductNote: This field may return null, indicating that no valid values can be obtained.
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// Product codeNote: This field may return null, indicating that no valid values can be obtained.
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Original price
	// Note: This field may return null, indicating that no valid values can be obtained.
	OriginalCost *string `json:"OriginalCost,omitnil,omitempty" name:"OriginalCost"`

	// Voucher amount
	// Note: This field may return null, indicating that no valid values can be obtained.
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Daily deduction
	// Note: This field may return null, indicating that no valid values can be obtained.
	RICost *string `json:"RICost,omitnil,omitempty" name:"RICost"`

	// Total amount
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`
}

type BusinessSummaryOverviewItem struct {
	// Product code
	// Note: This field may return null, indicating that no valid values can be obtained.
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// Product name
	// Note: This field may return null, indicating that no valid values can be obtained.
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// List price accurate down to eight decimal places
	// Note: This field may return null, indicating that no valid values can be obtained.
	OriginalCost *string `json:"OriginalCost,omitnil,omitempty" name:"OriginalCost"`

	// The deducted voucher amount accurate down to eight decimal places
	// Note: This field may return null, indicating that no valid values can be obtained.
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Consumption amount accurate down to eight decimal places
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`
}

type CountryCodeItem struct {
	// Country/region name in English
	EnName *string `json:"EnName,omitnil,omitempty" name:"EnName"`

	// Country/region name in Chinese
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`


	IOS2 *string `json:"IOS2,omitnil,omitempty" name:"IOS2"`


	IOS3 *string `json:"IOS3,omitnil,omitempty" name:"IOS3"`

	// International dialing code
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`
}

// Predefined struct for user
type CreateAccountRequestParams struct {
	// Account type of a new customer. Valid values: `personal`, `company`.
	AccountType *string `json:"AccountType,omitnil,omitempty" name:"AccountType"`

	// Registered email address, which should be valid and correct.
	// For example, account@qq.com.
	Mail *string `json:"Mail,omitnil,omitempty" name:"Mail"`

	// Account password
	// Length limit: 8-20 characters
	// A password must contain numbers, letters, and symbols (!@#$%^&*()). Space is not allowed.
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// The confirmed password, which must be the same as that entered in the `Password` field.
	ConfirmPassword *string `json:"ConfirmPassword,omitnil,omitempty" name:"ConfirmPassword"`

	// Customer mobile number, which should be valid and correct.
	// A global mobile number within 1-32 digits is allowed, such as 18888888888.
	PhoneNum *string `json:"PhoneNum,omitnil,omitempty" name:"PhoneNum"`

	// Customer's country/region code, which can be obtained via the `GetCountryCodes` API, such as "852".
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// Customer's ISO2 standard country/region code, which can be obtained via the `GetCountryCodes` API. It should correspond to the `CountryCode` field, such as `HK`.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// Extension field, which is left empty by default.
	Extended *string `json:"Extended,omitnil,omitempty" name:"Extended"`
}

type CreateAccountRequest struct {
	*tchttp.BaseRequest
	
	// Account type of a new customer. Valid values: `personal`, `company`.
	AccountType *string `json:"AccountType,omitnil,omitempty" name:"AccountType"`

	// Registered email address, which should be valid and correct.
	// For example, account@qq.com.
	Mail *string `json:"Mail,omitnil,omitempty" name:"Mail"`

	// Account password
	// Length limit: 8-20 characters
	// A password must contain numbers, letters, and symbols (!@#$%^&*()). Space is not allowed.
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// The confirmed password, which must be the same as that entered in the `Password` field.
	ConfirmPassword *string `json:"ConfirmPassword,omitnil,omitempty" name:"ConfirmPassword"`

	// Customer mobile number, which should be valid and correct.
	// A global mobile number within 1-32 digits is allowed, such as 18888888888.
	PhoneNum *string `json:"PhoneNum,omitnil,omitempty" name:"PhoneNum"`

	// Customer's country/region code, which can be obtained via the `GetCountryCodes` API, such as "852".
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// Customer's ISO2 standard country/region code, which can be obtained via the `GetCountryCodes` API. It should correspond to the `CountryCode` field, such as `HK`.
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// Extension field, which is left empty by default.
	Extended *string `json:"Extended,omitnil,omitempty" name:"Extended"`
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
	delete(f, "AccountType")
	delete(f, "Mail")
	delete(f, "Password")
	delete(f, "ConfirmPassword")
	delete(f, "PhoneNum")
	delete(f, "CountryCode")
	delete(f, "Area")
	delete(f, "Extended")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccountResponseParams struct {
	// Account UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type CustomerBillDetailData struct {
	// Reseller account
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayerAccountId *int64 `json:"PayerAccountId,omitnil,omitempty" name:"PayerAccountId"`

	// Customer account
	// Note: This field may return null, indicating that no valid values can be obtained.
	OwnerAccountId *int64 `json:"OwnerAccountId,omitnil,omitempty" name:"OwnerAccountId"`

	// Operator account
	// Note: This field may return null, indicating that no valid values can be obtained.
	OperatorAccountId *int64 `json:"OperatorAccountId,omitnil,omitempty" name:"OperatorAccountId"`

	// Product name
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// Billing mode
	// `Monthly subscription` (Monthly subscription)
	// `Pay-As-You-Go resources` (Pay-as-you-go)
	// `Standard RI` (Reserved instance)
	// Note: This field may return null, indicating that no valid values can be obtained.
	BillingMode *string `json:"BillingMode,omitnil,omitempty" name:"BillingMode"`

	// Project name
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// Resource region
	// Note: This field may return null, indicating that no valid values can be obtained.
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Resource AZ
	// Note: This field may return null, indicating that no valid values can be obtained.
	AvailabilityZone *string `json:"AvailabilityZone,omitnil,omitempty" name:"AvailabilityZone"`

	// Instance ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Instance name
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Subproduct name
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubProductName *string `json:"SubProductName,omitnil,omitempty" name:"SubProductName"`

	// Settlement type
	// Note: This field may return null, indicating that no valid values can be obtained.
	TransactionType *string `json:"TransactionType,omitnil,omitempty" name:"TransactionType"`

	// Transaction ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	TransactionId *string `json:"TransactionId,omitnil,omitempty" name:"TransactionId"`

	// Settlement time
	// 
	// Note: This field may return null, indicating that no valid values can be obtained.
	TransactionTime *string `json:"TransactionTime,omitnil,omitempty" name:"TransactionTime"`

	// Start time of resource use
	// Note: This field may return null, indicating that no valid values can be obtained.
	UsageStartTime *string `json:"UsageStartTime,omitnil,omitempty" name:"UsageStartTime"`

	// End time of resource use
	// Note: This field may return null, indicating that no valid values can be obtained.
	UsageEndTime *string `json:"UsageEndTime,omitnil,omitempty" name:"UsageEndTime"`

	// Component
	// Note: This field may return null, indicating that no valid values can be obtained.
	ComponentType *string `json:"ComponentType,omitnil,omitempty" name:"ComponentType"`

	// Component name
	// Note: This field may return null, indicating that no valid values can be obtained.
	ComponentName *string `json:"ComponentName,omitnil,omitempty" name:"ComponentName"`

	// Component list price
	// Note: This field may return null, indicating that no valid values can be obtained.
	ComponentListPrice *string `json:"ComponentListPrice,omitnil,omitempty" name:"ComponentListPrice"`

	// Price unit
	// Note: This field may return null, indicating that no valid values can be obtained.
	ComponentPriceMeasurementUnit *string `json:"ComponentPriceMeasurementUnit,omitnil,omitempty" name:"ComponentPriceMeasurementUnit"`

	// Component usage
	// Note: This field may return null, indicating that no valid values can be obtained.
	ComponentUsage *string `json:"ComponentUsage,omitnil,omitempty" name:"ComponentUsage"`

	// Component usage unit
	// Note: This field may return null, indicating that no valid values can be obtained.
	ComponentUsageUnit *string `json:"ComponentUsageUnit,omitnil,omitempty" name:"ComponentUsageUnit"`

	// Resource usage duration
	// Note: This field may return null, indicating that no valid values can be obtained.
	UsageDuration *string `json:"UsageDuration,omitnil,omitempty" name:"UsageDuration"`

	// Duration unit
	// Note: This field may return null, indicating that no valid values can be obtained.
	DurationUnit *string `json:"DurationUnit,omitnil,omitempty" name:"DurationUnit"`

	// Original cost
	// Original cost = component list price * component usage * usage duration
	// Note: This field may return null, indicating that no valid values can be obtained.
	OriginalCost *string `json:"OriginalCost,omitnil,omitempty" name:"OriginalCost"`

	// Currency
	// Note: This field may return null, indicating that no valid values can be obtained.
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// Total cost = discounted total - voucher deduction
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// ID
	// Note: The return value may be null, indicating that no valid data can be obtained.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Tag informationNote: This field may return null, indicating that no valid values can be obtained.
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`
}

// Predefined struct for user
type DescribeBillDetailRequestParams struct {
	// The queried month in the format of “YYYY-MM”, such as 2023-01.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Page parameter: Indicates the number of entries per page. Value range: [1, 200]
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Page parameter: Indicates the current page number. The minimum value is 1.
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// Billing mode. Valid values: `prePay` (Monthly subscription), postPay` (Pay-As-You-Go resources).
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Transaction type. Valid values: `prepay_purchase` (Purchase), `prepay_renew` (Renewal), `prepay_modify` (Upgrade/Downgrade), `prepay_return` ( Monthly subscription refund), `postpay_deduct` (Pay-as-you-go), `postpay_deduct_h` (Hourly settlement), `postpay_deduct_d` (Daily settlement), `postpay_deduct_m` (Monthly settlement), `offline_deduct` (Offline project deduction), `online_deduct` (Offline product deduction), `recon_deduct` (Adjustment - deduction), `recon_increase` (Adjustment - compensation), `ripay_purchase` (One-off RI Fee), `postpay_deduct_s` (Spot), `ri_hour_pay` (Hourly RI fee), `prePurchase` (New monthly subscription), `preRenew` (Monthly subscription renewal), `preUpgrade` (Upgrade/Downgrade), `preDowngrade` (Upgrade/Downgrade), `svp_hour_pay` (Hourly Savings Plan fee), `recon_guarantee` (Minimum spend deduction), `pre_purchase` (New monthly subscription), `pre_renew` (Monthly subscription renewal), `pre_upgrade` (Upgrade/Downgrade), `pre_downgrade` (Upgrade/Downgrade).
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`
}

type DescribeBillDetailRequest struct {
	*tchttp.BaseRequest
	
	// The queried month in the format of “YYYY-MM”, such as 2023-01.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Page parameter: Indicates the number of entries per page. Value range: [1, 200]
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Page parameter: Indicates the current page number. The minimum value is 1.
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// Billing mode. Valid values: `prePay` (Monthly subscription), postPay` (Pay-As-You-Go resources).
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Transaction type. Valid values: `prepay_purchase` (Purchase), `prepay_renew` (Renewal), `prepay_modify` (Upgrade/Downgrade), `prepay_return` ( Monthly subscription refund), `postpay_deduct` (Pay-as-you-go), `postpay_deduct_h` (Hourly settlement), `postpay_deduct_d` (Daily settlement), `postpay_deduct_m` (Monthly settlement), `offline_deduct` (Offline project deduction), `online_deduct` (Offline product deduction), `recon_deduct` (Adjustment - deduction), `recon_increase` (Adjustment - compensation), `ripay_purchase` (One-off RI Fee), `postpay_deduct_s` (Spot), `ri_hour_pay` (Hourly RI fee), `prePurchase` (New monthly subscription), `preRenew` (Monthly subscription renewal), `preUpgrade` (Upgrade/Downgrade), `preDowngrade` (Upgrade/Downgrade), `svp_hour_pay` (Hourly Savings Plan fee), `recon_guarantee` (Minimum spend deduction), `pre_purchase` (New monthly subscription), `pre_renew` (Monthly subscription renewal), `pre_upgrade` (Upgrade/Downgrade), `pre_downgrade` (Upgrade/Downgrade).
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`
}

func (r *DescribeBillDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Month")
	delete(f, "PageSize")
	delete(f, "Page")
	delete(f, "PayMode")
	delete(f, "ActionType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillDetailResponseParams struct {
	// Data details
	// Note: This field may return null, indicating that no valid values can be obtained.
	DetailSet []*CustomerBillDetailData `json:"DetailSet,omitnil,omitempty" name:"DetailSet"`

	// Total number of data entries
	// Note: This field may return null, indicating that no valid values can be obtained.
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBillDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillDetailResponseParams `json:"Response"`
}

func (r *DescribeBillDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillDownloadUrlRequestParams struct {
	// Bill month in the format of "yyyy-mm"; the earliest month available for query is June, 2022. Current month's billing data may be inaccurate; please download the current month's bill after it is generated at 1:00 on the 5th of the next month.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Type of bill. Valid values: L2 or L3
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`
}

type DescribeBillDownloadUrlRequest struct {
	*tchttp.BaseRequest
	
	// Bill month in the format of "yyyy-mm"; the earliest month available for query is June, 2022. Current month's billing data may be inaccurate; please download the current month's bill after it is generated at 1:00 on the 5th of the next month.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Type of bill. Valid values: L2 or L3
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`
}

func (r *DescribeBillDownloadUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillDownloadUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Month")
	delete(f, "FileType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillDownloadUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillDownloadUrlResponseParams struct {
	// File download address, valid for one hour.
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBillDownloadUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillDownloadUrlResponseParams `json:"Response"`
}

func (r *DescribeBillDownloadUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillDownloadUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryByPayModeRequestParams struct {
	// Bill month in the format of "yyyy-MM"
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`

	// Customer UIN
	CustomerUin *int64 `json:"CustomerUin,omitnil,omitempty" name:"CustomerUin"`
}

type DescribeBillSummaryByPayModeRequest struct {
	*tchttp.BaseRequest
	
	// Bill month in the format of "yyyy-MM"
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`

	// Customer UIN
	CustomerUin *int64 `json:"CustomerUin,omitnil,omitempty" name:"CustomerUin"`
}

func (r *DescribeBillSummaryByPayModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryByPayModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BillMonth")
	delete(f, "CustomerUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillSummaryByPayModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryByPayModeResponseParams struct {
	// Payment mode details in the customer bill data totaled by payment mode
	// Note: This field may return null, indicating that no valid values can be obtained.
	SummaryOverview []*PayModeSummaryOverviewItem `json:"SummaryOverview,omitnil,omitempty" name:"SummaryOverview"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBillSummaryByPayModeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillSummaryByPayModeResponseParams `json:"Response"`
}

func (r *DescribeBillSummaryByPayModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryByPayModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryByProductRequestParams struct {
	// Bill month in the format of "yyyy-MM"
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`

	// Customer UIN
	CustomerUin *int64 `json:"CustomerUin,omitnil,omitempty" name:"CustomerUin"`
}

type DescribeBillSummaryByProductRequest struct {
	*tchttp.BaseRequest
	
	// Bill month in the format of "yyyy-MM"
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`

	// Customer UIN
	CustomerUin *int64 `json:"CustomerUin,omitnil,omitempty" name:"CustomerUin"`
}

func (r *DescribeBillSummaryByProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryByProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BillMonth")
	delete(f, "CustomerUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillSummaryByProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryByProductResponseParams struct {
	// Bill details from the product dimension
	// Note: This field may return null, indicating that no valid values can be obtained.
	SummaryOverview []*BusinessSummaryOverviewItem `json:"SummaryOverview,omitnil,omitempty" name:"SummaryOverview"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBillSummaryByProductResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillSummaryByProductResponseParams `json:"Response"`
}

func (r *DescribeBillSummaryByProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryByProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryByRegionRequestParams struct {
	// Bill month in the format of "yyyy-MM"
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`

	// Customer UIN
	CustomerUin *int64 `json:"CustomerUin,omitnil,omitempty" name:"CustomerUin"`
}

type DescribeBillSummaryByRegionRequest struct {
	*tchttp.BaseRequest
	
	// Bill month in the format of "yyyy-MM"
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`

	// Customer UIN
	CustomerUin *int64 `json:"CustomerUin,omitnil,omitempty" name:"CustomerUin"`
}

func (r *DescribeBillSummaryByRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryByRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BillMonth")
	delete(f, "CustomerUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillSummaryByRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryByRegionResponseParams struct {
	// Region details in the customer bill data totaled by region
	// Note: This field may return null, indicating that no valid values can be obtained.
	SummaryOverview []*RegionSummaryOverviewItem `json:"SummaryOverview,omitnil,omitempty" name:"SummaryOverview"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBillSummaryByRegionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillSummaryByRegionResponseParams `json:"Response"`
}

func (r *DescribeBillSummaryByRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryByRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryRequestParams struct {
	// Bill month in the format of "yyyy-mm".
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Billing dimension. Optional parameters: product, project, tag
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// Tag value list
	TagKey []*string `json:"TagKey,omitnil,omitempty" name:"TagKey"`
}

type DescribeBillSummaryRequest struct {
	*tchttp.BaseRequest
	
	// Bill month in the format of "yyyy-mm".
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Billing dimension. Optional parameters: product, project, tag
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// Tag value list
	TagKey []*string `json:"TagKey,omitnil,omitempty" name:"TagKey"`
}

func (r *DescribeBillSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Month")
	delete(f, "GroupType")
	delete(f, "TagKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryResponseParams struct {
	// Detailed summary by billing dimensionNote: This field may return null, indicating that no valid values can be obtained.
	SummaryDetail []*SummaryDetails `json:"SummaryDetail,omitnil,omitempty" name:"SummaryDetail"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBillSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillSummaryResponseParams `json:"Response"`
}

func (r *DescribeBillSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomerBillDetailRequestParams struct {
	// Customer UIN
	CustomerUin *uint64 `json:"CustomerUin,omitnil,omitempty" name:"CustomerUin"`

	// The queried month in “YYYY-MM” format, such as 2023-01.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Page parameter: Indicates the number of entries per page. Value range: [1, 200]
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Page parameter: Indicates the current page number. The minimum value is 1.
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// Billing mode. Valid values:
	// `prePay` (Monthly subscription)
	// `postPay` (Pay-as-you-go)
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Transaction type. Valid values:
	// `prepay_purchase` (Purchase)
	// `prepay_renew` (Renewal)
	// `prepay_modify` (Upgrade/Downgrade)
	// `prepay_return` ( Monthly subscription refund)
	// `postpay_deduct` (Pay-as-you-go)
	// `postpay_deduct_h` (Hourly settlement)
	// `postpay_deduct_d` (Daily settlement)
	// `postpay_deduct_m` (Monthly settlement)
	// `offline_deduct` (Offline project deduction)
	// `online_deduct` (Offline product deduction)
	// `recon_deduct` (Adjustment - deduction)
	// `recon_increase` (Adjustment - compensation)
	// `ripay_purchase` (One-off RI Fee)
	// `postpay_deduct_s` (Spot)
	// `ri_hour_pay` (Hourly RI fee)
	// `prePurchase` (New monthly subscription)
	// `preRenew` (Monthly subscription renewal)
	// `preUpgrade` (Upgrade/Downgrade)
	// `preDowngrade` (Upgrade/Downgrade)
	// `svp_hour_pay` (Hourly Savings Plan fee)
	// `recon_guarantee` (Minimum spend deduction)
	// `pre_purchase` (New monthly subscription)
	// `pre_renew` (Monthly subscription renewal)
	// `pre_upgrade` (Upgrade/Downgrade)
	// `pre_downgrade` (Upgrade/Downgrade)
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Payment status
	// `0`: N/A
	// `1`: Paid
	// `2`: Unpaid
	IsConfirmed *string `json:"IsConfirmed,omitnil,omitempty" name:"IsConfirmed"`
}

type DescribeCustomerBillDetailRequest struct {
	*tchttp.BaseRequest
	
	// Customer UIN
	CustomerUin *uint64 `json:"CustomerUin,omitnil,omitempty" name:"CustomerUin"`

	// The queried month in “YYYY-MM” format, such as 2023-01.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Page parameter: Indicates the number of entries per page. Value range: [1, 200]
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// Page parameter: Indicates the current page number. The minimum value is 1.
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// Billing mode. Valid values:
	// `prePay` (Monthly subscription)
	// `postPay` (Pay-as-you-go)
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Transaction type. Valid values:
	// `prepay_purchase` (Purchase)
	// `prepay_renew` (Renewal)
	// `prepay_modify` (Upgrade/Downgrade)
	// `prepay_return` ( Monthly subscription refund)
	// `postpay_deduct` (Pay-as-you-go)
	// `postpay_deduct_h` (Hourly settlement)
	// `postpay_deduct_d` (Daily settlement)
	// `postpay_deduct_m` (Monthly settlement)
	// `offline_deduct` (Offline project deduction)
	// `online_deduct` (Offline product deduction)
	// `recon_deduct` (Adjustment - deduction)
	// `recon_increase` (Adjustment - compensation)
	// `ripay_purchase` (One-off RI Fee)
	// `postpay_deduct_s` (Spot)
	// `ri_hour_pay` (Hourly RI fee)
	// `prePurchase` (New monthly subscription)
	// `preRenew` (Monthly subscription renewal)
	// `preUpgrade` (Upgrade/Downgrade)
	// `preDowngrade` (Upgrade/Downgrade)
	// `svp_hour_pay` (Hourly Savings Plan fee)
	// `recon_guarantee` (Minimum spend deduction)
	// `pre_purchase` (New monthly subscription)
	// `pre_renew` (Monthly subscription renewal)
	// `pre_upgrade` (Upgrade/Downgrade)
	// `pre_downgrade` (Upgrade/Downgrade)
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Payment status
	// `0`: N/A
	// `1`: Paid
	// `2`: Unpaid
	IsConfirmed *string `json:"IsConfirmed,omitnil,omitempty" name:"IsConfirmed"`
}

func (r *DescribeCustomerBillDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomerBillDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerUin")
	delete(f, "Month")
	delete(f, "PageSize")
	delete(f, "Page")
	delete(f, "PayMode")
	delete(f, "ActionType")
	delete(f, "IsConfirmed")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomerBillDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomerBillDetailResponseParams struct {
	// Total number of data entries
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// Data details
	// Note: This field may return null, indicating that no valid values can be obtained.
	DetailSet []*BillDetailData `json:"DetailSet,omitnil,omitempty" name:"DetailSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCustomerBillDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomerBillDetailResponseParams `json:"Response"`
}

func (r *DescribeCustomerBillDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomerBillDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomerBillSummaryRequestParams struct {
	// Customer UIN
	CustomerUin *uint64 `json:"CustomerUin,omitnil,omitempty" name:"CustomerUin"`

	// The queried month in “YYYY-MM” format, such as 2023-01.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Billing mode. Valid values:
	// `prePay` (Monthly subscription)
	// `postPay` (Pay-as-you-go)
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Transaction type. Valid values:
	// `prepay_purchase` (Purchase)
	// `prepay_renew` (Renewal)
	// `prepay_modify` (Upgrade/Downgrade)
	// `prepay_return` (Monthly subscription refund)
	// `postpay_deduct` (Pay-as-you-go)
	// `postpay_deduct_h` (Hourly settlement)
	// `postpay_deduct_d` (Daily settlement)
	// `postpay_deduct_m` (Monthly settlement)
	// `offline_deduct` (Offline project deduction)
	// `online_deduct` (Offline product deduction)
	// `recon_deduct` (Adjustment - deduction)
	// `recon_increase` (Adjustment - compensation)
	// `ripay_purchase` (One-off RI Fee)
	// `postpay_deduct_s` (Spot)
	// `ri_hour_pay` (Hourly RI fee)
	// `prePurchase` (New monthly subscription)
	// `preRenew` (Monthly subscription renewal)
	// `preUpgrade` (Upgrade/Downgrade)
	// `preDowngrade` (Upgrade/Downgrade)
	// `svp_hour_pay` (Hourly Savings Plan fee)
	// `recon_guarantee` (Minimum spend deduction)
	// `pre_purchase` (New monthly subscription)
	// `pre_renew` (Monthly subscription renewal)
	// `pre_upgrade` (Upgrade/Downgrade)
	// `pre_downgrade` (Upgrade/Downgrade)
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Payment status
	// `0`: N/A
	// `1`: Paid
	// `2`: Unpaid
	IsConfirmed *string `json:"IsConfirmed,omitnil,omitempty" name:"IsConfirmed"`
}

type DescribeCustomerBillSummaryRequest struct {
	*tchttp.BaseRequest
	
	// Customer UIN
	CustomerUin *uint64 `json:"CustomerUin,omitnil,omitempty" name:"CustomerUin"`

	// The queried month in “YYYY-MM” format, such as 2023-01.
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// Billing mode. Valid values:
	// `prePay` (Monthly subscription)
	// `postPay` (Pay-as-you-go)
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Transaction type. Valid values:
	// `prepay_purchase` (Purchase)
	// `prepay_renew` (Renewal)
	// `prepay_modify` (Upgrade/Downgrade)
	// `prepay_return` (Monthly subscription refund)
	// `postpay_deduct` (Pay-as-you-go)
	// `postpay_deduct_h` (Hourly settlement)
	// `postpay_deduct_d` (Daily settlement)
	// `postpay_deduct_m` (Monthly settlement)
	// `offline_deduct` (Offline project deduction)
	// `online_deduct` (Offline product deduction)
	// `recon_deduct` (Adjustment - deduction)
	// `recon_increase` (Adjustment - compensation)
	// `ripay_purchase` (One-off RI Fee)
	// `postpay_deduct_s` (Spot)
	// `ri_hour_pay` (Hourly RI fee)
	// `prePurchase` (New monthly subscription)
	// `preRenew` (Monthly subscription renewal)
	// `preUpgrade` (Upgrade/Downgrade)
	// `preDowngrade` (Upgrade/Downgrade)
	// `svp_hour_pay` (Hourly Savings Plan fee)
	// `recon_guarantee` (Minimum spend deduction)
	// `pre_purchase` (New monthly subscription)
	// `pre_renew` (Monthly subscription renewal)
	// `pre_upgrade` (Upgrade/Downgrade)
	// `pre_downgrade` (Upgrade/Downgrade)
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// Payment status
	// `0`: N/A
	// `1`: Paid
	// `2`: Unpaid
	IsConfirmed *string `json:"IsConfirmed,omitnil,omitempty" name:"IsConfirmed"`
}

func (r *DescribeCustomerBillSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomerBillSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerUin")
	delete(f, "Month")
	delete(f, "PayMode")
	delete(f, "ActionType")
	delete(f, "IsConfirmed")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomerBillSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomerBillSummaryResponseParams struct {
	// Total amount
	TotalCost *float64 `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCustomerBillSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomerBillSummaryResponseParams `json:"Response"`
}

func (r *DescribeCustomerBillSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomerBillSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomerInfoData struct {
	// Customer UIN Note: This field may return null, indicating that no valid values can be obtained.
	CustomerUin *string `json:"CustomerUin,omitnil,omitempty" name:"CustomerUin"`

	// Email Note: This field may return null, indicating that no valid values can be obtained.
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// Mobile number Note: This field may return null, indicating that no valid values can be obtained.
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// Remarks Note: This field may return null, indicating that no valid values can be obtained.
	Mark *string `json:"Mark,omitnil,omitempty" name:"Mark"`

	// Displayed name Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Binding time Note: This field may return null, indicating that no valid values can be obtained.
	BindTime *string `json:"BindTime,omitnil,omitempty" name:"BindTime"`

	// Account status
	// 0: Normal
	// 1: Forcibly mandatory (this function is not supported yet)
	// 2. Mandatory arrears
	// Note: The return value may be null, indicating that no valid data can be obtained.
	AccountStatus *string `json:"AccountStatus,omitnil,omitempty" name:"AccountStatus"`

	// Identity verification status
	// -1: Files not uploaded
	// 0: Not submitted for review
	// 1: Under review
	// 2: Review error
	// 3: Approved
	// Note: The return value may be null, indicating that no valid data can be obtained.
	AuthStatus *string `json:"AuthStatus,omitnil,omitempty" name:"AuthStatus"`
}

// Predefined struct for user
type DescribeCustomerInfoRequestParams struct {
	// List of customer UIN. Array length value: 1-20.
	CustomerUin []*int64 `json:"CustomerUin,omitnil,omitempty" name:"CustomerUin"`
}

type DescribeCustomerInfoRequest struct {
	*tchttp.BaseRequest
	
	// List of customer UIN. Array length value: 1-20.
	CustomerUin []*int64 `json:"CustomerUin,omitnil,omitempty" name:"CustomerUin"`
}

func (r *DescribeCustomerInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomerInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomerInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomerInfoResponseParams struct {
	// Customer information Note: This field may return null, indicating that no valid values can be obtained.
	Data []*DescribeCustomerInfoData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCustomerInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomerInfoResponseParams `json:"Response"`
}

func (r *DescribeCustomerInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomerInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomerUinData struct {
	// Customer UIN Note: This field may return null, indicating that no valid values can be obtained.
	CustomerUin *string `json:"CustomerUin,omitnil,omitempty" name:"CustomerUin"`
}

// Predefined struct for user
type DescribeCustomerUinRequestParams struct {
	// Page number
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// Number of data entries per page
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeCustomerUinRequest struct {
	*tchttp.BaseRequest
	
	// Page number
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// Number of data entries per page
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeCustomerUinRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomerUinRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Page")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomerUinRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomerUinResponseParams struct {
	// List of customer UINs Note: This field may return null, indicating that no valid values can be obtained.
	Data []*DescribeCustomerUinData `json:"Data,omitnil,omitempty" name:"Data"`

	// The number of customers
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCustomerUinResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomerUinResponseParams `json:"Response"`
}

func (r *DescribeCustomerUinResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomerUinResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCountryCodesRequestParams struct {

}

type GetCountryCodesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetCountryCodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCountryCodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCountryCodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCountryCodesResponseParams struct {
	// List of country/region codes
	Data []*CountryCodeItem `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetCountryCodesResponse struct {
	*tchttp.BaseResponse
	Response *GetCountryCodesResponseParams `json:"Response"`
}

func (r *GetCountryCodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCountryCodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClientRemarkRequestParams struct {
	// Customer UIN
	ClientUin *string `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`

	// New customer remarks
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ModifyClientRemarkRequest struct {
	*tchttp.BaseRequest
	
	// Customer UIN
	ClientUin *string `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`

	// New customer remarks
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *ModifyClientRemarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClientRemarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClientUin")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClientRemarkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClientRemarkResponseParams struct {
	// If successful, returns the new customer remarks
	ClientRemark *string `json:"ClientRemark,omitnil,omitempty" name:"ClientRemark"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClientRemarkResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClientRemarkResponseParams `json:"Response"`
}

func (r *ModifyClientRemarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClientRemarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PayModeSummaryOverviewItem struct {
	// Billing mode
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// Billing mode name
	// Note: This field may return null, indicating that no valid values can be obtained.
	PayModeName *string `json:"PayModeName,omitnil,omitempty" name:"PayModeName"`

	// The actual total consumption amount accurate down to eight decimal places
	// Note: This field may return null, indicating that no valid values can be obtained.
	OriginalCost *string `json:"OriginalCost,omitnil,omitempty" name:"OriginalCost"`

	// Bill details in each payment mode
	// Note: This field may return null, indicating that no valid values can be obtained.
	Detail []*ActionSummaryOverviewItem `json:"Detail,omitnil,omitempty" name:"Detail"`

	// The deducted voucher amount accurate down to eight decimal places
	// Note: This field may return null, indicating that no valid values can be obtained.
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Total consumption amount accurate down to eight decimal places
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`
}

// Predefined struct for user
type QueryAccountVerificationStatusRequestParams struct {
	// Customer UIN
	ClientUin *int64 `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`
}

type QueryAccountVerificationStatusRequest struct {
	*tchttp.BaseRequest
	
	// Customer UIN
	ClientUin *int64 `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`
}

func (r *QueryAccountVerificationStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryAccountVerificationStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClientUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryAccountVerificationStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryAccountVerificationStatusResponseParams struct {
	// Account verification status
	AccountStatus *bool `json:"AccountStatus,omitnil,omitempty" name:"AccountStatus"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryAccountVerificationStatusResponse struct {
	*tchttp.BaseResponse
	Response *QueryAccountVerificationStatusResponseParams `json:"Response"`
}

func (r *QueryAccountVerificationStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryAccountVerificationStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCreditAllocationHistoryData struct {
	// Allocation time
	AllocatedTime *string `json:"AllocatedTime,omitnil,omitempty" name:"AllocatedTime"`

	// Operator
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// Allocated credit value
	Credit *float64 `json:"Credit,omitnil,omitempty" name:"Credit"`

	// The allocated total credit
	AllocatedCredit *float64 `json:"AllocatedCredit,omitnil,omitempty" name:"AllocatedCredit"`

	// Available credits after allocation
	// Note: The return value may be null, indicating that no valid data can be obtained.
	ClientCreditAfter *float64 `json:"ClientCreditAfter,omitnil,omitempty" name:"ClientCreditAfter"`
}

// Predefined struct for user
type QueryCreditAllocationHistoryRequestParams struct {
	// Customer UIN
	ClientUin *uint64 `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`

	// Page number
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// Number of data entries per page
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type QueryCreditAllocationHistoryRequest struct {
	*tchttp.BaseRequest
	
	// Customer UIN
	ClientUin *uint64 `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`

	// Page number
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// Number of data entries per page
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *QueryCreditAllocationHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCreditAllocationHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClientUin")
	delete(f, "Page")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryCreditAllocationHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCreditAllocationHistoryResponseParams struct {
	// Total number of records
	// Note: This field may return null, indicating that no valid values can be obtained.
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// List of record details
	// Note: This field may return null, indicating that no valid values can be obtained.
	History []*QueryCreditAllocationHistoryData `json:"History,omitnil,omitempty" name:"History"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryCreditAllocationHistoryResponse struct {
	*tchttp.BaseResponse
	Response *QueryCreditAllocationHistoryResponseParams `json:"Response"`
}

func (r *QueryCreditAllocationHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCreditAllocationHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCreditByUinListRequestParams struct {
	// List of user. Array length value: 1-50.
	UinList []*uint64 `json:"UinList,omitnil,omitempty" name:"UinList"`
}

type QueryCreditByUinListRequest struct {
	*tchttp.BaseRequest
	
	// List of user. Array length value: 1-50.
	UinList []*uint64 `json:"UinList,omitnil,omitempty" name:"UinList"`
}

func (r *QueryCreditByUinListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCreditByUinListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UinList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryCreditByUinListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCreditByUinListResponseParams struct {
	// User information list
	Data []*QueryDirectCustomersCreditData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryCreditByUinListResponse struct {
	*tchttp.BaseResponse
	Response *QueryCreditByUinListResponseParams `json:"Response"`
}

func (r *QueryCreditByUinListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCreditByUinListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCreditQuotaRequestParams struct {

}

type QueryCreditQuotaRequest struct {
	*tchttp.BaseRequest
	
}

func (r *QueryCreditQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCreditQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryCreditQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCreditQuotaResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryCreditQuotaResponse struct {
	*tchttp.BaseResponse
	Response *QueryCreditQuotaResponseParams `json:"Response"`
}

func (r *QueryCreditQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCreditQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCustomersCreditData struct {
	// Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Type
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Mobile number
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// Email
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// Overdue payment flag
	Arrears *string `json:"Arrears,omitnil,omitempty" name:"Arrears"`

	// Binding time
	AssociationTime *string `json:"AssociationTime,omitnil,omitempty" name:"AssociationTime"`

	// Expiration time
	RecentExpiry *string `json:"RecentExpiry,omitnil,omitempty" name:"RecentExpiry"`

	// Customer UIN
	ClientUin *uint64 `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`

	// Credit allocated to a customer
	Credit *float64 `json:"Credit,omitnil,omitempty" name:"Credit"`

	// The remaining credit of a customer
	RemainingCredit *float64 `json:"RemainingCredit,omitnil,omitempty" name:"RemainingCredit"`

	// `0`: Identity not verified; `1`: Individual identity verified; `2`: Enterprise identity verified.
	IdentifyType *uint64 `json:"IdentifyType,omitnil,omitempty" name:"IdentifyType"`

	// Customer remarks
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Forced status
	Force *int64 `json:"Force,omitnil,omitempty" name:"Force"`
}

// Predefined struct for user
type QueryCustomersCreditRequestParams struct {
	// Search condition type. You can only search by customer ID, name, remarks, or email.
	FilterType *string `json:"FilterType,omitnil,omitempty" name:"FilterType"`

	// Search condition
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// A pagination parameter that specifies the current page number, with a value starting from 1.
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// A pagination parameter that specifies the number of entries per page.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// A sort parameter that specifies the sort order. Valid values: `desc` (descending order), or `asc` (ascending order) based on `AssociationTime`. The value will be `desc` if left empty.
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type QueryCustomersCreditRequest struct {
	*tchttp.BaseRequest
	
	// Search condition type. You can only search by customer ID, name, remarks, or email.
	FilterType *string `json:"FilterType,omitnil,omitempty" name:"FilterType"`

	// Search condition
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// A pagination parameter that specifies the current page number, with a value starting from 1.
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// A pagination parameter that specifies the number of entries per page.
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// A sort parameter that specifies the sort order. Valid values: `desc` (descending order), or `asc` (ascending order) based on `AssociationTime`. The value will be `desc` if left empty.
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

func (r *QueryCustomersCreditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCustomersCreditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FilterType")
	delete(f, "Filter")
	delete(f, "Page")
	delete(f, "PageSize")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryCustomersCreditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCustomersCreditResponseParams struct {
	// The list of queried customers
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*QueryCustomersCreditData `json:"Data,omitnil,omitempty" name:"Data"`

	// Number of customers
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryCustomersCreditResponse struct {
	*tchttp.BaseResponse
	Response *QueryCustomersCreditResponseParams `json:"Response"`
}

func (r *QueryCustomersCreditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCustomersCreditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryDirectCustomersCreditData struct {
	// User UIN
	Uin *uint64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Total credit
	TotalCredit *float64 `json:"TotalCredit,omitnil,omitempty" name:"TotalCredit"`

	// Remaining credit
	RemainingCredit *float64 `json:"RemainingCredit,omitnil,omitempty" name:"RemainingCredit"`
}

// Predefined struct for user
type QueryDirectCustomersCreditRequestParams struct {

}

type QueryDirectCustomersCreditRequest struct {
	*tchttp.BaseRequest
	
}

func (r *QueryDirectCustomersCreditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryDirectCustomersCreditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryDirectCustomersCreditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryDirectCustomersCreditResponseParams struct {
	// Direct customer information list
	Data []*QueryDirectCustomersCreditData `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryDirectCustomersCreditResponse struct {
	*tchttp.BaseResponse
	Response *QueryDirectCustomersCreditResponseParams `json:"Response"`
}

func (r *QueryDirectCustomersCreditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryDirectCustomersCreditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryPartnerCreditRequestParams struct {

}

type QueryPartnerCreditRequest struct {
	*tchttp.BaseRequest
	
}

func (r *QueryPartnerCreditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryPartnerCreditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryPartnerCreditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryPartnerCreditResponseParams struct {
	// Allocated credit
	AllocatedCredit *float64 `json:"AllocatedCredit,omitnil,omitempty" name:"AllocatedCredit"`

	// Total credit
	TotalCredit *float64 `json:"TotalCredit,omitnil,omitempty" name:"TotalCredit"`

	// Remaining credit
	RemainingCredit *float64 `json:"RemainingCredit,omitnil,omitempty" name:"RemainingCredit"`

	// Allocated quota for the client
	CustomerTotalCredit *float64 `json:"CustomerTotalCredit,omitnil,omitempty" name:"CustomerTotalCredit"`

	// Remaining quota for the client
	CustomerRemainingCredit *float64 `json:"CustomerRemainingCredit,omitnil,omitempty" name:"CustomerRemainingCredit"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryPartnerCreditResponse struct {
	*tchttp.BaseResponse
	Response *QueryPartnerCreditResponseParams `json:"Response"`
}

func (r *QueryPartnerCreditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryPartnerCreditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryVoucherAmountByUinItem struct {
	// Customer UIN
	ClientUin *int64 `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`

	// Voucher quota
	TotalAmount *float64 `json:"TotalAmount,omitnil,omitempty" name:"TotalAmount"`

	// Voucher amount
	RemainAmount *float64 `json:"RemainAmount,omitnil,omitempty" name:"RemainAmount"`
}

// Predefined struct for user
type QueryVoucherAmountByUinRequestParams struct {
	// Customer UIN list. Array length value: 1-20.
	ClientUins []*uint64 `json:"ClientUins,omitnil,omitempty" name:"ClientUins"`
}

type QueryVoucherAmountByUinRequest struct {
	*tchttp.BaseRequest
	
	// Customer UIN list. Array length value: 1-20.
	ClientUins []*uint64 `json:"ClientUins,omitnil,omitempty" name:"ClientUins"`
}

func (r *QueryVoucherAmountByUinRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryVoucherAmountByUinRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClientUins")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryVoucherAmountByUinRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryVoucherAmountByUinResponseParams struct {
	// Customer voucher quota information
	Data []*QueryVoucherAmountByUinItem `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryVoucherAmountByUinResponse struct {
	*tchttp.BaseResponse
	Response *QueryVoucherAmountByUinResponseParams `json:"Response"`
}

func (r *QueryVoucherAmountByUinResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryVoucherAmountByUinResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryVoucherListByUinItem struct {
	// Customer UIN
	ClientUin *int64 `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`

	// The total number of vouchers
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Voucher details
	Data []*QueryVoucherListByUinVoucherItem `json:"Data,omitnil,omitempty" name:"Data"`
}

// Predefined struct for user
type QueryVoucherListByUinRequestParams struct {
	// Customer UIN list. Array length value: 1-20.
	ClientUins []*uint64 `json:"ClientUins,omitnil,omitempty" name:"ClientUins"`

	// Voucher status. If this parameter is not passed in, all status will be queried by default. Valid values: `Unused`, `Used`, `Expired`.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type QueryVoucherListByUinRequest struct {
	*tchttp.BaseRequest
	
	// Customer UIN list. Array length value: 1-20.
	ClientUins []*uint64 `json:"ClientUins,omitnil,omitempty" name:"ClientUins"`

	// Voucher status. If this parameter is not passed in, all status will be queried by default. Valid values: `Unused`, `Used`, `Expired`.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *QueryVoucherListByUinRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryVoucherListByUinRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClientUins")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryVoucherListByUinRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryVoucherListByUinResponseParams struct {
	// Customer voucher information
	Data []*QueryVoucherListByUinItem `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryVoucherListByUinResponse struct {
	*tchttp.BaseResponse
	Response *QueryVoucherListByUinResponseParams `json:"Response"`
}

func (r *QueryVoucherListByUinResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryVoucherListByUinResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryVoucherListByUinVoucherItem struct {
	// Voucher ID
	VoucherId *string `json:"VoucherId,omitnil,omitempty" name:"VoucherId"`

	// Voucher status
	VoucherStatus *string `json:"VoucherStatus,omitnil,omitempty" name:"VoucherStatus"`

	// Voucher value
	TotalAmount *float64 `json:"TotalAmount,omitnil,omitempty" name:"TotalAmount"`

	// Balance
	RemainAmount *float64 `json:"RemainAmount,omitnil,omitempty" name:"RemainAmount"`
}

// Predefined struct for user
type QueryVoucherPoolRequestParams struct {

}

type QueryVoucherPoolRequest struct {
	*tchttp.BaseRequest
	
}

func (r *QueryVoucherPoolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryVoucherPoolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryVoucherPoolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryVoucherPoolResponseParams struct {
	// Reseller name
	AgentName *string `json:"AgentName,omitnil,omitempty" name:"AgentName"`

	// Reseller role type (1: Reseller; 2: Distributor; 3: Second-level reseller)
	AccountType *int64 `json:"AccountType,omitnil,omitempty" name:"AccountType"`

	// Total quota
	TotalQuota *float64 `json:"TotalQuota,omitnil,omitempty" name:"TotalQuota"`

	// Remaining quota
	RemainingQuota *float64 `json:"RemainingQuota,omitnil,omitempty" name:"RemainingQuota"`

	// The number of issued vouchers
	IssuedNum *int64 `json:"IssuedNum,omitnil,omitempty" name:"IssuedNum"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryVoucherPoolResponse struct {
	*tchttp.BaseResponse
	Response *QueryVoucherPoolResponseParams `json:"Response"`
}

func (r *QueryVoucherPoolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryVoucherPoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionSummaryOverviewItem struct {
	// Region ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// Region name
	// Note: This field may return null, indicating that no valid values can be obtained.
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// The actual total consumption amount accurate down to eight decimal places
	// Note: This field may return null, indicating that no valid values can be obtained.
	OriginalCost *string `json:"OriginalCost,omitnil,omitempty" name:"OriginalCost"`

	// The deducted voucher amount accurate down to eight decimal places
	// Note: This field may return null, indicating that no valid values can be obtained.
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Total consumption amount accurate down to eight decimal places
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`
}

type SummaryDetails struct {
	// Product information list
	// Note: This field may return null, indicating that no valid values can be obtained.
	Business []*BusinessInfo `json:"Business,omitnil,omitempty" name:"Business"`

	// Original price
	// Note: This field may return null, indicating that no valid values can be obtained.
	OriginalCost *string `json:"OriginalCost,omitnil,omitempty" name:"OriginalCost"`

	// Voucher amount
	// Note: This field may return null, indicating that no valid values can be obtained.
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// Daily deduction
	// Note: This field may return null, indicating that no valid values can be obtained.
	RICost *string `json:"RICost,omitnil,omitempty" name:"RICost"`

	// Total amount
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// Summary key by classification dimension Note: This field may return null, indicating that no valid values can be obtained.
	GroupKey *string `json:"GroupKey,omitnil,omitempty" name:"GroupKey"`

	//  Summary value by classification dimension
	// Note: This field may return null, indicating that no valid values can be obtained.
	GroupValue *string `json:"GroupValue,omitnil,omitempty" name:"GroupValue"`
}

type TagInfo struct {
	// Tag keyNote: This field may return null, indicating that no valid values can be obtained.
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// Tag valueNote: This field may return null, indicating that no valid values can be obtained.
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}