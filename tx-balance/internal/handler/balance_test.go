package handler_test

import (
	"context"
	"errors"
	"fmt"

	"github.com/Astemirdum/transactions/tx-balance/config"
	"github.com/Astemirdum/transactions/tx-balance/internal/handler"

	"github.com/Astemirdum/transactions/pkg"
	"github.com/Astemirdum/transactions/tx-balance/internal/handler/broker"
	models "github.com/Astemirdum/transactions/tx-balance/models/v1"

	"net/http"
	"net/http/httptest"
	"testing"

	balancev1 "github.com/Astemirdum/transactions/proto/balance/v1"
	"github.com/Astemirdum/transactions/proto/balance/v1/balancev1connect"
	"github.com/golang/mock/gomock"
	"go.uber.org/zap"

	service_mocks "github.com/Astemirdum/transactions/tx-balance/internal/handler/mocks"
	"github.com/bufbuild/connect-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBalanceCashOut(t *testing.T) {
	t.Parallel()

	type (
		mockBehavior func(auth *service_mocks.MockAuth,
			br *service_mocks.MockCashOutBroker, req *balancev1.CashOutRequest, session string)

		input struct {
			req     *balancev1.CashOutRequest
			session string
		}
	)

	testTable := []struct {
		name               string
		inp                input
		mockBehavior       mockBehavior
		expectedStatusCode connect.Code
		expectedResponse   *balancev1.CashOutResponse
		wantErr            bool
	}{
		{
			name: "ok. standard",
			inp: input{
				req:     &balancev1.CashOutRequest{Cash: 50},
				session: "asldlmkRtO",
			},
			mockBehavior: func(auth *service_mocks.MockAuth,
				br *service_mocks.MockCashOutBroker, req *balancev1.CashOutRequest, session string) {

				var userID int32 = 1
				auth.EXPECT().Auth(gomock.Any(), session).Return(userID, nil)
				subject := fmt.Sprintf(broker.SubjectTmp, userID)
				br.EXPECT().PublishCashOut(subject, &models.CashOutMsg{
					Cash: uint64(req.Cash),
				})
			},
			expectedStatusCode: 0,
			expectedResponse:   &balancev1.CashOutResponse{},
		},
		{
			name: "ko. no session",
			inp: input{
				req: &balancev1.CashOutRequest{Cash: 50},
			},
			mockBehavior: func(auth *service_mocks.MockAuth,
				br *service_mocks.MockCashOutBroker, req *balancev1.CashOutRequest, session string) {

			},
			expectedStatusCode: connect.CodeUnauthenticated,
			expectedResponse:   &balancev1.CashOutResponse{},
			wantErr:            true,
		},
		{
			name: "ko. auth reject",
			inp: input{
				req:     &balancev1.CashOutRequest{Cash: 50},
				session: "wqeeqwFEee",
			},
			mockBehavior: func(auth *service_mocks.MockAuth,
				br *service_mocks.MockCashOutBroker, req *balancev1.CashOutRequest, session string) {

				auth.EXPECT().Auth(gomock.Any(), session).Return(int32(0),
					connect.NewError(connect.CodeInternal, errors.New("some auth prob")))
			},
			expectedStatusCode: connect.CodeInternal,
			expectedResponse:   &balancev1.CashOutResponse{},
			wantErr:            true,
		},
		{
			name: "ko. invalid cashOut arg <= 0",
			inp: input{
				req:     &balancev1.CashOutRequest{Cash: -10},
				session: "asldlmkRtO",
			},
			mockBehavior: func(auth *service_mocks.MockAuth,
				br *service_mocks.MockCashOutBroker, req *balancev1.CashOutRequest, session string) {

			},
			expectedStatusCode: connect.CodeInvalidArgument,
			expectedResponse:   &balancev1.CashOutResponse{},
			wantErr:            true,
		},
	}

	for _, testCase := range testTable {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mb := service_mocks.NewMockCashOutBroker(c)
			auth := service_mocks.NewMockAuth(c)
			log := zap.NewExample().Named("cashOut")
			h := handler.NewHandlerTest(nil, mb, auth, log, &config.Config{})
			mux := http.NewServeMux()
			mux.Handle(balancev1connect.NewBalanceServiceHandler(h,
				connect.WithInterceptors(h.ValidateInterceptor()),
				connect.WithInterceptors(h.AuthInterceptor()),
			))
			server := httptest.NewUnstartedServer(mux)
			server.EnableHTTP2 = true
			server.Start()
			defer server.Close()

			client := balancev1connect.NewBalanceServiceClient(
				server.Client(),
				server.URL,
				connect.WithGRPC(),
			)
			testCase.mockBehavior(auth, mb, testCase.inp.req, testCase.inp.session)
			req := connect.NewRequest(testCase.inp.req)
			req.Header().Add(pkg.SessionKey, testCase.inp.session)
			_, err := client.CashOut(context.Background(), req)
			if testCase.wantErr {
				require.Error(t, err)
				assert.Equal(t, connect.CodeOf(err), testCase.expectedStatusCode)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
