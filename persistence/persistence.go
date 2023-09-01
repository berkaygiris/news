package persistence

import (
	"encoding/json"
	"io"
	"newsletter/domain"
	"os"
)

const Path = "data/data.json"
const SavedPath = "data/saved.json"

func Read() (domain.EntryCollection, error) {
	return readPath(Path)
}

func readPath(path string) (domain.EntryCollection, error) {
	var existingEntries []domain.Entry
	if _, err := os.Stat(path); err == nil {
		existingData, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(existingData, &existingEntries); err != nil {
			return nil, err
		}
	}
	return existingEntries, nil
}

func SaveAll(existingEntries []domain.Entry) error {
	// Save the updated data back to the JSON file
	updatedData, err := json.MarshalIndent(existingEntries, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(Path, updatedData, 0644); err != nil {
		return err
	}
	return nil
}

func Save() error {
	sourceFile, err := os.Open(Path)
	if err != nil {
		return err
	}
	defer func(sourceFile *os.File) {
		_ = sourceFile.Close()
	}(sourceFile)

	destFile, err := os.Create(SavedPath)
	if err != nil {
		return err
	}
	defer func(destFile *os.File) {
		_ = destFile.Close()
	}(destFile)

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}

func Clear() error {
	// Save the updated data back to the JSON file
	updatedData, err := json.MarshalIndent([]any{}, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(Path, updatedData, 0644); err != nil {
		return err
	}
	return nil
}

func Recover() error {

	saved, err := readPath(SavedPath)
	if err != nil {
		return err
	}

	err = SaveAll(saved)
	if err != nil {
		return err
	}
	return nil
}
