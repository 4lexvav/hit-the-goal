CREATE TABLE IF NOT EXISTS projects (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(500) NOT NULL,
    description VARCHAR(1000),

    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TRIGGER projects_moddatetime
    BEFORE UPDATE ON projects
    FOR EACH ROW
    EXECUTE PROCEDURE moddatetime(updated_at);