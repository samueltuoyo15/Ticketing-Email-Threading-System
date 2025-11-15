CREATE TABLE IF NOT EXISTS messages (
  id UUID PRIMARY KEY AUTOINCREMENT,
  ticket_id UUID NOT NULL,
  sender_email TEXT NOT NULL,
  recipient_email TEXT NOT NULL,
  subject TEXT NOT NULL,
  body TEXT NOT NULL,
  message_id TEXT UNIQUE NOT NULL,
  in_reply_to TEXT,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (ticket_id) REFERENCES tickets(id)
)
