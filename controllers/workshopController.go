package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joyal007/go-scan/models"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	
)


func GoWorkshop(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil{
		fmt.Println(err)
	}
	fmt.Println(user.UserID)

	var userFound models.Gos

	coll := mgm.Coll(&models.Gos{})

	_ = coll.FindByID(user.UserID, &userFound)
	if userFound.Name == ""{
		c.JSON(http.StatusBadRequest, gin.H{"error":"user not found"})
		return
	}
	if !userFound.PaymentCompleted{
		c.JSON(http.StatusBadRequest, gin.H{"error":"payment not completed"})
		return
	}
	userFound.CheckIn = true
	err :=mgm.Coll(&userFound).Update(&userFound)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"internal server error"})
		return
	}
	c.JSON(http.StatusOK,userFound)	
}
func ArWorkshop(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil{
		fmt.Println(err)
	}
	fmt.Println(user.UserID)

	var userFound models.Ars

	coll := mgm.Coll(&models.Ars{})

	_ = coll.FindByID(user.UserID, &userFound)
	if userFound.Name == ""{
		c.JSON(http.StatusBadRequest, gin.H{"error":"user not found"})
		return
	}
	if !userFound.PaymentCompleted{
		c.JSON(http.StatusBadRequest, gin.H{"error":"payment not completed"})
		return
	}
	userFound.CheckIn = true
	err :=mgm.Coll(&userFound).Update(&userFound)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"internal server error"})
		return
	}
	c.JSON(http.StatusOK,userFound)	
}
func RustWorkshop(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil{
		fmt.Println(err)
	}
	fmt.Println(user.UserID)

	var userFound models.Rusts

	coll := mgm.Coll(&models.Rusts{})

	_ = coll.FindByID(user.UserID, &userFound)
	if userFound.Name == ""{
		c.JSON(http.StatusBadRequest, gin.H{"error":"user not found"})
		return
	}
	if !userFound.PaymentCompleted{
		c.JSON(http.StatusBadRequest, gin.H{"error":"payment not completed"})
		return
	}
	userFound.CheckIn = true
	err :=mgm.Coll(&userFound).Update(&userFound)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"internal server error"})
		return
	}
	c.JSON(http.StatusOK,userFound)	
}
func WebdevWorkshop(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil{
		fmt.Println(err)
	}
	fmt.Println(user.UserID)

	var userFound models.Webs

	coll := mgm.Coll(&models.Webs{})

	_ = coll.FindByID(user.UserID, &userFound)
	if userFound.Name == ""{
		c.JSON(http.StatusBadRequest, gin.H{"error":"user not found"})
		return
	}
	if !userFound.PaymentCompleted{
		c.JSON(http.StatusBadRequest, gin.H{"error":"payment not completed"})
		return
	}
	userFound.CheckIn = true
	err :=mgm.Coll(&userFound).Update(&userFound)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"internal server error"})
		return
	}
	c.JSON(http.StatusOK,userFound)	
}
func FlutterWorkshop(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil{
		fmt.Println(err)
	}
	fmt.Println(user.UserID)

	var userFound models.Flutters

	coll := mgm.Coll(&models.Flutters{})

	_ = coll.FindByID(user.UserID, &userFound)
	if userFound.Name == ""{
		c.JSON(http.StatusBadRequest, gin.H{"error":"user not found"})
		return
	}
	if !userFound.PaymentCompleted{
		c.JSON(http.StatusBadRequest, gin.H{"error":"payment not completed"})
		return
	}
	userFound.CheckIn = true
	err :=mgm.Coll(&userFound).Update(&userFound)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"internal server error"})
		return
	}
	c.JSON(http.StatusOK,userFound)	
}

func DevopsWorkshop(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil{
		fmt.Println(err)
	}
	fmt.Println(user.UserID)

	var userFound models.Devops

	coll := mgm.Coll(&models.Devops{})

	_ = coll.FindByID(user.UserID, &userFound)
	if userFound.Name == ""{
		c.JSON(http.StatusBadRequest, gin.H{"error":"user not found"})
		return
	}
	if !userFound.PaymentCompleted{
		c.JSON(http.StatusBadRequest, gin.H{"error":"payment not completed"})
		return
	}
	userFound.CheckIn = true
	err :=mgm.Coll(&userFound).Update(&userFound)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"internal server error"})
		return
	}
	c.JSON(http.StatusOK,userFound)
}

