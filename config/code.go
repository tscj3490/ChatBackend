package config

import (
	"../model"
)

// Constatns for role of admin
const (
	MinRoleCode   = 1000
	AdminCode     = 100
	DeveloperCode = 101
	MarketerCode  = 102
	ReporterCode  = 103
)

// Enum for Status for every models
const (
	Disabled model.Status = iota
	Enabled
	Pending = 100
	Verified
	Offline = 200
	Online
)

// Enum for UserSetting
const (
	SettingNone int = iota
	SettingPenalty
	SettingSite
	SettingIllegal
	SettingBlacklist
	SettingShortcut
	SettingLocation
	SettingWifi
	SettingUsb
	SettingPrint
	SettingSleeptime
	SettingTimezone
	SettingLanguage
)
