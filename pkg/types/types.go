package types

type Config struct {
	UNAME          string `yaml:"DB_USERNAME"`
	Pwd            string `yaml:"DB_PASSWORD"`
	Host           string `yaml:"DB_HOST"`
	DB             string `yaml:"DB_NAME"`
	JWT_SECRET_KEY string `yaml:"JWT_SECRET_KEY"`
}

type User struct {
	UID   int
	UName string
	Hash  string
	Admin int
}

type Book struct {
	BookID   int    `json:"bookID"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type ListBook struct {
	TotalBooks int       `json:"totalBooks"`
	DiffBook   int       `json:"diffBooks"`
	TotalUsers int       `json:"totalUsers"`
	IssuedBook int       `json:"issuedBook"`
	Books      []Book    `json:"books"`
	Requests   []Request `json:"request"`
}

type ListRequest struct {
	Requests []Request `json:"request"`
}

type Request struct {
	RequestID int `json:"requestID"`
	UserID    int `json:"userID"`
	BookID    int `json:"bookID"`
	Status    int `json:"status"`
}
