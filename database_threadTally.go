package main

// func (s *PGStore) CreateTotalThreadTally() error {
// 	logInfo("Running CreateTotalThreadTally")
// 	query := `INSERT INTO threadtally
// 	(tallyID, count)
// 	values ($1, $2)`
// 	_, err := s.DB.Query(query, 0, 0)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (s *PGStore) CreateThreadTally(threadTally *ThreadTally) error {
// 	logInfo("Running CreateThreadTally")
// 	query := `INSERT INTO threadtally
// 	(tagID, count)
// 	values ($1, $2)`
// 	_, err := s.DB.Query(query,
// 		threadTally.TagID,
// 		threadTally.Count)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (s *PGStore) UpdateTotalThreadTally(amt int) error {
// 	logInfo("Running UpdateTotalThreadTally")
// 	query := `UPDATE threadtally
// 	SET count = count + $1
// 	WHERE tallyID = 0`
// 	_, err := s.DB.Query(query, amt)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (s *PGStore) UpdateTagThreadTally(tagId uuid.UUID, amt int) error {
// 	logInfo("Running UpdateTagThreadTally")
// 	query := `UPDATE threadtally
// 	SET count = count + $1
// 	WHERE tagID = $2`
// 	_, err := s.DB.Query(query, amt, tagId)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (s *PGStore) GetTotalThreadTally() (*TotalThreadTally, error) {
// 	logInfo("Running GetTotalThreadTally")
// 	query := `SELECT count FROM threadtally
// 	WHERE tallyID = 0;`
// 	rows, err := s.DB.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	for rows.Next() {
// 		return scanIntoTotalThreadTally(rows)
// 	}
// 	return nil, fmt.Errorf("total thread tally not found")
// }

// func (s *PGStore) GetTagThreadTally(tagId uuid.UUID) (*ThreadTally, error) {
// 	logInfo("Running GetTagThreadTally")
// 	query := `SELECT count FROM threadtally
// 	WHERE tagID = $1`
// 	rows, err := s.DB.Query(query, tagId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	for rows.Next() {
// 		return scanIntoThreadTally(rows)
// 	}
// 	return nil, fmt.Errorf("thread tally for tag [%v] not found", tagId)
// }

// //No need to delete as tags cannot be deleted

// func scanIntoTotalThreadTally(rows *sql.Rows) (*TotalThreadTally, error) {
// 	thread := new(TotalThreadTally)
// 	err := rows.Scan(
// 		&thread.Count)
// 	return thread, err
// }

// func scanIntoThreadTally(rows *sql.Rows) (*ThreadTally, error) {
// 	thread := new(ThreadTally)
// 	err := rows.Scan(
// 		&thread.Count)
// 	return thread, err
// }
