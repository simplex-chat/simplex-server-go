CREATE TABLE uris (
  uri text PRIMARY KEY
);

/* here and further "ON <operation> RESTRICT" options make foreign key checks non deferrable,
 * so that one cannot change data even if application stack is hypothetically compromised
 */
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
  id uuid NOT NULL,
  ts timestamp NOT NULL,
  msg bytea NOT NULL,
  PRIMARY KEY (recipient_uri, id)
);
