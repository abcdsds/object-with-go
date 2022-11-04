package ticket

import "time"

type invitation struct {
	when time.Time
}

type ticket struct {
	Fee int64
}

type bag struct {
	Amount     int64
	Invitation *invitation
	Ticket     ticket
}

func (b bag) hasTicket() bool {
	return b.Invitation != nil
}

func (b bag) minusAmount(amount int64) {
	b.Amount -= amount
}

func (b bag) plusAmount(amount int64) {
	b.Amount += amount
}

func newBagWithAmount(amount int64) bag {
	return bag{
		Amount: amount,
	}
}

func newBagWithInvitationAndAmount(invitation invitation, amount int64) bag {
	return bag{
		Invitation: &invitation,
		Amount:     amount,
	}
}

type audience struct {
	Bag bag
}

func newAudience(bag bag) audience {
	return audience{
		Bag: bag,
	}
}

type ticketOffice struct {
	amount  int64
	tickets []ticket
}

func (o ticketOffice) getTicket() ticket {
	var ticket ticket
	ticket, o.tickets = o.tickets[0], o.tickets[1:]
	return ticket
}

func (o ticketOffice) minusAmount(amount int64) {
	o.amount -= amount
}

func (o ticketOffice) plusAmount(amount int64) {
	o.amount += amount
}

type ticketSeller struct {
	TicketOffice ticketOffice
}

type theater struct {
	TicketSeller ticketSeller
}

func (t theater) enter(audience audience) {
	if audience.Bag.hasTicket() {
		ticket := t.TicketSeller.TicketOffice.getTicket()
		audience.Bag.Ticket = ticket
	} else {
		ticket := t.TicketSeller.TicketOffice.getTicket()
		audience.Bag.minusAmount(ticket.Fee)
		t.TicketSeller.TicketOffice.plusAmount(ticket.Fee)
		audience.Bag.Ticket = ticket

	}
}
