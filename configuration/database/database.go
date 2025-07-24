package database

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Message struct {
	Topic string      `json:"topic"`
	Data  MessageData `json:"data"`
}

type MessageData struct {
	Notification *Notification  `json:"notification,omitempty"`
	Balance      *BalanceUpdate `json:"balance,omitempty"`
}

type Notification struct {
	ID                 int       `json:"id"`
	MemberID           int       `json:"member_id"`
	Subject            string    `json:"subject"`
	Context            string    `json:"context"`
	IconID             int       `json:"icon_id"`
	NotificationTypeID int       `json:"notification_type_id"`
	Description        string    `json:"description"`
	UpdatedAt          time.Time `json:"updated_at"`
	Action             string    `json:"action"`
}

type BalanceUpdate struct {
	MemberID   int       `json:"member_id"`
	CurrencyID int       `json:"currency_id"`
	Balance    float64   `json:"balance"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// DB INITIALIZATION
var (
	once    sync.Once
	db_pool *sqlx.DB
)

func initializeDB() {
	DATABASE_URL := os.Getenv("DATABASE_URL")
	var err_db error

	db_pool, err_db = sqlx.Connect("postgres", DATABASE_URL)
	if err_db != nil {
		log.Fatalln("Error connecting to the database:", err_db)
	}

	go listenForNotifications(DATABASE_URL)

	if err := db_pool.Ping(); err != nil {
		defer db_pool.Close()
		log.Fatalf("Failed to ping the database: %v", err)
	}

	// SET CONNECTION POOL SETTINGS
	db_pool.SetMaxIdleConns(10)
	db_pool.SetMaxOpenConns(10)
	db_pool.SetConnMaxLifetime(0)
}

func GetDB() *sqlx.DB {
	once.Do(initializeDB)
	return db_pool
}

// NOTIFICATION LISTENER
func listenForNotifications(dsn string) {
	listener := pq.NewListener(dsn, 10*time.Second, time.Minute, func(ev pq.ListenerEventType, err error) {
		if err != nil {
			log.Println("PostgreSQL listener error:", err)
		}
	})
	defer listener.Close()

	err := listener.Listen("member_notification_inserts_or_updates")
	if err != nil {
		log.Fatal("Failed to LISTEN on channel:", err)
	}

	fmt.Println("Listening for notifications...")

	for {
		select {
		case notification := <-listener.Notify:
			if notification != nil {
				fmt.Println("Received notification:", notification.Extra)

				var notification_message Message
				if err := json.Unmarshal([]byte(notification.Extra), &notification_message); err != nil {
					fmt.Printf("Error unmarshaling data: %s\n", err)
					continue
				}

				fmt.Println("Member ID:", notification_message.Data.Notification.MemberID)
			}

		case <-time.After(30 * time.Second):
			if err := listener.Ping(); err != nil {
				log.Println("PostgreSQL listener ping error:", err)
			}
		}
	}
}
