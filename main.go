main.go package main
import (
    "log"
    "net/http"
)
func homePage(w http.ResponseWriter, r *http.Request) {
    // Render the home html page from static folder
    http.ServeFile(w, r, "static/home.html")
}
func coursePage(w http.ResponseWriter, r *http.Request) {
    // Render the course html page
    http.ServeFile(w, r, "static/courses.html")
}
func aboutPage(w http.ResponseWriter, r *http.Request) {
    // Render the about html page
    http.ServeFile(w, r, "static/about.html")
}
func contactPage(w http.ResponseWriter, r *http.Request) {
    // Render the contact html page
    http.ServeFile(w, r, "static/contact.html")
}
func main() {
    http.HandleFunc("/home", homePage)
    http.HandleFunc("/courses", coursePage)
    http.HandleFunc("/about", aboutPage)
    http.HandleFunc("/contact", contactPage)
    err := http.ListenAndServe("0.0.0.0:8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}
main test.go // Test the main function
package main
import (
    "net/http"
    "net/http/httptest"
    "testing"
)
func TestMain(t *testing.T) {
    req, err := http.NewRequest("GET", "/home", nil)
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(homePage)
    handler.ServeHTTP(rr, req)
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }
    // Just verify the code not html content
    expected := "text/html; charset=utf-8"
    if contentType := rr.Header().Get("Content-Type"); contentType != expected {
        t.Errorf("handler returned unexpected content type: got %v want %v",
            contentType, expected)
    }
}
