// Load from file/env (port, host key path, etc.)

package config

type Config struct {
	Port        int
	SSHIdent    string
	LoginBanner string
	Hostname    string
}

func Default() Config {
	return Config{
		Port:        2222,
		SSHIdent:    "OpenSSH_8.9p1 Ubuntu-3ubuntu0.3",
		LoginBanner: "Welcome to Ubuntu 22.04 LTS (GNU/Linux 5.15.0-1034-aws x86_64)",
		Hostname:    "ubuntu",
	}
}
