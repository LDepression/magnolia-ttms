/**
 * @Author: lenovo
 * @Description:
 * @File:  qr
 * @Version: 1.0.0
 * @Date: 2023/06/07 12:29
 */

package manager

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"mognolia/internal/dao"
	"mognolia/internal/global"
	"mognolia/internal/pkg/FuncMgr"
	"sync"
)

var QRManager *QRMap

type QRMap struct {
	l       sync.RWMutex
	Manager map[uuid.UUID]*QR
	ctx     context.Context
}

func NewQRManager() *QRMap {
	once.Do(func() {
		QRManager = &QRMap{
			l:       sync.RWMutex{},
			Manager: make(map[uuid.UUID]*QR),
			ctx:     context.Background(),
		}

	})
	return QRManager
}

func (m *QRMap) Del(orderID uuid.UUID) {
	m.l.Lock()
	defer m.l.Unlock()
	delete(m.Manager, orderID)
}

func (m *QRMap) Set(orderID uuid.UUID) bool {
	m.l.Lock()
	defer m.l.Unlock()
	if _, ok := m.Manager[orderID]; ok {
		return false
	}
	ctx, cancel := context.WithTimeout(context.Background(), global.Settings.Rule.LockOrderTime)
	m.Manager[orderID] = &QR{
		ctx:     ctx,
		ctxFunc: cancel,
	}
	go func() {
		<-ctx.Done()
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			mgr := FuncMgr.NewFuncMgr()
			if err := mgr.DeleteOutTimeOrder(dao.Group.DB, orderID); err != nil {
				zap.S().Info("failed to delete ")
				return
			}
		}
		m.Del(orderID)
	}()
	return true
}

// SetOrderWithCtx 将订单与用户锁票座位关联到一起
//func (m *QRMap) SetOrderWithCtx(planID uint, userID uint, seatIDs []uint) {
//	m.l.Lock()
//	defer m.l.Unlock()
//	//此时的用户应该是锁上了票
//	ctx, cancel := context.WithTimeout(m.ctx, global.Settings.Rule.LockTicketTime)
//
//	Manager().PlanMgr[planID].LockSeatsWithCtx(ctx, userID, seatIDs)
//
//	m.Manager[]
//
//}

type QR struct {
	ctx     context.Context
	ctxFunc context.CancelFunc
}
