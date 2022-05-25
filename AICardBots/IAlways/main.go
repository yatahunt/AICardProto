package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	pb "AICard/AICardProto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:55055", "the address to connect to")
)

func main() {
	log.Printf("IAlways starting")
	var name = "IAlways"
	var command = pb.PlayerStatus_PlayerStatusCall
	var actionString = "Call"
	if len(os.Args) == 2 {
		switch os.Args[1] {

		case "call":
			command = pb.PlayerStatus_PlayerStatusCall
			actionString = "Call"
		case "raise":
			command = pb.PlayerStatus_PlayerStatusRaised
			actionString = "Raise"
		case "fold":
			command = pb.PlayerStatus_PlayerStatusFold
			actionString = "Fold"
		default:
			command = pb.PlayerStatus_PlayerStatusCall
			actionString = "Call"
		}
	}
	name = name + actionString
	log.Print("I will always " + actionString)
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGameServicesClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	var player = &pb.PlayerDefinition{Name: name, Secret: "donttellanyone"}
	log.Printf("asking to Join")
	var playerInscription *pb.PlayerInscription
	playerInscription, err = c.Join(ctx, player)
	if err != nil {
		log.Fatalf("could not join: %v", err)
	} else {
		log.Printf("Sending WaitingToStart")
		player = playerInscription.GetPlayer()
		var amount = playerInscription.Game.GetMinRaise()
		fmt.Print("amount is " + amount.String())
		var gameHistory = &pb.GameHistory{}
		gameHistory, err = c.Interact(ctx, &pb.GameInteraction{PlayerID: player.ID, Action: pb.PlayerStatus_PlayerStatusJoined, Amount: amount})
		var continueFor bool = (err == nil)
		for continueFor {
			var states = gameHistory.State
			var CurrentState = states[len(states)-1]

			if CurrentState.Status == pb.GameStatus_GameFinished {
				continueFor = false
			} else {
				fmt.Printf("Current state is %s \n", CurrentState.Status.String())
			}
			if continueFor {
				if CurrentState.BettedThisStage.Chips == 0 && actionString == "Call" {
					log.Printf("Sending Check")
					gameHistory, err = c.Interact(ctx, &pb.GameInteraction{PlayerID: player.ID, Action: pb.PlayerStatus_PlayerStatusCheck, Amount: amount})
				} else {
					log.Printf("Sending " + actionString)
					gameHistory, err = c.Interact(ctx, &pb.GameInteraction{PlayerID: player.ID, Action: command, Amount: amount})
				}
				if err != nil {
					log.Fatalf("Error Interacting with Server: %v", err)
					continueFor = false
				} else {
					log.Printf("Output Received")
				}

			}
		}
		if err != nil {
			log.Fatalf("could not interact: %v", err)
		}
	}

}
