package entity

type SecretCard struct {
	Name     string
	URL      string
	Login    string
	Password string
	Text     string
	Files    map[string]string
	Meta     map[string]string
}

// type JSONSecretCard struct {
// 	Name     string
// 	URL      url.URL
// 	Login    string
// 	Password string
// 	Text     string
// 	File     []byte
// 	Meta     map[string]string
// }
