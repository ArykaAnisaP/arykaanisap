package controller

import (
	"net/http"

	"github.com/ArykaAnisaP/arykaanisap/config"
	inimodel "github.com/ArykaAnisaP/Penggajian/model"
	inimodule "github.com/ArykaAnisaP/Penggajian/module"
	"github.com/aiteung/musik"
	cek "github.com/aiteung/presensi"
	"github.com/gofiber/fiber/v2"
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
