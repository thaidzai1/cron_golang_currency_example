package main

import (
	"fmt"
	"os"

	editpricejobs "gido.vn/gic/cron/auto-edit-price"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("env.yml")
	fmt.Println(os.Getenv("HOST"))
	editpricejobs.Main()
}
