/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import "time"

// Manifest use to store server manifest data
// All string slice is multi language and in order by ManifestLanguages order
type Manifest struct {
	AppID               [16]byte // Application ID
	Domain              string
	Email               string
	Icon                string
	AuthorizedAppDomain []string // Just accept request from these domains, neither guest nor users!
	SupportedLanguages  []uint32
	ManifestLanguages   []uint32
	Organization        []string
	Name                []string
	Description         []string
	TermsOfService      []string
	Licence             []string
	TAGS                []string // Use to categorized apps e.g. Music, GPS, ...
	RequestedPermission []uint32 // ServiceIDs from PersiaOS services e.g. InternetInBackground, Notification, ...
	TechnicalInfo       TechnicalInfo
}

// TechnicalInfo store some technical information but may different from really server condition!
type TechnicalInfo struct {
	UseAI               bool   // false: Open the lot of security concerns || true: use more resource.
	AuthorizationServer string // Domain name that have sRPC needed store connection data. default is "sabz.city"

	// Shutdown settings
	ShutdownDelay time.Duration // the server will wait for at least this amount of time for active streams to finish!

	// Server Overal rate limit
	MaxOpenConnection     uint64 // The maximum number of concurrent connections the app may serve.
	ConnectionIdleTimeout time.Duration
	MaxStreamHeaderSize   uint64 // For stream protocols with variable header size like HTTP
	MaxStreamPayloadSize  uint64 // For stream protocols with variable payload size like sRPC, HTTP, ...

	// Guest rete limit - Connection.OwnerType==0
	GuestMaxConnections            uint64 // 0 means App not accept guest connection.
	GuestMaxUserConnectionsPerAddr uint64
	GuestMaxConcurrentStreams      uint32
	GuestMaxStreamConnectionDaily  uint32 // Max open stream per day for a guest connection. overflow will drop on creation!
	GuestMaxServiceCallDaily       uint64 // 0 means no limit and good for PayAsGo strategy!
	GuestMaxBytesSendDaily         uint64
	GuestMaxBytesReceiveDaily      uint64
	GuestMaxPacketsSendDaily       uint64
	GuestMaxPacketsReceiveDaily    uint64

	// Registered rate limit - Connection.OwnerType==1
	RegisteredMaxConnections            uint64
	RegisteredMaxUserConnectionsPerAddr uint64
	RegisteredMaxConcurrentStreams      uint32
	RegisteredMaxStreamConnectionDaily  uint32 // Max open stream per day for a Registered user connection. overflow will drop on creation!
	RegisteredMaxServiceCallDaily       uint64 // 0 means no limit and good for PayAsGo strategy!
	RegisteredMaxBytesSendDaily         uint64
	RegisteredMaxBytesReceiveDaily      uint64
	RegisteredMaxPacketsSendDaily       uint64
	RegisteredMaxPacketsReceiveDaily    uint64

	// If you want to know Connection.OwnerType>1 rate limit strategy, You must read server codes!!

	// Minimum hardware specification for each instance of application.
	CPU              uint64 // Hz
	RAM              uint64 // Byte
	GPU              uint64 // Hz
	NetworkBandwidth uint64 // Byte
	HDD              uint64 // Byte, Hard disk drive as storage engine. Act as object+block storage.
	SSD              uint64 // Byte, Solid state drive as storage engine. Act as object+block storage..
	// 4 type of devices trend now:
	// Out of society scope - Cloud Computing - Cheapest scenario - Primary Datastore	::/1
	// Adjacent Router or Inside same society - Fog Computing - Caching Datastore		::/32
	// Inside Router scope - Edge Computing - Caching Datastore						::/64
	// End user device - 															::/128
	MaxNetworkScalability uint8 // If ::/32 it mean also ::/1 but not ::/64!!

	// DataStore
	DataCenterClassForDataStore uint64 //
	ReplicationNumber           uint64 // deafult is 3
	NodeNumber                  uint64 // default is 3
}
