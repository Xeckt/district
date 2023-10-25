package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var db *sql.DB

func connectToDB() *sql.DB {
	connStr := "user=district password=abcd dbname=district sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func updateMessageCount(db *sql.DB, serverID, userID string) {

	// Ensure the server exists
	_, err := db.Exec("INSERT INTO servers (server_id, server_name) VALUES ($1, $2) ON CONFLICT (server_id) DO NOTHING;", serverID, serverID)
	if err != nil {
		log.Fatal(err)
	}

	// Ensure the user exists
	_, err = db.Exec("INSERT INTO users (user_id, username) VALUES ($1, $2) ON CONFLICT (user_id) DO NOTHING;", userID, userID)
	if err != nil {
		log.Fatal(err)
	}

	// Insert or update message count
	_, err = db.Exec(`
		INSERT INTO metrics (server_id, user_id, metric_date, message_count)
		VALUES ($1, $2, CURRENT_DATE, 1)
		ON CONFLICT (server_id, user_id, metric_date)
		DO UPDATE SET message_count = metrics.message_count + 1;
	`, serverID, userID)
	if err != nil {
		log.Fatal(err)
	}
}

func queryServer(s *discordgo.Session, m *discordgo.MessageCreate) {
	var cnt string
	var ok bool
	if cnt, ok = strings.CutPrefix(m.Content, "!server"); !ok {
		return
	}
	cnt = strings.TrimSpace(cnt)
	days, _ := strconv.Atoi(cnt)

	query := fmt.Sprintf("SELECT SUM(message_count) FROM metrics WHERE server_id = $1 AND metric_date BETWEEN CURRENT_DATE - interval '%v days' AND CURRENT_DATE;", days)

	var count int
	err := db.QueryRow(query, m.GuildID).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%v", count))
}

func queryUser(s *discordgo.Session, m *discordgo.MessageCreate) {
	var cnt string
	var ok bool
	if cnt, ok = strings.CutPrefix(m.Content, "!user"); !ok {
		return
	}
	cnt = strings.TrimSpace(cnt)
	days, _ := strconv.Atoi(cnt)

	query := fmt.Sprintf("SELECT SUM(message_count) FROM metrics WHERE server_id = $1 AND user_id = $2 AND metric_date BETWEEN CURRENT_DATE - interval '%v days' AND CURRENT_DATE;", days)

	var count int
	err := db.QueryRow(query, m.GuildID, m.Author.ID).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%v", count))
}

func init() {
	db = connectToDB()
}
