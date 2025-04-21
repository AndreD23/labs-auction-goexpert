package auction

import (
	"context"
	"fmt"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
)

func (ar *AuctionRepository) CloseAuctions(ctx context.Context, auctions []auction_entity.Auction) *internal_error.InternalError {
	for _, a := range auctions {
		updateOperators := bson.M{"$set": bson.M{"status": auction_entity.Completed}}
		_, err := ar.Collection.UpdateByID(ctx, a.Id, updateOperators)

		if err != nil {
			msg := fmt.Sprintf("err close auction. id: %s, err: %v", a.Id, err)
			return internal_error.NewInternalServerError(msg)
		}
	}

	return nil
}
