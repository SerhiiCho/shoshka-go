package system

import (
	"github.com/SerhiiCho/shoshka_go/utils"
	"github.com/joho/godotenv"
)

/*
LoadEnvPackage loads dot env package
*/
func LoadEnvPackage() {
	err := godotenv.Load()
	utils.HandleError(err, "Error loading .env file")
}
