package main
import "github.com/bootcamp-go/hackaton-go-bases/internal/service"
import "os"
import ""

func main() {
	var tickets []service.Ticket
	file := *File{
		path : "tickets.csv",
	}
	tickets, err := file.Read()
	if err!= nil {
        panic(err)
    }


	


	service.NewBookings(tickets)
}
