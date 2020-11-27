package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type UserAccount struct {
	Type string `json:"type" binding:"required"`
	SpecVersion string `json:"spec_version,omitempty"`
	ID string `json:"id" binding:"required"`
	ObjectMarkingRefs []string `json:"object_marking_refs,omitempty"`
	GranularMarkings [] GranularMarking `json:"granular-markings,omitempty"`
	Defanged bool `json:"defanged,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`

	UserId string `json:"user_id,omitempty"`
	Credential string `json:"credential,omitempty"`
	AccountLogin string `json:"account_login,omitempty"`
	AccountType OpenVocab `json:"account_type,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	IsServiceAccount bool `json:"is_service_account,omitempty"`
	IsPrivileged bool `json:"is_privileged,omitempty"`
	CanEscalatePrivs bool `json:"can_escalate_privs,omitempty"`
	IsDisabled bool `json:"is_disabled,omitempty"`
	AccountCreated time.Time `json:"account_created,omitempty"`
	AccountExpired time.Time `json:"account_expired,omitempty"`
	CredentialLastChanged time.Time `json:"credential_last_changed,omitempty"`
	AccountFirstLogin time.Time `json:"account_first_login,omitempty"`
	AccountLastLogin time.Time `json:"account_last_login,omitempty"`
}

func printUserAccount(userAccount UserAccount) {
	fmt.Println("User Account:\n", marshalUserAccount(userAccount))
}

func marshalUserAccount(userAccount UserAccount) (jsonData string){
	data, e := json.MarshalIndent(userAccount, "", "  ")
	check(e)
	return string(data)
}

func unmarshalUserAccount(obj json.RawMessage) (userAccount UserAccount) {
	json.Unmarshal(obj, &userAccount)
	return userAccount
}
