package src

import "testing"

// Exemplo de Entrada           Exemplos de saída
// 3								
// 2 3 5						  N

// 5
// 5 5 5 						  S

// 9
// 15 9 10						  S




func TestCheckHasFit(t *testing.T) {
	tests := []struct {
		name     string
		ball     Ball
		box      Box
		expected bool
	}{
		{
			name:     "Ball fits 1",
			ball:     Ball{Diameter: 5},
			box:      Box{Height: 5, Width: 5, Length: 5},
			expected: true,
		},
		{
			name:     "Ball fits 2",
			ball:     Ball{Diameter: 9},
			box:      Box{Height: 15, Width: 9, Length: 10},
			expected: true,
		},
		{
			name:     "Ball large 1",
			ball:     Ball{Diameter: 25},
			box:      Box{Height: 30, Width: 20, Length: 40},
			expected: false,
		},
		{
			name:     "Ball large 2",
			ball:     Ball{Diameter: 3},
			box:      Box{Height: 2, Width: 3, Length: 5},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CheckHasFit(tt.ball, tt.box)
			if result != tt.expected {
				t.Errorf("CheckHasFit(%v, %v) = %v; want %v", tt.ball, tt.box, result, tt.expected)
			}
		})
	}
}
