package shopeepay

type Env int8
type EnvRegion int8

const (
	_ Env = iota

	// Sandbox : represent sandbox environment
	Sandbox

	// Production : represent production environment
	Production
)

const (
	_ EnvRegion = iota

	ID
	MY
	PH
	SG
)

var region = map[EnvRegion]string{
	ID: "co.id",
	MY: "com.my",
	PH: "com.ph",
	SG: "sg",
}

var typeString = map[Env]string{
	Sandbox:    "https://api.uat.wallet.airpay.",
	Production: "https://api.wallet.airpay.",
}

// implement stringer
func (e Env) String() string {
	for k, v := range typeString {
		if k == e {
			return v
		}
	}
	return "undefined"
}

// implement stringer
func (e EnvRegion) String() string {
	for k, v := range region {
		if k == e {
			return v
		}
	}
	return "undefined"
}
