package config

import (
	"fmt"
	"time"
)

// MinecraftTime is a JSON representation of time in Minecraft.
// "forever" is decoded as a zero time and vice versa..
type MinecraftTime struct {
	time.Time
}

const forever = "\"forever\""
const layout = "\"2006-01-02 15:04:05 -0700\""

// UnmarshalJSON decodes a json byte array into a MinecrafTime
func (mt *MinecraftTime) UnmarshalJSON(byteJSON []byte) (err error) {
	strJSON := string(byteJSON)
	if strJSON == forever {
		mt.Time = time.Time{}
		return nil
	}
	mt.Time, err = time.Parse(layout, strJSON)
	return err
}

// MarshalJSON encodes a MinecrafTime into JSON
func (mt MinecraftTime) MarshalJSON() ([]byte, error) {
	if mt.Time.IsZero() {
		return []byte(forever), nil
	}
	return []byte(fmt.Sprintf("%s", mt.Time.Format(layout))), nil
}
