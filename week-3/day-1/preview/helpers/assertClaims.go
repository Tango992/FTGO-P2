package helpers

import "preview-week3/dto"

func AssertClaims(claims any) dto.Claims {
	claimsTemp := claims.(map[string]any)
	
	return dto.Claims{
		Id: uint(claimsTemp["id"].(float64)),
		Email: claimsTemp["email"].(string),
		Username: claimsTemp["username"].(string),
	}
}