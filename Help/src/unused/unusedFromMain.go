package unused

//func generateSessionId() string {
//	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
//	b := make([]rune, 32)
//	for i := range b {
//		b[i] = letters[rand.Intn(len(letters))]
//	}
//	return string(b)
//}

//func FirstCookies(w http.ResponseWriter, r *http.Request) {
//	cookie, err := r.Cookie("Test-cookie")
//	fmt.Println("Cookie:", cookie, "Error:", err)
//	if err != nil {
//		fmt.Println("Cookies not found")
//		cookie = &http.Cookie{
//			Name:     "Test-cookie",
//			Value:    "Test-value",
//			HttpOnly: true,
//		}
//		http.SetCookie(w, cookie)
//	}
//	erro := tmpl.ExecuteTemplate(w, "register.html", nil)
//	if erro != nil {
//		return
//	}
//}

//func CompareCheck(w http.ResponseWriter, r *http.Request) {
//	r.ParseForm()
//	username := r.Form.Get("username")
//	password := r.Form.Get("password")
//	var user Users
//	var err error
//	err = db.QueryRow("SELECT * FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password)
//	if err != nil {
//		fmt.Fprintln(w, "Username or password is incorrect")
//		return
//	}
//	if user.Password != password {
//		fmt.Fprintln(w, "Username or password is incorrect")
//		return
//	}
//	fmt.Fprintln(w, "You are logged in")
//	http.Redirect(w, r, "/signin", http.StatusSeeOther)
//}

//func CompareForm(w http.ResponseWriter, r *http.Request) {
//	if r.Method != http.MethodPost {
//		return
//	}
//	if r.FormValue("password") != r.FormValue("passwordo") {
//		http.Redirect(w, r, "/", http.StatusSeeOther)
//	} else {
//		albID, err := addUser(Users{
//			Username:  r.FormValue("username"),
//			Password:  r.FormValue("password"),
//			Passwordo: r.FormValue("passwordo"),
//			Email:     r.FormValue("email"),
//		})
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Printf("ID of added Users: %v\n", albID)
//	}
//}
