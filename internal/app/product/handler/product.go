package handler

import "github.com/gofiber/fiber/v2"

func (h *HProduct) All(c *fiber.Ctx) error {
	ctx := c.UserContext()
	keyword := c.Query("keyword")

	p, err := h.up.All(ctx, keyword)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"payload": p,
	})
}

func (h *HProduct) Get(c *fiber.Ctx) error {
	ctx := c.UserContext()
	id := c.Params("productID")

	p, err := h.up.Get(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"payload": p,
	})
}
