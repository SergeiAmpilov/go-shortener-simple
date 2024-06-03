package storage

import (
	"errors"
	"sync"
)

var (
	errLinkNotFound = errors.New("Link not found")
)

type LinkStorage struct {
	storage map[string]Link
	mu      sync.Mutex
}

func (l *LinkStorage) CreateLink(internalLink string) (Link, error) {
	newUrl := "12123123"
	newLink := Link{
		Internal: internalLink,
		External: newUrl,
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	l.storage[newUrl] = newLink

	return newLink, nil
}

func (l *LinkStorage) GetLink(externalLink string) (Link, error) {

	l.mu.Lock()
	defer l.mu.Unlock()

	link, ok := l.storage[externalLink]

	if !ok {
		return Link{}, errLinkNotFound
	}

	return link, nil
}
