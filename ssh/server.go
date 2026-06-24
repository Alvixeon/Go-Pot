package ssh

// SSH ServerConfig setup, listener loop

import (
	"crypto/rand"
	"fmt"
	"log"

	"GO-POT/core/config"
	"GO-POT/ssh/shell"

	"github.com/gliderlabs/ssh"
)

type ConnectorValues struct {
	Username string
	Address  string
	Password string
}

func Listner() bool {
	cfg := config.Default()
	addr := fmt.Sprintf(":%d", cfg.Port)
	ssh.Handle(func(s ssh.Session) {
		userVars := ConnectorValues{
			Username: s.User(),
			Address:  s.RemoteAddr().String(),
		}
		b := make([]byte, 4)
		rand.Read(b)
		id := fmt.Sprintf("%x", b)
		s.Context().SetValue("sessionID", id)
		log.Printf("Connection with username: %s | Remote Address: %s", userVars.Username, userVars.Address)
		_, _, ok := s.Pty()
		log.Printf("[%s SESSION META] Pty=%v Subsystem=%v", id, ok, s.Subsystem())
		shell.HandleSession(s, cfg.SSHIdent, cfg.LoginBanner)

	})
	passwordOpt := ssh.PasswordAuth(func(ctx ssh.Context, password string) bool {
		log.Printf("Auth attempt — user: %s | password: %s | addr: %s",
			ctx.User(), password, ctx.RemoteAddr().String())
		return true // always accept, or false to block
	})

	log.Println("Served Identity: ", cfg.SSHIdent)
	log.Println("Served Banner: ", cfg.LoginBanner)
	log.Println("Starting ssh server on port: ", cfg.Port)

	log.Fatal(ssh.ListenAndServe(addr, nil, passwordOpt, ssh.Option(func(srv *ssh.Server) error {
		srv.Version = "OpenSSH_8.9p1 Ubuntu-3ubuntu0.3"
		return nil
	})))
	return true
}

func Start() {
	Listner()
}
