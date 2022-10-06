package file

import (
	"errors"
	
	"os"
	"encoding/csv"
	"strconv"
	"github.com/rgattas-meli/backpack-bcgow6-roberto-gattas/go_bases/hackaton-go-bases-main/internal/service"
	"github.com/gocarina/gocsv"
)

type File struct {
	path string
}

func (f *File) Read() ([]service.Ticket, error) {
	file, err := os.Open(f.path)
	defer file.Close()
	if err!= nil {
        return nil, err
    }
	if _, err := file.Seek(0, 0); err != nil { // Go to the start of the file
		return nil, errors.New("error en cargar 0,0")
	}
	var tickets []service.Ticket
	records, err := csv.NewReader(file).ReadAll()
	for _, record := range records {
		price1,_ := strconv.Atoi(record[5])
		id1,_ := strconv.Atoi(record[0])
		data := service.Ticket{
			Id: id1,
			Names: record[1],
			Email: record[2],
			Destination: record[3],
			Date : record[4],
			Price : price1,
		}
		tickets = append(tickets, data)

	}
	if err != nil {
		return nil, errors.New("error en cargar personas")
	}

	return tickets, err
}

func (f *File) Write(service.Ticket) error {
	file, err := os.Open(f.path)
    defer file.Close()
    if err!= nil {
        return err
    }
	csvwriter := csv.NewWriter(file)
	row := []string{strconv.Itoa(service.Ticket.Id), service.Ticket.Names, service.Ticket.Email, service.Ticket.Destination, service.Ticket.Date,strconv.Itoa(service.Ticket.Price)}
	if err := csvwriter.Write(row); err != nil {
		return errors.New("error en cargar 0,0")	}
	csvwriter.Flush()
	return err
}
