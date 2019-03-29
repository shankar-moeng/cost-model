// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// AUTO-GENERATED CODE. DO NOT EDIT.

// Package digitalassetlinks provides access to the Digital Asset Links API.
//
// See https://developers.google.com/digital-asset-links/
//
// Usage example:
//
//   import "google.golang.org/api/digitalassetlinks/v1"
//   ...
//   digitalassetlinksService, err := digitalassetlinks.New(oauthHttpClient)
package digitalassetlinks // import "google.golang.org/api/digitalassetlinks/v1"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled

const apiId = "digitalassetlinks:v1"
const apiName = "digitalassetlinks"
const apiVersion = "v1"
const basePath = "https://digitalassetlinks.googleapis.com/"

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Assetlinks = NewAssetlinksService(s)
	s.Statements = NewStatementsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Assetlinks *AssetlinksService

	Statements *StatementsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewAssetlinksService(s *Service) *AssetlinksService {
	rs := &AssetlinksService{s: s}
	return rs
}

type AssetlinksService struct {
	s *Service
}

func NewStatementsService(s *Service) *StatementsService {
	rs := &StatementsService{s: s}
	return rs
}

type StatementsService struct {
	s *Service
}

// AndroidAppAsset: Describes an android app asset.
type AndroidAppAsset struct {
	// Certificate: Because there is no global enforcement of package name
	// uniqueness, we also
	// require a signing certificate, which in combination with the package
	// name
	// uniquely identifies an app.
	//
	// Some apps' signing keys are rotated, so they may be signed by
	// different
	// keys over time.  We treat these as distinct assets, since we use
	// (package
	// name, cert) as the unique ID.  This should not normally pose any
	// problems
	// as both versions of the app will make the same or similar
	// statements.
	// Other assets making statements about the app will have to be updated
	// when a
	// key is rotated, however.
	//
	// (Note that the syntaxes for publishing and querying for statements
	// contain
	// syntactic sugar to easily let you specify apps that are known by
	// multiple
	// certificates.)
	// REQUIRED
	Certificate *CertificateInfo `json:"certificate,omitempty"`

	// PackageName: Android App assets are naturally identified by their
	// Java package name.
	// For example, the Google Maps app uses the package
	// name
	// `com.google.android.apps.maps`.
	// REQUIRED
	PackageName string `json:"packageName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Certificate") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Certificate") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AndroidAppAsset) MarshalJSON() ([]byte, error) {
	type NoMethod AndroidAppAsset
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Asset: Uniquely identifies an asset.
//
// A digital asset is an identifiable and addressable online entity
// that
// typically provides some service or content.  Examples of assets are
// websites,
// Android apps, Twitter feeds, and Plus Pages.
type Asset struct {
	// AndroidApp: Set if this is an Android App asset.
	AndroidApp *AndroidAppAsset `json:"androidApp,omitempty"`

	// Web: Set if this is a web asset.
	Web *WebAsset `json:"web,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AndroidApp") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AndroidApp") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Asset) MarshalJSON() ([]byte, error) {
	type NoMethod Asset
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CertificateInfo: Describes an X509 certificate.
type CertificateInfo struct {
	// Sha256Fingerprint: The uppercase SHA-265 fingerprint of the
	// certificate.  From the PEM
	//  certificate, it can be acquired like this:
	//
	//     $ keytool -printcert -file $CERTFILE | grep SHA256:
	//     SHA256: 14:6D:E9:83:C5:73:06:50:D8:EE:B9:95:2F:34:FC:64:16:A0:83:
	// \
	//         42:E6:1D:BE:A8:8A:04:96:B2:3F:CF:44:E5
	//
	// or like this:
	//
	//     $ openssl x509 -in $CERTFILE -noout -fingerprint -sha256
	//     SHA256
	// Fingerprint=14:6D:E9:83:C5:73:06:50:D8:EE:B9:95:2F:34:FC:64: \
	//         16:A0:83:42:E6:1D:BE:A8:8A:04:96:B2:3F:CF:44:E5
	//
	// In this example, the contents of this field would be
	// `14:6D:E9:83:C5:73:
	// 06:50:D8:EE:B9:95:2F:34:FC:64:16:A0:83:42:E6:1D:BE
	// :A8:8A:04:96:B2:3F:CF:
	// 44:E5`.
	//
	// If these tools are not available to you, you can convert the
	// PEM
	// certificate into the DER format, compute the SHA-256 hash of that
	// string
	// and represent the result as a hexstring (that is, uppercase
	// hexadecimal
	// representations of each octet, separated by colons).
	Sha256Fingerprint string `json:"sha256Fingerprint,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Sha256Fingerprint")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Sha256Fingerprint") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *CertificateInfo) MarshalJSON() ([]byte, error) {
	type NoMethod CertificateInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CheckResponse: Response message for the CheckAssetLinks call.
type CheckResponse struct {
	// DebugString: Human-readable message containing information intended
	// to help end users
	// understand, reproduce and debug the result.
	//
	//
	// The message will be in English and we are currently not planning to
	// offer
	// any translations.
	//
	// Please note that no guarantees are made about the contents or format
	// of
	// this string.  Any aspect of it may be subject to change without
	// notice.
	// You should not attempt to programmatically parse this data.
	// For
	// programmatic access, use the error_code field below.
	DebugString string `json:"debugString,omitempty"`

	// ErrorCode: Error codes that describe the result of the Check
	// operation.
	//
	// Possible values:
	//   "ERROR_CODE_UNSPECIFIED"
	//   "ERROR_CODE_INVALID_QUERY" - Unable to parse query.
	//   "ERROR_CODE_FETCH_ERROR" - Unable to fetch the asset links data.
	//   "ERROR_CODE_FAILED_SSL_VALIDATION" - Invalid HTTPS certificate .
	//   "ERROR_CODE_REDIRECT" - HTTP redirects (e.g, 301) are not allowed.
	//   "ERROR_CODE_TOO_LARGE" - Asset links data exceeds maximum size.
	//   "ERROR_CODE_MALFORMED_HTTP_RESPONSE" - Can't parse HTTP response.
	//   "ERROR_CODE_WRONG_CONTENT_TYPE" - HTTP Content-type should be
	// application/json.
	//   "ERROR_CODE_MALFORMED_CONTENT" - JSON content is malformed.
	//   "ERROR_CODE_SECURE_ASSET_INCLUDES_INSECURE" - A secure asset
	// includes an insecure asset (security downgrade).
	//   "ERROR_CODE_FETCH_BUDGET_EXHAUSTED" - Too many includes (maybe a
	// loop).
	ErrorCode []string `json:"errorCode,omitempty"`

	// Linked: Set to true if the assets specified in the request are linked
	// by the
	// relation specified in the request.
	Linked bool `json:"linked,omitempty"`

	// MaxAge: From serving time, how much longer the response should be
	// considered valid
	// barring further updates.
	// REQUIRED
	MaxAge string `json:"maxAge,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "DebugString") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DebugString") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CheckResponse) MarshalJSON() ([]byte, error) {
	type NoMethod CheckResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListResponse: Response message for the List call.
type ListResponse struct {
	// DebugString: Human-readable message containing information intended
	// to help end users
	// understand, reproduce and debug the result.
	//
	//
	// The message will be in English and we are currently not planning to
	// offer
	// any translations.
	//
	// Please note that no guarantees are made about the contents or format
	// of
	// this string.  Any aspect of it may be subject to change without
	// notice.
	// You should not attempt to programmatically parse this data.
	// For
	// programmatic access, use the error_code field below.
	DebugString string `json:"debugString,omitempty"`

	// ErrorCode: Error codes that describe the result of the List
	// operation.
	//
	// Possible values:
	//   "ERROR_CODE_UNSPECIFIED"
	//   "ERROR_CODE_INVALID_QUERY" - Unable to parse query.
	//   "ERROR_CODE_FETCH_ERROR" - Unable to fetch the asset links data.
	//   "ERROR_CODE_FAILED_SSL_VALIDATION" - Invalid HTTPS certificate .
	//   "ERROR_CODE_REDIRECT" - HTTP redirects (e.g, 301) are not allowed.
	//   "ERROR_CODE_TOO_LARGE" - Asset links data exceeds maximum size.
	//   "ERROR_CODE_MALFORMED_HTTP_RESPONSE" - Can't parse HTTP response.
	//   "ERROR_CODE_WRONG_CONTENT_TYPE" - HTTP Content-type should be
	// application/json.
	//   "ERROR_CODE_MALFORMED_CONTENT" - JSON content is malformed.
	//   "ERROR_CODE_SECURE_ASSET_INCLUDES_INSECURE" - A secure asset
	// includes an insecure asset (security downgrade).
	//   "ERROR_CODE_FETCH_BUDGET_EXHAUSTED" - Too many includes (maybe a
	// loop).
	ErrorCode []string `json:"errorCode,omitempty"`

	// MaxAge: From serving time, how much longer the response should be
	// considered valid
	// barring further updates.
	// REQUIRED
	MaxAge string `json:"maxAge,omitempty"`

	// Statements: A list of all the matching statements that have been
	// found.
	Statements []*Statement `json:"statements,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "DebugString") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DebugString") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Statement: Describes a reliable statement that has been made about
// the relationship
// between a source asset and a target asset.
//
// Statements are always made by the source asset, either directly or
// by
// delegating to a statement list that is stored elsewhere.
//
// For more detailed definitions of statements and assets, please
// refer
// to our [API documentation
// landing
// page](/digital-asset-links/v1/getting-started).
type Statement struct {
	// Relation: The relation identifies the use of the statement as
	// intended by the source
	// asset's owner (that is, the person or entity who issued the
	// statement).
	// Every complete statement has a relation.
	//
	// We identify relations with strings of the format `<kind>/<detail>`,
	// where
	// `<kind>` must be one of a set of pre-defined purpose categories,
	// and
	// `<detail>` is a free-form lowercase alphanumeric string that
	// describes the
	// specific use case of the statement.
	//
	// Refer to [our API
	// documentation](/digital-asset-links/v1/relation-strings)
	// for the current list of supported relations.
	//
	// Example: `delegate_permission/common.handle_all_urls`
	// REQUIRED
	Relation string `json:"relation,omitempty"`

	// Source: Every statement has a source asset.
	// REQUIRED
	Source *Asset `json:"source,omitempty"`

	// Target: Every statement has a target asset.
	// REQUIRED
	Target *Asset `json:"target,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Relation") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Relation") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Statement) MarshalJSON() ([]byte, error) {
	type NoMethod Statement
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// WebAsset: Describes a web asset.
type WebAsset struct {
	// Site: Web assets are identified by a URL that contains only the
	// scheme, hostname
	// and port parts.  The format is
	//
	//     http[s]://<hostname>[:<port>]
	//
	// Hostnames must be fully qualified: they must end in a single
	// period
	// (".").
	//
	// Only the schemes "http" and "https" are currently allowed.
	//
	// Port numbers are given as a decimal number, and they must be omitted
	// if the
	// standard port numbers are used: 80 for http and 443 for https.
	//
	// We call this limited URL the "site".  All URLs that share the same
	// scheme,
	// hostname and port are considered to be a part of the site and thus
	// belong
	// to the web asset.
	//
	// Example: the asset with the site `https://www.google.com` contains
	// all
	// these URLs:
	//
	//   *   `https://www.google.com/`
	//   *   `https://www.google.com:443/`
	//   *   `https://www.google.com/foo`
	//   *   `https://www.google.com/foo?bar`
	//   *   `https://www.google.com/foo#bar`
	//   *   `https://user@password:www.google.com/`
	//
	// But it does not contain these URLs:
	//
	//   *   `http://www.google.com/`       (wrong scheme)
	//   *   `https://google.com/`          (hostname does not match)
	//   *   `https://www.google.com:444/`  (port does not match)
	// REQUIRED
	Site string `json:"site,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Site") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Site") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *WebAsset) MarshalJSON() ([]byte, error) {
	type NoMethod WebAsset
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "digitalassetlinks.assetlinks.check":

type AssetlinksCheckCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Check: Determines whether the specified (directional) relationship
// exists between
// the specified source and target assets.
//
// The relation describes the intent of the link between the two assets
// as
// claimed by the source asset.  An example for such relationships is
// the
// delegation of privileges or permissions.
//
// This command is most often used by infrastructure systems to
// check
// preconditions for an action.  For example, a client may want to know
// if it
// is OK to send a web URL to a particular mobile app instead. The
// client can
// check for the relevant asset link from the website to the mobile app
// to
// decide if the operation should be allowed.
//
// A note about security: if you specify a secure asset as the source,
// such as
// an HTTPS website or an Android app, the API will ensure that
// any
// statements used to generate the response have been made in a secure
// way by
// the owner of that asset.  Conversely, if the source asset is an
// insecure
// HTTP website (that is, the URL starts with `http://` instead of
// `https://`),
// the API cannot verify its statements securely, and it is not possible
// to
// ensure that the website's statements have not been altered by a
// third
// party.  For more information, see the [Digital Asset Links technical
// design
// specification](https://github.com/google/digitalassetlinks/blob
// /master/well-known/details.md).
func (r *AssetlinksService) Check() *AssetlinksCheckCall {
	c := &AssetlinksCheckCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// Relation sets the optional parameter "relation": Query string for the
// relation.
//
// We identify relations with strings of the format `<kind>/<detail>`,
// where
// `<kind>` must be one of a set of pre-defined purpose categories,
// and
// `<detail>` is a free-form lowercase alphanumeric string that
// describes the
// specific use case of the statement.
//
// Refer to [our API
// documentation](/digital-asset-links/v1/relation-strings)
// for the current list of supported relations.
//
// For a query to match an asset link, both the query's and the asset
// link's
// relation strings must match exactly.
//
// Example: A query with relation
// `delegate_permission/common.handle_all_urls`
// matches an asset link with
// relation
// `delegate_permission/common.handle_all_urls`.
func (c *AssetlinksCheckCall) Relation(relation string) *AssetlinksCheckCall {
	c.urlParams_.Set("relation", relation)
	return c
}

// SourceAndroidAppCertificateSha256Fingerprint sets the optional
// parameter "source.androidApp.certificate.sha256Fingerprint": The
// uppercase SHA-265 fingerprint of the certificate.  From the PEM
//  certificate, it can be acquired like this:
//
//     $ keytool -printcert -file $CERTFILE | grep SHA256:
//     SHA256: 14:6D:E9:83:C5:73:06:50:D8:EE:B9:95:2F:34:FC:64:16:A0:83:
// \
//         42:E6:1D:BE:A8:8A:04:96:B2:3F:CF:44:E5
//
// or like this:
//
//     $ openssl x509 -in $CERTFILE -noout -fingerprint -sha256
//     SHA256
// Fingerprint=14:6D:E9:83:C5:73:06:50:D8:EE:B9:95:2F:34:FC:64: \
//         16:A0:83:42:E6:1D:BE:A8:8A:04:96:B2:3F:CF:44:E5
//
// In this example, the contents of this field would be
// `14:6D:E9:83:C5:73:
// 06:50:D8:EE:B9:95:2F:34:FC:64:16:A0:83:42:E6:1D:BE
// :A8:8A:04:96:B2:3F:CF:
// 44:E5`.
//
// If these tools are not available to you, you can convert the
// PEM
// certificate into the DER format, compute the SHA-256 hash of that
// string
// and represent the result as a hexstring (that is, uppercase
// hexadecimal
// representations of each octet, separated by colons).
func (c *AssetlinksCheckCall) SourceAndroidAppCertificateSha256Fingerprint(sourceAndroidAppCertificateSha256Fingerprint string) *AssetlinksCheckCall {
	c.urlParams_.Set("source.androidApp.certificate.sha256Fingerprint", sourceAndroidAppCertificateSha256Fingerprint)
	return c
}

// SourceAndroidAppPackageName sets the optional parameter
// "source.androidApp.packageName": Android App assets are naturally
// identified by their Java package name.
// For example, the Google Maps app uses the package
// name
// `com.google.android.apps.maps`.
// REQUIRED
func (c *AssetlinksCheckCall) SourceAndroidAppPackageName(sourceAndroidAppPackageName string) *AssetlinksCheckCall {
	c.urlParams_.Set("source.androidApp.packageName", sourceAndroidAppPackageName)
	return c
}

// SourceWebSite sets the optional parameter "source.web.site": Web
// assets are identified by a URL that contains only the scheme,
// hostname
// and port parts.  The format is
//
//     http[s]://<hostname>[:<port>]
//
// Hostnames must be fully qualified: they must end in a single
// period
// (".").
//
// Only the schemes "http" and "https" are currently allowed.
//
// Port numbers are given as a decimal number, and they must be omitted
// if the
// standard port numbers are used: 80 for http and 443 for https.
//
// We call this limited URL the "site".  All URLs that share the same
// scheme,
// hostname and port are considered to be a part of the site and thus
// belong
// to the web asset.
//
// Example: the asset with the site `https://www.google.com` contains
// all
// these URLs:
//
//   *   `https://www.google.com/`
//   *   `https://www.google.com:443/`
//   *   `https://www.google.com/foo`
//   *   `https://www.google.com/foo?bar`
//   *   `https://www.google.com/foo#bar`
//   *   `https://user@password:www.google.com/`
//
// But it does not contain these URLs:
//
//   *   `http://www.google.com/`       (wrong scheme)
//   *   `https://google.com/`          (hostname does not match)
//   *   `https://www.google.com:444/`  (port does not match)
// REQUIRED
func (c *AssetlinksCheckCall) SourceWebSite(sourceWebSite string) *AssetlinksCheckCall {
	c.urlParams_.Set("source.web.site", sourceWebSite)
	return c
}

// TargetAndroidAppCertificateSha256Fingerprint sets the optional
// parameter "target.androidApp.certificate.sha256Fingerprint": The
// uppercase SHA-265 fingerprint of the certificate.  From the PEM
//  certificate, it can be acquired like this:
//
//     $ keytool -printcert -file $CERTFILE | grep SHA256:
//     SHA256: 14:6D:E9:83:C5:73:06:50:D8:EE:B9:95:2F:34:FC:64:16:A0:83:
// \
//         42:E6:1D:BE:A8:8A:04:96:B2:3F:CF:44:E5
//
// or like this:
//
//     $ openssl x509 -in $CERTFILE -noout -fingerprint -sha256
//     SHA256
// Fingerprint=14:6D:E9:83:C5:73:06:50:D8:EE:B9:95:2F:34:FC:64: \
//         16:A0:83:42:E6:1D:BE:A8:8A:04:96:B2:3F:CF:44:E5
//
// In this example, the contents of this field would be
// `14:6D:E9:83:C5:73:
// 06:50:D8:EE:B9:95:2F:34:FC:64:16:A0:83:42:E6:1D:BE
// :A8:8A:04:96:B2:3F:CF:
// 44:E5`.
//
// If these tools are not available to you, you can convert the
// PEM
// certificate into the DER format, compute the SHA-256 hash of that
// string
// and represent the result as a hexstring (that is, uppercase
// hexadecimal
// representations of each octet, separated by colons).
func (c *AssetlinksCheckCall) TargetAndroidAppCertificateSha256Fingerprint(targetAndroidAppCertificateSha256Fingerprint string) *AssetlinksCheckCall {
	c.urlParams_.Set("target.androidApp.certificate.sha256Fingerprint", targetAndroidAppCertificateSha256Fingerprint)
	return c
}

// TargetAndroidAppPackageName sets the optional parameter
// "target.androidApp.packageName": Android App assets are naturally
// identified by their Java package name.
// For example, the Google Maps app uses the package
// name
// `com.google.android.apps.maps`.
// REQUIRED
func (c *AssetlinksCheckCall) TargetAndroidAppPackageName(targetAndroidAppPackageName string) *AssetlinksCheckCall {
	c.urlParams_.Set("target.androidApp.packageName", targetAndroidAppPackageName)
	return c
}

// TargetWebSite sets the optional parameter "target.web.site": Web
// assets are identified by a URL that contains only the scheme,
// hostname
// and port parts.  The format is
//
//     http[s]://<hostname>[:<port>]
//
// Hostnames must be fully qualified: they must end in a single
// period
// (".").
//
// Only the schemes "http" and "https" are currently allowed.
//
// Port numbers are given as a decimal number, and they must be omitted
// if the
// standard port numbers are used: 80 for http and 443 for https.
//
// We call this limited URL the "site".  All URLs that share the same
// scheme,
// hostname and port are considered to be a part of the site and thus
// belong
// to the web asset.
//
// Example: the asset with the site `https://www.google.com` contains
// all
// these URLs:
//
//   *   `https://www.google.com/`
//   *   `https://www.google.com:443/`
//   *   `https://www.google.com/foo`
//   *   `https://www.google.com/foo?bar`
//   *   `https://www.google.com/foo#bar`
//   *   `https://user@password:www.google.com/`
//
// But it does not contain these URLs:
//
//   *   `http://www.google.com/`       (wrong scheme)
//   *   `https://google.com/`          (hostname does not match)
//   *   `https://www.google.com:444/`  (port does not match)
// REQUIRED
func (c *AssetlinksCheckCall) TargetWebSite(targetWebSite string) *AssetlinksCheckCall {
	c.urlParams_.Set("target.web.site", targetWebSite)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AssetlinksCheckCall) Fields(s ...googleapi.Field) *AssetlinksCheckCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AssetlinksCheckCall) IfNoneMatch(entityTag string) *AssetlinksCheckCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AssetlinksCheckCall) Context(ctx context.Context) *AssetlinksCheckCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AssetlinksCheckCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AssetlinksCheckCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/assetlinks:check")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "digitalassetlinks.assetlinks.check" call.
// Exactly one of *CheckResponse or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *CheckResponse.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AssetlinksCheckCall) Do(opts ...googleapi.CallOption) (*CheckResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &CheckResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Determines whether the specified (directional) relationship exists between\nthe specified source and target assets.\n\nThe relation describes the intent of the link between the two assets as\nclaimed by the source asset.  An example for such relationships is the\ndelegation of privileges or permissions.\n\nThis command is most often used by infrastructure systems to check\npreconditions for an action.  For example, a client may want to know if it\nis OK to send a web URL to a particular mobile app instead. The client can\ncheck for the relevant asset link from the website to the mobile app to\ndecide if the operation should be allowed.\n\nA note about security: if you specify a secure asset as the source, such as\nan HTTPS website or an Android app, the API will ensure that any\nstatements used to generate the response have been made in a secure way by\nthe owner of that asset.  Conversely, if the source asset is an insecure\nHTTP website (that is, the URL starts with `http://` instead of `https://`),\nthe API cannot verify its statements securely, and it is not possible to\nensure that the website's statements have not been altered by a third\nparty.  For more information, see the [Digital Asset Links technical design\nspecification](https://github.com/google/digitalassetlinks/blob/master/well-known/details.md).",
	//   "flatPath": "v1/assetlinks:check",
	//   "httpMethod": "GET",
	//   "id": "digitalassetlinks.assetlinks.check",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "relation": {
	//       "description": "Query string for the relation.\n\nWe identify relations with strings of the format `\u003ckind\u003e/\u003cdetail\u003e`, where\n`\u003ckind\u003e` must be one of a set of pre-defined purpose categories, and\n`\u003cdetail\u003e` is a free-form lowercase alphanumeric string that describes the\nspecific use case of the statement.\n\nRefer to [our API documentation](/digital-asset-links/v1/relation-strings)\nfor the current list of supported relations.\n\nFor a query to match an asset link, both the query's and the asset link's\nrelation strings must match exactly.\n\nExample: A query with relation `delegate_permission/common.handle_all_urls`\nmatches an asset link with relation\n`delegate_permission/common.handle_all_urls`.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "source.androidApp.certificate.sha256Fingerprint": {
	//       "description": "The uppercase SHA-265 fingerprint of the certificate.  From the PEM\n certificate, it can be acquired like this:\n\n    $ keytool -printcert -file $CERTFILE | grep SHA256:\n    SHA256: 14:6D:E9:83:C5:73:06:50:D8:EE:B9:95:2F:34:FC:64:16:A0:83: \\\n        42:E6:1D:BE:A8:8A:04:96:B2:3F:CF:44:E5\n\nor like this:\n\n    $ openssl x509 -in $CERTFILE -noout -fingerprint -sha256\n    SHA256 Fingerprint=14:6D:E9:83:C5:73:06:50:D8:EE:B9:95:2F:34:FC:64: \\\n        16:A0:83:42:E6:1D:BE:A8:8A:04:96:B2:3F:CF:44:E5\n\nIn this example, the contents of this field would be `14:6D:E9:83:C5:73:\n06:50:D8:EE:B9:95:2F:34:FC:64:16:A0:83:42:E6:1D:BE:A8:8A:04:96:B2:3F:CF:\n44:E5`.\n\nIf these tools are not available to you, you can convert the PEM\ncertificate into the DER format, compute the SHA-256 hash of that string\nand represent the result as a hexstring (that is, uppercase hexadecimal\nrepresentations of each octet, separated by colons).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "source.androidApp.packageName": {
	//       "description": "Android App assets are naturally identified by their Java package name.\nFor example, the Google Maps app uses the package name\n`com.google.android.apps.maps`.\nREQUIRED",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "source.web.site": {
	//       "description": "Web assets are identified by a URL that contains only the scheme, hostname\nand port parts.  The format is\n\n    http[s]://\u003chostname\u003e[:\u003cport\u003e]\n\nHostnames must be fully qualified: they must end in a single period\n(\"`.`\").\n\nOnly the schemes \"http\" and \"https\" are currently allowed.\n\nPort numbers are given as a decimal number, and they must be omitted if the\nstandard port numbers are used: 80 for http and 443 for https.\n\nWe call this limited URL the \"site\".  All URLs that share the same scheme,\nhostname and port are considered to be a part of the site and thus belong\nto the web asset.\n\nExample: the asset with the site `https://www.google.com` contains all\nthese URLs:\n\n  *   `https://www.google.com/`\n  *   `https://www.google.com:443/`\n  *   `https://www.google.com/foo`\n  *   `https://www.google.com/foo?bar`\n  *   `https://www.google.com/foo#bar`\n  *   `https://user@password:www.google.com/`\n\nBut it does not contain these URLs:\n\n  *   `http://www.google.com/`       (wrong scheme)\n  *   `https://google.com/`          (hostname does not match)\n  *   `https://www.google.com:444/`  (port does not match)\nREQUIRED",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "target.androidApp.certificate.sha256Fingerprint": {
	//       "description": "The uppercase SHA-265 fingerprint of the certificate.  From the PEM\n certificate, it can be acquired like this:\n\n    $ keytool -printcert -file $CERTFILE | grep SHA256:\n    SHA256: 14:6D:E9:83:C5:73:06:50:D8:EE:B9:95:2F:34:FC:64:16:A0:83: \\\n        42:E6:1D:BE:A8:8A:04:96:B2:3F:CF:44:E5\n\nor like this:\n\n    $ openssl x509 -in $CERTFILE -noout -fingerprint -sha256\n    SHA256 Fingerprint=14:6D:E9:83:C5:73:06:50:D8:EE:B9:95:2F:34:FC:64: \\\n        16:A0:83:42:E6:1D:BE:A8:8A:04:96:B2:3F:CF:44:E5\n\nIn this example, the contents of this field would be `14:6D:E9:83:C5:73:\n06:50:D8:EE:B9:95:2F:34:FC:64:16:A0:83:42:E6:1D:BE:A8:8A:04:96:B2:3F:CF:\n44:E5`.\n\nIf these tools are not available to you, you can convert the PEM\ncertificate into the DER format, compute the SHA-256 hash of that string\nand represent the result as a hexstring (that is, uppercase hexadecimal\nrepresentations of each octet, separated by colons).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "target.androidApp.packageName": {
	//       "description": "Android App assets are naturally identified by their Java package name.\nFor example, the Google Maps app uses the package name\n`com.google.android.apps.maps`.\nREQUIRED",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "target.web.site": {
	//       "description": "Web assets are identified by a URL that contains only the scheme, hostname\nand port parts.  The format is\n\n    http[s]://\u003chostname\u003e[:\u003cport\u003e]\n\nHostnames must be fully qualified: they must end in a single period\n(\"`.`\").\n\nOnly the schemes \"http\" and \"https\" are currently allowed.\n\nPort numbers are given as a decimal number, and they must be omitted if the\nstandard port numbers are used: 80 for http and 443 for https.\n\nWe call this limited URL the \"site\".  All URLs that share the same scheme,\nhostname and port are considered to be a part of the site and thus belong\nto the web asset.\n\nExample: the asset with the site `https://www.google.com` contains all\nthese URLs:\n\n  *   `https://www.google.com/`\n  *   `https://www.google.com:443/`\n  *   `https://www.google.com/foo`\n  *   `https://www.google.com/foo?bar`\n  *   `https://www.google.com/foo#bar`\n  *   `https://user@password:www.google.com/`\n\nBut it does not contain these URLs:\n\n  *   `http://www.google.com/`       (wrong scheme)\n  *   `https://google.com/`          (hostname does not match)\n  *   `https://www.google.com:444/`  (port does not match)\nREQUIRED",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/assetlinks:check",
	//   "response": {
	//     "$ref": "CheckResponse"
	//   }
	// }

}

