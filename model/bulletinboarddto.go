package model

type BulletinBoardDTO struct {
	Bulletins []BulletinEntity `json:"bulletins"` // The bulletins in the bulletin board, that matches the supplied request.
	Generated string           `json:"generated"` // The timestamp when this report was generated.
}
