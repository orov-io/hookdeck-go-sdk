// This file was auto-generated by Fern from our API Definition.

package api

import (
	fmt "fmt"
	core "github.com/hookdeck/hookdeck-go-sdk/core"
	time "time"
)

type BookmarkCreateRequest struct {
	// ID of the event data to bookmark
	EventDataId string `json:"event_data_id"`
	// ID of the associated connection
	WebhookId string `json:"webhook_id"`
	// Descriptive name of the bookmark
	Label string `json:"label"`
	// A unique, human-friendly name for the bookmark
	Name *core.Optional[string] `json:"name,omitempty"`
}

type BookmarkListRequest struct {
	Id          []*string                   `json:"-"`
	Name        []*string                   `json:"-"`
	WebhookId   []*string                   `json:"-"`
	EventDataId []*string                   `json:"-"`
	Label       []*string                   `json:"-"`
	LastUsedAt  *time.Time                  `json:"-"`
	OrderBy     *BookmarkListRequestOrderBy `json:"-"`
	Dir         *BookmarkListRequestDir     `json:"-"`
	Limit       *int                        `json:"-"`
	Next        *string                     `json:"-"`
	Prev        *string                     `json:"-"`
}

type BookmarkTriggerRequest struct {
	// Bookmark target
	Target *core.Optional[BookmarkTriggerRequestTarget] `json:"target,omitempty"`
}

type BookmarkListRequestDir string

const (
	BookmarkListRequestDirAsc  BookmarkListRequestDir = "asc"
	BookmarkListRequestDirDesc BookmarkListRequestDir = "desc"
)

func NewBookmarkListRequestDirFromString(s string) (BookmarkListRequestDir, error) {
	switch s {
	case "asc":
		return BookmarkListRequestDirAsc, nil
	case "desc":
		return BookmarkListRequestDirDesc, nil
	}
	var t BookmarkListRequestDir
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (b BookmarkListRequestDir) Ptr() *BookmarkListRequestDir {
	return &b
}

type BookmarkListRequestOrderBy string

const (
	BookmarkListRequestOrderByCreatedAt BookmarkListRequestOrderBy = "created_at"
)

func NewBookmarkListRequestOrderByFromString(s string) (BookmarkListRequestOrderBy, error) {
	switch s {
	case "created_at":
		return BookmarkListRequestOrderByCreatedAt, nil
	}
	var t BookmarkListRequestOrderBy
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (b BookmarkListRequestOrderBy) Ptr() *BookmarkListRequestOrderBy {
	return &b
}

// Bookmark target
type BookmarkTriggerRequestTarget string

const (
	BookmarkTriggerRequestTargetHttp BookmarkTriggerRequestTarget = "http"
	BookmarkTriggerRequestTargetCli  BookmarkTriggerRequestTarget = "cli"
)

func NewBookmarkTriggerRequestTargetFromString(s string) (BookmarkTriggerRequestTarget, error) {
	switch s {
	case "http":
		return BookmarkTriggerRequestTargetHttp, nil
	case "cli":
		return BookmarkTriggerRequestTargetCli, nil
	}
	var t BookmarkTriggerRequestTarget
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (b BookmarkTriggerRequestTarget) Ptr() *BookmarkTriggerRequestTarget {
	return &b
}

type BookmarkUpdateRequest struct {
	// ID of the event data to bookmark
	EventDataId *core.Optional[string] `json:"event_data_id,omitempty"`
	// ID of the associated connection
	WebhookId *core.Optional[string] `json:"webhook_id,omitempty"`
	// Descriptive name of the bookmark
	Label *core.Optional[string] `json:"label,omitempty"`
	// A unique, human-friendly name for the bookmark
	Name *core.Optional[string] `json:"name,omitempty"`
}
