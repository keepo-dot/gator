# Gator

Gator is a command-line RSS feed aggregator built with Go and PostgreSQL. It allows multiple users to register, follow RSS feeds, continuously fetch updates in the background, and browse posts directly from the terminal.

## Features

- Multi-user Support: Register and switch between different users.
- Feed Management: Add RSS feeds and track which user added them.
- Feed Subscriptions: Follow and unfollow feeds per user.
- Continuous Aggregation: Concurrently fetch and parse RSS feeds at configurable intervals.
- Post Storage: Save and deduplicate parsed feed posts into PostgreSQL.
- Post Browsing: View recent posts from all feeds followed by the active user.

## Prerequisites

- Go (1.20+ recommended)
- PostgreSQL
- Goose (for database migrations)
- sqlc (optional, only needed if modifying SQL queries)

## Installation

1. Clone the repository:

```bash
git clone <https://github.com/><your-username>/gator.git
cd gator
```

1. Install the CLI binary:
go install .

Ensure your PATH contains your Go bin directory (typically ~/go/bin).

## Configuration

Create a configuration file at ~/.gatorconfig.json pointing to your PostgreSQL database:

```json
{
  "db_url": "postgres://<user>:<password>@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

## Database Setup

Run the migrations using goose against your local PostgreSQL instance:

```bash

cd sql/schema
goose postgres "postgres://<user>:<password>@localhost:5432/gator?sslmode=disable" up
cd ../..

```

## CLI Usage

### User Management

- Register a new user:

```bash
gator register <username>
```

- Login as an existing user:

```bash
gator login <username>
```

- List all registered users:

```bash
gator users
```

### Feeds and Subscriptions

- Add a new feed:

```bash
gator addfeed <name> <url>
```

- List all feeds in the system:

```bash
gator feeds
```

- Follow an existing feed:

```bash
gator follow <url>
```

- List all feeds followed by the current user:

```bash
gator following
```

- Unfollow a feed:

```bash
gator unfollow <url>
```

### Aggregation and Reading

- Run the feed aggregator:
Starts a long-running scraper loop that fetches feeds based on the given duration string (e.g., 1m, 30s, 1h):

```bash
gator agg 1m
```

- Browse posts:
View the newest posts from feeds you follow. Defaults to 2 posts if no limit is specified:

```bash
gator browse [limit]
```

### Administration

- Reset database (deletes all records):

```bash
gator reset
```

## Tech Stack

- Language: Go
- Database: PostgreSQL
- Query Generation: sqlc
- Migrations: goose
- Data Format: XML / RSS 2.0
