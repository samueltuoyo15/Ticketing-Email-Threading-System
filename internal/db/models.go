package db

import (
	"time"
)

type Ticket struct {
	ID        string    `db:"id"`
	Subject   string    `db:"subject"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Message struct {
	ID             string    `db:"id"`
	TicketID       string    `db:"ticket_id"`
	SenderEmail    string    `db:"sender_email"`
	RecipientEmail string    `db:"recipient_email"`
	Subject        string    `db:"subject"`
	Body           string    `db:"body"`
	MessageID      string    `db:"message_id"`
	InReplyTo      *string   `db:"in_reply_to"`
	CreatedAt      time.Time `db:"created_at"`
}
