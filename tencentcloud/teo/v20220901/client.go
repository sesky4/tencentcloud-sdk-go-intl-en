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

package v20220901

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2022-09-01"

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


func NewBindSecurityTemplateToEntityRequest() (request *BindSecurityTemplateToEntityRequest) {
    request = &BindSecurityTemplateToEntityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "BindSecurityTemplateToEntity")
    
    
    return
}

func NewBindSecurityTemplateToEntityResponse() (response *BindSecurityTemplateToEntityResponse) {
    response = &BindSecurityTemplateToEntityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindSecurityTemplateToEntity
// This API is used to bind/unbind a domain name to/from a specific policy template. 
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BindSecurityTemplateToEntity(request *BindSecurityTemplateToEntityRequest) (response *BindSecurityTemplateToEntityResponse, err error) {
    return c.BindSecurityTemplateToEntityWithContext(context.Background(), request)
}

// BindSecurityTemplateToEntity
// This API is used to bind/unbind a domain name to/from a specific policy template. 
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BindSecurityTemplateToEntityWithContext(ctx context.Context, request *BindSecurityTemplateToEntityRequest) (response *BindSecurityTemplateToEntityResponse, err error) {
    if request == nil {
        request = NewBindSecurityTemplateToEntityRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindSecurityTemplateToEntity require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindSecurityTemplateToEntityResponse()
    err = c.Send(request, response)
    return
}

func NewBindSharedCNAMERequest() (request *BindSharedCNAMERequest) {
    request = &BindSharedCNAMERequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "BindSharedCNAME")
    
    
    return
}

func NewBindSharedCNAMEResponse() (response *BindSharedCNAMEResponse) {
    response = &BindSharedCNAMEResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindSharedCNAME
// This API is used to bind/unbind a domain name to/from a shared CNAME. It is now only available to beta users.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) BindSharedCNAME(request *BindSharedCNAMERequest) (response *BindSharedCNAMEResponse, err error) {
    return c.BindSharedCNAMEWithContext(context.Background(), request)
}

// BindSharedCNAME
// This API is used to bind/unbind a domain name to/from a shared CNAME. It is now only available to beta users.
//
// error code that may be returned:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) BindSharedCNAMEWithContext(ctx context.Context, request *BindSharedCNAMERequest) (response *BindSharedCNAMEResponse, err error) {
    if request == nil {
        request = NewBindSharedCNAMERequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindSharedCNAME require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindSharedCNAMEResponse()
    err = c.Send(request, response)
    return
}

func NewBindZoneToPlanRequest() (request *BindZoneToPlanRequest) {
    request = &BindZoneToPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "BindZoneToPlan")
    
    
    return
}

func NewBindZoneToPlanResponse() (response *BindZoneToPlanResponse) {
    response = &BindZoneToPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindZoneToPlan
// This API is used to bind a site to a plan.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PLANNOTFOUND = "InvalidParameter.PlanNotFound"
//  INVALIDPARAMETER_ZONEHASBEENBOUND = "InvalidParameter.ZoneHasBeenBound"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) BindZoneToPlan(request *BindZoneToPlanRequest) (response *BindZoneToPlanResponse, err error) {
    return c.BindZoneToPlanWithContext(context.Background(), request)
}

// BindZoneToPlan
// This API is used to bind a site to a plan.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PLANNOTFOUND = "InvalidParameter.PlanNotFound"
//  INVALIDPARAMETER_ZONEHASBEENBOUND = "InvalidParameter.ZoneHasBeenBound"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) BindZoneToPlanWithContext(ctx context.Context, request *BindZoneToPlanRequest) (response *BindZoneToPlanResponse, err error) {
    if request == nil {
        request = NewBindZoneToPlanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindZoneToPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindZoneToPlanResponse()
    err = c.Send(request, response)
    return
}

func NewCheckCnameStatusRequest() (request *CheckCnameStatusRequest) {
    request = &CheckCnameStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CheckCnameStatus")
    
    
    return
}

func NewCheckCnameStatusResponse() (response *CheckCnameStatusResponse) {
    response = &CheckCnameStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckCnameStatus
// This API is used to query the CNAME status of a domain name.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) CheckCnameStatus(request *CheckCnameStatusRequest) (response *CheckCnameStatusResponse, err error) {
    return c.CheckCnameStatusWithContext(context.Background(), request)
}

