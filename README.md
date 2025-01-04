# Simple SQL-like Parser in Go

This project is a basic implementation of a SQL-like query parser written in Go. The parser is capable of recognizing and parsing `SELECT` statements, specifically identifying the columns to be selected, the `FROM` clause, and the table name. It handles basic tokenization and parsing to simulate how a real SQL parser might work.

## Features
- Tokenization of SQL query string into meaningful components (e.g., keywords, identifiers, commas).
- Basic support for `SELECT` queries, including columns and `FROM` clauses.
- Error handling for missing or incorrect SQL keywords.
  
## Technologies Used
- Go (Golang)
- `unicode` package for handling character types and spaces
- Simple string matching and parsing for SQL tokens

## How It Works

1. **Tokenizer**: Breaks down the input SQL query into tokens based on characters like keywords (`SELECT`, `FROM`), commas, and spaces.
2. **Parser**: Uses the generated tokens to parse the query according to the expected structure. The parser checks that `SELECT` is followed by identifiers (column names), and `FROM` is followed by the table name.
3. **Error Reporting**: If the structure doesn't match the expected pattern (e.g., missing `FROM`), an error is reported with a specific message and the position where the error occurred.

## Example Usage

### Valid Query
Input: 

```go
SELECT name, age FROM users
