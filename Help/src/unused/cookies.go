package unused

//
//var sessions = make(map[string]string)
//
//func Signin(w http.ResponseWriter, r *http.Request) {
//	username := r.FormValue("username")
//	password := r.FormValue("password")
//	if username != "username" || password != "password" {
//		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
//	}
//	sessionID := generateSessionID()
//	sessions[sessionID] = username
//	cookie := &http.Cookie{
//		Name:     "sessionID",
//		Value:    sessionID,
//		Path:     "/",
//		HttpOnly: true,
//		Expires:  time.Now().Add(time.Hour),
//	}
//	http.SetCookie(w, cookie)
//	http.Redirect(w, r, "/", http.StatusFound)
//}
//
//func generateSessionID() string {
//	return "sessionID"
//}
//
//func Refresh(w http.ResponseWriter, r *http.Request) {
//	cookie, err := r.Cookie("sessionID")
//	if err != nil {
//		http.Error(w, "Not logged in", http.StatusUnauthorized)
//		return
//	}
//	sessionID := cookie.Value
//	if _, ok := sessions[sessionID]; !ok {
//		http.Error(w, "Not logged in", http.StatusUnauthorized)
//		return
//	}
//	sessions[sessionID] = "username"
//	cookie.Expires = time.Now().Add(time.Hour)
//	http.SetCookie(w, cookie)
//	http.Redirect(w, r, "/8080", http.StatusFound)
//
//}
//func Signout(w http.ResponseWriter, r *http.Request) {
//	cookie, err := r.Cookie("sessionID")
//	if err != nil {
//		http.Error(w, "Not logged in", http.StatusUnauthorized)
//		return
//	}
//	delete(sessions, cookie.Value)
//	cookie.MaxAge = -1
//	http.SetCookie(w, cookie)
//	http.Redirect(w, r, "/", http.StatusFound)
//}
