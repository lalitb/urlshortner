package main
import(
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

func main_test1() {
    db := InitDB();
    id := insertLongUrl(db, "test");
    fmt.Println(id);
    fmt.Println(getLongUrl(db, id));

}
func InitDB() *sql.DB{

   db, err := sql.Open("sqlite3", "./urldb.sqlite");
   checkError(err);
   return db;

}

func insertLongUrl(db *sql.DB, longUrl string) int64 {

    fmt.Printf("test")
    stmt, err := db.Prepare("INSERT INTO urls(longUrl) values(?)");
    checkError(err);
    defer stmt.Close();
    _, err = stmt.Exec(longUrl);
    fmt.Printf("test" + longUrl)
    checkError(err);
    stmt, err = db.Prepare("SELECT MAX(ID) FROM urls WHERE longUrl = ?")
    var id int64 = -1
    rows, err1 := stmt.Query(longUrl);
    checkError(err1)
    defer rows.Close()
    rows.Next()
    err2 := rows.Scan(&id)
    fmt.Printf("Test: %d", id)
    checkError(err2);
    return id

}

func getLongUrl(db *sql.DB, id int64) string {
    var longUrl string;
    stmt, err := db.Prepare("SELECT longURL FROM urls WHERE ID = ?")
    checkError(err);
    defer stmt.Close()
    err = stmt.QueryRow(id).Scan(&longUrl);
    return longUrl ;
}

func checkError(err error) {
    if err != nil {
        log.Fatal(err);        
    }
}
