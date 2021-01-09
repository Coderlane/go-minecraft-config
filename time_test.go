package config

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestEncodeDecodeTime(t *testing.T) {
	testCases := []string{
		"forever",
		"2020-09-14 23:01:51 -0400",
	}
	for _, tc := range testCases {
		t.Run(tc, func(t *testing.T) {
			inputBytes := []byte("\"" + tc + "\"")

			var decodedTime MinecraftTime
			err := json.Unmarshal(inputBytes, &decodedTime)
			if err != nil {
				t.Fatal(err)
			}

			outputBytes, err := json.Marshal(decodedTime)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(inputBytes, outputBytes) {
				t.Errorf("Expected: %s Got: %s", inputBytes, outputBytes)
			}
		})
	}
}
