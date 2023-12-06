package main

import (
    "fmt"
    "math"
    "os"
    "reflect"
)

type Auditorium struct {
    rows, seats int
}

func NewAuditorium(row, seat int) *Auditorium {
    return &Auditorium{rows: row, seats: seat}
}

func (a *Auditorium) isSeatInputValid(row, seat int) bool {
    return seat < 0 || seat > a.seats && row < 0 || row > a.rows
}

func (a *Auditorium) totalTickets() int {
    return a.rows * a.seats
}

type SeatBooking struct {
    row, seat int
    price     float64
}

func NewSeatBooking(row, seat int, price float64) *SeatBooking {
    return &SeatBooking{row: row, seat: seat, price: price}
}

type Bookings []SeatBooking

func (b *Bookings) addBooking(seatBooking *SeatBooking) {
    *b = append(*b, *seatBooking)
}

func (b *Bookings) findBooking(row, seat int) SeatBooking {
    if b.getBooking() > 0 {
        for _, booking := range *b {
            if booking.row == row && booking.seat == seat {
                return booking
            }
        }
    }

    return SeatBooking{}
}

func (b *Bookings) getBooking() int {
    return len(*b)
}

func cinemaSeats(auditorium Auditorium, bookings Bookings) {
    fmt.Println("Cinema:")

    for row := 0; row <= auditorium.rows; row++ {
        for seat := 0; seat <= auditorium.seats; seat++ {
            if row == 0 && seat == 0 {
                fmt.Print("  ")
            }

            if seat == 0 && row > 0 {
                fmt.Printf("%d ", row)
            }

            if row == 0 && seat > 0 {
                fmt.Printf("%d ", seat)
            }

            if row > 0 && seat > 0 {
                if booking := bookings.findBooking(row, seat); &booking != nil {
                    switch {
                    case row == booking.row && seat == booking.seat:
                        fmt.Print("B ")
                    default:
                        fmt.Print("S ")
                    }
                }
            }
        }
        fmt.Println()
    }
}

func calculateTotalIncome(auditorium Auditorium) float64 {
    var price float64

    totalSeats := auditorium.rows * auditorium.seats

    if totalSeats <= 60 {
        price = float64(totalSeats * 10)
    } else {
        price = math.Floor(float64(auditorium.rows)/2) * float64(auditorium.seats) * 10
        price += math.Ceil(float64(auditorium.rows)/2) * float64(auditorium.seats) * 8
    }

    return price
}

func currentIncome(bookings Bookings) float64 {
    var total float64
    for _, booking := range bookings {
        total += booking.price
    }

    return total
}

func getSeatPrice(auditorium Auditorium, row int) float64 {
    totalSeats := auditorium.rows * auditorium.seats

    if totalSeats <= 60 {
        return 10
    } else {
        switch {
        case row <= auditorium.rows/2:
            return 10
        default:
            return 8
        }
    }
}

func menu() int {
    var option int

    fmt.Println("\n\n1. Show the seats")
    fmt.Println("2. Buy a ticket")
    fmt.Println("3. Statistics")
    fmt.Println("0. Exit")
    fmt.Println()
    fmt.Scan(&option)

    return option
}

func statistics(auditorium Auditorium, bookings Bookings) {
    totalBookings := bookings.getBooking()
    purchasedPercentage := (float64(totalBookings) / float64(auditorium.totalTickets())) * 100

    fmt.Printf("Number of purchased tickets: %d", totalBookings)
    fmt.Printf("\nPercentage: %.2f%%", purchasedPercentage)
    fmt.Printf("\nCurrent income: $%.f", currentIncome(bookings))
    fmt.Printf("\nTotal income: $%.f\n", calculateTotalIncome(auditorium))
}

func main() {
    var rows, seats, row, seat int
    bookings := Bookings{}

    fmt.Println("Enter the number of rows:")
    fmt.Scan(&rows)

    fmt.Println("Enter the number of seats in each row:")
    fmt.Scan(&seats)

    auditorium := *NewAuditorium(rows, seats)

    for {
        option := menu()

        switch option {
        case 0:
            os.Exit(0)
        case 1:
            cinemaSeats(auditorium, bookings)
        case 2:
            for {
                var seatPrice float64

                fmt.Println("\nEnter a row number:")
                fmt.Scan(&row)

                fmt.Println("Enter a seat number in that row:")
                fmt.Scan(&seat)

                if auditorium.isSeatInputValid(row, seat) {
                    fmt.Printf("\nWrong input!")
                } else {
                    if booking := bookings.findBooking(row, seat); reflect.DeepEqual(&booking, NewSeatBooking(0, 0, 0)) {
                        seatPrice = getSeatPrice(auditorium, row)

                        bookings.addBooking(NewSeatBooking(row, seat, seatPrice))

                        fmt.Printf("\nTicket price: $%.f", seatPrice)

                        break
                    } else {
                        fmt.Println("That ticket has already been purchased!")
                    }
                }
            }
        case 3:
            statistics(auditorium, bookings)
        }
    }
}
