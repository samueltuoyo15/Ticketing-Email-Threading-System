package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type Repository interface {
	CreateTicket(ticket *Ticket) error
	GetTicketByID(ticketID string) (*Ticket, error)
	UpdateTicketStatus(ticketID string, status string) error

	CreateMessage(message *Message) error
	GetMessageByTicketID(messageID string) (*Message, error)
	GetMessagesByTicketID(ticketID string) ([]*Message, error)
}

type SqlLiteRepository struct {
	db *sql.DB
}

func NewSqlLiteRepository(db *sql.DB) *SqlLiteRepository {
	return &SqlLiteRepository{db: db}
}

// Tickets
func (repo *SqlLiteRepository) CreateTicket(ticket *Ticket) error {
	if ticket.ID == "" {
		ticket.ID = uuid.New().String()
	}

	_, err := repo.db.Exec(`INSERT INTO tickets (id, subject, status, created_at, updated_at) VALUES (?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`, ticket.ID, ticket.Subject, ticket.Status)

	if err != nil {
		return fmt.Errorf("failed to create ticket: %w", err)
	}

	return nil
}

func (repo *SqlLiteRepository) GetTicketByID(ticketID string) (*Ticket, error) {
	var ticket Ticket

	row := repo.db.QueryRow(`SELECT id, subject, status, created_at, updated_at FROM tickets WHERE id = ?`, ticketID)
	if err := row.Scan(&ticket.ID, &ticket.Subject, &ticket.Status, &ticket.CreatedAt, &ticket.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
	}
	return &ticket, nil
}

func (repo *SqlLiteRepository) UpdateTicketStatus(ticketID string, status string) error {
	_, err := repo.db.Exec(`UPDATE tickets SET status = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`, status, ticketID)
	return err
}

// Messages
func (repo *SqlLiteRepository) CreateMessage(msg *Message) error {
	_, err := repo.db.Exec(`INSERT into messages (ticket_id, sender_email, recipient_email, subject, body, in_reply_to, created_at) VALUES (?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)`, msg.TicketID, msg.SenderEmail, msg.RecipientEmail, msg.Subject, msg.Body, msg.InReplyTo)
	return err
}

func (repo *SqlLiteRepository) GetMessageByTicketID(messageID string) (*Message, error) {
	var msg Message

	row := repo.db.QueryRow(`SELECT id, ticket_id, sender_email, recipient_email, subject, body, message_id, in_reply_to, created_at FROM messages WHERE id = ?`, messageID)
	if err := row.Scan(&msg.ID, &msg.TicketID, &msg.SenderEmail, &msg.RecipientEmail, &msg.Subject, &msg.Body, &msg.MessageID, &msg.InReplyTo, &msg.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get message: %w", err)
	}
	return &msg, nil
}

func (repo *SqlLiteRepository) GetMessagesByTicketID(ticketID string) ([]*Message, error) {
	rows, err := repo.db.Query(`SELECT id, ticket_id, sender_email, recipient_email, subject, body, message_id, in_reply_to, created_at FROM messages WHERE ticket_id = ?`, ticketID)
	if err != nil {
		return nil, fmt.Errorf("failed to get messages: %w", err)
	}
	defer rows.Close()

	var messages []*Message
	for rows.Next() {
		var message Message
		if err := rows.Scan(&message.ID, &message.TicketID, &message.SenderEmail, &message.RecipientEmail, &message.Subject, &message.Body, &message.MessageID, &message.InReplyTo, &message.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan message: %w", err)
		}
		messages = append(messages, &message)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over messages: %w", err)
	}

	return messages, nil
}
