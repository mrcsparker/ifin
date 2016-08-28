package model

type User struct {
	Id            string   `json:"id"`            // The id of the user.
	Dn            string   `json:"dn"`            // The dn of the user.
	UserName      string   `json:"userName"`      // The username. If it could not be extracted from the DN, this value will be the entire DN.
	UserGroup     string   `json:"userGroup"`     // The group this user belongs to.
	Justification string   `json:"justification"` // The justification for the user account.
	Creation      string   `json:"creation"`      // The timestamp when the user was created.
	Status        string   `json:"status"`        // The user status.
	LastVerified  string   `json:"lastVerified"`  // The timestamp the user authorities were verified.
	LastAccessed  string   `json:"lastAccessed"`  // The timestamp the user last accessed the system.
	Authorities   []string `json:"authorities"`   // The users authorities.
}
