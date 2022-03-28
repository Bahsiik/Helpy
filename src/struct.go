package src

type Users struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	FirstName  string `json:"first_name"`
	Password   string `json:"password"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Admin      bool   `json:"admin"`
	Status     bool   `json:"status"`
	Profil_pic string `json:"profil_pic"`
}

type Subjects struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	User_id     int    `json:"user_id"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}

type Topics struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Subject_id  int    `json:"subject_id"`
}

type Tags struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Topic_id    int    `json:"topic_id"`
}

type Replys struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Topic_id    int    `json:"topic_id"`
	User_id     int    `json:"user_id"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}

type Notifications struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	User_id     int    `json:"user_id"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}
