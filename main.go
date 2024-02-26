package main

import "AgenciaTurismo/objects"

const populateDatabase = false

func checkPopulateDB(db *objects.BancoDados) {
	if !populateDatabase {
		return
	}

	db.PopularBanco()
}

func main() {
	db := objects.BancoDados{}
	checkPopulateDB(&db)
}
