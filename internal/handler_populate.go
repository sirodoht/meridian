package internal

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-faker/faker/v4"
)

func (handlers *Handlers) HandlePopulate(
	w http.ResponseWriter,
	r *http.Request,
) {
	ctx := context.Background()

	go func() {
		fmt.Println("Populating database...")
		numUsers := 50
		numNotes := 100

		var minUserTime, maxUserTime, totalUserTime time.Duration
		var minNoteTime, maxNoteTime, totalNoteTime time.Duration

		users := make([]*User, 0, numUsers)

		// create x users
		for i := 0; i < numUsers; i++ {
			start := time.Now()
			res, err := handlers.api.Register(ctx, &RegisterRequest{
				Username:    faker.Username(),
				Password:    faker.Password(),
				Email:       faker.Email(),
				DisplayName: faker.Name(),
				Description: faker.Sentence(),
			})
			if err != nil {
				panic(err)
			}

			fmt.Printf("Creating user %v\n", res.User.Username)

			users = append(users, res.User)
			duration := time.Since(start)

			if minUserTime == 0 || duration < minUserTime {
				minUserTime = duration
			}
			if duration > maxUserTime {
				maxUserTime = duration
			}
			totalUserTime += duration
		}

		// create x notes per user
		for _, user := range users {
			fmt.Printf("Creating notes for user %v\n", user.Username)
			// start on 2023-01-01
			ts := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
			// add a random amount of time to the timestamp, up to 1 month
			ts = ts.Add(time.Duration(rand.Intn(60*60*24*30)) * time.Second)
			// calculate the interval between each note
			tsInterval := time.Since(ts) / time.Duration(numNotes)
			// create x notes
			for i := 0; i < numNotes; i++ {
				// start the timer for metrics
				start := time.Now()
				// add the interval to the timestamp
				ts = ts.Add(tsInterval)
				// create the note
				_, err := handlers.api.CreateNote(
					ctx,
					&CreateNoteRequest{
						Username:  user.Username,
						Content:   faker.Paragraph(),
						Timestamp: ts,
					},
				)
				if err != nil {
					panic(err)
				}

				// calculate the duration of the note creation
				duration := time.Since(start)
				if minNoteTime == 0 || duration < minNoteTime {
					minNoteTime = duration
				}
				if duration > maxNoteTime {
					maxNoteTime = duration
				}
				totalNoteTime += duration
			}
		}

		// Calculate average time for user generation and note creation
		avgUserTime := totalUserTime / time.Duration(numUsers)
		avgNoteTime := totalNoteTime / time.Duration(numNotes)

		// Print the min, max, and average time for user generation and note creation
		fmt.Printf("Populate complete\n\n")

		fmt.Printf("Created %v users\n", numUsers)
		fmt.Printf("Created %v notes per user\n\n", numNotes)

		fmt.Printf("User Generation - Min: %v, Max: %v, Avg: %v\n", minUserTime, maxUserTime, avgUserTime)
		fmt.Printf("Note Creation - Min: %v, Max: %v, Avg: %v\n", minNoteTime, maxNoteTime, avgNoteTime)
	}()
}
