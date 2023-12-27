package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"poll-app/ent"
	"poll-app/repository"
	"time"

	_ "github.com/lib/pq"
)

const port = 1234

type application struct {
	port               int
	persistenceContext *repository.PersistenceContext
	auth               *Auth
}

func createUser(app *application) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	user, err := app.persistenceContext.UserPersistence.Client().User.
		Create().
		SetFirstName("admin").
		SetLastName("admin").
		SetEmail("admin@example.com").
		SetPassword("password").
		SetRole("admin").
		Save(ctx)

	if err != nil {
		log.Println("user has not been created: ", err)
	} else {
		log.Println("user has been created: ", user)
	}
}

func openConnection() *ent.Client {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=poll sslmode=disable timezone=UTC connect_timeout=5")
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
	log.Println("Connecting to the database...")
	client := openConnection()

	app := application{
		port:               port,
		persistenceContext: repository.New(client),
	}

	log.Printf("Starting a server at port %d", app.port)
	createUser(&app)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", app.port), app.router()))
}
