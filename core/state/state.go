package state

type SSHState struct {
	User   string
	Host   string
	Cwd    string
	Symbol string
}

type FTPState struct {
	Cwd      string
	LoggedIn bool
	User     string
}

type HTTPState struct {
	RemoteAddr string
	UserAgent  string
	Method     string
}

type SessionState struct {
	ID     string
	User   string
	Host   string
	Cwd    string
	Symbol string
}