// CheckCnameStatus
// This API is used to query the CNAME status of a domain name.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) CheckCnameStatusWithContext(ctx context.Context, request *CheckCnameStatusRequest) (response *CheckCnameStatusResponse, err error) {
    if request == nil {
        request = NewCheckCnameStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckCnameStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckCnameStatusResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccelerationDomainRequest() (request *CreateAccelerationDomainRequest) {
    request = &CreateAccelerationDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateAccelerationDomain")
    
    
    return
}

func NewCreateAccelerationDomainResponse() (response *CreateAccelerationDomainResponse) {
    response = &CreateAccelerationDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAccelerationDomain
// This API is used to create an acceleration domain name. 
//
// 
//
// For sites connected via the CNAME, if you have not verified the ownership of the domain name, the ownership verification information of the domain name is returned. To verify your ownership of the domain name, see [Ownership Verification](https://intl.cloud.tencent.com/document/product/1552/70789?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INVALIDPARAMETER_CONFLICTHOSTORIGIN = "InvalidParameter.ConflictHostOrigin"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_INVALIDACCELERATETYPE = "InvalidParameter.InvalidAccelerateType"
//  INVALIDPARAMETER_INVALIDAWSPRIVATEACCESS = "InvalidParameter.InvalidAwsPrivateAccess"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDCLIENTIPORIGIN = "InvalidParameter.InvalidClientIpOrigin"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_INVALIDPRIVATEACCESSPARAMS = "InvalidParameter.InvalidPrivateAccessParams"
//  INVALIDPARAMETER_INVALIDPRIVATEACCESSSWITCH = "InvalidParameter.InvalidPrivateAccessSwitch"
//  INVALIDPARAMETER_INVALIDQUICBILLING = "InvalidParameter.InvalidQuicBilling"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGCLIENTIP = "InvalidParameter.InvalidStandardDebugClientIp"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_SPACENOTBINDORIGIN = "InvalidParameter.SpaceNotBindOrigin"
//  INVALIDPARAMETER_ZONEISGRAYPUBLISHING = "InvalidParameter.ZoneIsGrayPublishing"
//  INVALIDPARAMETERVALUE_CONFLICTRECORD = "InvalidParameterValue.ConflictRecord"
//  INVALIDPARAMETERVALUE_CONFLICTWITHDNSSEC = "InvalidParameterValue.ConflictWithDNSSEC"
//  INVALIDPARAMETERVALUE_CONFLICTWITHNSRECORD = "InvalidParameterValue.ConflictWithNSRecord"
//  INVALIDPARAMETERVALUE_CONTENTSAMEASNAME = "InvalidParameterValue.ContentSameAsName"
//  INVALIDPARAMETERVALUE_DOMAINNOTMATCHZONE = "InvalidParameterValue.DomainNotMatchZone"
//  INVALIDPARAMETERVALUE_INVALIDDNSCONTENT = "InvalidParameterValue.InvalidDNSContent"
//  INVALIDPARAMETERVALUE_INVALIDDNSNAME = "InvalidParameterValue.InvalidDNSName"
//  INVALIDPARAMETERVALUE_INVALIDDOMAINNAME = "InvalidParameterValue.InvalidDomainName"
//  INVALIDPARAMETERVALUE_INVALIDPROXYORIGIN = "InvalidParameterValue.InvalidProxyOrigin"
//  INVALIDPARAMETERVALUE_ORIGINGROUPNOTEXISTS = "InvalidParameterValue.OriginGroupNotExists"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDDISABLE = "OperationDenied.AccelerateMainlandDisable"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_DOMAINNOICP = "OperationDenied.DomainNoICP"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSEZONEAREA = "OperationDenied.InvalidAdvancedDefenseZoneArea"
//  OPERATIONDENIED_LOADBALANCESTATUSNOTINONLINE = "OperationDenied.LoadBalanceStatusNotInOnline"
//  OPERATIONDENIED_RECORDISFORBIDDEN = "OperationDenied.RecordIsForbidden"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_DNSRECORD = "ResourceInUse.DnsRecord"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_DOMAINALREADYEXISTS = "ResourceUnavailable.DomainAlreadyExists"
//  RESOURCESSOLDOUT_L7LACKOFRESOURCES = "ResourcesSoldOut.L7LackOfResources"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) CreateAccelerationDomain(request *CreateAccelerationDomainRequest) (response *CreateAccelerationDomainResponse, err error) {
    return c.CreateAccelerationDomainWithContext(context.Background(), request)
}

// CreateAccelerationDomain
// This API is used to create an acceleration domain name. 
//
// 
//
// For sites connected via the CNAME, if you have not verified the ownership of the domain name, the ownership verification information of the domain name is returned. To verify your ownership of the domain name, see [Ownership Verification](https://intl.cloud.tencent.com/document/product/1552/70789?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INVALIDPARAMETER_CONFLICTHOSTORIGIN = "InvalidParameter.ConflictHostOrigin"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_INVALIDACCELERATETYPE = "InvalidParameter.InvalidAccelerateType"
//  INVALIDPARAMETER_INVALIDAWSPRIVATEACCESS = "InvalidParameter.InvalidAwsPrivateAccess"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDCLIENTIPORIGIN = "InvalidParameter.InvalidClientIpOrigin"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_INVALIDPRIVATEACCESSPARAMS = "InvalidParameter.InvalidPrivateAccessParams"
//  INVALIDPARAMETER_INVALIDPRIVATEACCESSSWITCH = "InvalidParameter.InvalidPrivateAccessSwitch"
//  INVALIDPARAMETER_INVALIDQUICBILLING = "InvalidParameter.InvalidQuicBilling"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGCLIENTIP = "InvalidParameter.InvalidStandardDebugClientIp"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_SPACENOTBINDORIGIN = "InvalidParameter.SpaceNotBindOrigin"
//  INVALIDPARAMETER_ZONEISGRAYPUBLISHING = "InvalidParameter.ZoneIsGrayPublishing"
//  INVALIDPARAMETERVALUE_CONFLICTRECORD = "InvalidParameterValue.ConflictRecord"
//  INVALIDPARAMETERVALUE_CONFLICTWITHDNSSEC = "InvalidParameterValue.ConflictWithDNSSEC"
//  INVALIDPARAMETERVALUE_CONFLICTWITHNSRECORD = "InvalidParameterValue.ConflictWithNSRecord"
//  INVALIDPARAMETERVALUE_CONTENTSAMEASNAME = "InvalidParameterValue.ContentSameAsName"
//  INVALIDPARAMETERVALUE_DOMAINNOTMATCHZONE = "InvalidParameterValue.DomainNotMatchZone"
//  INVALIDPARAMETERVALUE_INVALIDDNSCONTENT = "InvalidParameterValue.InvalidDNSContent"
//  INVALIDPARAMETERVALUE_INVALIDDNSNAME = "InvalidParameterValue.InvalidDNSName"
//  INVALIDPARAMETERVALUE_INVALIDDOMAINNAME = "InvalidParameterValue.InvalidDomainName"
//  INVALIDPARAMETERVALUE_INVALIDPROXYORIGIN = "InvalidParameterValue.InvalidProxyOrigin"
//  INVALIDPARAMETERVALUE_ORIGINGROUPNOTEXISTS = "InvalidParameterValue.OriginGroupNotExists"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDDISABLE = "OperationDenied.AccelerateMainlandDisable"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_DOMAINNOICP = "OperationDenied.DomainNoICP"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSEZONEAREA = "OperationDenied.InvalidAdvancedDefenseZoneArea"
//  OPERATIONDENIED_LOADBALANCESTATUSNOTINONLINE = "OperationDenied.LoadBalanceStatusNotInOnline"
//  OPERATIONDENIED_RECORDISFORBIDDEN = "OperationDenied.RecordIsForbidden"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_DNSRECORD = "ResourceInUse.DnsRecord"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_DOMAINALREADYEXISTS = "ResourceUnavailable.DomainAlreadyExists"
//  RESOURCESSOLDOUT_L7LACKOFRESOURCES = "ResourcesSoldOut.L7LackOfResources"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) CreateAccelerationDomainWithContext(ctx context.Context, request *CreateAccelerationDomainRequest) (response *CreateAccelerationDomainResponse, err error) {
    if request == nil {
        request = NewCreateAccelerationDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAccelerationDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAccelerationDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAliasDomainRequest() (request *CreateAliasDomainRequest) {
    request = &CreateAliasDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateAliasDomain")
    
    
    return
}

func NewCreateAliasDomainResponse() (response *CreateAliasDomainResponse) {
    response = &CreateAliasDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAliasDomain
// This API is used to create an alias domain name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INVALIDPARAMETER_ALIASDOMAINNOTSUPPORTSMCERT = "InvalidParameter.AliasDomainNotSupportSMCert"
//  INVALIDPARAMETER_CERTNOTMATCHDOMAIN = "InvalidParameter.CertNotMatchDomain"
//  INVALIDPARAMETER_INVALIDHTTPSCERTINFO = "InvalidParameter.InvalidHttpsCertInfo"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDALIASDOMAINNAME = "InvalidParameterValue.InvalidAliasDomainName"
//  INVALIDPARAMETERVALUE_INVALIDALIASNAMESUFFIX = "InvalidParameterValue.InvalidAliasNameSuffix"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DOMAINISBLOCKED = "OperationDenied.DomainIsBlocked"
//  OPERATIONDENIED_DOMAINNOICP = "OperationDenied.DomainNoICP"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  RESOURCEINUSE_ALIASNAME = "ResourceInUse.AliasName"
//  RESOURCEINUSE_DUPLICATENAME = "ResourceInUse.DuplicateName"
//  RESOURCEINUSE_ZONE = "ResourceInUse.Zone"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_DOMAINALREADYEXISTS = "ResourceUnavailable.DomainAlreadyExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNSUPPORTEDOPERATION_TARGETNAMEORIGINTYPECOS = "UnsupportedOperation.TargetNameOriginTypeCos"
func (c *Client) CreateAliasDomain(request *CreateAliasDomainRequest) (response *CreateAliasDomainResponse, err error) {
    return c.CreateAliasDomainWithContext(context.Background(), request)
}

// CreateAliasDomain
// This API is used to create an alias domain name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INVALIDPARAMETER_ALIASDOMAINNOTSUPPORTSMCERT = "InvalidParameter.AliasDomainNotSupportSMCert"
//  INVALIDPARAMETER_CERTNOTMATCHDOMAIN = "InvalidParameter.CertNotMatchDomain"
//  INVALIDPARAMETER_INVALIDHTTPSCERTINFO = "InvalidParameter.InvalidHttpsCertInfo"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDALIASDOMAINNAME = "InvalidParameterValue.InvalidAliasDomainName"
//  INVALIDPARAMETERVALUE_INVALIDALIASNAMESUFFIX = "InvalidParameterValue.InvalidAliasNameSuffix"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DOMAINISBLOCKED = "OperationDenied.DomainIsBlocked"
//  OPERATIONDENIED_DOMAINNOICP = "OperationDenied.DomainNoICP"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  RESOURCEINUSE_ALIASNAME = "ResourceInUse.AliasName"
//  RESOURCEINUSE_DUPLICATENAME = "ResourceInUse.DuplicateName"
//  RESOURCEINUSE_ZONE = "ResourceInUse.Zone"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_DOMAINALREADYEXISTS = "ResourceUnavailable.DomainAlreadyExists"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNSUPPORTEDOPERATION_TARGETNAMEORIGINTYPECOS = "UnsupportedOperation.TargetNameOriginTypeCos"
func (c *Client) CreateAliasDomainWithContext(ctx context.Context, request *CreateAliasDomainRequest) (response *CreateAliasDomainResponse, err error) {
    if request == nil {
        request = NewCreateAliasDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAliasDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAliasDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationProxyRequest() (request *CreateApplicationProxyRequest) {
    request = &CreateApplicationProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateApplicationProxy")
    
    
    return
}

func NewCreateApplicationProxyResponse() (response *CreateApplicationProxyResponse) {
    response = &CreateApplicationProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApplicationProxy
// This API is on an earlier version. If you want to call it, please switch to the latest version [CreateL4Proxy] (https://intl.cloud.tencent.com/document/product/1552/103417?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETER_PROXYNAMEDUPLICATING = "InvalidParameter.ProxyNameDuplicating"
//  INVALIDPARAMETER_PROXYNAMENOTMATCHED = "InvalidParameter.ProxyNameNotMatched"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_USERQUOTALIMITED = "LimitExceeded.UserQuotaLimited"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_PLATTYPEIPACCELERATEMAINLANDNOTSUPPORT = "OperationDenied.PlatTypeIPAccelerateMainlandNotSupport"
//  OPERATIONDENIED_ZONENOTACTIVE = "OperationDenied.ZoneNotActive"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateApplicationProxy(request *CreateApplicationProxyRequest) (response *CreateApplicationProxyResponse, err error) {
    return c.CreateApplicationProxyWithContext(context.Background(), request)
}

// CreateApplicationProxy
// This API is on an earlier version. If you want to call it, please switch to the latest version [CreateL4Proxy] (https://intl.cloud.tencent.com/document/product/1552/103417?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETER_PROXYNAMEDUPLICATING = "InvalidParameter.ProxyNameDuplicating"
//  INVALIDPARAMETER_PROXYNAMENOTMATCHED = "InvalidParameter.ProxyNameNotMatched"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_USERQUOTALIMITED = "LimitExceeded.UserQuotaLimited"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_PLATTYPEIPACCELERATEMAINLANDNOTSUPPORT = "OperationDenied.PlatTypeIPAccelerateMainlandNotSupport"
//  OPERATIONDENIED_ZONENOTACTIVE = "OperationDenied.ZoneNotActive"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateApplicationProxyWithContext(ctx context.Context, request *CreateApplicationProxyRequest) (response *CreateApplicationProxyResponse, err error) {
    if request == nil {
        request = NewCreateApplicationProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplicationProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApplicationProxyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationProxyRuleRequest() (request *CreateApplicationProxyRuleRequest) {
    request = &CreateApplicationProxyRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateApplicationProxyRule")
    
    
    return
}

func NewCreateApplicationProxyRuleResponse() (response *CreateApplicationProxyRuleResponse) {
    response = &CreateApplicationProxyRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApplicationProxyRule
// This API is on an earlier version. If you want to call it, please switch to the latest version. For details, see [CreateL4ProxyRules] (https://intl.cloud.tencent.com/document/product/1552/103416?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_INVALIDRULEPROTO = "InvalidParameter.InvalidRuleProto"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETER_RULEORIGINFORMATERROR = "InvalidParameter.RuleOriginFormatError"
//  INVALIDPARAMETER_RULEORIGINMULTIDOMAIN = "InvalidParameter.RuleOriginMultiDomain"
//  INVALIDPARAMETER_RULEORIGINPORTINTEGER = "InvalidParameter.RuleOriginPortInteger"
//  INVALIDPARAMETER_RULEORIGINVALUEERROR = "InvalidParameter.RuleOriginValueError"
//  INVALIDPARAMETER_RULEPORTDUPLICATING = "InvalidParameter.RulePortDuplicating"
//  INVALIDPARAMETER_RULEPORTGROUP = "InvalidParameter.RulePortGroup"
//  INVALIDPARAMETER_RULEPORTINTEGER = "InvalidParameter.RulePortInteger"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateApplicationProxyRule(request *CreateApplicationProxyRuleRequest) (response *CreateApplicationProxyRuleResponse, err error) {
    return c.CreateApplicationProxyRuleWithContext(context.Background(), request)
}

// CreateApplicationProxyRule
// This API is on an earlier version. If you want to call it, please switch to the latest version. For details, see [CreateL4ProxyRules] (https://intl.cloud.tencent.com/document/product/1552/103416?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_INVALIDRULEPROTO = "InvalidParameter.InvalidRuleProto"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETER_RULEORIGINFORMATERROR = "InvalidParameter.RuleOriginFormatError"
//  INVALIDPARAMETER_RULEORIGINMULTIDOMAIN = "InvalidParameter.RuleOriginMultiDomain"
//  INVALIDPARAMETER_RULEORIGINPORTINTEGER = "InvalidParameter.RuleOriginPortInteger"
//  INVALIDPARAMETER_RULEORIGINVALUEERROR = "InvalidParameter.RuleOriginValueError"
//  INVALIDPARAMETER_RULEPORTDUPLICATING = "InvalidParameter.RulePortDuplicating"
//  INVALIDPARAMETER_RULEPORTGROUP = "InvalidParameter.RulePortGroup"
//  INVALIDPARAMETER_RULEPORTINTEGER = "InvalidParameter.RulePortInteger"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateApplicationProxyRuleWithContext(ctx context.Context, request *CreateApplicationProxyRuleRequest) (response *CreateApplicationProxyRuleResponse, err error) {
    if request == nil {
        request = NewCreateApplicationProxyRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplicationProxyRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApplicationProxyRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCLSIndexRequest() (request *CreateCLSIndexRequest) {
    request = &CreateCLSIndexRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateCLSIndex")
    
    
    return
}

func NewCreateCLSIndexResponse() (response *CreateCLSIndexResponse) {
    response = &CreateCLSIndexResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCLSIndex
// This API is used to create key-value indexes for relevant delivered log fields in the corresponding Tencent Cloud CLS log topic for a specified real-time log delivery task (task-id). If such indexes have been created in CLS, this API will append indexes through merging.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATECLSLOGSETFAILED = "FailedOperation.CreateClsLogSetFailed"
func (c *Client) CreateCLSIndex(request *CreateCLSIndexRequest) (response *CreateCLSIndexResponse, err error) {
    return c.CreateCLSIndexWithContext(context.Background(), request)
}

// CreateCLSIndex
// This API is used to create key-value indexes for relevant delivered log fields in the corresponding Tencent Cloud CLS log topic for a specified real-time log delivery task (task-id). If such indexes have been created in CLS, this API will append indexes through merging.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATECLSLOGSETFAILED = "FailedOperation.CreateClsLogSetFailed"
func (c *Client) CreateCLSIndexWithContext(ctx context.Context, request *CreateCLSIndexRequest) (response *CreateCLSIndexResponse, err error) {
    if request == nil {
        request = NewCreateCLSIndexRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCLSIndex require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCLSIndexResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConfigGroupVersionRequest() (request *CreateConfigGroupVersionRequest) {
    request = &CreateConfigGroupVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateConfigGroupVersion")
    
    
    return
}

func NewCreateConfigGroupVersionResponse() (response *CreateConfigGroupVersionResponse) {
    response = &CreateConfigGroupVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateConfigGroupVersion
// This API is used to create a new version for the specified configuration group in version management mode. The version management feature is currently undergoing beta testing and is accessible only to users on the whitelist.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CONFIGCONDITIONSYNTAXERROR = "FailedOperation.ConfigConditionSyntaxError"
//  FAILEDOPERATION_CONFIGCONDITIONUNKNOWNTARGET = "FailedOperation.ConfigConditionUnknownTarget"
//  FAILEDOPERATION_CONFIGCONDITIONVALUEEMPTYERROR = "FailedOperation.ConfigConditionValueEmptyError"
//  FAILEDOPERATION_CONFIGFIELDTYPEERROR = "FailedOperation.ConfigFieldTypeError"
//  FAILEDOPERATION_CONFIGFORMATERROR = "FailedOperation.ConfigFormatError"
//  FAILEDOPERATION_CONFIGMALFORMEDCONTENT = "FailedOperation.ConfigMalformedContent"
//  FAILEDOPERATION_CONFIGPARAMVALIDATEERRORS = "FailedOperation.ConfigParamValidateErrors"
//  FAILEDOPERATION_CONFIGUNKNOWNFIELD = "FailedOperation.ConfigUnknownField"
//  FAILEDOPERATION_CONFIGUNSUPPORTEDFORMATVERSION = "FailedOperation.ConfigUnsupportedFormatVersion"
//  FAILEDOPERATION_MISSINGCONFIGCHUNK = "FailedOperation.MissingConfigChunk"
//  FAILEDOPERATION_UNKNOWNCONFIGGROUPTYPE = "FailedOperation.UnknownConfigGroupType"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DUPLICATERULE = "InvalidParameter.DuplicateRule"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGACTION = "InvalidParameter.InvalidCacheKeyQueryStringAction"
//  INVALIDPARAMETER_INVALIDCONDITIONS = "InvalidParameter.InvalidConditions"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGEFOLLOWORIGIN = "InvalidParameter.InvalidMaxAgeFollowOrigin"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUG = "InvalidParameter.InvalidStandardDebug"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_STATUSCODECACHEINVALIDSTATUSCODE = "InvalidParameter.StatusCodeCacheInvalidStatusCode"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RULELIMITEXCEEDED = "LimitExceeded.RuleLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDIPV6CONFLICT = "OperationDenied.AccelerateMainlandIpv6Conflict"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_NOTINVERSIONCONTROLWHITELIST = "OperationDenied.NotInVersionControlWhiteList"
//  OPERATIONDENIED_OWNERSHIPVERIFICATIONNOTPASSED = "OperationDenied.OwnershipVerificationNotPassed"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  OPERATIONDENIED_WORKMODENOTINVERSIONCONTROL = "OperationDenied.WorkModeNotInVersionControl"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_FUNCTIONNOTFOUND = "ResourceUnavailable.FunctionNotFound"
func (c *Client) CreateConfigGroupVersion(request *CreateConfigGroupVersionRequest) (response *CreateConfigGroupVersionResponse, err error) {
    return c.CreateConfigGroupVersionWithContext(context.Background(), request)
}

// CreateConfigGroupVersion
// This API is used to create a new version for the specified configuration group in version management mode. The version management feature is currently undergoing beta testing and is accessible only to users on the whitelist.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CONFIGCONDITIONSYNTAXERROR = "FailedOperation.ConfigConditionSyntaxError"
//  FAILEDOPERATION_CONFIGCONDITIONUNKNOWNTARGET = "FailedOperation.ConfigConditionUnknownTarget"
//  FAILEDOPERATION_CONFIGCONDITIONVALUEEMPTYERROR = "FailedOperation.ConfigConditionValueEmptyError"
//  FAILEDOPERATION_CONFIGFIELDTYPEERROR = "FailedOperation.ConfigFieldTypeError"
//  FAILEDOPERATION_CONFIGFORMATERROR = "FailedOperation.ConfigFormatError"
//  FAILEDOPERATION_CONFIGMALFORMEDCONTENT = "FailedOperation.ConfigMalformedContent"
//  FAILEDOPERATION_CONFIGPARAMVALIDATEERRORS = "FailedOperation.ConfigParamValidateErrors"
//  FAILEDOPERATION_CONFIGUNKNOWNFIELD = "FailedOperation.ConfigUnknownField"
//  FAILEDOPERATION_CONFIGUNSUPPORTEDFORMATVERSION = "FailedOperation.ConfigUnsupportedFormatVersion"
//  FAILEDOPERATION_MISSINGCONFIGCHUNK = "FailedOperation.MissingConfigChunk"
//  FAILEDOPERATION_UNKNOWNCONFIGGROUPTYPE = "FailedOperation.UnknownConfigGroupType"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DUPLICATERULE = "InvalidParameter.DuplicateRule"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGACTION = "InvalidParameter.InvalidCacheKeyQueryStringAction"
//  INVALIDPARAMETER_INVALIDCONDITIONS = "InvalidParameter.InvalidConditions"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGEFOLLOWORIGIN = "InvalidParameter.InvalidMaxAgeFollowOrigin"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUG = "InvalidParameter.InvalidStandardDebug"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_STATUSCODECACHEINVALIDSTATUSCODE = "InvalidParameter.StatusCodeCacheInvalidStatusCode"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_RULELIMITEXCEEDED = "LimitExceeded.RuleLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDIPV6CONFLICT = "OperationDenied.AccelerateMainlandIpv6Conflict"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_NOTINVERSIONCONTROLWHITELIST = "OperationDenied.NotInVersionControlWhiteList"
//  OPERATIONDENIED_OWNERSHIPVERIFICATIONNOTPASSED = "OperationDenied.OwnershipVerificationNotPassed"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  OPERATIONDENIED_WORKMODENOTINVERSIONCONTROL = "OperationDenied.WorkModeNotInVersionControl"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_FUNCTIONNOTFOUND = "ResourceUnavailable.FunctionNotFound"
func (c *Client) CreateConfigGroupVersionWithContext(ctx context.Context, request *CreateConfigGroupVersionRequest) (response *CreateConfigGroupVersionResponse, err error) {
    if request == nil {
        request = NewCreateConfigGroupVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConfigGroupVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConfigGroupVersionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateL4ProxyRequest() (request *CreateL4ProxyRequest) {
    request = &CreateL4ProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateL4Proxy")
    
    
    return
}

func NewCreateL4ProxyResponse() (response *CreateL4ProxyResponse) {
    response = &CreateL4ProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateL4Proxy
// This API is used to create Layer 4 proxy instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_PROXYNAMEDUPLICATING = "InvalidParameter.ProxyNameDuplicating"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDDISABLE = "OperationDenied.AccelerateMainlandDisable"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4LACKOFRESOURCES = "OperationDenied.L4LackOfResources"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateL4Proxy(request *CreateL4ProxyRequest) (response *CreateL4ProxyResponse, err error) {
    return c.CreateL4ProxyWithContext(context.Background(), request)
}

// CreateL4Proxy
// This API is used to create Layer 4 proxy instances.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_PROXYNAMEDUPLICATING = "InvalidParameter.ProxyNameDuplicating"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDDISABLE = "OperationDenied.AccelerateMainlandDisable"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4LACKOFRESOURCES = "OperationDenied.L4LackOfResources"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateL4ProxyWithContext(ctx context.Context, request *CreateL4ProxyRequest) (response *CreateL4ProxyResponse, err error) {
    if request == nil {
        request = NewCreateL4ProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateL4Proxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateL4ProxyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateL4ProxyRulesRequest() (request *CreateL4ProxyRulesRequest) {
    request = &CreateL4ProxyRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateL4ProxyRules")
    
    
    return
}

func NewCreateL4ProxyRulesResponse() (response *CreateL4ProxyRulesResponse) {
    response = &CreateL4ProxyRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateL4ProxyRules
// This API is used to create Layer 4 proxy instance rules, supporting both individual and batch creation.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_INVALIDORIGINVALUE = "InvalidParameter.InvalidOriginValue"
//  INVALIDPARAMETER_RULEPORTDUPLICATING = "InvalidParameter.RulePortDuplicating"
//  LIMITEXCEEDED_PROXYRULESLIMITEXCEEDED = "LimitExceeded.ProxyRulesLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_L4PORTLACKOFRESOURCES = "OperationDenied.L4PortLackOfResources"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateL4ProxyRules(request *CreateL4ProxyRulesRequest) (response *CreateL4ProxyRulesResponse, err error) {
    return c.CreateL4ProxyRulesWithContext(context.Background(), request)
}

// CreateL4ProxyRules
// This API is used to create Layer 4 proxy instance rules, supporting both individual and batch creation.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_INVALIDORIGINVALUE = "InvalidParameter.InvalidOriginValue"
//  INVALIDPARAMETER_RULEPORTDUPLICATING = "InvalidParameter.RulePortDuplicating"
//  LIMITEXCEEDED_PROXYRULESLIMITEXCEEDED = "LimitExceeded.ProxyRulesLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_L4PORTLACKOFRESOURCES = "OperationDenied.L4PortLackOfResources"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateL4ProxyRulesWithContext(ctx context.Context, request *CreateL4ProxyRulesRequest) (response *CreateL4ProxyRulesResponse, err error) {
    if request == nil {
        request = NewCreateL4ProxyRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateL4ProxyRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateL4ProxyRulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOriginGroupRequest() (request *CreateOriginGroupRequest) {
    request = &CreateOriginGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateOriginGroup")
    
    
    return
}

func NewCreateOriginGroupResponse() (response *CreateOriginGroupResponse) {
    response = &CreateOriginGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOriginGroup
// This API is used to create an origin group for easy management. The created origin server group can be used for **adding acceleration domain names** and **layer-4 proxy configuration**.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INVALIDPARAMETER_HOSTHEADERINVALID = "InvalidParameter.HostHeaderInvalid"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDORIGINGROUPTYPE = "InvalidParameter.InvalidOriginGroupType"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_ORIGINNAMEEXISTS = "InvalidParameter.OriginNameExists"
//  INVALIDPARAMETER_ORIGINRECORDFORMATERROR = "InvalidParameter.OriginRecordFormatError"
//  INVALIDPARAMETER_ORIGINRECORDWEIGHTVALUE = "InvalidParameter.OriginRecordWeightValue"
//  INVALIDPARAMETER_ORIGINTHIRDPARTYPARAMFORMATERROR = "InvalidParameter.OriginThirdPartyParamFormatError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_LOADBALANCINGZONEISNOTACTIVE = "OperationDenied.LoadBalancingZoneIsNotActive"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) CreateOriginGroup(request *CreateOriginGroupRequest) (response *CreateOriginGroupResponse, err error) {
    return c.CreateOriginGroupWithContext(context.Background(), request)
}

// CreateOriginGroup
// This API is used to create an origin group for easy management. The created origin server group can be used for **adding acceleration domain names** and **layer-4 proxy configuration**.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INVALIDPARAMETER_HOSTHEADERINVALID = "InvalidParameter.HostHeaderInvalid"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDORIGINGROUPTYPE = "InvalidParameter.InvalidOriginGroupType"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_ORIGINNAMEEXISTS = "InvalidParameter.OriginNameExists"
//  INVALIDPARAMETER_ORIGINRECORDFORMATERROR = "InvalidParameter.OriginRecordFormatError"
//  INVALIDPARAMETER_ORIGINRECORDWEIGHTVALUE = "InvalidParameter.OriginRecordWeightValue"
//  INVALIDPARAMETER_ORIGINTHIRDPARTYPARAMFORMATERROR = "InvalidParameter.OriginThirdPartyParamFormatError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_LOADBALANCINGZONEISNOTACTIVE = "OperationDenied.LoadBalancingZoneIsNotActive"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) CreateOriginGroupWithContext(ctx context.Context, request *CreateOriginGroupRequest) (response *CreateOriginGroupResponse, err error) {
    if request == nil {
        request = NewCreateOriginGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOriginGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOriginGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePlanForZoneRequest() (request *CreatePlanForZoneRequest) {
    request = &CreatePlanForZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreatePlanForZone")
    
    
    return
}

func NewCreatePlanForZoneResponse() (response *CreatePlanForZoneResponse) {
    response = &CreatePlanForZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePlanForZone
// This API is used to purchase a plan for a new site.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSUFFICIENTACCOUNTBALANCE = "FailedOperation.InsufficientAccountBalance"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_ZONEHASBEENBOUND = "InvalidParameter.ZoneHasBeenBound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DOMAINNOICP = "OperationDenied.DomainNoICP"
func (c *Client) CreatePlanForZone(request *CreatePlanForZoneRequest) (response *CreatePlanForZoneResponse, err error) {
    return c.CreatePlanForZoneWithContext(context.Background(), request)
}

// CreatePlanForZone
// This API is used to purchase a plan for a new site.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INSUFFICIENTACCOUNTBALANCE = "FailedOperation.InsufficientAccountBalance"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_ZONEHASBEENBOUND = "InvalidParameter.ZoneHasBeenBound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DOMAINNOICP = "OperationDenied.DomainNoICP"
func (c *Client) CreatePlanForZoneWithContext(ctx context.Context, request *CreatePlanForZoneRequest) (response *CreatePlanForZoneResponse, err error) {
    if request == nil {
        request = NewCreatePlanForZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePlanForZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePlanForZoneResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrefetchTaskRequest() (request *CreatePrefetchTaskRequest) {
    request = &CreatePrefetchTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreatePrefetchTask")
    
    
    return
}

func NewCreatePrefetchTaskResponse() (response *CreatePrefetchTaskResponse) {
    response = &CreatePrefetchTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePrefetchTask
// This API is used to create a pre-warming task.
//
// error code that may be returned:
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DOMAINCONFIG = "InternalError.DomainConfig"
//  INTERNALERROR_FAILEDTOGENERATEURL = "InternalError.FailedToGenerateUrl"
//  INTERNALERROR_QUOTASYSTEM = "InternalError.QuotaSystem"
//  INVALIDPARAMETER_DOMAINNOTFOUND = "InvalidParameter.DomainNotFound"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETER_TARGET = "InvalidParameter.Target"
//  INVALIDPARAMETER_TASKNOTGENERATED = "InvalidParameter.TaskNotGenerated"
//  INVALIDPARAMETER_UPLOADURL = "InvalidParameter.UploadUrl"
//  LIMITEXCEEDED_BATCHQUOTA = "LimitExceeded.BatchQuota"
//  LIMITEXCEEDED_DAILYQUOTA = "LimitExceeded.DailyQuota"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) CreatePrefetchTask(request *CreatePrefetchTaskRequest) (response *CreatePrefetchTaskResponse, err error) {
    return c.CreatePrefetchTaskWithContext(context.Background(), request)
}

// CreatePrefetchTask
// This API is used to create a pre-warming task.
//
// error code that may be returned:
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DOMAINCONFIG = "InternalError.DomainConfig"
//  INTERNALERROR_FAILEDTOGENERATEURL = "InternalError.FailedToGenerateUrl"
//  INTERNALERROR_QUOTASYSTEM = "InternalError.QuotaSystem"
//  INVALIDPARAMETER_DOMAINNOTFOUND = "InvalidParameter.DomainNotFound"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETER_TARGET = "InvalidParameter.Target"
//  INVALIDPARAMETER_TASKNOTGENERATED = "InvalidParameter.TaskNotGenerated"
//  INVALIDPARAMETER_UPLOADURL = "InvalidParameter.UploadUrl"
//  LIMITEXCEEDED_BATCHQUOTA = "LimitExceeded.BatchQuota"
//  LIMITEXCEEDED_DAILYQUOTA = "LimitExceeded.DailyQuota"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) CreatePrefetchTaskWithContext(ctx context.Context, request *CreatePrefetchTaskRequest) (response *CreatePrefetchTaskResponse, err error) {
    if request == nil {
        request = NewCreatePrefetchTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrefetchTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrefetchTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePurgeTaskRequest() (request *CreatePurgeTaskRequest) {
    request = &CreatePurgeTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreatePurgeTask")
    
    
    return
}

func NewCreatePurgeTaskResponse() (response *CreatePurgeTaskResponse) {
    response = &CreatePurgeTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePurgeTask
// When there are resources updated on the origin with the TTL remaining valid, users cannot access the latest resources. In this case, you can purge the cache using this API. There are two methods: <li>Delete: This method deletes the node cache without verification and retrieves the latest resources from the origin when receiving a request.</li><li>Invalidate: This method marks the node cache as invalid and sends a request with the If-None-Match and If-Modified-Since headers to the origin. If the origin responses with 200, the latest resources are retrieved to be cached on the node. If a 304 response is returned, the latest resources are not cached on the node.
//
// 
//
// </li>For more details, see [Cache Purge](https://intl.cloud.tencent.com/document/product/1552/70759?from_cn_redirect=1). </li>
//
// error code that may be returned:
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DOMAINCONFIG = "InternalError.DomainConfig"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_QUOTASYSTEM = "InternalError.QuotaSystem"
//  INVALIDPARAMETER_DOMAINNOTFOUND = "InvalidParameter.DomainNotFound"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETER_TARGET = "InvalidParameter.Target"
//  INVALIDPARAMETER_TASKNOTGENERATED = "InvalidParameter.TaskNotGenerated"
//  INVALIDPARAMETER_UPLOADURL = "InvalidParameter.UploadUrl"
//  LIMITEXCEEDED_BATCHQUOTA = "LimitExceeded.BatchQuota"
//  LIMITEXCEEDED_DAILYQUOTA = "LimitExceeded.DailyQuota"
//  LIMITEXCEEDED_PACKNOTALLOW = "LimitExceeded.PackNotAllow"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) CreatePurgeTask(request *CreatePurgeTaskRequest) (response *CreatePurgeTaskResponse, err error) {
    return c.CreatePurgeTaskWithContext(context.Background(), request)
}

// CreatePurgeTask
// When there are resources updated on the origin with the TTL remaining valid, users cannot access the latest resources. In this case, you can purge the cache using this API. There are two methods: <li>Delete: This method deletes the node cache without verification and retrieves the latest resources from the origin when receiving a request.</li><li>Invalidate: This method marks the node cache as invalid and sends a request with the If-None-Match and If-Modified-Since headers to the origin. If the origin responses with 200, the latest resources are retrieved to be cached on the node. If a 304 response is returned, the latest resources are not cached on the node.
//
// 
//
// </li>For more details, see [Cache Purge](https://intl.cloud.tencent.com/document/product/1552/70759?from_cn_redirect=1). </li>
//
// error code that may be returned:
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DOMAINCONFIG = "InternalError.DomainConfig"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_QUOTASYSTEM = "InternalError.QuotaSystem"
//  INVALIDPARAMETER_DOMAINNOTFOUND = "InvalidParameter.DomainNotFound"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETER_TARGET = "InvalidParameter.Target"
//  INVALIDPARAMETER_TASKNOTGENERATED = "InvalidParameter.TaskNotGenerated"
//  INVALIDPARAMETER_UPLOADURL = "InvalidParameter.UploadUrl"
//  LIMITEXCEEDED_BATCHQUOTA = "LimitExceeded.BatchQuota"
//  LIMITEXCEEDED_DAILYQUOTA = "LimitExceeded.DailyQuota"
//  LIMITEXCEEDED_PACKNOTALLOW = "LimitExceeded.PackNotAllow"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) CreatePurgeTaskWithContext(ctx context.Context, request *CreatePurgeTaskRequest) (response *CreatePurgeTaskResponse, err error) {
    if request == nil {
        request = NewCreatePurgeTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePurgeTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePurgeTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRealtimeLogDeliveryTaskRequest() (request *CreateRealtimeLogDeliveryTaskRequest) {
    request = &CreateRealtimeLogDeliveryTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateRealtimeLogDeliveryTask")
    
    
    return
}

func NewCreateRealtimeLogDeliveryTaskResponse() (response *CreateRealtimeLogDeliveryTaskResponse) {
    response = &CreateRealtimeLogDeliveryTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRealtimeLogDeliveryTask
// This API is used to create real-time log delivery tasks. This API has the following restrictions:
//
// Under the same combination of data delivery type (LogType) and data delivery area (Area), an entity (Layer 7 domain or Layer 4 proxy instance) can be added to only one real-time log delivery task. It is recommended to first query the real-time log delivery task list by entity through the [DescribeRealtimeLogDeliveryTasks](https://tcloud4api.woa.com/document/product/1657/343539?!preview&!document=1) API to check whether an entity has been added to a real-time log delivery task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATECLSLOGSETFAILED = "FailedOperation.CreateClsLogSetFailed"
//  FAILEDOPERATION_CREATECLSLOGTOPICTASKFAILED = "FailedOperation.CreateClsLogTopicTaskFailed"
//  FAILEDOPERATION_CREATELOGTOPICTASKAUTHFAILURE = "FailedOperation.CreateLogTopicTaskAuthFailure"
//  FAILEDOPERATION_REALTIMELOGAUTHFAILURE = "FailedOperation.RealtimeLogAuthFailure"
//  FAILEDOPERATION_REALTIMELOGNOTFOUND = "FailedOperation.RealtimeLogNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_REALTIMELOGENTITYALREADYCREATED = "InvalidParameter.RealtimeLogEntityAlreadyCreated"
//  INVALIDPARAMETER_REALTIMELOGINVALIDDELIVERYAREA = "InvalidParameter.RealtimeLogInvalidDeliveryArea"
//  INVALIDPARAMETER_REALTIMELOGINVALIDLOGTYPE = "InvalidParameter.RealtimeLogInvalidLogType"
//  INVALIDPARAMETER_REALTIMELOGINVALIDTASKTYPE = "InvalidParameter.RealtimeLogInvalidTaskType"
//  INVALIDPARAMETER_REALTIMELOGNUMSEXCEEDLIMIT = "InvalidParameter.RealtimeLogNumsExceedLimit"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) CreateRealtimeLogDeliveryTask(request *CreateRealtimeLogDeliveryTaskRequest) (response *CreateRealtimeLogDeliveryTaskResponse, err error) {
    return c.CreateRealtimeLogDeliveryTaskWithContext(context.Background(), request)
}

// CreateRealtimeLogDeliveryTask
// This API is used to create real-time log delivery tasks. This API has the following restrictions:
//
// Under the same combination of data delivery type (LogType) and data delivery area (Area), an entity (Layer 7 domain or Layer 4 proxy instance) can be added to only one real-time log delivery task. It is recommended to first query the real-time log delivery task list by entity through the [DescribeRealtimeLogDeliveryTasks](https://tcloud4api.woa.com/document/product/1657/343539?!preview&!document=1) API to check whether an entity has been added to a real-time log delivery task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATECLSLOGSETFAILED = "FailedOperation.CreateClsLogSetFailed"
//  FAILEDOPERATION_CREATECLSLOGTOPICTASKFAILED = "FailedOperation.CreateClsLogTopicTaskFailed"
//  FAILEDOPERATION_CREATELOGTOPICTASKAUTHFAILURE = "FailedOperation.CreateLogTopicTaskAuthFailure"
//  FAILEDOPERATION_REALTIMELOGAUTHFAILURE = "FailedOperation.RealtimeLogAuthFailure"
//  FAILEDOPERATION_REALTIMELOGNOTFOUND = "FailedOperation.RealtimeLogNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_REALTIMELOGENTITYALREADYCREATED = "InvalidParameter.RealtimeLogEntityAlreadyCreated"
//  INVALIDPARAMETER_REALTIMELOGINVALIDDELIVERYAREA = "InvalidParameter.RealtimeLogInvalidDeliveryArea"
//  INVALIDPARAMETER_REALTIMELOGINVALIDLOGTYPE = "InvalidParameter.RealtimeLogInvalidLogType"
//  INVALIDPARAMETER_REALTIMELOGINVALIDTASKTYPE = "InvalidParameter.RealtimeLogInvalidTaskType"
//  INVALIDPARAMETER_REALTIMELOGNUMSEXCEEDLIMIT = "InvalidParameter.RealtimeLogNumsExceedLimit"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) CreateRealtimeLogDeliveryTaskWithContext(ctx context.Context, request *CreateRealtimeLogDeliveryTaskRequest) (response *CreateRealtimeLogDeliveryTaskResponse, err error) {
    if request == nil {
        request = NewCreateRealtimeLogDeliveryTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRealtimeLogDeliveryTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRealtimeLogDeliveryTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRuleRequest() (request *CreateRuleRequest) {
    request = &CreateRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateRule")
    
    
    return
}

func NewCreateRuleResponse() (response *CreateRuleResponse) {
    response = &CreateRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRule
// This API is used to create a rule in the rule engine.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACCESSREDIRECTREGEXERROR = "InvalidParameter.AccessRedirectRegexError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_ERRACTIONUNSUPPORTTARGET = "InvalidParameter.ErrActionUnsupportTarget"
//  INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONORIGINPRIVATEADDRESS = "InvalidParameter.ErrInvalidActionOriginPrivateAddress"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMACTION = "InvalidParameter.ErrInvalidActionParamAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMBADVALUETYPE = "InvalidParameter.ErrInvalidActionParamBadValueType"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMDUPLICATENAME = "InvalidParameter.ErrInvalidActionParamDuplicateName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMNAME = "InvalidParameter.ErrInvalidActionParamName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_ERRINVALIDACTIONTYPE = "InvalidParameter.ErrInvalidActionType"
//  INVALIDPARAMETER_ERRINVALIDCONDITION = "InvalidParameter.ErrInvalidCondition"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONCANNOTONLYCONTAINHOSTWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidConditionCannotOnlyContainHostWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONIGNORECASE = "InvalidParameter.ErrInvalidConditionIgnoreCase"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMEBADNAME = "InvalidParameter.ErrInvalidConditionNameBadName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADREGULAR = "InvalidParameter.ErrInvalidConditionValueBadRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADURL = "InvalidParameter.ErrInvalidConditionValueBadUrl"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUECONTAINFILENAMEEXTENSION = "InvalidParameter.ErrInvalidConditionValueBadValueContainFileNameExtension"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOLONGVALUE = "InvalidParameter.ErrInvalidConditionValueTooLongValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYREGULAR = "InvalidParameter.ErrInvalidConditionValueTooManyRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYVALUES = "InvalidParameter.ErrInvalidConditionValueTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYWILDCARD = "InvalidParameter.ErrInvalidConditionValueTooManyWildcard"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEZEROLENGTH = "InvalidParameter.ErrInvalidConditionValueZeroLength"
//  INVALIDPARAMETER_ERRNILCONDITION = "InvalidParameter.ErrNilCondition"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"
//  INVALIDPARAMETER_INVALIDAUTHENTICATION = "InvalidParameter.InvalidAuthentication"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPEEXPIRETIME = "InvalidParameter.InvalidAuthenticationTypeExpireTime"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"
//  INVALIDPARAMETER_INVALIDCACHECONFIGCACHE = "InvalidParameter.InvalidCacheConfigCache"
//  INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGVALUE = "InvalidParameter.InvalidCacheKeyQueryStringValue"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINEBILLING = "InvalidParameter.InvalidDynamicRoutineBilling"
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEXPIRETIMELIMIT = "InvalidParameter.InvalidStandardDebugExpireTimeLimit"
//  INVALIDPARAMETER_INVALIDUPSTREAMREQUESTQUERYSTRINGVALUE = "InvalidParameter.InvalidUpstreamRequestQueryStringValue"
//  INVALIDPARAMETER_INVALIDURLREDIRECT = "InvalidParameter.InvalidUrlRedirect"
//  INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_KEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.KeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_LOADBALANCEINSTANCEIDISREQUIRED = "InvalidParameter.LoadBalanceInstanceIdIsRequired"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ORIGINORIGINGROUPIDISREQUIRED = "InvalidParameter.OriginOriginGroupIdIsRequired"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_TASKSYSTEMERROR = "InvalidParameter.TaskSystemError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDMULTIPLYLAYERCONFLICT = "OperationDenied.AccelerateMainlandMultiplyLayerConflict"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSESECURITYTYPE = "OperationDenied.InvalidAdvancedDefenseSecurityType"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateRule(request *CreateRuleRequest) (response *CreateRuleResponse, err error) {
    return c.CreateRuleWithContext(context.Background(), request)
}

// CreateRule
// This API is used to create a rule in the rule engine.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACCESSREDIRECTREGEXERROR = "InvalidParameter.AccessRedirectRegexError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_ERRACTIONUNSUPPORTTARGET = "InvalidParameter.ErrActionUnsupportTarget"
//  INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONORIGINPRIVATEADDRESS = "InvalidParameter.ErrInvalidActionOriginPrivateAddress"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMACTION = "InvalidParameter.ErrInvalidActionParamAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMBADVALUETYPE = "InvalidParameter.ErrInvalidActionParamBadValueType"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMDUPLICATENAME = "InvalidParameter.ErrInvalidActionParamDuplicateName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMNAME = "InvalidParameter.ErrInvalidActionParamName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_ERRINVALIDACTIONTYPE = "InvalidParameter.ErrInvalidActionType"
//  INVALIDPARAMETER_ERRINVALIDCONDITION = "InvalidParameter.ErrInvalidCondition"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONCANNOTONLYCONTAINHOSTWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidConditionCannotOnlyContainHostWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONIGNORECASE = "InvalidParameter.ErrInvalidConditionIgnoreCase"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMEBADNAME = "InvalidParameter.ErrInvalidConditionNameBadName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADREGULAR = "InvalidParameter.ErrInvalidConditionValueBadRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADURL = "InvalidParameter.ErrInvalidConditionValueBadUrl"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUECONTAINFILENAMEEXTENSION = "InvalidParameter.ErrInvalidConditionValueBadValueContainFileNameExtension"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOLONGVALUE = "InvalidParameter.ErrInvalidConditionValueTooLongValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYREGULAR = "InvalidParameter.ErrInvalidConditionValueTooManyRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYVALUES = "InvalidParameter.ErrInvalidConditionValueTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYWILDCARD = "InvalidParameter.ErrInvalidConditionValueTooManyWildcard"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEZEROLENGTH = "InvalidParameter.ErrInvalidConditionValueZeroLength"
//  INVALIDPARAMETER_ERRNILCONDITION = "InvalidParameter.ErrNilCondition"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"
//  INVALIDPARAMETER_INVALIDAUTHENTICATION = "InvalidParameter.InvalidAuthentication"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPEEXPIRETIME = "InvalidParameter.InvalidAuthenticationTypeExpireTime"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"
//  INVALIDPARAMETER_INVALIDCACHECONFIGCACHE = "InvalidParameter.InvalidCacheConfigCache"
//  INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGVALUE = "InvalidParameter.InvalidCacheKeyQueryStringValue"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINEBILLING = "InvalidParameter.InvalidDynamicRoutineBilling"
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEXPIRETIMELIMIT = "InvalidParameter.InvalidStandardDebugExpireTimeLimit"
//  INVALIDPARAMETER_INVALIDUPSTREAMREQUESTQUERYSTRINGVALUE = "InvalidParameter.InvalidUpstreamRequestQueryStringValue"
//  INVALIDPARAMETER_INVALIDURLREDIRECT = "InvalidParameter.InvalidUrlRedirect"
//  INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_KEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.KeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_LOADBALANCEINSTANCEIDISREQUIRED = "InvalidParameter.LoadBalanceInstanceIdIsRequired"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ORIGINORIGINGROUPIDISREQUIRED = "InvalidParameter.OriginOriginGroupIdIsRequired"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_TASKSYSTEMERROR = "InvalidParameter.TaskSystemError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDMULTIPLYLAYERCONFLICT = "OperationDenied.AccelerateMainlandMultiplyLayerConflict"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSESECURITYTYPE = "OperationDenied.InvalidAdvancedDefenseSecurityType"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateRuleWithContext(ctx context.Context, request *CreateRuleRequest) (response *CreateRuleResponse, err error) {
    if request == nil {
        request = NewCreateRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecurityIPGroupRequest() (request *CreateSecurityIPGroupRequest) {
    request = &CreateSecurityIPGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateSecurityIPGroup")
    
    
    return
}

func NewCreateSecurityIPGroupResponse() (response *CreateSecurityIPGroupResponse) {
    response = &CreateSecurityIPGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSecurityIPGroup
// This API is used to create a security IP group.
//
// error code that may be returned:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateSecurityIPGroup(request *CreateSecurityIPGroupRequest) (response *CreateSecurityIPGroupResponse, err error) {
    return c.CreateSecurityIPGroupWithContext(context.Background(), request)
}

// CreateSecurityIPGroup
// This API is used to create a security IP group.
//
// error code that may be returned:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) CreateSecurityIPGroupWithContext(ctx context.Context, request *CreateSecurityIPGroupRequest) (response *CreateSecurityIPGroupResponse, err error) {
    if request == nil {
        request = NewCreateSecurityIPGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSecurityIPGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSecurityIPGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSharedCNAMERequest() (request *CreateSharedCNAMERequest) {
    request = &CreateSharedCNAMERequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateSharedCNAME")
    
    
    return
}

func NewCreateSharedCNAMEResponse() (response *CreateSharedCNAMEResponse) {
    response = &CreateSharedCNAMEResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSharedCNAME
// This API is used to create a shared CNAME. It is now only available to beta users.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_NOTALLOWEDWILDCARDSHAREDCNAME = "InvalidParameterValue.NotAllowedWildcardSharedCNAME"
//  INVALIDPARAMETERVALUE_SHAREDCNAMEPREFIXNOTMATCH = "InvalidParameterValue.SharedCNAMEPrefixNotMatch"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  RESOURCEUNAVAILABLE_SHAREDCNAMEALREADYEXISTS = "ResourceUnavailable.SharedCNAMEAlreadyExists"
func (c *Client) CreateSharedCNAME(request *CreateSharedCNAMERequest) (response *CreateSharedCNAMEResponse, err error) {
    return c.CreateSharedCNAMEWithContext(context.Background(), request)
}

// CreateSharedCNAME
// This API is used to create a shared CNAME. It is now only available to beta users.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_NOTALLOWEDWILDCARDSHAREDCNAME = "InvalidParameterValue.NotAllowedWildcardSharedCNAME"
//  INVALIDPARAMETERVALUE_SHAREDCNAMEPREFIXNOTMATCH = "InvalidParameterValue.SharedCNAMEPrefixNotMatch"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  RESOURCEUNAVAILABLE_SHAREDCNAMEALREADYEXISTS = "ResourceUnavailable.SharedCNAMEAlreadyExists"
func (c *Client) CreateSharedCNAMEWithContext(ctx context.Context, request *CreateSharedCNAMERequest) (response *CreateSharedCNAMEResponse, err error) {
    if request == nil {
        request = NewCreateSharedCNAMERequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSharedCNAME require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSharedCNAMEResponse()
    err = c.Send(request, response)
    return
}

func NewCreateZoneRequest() (request *CreateZoneRequest) {
    request = &CreateZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "CreateZone")
    
    
    return
}

func NewCreateZoneResponse() (response *CreateZoneResponse) {
    response = &CreateZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateZone
// This API is used to create a site. After you create the site, you can connect it to EdgeOne via the CNAME or NS (see [Quick Start](https://intl.cloud.tencent.com/document/product/1552/87601?from_cn_redirect=1)), or connect it without a domain name (see [Quick Access to L4 Proxy Service](https://intl.cloud.tencent.com/document/product/1552/96051?from_cn_redirect=1)).
//
// 
//
// If there are already EdgeOne plans under the current account, it is recommended to pass in the `PlanId` to bind the site with the plan directly. If `PlanId` is not passed in, the created site is not activated. You need to call [BindZoneToPlan](https://intl.cloud.tencent.com/document/product/1552/83042?from_cn_redirect=1) to bind the site with a plan. To purchase a plan, please go to the EdgeOne console.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_LENGTHEXCEEDSLIMIT = "InvalidParameter.LengthExceedsLimit"
//  INVALIDPARAMETER_PLANNOTFOUND = "InvalidParameter.PlanNotFound"
//  INVALIDPARAMETER_ZONEHASBEENBOUND = "InvalidParameter.ZoneHasBeenBound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACCESSBLACKLIST = "InvalidParameterValue.AccessBlacklist"
//  INVALIDPARAMETERVALUE_INVALIDTAGVALUE = "InvalidParameterValue.InvalidTagValue"
//  INVALIDPARAMETERVALUE_TOPLEVELDOMAINNOTSUPPORT = "InvalidParameterValue.TopLevelDomainNotSupport"
//  INVALIDPARAMETERVALUE_ZONENAMEINVALID = "InvalidParameterValue.ZoneNameInvalid"
//  INVALIDPARAMETERVALUE_ZONENAMENOTSUPPORTPUNYCODE = "InvalidParameterValue.ZoneNameNotSupportPunyCode"
//  INVALIDPARAMETERVALUE_ZONENAMENOTSUPPORTSUBDOMAIN = "InvalidParameterValue.ZoneNameNotSupportSubDomain"
//  INVALIDPARAMETERVALUE_ZONESAMEASNAME = "InvalidParameterValue.ZoneSameAsName"
//  LIMITEXCEEDED_ZONEBINDPLAN = "LimitExceeded.ZoneBindPlan"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DOMAINISBLOCKED = "OperationDenied.DomainIsBlocked"
//  OPERATIONDENIED_RECORDISFORBIDDEN = "OperationDenied.RecordIsForbidden"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_ALIASDOMAIN = "ResourceInUse.AliasDomain"
//  RESOURCEINUSE_CNAME = "ResourceInUse.Cname"
//  RESOURCEINUSE_DNS = "ResourceInUse.Dns"
//  RESOURCEINUSE_GENERICHOST = "ResourceInUse.GenericHost"
//  RESOURCEINUSE_HOST = "ResourceInUse.Host"
//  RESOURCEINUSE_NS = "ResourceInUse.NS"
//  RESOURCEINUSE_OTHERS = "ResourceInUse.Others"
//  RESOURCEINUSE_OTHERSALIASDOMAIN = "ResourceInUse.OthersAliasDomain"
//  RESOURCEINUSE_OTHERSCNAME = "ResourceInUse.OthersCname"
//  RESOURCEINUSE_OTHERSHOST = "ResourceInUse.OthersHost"
//  RESOURCEINUSE_OTHERSNS = "ResourceInUse.OthersNS"
//  RESOURCEINUSE_SELFANDOTHERSCNAME = "ResourceInUse.SelfAndOthersCname"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) CreateZone(request *CreateZoneRequest) (response *CreateZoneResponse, err error) {
    return c.CreateZoneWithContext(context.Background(), request)
}

// CreateZone
// This API is used to create a site. After you create the site, you can connect it to EdgeOne via the CNAME or NS (see [Quick Start](https://intl.cloud.tencent.com/document/product/1552/87601?from_cn_redirect=1)), or connect it without a domain name (see [Quick Access to L4 Proxy Service](https://intl.cloud.tencent.com/document/product/1552/96051?from_cn_redirect=1)).
//
// 
//
// If there are already EdgeOne plans under the current account, it is recommended to pass in the `PlanId` to bind the site with the plan directly. If `PlanId` is not passed in, the created site is not activated. You need to call [BindZoneToPlan](https://intl.cloud.tencent.com/document/product/1552/83042?from_cn_redirect=1) to bind the site with a plan. To purchase a plan, please go to the EdgeOne console.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_LENGTHEXCEEDSLIMIT = "InvalidParameter.LengthExceedsLimit"
//  INVALIDPARAMETER_PLANNOTFOUND = "InvalidParameter.PlanNotFound"
//  INVALIDPARAMETER_ZONEHASBEENBOUND = "InvalidParameter.ZoneHasBeenBound"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACCESSBLACKLIST = "InvalidParameterValue.AccessBlacklist"
//  INVALIDPARAMETERVALUE_INVALIDTAGVALUE = "InvalidParameterValue.InvalidTagValue"
//  INVALIDPARAMETERVALUE_TOPLEVELDOMAINNOTSUPPORT = "InvalidParameterValue.TopLevelDomainNotSupport"
//  INVALIDPARAMETERVALUE_ZONENAMEINVALID = "InvalidParameterValue.ZoneNameInvalid"
//  INVALIDPARAMETERVALUE_ZONENAMENOTSUPPORTPUNYCODE = "InvalidParameterValue.ZoneNameNotSupportPunyCode"
//  INVALIDPARAMETERVALUE_ZONENAMENOTSUPPORTSUBDOMAIN = "InvalidParameterValue.ZoneNameNotSupportSubDomain"
//  INVALIDPARAMETERVALUE_ZONESAMEASNAME = "InvalidParameterValue.ZoneSameAsName"
//  LIMITEXCEEDED_ZONEBINDPLAN = "LimitExceeded.ZoneBindPlan"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DOMAINISBLOCKED = "OperationDenied.DomainIsBlocked"
//  OPERATIONDENIED_RECORDISFORBIDDEN = "OperationDenied.RecordIsForbidden"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_ALIASDOMAIN = "ResourceInUse.AliasDomain"
//  RESOURCEINUSE_CNAME = "ResourceInUse.Cname"
//  RESOURCEINUSE_DNS = "ResourceInUse.Dns"
//  RESOURCEINUSE_GENERICHOST = "ResourceInUse.GenericHost"
//  RESOURCEINUSE_HOST = "ResourceInUse.Host"
//  RESOURCEINUSE_NS = "ResourceInUse.NS"
//  RESOURCEINUSE_OTHERS = "ResourceInUse.Others"
//  RESOURCEINUSE_OTHERSALIASDOMAIN = "ResourceInUse.OthersAliasDomain"
//  RESOURCEINUSE_OTHERSCNAME = "ResourceInUse.OthersCname"
//  RESOURCEINUSE_OTHERSHOST = "ResourceInUse.OthersHost"
//  RESOURCEINUSE_OTHERSNS = "ResourceInUse.OthersNS"
//  RESOURCEINUSE_SELFANDOTHERSCNAME = "ResourceInUse.SelfAndOthersCname"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) CreateZoneWithContext(ctx context.Context, request *CreateZoneRequest) (response *CreateZoneResponse, err error) {
    if request == nil {
        request = NewCreateZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateZoneResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccelerationDomainsRequest() (request *DeleteAccelerationDomainsRequest) {
    request = &DeleteAccelerationDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteAccelerationDomains")
    
    
    return
}

func NewDeleteAccelerationDomainsResponse() (response *DeleteAccelerationDomainsResponse) {
    response = &DeleteAccelerationDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAccelerationDomains
// This API is used to batch remove accelerated domain names.
//
// error code that may be returned:
//  INVALIDPARAMETER_ZONEISGRAYPUBLISHING = "InvalidParameter.ZoneIsGrayPublishing"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  INVALIDPARAMETERVALUE_INVALIDDOMAINSTATUS = "InvalidParameterValue.InvalidDomainStatus"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_DOMAINNOTFOUND = "ResourceUnavailable.DomainNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteAccelerationDomains(request *DeleteAccelerationDomainsRequest) (response *DeleteAccelerationDomainsResponse, err error) {
    return c.DeleteAccelerationDomainsWithContext(context.Background(), request)
}

// DeleteAccelerationDomains
// This API is used to batch remove accelerated domain names.
//
// error code that may be returned:
//  INVALIDPARAMETER_ZONEISGRAYPUBLISHING = "InvalidParameter.ZoneIsGrayPublishing"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  INVALIDPARAMETERVALUE_INVALIDDOMAINSTATUS = "InvalidParameterValue.InvalidDomainStatus"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_DOMAINNOTFOUND = "ResourceUnavailable.DomainNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteAccelerationDomainsWithContext(ctx context.Context, request *DeleteAccelerationDomainsRequest) (response *DeleteAccelerationDomainsResponse, err error) {
    if request == nil {
        request = NewDeleteAccelerationDomainsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAccelerationDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAccelerationDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAliasDomainRequest() (request *DeleteAliasDomainRequest) {
    request = &DeleteAliasDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteAliasDomain")
    
    
    return
}

func NewDeleteAliasDomainResponse() (response *DeleteAliasDomainResponse) {
    response = &DeleteAliasDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAliasDomain
// This API is used to delete an alias domain name.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteAliasDomain(request *DeleteAliasDomainRequest) (response *DeleteAliasDomainResponse, err error) {
    return c.DeleteAliasDomainWithContext(context.Background(), request)
}

// DeleteAliasDomain
// This API is used to delete an alias domain name.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteAliasDomainWithContext(ctx context.Context, request *DeleteAliasDomainRequest) (response *DeleteAliasDomainResponse, err error) {
    if request == nil {
        request = NewDeleteAliasDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAliasDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAliasDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApplicationProxyRequest() (request *DeleteApplicationProxyRequest) {
    request = &DeleteApplicationProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteApplicationProxy")
    
    
    return
}

func NewDeleteApplicationProxyResponse() (response *DeleteApplicationProxyResponse) {
    response = &DeleteApplicationProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteApplicationProxy
// This API is on an earlier version. If you want to call it, please switch to the latest version. For details, see [DeleteL4Proxy] (https://intl.cloud.tencent.com/document/product/1552/103415?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteApplicationProxy(request *DeleteApplicationProxyRequest) (response *DeleteApplicationProxyResponse, err error) {
    return c.DeleteApplicationProxyWithContext(context.Background(), request)
}

// DeleteApplicationProxy
// This API is on an earlier version. If you want to call it, please switch to the latest version. For details, see [DeleteL4Proxy] (https://intl.cloud.tencent.com/document/product/1552/103415?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteApplicationProxyWithContext(ctx context.Context, request *DeleteApplicationProxyRequest) (response *DeleteApplicationProxyResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApplicationProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApplicationProxyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApplicationProxyRuleRequest() (request *DeleteApplicationProxyRuleRequest) {
    request = &DeleteApplicationProxyRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteApplicationProxyRule")
    
    
    return
}

func NewDeleteApplicationProxyRuleResponse() (response *DeleteApplicationProxyRuleResponse) {
    response = &DeleteApplicationProxyRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteApplicationProxyRule
// This API is on an earlier version. If you want to call it, please switch to the latest version. For details, see [DeleteL4ProxyRules] (https://intl.cloud.tencent.com/document/product/1552/103414?from_cn_redirect=1).
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteApplicationProxyRule(request *DeleteApplicationProxyRuleRequest) (response *DeleteApplicationProxyRuleResponse, err error) {
    return c.DeleteApplicationProxyRuleWithContext(context.Background(), request)
}

// DeleteApplicationProxyRule
// This API is on an earlier version. If you want to call it, please switch to the latest version. For details, see [DeleteL4ProxyRules] (https://intl.cloud.tencent.com/document/product/1552/103414?from_cn_redirect=1).
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteApplicationProxyRuleWithContext(ctx context.Context, request *DeleteApplicationProxyRuleRequest) (response *DeleteApplicationProxyRuleResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationProxyRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApplicationProxyRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApplicationProxyRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteL4ProxyRequest() (request *DeleteL4ProxyRequest) {
    request = &DeleteL4ProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteL4Proxy")
    
    
    return
}

func NewDeleteL4ProxyResponse() (response *DeleteL4ProxyResponse) {
    response = &DeleteL4ProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteL4Proxy
// This API is used to delete a Layer 4 proxy instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteL4Proxy(request *DeleteL4ProxyRequest) (response *DeleteL4ProxyResponse, err error) {
    return c.DeleteL4ProxyWithContext(context.Background(), request)
}

// DeleteL4Proxy
// This API is used to delete a Layer 4 proxy instance.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteL4ProxyWithContext(ctx context.Context, request *DeleteL4ProxyRequest) (response *DeleteL4ProxyResponse, err error) {
    if request == nil {
        request = NewDeleteL4ProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteL4Proxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteL4ProxyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteL4ProxyRulesRequest() (request *DeleteL4ProxyRulesRequest) {
    request = &DeleteL4ProxyRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteL4ProxyRules")
    
    
    return
}

func NewDeleteL4ProxyRulesResponse() (response *DeleteL4ProxyRulesResponse) {
    response = &DeleteL4ProxyRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteL4ProxyRules
// This API is used to delete Layer 4 proxy forwarding rules, supporting both individual and batch operation.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteL4ProxyRules(request *DeleteL4ProxyRulesRequest) (response *DeleteL4ProxyRulesResponse, err error) {
    return c.DeleteL4ProxyRulesWithContext(context.Background(), request)
}

// DeleteL4ProxyRules
// This API is used to delete Layer 4 proxy forwarding rules, supporting both individual and batch operation.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteL4ProxyRulesWithContext(ctx context.Context, request *DeleteL4ProxyRulesRequest) (response *DeleteL4ProxyRulesResponse, err error) {
    if request == nil {
        request = NewDeleteL4ProxyRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteL4ProxyRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteL4ProxyRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOriginGroupRequest() (request *DeleteOriginGroupRequest) {
    request = &DeleteOriginGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteOriginGroup")
    
    
    return
}

func NewDeleteOriginGroupResponse() (response *DeleteOriginGroupResponse) {
    response = &DeleteOriginGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteOriginGroup
// This API is used to delete an origin group. Note that an origin group can not be deleted if it is referenced by services (e.g. L4 Proxy, domain name service, load balancing, rule engines). 
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ORIGINGROUPACCELERATIONDOMAINUSED = "OperationDenied.OriginGroupAccelerationDomainUsed"
//  OPERATIONDENIED_ORIGINGROUPL4USED = "OperationDenied.OriginGroupL4Used"
//  OPERATIONDENIED_ORIGINGROUPLBUSED = "OperationDenied.OriginGroupLBUsed"
//  OPERATIONDENIED_ORIGINGROUPRULEENGINEUSED = "OperationDenied.OriginGroupRuleEngineUsed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteOriginGroup(request *DeleteOriginGroupRequest) (response *DeleteOriginGroupResponse, err error) {
    return c.DeleteOriginGroupWithContext(context.Background(), request)
}

// DeleteOriginGroup
// This API is used to delete an origin group. Note that an origin group can not be deleted if it is referenced by services (e.g. L4 Proxy, domain name service, load balancing, rule engines). 
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ORIGINGROUPACCELERATIONDOMAINUSED = "OperationDenied.OriginGroupAccelerationDomainUsed"
//  OPERATIONDENIED_ORIGINGROUPL4USED = "OperationDenied.OriginGroupL4Used"
//  OPERATIONDENIED_ORIGINGROUPLBUSED = "OperationDenied.OriginGroupLBUsed"
//  OPERATIONDENIED_ORIGINGROUPRULEENGINEUSED = "OperationDenied.OriginGroupRuleEngineUsed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteOriginGroupWithContext(ctx context.Context, request *DeleteOriginGroupRequest) (response *DeleteOriginGroupResponse, err error) {
    if request == nil {
        request = NewDeleteOriginGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOriginGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOriginGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRealtimeLogDeliveryTaskRequest() (request *DeleteRealtimeLogDeliveryTaskRequest) {
    request = &DeleteRealtimeLogDeliveryTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteRealtimeLogDeliveryTask")
    
    
    return
}

func NewDeleteRealtimeLogDeliveryTaskResponse() (response *DeleteRealtimeLogDeliveryTaskResponse) {
    response = &DeleteRealtimeLogDeliveryTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRealtimeLogDeliveryTask
// This API is used to delete a real-time log delivery task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteRealtimeLogDeliveryTask(request *DeleteRealtimeLogDeliveryTaskRequest) (response *DeleteRealtimeLogDeliveryTaskResponse, err error) {
    return c.DeleteRealtimeLogDeliveryTaskWithContext(context.Background(), request)
}

// DeleteRealtimeLogDeliveryTask
// This API is used to delete a real-time log delivery task.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteRealtimeLogDeliveryTaskWithContext(ctx context.Context, request *DeleteRealtimeLogDeliveryTaskRequest) (response *DeleteRealtimeLogDeliveryTaskResponse, err error) {
    if request == nil {
        request = NewDeleteRealtimeLogDeliveryTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRealtimeLogDeliveryTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRealtimeLogDeliveryTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRulesRequest() (request *DeleteRulesRequest) {
    request = &DeleteRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteRules")
    
    
    return
}

func NewDeleteRulesResponse() (response *DeleteRulesResponse) {
    response = &DeleteRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRules
// This API is used to batch delete rules from the rule engine.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteRules(request *DeleteRulesRequest) (response *DeleteRulesResponse, err error) {
    return c.DeleteRulesWithContext(context.Background(), request)
}

// DeleteRules
// This API is used to batch delete rules from the rule engine.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteRulesWithContext(ctx context.Context, request *DeleteRulesRequest) (response *DeleteRulesResponse, err error) {
    if request == nil {
        request = NewDeleteRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecurityIPGroupRequest() (request *DeleteSecurityIPGroupRequest) {
    request = &DeleteSecurityIPGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteSecurityIPGroup")
    
    
    return
}

func NewDeleteSecurityIPGroupResponse() (response *DeleteSecurityIPGroupResponse) {
    response = &DeleteSecurityIPGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSecurityIPGroup
// This API is used to delete a specified security IP group. Note that the security IP group cannot be deleted if it is referenced in a rule.
//
// error code that may be returned:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteSecurityIPGroup(request *DeleteSecurityIPGroupRequest) (response *DeleteSecurityIPGroupResponse, err error) {
    return c.DeleteSecurityIPGroupWithContext(context.Background(), request)
}

// DeleteSecurityIPGroup
// This API is used to delete a specified security IP group. Note that the security IP group cannot be deleted if it is referenced in a rule.
//
// error code that may be returned:
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteSecurityIPGroupWithContext(ctx context.Context, request *DeleteSecurityIPGroupRequest) (response *DeleteSecurityIPGroupResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityIPGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSecurityIPGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSecurityIPGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSharedCNAMERequest() (request *DeleteSharedCNAMERequest) {
    request = &DeleteSharedCNAMERequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteSharedCNAME")
    
    
    return
}

func NewDeleteSharedCNAMEResponse() (response *DeleteSharedCNAMEResponse) {
    response = &DeleteSharedCNAMEResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSharedCNAME
// This API is used to delete a shared CNAME. It is now only available to beta users.
//
// error code that may be returned:
//  RESOURCEINUSE_SHAREDCNAME = "ResourceInUse.SharedCNAME"
func (c *Client) DeleteSharedCNAME(request *DeleteSharedCNAMERequest) (response *DeleteSharedCNAMEResponse, err error) {
    return c.DeleteSharedCNAMEWithContext(context.Background(), request)
}

// DeleteSharedCNAME
// This API is used to delete a shared CNAME. It is now only available to beta users.
//
// error code that may be returned:
//  RESOURCEINUSE_SHAREDCNAME = "ResourceInUse.SharedCNAME"
func (c *Client) DeleteSharedCNAMEWithContext(ctx context.Context, request *DeleteSharedCNAMERequest) (response *DeleteSharedCNAMEResponse, err error) {
    if request == nil {
        request = NewDeleteSharedCNAMERequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSharedCNAME require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSharedCNAMEResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteZoneRequest() (request *DeleteZoneRequest) {
    request = &DeleteZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeleteZone")
    
    
    return
}

func NewDeleteZoneResponse() (response *DeleteZoneResponse) {
    response = &DeleteZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteZone
// This API is used to delete a site.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DELETEZONEPRECHECKFAILED = "OperationDenied.DeleteZonePreCheckFailed"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ZONEISBINDINGSHAREDCNAME = "OperationDenied.ZoneIsBindingSharedCNAME"
//  OPERATIONDENIED_ZONEISREFERENCECUSTOMERRORPAGE = "OperationDenied.ZoneIsReferenceCustomErrorPage"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteZone(request *DeleteZoneRequest) (response *DeleteZoneResponse, err error) {
    return c.DeleteZoneWithContext(context.Background(), request)
}

// DeleteZone
// This API is used to delete a site.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DELETEZONEPRECHECKFAILED = "OperationDenied.DeleteZonePreCheckFailed"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ZONEISBINDINGSHAREDCNAME = "OperationDenied.ZoneIsBindingSharedCNAME"
//  OPERATIONDENIED_ZONEISREFERENCECUSTOMERRORPAGE = "OperationDenied.ZoneIsReferenceCustomErrorPage"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DeleteZoneWithContext(ctx context.Context, request *DeleteZoneRequest) (response *DeleteZoneResponse, err error) {
    if request == nil {
        request = NewDeleteZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteZoneResponse()
    err = c.Send(request, response)
    return
}

func NewDeployConfigGroupVersionRequest() (request *DeployConfigGroupVersionRequest) {
    request = &DeployConfigGroupVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DeployConfigGroupVersion")
    
    
    return
}

func NewDeployConfigGroupVersionResponse() (response *DeployConfigGroupVersionResponse) {
    response = &DeployConfigGroupVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeployConfigGroupVersion
// This API is used to release versions in version management mode. Users can deploy the version to either the testing environment or the production environment by specifying the EnvId parameter. The version management feature is currently undergoing beta testing and is accessible only to users on the whitelist.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFIGCONDITIONSYNTAXERROR = "FailedOperation.ConfigConditionSyntaxError"
//  FAILEDOPERATION_CONFIGCONDITIONUNKNOWNTARGET = "FailedOperation.ConfigConditionUnknownTarget"
//  FAILEDOPERATION_CONFIGCONDITIONVALUEEMPTYERROR = "FailedOperation.ConfigConditionValueEmptyError"
//  FAILEDOPERATION_CONFIGFIELDTYPEERROR = "FailedOperation.ConfigFieldTypeError"
//  FAILEDOPERATION_CONFIGFORMATERROR = "FailedOperation.ConfigFormatError"
//  FAILEDOPERATION_CONFIGMALFORMEDCONTENT = "FailedOperation.ConfigMalformedContent"
//  FAILEDOPERATION_CONFIGPARAMVALIDATEERRORS = "FailedOperation.ConfigParamValidateErrors"
//  FAILEDOPERATION_CONFIGUNKNOWNFIELD = "FailedOperation.ConfigUnknownField"
//  FAILEDOPERATION_CONFIGUNSUPPORTEDFORMATVERSION = "FailedOperation.ConfigUnsupportedFormatVersion"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DOMAINSTATUSUNSTABLE = "OperationDenied.DomainStatusUnstable"
//  OPERATIONDENIED_ENVNOTREADY = "OperationDenied.EnvNotReady"
//  OPERATIONDENIED_NOTINVERSIONCONTROLWHITELIST = "OperationDenied.NotInVersionControlWhiteList"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  OPERATIONDENIED_WORKMODENOTINVERSIONCONTROL = "OperationDenied.WorkModeNotInVersionControl"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_VERSIONNOTFOUND = "ResourceNotFound.VersionNotFound"
func (c *Client) DeployConfigGroupVersion(request *DeployConfigGroupVersionRequest) (response *DeployConfigGroupVersionResponse, err error) {
    return c.DeployConfigGroupVersionWithContext(context.Background(), request)
}

// DeployConfigGroupVersion
// This API is used to release versions in version management mode. Users can deploy the version to either the testing environment or the production environment by specifying the EnvId parameter. The version management feature is currently undergoing beta testing and is accessible only to users on the whitelist.
//
// error code that may be returned:
//  FAILEDOPERATION_CONFIGCONDITIONSYNTAXERROR = "FailedOperation.ConfigConditionSyntaxError"
//  FAILEDOPERATION_CONFIGCONDITIONUNKNOWNTARGET = "FailedOperation.ConfigConditionUnknownTarget"
//  FAILEDOPERATION_CONFIGCONDITIONVALUEEMPTYERROR = "FailedOperation.ConfigConditionValueEmptyError"
//  FAILEDOPERATION_CONFIGFIELDTYPEERROR = "FailedOperation.ConfigFieldTypeError"
//  FAILEDOPERATION_CONFIGFORMATERROR = "FailedOperation.ConfigFormatError"
//  FAILEDOPERATION_CONFIGMALFORMEDCONTENT = "FailedOperation.ConfigMalformedContent"
//  FAILEDOPERATION_CONFIGPARAMVALIDATEERRORS = "FailedOperation.ConfigParamValidateErrors"
//  FAILEDOPERATION_CONFIGUNKNOWNFIELD = "FailedOperation.ConfigUnknownField"
//  FAILEDOPERATION_CONFIGUNSUPPORTEDFORMATVERSION = "FailedOperation.ConfigUnsupportedFormatVersion"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DOMAINSTATUSUNSTABLE = "OperationDenied.DomainStatusUnstable"
//  OPERATIONDENIED_ENVNOTREADY = "OperationDenied.EnvNotReady"
//  OPERATIONDENIED_NOTINVERSIONCONTROLWHITELIST = "OperationDenied.NotInVersionControlWhiteList"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  OPERATIONDENIED_WORKMODENOTINVERSIONCONTROL = "OperationDenied.WorkModeNotInVersionControl"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_VERSIONNOTFOUND = "ResourceNotFound.VersionNotFound"
func (c *Client) DeployConfigGroupVersionWithContext(ctx context.Context, request *DeployConfigGroupVersionRequest) (response *DeployConfigGroupVersionResponse, err error) {
    if request == nil {
        request = NewDeployConfigGroupVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeployConfigGroupVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeployConfigGroupVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccelerationDomainsRequest() (request *DescribeAccelerationDomainsRequest) {
    request = &DescribeAccelerationDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeAccelerationDomains")
    
    
    return
}

func NewDescribeAccelerationDomainsResponse() (response *DescribeAccelerationDomainsResponse) {
    response = &DescribeAccelerationDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccelerationDomains
// This API is used to query domain name information of a site, including the acceleration domain name, origin, and domain name status. You can query the information of all domain names, or specific domain names by specifying filters information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_DOMAINONTRAFFICSCHEDULING = "InvalidParameter.DomainOnTrafficScheduling"
//  INVALIDPARAMETER_TOOMANYFILTERVALUES = "InvalidParameter.TooManyFilterValues"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeAccelerationDomains(request *DescribeAccelerationDomainsRequest) (response *DescribeAccelerationDomainsResponse, err error) {
    return c.DescribeAccelerationDomainsWithContext(context.Background(), request)
}

// DescribeAccelerationDomains
// This API is used to query domain name information of a site, including the acceleration domain name, origin, and domain name status. You can query the information of all domain names, or specific domain names by specifying filters information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_DOMAINONTRAFFICSCHEDULING = "InvalidParameter.DomainOnTrafficScheduling"
//  INVALIDPARAMETER_TOOMANYFILTERVALUES = "InvalidParameter.TooManyFilterValues"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeAccelerationDomainsWithContext(ctx context.Context, request *DescribeAccelerationDomainsRequest) (response *DescribeAccelerationDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeAccelerationDomainsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccelerationDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccelerationDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAliasDomainsRequest() (request *DescribeAliasDomainsRequest) {
    request = &DescribeAliasDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeAliasDomains")
    
    
    return
}

func NewDescribeAliasDomainsResponse() (response *DescribeAliasDomainsResponse) {
    response = &DescribeAliasDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAliasDomains
// This API is used to query the information of alias domain names.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeAliasDomains(request *DescribeAliasDomainsRequest) (response *DescribeAliasDomainsResponse, err error) {
    return c.DescribeAliasDomainsWithContext(context.Background(), request)
}

// DescribeAliasDomains
// This API is used to query the information of alias domain names.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeAliasDomainsWithContext(ctx context.Context, request *DescribeAliasDomainsRequest) (response *DescribeAliasDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeAliasDomainsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAliasDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAliasDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationProxiesRequest() (request *DescribeApplicationProxiesRequest) {
    request = &DescribeApplicationProxiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeApplicationProxies")
    
    
    return
}

func NewDescribeApplicationProxiesResponse() (response *DescribeApplicationProxiesResponse) {
    response = &DescribeApplicationProxiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApplicationProxies
// This API is on an earlier version. If you want to call it, please switch to the latest version. In the latest version, this API has been split into two APIs: one for querying the Layer 4 proxy instance list and the other for querying Layer 4 forwarding rules. For details, see [DescribeL4Proxy] (https://intl.cloud.tencent.com/document/product/1552/103413?from_cn_redirect=1) and [DescribeL4ProxyRules] (https://intl.cloud.tencent.com/document/product/1552/103412?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeApplicationProxies(request *DescribeApplicationProxiesRequest) (response *DescribeApplicationProxiesResponse, err error) {
    return c.DescribeApplicationProxiesWithContext(context.Background(), request)
}

// DescribeApplicationProxies
// This API is on an earlier version. If you want to call it, please switch to the latest version. In the latest version, this API has been split into two APIs: one for querying the Layer 4 proxy instance list and the other for querying Layer 4 forwarding rules. For details, see [DescribeL4Proxy] (https://intl.cloud.tencent.com/document/product/1552/103413?from_cn_redirect=1) and [DescribeL4ProxyRules] (https://intl.cloud.tencent.com/document/product/1552/103412?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeApplicationProxiesWithContext(ctx context.Context, request *DescribeApplicationProxiesRequest) (response *DescribeApplicationProxiesResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationProxiesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationProxies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationProxiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAvailablePlansRequest() (request *DescribeAvailablePlansRequest) {
    request = &DescribeAvailablePlansRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeAvailablePlans")
    
    
    return
}

func NewDescribeAvailablePlansResponse() (response *DescribeAvailablePlansResponse) {
    response = &DescribeAvailablePlansResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAvailablePlans
// This API is used to query plan options available for purchase.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
func (c *Client) DescribeAvailablePlans(request *DescribeAvailablePlansRequest) (response *DescribeAvailablePlansResponse, err error) {
    return c.DescribeAvailablePlansWithContext(context.Background(), request)
}

// DescribeAvailablePlans
// This API is used to query plan options available for purchase.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
func (c *Client) DescribeAvailablePlansWithContext(ctx context.Context, request *DescribeAvailablePlansRequest) (response *DescribeAvailablePlansResponse, err error) {
    if request == nil {
        request = NewDescribeAvailablePlansRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAvailablePlans require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAvailablePlansResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillingDataRequest() (request *DescribeBillingDataRequest) {
    request = &DescribeBillingDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeBillingData")
    
    
    return
}

func NewDescribeBillingDataResponse() (response *DescribeBillingDataResponse) {
    response = &DescribeBillingDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillingData
// This API is used to query billing data.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_INVALIDINTERVAL = "InvalidParameter.InvalidInterval"
//  INVALIDPARAMETER_INVALIDMETRIC = "InvalidParameter.InvalidMetric"
//  INVALIDPARAMETER_ZONEHASNOTBEENBOUNDTOPLAN = "InvalidParameter.ZoneHasNotBeenBoundToPlan"
func (c *Client) DescribeBillingData(request *DescribeBillingDataRequest) (response *DescribeBillingDataResponse, err error) {
    return c.DescribeBillingDataWithContext(context.Background(), request)
}

// DescribeBillingData
// This API is used to query billing data.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_INVALIDINTERVAL = "InvalidParameter.InvalidInterval"
//  INVALIDPARAMETER_INVALIDMETRIC = "InvalidParameter.InvalidMetric"
//  INVALIDPARAMETER_ZONEHASNOTBEENBOUNDTOPLAN = "InvalidParameter.ZoneHasNotBeenBoundToPlan"
func (c *Client) DescribeBillingDataWithContext(ctx context.Context, request *DescribeBillingDataRequest) (response *DescribeBillingDataResponse, err error) {
    if request == nil {
        request = NewDescribeBillingDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillingData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillingDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigGroupVersionDetailRequest() (request *DescribeConfigGroupVersionDetailRequest) {
    request = &DescribeConfigGroupVersionDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeConfigGroupVersionDetail")
    
    
    return
}

func NewDescribeConfigGroupVersionDetailResponse() (response *DescribeConfigGroupVersionDetailResponse) {
    response = &DescribeConfigGroupVersionDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConfigGroupVersionDetail
// This API is used to obtain detailed information about a version in version management mode. The response includes the version ID, description, status, creation time, configuration group information, and the content of the version configuration file. The version management feature is currently undergoing beta testing and is accessible only to users on the whitelist.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeConfigGroupVersionDetail(request *DescribeConfigGroupVersionDetailRequest) (response *DescribeConfigGroupVersionDetailResponse, err error) {
    return c.DescribeConfigGroupVersionDetailWithContext(context.Background(), request)
}

// DescribeConfigGroupVersionDetail
// This API is used to obtain detailed information about a version in version management mode. The response includes the version ID, description, status, creation time, configuration group information, and the content of the version configuration file. The version management feature is currently undergoing beta testing and is accessible only to users on the whitelist.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeConfigGroupVersionDetailWithContext(ctx context.Context, request *DescribeConfigGroupVersionDetailRequest) (response *DescribeConfigGroupVersionDetailResponse, err error) {
    if request == nil {
        request = NewDescribeConfigGroupVersionDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigGroupVersionDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigGroupVersionDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigGroupVersionsRequest() (request *DescribeConfigGroupVersionsRequest) {
    request = &DescribeConfigGroupVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeConfigGroupVersions")
    
    
    return
}

func NewDescribeConfigGroupVersionsResponse() (response *DescribeConfigGroupVersionsResponse) {
    response = &DescribeConfigGroupVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConfigGroupVersions
// This API is used to query the version list for the specified configuration group in version management mode. The version management feature is currently undergoing beta testing and is accessible only to users on the whitelist.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeConfigGroupVersions(request *DescribeConfigGroupVersionsRequest) (response *DescribeConfigGroupVersionsResponse, err error) {
    return c.DescribeConfigGroupVersionsWithContext(context.Background(), request)
}

// DescribeConfigGroupVersions
// This API is used to query the version list for the specified configuration group in version management mode. The version management feature is currently undergoing beta testing and is accessible only to users on the whitelist.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeConfigGroupVersionsWithContext(ctx context.Context, request *DescribeConfigGroupVersionsRequest) (response *DescribeConfigGroupVersionsResponse, err error) {
    if request == nil {
        request = NewDescribeConfigGroupVersionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConfigGroupVersions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConfigGroupVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContentQuotaRequest() (request *DescribeContentQuotaRequest) {
    request = &DescribeContentQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeContentQuota")
    
    
    return
}

func NewDescribeContentQuotaResponse() (response *DescribeContentQuotaResponse) {
    response = &DescribeContentQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeContentQuota
// This API is used to query content management quotas.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeContentQuota(request *DescribeContentQuotaRequest) (response *DescribeContentQuotaResponse, err error) {
    return c.DescribeContentQuotaWithContext(context.Background(), request)
}

// DescribeContentQuota
// This API is used to query content management quotas.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeContentQuotaWithContext(ctx context.Context, request *DescribeContentQuotaRequest) (response *DescribeContentQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeContentQuotaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeContentQuota require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeContentQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSAttackDataRequest() (request *DescribeDDoSAttackDataRequest) {
    request = &DescribeDDoSAttackDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDDoSAttackData")
    
    
    return
}

func NewDescribeDDoSAttackDataResponse() (response *DescribeDDoSAttackDataResponse) {
    response = &DescribeDDoSAttackDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDDoSAttackData
// This API is used to query the time-series data of DDoS attacks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeDDoSAttackData(request *DescribeDDoSAttackDataRequest) (response *DescribeDDoSAttackDataResponse, err error) {
    return c.DescribeDDoSAttackDataWithContext(context.Background(), request)
}

// DescribeDDoSAttackData
// This API is used to query the time-series data of DDoS attacks.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeDDoSAttackDataWithContext(ctx context.Context, request *DescribeDDoSAttackDataRequest) (response *DescribeDDoSAttackDataResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSAttackDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSAttackData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDoSAttackDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSAttackEventRequest() (request *DescribeDDoSAttackEventRequest) {
    request = &DescribeDDoSAttackEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDDoSAttackEvent")
    
    
    return
}

func NewDescribeDDoSAttackEventResponse() (response *DescribeDDoSAttackEventResponse) {
    response = &DescribeDDoSAttackEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDDoSAttackEvent
// This API is used to query DDoS attack events.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDDoSAttackEvent(request *DescribeDDoSAttackEventRequest) (response *DescribeDDoSAttackEventResponse, err error) {
    return c.DescribeDDoSAttackEventWithContext(context.Background(), request)
}

// DescribeDDoSAttackEvent
// This API is used to query DDoS attack events.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDDoSAttackEventWithContext(ctx context.Context, request *DescribeDDoSAttackEventRequest) (response *DescribeDDoSAttackEventResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSAttackEventRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSAttackEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDoSAttackEventResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSAttackTopDataRequest() (request *DescribeDDoSAttackTopDataRequest) {
    request = &DescribeDDoSAttackTopDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDDoSAttackTopData")
    
    
    return
}

func NewDescribeDDoSAttackTopDataResponse() (response *DescribeDDoSAttackTopDataResponse) {
    response = &DescribeDDoSAttackTopDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDDoSAttackTopData
// This API is used to query the top-ranked DDoS attack data.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDDoSAttackTopData(request *DescribeDDoSAttackTopDataRequest) (response *DescribeDDoSAttackTopDataResponse, err error) {
    return c.DescribeDDoSAttackTopDataWithContext(context.Background(), request)
}

// DescribeDDoSAttackTopData
// This API is used to query the top-ranked DDoS attack data.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDDoSAttackTopDataWithContext(ctx context.Context, request *DescribeDDoSAttackTopDataRequest) (response *DescribeDDoSAttackTopDataResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSAttackTopDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSAttackTopData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDoSAttackTopDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDefaultCertificatesRequest() (request *DescribeDefaultCertificatesRequest) {
    request = &DescribeDefaultCertificatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDefaultCertificates")
    
    
    return
}

func NewDescribeDefaultCertificatesResponse() (response *DescribeDefaultCertificatesResponse) {
    response = &DescribeDefaultCertificatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDefaultCertificates
// This API is used to query a list of default certificates.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeDefaultCertificates(request *DescribeDefaultCertificatesRequest) (response *DescribeDefaultCertificatesResponse, err error) {
    return c.DescribeDefaultCertificatesWithContext(context.Background(), request)
}

// DescribeDefaultCertificates
// This API is used to query a list of default certificates.
//
// error code that may be returned:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeDefaultCertificatesWithContext(ctx context.Context, request *DescribeDefaultCertificatesRequest) (response *DescribeDefaultCertificatesResponse, err error) {
    if request == nil {
        request = NewDescribeDefaultCertificatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDefaultCertificates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDefaultCertificatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeployHistoryRequest() (request *DescribeDeployHistoryRequest) {
    request = &DescribeDeployHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDeployHistory")
    
    
    return
}

func NewDescribeDeployHistoryResponse() (response *DescribeDeployHistoryResponse) {
    response = &DescribeDeployHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeployHistory
// This API is used to query the release history of versions in the production or test environment in version management mode. The version management feature is currently undergoing beta testing and is accessible only to users on the whitelist.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDeployHistory(request *DescribeDeployHistoryRequest) (response *DescribeDeployHistoryResponse, err error) {
    return c.DescribeDeployHistoryWithContext(context.Background(), request)
}

// DescribeDeployHistory
// This API is used to query the release history of versions in the production or test environment in version management mode. The version management feature is currently undergoing beta testing and is accessible only to users on the whitelist.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDeployHistoryWithContext(ctx context.Context, request *DescribeDeployHistoryRequest) (response *DescribeDeployHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeDeployHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeployHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeployHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvironmentsRequest() (request *DescribeEnvironmentsRequest) {
    request = &DescribeEnvironmentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeEnvironments")
    
    
    return
}

func NewDescribeEnvironmentsResponse() (response *DescribeEnvironmentsResponse) {
    response = &DescribeEnvironmentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEnvironments
// This API is used to query environment information in version management mode. The response includes the environment ID, type, and current effective version. The version management feature is currently undergoing beta testing and is accessible only to users on the whitelist.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEnvironments(request *DescribeEnvironmentsRequest) (response *DescribeEnvironmentsResponse, err error) {
    return c.DescribeEnvironmentsWithContext(context.Background(), request)
}

// DescribeEnvironments
// This API is used to query environment information in version management mode. The response includes the environment ID, type, and current effective version. The version management feature is currently undergoing beta testing and is accessible only to users on the whitelist.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEnvironmentsWithContext(ctx context.Context, request *DescribeEnvironmentsRequest) (response *DescribeEnvironmentsResponse, err error) {
    if request == nil {
        request = NewDescribeEnvironmentsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnvironments require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnvironmentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostsSettingRequest() (request *DescribeHostsSettingRequest) {
    request = &DescribeHostsSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeHostsSetting")
    
    
    return
}

func NewDescribeHostsSettingResponse() (response *DescribeHostsSettingResponse) {
    response = &DescribeHostsSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHostsSetting
// This API is used to query detailed domain name configuration.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeHostsSetting(request *DescribeHostsSettingRequest) (response *DescribeHostsSettingResponse, err error) {
    return c.DescribeHostsSettingWithContext(context.Background(), request)
}

// DescribeHostsSetting
// This API is used to query detailed domain name configuration.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeHostsSettingWithContext(ctx context.Context, request *DescribeHostsSettingRequest) (response *DescribeHostsSettingResponse, err error) {
    if request == nil {
        request = NewDescribeHostsSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostsSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostsSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIPRegionRequest() (request *DescribeIPRegionRequest) {
    request = &DescribeIPRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeIPRegion")
    
    
    return
}

func NewDescribeIPRegionResponse() (response *DescribeIPRegionResponse) {
    response = &DescribeIPRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIPRegion
// This API is used to check if the IP is an EdgeOne IP.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
func (c *Client) DescribeIPRegion(request *DescribeIPRegionRequest) (response *DescribeIPRegionResponse, err error) {
    return c.DescribeIPRegionWithContext(context.Background(), request)
}

// DescribeIPRegion
// This API is used to check if the IP is an EdgeOne IP.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
func (c *Client) DescribeIPRegionWithContext(ctx context.Context, request *DescribeIPRegionRequest) (response *DescribeIPRegionResponse, err error) {
    if request == nil {
        request = NewDescribeIPRegionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIPRegion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIPRegionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIdentificationsRequest() (request *DescribeIdentificationsRequest) {
    request = &DescribeIdentificationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeIdentifications")
    
    
    return
}

func NewDescribeIdentificationsResponse() (response *DescribeIdentificationsResponse) {
    response = &DescribeIdentificationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIdentifications
// This API is used to query the verification information of a site.
//
// error code that may be returned:
//  OPERATIONDENIED_DOMAININSHARECNAMEGROUP = "OperationDenied.DomainInShareCnameGroup"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeIdentifications(request *DescribeIdentificationsRequest) (response *DescribeIdentificationsResponse, err error) {
    return c.DescribeIdentificationsWithContext(context.Background(), request)
}

// DescribeIdentifications
// This API is used to query the verification information of a site.
//
// error code that may be returned:
//  OPERATIONDENIED_DOMAININSHARECNAMEGROUP = "OperationDenied.DomainInShareCnameGroup"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeIdentificationsWithContext(ctx context.Context, request *DescribeIdentificationsRequest) (response *DescribeIdentificationsResponse, err error) {
    if request == nil {
        request = NewDescribeIdentificationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIdentifications require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIdentificationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeL4ProxyRequest() (request *DescribeL4ProxyRequest) {
    request = &DescribeL4ProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeL4Proxy")
    
    
    return
}

func NewDescribeL4ProxyResponse() (response *DescribeL4ProxyResponse) {
    response = &DescribeL4ProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeL4Proxy
// This API is used to query a Layer 4 proxy instance list.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeL4Proxy(request *DescribeL4ProxyRequest) (response *DescribeL4ProxyResponse, err error) {
    return c.DescribeL4ProxyWithContext(context.Background(), request)
}

// DescribeL4Proxy
// This API is used to query a Layer 4 proxy instance list.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeL4ProxyWithContext(ctx context.Context, request *DescribeL4ProxyRequest) (response *DescribeL4ProxyResponse, err error) {
    if request == nil {
        request = NewDescribeL4ProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeL4Proxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeL4ProxyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeL4ProxyRulesRequest() (request *DescribeL4ProxyRulesRequest) {
    request = &DescribeL4ProxyRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeL4ProxyRules")
    
    
    return
}

func NewDescribeL4ProxyRulesResponse() (response *DescribeL4ProxyRulesResponse) {
    response = &DescribeL4ProxyRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeL4ProxyRules
// This API is used to query the forwarding rule list under a Layer 4 proxy instance.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeL4ProxyRules(request *DescribeL4ProxyRulesRequest) (response *DescribeL4ProxyRulesResponse, err error) {
    return c.DescribeL4ProxyRulesWithContext(context.Background(), request)
}

// DescribeL4ProxyRules
// This API is used to query the forwarding rule list under a Layer 4 proxy instance.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeL4ProxyRulesWithContext(ctx context.Context, request *DescribeL4ProxyRulesRequest) (response *DescribeL4ProxyRulesResponse, err error) {
    if request == nil {
        request = NewDescribeL4ProxyRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeL4ProxyRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeL4ProxyRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOriginGroupRequest() (request *DescribeOriginGroupRequest) {
    request = &DescribeOriginGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeOriginGroup")
    
    
    return
}

func NewDescribeOriginGroupResponse() (response *DescribeOriginGroupResponse) {
    response = &DescribeOriginGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOriginGroup
// This API is used to obtain a list of origin groups.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeOriginGroup(request *DescribeOriginGroupRequest) (response *DescribeOriginGroupResponse, err error) {
    return c.DescribeOriginGroupWithContext(context.Background(), request)
}

// DescribeOriginGroup
// This API is used to obtain a list of origin groups.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeOriginGroupWithContext(ctx context.Context, request *DescribeOriginGroupRequest) (response *DescribeOriginGroupResponse, err error) {
    if request == nil {
        request = NewDescribeOriginGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOriginGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOriginGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOriginProtectionRequest() (request *DescribeOriginProtectionRequest) {
    request = &DescribeOriginProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeOriginProtection")
    
    
    return
}

func NewDescribeOriginProtectionResponse() (response *DescribeOriginProtectionResponse) {
    response = &DescribeOriginProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOriginProtection
// This API is used to query the origin protection configuration.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeOriginProtection(request *DescribeOriginProtectionRequest) (response *DescribeOriginProtectionResponse, err error) {
    return c.DescribeOriginProtectionWithContext(context.Background(), request)
}

// DescribeOriginProtection
// This API is used to query the origin protection configuration.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeOriginProtectionWithContext(ctx context.Context, request *DescribeOriginProtectionRequest) (response *DescribeOriginProtectionResponse, err error) {
    if request == nil {
        request = NewDescribeOriginProtectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOriginProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOriginProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOverviewL7DataRequest() (request *DescribeOverviewL7DataRequest) {
    request = &DescribeOverviewL7DataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeOverviewL7Data")
    
    
    return
}

func NewDescribeOverviewL7DataResponse() (response *DescribeOverviewL7DataResponse) {
    response = &DescribeOverviewL7DataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOverviewL7Data
// This API is used to query the time sequence traffic data of the monitoring category in L7. This API is to be discarded. Please use the API <a href="https://intl.cloud.tencent.com/document/product/1552/80648?from_cn_redirect=1">DescribeTimingL7AnalysisData</a>.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeOverviewL7Data(request *DescribeOverviewL7DataRequest) (response *DescribeOverviewL7DataResponse, err error) {
    return c.DescribeOverviewL7DataWithContext(context.Background(), request)
}

// DescribeOverviewL7Data
// This API is used to query the time sequence traffic data of the monitoring category in L7. This API is to be discarded. Please use the API <a href="https://intl.cloud.tencent.com/document/product/1552/80648?from_cn_redirect=1">DescribeTimingL7AnalysisData</a>.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeOverviewL7DataWithContext(ctx context.Context, request *DescribeOverviewL7DataRequest) (response *DescribeOverviewL7DataResponse, err error) {
    if request == nil {
        request = NewDescribeOverviewL7DataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOverviewL7Data require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOverviewL7DataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrefetchTasksRequest() (request *DescribePrefetchTasksRequest) {
    request = &DescribePrefetchTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribePrefetchTasks")
    
    
    return
}

func NewDescribePrefetchTasksResponse() (response *DescribePrefetchTasksResponse) {
    response = &DescribePrefetchTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePrefetchTasks
// DescribePrefetchTasks is used to query the submission history and execution progress of preheating tasks. This interface can be used to query the tasks submitted by the CreatePrefetchTasks interface.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_DOMAINEMPTY = "UnauthorizedOperation.DomainEmpty"
func (c *Client) DescribePrefetchTasks(request *DescribePrefetchTasksRequest) (response *DescribePrefetchTasksResponse, err error) {
    return c.DescribePrefetchTasksWithContext(context.Background(), request)
}

// DescribePrefetchTasks
// DescribePrefetchTasks is used to query the submission history and execution progress of preheating tasks. This interface can be used to query the tasks submitted by the CreatePrefetchTasks interface.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_DOMAINEMPTY = "UnauthorizedOperation.DomainEmpty"
func (c *Client) DescribePrefetchTasksWithContext(ctx context.Context, request *DescribePrefetchTasksRequest) (response *DescribePrefetchTasksResponse, err error) {
    if request == nil {
        request = NewDescribePrefetchTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrefetchTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrefetchTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePurgeTasksRequest() (request *DescribePurgeTasksRequest) {
    request = &DescribePurgeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribePurgeTasks")
    
    
    return
}

func NewDescribePurgeTasksResponse() (response *DescribePurgeTasksResponse) {
    response = &DescribePurgeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePurgeTasks
// DescribePurgeTasks is used to query the submitted URL refreshing and directory refreshing records and execution progress. This interface can be used to query the tasks submitted by the CreatePurgeTasks API.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribePurgeTasks(request *DescribePurgeTasksRequest) (response *DescribePurgeTasksResponse, err error) {
    return c.DescribePurgeTasksWithContext(context.Background(), request)
}

// DescribePurgeTasks
// DescribePurgeTasks is used to query the submitted URL refreshing and directory refreshing records and execution progress. This interface can be used to query the tasks submitted by the CreatePurgeTasks API.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribePurgeTasksWithContext(ctx context.Context, request *DescribePurgeTasksRequest) (response *DescribePurgeTasksResponse, err error) {
    if request == nil {
        request = NewDescribePurgeTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePurgeTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePurgeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRealtimeLogDeliveryTasksRequest() (request *DescribeRealtimeLogDeliveryTasksRequest) {
    request = &DescribeRealtimeLogDeliveryTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeRealtimeLogDeliveryTasks")
    
    
    return
}

func NewDescribeRealtimeLogDeliveryTasksResponse() (response *DescribeRealtimeLogDeliveryTasksResponse) {
    response = &DescribeRealtimeLogDeliveryTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRealtimeLogDeliveryTasks
// This API is used to query the real-time log delivery task list.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeRealtimeLogDeliveryTasks(request *DescribeRealtimeLogDeliveryTasksRequest) (response *DescribeRealtimeLogDeliveryTasksResponse, err error) {
    return c.DescribeRealtimeLogDeliveryTasksWithContext(context.Background(), request)
}

// DescribeRealtimeLogDeliveryTasks
// This API is used to query the real-time log delivery task list.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeRealtimeLogDeliveryTasksWithContext(ctx context.Context, request *DescribeRealtimeLogDeliveryTasksRequest) (response *DescribeRealtimeLogDeliveryTasksResponse, err error) {
    if request == nil {
        request = NewDescribeRealtimeLogDeliveryTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRealtimeLogDeliveryTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRealtimeLogDeliveryTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRulesRequest() (request *DescribeRulesRequest) {
    request = &DescribeRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeRules")
    
    
    return
}

func NewDescribeRulesResponse() (response *DescribeRulesResponse) {
    response = &DescribeRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRules
// This API is used to query the rules in the rule engine.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeRules(request *DescribeRulesRequest) (response *DescribeRulesResponse, err error) {
    return c.DescribeRulesWithContext(context.Background(), request)
}

// DescribeRules
// This API is used to query the rules in the rule engine.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeRulesWithContext(ctx context.Context, request *DescribeRulesRequest) (response *DescribeRulesResponse, err error) {
    if request == nil {
        request = NewDescribeRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRulesSettingRequest() (request *DescribeRulesSettingRequest) {
    request = &DescribeRulesSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeRulesSetting")
    
    
    return
}

func NewDescribeRulesSettingResponse() (response *DescribeRulesSettingResponse) {
    response = &DescribeRulesSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRulesSetting
// This API is used to return the list of the settings of the rule engine that can be used for request match and their detailed recommended configuration information.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeRulesSetting(request *DescribeRulesSettingRequest) (response *DescribeRulesSettingResponse, err error) {
    return c.DescribeRulesSettingWithContext(context.Background(), request)
}

// DescribeRulesSetting
// This API is used to return the list of the settings of the rule engine that can be used for request match and their detailed recommended configuration information.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeRulesSettingWithContext(ctx context.Context, request *DescribeRulesSettingRequest) (response *DescribeRulesSettingResponse, err error) {
    if request == nil {
        request = NewDescribeRulesSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRulesSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRulesSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityTemplateBindingsRequest() (request *DescribeSecurityTemplateBindingsRequest) {
    request = &DescribeSecurityTemplateBindingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeSecurityTemplateBindings")
    
    
    return
}

func NewDescribeSecurityTemplateBindingsResponse() (response *DescribeSecurityTemplateBindingsResponse) {
    response = &DescribeSecurityTemplateBindingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityTemplateBindings
// This API is used to query bindings of a policy template.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeSecurityTemplateBindings(request *DescribeSecurityTemplateBindingsRequest) (response *DescribeSecurityTemplateBindingsResponse, err error) {
    return c.DescribeSecurityTemplateBindingsWithContext(context.Background(), request)
}

// DescribeSecurityTemplateBindings
// This API is used to query bindings of a policy template.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeSecurityTemplateBindingsWithContext(ctx context.Context, request *DescribeSecurityTemplateBindingsRequest) (response *DescribeSecurityTemplateBindingsResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityTemplateBindingsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityTemplateBindings require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityTemplateBindingsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTimingL4DataRequest() (request *DescribeTimingL4DataRequest) {
    request = &DescribeTimingL4DataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeTimingL4Data")
    
    
    return
}

func NewDescribeTimingL4DataResponse() (response *DescribeTimingL4DataResponse) {
    response = &DescribeTimingL4DataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTimingL4Data
// This API is used to query the list of L4 traffic data recorded over time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTimingL4Data(request *DescribeTimingL4DataRequest) (response *DescribeTimingL4DataResponse, err error) {
    return c.DescribeTimingL4DataWithContext(context.Background(), request)
}

// DescribeTimingL4Data
// This API is used to query the list of L4 traffic data recorded over time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTimingL4DataWithContext(ctx context.Context, request *DescribeTimingL4DataRequest) (response *DescribeTimingL4DataResponse, err error) {
    if request == nil {
        request = NewDescribeTimingL4DataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTimingL4Data require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTimingL4DataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTimingL7AnalysisDataRequest() (request *DescribeTimingL7AnalysisDataRequest) {
    request = &DescribeTimingL7AnalysisDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeTimingL7AnalysisData")
    
    
    return
}

func NewDescribeTimingL7AnalysisDataResponse() (response *DescribeTimingL7AnalysisDataResponse) {
    response = &DescribeTimingL7AnalysisDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTimingL7AnalysisData
// This API is used to query the L7 data recorded over time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTimingL7AnalysisData(request *DescribeTimingL7AnalysisDataRequest) (response *DescribeTimingL7AnalysisDataResponse, err error) {
    return c.DescribeTimingL7AnalysisDataWithContext(context.Background(), request)
}

// DescribeTimingL7AnalysisData
// This API is used to query the L7 data recorded over time.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTimingL7AnalysisDataWithContext(ctx context.Context, request *DescribeTimingL7AnalysisDataRequest) (response *DescribeTimingL7AnalysisDataResponse, err error) {
    if request == nil {
        request = NewDescribeTimingL7AnalysisDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTimingL7AnalysisData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTimingL7AnalysisDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTimingL7CacheDataRequest() (request *DescribeTimingL7CacheDataRequest) {
    request = &DescribeTimingL7CacheDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeTimingL7CacheData")
    
    
    return
}

func NewDescribeTimingL7CacheDataResponse() (response *DescribeTimingL7CacheDataResponse) {
    response = &DescribeTimingL7CacheDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTimingL7CacheData
// This API is used to query the time-series L7 cached data.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTimingL7CacheData(request *DescribeTimingL7CacheDataRequest) (response *DescribeTimingL7CacheDataResponse, err error) {
    return c.DescribeTimingL7CacheDataWithContext(context.Background(), request)
}

// DescribeTimingL7CacheData
// This API is used to query the time-series L7 cached data.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTimingL7CacheDataWithContext(ctx context.Context, request *DescribeTimingL7CacheDataRequest) (response *DescribeTimingL7CacheDataResponse, err error) {
    if request == nil {
        request = NewDescribeTimingL7CacheDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTimingL7CacheData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTimingL7CacheDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopL7AnalysisDataRequest() (request *DescribeTopL7AnalysisDataRequest) {
    request = &DescribeTopL7AnalysisDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeTopL7AnalysisData")
    
    
    return
}

func NewDescribeTopL7AnalysisDataResponse() (response *DescribeTopL7AnalysisDataResponse) {
    response = &DescribeTopL7AnalysisDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopL7AnalysisData
// This API is used to query the top-ranked L7 traffic data.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeTopL7AnalysisData(request *DescribeTopL7AnalysisDataRequest) (response *DescribeTopL7AnalysisDataResponse, err error) {
    return c.DescribeTopL7AnalysisDataWithContext(context.Background(), request)
}

// DescribeTopL7AnalysisData
// This API is used to query the top-ranked L7 traffic data.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeTopL7AnalysisDataWithContext(ctx context.Context, request *DescribeTopL7AnalysisDataRequest) (response *DescribeTopL7AnalysisDataResponse, err error) {
    if request == nil {
        request = NewDescribeTopL7AnalysisDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopL7AnalysisData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopL7AnalysisDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopL7CacheDataRequest() (request *DescribeTopL7CacheDataRequest) {
    request = &DescribeTopL7CacheDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeTopL7CacheData")
    
    
    return
}

func NewDescribeTopL7CacheDataResponse() (response *DescribeTopL7CacheDataResponse) {
    response = &DescribeTopL7CacheDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopL7CacheData
// This API is used to query the cached L7 top-ranked data.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTopL7CacheData(request *DescribeTopL7CacheDataRequest) (response *DescribeTopL7CacheDataResponse, err error) {
    return c.DescribeTopL7CacheDataWithContext(context.Background(), request)
}

// DescribeTopL7CacheData
// This API is used to query the cached L7 top-ranked data.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED_QUERYTIMELIMITEXCEEDED = "LimitExceeded.QueryTimeLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeTopL7CacheDataWithContext(ctx context.Context, request *DescribeTopL7CacheDataRequest) (response *DescribeTopL7CacheDataResponse, err error) {
    if request == nil {
        request = NewDescribeTopL7CacheDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopL7CacheData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopL7CacheDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZoneSettingRequest() (request *DescribeZoneSettingRequest) {
    request = &DescribeZoneSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeZoneSetting")
    
    
    return
}

func NewDescribeZoneSettingResponse() (response *DescribeZoneSettingResponse) {
    response = &DescribeZoneSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeZoneSetting
// This API is used to query the site configuration.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeZoneSetting(request *DescribeZoneSettingRequest) (response *DescribeZoneSettingResponse, err error) {
    return c.DescribeZoneSettingWithContext(context.Background(), request)
}

// DescribeZoneSetting
// This API is used to query the site configuration.
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeZoneSettingWithContext(ctx context.Context, request *DescribeZoneSettingRequest) (response *DescribeZoneSettingResponse, err error) {
    if request == nil {
        request = NewDescribeZoneSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZoneSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZoneSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZonesRequest() (request *DescribeZonesRequest) {
    request = &DescribeZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DescribeZones")
    
    
    return
}

func NewDescribeZonesResponse() (response *DescribeZonesResponse) {
    response = &DescribeZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeZones
// This API is used to query the information of sites that you have access to. You can filter sites based on different query criteria.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    return c.DescribeZonesWithContext(context.Background(), request)
}

// DescribeZones
// This API is used to query the information of sites that you have access to. You can filter sites based on different query criteria.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeZonesWithContext(ctx context.Context, request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    if request == nil {
        request = NewDescribeZonesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZones require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZonesResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadL4LogsRequest() (request *DownloadL4LogsRequest) {
    request = &DownloadL4LogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DownloadL4Logs")
    
    
    return
}

func NewDownloadL4LogsResponse() (response *DownloadL4LogsResponse) {
    response = &DownloadL4LogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DownloadL4Logs
// This API is used to download L4 logs.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DownloadL4Logs(request *DownloadL4LogsRequest) (response *DownloadL4LogsResponse, err error) {
    return c.DownloadL4LogsWithContext(context.Background(), request)
}

// DownloadL4Logs
// This API is used to download L4 logs.
//
// error code that may be returned:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DownloadL4LogsWithContext(ctx context.Context, request *DownloadL4LogsRequest) (response *DownloadL4LogsResponse, err error) {
    if request == nil {
        request = NewDownloadL4LogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadL4Logs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDownloadL4LogsResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadL7LogsRequest() (request *DownloadL7LogsRequest) {
    request = &DownloadL7LogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "DownloadL7Logs")
    
    
    return
}

func NewDownloadL7LogsResponse() (response *DownloadL7LogsResponse) {
    response = &DownloadL7LogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DownloadL7Logs
// This API is used to download L7 logs.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DownloadL7Logs(request *DownloadL7LogsRequest) (response *DownloadL7LogsResponse, err error) {
    return c.DownloadL7LogsWithContext(context.Background(), request)
}

// DownloadL7Logs
// This API is used to download L7 logs.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DownloadL7LogsWithContext(ctx context.Context, request *DownloadL7LogsRequest) (response *DownloadL7LogsResponse, err error) {
    if request == nil {
        request = NewDownloadL7LogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadL7Logs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDownloadL7LogsResponse()
    err = c.Send(request, response)
    return
}

func NewIdentifyZoneRequest() (request *IdentifyZoneRequest) {
    request = &IdentifyZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "IdentifyZone")
    
    
    return
}

func NewIdentifyZoneResponse() (response *IdentifyZoneResponse) {
    response = &IdentifyZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// IdentifyZone
// This API is used to verify ownership of the site.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) IdentifyZone(request *IdentifyZoneRequest) (response *IdentifyZoneResponse, err error) {
    return c.IdentifyZoneWithContext(context.Background(), request)
}

// IdentifyZone
// This API is used to verify ownership of the site.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) IdentifyZoneWithContext(ctx context.Context, request *IdentifyZoneRequest) (response *IdentifyZoneResponse, err error) {
    if request == nil {
        request = NewIdentifyZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("IdentifyZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewIdentifyZoneResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccelerationDomainRequest() (request *ModifyAccelerationDomainRequest) {
    request = &ModifyAccelerationDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyAccelerationDomain")
    
    
    return
}

func NewModifyAccelerationDomainResponse() (response *ModifyAccelerationDomainResponse) {
    response = &ModifyAccelerationDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAccelerationDomain
// This API is used to modify an accelerated domain name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_CONFLICTHOSTORIGIN = "InvalidParameter.ConflictHostOrigin"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDAWSSECRETKEY = "InvalidParameter.InvalidAwsSecretKey"
//  INVALIDPARAMETER_INVALIDCLIENTIPORIGIN = "InvalidParameter.InvalidClientIpOrigin"
//  INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_INVALIDPRIVATEACCESSPARAMS = "InvalidParameter.InvalidPrivateAccessParams"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGCLIENTIP = "InvalidParameter.InvalidStandardDebugClientIp"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_ORIGINISINNERIP = "InvalidParameter.OriginIsInnerIp"
//  INVALIDPARAMETER_SPACENOTBINDORIGIN = "InvalidParameter.SpaceNotBindOrigin"
//  INVALIDPARAMETER_ZONEISGRAYPUBLISHING = "InvalidParameter.ZoneIsGrayPublishing"
//  INVALIDPARAMETERVALUE_CONFLICTRECORD = "InvalidParameterValue.ConflictRecord"
//  INVALIDPARAMETERVALUE_DOMAINNOTMATCHZONE = "InvalidParameterValue.DomainNotMatchZone"
//  INVALIDPARAMETERVALUE_INVALIDDOMAINSTATUS = "InvalidParameterValue.InvalidDomainStatus"
//  OPERATIONDENIED_DOMAINNOICP = "OperationDenied.DomainNoICP"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_DNSRECORD = "ResourceInUse.DnsRecord"
//  RESOURCEUNAVAILABLE_DOMAINNOTFOUND = "ResourceUnavailable.DomainNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyAccelerationDomain(request *ModifyAccelerationDomainRequest) (response *ModifyAccelerationDomainResponse, err error) {
    return c.ModifyAccelerationDomainWithContext(context.Background(), request)
}

// ModifyAccelerationDomain
// This API is used to modify an accelerated domain name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_CONFLICTHOSTORIGIN = "InvalidParameter.ConflictHostOrigin"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDAWSSECRETKEY = "InvalidParameter.InvalidAwsSecretKey"
//  INVALIDPARAMETER_INVALIDCLIENTIPORIGIN = "InvalidParameter.InvalidClientIpOrigin"
//  INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_INVALIDPRIVATEACCESSPARAMS = "InvalidParameter.InvalidPrivateAccessParams"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGCLIENTIP = "InvalidParameter.InvalidStandardDebugClientIp"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_ORIGINISINNERIP = "InvalidParameter.OriginIsInnerIp"
//  INVALIDPARAMETER_SPACENOTBINDORIGIN = "InvalidParameter.SpaceNotBindOrigin"
//  INVALIDPARAMETER_ZONEISGRAYPUBLISHING = "InvalidParameter.ZoneIsGrayPublishing"
//  INVALIDPARAMETERVALUE_CONFLICTRECORD = "InvalidParameterValue.ConflictRecord"
//  INVALIDPARAMETERVALUE_DOMAINNOTMATCHZONE = "InvalidParameterValue.DomainNotMatchZone"
//  INVALIDPARAMETERVALUE_INVALIDDOMAINSTATUS = "InvalidParameterValue.InvalidDomainStatus"
//  OPERATIONDENIED_DOMAINNOICP = "OperationDenied.DomainNoICP"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_DNSRECORD = "ResourceInUse.DnsRecord"
//  RESOURCEUNAVAILABLE_DOMAINNOTFOUND = "ResourceUnavailable.DomainNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyAccelerationDomainWithContext(ctx context.Context, request *ModifyAccelerationDomainRequest) (response *ModifyAccelerationDomainResponse, err error) {
    if request == nil {
        request = NewModifyAccelerationDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccelerationDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccelerationDomainResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccelerationDomainStatusesRequest() (request *ModifyAccelerationDomainStatusesRequest) {
    request = &ModifyAccelerationDomainStatusesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyAccelerationDomainStatuses")
    
    
    return
}

func NewModifyAccelerationDomainStatusesResponse() (response *ModifyAccelerationDomainStatusesResponse) {
    response = &ModifyAccelerationDomainStatusesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAccelerationDomainStatuses
// This API is used to batch modify the status of accelerated domains.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_INVALIDERRORPAGE = "InvalidParameter.InvalidErrorPage"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_ZONEISGRAYPUBLISHING = "InvalidParameter.ZoneIsGrayPublishing"
//  INVALIDPARAMETERVALUE_DOMAINNOTMATCHZONE = "InvalidParameterValue.DomainNotMatchZone"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_DOMAINNOTFOUND = "ResourceUnavailable.DomainNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifyAccelerationDomainStatuses(request *ModifyAccelerationDomainStatusesRequest) (response *ModifyAccelerationDomainStatusesResponse, err error) {
    return c.ModifyAccelerationDomainStatusesWithContext(context.Background(), request)
}

// ModifyAccelerationDomainStatuses
// This API is used to batch modify the status of accelerated domains.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_INVALIDERRORPAGE = "InvalidParameter.InvalidErrorPage"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_ZONEISGRAYPUBLISHING = "InvalidParameter.ZoneIsGrayPublishing"
//  INVALIDPARAMETERVALUE_DOMAINNOTMATCHZONE = "InvalidParameterValue.DomainNotMatchZone"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"
//  OPERATIONDENIED_VERSIONCONTROLLOCKED = "OperationDenied.VersionControlLocked"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_DOMAINNOTFOUND = "ResourceUnavailable.DomainNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifyAccelerationDomainStatusesWithContext(ctx context.Context, request *ModifyAccelerationDomainStatusesRequest) (response *ModifyAccelerationDomainStatusesResponse, err error) {
    if request == nil {
        request = NewModifyAccelerationDomainStatusesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccelerationDomainStatuses require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccelerationDomainStatusesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAliasDomainRequest() (request *ModifyAliasDomainRequest) {
    request = &ModifyAliasDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyAliasDomain")
    
    
    return
}

func NewModifyAliasDomainResponse() (response *ModifyAliasDomainResponse) {
    response = &ModifyAliasDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAliasDomain
// This API is used to modify an alias domain name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CERTNOTMATCHDOMAIN = "InvalidParameter.CertNotMatchDomain"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyAliasDomain(request *ModifyAliasDomainRequest) (response *ModifyAliasDomainResponse, err error) {
    return c.ModifyAliasDomainWithContext(context.Background(), request)
}

// ModifyAliasDomain
// This API is used to modify an alias domain name.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CERTNOTMATCHDOMAIN = "InvalidParameter.CertNotMatchDomain"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyAliasDomainWithContext(ctx context.Context, request *ModifyAliasDomainRequest) (response *ModifyAliasDomainResponse, err error) {
    if request == nil {
        request = NewModifyAliasDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAliasDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAliasDomainResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAliasDomainStatusRequest() (request *ModifyAliasDomainStatusRequest) {
    request = &ModifyAliasDomainStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyAliasDomainStatus")
    
    
    return
}

func NewModifyAliasDomainStatusResponse() (response *ModifyAliasDomainStatusResponse) {
    response = &ModifyAliasDomainStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAliasDomainStatus
// This API is used to modify the status of an alias domain name.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) ModifyAliasDomainStatus(request *ModifyAliasDomainStatusRequest) (response *ModifyAliasDomainStatusResponse, err error) {
    return c.ModifyAliasDomainStatusWithContext(context.Background(), request)
}

// ModifyAliasDomainStatus
// This API is used to modify the status of an alias domain name.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) ModifyAliasDomainStatusWithContext(ctx context.Context, request *ModifyAliasDomainStatusRequest) (response *ModifyAliasDomainStatusResponse, err error) {
    if request == nil {
        request = NewModifyAliasDomainStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAliasDomainStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAliasDomainStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationProxyRequest() (request *ModifyApplicationProxyRequest) {
    request = &ModifyApplicationProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyApplicationProxy")
    
    
    return
}

func NewModifyApplicationProxyResponse() (response *ModifyApplicationProxyResponse) {
    response = &ModifyApplicationProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApplicationProxy
// This API is on an earlier version. If you want to call it, please switch to the latest version. For details, see [ModifyL4Proxy
//
// ] (https://intl.cloud.tencent.com/document/product/1552/103411?from_cn_redirect=1).
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4LACKOFRESOURCES = "OperationDenied.L4LackOfResources"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyApplicationProxy(request *ModifyApplicationProxyRequest) (response *ModifyApplicationProxyResponse, err error) {
    return c.ModifyApplicationProxyWithContext(context.Background(), request)
}

// ModifyApplicationProxy
// This API is on an earlier version. If you want to call it, please switch to the latest version. For details, see [ModifyL4Proxy
//
// ] (https://intl.cloud.tencent.com/document/product/1552/103411?from_cn_redirect=1).
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4LACKOFRESOURCES = "OperationDenied.L4LackOfResources"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyApplicationProxyWithContext(ctx context.Context, request *ModifyApplicationProxyRequest) (response *ModifyApplicationProxyResponse, err error) {
    if request == nil {
        request = NewModifyApplicationProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationProxyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationProxyRuleRequest() (request *ModifyApplicationProxyRuleRequest) {
    request = &ModifyApplicationProxyRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyApplicationProxyRule")
    
    
    return
}

func NewModifyApplicationProxyRuleResponse() (response *ModifyApplicationProxyRuleResponse) {
    response = &ModifyApplicationProxyRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApplicationProxyRule
// This API is on an earlier version. If you want to call it, please switch to the latest version. For details, see [ModifyL4ProxyRules] (https://intl.cloud.tencent.com/document/product/1552/103410?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyApplicationProxyRule(request *ModifyApplicationProxyRuleRequest) (response *ModifyApplicationProxyRuleResponse, err error) {
    return c.ModifyApplicationProxyRuleWithContext(context.Background(), request)
}

// ModifyApplicationProxyRule
// This API is on an earlier version. If you want to call it, please switch to the latest version. For details, see [ModifyL4ProxyRules] (https://intl.cloud.tencent.com/document/product/1552/103410?from_cn_redirect=1).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyApplicationProxyRuleWithContext(ctx context.Context, request *ModifyApplicationProxyRuleRequest) (response *ModifyApplicationProxyRuleResponse, err error) {
    if request == nil {
        request = NewModifyApplicationProxyRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationProxyRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationProxyRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationProxyRuleStatusRequest() (request *ModifyApplicationProxyRuleStatusRequest) {
    request = &ModifyApplicationProxyRuleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyApplicationProxyRuleStatus")
    
    
    return
}

func NewModifyApplicationProxyRuleStatusResponse() (response *ModifyApplicationProxyRuleStatusResponse) {
    response = &ModifyApplicationProxyRuleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApplicationProxyRuleStatus
// This API is on an earlier version. If you want to call it, please switch to the latest version. For details, see [ModifyL4ProxyRulesStatus
//
// ] (https://intl.cloud.tencent.com/document/product/1552/103409?from_cn_redirect=1).
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyApplicationProxyRuleStatus(request *ModifyApplicationProxyRuleStatusRequest) (response *ModifyApplicationProxyRuleStatusResponse, err error) {
    return c.ModifyApplicationProxyRuleStatusWithContext(context.Background(), request)
}

// ModifyApplicationProxyRuleStatus
// This API is on an earlier version. If you want to call it, please switch to the latest version. For details, see [ModifyL4ProxyRulesStatus
//
// ] (https://intl.cloud.tencent.com/document/product/1552/103409?from_cn_redirect=1).
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyApplicationProxyRuleStatusWithContext(ctx context.Context, request *ModifyApplicationProxyRuleStatusRequest) (response *ModifyApplicationProxyRuleStatusResponse, err error) {
    if request == nil {
        request = NewModifyApplicationProxyRuleStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationProxyRuleStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationProxyRuleStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationProxyStatusRequest() (request *ModifyApplicationProxyStatusRequest) {
    request = &ModifyApplicationProxyStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyApplicationProxyStatus")
    
    
    return
}

func NewModifyApplicationProxyStatusResponse() (response *ModifyApplicationProxyStatusResponse) {
    response = &ModifyApplicationProxyStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApplicationProxyStatus
// This API is on an earlier version. If you want to call it, please switch to the latest version. For details, see [ModifyL4ProxyStatus] (https://intl.cloud.tencent.com/document/product/1552/103408?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyApplicationProxyStatus(request *ModifyApplicationProxyStatusRequest) (response *ModifyApplicationProxyStatusResponse, err error) {
    return c.ModifyApplicationProxyStatusWithContext(context.Background(), request)
}

// ModifyApplicationProxyStatus
// This API is on an earlier version. If you want to call it, please switch to the latest version. For details, see [ModifyL4ProxyStatus] (https://intl.cloud.tencent.com/document/product/1552/103408?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyApplicationProxyStatusWithContext(ctx context.Context, request *ModifyApplicationProxyStatusRequest) (response *ModifyApplicationProxyStatusResponse, err error) {
    if request == nil {
        request = NewModifyApplicationProxyStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationProxyStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationProxyStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyHostsCertificateRequest() (request *ModifyHostsCertificateRequest) {
    request = &ModifyHostsCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyHostsCertificate")
    
    
    return
}

func NewModifyHostsCertificateResponse() (response *ModifyHostsCertificateResponse) {
    response = &ModifyHostsCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyHostsCertificate
// This API is used to configure the certificate of a site. You can use your own certificate or [apply for a free certificate](https://intl.cloud.tencent.com/document/product/1552/90437?from_cn_redirect=1).
//
// To use an external certificate, upload the certificate to [SSL Certificates Console](https://console.cloud.tencent.com/certoview) first, and then input the certificate ID in this API. For details, see [Deploying Own Certificates to EdgeOne Domains](https://intl.cloud.tencent.com/document/product/1552/88874?from_cn_redirect=1).
//
//  
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATEHASEXPIRED = "FailedOperation.CertificateHasExpired"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDZONESTATUS = "FailedOperation.InvalidZoneStatus"
//  FAILEDOPERATION_MODIFYFAILED = "FailedOperation.ModifyFailed"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWERROR = "InternalError.UnknowError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_CERTNOTMATCHDOMAIN = "InvalidParameter.CertNotMatchDomain"
//  INVALIDPARAMETER_CERTTOEXPIRE = "InvalidParameter.CertToExpire"
//  INVALIDPARAMETER_CERTTOOSHORTKEYSIZE = "InvalidParameter.CertTooShortKeySize"
//  INVALIDPARAMETER_CNAMEWILDHOSTNOTALLOWAPPLYCERTIFICATE = "InvalidParameter.CnameWildHostNotAllowApplyCertificate"
//  INVALIDPARAMETER_HOSTSTATUSNOTALLOWAPPLYCERTIFICATE = "InvalidParameter.HostStatusNotAllowApplyCertificate"
//  INVALIDPARAMETER_INVALIDCERTINFO = "InvalidParameter.InvalidCertInfo"
//  INVALIDPARAMETER_INVALIDHTTPSCERTINFO = "InvalidParameter.InvalidHttpsCertInfo"
//  INVALIDPARAMETER_INVALIDHTTPSTLSVERSION = "InvalidParameter.InvalidHttpsTlsVersion"
//  INVALIDPARAMETER_ZONEISGRAYPUBLISHING = "InvalidParameter.ZoneIsGrayPublishing"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_CERTNOTFOUND = "ResourceUnavailable.CertNotFound"
//  RESOURCEUNAVAILABLE_HOSTNOTFOUND = "ResourceUnavailable.HostNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) ModifyHostsCertificate(request *ModifyHostsCertificateRequest) (response *ModifyHostsCertificateResponse, err error) {
    return c.ModifyHostsCertificateWithContext(context.Background(), request)
}

// ModifyHostsCertificate
// This API is used to configure the certificate of a site. You can use your own certificate or [apply for a free certificate](https://intl.cloud.tencent.com/document/product/1552/90437?from_cn_redirect=1).
//
// To use an external certificate, upload the certificate to [SSL Certificates Console](https://console.cloud.tencent.com/certoview) first, and then input the certificate ID in this API. For details, see [Deploying Own Certificates to EdgeOne Domains](https://intl.cloud.tencent.com/document/product/1552/88874?from_cn_redirect=1).
//
//  
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATEHASEXPIRED = "FailedOperation.CertificateHasExpired"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  FAILEDOPERATION_INVALIDZONESTATUS = "FailedOperation.InvalidZoneStatus"
//  FAILEDOPERATION_MODIFYFAILED = "FailedOperation.ModifyFailed"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWERROR = "InternalError.UnknowError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_CERTNOTMATCHDOMAIN = "InvalidParameter.CertNotMatchDomain"
//  INVALIDPARAMETER_CERTTOEXPIRE = "InvalidParameter.CertToExpire"
//  INVALIDPARAMETER_CERTTOOSHORTKEYSIZE = "InvalidParameter.CertTooShortKeySize"
//  INVALIDPARAMETER_CNAMEWILDHOSTNOTALLOWAPPLYCERTIFICATE = "InvalidParameter.CnameWildHostNotAllowApplyCertificate"
//  INVALIDPARAMETER_HOSTSTATUSNOTALLOWAPPLYCERTIFICATE = "InvalidParameter.HostStatusNotAllowApplyCertificate"
//  INVALIDPARAMETER_INVALIDCERTINFO = "InvalidParameter.InvalidCertInfo"
//  INVALIDPARAMETER_INVALIDHTTPSCERTINFO = "InvalidParameter.InvalidHttpsCertInfo"
//  INVALIDPARAMETER_INVALIDHTTPSTLSVERSION = "InvalidParameter.InvalidHttpsTlsVersion"
//  INVALIDPARAMETER_ZONEISGRAYPUBLISHING = "InvalidParameter.ZoneIsGrayPublishing"
//  LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_CERTNOTFOUND = "ResourceUnavailable.CertNotFound"
//  RESOURCEUNAVAILABLE_HOSTNOTFOUND = "ResourceUnavailable.HostNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) ModifyHostsCertificateWithContext(ctx context.Context, request *ModifyHostsCertificateRequest) (response *ModifyHostsCertificateResponse, err error) {
    if request == nil {
        request = NewModifyHostsCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyHostsCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyHostsCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyL4ProxyRequest() (request *ModifyL4ProxyRequest) {
    request = &ModifyL4ProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyL4Proxy")
    
    
    return
}

func NewModifyL4ProxyResponse() (response *ModifyL4ProxyResponse) {
    response = &ModifyL4ProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyL4Proxy
// This API is used to modify the configuration of a Layer 4 proxy instance.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDDISABLE = "OperationDenied.AccelerateMainlandDisable"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4LACKOFRESOURCES = "OperationDenied.L4LackOfResources"
//  OPERATIONDENIED_L4PROXYINOFFLINESTATUS = "OperationDenied.L4ProxyInOfflineStatus"
//  OPERATIONDENIED_L4PROXYINPROCESSSTATUS = "OperationDenied.L4ProxyInProcessStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) ModifyL4Proxy(request *ModifyL4ProxyRequest) (response *ModifyL4ProxyResponse, err error) {
    return c.ModifyL4ProxyWithContext(context.Background(), request)
}

// ModifyL4Proxy
// This API is used to modify the configuration of a Layer 4 proxy instance.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDDISABLE = "OperationDenied.AccelerateMainlandDisable"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4LACKOFRESOURCES = "OperationDenied.L4LackOfResources"
//  OPERATIONDENIED_L4PROXYINOFFLINESTATUS = "OperationDenied.L4ProxyInOfflineStatus"
//  OPERATIONDENIED_L4PROXYINPROCESSSTATUS = "OperationDenied.L4ProxyInProcessStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) ModifyL4ProxyWithContext(ctx context.Context, request *ModifyL4ProxyRequest) (response *ModifyL4ProxyResponse, err error) {
    if request == nil {
        request = NewModifyL4ProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyL4Proxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyL4ProxyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyL4ProxyRulesRequest() (request *ModifyL4ProxyRulesRequest) {
    request = &ModifyL4ProxyRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyL4ProxyRules")
    
    
    return
}

func NewModifyL4ProxyRulesResponse() (response *ModifyL4ProxyRulesResponse) {
    response = &ModifyL4ProxyRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyL4ProxyRules
// This API is used to modify Layer 4 proxy forwarding rules, supporting both individual and batch modification.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyL4ProxyRules(request *ModifyL4ProxyRulesRequest) (response *ModifyL4ProxyRulesResponse, err error) {
    return c.ModifyL4ProxyRulesWithContext(context.Background(), request)
}

// ModifyL4ProxyRules
// This API is used to modify Layer 4 proxy forwarding rules, supporting both individual and batch modification.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyL4ProxyRulesWithContext(ctx context.Context, request *ModifyL4ProxyRulesRequest) (response *ModifyL4ProxyRulesResponse, err error) {
    if request == nil {
        request = NewModifyL4ProxyRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyL4ProxyRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyL4ProxyRulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyL4ProxyRulesStatusRequest() (request *ModifyL4ProxyRulesStatusRequest) {
    request = &ModifyL4ProxyRulesStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyL4ProxyRulesStatus")
    
    
    return
}

func NewModifyL4ProxyRulesStatusResponse() (response *ModifyL4ProxyRulesStatusResponse) {
    response = &ModifyL4ProxyRulesStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyL4ProxyRulesStatus
// This API is used to start or stop Layer 4 proxy forwarding rules, supporting both individual and batch operation.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyL4ProxyRulesStatus(request *ModifyL4ProxyRulesStatusRequest) (response *ModifyL4ProxyRulesStatusResponse, err error) {
    return c.ModifyL4ProxyRulesStatusWithContext(context.Background(), request)
}

// ModifyL4ProxyRulesStatus
// This API is used to start or stop Layer 4 proxy forwarding rules, supporting both individual and batch operation.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyL4ProxyRulesStatusWithContext(ctx context.Context, request *ModifyL4ProxyRulesStatusRequest) (response *ModifyL4ProxyRulesStatusResponse, err error) {
    if request == nil {
        request = NewModifyL4ProxyRulesStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyL4ProxyRulesStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyL4ProxyRulesStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyL4ProxyStatusRequest() (request *ModifyL4ProxyStatusRequest) {
    request = &ModifyL4ProxyStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyL4ProxyStatus")
    
    
    return
}

func NewModifyL4ProxyStatusResponse() (response *ModifyL4ProxyStatusResponse) {
    response = &ModifyL4ProxyStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyL4ProxyStatus
// This API is used to enable or disable a Layer 4 proxy instance.
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyL4ProxyStatus(request *ModifyL4ProxyStatusRequest) (response *ModifyL4ProxyStatusResponse, err error) {
    return c.ModifyL4ProxyStatusWithContext(context.Background(), request)
}

// ModifyL4ProxyStatus
// This API is used to enable or disable a Layer 4 proxy instance.
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_L4PROXYINBANNEDSTATUS = "OperationDenied.L4ProxyInBannedStatus"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyL4ProxyStatusWithContext(ctx context.Context, request *ModifyL4ProxyStatusRequest) (response *ModifyL4ProxyStatusResponse, err error) {
    if request == nil {
        request = NewModifyL4ProxyStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyL4ProxyStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyL4ProxyStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyOriginGroupRequest() (request *ModifyOriginGroupRequest) {
    request = &ModifyOriginGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyOriginGroup")
    
    
    return
}

func NewModifyOriginGroupResponse() (response *ModifyOriginGroupResponse) {
    response = &ModifyOriginGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyOriginGroup
// This API is used to modify the configuration of an origin group. The original configuration will be overwritten. 
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_HOSTHEADERINVALID = "InvalidParameter.HostHeaderInvalid"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_ORIGINL4RECORDIPV4MIXDOMAIN = "InvalidParameter.OriginL4RecordIPV4MixDomain"
//  INVALIDPARAMETER_ORIGINL4RECORDMULTIDOMAIN = "InvalidParameter.OriginL4RecordMultiDomain"
//  INVALIDPARAMETER_ORIGINNAMEEXISTS = "InvalidParameter.OriginNameExists"
//  INVALIDPARAMETER_ORIGINRECORDFORMATERROR = "InvalidParameter.OriginRecordFormatError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATIONDOMAINSTATUSNOTINONLINE = "OperationDenied.AccelerationDomainStatusNotInOnline"
//  OPERATIONDENIED_L4STATUSNOTINONLINE = "OperationDenied.L4StatusNotInOnline"
//  OPERATIONDENIED_LOADBALANCESTATUSNOTINONLINE = "OperationDenied.LoadBalanceStatusNotInOnline"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifyOriginGroup(request *ModifyOriginGroupRequest) (response *ModifyOriginGroupResponse, err error) {
    return c.ModifyOriginGroupWithContext(context.Background(), request)
}

// ModifyOriginGroup
// This API is used to modify the configuration of an origin group. The original configuration will be overwritten. 
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_HOSTHEADERINVALID = "InvalidParameter.HostHeaderInvalid"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_ORIGINL4RECORDIPV4MIXDOMAIN = "InvalidParameter.OriginL4RecordIPV4MixDomain"
//  INVALIDPARAMETER_ORIGINL4RECORDMULTIDOMAIN = "InvalidParameter.OriginL4RecordMultiDomain"
//  INVALIDPARAMETER_ORIGINNAMEEXISTS = "InvalidParameter.OriginNameExists"
//  INVALIDPARAMETER_ORIGINRECORDFORMATERROR = "InvalidParameter.OriginRecordFormatError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATIONDOMAINSTATUSNOTINONLINE = "OperationDenied.AccelerationDomainStatusNotInOnline"
//  OPERATIONDENIED_L4STATUSNOTINONLINE = "OperationDenied.L4StatusNotInOnline"
//  OPERATIONDENIED_LOADBALANCESTATUSNOTINONLINE = "OperationDenied.LoadBalanceStatusNotInOnline"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifyOriginGroupWithContext(ctx context.Context, request *ModifyOriginGroupRequest) (response *ModifyOriginGroupResponse, err error) {
    if request == nil {
        request = NewModifyOriginGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyOriginGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyOriginGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRealtimeLogDeliveryTaskRequest() (request *ModifyRealtimeLogDeliveryTaskRequest) {
    request = &ModifyRealtimeLogDeliveryTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyRealtimeLogDeliveryTask")
    
    
    return
}

func NewModifyRealtimeLogDeliveryTaskResponse() (response *ModifyRealtimeLogDeliveryTaskResponse) {
    response = &ModifyRealtimeLogDeliveryTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRealtimeLogDeliveryTask
// This API is used to modify the real-time log delivery task configuration. This API has the following restrictions:<li>Does not support modifying the destination type of the real-time log delivery task (TaskType);</li><li>Does not support modifying the data delivery type (LogType)</li><li>Does not support modifying the data delivery area (Area)</li><li>Does not support modifying the detailed destination configuration, such as log set and log topic, when the destination of the original real-time log delivery task is Tencent Cloud CLS.</li>
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATELOGTOPICTASKAUTHFAILURE = "FailedOperation.CreateLogTopicTaskAuthFailure"
func (c *Client) ModifyRealtimeLogDeliveryTask(request *ModifyRealtimeLogDeliveryTaskRequest) (response *ModifyRealtimeLogDeliveryTaskResponse, err error) {
    return c.ModifyRealtimeLogDeliveryTaskWithContext(context.Background(), request)
}

// ModifyRealtimeLogDeliveryTask
// This API is used to modify the real-time log delivery task configuration. This API has the following restrictions:<li>Does not support modifying the destination type of the real-time log delivery task (TaskType);</li><li>Does not support modifying the data delivery type (LogType)</li><li>Does not support modifying the data delivery area (Area)</li><li>Does not support modifying the detailed destination configuration, such as log set and log topic, when the destination of the original real-time log delivery task is Tencent Cloud CLS.</li>
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CREATELOGTOPICTASKAUTHFAILURE = "FailedOperation.CreateLogTopicTaskAuthFailure"
func (c *Client) ModifyRealtimeLogDeliveryTaskWithContext(ctx context.Context, request *ModifyRealtimeLogDeliveryTaskRequest) (response *ModifyRealtimeLogDeliveryTaskResponse, err error) {
    if request == nil {
        request = NewModifyRealtimeLogDeliveryTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRealtimeLogDeliveryTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRealtimeLogDeliveryTaskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRuleRequest() (request *ModifyRuleRequest) {
    request = &ModifyRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyRule")
    
    
    return
}

func NewModifyRuleResponse() (response *ModifyRuleResponse) {
    response = &ModifyRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRule
// This API is used to modify a rule in the rule engine.
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGTOOMANYVALUE = "InvalidParameter.CacheKeyQueryStringTooManyValue"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_COMPRESSIONINVALIDALGORITHMS = "InvalidParameter.CompressionInvalidAlgorithms"
//  INVALIDPARAMETER_ERRACTIONUNSUPPORTTARGET = "InvalidParameter.ErrActionUnsupportTarget"
//  INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONDUPLICATEACTION = "InvalidParameter.ErrInvalidActionDuplicateAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMACTION = "InvalidParameter.ErrInvalidActionParamAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMDUPLICATENAME = "InvalidParameter.ErrInvalidActionParamDuplicateName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMNAME = "InvalidParameter.ErrInvalidActionParamName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMTOOMANYVALUES = "InvalidParameter.ErrInvalidActionParamTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_ERRINVALIDACTIONTYPE = "InvalidParameter.ErrInvalidActionType"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONHOSTTOOMANYWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidConditionHostTooManyWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONIGNORECASE = "InvalidParameter.ErrInvalidConditionIgnoreCase"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMEBADNAME = "InvalidParameter.ErrInvalidConditionNameBadName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADREGULAR = "InvalidParameter.ErrInvalidConditionValueBadRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADURL = "InvalidParameter.ErrInvalidConditionValueBadUrl"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUECONTAINFILENAMEEXTENSION = "InvalidParameter.ErrInvalidConditionValueBadValueContainFileNameExtension"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOLONGVALUE = "InvalidParameter.ErrInvalidConditionValueTooLongValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYVALUES = "InvalidParameter.ErrInvalidConditionValueTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYWILDCARD = "InvalidParameter.ErrInvalidConditionValueTooManyWildcard"
//  INVALIDPARAMETER_ERRINVALIDELSEWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidElseWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRNILCONDITION = "InvalidParameter.ErrNilCondition"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESECRETKEY = "InvalidParameter.InvalidAuthenticationTypeSecretKey"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPETIMEFORMAT = "InvalidParameter.InvalidAuthenticationTypeTimeFormat"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPETIMEPARAM = "InvalidParameter.InvalidAuthenticationTypeTimeParam"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"
//  INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"
//  INVALIDPARAMETER_INVALIDCACHEKEYCOOKIE = "InvalidParameter.InvalidCacheKeyCookie"
//  INVALIDPARAMETER_INVALIDCACHEKEYIGNORECASE = "InvalidParameter.InvalidCacheKeyIgnoreCase"
//  INVALIDPARAMETER_INVALIDCACHEKEYSCHEME = "InvalidParameter.InvalidCacheKeyScheme"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDCOSDOMAIN = "InvalidParameter.InvalidCosDomain"
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDUPSTREAMREQUESTQUERYSTRINGVALUE = "InvalidParameter.InvalidUpstreamRequestQueryStringValue"
//  INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_KEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.KeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_LOADBALANCEINSTANCEIDISREQUIRED = "InvalidParameter.LoadBalanceInstanceIdIsRequired"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ORIGINORIGINGROUPIDISREQUIRED = "InvalidParameter.OriginOriginGroupIdIsRequired"
//  INVALIDPARAMETER_ORIGINPULLPROTOCOLISREQUIRED = "InvalidParameter.OriginPullProtocolIsRequired"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_RESPONSEHEADERCACHECONTROLNOTALLOWDELETE = "InvalidParameter.ResponseHeaderCacheControlNotAllowDelete"
//  INVALIDPARAMETER_STATUSCODECACHEINVALIDSTATUSCODE = "InvalidParameter.StatusCodeCacheInvalidStatusCode"
//  INVALIDPARAMETER_TLSVERSIONNOTINSEQUENCE = "InvalidParameter.TlsVersionNotInSequence"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSESECURITYTYPE = "OperationDenied.InvalidAdvancedDefenseSecurityType"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifyRule(request *ModifyRuleRequest) (response *ModifyRuleResponse, err error) {
    return c.ModifyRuleWithContext(context.Background(), request)
}

// ModifyRule
// This API is used to modify a rule in the rule engine.
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGTOOMANYVALUE = "InvalidParameter.CacheKeyQueryStringTooManyValue"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_COMPRESSIONINVALIDALGORITHMS = "InvalidParameter.CompressionInvalidAlgorithms"
//  INVALIDPARAMETER_ERRACTIONUNSUPPORTTARGET = "InvalidParameter.ErrActionUnsupportTarget"
//  INVALIDPARAMETER_ERRINVALIDACTION = "InvalidParameter.ErrInvalidAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONDUPLICATEACTION = "InvalidParameter.ErrInvalidActionDuplicateAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAM = "InvalidParameter.ErrInvalidActionParam"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMACTION = "InvalidParameter.ErrInvalidActionParamAction"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMDUPLICATENAME = "InvalidParameter.ErrInvalidActionParamDuplicateName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMNAME = "InvalidParameter.ErrInvalidActionParamName"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMTOOMANYVALUES = "InvalidParameter.ErrInvalidActionParamTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDACTIONPARAMVALUE = "InvalidParameter.ErrInvalidActionParamValue"
//  INVALIDPARAMETER_ERRINVALIDACTIONTYPE = "InvalidParameter.ErrInvalidActionType"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONHOSTTOOMANYWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidConditionHostTooManyWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONIGNORECASE = "InvalidParameter.ErrInvalidConditionIgnoreCase"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMEBADNAME = "InvalidParameter.ErrInvalidConditionNameBadName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONNAMETARGETNOTSUPPORTNAME = "InvalidParameter.ErrInvalidConditionNameTargetNotSupportName"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADREGULAR = "InvalidParameter.ErrInvalidConditionValueBadRegular"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADURL = "InvalidParameter.ErrInvalidConditionValueBadUrl"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUE = "InvalidParameter.ErrInvalidConditionValueBadValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUEBADVALUECONTAINFILENAMEEXTENSION = "InvalidParameter.ErrInvalidConditionValueBadValueContainFileNameExtension"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOLONGVALUE = "InvalidParameter.ErrInvalidConditionValueTooLongValue"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYVALUES = "InvalidParameter.ErrInvalidConditionValueTooManyValues"
//  INVALIDPARAMETER_ERRINVALIDCONDITIONVALUETOOMANYWILDCARD = "InvalidParameter.ErrInvalidConditionValueTooManyWildcard"
//  INVALIDPARAMETER_ERRINVALIDELSEWHENMODIFYORIGINACTIONCONFIGURED = "InvalidParameter.ErrInvalidElseWhenModifyOriginActionConfigured"
//  INVALIDPARAMETER_ERRNILCONDITION = "InvalidParameter.ErrNilCondition"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_HOSTNOTFOUND = "InvalidParameter.HostNotFound"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESECRETKEY = "InvalidParameter.InvalidAuthenticationTypeSecretKey"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPESIGNPARAM = "InvalidParameter.InvalidAuthenticationTypeSignParam"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPETIMEFORMAT = "InvalidParameter.InvalidAuthenticationTypeTimeFormat"
//  INVALIDPARAMETER_INVALIDAUTHENTICATIONTYPETIMEPARAM = "InvalidParameter.InvalidAuthenticationTypeTimeParam"
//  INVALIDPARAMETER_INVALIDAWSREGION = "InvalidParameter.InvalidAwsRegion"
//  INVALIDPARAMETER_INVALIDBACKUPSERVERNAME = "InvalidParameter.InvalidBackupServerName"
//  INVALIDPARAMETER_INVALIDCACHEKEY = "InvalidParameter.InvalidCacheKey"
//  INVALIDPARAMETER_INVALIDCACHEKEYCOOKIE = "InvalidParameter.InvalidCacheKeyCookie"
//  INVALIDPARAMETER_INVALIDCACHEKEYIGNORECASE = "InvalidParameter.InvalidCacheKeyIgnoreCase"
//  INVALIDPARAMETER_INVALIDCACHEKEYSCHEME = "InvalidParameter.InvalidCacheKeyScheme"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDCOSDOMAIN = "InvalidParameter.InvalidCosDomain"
//  INVALIDPARAMETER_INVALIDERRORPAGEREDIRECTURL = "InvalidParameter.InvalidErrorPageRedirectUrl"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAMEXFF = "InvalidParameter.InvalidRequestHeaderNameXff"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERVALUE = "InvalidParameter.InvalidRequestHeaderValue"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERNAME = "InvalidParameter.InvalidResponseHeaderName"
//  INVALIDPARAMETER_INVALIDRESPONSEHEADERVALUE = "InvalidParameter.InvalidResponseHeaderValue"
//  INVALIDPARAMETER_INVALIDRULEENGINEACTION = "InvalidParameter.InvalidRuleEngineAction"
//  INVALIDPARAMETER_INVALIDRULEENGINENOTFOUND = "InvalidParameter.InvalidRuleEngineNotFound"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGET = "InvalidParameter.InvalidRuleEngineTarget"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSEXTENSION = "InvalidParameter.InvalidRuleEngineTargetsExtension"
//  INVALIDPARAMETER_INVALIDRULEENGINETARGETSURL = "InvalidParameter.InvalidRuleEngineTargetsUrl"
//  INVALIDPARAMETER_INVALIDSERVERNAME = "InvalidParameter.InvalidServerName"
//  INVALIDPARAMETER_INVALIDUPSTREAMREQUESTQUERYSTRINGVALUE = "InvalidParameter.InvalidUpstreamRequestQueryStringValue"
//  INVALIDPARAMETER_INVALIDURLREDIRECTHOST = "InvalidParameter.InvalidUrlRedirectHost"
//  INVALIDPARAMETER_INVALIDURLREDIRECTURL = "InvalidParameter.InvalidUrlRedirectUrl"
//  INVALIDPARAMETER_KEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.KeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_LOADBALANCEINSTANCEIDISREQUIRED = "InvalidParameter.LoadBalanceInstanceIdIsRequired"
//  INVALIDPARAMETER_NOTSUPPORTTHISPRESET = "InvalidParameter.NotSupportThisPreset"
//  INVALIDPARAMETER_ORIGINORIGINGROUPIDISREQUIRED = "InvalidParameter.OriginOriginGroupIdIsRequired"
//  INVALIDPARAMETER_ORIGINPULLPROTOCOLISREQUIRED = "InvalidParameter.OriginPullProtocolIsRequired"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_RESPONSEHEADERCACHECONTROLNOTALLOWDELETE = "InvalidParameter.ResponseHeaderCacheControlNotAllowDelete"
//  INVALIDPARAMETER_STATUSCODECACHEINVALIDSTATUSCODE = "InvalidParameter.StatusCodeCacheInvalidStatusCode"
//  INVALIDPARAMETER_TLSVERSIONNOTINSEQUENCE = "InvalidParameter.TlsVersionNotInSequence"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_INVALIDADVANCEDDEFENSESECURITYTYPE = "OperationDenied.InvalidAdvancedDefenseSecurityType"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifyRuleWithContext(ctx context.Context, request *ModifyRuleRequest) (response *ModifyRuleResponse, err error) {
    if request == nil {
        request = NewModifyRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecurityIPGroupRequest() (request *ModifySecurityIPGroupRequest) {
    request = &ModifySecurityIPGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifySecurityIPGroup")
    
    
    return
}

func NewModifySecurityIPGroupResponse() (response *ModifySecurityIPGroupResponse) {
    response = &ModifySecurityIPGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySecurityIPGroup
// This API is used to modify a security IP group.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifySecurityIPGroup(request *ModifySecurityIPGroupRequest) (response *ModifySecurityIPGroupResponse, err error) {
    return c.ModifySecurityIPGroupWithContext(context.Background(), request)
}

// ModifySecurityIPGroup
// This API is used to modify a security IP group.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) ModifySecurityIPGroupWithContext(ctx context.Context, request *ModifySecurityIPGroupRequest) (response *ModifySecurityIPGroupResponse, err error) {
    if request == nil {
        request = NewModifySecurityIPGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecurityIPGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySecurityIPGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecurityPolicyRequest() (request *ModifySecurityPolicyRequest) {
    request = &ModifySecurityPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifySecurityPolicy")
    
    
    return
}

func NewModifySecurityPolicyResponse() (response *ModifySecurityPolicyResponse) {
    response = &ModifySecurityPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySecurityPolicy
// This API is used to modify the web and bot security configurations.
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  LIMITEXCEEDED_SECURITY = "LimitExceeded.Security"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifySecurityPolicy(request *ModifySecurityPolicyRequest) (response *ModifySecurityPolicyResponse, err error) {
    return c.ModifySecurityPolicyWithContext(context.Background(), request)
}

// ModifySecurityPolicy
// This API is used to modify the web and bot security configurations.
//
// error code that may be returned:
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_SECURITY = "InvalidParameter.Security"
//  LIMITEXCEEDED_SECURITY = "LimitExceeded.Security"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifySecurityPolicyWithContext(ctx context.Context, request *ModifySecurityPolicyRequest) (response *ModifySecurityPolicyResponse, err error) {
    if request == nil {
        request = NewModifySecurityPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecurityPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySecurityPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyZoneRequest() (request *ModifyZoneRequest) {
    request = &ModifyZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyZone")
    
    
    return
}

func NewModifyZoneResponse() (response *ModifyZoneResponse) {
    response = &ModifyZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyZone
// This API is used to modify a site.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_LENGTHEXCEEDSLIMIT = "InvalidParameter.LengthExceedsLimit"
//  INVALIDPARAMETER_ZONENAMEISREQUIRED = "InvalidParameter.ZoneNameIsRequired"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ZONESAMEASNAME = "InvalidParameterValue.ZoneSameAsName"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_DOMAININSHARECNAMEGROUP = "OperationDenied.DomainInShareCnameGroup"
//  OPERATIONDENIED_DOMAINNUMBERISNOTZERO = "OperationDenied.DomainNumberIsNotZero"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_MULTIPLECNAMEZONE = "OperationDenied.MultipleCnameZone"
//  OPERATIONDENIED_NSNOTALLOWTRAFFICSTRATEGY = "OperationDenied.NSNotAllowTrafficStrategy"
//  OPERATIONDENIED_NODOMAINACCESSZONEONLYALLOWMODIFIEDTOCNAME = "OperationDenied.NoDomainAccessZoneOnlyAllowModifiedToCNAME"
//  OPERATIONDENIED_NODOMAINACCESSZONEONLYSUPPORTMODIFYTYPE = "OperationDenied.NoDomainAccessZoneOnlySupportModifyType"
//  OPERATIONDENIED_PLANNOTSUPPORTMODIFYZONEAREA = "OperationDenied.PlanNotSupportModifyZoneArea"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  RESOURCEINUSE_CNAME = "ResourceInUse.Cname"
//  RESOURCEINUSE_DNS = "ResourceInUse.Dns"
//  RESOURCEINUSE_GENERICHOST = "ResourceInUse.GenericHost"
//  RESOURCEINUSE_NS = "ResourceInUse.NS"
//  RESOURCEINUSE_OTHERS = "ResourceInUse.Others"
//  RESOURCEINUSE_OTHERSALIASDOMAIN = "ResourceInUse.OthersAliasDomain"
//  RESOURCEINUSE_OTHERSCNAME = "ResourceInUse.OthersCname"
//  RESOURCEINUSE_OTHERSNS = "ResourceInUse.OthersNS"
//  RESOURCEINUSE_SELFANDOTHERSCNAME = "ResourceInUse.SelfAndOthersCname"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyZone(request *ModifyZoneRequest) (response *ModifyZoneResponse, err error) {
    return c.ModifyZoneWithContext(context.Background(), request)
}

// ModifyZone
// This API is used to modify a site.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INVALIDPARAMETER_INVALIDORIGINIP = "InvalidParameter.InvalidOriginIp"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_LENGTHEXCEEDSLIMIT = "InvalidParameter.LengthExceedsLimit"
//  INVALIDPARAMETER_ZONENAMEISREQUIRED = "InvalidParameter.ZoneNameIsRequired"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ZONESAMEASNAME = "InvalidParameterValue.ZoneSameAsName"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_DOMAININSHARECNAMEGROUP = "OperationDenied.DomainInShareCnameGroup"
//  OPERATIONDENIED_DOMAINNUMBERISNOTZERO = "OperationDenied.DomainNumberIsNotZero"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_MULTIPLECNAMEZONE = "OperationDenied.MultipleCnameZone"
//  OPERATIONDENIED_NSNOTALLOWTRAFFICSTRATEGY = "OperationDenied.NSNotAllowTrafficStrategy"
//  OPERATIONDENIED_NODOMAINACCESSZONEONLYALLOWMODIFIEDTOCNAME = "OperationDenied.NoDomainAccessZoneOnlyAllowModifiedToCNAME"
//  OPERATIONDENIED_NODOMAINACCESSZONEONLYSUPPORTMODIFYTYPE = "OperationDenied.NoDomainAccessZoneOnlySupportModifyType"
//  OPERATIONDENIED_PLANNOTSUPPORTMODIFYZONEAREA = "OperationDenied.PlanNotSupportModifyZoneArea"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  RESOURCEINUSE_CNAME = "ResourceInUse.Cname"
//  RESOURCEINUSE_DNS = "ResourceInUse.Dns"
//  RESOURCEINUSE_GENERICHOST = "ResourceInUse.GenericHost"
//  RESOURCEINUSE_NS = "ResourceInUse.NS"
//  RESOURCEINUSE_OTHERS = "ResourceInUse.Others"
//  RESOURCEINUSE_OTHERSALIASDOMAIN = "ResourceInUse.OthersAliasDomain"
//  RESOURCEINUSE_OTHERSCNAME = "ResourceInUse.OthersCname"
//  RESOURCEINUSE_OTHERSNS = "ResourceInUse.OthersNS"
//  RESOURCEINUSE_SELFANDOTHERSCNAME = "ResourceInUse.SelfAndOthersCname"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyZoneWithContext(ctx context.Context, request *ModifyZoneRequest) (response *ModifyZoneResponse, err error) {
    if request == nil {
        request = NewModifyZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyZoneResponse()
    err = c.Send(request, response)
    return
}

func NewModifyZoneSettingRequest() (request *ModifyZoneSettingRequest) {
    request = &ModifyZoneSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyZoneSetting")
    
    
    return
}

func NewModifyZoneSettingResponse() (response *ModifyZoneSettingResponse) {
    response = &ModifyZoneSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyZoneSetting
// This API is used to modify the site configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWERROR = "InternalError.UnknowError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGREQUIRESFULLURLCACHEOFF = "InvalidParameter.CacheKeyQueryStringRequiresFullUrlCacheOff"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGTOOMANYVALUE = "InvalidParameter.CacheKeyQueryStringTooManyValue"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_CLIENTIPCOUNTRYCONFLICTSWITHIPV6 = "InvalidParameter.ClientIpCountryConflictsWithIpv6"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_INVALIDAWSPRIVATEACCESS = "InvalidParameter.InvalidAwsPrivateAccess"
//  INVALIDPARAMETER_INVALIDCACHECONFIGFOLLOWORIGIN = "InvalidParameter.InvalidCacheConfigFollowOrigin"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGVALUE = "InvalidParameter.InvalidCacheKeyQueryStringValue"
//  INVALIDPARAMETER_INVALIDCACHEONLYONSWITCH = "InvalidParameter.InvalidCacheOnlyOnSwitch"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPCOUNTRYHEADERNAME = "InvalidParameter.InvalidClientIpCountryHeaderName"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINEBILLING = "InvalidParameter.InvalidDynamicRoutineBilling"
//  INVALIDPARAMETER_INVALIDFORCEREDIRECTTYPE = "InvalidParameter.InvalidForceRedirectType"
//  INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"
//  INVALIDPARAMETER_INVALIDHTTPSCERTINFO = "InvalidParameter.InvalidHttpsCertInfo"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDHTTPSTLSVERSION = "InvalidParameter.InvalidHttpsTlsVersion"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDORIGINTYPE = "InvalidParameter.InvalidOriginType"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPOSTMAXSIZEBILLING = "InvalidParameter.InvalidPostMaxSizeBilling"
//  INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDRESOURCEIDBILLING = "InvalidParameter.InvalidResourceIdBilling"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGCLIENTIP = "InvalidParameter.InvalidStandardDebugClientIp"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEMPTYLIST = "InvalidParameter.InvalidStandardDebugEmptyList"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEXPIRETIMELIMIT = "InvalidParameter.InvalidStandardDebugExpireTimeLimit"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_MULTIPLYLAYERNOTSUPPORTSMARTROUTING = "InvalidParameter.MultiplyLayerNotSupportSmartRouting"
//  INVALIDPARAMETER_OCDIRECTORIGINREQUIRESSMARTROUTING = "InvalidParameter.OCDirectOriginRequiresSmartRouting"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"
//  INVALIDPARAMETER_ZONEISGRAYPUBLISHING = "InvalidParameter.ZoneIsGrayPublishing"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDDISABLE = "OperationDenied.AccelerateMainlandDisable"
//  OPERATIONDENIED_ACCELERATEMAINLANDIPV6CONFLICT = "OperationDenied.AccelerateMainlandIpv6Conflict"
//  OPERATIONDENIED_ACCELERATEMAINLANDMULTIPLYLAYERCONFLICT = "OperationDenied.AccelerateMainlandMultiplyLayerConflict"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_SHAREDCNAMEUNSUPPORTEDACCELERATEMAINLAND = "OperationDenied.SharedCNAMEUnsupportedAccelerateMainland"
//  OPERATIONDENIED_SHAREDCNAMEUNSUPPORTEDIPV6 = "OperationDenied.SharedCNAMEUnsupportedIPv6"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_POSTMAXSIZEQUOTANOTFOUND = "ResourceNotFound.PostMaxSizeQuotaNotFound"
//  RESOURCEUNAVAILABLE_CERTNOTFOUND = "ResourceUnavailable.CertNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyZoneSetting(request *ModifyZoneSettingRequest) (response *ModifyZoneSettingResponse, err error) {
    return c.ModifyZoneSettingWithContext(context.Background(), request)
}

// ModifyZoneSetting
// This API is used to modify the site configuration.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  INTERNALERROR_CONFIGLOCKED = "InternalError.ConfigLocked"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWERROR = "InternalError.UnknowError"
//  INVALIDPARAMETER_ACTIONINPROGRESS = "InvalidParameter.ActionInProgress"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGREQUIRESFULLURLCACHEOFF = "InvalidParameter.CacheKeyQueryStringRequiresFullUrlCacheOff"
//  INVALIDPARAMETER_CACHEKEYQUERYSTRINGTOOMANYVALUE = "InvalidParameter.CacheKeyQueryStringTooManyValue"
//  INVALIDPARAMETER_CERTSYSTEMERROR = "InvalidParameter.CertSystemError"
//  INVALIDPARAMETER_CLIENTIPCOUNTRYCONFLICTSWITHIPV6 = "InvalidParameter.ClientIpCountryConflictsWithIpv6"
//  INVALIDPARAMETER_GRPCREQUIREHTTP2 = "InvalidParameter.GrpcRequireHttp2"
//  INVALIDPARAMETER_INVALIDAWSPRIVATEACCESS = "InvalidParameter.InvalidAwsPrivateAccess"
//  INVALIDPARAMETER_INVALIDCACHECONFIGFOLLOWORIGIN = "InvalidParameter.InvalidCacheConfigFollowOrigin"
//  INVALIDPARAMETER_INVALIDCACHEKEYQUERYSTRINGVALUE = "InvalidParameter.InvalidCacheKeyQueryStringValue"
//  INVALIDPARAMETER_INVALIDCACHEONLYONSWITCH = "InvalidParameter.InvalidCacheOnlyOnSwitch"
//  INVALIDPARAMETER_INVALIDCACHETIME = "InvalidParameter.InvalidCacheTime"
//  INVALIDPARAMETER_INVALIDCLIENTIPCOUNTRYHEADERNAME = "InvalidParameter.InvalidClientIpCountryHeaderName"
//  INVALIDPARAMETER_INVALIDCLIENTIPHEADERNAME = "InvalidParameter.InvalidClientIpHeaderName"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINE = "InvalidParameter.InvalidDynamicRoutine"
//  INVALIDPARAMETER_INVALIDDYNAMICROUTINEBILLING = "InvalidParameter.InvalidDynamicRoutineBilling"
//  INVALIDPARAMETER_INVALIDFORCEREDIRECTTYPE = "InvalidParameter.InvalidForceRedirectType"
//  INVALIDPARAMETER_INVALIDHTTPS = "InvalidParameter.InvalidHttps"
//  INVALIDPARAMETER_INVALIDHTTPSCERTINFO = "InvalidParameter.InvalidHttpsCertInfo"
//  INVALIDPARAMETER_INVALIDHTTPSCIPHERSUITEANDTLSVERSION = "InvalidParameter.InvalidHttpsCipherSuiteAndTlsVersion"
//  INVALIDPARAMETER_INVALIDHTTPSHSTSMAXAGE = "InvalidParameter.InvalidHttpsHstsMaxAge"
//  INVALIDPARAMETER_INVALIDHTTPSTLSVERSION = "InvalidParameter.InvalidHttpsTlsVersion"
//  INVALIDPARAMETER_INVALIDIPV6SWITCH = "InvalidParameter.InvalidIpv6Switch"
//  INVALIDPARAMETER_INVALIDMAXAGETIME = "InvalidParameter.InvalidMaxAgeTime"
//  INVALIDPARAMETER_INVALIDORIGIN = "InvalidParameter.InvalidOrigin"
//  INVALIDPARAMETER_INVALIDORIGINTYPE = "InvalidParameter.InvalidOriginType"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_INVALIDPOSTMAXSIZEBILLING = "InvalidParameter.InvalidPostMaxSizeBilling"
//  INVALIDPARAMETER_INVALIDPOSTSIZEVALUE = "InvalidParameter.InvalidPostSizeValue"
//  INVALIDPARAMETER_INVALIDRANGEORIGINPULL = "InvalidParameter.InvalidRangeOriginPull"
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_INVALIDRESOURCEIDBILLING = "InvalidParameter.InvalidResourceIdBilling"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGCLIENTIP = "InvalidParameter.InvalidStandardDebugClientIp"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEMPTYLIST = "InvalidParameter.InvalidStandardDebugEmptyList"
//  INVALIDPARAMETER_INVALIDSTANDARDDEBUGEXPIRETIMELIMIT = "InvalidParameter.InvalidStandardDebugExpireTimeLimit"
//  INVALIDPARAMETER_INVALIDWEBSOCKETTIMEOUT = "InvalidParameter.InvalidWebSocketTimeout"
//  INVALIDPARAMETER_MULTIPLYLAYERNOTSUPPORTSMARTROUTING = "InvalidParameter.MultiplyLayerNotSupportSmartRouting"
//  INVALIDPARAMETER_OCDIRECTORIGINREQUIRESSMARTROUTING = "InvalidParameter.OCDirectOriginRequiresSmartRouting"
//  INVALIDPARAMETER_POSTMAXSIZELIMITEXCEEDED = "InvalidParameter.PostMaxSizeLimitExceeded"
//  INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"
//  INVALIDPARAMETER_ZONEISGRAYPUBLISHING = "InvalidParameter.ZoneIsGrayPublishing"
//  INVALIDPARAMETER_ZONENOTFOUND = "InvalidParameter.ZoneNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCELERATEMAINLANDDISABLE = "OperationDenied.AccelerateMainlandDisable"
//  OPERATIONDENIED_ACCELERATEMAINLANDIPV6CONFLICT = "OperationDenied.AccelerateMainlandIpv6Conflict"
//  OPERATIONDENIED_ACCELERATEMAINLANDMULTIPLYLAYERCONFLICT = "OperationDenied.AccelerateMainlandMultiplyLayerConflict"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_ERRZONEISALREADYPAUSED = "OperationDenied.ErrZoneIsAlreadyPaused"
//  OPERATIONDENIED_SHAREDCNAMEUNSUPPORTEDACCELERATEMAINLAND = "OperationDenied.SharedCNAMEUnsupportedAccelerateMainland"
//  OPERATIONDENIED_SHAREDCNAMEUNSUPPORTEDIPV6 = "OperationDenied.SharedCNAMEUnsupportedIPv6"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_POSTMAXSIZEQUOTANOTFOUND = "ResourceNotFound.PostMaxSizeQuotaNotFound"
//  RESOURCEUNAVAILABLE_CERTNOTFOUND = "ResourceUnavailable.CertNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyZoneSettingWithContext(ctx context.Context, request *ModifyZoneSettingRequest) (response *ModifyZoneSettingResponse, err error) {
    if request == nil {
        request = NewModifyZoneSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyZoneSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyZoneSettingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyZoneStatusRequest() (request *ModifyZoneStatusRequest) {
    request = &ModifyZoneStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "ModifyZoneStatus")
    
    
    return
}

func NewModifyZoneStatusResponse() (response *ModifyZoneStatusResponse) {
    response = &ModifyZoneStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyZoneStatus
// This API is used to change the site status.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_L4PROXYINPROGRESSSTATUS = "OperationDenied.L4ProxyInProgressStatus"
//  OPERATIONDENIED_L4PROXYINSTOPPINGSTATUS = "OperationDenied.L4ProxyInStoppingStatus"
//  OPERATIONDENIED_L7HOSTINPROCESSSTATUS = "OperationDenied.L7HostInProcessStatus"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) ModifyZoneStatus(request *ModifyZoneStatusRequest) (response *ModifyZoneStatusResponse, err error) {
    return c.ModifyZoneStatusWithContext(context.Background(), request)
}

// ModifyZoneStatus
// This API is used to change the site status.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_CONFIGLOCKED = "OperationDenied.ConfigLocked"
//  OPERATIONDENIED_DISABLEZONENOTCOMPLETED = "OperationDenied.DisableZoneNotCompleted"
//  OPERATIONDENIED_L4PROXYINPROGRESSSTATUS = "OperationDenied.L4ProxyInProgressStatus"
//  OPERATIONDENIED_L4PROXYINSTOPPINGSTATUS = "OperationDenied.L4ProxyInStoppingStatus"
//  OPERATIONDENIED_L7HOSTINPROCESSSTATUS = "OperationDenied.L7HostInProcessStatus"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  OPERATIONDENIED_VERSIONCONTROLISGRAYING = "OperationDenied.VersionControlIsGraying"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) ModifyZoneStatusWithContext(ctx context.Context, request *ModifyZoneStatusRequest) (response *ModifyZoneStatusResponse, err error) {
    if request == nil {
        request = NewModifyZoneStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyZoneStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyZoneStatusResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyOwnershipRequest() (request *VerifyOwnershipRequest) {
    request = &VerifyOwnershipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("teo", APIVersion, "VerifyOwnership")
    
    
    return
}

func NewVerifyOwnershipResponse() (response *VerifyOwnershipResponse) {
    response = &VerifyOwnershipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// VerifyOwnership
// This API is used to verify your ownership of a site or domain name. It's required in the CNAME access mode. After a site is verified, you don't need to verify the ownership again for domain names added to it in the future. For details, see [Ownership Verification](https://intl.cloud.tencent.com/document/product/1552/70789?from_cn_redirect=1).
//
// 
//
// For sites connected via the NS, you can query whether the NS is successfully switched through this API. For details, see [Modifying DNS Servers](https://intl.cloud.tencent.com/document/product/1552/90452?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) VerifyOwnership(request *VerifyOwnershipRequest) (response *VerifyOwnershipResponse, err error) {
    return c.VerifyOwnershipWithContext(context.Background(), request)
}

// VerifyOwnership
// This API is used to verify your ownership of a site or domain name. It's required in the CNAME access mode. After a site is verified, you don't need to verify the ownership again for domain names added to it in the future. For details, see [Ownership Verification](https://intl.cloud.tencent.com/document/product/1552/70789?from_cn_redirect=1).
//
// 
//
// For sites connected via the NS, you can query whether the NS is successfully switched through this API. For details, see [Modifying DNS Servers](https://intl.cloud.tencent.com/document/product/1552/90452?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  OPERATIONDENIED_RESOURCELOCKEDTEMPORARY = "OperationDenied.ResourceLockedTemporary"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) VerifyOwnershipWithContext(ctx context.Context, request *VerifyOwnershipRequest) (response *VerifyOwnershipResponse, err error) {
    if request == nil {
        request = NewVerifyOwnershipRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyOwnership require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyOwnershipResponse()
    err = c.Send(request, response)
    return
}
