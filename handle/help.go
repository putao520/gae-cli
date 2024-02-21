package handle

import "fmt"

func ArgHelp() {
	println("Commands")
	fmt.Printf("\t%-20s %s\n", "run, r", "Run GAE's script")
	fmt.Printf("\t%-20s %s\n", "version, v", "Show the GAE version information")
	fmt.Printf("\t%-20s %s\n", "help, h", "Show the GAE help invformation")
	fmt.Printf("\t%-20s %s\n", "service, s", "Create a gsc mircoservice maven project")
	fmt.Printf("\t%-20s %s\n", "provider", "Manage the GAE provider")
	fmt.Printf("\t\t%-20s %s\n", "add", "(name, host)")
	fmt.Printf("\t\t%-20s %s\n", "remove", "(name)")
	fmt.Printf("\t\t%-20s %s\n", "update", "(name, \"host\",new_host)")
	fmt.Printf("\t\t%-20s %s\n", "update", "(name, \"authKey\",new_authKey)")
	fmt.Printf("\t\t%-20s %s\n", "default", "(name)")
	fmt.Printf("\t%-20s %s\n", "login", "Login user to default GAE provider")
	fmt.Printf("\t%-20s %s\n", "logout", "Logout loginned user to default GAE provider")
}
