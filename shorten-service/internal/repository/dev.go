package repository

import (
    "errors"
    "sync"
	"time"

    "shorten-service/internal/util"
)

var ErrNotFound = errors.New("not found")

type DevRepo struct {
    mutex   sync.RWMutex
    counter uint64
    store   map[string]Link
}

func NewDevRepo() *DevRepo {
    return &DevRepo{
        store: make(map[string]Link),
    }
}

func (repo *DevRepo) Save(originalURL string) string {
    repo.mutex.Lock()
    defer repo.mutex.Unlock()

    repo.counter++
    code := util.EncodeBase62(repo.counter)
    link := Link{
        Code:        code,
        OriginalURL: originalURL,
        CreatedAt:   time.Now(),
    }
    repo.store[link.Code] = link

    return code
}

func (repo *DevRepo) FindByCode(code string) (*Link, error) {
    repo.mutex.RLock()
    defer repo.mutex.RUnlock()

    link, ok := repo.store[code]
    if !ok {
        return nil, ErrNotFound
    }
    return &link, nil
}
