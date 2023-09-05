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

package v20191205

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

// Predefined struct for user
type ApplyCertificateRequestParams struct {
	// Validation type. `DNS_AUTO`: automatic DNS validation; `DNS`: manual DNS validation; `FILE`: file validation
	DvAuthMethod *string `json:"DvAuthMethod,omitnil" name:"DvAuthMethod"`

	// Domain name
	DomainName *string `json:"DomainName,omitnil" name:"DomainName"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Certificate type. Currently, the only supported value is 2, which indicates TrustAsia TLS RSA CA.
	PackageType *string `json:"PackageType,omitnil" name:"PackageType"`

	// Email address
	ContactEmail *string `json:"ContactEmail,omitnil" name:"ContactEmail"`

	// Mobile number
	ContactPhone *string `json:"ContactPhone,omitnil" name:"ContactPhone"`

	// Validity period. The default value is 12 months, which is the only supported value currently.
	ValidityPeriod *string `json:"ValidityPeriod,omitnil" name:"ValidityPeriod"`

	// Encryption algorithm. RSA and ECC are supported.
	CsrEncryptAlgo *string `json:"CsrEncryptAlgo,omitnil" name:"CsrEncryptAlgo"`

	// Key pair parameter. RSA supports only the 2048-bit key and ECC supports only prime256v1.
	CsrKeyParameter *string `json:"CsrKeyParameter,omitnil" name:"CsrKeyParameter"`

	// CSR encryption password
	CsrKeyPassword *string `json:"CsrKeyPassword,omitnil" name:"CsrKeyPassword"`

	// Alias
	Alias *string `json:"Alias,omitnil" name:"Alias"`

	// Original certificate ID, which is used to apply for a new certificate.
	OldCertificateId *string `json:"OldCertificateId,omitnil" name:"OldCertificateId"`

	// Benefit package ID, which is used to expand the free certificate package
	PackageId *string `json:"PackageId,omitnil" name:"PackageId"`

	// Whether to delete the automatic domain name verification record after issuance, which is no by default. This parameter can be passed in only for domain names of the DNS_AUTO verification type.
	DeleteDnsAutoRecord *bool `json:"DeleteDnsAutoRecord,omitnil" name:"DeleteDnsAutoRecord"`
}

type ApplyCertificateRequest struct {
	*tchttp.BaseRequest
	
	// Validation type. `DNS_AUTO`: automatic DNS validation; `DNS`: manual DNS validation; `FILE`: file validation
	DvAuthMethod *string `json:"DvAuthMethod,omitnil" name:"DvAuthMethod"`

	// Domain name
	DomainName *string `json:"DomainName,omitnil" name:"DomainName"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Certificate type. Currently, the only supported value is 2, which indicates TrustAsia TLS RSA CA.
	PackageType *string `json:"PackageType,omitnil" name:"PackageType"`

	// Email address
	ContactEmail *string `json:"ContactEmail,omitnil" name:"ContactEmail"`

	// Mobile number
	ContactPhone *string `json:"ContactPhone,omitnil" name:"ContactPhone"`

	// Validity period. The default value is 12 months, which is the only supported value currently.
	ValidityPeriod *string `json:"ValidityPeriod,omitnil" name:"ValidityPeriod"`

	// Encryption algorithm. RSA and ECC are supported.
	CsrEncryptAlgo *string `json:"CsrEncryptAlgo,omitnil" name:"CsrEncryptAlgo"`

	// Key pair parameter. RSA supports only the 2048-bit key and ECC supports only prime256v1.
	CsrKeyParameter *string `json:"CsrKeyParameter,omitnil" name:"CsrKeyParameter"`

	// CSR encryption password
	CsrKeyPassword *string `json:"CsrKeyPassword,omitnil" name:"CsrKeyPassword"`

	// Alias
	Alias *string `json:"Alias,omitnil" name:"Alias"`

	// Original certificate ID, which is used to apply for a new certificate.
	OldCertificateId *string `json:"OldCertificateId,omitnil" name:"OldCertificateId"`

	// Benefit package ID, which is used to expand the free certificate package
	PackageId *string `json:"PackageId,omitnil" name:"PackageId"`

	// Whether to delete the automatic domain name verification record after issuance, which is no by default. This parameter can be passed in only for domain names of the DNS_AUTO verification type.
	DeleteDnsAutoRecord *bool `json:"DeleteDnsAutoRecord,omitnil" name:"DeleteDnsAutoRecord"`
}

