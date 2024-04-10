// adminroutes.go

package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *Server) setupAdminRoutes() {
	admin := s.app.Group("/admin")

	admin.Post("/login", s.loginAdmin)
	admin.Get("/profile", s.getAdminProfile)
}

func (s *Server) loginAdmin(c *fiber.Ctx) error {
	adminID := c.FormValue("admin_id")
	password := c.FormValue("password")

	admin, err := s.adminController.AuthenticateAdmin(adminID, password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(admin)
}

func (s *Server) getAdminProfile(c *fiber.Ctx) error {
	adminID := c.Query("admin_id")

	admin, err := s.adminController.GetAdminProfile(adminID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(admin)
}
