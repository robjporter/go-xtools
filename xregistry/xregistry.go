package xregsitry

import (
	"time"
	"sync"
)

var secondsAfter time.Duration = 1

type value struct {
	expiry  time.Time
	access  time.Time
	val     interface{}
	expires time.Duration
}

type Xregistry struct {
	store        map[interface{}]*value
	Expiry       time.Duration
	cleanupTimer *time.Timer
	cleanupLock  sync.Mutex
	cacheLock    sync.RWMutex
	OnExpired    func(key interface{}, val interface{})
}

func New(expiry ...time.Duration) *Xregistry {
	c := &Xregistry{store: make(map[interface{}]*value)}
	if expiry != nil {
		c.Expiry = expiry[0]
	} else {
		c.Expiry = 1<<63 - 1
	}
	return c
}

func (s *Xregistry) Get(key interface{}) interface{} {
	s.cacheLock.RLock()
	valueHolder := s.store[key]
	s.cacheLock.RUnlock()
	if valueHolder != nil {
		n := time.Now()
		valueHolder.access = n
		return valueHolder.val

	}
	return nil
}

func (s *Xregistry) GetReset(key interface{}) interface{} {
	s.cacheLock.RLock()
	valueHolder := s.store[key]
	s.cacheLock.RUnlock()
	if valueHolder != nil {
		n := time.Now()
		valueHolder.access = n
		valueHolder.expiry = valueHolder.access.Add(valueHolder.expires)
		return valueHolder.val

	}
	return nil
}

func (s *Xregistry) Remove(key interface{}) bool {
	s.cacheLock.Lock()
	session := s.store[key]
	if session != nil {
		delete(s.store, key)
		s.cacheLock.Unlock()
		return true
	} else {
		s.cacheLock.Unlock()
		return false
	}
}

func (s *Xregistry) SetCustom(key interface{}, val interface{}, expiry time.Duration) {
	n := time.Now()
	exp := n.Add(expiry)
	session := &value{expiry: exp, access: n, val: val, expires: expiry}
	s.cacheLock.Lock()
	s.store[key] = session
	s.cacheLock.Unlock()
	s.startCleanup(expiry + (secondsAfter * time.Second))
}

func (s *Xregistry) Set(key interface{}, val interface{}) {
	n := time.Now()
	expiry := n.Add(s.Expiry)
	session := &value{expiry: expiry, access: n, val: val}
	s.cacheLock.Lock()
	s.store[key] = session
	s.cacheLock.Unlock()
	s.startCleanup(s.Expiry + (secondsAfter * time.Second))
}

func (s *Xregistry) cleanupScheduler() {
	n := time.Now()
	s.cleanupLock.Lock()
	var minExpiry time.Time = n.Add(s.Expiry)
	expiredSessions := make(map[interface{}]*value)
	s.cacheLock.Lock()
	var key interface{}
	var val *value
	for key = range s.store {
		val = s.store[key]
		if n.After(val.expiry) {
			expiredSessions[key] = val
		} else if val.expiry.Before(minExpiry) {
			minExpiry = val.expiry
		}
	}
	for key = range expiredSessions {
		if s.OnExpired != nil {
			s.OnExpired(key, expiredSessions[key].val);
		}
		delete(s.store, key)
	}
	s.cacheLock.Unlock()
	for key = range expiredSessions {
		val = expiredSessions[key]
	}
	sessionsLength := len(s.store)
	if sessionsLength > 0 {
		nextRunIn := minExpiry.Sub(n) + (secondsAfter * time.Second)
		s.cleanupTimer = time.AfterFunc(nextRunIn, s.cleanupScheduler)
	} else {
		s.cleanupTimer = nil
	}
	s.cleanupLock.Unlock()
}

func (s *Xregistry) startCleanup(runAfter time.Duration) {
	if s.cleanupTimer == nil {
		s.cleanupLock.Lock()
		if s.cleanupTimer == nil {
			s.cleanupTimer = time.AfterFunc(runAfter, s.cleanupScheduler)
		}
		s.cleanupLock.Unlock()
	}
}

func (s *Xregistry) stopCleanup() {
	s.cleanupLock.Lock()
	if s.cleanupTimer != nil {
		s.cleanupTimer.Stop()
		s.cleanupTimer = nil
	}
	s.cleanupLock.Unlock()
}

func (s *Xregistry) Close() {
	s.stopCleanup()
}
