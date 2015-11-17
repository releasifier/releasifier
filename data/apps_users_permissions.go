package data

type AppsUsersPermissions struct {
	ID         int64      `db:"id,omitempty,pk" json:"_"`
	AppID      int64      `db:"app_id" json:"_"`
	UserID     int64      `db:"user_id" json:"_"`
	Permission Permission `db:"permission" json:"permission"`
}

//CollectionName returns collection name in database
func (a *AppsUsersPermissions) CollectionName() string {
	return `apps_users_permissions`
}
