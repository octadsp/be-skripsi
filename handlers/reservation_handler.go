package handlers

import (
	reservationsdto "be-skripsi/dto/reservations"
	dto "be-skripsi/dto/results"
	"be-skripsi/models"
	"be-skripsi/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerReservation struct {
	ReservationRepository repositories.ReservationRepository
}

func HandlerReservation(ReservationRepository repositories.ReservationRepository) *handlerReservation {
	return &handlerReservation{ReservationRepository}
}

func (h *handlerReservation) FindReservations(c echo.Context) error {
	s, err := h.ReservationRepository.FindReservations()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: s})
}

func (h *handlerReservation) FindReservationsDone(c echo.Context) error {
	s, err := h.ReservationRepository.FindReservationsDone()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: s})
}

// func (h *handlerReservation) FindReservationsStatus(c echo.Context) error {
// 	status := c.QueryParam("status") // Ambil nilai dari query parameter "status"
// 	if status == "" {                // Periksa jika status kosong
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Missing status parameter"})
// 	}

// 	reservations, err := h.ReservationRepository.FindReservationsStatus(status)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
// 	}

//		return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: reservations})
//	}
func (h *handlerReservation) FindReservationsStatus(c echo.Context) error {
	// Ambil nilai dari query parameter "status"
	status := c.QueryParam("status")
	if status == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Missing status parameter"})
	}

	// Ambil nilai dari query parameter "date"
	dateStr := c.QueryParam("date")
	var date time.Time
	var err error
	// Periksa jika query parameter "date" tidak kosong
	if dateStr != "" {
		// Konversi string tanggal menjadi tipe time.Time
		date, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid date format"})
		}
	}

	// Panggil metode repository untuk mencari reservasi berdasarkan status dan tanggal
	var reservations []models.Reservation
	if date.IsZero() {
		reservations, err = h.ReservationRepository.FindReservationsStatus(status, date)
	} else {
		reservations, err = h.ReservationRepository.FindReservationsStatus(status, date)
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: reservations})
}

func (h *handlerReservation) FindReservationsStatusFromAndUntil(c echo.Context) error {
	// Ambil nilai status dari query parameter "status"
	status := c.QueryParam("status")

	// Ambil nilai from dan until dari query parameter "from" dan "until"
	fromStr := c.QueryParam("from")
	untilStr := c.QueryParam("until")

	// Parsing nilai "from" dan "until" ke dalam tipe data time.Time
	from, err := time.Parse("2006-01-02", fromStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid 'from' parameter format"})
	}

	until, err := time.Parse("2006-01-02", untilStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid 'until' parameter format"})
	}

	// Panggil fungsi repository untuk mencari reservasi berdasarkan status, dari, dan sampai
	reservations, err := h.ReservationRepository.FindReservationsStatusFromAndUntil(status, from, until)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	// Kirim hasil pencarian sebagai JSON response
	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: reservations})
}

func (h *handlerReservation) FindReservationsStatusFromAndUntilChart(c echo.Context) error {
	// Ambil nilai status dari query parameter "status"
	status := c.QueryParam("status")

	// Ambil nilai from dan until dari query parameter "from" dan "until"
	fromStr := c.QueryParam("from")
	untilStr := c.QueryParam("until")

	// Parsing nilai "from" dan "until" ke dalam tipe data time.Time
	from, err := time.Parse("2006-01-02", fromStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid 'from' parameter format"})
	}

	until, err := time.Parse("2006-01-02", untilStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: "Invalid 'until' parameter format"})
	}

	// Panggil fungsi repository untuk mencari reservasi berdasarkan status, dari, dan sampai
	reservations, err := h.ReservationRepository.FindReservationsStatusFromAndUntilChart(status, from, until)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}
	// Kirim hasil pencarian sebagai JSON response
	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: respReservationChart(reservations)})
}

func (h *handlerReservation) GetReservSubStatus(c echo.Context) error {
	status := c.Param("substatus")

	reserv, err := h.ReservationRepository.GetReservSubStatus(status)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: reserv})
}

func (h *handlerReservation) GetReservation(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// userLogin := c.Get("userLogin")
	// userId := userLogin.(jwt.MapClaims)["id"].(float64)

	reserv, err := h.ReservationRepository.GetReservation(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: reserv})
}

func (h *handlerReservation) GetReservationCountByDate(c echo.Context) error {
	dateParam := c.QueryParam("date")
	date, err := time.Parse("2006-01-02", dateParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	count, err := h.ReservationRepository.GetReservationCountByDate(date)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]int64{"count": count})
}

