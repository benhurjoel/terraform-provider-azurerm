package afdendpoints

import "fmt"

const defaultApiVersion = "2021-06-01"

func userAgent() string {
	return fmt.Sprintf("pandora/afdendpoints/%s", defaultApiVersion)
}