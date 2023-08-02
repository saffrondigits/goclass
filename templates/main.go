package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server listening on http://localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := Person{
		Name: "Gopher",
		Age:  12,
	}

	// HTML template with if-else condition to display personalized message
	htmlTemplate := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Go Web Page</title>
		</head>
		<body>
			{{if .Age}}
				<h1>Hello, {{.Name}}!</h1>
					{{if ge .Age 18}}
						<p>You are an adult.</p>
					{{else}}
						<p>You are a minor.</p>
					{{end}}
			{{else}}
				<p>Invalid age.</p>
			{{end}}
		</body>
		</html>
	`

	// Parse the template
	tmpl, err := template.New("example").Parse(htmlTemplate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template with the data
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
