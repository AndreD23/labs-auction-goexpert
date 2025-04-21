package mocks

import (
	"context"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/entity/bid_entity"
	"fullcycle-auction_go/internal/internal_error"
)

type MockAuctionRepository struct {
	Error map[string]*internal_error.InternalError
}

func NewMockAuctionRepository(error map[string]*internal_error.InternalError) *MockAuctionRepository {
	return &MockAuctionRepository{
		Error: error,
	}
}

func (m *MockAuctionRepository) CreateAuction(ctx context.Context, auction *auction_entity.Auction) *internal_error.InternalError {
	if err, ok := m.Error["create_auction"]; ok {
		return err
	}

	return nil
}

func (m *MockAuctionRepository) CloseAuctions(ctx context.Context, auctions []auction_entity.Auction) *internal_error.InternalError {
	if err, ok := m.Error["create_auction"]; ok {
		return err
	}

	return nil
}

func (m *MockAuctionRepository) FindAuctionById(ctx context.Context, id string) (*auction_entity.Auction, *internal_error.InternalError) {
	if err, ok := m.Error["find_auctions_by_id"]; ok {
		return &auction_entity.Auction{}, err
	}

	return &auction_entity.Auction{Id: id}, nil
}

func (m *MockAuctionRepository) FindAuctions(ctx context.Context, status auction_entity.AuctionStatus, category string, productName string) ([]auction_entity.Auction, *internal_error.InternalError) {
	if err, ok := m.Error["find_auctions"]; ok {
		return []auction_entity.Auction{}, err
	}

	return []auction_entity.Auction{}, nil
}

type MockBidRepository struct {
	Error map[string]*internal_error.InternalError
}

func NewMockBidRepository(error map[string]*internal_error.InternalError) *MockBidRepository {
	return &MockBidRepository{
		Error: error,
	}
}

func (m *MockBidRepository) CreateBid(ctx context.Context, bidEntities []bid_entity.Bid) *internal_error.InternalError {
	if err, ok := m.Error["create_bid"]; ok {
		return err
	}

	return nil
}

func (m *MockBidRepository) FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*bid_entity.Bid, *internal_error.InternalError) {
	if err, ok := m.Error["find_winning_bid"]; ok {
		return &bid_entity.Bid{}, err
	}

	return &bid_entity.Bid{Id: "1"}, nil
}

func (m *MockBidRepository) FindBidByAuctionId(ctx context.Context, auctionId string) ([]bid_entity.Bid, *internal_error.InternalError) {
	if err, ok := m.Error["find_bid_by_auction"]; ok {
		return []bid_entity.Bid{}, err
	}

	return []bid_entity.Bid{}, nil
}
