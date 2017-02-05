package main
import(
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

func InitDB() {
    db := OpenDB()
    rows, err := db.Query(" SELECT COUNT(*) FROM  sqlite_master WHERE  name= 'urls'")
    checkError(err);
    defer db.Close()
    var id = 0;
    rows.Next()
    err = rows.Scan(&id)
    checkError(err)
    rows.Close()
    if (id == 0 ){
        _, err1 := db.Exec( `CREATE TABLE urls (
                                id INTEGER PRIMARY KEY AUTOINCREMENT,
                                longurl text )` )
        checkError(err1)
        _, err2 := db.Exec("INSERT INTO sqlite_sequence VALUES('urls',10000000)")
        checkError(err2)
    }
}

func OpenDB() *sql.DB{
   db, err := sql.Open("sqlite3", "./urldb.sqlite");
   checkError(err);
   return db;

}

func insertLongUrl(db *sql.DB, longUrl string) int64 {
    stmt, err := db.Prepare("INSERT INTO urls(longUrl) values(?)");
    checkError(err);
    defer stmt.Close();
    _, err = stmt.Exec(longUrl);
    checkError(err);
    stmt, err = db.Prepare("SELECT MAX(ID) FROM urls WHERE longUrl = ?")
    var id int64 = -1
    rows, err1 := stmt.Query(longUrl);
    checkError(err1)
    defer rows.Close()
    rows.Next()
    err2 := rows.Scan(&id)
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
