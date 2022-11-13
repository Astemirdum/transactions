package handler_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Astemirdum/transactions/tx-user/internal/handler"

	userv1 "github.com/Astemirdum/transactions/proto/user/v1"
	"github.com/Astemirdum/transactions/proto/user/v1/userv1connect"
	"github.com/Astemirdum/transactions/tx-user/internal/service"
	"github.com/golang/mock/gomock"
	"go.uber.org/zap"

	service_mocks "github.com/Astemirdum/transactions/tx-user/internal/handler/mocks"
	"github.com/bufbuild/connect-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserSignUp(t *testing.T) {
	t.Parallel()

	type mockBehavior func(r *service_mocks.MockUserService, user *service.User)

	testTable := []struct {
		name               string
		req                *userv1.SignUpRequest
		mockBehavior       mockBehavior
		expectedStatusCode connect.Code
		expectedResponse   *userv1.SignUpResponse
		wantErr            bool
	}{
		{
			name: "ok. standard",
			mockBehavior: func(r *service_mocks.MockUserService, user *service.User) {
				r.EXPECT().CreateAccount(gomock.Any(), user).Return(&service.Account{
					UserID: 1,
					Cash:   1000,
				}, nil)
			},
			req: &userv1.SignUpRequest{
				User: &userv1.User{
					Email:    "lol@kek.ru",
					Password: "12345",
				},
			},
			expectedStatusCode: 0,
			expectedResponse: &userv1.SignUpResponse{
				UserId:  1,
				Balance: &userv1.Balance{Cash: 1000},
			},
		},
		{
			name: "ko. internal",
			mockBehavior: func(r *service_mocks.MockUserService, user *service.User) {
				r.EXPECT().CreateAccount(gomock.Any(), user).
					Return(nil, errors.New("internal"))
			},
			req: &userv1.SignUpRequest{
				User: &userv1.User{
					Email:    "lol@kek.ru",
					Password: "12345",
				},
			},
			expectedStatusCode: connect.CodeInternal,
			expectedResponse:   nil,
			wantErr:            true,
		},
		{
			name: "ko. login exist",
			mockBehavior: func(r *service_mocks.MockUserService, user *service.User) {
				r.EXPECT().CreateAccount(gomock.Any(), user).
					Return(nil, service.ErrEmailAlreadyExists)
			},
			req: &userv1.SignUpRequest{
				User: &userv1.User{
					Email:    "lol@kek.ru",
					Password: "12345",
				},
			},
			expectedStatusCode: connect.CodeAlreadyExists,
			expectedResponse:   nil,
			wantErr:            true,
		},
		{
			name:         "ko. invalid email",
			mockBehavior: func(r *service_mocks.MockUserService, user *service.User) {},
			req: &userv1.SignUpRequest{
				User: &userv1.User{
					Email:    "lolkek.ru",
					Password: "12345",
				},
			},
			expectedStatusCode: connect.CodeInvalidArgument,
			expectedResponse:   nil,
			wantErr:            true,
		},
		{
			name:         "ko. invalid password",
			mockBehavior: func(r *service_mocks.MockUserService, user *service.User) {},
			req: &userv1.SignUpRequest{
				User: &userv1.User{
					Email:    "lol@kek.ru",
					Password: "123",
				},
			},
			expectedStatusCode: connect.CodeInvalidArgument,
			expectedResponse:   nil,
			wantErr:            true,
		},
	}

	for _, testCase := range testTable {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			log := zap.NewExample().Named("signUp")
			svc := service_mocks.NewMockUserService(c)
			h := handler.NewHandlerTest(svc, log)
			mux := http.NewServeMux()
			mux.Handle(userv1connect.NewUserServiceHandler(h, connect.WithInterceptors(h.ValidateInterceptor())))
			server := httptest.NewUnstartedServer(mux)
			server.EnableHTTP2 = true
			server.Start()
			defer server.Close()

			client := userv1connect.NewUserServiceClient(
				server.Client(),
				server.URL,
				connect.WithGRPC(),
			)
			testCase.mockBehavior(svc, &service.User{
				Email:    testCase.req.User.Email,
				Password: testCase.req.User.Password,
			})
			result, err := client.SignUp(context.Background(), connect.NewRequest(testCase.req))
			if testCase.wantErr {
				require.Error(t, err)
				assert.Equal(t, connect.CodeOf(err), testCase.expectedStatusCode)
			} else {
				require.NoError(t, err)
				assert.Equal(t, result.Msg.UserId, testCase.expectedResponse.UserId)
				assert.Equal(t, result.Msg.Balance.Cash, testCase.expectedResponse.Balance.Cash)
			}
		})
	}
}
