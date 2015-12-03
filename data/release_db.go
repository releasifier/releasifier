package data

import "time"

//Release structure for data represents in database
type Release struct {
	ID       int64     `db:"id,omitempty,pk" json:"id"`
	AppID    int64     `db:"app_id" json:"-"`
	Platform Platform  `db:"platform" json:"platform"`
	Note     string    `db:"note" json:"note"`
	Version  Version   `db:"version" json:"version"`
	CreateAt time.Time `db:"created_at" json:"created_at" bondb:",utc"`
	Private  bool      `db:"private" json:"private"`
}

//CollectionName returns collection name in database
func (r *Release) CollectionName() string {
	return `releases`
}
