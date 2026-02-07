package main

import (
	"context"
	"fmt"
	"time"

	"github.com/CybrRonin/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("usage: %s <feed name> <feed URL>", cmd.Name)
	}

	currUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	values := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    currUser.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), values)
	if err != nil {
		return fmt.Errorf("unable to generate feed: %w", err)
	}

	fmt.Println("New RSS feed added...")
	printFeed(feed)
	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error retrieving feeds from database: %w\n", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("unable to find user with that ID: %w\n", err)
		}
		printFeed(feed)
		fmt.Printf("* User Name: %s\n\n", user.Name)
	}

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID: 			%d\n", feed.ID)
	fmt.Printf("* Created at: 	%v\n", feed.CreatedAt)
	fmt.Printf("* Updated at: 	%v\n", feed.UpdatedAt)
	fmt.Printf("* Feed Name at: %s\n", feed.Name)
	fmt.Printf("* Feed URL at: 	%s\n", feed.Url)
	fmt.Printf("* User ID: 		%d\n", feed.UserID)
}
