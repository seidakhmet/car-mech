package person

import (
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

// Person structure
type Person struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

// GetPerson function
func GetPerson(c *fiber.Ctx) {
	c.Send("Test handler")
}
