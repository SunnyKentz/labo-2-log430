package main

import (
	LOGIS "caisse-app-scaled/caisse_app_scaled/centre_logistique/api"
	MAG "caisse-app-scaled/caisse_app_scaled/magasin/api"
	MERE "caisse-app-scaled/caisse_app_scaled/maison_mere/api"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	os.Setenv("DB_PASSWORD", "test_password")
	os.Setenv("DB_USER", "test_user")
	os.Setenv("DB_PORT", "5435")
	go MERE.NewApp()
	go MAG.NewApp()
	go LOGIS.NewApp()
	time.Sleep(time.Second * 2) // time for database to startup
	login("magasin", "http://localhost:8080")
	login("maison_mere", "http://localhost:8090")
	login("logistique", "http://localhost:8091")
	<-make(chan string)
}
func login(service, baseURL string) string {
	resp, err := http.PostForm(baseURL+"/login", map[string][]string{
		"username": {"Bob"},
		"caisse":   {"Caisse 1"},
		"magasin":  {"Magasin 1"},
	})
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if !(resp.StatusCode >= 400) {
		return "ok"
	}
	fmt.Println(service+" failed logged in", resp.StatusCode)
	return ""
}
