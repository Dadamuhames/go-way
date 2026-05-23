# go-way

A lightweight database migration library for Go, supporting PostgreSQL and MySQL.

---

## Installation

```bash
go get github.com/Dadamuhames/go-way
```

---

## Usage

Specify the SQL dialect in the enviroment:

GOWAY_DIALECT=postgres (or mysql)

### Programmatic

Import the library and pass your `*sql.DB` instance:

```go
import (
    goway "github.com/Dadamuhames/go-way/migrate"
)

func main() {
    db := database.GetInstance() // your *sql.DB
    if err := goway.Migrate(db); err != nil {
        log.Fatal(err)
    }
}
```

### CLI (via Cobra)

You can also use go-way as a CLI command with [Cobra](https://github.com/spf13/cobra):

```go
import (
    "github.com/spf13/cobra"
    goway "github.com/Dadamuhames/go-way/migrate"
)

func main() {
    var rootCmd = &cobra.Command{Use: "app"}
    rootCmd.AddCommand(goway.MigrationCommand())
    rootCmd.Execute()
}
```

Then run migrations via:

```bash
./app migrate
```

---

## Configuration

When using as CLI, go-way reads database connection settings from environment variables:

| Variable           | Description          | Example          |
|--------------------|----------------------|------------------|
| `GOWAY_DB_HOST`     | Database host        | `localhost`      |
| `GOWAY_DB_PORT`     | Database port        | `5432`           |
| `GOWAY_DB_DATABASE` | Database name        | `my_db`          |
| `GOWAY_DB_USERNAME` | Database user        | `root`           |
| `GOWAY_DB_PASSWORD` | Database password    | `root_password`  |
| `GOWAY_DB_SCHEMA`   | Schema (PostgreSQL)  | `public`         |

Example `.env`:

```env
GOWAY_DB_HOST=localhost
GOWAY_DB_PORT=5432
GOWAY_DB_DATABASE=my_db
GOWAY_DB_USERNAME=mark
GOWAY_DB_PASSWORD=root_password
GOWAY_DB_SCHEMA=public
```

---

## Migration Files

Place your migration scripts in the following directory at the root of your project:

```
resource/db/migration/
```

### Naming Convention

| Prefix | Type        | Example                        |
|--------|-------------|--------------------------------|
| `V`    | Versioned   | `V1__create_users.sql`         |
| `R`    | Repeatable  | `R__seed_data.sql`             |

- **Versioned** migrations (`V`) run once and are tracked by version number.
- **Repeatable** migrations (`R`) re-run whenever their content changes (checksum-based).

### Example structure

```
resource/
└── db/
    └── migration/
        ├── V1__init.sql
        ├── V2__add_users_table.sql
        └── R__create_users.sql
```

---

## How It Works

go-way creates a `goway_schema_history` table in your database to track which migrations have been applied, their checksums, and whether they succeeded. This ensures migrations are never run twice and repeatable migrations are re-applied when changed.

---

## License

MIT
