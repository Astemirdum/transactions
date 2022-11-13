package broker

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/Astemirdum/transactions/tx-balance/internal/repository"
	models "github.com/Astemirdum/transactions/tx-balance/models/v1"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

func (b *Broker) cashOutProcess(
	ctx context.Context,
	msg *nats.Msg,
	unSubCh chan int,
	userID int,
) error {

	var cashMsg models.CashOutMsg
	if err := json.Unmarshal(msg.Data, &cashMsg); err != nil {
		b.log.Error("unmarshal CashOutMsg", zap.Error(err))
		return err
	}
	remainCash, err := b.cashOut(ctx, &models.CashOut{
		UserID: userID,
		Cash:   cashMsg.Cash,
	})
	switch {
	case err == nil:
		if e := msg.AckSync(); e != nil {
			b.log.Warn("failed to ack batch", zap.Error(e))
		}
		b.log.Debug("cashOutProcess",
			zap.Int64("remaining cash", remainCash),
			zap.Int("userID", userID))
	case errors.Is(err, repository.ErrOverdraft):
		if e := msg.AckSync(); e != nil {
			b.log.Warn("failed to ack batch", zap.Error(e))
		}
		b.log.Debug("block: lack of money",
			zap.Error(err),
			zap.Int64("remaining cash", remainCash),
			zap.Int("userID", userID))

		// unsubscribe (no money available)
		unSubCh <- userID
		return nil
	default:
		if e := msg.NakWithDelay(time.Second * 3); e != nil {
			b.log.Warn("failed to nak failed batch", zap.Error(e))
		}
		b.log.Error("cashOut", zap.Error(err))
		return err
	}
	return nil
}

func (b *Broker) PublishCashOut(subjectName string, cashMsg *models.CashOutMsg) error {
	data, err := json.Marshal(cashMsg)
	if err != nil {
		return err
	}
	_, err = b.js.Publish(subjectName, data)
	return err
}
