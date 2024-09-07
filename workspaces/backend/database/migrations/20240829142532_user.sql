-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id                varchar(36)                            NOT NULL,
    created_at        timestamp(6)                           NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    updated_at        timestamp(6)                           NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    name              varchar(255)                           NOT NULL,
    email             varchar(255)                           NOT NULL,
    password          varchar(255)                           NULL,
    status            enum ('active', 'inActive', 'request') NOT NULL DEFAULT 'inActive',
    social_provider   enum ('google', 'linkedin')            NULL,
    image             text                                   NULL,
    verify            tinyint                                NOT NULL DEFAULT 0,
    verification_code varchar(6)                             NULL,
    UNIQUE INDEX IDX_user_email (email),
    PRIMARY KEY (id)
) ENGINE = InnoDB;
-- +goose StatementEnd

-- +goose Down
DROP INDEX IDX_user_email ON users;
DROP TABLE IF EXISTS users;