// method id "digitalassetlinks.statements.list":

type StatementsListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Retrieves a list of all statements from a given source that
// match the
// specified target and statement string.
//
// The API guarantees that all statements with secure source assets,
// such as
// HTTPS websites or Android apps, have been made in a secure way by the
// owner
// of those assets, as described in the [Digital Asset Links technical
// design
// specification](https://github.com/google/digitalassetlinks/blob
// /master/well-known/details.md).
// Specifically, you should consider that for insecure websites (that
// is,
// where the URL starts with `http://` instead of `https://`), this
// guarantee
// cannot be made.
//
// The `List` command is most useful in cases where the API client wants
// to
// know all the ways in which two assets are related, or enumerate all
// the
// relationships from a particular source asset.  Example: a feature
// that
// helps users navigate to related items.  When a mobile app is running
// on a
// device, the feature would make it easy to navigate to the
// corresponding web
// site or Google+ profile.
func (r *StatementsService) List() *StatementsListCall {
	c := &StatementsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// Relation sets the optional parameter "relation": Use only
// associations that match the specified relation.
//
// See the [`Statement`](#Statement) message for a detailed definition
// of
// relation strings.
//
// For a query to match a statement, one of the following must be
// true:
//
// *    both the query's and the statement's relation strings match
// exactly,
//      or
// *    the query's relation string is empty or missing.
//
// Example: A query with relation
// `delegate_permission/common.handle_all_urls`
// matches an asset link with
// relation
// `delegate_permission/common.handle_all_urls`.
func (c *StatementsListCall) Relation(relation string) *StatementsListCall {
	c.urlParams_.Set("relation", relation)
	return c
}

// SourceAndroidAppCertificateSha256Fingerprint sets the optional
// parameter "source.androidApp.certificate.sha256Fingerprint": The
// uppercase SHA-265 fingerprint of the certificate.  From the PEM
//  certificate, it can be acquired like this:
//
//     $ keytool -printcert -file $CERTFILE | grep SHA256:
//     SHA256: 14:6D:E9:83:C5:73:06:50:D8:EE:B9:95:2F:34:FC:64:16:A0:83:
// \
//         42:E6:1D:BE:A8:8A:04:96:B2:3F:CF:44:E5
//
// or like this:
//
//     $ openssl x509 -in $CERTFILE -noout -fingerprint -sha256
//     SHA256
// Fingerprint=14:6D:E9:83:C5:73:06:50:D8:EE:B9:95:2F:34:FC:64: \
//         16:A0:83:42:E6:1D:BE:A8:8A:04:96:B2:3F:CF:44:E5
//
// In this example, the contents of this field would be
// `14:6D:E9:83:C5:73:
// 06:50:D8:EE:B9:95:2F:34:FC:64:16:A0:83:42:E6:1D:BE
// :A8:8A:04:96:B2:3F:CF:
// 44:E5`.
//
// If these tools are not available to you, you can convert the
// PEM
// certificate into the DER format, compute the SHA-256 hash of that
// string
// and represent the result as a hexstring (that is, uppercase
// hexadecimal
// representations of each octet, separated by colons).
func (c *StatementsListCall) SourceAndroidAppCertificateSha256Fingerprint(sourceAndroidAppCertificateSha256Fingerprint string) *StatementsListCall {
	c.urlParams_.Set("source.androidApp.certificate.sha256Fingerprint", sourceAndroidAppCertificateSha256Fingerprint)
	return c
}

// SourceAndroidAppPackageName sets the optional parameter
// "source.androidApp.packageName": Android App assets are naturally
// identified by their Java package name.
// For example, the Google Maps app uses the package
// name
// `com.google.android.apps.maps`.
// REQUIRED
func (c *StatementsListCall) SourceAndroidAppPackageName(sourceAndroidAppPackageName string) *StatementsListCall {
	c.urlParams_.Set("source.androidApp.packageName", sourceAndroidAppPackageName)
	return c
}

// SourceWebSite sets the optional parameter "source.web.site": Web
// assets are identified by a URL that contains only the scheme,
// hostname
// and port parts.  The format is
//
//     http[s]://<hostname>[:<port>]
//
// Hostnames must be fully qualified: they must end in a single
// period
// (".").
//
// Only the schemes "http" and "https" are currently allowed.
//
// Port numbers are given as a decimal number, and they must be omitted
// if the
// standard port numbers are used: 80 for http and 443 for https.
//
// We call this limited URL the "site".  All URLs that share the same
// scheme,
// hostname and port are considered to be a part of the site and thus
// belong
// to the web asset.
//
// Example: the asset with the site `https://www.google.com` contains
// all
// these URLs:
//
//   *   `https://www.google.com/`
//   *   `https://www.google.com:443/`
//   *   `https://www.google.com/foo`
//   *   `https://www.google.com/foo?bar`
//   *   `https://www.google.com/foo#bar`
//   *   `https://user@password:www.google.com/`
//
// But it does not contain these URLs:
//
//   *   `http://www.google.com/`       (wrong scheme)
//   *   `https://google.com/`          (hostname does not match)
//   *   `https://www.google.com:444/`  (port does not match)
// REQUIRED
func (c *StatementsListCall) SourceWebSite(sourceWebSite string) *StatementsListCall {
	c.urlParams_.Set("source.web.site", sourceWebSite)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *StatementsListCall) Fields(s ...googleapi.Field) *StatementsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *StatementsListCall) IfNoneMatch(entityTag string) *StatementsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *StatementsListCall) Context(ctx context.Context) *StatementsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *StatementsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *StatementsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/statements:list")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "digitalassetlinks.statements.list" call.
