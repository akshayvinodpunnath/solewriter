package main

import (
	"fmt"
	"github.com/akshayvinodpunnath/solewriter/pkg/Blog"
	"github.com/akshayvinodpunnath/solewriter/pkg/ByteOperation"
	"github.com/akshayvinodpunnath/solewriter/pkg/FileOperations"
	"github.com/akshayvinodpunnath/solewriter/pkg/Journel"
	"github.com/akshayvinodpunnath/solewriter/pkg/Log"
)

func main() {
	blogManager := &Blog.Manager{}
	blogManager.NewBlogEntry("Power trains", "ICE, HEV, BEV")
	blogManager.NewBlogEntry("Engine configurations", "V, Inline, InLine Six, Twin-Cylinder, Boxer, W")
	blogEntry, err := blogManager.GetBlogEntry(0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(blogEntry)

	journalManager := &Journel.Manager{}
	journalManager.NewJournalEntry("Today I watched cricket")
	journalManager.NewJournalEntry("Today i tried to learn go")
	journalManager.NewJournalEntry("Today we had good food")

	journalEntry, err := journalManager.GetJournalEntry(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(journalEntry)

	logManager := Log.Manager{}
	logManager.AddLogEntry("null point exception", "Workspace Services", Log.LevelError)
	logManager.AddLogEntry("Performance log - 3sec", "Upgrade services", Log.LevelInfo)
	logManager.AddLogEntry("Node Def missing", "NodeFacade", Log.LevelDebug)

	logEntry, err := logManager.GetLogEntry(3)
	if err != nil {
		fmt.Println(err)
	}
	logEntry, err = logManager.GetLogEntry(2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(logEntry)

	byteArray := ByteOperation.ToByteArray(logManager.Entries, "::", "===\n")
	logFileOps := FileOperations.FileOperation{
		Input:    byteArray,
		FileName: "LogEntries",
		Path:     "./output/log",
	}
	err = logFileOps.SaveFile()
	if err != nil {
		fmt.Println(err)
	}

	byteArray = ByteOperation.ToByteArray(blogManager.Entries, ";", "\n\n")
	blogFileOps := FileOperations.FileOperation{
		Input:    byteArray,
		FileName: "BlogEntries",
		Path:     "./output/log",
	}
	err = blogFileOps.SaveFile()
	if err != nil {
		fmt.Println(err)
	}

	byteArray = ByteOperation.ToByteArray(journalManager.Entries, "*", "=\n")
	journalFileOps := FileOperations.FileOperation{
		Input:    byteArray,
		FileName: "JournalEntries",
		Path:     "./output/log",
	}
	err = journalFileOps.SaveFile()
	if err != nil {
		fmt.Println(err)
	}

	readFileJournal := FileOperations.FileOperation{
		Input:    byteArray,
		FileName: "JournalEntries",
		Path:     "./output/log",
	}
	journalData, err := readFileJournal.ReadFile()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ByteOperation.ToArray(journalData))

	invalidFileJournal := FileOperations.FileOperation{
		Input:    byteArray,
		FileName: "JournalEntries1",
		Path:     "./output/log",
	}
	_, err = invalidFileJournal.ReadFile()
	if err != nil {
		fmt.Println(err)
	}
}
