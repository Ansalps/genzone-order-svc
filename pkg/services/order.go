package services

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	cartpb "github.com/Ansalps/genzone-order-svc/pkg/cart/pb"
	"github.com/Ansalps/genzone-order-svc/pkg/db"
	"github.com/Ansalps/genzone-order-svc/pkg/models"
	orderpb "github.com/Ansalps/genzone-order-svc/pkg/pb"
	"google.golang.org/grpc"
)

type Server struct {
	H db.Handler
	orderpb.UnimplementedOrderServiceServer
}

func (s *Server) CreateOrder(ctx context.Context, req *orderpb.CreateOrderRequest) (*orderpb.CreateOrderResponse, error) {
	var order models.Order
	fmt.Println("req", req)
	order.UserID = req.Userid
	order.AddressID = req.Addressid
	//cart.Qty = uint(req.Quantity)
	// Fetch product details
	cart, err := getCartDetails(req.Userid)
	if err != nil {
		log.Printf("Error fetching product details: %v", err)
	}
	// Calculate total amount
	// cart.Price = product.Price
	// totalAmount := product.Price * float64(req.Quantity)
	// cart.Amount = totalAmount
	//var product models.Product
	if result := s.H.DB.Create(&cart); result.Error != nil {
		return &orderpb.CreateOrderResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}
	return &orderpb.CreateOrderResponse{
		Status: http.StatusCreated,
		Id:     int64(cart.ID),
	}, nil
}
func getCartDetails(userID string) (*cartpb.GetCartResponse, error) {
	// Connect to Product Service
	conn, err := grpc.Dial("localhost:50055", grpc.WithInsecure()) // Replace with proper address
	if err != nil {
		return nil, fmt.Errorf("failed to connect to product service: %v", err)
	}
	defer conn.Close()

	client := cartpb.NewCartServiceClient(conn)

	// Call GetProduct RPC
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := client.GetCart(ctx, &cartpb.GetCartRequest{UserId: userID})
	if err != nil {
		return nil, fmt.Errorf("failed to get cart details: %v", err)
	}
	//fmt.Println("price",response.Price)
	return response, nil
}
