package reservationmastersdto

import (
	"time"
)

type ReservationMasterReq struct {
	KodeOrder              string    `json:"kode_order" form:"kode_order" validate:"required"`
	Status                 string    `json:"status" form:"status" validate:"required"`
	OrderMasuk             time.Time `json:"order_masuk" form:"order_masuk" validate:"required"`
	UserID                 uint32    `json:"user_id" form:"user_id" validate:"required"`
	ReservationVehicleID   uint32    `json:"reservation_vehicle_id" form:"reservation_vehicle_id" validate:"required"`
	ReservationInsuranceID uint32    `json:"reservation_insurance_id" form:"reservation_insurance_id" validate:"required"`
	ReservationItemID      uint32    `json:"reservation_item_id" form:"reservation_item_id" validate:"required"`
}

type ReservationMasterReqUpdate struct {
	Status                 string    `json:"status" form:"status"`
	OrderProses             time.Time `json:"order_masuk" form:"order_masuk"`
	UserID                 uint32    `json:"user_id" form:"user_id"`
	ReservationVehicleID   uint32    `json:"reservation_vehicle_id" form:"reservation_vehicle_id"`
	ReservationInsuranceID uint32    `json:"reservation_insurance_id" form:"reservation_insurance_id"`
	ReservationItemID      uint32    `json:"reservation_item_id" form:"reservation_item_id"`
}
