package main

import (
	"context"
	"database/sql"
	vlt "github.com/assets-atlas/cryptography"
	vaultidentity "github.com/assets-atlas/vault-identity"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hashicorp/vault-client-go"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func TokenValidate(vc vaultidentity.VaultClientWrapper, token string) (bool, error) {
	isValid, err := vaultidentity.ValidateToken(&vc, token)
	if err != nil {
		log.Fatalf("Token validation failed: %v", err)
		return false, err
	}

	return isValid, nil
}

func AuthenticationMiddleware(db *sql.DB, vc vaultidentity.VaultClientWrapper, vaultClient *vault.Client, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the Authorization header
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Validate the token with Vault
		valid, err := TokenValidate(vc, tokenString)
		if err != nil || !valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Parse the token without verification
		token, _, err := jwt.NewParser().ParseUnverified(tokenString, jwt.MapClaims{})
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			log.Println("Error parsing token:", err)
			return
		}

		// Extract claims from the token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Get the email claim
		email, ok := claims["email"].(string)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Get Vault key name from environment
		keyName := os.Getenv("VAULT_TRANSIT_KEY_NAME")

		// Prepare the data payload for encryption
		payload := []vlt.Data{
			{
				Plaintext: email,
				Context:   "email",
			},
		}

		// Encrypt the email using Vault
		encPayload, err := vlt.Encrypt(vaultClient, keyName, payload)
		if err != nil {
			log.Printf("Error encrypting data: %s\n", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Assert the encrypted payload as []string
		encryptedEmail, ok := encPayload.([]string)
		if !ok || len(encryptedEmail) == 0 {
			log.Println("Error: encrypted payload format incorrect")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Look up the user in the database using the encrypted email
		var userID int
		err = db.QueryRow("SELECT id FROM users WHERE email = $1", encryptedEmail[0]).Scan(&userID)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			log.Println("Database lookup error:", err)
			return
		}

		// Add the user ID to the request context
		ctx := context.WithValue(r.Context(), "userID", userID)

		// Call the next handler with the new context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
