package config

import (
	"net"
)

// User represents a minecraft user
type User struct {
	// Name is the display name of the minecraft user
	Name string `json:"name" firebase:"name"`
	// UUID is the unique identifier of the minecraft user
	UUID string `json:"uuid" firebase:"uuid"`
}

// OperatorUser represents an operator in the OperatorUserList
type OperatorUser struct {
	User

	// Level represents the operator level which determines what they can modify
	Level               int  `json:"level" firebase:"level"`
	BypassesPlayerLimit bool `json:"bypassesPlayerLimit" firebase:"bypasses_player_limit"`
}

// OperatorUserList represents a list of operators of a minecraft server
// IE: The list of operators from `ops.json`
type OperatorUserList []OperatorUser

// AllowUserList represents a list of users allowed to join a minecraft server.
// IE: The list of users from `whitelist.json`
type AllowUserList []User

// Deny represents why a `User` or `net.IP` was denied access to a server.
type Deny struct {
	// Created is when the ban was created
	Created MinecraftTime `json:"created" firebase:"created"`
	// Source represents where the ban originated from. IE: user or console
	Source string `json:"source" firebase:"source"`
	// Expires represents when the ban expires, usually "forever"
	Expires MinecraftTime `json:"expires" firebase:"expires"`
	// Reason represents why a ban was created
	Reason string `json:"reason" firebase:"reason"`
}

// DenyIP represents why a `net.IP` was denied access to a server.
type DenyIP struct {
	// IP is the `net.IP` that was banned, typically IPv4
	IP net.IP `json:"ip" firebase:"ip"`
	Deny
}

// DenyIPList represents a list of `net.IP` that were banned from a server and
// why. IE: The list of IPs from `banned-ips.json`
type DenyIPList []DenyIP

// DenyUser represents why a `User` was denied access to a server.
type DenyUser struct {
	User
	Deny
}

// DenyUserList represents a list of `User` that were banned from a server and
// why. IE: The list of Users from `banned-ips.json`
type DenyUserList []DenyUser
