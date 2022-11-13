package session

import (
	"context"
	"encoding/json"
	"errors"
	"math/rand"
	"net"
	"time"

	"github.com/Astemirdum/transactions/tx-user/config"
	"github.com/go-redis/redis/v8"
)

type Session struct {
	Login  string
	UserID int
}

type ID struct {
	ID string
}

const (
	sessKeyLen = 10
	sessionKey = "sessions_"
)

type Manager struct {
	client     *redis.Client
	sessionTTL time.Duration
}

func NewManager(ctx context.Context, cfg config.Redis) (*Manager, error) {
	client := redis.NewClient(
		&redis.Options{
			Addr:     net.JoinHostPort(cfg.Host, cfg.Port),
			Password: cfg.Password,
		})
	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, err
	}
	return &Manager{
		client:     client,
		sessionTTL: cfg.SessionTTL,
	}, nil
}

func (sm *Manager) Create(ctx context.Context, in *Session) (*ID, error) {
	id := ID{randStringRunes(sessKeyLen)}
	dataSerialized, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	key := sessionKey + id.ID
	res := sm.client.Set(ctx, key, dataSerialized, sm.sessionTTL)
	if err := res.Err(); err != nil {
		return nil, err
	}
	return &id, nil
}

var ErrSessionNotFound = errors.New("sessionID not found")

func (sm *Manager) Check(ctx context.Context, in *ID) (*Session, error) {
	key := sessionKey + in.ID
	res := sm.client.Get(ctx, key)

	if err := res.Err(); err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, ErrSessionNotFound
		}
		return nil, err
	}
	data, err := res.Bytes()
	if err != nil {
		return nil, err
	}
	var sess Session
	err = json.Unmarshal(data, &sess)
	if err != nil {
		return nil, err
	}
	return &sess, nil
}

func (sm *Manager) Delete(ctx context.Context, in *ID) error {
	key := sessionKey + in.ID
	return sm.client.Del(ctx, key).Err()
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))] //nolint:gosec
	}
	return string(b)
}
