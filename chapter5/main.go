package main

import (
	chapter5 "chapter5/gen"
	"chapter5/pkg"
	"context"
	"database/sql"
	"embed"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"time"
)

var (
	Version string = strings.TrimSpace(version)
	//go:embed version/version.txt
	version string

	//go:embed static
	staticEmbed embed.FS

	//go:embed css/*
	cssEmbed embed.FS

	//go:embed tmpl/*.html
	tmplEmbed embed.FS

	dbQuery *chapter5.Queries

	store = sessions.NewCookieStore([]byte("forDemo"))
)

//renderFiles renders file and push data (d) into the templates to be rendered
func renderFiles(tmpl string, w http.ResponseWriter, d interface{}) {
	t, err := template.ParseFS(tmplEmbed, fmt.Sprintf("tmpl/%s.html", tmpl))
	if err != nil {
		log.Fatal(err)
	}

	if err := t.Execute(w, d); err != nil {
		log.Fatal(err)
	}
}

//securityMiddleware is middleware to make sure all request has a
//valid session and authenticated
func securityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//first of all all request MUST have a valid session
		if sessionValid(w, r) {
			//login path will be let through, otherwise it won't be served
			//to the front end
			if r.URL.Path == "/login" {
				next.ServeHTTP(w, r)
				return
			}
		}

		//if it does have a valid session make sure it has been authenticated
		if hasBeenAuthenticated(w, r) {
			next.ServeHTTP(w, r)
			return
		}

		//otherwise it will need to be redirected to /login
		storeAuthenticated(w, r, false)
		http.Redirect(w, r, "/login", 307)
	})
}

//sessionValid check whether the session is a valid session
func sessionValid(w http.ResponseWriter, r *http.Request) bool {
	session, _ := store.Get(r, "session_token")
	return !session.IsNew
}

//authenticationHandler handles authentication
func authenticationHandler(w http.ResponseWriter, r *http.Request) {
	result := "Login "
	r.ParseForm()

	if validateUser(r.FormValue("username"), r.FormValue("password")) {
		storeAuthenticated(w, r, true)
		result = result + "successfull"
	} else {
		result = result + "unsuccessful"
	}

	renderFiles("msg", w, result)
}

//hasBeenAuthenticated checks whether the session contain the flag to indicate
//that the session has gone through authentication process
func hasBeenAuthenticated(w http.ResponseWriter, r *http.Request) bool {
	session, _ := store.Get(r, "session_token")
	a, _ := session.Values["authenticated"]

	if a == nil {
		return false
	}

	return a.(bool)
}

//storeAuthenticated to store authenticated value
func storeAuthenticated(w http.ResponseWriter, r *http.Request, v bool) {
	session, _ := store.Get(r, "session_token")

	session.Values["authenticated"] = v
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//validateUser check whether username/password exist in database
func validateUser(username string, password string) bool {
	//query the data from database
	ctx := context.Background()
	u, _ := dbQuery.GetUserByName(ctx, username)

	//username does not exist
	if u.UserName != username {
		return false
	}

	return pkg.CheckPasswordHash(password, u.PassWordHash)
}

func basicMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		log.Println("Middleware called on", req.URL.Path)
		// do stuff
		h.ServeHTTP(wr, req)
	})
}

func main() {
	log.Println("Server Version :", Version)
	initDatabase()

	router := mux.NewRouter()

	//POST handler for /login
	router.HandleFunc("/login", authenticationHandler).Methods("POST")

	//embed handler for /css path
	ccssontentStatic, _ := fs.Sub(cssEmbed, "css")
	css := http.FileServer(http.FS(ccssontentStatic))
	router.PathPrefix("/css").Handler(http.StripPrefix("/css", css))

	//embed handler for /app path
	contentStatic, _ := fs.Sub(staticEmbed, "static")
	static := http.FileServer(http.FS(contentStatic))
	router.PathPrefix("/app").Handler(securityMiddleware(http.StripPrefix("/app", static)))

	//add /login path
	router.PathPrefix("/login").Handler(securityMiddleware(http.StripPrefix("/login", static)))

	//root will redirect to /apo
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/app", http.StatusPermanentRedirect)
	})

	// Use our basicMiddleware
	router.Use(basicMiddleware)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:3333",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func initDatabase() {
	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		GetAsString("DB_USER", "postgres"),
		GetAsString("DB_PASSWORD", "mysecretpassword"),
		GetAsString("DB_HOST", "localhost"),
		GetAsInt("DB_PORT", 5432),
		GetAsString("DB_NAME", "postgres"),
	)

	// Open the database
	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		panic(err)
	}

	// Connectivity check
	if err := db.Ping(); err != nil {
		log.Fatalln("Error from database ping:", err)
	}

	// Create the store
	dbQuery = chapter5.New(db)

	ctx := context.Background()

	createUserDb(ctx)

	if err != nil {
		os.Exit(1)
	}
}

func createUserDb(ctx context.Context) {
	//has the user been created
	u, _ := dbQuery.GetUserByName(ctx, "user@user")

	if u.UserName == "user@user" {
		log.Println("user@user exist...")
		return
	}
	log.Println("Creating user@user...")
	hashPwd, _ := pkg.HashPassword("password")
	_, err := dbQuery.CreateUsers(ctx, chapter5.CreateUsersParams{
		UserName:     "user@user",
		PassWordHash: hashPwd,
		Name:         "Dummy user",
	})
	if err != nil {
		log.Println("error getting user@dummyuser.domain ", err)
		os.Exit(1)
	}
}
