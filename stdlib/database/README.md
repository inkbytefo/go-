# GO-Minus Database Package

This package provides database access and management capabilities for the GO-Minus programming language. It includes a generic database interface and specific implementations for various database systems.

## Features

- Generic database interface
- SQL database support
- NoSQL database support
- Connection pooling
- Transaction management
- Query building
- Prepared statements
- Parameter binding
- Result set handling
- Error handling
- Migration support
- ORM (Object-Relational Mapping) capabilities

## Supported Databases

- **SQL**
  - MySQL
  - PostgreSQL
  - SQLite
  - Microsoft SQL Server
  - Oracle

- **NoSQL**
  - MongoDB
  - Redis
  - Cassandra
  - Elasticsearch

## Usage

### Basic SQL Usage

```go
import (
    "database/sql"
    "fmt"
)

func main() {
    // Open a database connection
    db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/mydb")
    if err != nil {
        fmt.Println("Error opening database:", err)
        return
    }
    defer db.Close()
    
    // Execute a query
    rows, err := db.Query("SELECT id, name FROM users WHERE age > ?", 18)
    if err != nil {
        fmt.Println("Error executing query:", err)
        return
    }
    defer rows.Close()
    
    // Iterate over the results
    for rows.Next() {
        var id int
        var name string
        
        err := rows.Scan(&id, &name)
        if err != nil {
            fmt.Println("Error scanning row:", err)
            continue
        }
        
        fmt.Printf("User: %d, %s\n", id, name)
    }
    
    // Check for errors during iteration
    if err := rows.Err(); err != nil {
        fmt.Println("Error during iteration:", err)
    }
}
```

### Using Transactions

```go
import (
    "database/sql"
    "fmt"
)

func transferMoney(db *sql.DB, fromAccount, toAccount int, amount float64) error {
    // Start a transaction
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    
    // Defer a rollback in case anything fails
    defer tx.Rollback()
    
    // Update the first account
    _, err = tx.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ?", amount, fromAccount)
    if err != nil {
        return err
    }
    
    // Update the second account
    _, err = tx.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, toAccount)
    if err != nil {
        return err
    }
    
    // Commit the transaction
    return tx.Commit()
}

func main() {
    db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/mydb")
    if err != nil {
        fmt.Println("Error opening database:", err)
        return
    }
    defer db.Close()
    
    err = transferMoney(db, 1, 2, 100.00)
    if err != nil {
        fmt.Println("Error transferring money:", err)
        return
    }
    
    fmt.Println("Money transferred successfully")
}
```

### Using Prepared Statements

```go
import (
    "database/sql"
    "fmt"
)

func main() {
    db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/mydb")
    if err != nil {
        fmt.Println("Error opening database:", err)
        return
    }
    defer db.Close()
    
    // Prepare a statement
    stmt, err := db.Prepare("INSERT INTO users(name, age) VALUES(?, ?)")
    if err != nil {
        fmt.Println("Error preparing statement:", err)
        return
    }
    defer stmt.Close()
    
    // Execute the statement multiple times
    users := []struct {
        name string
        age  int
    }{
        {"Alice", 25},
        {"Bob", 30},
        {"Charlie", 35},
    }
    
    for _, user := range users {
        _, err := stmt.Exec(user.name, user.age)
        if err != nil {
            fmt.Printf("Error inserting user %s: %v\n", user.name, err)
            continue
        }
        fmt.Printf("User %s inserted successfully\n", user.name)
    }
}
```

### Using the Query Builder

```go
import (
    "database/sql"
    "database/sql/querybuilder"
    "fmt"
)

func main() {
    db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/mydb")
    if err != nil {
        fmt.Println("Error opening database:", err)
        return
    }
    defer db.Close()
    
    // Create a query builder
    qb := querybuilder.New(db)
    
    // Build a SELECT query
    query := qb.Select("id", "name", "age").
              From("users").
              Where("age > ?", 18).
              OrderBy("name ASC").
              Limit(10)
    
    // Execute the query
    rows, err := query.Execute()
    if err != nil {
        fmt.Println("Error executing query:", err)
        return
    }
    defer rows.Close()
    
    // Iterate over the results
    for rows.Next() {
        var id int
        var name string
        var age int
        
        err := rows.Scan(&id, &name, &age)
        if err != nil {
            fmt.Println("Error scanning row:", err)
            continue
        }
        
        fmt.Printf("User: %d, %s, %d\n", id, name, age)
    }
}
```

### Using the ORM

