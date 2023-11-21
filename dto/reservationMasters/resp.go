package reservationmastersdto

import (
	"time"
)

type ReservationMasterResp struct {
	KodeOrder            string    `json:"kode_order" form:"kode_order"`
	Status               string    `json:"status" form:"status"`
	OrderMasuk           time.Time `json:"order_masuk" form:"order_masuk"`
	OrderProses          time.Time `json:"order_proses" form:"order_proses"`
	OrderSelesai         time.Time `json:"order_selesai" form:"order_selesai"`
	UserID               uint32    `json:"user_id" form:"user_id"`
	ReservationVehicle   uint32    `json:"reservation_vehicle_id" form:"reservation_vehicle_id"`
	ReservationInsurance uint32    `json:"reservation_insurance_id" form:"reservation_insurance_id"`
	ReservationItem      uint32    `json:"reservation_item_id" form:"reservation_item_id"`
}
