package data

type AppsUsersPermissions struct {
	AppID  int64 `db:"id,omitempty" json:"id"`
	UserID int64 `db:"id,omitempty" json:"id"`
}
