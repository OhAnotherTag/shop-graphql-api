package main

import (
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

	"github.com/OhAnotherTag/shop-gql-api/config/database"
	"github.com/OhAnotherTag/shop-gql-api/graph"
	"github.com/OhAnotherTag/shop-gql-api/graph/generated"
	"github.com/OhAnotherTag/shop-gql-api/graph/model"
)

const defaultPort = "7000"

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	port := os.Getenv("PORT")
	
	if port == "" {
		port = defaultPort
	}

	database.ConnectDB()
	database.DB.AutoMigrate(&model.User{}, &model.Product{}, &model.Category{})

	// populateDb()

	r := gin.Default()

	r.Use(CORS())

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	r.Run(":" + defaultPort)
}

func populateDb() {
	categories := []model.Category{
		{
			Name:     "Baby",
			Products: []*model.Product{},
		},
		{
			Name:     "Home",
			Products: []*model.Product{},
		},
		{
			Name:     "Toys",
			Products: []*model.Product{},
		},
		{
			Name:     "Books",
			Products: []*model.Product{},
		},
		{
			Name:     "Industrial",
			Products: []*model.Product{},
		},
	}
	database.DB.Create(&categories)

	products := []model.Product{
		{
			Title:       "Refined Metal Chicken",
			Price:       589.00,
			Description: "New range of formal shirts are designed keeping you in mind. With fits and styling that will make you stand apart",
			CategoryID:  1,
		},
		{
			Title:       "Incredible Soft Tuna",
			Price:       779.00,
			Description: "New ABC 13 9370, 13.3, 5th Gen CoreA5-8250U, 8GB RAM, 256GB SSD, power UHD Graphics, OS 10 Home, OS Office A & J 2016",
			CategoryID:  2,
		},
		{
			Title:       "Licensed Rubber Chicken",
			Price:       87.00,
			Description: "The beautiful range of Apple Naturalé that has an exciting mix of natural ingredients. With the Goodness of 100% Natural Ingredients",
			CategoryID:  3,
		},
		{
			Title:       "Rustic Frozen Soap",
			Price:       302.00,
			Description: "Andy shoes are designed to keeping in mind durability as well as trends, the most stylish range of shoes & sandals",
			CategoryID:  4,
		},
		{
			Title:       "Small Granite Bacon",
			Price:       84.00,
			Description: "Andy shoes are designed to keeping in mind durability as well as trends, the most stylish range of shoes & sandals",
			CategoryID:  5,
		},
		{
			Title:       "Generic Rubber Towels",
			Price:       254.00,
			Description: "Carbonite web goalkeeper gloves are ergonomically designed to give easy fit",
			CategoryID:  1,
		},
		{
			Title:       "Tasty Fresh Car",
			Price:       48.00,
			Description: "New range of formal shirts are designed keeping you in mind. With fits and styling that will make you stand apart",
			CategoryID:  2,
		},
		{
			Title:       "Fantastic Granite Computer",
			Price:       9.00,
			Description: "The slim & simple Maple Gaming Keyboard from Dev Byte comes with a sleek body and 7- Color RGB LED Back-lighting for smart functionality",
			CategoryID:  3,
		},
		{
			Title:       "Gorgeous Cotton Keyboard",
			Price:       321.00,
			Description: "The Nagasaki Lander is the trademarked name of several series of Nagasaki sport bikes, that started with the 1984 ABC800J",
			CategoryID:  4,
		},
		{
			Title:       "Refined Concrete Chair",
			Price:       970.00,
			Description: "Carbonite web goalkeeper gloves are ergonomically designed to give easy fit",
			CategoryID:  5,
		},
		{
			Title:       "Refined Concrete Pants",
			Price:       510.00,
			Description: "The Football Is Good For Training And Recreational Purposes",
			CategoryID:  1,
		},
		{
			Title:       "Handmade Concrete Table",
			Price:       291.00,
			Description: "Carbonite web goalkeeper gloves are ergonomically designed to give easy fit",
			CategoryID:  1,
		},
		{
			Title:       "Ergonomic Granite Bacon",
			Price:       430.00,
			Description: "New ABC 13 9370, 13.3, 5th Gen CoreA5-8250U, 8GB RAM, 256GB SSD, power UHD Graphics, OS 10 Home, OS Office A & J 2016",
			CategoryID:  2,
		},
		{
			Title:       "Licensed Concrete Towels",
			Price:       16.00,
			Description: "New range of formal shirts are designed keeping you in mind. With fits and styling that will make you stand apart",
			CategoryID:  4,
		},
		{
			Title:       "Ergonomic Rubber Tuna",
			Price:       229.00,
			Description: "Ergonomic executive chair upholstered in bonded black leather and PVC padded seat and back for all-day comfort and support",
			CategoryID:  3,
		},
		{
			Title:       "Refined Cotton Cheese",
			Price:       202.00,
			Description: "The slim & simple Maple Gaming Keyboard from Dev Byte comes with a sleek body and 7- Color RGB LED Back-lighting for smart functionality",
			CategoryID:  1,
		},
		{
			Title:       "Small Metal Table",
			Price:       911.00,
			Description: "The slim & simple Maple Gaming Keyboard from Dev Byte comes with a sleek body and 7- Color RGB LED Back-lighting for smart functionality",
			CategoryID:  4,
		},
		{
			Title:       "Refined Concrete Pizza",
			Price:       964.00,
			Description: "Andy shoes are designed to keeping in mind durability as well as trends, the most stylish range of shoes & sandals",
			CategoryID:  5,
		},
		{
			Title:       "Intelligent Steel Pants",
			Price:       903.00,
			Description: "The beautiful range of Apple Naturalé that has an exciting mix of natural ingredients. With the Goodness of 100% Natural Ingredients",
			CategoryID:  2,
		},
		{
			Title:       "Handmade Plastic Bike",
			Price:       824.00,
			Description: "The Nagasaki Lander is the trademarked name of several series of Nagasaki sport bikes, that started with the 1984 ABC800J",
			CategoryID:  2,
		},
	}

	database.DB.Create(&products)
}
