package repository

import (
    "errors"
    "sync"
	"time"

    "shorten-service/internal/model"
    "shorten-service/internal/util"
)

var notFoundError = errors.New("not found")

type MemoryRepo struct {
    mutex   sync.RWMutex
    counter uint64
    store   map[string]model.Link
}

func NewMemoryRepo() *MemoryRepo {
    return &MemoryRepo{
        store: make(map[string]model.Link),
    }
}

func (repo *MemoryRepo) Save(originalURL string) string {
    repo.mutex.Lock()
    defer repo.mutex.Unlock()

    repo.counter++
    code := util.EncodeBase62(repo.counter)
    link := model.Link{
        Code:        code,
        OriginalURL: originalURL,
        CreatedAt:   time.Now(),
    }
    repo.store[link.Code] = link

    return code
}

func (repo *MemoryRepo) FindByCode(code string) (*model.Link, error) {
    repo.mutex.RLock()
    defer repo.mutex.RUnlock()

    link, ok := repo.store[code]
    if !ok {
        return nil, notFoundError
    }
    return &link, nil
}
