package utils

import (
	"os"
)

func GetEnvVariable(key string) string {
	// if err := godotenv.Load(); err != nil {
	// 	fmt.Println(err.Error())
	// }

	return os.Getenv(key)
}
