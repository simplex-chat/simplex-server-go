CREATE TABLE ids (
  id text PRIMARY KEY
);

CREATE TABLE connections (
  recipient_id text PRIMARY KEY
    REFERENCES ids (id)
    ON DELETE CASCADE
    ON UPDATE RESTRICT,
  sender_id text NOT NULL
    REFERENCES ids (id)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT,
  recipient_key bytea NOT NULL,
  sender_key bytea
);

CREATE TABLE messages (
  recipient_id text NOT NULL
    REFERENCES connections
    ON DELETE CASCADE
    ON UPDATE RESTRICT,
  id uuid NOT NULL,
  ts timestamp NOT NULL,
  msg bytea NOT NULL,
  PRIMARY KEY (recipient_id, id)
);
