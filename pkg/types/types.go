package types

type Config struct {
	UserNAME       string `yaml:"DB_USERNAME"`
	Password       string `yaml:"DB_PASSWORD"`
	Host           string `yaml:"DB_HOST"`
	DB             string `yaml:"DB_NAME"`
	JWT_SECRET_KEY string `yaml:"JWT_SECRET_KEY"`
}

type User struct {
	UserID   int
	UserName string
	Hash     string
	Admin    int
}

type Book struct {
	BookID   int    `json:"bookID"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Issued   int    `json:"issued"`
}

type ListBook struct {
	TotalBooks    int       `json:"totalBooks"`
	DifferentBook int       `json:"differentBooks"`
	TotalUsers    int       `json:"totalUsers"`
	IssuedBook    int       `json:"issuedBook"`
	Books         []Book    `json:"books"`
	Requests      []Request `json:"request"`
	IssuedAmount  []Request `json:"issuedAmount"`
	UserRequest   []Request `json:"userRequest"`
}

type ListRequest struct {
	Requests []Request `json:"request"`
}

type Request struct {
	RequestID int    `json:"requestID"`
	UserID    int    `json:"userID"`
	BookID    int    `json:"bookID"`
	Status    int    `json:"status"`
	BookName  string `json:"bookName"`
}
