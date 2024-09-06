-- +goose Up
CREATE TABLE tokens (
    id varchar(36) NOT NULL,
    created_at timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    updated_at timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    user_id varchar(255) NOT NULL,
    session varchar(255) NOT NULL,
    refresh_token text NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB;

ALTER TABLE tokens ADD CONSTRAINT FK_tokens_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE NO ACTION ON UPDATE NO ACTION;

-- +goose Down
ALTER TABLE tokens DROP FOREIGN KEY FK_tokens_user;

DROP TABLE tokens;
