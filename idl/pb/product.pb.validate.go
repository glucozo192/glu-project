// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: product.proto

package pb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Product with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Product) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Product with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ProductMultiError, or nil if none found.
func (m *Product) ValidateAll() error {
	return m.validate(true)
}

func (m *Product) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Sku

	// no validation rules for Name

	// no validation rules for Description

	// no validation rules for ProductStatusId

	// no validation rules for RegularPrice

	// no validation rules for DiscountPrice

	// no validation rules for Quantity

	// no validation rules for Taxable

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProductValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProductValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProductValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProductValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProductValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProductValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ProductMultiError(errors)
	}

	return nil
}

// ProductMultiError is an error wrapping multiple validation errors returned
// by Product.ValidateAll() if the designated constraints aren't met.
type ProductMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProductMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProductMultiError) AllErrors() []error { return m }

// ProductValidationError is the validation error returned by Product.Validate
// if the designated constraints aren't met.
type ProductValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductValidationError) ErrorName() string { return "ProductValidationError" }

// Error satisfies the builtin error interface
func (e ProductValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProduct.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductValidationError{}

// Validate checks the field values on InsertProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *InsertProductRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on InsertProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// InsertProductRequestMultiError, or nil if none found.
func (m *InsertProductRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *InsertProductRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetProduct()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, InsertProductRequestValidationError{
					field:  "Product",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, InsertProductRequestValidationError{
					field:  "Product",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProduct()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return InsertProductRequestValidationError{
				field:  "Product",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return InsertProductRequestMultiError(errors)
	}

	return nil
}

// InsertProductRequestMultiError is an error wrapping multiple validation
// errors returned by InsertProductRequest.ValidateAll() if the designated
// constraints aren't met.
type InsertProductRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InsertProductRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InsertProductRequestMultiError) AllErrors() []error { return m }

// InsertProductRequestValidationError is the validation error returned by
// InsertProductRequest.Validate if the designated constraints aren't met.
type InsertProductRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InsertProductRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InsertProductRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InsertProductRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InsertProductRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InsertProductRequestValidationError) ErrorName() string {
	return "InsertProductRequestValidationError"
}

// Error satisfies the builtin error interface
func (e InsertProductRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInsertProductRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InsertProductRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InsertProductRequestValidationError{}

// Validate checks the field values on InsertProductResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *InsertProductResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on InsertProductResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// InsertProductResponseMultiError, or nil if none found.
func (m *InsertProductResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *InsertProductResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetProduct()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, InsertProductResponseValidationError{
					field:  "Product",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, InsertProductResponseValidationError{
					field:  "Product",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProduct()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return InsertProductResponseValidationError{
				field:  "Product",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return InsertProductResponseMultiError(errors)
	}

	return nil
}

// InsertProductResponseMultiError is an error wrapping multiple validation
// errors returned by InsertProductResponse.ValidateAll() if the designated
// constraints aren't met.
type InsertProductResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InsertProductResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InsertProductResponseMultiError) AllErrors() []error { return m }

// InsertProductResponseValidationError is the validation error returned by
// InsertProductResponse.Validate if the designated constraints aren't met.
type InsertProductResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InsertProductResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InsertProductResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InsertProductResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InsertProductResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InsertProductResponseValidationError) ErrorName() string {
	return "InsertProductResponseValidationError"
}

// Error satisfies the builtin error interface
func (e InsertProductResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInsertProductResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InsertProductResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InsertProductResponseValidationError{}

// Validate checks the field values on ListProductsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListProductsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListProductsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListProductsRequestMultiError, or nil if none found.
func (m *ListProductsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListProductsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ProductId

	if len(errors) > 0 {
		return ListProductsRequestMultiError(errors)
	}

	return nil
}

// ListProductsRequestMultiError is an error wrapping multiple validation
// errors returned by ListProductsRequest.ValidateAll() if the designated
// constraints aren't met.
type ListProductsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListProductsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListProductsRequestMultiError) AllErrors() []error { return m }

// ListProductsRequestValidationError is the validation error returned by
// ListProductsRequest.Validate if the designated constraints aren't met.
type ListProductsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListProductsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListProductsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListProductsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListProductsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListProductsRequestValidationError) ErrorName() string {
	return "ListProductsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListProductsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListProductsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListProductsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListProductsRequestValidationError{}

// Validate checks the field values on ListProductsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListProductsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListProductsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListProductsResponseMultiError, or nil if none found.
func (m *ListProductsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListProductsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetProduct()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListProductsResponseValidationError{
					field:  "Product",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListProductsResponseValidationError{
					field:  "Product",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProduct()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListProductsResponseValidationError{
				field:  "Product",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ListProductsResponseMultiError(errors)
	}

	return nil
}

// ListProductsResponseMultiError is an error wrapping multiple validation
// errors returned by ListProductsResponse.ValidateAll() if the designated
// constraints aren't met.
type ListProductsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListProductsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListProductsResponseMultiError) AllErrors() []error { return m }

// ListProductsResponseValidationError is the validation error returned by
// ListProductsResponse.Validate if the designated constraints aren't met.
type ListProductsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListProductsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListProductsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListProductsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListProductsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListProductsResponseValidationError) ErrorName() string {
	return "ListProductsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListProductsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListProductsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListProductsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListProductsResponseValidationError{}

// Validate checks the field values on ListAllProductsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListAllProductsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListAllProductsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListAllProductsRequestMultiError, or nil if none found.
func (m *ListAllProductsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListAllProductsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListAllProductsRequestMultiError(errors)
	}

	return nil
}

// ListAllProductsRequestMultiError is an error wrapping multiple validation
// errors returned by ListAllProductsRequest.ValidateAll() if the designated
// constraints aren't met.
type ListAllProductsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListAllProductsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListAllProductsRequestMultiError) AllErrors() []error { return m }

// ListAllProductsRequestValidationError is the validation error returned by
// ListAllProductsRequest.Validate if the designated constraints aren't met.
type ListAllProductsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAllProductsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAllProductsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAllProductsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAllProductsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAllProductsRequestValidationError) ErrorName() string {
	return "ListAllProductsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListAllProductsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAllProductsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAllProductsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAllProductsRequestValidationError{}

// Validate checks the field values on ListAllProductsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListAllProductsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListAllProductsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListAllProductsResponseMultiError, or nil if none found.
func (m *ListAllProductsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListAllProductsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetProducts() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListAllProductsResponseValidationError{
						field:  fmt.Sprintf("Products[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListAllProductsResponseValidationError{
						field:  fmt.Sprintf("Products[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListAllProductsResponseValidationError{
					field:  fmt.Sprintf("Products[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListAllProductsResponseMultiError(errors)
	}

	return nil
}

// ListAllProductsResponseMultiError is an error wrapping multiple validation
// errors returned by ListAllProductsResponse.ValidateAll() if the designated
// constraints aren't met.
type ListAllProductsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListAllProductsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListAllProductsResponseMultiError) AllErrors() []error { return m }

// ListAllProductsResponseValidationError is the validation error returned by
// ListAllProductsResponse.Validate if the designated constraints aren't met.
type ListAllProductsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAllProductsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAllProductsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAllProductsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAllProductsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAllProductsResponseValidationError) ErrorName() string {
	return "ListAllProductsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListAllProductsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAllProductsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAllProductsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAllProductsResponseValidationError{}

// Validate checks the field values on DeleteProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteProductRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteProductRequestMultiError, or nil if none found.
func (m *DeleteProductRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteProductRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteProductRequestMultiError(errors)
	}

	return nil
}

// DeleteProductRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteProductRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteProductRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteProductRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteProductRequestMultiError) AllErrors() []error { return m }

// DeleteProductRequestValidationError is the validation error returned by
// DeleteProductRequest.Validate if the designated constraints aren't met.
type DeleteProductRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteProductRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteProductRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteProductRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteProductRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteProductRequestValidationError) ErrorName() string {
	return "DeleteProductRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteProductRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteProductRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteProductRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteProductRequestValidationError{}

// Validate checks the field values on DeleteProductResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteProductResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteProductResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteProductResponseMultiError, or nil if none found.
func (m *DeleteProductResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteProductResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if len(errors) > 0 {
		return DeleteProductResponseMultiError(errors)
	}

	return nil
}

// DeleteProductResponseMultiError is an error wrapping multiple validation
// errors returned by DeleteProductResponse.ValidateAll() if the designated
// constraints aren't met.
type DeleteProductResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteProductResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteProductResponseMultiError) AllErrors() []error { return m }

// DeleteProductResponseValidationError is the validation error returned by
// DeleteProductResponse.Validate if the designated constraints aren't met.
type DeleteProductResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteProductResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteProductResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteProductResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteProductResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteProductResponseValidationError) ErrorName() string {
	return "DeleteProductResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteProductResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteProductResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteProductResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteProductResponseValidationError{}
