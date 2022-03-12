package config

// ## Configuration file for a typical i2pd user
// ## See https://i2pd.readthedocs.io/en/latest/user-guide/configuration/
type ConfigStruct struct {
	Base           *BaseConfigStruct           `json:"base"`
	HTTP           *HTTPConfigStruct           `json:"http"`
	HTTPProxy      *HTTPProxyConfigStruct      `json:"http_proxy"`
	SOCKSProxy     *SOCKSProxyConfigStruct     `json:"socks_proxy"`
	SAM            *SAMConfigStruct            `json:"sam"`
	BOB            *BOBConfigStruct            `json:"bob"`
	I2CP           *I2CPConfigStruct           `json:"i2cp"`
	I2PControl     *I2PControlConfigStruct     `json:"i2p_control"`
	Precomputation *PrecomputationConfigStruct `json:"precomputation"`
	UPnP           *UPnPConfigStruct           `json:"upnp"`
	Meshnets       *MeshnetsConfigStruct       `json:"meshnets"`
	Reseed         *ReseedConfigStruct         `json:"reseed"`
	AddressBook    *AddressBookConfigStruct    `json:"addressbook"`
	Limits         *LimitsConfigStruct         `json:"limits"`
	Trust          *TrustConfigStruct          `json:"trust"`
	Exploratory    *ExploratoryConfigStruct    `json:"exploratory"`
	Persist        *PersistConfigStruct        `json:"persist"`
	CPUExt         *CPUExtConfigStruct         `json:"cpu_ext"`
}

type BaseConfigStruct struct {
	TunConf    string `json:"tun_conf" default:"~/.i2pd/tunnels.conf"`  // Default: ~/.i2pd/tunnels.conf or /var/lib/i2pd/tunnels.conf
	TunnelsDir string `json:"tunnels_dir" default:"~/.i2pd/tunnels.d"`  // Default: ~/.i2pd/tunnels.d or /var/lib/i2pd/tunnels.d
	CertsDir   string `json:"certs_dir" default:"~/.i2pd/certificates"` // Default: ~/.i2pd/certificates or /var/lib/i2pd/certificates
	PidFile    string `json:"pid_file"`                                 // /run/i2pd.pid (default: i2pd.pid, not used in Windows)
	Daemon     bool   `json:"deamon"`                                   // Daemon mode. Router will go to background after start. Ignored on Windows
	Family     string `json:"family"`                                   // Specify a family, router belongs to (default - none)
	IfName     string `json:"if_name"`                                  // Updates address4/6 options if they are not set
	IfName4    string `json:"if_name4"`                                 // You can specify different interfaces for IPv4
	IfName6    string `json:"if_name6"`                                 // You can specify different interfaces for IPv6
	Address4   string `json:"address4"`                                 // Local address to bind transport sockets to Overrides host option if: For ipv 4: if ipv4 = true and nat = false
	Address6   string `json:"address6"`                                 // Local address to bind transport sockets to Overrides host option if: For ipv6: if 'host' is not set or ipv4 = true
	// Host External IPv4 or IPv6 address to listen for connections
	// By default i2pd sets IP automatically
	// Sets published NTCP2v4/SSUv4 address to 'host' value if nat = true
	// Sets published NTCP2v6/SSUv6 address to 'host' value if ipv4 = false
	Host string `json:"host"`
	// Port to listen for connections
	// By default i2pd picks random port. You MUST pick a random number too,
	// don't just uncomment this
	Port int  `json:"port"`
	IPv4 bool `json:"ipv4" default:"true"`  // Default: true; Enable communication through ipv4
	IPv6 bool `json:"ipv6" default:"false"` // Default: false; Enable communication through ipv6
	SSU  bool `json:"ssu" default:"true"`   // Default: true; Enable SSU transport
	// Bandwidth configuration
	// L limit bandwidth to 32KBs/sec, O - to 256KBs/sec, P - to 2048KBs/sec,
	// X - unlimited
	// Default is L (regular node) and X if floodfill mode enabled. If you want to
	// share more bandwidth without floodfill mode, uncomment that line and adjust
	// value to your possibilities
	Bandwidth string `json:"bandwidth" default:"L"`      // Default: L
	Share     int    `json:"share" default:"100"`        // Default: 100; Max % of bandwidth limit for transit. 0-100
	NoTransit bool   `json:"no_transit" default:"false"` // Default: false; Router will not accept transit tunnels, disabling transit traffic completely
	Floodfill bool   `json:"floodfill" default:"false"`  // Default: false; Router will be floodfill; Note: that mode uses much more network connections and CPU!
}

