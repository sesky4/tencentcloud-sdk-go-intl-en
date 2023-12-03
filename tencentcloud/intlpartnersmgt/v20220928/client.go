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
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2022-09-28"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewAllocateCustomerCreditRequest() (request *AllocateCustomerCreditRequest) {
    request = &AllocateCustomerCreditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "AllocateCustomerCredit")
    
    
    return
}

func NewAllocateCustomerCreditResponse() (response *AllocateCustomerCreditResponse) {
    response = &AllocateCustomerCreditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AllocateCustomerCredit
// This API is used for a partner to set credit for a customer, such as increasing or lowering the credit and setting it to 0.
//
// 1. The credit is valid permanently and will not be zeroed regularly.
//
// 2. The customer's service will be suspended when its available credit is set to 0, so caution should be exercised with this operation.
//
// 3. To prevent the customer from making new purchases without affecting their use of previously purchased products, the partner can set their available credit to 0 after obtaining the non-stop feature privilege from the channel manager.
//
// 4. The set credit is an increment of the current available credit and cannot exceed the remaining allocable credit. Setting the credit to a negative value indicates that it will be repossessed. The available credit can be set to 0 at the minimum.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_CREDITAMOUNTOUTOFRANGE = "InvalidParameterValue.CreditAmountOutOfRange"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) AllocateCustomerCredit(request *AllocateCustomerCreditRequest) (response *AllocateCustomerCreditResponse, err error) {
    return c.AllocateCustomerCreditWithContext(context.Background(), request)
}

// AllocateCustomerCredit
// This API is used for a partner to set credit for a customer, such as increasing or lowering the credit and setting it to 0.
//
// 1. The credit is valid permanently and will not be zeroed regularly.
//
// 2. The customer's service will be suspended when its available credit is set to 0, so caution should be exercised with this operation.
//
// 3. To prevent the customer from making new purchases without affecting their use of previously purchased products, the partner can set their available credit to 0 after obtaining the non-stop feature privilege from the channel manager.
//
// 4. The set credit is an increment of the current available credit and cannot exceed the remaining allocable credit. Setting the credit to a negative value indicates that it will be repossessed. The available credit can be set to 0 at the minimum.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_CREDITAMOUNTOUTOFRANGE = "InvalidParameterValue.CreditAmountOutOfRange"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) AllocateCustomerCreditWithContext(ctx context.Context, request *AllocateCustomerCreditRequest) (response *AllocateCustomerCreditResponse, err error) {
    if request == nil {
        request = NewAllocateCustomerCreditRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AllocateCustomerCredit require credential")
    }

    request.SetContext(ctx)
    
    response = NewAllocateCustomerCreditResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccountRequest() (request *CreateAccountRequest) {
    request = &CreateAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "CreateAccount")
    
    
    return
}

