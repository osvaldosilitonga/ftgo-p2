package entitiy

type User struct {
	ID         int    `json:"id"`
	Email      string `json:"email" required:"true"`
	Password   string `json:"password" required:"true"`
	FullName   string `json:"full_name" required:"true"`
	Age        int    `json:"age" required:"true"`
	Occupation string `json:"occupation" required:"true"`
	Role       string `json:"role"`
}

type UserLogin struct {
	Email    string `json:"email" required:"true"`
	Password string `json:"password" required:"true"`
	Role     string `json:"role"`
}
