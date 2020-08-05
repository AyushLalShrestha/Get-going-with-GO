package gorooutines

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Response is a type of HTTP handler response
type Response struct {
	Status  string
	Message string
}

func handleCustomerTransaction(customerID int, isDone chan bool) {
	log.Println("ID:", customerID, " Recieved booking request...")
	status := make(chan string)
	transactionSuccess := make(chan bool)
	// It returns a channel whose value is going to be written after 15 seconds for timeout
	timeout := time.After(15 * time.Second)

	// This function if not a go routine blocks the below code
	go func() {
		status <- " Seat selection going on..."
		// Use customer details to make DB queries, third party API/Service call
		time.Sleep(5 * time.Second)
		status <- " Making payments from bank..."
		time.Sleep(5 * time.Second)
		// Everything looks good. Notify customer
		transactionSuccess <- true
		defer close(transactionSuccess)
	}()

	for {
		select {
		case update := <-status:
			log.Println("ID:", customerID, update)
		case <-timeout:
			close(status)
			log.Println("Operation timed out!")
			isDone <- false
			return
		case <-transactionSuccess:
			log.Println("ID:", customerID, "Successfully booked ticket!")
			isDone <- true
			return
		}
	}
}

func bookingHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/v1/flight-booking/")
	isDone := make(chan bool)
	customerID, _ := strconv.Atoi(id)

	// Call our magic concurrent function
	go handleCustomerTransaction(customerID, isDone)
	w.Header().Set("Content-Type", "application/json")
	response := Response{}
	// Check whether operation is success of a failure
	if <-isDone {
		response.Status = "success"
		response.Message = "Dear passenger, your ticket is booked sucessfully!"
		jsonResponse, _ := json.Marshal(response)
		w.Write(jsonResponse)
	} else {
		response.Status = "failure"
		response.Message = "Operation timed out! Please try again"
		jsonResponse, _ := json.Marshal(response)
		w.Write(jsonResponse)
	}
	defer close(isDone)
}

func main() {
	// Handle the API path
	http.HandleFunc("/v1/flight-booking/", bookingHandler)
	http.ListenAndServe(":8000", nil)
}
