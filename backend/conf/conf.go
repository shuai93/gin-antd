package conf

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func Init() {
	_ = godotenv.Load("conf/dev.env")
	fmt.Println("name: ", os.Getenv("name"))
	fmt.Println("age: ", os.Getenv("age"))
	fmt.Println("city: ", os.Getenv("city"))

}
