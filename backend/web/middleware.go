package web

import (
	"context"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

func FirestoreMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()

		location, exists := os.LookupEnv("FIRESTORE_SERVICE_ACCOUNT_PATH")

		if !exists {
			log.Fatalf("FIRESTORE_SERVICE_ACCOUNT_PATH not set in environment. \n")
			return c.String(http.StatusInternalServerError, "Improperly configured server.")
		}

		opt := option.WithCredentialsFile(location)

		app, err := firebase.NewApp(ctx, nil, opt)
		if err != nil {
			log.Fatalf("Error initializing Firebase app: %v", err)
		}

		firestore, err := app.Firestore(ctx)
		if err != nil {
			log.Fatalf("Error creating Firestore client: %v", err)
		}
		defer firestore.Close()

		auth, err := app.Auth(ctx)
		if err != nil {
			log.Fatalf("error getting Auth client: %v\n", err)
		}

		c.Set("firestore", firestore)
		c.Set("auth", auth)

		return next(c)
	}
}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Get("auth").(*auth.Client)

		if c.Request().Header.Get("Authorization") == "" {
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}

		token, err := auth.VerifyIDToken(context.Background(), c.Request().Header.Get("Authorization"))
		c.Set("userID", token.UID)
		if err != nil {
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}

		return next(c)
	}
}
