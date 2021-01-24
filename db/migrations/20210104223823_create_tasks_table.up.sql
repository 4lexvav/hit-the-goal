CREATE TABLE IF NOT EXISTS tasks (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(500) NOT NULL,
    description VARCHAR(5000),
    position smallint NOT NULL DEFAULT 0,

    list_id BIGINT REFERENCES lists(id) ON DELETE CASCADE,

    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX tasks_list_id_index ON tasks USING btree(list_id);

CREATE TRIGGER tasks_moddatetime
    BEFORE UPDATE ON tasks
    FOR EACH ROW
    EXECUTE PROCEDURE moddatetime(updated_at);