package main

import (
    // "context"
    "log"
    "io"
    "io/ioutil"
    "strings"
    "fmt"
    "errors"
    "reflect"
    "runtime/debug"
    "net/http"
    // "net/url"
    "net/mail"
    "database/sql"
    // "encoding/json"
    "os"
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
)

func signup(email string, password string) error {
    log.Printf("[INFO] [%v] Attempting to sign up.\n", email)
    
    // see if email is valid
    _, err := mail.ParseAddress(email)
    if err != nil { return err }
    
    // see if the email is already registered
    var exists int
    err = db.QueryRow("select exists(select 1 from User where Email=$1);", email).Scan(&exists)
    fatal(err, "[signup] DB query failed")
    if exists == 1 { return errors.New("Email address is already registered.") }
    
    // TODO: check that password is ok
    
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
    if err == sql.ErrNoRows { return errors.New("User is not registered") }
    fatal(err, "[login] Could not read password from database");
    
    err = bcrypt.CompareHashAndPassword(hash, []byte(password))
    if err != nil { return errors.New("Entered incorrect password") }
    
    log.Printf("[INFO] [%v] Signed in.\n", email)
    return nil
}

func getNumberOfGames(userID int) string {
    log.Printf("[INFO] Attempting to query number of games for user [%d]", userID);

    var result string
    var err error
    err = db.QueryRow("SELECT COUNT(GameID) FROM UserGame WHERE UserID = $1", userID).Scan(&result)
    if err == sql.ErrNoRows { return "0"}
    fatal(err, "[query] Could not read games from database");

    log.Printf("[INFO] Number of games [%s]", result);
    return result
}

func getNumberOfGamesWon(userID int) string {
    log.Printf("[INFO] Attempting to query number of games won for user [%d]", userID);

    var result string
    var err error
    err = db.QueryRow("SELECT COUNT(U.GameID) FROM UserGame U, Game G WHERE G.GameID = U.GameID AND UserID = $1 AND G.WinOrLoss = 1", userID).Scan(&result)
    if err == sql.ErrNoRows { return "0"}
    fatal(err, "[query] Could not read games from database");

    log.Printf("[INFO] Number of games won [%s]", result);
    return result
}

func getRecentWords(userID int) string {
    log.Printf("[INFO] Attempting to query 10 most recent words");

    var results []string
    var err error
    rows, err := db.Query("SELECT W.Letters FROM UserGame U, Game G, Word W WHERE U.UserID = $1 AND U.GameID = G.GameID AND G.WordID = W.WordID ORDER BY U.Date DESC LIMIT 10;", userID);
    if err == sql.ErrNoRows { return ""}
    fatal(err, "[query] Could not read words from database");

    for rows.Next() {
        var result string

        // Scan the current row into the result variable
        err := rows.Scan(&result)
        if err != nil {
            log.Fatal(err)
        }

        results = append(results, result)
    }

    log.Printf("[INFO] Recent words: [%s]", results);
    return strings.Join(results, "\n")
}

func getWords(userID int) string {
    log.Printf("[INFO] Attempting to query 10 most recent words");

    var results []string
    var err error
    rows, err := db.Query("SELECT W.Letters FROM UserGame U, Game G, Word W WHERE U.UserID = $1 AND U.GameID = G.GameID AND G.WordID = W.WordID ORDER BY U.Date DESC;", userID);
    if err == sql.ErrNoRows { return ""}
    fatal(err, "[query] Could not read words from database");

    for rows.Next() {
        var result string

        // Scan the current row into the result variable
        err := rows.Scan(&result)
        if err != nil {
            log.Fatal(err)
        }

        results = append(results, result)
    }

    log.Printf("[INFO] Recent words: [%s]", results);
    return strings.Join(results, "\n")
}

func getGames(userID int) string {
    log.Printf("[INFO] Attempting to get games for user [%d]", userID);

    var results []string
    var err error
    rows, err := db.Query("SELECT U.Date, COUNT(U.GameID) FROM UserGame U, Game G WHERE U.GameID = G.GameID AND U.UserID = $1 AND G.WinOrLoss = 0 GROUP BY U.Date HAVING U.Date >= DATE('now', '-30 days');", userID);
    if err == sql.ErrNoRows { log.Printf("No rows"); return ""}
    fatal(err, "[query] Could not read games from database");

    for rows.Next() {
        var date string
        var number string

        // Scan the current row into the result variable
        err := rows.Scan(&date, &number)
        if err != nil {
            log.Fatal(err)
        }

        results = append(results, date)
        results = append(results, number)
    }

    return strings.Join(results, "\n")
}

