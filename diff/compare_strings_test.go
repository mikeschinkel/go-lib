package diff_test

import (
	"testing"

	"github.com/mikeschinkel/go-lib/diff"
)

func TestCompareStrings(t *testing.T) {
	type args struct {
		s1  string
		s2  string
		pad int
	}
	var tests = []struct {
		name string
		args args
		want string
	}{
		{
			name: "S1 and S2 are empty",
		},
		{
			name: "S1 is empty",
			args: args{
				s2: "ABC",
			},
			want: "(/ABC)",
		},
		{
			name: "S2 is empty",
			args: args{
				s1: "ABC",
			},
			want: "(ABC/)",
		},
		{
			name: "S1 and S2 are completely different",
			args: args{
				s1: "ABC",
				s2: "XYZ",
			},
			want: "(ABC/XYZ)",
		},
		{
			name: "S1 and S2 start the same, but end different",
			args: args{
				s1: "ABCDEF",
				s2: "ABCDXYZ",
			},
			want: "ABCD(EF/XYZ)",
		},
		{
			name: "S1 and S2 start different but end the same",
			args: args{
				s1: "ABCDXYZ",
				s2: "123XYZ",
			},
			want: "(ABCD/123)XYZ",
		},
		{
			name: "S1 has extra middle chars",
			args: args{
				s1:  "ABCDEF123GHIJKLMNOP",
				s2:  "ABCDEFGHIJKLMNOP",
				pad: 5,
			},
			want: "BCDEF(123/)GHIJK",
		},
		{
			name: "S1 has prefix and suffix that S2 does not have",
			args: args{
				s1: "123GHI456",
				s2: "GHI",
			},
			want: "(123/)GHI(456/)",
		},
		{
			name: "S1 and S2 share a middle, differ on the ends",
			args: args{
				s1: "123GHI789",
				s2: "987GHI321",
			},
			want: "(123/987)GHI(789/321)",
		},
		{
			name: "S1 has two sets of extra middle chars",
			args: args{
				s1:  "ABCDEF123GHI456JKLMNOP",
				s2:  "ABCDEFGHIJKLMNOP",
				pad: 5,
			},
			want: "BCDEF(123/)GHI(456/)JKLMN",
		},
		//{
		//	name: "Lorem Ipsum",
		//	args: args{
		//		s1:  "In publishing and graphic design, Lorem ipsum is a placeholder text commonly used to demonstrate the visual form of a document or a typeface without relying on meaningful content. Lorem ipsum may be used as a placeholder before final copy is available.",
		//		s2:  "In publishing & graphic design, Lorem ipsum is a commonly used text placeholder to demonstrate a document in its visual form, or a typeface sans meaningful content. Lorem ipsum is often used as a placeholder awaiting final copy.",
		//		pad: 25,
		//	},
		//	want: "",
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diff.CompareStrings(tt.args.s1, tt.args.s2, tt.args.pad); got != tt.want {
				t.Errorf("\ndiff.CompareStrings(s1,s2):\n\t got: %v\n\twant: %v\n", got, tt.want)
			}
		})
	}
}
