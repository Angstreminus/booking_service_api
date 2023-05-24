package mocks

import (
	"booking_service_api/models"
	"context"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

// MockTokenService is a mock type for model.TokenService
type MockTokenService struct {
	mock.Mock
}

// NewPairFromUser mocks concrete NewPairFromUser
func (m *MockTokenService) NewPairFromUser(ctx context.Context, u *models.User, prevTokenID string) (*models.TokenPair, error) {
	ret := m.Called(ctx, u, prevTokenID)

	// first value passed to "Return"
	var r0 *models.TokenPair
	if ret.Get(0) != nil {
		// we can just return this if we know we won't be passing function to "Return"
		r0 = ret.Get(0).(*models.TokenPair)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

// Signout mocks concrete Signout
func (m *MockTokenService) Signout(ctx context.Context, uid uuid.UUID) error {
	ret := m.Called(ctx, uid)
	var r0 error

	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

// ValidateIDToken mocks concrete ValidateIDToken
func (m *MockTokenService) ValidateIDToken(tokenString string) (*models.User, error) {
	ret := m.Called(tokenString)

	// first value passed to "Return"
	var r0 *models.User
	if ret.Get(0) != nil {
		// we can just return this if we know we won't be passing function to "Return"
		r0 = ret.Get(0).(*models.User)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

// ValidateRefreshToken mocks concrete ValidateRefreshToken
func (m *MockTokenService) ValidateRefreshToken(refreshTokenString string) (*models.RefreshToken, error) {
	ret := m.Called(refreshTokenString)

	var r0 *models.RefreshToken
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*models.RefreshToken)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}