type HTTPConfigStruct struct {
	Enabled       bool   `json:"enabled" default:"true"`        // Default: true; set to 'false' to disable Web Console
	Address       string `json:"address" default:"127.0.0.1"`   // Default: 127.0.0.1; Address service will listen on
	Port          int    `json:"port" default:"7070"`           // Default: 7070; Port service will listen on
	Webroot       string `json:"webroot" default:"/"`           // Default: /; Path to web console
	Auth          bool   `json:"auth" default:"true"`           // Default: true; enable/disable Web Console authentication
	User          string `json:"user" default:"i2pd"`           // Default: i2pd;
	Password      string `json:"password" default:"password"`   // Default: password;
	Lang          string `json:"lang" default:"en"`             // Default: en;
	StrictHeaders bool   `json:"strict_headers" default:"true"` // Default: true; Enable strict host checking on WebUI
	Hostname      string `json:"hostname" default:"localhost"`  // Default: localhost; Expected hostname for WebUI
}

type HTTPProxyConfigStruct struct {
	Enabled       bool   `json:"enabled" default:"true"`             // Default: true; set to 'false' to disable HTTP Proxy
	Address       string `json:"address" default:"127.0.0.1"`        // Default: 127.0.0.1; Address service will listen on
	Port          int    `json:"port" default:"4444"`                // Default: 4444; Port service will listen on
	Keys          string `json:"keys" default:"http-proxy-keys.dat"` // Default: http-proxy-keys.dat; Optional keys file for proxy local destination
	AddressHelper bool   `json:"address_helper" default:"true"`      // Default: true; Enable/Disable address helper for adding .i2p domains with "jump URLs"
}

type SOCKSProxyConfigStruct struct {
	Enabled         bool   `json:"enabled" default:"true"`              // Default: true; set to 'false' to disable SOCKS Proxy
	Address         string `json:"address" default:"127.0.0.1"`         // Default: 127.0.0.1; Address service will listen on
	Port            int    `json:"port" default:"4447"`                 // Default: 4447; Port service will listen on
	Keys            string `json:"keys" default:"socks-proxy-keys.dat"` // Default: socks-proxy-keys.dat; Optional keys file for proxy local destination
	OutproxyEnabled bool   `json:"outproxy_enabled" default:"false"`    // Default: false; Socks outproxy. Example below is set to use Tor for all connections except i2p
	Outproxy        string `json:"outproxy" default:"127.0.0.1"`        // Default: 127.0.0.1; Address of outproxy
	OutproxyPort    int    `json:"outproxy_port" default:"9050"`        // Default: 9050; Port of outproxy
}

type SAMConfigStruct struct {
	Enabled      bool   `json:"enabled" default:"true"`       // Default: true; set to 'false' to disable SAM Bridge
	Address      string `json:"address" default:"127.0.0.1"`  // Default: 127.0.0.1; Address service will listen on
	Port         int    `json:"port" default:"7656"`          // Default: 7656; Port service will listen on
	SingleThread bool   `json:"single_thread" default:"true"` // Default: true; If false every SAM session runs in own thread
}

type BOBConfigStruct struct {
	Enabled bool   `json:"enabled" default:"false"`     // Default: false; set to 'true' to disable BOB command channel
	Address string `json:"address" default:"127.0.0.1"` // Default: 127.0.0.1; Address service will listen on
	Port    int    `json:"port" default:"2827"`         // Default: 2827; Port service will listen on
}

type I2CPConfigStruct struct {
	Enabled bool   `json:"enabled" default:"false"`     // Default: false; set to 'true' to enable I2CP protocol
	Address string `json:"address" default:"127.0.0.1"` // Default: 127.0.0.1; Address service will listen on
	Port    int    `json:"port" default:"7654"`         // Default: 7654; Port service will listen on
}

