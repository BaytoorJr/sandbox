package private

import (
	"context"
	"encoding/json"
	"fmt"
	pb "gitlab.globerce.com/freedom-business/libs/protobuf-files/companypb"

	"google.golang.org/grpc"
	"log"
	"net/http"
)

func EventsHandler(w http.ResponseWriter, r *http.Request) {
	// Set headers for SSE
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Create a gRPC connection
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		http.Error(w, "Failed to connect to gRPC server", http.StatusInternalServerError)
		log.Println("Failed to connect to gRPC server:", err)
		return
	}
	defer conn.Close()

	client := pb.NewSellingPointServiceClient(conn)
	//client := localpb.NewEventServiceClient(conn)

	// Stream events from the gRPC server
	stream, err := client.ListenPaymentChanges(context.Background(), &pb.ListenPaymentChangesRequest{
		QrID: "b435deca-4c97-4980-9a56-43b7862d9fbf",
	})

	//stream, err := client.StreamEvents(context.Background(), &localpb.Empty{})
	//if err != nil {
	//	return
	//}

	if err != nil {
		http.Error(w, "Failed to start streaming events", http.StatusInternalServerError)
		log.Println("Failed to start streaming events:", err)
		return
	}

	ctx := r.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("Client disconnected")
			return
		default:
			event, err := stream.Recv()
			if err != nil {
				log.Println("Error receiving event from gRPC:", err)
				return
			}

			log.Println("Event:", event)

			// Convert the gRPC message to JSON
			jsonData, err := json.Marshal(event)
			if err != nil {
				http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
				return
			}

			// Send the JSON data to the client
			fmt.Fprintf(w, "data: %s\n\n", jsonData)
			w.(http.Flusher).Flush()
		}
	}
}
