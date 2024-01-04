package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"poll-app/ent"
	"poll-app/repository"
	"time"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

const port = 1234

type application struct {
	port               int
	persistenceContext *repository.PersistenceContext
	auth               Auth
	JWTSecret          string
	JWTIssuer          string
	JWTAudience        string
	CookieDomain       string
}

func createUser(app *application) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), 1)
	if err != nil {
		log.Fatalf("User cannot be created")
	}

	user, err := app.persistenceContext.UserPersistence.Client().User.
		Create().
		SetFirstName("admin").
		SetLastName("admin").
		SetEmail("admin@example.com").
		SetPassword(string(hashedPassword)).
		SetRole("admin").
		Save(ctx)

	if err != nil {
		log.Println("user has not been created: ", err)
	} else {
		log.Println("user has been created: ", user)
	}
}

func openConnection(connectionString string) *ent.Client {
	client, err := ent.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}

func main() {
	var host, username, password string
	var jwtSecret, jwtIssuer, jwtAudience, cookieDomain string

	flag.StringVar(
		&host,
		"host",
		"localhost",
		"db host",
	)

	flag.StringVar(
		&username,
		"username",
		"poll",
		"db user",
	)

	flag.StringVar(
		&password,
		"password",
		"pass123",
		"db user's password",
	)

	flag.StringVar(
		&jwtSecret,
		"jwt-secret",
		"veryverysecret",
		"signing secret",
	)

	flag.StringVar(
		&jwtIssuer,
		"jwt-issuer",
		"example.com",
		"signing-issuer",
	)

	flag.StringVar(
		&jwtAudience,
		"jwt-audience",
		"example.com",
		"signing audience ",
	)

	flag.StringVar(
		&cookieDomain,
		"cookie-domain",
		"localhost",
		"cookie domain",
	)

	flag.Parse()

	connectionString := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=poll sslmode=disable timezone=UTC connect_timeout=5", host, username, password)

	log.Printf("Connecting to the database - %s...", connectionString)
	client := openConnection(connectionString)

	app := application{
		port:               port,
		persistenceContext: repository.New(client),
		JWTSecret:          jwtSecret,
		JWTIssuer:          jwtIssuer,
		JWTAudience:        jwtAudience,
		CookieDomain:       cookieDomain,
	}

	app.auth = Auth{
		Issuer:        app.JWTIssuer,
		Audience:      app.JWTAudience,
		Secret:        app.JWTSecret,
		TokenExpiry:   time.Minute * 15,
		RefreshExpiry: time.Hour * 24,
		CookiePath:    "/",
		CookieDomain:  app.CookieDomain,
		CookieName:    "__Host-refresh-token",
	}

	log.Printf("Starting a server at port %d", app.port)
	createUser(&app)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", app.port), app.router()))
}
