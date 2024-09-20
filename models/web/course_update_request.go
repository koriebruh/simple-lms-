package web

// <-- maybe letter w add validation

type CourseUpdateRequest struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
	Teacher     int    `json:"teacher"`
}
