package config

import (
	"path"

	"github.com/magiconair/properties"
)

// Config represents the Minecraft server config
type Config struct {
	AllowFlight                    bool   `properties:"allow-flight,default=false" json:"allow_flight" firebase:"allow_flight"`
	AllowNether                    bool   `properties:"allow-nether,default=true" json:"allow_nether" firebase:"allow_nether"`
	BroadcastConsoleToOps          bool   `properties:"broadcast-console-to-ops,default=true" json:"broadcast_console_to_ops" firebase:"broadcast_console_to_ops"`
	BroadcastRconToOps             bool   `properties:"broadcast-rcon-to-ops,default=true" json:"broadcast_rcon_to_ops" firebase:"broadcast_rcon_to_ops"`
	Difficulty                     string `properties:"difficulty,default=easy" json:"difficulty" firebase:"difficulty"`
	EnableCommandBlock             bool   `properties:"enable-command-block,default=false" json:"enable_command_block" firebase:"enable_command_block"`
	EnableJmxMonitoring            bool   `properties:"enable-jmx-monitoring,default=false" json:"enable_jmx_monitoring" firebase:"enable_jmx_monitoring"`
	EnableRcon                     bool   `properties:"enable-rcon,default=false" json:"enable_rcon" firebase:"enable_rcon"`
	SyncChunkWrites                bool   `properties:"sync-chunk-writes,default=true" json:"sync_chunk_writes" firebase:"sync_chunk_writes"`
	EnableStatus                   bool   `properties:"enable-status,default=true" json:"enable_status" firebase:"enable_status"`
	EnableQuery                    bool   `properties:"enable-query,default=false" json:"enable_query" firebase:"enable_query"`
	EntityBroadcastRangePercentage int    `properties:"entity-broadcast-range-percentage,default=100" json:"entity_broadcast_range_percentage" firebase:"entity_broadcast_range_percentage"`
	ForceGamemode                  bool   `properties:"force-gamemode,default=false" json:"force_gamemode" firebase:"force_gamemode"`
	Gamemode                       string `properties:"gamemode,default=survival" json:"gamemode" firebase:"gamemode"`
	GenerateStructures             bool   `properties:"generate-structures,default=true" json:"generate_structures" firebase:"generate_structures"`
	GeneratorSettings              string `properties:"generator-settings,default=" json:"generator_settings" firebase:"generator_settings"`
	Hardcore                       bool   `properties:"hardcore,default=false" json:"hardcore" firebase:"hardcore"`
	LevelName                      string `properties:"level-name,default=world" json:"level_name" firebase:"level_name"`
	LevelSeed                      string `properties:"level-seed,default=" json:"level_seed" firebase:"level_seed"`
	LevelType                      string `properties:"level-type,default=default" json:"level_type" firebase:"level_type"`
	MaxBuildHeight                 int    `properties:"max-build-height,default=256" json:"max_build_height" firebase:"max_build_height"`
	MaxPlayers                     int    `properties:"max-players,default=20" json:"max_players" firebase:"max_players"`
	MaxTickTime                    int    `properties:"max-tick-time,default=60000" json:"max_tick_time" firebase:"max_tick_time"`
	MaxWorldSize                   int    `properties:"max-world-size,default=29999984" json:"max_world_size" firebase:"max_world_size"`
	MotD                           string `properties:"motd,default=A Minecraft Server" json:"motd" firebase:"motd"`
	NetworkCompressionThreshold    int    `properties:"network-compression-threshold,default=256" json:"network_compression_threshold" firebase:"network_compression_threshold"`
	OnlineMode                     bool   `properties:"online-mode,default=true" json:"online_mode" firebase:"online_mode"`
	OpPermissionLevel              int    `properties:"op-permission-level,default=4" json:"op_permission_level" firebase:"op_permission_level"`
	PlayerIdleTimeout              int    `properties:"player-idle-timeout,default=0" json:"player_idle_timeout" firebase:"player_idle_timeout"`
	PreventProxyConnections        bool   `properties:"prevent-proxy-connections,default=false" json:"prevent_proxy_connections" firebase:"prevent_proxy_connections"`
	PvP                            bool   `properties:"pvp,default=true" json:"pvp" firebase:"pvp"`
	QueryPort                      int    `properties:"query.port,default=25565" json:"query_port" firebase:"query_port"`
	RateLimit                      int    `properties:"rate-limit,default=0" json:"rate_limit" firebase:"rate_limit"`
	RconPassword                   string `properties:"rcon.password,default=" json:"rcon_password" firebase:"rcon_password"`
	RconPort                       int    `properties:"rcon.port,default=25575" json:"rcon_port" firebase:"rcon_port"`
	ResourcePack                   string `properties:"resource-pack,default=" json:"resource_pack" firebase:"resource_pack"`
	ResourcePackSHA1               string `properties:"resource-pack-sha1,default=" json:"resource_pack_sha1" firebase:"resource_pack_sha1"`
	RequireResourcePack            bool   `properties:"require-resource-pack,default=false" json:"require_resource_pack" firebase:"require_resource_pack"`
	ServerIP                       string `properties:"server-ip,default=" json:"server_ip" firebase:"server_ip"`
	ServerPort                     int    `properties:"server-port,default=25565" json:"server_port" firebase:"server_port"`
	SnooperEnabled                 bool   `properties:"snooper-enabled,default=true" json:"snooper_enabled" firebase:"snooper_enabled"`
	SpawnAnimals                   bool   `properties:"spawn-animals,default=true" json:"spawn_animals" firebase:"spawn_animals"`
	SpawnMonsters                  bool   `properties:"spawn-monsters,default=true" json:"spawn_monsters" firebase:"spawn_monsters"`
	SpawnNpcs                      bool   `properties:"spawn-npcs,default=true" json:"spawn_npcs" firebase:"spawn_npcs"`
	ViewDistance                   int    `properties:"view-distance,default=10" json:"view_distance" firebase:"view_distance"`
	WhiteList                      bool   `properties:"white-list,default=false" json:"white_list" firebase:"white_list"`
	EnforceWhitelist               bool   `properties:"enforce-whitelist,default=false" json:"enforce_whitelist" firebase:"enforce_whitelist"`
}

// LoadConfig loads server config from the provided directory
func LoadConfig(dir string) (*Config, error) {
	cfgFile := path.Join(dir, "server.properties")

	props, err := properties.LoadFile(cfgFile, properties.UTF8)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err = props.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
