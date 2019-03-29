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

// Package cloudscheduler provides access to the Cloud Scheduler API.
//
// See https://cloud.google.com/scheduler/
//
// Usage example:
//
//   import "google.golang.org/api/cloudscheduler/v1beta1"
//   ...
//   cloudschedulerService, err := cloudscheduler.New(oauthHttpClient)
package cloudscheduler // import "google.golang.org/api/cloudscheduler/v1beta1"

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

const apiId = "cloudscheduler:v1beta1"
const apiName = "cloudscheduler"
const apiVersion = "v1beta1"
const basePath = "https://cloudscheduler.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Projects = NewProjectsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Projects *ProjectsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Locations = NewProjectsLocationsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Locations *ProjectsLocationsService
}

func NewProjectsLocationsService(s *Service) *ProjectsLocationsService {
	rs := &ProjectsLocationsService{s: s}
	rs.Jobs = NewProjectsLocationsJobsService(s)
	return rs
}

type ProjectsLocationsService struct {
	s *Service

	Jobs *ProjectsLocationsJobsService
}

func NewProjectsLocationsJobsService(s *Service) *ProjectsLocationsJobsService {
	rs := &ProjectsLocationsJobsService{s: s}
	return rs
}

type ProjectsLocationsJobsService struct {
	s *Service
}

