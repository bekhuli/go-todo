# Go Todo

## Introduction

A simple todolist application written in Go 

## Requirements
* MySQL installed
* Go installed

## Installation

* Clone this repo 

```bash
git clone https://github.com/bekhuli/go-todo.git
```

* Change Directory

```bash
cd go-todo
```

* Initiate `.env` file

```bash
cp .env.example .env
```

* Modify `.env` file with your correct database credentials and desired Port

## Usage

To run this application, execute:

```bash
go run cmd/todo/main.go
```

You should be able to access this application at `http://127.0.0.1:8080`

>**NOTE**<br>
>If you modified the port in the `.env` file, you should access the application for the port you set
