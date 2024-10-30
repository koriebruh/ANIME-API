package cnf

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"strconv"
)

func InsertDataCSV() {
	// Koneksi ke MySQL
	cnf := GetConfig()
	dsn := cnf.DataBase.User + ":" + cnf.DataBase.Pass + "@tcp(" + cnf.DataBase.Host + ":" + cnf.DataBase.Port + ")/" + cnf.DataBase.Name

	// Pastikan untuk menyertakan driver name ("mysql") sebagai parameter pertama
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Buka file CSV
	file, err := os.Open("processed_anime_dataset.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Baca file CSV
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Loop untuk memasukkan data ke database
	for i, record := range records {
		if i == 0 {
			continue // Lewati header
		}

		animeID, _ := strconv.Atoi(record[0])
		score, _ := strconv.ParseFloat(record[4], 64)
		rank, _ := strconv.ParseFloat(record[18], 64)
		popularity, _ := strconv.Atoi(record[19])
		favorites, _ := strconv.Atoi(record[20])
		members, _ := strconv.Atoi(record[22])

		// Query untuk memasukkan data
		_, err = db.Exec(`
    INSERT INTO anime_info (
        anime_id, name, english_name, other_name, score, genres, synopsis, type, episodes,
        aired, premiered, status, producers, licensors, studios, source, duration,
        rating, `+"`rank`"+`, popularity, favorites, scored_by, members, image_url
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			animeID, record[1], record[2], record[3], score, record[5], record[6], record[7],
			record[8], record[9], record[10], record[11], record[12], record[13], record[14],
			record[15], record[16], record[17], rank, popularity, favorites, record[21], members,
			record[23],
		)
		log.Print("insert ke ", i)
	}

	fmt.Println("Data successfully inserted!")
}