// Exactly one of *ListResponse or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *ListResponse.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *StatementsListCall) Do(opts ...googleapi.CallOption) (*ListResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of all statements from a given source that match the\nspecified target and statement string.\n\nThe API guarantees that all statements with secure source assets, such as\nHTTPS websites or Android apps, have been made in a secure way by the owner\nof those assets, as described in the [Digital Asset Links technical design\nspecification](https://github.com/google/digitalassetlinks/blob/master/well-known/details.md).\nSpecifically, you should consider that for insecure websites (that is,\nwhere the URL starts with `http://` instead of `https://`), this guarantee\ncannot be made.\n\nThe `List` command is most useful in cases where the API client wants to\nknow all the ways in which two assets are related, or enumerate all the\nrelationships from a particular source asset.  Example: a feature that\nhelps users navigate to related items.  When a mobile app is running on a\ndevice, the feature would make it easy to navigate to the corresponding web\nsite or Google+ profile.",
	//   "flatPath": "v1/statements:list",
	//   "httpMethod": "GET",
	//   "id": "digitalassetlinks.statements.list",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "relation": {
	//       "description": "Use only associations that match the specified relation.\n\nSee the [`Statement`](#Statement) message for a detailed definition of\nrelation strings.\n\nFor a query to match a statement, one of the following must be true:\n\n*    both the query's and the statement's relation strings match exactly,\n     or\n*    the query's relation string is empty or missing.\n\nExample: A query with relation `delegate_permission/common.handle_all_urls`\nmatches an asset link with relation\n`delegate_permission/common.handle_all_urls`.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "source.androidApp.certificate.sha256Fingerprint": {
	//       "description": "The uppercase SHA-265 fingerprint of the certificate.  From the PEM\n certificate, it can be acquired like this:\n\n    $ keytool -printcert -file $CERTFILE | grep SHA256:\n    SHA256: 14:6D:E9:83:C5:73:06:50:D8:EE:B9:95:2F:34:FC:64:16:A0:83: \\\n        42:E6:1D:BE:A8:8A:04:96:B2:3F:CF:44:E5\n\nor like this:\n\n    $ openssl x509 -in $CERTFILE -noout -fingerprint -sha256\n    SHA256 Fingerprint=14:6D:E9:83:C5:73:06:50:D8:EE:B9:95:2F:34:FC:64: \\\n        16:A0:83:42:E6:1D:BE:A8:8A:04:96:B2:3F:CF:44:E5\n\nIn this example, the contents of this field would be `14:6D:E9:83:C5:73:\n06:50:D8:EE:B9:95:2F:34:FC:64:16:A0:83:42:E6:1D:BE:A8:8A:04:96:B2:3F:CF:\n44:E5`.\n\nIf these tools are not available to you, you can convert the PEM\ncertificate into the DER format, compute the SHA-256 hash of that string\nand represent the result as a hexstring (that is, uppercase hexadecimal\nrepresentations of each octet, separated by colons).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "source.androidApp.packageName": {
	//       "description": "Android App assets are naturally identified by their Java package name.\nFor example, the Google Maps app uses the package name\n`com.google.android.apps.maps`.\nREQUIRED",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "source.web.site": {
	//       "description": "Web assets are identified by a URL that contains only the scheme, hostname\nand port parts.  The format is\n\n    http[s]://\u003chostname\u003e[:\u003cport\u003e]\n\nHostnames must be fully qualified: they must end in a single period\n(\"`.`\").\n\nOnly the schemes \"http\" and \"https\" are currently allowed.\n\nPort numbers are given as a decimal number, and they must be omitted if the\nstandard port numbers are used: 80 for http and 443 for https.\n\nWe call this limited URL the \"site\".  All URLs that share the same scheme,\nhostname and port are considered to be a part of the site and thus belong\nto the web asset.\n\nExample: the asset with the site `https://www.google.com` contains all\nthese URLs:\n\n  *   `https://www.google.com/`\n  *   `https://www.google.com:443/`\n  *   `https://www.google.com/foo`\n  *   `https://www.google.com/foo?bar`\n  *   `https://www.google.com/foo#bar`\n  *   `https://user@password:www.google.com/`\n\nBut it does not contain these URLs:\n\n  *   `http://www.google.com/`       (wrong scheme)\n  *   `https://google.com/`          (hostname does not match)\n  *   `https://www.google.com:444/`  (port does not match)\nREQUIRED",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/statements:list",
	//   "response": {
	//     "$ref": "ListResponse"
	//   }
	// }

}
