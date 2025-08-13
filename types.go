package main

type CopilotToken struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
	Endpoints struct {
		API string `json:"api"`
	} `json:"endpoints"`
}

type QuotaSnapshots struct {
	PremiumInteractions *struct {
		Entitlement      int     `json:"entitlement"`
		Remaining        int     `json:"remaining"`
		PercentRemaining float64 `json:"percent_remaining"`
		Unlimited        bool    `json:"unlimited"`
		OveragePermitted bool    `json:"overage_permitted"`
	} `json:"premium_interactions"`
	Chat *struct {
		Entitlement int  `json:"entitlement"`
		Remaining   int  `json:"remaining"`
		Unlimited   bool `json:"unlimited"`
	} `json:"chat"`
	Completions *struct {
		Entitlement int  `json:"entitlement"`
		Remaining   int  `json:"remaining"`
		Unlimited   bool `json:"unlimited"`
	} `json:"completions"`
}

type CopilotStats struct {
	QuotaSnapshots QuotaSnapshots `json:"quota_snapshots"`
	QuotaResetDate string         `json:"quota_reset_date"`
}
