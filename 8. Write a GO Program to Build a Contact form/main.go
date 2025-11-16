package main

import (
	"fmt"
	"net/http"
)

const formHTML = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Simple Form Example</title>
</head>
<body>
    <h2>Submit Form</h2>
    <form action="/" method="post">
        <label>Name:</label><br>
        <input type="text" name="name" required><br><br>
        <label>Email:</label><br>
        <input type="email" name="email" required><br><br>
        <label>Message:</label><br>
        <textarea name="message" rows="4" cols="40" required></textarea><br><br>
        <input type="reset" value="Reset">
        <input type="submit" value="Submit">
    </form>
</body>
</html>
`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		if r.Method == http.MethodGet {
			fmt.Fprint(w, formHTML)
			return
		}
		r.ParseForm()
		name := r.FormValue("name")
		email := r.FormValue("email")
		message := r.FormValue("message")
		fmt.Fprintf(w, `
        <html>
        <body style="font-family: Arial; margin: 40px;">
            <h2 style="color: green;">Form Submitted Successfully!</h2>
            <p><b>Name:</b> %s</p>
            <p><b>Email:</b> %s</p>
            <p><b>Message:</b> %s</p>
            <br>
            <a href="/"> Go Back</a>
        </body>
        </html>
        `, name, email, message)
	})
	fmt.Println("🚀 Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