func NewCreateAccountResponse() (response *CreateAccountResponse) {
    response = &CreateAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAccount
// This API is used to create a Tencent Cloud account on the partner platform for a customer. After registration, the customer will be automatically bound to the partner account.
//
// 
//
// Notes:<br>
//
// 1. The partner should verify the entered email address and mobile number for creating a Tencent Cloud account.<br>
//
// 2. The customer needs to complete personal information after the first login.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_MAILISREGISTERED = "FailedOperation.MailIsRegistered"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ACCOUNTTYPECONTENTINCORRECT = "InvalidParameter.AccountTypeContentIncorrect"
//  INVALIDPARAMETER_AREACONTENTINCORRECT = "InvalidParameter.AreaContentIncorrect"
//  INVALIDPARAMETER_AREAFORMATINCORRECT = "InvalidParameter.AreaFormatIncorrect"
//  INVALIDPARAMETER_CONFIRMPASSWORDCONTENTINCORRECT = "InvalidParameter.ConfirmPasswordContentIncorrect"
//  INVALIDPARAMETER_COUNTRYCODECONTENTINCORRECT = "InvalidParameter.CountryCodeContentIncorrect"
//  INVALIDPARAMETER_COUNTRYCODEFORMATINCORRECT = "InvalidParameter.CountryCodeFormatIncorrect"
//  INVALIDPARAMETER_MAILFORMATINCORRECT = "InvalidParameter.MailFormatIncorrect"
//  INVALIDPARAMETER_PASSWORDCONTENTINCORRECT = "InvalidParameter.PasswordContentIncorrect"
//  INVALIDPARAMETER_PASSWORDFORMATINCORRECT = "InvalidParameter.PasswordFormatIncorrect"
//  INVALIDPARAMETER_PHONENUMFORMATINCORRECT = "InvalidParameter.PhoneNumFormatIncorrect"
//  INVALIDPARAMETERVALUE_ACCOUNTTYPEEMPTY = "InvalidParameterValue.AccountTypeEmpty"
//  INVALIDPARAMETERVALUE_AREAEMPTY = "InvalidParameterValue.AreaEmpty"
//  INVALIDPARAMETERVALUE_COUNTRYCODEEMPTY = "InvalidParameterValue.CountryCodeEmpty"
//  INVALIDPARAMETERVALUE_MAILEMPTY = "InvalidParameterValue.MailEmpty"
//  INVALIDPARAMETERVALUE_PASSWORDEMPTY = "InvalidParameterValue.PasswordEmpty"
//  INVALIDPARAMETERVALUE_PHONENUMEMPTY = "InvalidParameterValue.PhoneNumEmpty"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAccount(request *CreateAccountRequest) (response *CreateAccountResponse, err error) {
    return c.CreateAccountWithContext(context.Background(), request)
}

// CreateAccount
// This API is used to create a Tencent Cloud account on the partner platform for a customer. After registration, the customer will be automatically bound to the partner account.
//
// 
//
// Notes:<br>
//
// 1. The partner should verify the entered email address and mobile number for creating a Tencent Cloud account.<br>
//
// 2. The customer needs to complete personal information after the first login.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_MAILISREGISTERED = "FailedOperation.MailIsRegistered"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ACCOUNTTYPECONTENTINCORRECT = "InvalidParameter.AccountTypeContentIncorrect"
//  INVALIDPARAMETER_AREACONTENTINCORRECT = "InvalidParameter.AreaContentIncorrect"
//  INVALIDPARAMETER_AREAFORMATINCORRECT = "InvalidParameter.AreaFormatIncorrect"
//  INVALIDPARAMETER_CONFIRMPASSWORDCONTENTINCORRECT = "InvalidParameter.ConfirmPasswordContentIncorrect"
//  INVALIDPARAMETER_COUNTRYCODECONTENTINCORRECT = "InvalidParameter.CountryCodeContentIncorrect"
//  INVALIDPARAMETER_COUNTRYCODEFORMATINCORRECT = "InvalidParameter.CountryCodeFormatIncorrect"
//  INVALIDPARAMETER_MAILFORMATINCORRECT = "InvalidParameter.MailFormatIncorrect"
//  INVALIDPARAMETER_PASSWORDCONTENTINCORRECT = "InvalidParameter.PasswordContentIncorrect"
//  INVALIDPARAMETER_PASSWORDFORMATINCORRECT = "InvalidParameter.PasswordFormatIncorrect"
//  INVALIDPARAMETER_PHONENUMFORMATINCORRECT = "InvalidParameter.PhoneNumFormatIncorrect"
//  INVALIDPARAMETERVALUE_ACCOUNTTYPEEMPTY = "InvalidParameterValue.AccountTypeEmpty"
//  INVALIDPARAMETERVALUE_AREAEMPTY = "InvalidParameterValue.AreaEmpty"
//  INVALIDPARAMETERVALUE_COUNTRYCODEEMPTY = "InvalidParameterValue.CountryCodeEmpty"
//  INVALIDPARAMETERVALUE_MAILEMPTY = "InvalidParameterValue.MailEmpty"
//  INVALIDPARAMETERVALUE_PASSWORDEMPTY = "InvalidParameterValue.PasswordEmpty"
//  INVALIDPARAMETERVALUE_PHONENUMEMPTY = "InvalidParameterValue.PhoneNumEmpty"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateAccountWithContext(ctx context.Context, request *CreateAccountRequest) (response *CreateAccountResponse, err error) {
    if request == nil {
        request = NewCreateAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillDetailRequest() (request *DescribeBillDetailRequest) {
    request = &DescribeBillDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "DescribeBillDetail")
    
    
    return
}

