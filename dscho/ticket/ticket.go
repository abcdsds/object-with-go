package ticket

import "time"

type invitation struct {
	when time.Time
}

type ticket struct {
	fee int64
}

type bag struct {
	amount     int64
	invitation *invitation
	ticket     ticket
}

func (b bag) hasTicket() bool {
	return b.invitation != nil
}

func (b bag) minusAmount(amount int64) {
	b.amount -= amount
}

func (b bag) plusAmount(amount int64) {
	b.amount += amount
}

func (b bag) hold(ticket ticket) int64 {
	if b.invitation != nil {
		b.ticket = ticket
		return 0
	} else {
		b.ticket = ticket
		b.minusAmount(ticket.fee)
		return ticket.fee
	}
}

func newBagWithAmount(amount int64) bag {
	return bag{
		amount: amount,
	}
}

func newBagWithInvitationAndAmount(invitation invitation, amount int64) bag {
	return bag{
		invitation: &invitation,
		amount:     amount,
	}
}

type audience struct {
	bag bag
}

func (a audience) buy(ticket ticket) int64 {
	return a.bag.hold(ticket)
}

func newAudience(bag bag) audience {
	return audience{
		bag: bag,
	}
}

type ticketOffice struct {
	amount  int64
	tickets []ticket
}

func (o ticketOffice) sellTicketTo(audience audience) {
	o.plusAmount(audience.buy(o.getTicket()))
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
	ticketOffice ticketOffice
}

func (s ticketSeller) sellTo(audience audience) {
	s.ticketOffice.sellTicketTo(audience)
}

type theater struct {
	ticketSeller ticketSeller
}

func (t theater) enter(audience audience) {
	t.ticketSeller.sellTo(audience)
}
