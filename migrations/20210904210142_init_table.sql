-- +goose Up
-- +goose StatementBegin
CREATE table reasons (
                                id bigserial primary key,
                                user_id bigint not null,
                                action_id bigint not null,
                                why varchar(50) not null
);

INSERT INTO reasons (user_id, action_id, why)
VALUES
    (1, 1, 'lost the keys'),
    (1, 2, 'want more money'),
    (5, 3, 'just because you are stupid');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE reasons;
-- +goose StatementEnd
