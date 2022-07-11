// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/data/accesslog/v3/accesslog.proto

package envoy_data_accesslog_v3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"

	v3 "github.com/kabakaev/envoyproxy-go-control-plane/envoy/config/core/v3"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}

	_ = v3.RequestMethod(0)
)

// Validate checks the field values on TCPAccessLogEntry with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *TCPAccessLogEntry) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetCommonProperties()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TCPAccessLogEntryValidationError{
				field:  "CommonProperties",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetConnectionProperties()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TCPAccessLogEntryValidationError{
				field:  "ConnectionProperties",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// TCPAccessLogEntryValidationError is the validation error returned by
// TCPAccessLogEntry.Validate if the designated constraints aren't met.
type TCPAccessLogEntryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TCPAccessLogEntryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TCPAccessLogEntryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TCPAccessLogEntryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TCPAccessLogEntryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TCPAccessLogEntryValidationError) ErrorName() string {
	return "TCPAccessLogEntryValidationError"
}

// Error satisfies the builtin error interface
func (e TCPAccessLogEntryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTCPAccessLogEntry.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TCPAccessLogEntryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TCPAccessLogEntryValidationError{}

// Validate checks the field values on HTTPAccessLogEntry with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HTTPAccessLogEntry) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetCommonProperties()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HTTPAccessLogEntryValidationError{
				field:  "CommonProperties",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ProtocolVersion

	if v, ok := interface{}(m.GetRequest()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HTTPAccessLogEntryValidationError{
				field:  "Request",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetResponse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HTTPAccessLogEntryValidationError{
				field:  "Response",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// HTTPAccessLogEntryValidationError is the validation error returned by
// HTTPAccessLogEntry.Validate if the designated constraints aren't met.
type HTTPAccessLogEntryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HTTPAccessLogEntryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HTTPAccessLogEntryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HTTPAccessLogEntryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HTTPAccessLogEntryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HTTPAccessLogEntryValidationError) ErrorName() string {
	return "HTTPAccessLogEntryValidationError"
}

// Error satisfies the builtin error interface
func (e HTTPAccessLogEntryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHTTPAccessLogEntry.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HTTPAccessLogEntryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HTTPAccessLogEntryValidationError{}

// Validate checks the field values on ConnectionProperties with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ConnectionProperties) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ReceivedBytes

	// no validation rules for SentBytes

	return nil
}

// ConnectionPropertiesValidationError is the validation error returned by
// ConnectionProperties.Validate if the designated constraints aren't met.
type ConnectionPropertiesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConnectionPropertiesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConnectionPropertiesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConnectionPropertiesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConnectionPropertiesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConnectionPropertiesValidationError) ErrorName() string {
	return "ConnectionPropertiesValidationError"
}

// Error satisfies the builtin error interface
func (e ConnectionPropertiesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConnectionProperties.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConnectionPropertiesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConnectionPropertiesValidationError{}

// Validate checks the field values on AccessLogCommon with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *AccessLogCommon) Validate() error {
	if m == nil {
		return nil
	}

	if val := m.GetSampleRate(); val <= 0 || val > 1 {
		return AccessLogCommonValidationError{
			field:  "SampleRate",
			reason: "value must be inside range (0, 1]",
		}
	}

	if v, ok := interface{}(m.GetDownstreamRemoteAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccessLogCommonValidationError{
				field:  "DownstreamRemoteAddress",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDownstreamLocalAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccessLogCommonValidationError{
				field:  "DownstreamLocalAddress",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTlsProperties()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccessLogCommonValidationError{
				field:  "TlsProperties",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetStartTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccessLogCommonValidationError{
				field:  "StartTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTimeToLastRxByte()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccessLogCommonValidationError{
				field:  "TimeToLastRxByte",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTimeToFirstUpstreamTxByte()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccessLogCommonValidationError{
				field:  "TimeToFirstUpstreamTxByte",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTimeToLastUpstreamTxByte()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccessLogCommonValidationError{
				field:  "TimeToLastUpstreamTxByte",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTimeToFirstUpstreamRxByte()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccessLogCommonValidationError{
				field:  "TimeToFirstUpstreamRxByte",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTimeToLastUpstreamRxByte()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccessLogCommonValidationError{
				field:  "TimeToLastUpstreamRxByte",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTimeToFirstDownstreamTxByte()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccessLogCommonValidationError{
				field:  "TimeToFirstDownstreamTxByte",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTimeToLastDownstreamTxByte()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccessLogCommonValidationError{
				field:  "TimeToLastDownstreamTxByte",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpstreamRemoteAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccessLogCommonValidationError{
				field:  "UpstreamRemoteAddress",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpstreamLocalAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccessLogCommonValidationError{
				field:  "UpstreamLocalAddress",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for UpstreamCluster

	if v, ok := interface{}(m.GetResponseFlags()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccessLogCommonValidationError{
				field:  "ResponseFlags",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccessLogCommonValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for UpstreamTransportFailureReason

	// no validation rules for RouteName

	if v, ok := interface{}(m.GetDownstreamDirectRemoteAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccessLogCommonValidationError{
				field:  "DownstreamDirectRemoteAddress",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for key, val := range m.GetFilterStateObjects() {
		_ = val

		// no validation rules for FilterStateObjects[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AccessLogCommonValidationError{
					field:  fmt.Sprintf("FilterStateObjects[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// AccessLogCommonValidationError is the validation error returned by
// AccessLogCommon.Validate if the designated constraints aren't met.
type AccessLogCommonValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AccessLogCommonValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AccessLogCommonValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AccessLogCommonValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AccessLogCommonValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AccessLogCommonValidationError) ErrorName() string { return "AccessLogCommonValidationError" }

// Error satisfies the builtin error interface
func (e AccessLogCommonValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccessLogCommon.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AccessLogCommonValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AccessLogCommonValidationError{}

// Validate checks the field values on ResponseFlags with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ResponseFlags) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for FailedLocalHealthcheck

	// no validation rules for NoHealthyUpstream

	// no validation rules for UpstreamRequestTimeout

	// no validation rules for LocalReset

	// no validation rules for UpstreamRemoteReset

	// no validation rules for UpstreamConnectionFailure

	// no validation rules for UpstreamConnectionTermination

	// no validation rules for UpstreamOverflow

	// no validation rules for NoRouteFound

	// no validation rules for DelayInjected

	// no validation rules for FaultInjected

	// no validation rules for RateLimited

	if v, ok := interface{}(m.GetUnauthorizedDetails()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ResponseFlagsValidationError{
				field:  "UnauthorizedDetails",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for RateLimitServiceError

	// no validation rules for DownstreamConnectionTermination

	// no validation rules for UpstreamRetryLimitExceeded

	// no validation rules for StreamIdleTimeout

	// no validation rules for InvalidEnvoyRequestHeaders

	// no validation rules for DownstreamProtocolError

	// no validation rules for UpstreamMaxStreamDurationReached

	// no validation rules for ResponseFromCacheFilter

	// no validation rules for NoFilterConfigFound

	// no validation rules for DurationTimeout

	// no validation rules for UpstreamProtocolError

	// no validation rules for NoClusterFound

	return nil
}

// ResponseFlagsValidationError is the validation error returned by
// ResponseFlags.Validate if the designated constraints aren't met.
type ResponseFlagsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResponseFlagsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResponseFlagsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResponseFlagsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResponseFlagsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResponseFlagsValidationError) ErrorName() string { return "ResponseFlagsValidationError" }

// Error satisfies the builtin error interface
func (e ResponseFlagsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResponseFlags.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResponseFlagsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResponseFlagsValidationError{}

// Validate checks the field values on TLSProperties with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *TLSProperties) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for TlsVersion

	if v, ok := interface{}(m.GetTlsCipherSuite()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TLSPropertiesValidationError{
				field:  "TlsCipherSuite",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for TlsSniHostname

	if v, ok := interface{}(m.GetLocalCertificateProperties()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TLSPropertiesValidationError{
				field:  "LocalCertificateProperties",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetPeerCertificateProperties()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TLSPropertiesValidationError{
				field:  "PeerCertificateProperties",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for TlsSessionId

	return nil
}

// TLSPropertiesValidationError is the validation error returned by
// TLSProperties.Validate if the designated constraints aren't met.
type TLSPropertiesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TLSPropertiesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TLSPropertiesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TLSPropertiesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TLSPropertiesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TLSPropertiesValidationError) ErrorName() string { return "TLSPropertiesValidationError" }

// Error satisfies the builtin error interface
func (e TLSPropertiesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTLSProperties.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TLSPropertiesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TLSPropertiesValidationError{}

// Validate checks the field values on HTTPRequestProperties with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HTTPRequestProperties) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := v3.RequestMethod_name[int32(m.GetRequestMethod())]; !ok {
		return HTTPRequestPropertiesValidationError{
			field:  "RequestMethod",
			reason: "value must be one of the defined enum values",
		}
	}

	// no validation rules for Scheme

	// no validation rules for Authority

	if v, ok := interface{}(m.GetPort()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HTTPRequestPropertiesValidationError{
				field:  "Port",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Path

	// no validation rules for UserAgent

	// no validation rules for Referer

	// no validation rules for ForwardedFor

	// no validation rules for RequestId

	// no validation rules for OriginalPath

	// no validation rules for RequestHeadersBytes

	// no validation rules for RequestBodyBytes

	// no validation rules for RequestHeaders

	return nil
}

// HTTPRequestPropertiesValidationError is the validation error returned by
// HTTPRequestProperties.Validate if the designated constraints aren't met.
type HTTPRequestPropertiesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HTTPRequestPropertiesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HTTPRequestPropertiesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HTTPRequestPropertiesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HTTPRequestPropertiesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HTTPRequestPropertiesValidationError) ErrorName() string {
	return "HTTPRequestPropertiesValidationError"
}

// Error satisfies the builtin error interface
func (e HTTPRequestPropertiesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHTTPRequestProperties.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HTTPRequestPropertiesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HTTPRequestPropertiesValidationError{}

// Validate checks the field values on HTTPResponseProperties with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HTTPResponseProperties) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetResponseCode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HTTPResponsePropertiesValidationError{
				field:  "ResponseCode",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ResponseHeadersBytes

	// no validation rules for ResponseBodyBytes

	// no validation rules for ResponseHeaders

	// no validation rules for ResponseTrailers

	// no validation rules for ResponseCodeDetails

	return nil
}

// HTTPResponsePropertiesValidationError is the validation error returned by
// HTTPResponseProperties.Validate if the designated constraints aren't met.
type HTTPResponsePropertiesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HTTPResponsePropertiesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HTTPResponsePropertiesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HTTPResponsePropertiesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HTTPResponsePropertiesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HTTPResponsePropertiesValidationError) ErrorName() string {
	return "HTTPResponsePropertiesValidationError"
}

// Error satisfies the builtin error interface
func (e HTTPResponsePropertiesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHTTPResponseProperties.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HTTPResponsePropertiesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HTTPResponsePropertiesValidationError{}

// Validate checks the field values on ResponseFlags_Unauthorized with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ResponseFlags_Unauthorized) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Reason

	return nil
}

// ResponseFlags_UnauthorizedValidationError is the validation error returned
// by ResponseFlags_Unauthorized.Validate if the designated constraints aren't met.
type ResponseFlags_UnauthorizedValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResponseFlags_UnauthorizedValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResponseFlags_UnauthorizedValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResponseFlags_UnauthorizedValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResponseFlags_UnauthorizedValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResponseFlags_UnauthorizedValidationError) ErrorName() string {
	return "ResponseFlags_UnauthorizedValidationError"
}

// Error satisfies the builtin error interface
func (e ResponseFlags_UnauthorizedValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResponseFlags_Unauthorized.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResponseFlags_UnauthorizedValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResponseFlags_UnauthorizedValidationError{}

// Validate checks the field values on TLSProperties_CertificateProperties with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *TLSProperties_CertificateProperties) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetSubjectAltName() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TLSProperties_CertificatePropertiesValidationError{
					field:  fmt.Sprintf("SubjectAltName[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Subject

	return nil
}

// TLSProperties_CertificatePropertiesValidationError is the validation error
// returned by TLSProperties_CertificateProperties.Validate if the designated
// constraints aren't met.
type TLSProperties_CertificatePropertiesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TLSProperties_CertificatePropertiesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TLSProperties_CertificatePropertiesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TLSProperties_CertificatePropertiesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TLSProperties_CertificatePropertiesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TLSProperties_CertificatePropertiesValidationError) ErrorName() string {
	return "TLSProperties_CertificatePropertiesValidationError"
}

// Error satisfies the builtin error interface
func (e TLSProperties_CertificatePropertiesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTLSProperties_CertificateProperties.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TLSProperties_CertificatePropertiesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TLSProperties_CertificatePropertiesValidationError{}

// Validate checks the field values on
// TLSProperties_CertificateProperties_SubjectAltName with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *TLSProperties_CertificateProperties_SubjectAltName) Validate() error {
	if m == nil {
		return nil
	}

	switch m.San.(type) {

	case *TLSProperties_CertificateProperties_SubjectAltName_Uri:
		// no validation rules for Uri

	case *TLSProperties_CertificateProperties_SubjectAltName_Dns:
		// no validation rules for Dns

	}

	return nil
}

// TLSProperties_CertificateProperties_SubjectAltNameValidationError is the
// validation error returned by
// TLSProperties_CertificateProperties_SubjectAltName.Validate if the
// designated constraints aren't met.
type TLSProperties_CertificateProperties_SubjectAltNameValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TLSProperties_CertificateProperties_SubjectAltNameValidationError) Field() string {
	return e.field
}

// Reason function returns reason value.
func (e TLSProperties_CertificateProperties_SubjectAltNameValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e TLSProperties_CertificateProperties_SubjectAltNameValidationError) Cause() error {
	return e.cause
}

// Key function returns key value.
func (e TLSProperties_CertificateProperties_SubjectAltNameValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TLSProperties_CertificateProperties_SubjectAltNameValidationError) ErrorName() string {
	return "TLSProperties_CertificateProperties_SubjectAltNameValidationError"
}

// Error satisfies the builtin error interface
func (e TLSProperties_CertificateProperties_SubjectAltNameValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTLSProperties_CertificateProperties_SubjectAltName.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TLSProperties_CertificateProperties_SubjectAltNameValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TLSProperties_CertificateProperties_SubjectAltNameValidationError{}
