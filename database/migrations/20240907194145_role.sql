-- +goose Up
CREATE TABLE roles
(
    id          varchar(36)  NOT NULL,
    created_at  timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    updated_at  timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    name        varchar(255) NOT NULL,
    description text         NULL,
    can_modify  tinyint      NOT NULL DEFAULT 1,
    UNIQUE INDEX IDX_role_name (name),
    PRIMARY KEY (id)
) ENGINE = InnoDB;

CREATE TABLE roles_permissions
(
    role_id       varchar(36) NOT NULL,
    permission_id varchar(36) NOT NULL,
    INDEX IDX_role_id_role (role_id),
    INDEX IDX_permission_id_permission (permission_id),
    PRIMARY KEY (role_id, permission_id)
) ENGINE = InnoDB;

ALTER TABLE users
    ADD role_id varchar(36) NULL;
ALTER TABLE users
    ADD CONSTRAINT FK_users_roles FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE roles_permissions
    ADD CONSTRAINT FK_roles_permissions_roles FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE roles_permissions
    ADD CONSTRAINT FK_roles_permissions_permissions FOREIGN KEY (permission_id) REFERENCES permissions (id) ON DELETE CASCADE ON UPDATE CASCADE;

-- +goose Down
ALTER TABLE roles_permissions
    DROP FOREIGN KEY FK_roles_permissions_roles;
ALTER TABLE roles_permissions
    DROP FOREIGN KEY FK_roles_permissions_permissions;
ALTER TABLE users
    DROP FOREIGN KEY FK_users_roles;
ALTER TABLE users
    DROP COLUMN role_id;
DROP INDEX IDX_permission_id_permission ON roles_permissions;
DROP INDEX IDX_role_id_role ON roles_permissions;
DROP TABLE roles_permissions;
DROP INDEX IDX_role_name ON roles;
DROP TABLE roles;

