package cmd

import(
	"log"

	"github.com/jeanebaebae/suplaihub-app/backend/cmd/api"
	"github.com/jeanebaebae/suplaihub-app/backend/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found")
	}

	cfg := config.LoadConfig()

	server := api.NewServer(cfg.Port)
	if err := server.Run(); err != nil {
		log.Fatalf("Server not found: %v", err)
	}
}