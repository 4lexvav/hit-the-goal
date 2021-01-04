CREATE TABLE IF NOT EXISTS comments (
    id BIGSERIAL PRIMARY KEY,
    text VARCHAR(5000) NOT NULL,

    task_id BIGINT REFERENCES tasks(id) ON DELETE CASCADE,

    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX comments_task_id_index ON comments (task_id);

CREATE TRIGGER comments_moddatetime
    BEFORE UPDATE ON comments
    FOR EACH ROW
    EXECUTE PROCEDURE moddatetime(updated_at);