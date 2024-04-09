package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *Server) setupStudentRoutes() {
	// Define routes for student functionality
	// For example:
	s.app.Get("/student/profile/:regNo", func(c *fiber.Ctx) error {
		// Get the regNo parameter from the request
		regNo := c.Params("regNo")

		// Call GetStudentProfile method with the regNo parameter
		profile, err := s.studentController.GetStudentProfile(regNo)
		if err != nil {
			return err
		}

		// Return the profile as JSON response
		return c.JSON(profile)
	})

	// Add more routes as needed
}
