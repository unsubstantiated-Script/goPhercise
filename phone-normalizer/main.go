package main

import (
	"fmt"
	_ "github.com/lib/pq"
	phonedb "goPhercise/phone-normalizer/db"
	"regexp"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "gophercises_phone"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	must(phonedb.Reset("postgres", psqlInfo, dbname))

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	must(phonedb.Migrate("postgres", psqlInfo))

	db, err := phonedb.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()

	err = db.Seed()
	must(err)

	phones, err := db.AllPhones()
	must(err)
	for _, p := range phones {
		fmt.Printf("Working on... %+v\n", p)

		number := normalize(p.Number)

		if number != p.Number {
			fmt.Println("Updating or removing...", number)
			existing, err := db.FindPhone(number)
			must(err)
			if existing != nil {
				must(db.DeletePhone(p.ID))
			} else {
				p.Number = number
				must(db.UpdatePhone(&p))
			}
		} else {
			fmt.Println("No changes required")
		}
	}

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// REGEX version
func normalize(phone string) string {
	// If not caught with testing, bad stuff can happen with this MustCompile
	//re := regexp.MustCompile("[^0-9]")
	//Another way of Regexing digits double backslash is to escape backslash. Capitol D is any non-digits
	re := regexp.MustCompile("\\D")
	return re.ReplaceAllString(phone, "")
}
