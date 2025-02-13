package main

import (
	"auction_go/configuration/database/mongodb"
	"auction_go/internal/infra/api/web/controller/auction_controller"
	"auction_go/internal/infra/api/web/controller/bid_controller"
	"auction_go/internal/infra/api/web/controller/user_controller"
	"auction_go/internal/infra/database/auction"
	"auction_go/internal/infra/database/bid"
	"auction_go/internal/infra/database/user"
	"auction_go/internal/usecase/auction_usecase"
	"auction_go/internal/usecase/bid_usecase"
	"auction_go/internal/usecase/user_usecase"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	ctx := context.Background()
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error trying to load .env file")
		return
	}

	databaseConnection, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	router := gin.Default()

	userController, bidController, auctionsController := initDependencies(databaseConnection)

	router.GET("/auctions", auctionsController.FindAuctions)
	router.GET("/auctions/:auctionId", auctionsController.FindAuctionById)
	router.POST("/auctions", auctionsController.CreateAuction)
	router.GET("/auction/winner/:auctionId", auctionsController.FindWinningBidByAuctionId)

	router.POST("/bid", bidController.CreateBid)
	router.GET("/bid/:auctionId", bidController.FindBidByAuctionId)

	router.GET("/user/:userId", userController.FindUserById)

	router.Run(":8080")
}

func initDependencies(database *mongo.Database) (
	userController *user_controller.UserController,
	bidController *bid_controller.BidController,
	auctionController *auction_controller.AuctionController) {

	auctionRepository := auction.NewAuctionRepository(database)
	bidRepository := bid.NewBidRepository(database, auctionRepository)
	userRepository := user.NewUserRepository(database)

	userController = user_controller.NewUserController(user_usecase.NewUserUseCase(userRepository))
	auctionController = auction_controller.NewAuctionController(auction_usecase.NewAuctionUseCase(auctionRepository, bidRepository))
	bidController = bid_controller.NewBidController(bid_usecase.NewBidUseCase(bidRepository))

	return
}
