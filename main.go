package main

/*import (
	"context"
	"controllers/backend/controllers"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Select database
	db := client.Database("STU")

	// Initialize AttendanceController
	attendanceController := controllers.NewAttendanceController(db)

	// Call AddAttendance

	regNo := "2023001"
	attendanceRecords, err := attendanceController.GetAttendance(regNo)
	if err != nil {
		log.Fatal(err)
	}

	// Print the attendance records
	fmt.Println("Attendance records for", regNo)
	for _, record := range attendanceRecords {
		fmt.Println("Date:", record.Date.Format("2006-01-02"), "Status:", record.Status)
	}

}
*/
