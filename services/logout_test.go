package services

import (
	"atm-simulation/datasource/mock"
	"atm-simulation/schemas"
	"atm-simulation/utils"
	"testing"

	"go.uber.org/mock/gomock"
)

func Test_logout_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDatasource := mock.NewMockDatasources(ctrl)

	type args struct {
		cmd *schemas.Command
	}
	tests := []struct {
		name    string
		mocks   []*gomock.Call
		args    args
		wantErr bool
	}{
		{
			name: "Should return error when failed logout",
			mocks: []*gomock.Call{
				mockDatasource.EXPECT().Logout().
					Times(1).
					Return(utils.ErrorInvalidAccount),
			},
			wantErr: true,
		},
		{
			name: "Should return success",
			mocks: []*gomock.Call{
				mockDatasource.EXPECT().Logout().
					Times(1).
					Return(nil),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pl := &logout{
				repo: mockDatasource,
			}
			if err := pl.Execute(tt.args.cmd); (err != nil) != tt.wantErr {
				t.Errorf("logout.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
