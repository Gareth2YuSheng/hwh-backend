package main

import (
	"database/sql"
)

func (s *PGStore) CreateTag(tag *Tag) error {
	logInfo("Running: Database - CreateTag")
	//Create Tag
	query := `INSERT INTO tags
	(tagID, name)
	values ($1, $2);`
	_, err := s.DB.Query(query,
		tag.TagID,
		tag.Name)
	if err != nil {
		return err
	}
	return nil
}

func (s *PGStore) GetAllTags() ([]*Tag, error) {
	logInfo("Running: Database - GetAllTags")
	query := `SELECT * FROM tags ORDER BY name ASC;`
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	tags := []*Tag{}
	for rows.Next() {
		tag, err := scanIntoTag(rows)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func scanIntoTag(rows *sql.Rows) (*Tag, error) {
	tag := new(Tag)
	err := rows.Scan(
		&tag.TagID,
		&tag.Name)
	return tag, err
}
