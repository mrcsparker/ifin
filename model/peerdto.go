package model

type PeerDTO struct {
	Hostname      string `json:"hostname"`      // The hostname of this peer.
	Port          int32  `json:"port"`          // The port number of this peer.
	Secure        bool   `json:"secure"`        // Returns if this peer connection is secure.
	FlowFileCount int32  `json:"flowFileCount"` // The number of flowFiles this peer holds.
}
