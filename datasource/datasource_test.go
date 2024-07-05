package datasource

import (
	"reflect"
	"testing"
)

var ds = new(datasource)

func Test_datasource_GetUserByAccountNumber(t *testing.T) {

	type args struct {
		accountNumber string
	}
	tests := []struct {
		name       string
		d          datasource
		args       args
		wantResult User
		wantErr    bool
	}{
		{
			name: "Should return error when account not found",
			args: args{
				accountNumber: "123456",
			},
			d:          *ds,
			wantErr:    true,
			wantResult: User{},
		},
		{
			name: "Should return success",
			args: args{
				accountNumber: "112233",
			},
			d:          *ds,
			wantErr:    false,
			wantResult: userAccounts[0],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := datasource{}
			gotResult, err := d.GetUserByAccountNumber(tt.args.accountNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("datasource.GetUserByAccountNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("datasource.GetUserByAccountNumber() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_datasource_GetLoggedUser(t *testing.T) {
	tests := []struct {
		name     string
		d        datasource
		wantUser User
		wantErr  bool
	}{
		{
			name:     "Should return error when account not found",
			d:        *ds,
			wantErr:  true,
			wantUser: User{},
		},
		{
			name:     "Should return success",
			d:        *ds,
			wantErr:  false,
			wantUser: userAccounts[0],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := datasource{}

			userEmpty := User{}
			if tt.wantUser != userEmpty {
				LoggedUser = &userAccounts[0]
			}

			gotUser, err := d.GetLoggedUser()
			if (err != nil) != tt.wantErr {
				t.Errorf("datasource.GetLoggedUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("datasource.GetLoggedUser() = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}

func Test_datasource_UpdateUserBalance(t *testing.T) {
	type args struct {
		id      int
		balance int64
	}
	tests := []struct {
		name    string
		d       datasource
		args    args
		wantErr bool
	}{
		{
			name: "Should return error when account not found",
			args: args{
				id:      -1,
				balance: 1000,
			},
			d:       *ds,
			wantErr: true,
		},
		{
			name: "Should return success",
			args: args{
				id:      1,
				balance: 1000,
			},
			d:       *ds,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := datasource{}
			if err := d.UpdateUserBalance(tt.args.id, tt.args.balance); (err != nil) != tt.wantErr {
				t.Errorf("datasource.UpdateUserBalance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_datasource_Login(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		d       datasource
		args    args
		wantErr bool
	}{
		{
			name: "Should return error when account not found",
			args: args{
				id: -1,
			},
			d:       *ds,
			wantErr: true,
		},
		{
			name: "Should return success",
			args: args{
				id: 1,
			},
			d:       *ds,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := datasource{}

			if !tt.wantErr {
				LoggedUser = &userAccounts[0]
			}

			if err := d.Login(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("datasource.Login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_datasource_Logout(t *testing.T) {
	tests := []struct {
		name    string
		d       datasource
		wantErr bool
	}{
		{
			name:    "Should return success",
			d:       *ds,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := datasource{}
			if err := d.Logout(); (err != nil) != tt.wantErr {
				t.Errorf("datasource.Logout() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
