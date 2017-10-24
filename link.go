package link

import "github.com/google/uuid"

const baseHRef = "http://sego.ln/"

// The link represents a link in storage
type link struct {
	ID           uuid.UUID
	OriginalLink string
	ShortedLink  string
	Path         []byte
}

//GetShort creates short link and return it
func GetShort(l string) string {
	var r iRepository = repository{}
	var existedLink, err = r.GetByOriginalLink(l)
	if err == nil {
		return existedLink.ShortedLink
	}
	shorted, b := makeShort(l)
	link := link{
		ID:           uuid.New(),
		OriginalLink: l,
		ShortedLink:  shorted,
		Path:         b,
	}
	r.CreateOrUpdate(&link)
	return shorted
}

//GetOriginal returns an original link for a short one
func GetOriginal(l string) (string, error) {
	var repository iRepository = repository{}
	var existedLink, err = repository.GetByShortLink(l)
	if err == nil {
		return existedLink.OriginalLink, nil
	}
	return "", err
}

func makeShort(l string) (string, []byte) {
	var r iRepository = repository{}
	var maxByte = r.GetMaxByte()
	maxByte = append(maxByte, 1)
	return baseHRef + string(maxByte), maxByte
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
