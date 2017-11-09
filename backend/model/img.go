package model

type Img struct {
	ID        int64  `db:"id" json:"id"`
	Src       []byte `db:"src" json:"src"`
	CreatedAt string `db:"createdAt" json:"createdAt"`
}
