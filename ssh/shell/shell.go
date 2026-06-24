// Fake shell emulation

package shell

import (
	"GO-POT/core/state"
	"GO-POT/ssh/command"

	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/gliderlabs/ssh"
)

func formatPath(state *state.SessionState) string {
	home := "/home/" + state.User
	if state.User == "root" {
		home = "/root"
	}
	if state.Cwd == home {
		return "~"
	}
	return state.Cwd
}

func HandleSession(s ssh.Session, ident string, banner string) {

	defer s.Close()

	fmt.Fprintf(s, "%s\n%s \n", ident, banner)

	state := state.SessionState{
		ID:     s.Context().Value("sessionID").(string),
		User:   s.User(),
		Host:   "ubuntu",
		Cwd:    "/", // default
		Symbol: "$",
	}

	if s.User() == "root" {
		state.Symbol = "#"
		state.Cwd = "/root"
	}
	shellName := fmt.Sprintf("%s@%s:~%s ", s.User(), state.Host, state.Symbol)

	io.WriteString(s, shellName)

	buf := make([]byte, 1)
	line := ""

	for {
		n, err := s.Read(buf)
		if err != nil {
			log.Printf("[SESSION %s] | Read error: %s", state.ID, err)
		}
		if n == 0 {
			log.Printf("[SESSION %s] | Connection closed by client", state.ID)
			return
		}

		b := buf[0]
		switch b {
		case '\r', '\n':
			io.WriteString(s, "\r\n")
			cmd := strings.TrimSpace(
				strings.ReplaceAll(line, "\r", ""),
			)
			log.Printf("[SESSION %s] | [CMD] user=%s cmd=%s | addr=%s", state.ID, state.User, cmd, s.RemoteAddr())
			line = ""

			time.Sleep(time.Millisecond * 300)
			command.Execute(s, &state, cmd)

			shellName = fmt.Sprintf("%s@%s:%s%s ", state.User, state.Host, formatPath(&state), state.Symbol)
			io.WriteString(s, shellName)
		case 127, '\b':
			if len(line)-1 < 0 {
				continue
			} else {
				line = line[:len(line)-1]
				io.WriteString(s, "\b \b")
			}

		default:
			// echo the character and add to line buffer
			line += string(b)
			s.Write([]byte{b})
		}
	}
}
