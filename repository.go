package link

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

// iRepository provides access to storage
type iRepository interface {
	CreateOrUpdate(link *link) uuid.UUID
	Read(ids []uuid.UUID) []link
	Update(id uuid.UUID, l *link)
	Delete(id uuid.UUID)
	GetByOriginalLink(l string) (link, error)
	GetByShortLink(l string) (link, error)
	GetMaxByte() []byte
}

// repository implements IRepository interface
type repository struct{}

// CreateOrUpdate create new value if value with such ID is not exist
func (r repository) CreateOrUpdate(l *link) uuid.UUID {
	var links []link
	var fp fileProvider
	fp.Get("link_repository.json", &links)
	if len(links) == 0 {
		links = append(links, *l)
	} else {
		if r.exist(l) {
			r.Update(l.ID, l)
			return l.ID
		}
		links = append(links, *l)
	}
	fp.Save("link_repository.json", &links)
	return l.ID
}

// Read returns all links meet criteria
func (repository) Read(ids []uuid.UUID) []link {
	var links []link
	var fp fileProvider
	fp.Get("link_repository.json", &links)
	var r []link
	for _, link := range links {
		if has(ids, link.ID) {
			r = append(r, link)
		}
	}

	return r
}

func has(ids []uuid.UUID, id uuid.UUID) bool {
	for _, a := range ids {
		if a == id {
			return true
		}
	}
	return false
}

func (repository) Update(id uuid.UUID, l *link) {
	fmt.Printf("Update %v \n", id)
}

func (repository) Delete(id uuid.UUID) {
	var links []link
	var fp fileProvider
	fp.Get("link_repository.json", &links)
	for _, link := range links {
		if link.ID != id {
			links = append(links, link)
		}
	}
	fp.Save("link_repository.json", &links)
}

func (repository) DeleteAll() {
	var fp fileProvider
	fp.Save("link_repository.json", "")
}

func (repository) GetByOriginalLink(s string) (link, error) {
	var links []link
	var fp fileProvider
	fp.Get("link_repository.json", &links)
	if len(links) == 0 {
		return link{}, errors.New("link not found")
	}
	for _, l := range links {
		if l.OriginalLink == s {
			return l, nil
		}
	}
	return link{}, errors.New("link not found")
}

func (repository) GetByShortLink(s string) (link, error) {
	var links []link
	var fp fileProvider
	fp.Get("link_repository.json", &links)
	for _, l := range links {
		if l.ShortedLink == s {
			return l, nil
		}
	}
	return link{}, errors.New("link not found")
}

func (repository) GetMaxByte() []byte {
	var links []link
	var fp fileProvider
	var max []byte
	fp.Get("link_repository.json", &links)
	for _, l := range links {
		var b = []byte(l.Path)
		if bytes.Compare(b, max) > 0 {
			max = b
		}
	}
	return max
}

func (repository) exist(l *link) bool {
	var links []link
	var fp fileProvider
	fp.Get("link_repository.json", &links)
	for _, i := range links {
		if i.ID == l.ID {
			return true
		}
	}
	return false
}