func getGamesWon(userID int) string {
    log.Printf("[INFO] Attempting to get games for user [%d]", userID);

    var results []string
    var err error
    rows, err := db.Query("SELECT U.Date, COUNT(U.GameID) FROM UserGame U, Game G WHERE U.GameID = G.GameID AND U.UserID = $1 AND G.WinOrLoss = 1 GROUP BY U.Date HAVING U.Date >= DATE('now', '-30 days');", userID);
    if err == sql.ErrNoRows { log.Printf("No rows"); return ""}
    fatal(err, "[query] Could not read games from database");

    for rows.Next() {
        var date string
        var number string

        // Scan the current row into the result variable
        err := rows.Scan(&date, &number)
        if err != nil {
            log.Fatal(err)
        }

        results = append(results, date)
        results = append(results, number)
    }

    return strings.Join(results, "\n")
}

func main() {
    var err error
    
    // init database
    db, err = sql.Open("sqlite3", "fourdle.db")
    if err != nil { log.Fatal(err) }
    defer db.Close()
    execSqlFile("schema.sql")
    var wordCount int
    err = db.QueryRow("select Count(*) from Word").Scan(&wordCount)
    fatal(err, "query failed")
    if wordCount == 0 {
        // populate word table
        log.Println("Populating Fourdles...")
        
        stmt, err := db.Prepare("insert into Word (Letters) values ($1);")
        fatal(err, "could not prepare statement")
        
        file, err := ioutil.ReadFile("fourdles.txt")
        fatal(err, "could not get fourdles")
        
        fourdles := strings.Split(string(file), "\n")
        for _, fourdle := range fourdles {
            _, err := stmt.Exec(fourdle)
            fatal(err, "Could not insert fourdle", fourdle)
        }
    }

    // set up http endpoints
    
    fs := http.FileServer(http.Dir("./static/"))
    http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
        log.Printf("[INFO] Request to `/` %v\n", pp(*req))
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
        
        // TODO: return session token in response
        http.Redirect(resp, req, "/", http.StatusFound)
    })
    
    http.HandleFunc("/sign-up", func(resp http.ResponseWriter, req *http.Request) {
        resp.Header().Set("Content-Type", "text/plain; charset=utf-8")
        
        query := req.URL.Query()
        email, password := query.Get("email"), query.Get("password")
        
        if err = signup(email, password); err != nil {
            log.Printf("[INFO] [%v] Could not sign up: %v\n", email, err)
            resp.WriteHeader(http.StatusBadRequest)
            io.WriteString(resp, err.Error())
            return
        }
        
        fatal(login(email, password), "Could not log in to new account")
        
        // TODO: return session token in response
        http.Redirect(resp, req, "/", http.StatusFound)
    })

    http.HandleFunc("/get-number-of-games", func(resp http.ResponseWriter, req *http.Request) {
        resp.Header().Set("Content-Type", "text/plain; charset=utf-8")

        userID := 1

        var result string
        result = getNumberOfGames(userID);

        resp.Write([]byte(result));
        // TODO: return session token in response
        // http.Redirect(resp, req, "/", http.StatusFound)
    })

    http.HandleFunc("/get-number-of-games-won", func(resp http.ResponseWriter, req *http.Request) {
        resp.Header().Set("Content-Type", "text/plain; charset=utf-8")

        userID := 1

        var result string
        result = getNumberOfGamesWon(userID);

        resp.Write([]byte(result));
    })

    http.HandleFunc("/get-recent-words", func(resp http.ResponseWriter, req *http.Request) {
        resp.Header().Set("Content-Type", "text/plain; charset=utf-8")

        userID := 1

        var result string
        result = getRecentWords(userID);

        resp.Write([]byte(result));
    })

    http.HandleFunc("/get-words", func(resp http.ResponseWriter, req *http.Request) {
        resp.Header().Set("Content-Type", "text/plain; charset=utf-8")

        userID := 1

        var result string
        result = getWords(userID);

        resp.Write([]byte(result));
    })

     http.HandleFunc("/get-games", func(resp http.ResponseWriter, req *http.Request) {
            resp.Header().Set("Content-Type", "text/plain; charset=utf-8")

            userID := 1

            var result string
            result = getGames(userID);

            resp.Write([]byte(result));
        })

     http.HandleFunc("/get-games-won", func(resp http.ResponseWriter, req *http.Request) {
         resp.Header().Set("Content-Type", "text/plain; charset=utf-8")

         userID := 1

         var result string
         result = getGamesWon(userID);

         resp.Write([]byte(result));
     })

    // profit
    log.Println("Serving at http://localhost"+port)
    fatal(http.ListenAndServe(port, nil), "Could not start server")
}

func execSqlFile(path string) {
    file, err := ioutil.ReadFile(path)
    fatal(err, "Could not read sql script", path)
    
    requests := strings.Split(string(file), ";")
    for _, request := range requests {
        _, err := db.Exec(request)
        fatal(err, "Could not execute script", path)
    }
}

func fatal(err error, args ...any) {
    if err != nil {
        log.Println("[FATAL] " + fmt.Sprint(args...) + ": " + err.Error())
        debug.PrintStack()
        os.Exit(1)
    }
}

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
