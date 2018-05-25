package appetize

import (
	"time"
)

// create/upload と list/show/update とで共通する部分をここに定義
type AppItem struct {
	PublicKey   string    `json:publicKey`
	PrivateKey  string    `json:privateKey`
	Created     time.Time `json:created`
	Updated     time.Time `json:updated`
	Platform    string    `json:platform`
	VersionCode int       `json:versionCode`
	Email       string    `json:email`

	// optional...
	LaunchUrl string `json:launchUrl`
	Timeout   int    `json:timeout`
	Disabled  bool   `json:disabled`
	Note      string `json:note`

	// not implemented yet...
	//AppPermissions       `json:appPermissions`
	//Architectures        `json:architectures`
}

// create/upload のレスポンス
type CreatedAppItem struct {
	*AppItem

	PublicUrl string `json:publicURL`
	AppUrl    string `json:appURL`
	ManageUrl string `json:manageURL`
}

// list/show/update のレスポンス
type DetailedAppItem struct {
	*AppItem

	AppVersionCode string `json:appVersionCode`
	AppVersionName string `json:appVersionName`
	Bundle         string `json:bundle`
	IconUrl        string `json:iconUrl`
	Name           string `json:name`
}
