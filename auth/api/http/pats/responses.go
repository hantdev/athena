package pats

import (
	"net/http"

	"github.com/hantdev/athena"
	"github.com/hantdev/athena/auth"
)

var (
	_ athena.Response = (*createPatRes)(nil)
	_ athena.Response = (*retrievePatRes)(nil)
	_ athena.Response = (*updatePatNameRes)(nil)
	_ athena.Response = (*updatePatDescriptionRes)(nil)
	_ athena.Response = (*deletePatRes)(nil)
	_ athena.Response = (*resetPatSecretRes)(nil)
	_ athena.Response = (*revokePatSecretRes)(nil)
	_ athena.Response = (*scopeRes)(nil)
	_ athena.Response = (*clearAllRes)(nil)
)

type createPatRes struct {
	auth.PAT `json:",inline"`
}

func (res createPatRes) Code() int {
	return http.StatusCreated
}

func (res createPatRes) Headers() map[string]string {
	return map[string]string{}
}

func (res createPatRes) Empty() bool {
	return false
}

type retrievePatRes struct {
	auth.PAT `json:",inline"`
}

func (res retrievePatRes) Code() int {
	return http.StatusOK
}

func (res retrievePatRes) Headers() map[string]string {
	return map[string]string{}
}

func (res retrievePatRes) Empty() bool {
	return false
}

type updatePatNameRes struct {
	auth.PAT `json:",inline"`
}

func (res updatePatNameRes) Code() int {
	return http.StatusAccepted
}

func (res updatePatNameRes) Headers() map[string]string {
	return map[string]string{}
}

func (res updatePatNameRes) Empty() bool {
	return false
}

type updatePatDescriptionRes struct {
	auth.PAT `json:",inline"`
}

func (res updatePatDescriptionRes) Code() int {
	return http.StatusAccepted
}

func (res updatePatDescriptionRes) Headers() map[string]string {
	return map[string]string{}
}

func (res updatePatDescriptionRes) Empty() bool {
	return false
}

type listPatsRes struct {
	auth.PATSPage `json:",inline"`
}

func (res listPatsRes) Code() int {
	return http.StatusOK
}

func (res listPatsRes) Headers() map[string]string {
	return map[string]string{}
}

func (res listPatsRes) Empty() bool {
	return false
}

type deletePatRes struct{}

func (res deletePatRes) Code() int {
	return http.StatusNoContent
}

func (res deletePatRes) Headers() map[string]string {
	return map[string]string{}
}

func (res deletePatRes) Empty() bool {
	return true
}

type resetPatSecretRes struct {
	auth.PAT `json:",inline"`
}

func (res resetPatSecretRes) Code() int {
	return http.StatusOK
}

func (res resetPatSecretRes) Headers() map[string]string {
	return map[string]string{}
}

func (res resetPatSecretRes) Empty() bool {
	return false
}

type revokePatSecretRes struct{}

func (res revokePatSecretRes) Code() int {
	return http.StatusNoContent
}

func (res revokePatSecretRes) Headers() map[string]string {
	return map[string]string{}
}

func (res revokePatSecretRes) Empty() bool {
	return true
}

type scopeRes struct{}

func (res scopeRes) Code() int {
	return http.StatusOK
}

func (res scopeRes) Headers() map[string]string {
	return map[string]string{}
}

func (res scopeRes) Empty() bool {
	return true
}

type clearAllRes struct{}

func (res clearAllRes) Code() int {
	return http.StatusOK
}

func (res clearAllRes) Headers() map[string]string {
	return map[string]string{}
}

func (res clearAllRes) Empty() bool {
	return true
}

type listScopeRes struct {
	auth.ScopesPage `json:",inline"`
}

func (res listScopeRes) Code() int {
	return http.StatusOK
}

func (res listScopeRes) Headers() map[string]string {
	return map[string]string{}
}

func (res listScopeRes) Empty() bool {
	return false
}