// AppEngineHttpTarget: App Engine target. The job will be pushed to a
// job handler by means
// of an HTTP request via an http_method such
// as HTTP POST, HTTP GET, etc. The job is acknowledged by means of
// an
// HTTP response code in the range [200 - 299]. Error 503 is
// considered an App Engine system error instead of an
// application
// error. Requests returning error 503 will be retried regardless
// of
// retry configuration and not counted against retry counts. Any
// other
// response code, or a failure to receive a response before
// the
// deadline, constitutes a failed attempt.
type AppEngineHttpTarget struct {
	// AppEngineRouting: App Engine Routing setting for the job.
	AppEngineRouting *AppEngineRouting `json:"appEngineRouting,omitempty"`

	// Body: Body.
	//
	// HTTP request body. A request body is allowed only if the HTTP method
	// is
	// POST or PUT. It will result in invalid argument error to set a body
	// on a
	// job with an incompatible HttpMethod.
	Body string `json:"body,omitempty"`

	// Headers: HTTP request headers.
	//
	// This map contains the header field names and values. Headers can be
	// set
	// when the job is created.
	//
	// Cloud Scheduler sets some headers to default values:
	//
	// * `User-Agent`: By default, this header is
	//   "AppEngine-Google; (+http://code.google.com/appengine)".
	//   This header can be modified, but Cloud Scheduler will append
	//   "AppEngine-Google; (+http://code.google.com/appengine)" to the
	//   modified `User-Agent`.
	//
	// If the job has an body, Cloud Scheduler sets the
	// following headers:
	//
	// * `Content-Type`: By default, the `Content-Type` header is set to
	//   "application/octet-stream". The default can be overridden by
	// explictly
	//   setting `Content-Type` to a particular media type when the job is
	//   created.
	//   For example, `Content-Type` can be set to "application/json".
	// * `Content-Length`: This is computed by Cloud Scheduler. This value
	// is
	//   output only. It cannot be changed.
	//
	// The headers below are output only. They cannot be set or
	// overridden:
	//
	// * `X-Google-*`: For Google internal use only.
	// * `X-AppEngine-*`: For Google internal use only. See
	//   [Reading request
	// headers](https://cloud.google.com/appengine/docs/python/taskqueue/push
	// /creating-handlers#reading_request_headers).
	//
	// In addition, some App Engine headers, which contain
	// job-specific information, are also be sent to the job handler;
	// see
	// [request
	// headers](https://cloud.google.comappengine/docs/standard/python/config
	// /cron#securing_urls_for_cron).
	Headers map[string]string `json:"headers,omitempty"`

	// HttpMethod: The HTTP method to use for the request. PATCH and OPTIONS
	// are not
	// permitted.
	//
	// Possible values:
	//   "HTTP_METHOD_UNSPECIFIED" - HTTP method unspecified. Defaults to
	// POST.
	//   "POST" - HTTP POST
	//   "GET" - HTTP GET
	//   "HEAD" - HTTP HEAD
	//   "PUT" - HTTP PUT
	//   "DELETE" - HTTP DELETE
	//   "PATCH" - HTTP PATCH
	//   "OPTIONS" - HTTP OPTIONS
	HttpMethod string `json:"httpMethod,omitempty"`

	// RelativeUri: The relative URI.
	//
	// The relative URL must begin with "/" and must be a valid HTTP
	// relative URL.
	// It can contain a path, query string arguments, and `#` fragments.
	// If the relative URL is empty, then the root path "/" will be used.
	// No spaces are allowed, and the maximum length allowed is 2083
	// characters.
	RelativeUri string `json:"relativeUri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AppEngineRouting") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AppEngineRouting") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *AppEngineHttpTarget) MarshalJSON() ([]byte, error) {
	type NoMethod AppEngineHttpTarget
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AppEngineRouting: App Engine Routing.
//
// For more information about services, versions, and instances see
// [An Overview of App
// Engine](https://cloud.google.com/appengine/docs/python/an-overview-of-
// app-engine),
// [Microservices Architecture on Google App
// Engine](https://cloud.google.com/appengine/docs/python/microservices-o
// n-app-engine),
// [App Engine Standard request
// routing](https://cloud.google.com/appengine/docs/standard/python/how-r
// equests-are-routed), and
// [App Engine Flex request
// routing](https://cloud.google.com/appengine/docs/flexible/python/how-r
// equests-are-routed).
type AppEngineRouting struct {
	// Host: Output only. The host that the job is sent to.
	//
	// For more information about how App Engine requests are routed,
	// see
	// [here](https://cloud.google.com/appengine/docs/standard/python/how
	// -requests-are-routed).
	//
	// The host is constructed as:
	//
	//
	// * `host = [application_domain_name]`</br>
	//   `| [service] + '.' + [application_domain_name]`</br>
	//   `| [version] + '.' + [application_domain_name]`</br>
	//   `| [version_dot_service]+ '.' + [application_domain_name]`</br>
	//   `| [instance] + '.' + [application_domain_name]`</br>
	//   `| [instance_dot_service] + '.' + [application_domain_name]`</br>
	//   `| [instance_dot_version] + '.' + [application_domain_name]`</br>
	//   `| [instance_dot_version_dot_service] + '.' +
	// [application_domain_name]`
	//
	// * `application_domain_name` = The domain name of the app, for
	//   example <app-id>.appspot.com, which is associated with the
	//   job's project ID.
	//
	// * `service =` service
	//
	// * `version =` version
	//
	// * `version_dot_service =`
	//   version `+ '.' +`
	//   service
	//
	// * `instance =` instance
	//
	// * `instance_dot_service =`
	//   instance `+ '.' +`
	//   service
	//
	// * `instance_dot_version =`
	//   instance `+ '.' +`
	//   version
	//
	// * `instance_dot_version_dot_service =`
	//   instance `+ '.' +`
	//   version `+ '.' +`
	//   service
	//
	//
	// If service is empty, then the job will be sent
	// to the service which is the default service when the job is
	// attempted.
	//
	// If version is empty, then the job will be sent
	// to the version which is the default version when the job is
	// attempted.
	//
	// If instance is empty, then the job will be
	// sent to an instance which is available when the job is attempted.
	//
	// If service,
	// version, or
	// instance is invalid, then the job will be sent
	// to the default version of the default service when the job is
	// attempted.
	Host string `json:"host,omitempty"`

	// Instance: App instance.
	//
	// By default, the job is sent to an instance which is available
	// when
	// the job is attempted.
	//
	// Requests can only be sent to a specific instance if
	// [manual scaling is used in App Engine
	// Standard](https://cloud.google.com/appengine/docs/python/an-overview-o
	// f-app-engine?hl=en_US#scaling_types_and_instance_classes).
	// App Engine Flex does not support instances. For more information,
	// see
	// [App Engine Standard request
	// routing](https://cloud.google.com/appengine/docs/standard/python/how-r
	// equests-are-routed) and
	// [App Engine Flex request
	// routing](https://cloud.google.com/appengine/docs/flexible/python/how-r
	// equests-are-routed).
	Instance string `json:"instance,omitempty"`

	// Service: App service.
	//
	// By default, the job is sent to the service which is the
	// default
	// service when the job is attempted.
	Service string `json:"service,omitempty"`

	// Version: App version.
	//
	// By default, the job is sent to the version which is the
	// default
	// version when the job is attempted.
	Version string `json:"version,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Host") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Host") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AppEngineRouting) MarshalJSON() ([]byte, error) {
	type NoMethod AppEngineRouting
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated
// empty messages in your APIs. A typical example is to use it as the
// request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// HttpTarget: Http target. The job will be pushed to the job handler by
// means of
// an HTTP request via an http_method such as HTTP
// POST, HTTP GET, etc. The job is acknowledged by means of an
// HTTP
// response code in the range [200 - 299]. A failure to receive a
// response
// constitutes a failed execution. For a redirected request, the
// response
// returned by the redirected request is considered.
type HttpTarget struct {
	// Body: HTTP request body. A request body is allowed only if the
	// HTTP
	// method is POST, PUT, or PATCH. It is an error to set body on a job
	// with an
	// incompatible HttpMethod.
	Body string `json:"body,omitempty"`

	// Headers: The user can specify HTTP request headers to send with the
	// job's
	// HTTP request. This map contains the header field names and
	// values. Repeated headers are not supported, but a header value
	// can
	// contain commas. These headers represent a subset of the headers
	// that will accompany the job's HTTP request. Some HTTP request
	// headers will be ignored or replaced. A partial list of headers
	// that
	// will be ignored or replaced is below:
	// - Host: This will be computed by Cloud Scheduler and derived
	// from
	// uri.
	// * `Content-Length`: This will be computed by Cloud Scheduler.
	// * `User-Agent`: This will be set to "Google-Cloud-Scheduler".
	// * `X-Google-*`: Google internal use only.
	// * `X-AppEngine-*`: Google internal use only.
	//
	// The total size of headers must be less than 80KB.
	Headers map[string]string `json:"headers,omitempty"`

	// HttpMethod: Which HTTP method to use for the request.
	//
	// Possible values:
	//   "HTTP_METHOD_UNSPECIFIED" - HTTP method unspecified. Defaults to
	// POST.
	//   "POST" - HTTP POST
	//   "GET" - HTTP GET
	//   "HEAD" - HTTP HEAD
	//   "PUT" - HTTP PUT
	//   "DELETE" - HTTP DELETE
	//   "PATCH" - HTTP PATCH
	//   "OPTIONS" - HTTP OPTIONS
	HttpMethod string `json:"httpMethod,omitempty"`

	// Uri: Required.
	//
	// The full URI path that the request will be sent to. This string
	// must begin with either "http://" or "https://". Some examples
	// of
	// valid values for uri are:
	// `http://acme.com` and `https://acme.com/sales:8080`. Cloud Scheduler
	// will
	// encode some characters for safety and compatibility. The maximum
	// allowed
	// URL length is 2083 characters after encoding.
	Uri string `json:"uri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Body") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Body") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *HttpTarget) MarshalJSON() ([]byte, error) {
	type NoMethod HttpTarget
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Job: Configuration for a job.
// The maximum allowed size for a job is 100KB.
type Job struct {
	// AppEngineHttpTarget: App Engine HTTP target.
	AppEngineHttpTarget *AppEngineHttpTarget `json:"appEngineHttpTarget,omitempty"`

	// Description: A human-readable description for the job. This string
	// must not contain
	// more than 500 characters.
	Description string `json:"description,omitempty"`

	// HttpTarget: HTTP target.
	HttpTarget *HttpTarget `json:"httpTarget,omitempty"`

	// LastAttemptTime: Output only. The time the last job attempt started.
	LastAttemptTime string `json:"lastAttemptTime,omitempty"`

	// Name: The job name. For
	// example:
	// `projects/PROJECT_ID/locations/LOCATION_ID/jobs/JOB_ID`.
	//
	// * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]),
	//    hyphens (-), colons (:), or periods (.).
	//    For more information, see
	//    [Identifying
	// projects](https://cloud.google.com/resource-manager/docs/creating-mana
	// ging-projects#identifying_projects)
	// * `LOCATION_ID` is the canonical ID for the job's location.
	//    The list of available locations can be obtained by calling
	//    ListLocations.
	//    For more information, see
	// https://cloud.google.com/about/locations/.
	// * `JOB_ID` can contain only letters ([A-Za-z]), numbers ([0-9]),
	//    hyphens (-), or underscores (_). The maximum length is 500
	// characters.
	Name string `json:"name,omitempty"`

	// PubsubTarget: Pub/Sub target.
	PubsubTarget *PubsubTarget `json:"pubsubTarget,omitempty"`

	// RetryConfig: Settings that determine the retry behavior.
	RetryConfig *RetryConfig `json:"retryConfig,omitempty"`

	// Schedule: Required.
	//
	// Describes the schedule on which the job will be executed.
	//
	// As a general rule, execution `n + 1` of a job will not begin
	// until execution `n` has finished. Cloud Scheduler will never
	// allow two simultaneously outstanding executions. For example,
	// this implies that if the `n+1`th execution is scheduled to run
	// at
	// 16:00 but the `n`th execution takes until 16:15, the
	// `n+1`th
	// execution will not start until `16:15`.
	// A scheduled start time will be delayed if the previous
	// execution has not ended when its scheduled time occurs.
	//
	// If retry_count > 0 and a job attempt fails,
	// the job will be tried a total of retry_count
	// times, with exponential backoff, until the next scheduled
	// start
	// time.
	//
	// The schedule can be either of the following types:
	//
	// * [Crontab](http://en.wikipedia.org/wiki/Cron#Overview)
	// * English-like
	// [schedule](https://cloud.google.com/scheduler/docs/configuring/cron-jo
	// b-schedules)
	Schedule string `json:"schedule,omitempty"`

	// ScheduleTime: Output only. The next time the job is scheduled. Note
	// that this may be a
	// retry of a previously failed attempt or the next execution
	// time
	// according to the schedule.
	ScheduleTime string `json:"scheduleTime,omitempty"`

	// State: Output only. State of the job.
	//
	// Possible values:
	//   "STATE_UNSPECIFIED" - Unspecified state.
	//   "ENABLED" - The job is executing normally.
	//   "PAUSED" - The job is paused by the user. It will not execute. A
	// user can
	// intentionally pause the job using
	// PauseJobRequest.
	//   "DISABLED" - The job is disabled by the system due to error. The
	// user
	// cannot directly set a job to be disabled.
	//   "UPDATE_FAILED" - The job state resulting from a failed
	// CloudScheduler.UpdateJob
	// operation. To recover a job from this state,
	// retry
	// CloudScheduler.UpdateJob until a successful response is received.
	State string `json:"state,omitempty"`

	// Status: Output only. The response from the target for the last
	// attempted execution.
	Status *Status `json:"status,omitempty"`

	// TimeZone: Specifies the time zone to be used in
	// interpreting
	// schedule. The value of this field must be a time
	// zone name from the [tz
	// database](http://en.wikipedia.org/wiki/Tz_database).
	//
	// Note that some time zones include a provision for
	// daylight savings time. The rules for daylight saving time
	// are
	// determined by the chosen tz. For UTC use the string "utc". If a
	// time zone is not specified, the default will be in UTC (also known
	// as GMT).
	TimeZone string `json:"timeZone,omitempty"`

	// UserUpdateTime: Output only. The creation time of the job.
	UserUpdateTime string `json:"userUpdateTime,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AppEngineHttpTarget")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AppEngineHttpTarget") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Job) MarshalJSON() ([]byte, error) {
	type NoMethod Job
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListJobsResponse: Response message for listing jobs using ListJobs.
type ListJobsResponse struct {
	// Jobs: The list of jobs.
	Jobs []*Job `json:"jobs,omitempty"`

	// NextPageToken: A token to retrieve next page of results. Pass this
	// value in the
	// page_token field in the subsequent call to
	// ListJobs to retrieve the next page of results.
	// If this is empty it indicates that there are no more results
	// through which to paginate.
	//
	// The page token is valid for only 2 hours.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Jobs") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Jobs") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListJobsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListJobsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListLocationsResponse: The response message for
// Locations.ListLocations.
type ListLocationsResponse struct {
	// Locations: A list of locations that matches the specified filter in
	// the request.
	Locations []*Location `json:"locations,omitempty"`

	// NextPageToken: The standard List next-page token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Locations") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Locations") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListLocationsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListLocationsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Location: A resource that represents Google Cloud Platform location.
type Location struct {
	// DisplayName: The friendly name for this location, typically a nearby
	// city name.
	// For example, "Tokyo".
	DisplayName string `json:"displayName,omitempty"`

	// Labels: Cross-service attributes for the location. For example
	//
	//     {"cloud.googleapis.com/region": "us-east1"}
	Labels map[string]string `json:"labels,omitempty"`

	// LocationId: The canonical id for this location. For example:
	// "us-east1".
	LocationId string `json:"locationId,omitempty"`

	// Metadata: Service-specific metadata. For example the available
	// capacity at the given
	// location.
	Metadata googleapi.RawMessage `json:"metadata,omitempty"`

	// Name: Resource name for the location, which may vary between
	// implementations.
	// For example: "projects/example-project/locations/us-east1"
	Name string `json:"name,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "DisplayName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DisplayName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Location) MarshalJSON() ([]byte, error) {
	type NoMethod Location
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PauseJobRequest: Request message for PauseJob.
type PauseJobRequest struct {
}

// PubsubMessage: A message that is published by publishers and consumed
// by subscribers. The
// message must contain either a non-empty data field or at least one
// attribute.
// Note that client libraries represent this object
// differently
// depending on the language. See the corresponding
// <a
// href="https://cloud.google.com/pubsub/docs/reference/libraries">client
//
// library documentation</a> for more information. See
// <a href="https://cloud.google.com/pubsub/quotas">Quotas and
// limits</a>
// for more information about message limits.
type PubsubMessage struct {
	// Attributes: Optional attributes for this message.
	Attributes map[string]string `json:"attributes,omitempty"`

	// Data: The message data field. If this field is empty, the message
	// must contain
	// at least one attribute.
	Data string `json:"data,omitempty"`

	// MessageId: ID of this message, assigned by the server when the
	// message is published.
	// Guaranteed to be unique within the topic. This value may be read by
	// a
	// subscriber that receives a `PubsubMessage` via a `Pull` call or a
	// push
	// delivery. It must not be populated by the publisher in a `Publish`
	// call.
	MessageId string `json:"messageId,omitempty"`

	// PublishTime: The time at which the message was published, populated
	// by the server when
	// it receives the `Publish` call. It must not be populated by
	// the
	// publisher in a `Publish` call.
	PublishTime string `json:"publishTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Attributes") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Attributes") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PubsubMessage) MarshalJSON() ([]byte, error) {
	type NoMethod PubsubMessage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PubsubTarget: Pub/Sub target. The job will be delivered by publishing
// a message to
// the given Pub/Sub topic.
type PubsubTarget struct {
	// Attributes: Attributes for PubsubMessage.
	//
	// Pubsub message must contain either non-empty data, or at least
	// one
	// attribute.
	Attributes map[string]string `json:"attributes,omitempty"`

	// Data: The message payload for PubsubMessage.
	//
	// Pubsub message must contain either non-empty data, or at least
	// one
	// attribute.
	Data string `json:"data,omitempty"`

	// TopicName: Required.
	//
	// The name of the Cloud Pub/Sub topic to which messages will
	// be published when a job is delivered. The topic name must be in
	// the
	// same format as required by
	// PubSub's
	// [PublishRequest.name](https://cloud.google.com/pubsub/docs/re
	// ference/rpc/google.pubsub.v1#publishrequest),
	// for example `projects/PROJECT_ID/topics/TOPIC_ID`.
	//
	// The topic must be in the same project as the Cloud Scheduler job.
	TopicName string `json:"topicName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Attributes") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Attributes") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PubsubTarget) MarshalJSON() ([]byte, error) {
	type NoMethod PubsubTarget
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ResumeJobRequest: Request message for ResumeJob.
type ResumeJobRequest struct {
}

// RetryConfig: Settings that determine the retry behavior.
//
// By default, if a job does not complete successfully (meaning that
// an acknowledgement is not received from the handler, then it will be
// retried
// with exponential backoff according to the settings in RetryConfig.
type RetryConfig struct {
	// MaxBackoffDuration: The maximum amount of time to wait before
	// retrying a job after
	// it fails.
	//
	// The default value of this field is 1 hour.
	MaxBackoffDuration string `json:"maxBackoffDuration,omitempty"`

	// MaxDoublings: The time between retries will double `max_doublings`
	// times.
	//
	// A job's retry interval starts at
	// min_backoff_duration, then doubles
	// `max_doublings` times, then increases linearly, and finally
	// retries retries at intervals of
	// max_backoff_duration up to
	// retry_count times.
	//
	// For example, if min_backoff_duration is
	// 10s, max_backoff_duration is 300s, and
	// `max_doublings` is 3, then the a job will first be retried in 10s.
	// The
	// retry interval will double three times, and then increase linearly
	// by
	// 2^3 * 10s.  Finally, the job will retry at intervals
	// of
	// max_backoff_duration until the job has
	// been attempted retry_count times. Thus, the
	// requests will retry at 10s, 20s, 40s, 80s, 160s, 240s, 300s, 300s,
	// ....
	//
	// The default value of this field is 5.
	MaxDoublings int64 `json:"maxDoublings,omitempty"`

	// MaxRetryDuration: The time limit for retrying a failed job, measured
	// from time when an
	// execution was first attempted. If specified with
	// retry_count, the job will be retried until both limits
	// are
	// reached.
	//
	// The default value for max_retry_duration is zero, which means
	// retry
	// duration is unlimited.
	MaxRetryDuration string `json:"maxRetryDuration,omitempty"`

	// MinBackoffDuration: The minimum amount of time to wait before
	// retrying a job after
	// it fails.
	//
	// The default value of this field is 5 seconds.
	MinBackoffDuration string `json:"minBackoffDuration,omitempty"`

	// RetryCount: The number of attempts that the system will make to run a
	// job using the
	// exponential backoff procedure described by
	// max_doublings.
	//
	// The default value of retry_count is zero.
	//
	// If retry_count is zero, a job attempt will *not* be retried if
	// it fails. Instead the Cloud Scheduler system will wait for the
	// next scheduled execution time.
	//
	// If retry_count is set to a non-zero number then Cloud Scheduler
	// will retry failed attempts, using exponential backoff,
	// retry_count times, or until the next scheduled execution
	// time,
	// whichever comes first.
	//
	// Values greater than 5 and negative values are not allowed.
	RetryCount int64 `json:"retryCount,omitempty"`

	// ForceSendFields is a list of field names (e.g. "MaxBackoffDuration")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "MaxBackoffDuration") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *RetryConfig) MarshalJSON() ([]byte, error) {
	type NoMethod RetryConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RunJobRequest: Request message for forcing a job to run now
// using
// RunJob.
type RunJobRequest struct {
}

// Status: The `Status` type defines a logical error model that is
// suitable for different
// programming environments, including REST APIs and RPC APIs. It is
// used by
// [gRPC](https://github.com/grpc). The error model is designed to
// be:
//
// - Simple to use and understand for most users
// - Flexible enough to meet unexpected needs
//
// # Overview
//
// The `Status` message contains three pieces of data: error code, error
// message,
// and error details. The error code should be an enum value
// of
// google.rpc.Code, but it may accept additional error codes if needed.
// The
// error message should be a developer-facing English message that
// helps
// developers *understand* and *resolve* the error. If a localized
// user-facing
// error message is needed, put the localized message in the error
// details or
// localize it in the client. The optional error details may contain
// arbitrary
// information about the error. There is a predefined set of error
// detail types
// in the package `google.rpc` that can be used for common error
// conditions.
//
// # Language mapping
//
// The `Status` message is the logical representation of the error
// model, but it
// is not necessarily the actual wire format. When the `Status` message
// is
// exposed in different client libraries and different wire protocols,
// it can be
// mapped differently. For example, it will likely be mapped to some
// exceptions
// in Java, but more likely mapped to some error codes in C.
//
// # Other uses
//
// The error model and the `Status` message can be used in a variety
// of
// environments, either with or without APIs, to provide a
// consistent developer experience across different
// environments.
//
// Example uses of this error model include:
//
// - Partial errors. If a service needs to return partial errors to the
// client,
//     it may embed the `Status` in the normal response to indicate the
// partial
//     errors.
//
// - Workflow errors. A typical workflow has multiple steps. Each step
// may
//     have a `Status` message for error reporting.
//
// - Batch operations. If a client uses batch request and batch
// response, the
//     `Status` message should be used directly inside batch response,
// one for
//     each error sub-response.
//
// - Asynchronous operations. If an API call embeds asynchronous
// operation
//     results in its response, the status of those operations should
// be
//     represented directly using the `Status` message.
//
// - Logging. If some API errors are stored in logs, the message
// `Status` could
//     be used directly after any stripping needed for security/privacy
// reasons.
type Status struct {
	// Code: The status code, which should be an enum value of
	// google.rpc.Code.
	Code int64 `json:"code,omitempty"`

	// Details: A list of messages that carry the error details.  There is a
	// common set of
	// message types for APIs to use.
	Details []googleapi.RawMessage `json:"details,omitempty"`

	// Message: A developer-facing error message, which should be in
	// English. Any
	// user-facing error message should be localized and sent in
	// the
	// google.rpc.Status.details field, or localized by the client.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Status) MarshalJSON() ([]byte, error) {
	type NoMethod Status
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "cloudscheduler.projects.locations.get":

type ProjectsLocationsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets information about a location.
func (r *ProjectsLocationsService) Get(name string) *ProjectsLocationsGetCall {
	c := &ProjectsLocationsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsGetCall) Fields(s ...googleapi.Field) *ProjectsLocationsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsLocationsGetCall) IfNoneMatch(entityTag string) *ProjectsLocationsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsGetCall) Context(ctx context.Context) *ProjectsLocationsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsGetCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudscheduler.projects.locations.get" call.
// Exactly one of *Location or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Location.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsLocationsGetCall) Do(opts ...googleapi.CallOption) (*Location, error) {
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
	ret := &Location{
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
	//   "description": "Gets information about a location.",
	//   "flatPath": "v1beta1/projects/{projectsId}/locations/{locationsId}",
	//   "httpMethod": "GET",
	//   "id": "cloudscheduler.projects.locations.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Resource name for the location.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "response": {
	//     "$ref": "Location"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudscheduler.projects.locations.list":

type ProjectsLocationsListCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists information about the supported locations for this
// service.
func (r *ProjectsLocationsService) List(name string) *ProjectsLocationsListCall {
	c := &ProjectsLocationsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Filter sets the optional parameter "filter": The standard list
// filter.
func (c *ProjectsLocationsListCall) Filter(filter string) *ProjectsLocationsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": The standard list
// page size.
func (c *ProjectsLocationsListCall) PageSize(pageSize int64) *ProjectsLocationsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The standard list
// page token.
func (c *ProjectsLocationsListCall) PageToken(pageToken string) *ProjectsLocationsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsListCall) Fields(s ...googleapi.Field) *ProjectsLocationsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsLocationsListCall) IfNoneMatch(entityTag string) *ProjectsLocationsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsListCall) Context(ctx context.Context) *ProjectsLocationsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsListCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}/locations")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudscheduler.projects.locations.list" call.
// Exactly one of *ListLocationsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListLocationsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsLocationsListCall) Do(opts ...googleapi.CallOption) (*ListLocationsResponse, error) {
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
	ret := &ListLocationsResponse{
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
	//   "description": "Lists information about the supported locations for this service.",
	//   "flatPath": "v1beta1/projects/{projectsId}/locations",
	//   "httpMethod": "GET",
	//   "id": "cloudscheduler.projects.locations.list",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "The standard list filter.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The resource that owns the locations collection, if applicable.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The standard list page size.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The standard list page token.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}/locations",
	//   "response": {
	//     "$ref": "ListLocationsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsLocationsListCall) Pages(ctx context.Context, f func(*ListLocationsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "cloudscheduler.projects.locations.jobs.create":

type ProjectsLocationsJobsCreateCall struct {
	s          *Service
	parent     string
	job        *Job
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Create: Creates a job.
func (r *ProjectsLocationsJobsService) Create(parent string, job *Job) *ProjectsLocationsJobsCreateCall {
	c := &ProjectsLocationsJobsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.job = job
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsJobsCreateCall) Fields(s ...googleapi.Field) *ProjectsLocationsJobsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsJobsCreateCall) Context(ctx context.Context) *ProjectsLocationsJobsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsJobsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsJobsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.job)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+parent}/jobs")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudscheduler.projects.locations.jobs.create" call.
// Exactly one of *Job or error will be non-nil. Any non-2xx status code
// is an error. Response headers are in either
// *Job.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsLocationsJobsCreateCall) Do(opts ...googleapi.CallOption) (*Job, error) {
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
	ret := &Job{
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
	//   "description": "Creates a job.",
	//   "flatPath": "v1beta1/projects/{projectsId}/locations/{locationsId}/jobs",
	//   "httpMethod": "POST",
	//   "id": "cloudscheduler.projects.locations.jobs.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required.\n\nThe location name. For example:\n`projects/PROJECT_ID/locations/LOCATION_ID`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+parent}/jobs",
	//   "request": {
	//     "$ref": "Job"
	//   },
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudscheduler.projects.locations.jobs.delete":

type ProjectsLocationsJobsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes a job.
func (r *ProjectsLocationsJobsService) Delete(name string) *ProjectsLocationsJobsDeleteCall {
	c := &ProjectsLocationsJobsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsJobsDeleteCall) Fields(s ...googleapi.Field) *ProjectsLocationsJobsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsJobsDeleteCall) Context(ctx context.Context) *ProjectsLocationsJobsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsJobsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsJobsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("DELETE", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudscheduler.projects.locations.jobs.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsLocationsJobsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
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
	ret := &Empty{
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
	//   "description": "Deletes a job.",
	//   "flatPath": "v1beta1/projects/{projectsId}/locations/{locationsId}/jobs/{jobsId}",
	//   "httpMethod": "DELETE",
	//   "id": "cloudscheduler.projects.locations.jobs.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required.\n\nThe job name. For example:\n`projects/PROJECT_ID/locations/LOCATION_ID/jobs/JOB_ID`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/jobs/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudscheduler.projects.locations.jobs.get":

type ProjectsLocationsJobsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets a job.
func (r *ProjectsLocationsJobsService) Get(name string) *ProjectsLocationsJobsGetCall {
	c := &ProjectsLocationsJobsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsJobsGetCall) Fields(s ...googleapi.Field) *ProjectsLocationsJobsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsLocationsJobsGetCall) IfNoneMatch(entityTag string) *ProjectsLocationsJobsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsJobsGetCall) Context(ctx context.Context) *ProjectsLocationsJobsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsJobsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsJobsGetCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudscheduler.projects.locations.jobs.get" call.
// Exactly one of *Job or error will be non-nil. Any non-2xx status code
// is an error. Response headers are in either
// *Job.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsLocationsJobsGetCall) Do(opts ...googleapi.CallOption) (*Job, error) {
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
	ret := &Job{
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
	//   "description": "Gets a job.",
	//   "flatPath": "v1beta1/projects/{projectsId}/locations/{locationsId}/jobs/{jobsId}",
	//   "httpMethod": "GET",
	//   "id": "cloudscheduler.projects.locations.jobs.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required.\n\nThe job name. For example:\n`projects/PROJECT_ID/locations/LOCATION_ID/jobs/JOB_ID`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/jobs/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudscheduler.projects.locations.jobs.list":

type ProjectsLocationsJobsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists jobs.
func (r *ProjectsLocationsJobsService) List(parent string) *ProjectsLocationsJobsListCall {
	c := &ProjectsLocationsJobsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page
// size.
//
// The maximum page size is 500. If unspecified, the page size will
// be the maximum. Fewer jobs than requested might be returned,
// even if more jobs exist; use next_page_token to determine if
// more
// jobs exist.
func (c *ProjectsLocationsJobsListCall) PageSize(pageSize int64) *ProjectsLocationsJobsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server will return. To
// request the first page results, page_token must be empty. To
// request the next page of results, page_token must be the value
// of
// next_page_token returned from
// the previous call to ListJobs. It is an error to
// switch the value of filter or
// order_by while iterating through pages.
func (c *ProjectsLocationsJobsListCall) PageToken(pageToken string) *ProjectsLocationsJobsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsJobsListCall) Fields(s ...googleapi.Field) *ProjectsLocationsJobsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsLocationsJobsListCall) IfNoneMatch(entityTag string) *ProjectsLocationsJobsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsJobsListCall) Context(ctx context.Context) *ProjectsLocationsJobsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsJobsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsJobsListCall) doRequest(alt string) (*http.Response, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+parent}/jobs")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudscheduler.projects.locations.jobs.list" call.
// Exactly one of *ListJobsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListJobsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsLocationsJobsListCall) Do(opts ...googleapi.CallOption) (*ListJobsResponse, error) {
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
	ret := &ListJobsResponse{
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
	//   "description": "Lists jobs.",
	//   "flatPath": "v1beta1/projects/{projectsId}/locations/{locationsId}/jobs",
	//   "httpMethod": "GET",
	//   "id": "cloudscheduler.projects.locations.jobs.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Requested page size.\n\nThe maximum page size is 500. If unspecified, the page size will\nbe the maximum. Fewer jobs than requested might be returned,\neven if more jobs exist; use next_page_token to determine if more\njobs exist.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results the server will return. To\nrequest the first page results, page_token must be empty. To\nrequest the next page of results, page_token must be the value of\nnext_page_token returned from\nthe previous call to ListJobs. It is an error to\nswitch the value of filter or\norder_by while iterating through pages.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required.\n\nThe location name. For example:\n`projects/PROJECT_ID/locations/LOCATION_ID`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+parent}/jobs",
	//   "response": {
	//     "$ref": "ListJobsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsLocationsJobsListCall) Pages(ctx context.Context, f func(*ListJobsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "cloudscheduler.projects.locations.jobs.patch":

type ProjectsLocationsJobsPatchCall struct {
	s          *Service
	name       string
	job        *Job
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Patch: Updates a job.
//
// If successful, the updated Job is returned. If the job does
// not exist, `NOT_FOUND` is returned.
//
// If UpdateJob does not successfully return, it is possible for the
// job to be in an Job.State.UPDATE_FAILED state. A job in this state
// may
// not be executed. If this happens, retry the UpdateJob request
// until a successful response is received.
func (r *ProjectsLocationsJobsService) Patch(name string, job *Job) *ProjectsLocationsJobsPatchCall {
	c := &ProjectsLocationsJobsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.job = job
	return c
}

// UpdateMask sets the optional parameter "updateMask": A  mask used to
// specify which fields of the job are being updated.
func (c *ProjectsLocationsJobsPatchCall) UpdateMask(updateMask string) *ProjectsLocationsJobsPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsJobsPatchCall) Fields(s ...googleapi.Field) *ProjectsLocationsJobsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsJobsPatchCall) Context(ctx context.Context) *ProjectsLocationsJobsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsJobsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsJobsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.job)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("PATCH", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudscheduler.projects.locations.jobs.patch" call.
// Exactly one of *Job or error will be non-nil. Any non-2xx status code
// is an error. Response headers are in either
// *Job.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsLocationsJobsPatchCall) Do(opts ...googleapi.CallOption) (*Job, error) {
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
	ret := &Job{
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
	//   "description": "Updates a job.\n\nIf successful, the updated Job is returned. If the job does\nnot exist, `NOT_FOUND` is returned.\n\nIf UpdateJob does not successfully return, it is possible for the\njob to be in an Job.State.UPDATE_FAILED state. A job in this state may\nnot be executed. If this happens, retry the UpdateJob request\nuntil a successful response is received.",
	//   "flatPath": "v1beta1/projects/{projectsId}/locations/{locationsId}/jobs/{jobsId}",
	//   "httpMethod": "PATCH",
	//   "id": "cloudscheduler.projects.locations.jobs.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The job name. For example:\n`projects/PROJECT_ID/locations/LOCATION_ID/jobs/JOB_ID`.\n\n* `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]),\n   hyphens (-), colons (:), or periods (.).\n   For more information, see\n   [Identifying projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects)\n* `LOCATION_ID` is the canonical ID for the job's location.\n   The list of available locations can be obtained by calling\n   ListLocations.\n   For more information, see https://cloud.google.com/about/locations/.\n* `JOB_ID` can contain only letters ([A-Za-z]), numbers ([0-9]),\n   hyphens (-), or underscores (_). The maximum length is 500 characters.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/jobs/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "A  mask used to specify which fields of the job are being updated.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "request": {
	//     "$ref": "Job"
	//   },
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudscheduler.projects.locations.jobs.pause":

type ProjectsLocationsJobsPauseCall struct {
	s               *Service
	name            string
	pausejobrequest *PauseJobRequest
	urlParams_      gensupport.URLParams
	ctx_            context.Context
	header_         http.Header
}

// Pause: Pauses a job.
//
// If a job is paused then the system will stop executing the job
// until it is re-enabled via ResumeJob. The
// state of the job is stored in state; if paused it
// will be set to Job.State.PAUSED. A job must be in
// Job.State.ENABLED
// to be paused.
func (r *ProjectsLocationsJobsService) Pause(name string, pausejobrequest *PauseJobRequest) *ProjectsLocationsJobsPauseCall {
	c := &ProjectsLocationsJobsPauseCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.pausejobrequest = pausejobrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsJobsPauseCall) Fields(s ...googleapi.Field) *ProjectsLocationsJobsPauseCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsJobsPauseCall) Context(ctx context.Context) *ProjectsLocationsJobsPauseCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsJobsPauseCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsJobsPauseCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.pausejobrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}:pause")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudscheduler.projects.locations.jobs.pause" call.
