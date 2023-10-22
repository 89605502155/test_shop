-- +goose Up
-- +goose StatementBegin
INSERT INTO "warehouse" (name) VALUES ('А');
INSERT INTO "warehouse" (name) VALUES ('Б');
INSERT INTO "warehouse" (name) VALUES ('З');
INSERT INTO "warehouse" (name) VALUES ('В');
INSERT INTO "warehouse" (name) VALUES ('Ж');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