func NewDescribeBillDetailResponse() (response *DescribeBillDetailResponse) {
    response = &DescribeBillDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillDetail
// This API is used to query the customer bill details.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDMONTH = "InvalidParameterValue.InvalidMonth"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeBillDetail(request *DescribeBillDetailRequest) (response *DescribeBillDetailResponse, err error) {
    return c.DescribeBillDetailWithContext(context.Background(), request)
}

// DescribeBillDetail
// This API is used to query the customer bill details.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDMONTH = "InvalidParameterValue.InvalidMonth"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeBillDetailWithContext(ctx context.Context, request *DescribeBillDetailRequest) (response *DescribeBillDetailResponse, err error) {
    if request == nil {
        request = NewDescribeBillDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryByPayModeRequest() (request *DescribeBillSummaryByPayModeRequest) {
    request = &DescribeBillSummaryByPayModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "DescribeBillSummaryByPayMode")
    
    
    return
}

func NewDescribeBillSummaryByPayModeResponse() (response *DescribeBillSummaryByPayModeResponse) {
    response = &DescribeBillSummaryByPayModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillSummaryByPayMode
// This API is used to obtain the total amount of customer bills by payment mode.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDMONTH = "InvalidParameterValue.InvalidMonth"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeBillSummaryByPayMode(request *DescribeBillSummaryByPayModeRequest) (response *DescribeBillSummaryByPayModeResponse, err error) {
    return c.DescribeBillSummaryByPayModeWithContext(context.Background(), request)
}

// DescribeBillSummaryByPayMode
// This API is used to obtain the total amount of customer bills by payment mode.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDMONTH = "InvalidParameterValue.InvalidMonth"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeBillSummaryByPayModeWithContext(ctx context.Context, request *DescribeBillSummaryByPayModeRequest) (response *DescribeBillSummaryByPayModeResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByPayModeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillSummaryByPayMode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillSummaryByPayModeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryByProductRequest() (request *DescribeBillSummaryByProductRequest) {
    request = &DescribeBillSummaryByProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "DescribeBillSummaryByProduct")
    
    
    return
}

func NewDescribeBillSummaryByProductResponse() (response *DescribeBillSummaryByProductResponse) {
    response = &DescribeBillSummaryByProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillSummaryByProduct
// This API is used to obtain the total amount of customer bills by product.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDMONTH = "InvalidParameterValue.InvalidMonth"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeBillSummaryByProduct(request *DescribeBillSummaryByProductRequest) (response *DescribeBillSummaryByProductResponse, err error) {
    return c.DescribeBillSummaryByProductWithContext(context.Background(), request)
}

// DescribeBillSummaryByProduct
// This API is used to obtain the total amount of customer bills by product.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDMONTH = "InvalidParameterValue.InvalidMonth"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeBillSummaryByProductWithContext(ctx context.Context, request *DescribeBillSummaryByProductRequest) (response *DescribeBillSummaryByProductResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByProductRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillSummaryByProduct require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillSummaryByProductResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryByRegionRequest() (request *DescribeBillSummaryByRegionRequest) {
    request = &DescribeBillSummaryByRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "DescribeBillSummaryByRegion")
    
    
    return
}

func NewDescribeBillSummaryByRegionResponse() (response *DescribeBillSummaryByRegionResponse) {
    response = &DescribeBillSummaryByRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillSummaryByRegion
// This API is used to obtain the total amount of customer bills by region.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDMONTH = "InvalidParameterValue.InvalidMonth"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeBillSummaryByRegion(request *DescribeBillSummaryByRegionRequest) (response *DescribeBillSummaryByRegionResponse, err error) {
    return c.DescribeBillSummaryByRegionWithContext(context.Background(), request)
}

