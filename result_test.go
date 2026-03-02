package result_test

import (
	"errors"
	"testing"

	"github.com/njchilds90/go-result"
)

func TestResult(t *testing.T) {
	tests := []struct {
		name     string
		r        result.Result[int]
		wantOk   bool
		wantVal  int
		wantErr  error
	}{
		{
			name:    "success",
			r:       result.Ok(42),
			wantOk:  true,
			wantVal: 42,
		},
		{
			name:    "error",
			r:       result.Err[int](errors.New("boom")),
			wantOk:  false,
			wantErr: errors.New("boom"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.r.IsOk() != tt.wantOk {
				t.Errorf("IsOk() = %v, want %v", tt.r.IsOk(), tt.wantOk)
			}
			if tt.r.IsErr() == tt.wantOk {
				t.Errorf("IsErr() = %v, want %v", tt.r.IsErr(), !tt.wantOk)
			}

			if tt.wantOk {
				if got := tt.r.Value(); got != tt.wantVal {
					t.Errorf("Value() = %v, want %v", got, tt.wantVal)
				}
				if got := tt.r.UnwrapOr(999); got != tt.wantVal {
					t.Errorf("UnwrapOr() = %v, want %v", got, tt.wantVal)
				}
			} else {
				if got := tt.r.Error(); got.Error() != tt.wantErr.Error() {
					t.Errorf("Error() = %v, want %v", got, tt.wantErr)
				}
			}
		})
	}
}

func TestResult_Chain(t *testing.T) {
	double := func(x int) result.Result[int] {
		return result.Ok(x * 2)
	}
	fail := func(int) result.Result[int] {
		return result.Err[int](errors.New("failed"))
	}

	r1 := result.Ok(5).AndThen(double).AndThen(double)
	if !r1.IsOk() || r1.Value() != 20 {
		t.Errorf("expected 20, got %v", r1)
	}

	r2 := result.Ok(5).AndThen(double).AndThen(fail)
	if !r2.IsErr() || r2.Error().Error() != "failed" {
		t.Errorf("expected error, got %v", r2)
	}
}

func TestResult_Map(t *testing.T) {
	r := result.Ok(10).Map(func(x int) int { return x * 10 })
	if !r.IsOk() || r.Value() != 100 {
		t.Errorf("Map failed, got %v", r)
	}
}
