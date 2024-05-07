package model

type (
		Station struct {
			Name        string    	`json:"name"`
			EnName      string  	`json:"en_name"`
			Lat         float64 	`json:"lat"`
			Long        float64  	`json:"long"`
			Distance    float64  	`json:"distance"`
				
				
		}

		StationFilter struct {
			Lat float64 `query:"lat" validate:"omitempty"`
			Long float64 `query:"long" validate:"omitempty"`
			Count	int	`query:"count" validate:"omitempty,min=1"`
		}

		StationPaginateFilter struct {
			Lat 	float64 `query:"lat" validate:"omitempty"`
			Long 	float64 `query:"long" validate:"omitempty"`
			Count	int		`query:"count" validate:"omitempty,min=1"`
			Page 	int64  `query:"page" validate:"required,gt=0"`
			Size 	int64  `query:"size" validate:"required,min=1,max=20"`
		}





)

