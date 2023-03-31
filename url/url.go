package url

import (
	"github.com/ArykaAnisaP/arykaanisap/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/", controller.Homepage)                                    //ujicoba panggil package musik
	page.Get("/presensi", controller.GetAllPresensi)
	page.Get("/gaji", controller.GetAllGajiFromNamaKaryawan)
	page.Get("/uang", controller.GetAllGajiFromPresensi)
	page.Get("/presensi", controller.GetAllPresensi)    //menampilkan seluruh data presensi
	page.Get("/presensi/:id", controller.GetPresensiID) //menampilkan data presensi berdasarkan id
}
