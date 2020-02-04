CREATE SCHEMA server_schema;

CREATE TABLE server_schema.uris (
  uri text PRIMARY KEY
);

CREATE TABLE server_schema.connections (
  recipient_uri text PRIMARY KEY
    REFERENCES uris (uri)
    ON DELETE CASCADE,
  sender_uri text NOT NULL
    REFERENCES uris (uri),
  recipient_key bytea NOT NULL,
  sender_key bytea
);

CREATE TABLE server_schema.messages (
  recipient_uri text NOT NULL
    REFERENCES connections (recipient_uri)
    ON DELETE CASCADE,
  id uuid NOT NULL,
  ts timestamp NOT NULL,
  msg bytea NOT NULL,
  PRIMARY KEY (recipient_uri, id)
);

GRANT SELECT, INSERT ON ALL TABLES IN SCHEMA server_schema TO simplex_server;
GRANT UPDATE (sender_key) ON server_schema.connections TO simplex_server;
