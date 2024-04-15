package models

type Car struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Year       int     `json:"year"`
	Brand      string  `json:"brand"`
	Model      string  `json:"model"`
	HorsePower int     `json:"horsePower"`
	Colour     string  `json:"colour"`
	EngineCap  float32 `json:"engineCap"`
	CreatedAt  string  `json:"createdAt"`
	UpdatedAt  string  `json:"updatedAt"`
	DeleteAt   string  `json:"deleteAt"`
}

type GetAllCarsResponse struct {
	Cars  []Car `json:"cars"`
	Count int64 `json:"count"`
}

type GetAllCarsRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}

type DeleteCar struct {
	Id       string `json:"id"`
	DeleteAt string `json:"deleteAt"`
}

type UpdateCarRequest struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Year       int64   `json:"year"`
	Brand      string  `json:"brand"`
	Model      string  `json:"model"`
	HorsePower int     `json:"horsePower"`
	Colour     string  `json:"colour"`
	EngineCap  float32 `json:"engineCap"`
}
