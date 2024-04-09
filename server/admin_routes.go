package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *Server) setupAdminRoutes() {
	// Define routes for admin functionality
	// For example:
	s.app.Post("/admin/login", func(c *fiber.Ctx) error {
		// Implement logic for admin login
		// Return appropriate response based on the login result
		return c.SendString("Admin login handler")
	})

	s.app.Get("/admin/profile/:adminID", func(c *fiber.Ctx) error {
		// Get the adminID parameter from the request
		adminID := c.Params("adminID")

		// Call GetProfile method with the adminID parameter
		profile, err := s.adminController.GetAdminProfile(adminID)
		if err != nil {
			return err
		}

		// Return the profile as JSON response
		return c.JSON(profile)
	})
	// Add more routes as needed
}
