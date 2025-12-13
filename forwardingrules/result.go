package forwardingrules

type ForwardingRule struct {
	ID   string `json:"id,omitempty"`
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

type ListForwardingRulesResult struct {
	Links ForwardingRulesResponseLink `json:"links"`
	Rules []ForwardingRule             `json:"values"`
	Count int                          `json:"count"`
}

type ForwardingRulesResponseLink struct {
	Next string `json:"next"`
}

