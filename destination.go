// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/hookdeck/hookdeck-go-sdk/core"
	time "time"
)

type DestinationCreateRequest struct {
	// Name for the destination
	Name string `json:"name" url:"-"`
	// Description for the destination
	Description *core.Optional[string] `json:"description,omitempty" url:"-"`
	// Endpoint of the destination
	Url *core.Optional[string] `json:"url,omitempty" url:"-"`
	// Path for the CLI destination
	CliPath *core.Optional[string] `json:"cli_path,omitempty" url:"-"`
	// Limit of events to receive per period. Refered as Delivery Rate limit in the dashboard and documentation.
	RateLimit *core.Optional[int] `json:"rate_limit,omitempty" url:"-"`
	// Period to rate limit events by. Refered as Delivery Rate period in the dashboard and documentation.
	RateLimitPeriod        *core.Optional[DestinationCreateRequestRateLimitPeriod] `json:"rate_limit_period,omitempty" url:"-"`
	HttpMethod             *core.Optional[DestinationHttpMethod]                   `json:"http_method,omitempty" url:"-"`
	AuthMethod             *core.Optional[DestinationAuthMethodConfig]             `json:"auth_method,omitempty" url:"-"`
	PathForwardingDisabled *core.Optional[bool]                                    `json:"path_forwarding_disabled,omitempty" url:"-"`
}

type DestinationListRequest struct {
	Id         []*string                      `json:"-" url:"id,omitempty"`
	Name       *string                        `json:"-" url:"name,omitempty"`
	Disabled   *bool                          `json:"-" url:"disabled,omitempty"`
	DisabledAt *time.Time                     `json:"-" url:"disabled_at,omitempty"`
	Url        []*string                      `json:"-" url:"url,omitempty"`
	CliPath    *string                        `json:"-" url:"cli_path,omitempty"`
	OrderBy    *DestinationListRequestOrderBy `json:"-" url:"order_by,omitempty"`
	Dir        *DestinationListRequestDir     `json:"-" url:"dir,omitempty"`
	Limit      *int                           `json:"-" url:"limit,omitempty"`
	Next       *string                        `json:"-" url:"next,omitempty"`
	Prev       *string                        `json:"-" url:"prev,omitempty"`
}

// Period to rate limit events by. Refered as Delivery Rate period in the dashboard and documentation.
type DestinationCreateRequestRateLimitPeriod string

const (
	DestinationCreateRequestRateLimitPeriodSecond     DestinationCreateRequestRateLimitPeriod = "second"
	DestinationCreateRequestRateLimitPeriodMinute     DestinationCreateRequestRateLimitPeriod = "minute"
	DestinationCreateRequestRateLimitPeriodHour       DestinationCreateRequestRateLimitPeriod = "hour"
	DestinationCreateRequestRateLimitPeriodConcurrent DestinationCreateRequestRateLimitPeriod = "concurrent"
)

