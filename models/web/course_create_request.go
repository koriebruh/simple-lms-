package web

// <-- maybe letter w add validation

type CourseCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
	Teacher     int    `json:"teacher"`
}
