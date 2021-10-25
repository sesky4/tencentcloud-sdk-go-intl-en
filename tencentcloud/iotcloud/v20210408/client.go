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

package v20210408

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2021-04-08"

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


func NewCreateDeviceRequest() (request *CreateDeviceRequest) {
    request = &CreateDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "CreateDevice")
    return
}

func NewCreateDeviceResponse() (response *CreateDeviceResponse) {
    response = &CreateDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDevice
// This API is used to create an IoT Hub device. 
//
// error code that may be returned:
//  FAILEDOPERATION_TIDWHITELISTNOTOPEN = "FailedOperation.TidWhiteListNotOpen"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEFINEDPSKNOTBASE64 = "InvalidParameterValue.DefinedPskNotBase64"
//  INVALIDPARAMETERVALUE_DEVICEALREADYEXIST = "InvalidParameterValue.DeviceAlreadyExist"
//  LIMITEXCEEDED_DEVICEEXCEEDLIMIT = "LimitExceeded.DeviceExceedLimit"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_PRODUCTCANTHAVELORADEVICE = "UnauthorizedOperation.ProductCantHaveLoRaDevice"
//  UNAUTHORIZEDOPERATION_PRODUCTCANTHAVENORMALDEVICE = "UnauthorizedOperation.ProductCantHaveNormalDevice"
//  UNAUTHORIZEDOPERATION_PRODUCTCANTHAVENOTLORADEVICE = "UnauthorizedOperation.ProductCantHaveNotLoRaDevice"
//  UNAUTHORIZEDOPERATION_PRODUCTNOTSUPPORTPSK = "UnauthorizedOperation.ProductNotSupportPSK"
//  UNSUPPORTEDOPERATION_SUITETOKENNOCREATE = "UnsupportedOperation.SuiteTokenNoCreate"
func (c *Client) CreateDevice(request *CreateDeviceRequest) (response *CreateDeviceResponse, err error) {
    if request == nil {
        request = NewCreateDeviceRequest()
    }
    response = NewCreateDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrivateCARequest() (request *CreatePrivateCARequest) {
    request = &CreatePrivateCARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "CreatePrivateCA")
    return
}

func NewCreatePrivateCAResponse() (response *CreatePrivateCAResponse) {
    response = &CreatePrivateCAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePrivateCA
// This API is used to create a private CA certificate.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_CACERTINVALID = "InvalidParameterValue.CACertInvalid"
//  INVALIDPARAMETERVALUE_CACERTNOTMATCH = "InvalidParameterValue.CACertNotMatch"
//  LIMITEXCEEDED_CAREPEAT = "LimitExceeded.CARepeat"
func (c *Client) CreatePrivateCA(request *CreatePrivateCARequest) (response *CreatePrivateCAResponse, err error) {
    if request == nil {
        request = NewCreatePrivateCARequest()
    }
    response = NewCreatePrivateCAResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDeviceRequest() (request *DeleteDeviceRequest) {
    request = &DeleteDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DeleteDevice")
    return
}

func NewDeleteDeviceResponse() (response *DeleteDeviceResponse) {
    response = &DeleteDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDevice
// This API is used to delete an IoT Hub device. 
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNAUTHORIZEDOPERATION_DEVICEHASALREADYBINDGATEWAY = "UnauthorizedOperation.DeviceHasAlreadyBindGateway"
//  UNAUTHORIZEDOPERATION_GATEWAYHASBINDEDDEVICES = "UnauthorizedOperation.GatewayHasBindedDevices"
//  UNSUPPORTEDOPERATION_DEVICEOTATASKINPROGRESS = "UnsupportedOperation.DeviceOtaTaskInProgress"
func (c *Client) DeleteDevice(request *DeleteDeviceRequest) (response *DeleteDeviceResponse, err error) {
    if request == nil {
        request = NewDeleteDeviceRequest()
    }
    response = NewDeleteDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrivateCARequest() (request *DeletePrivateCARequest) {
    request = &DeletePrivateCARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DeletePrivateCA")
    return
}

func NewDeletePrivateCAResponse() (response *DeletePrivateCAResponse) {
    response = &DeletePrivateCAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePrivateCA
// This API is used to delete a private CA certificate.
//
// error code that may be returned:
//  LIMITEXCEEDED_CAALREADYBINDPRODUCT = "LimitExceeded.CAAlreadyBindProduct"
//  RESOURCENOTFOUND_CACERTNOTEXIST = "ResourceNotFound.CACertNotExist"
func (c *Client) DeletePrivateCA(request *DeletePrivateCARequest) (response *DeletePrivateCAResponse, err error) {
    if request == nil {
        request = NewDeletePrivateCARequest()
    }
    response = NewDeletePrivateCAResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProductRequest() (request *DeleteProductRequest) {
    request = &DeleteProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DeleteProduct")
    return
}

func NewDeleteProductResponse() (response *DeleteProductResponse) {
    response = &DeleteProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteProduct
// This API is used to delete an IoT Hub product.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_DEVICESEXISTUNDERPRODUCT = "UnauthorizedOperation.DevicesExistUnderProduct"
//  UNSUPPORTEDOPERATION_GATEWAYPRODUCTHASBINDEDPRODUCT = "UnsupportedOperation.GatewayProductHasBindedProduct"
//  UNSUPPORTEDOPERATION_PRODUCTHASBINDGATEWAY = "UnsupportedOperation.ProductHasBindGateway"
//  UNSUPPORTEDOPERATION_PRODUCTHASBINDEDGATEWAYPRODUCT = "UnsupportedOperation.ProductHasBindedGatewayProduct"
func (c *Client) DeleteProduct(request *DeleteProductRequest) (response *DeleteProductResponse, err error) {
    if request == nil {
        request = NewDeleteProductRequest()
    }
    response = NewDeleteProductResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceRequest() (request *DescribeDeviceRequest) {
    request = &DescribeDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeDevice")
    return
}

func NewDescribeDeviceResponse() (response *DescribeDeviceResponse) {
    response = &DescribeDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDevice
// This API is used to query device details.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) DescribeDevice(request *DescribeDeviceRequest) (response *DescribeDeviceResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceRequest()
    }
    response = NewDescribeDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDevicesRequest() (request *DescribeDevicesRequest) {
    request = &DescribeDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeDevices")
    return
}

func NewDescribeDevicesResponse() (response *DescribeDevicesResponse) {
    response = &DescribeDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDevices
// This API is used to get the list of IoT Hub devices. 
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) DescribeDevices(request *DescribeDevicesRequest) (response *DescribeDevicesResponse, err error) {
    if request == nil {
        request = NewDescribeDevicesRequest()
    }
    response = NewDescribeDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrivateCARequest() (request *DescribePrivateCARequest) {
    request = &DescribePrivateCARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribePrivateCA")
    return
}

func NewDescribePrivateCAResponse() (response *DescribePrivateCAResponse) {
    response = &DescribePrivateCAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrivateCA
// This API is used to query private CA certificate details.
//
// error code that may be returned:
//  RESOURCENOTFOUND_CACERTNOTEXIST = "ResourceNotFound.CACertNotExist"
func (c *Client) DescribePrivateCA(request *DescribePrivateCARequest) (response *DescribePrivateCAResponse, err error) {
    if request == nil {
        request = NewDescribePrivateCARequest()
    }
    response = NewDescribePrivateCAResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrivateCABindedProductsRequest() (request *DescribePrivateCABindedProductsRequest) {
    request = &DescribePrivateCABindedProductsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribePrivateCABindedProducts")
    return
}

func NewDescribePrivateCABindedProductsResponse() (response *DescribePrivateCABindedProductsResponse) {
    response = &DescribePrivateCABindedProductsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrivateCABindedProducts
// This API is used to query the products bound to a private CA certificate.
//
// error code that may be returned:
//  RESOURCENOTFOUND_CACERTNOTEXIST = "ResourceNotFound.CACertNotExist"
func (c *Client) DescribePrivateCABindedProducts(request *DescribePrivateCABindedProductsRequest) (response *DescribePrivateCABindedProductsResponse, err error) {
    if request == nil {
        request = NewDescribePrivateCABindedProductsRequest()
    }
    response = NewDescribePrivateCABindedProductsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrivateCAsRequest() (request *DescribePrivateCAsRequest) {
    request = &DescribePrivateCAsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribePrivateCAs")
    return
}

func NewDescribePrivateCAsResponse() (response *DescribePrivateCAsResponse) {
    response = &DescribePrivateCAsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrivateCAs
// This API is used to get the list of private CA certificates.
//
// error code that may be returned:
//  RESOURCENOTFOUND_CACERTNOTEXIST = "ResourceNotFound.CACertNotExist"
func (c *Client) DescribePrivateCAs(request *DescribePrivateCAsRequest) (response *DescribePrivateCAsResponse, err error) {
    if request == nil {
        request = NewDescribePrivateCAsRequest()
    }
    response = NewDescribePrivateCAsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductRequest() (request *DescribeProductRequest) {
    request = &DescribeProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeProduct")
    return
}

func NewDescribeProductResponse() (response *DescribeProductResponse) {
    response = &DescribeProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProduct
// This API is used to query product details.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) DescribeProduct(request *DescribeProductRequest) (response *DescribeProductResponse, err error) {
    if request == nil {
        request = NewDescribeProductRequest()
    }
    response = NewDescribeProductResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductCARequest() (request *DescribeProductCARequest) {
    request = &DescribeProductCARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "DescribeProductCA")
    return
}

func NewDescribeProductCAResponse() (response *DescribeProductCAResponse) {
    response = &DescribeProductCAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProductCA
// This API is used to query the CA certificates bound to a product.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) DescribeProductCA(request *DescribeProductCARequest) (response *DescribeProductCAResponse, err error) {
    if request == nil {
        request = NewDescribeProductCARequest()
    }
    response = NewDescribeProductCAResponse()
    err = c.Send(request, response)
    return
}

func NewSetProductsForbiddenStatusRequest() (request *SetProductsForbiddenStatusRequest) {
    request = &SetProductsForbiddenStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "SetProductsForbiddenStatus")
    return
}

func NewSetProductsForbiddenStatusResponse() (response *SetProductsForbiddenStatusResponse) {
    response = &SetProductsForbiddenStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetProductsForbiddenStatus
// This API is used to enable or disable multiple products at a time.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
func (c *Client) SetProductsForbiddenStatus(request *SetProductsForbiddenStatusRequest) (response *SetProductsForbiddenStatusResponse, err error) {
    if request == nil {
        request = NewSetProductsForbiddenStatusRequest()
    }
    response = NewSetProductsForbiddenStatusResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDeviceLogLevelRequest() (request *UpdateDeviceLogLevelRequest) {
    request = &UpdateDeviceLogLevelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "UpdateDeviceLogLevel")
    return
}

func NewUpdateDeviceLogLevelResponse() (response *UpdateDeviceLogLevelResponse) {
    response = &UpdateDeviceLogLevelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateDeviceLogLevel
// This API is used to set the device log level.  
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  UNAUTHORIZEDOPERATION_DEVICEISNOTENABLED = "UnauthorizedOperation.DeviceIsNotEnabled"
func (c *Client) UpdateDeviceLogLevel(request *UpdateDeviceLogLevelRequest) (response *UpdateDeviceLogLevelResponse, err error) {
    if request == nil {
        request = NewUpdateDeviceLogLevelRequest()
    }
    response = NewUpdateDeviceLogLevelResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDevicesEnableStateRequest() (request *UpdateDevicesEnableStateRequest) {
    request = &UpdateDevicesEnableStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "UpdateDevicesEnableState")
    return
}

func NewUpdateDevicesEnableStateResponse() (response *UpdateDevicesEnableStateResponse) {
    response = &UpdateDevicesEnableStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateDevicesEnableState
// This API is used to enable or disable multiple devices. 
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_PRODUCTTYPENOTSUPPORT = "InvalidParameterValue.ProductTypeNotSupport"
//  RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"
//  RESOURCENOTFOUND_PRODUCTNOTEXIST = "ResourceNotFound.ProductNotExist"
//  UNAUTHORIZEDOPERATION_DEVICEHASALREADYBINDGATEWAY = "UnauthorizedOperation.DeviceHasAlreadyBindGateway"
func (c *Client) UpdateDevicesEnableState(request *UpdateDevicesEnableStateRequest) (response *UpdateDevicesEnableStateResponse, err error) {
    if request == nil {
        request = NewUpdateDevicesEnableStateRequest()
    }
    response = NewUpdateDevicesEnableStateResponse()
    err = c.Send(request, response)
    return
}

func NewUpdatePrivateCARequest() (request *UpdatePrivateCARequest) {
    request = &UpdatePrivateCARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotcloud", APIVersion, "UpdatePrivateCA")
    return
}

func NewUpdatePrivateCAResponse() (response *UpdatePrivateCAResponse) {
    response = &UpdatePrivateCAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdatePrivateCA
// This API is used to update a private CA certificate.
//
// error code that may be returned:
//  LIMITEXCEEDED_CAREPEAT = "LimitExceeded.CARepeat"
func (c *Client) UpdatePrivateCA(request *UpdatePrivateCARequest) (response *UpdatePrivateCAResponse, err error) {
    if request == nil {
        request = NewUpdatePrivateCARequest()
    }
    response = NewUpdatePrivateCAResponse()
    err = c.Send(request, response)
    return
}
