package domain

import (
	"fmt"
	"sort"
)

type EntryCollection []Entry

type Group struct {
	Date    string  `json:"date"`
	Entries []Entry `json:"entries"`
}

func (e EntryCollection) Print() {
	fmt.Println(e.String())
}

func (e EntryCollection) String() string {
	var out = ""
	_groupsMaxSize := 3
	_archiveMaxSize := 10

	groupsPrinted := make(map[string]bool)
	archivePrintCount := 0
	archivePrinted := false
	sortedEntries := e.Sorted()
	for _, entry := range sortedEntries {
		maxGroupsReached := len(groupsPrinted) >= _groupsMaxSize
		if !archivePrinted && maxGroupsReached && !groupsPrinted[entry.Date] {
			archivePrinted = true
			groupsPrinted[entry.GroupKey()] = true
			out += "\n" + ArchiveHeader() + "\n"
		}
		if !groupsPrinted[entry.GroupKey()] && !archivePrinted {
			groupsPrinted[entry.GroupKey()] = true
			out += "\n" + DateHeader(entry.GroupKey()) + "\n"
		}
		if len(groupsPrinted) <= _groupsMaxSize {
			out += EntryText(entry) + "\n"
		} else if archivePrintCount < _archiveMaxSize {
			archivePrintCount++

			out += EntryText(entry) + "\n"
		}
	}
	return out
}

func (e EntryCollection) Sorted() EntryCollection {
	sort.Slice(e, func(i, j int) bool {
		// Sort by Entry.Date in descending order
		return e[i].Date > e[j].Date
	})
	return e
}
