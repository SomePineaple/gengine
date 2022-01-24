package utils

const DegreesToRadians = 0.017453292519943295

func ToRadians(deg float64) float64 {
	return deg * DegreesToRadians
}
