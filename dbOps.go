package main
import(
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

func InitDB() {
    db := OpenDB()
    rows, err := db.Query(" SELECT COUNT(*) FROM  sqlite_master WHERE  name= 'urls'")
    CheckError(err);
    defer db.Close()
    var id = 0;
    rows.Next()
    err = rows.Scan(&id)
    CheckError(err)
    rows.Close()
    if (id == 0 ){
        _, err1 := db.Exec( `CREATE TABLE urls (
                                id INTEGER PRIMARY KEY AUTOINCREMENT,
                                longurl text )` )
        CheckError(err1)
        _, err2 := db.Exec("INSERT INTO sqlite_sequence VALUES('urls',10000000)")
        CheckError(err2)
    }
}

func OpenDB() *sql.DB{
   db, err := sql.Open("sqlite3", "./urldb.sqlite");
   CheckError(err);
   return db;

}

func InsertLongUrl(db *sql.DB, longUrl string) int64 {
    stmt, err := db.Prepare("INSERT INTO urls(longUrl) values(?)");
    CheckError(err);
    defer stmt.Close();
    _, err = stmt.Exec(longUrl);
    CheckError(err);
    stmt, err = db.Prepare("SELECT MAX(ID) FROM urls WHERE longUrl = ?")
    var id int64 = -1
    rows, err1 := stmt.Query(longUrl);
    CheckError(err1)
    defer rows.Close()
    rows.Next()
    err2 := rows.Scan(&id)
    CheckError(err2);
    return id

}

func GetLongUrl(db *sql.DB, id int64) string {
    var longUrl string;
    stmt, err := db.Prepare("SELECT longURL FROM urls WHERE ID = ?")
    CheckError(err);
    defer stmt.Close()
    err = stmt.QueryRow(id).Scan(&longUrl);
    return longUrl ;
}

func CheckError(err error) {
    if err != nil {
        log.Fatal(err);        
    }
}
