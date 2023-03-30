package models

import (
	"github.com/joyal007/go-scan/config"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{
	UserID		primitive.ObjectID		`json:"userid"`
}

type Users struct{
	mgm.DefaultModel `bson:",inline"`
	Name		string		`json:"name" bson:"name"`
	Team		string		`json:"team_name" bson:"team_name"`
	YearOfStudy string		`json:"year_of_study" bson:"year_of_study" `
	Institution	string		`json:"institution" bson:"institution"`
	Number		string		`json:"number" bson:"number"`
	CheckIn		bool		`json:"checkin" bson:"checkin"`
	TeamStrength int		`json:"team_strength" bson:"team_strength"`
	PaymentCompleted bool	`json:"payment_completed" bson:"payment_completed"`
	Food		string		`json:"food" bson:"food"`
}

type Ars struct{
	mgm.DefaultModel `bson:",inline"`
	Name		string		`json:"name" `
	YearOfStudy string		`json:"year_of_study" bson:"year_of_study" `
	Branch		string		`json:"branch" `
	Institution	string		`json:"institution" `
	Number		string		`json:"number" `
	CheckIn		bool		`json:"checkin"`
	PaymentCompleted bool	`json:"payment_completed" bson:"payment_completed"`
}
type Devops struct{
	mgm.DefaultModel `bson:",inline"`
	Name		string		`json:"name" `
	YearOfStudy string		`json:"year_of_study" bson:"year_of_study" `
	Branch		string		`json:"branch" `
	Institution	string		`json:"institution" `
	Number		string		`json:"number" `
	CheckIn		bool		`json:"checkin"`
	PaymentCompleted bool	`json:"payment_completed" bson:"payment_completed"`
}
type Flutters struct{
	mgm.DefaultModel `bson:",inline"`
	Name		string		`json:"name" `
	YearOfStudy string		`json:"year_of_study" bson:"year_of_study" `
	Branch		string		`json:"branch" `
	Institution	string		`json:"institution" `
	Number		string		`json:"number" `
	CheckIn		bool		`json:"checkin"`
	PaymentCompleted bool	`json:"payment_completed" bson:"payment_completed"`
}
type Rusts struct{
	mgm.DefaultModel `bson:",inline"`
	Name		string		`json:"name" `
	YearOfStudy string		`json:"year_of_study" bson:"year_of_study" `
	Branch		string		`json:"branch" `
	Institution	string		`json:"institution" `
	Number		string		`json:"number" `
	CheckIn		bool		`json:"checkin"`
	PaymentCompleted bool	`json:"payment_completed" bson:"payment_completed"`
}
type Webs struct{
	mgm.DefaultModel `bson:",inline"`
	Name		string		`json:"name" `
	YearOfStudy string		`json:"year_of_study" bson:"year_of_study" `
	Branch		string		`json:"branch" `
	Institution	string		`json:"institution" `
	Number		string		`json:"number" `
	CheckIn		bool		`json:"checkin"`
	PaymentCompleted bool	`json:"payment_completed" bson:"payment_completed"`
}
type Gos struct{
	mgm.DefaultModel `bson:",inline"`
	Name		string		`json:"name" `
	YearOfStudy string		`json:"year_of_study" bson:"year_of_study" `
	Branch		string		`json:"branch" `
	Institution	string		`json:"institution" `
	Number		string		`json:"number" `
	CheckIn		bool		`json:"checkin"`
	PaymentCompleted bool	`json:"payment_completed" bson:"payment_completed"`
}


func init(){
	config.Connect()
}