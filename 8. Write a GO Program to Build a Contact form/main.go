package main 
import ( 
    "fmt" 
    "net/http" 
) 
func main() { 
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { 
        if r.Method != http.MethodPost { 
            http.ServeFile(w, r, "form.html") 
            return 
        } 
        r.ParseForm() 
        name := r.FormValue("name") 
        email := r.FormValue("email") 
        msg := r.FormValue("message") 
        fmt.Println("Form Submitted:") 
        fmt.Println("Name:", name) 
        fmt.Println("Email:", email) 
        fmt.Println("Message:", msg) 
        fmt.Fprintf(w, "Form Received Successfully!\n\nName: %s\nEmail: %s\nMessage: %s\n", name, email, msg) 
    }) 
    fmt.Println("Server running at: http://localhost:8080") 
    http.ListenAndServe(":8080", nil) 
}
