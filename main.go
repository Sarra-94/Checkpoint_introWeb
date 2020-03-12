package main
import (
     helper "./helper"
	"fmt"
	"net/http"
)
var (
	uName      = ""
	email      = ""
	pwd        = ""
	pwdConfirm = ""
)
func signUp(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	uName = r.FormValue("username")     // Data from the form
	email = r.FormValue("email")        // Data from the form
	pwd = r.FormValue("password")       // Data from the form
	pwdConfirm = r.FormValue("confirm") // Data from the form
	// Empty data checking
	uNameCheck := helper.IsEmpty(uName)
	emailCheck := helper.IsEmpty(email)
	pwdCheck := helper.IsEmpty(pwd)
	pwdConfirmCheck := helper.IsEmpty(pwdConfirm)
	if uNameCheck || emailCheck || pwdCheck || pwdConfirmCheck {
		fmt.Fprintf(w, "ErrorCode is -10 : There is empty data.")
		return
	}
	if pwd == pwdConfirm {
		// Save to database (username, email and password)
		fmt.Fprintln(w, "Registration successful.")
	} else {
		fmt.Fprintln(w, "Password information must be the same.")
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email = r.FormValue("email")  // Data from the form
	pwd = r.FormValue("password") // Data from the form
	// Empty data checking
	emailCheck := helper.IsEmpty(email)
	pwdCheck := helper.IsEmpty(pwd)
	if emailCheck || pwdCheck {
		fmt.Fprintf(w, "ErrorCode is -10 : There is empty data.")
		return
	}
	dbPwd := "1234!*."                   // DB simulation
	dbEmail := "test.test@mail.com" // DB simulation
	if email == dbEmail && pwd == dbPwd {
		fmt.Fprintln(w, "Login succesful!")
	} else {
		fmt.Fprintln(w, "Login failed!")
	}
}
func main() {

	mux := http.NewServeMux()

	// Signup
	mux.HandleFunc("/signup", signUp)

	// Login
	mux.HandleFunc("/login", login)

	http.ListenAndServe(":8080", mux)
}


