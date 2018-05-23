package appetize

import (
	"time"
)

type AppItem struct {
	PublicKey   string    `json:publicKey`
	Created     time.Time `json:created`
	Updated     time.Time `json:updated`
	Platform    string    `json:platform`
	VersionCode int       `json:versionCode`
}
