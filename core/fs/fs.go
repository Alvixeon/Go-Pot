// Fake Filesystem working package
package fs

import (
	"path"
)

type FakeFile struct {
	Name    string
	IsDir   bool
	Size    int
	ModTime string
	Owner   string
	Perms   string
}

var FakeFS = map[string][]FakeFile{
	"/": {
		{Name: "bin", IsDir: true, Size: 4096, ModTime: "Feb 10 14:22", Owner: "root", Perms: "drwxr-xr-x"},
		{Name: "etc", IsDir: true, Size: 4096, ModTime: "Feb 10 14:22", Owner: "root", Perms: "drwxr-xr-x"},
		{Name: "home", IsDir: true, Size: 4096, ModTime: "Feb 10 14:23", Owner: "root", Perms: "drwxr-xr-x"},
		{Name: "root", IsDir: true, Size: 4096, ModTime: "Feb 10 14:21", Owner: "root", Perms: "drwx------"},
		{Name: "var", IsDir: true, Size: 4096, ModTime: "Feb 10 14:24", Owner: "root", Perms: "drwxr-xr-x"},
		{Name: "usr", IsDir: true, Size: 4096, ModTime: "Feb 10 14:20", Owner: "root", Perms: "drwxr-xr-x"},
		{Name: "tmp", IsDir: true, Size: 4096, ModTime: "Feb 10 14:25", Owner: "root", Perms: "drwxrwxrwt"},
	},

	"/bin": {
		{Name: "bash", Size: 1037520, IsDir: false, ModTime: "Jan 18 09:12", Owner: "root", Perms: "-rwxr-xr-x"},
		{Name: "ls", Size: 146432, IsDir: false, ModTime: "Jan 18 09:12", Owner: "root", Perms: "-rwxr-xr-x"},
		{Name: "cat", Size: 130112, IsDir: false, ModTime: "Jan 18 09:12", Owner: "root", Perms: "-rwxr-xr-x"},
		{Name: "grep", Size: 223840, IsDir: false, ModTime: "Jan 18 09:12", Owner: "root", Perms: "-rwxr-xr-x"},
	},

	"/etc": {
		{Name: "passwd", IsDir: false, Size: 2341, ModTime: "Mar 01 12:00", Owner: "root", Perms: "-rw-r--r--"},
		{Name: "shadow", IsDir: false, Size: 1450, ModTime: "Mar 01 12:00", Owner: "root", Perms: "-rw-r-----"},
		{Name: "hosts", IsDir: false, Size: 312, ModTime: "Feb 28 18:42", Owner: "root", Perms: "-rw-r--r--"},
		{Name: "ssh", IsDir: true, Size: 4096, ModTime: "Feb 28 18:40", Owner: "root", Perms: "drwxr-xr-x"},
		{Name: "systemd", IsDir: true, Size: 4096, ModTime: "Feb 28 18:41", Owner: "root", Perms: "drwxr-xr-x"},
	},

	"/etc/ssh": {
		{Name: "sshd_config", Size: 3241, IsDir: false, ModTime: "Feb 28 18:40", Owner: "root", Perms: "-rw-r--r--"},
		{Name: "ssh_config", Size: 1832, IsDir: false, ModTime: "Feb 28 18:40", Owner: "root", Perms: "-rw-r--r--"},
	},

	"/home": {
		{Name: "ethan", IsDir: true, Size: 4096, ModTime: "Apr 20 10:11", Owner: "ethan", Perms: "drwxr-xr-x"},
		{Name: "guest", IsDir: true, Size: 4096, ModTime: "Apr 19 08:33", Owner: "guest", Perms: "drwxr-xr-x"},
		{Name: "admin", IsDir: true, Size: 4096, ModTime: "Apr 18 22:12", Owner: "admin", Perms: "drwxr-x---"},
	},

	"/home/ethan": {
		{Name: "documents", IsDir: true, Size: 4096, ModTime: "Apr 21 14:01", Owner: "ethan", Perms: "drwxr-xr-x"},
		{Name: "downloads", IsDir: true, Size: 4096, ModTime: "Apr 21 16:42", Owner: "ethan", Perms: "drwxr-xr-x"},
		{Name: "projects", IsDir: true, Size: 4096, ModTime: "Apr 21 18:10", Owner: "ethan", Perms: "drwxr-xr-x"},
		{Name: "notes.txt", IsDir: false, Size: 1834, ModTime: "Apr 22 09:44", Owner: "ethan", Perms: "-rw-r--r--"},
		{Name: "todo.md", IsDir: false, Size: 921, ModTime: "Apr 22 09:45", Owner: "ethan", Perms: "-rw-r--r--"},
	},

	"/home/ethan/documents": {
		{Name: "uni", IsDir: true, Size: 4096, ModTime: "Apr 22 11:22", Owner: "ethan", Perms: "drwxr-xr-x"},
		{Name: "taxes.pdf", IsDir: false, Size: 81234, ModTime: "Mar 12 17:03", Owner: "ethan", Perms: "-rw-r--r--"},
		{Name: "resume.docx", IsDir: false, Size: 43122, ModTime: "Feb 09 10:12", Owner: "ethan", Perms: "-rw-r--r--"},
	},

	"/root": {
		{Name: "secret.txt", IsDir: false, Size: 128, ModTime: "Jan 03 03:11", Owner: "root", Perms: "-rw-------"},
		{Name: ".bash_history", IsDir: false, Size: 2048, ModTime: "Apr 23 02:11", Owner: "root", Perms: "-rw-------"},
		{Name: "keys", IsDir: true, Size: 4096, ModTime: "Jan 03 03:10", Owner: "root", Perms: "drwx------"},
	},

	"/root/keys": {
		{Name: "id_rsa", IsDir: false, Size: 3243, ModTime: "Jan 03 03:10", Owner: "root", Perms: "-rw-------"},
		{Name: "id_rsa.pub", IsDir: false, Size: 743, ModTime: "Jan 03 03:10", Owner: "root", Perms: "-rw-r--r--"},
	},

	"/var": {
		{Name: "log", IsDir: true, Size: 4096, ModTime: "Apr 22 23:12", Owner: "root", Perms: "drwxr-xr-x"},
		{Name: "www", IsDir: true, Size: 4096, ModTime: "Apr 22 23:10", Owner: "www-data", Perms: "drwxr-xr-x"},
		{Name: "lib", IsDir: true, Size: 4096, ModTime: "Apr 22 23:11", Owner: "root", Perms: "drwxr-xr-x"},
	},

	"/var/log": {
		{Name: "auth.log", Size: 23412, ModTime: "Apr 23 08:50", Owner: "syslog", Perms: "-rw-r-----"},
		{Name: "syslog", Size: 98431, ModTime: "Apr 23 08:50", Owner: "syslog", Perms: "-rw-r-----"},
		{Name: "kern.log", Size: 55412, ModTime: "Apr 23 08:50", Owner: "syslog", Perms: "-rw-r-----"},
	},

	"/var/www": {
		{Name: "html", IsDir: true, Size: 4096, ModTime: "Apr 21 20:11", Owner: "www-data", Perms: "drwxr-xr-x"},
		{Name: "index.html", Size: 1243, ModTime: "Apr 21 20:10", Owner: "www-data", Perms: "-rw-r--r--"},
	},
}

func Resolve(cwd, input string) string {
	if input == "" {
		return cwd
	}
	if input == "~" {
		return "/home"
	}
	if input == "/" {
		return "/"
	}
	if input == ".." {
		return path.Dir(cwd)
	}
	if input[0] == '/' {
		return path.Clean(input)
	}

	return path.Clean(path.Join(cwd, input))
}
