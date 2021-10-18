package level

import "github.com/syndtr/goleveldb/leveldb"

func GetGConfigbyString() {
	db, err := leveldb.OpenFile("path/to/db", nil)
	defer db.Close()
}
