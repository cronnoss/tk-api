-- +goose Up
-- +goose Down
-- +goose StatementBegin
DROP TABLE places;
DROP TABLE events;
DROP TABLE shows;
-- +goose StatementEnd
