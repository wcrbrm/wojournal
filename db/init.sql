CREATE TABLE notes (
  "id" BIGSERIAL PRIMARY KEY,
  "tm" BIGINT NOT NULL DEFAULT '0',
  "realm" TEXT NOT NULL DEFAULT 'default',
  "tags" text[] not null default '{}',
  "content" TEXT
);

CREATE INDEX idx_notes_tm ON notes("tm");
CREATE INDEX idx_notes_realm ON notes("realm");
CREATE INDEX idx_notes_tags ON notes USING gin("tags");

