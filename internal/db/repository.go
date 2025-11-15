package db

import (
	"database/sql"
	"fmt"
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
	result, err := repo.db.Exec(`INSERT INTO tickets (subject, status, created_at, updated_at) VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`, ticket.Subject, ticket.Status)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	ticket.ID = fmt.Sprintf("%d", id)
	return nil
}
