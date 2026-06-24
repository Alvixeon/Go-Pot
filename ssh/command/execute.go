// Command Executor

package command

import (
	"GO-POT/core/fs"
	"GO-POT/core/state"
	"fmt"
	"io"
	"strings"

	"github.com/gliderlabs/ssh"
)

func Execute(s ssh.Session, state *state.SessionState, cmd string) {

	if strings.HasPrefix(cmd, "cd ") || cmd == "cd" {
		fields := strings.Fields(cmd)
		dTarget := ""
		if len(fields) > 1 {
			dTarget = fields[1]
		}
		runes := []rune(cmd)
		if len(runes) > 2 {
			if runes[2] != ' ' {
				fmt.Fprintf(s, "bash: %s: command not found\n", cmd)
				return
			}
		}

		newTarget := fs.Resolve(state.Cwd, dTarget)

		if _, ok := fs.FakeFS[newTarget]; ok {
			state.Cwd = newTarget
		} else {
			io.WriteString(s, "No such file or directory\n")
		}

	} else {
		switch cmd {
		case "":
			io.WriteString(s, "\r\n")
		case "whoami":
			if state.User != "root" {
				fmt.Fprintf(s, "%s\n", state.User)
			} else {
				io.WriteString(s, "root\n")
			}
		case "su":
			if state.User != "root" {
				io.WriteString(s, "Password: ")
				buff := make([]byte, 1)
				suLine := ""
				for {
					n, err := s.Read(buff)
					if err != nil || n == 0 {
						return
					}
					b := buff[0]
					if b == '\r' || b == '\n' {
						break
					}
					suLine += string(b)
				}
				state.User = "root"
				state.Symbol = "#"
				state.Cwd = "/root"
				io.WriteString(s, "\n")
			}
			io.WriteString(s, "")
		case "id":
			if state.User != "root" {
				fmt.Fprintf(s, "uid=1000(%s) gid=1000(%s) groups=1000(%s),27(sudo),3(sys),90(network),957(nopasswdlogin),979(rfkill),982(users),983(video),985(storage),989(lp),995(audio),998(wheel)\n",
					state.User, state.User, state.User)
			} else {
				io.WriteString(s, "uid=0(root) gid=0(root) groups=0(root)\n")
			}
		case "uname -a":
			io.WriteString(s, "Linux ubuntu 5.15.0-1034-aws #38-Ubuntu SMP Mon Mar 20 15:41:27 UTC 2023 x86_64 x86_64 x86_64 GNU/Linux\r\n")
		case "ls":
			if files, ok := fs.FakeFS[state.Cwd]; ok {
				names := make([]string, len(files))
				for i, f := range files {
					names[i] = f.Name
				}
				fmt.Fprintln(s, strings.Join(names, "  "))
			} else {
				io.WriteString(s, "\n")
			}

			//fmt.Fprintln(s, strings.Join(home, "  "))
		case "exit":
			s.Close()
			return
		default:
			fmt.Fprintf(s, "bash: %s: command not found\n", cmd)
		}
	}
}
