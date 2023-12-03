package main

import (
	// "context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"reflect"
	"runtime/debug"
	"strings"
	"time"

	// "net/url"
	"database/sql"
	"net/mail"

	// "encoding/json"
	"os"
	// "golang.org/x/oauth2"
	// "golang.org/x/oauth2/paypal"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"

	// "github.com/golang-jwt/jwt"
	_ "github.com/mattn/go-sqlite3"
)

var (
    port   = ":8080"             // TODO: should be from environment
    secret = []byte("secretKey") // TODO: should be from environment
    store = sessions.NewCookieStore(secret)
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

func login(email, password string) (int, error) {
    log.Printf("[INFO] [%v] Attempting to sign in.\n", email)
    var hash []byte
    var userID int
    var err error
    
    err = db.QueryRow("select PasswordSaltAndHash from User where Email=$1", email).Scan(&hash)
    if err == sql.ErrNoRows { return -1, errors.New("User is not registered") }
    fatal(err, "[login] Could not read password from database");
    
    err = bcrypt.CompareHashAndPassword(hash, []byte(password))
    if err != nil { return -1, errors.New("Entered incorrect password") }
    
    err = db.QueryRow("select UserID from User where Email=$1", email).Scan(&userID)
    if err != nil { return -1, errors.New("Could not get UserID") }

    log.Printf("[INFO] [%v] Signed in.\n", email)
    
    return userID, nil
}

func getNumberOfGames(userID int) string {

    var result string
    var err error
    err = db.QueryRow("SELECT COUNT(GameID) FROM UserGame WHERE UserID = $1", userID).Scan(&result)
    if err == sql.ErrNoRows { return "0"}
    fatal(err, "[query] Could not read games from database");

    return result
}

func getNumberOfGamesWon(userID int) string {

    var result string
    var err error
    err = db.QueryRow("SELECT COUNT(U.GameID) FROM UserGame U, Game G WHERE G.GameID = U.GameID AND UserID = $1 AND G.WinOrLoss = 1", userID).Scan(&result)
    if err == sql.ErrNoRows { return "0"}
    fatal(err, "[query] Could not read games from database");

    return result
}

func getRecentWords(userID int) string {

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
    return strings.Join(results, "\n")
}

func getWords(userID int) string {

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
    return strings.Join(results, "\n")
}

func getGames(userID int) string {

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
func editEmail(email string) error {
    log.Printf("[INFO] [%v] Attempting to edit email.\n", email)

    // check if email is valid
    _, err := mail.ParseAddress(email)
    if err != nil {
        return err
    }

    // update the table
    _, err = db.Exec("update User set Email = $1 where Email = $2;", email, email)
    if err != nil {
        return err
    }
    log.Printf("[INFO] [%v] Updated email.\n", email)
    return nil
}

func editPassword(email, password string) error {
    log.Printf("[INFO] [%v] Attempting to edit password.\n", email)

    // generate password hash
    hash, err := bcrypt.GenerateFromPassword([]byte(password), 8)
    if err != nil {
        return err
    }

    // update the table
    _, err = db.Exec("update User set PasswordSaltAndHash = $1 where Email = $2;", string(hash[:]), email)
    if err != nil {
        return err
    }
    log.Printf("[INFO] [%v] Updated password.\n", email)
    return nil
}

func isSessionValid(r *http.Request, name string) (bool, error) {
    session, err := store.Get(r, name)
    if err != nil {
        return false, err
    }

    // Check if user is authenticated
    if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
        return false, nil
    }

    return true, nil
}

func main() {
    var err error
    store.Options.MaxAge = 3600*3;

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
        
        stmt1, err := db.Prepare("insert into Word (Letters, WordOfDayDate) values ($1, $2);")
        fatal(err, "could not prepare statement")
        
        stmt2, err := db.Prepare("insert into Word (Letters) values ($1);")
        fatal(err, "could not prepare statement")
        
        file, err := ioutil.ReadFile("fourdles.txt")
        fatal(err, "could not get fourdles")
        
        fourdles := strings.Split(string(file), "\n")
        rand.Shuffle(len(fourdles), func(i, j int) { fourdles[i], fourdles[j] = fourdles[j], fourdles[i] })
        
        rn  := time.Now();
        day := time.Date(rn.Year(), rn.Month(), rn.Day(), 0, 0, 0, 0, rn.Location())
        for idx, fourdle := range fourdles {
            if idx < 365 * 5 {
                _, err := stmt1.Exec(fourdle, day)
                fatal(err, "Could not insert fourdle", fourdle)
                day = day.Add(time.Hour * 24)
            } else {
                _, err := stmt2.Exec(fourdle)
                fatal(err, "Could not insert fourdle", fourdle)
            }
        }
        
        stmt1.Close()
        stmt2.Close()
    }

    var errorValue error
    var exampleuser string
    errorValue = db.QueryRow("SELECT Email from User WHERE Email LIKE 'exampleuser@test.com'").Scan(&exampleuser)
    if errorValue == sql.ErrNoRows {
        // Example user doesn't exist, create them and add sample data
        execSqlFile("sampledata.sql")
    } else if errorValue != nil {
        fatal(errorValue, "query failed")
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
        //make cookie
        session, _ := store.Get(req, "fourdle-session")

        userID, err := login(email, password)
        if err != nil {
            log.Printf("[INFO] [%v] Rejected login: %v\n", email, err)
            resp.WriteHeader(http.StatusBadRequest)
            io.WriteString(resp, err.Error())
            return
        }

        //save session cookie
        session.Values["email"] = email
        session.Values["userID"] = userID
        session.Values["authenticated"] = true
        session.Save(req, resp)

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
        
        _, err := login(email, password)
        fatal(err, "Could not log in to new account")
        
        // TODO: return session token in response
        http.Redirect(resp, req, "/", http.StatusFound)
    })

    http.HandleFunc("/get-number-of-games", func(resp http.ResponseWriter, req *http.Request) {
        resp.Header().Set("Content-Type", "text/plain; charset=utf-8")

        session, _ := store.Get(req, "fourdle-session")
        userID, _ := session.Values["userID"].(int)
        log.Printf("[INFO] [%d] userID for get-number-of-games\n", userID)

        var result string
        result = getNumberOfGames(userID);

        resp.Write([]byte(result));
    })

    http.HandleFunc("/get-number-of-games-won", func(resp http.ResponseWriter, req *http.Request) {
        resp.Header().Set("Content-Type", "text/plain; charset=utf-8")

        session, _ := store.Get(req, "fourdle-session")
        userID, _ := session.Values["userID"].(int)

        var result string
        result = getNumberOfGamesWon(userID);

        resp.Write([]byte(result));
    })

    http.HandleFunc("/get-recent-words", func(resp http.ResponseWriter, req *http.Request) {
        resp.Header().Set("Content-Type", "text/plain; charset=utf-8")

        session, _ := store.Get(req, "fourdle-session")
        userID, _ := session.Values["userID"].(int)

        var result string
        result = getRecentWords(userID);

        resp.Write([]byte(result));
    })

    http.HandleFunc("/get-words", func(resp http.ResponseWriter, req *http.Request) {
        resp.Header().Set("Content-Type", "text/plain; charset=utf-8")

        session, _ := store.Get(req, "fourdle-session")
        userID, _ := session.Values["userID"].(int)

        var result string
        result = getWords(userID);

        resp.Write([]byte(result));
    })

     http.HandleFunc("/get-games", func(resp http.ResponseWriter, req *http.Request) {
        resp.Header().Set("Content-Type", "text/plain; charset=utf-8")

        session, _ := store.Get(req, "fourdle-session")
        userID, _ := session.Values["userID"].(int)

        var result string
        result = getGames(userID);

        resp.Write([]byte(result));
     })

     http.HandleFunc("/get-games-won", func(resp http.ResponseWriter, req *http.Request) {
         resp.Header().Set("Content-Type", "text/plain; charset=utf-8")

         session, _ := store.Get(req, "fourdle-session")
         userID, _ := session.Values["userID"].(int)

         var result string
         result = getGamesWon(userID);

         resp.Write([]byte(result));
     })

     http.HandleFunc("/edit-account", func(resp http.ResponseWriter, req *http.Request) {
        
        sessionErr := store.Get(req, "fourdle-session")
        if sessionErr != nil {
            http.Error(resp, "Access denied! User not logged in.", http.StatusUnauthorized)
            log.Printf("Access denied! User not logged in.", err)
            return
        }
        
        if req.Method != http.MethodPost {
            http.Error(resp, "Invalid request method", http.StatusMethodNotAllowed)
            return
        }
    
        err := req.ParseForm()
        if err != nil {
            http.Error(resp, "Failed to parse form", http.StatusBadRequest)
            return
        }
    
        email := req.FormValue("email")
        password := req.FormValue("password")
        confirmPassword := req.FormValue("confirm_password")
    
        if password != confirmPassword {
            http.Error(resp, "Passwords do not match", http.StatusBadRequest)
            return
        }
    
        if email != "" {
            // If email is provided, update email
            if err = editEmail(email); err != nil {
                log.Printf("[INFO] [%v] Failed to edit email: %v\n", email, err)
                resp.WriteHeader(http.StatusBadRequest)
                io.WriteString(resp, err.Error())
                return
            }
        }
    
        if password != "" {
            // If password is provided, update password
            if err = editPassword(email, password); err != nil {
                log.Printf("[INFO] [%v] Failed to edit password: %v\n", email, err)
                resp.WriteHeader(http.StatusBadRequest)
                io.WriteString(resp, err.Error())
                return
            }
        }
    
        http.Redirect(resp, req, "/", http.StatusFound)
    })

    http.HandleFunc("/check-cookie", func(resp http.ResponseWriter, req *http.Request) {
        isValid, err := isSessionValid(req, "fourdle-session")
        if err != nil || !isValid {
            http.Error(resp, "Invalid session", http.StatusUnauthorized)
            return
        }
        resp.WriteHeader(http.StatusOK)
    })
    
    http.HandleFunc("/get-email", func(resp http.ResponseWriter, req *http.Request) {
        session, err := store.Get(req, "fourdle-session")
        if err != nil {
            http.Error(resp, "Session not found", http.StatusUnauthorized)
            log.Printf("Session not found", err)
            return
        }
        email, ok := session.Values["email"]
        if !ok {
            http.Error(resp, "Email not found", http.StatusUnauthorized)
            log.Printf("Email not found", err)
            return
        }
        resp.Write([]byte(email.(string)))
    })

    // profit
    log.Println("Serving at http://localhost"+port)
    fatal(http.ListenAndServe(port, nil), "Could not start server")
}

func execSqlFile(path string) {
    file, err := ioutil.ReadFile(path)
    fatal(err, "Could not read sql script ", path)
    
    requests := strings.Split(string(file), ";")
    for _, request := range requests {
        _, err := db.Exec(request)
        fatal(err, "Could not execute script ", path)
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