// Exactly one of *Job or error will be non-nil. Any non-2xx status code
// is an error. Response headers are in either
// *Job.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsLocationsJobsPauseCall) Do(opts ...googleapi.CallOption) (*Job, error) {
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
	ret := &Job{
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
	//   "description": "Pauses a job.\n\nIf a job is paused then the system will stop executing the job\nuntil it is re-enabled via ResumeJob. The\nstate of the job is stored in state; if paused it\nwill be set to Job.State.PAUSED. A job must be in Job.State.ENABLED\nto be paused.",
	//   "flatPath": "v1beta1/projects/{projectsId}/locations/{locationsId}/jobs/{jobsId}:pause",
	//   "httpMethod": "POST",
	//   "id": "cloudscheduler.projects.locations.jobs.pause",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required.\n\nThe job name. For example:\n`projects/PROJECT_ID/locations/LOCATION_ID/jobs/JOB_ID`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/jobs/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}:pause",
	//   "request": {
	//     "$ref": "PauseJobRequest"
	//   },
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudscheduler.projects.locations.jobs.resume":

type ProjectsLocationsJobsResumeCall struct {
	s                *Service
	name             string
	resumejobrequest *ResumeJobRequest
	urlParams_       gensupport.URLParams
	ctx_             context.Context
	header_          http.Header
}

// Resume: Resume a job.
//
// This method reenables a job after it has been Job.State.PAUSED.
// The
// state of a job is stored in Job.state; after calling this method
// it
// will be set to Job.State.ENABLED. A job must be in
// Job.State.PAUSED to be resumed.
func (r *ProjectsLocationsJobsService) Resume(name string, resumejobrequest *ResumeJobRequest) *ProjectsLocationsJobsResumeCall {
	c := &ProjectsLocationsJobsResumeCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.resumejobrequest = resumejobrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsJobsResumeCall) Fields(s ...googleapi.Field) *ProjectsLocationsJobsResumeCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsJobsResumeCall) Context(ctx context.Context) *ProjectsLocationsJobsResumeCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsJobsResumeCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsJobsResumeCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.resumejobrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}:resume")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudscheduler.projects.locations.jobs.resume" call.
