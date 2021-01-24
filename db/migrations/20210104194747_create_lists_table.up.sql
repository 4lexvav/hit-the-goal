CREATE TYPE list_status AS ENUM ('ACTIVE', 'INACTIVE');

CREATE TABLE IF NOT EXISTS lists (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    status list_status NOT NULL DEFAULT 'ACTIVE',
    position smallint NOT NULL DEFAULT 0,

    project_id BIGINT REFERENCES projects(id) ON DELETE CASCADE,

    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    UNIQUE(name, project_id)
);

CREATE INDEX lists_project_id_index ON lists USING btree(project_id);

CREATE TRIGGER lists_moddatetime
    BEFORE UPDATE ON lists
    FOR EACH ROW
    EXECUTE PROCEDURE moddatetime(updated_at);