package services

import (
	"errors"

	"github.com/samueltuoyo15/Ticketing-Email-Threading-System/internal/db"
)

type MessagingService struct {
	repo db.Repository
}

func NewMessagingService(repo db.Repository) *MessagingService {
	return &MessagingService{repo: repo}
}

func (service *MessagingService) CreateMessage(
	ticketID string,
	sender_email string,
	recipient_email string,
	subject string,
	body string,
	inReplyTo *string,
) (*db.Message, error) {
	if sender_email == "" || recipient_email == "" || subject == "" || body == "" {
		return nil, errors.New("missing fields")
	}
	msg := &db.Message{
		TicketID:       ticketID,
		SenderEmail:    sender_email,
		RecipientEmail: recipient_email,
		Subject:        subject,
		Body:           body,
		InReplyTo:      inReplyTo,
	}
	err := service.repo.CreateMessage(msg)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
