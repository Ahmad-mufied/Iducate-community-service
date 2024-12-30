package data

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

type PaginatedFeedQuery struct {
	Limit    int    `json:"limit" validate:"gte=1,lte=20"`
	Offset   int    `json:"offset" validate:"gte=0"`
	SortType string `json:"sortType" validate:"oneof=trend latest"`
	Sort     string `json:"sort" validate:"oneof=asc desc"`
}

func (fq *PaginatedFeedQuery) Parse(c echo.Context) error {
	qs := c.QueryParams()

	// Helper function for parsing integers with default fallback
	parseInt := func(key string, defaultValue, minValue, maxValue int) int {
		if val := qs.Get(key); val != "" {
			if parsed, err := strconv.Atoi(val); err == nil && parsed >= minValue && parsed <= maxValue {
				return parsed
			}
		}
		return defaultValue
	}

	// Helper function for parsing strings with default fallback
	parseString := func(key, defaultValue string, validValues ...string) string {
		if val := qs.Get(key); val != "" {
			for _, v := range validValues {
				if val == v {
					return val
				}
			}
		}
		return defaultValue
	}

	// Parse limit (1–20, default: 10)
	fq.Limit = parseInt("limit", 10, 1, 20)

	// Parse offset (>= 0, default: 0)
	fq.Offset = parseInt("offset", 0, 0, int(^uint(0)>>1)) // Max int value for offset

	// Parse sortType (valid: "trend", "latest"; default: "trend")
	fq.SortType = parseString("sortType", "trend", "trend", "latest")

	// Parse sort (valid: "asc", "desc"; default: "desc")
	fq.Sort = parseString("sort", "desc", "asc", "desc")

	return nil
}
