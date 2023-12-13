package migration

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"strings"
	"time"
)

func MigratePayments(from, to *pgxpool.Pool) {

	// Get user mapping
	userMapping := GetUserIdMapping()

	ctx := context.Background()

	// Get all payments from supabase
	rows, err := from.Query(ctx, `SELECT id, user_id, date, amount FROM payments`)
	if err != nil {
		log.Fatalln(err)
	}

	// Iterate over payments
	for rows.Next() {
		var (
			id     string
			userId string
			date   time.Time
			amount string
		)

		err = rows.Scan(&id, &userId, &date, &amount)
		if err != nil {
			log.Fatalln(err)
		}

		newUserId, ok := userMapping[userId]
		if !ok {
			log.Fatalf("user not found: %v", userId)
		}

		log.Printf("Payment: %v", id)

		// Insert into new database
		_, err = to.Exec(ctx, `
			INSERT INTO payments (id, user_id, date, amount)
			VALUES ($1, $2, $3, $4)
			ON CONFLICT DO NOTHING
		`, id, newUserId, date, amount)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func MigrateOpenaiUsage(from, to *pgxpool.Pool) {

	// Get user mapping
	userMapping := GetUserIdMapping()

	ctx := context.Background()

	rows, err := from.Query(ctx, `
		SELECT id, user_id, created_at, model, input_tokens, output_tokens
		FROM openai_usage`)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var (
			id     string
			userId string
			date   time.Time
			model  string
			input  string
			output string
		)

		err = rows.Scan(&id, &userId, &date, &model, &input, &output)
		if err != nil {
			log.Fatalln(err)
		}

		newUserId, ok := userMapping[userId]
		if !ok {
			log.Fatalf("user not found: %v", userId)
		}

		log.Printf("Usage: %v", id)

		_, err = to.Exec(ctx, `
			INSERT INTO openai_usages (id, user_id, created_at, model, input_tokens, output_tokens)
			VALUES ($1, $2, $3, $4, $5, $6)
			ON CONFLICT DO NOTHING
		`, id, newUserId, date, model, input, output)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func MigrateCollections(from, to *pgxpool.Pool) {

	// Get user mapping
	userMapping := GetUserIdMapping()

	ctx := context.Background()

	rows, err := from.Query(ctx, `
		SELECT id, user_id, name
		FROM collections`)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var (
			id     string
			userId string
			name   string
		)

		err = rows.Scan(&id, &userId, &name)
		if err != nil {
			log.Fatalln(err)
		}

		newUserId, ok := userMapping[userId]
		if !ok {
			log.Fatalf("user not found: %v", userId)
		}

		log.Printf("Collection: %v %v", id, name)

		_, err = to.Exec(ctx, `
			INSERT INTO collections (id, user_id, name)
			VALUES ($1, $2, $3)
			ON CONFLICT DO NOTHING
		`, id, newUserId, name)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func MigrateChatMessages(from, to *pgxpool.Pool) {

	// Get user mapping
	userMapping := GetUserIdMapping()

	ctx := context.Background()

	rows, err := from.Query(ctx, `
		SELECT id, user_id, created_at, collection_id, prompt, completion
		FROM chat_message`)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var (
			id     string
			userId string
			date   time.Time
			collId string
			prompt string
			compl  string
		)

		err = rows.Scan(&id, &userId, &date, &collId, &prompt, &compl)
		if err != nil {
			log.Fatalln(err)
		}

		newUserId, ok := userMapping[userId]
		if !ok {
			log.Fatalf("user not found: %v", userId)
		}

		log.Printf("Chat message: %v", id)

		_, err = to.Exec(ctx, `
			INSERT INTO chat_messages (id, user_id, created_at, collection_id, prompt, completion)
			VALUES ($1, $2, $3, $4, $5, $6)
			ON CONFLICT DO NOTHING
		`, id, newUserId, date, collId, prompt, compl)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func MigrateDocuments(from, to *pgxpool.Pool) {

	// Get user mapping
	userMapping := GetUserIdMapping()

	ctx := context.Background()

	rows, err := from.Query(ctx, `
		SELECT id, user_id, filename, path, collection_id
		FROM documents`)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var (
			id       string
			userId   string
			filename string
			path     string
			collId   string
		)

		err = rows.Scan(&id, &userId, &filename, &path, &collId)
		if err != nil {
			log.Fatalln(err)
		}

		newUserId, ok := userMapping[userId]
		if !ok {
			log.Fatalf("user not found: %v", userId)
		}

		newPath := "documents/" + strings.Replace(path, userId, newUserId, 1)
		//log.Printf("Documents: %v %s --> %s", id, path, newPath)
		log.Printf("Documents: %v", id)

		_, err = to.Exec(ctx, `
			INSERT INTO documents (id, user_id, filename, path, collection_id)
			VALUES ($1, $2, $3, $4, $5)
			ON CONFLICT DO NOTHING
		`, id, newUserId, filename, newPath, collId)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
