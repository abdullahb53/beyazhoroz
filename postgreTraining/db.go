package database

//>docker exec -it server-db-1 psql -U root -W authServer
import (
	"database/sql" // add this
	"fmt"
	"os"

	"github.com/lib/pq"
	_ "github.com/lib/pq" // add this
)

var (
	// host     = os.Getenv("DB_HOST")
	// port     = os.Getenv("DB_PORT")
	// user     = os.Getenv("DB_USER")
	// password = os.Getenv("DB_PASSWORD")
	// dbname   = os.Getenv("DB_NAME")

	host = os.Getenv("DATABASE_HOST")
	// host     = "host.docker.internal"
	port     = os.Getenv("DATABASE_PORT")
	user     = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	dbname   = os.Getenv("POSTGRES_DB")
)

var connStr string

func CreateDBEngine() (*sql.DB, error) {
	// fmt.Println("os.Getenv(POSTGRES_USER)->", os.Getenv("POSTGRES_USER"))
	// fmt.Println("os.Getenv(DB_USER) ->", os.Getenv("DB_USER"))
	if os.Getenv("ENV") == "PROD" {
		connStr, err := pq.ParseURL(os.Getenv("DATABASE_URL"))
		if err != nil {
			return nil, fmt.Errorf(connStr, "Bu postgres connection string calismadi.")
		}

	} else {
		//fmt.Sprintf("postgres://%s:%s@%s:%s/%s",user,password,host,port,dbname)
		//connStr = fmt.Sprintf("postgresql://%s:%s@%s/%s:%s?sslmode=disable", user, password, host, dbname, port)
		connStr = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
		fmt.Println(connStr, "Bu connStr dota2 POSTGRESQL calistiii.", "os.Getenv(POSTGRES_USER)->", os.Getenv("POSTGRES_USER"), "os.Getenv(DB_USER) ->", os.Getenv("DB_USER"))

	}

	//connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	//err := godotenv.Load(filepath.Join("C:/Users/abdullah/beyazhoroz/server/apiServerService/database/", ".env"))

	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	// host := os.Getenv("DB_HOST")
	// port := 5432
	// user := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	// dbname := os.Getenv("DB_NAME")

	// API_PORT=8080
	// DB_HOST=localhost
	// DB_DRIVER=postgres
	// DB_USER=root
	// DB_PASSWORD=123456
	// DB_NAME=apiserverservice
	// DB_PORT=5432

	// now you can use os.Getenv ...

	// CreateProductTable ...

	_, err = db.Query(`CREATE TABLE IF NOT EXISTS userss (
		uid SERIAL PRIMARY KEY,
		username character varying(100) NOT NULL,
		city character varying(100) NOT NULL,
		country character varying(100) NOT NULL, 
		photo character varying(100) NOT NULL,
		explain character varying(500) NOT NULL
	)
	
	`)

	if err != nil {
		return nil, err
	}

	return db, nil
	// CREATE TABLE userinfo
	// (
	//     uid serial NOT NULL,
	//     username character varying(100) NOT NULL,
	//     city character varying(100) NOT NULL,
	//	   country character varying(100) NOT NULL,
	//     Created date,
	//     CONSTRAINT userinfo_pkey PRIMARY KEY (uid)
	// )

	// WITH (OIDS=FALSE);

}

type User struct {
	Uid      int    `json:"uid"`
	Username string `json:"username"`
	City     string `json:"city"`
	Country  string `json:"country"`
	Photo    string `json:"photo"`
	Explain  string `json:"explain"`
}

func GetUseCiCo(db *sql.DB, City string, Country string) ([]User, error) {

	fmt.Println("# <<Get Users With City and Countries<<")

	rows, err := db.Query("SELECT uid,username,city,country,photo,explain FROM userss WHERE city=$1 AND country=$2", City, Country)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	all_users := []User{}
	for rows.Next() {
		var p User
		if err := rows.Scan(&p.Uid, &p.Username, &p.City, &p.Country, &p.Photo, &p.Explain); err != nil {
			return nil, err
		}
		all_users = append(all_users, p)
	}

	return all_users, nil
}

func InsertUser(db *sql.DB, Username string, City string, Country string, Photo string, Explain string) error {

	fmt.Println("# <<Insert User<<")

	ins := User{
		Username: Username,
		City:     City,
		Country:  Country,
		Photo:    Photo,
		Explain:  Explain,
	}

	rows, err := db.Query("INSERT INTO userss(username, city, country, photo, explain) VALUES ($1,$2,$3,$4,$5)", ins.Username, ins.City, ins.Country, ins.Photo, ins.Explain)
	if err != nil {
		return err
	}

	defer rows.Close()

	fmt.Println("INSERT INTO -BASARILI-> ")

	return nil
}
