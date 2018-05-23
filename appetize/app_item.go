package appetize

import (
	"time"
)

type AppItem struct {
	PublicKey   string    `json:publicKey`
	PrivateKey  string    `json:privateKey`
	Created     time.Time `json:created`
	Updated     time.Time `json:updated`
	Platform    string    `json:platform`
	VersionCode int       `json:versionCode`
}
