package day02

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_parseInput(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    Sequences
		wantErr bool
	}{
		{
			name: "",
			args: args{
				reader: strings.NewReader(
					"abcdef\n" +
						"bababc\n" +
						"abbcde\n",
				),
			},
			want: Sequences{
				struct {
					characterCount characterCount
					tupleCount     tupleCount
					chars          []string
				}{
					characterCount: characterCount{
						count:   map[string]int{"a": 0, "b": 0, "c": 0, "d": 0, "e": 0, "f": 0, "g": 0, "h": 0, "i": 0, "j": 0, "k": 0, "l": 0, "m": 0, "n": 0, "o": 0, "p": 0, "q": 0, "r": 0, "s": 0, "t": 0, "u": 0, "v": 0, "w": 0, "x": 0, "y": 0, "z": 0},
						counted: false,
					},
					tupleCount: tupleCount{
						count:   map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0, 10: 0, 11: 0, 12: 0, 13: 0, 14: 0, 15: 0, 16: 0, 17: 0, 18: 0, 19: 0, 20: 0, 21: 0, 22: 0, 23: 0, 24: 0, 25: 0, 26: 0},
						counted: false,
					},
					chars: []string{
						"a", "b", "c", "d", "e", "f",
					},
				}, {
					characterCount: characterCount{
						count:   map[string]int{"a": 0, "b": 0, "c": 0, "d": 0, "e": 0, "f": 0, "g": 0, "h": 0, "i": 0, "j": 0, "k": 0, "l": 0, "m": 0, "n": 0, "o": 0, "p": 0, "q": 0, "r": 0, "s": 0, "t": 0, "u": 0, "v": 0, "w": 0, "x": 0, "y": 0, "z": 0},
						counted: false,
					},
					tupleCount: tupleCount{
						count:   map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0, 10: 0, 11: 0, 12: 0, 13: 0, 14: 0, 15: 0, 16: 0, 17: 0, 18: 0, 19: 0, 20: 0, 21: 0, 22: 0, 23: 0, 24: 0, 25: 0, 26: 0},
						counted: false,
					},
					chars: []string{
						"b", "a", "b", "a", "b", "c",
					},
				}, {
					characterCount: characterCount{
						count:   map[string]int{"a": 0, "b": 0, "c": 0, "d": 0, "e": 0, "f": 0, "g": 0, "h": 0, "i": 0, "j": 0, "k": 0, "l": 0, "m": 0, "n": 0, "o": 0, "p": 0, "q": 0, "r": 0, "s": 0, "t": 0, "u": 0, "v": 0, "w": 0, "x": 0, "y": 0, "z": 0},
						counted: false,
					},
					tupleCount: tupleCount{
						count:   map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0, 10: 0, 11: 0, 12: 0, 13: 0, 14: 0, 15: 0, 16: 0, 17: 0, 18: 0, 19: 0, 20: 0, 21: 0, 22: 0, 23: 0, 24: 0, 25: 0, 26: 0},
						counted: false,
					},
					chars: []string{
						"a", "b", "b", "c", "d", "e",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseInputAsSequences(tt.args.reader)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInputAsSequences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInput_solvePartOne(t *testing.T) {
	type fields struct {
		Filename     string
		TargetTuples []int
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{
			name: "Test solvePartOne",
			fields: fields{
				Filename:     "./testdata/testinput.txt",
				TargetTuples: []int{2, 3},
			},
			want:    12,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Input{
				Filename:     tt.fields.Filename,
				TargetTuples: tt.fields.TargetTuples,
			}
			input, err := openInput(i.Filename)
			if err != nil {
				t.Errorf("Error opening test files %v", err)
			}
			sequences := parseInputAsSequences(input)
			got, err := sequences.solvePartOne(i.TargetTuples)
			if (err != nil) != tt.wantErr {
				t.Errorf("Input.solvePartOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Input.solvePartOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkInput_solvePartOne(b *testing.B) {
	type fields struct {
		Filename     string
		TargetTuples []int
	}
	benchmarks := []struct {
		name   string
		fields fields
	}{
		{
			name: "Test solvePartOne",
			fields: fields{
				Filename:     "./testdata/testinput.txt",
				TargetTuples: []int{2, 3},
			},
		},
	}
	for _, bb := range benchmarks {
		b.Run(bb.name, func(b *testing.B) {
			i := Input{
				Filename:     bb.fields.Filename,
				TargetTuples: bb.fields.TargetTuples,
			}

			input, err := openInput(i.Filename)
			if err != nil {
				b.Errorf("Error opening test files %v", err)
			}
			sequences := parseInputAsSequences(input)

			for n := 0; n < b.N; n++ {
				_, _ = sequences.solvePartOne(i.TargetTuples)
			}
		})
	}
}

func TestSequence_howManyTuplesOf(t *testing.T) {
	type fields struct {
		characterCount characterCount
		tupleCount     tupleCount
		chars          []string
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "",
			fields: fields{
				characterCount: characterCount{
					count:   map[string]int{"a": 0, "b": 0, "c": 0, "d": 0, "e": 0, "f": 0, "g": 0, "h": 0, "i": 0, "j": 0, "k": 0, "l": 0, "m": 0, "n": 0, "o": 0, "p": 0, "q": 0, "r": 0, "s": 0, "t": 0, "u": 0, "v": 0, "w": 0, "x": 0, "y": 0, "z": 0},
					counted: false,
				},
				tupleCount: tupleCount{
					count:   map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0, 10: 0, 11: 0, 12: 0, 13: 0, 14: 0, 15: 0, 16: 0, 17: 0, 18: 0, 19: 0, 20: 0, 21: 0, 22: 0, 23: 0, 24: 0, 25: 0, 26: 0},
					counted: false,
				},
				chars: []string{"b", "a", "b", "a", "b", "c"},
			},
			args: args{
				n: 2,
			},
			want: 1,
		}, {
			name: "",
			fields: fields{
				characterCount: characterCount{
					count:   map[string]int{"a": 0, "b": 0, "c": 0, "d": 0, "e": 0, "f": 0, "g": 0, "h": 0, "i": 0, "j": 0, "k": 0, "l": 0, "m": 0, "n": 0, "o": 0, "p": 0, "q": 0, "r": 0, "s": 0, "t": 0, "u": 0, "v": 0, "w": 0, "x": 0, "y": 0, "z": 0},
					counted: false,
				},
				tupleCount: tupleCount{
					count:   map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0, 10: 0, 11: 0, 12: 0, 13: 0, 14: 0, 15: 0, 16: 0, 17: 0, 18: 0, 19: 0, 20: 0, 21: 0, 22: 0, 23: 0, 24: 0, 25: 0, 26: 0},
					counted: false,
				},
				chars: []string{"b", "a", "b", "a", "b", "c"},
			},
			args: args{
				n: 3,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Sequence{
				characterCount: tt.fields.characterCount,
				tupleCount:     tt.fields.tupleCount,
				chars:          tt.fields.chars,
			}
			if got := s.howManyTuplesOf(tt.args.n); got != tt.want {
				t.Errorf("Sequence.howManyTuplesOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSequence_howManyTuplesOf(b *testing.B) {
	type fields struct {
		characterCount characterCount
		tupleCount     tupleCount
		chars          []string
	}
	type args struct {
		n int
	}
	benchmarks := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "",
			fields: fields{
				characterCount: characterCount{
					count:   map[string]int{"a": 0, "b": 0, "c": 0, "d": 0, "e": 0, "f": 0, "g": 0, "h": 0, "i": 0, "j": 0, "k": 0, "l": 0, "m": 0, "n": 0, "o": 0, "p": 0, "q": 0, "r": 0, "s": 0, "t": 0, "u": 0, "v": 0, "w": 0, "x": 0, "y": 0, "z": 0},
					counted: false,
				},
				tupleCount: tupleCount{
					count:   map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0, 10: 0, 11: 0, 12: 0, 13: 0, 14: 0, 15: 0, 16: 0, 17: 0, 18: 0, 19: 0, 20: 0, 21: 0, 22: 0, 23: 0, 24: 0, 25: 0, 26: 0},
					counted: false,
				},
				chars: []string{"b", "a", "b", "a", "b", "c"},
			},
			args: args{
				n: 2,
			},
			want: 1,
		}, {
			name: "",
			fields: fields{
				characterCount: characterCount{
					count:   map[string]int{"a": 0, "b": 0, "c": 0, "d": 0, "e": 0, "f": 0, "g": 0, "h": 0, "i": 0, "j": 0, "k": 0, "l": 0, "m": 0, "n": 0, "o": 0, "p": 0, "q": 0, "r": 0, "s": 0, "t": 0, "u": 0, "v": 0, "w": 0, "x": 0, "y": 0, "z": 0},
					counted: false,
				},
				tupleCount: tupleCount{
					count:   map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0, 10: 0, 11: 0, 12: 0, 13: 0, 14: 0, 15: 0, 16: 0, 17: 0, 18: 0, 19: 0, 20: 0, 21: 0, 22: 0, 23: 0, 24: 0, 25: 0, 26: 0},
					counted: false,
				},
				chars: []string{"b", "a", "b", "a", "b", "c"},
			},
			args: args{
				n: 3,
			},
			want: 1,
		},
	}
	for _, bb := range benchmarks {
		b.Run(bb.name, func(b *testing.B) {
			s := Sequence{
				characterCount: bb.fields.characterCount,
				tupleCount:     bb.fields.tupleCount,
				chars:          bb.fields.chars,
			}
			for n := 0; n < b.N; n++ {
				s.howManyTuplesOf(bb.args.n)
			}
		})
	}
}

func Test_hamming(t *testing.T) {
	type args struct {
		a []string
		b []string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "",
			args: args{
				a: []string{"b", "a", "b", "a", "b", "c"},
				b: []string{"b", "a", "b", "a", "b", "d"},
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hamming(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("hamming() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("hamming() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSequence_findMatchForSequence(t *testing.T) {
	type fields struct {
		characterCount characterCount
		tupleCount     tupleCount
		chars          []string
	}
	type args struct {
		haystack []Sequence
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Sequence
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				characterCount: characterCount{},
				tupleCount:     tupleCount{},
				chars:          []string{"t","j","a","m","o","e","w","p","q","k","y","a","i","z","v","o","n","d","g","f","l","u","b","s","z","c"},

			},
			args: args{
				haystack: []Sequence{
					{
						characterCount: characterCount{},
						tupleCount:     tupleCount{},
						chars:          []string{"x","j","a","m","o","e","w","p","q","k","y","a","i","z","v","o","n","d","g","f","l","u","b","s","z","c"},
					}, {
						characterCount: characterCount{},
						tupleCount:     tupleCount{},
						chars:          []string{"x","u","a","m","o","e","w","p","q","k","y","a","i","z","v","o","n","d","g","f","l","u","b","s","z","c"},
					}, {
						characterCount: characterCount{},
						tupleCount:     tupleCount{},
						chars:          []string{"a","x","v","m","o","e","w","p","q","k","y","a","i","z","v","o","n","d","g","f","l","u","b","s","z","c"},
					},
				},
			},

			want: Sequence{
				characterCount: characterCount{},
				tupleCount:     tupleCount{},
				chars:          []string{"x","j","a","m","o","e","w","p","q","k","y","a","i","z","v","o","n","d","g","f","l","u","b","s","z","c"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Sequence{
				characterCount: tt.fields.characterCount,
				tupleCount:     tt.fields.tupleCount,
				chars:          tt.fields.chars,
			}
			got, err := s.findMatchForSequence(tt.args.haystack)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sequence.findMatchForSequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sequence.findMatchForSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
