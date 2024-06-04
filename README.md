# HTML Parser and MySQL Inserter

This project is a Go application that parses HTML pages with pagination and inserts data into a MySQL database. It extracts content from `span` elements with the class `profile-description`, saves the full content into the `name` column and a slug version of the first four words into the `slug` column of the `categories` table.

## Features

- Parses HTML pages with pagination
- Extracts content from `span.profile-description`
- Inserts the full content into the `name` column
- Generates a slug from the first four words and inserts into the `slug` column
- Supports deleting entries with an ID greater than 5

## Prerequisites

- Go (Golang) installed
- MySQL database with `categories` table
- Go packages:
    - `github.com/PuerkitoBio/goquery`
    - `github.com/go-sql-driver/mysql`

## Setup

1. Clone the repository:

```sh
git clone https://github.com/mefistofeles00/go_parse_to_data
cd go_parse_to_data

go get github.com/PuerkitoBio/goquery
go get github.com/go-sql-driver/mysql

Configure your MySQL database connection in the main.go file:

dsn := "username:password@tcp(host:port)/database"

Ensure your MySQL categories table is set up correctly:

CREATE TABLE categories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL
);
