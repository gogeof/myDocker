package main

import(
	"os/exec"
	"os"
	"log"
)

func main(){
	cmd := exec.Command("sh")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run();err != nil {
		log.Fatal(err)
	}
}
