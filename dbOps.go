package main
import(
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

func main() {
    db := initDB();
    id := insertLongURL(db, "test");
    fmt.Println(id);

}
func initDB() *sql.DB{

   db, err := sql.Open("sqlite3", "./urldb.sqlite");
   checkError(err);
   return db;
  // defer db.Close()

}

func insertLongURL(db *sql.DB, longURL string) int64 {

    stmt, err := db.Prepare("INSERT INTO urls(longURL) values(?)");
    checkError(err);
    defer stmt.Close();
    res, err := stmt.Exec(longURL);
    checkError(err);
    id, err := res.LastInsertId()
    checkError(err);
    return id

}

func getLongURL(db *sql.DB, id int64) string {
    var longURL string;
    stmt, err := db.Prepare("SELECT longURL FROM urls WHERE rowid = ?")
    checkError(err);
    defer stmt.Close()
    err = stmt.QueryRow(id).Scan(&longURL);
    return longURL ;

}

func checkError(err error) {
    if err != nil {
        log.Fatal(err);        
    }
}