type I2PControlConfigStruct struct {
	Enabled  bool   `json:"enabled" default:"false"`     // Default: false; set to 'true' to enable I2PControl protocol
	Address  string `json:"address" default:"127.0.0.1"` // Default: 127.0.0.1; Address service will listen on
	Port     int    `json:"port" default:"7650"`         // Default: 7650; Port service will listen on
	Password string `json:"password" default:"itoopie"`  // Default: "itoopie"; Authentication password
}

type PrecomputationConfigStruct struct {
	Elgamal bool `json:"elgamal" default:"true"` // Default: true; Enable or disable elgamal precomputation table; By default, enabled on i386 hosts
}

type UPnPConfigStruct struct {
	Enabled bool   `json:"enabled" default:"false"` // Default: false; Enable or disable UPnP: automatic port forwarding (enabled by default in WINDOWS, ANDROID)
	Name    string `json:"name" default:"I2Pd"`     // Default: I2Pd; Name i2pd appears in UPnP forwardings list
}

type MeshnetsConfigStruct struct {
	Yggdrasil  bool   `json:"yggdrasil" default:"false"` // Default: false; Enable connectivity over the Yggdrasil network
	YggAddress string `json:"yggaddress"`                // You can bind address from your Yggdrasil subnet 300::/64; The address must first be added to the network interface
}

type ReseedConfigStruct struct {
	Verify    bool   `json:"verify" default:"false"` // Default: false; Options for bootstrapping into I2P network, aka reseeding; Enable or disable reseed data verification.
	URLs      string `json:"urls"`                   // URLs to request reseed data from, separated by comma; Default: "mainline" I2P Network reseeds
	YggURLs   string `json:"ygg_urls"`               // Reseed URLs through the Yggdrasil, separated by comma
	File      string `json:"file"`                   // Path to local reseed data file (.su3) for manual reseeding or HTTPS URL to reseed from
	ZIPfile   string `json:"zip_file"`               // Path to local ZIP file or HTTPS URL to reseed from
	Proxy     string `json:"proxy"`                  // If you run i2pd behind a proxy server, set proxy server for reseeding here; Should be http://address:port or socks://address:port
	Threshold int    `json:"threshold" default:"25"` // Minimum number of known routers, below which i2pd triggers reseeding. 25 by default
}

type AddressBookConfigStruct struct {
	DefaultURL    string `json:"default_url" default:"reg.i2p"` // Default: reg.i2p at "mainline" I2P Network; AddressBook subscription URL for initial setup
	Subscriptions string `json:"subscriptions"`                 // Optional subscriptions URLs, separated by comma
}

type LimitsConfigStruct struct {
	TransitTunnels int `json:"transit_tunnels" default:"2500"` // Default: 2500; Maximum active transit sessions
	OpenFiles      int `json:"open_files" default:"0"`         // Default: 0; Limit number of open file descriptors (0 - use system limit)
	CoreSize       int `json:"core_size" default:"0"`          // Default: 0; Maximum size of corefile in Kb (0 - use system limit)
}

type TrustConfigStruct struct {
	Enabled bool   `json:"enabled" default:"false"` // Default: false; Enable explicit trust options
	Family  string `json:"family"`                  // Default ; Make direct I2P connections only to routers in specified Family
	Routers string `json:"routers"`                 // Default: ; Make direct I2P connections only to routers specified here. Comma separated list of base64 identities.
	Hidden  bool   `json:"hidden" default:"true"`   // Default: false; Should we hide our router from other routers?
}

// ExploratoryConfigStruct. Exploratory tunnels settings with default values
type ExploratoryConfigStruct struct {
	InboundLength    int `json:"inbound_legnth" default:"2"`
	InboundQuantity  int `json:"inbound_quantity" default:"3"`
	OutboundLength   int `json:"outbound_legnth" default:"2"`
	OutboundQuantity int `json:"outbound_quantity" default:"3"`
}

type PersistConfigStruct struct {
	Profiles    bool `json:"profiles" default:"true"`    // Default: true; Save peer profiles on disk
	AddressBook bool `json:"addressbook" default:"true"` // Default: true; Save full addresses on disk
}

type CPUExtConfigStruct struct {
	AESNi bool `json:"aesni" default:"true"` // Default: true; Use CPU AES-NI instructions set when work with cryptography when available
	AVX   bool `json:"avx" default:"true"`   // Default: true; Use CPU AVX instructions set when work with cryptography when available
}
