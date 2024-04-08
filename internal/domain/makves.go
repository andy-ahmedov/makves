package domain

import (
	"errors"
)

type MakvesPtr struct {
	InternalID                int     `json:"internal_id" db:"internal_id"`
	ID                        int     `json:"id" db:"id"`
	UID                       string  `json:"uid" db:"uid"`
	Domain                    string  `json:"domain" db:"domain"`
	CN                        string  `json:"cn" db:"cn"`
	Department                *string `json:"department" db:"department"`
	Title                     *string `json:"title" db:"title"`
	Who                       *string `json:"who" db:"who"`
	LogonCount                *int    `json:"logon_count" db:"logon_count"`
	NumLogons7                *int    `json:"num_logons7" db:"num_logons7"`
	NumShare7                 *int    `json:"num_share7" db:"num_share7"`
	NumFile7                  *int    `json:"num_file7" db:"num_file7"`
	NumAd7                    *int    `json:"num_ad7" db:"num_ad7"`
	NumN7                     *int    `json:"num_n7" db:"num_n7"`
	NumLogons14               *int    `json:"num_logons14" db:"num_logons14"`
	NumShare14                *int    `json:"num_share14" db:"num_share14"`
	NumFile14                 *int    `json:"num_file14" db:"num_file14"`
	NumAd14                   *int    `json:"num_ad14" db:"num_ad14"`
	NumN14                    *int    `json:"num_n14" db:"num_n14"`
	NumLogons30               *int    `json:"num_logons30" db:"num_logons30"`
	NumShare30                *int    `json:"num_share30" db:"num_share30"`
	NumFile30                 *int    `json:"num_file30" db:"num_file30"`
	NumAd30                   *int    `json:"num_ad30" db:"num_ad30"`
	NumN30                    *int    `json:"num_n30" db:"num_n30"`
	NumLogons150              *int    `json:"num_logons150" db:"num_logons150"`
	NumShare150               *int    `json:"num_share150" db:"num_share150"`
	NumFile150                *int    `json:"num_file150" db:"num_file150"`
	NumAd150                  *int    `json:"num_ad150" db:"num_ad150"`
	NumN150                   *int    `json:"num_n150" db:"num_n150"`
	NumLogons365              *int    `json:"num_logons365" db:"num_logons365"`
	NumShare365               *int    `json:"num_share365" db:"num_share365"`
	NumFile365                *int    `json:"num_file365" db:"num_file365"`
	NumAd365                  *int    `json:"num_ad365" db:"num_ad365"`
	NumN365                   *int    `json:"num_n365" db:"num_n365"`
	HasUserPrincipalName      *int    `json:"has_user_principal_name" db:"has_user_principal_name"`
	HasMail                   *int    `json:"has_mail" db:"has_mail"`
	HasPhone                  *int    `json:"has_phone" db:"has_phone"`
	FlagDisabled              *int    `json:"flag_disabled" db:"flag_disabled"`
	FlagLockout               *int    `json:"flag_lockout" db:"flag_lockout"`
	FlagPasswordNotRequired   *int    `json:"flag_password_not_required" db:"flag_password_not_required"`
	FlagPasswordCantChange    *int    `json:"flag_password_cant_change" db:"flag_password_cant_change"`
	FlagDontExpirePassword    *int    `json:"flag_dont_expire_password" db:"flag_dont_expire_password"`
	OwnedFiles                *int    `json:"owned_files" db:"owned_files"`
	NumMailboxes              *int    `json:"num_mailboxes" db:"num_mailboxes"`
	NumMemberOfGroups         *int    `json:"num_member_of_groups" db:"num_member_of_groups"`
	NumMemberOfIndirectGroups *int    `json:"num_member_of_indirect_groups" db:"num_member_of_indirect_groups"`
	MemberOfIndirectGroupsIds *string `json:"member_of_indirect_groups_ids" db:"member_of_indirect_groups_ids"`
	MemberOfGroupsIds         *string `json:"member_of_groups_ids" db:"member_of_groups_ids"`
	IsAdmin                   *int    `json:"is_admin" db:"is_admin"`
	IsService                 *int    `json:"is_service" db:"is_service"`
}

