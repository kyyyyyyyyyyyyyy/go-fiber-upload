package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	enggine := html.New("./", ".html")

	app := fiber.New(fiber.Config{Views: enggine})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	app.Post("/upload", func(c *fiber.Ctx) error {
		var Input struct {
			Name string
		}

		if err := c.BodyParser(&Input); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		gambar, err := c.FormFile("image")
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		//mengambil nama file
		fmt.Println("nama file: \n", gambar.Filename)
		//mengambil ukuran file
		fmt.Println("ukuran file (bytes): \n", gambar.Size)
		fmt.Println("ukuran file (Kilobytes): \n", gambar.Size/1024)
		fmt.Println("ukuran file (Kilobytes): \n", float64((gambar.Size)/1024)/1024)
		//mengambil nime type
		fmt.Println("nime type: \n", gambar.Header.Get("Content-Type"))
		//mengambil ekstensi file
		splitDot := strings.Split(gambar.Filename, ".")
		ext := splitDot[len(splitDot)-1]
		fmt.Println(ext)

		//mengubah nama file
		chanefile := fmt.Sprintf("%s.%s", time.Now().Format("2006-01-02-15-04-05"), ext)
		fmt.Println(chanefile)

		//mengambil ukuran gambar
		fileHeader, _ := gambar.Open()
		defer fileHeader.Close()

		imageConfig, _, err := image.DecodeConfig(fileHeader)
		if err != nil {
			log.Print(err)
		}

		width := imageConfig.Width
		height := imageConfig.Height
		fmt.Println("leabar: \n", width)
		fmt.Println("leabar: \n", height)

		//menyimpan file ke folder
		if err := c.SaveFile(gambar, "./uploads/"+chanefile); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"title":   Input.Name,
			"image":   chanefile,
			"message": "upload successfully",
		})
	})

	app.Listen(":8080")

}
