package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// UserSession represents the JSON object which is saved under the session in our cache
type UserSession struct {
	UserID    int32  `json:"userId"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// SessionManager declares the interface for managing user sessions
type SessionManager interface {

	// SetSession saves the current user session details into cache, under the sessionId key.
	SetSession(ctx context.Context, sessionId string, sessionDetails UserSession) error

	// GetSession queries the cache for a session stored under a given sessionId
	// and parses the stored session JSON back into a usable struct
	GetSession(ctx context.Context, sessionId string) (*UserSession, error)

	// ClearSession deletes a session saved in cache under a given sessionId key
	ClearSession(ctx context.Context, sessionId string) error
}

// RedisSessionManager implements SessionManager, with Redis as the store
type RedisSessionManager struct {
	Client *redis.Client
}

func (sessionManager RedisSessionManager) SetSession(ctx context.Context, sessionId string, sessionDetails UserSession) error {
	sessionString, err := json.Marshal(sessionDetails)
	if err != nil {
		return err
	}

	cacheKey := fmt.Sprintf("session:%v", sessionId)
	err = sessionManager.Client.Set(ctx, cacheKey, sessionString, 10*time.Minute).Err()

	if err != nil {
		return err
	}

	return nil
}

func (sessionManager RedisSessionManager) GetSession(ctx context.Context, sessionId string) (*UserSession, error) {
	sessionDetails := UserSession{}
	cacheKey := fmt.Sprintf("session:%v", sessionId)

	value, err := sessionManager.Client.Get(ctx, cacheKey).Result()

	if err != nil {
		return nil, err
	}

	// unmarshall redis JSON object into a usable struct
	err = json.Unmarshal([]byte(value), &sessionDetails)
	if err != nil {
		return nil, err
	}

	return &sessionDetails, nil

}

func (sessionManager RedisSessionManager) ClearSession(ctx context.Context, sessionId string) error {
	cacheKey := fmt.Sprintf("session:%v", sessionId)
	err := sessionManager.Client.Del(ctx, cacheKey).Err()
	return err
}
