# Game Review & Rating Platform API

## Core Idea
Think of a simplified Metacritic or IGN. Users can add games, write reviews, and assign ratings. Other users can see overall scores.

## CRUD Breakdown

- Create – Add new game entries, write new reviews.
- Read – List all games, view reviews and ratings.
- Update – Update review text or ratings, correct game details.
- Delete – Remove a game (admin function) or delete a user review.

## Possible Data Models

- Game: id, title, description, genre, releaseDate, developer
- Review: id, gameID, userID, rating, reviewText, dateCreated
- User: id, username, email, passwordHash

## Features to Consider

- Average Rating: Each game has an aggregate rating from all user reviews.
- Filtering & Sorting: Sort by highest-rated, newest, or most-reviewed.
- Comment on Reviews: Let users comment on or upvote/downvote reviews.
- Role-Based Access: Admins can hide or delete inappropriate reviews.