package config

import (
	"io/ioutil"
	"path"
	"testing"
)

const defaultServerProperties = `
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
level-type=default
text-filtering-config=
spawn-monsters=true
enforce-whitelist=false
resource-pack-sha1=
spawn-protection=16
max-world-size=29999984
`

func TestLoadDeafultConfig(t *testing.T) {
	testDir := t.TempDir()
	testFile := path.Join(testDir, "server.properties")
	err := ioutil.WriteFile(testFile, []byte(defaultServerProperties), 0600)
	if err != nil {
		t.Fatal(err)
	}

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
	testFile := path.Join(testDir, "server.properties")
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
