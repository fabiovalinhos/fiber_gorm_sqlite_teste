package routes

import (
	"time"

	"github.com/fabiovalinhos/fiber-api/database"
	"github.com/fabiovalinhos/fiber-api/models"
	"github.com/gofiber/fiber/v2"
)

// {
// 	id: 1,
// 	user: {
// 		id: 23,
// 		first_name: "Bob",
// 		last_name: "Jonatas"
// 	},
// 	product: {
// 		id: 24,
// 		name: "Macbook",
// 		serial_number: "323232323"
// 	}
// }

type Order struct {
	ID       uint      `json:"id"`
	User     User      `json:"user"`
	Product  Product   `json:"product"`
	CreateAt time.Time `json:"order_date"`
}

func CreateResponseOrder(order models.Order, user User, product Product) Order {
	return Order{
		ID:       order.ID,
		User:     user,
		Product:  product,
		CreateAt: order.CreateAt,
	}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	if err := findUser(order.UserRefer, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var product models.Product
	if err := findProduct(order.ProductRefer, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&order)

	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

	return c.Status(200).JSON(responseOrder)
}