func (r *ApplyCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DvAuthMethod")
	delete(f, "DomainName")
	delete(f, "ProjectId")
	delete(f, "PackageType")
	delete(f, "ContactEmail")
	delete(f, "ContactPhone")
	delete(f, "ValidityPeriod")
	delete(f, "CsrEncryptAlgo")
	delete(f, "CsrKeyParameter")
	delete(f, "CsrKeyPassword")
	delete(f, "Alias")
	delete(f, "OldCertificateId")
	delete(f, "PackageId")
	delete(f, "DeleteDnsAutoRecord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyCertificateResponseParams struct {
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil" name:"CertificateId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ApplyCertificateResponse struct {
	*tchttp.BaseResponse
	Response *ApplyCertificateResponseParams `json:"Response"`
}

func (r *ApplyCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelCertificateOrderRequestParams struct {
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil" name:"CertificateId"`
}

type CancelCertificateOrderRequest struct {
	*tchttp.BaseRequest
	
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil" name:"CertificateId"`
}

func (r *CancelCertificateOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelCertificateOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelCertificateOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelCertificateOrderResponseParams struct {
	// ID of the certificate whose order has been successfully cancelled
	CertificateId *string `json:"CertificateId,omitnil" name:"CertificateId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CancelCertificateOrderResponse struct {
	*tchttp.BaseResponse
	Response *CancelCertificateOrderResponseParams `json:"Response"`
}

func (r *CancelCertificateOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelCertificateOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CertificateExtra struct {
	// Number of domain names which can be associated with the certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	DomainNumber *string `json:"DomainNumber,omitnil" name:"DomainNumber"`

	// Original certificate ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	OriginCertificateId *string `json:"OriginCertificateId,omitnil" name:"OriginCertificateId"`

	// Original ID of the new certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	ReplacedBy *string `json:"ReplacedBy,omitnil" name:"ReplacedBy"`

	// New ID of the new certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	ReplacedFor *string `json:"ReplacedFor,omitnil" name:"ReplacedFor"`

	// Certificate ID of the new order
	// Note: this field may return null, indicating that no valid values can be obtained.
	RenewOrder *string `json:"RenewOrder,omitnil" name:"RenewOrder"`

	// Whether the certificate is a Chinese SM certificate.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SMCert *int64 `json:"SMCert,omitnil" name:"SMCert"`
}

type Certificates struct {
	// User UIN
	// Note: this field may return null, indicating that no valid values can be obtained.
	OwnerUin *string `json:"OwnerUin,omitnil" name:"OwnerUin"`

	// Project ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// Certificate source
	// Note: this field may return null, indicating that no valid values can be obtained.
	From *string `json:"From,omitnil" name:"From"`

	// The certificate plan type. Valid values:
	// null: Certificates uploaded by users (no plan type)
	// `1`: GeoTrust DV SSL CA - G3; `2`: TrustAsia TLS RSA CA; `3`: SecureSite EV Pro; `4`: SecureSite EV; `5`: SecureSite OV Pro; `6`: SecureSite OV; `7`: SecureSite OV wildcard; `8`: GeoTrust EV; `9`: GeoTrust OV; `10`: GeoTrust OV wildcard; `11`: TrustAsia DV multi-domain; `12`: TrustAsia DV wildcard; `13`: TrustAsia OV wildcard D3; `14`: TrustAsia OV D3; `15`: TrustAsia OV multi-domain D3; `16`: TrustAsia EV D3; `17`: TrustAsia EV multi-domain D3; `18`: GlobalSign OV; `19`: GlobalSign OV wildcard; `20`: GlobalSign EV; `21`: TrustAsia OV wildcard multi-domain D3; `22`: GlobalSign OV multi-domain; `23`: GlobalSign OV wildcard multi-domain; `24`: GlobalSign EV multi-domain
	// Note: This field may return null, indicating that no valid values can be obtained.
	PackageType *string `json:"PackageType,omitnil" name:"PackageType"`

	// Certificate type. `CA`: client certificate; `SVR`: server certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificateType *string `json:"CertificateType,omitnil" name:"CertificateType"`

	// Issuer
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProductZhName *string `json:"ProductZhName,omitnil" name:"ProductZhName"`

	// Primary domain name
	// Note: this field may return null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Alias
	// Note: this field may return null, indicating that no valid values can be obtained.
	Alias *string `json:"Alias,omitnil" name:"Alias"`

	// Status. `0`: Reviewing; `1`: Approved; `2`: Unapproved; `3`: Expired; `4`: DNS record added for domain names of the DNS_AUTO verification type; `5`: Enterprise-grade certificate, pending submission; `6`: Canceling order; `7`: Canceled; `8`: Information submitted, pending confirmation letter upload; `9`: Revoking certificate; `10`: Revoked; `11`: Reissuing; `12`: Pending revocation confirmation letter upload; `13`: Pending information submission for the free certificate; `14`: Refunded.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// Extended information of the certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificateExtra *CertificateExtra `json:"CertificateExtra,omitnil" name:"CertificateExtra"`

	// Vulnerability scanning status. `INACTIVE`: not activated; `ACTIVE`: activated
	// Note: this field may return null, indicating that no valid values can be obtained.
	VulnerabilityStatus *string `json:"VulnerabilityStatus,omitnil" name:"VulnerabilityStatus"`

	// Status information
	// Note: this field may return null, indicating that no valid values can be obtained.
	StatusMsg *string `json:"StatusMsg,omitnil" name:"StatusMsg"`

	// Validation type. `DNS_AUTO`: automatic DNS validation; `DNS`: manual DNS validation; `FILE`: file validation; `EMAIL`: email validation
	// Note: this field may return null, indicating that no valid values can be obtained.
	VerifyType *string `json:"VerifyType,omitnil" name:"VerifyType"`

	// Time when the certificate takes effect
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertBeginTime *string `json:"CertBeginTime,omitnil" name:"CertBeginTime"`

	// Time when the certificate expires
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertEndTime *string `json:"CertEndTime,omitnil" name:"CertEndTime"`

	// Validity period of the certificate, in months
	// Note: this field may return null, indicating that no valid values can be obtained.
	ValidityPeriod *string `json:"ValidityPeriod,omitnil" name:"ValidityPeriod"`

	// Creation time
	// Note: this field may return null, indicating that no valid values can be obtained.
	InsertTime *string `json:"InsertTime,omitnil" name:"InsertTime"`

	// Certificate ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificateId *string `json:"CertificateId,omitnil" name:"CertificateId"`

	// Domain names associated with the certificate (including the primary domain name)
	// Note: this field may return null, indicating that no valid values can be obtained.
	SubjectAltName []*string `json:"SubjectAltName,omitnil" name:"SubjectAltName"`

	// Certificate type name
	// Note: this field may return null, indicating that no valid values can be obtained.
	PackageTypeName *string `json:"PackageTypeName,omitnil" name:"PackageTypeName"`

	// Status description
	// Note: this field may return null, indicating that no valid values can be obtained.
	StatusName *string `json:"StatusName,omitnil" name:"StatusName"`

	// Whether the customer is a VIP customer
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsVip *bool `json:"IsVip,omitnil" name:"IsVip"`

	// Whether the certificate is a DV certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsDv *bool `json:"IsDv,omitnil" name:"IsDv"`

	// Whether the certificate is a wildcard certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsWildcard *bool `json:"IsWildcard,omitnil" name:"IsWildcard"`

	// Whether the vulnerability scanning feature is enabled
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsVulnerability *bool `json:"IsVulnerability,omitnil" name:"IsVulnerability"`

	// Whether the certificate can be reissued
	// Note: this field may return null, indicating that no valid values can be obtained.
	RenewAble *bool `json:"RenewAble,omitnil" name:"RenewAble"`

	// Project information
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectInfo *ProjectInfo `json:"ProjectInfo,omitnil" name:"ProjectInfo"`

	// Associated Tencent Cloud services. Currently, this parameter is unavailable.
	// Note: this field may return null, indicating that no valid values can be obtained.
	BoundResource []*string `json:"BoundResource,omitnil" name:"BoundResource"`

	// Whether the certificate can be deployed
	// Note: this field may return null, indicating that no valid values can be obtained.
	Deployable *bool `json:"Deployable,omitnil" name:"Deployable"`

	// List of tags
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Tags []*Tags `json:"Tags,omitnil" name:"Tags"`

	// Whether the expiration notification was ignored
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsIgnore *bool `json:"IsIgnore,omitnil" name:"IsIgnore"`

	// Whether the certificate is a Chinese SM certificate
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsSM *bool `json:"IsSM,omitnil" name:"IsSM"`

	// Certificate algorithm
	// Note: This field may return null, indicating that no valid values can be obtained.
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitnil" name:"EncryptAlgorithm"`

	// Encryption algorithm of the uploaded CA certificate
	// Note: This field may return null, indicating that no valid values can be obtained.
	CAEncryptAlgorithms []*string `json:"CAEncryptAlgorithms,omitnil" name:"CAEncryptAlgorithms"`

	// Expiration time of the uploaded CA certificate
	// Note: This field may return null, indicating that no valid values can be obtained.
	CAEndTimes []*string `json:"CAEndTimes,omitnil" name:"CAEndTimes"`

	// Generic name of the uploaded CA certificate
	// Note: This field may return null, indicating that no valid values can be obtained.
	CACommonNames []*string `json:"CACommonNames,omitnil" name:"CACommonNames"`

	// Prereview information of the certificate
	// Note: This field may return null, indicating that no valid values can be obtained.
	PreAuditInfo *PreAuditInfo `json:"PreAuditInfo,omitnil" name:"PreAuditInfo"`

	// Whether auto-renewal is enabled.
	// Note: This field may return null, indicating that no valid value can be obtained.
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`
}

// Predefined struct for user
type CommitCertificateInformationRequestParams struct {
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil" name:"CertificateId"`

	// Domain validation method
	VerifyType *string `json:"VerifyType,omitnil" name:"VerifyType"`
}

type CommitCertificateInformationRequest struct {
	*tchttp.BaseRequest
	
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil" name:"CertificateId"`

	// Domain validation method
	VerifyType *string `json:"VerifyType,omitnil" name:"VerifyType"`
}

func (r *CommitCertificateInformationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CommitCertificateInformationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "VerifyType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CommitCertificateInformationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CommitCertificateInformationResponseParams struct {
	// TrustAsia order ID
	OrderId *string `json:"OrderId,omitnil" name:"OrderId"`

	// Certificate status. `0`: reviewing; `1`: approved; `2`: unapproved; `3`: expired; `4`: DNS record added; `5`: enterprise-grade certificate, pending submission; `6`: canceling order; `7`: canceled; `8`: information submitted, pending confirmation letter upload; `9`: revoking certificate; `10`: revoked; `11`: reissuing; `12`: pending revocation confirmation letter upload
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CommitCertificateInformationResponse struct {
	*tchttp.BaseResponse
	Response *CommitCertificateInformationResponseParams `json:"Response"`
}

func (r *CommitCertificateInformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CommitCertificateInformationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCertificateRequestParams struct {
	// Certificate product ID. `3`: SecureSite EV Pro; `4`: SecureSite EV; `5`: SecureSite OV Pro; `6`: SecureSite OV; `7`: SecureSite OV wildcard; `8`: GeoTrust EV; `9`: GeoTrust OV; `10`: GeoTrust OV wildcard; `11`: TrustAsia DV multi-domain; `12`: TrustAsia DV wildcard; `13`: TrustAsia OV wildcard D3; `14`: TrustAsia OV D3; `15`: TrustAsia OV multi-domain D3; `16`: TrustAsia EV D3; `17`: TrustAsia EV multi-domain D3; `18`: GlobalSign OV; `19`: GlobalSign OV wildcard; `20`: GlobalSign EV; `21`: TrustAsia OV wildcard multi-domain D3; `22`: GlobalSign OV multi-domain; `23`: GlobalSign OV wildcard multi-domain; `24`: GlobalSign EV multi-domain; `25` WoTrus DV; `26`: WoTrus DV multi-domain; `27`: WoTrus DV wildcard; `28`: WoTrus OV; `29`: WoTrus OV multi-domain; `30`: WoTrus OV wildcard; `31`: WoTrus EV; `32`: WoTrus EV multi-domain; `33`: DNSPod SM2 DV; `34`: DNSPod SM2 DV multi-domain; `35`: DNSPod SM2 DV wildcard; `37`: DNSPod SM2 OV; `38`: DNSPod SM2 OV multi-domain; `39`: DNSPod SM2 OV wildcard: `40`: DNSPod SM2 EV; `41`: DNSPod SM2 EV multi-domain; `42`: TrustAsia DV wildcard multi-domain.
	ProductId *int64 `json:"ProductId,omitnil" name:"ProductId"`

	// Number of domains associated with the certificate
	DomainNum *int64 `json:"DomainNum,omitnil" name:"DomainNum"`

	// Certificate validity period. Currently, you can only purchase 1-year certificates.
	TimeSpan *int64 `json:"TimeSpan,omitnil" name:"TimeSpan"`
}

type CreateCertificateRequest struct {
	*tchttp.BaseRequest
	
	// Certificate product ID. `3`: SecureSite EV Pro; `4`: SecureSite EV; `5`: SecureSite OV Pro; `6`: SecureSite OV; `7`: SecureSite OV wildcard; `8`: GeoTrust EV; `9`: GeoTrust OV; `10`: GeoTrust OV wildcard; `11`: TrustAsia DV multi-domain; `12`: TrustAsia DV wildcard; `13`: TrustAsia OV wildcard D3; `14`: TrustAsia OV D3; `15`: TrustAsia OV multi-domain D3; `16`: TrustAsia EV D3; `17`: TrustAsia EV multi-domain D3; `18`: GlobalSign OV; `19`: GlobalSign OV wildcard; `20`: GlobalSign EV; `21`: TrustAsia OV wildcard multi-domain D3; `22`: GlobalSign OV multi-domain; `23`: GlobalSign OV wildcard multi-domain; `24`: GlobalSign EV multi-domain; `25` WoTrus DV; `26`: WoTrus DV multi-domain; `27`: WoTrus DV wildcard; `28`: WoTrus OV; `29`: WoTrus OV multi-domain; `30`: WoTrus OV wildcard; `31`: WoTrus EV; `32`: WoTrus EV multi-domain; `33`: DNSPod SM2 DV; `34`: DNSPod SM2 DV multi-domain; `35`: DNSPod SM2 DV wildcard; `37`: DNSPod SM2 OV; `38`: DNSPod SM2 OV multi-domain; `39`: DNSPod SM2 OV wildcard: `40`: DNSPod SM2 EV; `41`: DNSPod SM2 EV multi-domain; `42`: TrustAsia DV wildcard multi-domain.
	ProductId *int64 `json:"ProductId,omitnil" name:"ProductId"`

	// Number of domains associated with the certificate
	DomainNum *int64 `json:"DomainNum,omitnil" name:"DomainNum"`

	// Certificate validity period. Currently, you can only purchase 1-year certificates.
	TimeSpan *int64 `json:"TimeSpan,omitnil" name:"TimeSpan"`
}

func (r *CreateCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DomainNum")
	delete(f, "TimeSpan")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCertificateResponseParams struct {
	// List of certificate IDs
	CertificateIds []*string `json:"CertificateIds,omitnil" name:"CertificateIds"`

	// List of order IDs
	DealIds []*string `json:"DealIds,omitnil" name:"DealIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCertificateResponse struct {
	*tchttp.BaseResponse
	Response *CreateCertificateResponseParams `json:"Response"`
}

func (r *CreateCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCertificateRequestParams struct {
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil" name:"CertificateId"`
}

type DeleteCertificateRequest struct {
	*tchttp.BaseRequest
	
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil" name:"CertificateId"`
}

func (r *DeleteCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCertificateResponseParams struct {
	// Deletion result
	DeleteResult *bool `json:"DeleteResult,omitnil" name:"DeleteResult"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteCertificateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCertificateResponseParams `json:"Response"`
}

func (r *DeleteCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificateDetailRequestParams struct {
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil" name:"CertificateId"`
}

type DescribeCertificateDetailRequest struct {
	*tchttp.BaseRequest
	
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil" name:"CertificateId"`
}

func (r *DescribeCertificateDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCertificateDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificateDetailResponseParams struct {
	// User UIN
	// Note: this field may return null, indicating that no valid values can be obtained.
	OwnerUin *string `json:"OwnerUin,omitnil" name:"OwnerUin"`

	// Project ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// Certificate source. `trustasia`: TrustAsia; `upload`: certificate uploaded by users
	// Note: this field may return null, indicating that no valid values can be obtained.
	From *string `json:"From,omitnil" name:"From"`

	// Certificate type. `CA`: client certificate; `SVR`: server certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificateType *string `json:"CertificateType,omitnil" name:"CertificateType"`

	// Certificate plan type. null: User-uploaded certificate (no plan type); `1`: GeoTrust DV SSL CA - G3; `2`: TrustAsia TLS RSA CA; `3`: SecureSite EV Pro; `4`: SecureSite EV; `5`: SecureSite OV Pro; `6`: SecureSite OV; `7`: SecureSite OV wildcard; `8`: GeoTrust EV; `9`: GeoTrust OV; `10`: GeoTrust OV wildcard; `11`: TrustAsia DV multi-domain; `12`: TrustAsia DV wildcard; `13`: TrustAsia OV wildcard D3; `14`: TrustAsia OV D3; `15`: TrustAsia OV multi-domain D3; `16`: TrustAsia EV D3; `17`: TrustAsia EV multi-domain D3; `18`: GlobalSign OV; `19`: GlobalSign OV wildcard; `20`: GlobalSign EV; `21`: TrustAsia OV wildcard multi-domain D3; `22`: GlobalSign OV multi-domain; `23`: GlobalSign OV wildcard multi-domain; `24`: GlobalSign EV multi-domain; `25` WoTrus DV; `26`: WoTrus DV multi-domain; `27`: WoTrus DV wildcard; `28`: WoTrus OV; `29`: WoTrus OV multi-domain; `30`: WoTrus OV wildcard; `31`: WoTrus EV; `32`: WoTrus EV multi-domain; `33`: DNSPod SM2 DV; `34`: DNSPod SM2 DV multi-domain; `35`: DNSPod SM2 DV wildcard; `37`: DNSPod SM2 OV; `38`: DNSPod SM2 OV multi-domain; `39`: DNSPod SM2 OV wildcard: `40`: DNSPod SM2 EV; `41`: DNSPod SM2 EV multi-domain; `42`: TrustAsia DV wildcard multi-domain.
	PackageType *string `json:"PackageType,omitnil" name:"PackageType"`

	// Issuer
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProductZhName *string `json:"ProductZhName,omitnil" name:"ProductZhName"`

	// Domain name
	// Note: this field may return null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Alias
	// Note: this field may return null, indicating that no valid values can be obtained.
	Alias *string `json:"Alias,omitnil" name:"Alias"`

	// Certificate status. `0`: reviewing; `1`: approved; `2`: unapproved; `3`: expired; `4`: DNS record added; `5`: enterprise-grade certificate, pending submission; `6`: canceling order; `7`: canceled; `8`: information submitted, pending confirmation letter upload; `9`: revoking certificate; `10`: revoked; `11`: reissuing; `12`: pending revocation confirmation letter upload
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// Status information
	// Note: this field may return null, indicating that no valid values can be obtained.
	StatusMsg *string `json:"StatusMsg,omitnil" name:"StatusMsg"`

	// Validation type. `DNS_AUTO`: automatic DNS validation; `DNS`: manual DNS validation; `FILE`: file validation; `EMAIL`: email validation
	// Note: this field may return null, indicating that no valid values can be obtained.
	VerifyType *string `json:"VerifyType,omitnil" name:"VerifyType"`

	// Vulnerability scanning status
	// Note: this field may return null, indicating that no valid values can be obtained.
	VulnerabilityStatus *string `json:"VulnerabilityStatus,omitnil" name:"VulnerabilityStatus"`

	// Time when the certificate takes effect
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertBeginTime *string `json:"CertBeginTime,omitnil" name:"CertBeginTime"`

	// Time when the certificate expires
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertEndTime *string `json:"CertEndTime,omitnil" name:"CertEndTime"`

	// Validity period of the certificate, in months
	// Note: this field may return null, indicating that no valid values can be obtained.
	ValidityPeriod *string `json:"ValidityPeriod,omitnil" name:"ValidityPeriod"`

	// Application time
	// Note: this field may return null, indicating that no valid values can be obtained.
	InsertTime *string `json:"InsertTime,omitnil" name:"InsertTime"`

	// Order ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrderId *string `json:"OrderId,omitnil" name:"OrderId"`

	// Extended information of the certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificateExtra *CertificateExtra `json:"CertificateExtra,omitnil" name:"CertificateExtra"`

	// Private key of the certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificatePrivateKey *string `json:"CertificatePrivateKey,omitnil" name:"CertificatePrivateKey"`

	// Public key of the certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificatePublicKey *string `json:"CertificatePublicKey,omitnil" name:"CertificatePublicKey"`

	// DV authentication information
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthDetail *DvAuthDetail `json:"DvAuthDetail,omitnil" name:"DvAuthDetail"`

	// Vulnerability scanning assessment report
	// Note: this field may return null, indicating that no valid values can be obtained.
	VulnerabilityReport *string `json:"VulnerabilityReport,omitnil" name:"VulnerabilityReport"`

	// Certificate ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificateId *string `json:"CertificateId,omitnil" name:"CertificateId"`

	// Certificate type name
	// Note: this field may return null, indicating that no valid values can be obtained.
	TypeName *string `json:"TypeName,omitnil" name:"TypeName"`

	// Status description
	// Note: this field may return null, indicating that no valid values can be obtained.
	StatusName *string `json:"StatusName,omitnil" name:"StatusName"`

	// Multiple domain names included in the certificate (excluding the primary domain name, which uses the `Domain` field)
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubjectAltName []*string `json:"SubjectAltName,omitnil" name:"SubjectAltName"`

	// Whether the certificate is a paid one.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsVip *bool `json:"IsVip,omitnil" name:"IsVip"`

	// Whether the certificate is a wildcard certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsWildcard *bool `json:"IsWildcard,omitnil" name:"IsWildcard"`

	// Whether the certificate is a DV certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsDv *bool `json:"IsDv,omitnil" name:"IsDv"`

	// Whether the vulnerability scanning feature is enabled
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsVulnerability *bool `json:"IsVulnerability,omitnil" name:"IsVulnerability"`

	// Submitted data
	// Note: this field may return null, indicating that no valid values can be obtained.
	SubmittedData *SubmittedData `json:"SubmittedData,omitnil" name:"SubmittedData"`

	// Whether the certificate can be renewed.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RenewAble *bool `json:"RenewAble,omitnil" name:"RenewAble"`

	// Whether the certificate can be deployed
	// Note: this field may return null, indicating that no valid values can be obtained.
	Deployable *bool `json:"Deployable,omitnil" name:"Deployable"`

	// List of associated tags
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Tags []*Tags `json:"Tags,omitnil" name:"Tags"`

	// Root certificate.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RootCert *RootCertificates `json:"RootCert,omitnil" name:"RootCert"`

	// Chinese SM encryption certificate
	// Note: This field may return null, indicating that no valid values can be obtained.
	EncryptCert *string `json:"EncryptCert,omitnil" name:"EncryptCert"`

	// Private key of Chinese SM encryption
	// Note: This field may return null, indicating that no valid values can be obtained.
	EncryptPrivateKey *string `json:"EncryptPrivateKey,omitnil" name:"EncryptPrivateKey"`

	// SHA1 fingerprint of the signature certificate
	// Note: This field may return null, indicating that no valid values can be obtained.
	CertFingerprint *string `json:"CertFingerprint,omitnil" name:"CertFingerprint"`

	// SHA1 fingerprint of the encryption certificate (for Chinese SM certificates only)
	// Note: This field may return null, indicating that no valid values can be obtained.
	EncryptCertFingerprint *string `json:"EncryptCertFingerprint,omitnil" name:"EncryptCertFingerprint"`

	// Certificate algorithm
	// Note: This field may return null, indicating that no valid values can be obtained.
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitnil" name:"EncryptAlgorithm"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCertificateDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCertificateDetailResponseParams `json:"Response"`
}

func (r *DescribeCertificateDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificateOperateLogsRequestParams struct {
	// Offset. The default value is 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of requested logs. The default value is 20.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Start time. The default value is 15 days ago.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End time. The default value is the current time.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

type DescribeCertificateOperateLogsRequest struct {
	*tchttp.BaseRequest
	
	// Offset. The default value is 0.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of requested logs. The default value is 20.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Start time. The default value is 15 days ago.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// End time. The default value is the current time.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

func (r *DescribeCertificateOperateLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateOperateLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCertificateOperateLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificateOperateLogsResponseParams struct {
	// Total number of logs that meet query conditions
	AllTotal *uint64 `json:"AllTotal,omitnil" name:"AllTotal"`

	// Number of logs returned for this request
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Certificate operation log list
	// Note: this field may return null, indicating that no valid values can be obtained.
	OperateLogs []*OperationLog `json:"OperateLogs,omitnil" name:"OperateLogs"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCertificateOperateLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCertificateOperateLogsResponseParams `json:"Response"`
}

func (r *DescribeCertificateOperateLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateOperateLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificateRequestParams struct {
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil" name:"CertificateId"`
}

type DescribeCertificateRequest struct {
	*tchttp.BaseRequest
	
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil" name:"CertificateId"`
}

func (r *DescribeCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificateResponseParams struct {
	// User UIN
	// Note: this field may return null, indicating that no valid values can be obtained.
	OwnerUin *string `json:"OwnerUin,omitnil" name:"OwnerUin"`

	// Project ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// Certificate source. `trustasia`: TrustAsia; `upload`: certificate uploaded by users
	// Note: this field may return null, indicating that no valid values can be obtained.
	From *string `json:"From,omitnil" name:"From"`

	// Certificate type. `CA`: client certificate; `SVR`: server certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificateType *string `json:"CertificateType,omitnil" name:"CertificateType"`

	// Certificate plan type. `1`: GeoTrust DV SSL CA - G3; `2`: TrustAsia TLS RSA CA; `3`: SecureSite EV Pro; `4`: SecureSite EV; `5`: SecureSite OV Pro; `6`: SecureSite OV; `7`: SecureSite OV wildcard; `8`: GeoTrust EV; `9`: GeoTrust OV; `10`: GeoTrust OV wildcard; `11`: TrustAsia DV multi-domain; `12`: TrustAsia DV wildcard; `13`: TrustAsia OV wildcard D3; `14`: TrustAsia OV D3; `15`: TrustAsia OV multi-domain D3; `16`: TrustAsia EV D3; `17`: TrustAsia EV multi-domain D3; `18`: GlobalSign OV; `19`: GlobalSign OV wildcard; `20`: GlobalSign EV; `21`: TrustAsia OV wildcard multi-domain D3; `22`: GlobalSign OV multi-domain; `23`: GlobalSign OV wildcard multi-domain; `24`: GlobalSign EV multi-domain
	// Note: this field may return null, indicating that no valid values can be obtained.
	PackageType *string `json:"PackageType,omitnil" name:"PackageType"`

	// Name of the certificate issuer
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProductZhName *string `json:"ProductZhName,omitnil" name:"ProductZhName"`

	// Domain name
	// Note: this field may return null, indicating that no valid values can be obtained.
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Alias
	// Note: this field may return null, indicating that no valid values can be obtained.
	Alias *string `json:"Alias,omitnil" name:"Alias"`

	// Certificate status. `0`: reviewing; `1`: approved; `2`: unapproved; `3`: expired; `4`: DNS record added; `5`: enterprise-grade certificate, pending submission; `6`: canceling order; `7`: canceled; `8`: information submitted, pending confirmation letter upload; `9`: revoking certificate; `10`: revoked; `11`: reissuing; `12`: pending revocation confirmation letter upload
	// Note: this field may return null, indicating that no valid values can be obtained.
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// Status information
	// Note: this field may return null, indicating that no valid values can be obtained.
	StatusMsg *string `json:"StatusMsg,omitnil" name:"StatusMsg"`

	// Validation type. `DNS_AUTO`: automatic DNS validation; `DNS`: manual DNS validation; `FILE`: file validation; `EMAIL`: email validation
	// Note: this field may return null, indicating that no valid values can be obtained.
	VerifyType *string `json:"VerifyType,omitnil" name:"VerifyType"`

	// Vulnerability scanning status
	// Note: this field may return null, indicating that no valid values can be obtained.
	VulnerabilityStatus *string `json:"VulnerabilityStatus,omitnil" name:"VulnerabilityStatus"`

	// Time when the certificate takes effect
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertBeginTime *string `json:"CertBeginTime,omitnil" name:"CertBeginTime"`

	// Time when the certificate expires
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertEndTime *string `json:"CertEndTime,omitnil" name:"CertEndTime"`

	// Validity period of the certificate, in months
	// Note: this field may return null, indicating that no valid values can be obtained.
	ValidityPeriod *string `json:"ValidityPeriod,omitnil" name:"ValidityPeriod"`

	// Application time
	// Note: this field may return null, indicating that no valid values can be obtained.
	InsertTime *string `json:"InsertTime,omitnil" name:"InsertTime"`

	// Order ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrderId *string `json:"OrderId,omitnil" name:"OrderId"`

	// Extended information of the certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificateExtra *CertificateExtra `json:"CertificateExtra,omitnil" name:"CertificateExtra"`

	// DV authentication information
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthDetail *DvAuthDetail `json:"DvAuthDetail,omitnil" name:"DvAuthDetail"`

	// Vulnerability scanning assessment report
	// Note: this field may return null, indicating that no valid values can be obtained.
	VulnerabilityReport *string `json:"VulnerabilityReport,omitnil" name:"VulnerabilityReport"`

	// Certificate ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificateId *string `json:"CertificateId,omitnil" name:"CertificateId"`

	// Certificate type name
	// Note: this field may return null, indicating that no valid values can be obtained.
	PackageTypeName *string `json:"PackageTypeName,omitnil" name:"PackageTypeName"`

	// Status description
	// Note: this field may return null, indicating that no valid values can be obtained.
	StatusName *string `json:"StatusName,omitnil" name:"StatusName"`

	// Domain names associated with the certificate (including the primary domain name)
	// Note: this field may return null, indicating that no valid values can be obtained.
	SubjectAltName []*string `json:"SubjectAltName,omitnil" name:"SubjectAltName"`

	// Whether the customer is a VIP customer
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsVip *bool `json:"IsVip,omitnil" name:"IsVip"`

	// Whether the certificate is a wildcard certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsWildcard *bool `json:"IsWildcard,omitnil" name:"IsWildcard"`

	// Whether the certificate is a DV certificate
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsDv *bool `json:"IsDv,omitnil" name:"IsDv"`

	// Whether the vulnerability scanning feature is enabled
	// Note: this field may return null, indicating that no valid values can be obtained.
	IsVulnerability *bool `json:"IsVulnerability,omitnil" name:"IsVulnerability"`

	// Whether the certificate can be reissued
	// Note: this field may return null, indicating that no valid values can be obtained.
	RenewAble *bool `json:"RenewAble,omitnil" name:"RenewAble"`

	// Submitted data
	// Note: this field may return null, indicating that no valid values can be obtained.
	SubmittedData *SubmittedData `json:"SubmittedData,omitnil" name:"SubmittedData"`

	// Whether the certificate can be deployed
	// Note: this field may return null, indicating that no valid values can be obtained.
	Deployable *bool `json:"Deployable,omitnil" name:"Deployable"`

	// List of tags
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Tags []*Tags `json:"Tags,omitnil" name:"Tags"`

	// All encryption algorithms of a CA certificate
	// Note: This field may return null, indicating that no valid values can be obtained.
	CAEncryptAlgorithms []*string `json:"CAEncryptAlgorithms,omitnil" name:"CAEncryptAlgorithms"`

	// All common names of a CA certificate
	// Note: This field may return null, indicating that no valid values can be obtained.
	CACommonNames []*string `json:"CACommonNames,omitnil" name:"CACommonNames"`

	// All expiration time of a CA certificate
	// Note: This field may return null, indicating that no valid values can be obtained.
	CAEndTimes []*string `json:"CAEndTimes,omitnil" name:"CAEndTimes"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCertificateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCertificateResponseParams `json:"Response"`
}

func (r *DescribeCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificatesRequestParams struct {
	// Pagination offset, starting from 0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of entries per page. Default value: `20`. Maximum value: `1000`.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Keyword for search, which can be a certificate ID, alias, or domain name, for example, a8xHcaIs
	SearchKey *string `json:"SearchKey,omitnil" name:"SearchKey"`

	// Certificate type. `CA`: client certificate; `SVR`: server certificate
	CertificateType *string `json:"CertificateType,omitnil" name:"CertificateType"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Sorting by expiration time. `DESC`: descending; `ASC`: ascending
	ExpirationSort *string `json:"ExpirationSort,omitnil" name:"ExpirationSort"`

	// Certificate status. `0`: Reviewing; `1`: Approved; `2`: Unapproved; `3`: Expired; `4`: DNS record added; `5`: Enterprise-grade certificate, pending submission; `6`: Canceling order; `7`: Canceled; `8`: Information submitted, pending confirmation letter upload; `9`: Revoking certificate; `10`: Revoked; `11`: Reissuing; `12`: Pending revocation confirmation letter upload; `13`: Pending information submission for the free certificate.
	CertificateStatus []*uint64 `json:"CertificateStatus,omitnil" name:"CertificateStatus"`

	// Whether the certificate can be deployed. `1`: yes; `0`: no
	Deployable *uint64 `json:"Deployable,omitnil" name:"Deployable"`

	// Whether to filter uploaded hosted certificates. `1`: Yes; `0`: No.
	Upload *int64 `json:"Upload,omitnil" name:"Upload"`

	// Whether to filter renewable certificates. `1`: Yes; `0`: No.
	Renew *int64 `json:"Renew,omitnil" name:"Renew"`

	// Filter by source. `upload`: Uploaded certificate; `buy`: Tencent Cloud certificate. If this parameter is left empty, all certificates will be queried.
	FilterSource *string `json:"FilterSource,omitnil" name:"FilterSource"`

	// Whether to filter Chinese SM certificates. `1`: Yes; `0`: No.
	IsSM *int64 `json:"IsSM,omitnil" name:"IsSM"`

	// Whether to filter expiring certificates. `1`: Yes; `0`: No.
	FilterExpiring *uint64 `json:"FilterExpiring,omitnil" name:"FilterExpiring"`
}

type DescribeCertificatesRequest struct {
	*tchttp.BaseRequest
	
	// Pagination offset, starting from 0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of entries per page. Default value: `20`. Maximum value: `1000`.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Keyword for search, which can be a certificate ID, alias, or domain name, for example, a8xHcaIs
	SearchKey *string `json:"SearchKey,omitnil" name:"SearchKey"`

	// Certificate type. `CA`: client certificate; `SVR`: server certificate
	CertificateType *string `json:"CertificateType,omitnil" name:"CertificateType"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// Sorting by expiration time. `DESC`: descending; `ASC`: ascending
	ExpirationSort *string `json:"ExpirationSort,omitnil" name:"ExpirationSort"`

	// Certificate status. `0`: Reviewing; `1`: Approved; `2`: Unapproved; `3`: Expired; `4`: DNS record added; `5`: Enterprise-grade certificate, pending submission; `6`: Canceling order; `7`: Canceled; `8`: Information submitted, pending confirmation letter upload; `9`: Revoking certificate; `10`: Revoked; `11`: Reissuing; `12`: Pending revocation confirmation letter upload; `13`: Pending information submission for the free certificate.
	CertificateStatus []*uint64 `json:"CertificateStatus,omitnil" name:"CertificateStatus"`

	// Whether the certificate can be deployed. `1`: yes; `0`: no
	Deployable *uint64 `json:"Deployable,omitnil" name:"Deployable"`

	// Whether to filter uploaded hosted certificates. `1`: Yes; `0`: No.
	Upload *int64 `json:"Upload,omitnil" name:"Upload"`

	// Whether to filter renewable certificates. `1`: Yes; `0`: No.
	Renew *int64 `json:"Renew,omitnil" name:"Renew"`

	// Filter by source. `upload`: Uploaded certificate; `buy`: Tencent Cloud certificate. If this parameter is left empty, all certificates will be queried.
	FilterSource *string `json:"FilterSource,omitnil" name:"FilterSource"`

	// Whether to filter Chinese SM certificates. `1`: Yes; `0`: No.
	IsSM *int64 `json:"IsSM,omitnil" name:"IsSM"`

	// Whether to filter expiring certificates. `1`: Yes; `0`: No.
	FilterExpiring *uint64 `json:"FilterExpiring,omitnil" name:"FilterExpiring"`
}

func (r *DescribeCertificatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchKey")
	delete(f, "CertificateType")
	delete(f, "ProjectId")
	delete(f, "ExpirationSort")
	delete(f, "CertificateStatus")
	delete(f, "Deployable")
	delete(f, "Upload")
	delete(f, "Renew")
	delete(f, "FilterSource")
	delete(f, "IsSM")
	delete(f, "FilterExpiring")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCertificatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificatesResponseParams struct {
	// Total number
	// Note: this field may return null, indicating that no valid values can be obtained.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// List
	// Note: this field may return null, indicating that no valid values can be obtained.
	Certificates []*Certificates `json:"Certificates,omitnil" name:"Certificates"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCertificatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCertificatesResponseParams `json:"Response"`
}

func (r *DescribeCertificatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadCertificateRequestParams struct {
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil" name:"CertificateId"`
}

type DownloadCertificateRequest struct {
	*tchttp.BaseRequest
	
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil" name:"CertificateId"`
}

func (r *DownloadCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadCertificateResponseParams struct {
	// ZIP content encoded by using Base64. After the content is decoded by using Base64, it can be saved as a ZIP file.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Content *string `json:"Content,omitnil" name:"Content"`

	// MIME type. `application/zip`: ZIP file
	// Note: this field may return null, indicating that no valid values can be obtained.
	ContentType *string `json:"ContentType,omitnil" name:"ContentType"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DownloadCertificateResponse struct {
	*tchttp.BaseResponse
	Response *DownloadCertificateResponseParams `json:"Response"`
}

func (r *DownloadCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DvAuthDetail struct {
	// DV authentication key
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthKey *string `json:"DvAuthKey,omitnil" name:"DvAuthKey"`

	// DV authentication value
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthValue *string `json:"DvAuthValue,omitnil" name:"DvAuthValue"`

	// Domain name of the DV authentication value
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthDomain *string `json:"DvAuthDomain,omitnil" name:"DvAuthDomain"`

	// Path of the DV authentication value
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthPath *string `json:"DvAuthPath,omitnil" name:"DvAuthPath"`

	// DV authentication sub-domain name
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthKeySubDomain *string `json:"DvAuthKeySubDomain,omitnil" name:"DvAuthKeySubDomain"`

	// DV authentication information
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuths []*DvAuths `json:"DvAuths,omitnil" name:"DvAuths"`
}

type DvAuths struct {
	// DV authentication key
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthKey *string `json:"DvAuthKey,omitnil" name:"DvAuthKey"`

	// DV authentication value
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthValue *string `json:"DvAuthValue,omitnil" name:"DvAuthValue"`

	// Domain name of the DV authentication value
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthDomain *string `json:"DvAuthDomain,omitnil" name:"DvAuthDomain"`

	// Path of the DV authentication value
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthPath *string `json:"DvAuthPath,omitnil" name:"DvAuthPath"`

	// DV authentication sub-domain name
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthSubDomain *string `json:"DvAuthSubDomain,omitnil" name:"DvAuthSubDomain"`

	// DV authentication type
	// Note: this field may return null, indicating that no valid values can be obtained.
	DvAuthVerifyType *string `json:"DvAuthVerifyType,omitnil" name:"DvAuthVerifyType"`
}

// Predefined struct for user
type ModifyCertificateAliasRequestParams struct {
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil" name:"CertificateId"`

	// Alias
	Alias *string `json:"Alias,omitnil" name:"Alias"`
}

type ModifyCertificateAliasRequest struct {
	*tchttp.BaseRequest
	
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil" name:"CertificateId"`

	// Alias
	Alias *string `json:"Alias,omitnil" name:"Alias"`
}

func (r *ModifyCertificateAliasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificateAliasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "Alias")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCertificateAliasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCertificateAliasResponseParams struct {
	// ID of the successfully modified certificate
	CertificateId *string `json:"CertificateId,omitnil" name:"CertificateId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCertificateAliasResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCertificateAliasResponseParams `json:"Response"`
}

func (r *ModifyCertificateAliasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificateAliasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCertificateProjectRequestParams struct {
	// ID list of certificates whose projects need to be modified. A maximum of 100 certificate IDs are supported.
	CertificateIdList []*string `json:"CertificateIdList,omitnil" name:"CertificateIdList"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`
}

type ModifyCertificateProjectRequest struct {
	*tchttp.BaseRequest
	
	// ID list of certificates whose projects need to be modified. A maximum of 100 certificate IDs are supported.
	CertificateIdList []*string `json:"CertificateIdList,omitnil" name:"CertificateIdList"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`
}

func (r *ModifyCertificateProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificateProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateIdList")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCertificateProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCertificateProjectResponseParams struct {
	// List of certificates whose projects were modified successfully
	// Note: this field may return null, indicating that no valid values can be obtained.
	SuccessCertificates []*string `json:"SuccessCertificates,omitnil" name:"SuccessCertificates"`

	// List of certificates whose projects failed to be modified
	// Note: this field may return null, indicating that no valid values can be obtained.
	FailCertificates []*string `json:"FailCertificates,omitnil" name:"FailCertificates"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCertificateProjectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCertificateProjectResponseParams `json:"Response"`
}

func (r *ModifyCertificateProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificateProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperationLog struct {
	// Action performed on logs
	Action *string `json:"Action,omitnil" name:"Action"`

	// Time when the action is performed
	CreatedOn *string `json:"CreatedOn,omitnil" name:"CreatedOn"`
}

type PreAuditInfo struct {
	// Total number of years of the certificate
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalPeriod *int64 `json:"TotalPeriod,omitnil" name:"TotalPeriod"`

	// Current year of the certificate
	// Note: This field may return null, indicating that no valid values can be obtained.
	NowPeriod *int64 `json:"NowPeriod,omitnil" name:"NowPeriod"`

	// Certificate prereview manager ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	ManagerId *string `json:"ManagerId,omitnil" name:"ManagerId"`
}

type ProjectInfo struct {
	// Project name
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectName *string `json:"ProjectName,omitnil" name:"ProjectName"`

	// UIN of the project creator
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectCreatorUin *uint64 `json:"ProjectCreatorUin,omitnil" name:"ProjectCreatorUin"`

	// Project creation time
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectCreateTime *string `json:"ProjectCreateTime,omitnil" name:"ProjectCreateTime"`

	// Brief project information
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectResume *string `json:"ProjectResume,omitnil" name:"ProjectResume"`

	// User UIN
	// Note: this field may return null, indicating that no valid values can be obtained.
	OwnerUin *uint64 `json:"OwnerUin,omitnil" name:"OwnerUin"`

	// Project ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`
}

// Predefined struct for user
type ReplaceCertificateRequestParams struct {
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil" name:"CertificateId"`

	// Validation type. `DNS_AUTO`: automatic DNS validation; `DNS`: manual DNS validation; `FILE`: file validation
	ValidType *string `json:"ValidType,omitnil" name:"ValidType"`

	// Type. `original`: original certificate CSR; `upload`: uploaded manually; `online`: generated online. The default value is original.
	CsrType *string `json:"CsrType,omitnil" name:"CsrType"`

	// CSR content
	CsrContent *string `json:"CsrContent,omitnil" name:"CsrContent"`

	// Password of the key
	CsrkeyPassword *string `json:"CsrkeyPassword,omitnil" name:"CsrkeyPassword"`

	// Reissue reason
	Reason *string `json:"Reason,omitnil" name:"Reason"`
}

type ReplaceCertificateRequest struct {
	*tchttp.BaseRequest
	
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil" name:"CertificateId"`

	// Validation type. `DNS_AUTO`: automatic DNS validation; `DNS`: manual DNS validation; `FILE`: file validation
	ValidType *string `json:"ValidType,omitnil" name:"ValidType"`

	// Type. `original`: original certificate CSR; `upload`: uploaded manually; `online`: generated online. The default value is original.
	CsrType *string `json:"CsrType,omitnil" name:"CsrType"`

	// CSR content
	CsrContent *string `json:"CsrContent,omitnil" name:"CsrContent"`

	// Password of the key
	CsrkeyPassword *string `json:"CsrkeyPassword,omitnil" name:"CsrkeyPassword"`

	// Reissue reason
	Reason *string `json:"Reason,omitnil" name:"Reason"`
}

func (r *ReplaceCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "ValidType")
	delete(f, "CsrType")
	delete(f, "CsrContent")
	delete(f, "CsrkeyPassword")
	delete(f, "Reason")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReplaceCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceCertificateResponseParams struct {
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil" name:"CertificateId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ReplaceCertificateResponse struct {
	*tchttp.BaseResponse
	Response *ReplaceCertificateResponseParams `json:"Response"`
}

func (r *ReplaceCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RootCertificates struct {
	// Chinese SM signature certificate
	// Note: This field may return null, indicating that no valid values can be obtained.
	Sign *string `json:"Sign,omitnil" name:"Sign"`

	// Chinese SM encryption certificate
	// Note: This field may return null, indicating that no valid values can be obtained.
	Encrypt *string `json:"Encrypt,omitnil" name:"Encrypt"`

	// Standard certificate
	// Note: This field may return null, indicating that no valid values can be obtained.
	Standard *string `json:"Standard,omitnil" name:"Standard"`
}

// Predefined struct for user
type SubmitCertificateInformationRequestParams struct {
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil" name:"CertificateId"`

	// CSR generation mode. `online`: generated online; `parse`: uploaded manually
	CsrType *string `json:"CsrType,omitnil" name:"CsrType"`

	// Uploaded CSR content
	CsrContent *string `json:"CsrContent,omitnil" name:"CsrContent"`

	// Domain name bound with the certificate
	CertificateDomain *string `json:"CertificateDomain,omitnil" name:"CertificateDomain"`

	// Uploaded domain name array (can be uploaded for a multi-domain certificate)
	DomainList []*string `json:"DomainList,omitnil" name:"DomainList"`

	// Password of the private key
	KeyPassword *string `json:"KeyPassword,omitnil" name:"KeyPassword"`

	// Organization name
	OrganizationName *string `json:"OrganizationName,omitnil" name:"OrganizationName"`

	// Division name
	OrganizationDivision *string `json:"OrganizationDivision,omitnil" name:"OrganizationDivision"`

	// Detailed address of the organization
	OrganizationAddress *string `json:"OrganizationAddress,omitnil" name:"OrganizationAddress"`

	// Country where the organization is located, for example, CN (China)
	OrganizationCountry *string `json:"OrganizationCountry,omitnil" name:"OrganizationCountry"`

	// City where the organization is located
	OrganizationCity *string `json:"OrganizationCity,omitnil" name:"OrganizationCity"`

	// Province where the organization is located
	OrganizationRegion *string `json:"OrganizationRegion,omitnil" name:"OrganizationRegion"`

	// Postal code of the organization
	PostalCode *string `json:"PostalCode,omitnil" name:"PostalCode"`

	// Area code of the fixed-line phone number of the organization
	PhoneAreaCode *string `json:"PhoneAreaCode,omitnil" name:"PhoneAreaCode"`

	// Fixed-line phone number of the organization
	PhoneNumber *string `json:"PhoneNumber,omitnil" name:"PhoneNumber"`

	// Certificate validation method
	VerifyType *string `json:"VerifyType,omitnil" name:"VerifyType"`

	// Last name of the administrator
	AdminFirstName *string `json:"AdminFirstName,omitnil" name:"AdminFirstName"`

	// First name of the administrator
	AdminLastName *string `json:"AdminLastName,omitnil" name:"AdminLastName"`

	// Mobile number of the administrator
	AdminPhoneNum *string `json:"AdminPhoneNum,omitnil" name:"AdminPhoneNum"`

	// Email of the administrator
	AdminEmail *string `json:"AdminEmail,omitnil" name:"AdminEmail"`

	// Position of the administrator
	AdminPosition *string `json:"AdminPosition,omitnil" name:"AdminPosition"`

	// Last name of the contact
	ContactFirstName *string `json:"ContactFirstName,omitnil" name:"ContactFirstName"`

	// First name of the contact
	ContactLastName *string `json:"ContactLastName,omitnil" name:"ContactLastName"`

	// Email of the contact
	ContactEmail *string `json:"ContactEmail,omitnil" name:"ContactEmail"`

	// Mobile number of the contact
	ContactNumber *string `json:"ContactNumber,omitnil" name:"ContactNumber"`

	// Position of the contact
	ContactPosition *string `json:"ContactPosition,omitnil" name:"ContactPosition"`
}

type SubmitCertificateInformationRequest struct {
	*tchttp.BaseRequest
	
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil" name:"CertificateId"`

	// CSR generation mode. `online`: generated online; `parse`: uploaded manually
	CsrType *string `json:"CsrType,omitnil" name:"CsrType"`

	// Uploaded CSR content
	CsrContent *string `json:"CsrContent,omitnil" name:"CsrContent"`

	// Domain name bound with the certificate
	CertificateDomain *string `json:"CertificateDomain,omitnil" name:"CertificateDomain"`

	// Uploaded domain name array (can be uploaded for a multi-domain certificate)
	DomainList []*string `json:"DomainList,omitnil" name:"DomainList"`

	// Password of the private key
	KeyPassword *string `json:"KeyPassword,omitnil" name:"KeyPassword"`

	// Organization name
	OrganizationName *string `json:"OrganizationName,omitnil" name:"OrganizationName"`

	// Division name
	OrganizationDivision *string `json:"OrganizationDivision,omitnil" name:"OrganizationDivision"`

	// Detailed address of the organization
	OrganizationAddress *string `json:"OrganizationAddress,omitnil" name:"OrganizationAddress"`

	// Country where the organization is located, for example, CN (China)
	OrganizationCountry *string `json:"OrganizationCountry,omitnil" name:"OrganizationCountry"`

	// City where the organization is located
	OrganizationCity *string `json:"OrganizationCity,omitnil" name:"OrganizationCity"`

	// Province where the organization is located
	OrganizationRegion *string `json:"OrganizationRegion,omitnil" name:"OrganizationRegion"`

	// Postal code of the organization
	PostalCode *string `json:"PostalCode,omitnil" name:"PostalCode"`

	// Area code of the fixed-line phone number of the organization
	PhoneAreaCode *string `json:"PhoneAreaCode,omitnil" name:"PhoneAreaCode"`

	// Fixed-line phone number of the organization
	PhoneNumber *string `json:"PhoneNumber,omitnil" name:"PhoneNumber"`

	// Certificate validation method
	VerifyType *string `json:"VerifyType,omitnil" name:"VerifyType"`

	// Last name of the administrator
	AdminFirstName *string `json:"AdminFirstName,omitnil" name:"AdminFirstName"`

	// First name of the administrator
	AdminLastName *string `json:"AdminLastName,omitnil" name:"AdminLastName"`

	// Mobile number of the administrator
	AdminPhoneNum *string `json:"AdminPhoneNum,omitnil" name:"AdminPhoneNum"`

	// Email of the administrator
	AdminEmail *string `json:"AdminEmail,omitnil" name:"AdminEmail"`

	// Position of the administrator
	AdminPosition *string `json:"AdminPosition,omitnil" name:"AdminPosition"`

	// Last name of the contact
	ContactFirstName *string `json:"ContactFirstName,omitnil" name:"ContactFirstName"`

	// First name of the contact
	ContactLastName *string `json:"ContactLastName,omitnil" name:"ContactLastName"`

	// Email of the contact
	ContactEmail *string `json:"ContactEmail,omitnil" name:"ContactEmail"`

	// Mobile number of the contact
	ContactNumber *string `json:"ContactNumber,omitnil" name:"ContactNumber"`

	// Position of the contact
	ContactPosition *string `json:"ContactPosition,omitnil" name:"ContactPosition"`
}

func (r *SubmitCertificateInformationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitCertificateInformationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "CsrType")
	delete(f, "CsrContent")
	delete(f, "CertificateDomain")
	delete(f, "DomainList")
	delete(f, "KeyPassword")
	delete(f, "OrganizationName")
	delete(f, "OrganizationDivision")
	delete(f, "OrganizationAddress")
	delete(f, "OrganizationCountry")
	delete(f, "OrganizationCity")
	delete(f, "OrganizationRegion")
	delete(f, "PostalCode")
	delete(f, "PhoneAreaCode")
	delete(f, "PhoneNumber")
	delete(f, "VerifyType")
	delete(f, "AdminFirstName")
	delete(f, "AdminLastName")
	delete(f, "AdminPhoneNum")
	delete(f, "AdminEmail")
	delete(f, "AdminPosition")
	delete(f, "ContactFirstName")
	delete(f, "ContactLastName")
	delete(f, "ContactEmail")
	delete(f, "ContactNumber")
	delete(f, "ContactPosition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitCertificateInformationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitCertificateInformationResponseParams struct {
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil" name:"CertificateId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SubmitCertificateInformationResponse struct {
	*tchttp.BaseResponse
	Response *SubmitCertificateInformationResponseParams `json:"Response"`
}

func (r *SubmitCertificateInformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitCertificateInformationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubmittedData struct {
	// CSR type. `online`: CSR generated online; `parse`: CSR pasted
	// Note: this field may return null, indicating that no valid values can be obtained.
	CsrType *string `json:"CsrType,omitnil" name:"CsrType"`

	// CSR content
	// Note: this field may return null, indicating that no valid values can be obtained.
	CsrContent *string `json:"CsrContent,omitnil" name:"CsrContent"`

	// Domain name information
	// Note: this field may return null, indicating that no valid values can be obtained.
	CertificateDomain *string `json:"CertificateDomain,omitnil" name:"CertificateDomain"`

	// DNS information
	// Note: this field may return null, indicating that no valid values can be obtained.
	DomainList []*string `json:"DomainList,omitnil" name:"DomainList"`

	// Password of the private key
	// Note: this field may return null, indicating that no valid values can be obtained.
	KeyPassword *string `json:"KeyPassword,omitnil" name:"KeyPassword"`

	// Enterprise or unit name
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrganizationName *string `json:"OrganizationName,omitnil" name:"OrganizationName"`

	// Division
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrganizationDivision *string `json:"OrganizationDivision,omitnil" name:"OrganizationDivision"`

	// Address
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrganizationAddress *string `json:"OrganizationAddress,omitnil" name:"OrganizationAddress"`

	// Country
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrganizationCountry *string `json:"OrganizationCountry,omitnil" name:"OrganizationCountry"`

	// City
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrganizationCity *string `json:"OrganizationCity,omitnil" name:"OrganizationCity"`

	// Province
	// Note: this field may return null, indicating that no valid values can be obtained.
	OrganizationRegion *string `json:"OrganizationRegion,omitnil" name:"OrganizationRegion"`

	// Postal code
	// Note: this field may return null, indicating that no valid values can be obtained.
	PostalCode *string `json:"PostalCode,omitnil" name:"PostalCode"`

	// Area code of the fixed-line phone number
	// Note: this field may return null, indicating that no valid values can be obtained.
	PhoneAreaCode *string `json:"PhoneAreaCode,omitnil" name:"PhoneAreaCode"`

	// Fixed-line phone number
	// Note: this field may return null, indicating that no valid values can be obtained.
	PhoneNumber *string `json:"PhoneNumber,omitnil" name:"PhoneNumber"`

	// First name of the administrator
	// Note: this field may return null, indicating that no valid values can be obtained.
	AdminFirstName *string `json:"AdminFirstName,omitnil" name:"AdminFirstName"`

	// Last name of the administrator
	// Note: this field may return null, indicating that no valid values can be obtained.
	AdminLastName *string `json:"AdminLastName,omitnil" name:"AdminLastName"`

	// Phone number of the administrator
	// Note: this field may return null, indicating that no valid values can be obtained.
	AdminPhoneNum *string `json:"AdminPhoneNum,omitnil" name:"AdminPhoneNum"`

	// Email of the administrator
	// Note: this field may return null, indicating that no valid values can be obtained.
	AdminEmail *string `json:"AdminEmail,omitnil" name:"AdminEmail"`

	// Position of the administrator
	// Note: this field may return null, indicating that no valid values can be obtained.
	AdminPosition *string `json:"AdminPosition,omitnil" name:"AdminPosition"`

	// First name of the contact
	// Note: this field may return null, indicating that no valid values can be obtained.
	ContactFirstName *string `json:"ContactFirstName,omitnil" name:"ContactFirstName"`

	// Last name of the contact
	// Note: this field may return null, indicating that no valid values can be obtained.
	ContactLastName *string `json:"ContactLastName,omitnil" name:"ContactLastName"`

	// Phone number of the contact
	// Note: this field may return null, indicating that no valid values can be obtained.
	ContactNumber *string `json:"ContactNumber,omitnil" name:"ContactNumber"`

	// Email of the contact
	// Note: this field may return null, indicating that no valid values can be obtained.
	ContactEmail *string `json:"ContactEmail,omitnil" name:"ContactEmail"`

	// Position of the contact
	// Note: this field may return null, indicating that no valid values can be obtained.
	ContactPosition *string `json:"ContactPosition,omitnil" name:"ContactPosition"`

	// Validation type
	// Note: this field may return null, indicating that no valid values can be obtained.
	VerifyType *string `json:"VerifyType,omitnil" name:"VerifyType"`
}

type Tags struct {
	// Tag key
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// Tag value
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}

// Predefined struct for user
type UploadCertificateRequestParams struct {
	// Public key of the certificate
	CertificatePublicKey *string `json:"CertificatePublicKey,omitnil" name:"CertificatePublicKey"`

	// Private key content. This parameter is required when the certificate type is SVR, and not required when the certificate type is CA.
	CertificatePrivateKey *string `json:"CertificatePrivateKey,omitnil" name:"CertificatePrivateKey"`

	// Certificate type. Valid values: `CA` (CA certificate) and `SVR` (server certificate). Default value: `SVR`
	CertificateType *string `json:"CertificateType,omitnil" name:"CertificateType"`

	// Alias
	Alias *string `json:"Alias,omitnil" name:"Alias"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`


	CertificateUse *string `json:"CertificateUse,omitnil" name:"CertificateUse"`

	// Whether a certificate can be repeatedly uploaded.
	Repeatable *bool `json:"Repeatable,omitnil" name:"Repeatable"`
}

type UploadCertificateRequest struct {
	*tchttp.BaseRequest
	
	// Public key of the certificate
	CertificatePublicKey *string `json:"CertificatePublicKey,omitnil" name:"CertificatePublicKey"`

	// Private key content. This parameter is required when the certificate type is SVR, and not required when the certificate type is CA.
	CertificatePrivateKey *string `json:"CertificatePrivateKey,omitnil" name:"CertificatePrivateKey"`

	// Certificate type. Valid values: `CA` (CA certificate) and `SVR` (server certificate). Default value: `SVR`
	CertificateType *string `json:"CertificateType,omitnil" name:"CertificateType"`

	// Alias
	Alias *string `json:"Alias,omitnil" name:"Alias"`

	// Project ID
	ProjectId *uint64 `json:"ProjectId,omitnil" name:"ProjectId"`

	CertificateUse *string `json:"CertificateUse,omitnil" name:"CertificateUse"`

	// Whether a certificate can be repeatedly uploaded.
	Repeatable *bool `json:"Repeatable,omitnil" name:"Repeatable"`
}

func (r *UploadCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificatePublicKey")
	delete(f, "CertificatePrivateKey")
	delete(f, "CertificateType")
	delete(f, "Alias")
	delete(f, "ProjectId")
	delete(f, "CertificateUse")
	delete(f, "Repeatable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadCertificateResponseParams struct {
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil" name:"CertificateId"`

	// The ID of the repeatedly uploaded certificate.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RepeatCertId *string `json:"RepeatCertId,omitnil" name:"RepeatCertId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UploadCertificateResponse struct {
	*tchttp.BaseResponse
	Response *UploadCertificateResponseParams `json:"Response"`
}

func (r *UploadCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadConfirmLetterRequestParams struct {
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil" name:"CertificateId"`

	// Base64-encoded confirmation letter file, which must be a JPG, JPEG, PNG, or PDF file of 1 KB to 1.4 MB
	ConfirmLetter *string `json:"ConfirmLetter,omitnil" name:"ConfirmLetter"`
}

type UploadConfirmLetterRequest struct {
	*tchttp.BaseRequest
	
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil" name:"CertificateId"`

	// Base64-encoded confirmation letter file, which must be a JPG, JPEG, PNG, or PDF file of 1 KB to 1.4 MB
	ConfirmLetter *string `json:"ConfirmLetter,omitnil" name:"ConfirmLetter"`
}

func (r *UploadConfirmLetterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadConfirmLetterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "ConfirmLetter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadConfirmLetterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadConfirmLetterResponseParams struct {
	// Certificate ID
	CertificateId *string `json:"CertificateId,omitnil" name:"CertificateId"`

	// Whether the operation is successful
	IsSuccess *bool `json:"IsSuccess,omitnil" name:"IsSuccess"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UploadConfirmLetterResponse struct {
	*tchttp.BaseResponse
	Response *UploadConfirmLetterResponseParams `json:"Response"`
}

func (r *UploadConfirmLetterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadConfirmLetterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}