package configs

type Component struct {
	ID        int            `json:"id"`
	Component string         `json:"component"`
	Model     map[string]any `json:"model"`
}

var Layout = []Component{
	{
		ID:        1,
		Component: "crypto_btc",
		Model:     map[string]any{},
	},
	{
		ID:        2,
		Component: "crypto_eth",
		Model:     map[string]any{},
	},
	{
		ID:        3,
		Component: "crypto_xrp",
		Model:     map[string]any{},
	},
}
