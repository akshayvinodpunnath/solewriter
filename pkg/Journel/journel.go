package Journel

import (
	"errors"
	"fmt"
)

var journalCounter = 0

type Journal struct {
	EntryNumber int
	Entry       string
}

func (j Journal) String() string {
	return fmt.Sprintf("%d;%s", j.EntryNumber, j.Entry)
}

type Manager struct {
	Entries []Journal
}

func (jb *Manager) NewJournalEntry(entry string) {
	journalCounter++
	journal := Journal{
		EntryNumber: journalCounter,
		Entry:       entry,
	}
	jb.Entries = append(jb.Entries, journal)
}

func (jb *Manager) RemoveJournalEntry(index int) error {
	if index < 0 || index >= len(jb.Entries) {
		return errors.New("invalid index: out of range")
	}
	jb.Entries = append(jb.Entries[:index], jb.Entries[index+1:]...)
	return nil
}

func (jb *Manager) GetJournalEntry(index int) (Journal, error) {
	if index < 0 || index >= len(jb.Entries) {
		return Journal{}, errors.New("invalid index: out of range")
	}
	return jb.Entries[index], nil
}
