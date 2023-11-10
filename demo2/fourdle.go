package main

import (
    "log"
    "net/http"
    "net/mail"
    "database/sql"
    // "os"
    "golang.org/x/crypto/bcrypt"
    _  "github.com/mattn/go-sqlite3"
)

const (
    port = ":8080"
)

var (
    db *sql.DB
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

func userSignIn(resp http.ResponseWriter, req *http.Request) {
    query := req.URL.Query()
    email, password := query.Get("email"), query.Get("password")
    log.Printf("[INFO] [%v] Attempting to sign in.\n", email)
    
    var hash []byte
    if err := db.QueryRow("select PasswordSaltAndHash from User where Email=$1", email).Scan(&hash); err == sql.ErrNoRows {
        // TODO: invalid form or whatever
        log.Printf("[INFO] [%v] User is not registered.\n", email)
        return
    } else if err != nil {
        // TODO: internal server error
        log.Printf("[ERRO] [%v] Could not read password from database: %v\n", email, err)
        return
    }
    
    if err := bcrypt.CompareHashAndPassword(hash, []byte(password)); err != nil {
        // TODO: invalid form or whatever
        log.Printf("[INFO] [%v] Entered incorrect password.\n", email)
        return
    }
    
    log.Printf("[INFO] [%v] Signed in.\n", email)
    
    // TODO: send back some sort of session token
    http.ServeFile(resp, req, "./static/index.html")
}

func userSignUp(resp http.ResponseWriter, req *http.Request) {
    query := req.URL.Query()
    email, password := query.Get("email"), query.Get("password")
    log.Printf("[INFO] [%v] Attempting to sign up.\n", email)
    
    // see if email is valid
    _, err := mail.ParseAddress(email)
    if err != nil {
        // TODO: invalid form or whatever
        log.Printf("[INFO] [%v] Email is invalid: %v.\n", email, err)
        return
    }
    
    // see if the email is already registered
    var exists int
    if err := db.QueryRow("select exists(select 1 from User where Email=$1);", email).Scan(&exists); err != nil {
        // TODO: internal server error
        log.Printf("[ERRO] [%v] Could not check database for user: %v\n", email, err)
        return
    } else if exists == 1 {
        // TODO: invalid form or whatever
        log.Printf("[INFO] [%v] User is already registered.\n", email)
        return
    }
    
    // TODO: stinky paypal will not let me make a developer account........
    
    // generate password hash
    hash, err := bcrypt.GenerateFromPassword([]byte(password), 8)
    if err != nil {
        // TODO: internal server error
        log.Printf("[ERRO] [%v] Could not generate passord hash:", email, err)
        return
    }
    
    _, err  = db.Exec("insert into User (Email, PasswordSaltAndHash) values ($1, $2);", email, string(hash[:]))
    if err != nil {
        // TODO: internal server error
        log.Printf("[ERRO] [%v] Could not insert user into table: %v\n", email, err)
        return
    }
    
    log.Printf("[INFO] [%v] Registered user.\n", email)
    http.ServeFile(resp, req, "./static/index.html")
}

func main() {
    // init database
    var err error
    db, err = sql.Open("sqlite3", "fourdle.db")
    if err != nil { log.Fatal(err) }
    defer db.Close()
    for i, statement := range sqlInitTables {
        _, err  = db.Exec(statement)
        if err != nil { log.Fatalf("Could not init table %v: %v.\n", i, err) }
    }
    
    // set http handlers
    http.Handle("/", http.FileServer(http.Dir("./static/")))
    http.HandleFunc("/sign-in", userSignIn)
    http.HandleFunc("/sign-up", userSignUp)
    
    // profit
    log.Fatal(http.ListenAndServe(port, nil))
}

