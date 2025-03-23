package server

import (
	"fmt"
	"giva-url-shortner/database"
	"giva-url-shortner/utils"
	"net/url"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func createGoShort(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	var goshort database.GoShort
	err := ctx.BodyParser(&goshort)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing JSON " + err.Error(),
		})
	}

	// Validate URL
	_, err = url.ParseRequestURI(goshort.Redirect)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error validating url " + err.Error(),
		})
	}
	dl, err := database.FindByRedirectUrl(goshort.Redirect)
	if err == nil {
		return ctx.Status(fiber.StatusOK).
			JSON(fiber.Map{"alreadyExist": true, "short-url": dl.Goshort})
	}
	custom := ctx.Query("alias")
	if len(custom) != 0 {
		goshort.Goshort = custom
	} else {
		lenght, err := strconv.Atoi(os.Getenv("LENGTH"))
		if err != nil {
			lenght = 8
		}
		goshort.Goshort = utils.RandomURL(lenght)
	}

	goshort, err = database.CreateGoShort(goshort)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Try Again. This short url already exist.",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(goshort.Goshort)
}

func redirect(ctx *fiber.Ctx) error {
	goshortUrl := ctx.Params("redirect")

	goshort, err := database.FindByGoShortUrl(goshortUrl)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "could not find goshort in DB " + err.Error(),
		})
	}

	// Stats: update times url clicked
	goshort.Clicked += 1
	_, err = database.UpdateGoShort(goshort)
	if err != nil {
		fmt.Printf("error updating stats: %v\n", err)
	}

	return ctx.Redirect(goshort.Redirect, fiber.StatusTemporaryRedirect)
}

func getStats(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")
	goshortUrl := ctx.Params("redirect")
	goshort, err := database.FindByGoShortUrl(goshortUrl)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "could not find goshort in DB " + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"no_of_clicks": goshort.Clicked})
}

func RunServer() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Post("/shorten", createGoShort)

	app.Get("/:redirect", redirect)
	app.Get("/stats/:redirect", getStats)
	app.Listen(":3000")
}
