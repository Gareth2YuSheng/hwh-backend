package main

func (s *PGStore) CreateImage(image *Image) error {
	logInfo("Running: Database - CreateImage")
	//Create Tag
	query := `INSERT INTO images
	(imageID, threadID, cloudinaryURL)
	values ($1, $2, $3);`
	_, err := s.DB.Query(query,
		image.ImageID,
		image.ThreadID,
		image.CloudinaryURL)
	if err != nil {
		return err
	}
	return nil
}
