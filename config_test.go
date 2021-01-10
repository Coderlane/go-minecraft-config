package config

import (
	"io/ioutil"
	"path"
	"testing"
)

const testServerProperties = `
#Minecraft server properties
#Sat Jan 09 18:23:20 EST 2021
enable-jmx-monitoring=false
rcon.port=25575
level-seed=
gamemode=survival
enable-command-block=false
enable-query=false
generator-settings=
level-name=world
motd=A Minecraft Server
query.port=25565
pvp=true
generate-structures=true
difficulty=easy
network-compression-threshold=256
max-tick-time=60000
max-players=20
use-native-transport=true
online-mode=true
enable-status=true
allow-flight=false
broadcast-rcon-to-ops=true
view-distance=10
max-build-height=256
server-ip=
allow-nether=true
server-port=25565
enable-rcon=false
sync-chunk-writes=true
op-permission-level=4
prevent-proxy-connections=false
resource-pack=
entity-broadcast-range-percentage=100
rcon.password=
player-idle-timeout=0
force-gamemode=false
rate-limit=0
hardcore=false
white-list=false
broadcast-console-to-ops=true
spawn-npcs=true
spawn-animals=true
snooper-enabled=true
function-permission-level=2
level-type=test
text-filtering-config=
spawn-monsters=true
enforce-whitelist=false
resource-pack-sha1=
spawn-protection=16
max-world-size=29999984
`

const testDenyIPs = `[
  {
    "ip": "127.0.0.134",
    "created": "2020-09-14 23:05:05 -0400",
    "source": "Server",
    "expires": "forever",
    "reason": "Deny by an operator."
  }
]`

const testDenyUsers = `[
  {
    "uuid": "9b15dea6-606e-47a4-a241-000000000000",
    "name": "Test",
    "created": "2020-09-14 23:01:51 -0400",
    "source": "Server",
    "expires": "forever",
    "reason": "Banned by an operator."
  }
]`

const testAllowUsers = `[
  {
    "uuid": "9b15dea6-606e-47a4-a241-000000000000",
    "name": "Test"
  }
]`

const testOPs = `[
  {
    "uuid": "9b15dea6-606e-47a4-a241-000000000000",
    "name": "Test",
    "level": 4,
    "bypassesPlayerLimit": false
  }
]`

func writeTestFile(t *testing.T, dir, file, data string) {
	testFile := path.Join(dir, file)
	err := ioutil.WriteFile(testFile, []byte(data), 0600)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLoadDeafultConfig(t *testing.T) {
	testDir := t.TempDir()
	writeTestFile(t, testDir, filenameServerProperties, testServerProperties)
	cfg, err := LoadConfig(testDir)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", cfg)
}

func TestLoadMissingConfig(t *testing.T) {
	cfg, err := LoadConfig(t.TempDir())
	if err == nil {
		t.Errorf("Expected an error")
	}
	if cfg != nil {
		t.Errorf("Expected config to be nil")
	}
	t.Log(err)
}

func TestLoadInvalidConfig(t *testing.T) {
	testDir := t.TempDir()
	testFile := path.Join(testDir, filenameServerProperties)
	err := ioutil.WriteFile(testFile, []byte("query.port=false"), 0600)
	if err != nil {
		t.Fatal(err)
	}

	cfg, err := LoadConfig(testDir)
	if err == nil {
		t.Errorf("Expected an error")
	}
	if cfg != nil {
		t.Errorf("Expected config to be nil")
	}
	t.Log(err)
}

func TestLoadDenyIPList(t *testing.T) {
	testDir := t.TempDir()
	writeTestFile(t, testDir, filenameBannedIPs, testDenyIPs)

	cfg, err := LoadDenyIPList(testDir)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", cfg)
}

func TestLoadDenyUserList(t *testing.T) {
	testDir := t.TempDir()
	writeTestFile(t, testDir, filenameBannedPlayers, testDenyUsers)

	cfg, err := LoadDenyUserList(testDir)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", cfg)
}

func TestLoadAllowUserList(t *testing.T) {
	testDir := t.TempDir()
	writeTestFile(t, testDir, filenameWhitelist, testAllowUsers)

	cfg, err := LoadAllowUserList(testDir)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", cfg)
}

func TestLoadOperatorUserList(t *testing.T) {
	testDir := t.TempDir()
	writeTestFile(t, testDir, filenameOPs, testOPs)

	cfg, err := LoadOperatorUserList(testDir)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", cfg)
}

func TestLoadMissingOperatorUserList(t *testing.T) {
	testDir := t.TempDir()
	_, err := LoadOperatorUserList(testDir)
	if err == nil {
		t.Fatalf("Expected an error")
	}
	t.Logf("%+v\n", err)
}

func TestLoadCorruptedOperatorUserList(t *testing.T) {
	testDir := t.TempDir()
	writeTestFile(t, testDir, filenameOPs, "invalid")

	_, err := LoadOperatorUserList(testDir)
	if err == nil {
		t.Fatalf("Expected an error")
	}
	t.Logf("%+v\n", err)
}
