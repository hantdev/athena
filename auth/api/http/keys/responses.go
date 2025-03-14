package keys

import (
	"net/http"
	"time"

	"github.com/hantdev/mitras"
	"github.com/hantdev/mitras/auth"
)

var (
	_ mitras.Response = (*issueKeyRes)(nil)
	_ mitras.Response = (*revokeKeyRes)(nil)
)

type issueKeyRes struct {
	ID        string     `json:"id,omitempty"`
	Value     string     `json:"value,omitempty"`
	IssuedAt  time.Time  `json:"issued_at,omitempty"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

func (res issueKeyRes) Code() int {
	return http.StatusCreated
}

func (res issueKeyRes) Headers() map[string]string {
	return map[string]string{}
}

func (res issueKeyRes) Empty() bool {
	return res.Value == ""
}

type retrieveKeyRes struct {
	ID        string       `json:"id,omitempty"`
	IssuerID  string       `json:"issuer_id,omitempty"`
	Subject   string       `json:"subject,omitempty"`
	Type      auth.KeyType `json:"type,omitempty"`
	IssuedAt  time.Time    `json:"issued_at,omitempty"`
	ExpiresAt *time.Time   `json:"expires_at,omitempty"`
}

func (res retrieveKeyRes) Code() int {
	return http.StatusOK
}

func (res retrieveKeyRes) Headers() map[string]string {
	return map[string]string{}
}

func (res retrieveKeyRes) Empty() bool {
	return false
}

type revokeKeyRes struct{}

func (res revokeKeyRes) Code() int {
	return http.StatusNoContent
}

func (res revokeKeyRes) Headers() map[string]string {
	return map[string]string{}
}

func (res revokeKeyRes) Empty() bool {
	return true
}
