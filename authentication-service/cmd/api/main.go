package main

import(
	"database/sql"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)


const webPort = "80"

var counts int64

type Config struct {
	DB *sql.DB
	Models data.Models
}

func main() {
log.Println("Starting authentication service")

 // TODO connect to DB

 conn := connectToDB()
 if conn == nil {
	log.panic("can't connect to Postgres!")
 }

 // set up config
 app := Config{
	DB: conn,
	Models: data.New(conn),
 }

 srv := &http.Server {
	Addr: fmt.sprintf(":%s", webPort),
	Handler: app.routes(),
 }

 err := srv.ListenAndServe()
 if err != nil{
	log.panic(err)
 }
}

func openDB(dsn string) (*sql.DB, error) {  // dsn is connection string for db 
db, err := sql.Open("pgx", dsn)
if err != nil {
	return nil, err
}
   err = db.Ping()
   if err != nil {
	return nil, err
}

return db, nil
   
} 

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	for{
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres is not ready yet")
			counts++
		}else{
			log.Println("Connected to Postgres")
			return connection
		}

		if counts > 10{
			log.Println(err)
			return nil
		}

		log.Println("Back off 2 seconds")
		time.Sleep(2 * time.seconds)
		continue
	}
}
