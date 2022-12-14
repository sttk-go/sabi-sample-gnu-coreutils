package main

import (
	"fmt"
)

type ConsoleDax struct {
}

func NewConsoleDax() ConsoleDax {
	return ConsoleDax{}
}

func (dax ConsoleDax) PrintUserName(userName string) {
	fmt.Println(userName)
}

func (dax ConsoleDax) PrintVersion() {
	fmt.Print(`whoami 1.0
Copyright (C) 2022 sttk-go project.
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Written by Takayuki Sato.
`)
}

func (dax ConsoleDax) PrintHelp() {
	fmt.Print(`Usage: ./whoami [OPTION]...
Print the user name associated with the current effective user ID.
Same as id -un.

      --help        display this help and exit
      --version     output version information and exit
`)
}
