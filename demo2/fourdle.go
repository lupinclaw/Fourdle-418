package main

import (
    // "context"
    "log"
    "io"
    "fmt"
    "errors"
    "reflect"
    "net/http"
    "net/url"
    "net/mail"
    "database/sql"
    // "encoding/json"
    // "os"
    // "golang.org/x/oauth2"
    // "golang.org/x/oauth2/paypal"
    "golang.org/x/crypto/bcrypt"
    // "github.com/golang-jwt/jwt"
    _  "github.com/mattn/go-sqlite3"
)

var (
    port   = ":8080"             // TODO: should be from environment
    secret = []byte("secretKey") // TODO: should be from environment
    db *sql.DB
    
    // TODO: this is stupid stop this cut it out
    sqlInitTables = []string{`
create table if not exists User (
   UserID              integer  primary key autoincrement,
   Email               text not null,
   PasswordSaltAndHash text not null
);`, `
create table if not exists SubscriptionType (
    SubscriptionTypeID integer  primary key autoincrement,
    Type               text not null,
    Fee                int  not null
);`, `
create table if not exists Subscription (
    SubscriptionID     integer  primary key autoincrement,
    UserID             int  not null,
    SubscriptionTypeID int  not null,
    Active             int  not null,
    LastInvoiceDate    text not null,
    SignUpDate         text nor null,
    
    foreign key (UserID) references User (UserID),
    foreign key (SubscriptionTypeID) references SubscriptionType (SubscriptionTypeID)
);`, `
create table if not exists Word (
    WordID  integer  primary key autoincrement,
    Letters text not null
);`, `
create table if not exists Game (
    GameID        integer  primary key autoincrement,
    WordID        int  not null,
    WinOrLoss     int  not null,
    WordOfDayDate text,
    
    foreign key (WordID) references Word (WordID)
);`, `
create table if not exists UserMove (
    UserMoveID integer  primary key autoincrement,
    UserID     int  not null,
    GameID     int  not null,
    Guess      text not null,
    Date       text not null,
    
    foreign key (UserID) references User (UserID),
    foreign key (GameID) references Game (GameID)
);`}
)

func signup(email string, password string) error {
    log.Printf("[INFO] [%v] Attempting to sign up.\n", email)
    
    // see if email is valid
    _, err := mail.ParseAddress(email)
    if err != nil { return err }
    
    // see if the email is already registered
    var exists int
    err = db.QueryRow("select exists(select 1 from User where Email=$1);", email).Scan(&exists)
    if err    != nil { log.Fatal("signup DB query failed:", err) }
    if exists == 1   { return errors.New("Email address is already registered.") }
    
    // TODO: check that password is ok
    
    // TODO: temporary redirect to paypal (?)
    
    // generate password hash
    hash, err := bcrypt.GenerateFromPassword([]byte(password), 8)
    fatal(err, "[signup] could not encrypt password", password)
    
    // insert into table
    _, err  = db.Exec("insert into User (Email, PasswordSaltAndHash) values ($1, $2);", email, string(hash[:]))
    fatal(err, "[signup] could not insert user", email)
    
    log.Printf("[INFO] [%v] Registered user.\n", email)
    return nil;
}

func login(email, password string) error {
    log.Printf("[INFO] [%v] Attempting to sign in.\n", email)
    var hash []byte
    var err error
    
    err = db.QueryRow("select PasswordSaltAndHash from User where Email=$1", email).Scan(&hash)
    if err == sql.ErrNoRows { return errors.New("User is not registered.") }
    fatal(err, "[login] Could not read password from database");
    
    err = bcrypt.CompareHashAndPassword(hash, []byte(password))
    if err != nil { return errors.New("Entered incorrect password.") }
    
    log.Printf("[INFO] [%v] Signed in.\n", email)
    return nil
}

func main() {
    var err error
    
    // init database
    db, err = sql.Open("sqlite3", "fourdle.db")
    if err != nil { log.Fatal(err) }
    defer db.Close()
    for i, statement := range sqlInitTables {
        _, err  = db.Exec(statement)
        if err != nil { log.Fatalf("Could not init table %v: %v.\n", i, err) }
    }
    
    // set up http endpoints
    
    fs := http.FileServer(http.Dir("./static/"))
    http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
        log.Printf("REQUEST: %v\n", pp(*req))
        fs.ServeHTTP(resp, req)
    })
    
    http.HandleFunc("/sign-in", func(resp http.ResponseWriter, req *http.Request) {
        query := req.URL.Query()
        email, password := query.Get("email"), query.Get("password")
        
        if err = login(email, password); err != nil {
            log.Printf("[INFO] [%v] Rejected login: %v\n", email, err)
            resp.WriteHeader(http.StatusBadRequest)
            io.WriteString(resp, err.Error())
            return
        }
        
        http.Redirect(resp, req, "/", http.StatusFound)
    })
    
    http.HandleFunc("/sign-up", func(resp http.ResponseWriter, req *http.Request) {
        resp.Header().Set("Content-Type", "text/plain; charset=utf-8")
        
        query := req.URL.Query()
        email, password := query.Get("email"), query.Get("password")
        
        if err = signup(email, password); err != nil {
            log.Printf("[ERRO] [%v] Could not sign up: %v\n", email, err)
            resp.WriteHeader(http.StatusBadRequest)
            io.WriteString(resp, err.Error())
            return
        }
        
        http.Redirect(resp, req, fmt.Sprintf("/sign-in?email=%v&password=%v", url.PathEscape(email), url.PathEscape(password)), http.StatusFound)
    })
    
    // profit
    log.Println("Starting server...")
    fatal(http.ListenAndServe(port, nil), "Could not start server")
}

// func execSql(stmn string, args ...any) { ... }
// func execSqlFile(path string) {
//     file, err := ioutil.ReadFile("/some/path/to/file")
//     if err != nil {
//         // handle error
//     }
//     requests := strings.Split(string(file), ";")
//     for _, request := range requests {
//         _, err := db.Exec(request)
//         // do whatever you need with result and error
//     }
// }
func fatal(err error, args ...any) { if err != nil { log.Fatal("[ERRO] " + fmt.Sprint(args...) + ": " + err.Error()) } }
func pp[T any](a T) string {
    r := reflect.ValueOf(&a).Elem()
    t := r.Type()
    s := fmt.Sprintf("%v {\n", t);
    for i := 0; i < r.NumField(); i++ {
        f := r.Field(i)
        s += fmt.Sprintf("    %s %s = %v\n", t.Field(i).Name, f.Type(), f)
    }
    s += "}"
    return s;
}
