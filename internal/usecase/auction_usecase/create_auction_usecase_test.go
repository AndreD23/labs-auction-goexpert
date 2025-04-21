package auction_usecase

import (
	"context"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/infra/database/mocks"
	"fullcycle-auction_go/internal/internal_error"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuctionUseCase_CreateAuction(t *testing.T) {
	errT := internal_error.NewInternalServerError("teste")
	ea := map[string]*internal_error.InternalError{
		"teste": errT,
	}
	auctionRepository := mocks.NewMockAuctionRepository(ea)
	bidRepository := mocks.NewMockBidRepository(ea)

	usecase := NewAuctionUseCase(auctionRepository, bidRepository)

	dto := AuctionInputDTO{
		ProductName: "Casa Grande",
		Category:    "Imóveis",
		Description: "Casa com 3 quartos",
		Condition:   ProductCondition(auction_entity.New),
	}

	output := usecase.CreateAuction(context.Background(), dto)
	assert.Equal(t, output, (*internal_error.InternalError)(nil))

	ea = map[string]*internal_error.InternalError{
		"create_auction": errT,
	}
	auctionRepository = mocks.NewMockAuctionRepository(ea)
	usecase = NewAuctionUseCase(auctionRepository, bidRepository)

	output = usecase.CreateAuction(context.Background(), dto)
	assert.Equal(t, output, errT)
}