// DescribeBillSummaryByRegion
// This API is used to obtain the total amount of customer bills by region.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDMONTH = "InvalidParameterValue.InvalidMonth"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeBillSummaryByRegionWithContext(ctx context.Context, request *DescribeBillSummaryByRegionRequest) (response *DescribeBillSummaryByRegionResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByRegionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillSummaryByRegion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillSummaryByRegionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomerBillDetailRequest() (request *DescribeCustomerBillDetailRequest) {
    request = &DescribeCustomerBillDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "DescribeCustomerBillDetail")
    
    
    return
}

func NewDescribeCustomerBillDetailResponse() (response *DescribeCustomerBillDetailResponse) {
    response = &DescribeCustomerBillDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCustomerBillDetail
// This API is used to query the customer bill details.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDMONTH = "InvalidParameterValue.InvalidMonth"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_NOTCUSTOMERUIN = "UnauthorizedOperation.NotCustomerUin"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeCustomerBillDetail(request *DescribeCustomerBillDetailRequest) (response *DescribeCustomerBillDetailResponse, err error) {
    return c.DescribeCustomerBillDetailWithContext(context.Background(), request)
}

// DescribeCustomerBillDetail
// This API is used to query the customer bill details.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDMONTH = "InvalidParameterValue.InvalidMonth"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_NOTCUSTOMERUIN = "UnauthorizedOperation.NotCustomerUin"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeCustomerBillDetailWithContext(ctx context.Context, request *DescribeCustomerBillDetailRequest) (response *DescribeCustomerBillDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCustomerBillDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomerBillDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomerBillDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomerBillSummaryRequest() (request *DescribeCustomerBillSummaryRequest) {
    request = &DescribeCustomerBillSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "DescribeCustomerBillSummary")
    
    
    return
}

func NewDescribeCustomerBillSummaryResponse() (response *DescribeCustomerBillSummaryResponse) {
    response = &DescribeCustomerBillSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCustomerBillSummary
// This API is used to query the total amount of customer bills.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDMONTH = "InvalidParameterValue.InvalidMonth"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeCustomerBillSummary(request *DescribeCustomerBillSummaryRequest) (response *DescribeCustomerBillSummaryResponse, err error) {
    return c.DescribeCustomerBillSummaryWithContext(context.Background(), request)
}

// DescribeCustomerBillSummary
// This API is used to query the total amount of customer bills.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDMONTH = "InvalidParameterValue.InvalidMonth"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeCustomerBillSummaryWithContext(ctx context.Context, request *DescribeCustomerBillSummaryRequest) (response *DescribeCustomerBillSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeCustomerBillSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomerBillSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomerBillSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomerInfoRequest() (request *DescribeCustomerInfoRequest) {
    request = &DescribeCustomerInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "DescribeCustomerInfo")
    
    
    return
}

func NewDescribeCustomerInfoResponse() (response *DescribeCustomerInfoResponse) {
    response = &DescribeCustomerInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCustomerInfo
// This API is used to query the customer information.
//
// error code that may be returned:
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeCustomerInfo(request *DescribeCustomerInfoRequest) (response *DescribeCustomerInfoResponse, err error) {
    return c.DescribeCustomerInfoWithContext(context.Background(), request)
}

// DescribeCustomerInfo
// This API is used to query the customer information.
//
// error code that may be returned:
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeCustomerInfoWithContext(ctx context.Context, request *DescribeCustomerInfoRequest) (response *DescribeCustomerInfoResponse, err error) {
    if request == nil {
        request = NewDescribeCustomerInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomerInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomerInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomerUinRequest() (request *DescribeCustomerUinRequest) {
    request = &DescribeCustomerUinRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "DescribeCustomerUin")
    
    
    return
}

func NewDescribeCustomerUinResponse() (response *DescribeCustomerUinResponse) {
    response = &DescribeCustomerUinResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCustomerUin
// This API is used to query the list of customer UINs.
//
// error code that may be returned:
//  INVALIDPARAMETER_PAGE = "InvalidParameter.Page"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeCustomerUin(request *DescribeCustomerUinRequest) (response *DescribeCustomerUinResponse, err error) {
    return c.DescribeCustomerUinWithContext(context.Background(), request)
}

// DescribeCustomerUin
// This API is used to query the list of customer UINs.
//
// error code that may be returned:
//  INVALIDPARAMETER_PAGE = "InvalidParameter.Page"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) DescribeCustomerUinWithContext(ctx context.Context, request *DescribeCustomerUinRequest) (response *DescribeCustomerUinResponse, err error) {
    if request == nil {
        request = NewDescribeCustomerUinRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomerUin require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomerUinResponse()
    err = c.Send(request, response)
    return
}

func NewGetCountryCodesRequest() (request *GetCountryCodesRequest) {
    request = &GetCountryCodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "GetCountryCodes")
    
    
    return
}

func NewGetCountryCodesResponse() (response *GetCountryCodesResponse) {
    response = &GetCountryCodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetCountryCodes
// This API is used to obtain country/region codes.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) GetCountryCodes(request *GetCountryCodesRequest) (response *GetCountryCodesResponse, err error) {
    return c.GetCountryCodesWithContext(context.Background(), request)
}

// GetCountryCodes
// This API is used to obtain country/region codes.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) GetCountryCodesWithContext(ctx context.Context, request *GetCountryCodesRequest) (response *GetCountryCodesResponse, err error) {
    if request == nil {
        request = NewGetCountryCodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetCountryCodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetCountryCodesResponse()
    err = c.Send(request, response)
    return
}

func NewQueryAccountVerificationStatusRequest() (request *QueryAccountVerificationStatusRequest) {
    request = &QueryAccountVerificationStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "QueryAccountVerificationStatus")
    
    
    return
}

