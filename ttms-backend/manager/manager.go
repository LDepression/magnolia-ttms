/**
 * @Author: lenovo
 * @Description:
 * @File:  FuncMgr
 * @Version: 1.0.0
 * @Date: 2023/06/05 20:48
 */

package manager

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"mognolia/internal/dao"
	"mognolia/internal/global"
	"mognolia/internal/pkg/FuncMgr"
	"mognolia/internal/pkg/goroutine/pattern"
	"sync"
	"time"
)

var once sync.Once
var globalMgr *PlanManager

type PlanManager struct {
	l       sync.RWMutex
	PlanMgr map[uint]*plan
}

type plan struct {
	planID    uint
	l         sync.RWMutex
	ctx       context.Context
	cancel    context.CancelFunc
	ticketMgr map[uint]*ticket
}

type ticket struct {
	userID uint
	l      sync.RWMutex
	ctx    context.Context
	cancel context.CancelFunc
}

func Manager() *PlanManager {
	once.Do(func() {
		globalMgr = &PlanManager{
			l:       sync.RWMutex{},
			PlanMgr: make(map[uint]*plan)}
	})
	return globalMgr
}

func (m *PlanManager) Set(planID uint, timeout time.Duration) {
	m.l.Lock()
	defer m.l.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	m.PlanMgr[planID] = &plan{
		planID:    planID,
		l:         sync.RWMutex{},
		ctx:       ctx,
		cancel:    cancel,
		ticketMgr: make(map[uint]*ticket),
	}
}

func (m *PlanManager) Del(planIDs []uint) {
	m.l.Lock()
	defer m.l.Unlock()
	for _, planID := range planIDs {
		if m.PlanMgr[planID] != nil {
			data := m.PlanMgr[planID]
			data.cancel() //终止这个任务,下游有这个通知之后,就会终止这个任务
			delete(m.PlanMgr, planID)
		}
	}

}

// LockTicket 锁定单张票
func (p *plan) LockTicket(userID, seatID uint) bool {
	p.l.Lock()
	defer p.l.Unlock()
	if p.ticketMgr[seatID] != nil {
		return false
	}
	ctx, cancel := context.WithTimeout(context.Background(), global.Settings.Rule.LockTicketTime)

	p.ticketMgr[seatID] = &ticket{
		userID: userID,
		l:      sync.RWMutex{},
		ctx:    ctx,
		cancel: cancel,
	}
	go func() {
		<-ctx.Done()
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			carrier := FuncMgr.NewFuncMgr()
			if err := carrier.DeleteOutTimeTicket(dao.Group.DB, p.planID, seatID, userID); err != nil {
				zap.S().Info("delete out time ticket failed")
				return
			}
		}
		p.Del(seatID)
	}()
	return true
}

// LockList 锁住多个座位
func (p *plan) LockList(userID uint, seatIDs []uint) bool {
	p.l.Lock()
	defer p.l.Unlock()

	for _, seatID := range seatIDs {
		if p.ticketMgr[seatID] != nil {
			return false
		}
		ctx, cancel := context.WithTimeout(context.Background(), global.Settings.Rule.LockTicketTime)
		p.ticketMgr[seatID] = &ticket{
			userID: userID,
			l:      sync.RWMutex{},
			ctx:    ctx,
			cancel: cancel,
		}
		id := seatID
		go func() {
			<-ctx.Done()
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				carrier := FuncMgr.NewFuncMgr()
				if err := carrier.DeleteOutTimeTicket(dao.Group.DB, p.planID, id, userID); err != nil {
					zap.S().Info("delete out time ticket failed")
					return
				}
			}
			p.Del(id)
		}()
	}
	return true
}

func (p *plan) LockSeatsWithCtx(orderCtx context.Context, userID uint, seatIDs []uint) {
	p.l.Lock()
	defer p.l.Unlock()
	//自生的plan与orderID建立上了关联

	for _, seatID := range seatIDs {
		ctx := pattern.Or(p.ctx, orderCtx, p.ticketMgr[seatID].ctx)
		go func(seatID uint) {
			<-ctx.Done()
			p.Del(seatID)
		}(seatID)
	}
}

func (p *plan) Get(seatID uint) *ticket {
	p.l.RLock()
	defer p.l.RUnlock()
	if v, ok := p.ticketMgr[seatID]; ok {
		return v
	}
	return nil
}

func (p *plan) Del(seatID uint) {
	p.l.Lock()
	defer p.l.Unlock()
	delete(p.ticketMgr, seatID)
}

// GetTicketList 获取锁票的信息
func (p *plan) GetTicketList(seatIDs []uint) []*ticket {
	p.l.RLock()
	defer p.l.RUnlock()

	var res []*ticket
	for _, seatID := range seatIDs {
		if v, ok := p.ticketMgr[seatID]; ok {
			res = append(res, v)
		}
	}
	return res
}

func (t *ticket) GetUserID() uint {
	return t.userID
}
