package server

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) setupMarkRoutes() {
	
	s.app.Post("/mark/add", func(c *fiber.Ctx) error {
		// Parse request body to get parameters
		var body struct {
			RegNo       string `json:"regNo"`
			ExamNo      int    `json:"examNo"`
			SubjectCode string `json:"subjectCode"`
			Mark        int    `json:"mark"`
		}
		if err := c.BodyParser(&body); err != nil {
			return err
		}

		// Call AddMark method with parsed parameters
		if err := s.markController.AddMark(body.RegNo, body.ExamNo, body.SubjectCode, body.Mark); err != nil {
			return err
		}

		// Return success response
		return c.SendStatus(fiber.StatusCreated)
	})

	s.app.Get("/mark/get/:regNo/:examNo", func(c *fiber.Ctx) error {
		// Get the regNo and examNo parameters from the request
		regNo := c.Params("regNo")
		examNoStr := c.Params("examNo")

		// Convert examNo to int
		examNo, err := strconv.Atoi(examNoStr)
		if err != nil {
			return err
		}

		// Call GetMarks method with the regNo and examNo parameters
		marks, err := s.markController.GetMarks(regNo, examNo)
		if err != nil {
			return err
		}

		// Return the marks as JSON response
		return c.JSON(marks)
	})
	// Add more routes as needed
}
