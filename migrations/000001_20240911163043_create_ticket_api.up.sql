-- +goose Up
-- +goose StatementBegin
CREATE TABLE shows
(
    id         serial                                 NOT NULL PRIMARY KEY,
    name       text                                   NOT NULL UNIQUE,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone
);

CREATE TABLE events
(
    id         serial                                 NOT NULL PRIMARY KEY,
    show_id    integer,
    date       timestamp with time zone               NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone,

    FOREIGN KEY (show_id) REFERENCES shows (id)
);

CREATE TABLE places
(
    id           serial  NOT NULL PRIMARY KEY,
    x            float   NOT NULL,
    y            float   NOT NULL,
    width        float   NOT NULL,
    height       float   NOT NULL,
    is_available boolean NOT NULL         DEFAULT true,
    created_at   timestamp with time zone DEFAULT now() NOT NULL,
    updated_at   timestamp with time zone
);
-- +goose StatementEnd
