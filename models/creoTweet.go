package models

import "time"

/*CreoTweet es la estructura que tendrá nuestro CreoTweet en la base datos */
type CreoTweet struct {
	UserID  string    `bson:"userid" json:"userid,omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}
