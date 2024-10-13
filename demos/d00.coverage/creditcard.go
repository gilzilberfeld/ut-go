package demos

import (
	"time"
)

type CreditCard struct {
    expirationDate time.Time
}

func (card *CreditCard) IsExpired() bool {
    localExpDate := card.expirationDate.Local()
    today := time.Now()
    if today.Before(localExpDate) {
		return true 
	} else {
		return false
	}
}
