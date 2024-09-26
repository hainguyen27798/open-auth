-- +goose Up
CREATE TABLE user_statistics
(
    user_id          varchar(36)  NOT NULL,
    created_at       timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    deleted_at       timestamp(6)          DEFAULT NULL,
    updated_at       timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    number_of_logins int,
    PRIMARY KEY (user_id)
) ENGINE = InnoDB;

ALTER TABLE user_statistics
    ADD CONSTRAINT FK_user_statistic_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE NO ACTION ON UPDATE NO ACTION;

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_statistics;
-- +goose StatementEnd
