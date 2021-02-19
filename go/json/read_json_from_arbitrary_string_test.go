package json

import "testing"

func TestReadJSONFromArbitraryString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		output  Data
		expect  Data
		wantErr bool
	}{
		{
			name:   "only valid json",
			input:  `{"foo":"foo","bar":"bar"}`,
			output: Data{},
			expect: Data{
				Foo: "foo",
				Bar: "bar",
			},
			wantErr: false,
		},
		{
			name:   "valid json with arbitrary suffix",
			input:  `{"foo":"foo","bar":"bar"} some suffix`,
			output: Data{},
			expect: Data{
				Foo: "foo",
				Bar: "bar",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReadJsonFromArbitraryString(tt.input, &tt.output); (err != nil) != tt.wantErr {
				t.Errorf("ReadJSONFromArbitraryString().  error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.output.Foo != tt.expect.Foo || tt.output.Bar != tt.expect.Bar {
				t.Errorf("Output mismatched. Expected: %v, Found: %v", tt.expect, tt.output)
			}
		})
	}
}
