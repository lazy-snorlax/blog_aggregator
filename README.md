# Gator
A CLI tool for aggregating RSS feeds and viewing the posts.


## Installation
You will need to have Postgres and Go installed on your local system.
In the root directory of the project, run 
```bash
go install
```

## Config
Create a `.gatorconfig.json` file in your home directory with the following structure
```json
{
    "db_url": "postgres://username:@localhost:5432/database?sslmode=disable"
}
```

## Usage
Create a new user:
```bash
gator register <name>
```

Add new feed:
```bash
gator addfeed <url>
```

Start aggregator:
```bash
gator agg 30s
```

View saved posts:
```bash
gator browse [limit]
```

Here are a few other commands you'll need as well:
- `gator login <name>` - Log in as a user that already exists
- `gator users` - List all users
- `gator feeds` - List all feeds
- `gator follow <url>` - Follow a feed that already exists in the database
- `gator unfollow <url>` - Unfollow a feed that already exists in the database
