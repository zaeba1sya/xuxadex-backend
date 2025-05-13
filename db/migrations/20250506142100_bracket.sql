-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS bracket_entity (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    tournament_id UUID NOT NULL REFERENCES tournament_entity (id),
    created_at TIMESTAMP NOT NULL DEFAULT now (),
    updated_at TIMESTAMP NOT NULL DEFAULT now (),
    deleted_at TIMESTAMP NULL DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS bracket_team_rel_entity (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    bracket_id UUID NOT NULL REFERENCES bracket_entity (id),
    team_id UUID NOT NULL REFERENCES team_entity (id),
    created_at TIMESTAMP NOT NULL DEFAULT now (),
    updated_at TIMESTAMP NOT NULL DEFAULT now (),
    deleted_at TIMESTAMP NULL DEFAULT NULL
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS bracket_team_rel_entity CASCADE;

DROP TABLE IF EXISTS bracket_entity CASCADE;

-- +goose StatementEnd
