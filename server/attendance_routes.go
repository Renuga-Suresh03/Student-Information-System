package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *Server) setupAttendanceRoutes() {

	s.app.Post("/attendance/add", func(c *fiber.Ctx) error {
		var body struct {
			RegNo  string `json:"regNo"`
			Status string `json:"status"`
		}
		if err := c.BodyParser(&body); err != nil {
			return err
		}

		if err := s.attendanceController.AddAttendance(body.RegNo, body.Status); err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusCreated)
	})

	s.app.Get("/attendance/get/:regNo", func(c *fiber.Ctx) error {
		// Get the regNo parameter from the request
		regNo := c.Params("regNo")

		// Call GetAttendance method with the regNo parameter
		attendance, err := s.attendanceController.GetAttendance(regNo)
		if err != nil {
			return err
		}

		// Return the attendance details as JSON response
		return c.JSON(attendance)
	})
	// Add more routes as needed
}
