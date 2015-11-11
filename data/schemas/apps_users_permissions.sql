DROP TABLE IF EXISTS apps_users_permissions;

CREATE TABLE apps_users_permissions (
    app_id bigint NOT NULL,
    user_id bigint NOT NULL,
    permissions int NOT NULL
);

ALTER TABLE apps_users_permissions ADD FOREIGN KEY ("app_id") REFERENCES apps("id") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE apps_users_permissions ADD FOREIGN KEY ("user_id") REFERENCES users("id") ON DELETE CASCADE ON UPDATE CASCADE;
