-- +goose Up
-- +goose StatementBegin

CREATE OR REPLACE FUNCTION update_modified_column()
    RETURNS TRIGGER AS $$
        BEGIN
        NEW.updated_at = now();
        RETURN NEW;
        END;
    $$ language 'plpgsql';

CREATE TABLE IF NOT EXISTS user_entity (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nickname VARCHAR(255) NOT NULL,
    avatar VARCHAR(255) NOT NULL,
    wallet VARCHAR(255) UNIQUE NOT NULL,
    steam_id VARCHAR(255) DEFAULT NULL,
    last_login TIMESTAMP NOT NULL DEFAULT now(),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP DEFAULT NULL
);
CREATE TRIGGER update_modified_time BEFORE UPDATE ON user_entity FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

CREATE INDEX IF NOT EXISTS user_wallet_idx ON user_entity(LOWER(wallet));
CREATE INDEX IF NOT EXISTS user_steam_id_idx ON user_entity(steam_id);

CREATE TABLE IF NOT EXISTS game_entity (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    icon VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP DEFAULT NULL
);
CREATE TRIGGER update_modified_time BEFORE UPDATE ON game_entity FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

INSERT INTO game_entity (id, name, icon) VALUES ('066f3ae5-8919-4a3e-8933-744767ddeef7', 'Counter-Strike 2', '/static/games/icons/cs2.svg');

CREATE TABLE IF NOT EXISTS game_mode_entity (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    game_id UUID REFERENCES game_entity(id) DEFAULT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP DEFAULT NULL
);
CREATE TRIGGER update_modified_time BEFORE UPDATE ON game_mode_entity FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

INSERT INTO game_mode_entity (id, name, game_id) VALUES ('ced861d0-6f1f-4c9c-8643-ec997f6d9fd3', 'Deathmatch', '066f3ae5-8919-4a3e-8933-744767ddeef7');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER update_modified_time ON user_entity;
DROP TABLE IF EXISTS user_entity CASCADE;
DROP INDEX IF EXISTS user_wallet_idx;
DROP INDEX IF EXISTS user_steam_id_idx;

DROP TABLE IF EXISTS game_mode_entity CASCADE;
DROP TABLE IF EXISTS game_entity CASCADE;

DROP FUNCTION IF EXISTS update_modified_column();

-- +goose StatementEnd
