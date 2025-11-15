package db

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Repository interface {
	CreateTicket(ticket *Ticket) error
	GetTicketByID(id string) (*Ticket, error)
	UpdateTicketStatus(id string, status string) error

	CreateMessage(message *Message) error
	GetMessageByID(id string) (*Message, error)
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
