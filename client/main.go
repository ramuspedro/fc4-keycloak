package main

import (
	"context";
	"log";
	"net/http";
	"encoding/json";
	"golang.org/x/oauth2";
	oidc "github.com/coreos/go-oidc"
)

var (
	clientID = "app"
	clientSecret = "bc0ad164-8372-4bb9-8fea-07a51f88e658"
)

func main() {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "http://localhost:8080/auth/realms/demo")

	if err != nil {
		log.Fatal(err)
	}

	config := oauth2.Config {
		ClientID: 		clientID,
		ClientSecret: 	clientSecret,
		Endpoint: 		provider.Endpoint(),
		RedirectURL: 	"http://localhost:8081/auth/callback",
		Scopes: 		[]string{oidc.ScopeOpenID, "profile", "email", "roles"},
	}

	state := "magica"

	http.HandleFunc("/auth/callback", func(w http.ResponseWriter, r *http.Request) {
        if(r.URL.Query().Get("state") != state) {
            http.Error(w, "State did not match", http.StatusBadRequest)
            return
        }

        oauth2Token, err := config.Exchange(ctx, r.URL.Query().Get("code"))

        if err != nil {
            http.Error(w, "Failed to exchange code", http.StatusBadRequest)
            return
        }

        rawIDToken, ok := oauth2Token.Extra("id_token").(string)

        if !ok {
            http.Error(w, "No id_token", http.StatusBadRequest)
            return
        }

        resp := struct {
                Oauth2Token *oauth2.Token
                RawIDToken string
            } {
                oauth2Token, rawIDToken,
            }

        data, err := json.MarshalIndent(resp, "", "  ")

        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        w.Write(data)
    })

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, config.AuthCodeURL(state), http.StatusFound)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}