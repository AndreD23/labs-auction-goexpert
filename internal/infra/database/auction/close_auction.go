package auction

import (
	"context"
	"fmt"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/internal_error"
)

func (ar *AuctionRepository) CloseAuctions(ctx context.Context, auctions []auction_entity.Auction) *internal_error.InternalError {
	for _, a := range auctions {
		fmt.Printf("auction %s completed\n", a.Id)
	}

	return nil
}
