-- +goose Up
CREATE TABLE IF NOT EXISTS "department"(
"id" BIGSERIAL PRIMARY KEY,
"name" TEXT NOT NULL, -- CHECK (LENGTH(name) >= 1 AND LENGTH(name)<=200) --
"parent_id" BIGINT,
"created_at" TIMESTAMP NOT NULL DEFAULT now(),
FOREIGN KEY ("parent_id") REFERENCES "department"("id") ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "employee"(
"id" BIGSERIAL PRIMARY KEY,
"department_id" BIGINT NOT NULL,
"full_name" TEXT NOT NULL,
"position" TEXT NOT NULL,
"hired_at" TIMESTAMP,
"created_at" TIMESTAMP NOT NULL DEFAULT now(),
FOREIGN KEY ("department_id") REFERENCES "department"("id") ON DELETE CASCADE
);
-- +goose Down
DROP TABLE employee;
DROP TABLE department;
 