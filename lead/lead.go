package lead

import (
	"github.com/achintya-7/crm-go-basic/database"
	"github.com/gofiber/fiber"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

/* 
	* Demo Post Body
	{
  		"id": 1,
  		"name": "Achintyta",
  		"company": "Self",
  		"email": "achintya22052000@gmail.com",
  		"phone": 9810974477
	}
*/

type Lead struct {
	ID 			 int 		`json:"id"`
	Name         string     `json:"name"`  // * it says for json, Name will be name
	Company      string     `json:"company"`
	Email        string     `json:"email"`
	Phone        int        `json:"phone"`
}


func GetLeads(c *fiber.Ctx) {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}

func GetLead(c *fiber.Ctx)  {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
}

func NewLead(c *fiber.Ctx) {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var lead Lead
	db.First(&lead, id)
	if lead.Name == ""{
		c.Status(503).Send("No lead found with ID")
		return
	}
	db.Delete(&lead)
	c.Send("Lead Deleted")
}