```go
import (
    "database/sql/orm"
    "fmt"
)

// Define a model
class User {
    public:
        var ID int `db:"id,primarykey,autoincrement"`
        var Name string `db:"name"`
        var Age int `db:"age"`
        var Email string `db:"email"`
}

func main() {
    // Create an ORM instance
    o := orm.New("mysql", "user:password@tcp(localhost:3306)/mydb")
    
    // Register models
    o.Register(User{})
    
    // Create tables if they don't exist
    err := o.CreateTables()
    if err != nil {
        fmt.Println("Error creating tables:", err)
        return
    }
    
    // Create a new user
    user := User{
        Name: "Alice",
        Age: 25,
        Email: "alice@example.com",
    }
    
    // Insert the user
    err = o.Insert(&user)
    if err != nil {
        fmt.Println("Error inserting user:", err)
        return
    }
    
    fmt.Println("User inserted with ID:", user.ID)
    
    // Query users
    var users []User
    err = o.Query().Where("age > ?", 18).OrderBy("name").Limit(10).Find(&users)
    if err != nil {
        fmt.Println("Error querying users:", err)
        return
    }
    
    for _, u := range users {
        fmt.Printf("User: %d, %s, %d, %s\n", u.ID, u.Name, u.Age, u.Email)
    }
    
    // Update a user
    user.Age = 26
    err = o.Update(&user)
    if err != nil {
        fmt.Println("Error updating user:", err)
        return
    }
    
    // Delete a user
    err = o.Delete(&user)
    if err != nil {
        fmt.Println("Error deleting user:", err)
        return
    }
}
```

### Using NoSQL Databases

```go
import (
    "database/nosql"
    "fmt"
)

func main() {
    // Connect to MongoDB
    client, err := nosql.Connect("mongodb", "mongodb://localhost:27017")
    if err != nil {
        fmt.Println("Error connecting to MongoDB:", err)
        return
    }
    defer client.Close()
    
    // Get a database
    db := client.Database("mydb")
    
    // Get a collection
    collection := db.Collection("users")
    
    // Insert a document
    doc := map[string]interface{}{
        "name": "Alice",
        "age": 25,
        "email": "alice@example.com",
    }
    
    result, err := collection.InsertOne(doc)
    if err != nil {
        fmt.Println("Error inserting document:", err)
        return
    }
    
    fmt.Println("Document inserted with ID:", result.InsertedID)
    
    // Find documents
    filter := map[string]interface{}{
        "age": map[string]interface{}{
            "$gt": 18,
        },
    }
    
    cursor, err := collection.Find(filter)
    if err != nil {
        fmt.Println("Error finding documents:", err)
        return
    }
    defer cursor.Close()
    
    // Iterate over the results
    for cursor.Next() {
        var result map[string]interface{}
        err := cursor.Decode(&result)
        if err != nil {
            fmt.Println("Error decoding document:", err)
            continue
        }
        
        fmt.Printf("User: %v, %v, %v\n", result["name"], result["age"], result["email"])
    }
}
```

## Classes and Interfaces

### SQL Package

The `sql` package provides interfaces and classes for SQL database access.

```go
// DB represents a database connection.
class DB {
    // Open opens a database connection.
    static func Open(driverName, dataSourceName string) (*DB, error)
    
    // Close closes the database connection.
    func Close() error
    
    // Ping verifies a connection to the database is still alive.
    func Ping() error
    
    // Begin starts a transaction.
    func Begin() (*Tx, error)
    
    // Exec executes a query without returning any rows.
    func Exec(query string, args ...interface{}) (Result, error)
    
    // Query executes a query that returns rows.
    func Query(query string, args ...interface{}) (*Rows, error)
    
    // QueryRow executes a query that is expected to return at most one row.
    func QueryRow(query string, args ...interface{}) *Row
    
    // Prepare creates a prepared statement for later queries or executions.
    func Prepare(query string) (*Stmt, error)
}
```

### NoSQL Package

The `nosql` package provides interfaces and classes for NoSQL database access.

```go
// Client represents a NoSQL database client.
class Client {
    // Connect connects to a NoSQL database.
    static func Connect(driverName, connectionString string) (*Client, error)
    
    // Close closes the client connection.
    func Close() error
    
    // Database returns a handle to a database.
    func Database(name string) *Database
}
```

## Error Handling

The database package uses GO-Minus's exception handling mechanism for error handling.

```go
import (
    "database/sql"
    "fmt"
)

func main() {
    try {
        db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/mydb")
        if err != nil {
            throw err
        }
        defer db.Close()
        
        rows, err := db.Query("SELECT id, name FROM users WHERE age > ?", 18)
        if err != nil {
            throw err
        }
        defer rows.Close()
        
        for rows.Next() {
            var id int
            var name string
            
            err := rows.Scan(&id, &name)
            if err != nil {
                fmt.Println("Error scanning row:", err)
                continue
            }
            
            fmt.Printf("User: %d, %s\n", id, name)
        }
        
        if err := rows.Err(); err != nil {
            throw err
        }
    } catch (err) {
        fmt.Println("Database error:", err)
    }
}
```
