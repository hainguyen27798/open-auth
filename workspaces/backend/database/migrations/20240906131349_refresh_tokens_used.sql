-- +goose Up
CREATE TABLE refresh_tokens_used
(
    id            varchar(36)  NOT NULL,
    created_at    timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    updated_at    timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    token_id      varchar(255) NOT NULL,
    refresh_token text         NOT NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB;

ALTER TABLE refresh_tokens_used
    ADD CONSTRAINT FK_refresh_tokens_used_token FOREIGN KEY (token_id) REFERENCES tokens (id) ON DELETE CASCADE ON UPDATE NO ACTION;

-- +goose Down
ALTER TABLE refresh_tokens_used
    DROP FOREIGN KEY FK_refresh_tokens_used_token;

DROP TABLE refresh_tokens_used;
