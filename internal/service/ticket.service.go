package services

import (
	"errors"
	"time"

	"github.com/samueltuoyo15/Ticketing-Email-Threading-System/internal/db"
)

type TicketService struct {
	repo db.Repository
}

func NewTicketService(repo db.Repository) *TicketService {
	return &TicketService{repo: repo}
}

func (service *TicketService) CreateTicket(subject string) (*db.Ticket, error) {
	if subject == "" {
		return nil, errors.New("subject required")
	}

	ticket := &db.Ticket{
		Subject: subject,
		Status:  "open",
	}

	err := service.repo.CreateTicket(ticket)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

func (service *TicketService) GetTicket(ticketID string) (*db.Ticket, error) {
	return service.repo.GetTicketByID(ticketID)
}

func (service *TicketService) CloseTicket(ticketID string) error {
	return service.repo.UpdateTicketStatus(ticketID, "closed")
}

func (service *TicketService) ReopenTicket(ticketID string) error {
	return service.repo.UpdateTicketStatus(ticketID, "open")
}

func (service *TicketService) SetPending(ticketID string) error {
	return service.repo.UpdateTicketStatus(ticketID, "pending")
}

func (service *TicketService) UpdateUpdatedAt(ticketID string) error {
	ticket, err := service.repo.GetTicketByID(ticketID)

	if err != nil {
		return err
	}

	if ticket == nil {
		return errors.New("ticket not found")
	}

	ticket.UpdatedAt = time.Now()
	return nil
}
