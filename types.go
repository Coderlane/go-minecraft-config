package config

import (
	"net"
)

type User struct {
	Name string `json:"name" firebase:"name"`
	UUID string `json:"uuid" firebase:"uuid"`
}

type OPUser struct {
	User

	Level               int  `json:"level" firebase:"level"`
	BypassesPlayerLimit bool `json:"bypassesPlayerLimit" firebase:"bypasses_player_limit"`
}

type AllowList []User

type Deny struct {
	Created MinecraftTime `json:"created" firebase:"created"`
	Source  string        `json:"source" firebase:"source"`
	Expires MinecraftTime `json:"expires" firebase:"expires"`
	Reason  string        `json:"reason" firebase:"reason"`
}

type DenyIP struct {
	IP net.IP `json:"ip" firebase:"ip"`
	Deny
}

type DenyUser struct {
	User
	Deny
}
