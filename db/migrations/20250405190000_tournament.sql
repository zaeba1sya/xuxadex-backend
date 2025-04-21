-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tournament_status_entity (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    name VARCHAR(255) NOT NULL,
    can_join BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT now (),
    updated_at TIMESTAMP NOT NULL DEFAULT now (),
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS status_name_idx ON tournament_status_entity (name);

CREATE TRIGGER update_modified_time BEFORE
UPDATE ON tournament_status_entity FOR EACH ROW EXECUTE PROCEDURE update_modified_column ();

INSERT INTO
    tournament_status_entity (id, name, can_join)
VALUES
    (
        '2c60ea0a-f337-4752-8556-b630e6f6cd32',
        'PREPARATION',
        TRUE
    ),
    (
        'd4114c4b-ebd8-4aca-84eb-dafa5ca474f7',
        'IN_PROGRESS',
        FALSE
    ),
    (
        '80261546-b3be-4420-af21-c6b529e4276e',
        'FINISHED',
        FALSE
    ),
    (
        '0a04a3ba-4126-4e81-b39c-d3eec7408cee',
        'CANCELLED',
        FALSE
    );

CREATE TABLE IF NOT EXISTS tournament_entity (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    title VARCHAR(255) NOT NULL,
    description TEXT DEFAULT NULL,
    creator_id UUID NOT NULL REFERENCES user_entity (id),
    entrance_fee DECIMAL(10, 2) NOT NULL DEFAULT 0,
    teams_count INTEGER NOT NULL DEFAULT 2,
    teams_size INTEGER NOT NULL DEFAULT 1,
    status_id UUID NOT NULL REFERENCES tournament_status_entity (id) DEFAULT '2c60ea0a-f337-4752-8556-b630e6f6cd32',
    start_timestamp TIMESTAMP NOT NULL,
    game_id UUID NOT NULL REFERENCES game_entity (id),
    game_mode_id UUID NOT NULL REFERENCES game_mode_entity (id),
    created_at TIMESTAMP NOT NULL DEFAULT now (),
    updated_at TIMESTAMP NOT NULL DEFAULT now (),
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TRIGGER update_modified_time BEFORE
UPDATE ON tournament_entity FOR EACH ROW EXECUTE PROCEDURE update_modified_column ();

CREATE INDEX IF NOT EXISTS tournament_idx ON tournament_entity (title, creator_id, status_id);

CREATE TABLE IF NOT EXISTS tournament_uploading_type_entity (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now (),
    updated_at TIMESTAMP NOT NULL DEFAULT now (),
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS uploading_type_name_idx ON tournament_uploading_type_entity (name);

INSERT INTO
    tournament_uploading_type_entity (id, name)
VALUES
    ('504e3166-b8db-407f-a7d1-5825413566a9', 'BANNER'),
    ('e6ffee14-ef02-40ee-9c20-af81ab9e3ca1', 'PREVIEW');

CREATE TABLE IF NOT EXISTS tournament_uploading_entity (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    tournament_id UUID NOT NULL REFERENCES tournament_entity (id),
    type UUID NOT NULL REFERENCES tournament_uploading_type_entity (id) DEFAULT 'e6ffee14-ef02-40ee-9c20-af81ab9e3ca1',
    file_name VARCHAR(255) NOT NULL,
    file_uuid UUID NOT NULL,
    extension VARCHAR(255) NOT NULL,
    size BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now (),
    updated_at TIMESTAMP NOT NULL DEFAULT now (),
    deleted_at TIMESTAMP DEFAULT NULL
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tournament_uploading_entity CASCADE;

DROP TABLE IF EXISTS tournament_uploading_type_entity CASCADE;

DROP INDEX IF EXISTS uploading_type_name_idx;

DROP TABLE IF EXISTS tournament_entity CASCADE;

DROP INDEX IF EXISTS tournament_idx;

DROP TABLE IF EXISTS tournament_status_entity CASCADE;

-- +goose StatementEnd
