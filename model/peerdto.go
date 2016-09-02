package model

type PeerDTO struct {
	Hostname      string `json:"hostname"`      // The hostname of this peer.
	Port          int    `json:"port"`          // The port number of this peer.
	Secure        bool   `json:"secure"`        // Returns if this peer connection is secure.
	FlowFileCount int    `json:"flowFileCount"` // The number of flowFiles this peer holds.
}
