package broker

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/Astemirdum/transactions/tx-balance/config"
	"github.com/Astemirdum/transactions/tx-balance/internal/repository"
	models "github.com/Astemirdum/transactions/tx-balance/models/v1"
	nats "github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

const (
	streamSubjects = "CASH_OUT.*"
	streamName     = "CASH_OUT"
	SubjectTmp     = "CASH_OUT.cash_out_%d"
	durConsumer    = "balance"
)

type Broker struct {
	nc      *nats.Conn
	js      nats.JetStreamContext
	log     *zap.Logger
	cashOut CashOut

	subs sync.Map
}

type CashOut func(ctx context.Context, balance *models.CashOut) (int64, error)

func (b *Broker) SetCashOutHandler(cashOut CashOut) {
	b.cashOut = cashOut
}

func NewBroker(cfg config.JS, log *zap.Logger) (*Broker, error) {
	connOpts := []nats.Option{
		nats.MaxReconnects(cfg.MaxReconnects),
		nats.ReconnectWait(cfg.ReconnectWait),
		nats.Timeout(cfg.ConnectTimeout),
	}
	nc, err := nats.Connect(cfg.URL, connOpts...)
	if err != nil {
		return nil, err
	}
	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}
	b := &Broker{
		nc:  nc,
		js:  js,
		log: log,
	}
	if err := b.createStream(streamName); err != nil {
		return nil, err
	}
	return b, nil
}

func (b *Broker) StartCashOutByUser(ctx context.Context, userID int, unSubCh chan int) error {
	dur := durConsumer + strconv.Itoa(userID)
	subject := fmt.Sprintf(SubjectTmp, userID)
	sub, err := b.js.Subscribe(subject, func(msg *nats.Msg) {
		if err := b.cashOutProcess(ctx, msg, unSubCh, userID); err != nil {
			if errors.Is(err, repository.ErrOverdraft) {
				b.log.Warn("stop sub ErrOverdraft", zap.Int("userID", userID))
				return
			}
		}
	}, nats.AckExplicit(), nats.Durable(dur))
	if err != nil {
		return err
	}
	b.subs.Store(userID, sub)
	/*
		sub, err := b.js.PullSubscribe(subject,
			dur,
			nats.AckExplicit())
		if err != nil {
			return err
		}
		b.subs.Store(userID, sub)
		const (
			fetchAwait   = time.Millisecond * 300
			connectAwait = time.Second * 5
		)
		go func() {
			for {
				select {
				case <-ctx.Done():
					b.log.Error("ctx done")
					return
				default:
					msgs, err := sub.Fetch(1, nats.MaxWait(fetchAwait))
					if err != nil {
						if !errors.Is(err, context.DeadlineExceeded) &&
							!errors.Is(err, context.Canceled) &&
							!errors.Is(err, nats.ErrTimeout) {
							b.log.Error("fetch batches", zap.Error(err))
						}
						time.Sleep(connectAwait)
						continue
					}
					for i := range msgs {
						if err := b.cashOutProcess(ctx, msgs[i], unSubCh, userID); err != nil {
							if errors.Is(err, repository.ErrOverdraft) {
								b.log.Warn("errors.Is(err, repository.ErrOverdraft)")
								return
							}
						}
					}
				}
			}
		}()
	*/
	b.log.Debug("Subscribe on user", zap.Int("userID", userID))
	return nil
}

var ErrNoUserSub = errors.New("no user sub")

func (b *Broker) UnsubscribeUser(userID int) error {
	sub, ok := b.subs.Load(userID)
	if !ok {
		return ErrNoUserSub
	}
	{
		sub, _ := sub.(*nats.Subscription) //nolint:errcheck
		if err := sub.Unsubscribe(); err != nil {
			b.log.Error("failed to Unsubscribe",
				zap.Error(err),
				zap.Int("userID", userID))
			return err
		}
		b.log.Debug("done Unsubscribe",
			zap.Int("userID", userID))
		b.subs.Delete(userID)
	}
	return nil
}

func (b *Broker) createStream(streamName string) error {
	stream, err := b.js.StreamInfo(streamName)
	if err != nil {
		b.log.Debug("StreamInfo", zap.Error(err))
	}
	if stream == nil {
		b.log.Debug("creating stream", zap.String("streamName", streamName))
		_, err = b.js.AddStream(&nats.StreamConfig{
			Name:     streamName,
			Subjects: []string{streamSubjects},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Broker) Close() {
	b.nc.Close()
}
