 

import grpc

from AICardProto import AICardProto_pb2, AICardProto_pb2_grpc


def run():
    with grpc.insecure_channel('localhost:55055') as channel:
        stub = AICardProto_pb2_grpc.GameServicesStub(channel)
        player = AICardProto_pb2.PlayerDefinition(Name='PythonAlwaysRaises'  , Secret= "donttellanyone")
        print("asking To Join")
        response   = stub.Join(player)
        
        print("Player Inscription ID: " + str(response.Player.ID.ID))
        minRaise = response.Game.MinRaise
        print("MinRaise: " + str(minRaise))
        playerDefinition = response.Player
        gameInteraction = AICardProto_pb2.GameInteraction(PlayerID=playerDefinition.ID, Action=AICardProto_pb2.PlayerStatusJoined, Amount=minRaise)
        stub = AICardProto_pb2_grpc.GameServicesStub(channel)
        gameHistory = stub.Interact(gameInteraction) 
        validStatus = True
        while validStatus:
            lastState = len(gameHistory.State) - 1
            print("Game Status: " +  str(gameHistory.State[lastState].Status))
            if gameHistory.State[lastState].Status == AICardProto_pb2.GameFinished:
                validStatus = False
            else:
                print("Waiting for Game to start")
                gameInteraction = AICardProto_pb2.GameInteraction(PlayerID=playerDefinition.ID, Action=AICardProto_pb2.PlayerStatusRaised, Amount=minRaise)
                gameHistory = stub.Interact(gameInteraction) 
            
        
 
    print("End of program, did we win?")

run() 
