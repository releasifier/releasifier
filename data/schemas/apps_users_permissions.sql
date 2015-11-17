DROP TABLE IF EXISTS apps_users_permissions;
DROP SEQUENCE IF EXISTS apps_users_permission_id_seq;

CREATE SEQUENCE apps_users_permission_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE apps_users_permissions (
    id bigint DEFAULT nextval('apps_users_permission_id_seq'::regclass) NOT NULL,
    app_id bigint NOT NULL,
    user_id bigint NOT NULL,
    permission int NOT NULL
);

ALTER TABLE ONLY apps_users_permissions ADD CONSTRAINT apps_users_permissions_pkey PRIMARY KEY (id);
ALTER TABLE apps_users_permissions ADD FOREIGN KEY ("app_id") REFERENCES apps("id") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE apps_users_permissions ADD FOREIGN KEY ("user_id") REFERENCES users("id") ON DELETE CASCADE ON UPDATE CASCADE;