// Exactly one of *Job or error will be non-nil. Any non-2xx status code
// is an error. Response headers are in either
// *Job.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsLocationsJobsResumeCall) Do(opts ...googleapi.CallOption) (*Job, error) {
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
	ret := &Job{
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
	//   "description": "Resume a job.\n\nThis method reenables a job after it has been Job.State.PAUSED. The\nstate of a job is stored in Job.state; after calling this method it\nwill be set to Job.State.ENABLED. A job must be in\nJob.State.PAUSED to be resumed.",
	//   "flatPath": "v1beta1/projects/{projectsId}/locations/{locationsId}/jobs/{jobsId}:resume",
	//   "httpMethod": "POST",
	//   "id": "cloudscheduler.projects.locations.jobs.resume",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required.\n\nThe job name. For example:\n`projects/PROJECT_ID/locations/LOCATION_ID/jobs/JOB_ID`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/jobs/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}:resume",
	//   "request": {
	//     "$ref": "ResumeJobRequest"
	//   },
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudscheduler.projects.locations.jobs.run":

type ProjectsLocationsJobsRunCall struct {
	s             *Service
	name          string
	runjobrequest *RunJobRequest
	urlParams_    gensupport.URLParams
	ctx_          context.Context
	header_       http.Header
}

// Run: Forces a job to run now.
//
// When this method is called, Cloud Scheduler will dispatch the job,
// even
// if the job is already running.
func (r *ProjectsLocationsJobsService) Run(name string, runjobrequest *RunJobRequest) *ProjectsLocationsJobsRunCall {
	c := &ProjectsLocationsJobsRunCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.runjobrequest = runjobrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsJobsRunCall) Fields(s ...googleapi.Field) *ProjectsLocationsJobsRunCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsJobsRunCall) Context(ctx context.Context) *ProjectsLocationsJobsRunCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsJobsRunCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsJobsRunCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.runjobrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}:run")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudscheduler.projects.locations.jobs.run" call.
// Exactly one of *Job or error will be non-nil. Any non-2xx status code
// is an error. Response headers are in either
// *Job.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsLocationsJobsRunCall) Do(opts ...googleapi.CallOption) (*Job, error) {
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
	ret := &Job{
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
	//   "description": "Forces a job to run now.\n\nWhen this method is called, Cloud Scheduler will dispatch the job, even\nif the job is already running.",
	//   "flatPath": "v1beta1/projects/{projectsId}/locations/{locationsId}/jobs/{jobsId}:run",
	//   "httpMethod": "POST",
	//   "id": "cloudscheduler.projects.locations.jobs.run",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required.\n\nThe job name. For example:\n`projects/PROJECT_ID/locations/LOCATION_ID/jobs/JOB_ID`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/jobs/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}:run",
	//   "request": {
	//     "$ref": "RunJobRequest"
	//   },
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