func NewQueryAccountVerificationStatusResponse() (response *QueryAccountVerificationStatusResponse) {
    response = &QueryAccountVerificationStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryAccountVerificationStatus
// This API is used to query the account verification status.
//
// error code that may be returned:
//  FAILEDOPERATION_UININVALID = "FailedOperation.UinInvalid"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_NOTCUSTOMERUIN = "UnauthorizedOperation.NotCustomerUin"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) QueryAccountVerificationStatus(request *QueryAccountVerificationStatusRequest) (response *QueryAccountVerificationStatusResponse, err error) {
    return c.QueryAccountVerificationStatusWithContext(context.Background(), request)
}

// QueryAccountVerificationStatus
// This API is used to query the account verification status.
//
// error code that may be returned:
//  FAILEDOPERATION_UININVALID = "FailedOperation.UinInvalid"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_NOTCUSTOMERUIN = "UnauthorizedOperation.NotCustomerUin"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) QueryAccountVerificationStatusWithContext(ctx context.Context, request *QueryAccountVerificationStatusRequest) (response *QueryAccountVerificationStatusResponse, err error) {
    if request == nil {
        request = NewQueryAccountVerificationStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryAccountVerificationStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryAccountVerificationStatusResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCreditAllocationHistoryRequest() (request *QueryCreditAllocationHistoryRequest) {
    request = &QueryCreditAllocationHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "QueryCreditAllocationHistory")
    
    
    return
}

func NewQueryCreditAllocationHistoryResponse() (response *QueryCreditAllocationHistoryResponse) {
    response = &QueryCreditAllocationHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryCreditAllocationHistory
// This API is used to query all the credit allocation records of a single customer.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) QueryCreditAllocationHistory(request *QueryCreditAllocationHistoryRequest) (response *QueryCreditAllocationHistoryResponse, err error) {
    return c.QueryCreditAllocationHistoryWithContext(context.Background(), request)
}

// QueryCreditAllocationHistory
// This API is used to query all the credit allocation records of a single customer.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) QueryCreditAllocationHistoryWithContext(ctx context.Context, request *QueryCreditAllocationHistoryRequest) (response *QueryCreditAllocationHistoryResponse, err error) {
    if request == nil {
        request = NewQueryCreditAllocationHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryCreditAllocationHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryCreditAllocationHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCreditByUinListRequest() (request *QueryCreditByUinListRequest) {
    request = &QueryCreditByUinListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "QueryCreditByUinList")
    
    
    return
}

