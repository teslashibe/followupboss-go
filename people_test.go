package followupboss

import "testing"

func TestNormalizeFields(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"empty", "", ""},
		{"email alias", "email", "emails"},
		{"phone alias", "phone", "phones"},
		{"both aliases", "email,phone", "emails,phones"},
		{"mixed with other tokens", "firstName,email,stage,phone", "firstName,emails,stage,phones"},
		{"already plural untouched", "emails,phones", "emails,phones"},
		{"trims whitespace", " email , phone , stage ", "emails,phones,stage"},
		{"allFields untouched", "allFields", "allFields"},
		{"case variant Email", "Email", "emails"},
		{"case variant PHONE", "PHONE", "phones"},
		{"mixed case variants", "Email,PHONE", "emails,phones"},
		{"empty inner token", "email,,phone", "emails,phones"},
		{"trailing empty token", "email,", "emails"},
		{"case-sensitive projection preserved", "firstName,Email", "firstName,emails"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := normalizeFields(tc.in); got != tc.want {
				t.Errorf("normalizeFields(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
