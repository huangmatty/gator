# Gator

This is a WIP.
A multi-player command line tool for aggregating RSS feeds and viewing the posts.

## Installation

Make sure you have the latest [Go toolchain](https://golang.org/dl/) installed as well as a local Postgres database. You can then install `gator` with:

```bash
go install ...
```

## Config

Create a `.gatorconfig.json` file in your home directory with the following structure:

```json
{
  "db_url": "postgres://username:@localhost:5432/database?sslmode=disable"
}
```

Replace the values with your database connection string.

## Usage

- `gator register <name>` - Add a new user
- `gator login <name>` - Log in as a user that already exists
- `gator users` - List all users
- `gator feeds` - List all feeds
- `gator addfeed <url>` - Add a feed under the current user
- `gator agg <duration>` - Start the aggregator with a custom time interval
- `gator browse <optional_num_of_posts>` - View the current user's posts (defaults to 2 posts)
- `gator follow <url>` - Follow a feed that already exists in the database
- `gator unfollow <url>` - Unfollow a feed that already exists in the database