func NewQueryCreditByUinListResponse() (response *QueryCreditByUinListResponse) {
    response = &QueryCreditByUinListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryCreditByUinList
// This API is used to query the credit of users in the list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION_NOTCUSTOMERUIN = "UnauthorizedOperation.NotCustomerUin"
func (c *Client) QueryCreditByUinList(request *QueryCreditByUinListRequest) (response *QueryCreditByUinListResponse, err error) {
    return c.QueryCreditByUinListWithContext(context.Background(), request)
}

// QueryCreditByUinList
// This API is used to query the credit of users in the list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION_NOTCUSTOMERUIN = "UnauthorizedOperation.NotCustomerUin"
func (c *Client) QueryCreditByUinListWithContext(ctx context.Context, request *QueryCreditByUinListRequest) (response *QueryCreditByUinListResponse, err error) {
    if request == nil {
        request = NewQueryCreditByUinListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryCreditByUinList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryCreditByUinListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCreditQuotaRequest() (request *QueryCreditQuotaRequest) {
    request = &QueryCreditQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "QueryCreditQuota")
    
    
    return
}

func NewQueryCreditQuotaResponse() (response *QueryCreditQuotaResponse) {
    response = &QueryCreditQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryCreditQuota
// This API is used to query customer credits.
//
// error code that may be returned:
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
func (c *Client) QueryCreditQuota(request *QueryCreditQuotaRequest) (response *QueryCreditQuotaResponse, err error) {
    return c.QueryCreditQuotaWithContext(context.Background(), request)
}

// QueryCreditQuota
// This API is used to query customer credits.
//
// error code that may be returned:
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
func (c *Client) QueryCreditQuotaWithContext(ctx context.Context, request *QueryCreditQuotaRequest) (response *QueryCreditQuotaResponse, err error) {
    if request == nil {
        request = NewQueryCreditQuotaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryCreditQuota require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryCreditQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCustomersCreditRequest() (request *QueryCustomersCreditRequest) {
    request = &QueryCustomersCreditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "QueryCustomersCredit")
    
    
    return
}

func NewQueryCustomersCreditResponse() (response *QueryCustomersCreditResponse) {
    response = &QueryCustomersCreditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryCustomersCredit
// This API is used for a partner to the credits and basic information of cutomers.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) QueryCustomersCredit(request *QueryCustomersCreditRequest) (response *QueryCustomersCreditResponse, err error) {
    return c.QueryCustomersCreditWithContext(context.Background(), request)
}

// QueryCustomersCredit
// This API is used for a partner to the credits and basic information of cutomers.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) QueryCustomersCreditWithContext(ctx context.Context, request *QueryCustomersCreditRequest) (response *QueryCustomersCreditResponse, err error) {
    if request == nil {
        request = NewQueryCustomersCreditRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryCustomersCredit require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryCustomersCreditResponse()
    err = c.Send(request, response)
    return
}

func NewQueryDirectCustomersCreditRequest() (request *QueryDirectCustomersCreditRequest) {
    request = &QueryDirectCustomersCreditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "QueryDirectCustomersCredit")
    
    
    return
}

func NewQueryDirectCustomersCreditResponse() (response *QueryDirectCustomersCreditResponse) {
    response = &QueryDirectCustomersCreditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryDirectCustomersCredit
// This API is used to query the credits of direct customers.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) QueryDirectCustomersCredit(request *QueryDirectCustomersCreditRequest) (response *QueryDirectCustomersCreditResponse, err error) {
    return c.QueryDirectCustomersCreditWithContext(context.Background(), request)
}

// QueryDirectCustomersCredit
// This API is used to query the credits of direct customers.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) QueryDirectCustomersCreditWithContext(ctx context.Context, request *QueryDirectCustomersCreditRequest) (response *QueryDirectCustomersCreditResponse, err error) {
    if request == nil {
        request = NewQueryDirectCustomersCreditRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryDirectCustomersCredit require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryDirectCustomersCreditResponse()
    err = c.Send(request, response)
    return
}

func NewQueryPartnerCreditRequest() (request *QueryPartnerCreditRequest) {
    request = &QueryPartnerCreditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "QueryPartnerCredit")
    
    
    return
}

func NewQueryPartnerCreditResponse() (response *QueryPartnerCreditResponse) {
    response = &QueryPartnerCreditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryPartnerCredit
// This API is used for a partner to query its own total credit, available credit, and used credit in USD.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) QueryPartnerCredit(request *QueryPartnerCreditRequest) (response *QueryPartnerCreditResponse, err error) {
    return c.QueryPartnerCreditWithContext(context.Background(), request)
}

// QueryPartnerCredit
// This API is used for a partner to query its own total credit, available credit, and used credit in USD.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) QueryPartnerCreditWithContext(ctx context.Context, request *QueryPartnerCreditRequest) (response *QueryPartnerCreditResponse, err error) {
    if request == nil {
        request = NewQueryPartnerCreditRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryPartnerCredit require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryPartnerCreditResponse()
    err = c.Send(request, response)
    return
}

func NewQueryVoucherAmountByUinRequest() (request *QueryVoucherAmountByUinRequest) {
    request = &QueryVoucherAmountByUinRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "QueryVoucherAmountByUin")
    
    
    return
}

func NewQueryVoucherAmountByUinResponse() (response *QueryVoucherAmountByUinResponse) {
    response = &QueryVoucherAmountByUinResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryVoucherAmountByUin
// This API is used to query the voucher quota based on the customer UIN.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) QueryVoucherAmountByUin(request *QueryVoucherAmountByUinRequest) (response *QueryVoucherAmountByUinResponse, err error) {
    return c.QueryVoucherAmountByUinWithContext(context.Background(), request)
}

// QueryVoucherAmountByUin
// This API is used to query the voucher quota based on the customer UIN.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) QueryVoucherAmountByUinWithContext(ctx context.Context, request *QueryVoucherAmountByUinRequest) (response *QueryVoucherAmountByUinResponse, err error) {
    if request == nil {
        request = NewQueryVoucherAmountByUinRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryVoucherAmountByUin require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryVoucherAmountByUinResponse()
    err = c.Send(request, response)
    return
}

func NewQueryVoucherListByUinRequest() (request *QueryVoucherListByUinRequest) {
    request = &QueryVoucherListByUinRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "QueryVoucherListByUin")
    
    
    return
}