func Hackfit(c *gin.Context){
	var user models.User
	if err := c.BindJSON(&user); err != nil{
		fmt.Println(err)
	}
	fmt.Println(user.UserID)

	var userFound models.Users

	coll := mgm.Coll(&models.Users{})

	_ = coll.FindByID(user.UserID, &userFound)
	if userFound.Name == ""{
		c.JSON(http.StatusBadRequest, gin.H{"error":"user not found"})
		return
	}
	if !userFound.PaymentCompleted{
		c.JSON(http.StatusBadRequest, gin.H{"error":"payment not completed"})
		return
	}
	userFound.CheckIn = true
	err :=mgm.Coll(&userFound).Update(&userFound)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"internal server error"})
		return
	}
	c.JSON(http.StatusOK,userFound)
}

func HackfitCheckIn(c *gin.Context){
	coll := mgm.Coll(&models.Users{})
	cursor, err := coll.Find(context.TODO(), bson.M{"checkin":true,"payment_completed":true})
	if err != nil {
		log.Fatal(err)
	}

	var results []models.Users
	if err := cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}

	if(len(results)==0){
		c.JSON(http.StatusInternalServerError,gin.H{"error":"none checked in"})
		return
	}

	c.JSON(http.StatusOK, results)

}


func GoWorkshopCheckIn(c *gin.Context) {

	coll := mgm.Coll(&models.Gos{})
	cursor, err := coll.Find(context.TODO(), bson.M{"checkin":true,"payment_completed":true})
	if err != nil {
		log.Fatal(err)
	}

	var results []models.Gos
	if err := cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}

	if(len(results)==0){
		c.JSON(http.StatusInternalServerError,gin.H{"error":"none checked in"})
		return
	}

	c.JSON(http.StatusOK, results)

}
func ArWorkshopCheckIn(c *gin.Context) {

	coll := mgm.Coll(&models.Ars{})
	cursor, err := coll.Find(context.TODO(), bson.M{"checkin":true,"payment_completed":true})
	if err != nil {
		log.Fatal(err)
	}

	var results []models.Ars
	if err := cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}

	if(len(results)==0){
		c.JSON(http.StatusInternalServerError,gin.H{"error":"none checked in"})
		return
	}

	c.JSON(http.StatusOK, results)

}
func WebWorkshopCheckIn(c *gin.Context) {

	coll := mgm.Coll(&models.Webs{})
	cursor, err := coll.Find(context.TODO(), bson.M{"checkin":true,"payment_completed":true})
	if err != nil {
		log.Fatal(err)
	}

	var results []models.Webs
	if err := cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}

	if(len(results)==0){
		c.JSON(http.StatusInternalServerError,gin.H{"error":"none checked in"})
		return
	}

	c.JSON(http.StatusOK, results)

}
func FlutterWorkshopCheckIn(c *gin.Context) {

	coll := mgm.Coll(&models.Flutters{})
	cursor, err := coll.Find(context.TODO(), bson.M{"checkin":true,"payment_completed":true})
	if err != nil {
		log.Fatal(err)
	}

	var results []models.Flutters
	if err := cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}

	if(len(results)==0){
		c.JSON(http.StatusInternalServerError,gin.H{"error":"none checked in"})
		return
	}

	c.JSON(http.StatusOK, results)

}

func DevopsWorkshopCheckIn(c *gin.Context) {

	coll := mgm.Coll(&models.Devops{})
	cursor, err := coll.Find(context.TODO(), bson.M{"checkin":true,"payment_completed":true})
	if err != nil {
		log.Fatal(err)
	}

	var results []models.Devops
	if err := cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}

	if(len(results)==0){
		c.JSON(http.StatusInternalServerError,gin.H{"error":"none checked in"})
		return
	}

	c.JSON(http.StatusOK, results)

}
func RustWorkshopCheckIn(c *gin.Context) {

	coll := mgm.Coll(&models.Rusts{})
	cursor, err := coll.Find(context.TODO(), bson.M{"checkin":true,"payment_completed":true})
	if err != nil {
		log.Fatal(err)
	}

	var results []models.Rusts
	if err := cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}

	if(len(results)==0){
		c.JSON(http.StatusInternalServerError,gin.H{"error":"none checked in"})
		return
	}

	c.JSON(http.StatusOK, results)

}


