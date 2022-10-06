package main
import "github.com/rgattas-meli/backpack-bcgow6-roberto-gattas/go_bases/hackaton-go-bases-main/internal/service"

func main() {
	var tickets []service.Ticket
	// Funcion para obtener tickets del archivo csv
	service.NewBookings(tickets)
}
