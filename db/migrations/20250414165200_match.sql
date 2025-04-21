-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS team_entity (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    name VARCHAR(255) NOT NULL,
    max_players INTEGER NOT NULL DEFAULT 1,
    min_players INTEGER NOT NULL DEFAULT 1,
    created_at TIMESTAMP NOT NULL DEFAULT now (),
    updated_at TIMESTAMP NOT NULL DEFAULT now (),
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TRIGGER update_modified_time BEFORE
UPDATE ON team_entity FOR EACH ROW EXECUTE PROCEDURE update_modified_column ();

CREATE TABLE IF NOT EXISTS match_entity (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    creator_id UUID NOT NULL REFERENCES user_entity (id),
    tournament_id UUID REFERENCES tournament_entity (id) DEFAULT NULL,
    start_time TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now (),
    updated_at TIMESTAMP NOT NULL DEFAULT now (),
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TRIGGER update_modified_time BEFORE
UPDATE ON match_entity FOR EACH ROW EXECUTE PROCEDURE update_modified_column ();

CREATE TABLE IF NOT EXISTS player_entity (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    user_id UUID NOT NULL REFERENCES user_entity (id),
    team_id UUID NOT NULL REFERENCES team_entity (id),
    created_at TIMESTAMP NOT NULL DEFAULT now (),
    updated_at TIMESTAMP NOT NULL DEFAULT now (),
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TRIGGER update_modified_time BEFORE
UPDATE ON player_entity FOR EACH ROW EXECUTE PROCEDURE update_modified_column ();

CREATE TABLE IF NOT EXISTS match_team_entity (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    match_id UUID NOT NULL REFERENCES match_entity (id),
    team_id UUID NOT NULL REFERENCES team_entity (id),
    score INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT now (),
    updated_at TIMESTAMP NOT NULL DEFAULT now (),
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TRIGGER update_modified_time BEFORE
UPDATE ON match_team_entity FOR EACH ROW EXECUTE PROCEDURE update_modified_column ();

CREATE TABLE IF NOT EXISTS player_match_stats_entity (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    player_id UUID NOT NULL REFERENCES player_entity (id),
    match_id UUID NOT NULL REFERENCES match_entity (id),
    kill INTEGER NOT NULL DEFAULT 0,
    death INTEGER NOT NULL DEFAULT 0,
    assist INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT now (),
    updated_at TIMESTAMP NOT NULL DEFAULT now (),
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TRIGGER update_modified_time BEFORE
UPDATE ON player_match_stats_entity FOR EACH ROW EXECUTE PROCEDURE update_modified_column ();

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS player_entity CASCADE;

DROP TABLE IF EXISTS player_match_stats_entity CASCADE;

DROP TABLE IF EXISTS match_team_entity CASCADE;

DROP TABLE IF EXISTS match_entity CASCADE;

DROP TABLE IF EXISTS team_entity CASCADE;

-- +goose StatementEnd
