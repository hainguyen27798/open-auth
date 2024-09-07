-- +goose Up
-- +goose StatementBegin
CREATE TABLE permissions
(
    id           varchar(36)  NOT NULL,
    created_at   timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    updated_at   timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    service_name varchar(255) NOT NULL,
    resource     varchar(255) NOT NULL,
    action       varchar(255) NOT NULL,
    attributes   varchar(255) NOT NULL,
    description  text         NULL,
    PRIMARY KEY (id),
    UNIQUE INDEX IDX_permission_key (service_name, resource, action)
) ENGINE = InnoDB;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS permissions;
-- +goose StatementEnd
