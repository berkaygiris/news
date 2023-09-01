package domain

import "fmt"

func ArchiveHeader() string {
	return "## 💾 Archive\nMax 10 elements are listed here."
}

func DateHeader(date string) string {
	return "## 📆 " + date
}

func EntryText(e Entry) string {
	publisher := func() string {
		if e.Publisher == "" {
			return ""
		}
		return fmt.Sprintf(" (%s)", e.Publisher)
	}()
	name := fmt.Sprintf("%s%s", e.Title, publisher)
	itemFormat := "- [%s](%s) - %s `%s`"
	commentFormat := "  > %s"

	item := fmt.Sprintf(itemFormat, name, e.URL, e.Date, e.Type)
	comment := fmt.Sprintf(commentFormat, e.Comment)
	return item + "\n" + comment
}
