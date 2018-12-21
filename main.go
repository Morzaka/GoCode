package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	request(conn)

	//respond(conn)
}

func request(conn net.Conn) {
	i := 0 // just a little counter
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	m := strings.Fields(ln)[0]
	url := strings.Fields(ln)[1]
	fmt.Println(m)
	fmt.Println(url)

	//multiplexer
	if m == "GET" && url == "/" {
		index(conn)
	}
	if m == "GET" && url == "/about" {
		about(conn)
	}
	if m == "GET" && url == "/something" {
		something(conn)
	}
	if m == "GET" && url == "/apply" {
		apply(conn)
	}
	if m == "POST" && url == "/apply" {
		applyPost(conn)
	}
}

func index(conn net.Conn) {
	body := `
		<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title></title>
			</head>
			<body>
				<strong>Hello World</strong>
				<ul>
					<li><a href="/">Index</a></li>
					<li><a href="/about">About</a></li>
					<li><a href="/something">Something</a></li>
					<li><a href="/apply">Apply</a></li>
				</ul>
			</body>
		</html>`

	fmt.Fprint(conn, "HTTP /1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func about(conn net.Conn) {
	body := `
		<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title></title>
			</head>
			<body>
				<strong>ABOUT</strong>
				<ul>
					<li><a href="/">Index</a></li>
					<li><a href="/about">About</a></li>
					<li><a href="/something">Something</a></li>
					<li><a href="/apply">Apply</a></li>
				</ul>
				<p>Bla bla bla About</p>
			</body>
		</html>`

	fmt.Fprint(conn, "HTTP /1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func something(conn net.Conn) {
	body := `
		<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title></title>
			</head>
			<body>
				<h1>wow something</h1>
				<ul>
					<li><a href="/">Index</a></li>
					<li><a href="/about">About</a></li>
					<li><a href="/something">Something</a></li>
					<li><a href="/apply">Apply</a></li>
				</ul>
				<p>WOW SOMETHING</p>
			</body>
		</html>`

	fmt.Fprint(conn, "HTTP /1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func apply(conn net.Conn) {
	body := `
		<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title></title>
			</head>
			<body>
				<h1>Apply</h1>
				<ul>
					<li><a href="/">Index</a></li>
					<li><a href="/about">About</a></li>
					<li><a href="/something">Something</a></li>
					<li><a href="/apply">Apply</a></li>
				</ul>
				<p>Apply this ---> </p>
				<form method="post" action="/apply">
					<input type="Apply" value="Go to Google" />
				</form>
			</body>
		</html>`

	fmt.Fprint(conn, "HTTP /1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func applyPost(conn net.Conn) {
	body := `
		<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title></title>
			</head>
			<body>
				<h1>Congrats!!!</h1>
				<h2><a href="/">Go to Home page</a></h2>
			</body>
		</html>`

	fmt.Fprint(conn, "HTTP /1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
