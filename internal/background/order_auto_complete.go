package background

import (
	"github.com/Anwarjondev/fast-food/internal/db"
	"log"
	"time"
)

func AutoCompleteOrders() {
	ticker := time.NewTicker(1 * time.Minute)

	go func() {
		for range ticker.C {
			_, err := db.DB.Exec(`
				update orders
				set status = 'completed', delivered_at = now()
				where status = 'active' and now() -  created_at >interval '10 minutes'
			`)
			if err != nil {
				log.Println("Error auto completing orders:", err)
			}
			
		}
	}()
}
