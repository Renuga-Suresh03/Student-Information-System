//admin_test
/*

package main


import (
    "context"
    "controllers/backend/controllers"
    "fmt"
    "log"


    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)


func main() {
    mongoURI := "mongodb://localhost:27017"
    client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
    if err != nil {
        log.Fatal("Error creating MongoDB client:", err)
    }
    defer client.Disconnect(context.Background())
    err = client.Connect(context.Background())
    if err != nil {
        log.Fatal("Error connecting to MongoDB:", err)
    }
    adminController := controllers.NewAdminController(client.Database("STU"))
    adminID := "admin123"
    password := "secretpassword"
    admin, err := adminController.AuthenticateAdmin(adminID, password)
    if err != nil {
        fmt.Printf("Authentication failed: %v\n", err)
        return
    }
    fmt.Printf("Authentication successful! Admin Details: %+v\n", admin)
    adminProfile, err := adminController.GetAdminProfile(adminID)
    if err != nil {
        fmt.Printf("Error retrieving admin profile: %v\n", err)
        return
    }
    fmt.Printf("Admin Profile: %+v\n", adminProfile)
}
*/

//----------------------------------------------------------------------------------------------------------------------

//student test

/*package main

import (
	"context"
	"controllers/backend/controllers"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongoURI := "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Error creating MongoDB client:", err)
	}
	defer client.Disconnect(context.Background())
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	studentController := controllers.NewStudentController(client.Database("STU"))

	// Test AuthenticateStudent
	regNo := "2023001"
	dob := "1998-05-15"
	authenticatedStudent, err := studentController.AuthenticateStudent(regNo, dob)
	if err != nil {
		fmt.Printf("Authentication failed: %v\n", err)
		return
	}
	fmt.Printf("Authentication successful! Student Details: %+v\n", authenticatedStudent)

	// Test GetStudentProfile
	retrievedStudent, err := studentController.GetStudentProfile(regNo)
	if err != nil {
		fmt.Printf("Error retrieving student profile: %v\n", err)
		return
	}
	fmt.Printf("Student Profile: %+v\n", retrievedStudent)
}*/

//----------------------------------------------------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------------------------------------------------------

//mark test

//successful

/*package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"controllers/backend/controllers"
)

func main() {
	// Connect to MongoDB
	mongoURI := "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Error creating MongoDB client:", err)
	}
	defer client.Disconnect(context.Background())
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	// Initialize MarkController
	markController := controllers.NewMarkController(client.Database("STU"))

	// Test AddMark
	regNo := "2023001"
	examNo := 1
	subjectCode := "SOC103"
	subject := "Social"
	mark := 90
	err = markController.AddMark(regNo, examNo, subjectCode, subject, mark)
	if err != nil {
		fmt.Printf("Failed to add mark: %v\n", err)
		return
	}
	fmt.Println("Mark added successfully!")

	// Test GetMarks
	examNoToFetch := 1
	marks, err := markController.GetMarks(regNo, examNoToFetch)
	if err != nil {
		fmt.Printf("Failed to get marks: %v\n", err)
		return
	}

	// Print the marks
	fmt.Println("Marks for exam", examNoToFetch, ":")
	for _, m := range marks {
		for _, subjectMark := range m.Subjects {
			fmt.Printf("Subject Code: %s, Subject: %s, Mark: %d\n", subjectMark.SubjectCode, subjectMark.Subject, subjectMark.Mark)
		}
	}
}
*/

//----------------------------------------------------------------------------------------------------------------

// testing attendance pending

// main.go
/*package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"controllers/backend/controllers"
)

func main() {
	// Initialize MongoDB client
	mongoURI := "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Error creating MongoDB client:", err)
	}
	defer client.Disconnect(context.Background())
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	// Initialize AttendanceController
	attendanceController := controllers.NewAttendanceController(client.Database("STU"))

	// Test GetAttendance before adding new attendance
	regNo := "2023001"
	existingAttendance, err := attendanceController.GetAttendance(regNo)
	if err != nil {
		fmt.Printf("Failed to get existing attendance: %v\n", err)
		return
	}

	fmt.Println("Existing Attendance Record:")
	if existingAttendance != nil {
		fmt.Printf("Attendance Percentage: %d%%\n", existingAttendance.AttendancePercentage)
		fmt.Println("Attendance Records for the Last 7 Days:")
		for _, record := range existingAttendance.AttendanceRecords {
			fmt.Printf("Date: %s, Status: %s\n", record.Date.Format("2006-01-02"), record.Status)
		}
	} else {
		fmt.Println("No existing attendance records found.")
	}

	// Test AddAttendance
	status := "present"
	err = attendanceController.AddAttendance(regNo, status)
	if err != nil {
		fmt.Printf("Failed to add attendance: %v\n", err)
		return
	}
	fmt.Println("Attendance added successfully!")

	// Test GetAttendance after adding new attendance
	newAttendance, err := attendanceController.GetAttendance(regNo)
	if err != nil {
		fmt.Printf("Failed to get new attendance: %v\n", err)
		return
	}

	fmt.Println("New Attendance Record:")
	fmt.Printf("Attendance Percentage: %d%%\n", newAttendance.AttendancePercentage)
	fmt.Println("Attendance Records for the Last 7 Days:")
	for _, record := range newAttendance.AttendanceRecords {
		fmt.Printf("Date: %s, Status: %s\n", record.Date.Format("2006-01-02"), record.Status)
	}
}
*/