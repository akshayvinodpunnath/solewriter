package Blog

import (
	"errors"
	"fmt"
)

var blogCounter = 0

type Blog struct {
	BlogNumber  int
	Title       string
	Description string
}

func (b Blog) String() string {
	return fmt.Sprintf("%d;%s;%s", b.BlogNumber, b.Title, b.Description)
}

type Manager struct {
	Entries []Blog
}

func (bm *Manager) NewBlogEntry(title, desc string) {
	blogCounter++
	newEntry := Blog{
		BlogNumber:  blogCounter,
		Title:       title,
		Description: desc,
	}
	bm.Entries = append(bm.Entries, newEntry)
}

func (bm *Manager) RemoveBlogEntry(index int) error {
	if index < 0 || index >= len(bm.Entries) {
		return errors.New("invalid index: out of range")
	}
	bm.Entries = append(bm.Entries[:index], bm.Entries[index+1:]...)
	return nil
}

func (bm *Manager) GetBlogEntry(index int) (Blog, error) {
	if index < 0 || index >= len(bm.Entries) {
		return Blog{}, errors.New("invalid index: out of range")
	}
	return bm.Entries[index], nil
}
