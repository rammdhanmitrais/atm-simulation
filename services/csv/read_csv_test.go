package csv

import (
	"atm-simulation/datasource"
	"atm-simulation/datasource/mock"
	"atm-simulation/schemas"
	"atm-simulation/utils"
	"testing"

	"go.uber.org/mock/gomock"
)

func Test_readCsv_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDatasource := mock.NewMockUserDatasources(ctrl)
	type fields struct {
		repo datasource.UserDatasources
	}
	type args struct {
		cmd *schemas.Command
	}
	tests := []struct {
		name    string
		fields  fields
		mocks   []*gomock.Call
		args    args
		wantErr bool
	}{
		{
			name: "Should return error when user choose invalid option",
			fields: fields{
				repo: mockDatasource,
			},
			args: args{
				cmd: &schemas.Command{
					Arguments: schemas.Arguments{
						CsvArg: schemas.CsvArguments{
							Chosen: 3,
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Should return success when user choose default datasource",
			fields: fields{
				repo: mockDatasource,
			},
			mocks: []*gomock.Call{
				mockDatasource.EXPECT().InitiateDefaultUser().
					Times(1),
			},
			args: args{
				cmd: &schemas.Command{
					Arguments: schemas.Arguments{
						CsvArg: schemas.CsvArguments{
							Chosen: 1,
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Should return error when user choose read csv but path is empty",
			fields: fields{
				repo: mockDatasource,
			},
			args: args{
				cmd: &schemas.Command{
					Arguments: schemas.Arguments{
						CsvArg: schemas.CsvArguments{
							Chosen: 2,
							Path:   "",
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Should return error when user choose read csv but has invalid values",
			fields: fields{
				repo: mockDatasource,
			},
			args: args{
				cmd: &schemas.Command{
					Arguments: schemas.Arguments{
						CsvArg: schemas.CsvArguments{
							Chosen: 2,
							Path:   "../../utils/test_file/datasource_invalid.csv",
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Should return error when user choose read csv but has empty values",
			fields: fields{
				repo: mockDatasource,
			},
			args: args{
				cmd: &schemas.Command{
					Arguments: schemas.Arguments{
						CsvArg: schemas.CsvArguments{
							Chosen: 2,
							Path:   "../../utils/test_file/datasource_empty.csv",
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Should return error when user choose read csv but has duplicate account number",
			fields: fields{
				repo: mockDatasource,
			},
			args: args{
				cmd: &schemas.Command{
					Arguments: schemas.Arguments{
						CsvArg: schemas.CsvArguments{
							Chosen: 2,
							Path:   "../../utils/test_file/datasource_duplicate_account.csv",
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Should return error when fail to get user account number",
			fields: fields{
				repo: mockDatasource,
			},
			mocks: []*gomock.Call{
				mockDatasource.EXPECT().GetUserByAccountNumber(gomock.Any()).
					Times(1).
					Return(&datasource.User{}, nil),
			},
			args: args{
				cmd: &schemas.Command{
					Arguments: schemas.Arguments{
						CsvArg: schemas.CsvArguments{
							Chosen: 2,
							Path:   "../../utils/test_file/datasource.csv",
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Should return error when fail to insert user to datasource",
			fields: fields{
				repo: mockDatasource,
			},
			mocks: []*gomock.Call{
				mockDatasource.EXPECT().GetUserByAccountNumber(gomock.Any()).
					Times(1).
					Return(nil, nil),

				mockDatasource.EXPECT().InsertUser(gomock.Any()).
					Times(1).
					Return(utils.ErrorOccurs),
			},
			args: args{
				cmd: &schemas.Command{
					Arguments: schemas.Arguments{
						CsvArg: schemas.CsvArguments{
							Chosen: 2,
							Path:   "../../utils/test_file/datasource.csv",
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Should return success when user choose read csv",
			fields: fields{
				repo: mockDatasource,
			},
			mocks: []*gomock.Call{
				mockDatasource.EXPECT().GetUserByAccountNumber(gomock.Any()).
					Times(3).
					Return(nil, nil),

				mockDatasource.EXPECT().InsertUser(gomock.Any()).
					Times(3).
					Return(nil),
			},
			args: args{
				cmd: &schemas.Command{
					Arguments: schemas.Arguments{
						CsvArg: schemas.CsvArguments{
							Chosen: 2,
							Path:   "../../utils/test_file/datasource.csv",
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pl := &readCsv{
				repo: tt.fields.repo,
			}
			if err := pl.Execute(tt.args.cmd); (err != nil) != tt.wantErr {
				t.Errorf("readCsv.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
