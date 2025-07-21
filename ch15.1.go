
package main

import "fmt"

func (a *analytics) handleEmailBounce(em email) error {
	if err:=em.recipient.updateStatus(em.status);err!=nil{
		return fmt.Errorf("error updating user status: %w",err)
	}
	a.track(em.status)
	return nil
}

