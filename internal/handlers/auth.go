package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"savannah/internal/models"

	"github.com/go-chi/render"
	"golang.org/x/oauth2"
)

var (
	clientID     = "f10ce7eea34b581908c3"
	clientSecret = "b4d3faac2b371de41c53b20ae4ce16fb218a2308"
	redirectURL  = "http://192.168.1.229:8002/callback"
)

func GitAuth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		oauth2Config := oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://github.com/login/oauth/authorize",
				TokenURL: "https://github.com/login/oauth/access_token",
			},
			RedirectURL: redirectURL,
			Scopes:      []string{"user:email"},
		}
		authURL := oauth2Config.AuthCodeURL("", oauth2.AccessTypeOffline)
		http.Redirect(w, r, authURL, http.StatusFound)
	}
}

func (repo *HandlerRepo) HandleGitHubCallback() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		log.Printf("Returned code %v", code)

		if code == "" {
			http.Error(w, "Missing code parameter", http.StatusBadRequest)
			return
		}

		oauth2Config := oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Endpoint: oauth2.Endpoint{
				TokenURL: "https://github.com/login/oauth/access_token",
			},
			RedirectURL: redirectURL,
		}
		token, err := oauth2Config.Exchange(context.Background(), code)
		if err != nil {
			http.Error(w, "Failed to exchange code for token", http.StatusInternalServerError)
			log.Printf("OAuth2 token exchange failed: %v", err)
			return
		}

		// Use the access token to fetch the user's email
		client := oauth2Config.Client(context.Background(), token)
		resp, err := client.Get("https://api.github.com/user/emails")
		if err != nil {
			http.Error(w, "Failed to get user's email", http.StatusInternalServerError)
			log.Printf("GitHub API request failed: %v", err)
			return
		}
		defer resp.Body.Close()

		// Read the email from the response
		var emails []struct {
			Email string `json:"email"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&emails); err != nil {
			http.Error(w, "Failed to parse email response", http.StatusInternalServerError)
			log.Printf("JSON decoding failed: %v", err)
			return
		}

		if len(emails) > 0 {
			fmt.Fprintf(w, "Logged in with GitHub as %s and your token %s, use this to login", emails[0].Email, token.AccessToken)
			err := repo.store.UserCreate(r.Context(), models.User{
				Username:  emails[0].Email,
				AuthToken: token.AccessToken,
			})
			if err != nil {
				log.Println(err)
				e := NewBadRequestError(ErrorProcessingRequest)
				render.Respond(w, r, e)
				return
			}

		} else {
			fmt.Fprintf(w, "Failed to retrieve email from GitHub")
		}
	}
}