type Makves struct {
	InternalID                int    `json:"internal_id" db:"internal_id"`
	ID                        int    `json:"id" db:"id"`
	UID                       string `json:"uid" db:"uid"`
	Domain                    string `json:"domain" db:"domain"`
	CN                        string `json:"cn" db:"cn"`
	Department                string `json:"department" db:"department"`
	Title                     string `json:"title" db:"title"`
	Who                       string `json:"who" db:"who"`
	LogonCount                int    `json:"logon_count" db:"logon_count"`
	NumLogons7                int    `json:"num_logons7" db:"num_logons7"`
	NumShare7                 int    `json:"num_share7" db:"num_share7"`
	NumFile7                  int    `json:"num_file7" db:"num_file7"`
	NumAd7                    int    `json:"num_ad7" db:"num_ad7"`
	NumN7                     int    `json:"num_n7" db:"num_n7"`
	NumLogons14               int    `json:"num_logons14" db:"num_logons14"`
	NumShare14                int    `json:"num_share14" db:"num_share14"`
	NumFile14                 int    `json:"num_file14" db:"num_file14"`
	NumAd14                   int    `json:"num_ad14" db:"num_ad14"`
	NumN14                    int    `json:"num_n14" db:"num_n14"`
	NumLogons30               int    `json:"num_logons30" db:"num_logons30"`
	NumShare30                int    `json:"num_share30" db:"num_share30"`
	NumFile30                 int    `json:"num_file30" db:"num_file30"`
	NumAd30                   int    `json:"num_ad30" db:"num_ad30"`
	NumN30                    int    `json:"num_n30" db:"num_n30"`
	NumLogons150              int    `json:"num_logons150" db:"num_logons150"`
	NumShare150               int    `json:"num_share150" db:"num_share150"`
	NumFile150                int    `json:"num_file150" db:"num_file150"`
	NumAd150                  int    `json:"num_ad150" db:"num_ad150"`
	NumN150                   int    `json:"num_n150" db:"num_n150"`
	NumLogons365              int    `json:"num_logons365" db:"num_logons365"`
	NumShare365               int    `json:"num_share365" db:"num_share365"`
	NumFile365                int    `json:"num_file365" db:"num_file365"`
	NumAd365                  int    `json:"num_ad365" db:"num_ad365"`
	NumN365                   int    `json:"num_n365" db:"num_n365"`
	HasUserPrincipalName      int    `json:"has_user_principal_name" db:"has_user_principal_name"`
	HasMail                   int    `json:"has_mail" db:"has_mail"`
	HasPhone                  int    `json:"has_phone" db:"has_phone"`
	FlagDisabled              int    `json:"flag_disabled" db:"flag_disabled"`
	FlagLockout               int    `json:"flag_lockout" db:"flag_lockout"`
	FlagPasswordNotRequired   int    `json:"flag_password_not_required" db:"flag_password_not_required"`
	FlagPasswordCantChange    int    `json:"flag_password_cant_change" db:"flag_password_cant_change"`
	FlagDontExpirePassword    int    `json:"flag_dont_expire_password" db:"flag_dont_expire_password"`
	OwnedFiles                int    `json:"owned_files" db:"owned_files"`
	NumMailboxes              int    `json:"num_mailboxes" db:"num_mailboxes"`
	NumMemberOfGroups         int    `json:"num_member_of_groups" db:"num_member_of_groups"`
	NumMemberOfIndirectGroups int    `json:"num_member_of_indirect_groups" db:"num_member_of_indirect_groups"`
	MemberOfIndirectGroupsIds string `json:"member_of_indirect_groups_ids" db:"member_of_indirect_groups_ids"`
	MemberOfGroupsIds         string `json:"member_of_groups_ids" db:"member_of_groups_ids"`
	IsAdmin                   int    `json:"is_admin" db:"is_admin"`
	IsService                 int    `json:"is_service" db:"is_service"`
}

var (
	ErrDataNotFound = errors.New("data not fount")
)