func NewDestinationCreateRequestRateLimitPeriodFromString(s string) (DestinationCreateRequestRateLimitPeriod, error) {
	switch s {
	case "second":
		return DestinationCreateRequestRateLimitPeriodSecond, nil
	case "minute":
		return DestinationCreateRequestRateLimitPeriodMinute, nil
	case "hour":
		return DestinationCreateRequestRateLimitPeriodHour, nil
	case "concurrent":
		return DestinationCreateRequestRateLimitPeriodConcurrent, nil
	}
	var t DestinationCreateRequestRateLimitPeriod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (d DestinationCreateRequestRateLimitPeriod) Ptr() *DestinationCreateRequestRateLimitPeriod {
	return &d
}

type DestinationDeleteResponse struct {
	// ID of the destination
	Id string `json:"id" url:"id"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (d *DestinationDeleteResponse) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DestinationDeleteResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DestinationDeleteResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DestinationDeleteResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties

	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DestinationDeleteResponse) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := core.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type DestinationListRequestDir string

const (
	DestinationListRequestDirAsc  DestinationListRequestDir = "asc"
	DestinationListRequestDirDesc DestinationListRequestDir = "desc"
)

func NewDestinationListRequestDirFromString(s string) (DestinationListRequestDir, error) {
	switch s {
	case "asc":
		return DestinationListRequestDirAsc, nil
	case "desc":
		return DestinationListRequestDirDesc, nil
	}
	var t DestinationListRequestDir
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (d DestinationListRequestDir) Ptr() *DestinationListRequestDir {
	return &d
}

type DestinationListRequestOrderBy string

const (
	DestinationListRequestOrderByCreatedAt DestinationListRequestOrderBy = "created_at"
)

func NewDestinationListRequestOrderByFromString(s string) (DestinationListRequestOrderBy, error) {
	switch s {
	case "created_at":
		return DestinationListRequestOrderByCreatedAt, nil
	}
	var t DestinationListRequestOrderBy
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (d DestinationListRequestOrderBy) Ptr() *DestinationListRequestOrderBy {
	return &d
}

// Period to rate limit events by. Refered as Delivery Rate period in the dashboard and documentation.
type DestinationUpdateRequestRateLimitPeriod string

const (
	DestinationUpdateRequestRateLimitPeriodSecond     DestinationUpdateRequestRateLimitPeriod = "second"
	DestinationUpdateRequestRateLimitPeriodMinute     DestinationUpdateRequestRateLimitPeriod = "minute"
	DestinationUpdateRequestRateLimitPeriodHour       DestinationUpdateRequestRateLimitPeriod = "hour"
	DestinationUpdateRequestRateLimitPeriodConcurrent DestinationUpdateRequestRateLimitPeriod = "concurrent"
)

func NewDestinationUpdateRequestRateLimitPeriodFromString(s string) (DestinationUpdateRequestRateLimitPeriod, error) {
	switch s {
	case "second":
		return DestinationUpdateRequestRateLimitPeriodSecond, nil
	case "minute":
		return DestinationUpdateRequestRateLimitPeriodMinute, nil
	case "hour":
		return DestinationUpdateRequestRateLimitPeriodHour, nil
	case "concurrent":
		return DestinationUpdateRequestRateLimitPeriodConcurrent, nil
	}
	var t DestinationUpdateRequestRateLimitPeriod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (d DestinationUpdateRequestRateLimitPeriod) Ptr() *DestinationUpdateRequestRateLimitPeriod {
	return &d
}

// Period to rate limit events by. Refered as Delivery Rate period in the dashboard and documentation.
type DestinationUpsertRequestRateLimitPeriod string

const (
	DestinationUpsertRequestRateLimitPeriodSecond     DestinationUpsertRequestRateLimitPeriod = "second"
	DestinationUpsertRequestRateLimitPeriodMinute     DestinationUpsertRequestRateLimitPeriod = "minute"
	DestinationUpsertRequestRateLimitPeriodHour       DestinationUpsertRequestRateLimitPeriod = "hour"
	DestinationUpsertRequestRateLimitPeriodConcurrent DestinationUpsertRequestRateLimitPeriod = "concurrent"
)

func NewDestinationUpsertRequestRateLimitPeriodFromString(s string) (DestinationUpsertRequestRateLimitPeriod, error) {
	switch s {
	case "second":
		return DestinationUpsertRequestRateLimitPeriodSecond, nil
	case "minute":
		return DestinationUpsertRequestRateLimitPeriodMinute, nil
	case "hour":
		return DestinationUpsertRequestRateLimitPeriodHour, nil
	case "concurrent":
		return DestinationUpsertRequestRateLimitPeriodConcurrent, nil
	}
	var t DestinationUpsertRequestRateLimitPeriod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (d DestinationUpsertRequestRateLimitPeriod) Ptr() *DestinationUpsertRequestRateLimitPeriod {
	return &d
}

type DestinationUpdateRequest struct {
	// Name for the destination
	Name *core.Optional[string] `json:"name,omitempty" url:"-"`
	// Description for the destination
	Description *core.Optional[string] `json:"description,omitempty" url:"-"`
	// Endpoint of the destination
	Url *core.Optional[string] `json:"url,omitempty" url:"-"`
	// Path for the CLI destination
	CliPath *core.Optional[string] `json:"cli_path,omitempty" url:"-"`
	// Limit of events to receive per period. Refered as Delivery Rate limit in the dashboard and documentation.
	RateLimit *core.Optional[int] `json:"rate_limit,omitempty" url:"-"`
	// Period to rate limit events by. Refered as Delivery Rate period in the dashboard and documentation.
	RateLimitPeriod        *core.Optional[DestinationUpdateRequestRateLimitPeriod] `json:"rate_limit_period,omitempty" url:"-"`
	HttpMethod             *core.Optional[DestinationHttpMethod]                   `json:"http_method,omitempty" url:"-"`
	AuthMethod             *core.Optional[DestinationAuthMethodConfig]             `json:"auth_method,omitempty" url:"-"`
	PathForwardingDisabled *core.Optional[bool]                                    `json:"path_forwarding_disabled,omitempty" url:"-"`
}

type DestinationUpsertRequest struct {
	// Name for the destination
	Name string `json:"name" url:"-"`
	// Description for the destination
	Description *core.Optional[string] `json:"description,omitempty" url:"-"`
	// Endpoint of the destination
	Url *core.Optional[string] `json:"url,omitempty" url:"-"`
	// Path for the CLI destination
	CliPath *core.Optional[string] `json:"cli_path,omitempty" url:"-"`
	// Limit of events to receive per period. Refered as Delivery Rate limit in the dashboard and documentation.
	RateLimit *core.Optional[int] `json:"rate_limit,omitempty" url:"-"`
	// Period to rate limit events by. Refered as Delivery Rate period in the dashboard and documentation.
	RateLimitPeriod        *core.Optional[DestinationUpsertRequestRateLimitPeriod] `json:"rate_limit_period,omitempty" url:"-"`
	HttpMethod             *core.Optional[DestinationHttpMethod]                   `json:"http_method,omitempty" url:"-"`
	AuthMethod             *core.Optional[DestinationAuthMethodConfig]             `json:"auth_method,omitempty" url:"-"`
	PathForwardingDisabled *core.Optional[bool]                                    `json:"path_forwarding_disabled,omitempty" url:"-"`
}
