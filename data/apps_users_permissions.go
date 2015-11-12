package data

type AppsUsersPermissions struct {
	AppID  SecureID `db:"id,omitempty" json:"id"`
	UserID SecureID `db:"id,omitempty" json:"id"`
}
