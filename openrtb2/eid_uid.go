package openrtb2

import "encoding/json"

// Eid defines the contract for bidrequest.user.eids
// Responsible for the Universal User ID support: establishing pseudonymous IDs for users.
type Eid struct {
	Source string          `json:"source"`
	ID     string          `json:"id,omitempty"`
	Uids   []Uid           `json:"uids,omitempty"`
	Ext    json.RawMessage `json:"ext,omitempty"`
}

// Uid defines the contract for bidrequest.user.eids[i].uids[j]
type Uid struct {
	ID    string          `json:"id"`
	AType int             `json:"atype,omitempty"`
	Ext   json.RawMessage `json:"ext,omitempty"`
}
