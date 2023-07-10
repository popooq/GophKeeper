package types

type Entry struct {
	User     string `json:"user"`
	Service  string `json:"service"`
	Entry    string `json:"entry"`
	Metadata string `json:"metadata"`
}
