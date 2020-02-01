CREATE TABLE uris (
  uri text PRIMARY KEY
);

CREATE TABLE connections (
  recipient_uri text PRIMARY KEY
    REFERENCES uris (uri)
    ON DELETE CASCADE
    ON UPDATE RESTRICT,
  sender_uri text NOT NULL
    REFERENCES uris (uri)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT,
  recipient_key bytea NOT NULL,
  sender_key bytea
);

CREATE TABLE messages (
  recipient_uri text NOT NULL
    REFERENCES connections
    ON DELETE CASCADE
    ON UPDATE RESTRICT,
  id uuid NOT NULL UNIQUE,
  ts timestamp NOT NULL,
  msg bytea NOT NULL
);