func (h *handlerReservation) AddReservation(c echo.Context) error {
	request := new(reservationsdto.ReservationReq)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	reserv := models.Reservation{
		KodeOrder:  request.KodeOrder,
		Status:     request.Status,
		OrderMasuk: request.OrderMasuk,
		UserID:     request.UserID,

		CarBrand: request.CarBrand,
		CarType:  request.CarType,
		CarYear:  request.CarYear,
		CarColor: request.CarColor,

		IsInsurance: request.IsInsurance,

		InsuranceName:  request.InsuranceName,
		EventDate:      request.EventDate,
		Place:          request.Place,
		Time:           request.Time,
		DrivingSpeed:   request.DrivingSpeed,
		DriverName:     request.DriverName,
		DriverRelation: request.DriverRelation,
		DriverJob:      request.DriverJob,
		DriverAge:      request.DriverAge,
		DriverLicense:  request.DriverLicense,

		CreatedAt: time.Now(),
	}

	data, err := h.ReservationRepository.AddReservation(reserv)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerReservation) UpdateReservation(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	request := new(reservationsdto.ReservationReqUpdate)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	reserv, err := h.ReservationRepository.GetReservation(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	// Gunakan time.Time.Nil untuk menandakan bahwa field tersebut tidak diubah
	var orderProses, orderSelesai time.Time

	// Mengecek apakah request.OrderProses dan request.OrderSelesai berisi nilai yang tidak nol
	if !request.OrderProses.IsZero() {
		orderProses = request.OrderProses
	}

	if !request.OrderSelesai.IsZero() {
		orderSelesai = request.OrderSelesai
	}

	// Update nilai-nilai yang tidak nol
	reserv.OrderProses = orderProses
	reserv.OrderSelesai = orderSelesai

	if request.CarBrand != "" {
		reserv.CarBrand = request.CarBrand
	}

	if request.CarType != "" {
		reserv.CarType = request.CarType
	}

	if request.CarYear != "" {
		reserv.CarYear = request.CarYear
	}

	if request.CarColor != "" {
		reserv.CarColor = request.CarColor
	}

	if request.IsInsurance != 0 {
		reserv.IsInsurance = request.IsInsurance
	}

	if request.InsuranceName != "" {
		reserv.InsuranceName = request.InsuranceName
	}

	if request.EventDate != "" {
		reserv.EventDate = request.EventDate
	}
	if request.Place != "" {
		reserv.Place = request.Place
	}
	if request.Time != "" {
		reserv.Time = request.Time
	}
	if request.DrivingSpeed != "" {
		reserv.DrivingSpeed = request.DrivingSpeed
	}
	if request.DriverName != "" {
		reserv.DriverName = request.DriverName
	}
	if request.DriverRelation != "" {
		reserv.DriverRelation = request.DriverRelation
	}
	if request.DriverJob != "" {
		reserv.DriverJob = request.DriverJob
	}
	if request.DriverAge != "" {
		reserv.DriverAge = request.DriverAge
	}
	if request.DriverLicense != "" {
		reserv.DriverLicense = request.DriverLicense
	}

	reserv.UpdatedAt = time.Now()

	data, err := h.ReservationRepository.UpdateReservation(reserv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerReservation) UpdateStatusReserv(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	request := new(reservationsdto.ReservationReqUpdate)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	reserv, err := h.ReservationRepository.GetReservation(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	// Gunakan time.Time.Nil untuk menandakan bahwa field tersebut tidak diubah
	var orderProses, orderSelesai time.Time

	// Mengecek apakah request.OrderProses dan request.OrderSelesai berisi nilai yang tidak nol
	if !request.OrderProses.IsZero() {
		orderProses = request.OrderProses
	}

	if !request.OrderSelesai.IsZero() {
		orderSelesai = request.OrderSelesai
	}

	// Update nilai-nilai yang tidak nol
	reserv.OrderProses = orderProses
	reserv.OrderSelesai = orderSelesai

	if request.Status != "" {
		reserv.Status = request.Status
	}
	if request.SubStatus != "" {
		reserv.SubStatus = request.SubStatus
	}
	if request.TotalItem != 0 {
		reserv.TotalItem = request.TotalItem
	}
	if request.TotalPrice != 0 {
		reserv.TotalPrice = request.TotalPrice
	}

	reserv.UpdatedAt = time.Now()

	data, err := h.ReservationRepository.UpdateReservation(reserv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerReservation) DeleteReservation(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	brand, err := h.ReservationRepository.GetReservation(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.ReservationRepository.DeleteReservation(brand)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

// func respReservationChart(reservations []models.Reservation) []reservationsdto.ReservationChart {
// 	var charts []reservationsdto.ReservationChart

// 	for _, u := range reservations {
// 		month := time.Month(u.OrderMasuk.Month()).String()

// 		chart := reservationsdto.ReservationChart{
// 			MonthInt:   int(u.OrderMasuk.Month()),
// 			Month:      month,
// 			TotalItem:  u.TotalItem,
// 			TotalPrice: u.TotalPrice,
// 		}

// 		charts = append(charts, chart)
// 	}

// 	return charts
// }

// func respReservationChart(reservations []models.Reservation) map[int][]reservationsdto.ReservationChart {
// 	// Buat map untuk menyimpan data berdasarkan monthint
// 	chartMap := make(map[int][]reservationsdto.ReservationChart)

// 	// Iterasi melalui setiap reservasi
// 	for _, u := range reservations {
// 		// Ambil monthint dari reservasi
// 		monthInt := int(u.OrderMasuk.Month())

// 		// Buat struktur data ReservationChart
// 		chart := reservationsdto.ReservationChart{
// 			MonthInt:   monthInt,
// 			Month:      time.Month(monthInt).String(),
// 			TotalItem:  u.TotalItem,
// 			TotalPrice: u.TotalPrice,
// 		}

// 		// Tambahkan ReservationChart ke slice yang sesuai dengan monthint
// 		chartMap[monthInt] = append(chartMap[monthInt], chart)
// 	}

// 	return chartMap
// }

func respReservationChart(reservations []models.Reservation) []reservationsdto.ReservationChart {
	var charts []reservationsdto.ReservationChart

	// Iterasi melalui setiap reservasi
	for _, u := range reservations {
		// Ambil monthint dari reservasi
		monthInt := int(u.OrderMasuk.Month())

		// Buat struktur data ReservationChart
		chart := reservationsdto.ReservationChart{
			MonthInt:   monthInt,
			Month:      time.Month(monthInt).String(),
			TotalItem:  u.TotalItem,
			TotalPrice: u.TotalPrice,
		}

		// Tambahkan ReservationChart ke slice
		charts = append(charts, chart)
	}

	return charts
}
