package types

type Entry struct {
	User     string `json:"user"`
	Service  string `json:"service"`
	Entry    string `json:"entry"`
	Metadata string `json:"metadata"`
}

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	ID       int    `json:"id"`
}
