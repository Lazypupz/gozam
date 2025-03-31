package fingerprint

import (
	"database/sql"
	"fmt"
	"hash/fnv"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 8080
	user     = "gozam_user"
	password = "securepassword"
	dbname   = "gozam_db"
)

func GenFingerPrint(peaks [][]int) ([]byte, string) {

	hash := fnv.New64a()

	for _, peak := range peaks {

		peakStr := fmt.Sprintf("%d-%d", peak[0], peak[1])

		hash.Write([]byte(peakStr))
	}

	fingerprint := hash.Sum(nil)

	// Generate song ID as a hash of the concatenated peaks
	songIDHash := fnv.New64a()
	for _, peak := range peaks {
		peakStr := fmt.Sprintf("%d-%d", peak[0], peak[1])
		songIDHash.Write([]byte(peakStr))
	}

	return fingerprint, fmt.Sprintf("%x", songIDHash.Sum(nil))
}

func SaveFingerprintToDB(fingerprint []byte, songID string) error {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Insert fingerprint and songID into the database
	query := `INSERT INTO fingerprints (fingerprint, song_id) VALUES ($1, $2)`
	_, err = db.Exec(query, fingerprint, songID)
	if err != nil {
		return fmt.Errorf("failed to insert data: %v", err)
	}

	return nil
}
