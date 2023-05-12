package controller

import (
	"errors"
	"fmt"
	"net/http"

	inimodel "github.com/ArykaAnisaP/Penggajian/model"
	inimodule "github.com/ArykaAnisaP/Penggajian/module"
	"github.com/ArykaAnisaP/arykaanisap/config"
	"github.com/aiteung/musik"
	cek "github.com/aiteung/presensi"
	"github.com/gofiber/fiber/v2"
	inimodell "github.com/indrariksa/be_presensi/model"

	// inimodullatihan "github.com/indrariksa/be_presensi/model"
	inimodulee "github.com/indrariksa/be_presensi/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// func Home(c *fiber.Ctx) error {
// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"github_repo": "https://github.com/ArykaAnisaP/arykaanisap",
// 		"message":     "You are at the root endpoint 😉",
// 		"success":     true,
// 	})
// }

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetPresensi(c *fiber.Ctx) error {
	ps := cek.GetPresensiCurrentMonth(config.Ulbimongoconn)
	return c.JSON(ps)
}
func GetAllUang(c *fiber.Ctx) error {
	ps := inimodule.GetAllUang(config.Ulbimongoconn, "uang")
	return c.JSON(ps)
}

func GetAllGajiFromNamaKaryawan(c *fiber.Ctx) error {
	ps := inimodule.GetGajiFromNamaKaryawan("aryka", config.Ulbimongoconn, "uang")
	return c.JSON(ps)
}

func GetAllPresensiFromWaktu(c *fiber.Ctx) error {
	ps := inimodule.GetPresensiFromWaktu("12-03-2023", config.Ulbimongoconn, "presensi")
	return c.JSON(ps)
}

func InsertJamker(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var jamker inimodel.Jamker
	if err := c.BodyParser(&jamker); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID := inimodule.InsertJamker(db, "jamker",
		jamker.Jam_masuk,
		jamker.Jam_keluar,
		jamker.Hari,
		jamker.Shift,
		jamker.Piket_tim)
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func GetAllGajiFromPresensi(c *fiber.Ctx) error {
	ps := inimodule.GetGajiFromPresensi("0854632178", config.Ulbimongoconn, "uang")
	return c.JSON(ps)
}

func GetAllPresensi(c *fiber.Ctx) error {
	ps := inimodulee.GetAllPresensi(config.Ulbimongoconn, "presensi")
	return c.JSON(ps)
}

func GetPresensiID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ps, err := inimodulee.GetPresensiFromID(objID, config.Ulbimongoconn, "presensi")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(ps)
}

func InsertData(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var presensi inimodell.Presensi
	if err := c.BodyParser(&presensi); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := inimodulee.InsertPresensi(db, "presensi",
		presensi.Longitude,
		presensi.Latitude,
		presensi.Location,
		presensi.Phone_number,
		presensi.Checkin,
		presensi.Biodata)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}
