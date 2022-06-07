package model_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/tmrrwnxtsn/const-payments-api/internal/model"
	"reflect"
	"testing"
)

func TestParseStatus(t *testing.T) {
	type args struct {
		status string
	}
	testCases := []struct {
		name    string
		args    args
		want    model.Status
		wantErr bool
	}{
		{
			name: "valid (StatusNew)",
			args: args{
				status: "НОВЫЙ",
			},
			want:    model.StatusNew,
			wantErr: false,
		},
		{
			name: "valid (StatusSuccess)",
			args: args{
				status: "УСПЕХ",
			},
			want:    model.StatusSuccess,
			wantErr: false,
		},
		{
			name: "valid (StatusFailure)",
			args: args{
				status: "НЕУСПЕХ",
			},
			want:    model.StatusFailure,
			wantErr: false,
		},
		{
			name: "valid (StatusError)",
			args: args{
				status: "ОШИБКА",
			},
			want:    model.StatusError,
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				status: "НЕИЗВЕСТНО",
			},
			wantErr: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := model.ParseStatus(tc.args.status)
			if (err != nil) != tc.wantErr {
				t.Errorf("ParseStatus() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if got != tc.want {
				t.Errorf("ParseStatus() got = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestStatus_MarshalText(t *testing.T) {
	testCases := []struct {
		name    string
		s       model.Status
		want    []byte
		wantErr bool
	}{
		{
			name:    "valid (StatusNew)",
			s:       model.StatusNew,
			want:    []byte("НОВЫЙ"),
			wantErr: false,
		},
		{
			name:    "valid (StatusSuccess)",
			s:       model.StatusSuccess,
			want:    []byte("УСПЕХ"),
			wantErr: false,
		},
		{
			name:    "valid (StatusFailure)",
			s:       model.StatusFailure,
			want:    []byte("НЕУСПЕХ"),
			wantErr: false,
		},
		{
			name:    "valid (StatusError)",
			s:       model.StatusError,
			want:    []byte("ОШИБКА"),
			wantErr: false,
		},
		{
			name:    "invalid",
			s:       model.Status(55),
			wantErr: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.s.MarshalText()
			if (err != nil) != tc.wantErr {
				t.Errorf("MarshalText() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("MarshalText() got = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestStatus_String(t *testing.T) {
	testCases := []struct {
		name string
		s    model.Status
		want string
	}{
		{
			name: "valid",
			s:    model.StatusNew,
			want: "НОВЫЙ",
		},
		{
			name: "invalid",
			s:    model.Status(55),
			want: "НЕИЗВЕСТНО",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.s.String(); got != tc.want {
				t.Errorf("String() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestStatus_UnmarshalText(t *testing.T) {
	type args struct {
		text []byte
	}
	testCases := []struct {
		name    string
		s       model.Status
		args    args
		wantErr bool
	}{
		{
			name: "valid",
			s:    model.Status(0),
			args: args{
				text: []byte("НОВЫЙ"),
			},
			wantErr: false,
		},
		{
			name: "invalid",
			s:    model.Status(0),
			args: args{
				text: []byte("НЕИЗВЕСТНО"),
			},
			wantErr: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.s.UnmarshalText(tc.args.text); (err != nil) != tc.wantErr {
				t.Errorf("UnmarshalText() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestTransaction_Validate(t *testing.T) {
	testCases := []struct {
		name        string
		transaction func() *model.Transaction
		isValid     bool
	}{
		{
			name: "valid",
			transaction: func() *model.Transaction {
				return model.TestTransaction(t)
			},
			isValid: true,
		},
		{
			name: "wrong user email",
			transaction: func() *model.Transaction {
				testTransaction := model.TestTransaction(t)
				testTransaction.UserEmail = "tmrrwnxtsn"
				return testTransaction
			},
			isValid: false,
		},
		{
			name: "wrong currency code",
			transaction: func() *model.Transaction {
				testTransaction := model.TestTransaction(t)
				testTransaction.CurrencyCode = "ruble"
				return testTransaction
			},
			isValid: false,
		},
		{
			name: "wrong amount",
			transaction: func() *model.Transaction {
				testTransaction := model.TestTransaction(t)
				testTransaction.Amount = -1924.4
				return testTransaction
			},
			isValid: false,
		},
		{
			name: "empty fields",
			transaction: func() *model.Transaction {
				return &model.Transaction{}
			},
			isValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.transaction().Validate())
			} else {
				assert.Error(t, tc.transaction().Validate())
			}
		})
	}
}
