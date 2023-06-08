/**
 * @Author: lenovo
 * @Description:
 * @File:  ticket
 * @Version: 1.0.0
 * @Date: 2023/06/06 8:54
 */

package reply

import "mognolia/internal/model/automigrate"

type TicketInfo struct {
	SeatID       uint `gorm:"column:seat_id"`
	Row          int
	Col          int
	Price        float32
	CinemaID     uint `gorm:"column:cinema_id"`
	Status       automigrate.SeatStatus
	TicketStatus automigrate.TicketStatus `gorm:"columns:ticket_status"`
}

type ShowTicket struct {
	Tickets [][]TicketInfo
}