func NewQueryVoucherListByUinResponse() (response *QueryVoucherListByUinResponse) {
    response = &QueryVoucherListByUinResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryVoucherListByUin
// This API is used to query the voucher list based on the customer UIN.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) QueryVoucherListByUin(request *QueryVoucherListByUinRequest) (response *QueryVoucherListByUinResponse, err error) {
    return c.QueryVoucherListByUinWithContext(context.Background(), request)
}

// QueryVoucherListByUin
// This API is used to query the voucher list based on the customer UIN.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) QueryVoucherListByUinWithContext(ctx context.Context, request *QueryVoucherListByUinRequest) (response *QueryVoucherListByUinResponse, err error) {
    if request == nil {
        request = NewQueryVoucherListByUinRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryVoucherListByUin require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryVoucherListByUinResponse()
    err = c.Send(request, response)
    return
}

func NewQueryVoucherPoolRequest() (request *QueryVoucherPoolRequest) {
    request = &QueryVoucherPoolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("intlpartnersmgt", APIVersion, "QueryVoucherPool")
    
    
    return
}

func NewQueryVoucherPoolResponse() (response *QueryVoucherPoolResponse) {
    response = &QueryVoucherPoolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryVoucherPool
// This API is used to query the voucher quota pool.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) QueryVoucherPool(request *QueryVoucherPoolRequest) (response *QueryVoucherPoolResponse, err error) {
    return c.QueryVoucherPoolWithContext(context.Background(), request)
}

// QueryVoucherPool
// This API is used to query the voucher quota pool.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_SERVICEBUSY = "OperationDenied.ServiceBusy"
//  UNAUTHORIZEDOPERATION_UINNOAUTH = "UnauthorizedOperation.UinNoAuth"
func (c *Client) QueryVoucherPoolWithContext(ctx context.Context, request *QueryVoucherPoolRequest) (response *QueryVoucherPoolResponse, err error) {
    if request == nil {
        request = NewQueryVoucherPoolRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryVoucherPool require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryVoucherPoolResponse()
    err = c.Send(request, response)
    return
}
