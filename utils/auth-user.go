package utils

import "sync"

var tokenBlacklist = make(map[string]bool)
var blacklistLock sync.Mutex

func LogoutDriver(token string) error {
	blacklistLock.Lock()
	defer blacklistLock.Unlock()
	tokenBlacklist[token] = true
	return nil
}

func IsTokenBlacklisted(token string) (bool, error) {
	blacklistLock.Lock()
	defer blacklistLock.Unlock()
	isBlacklisted := tokenBlacklist[token]
	return isBlacklisted, nil
}
