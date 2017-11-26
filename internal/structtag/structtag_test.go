package structtag_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/apex/log"
	"github.com/matthewmueller/golly/internal/structtag"
)

func TestParse(t *testing.T) {
	test := []struct {
		name    string
		tag     string
		exp     []*structtag.Tag
		invalid bool
	}{
		{
			name: "empty tag",
			tag:  "",
		},
		{
			name:    "tag with one key (invalid)",
			tag:     "json",
			invalid: true,
		},
		{
			name: "tag with one key (valid)",
			tag:  `json:""`,
			exp: []*structtag.Tag{
				{
					Key: "json",
				},
			},
		},
		{
			name: "tag with one key and dash name",
			tag:  `json:"-"`,
			exp: []*structtag.Tag{
				{
					Key:  "json",
					Name: "-",
				},
			},
		},
		{
			name: "tag with key and name",
			tag:  `json:"foo"`,
			exp: []*structtag.Tag{
				{
					Key:  "json",
					Name: "foo",
				},
			},
		},
		{
			name: "tag with key, name and option",
			tag:  `json:"foo,omitempty"`,
			exp: []*structtag.Tag{
				{
					Key:     "json",
					Name:    "foo",
					Options: []string{"omitempty"},
				},
			},
		},
		{
			name: "tag with multiple keys",
			tag:  `json:"" hcl:""`,
			exp: []*structtag.Tag{
				{
					Key: "json",
				},
				{
					Key: "hcl",
				},
			},
		},
		{
			name: "tag with multiple keys and names",
			tag:  `json:"foo" hcl:"foo"`,
			exp: []*structtag.Tag{
				{
					Key:  "json",
					Name: "foo",
				},
				{
					Key:  "hcl",
					Name: "foo",
				},
			},
		},
		{
			name: "tag with multiple keys and names",
			tag:  `json:"foo" hcl:"foo"`,
			exp: []*structtag.Tag{
				{
					Key:  "json",
					Name: "foo",
				},
				{
					Key:  "hcl",
					Name: "foo",
				},
			},
		},
		{
			name: "tag with multiple keys and different names",
			tag:  `json:"foo" hcl:"bar"`,
			exp: []*structtag.Tag{
				{
					Key:  "json",
					Name: "foo",
				},
				{
					Key:  "hcl",
					Name: "bar",
				},
			},
		},
		{
			name: "tag with multiple keys, different names and options",
			tag:  `json:"foo,omitempty" structs:"bar,omitnested"`,
			exp: []*structtag.Tag{
				{
					Key:     "json",
					Name:    "foo",
					Options: []string{"omitempty"},
				},
				{
					Key:     "structs",
					Name:    "bar",
					Options: []string{"omitnested"},
				},
			},
		},
		{
			name: "tag with multiple keys, different names and options",
			tag:  `json:"foo" structs:"bar,omitnested" hcl:"-"`,
			exp: []*structtag.Tag{
				{
					Key:  "json",
					Name: "foo",
				},
				{
					Key:     "structs",
					Name:    "bar",
					Options: []string{"omitnested"},
				},
				{
					Key:  "hcl",
					Name: "-",
				},
			},
		},
	}

	for _, ts := range test {
		t.Run(ts.name, func(t *testing.T) {
			tags, err := structtag.Parse(ts.tag)
			invalid := err != nil

			if invalid != ts.invalid {
				t.Errorf("invalid case\n\twant: %+v\n\tgot : %+v\n\terr : %s", ts.invalid, invalid, err)
			}

			if invalid {
				return
			}

			got := tags.Tags()

			if !reflect.DeepEqual(ts.exp, got) {
				t.Errorf("parse\n\twant: %#v\n\tgot : %#v", ts.exp, got)
			}
		})
	}
}

func TestTags_Get(t *testing.T) {
	tag := `json:"foo,omitempty" structs:"bar,omitnested"`

	tags, err := structtag.Parse(tag)
	if err != nil {
		t.Fatal(err)
	}

	found, err := tags.Get("json")
	if err != nil {
		t.Fatal(err)
	}

	want := `json:"foo,omitempty"`

	if found.String() != want {
		t.Errorf("get\n\twant: %#v\n\tgot : %#v", want, found.String())
	}
}

func TestTags_Set(t *testing.T) {
	tag := `json:"foo,omitempty" structs:"bar,omitnested"`

	tags, err := structtag.Parse(tag)
	if err != nil {
		t.Fatal(err)
	}

	err = tags.Set(&structtag.Tag{
		Key:     "json",
		Name:    "bar",
		Options: []string{},
	})
	if err != nil {
		t.Fatal(err)
	}

	found, err := tags.Get("json")
	if err != nil {
		t.Fatal(err)
	}

	want := `json:"bar"`
	if found.String() != want {
		t.Errorf("set\n\twant: %#v\n\tgot : %#v", want, found.String())
	}
}

func TestTags_Set_Append(t *testing.T) {
	tag := `json:"foo,omitempty"`

	tags, err := structtag.Parse(tag)
	if err != nil {
		t.Fatal(err)
	}

	err = tags.Set(&structtag.Tag{
		Key:     "structs",
		Name:    "bar",
		Options: []string{"omitnested"},
	})
	if err != nil {
		t.Fatal(err)
	}

	found, err := tags.Get("structs")
	if err != nil {
		t.Fatal(err)
	}

	want := `structs:"bar,omitnested"`
	if found.String() != want {
		t.Errorf("set append\n\twant: %#v\n\tgot : %#v", want, found.String())
	}

	wantFull := `json:"foo,omitempty" structs:"bar,omitnested"`
	if tags.String() != wantFull {
		t.Errorf("set append\n\twant: %#v\n\tgot : %#v", wantFull, tags.String())
	}
}

func TestTags_Set_KeyDoesNotExist(t *testing.T) {
	tag := `json:"foo,omitempty" structs:"bar,omitnested"`

	tags, err := structtag.Parse(tag)
	if err != nil {
		t.Fatal(err)
	}

	err = tags.Set(&structtag.Tag{
		Key:     "",
		Name:    "bar",
		Options: []string{},
	})
	if err == nil {
		t.Fatal("setting tag with a nonexisting key should error")
	}

	if err != structtag.ErrKeyNotSet {
		t.Errorf("set\n\twant: %#v\n\tgot : %#v", structtag.ErrTagKeyMismatch, err)
	}
}

func TestTags_Delete(t *testing.T) {
	tag := `json:"foo,omitempty" structs:"bar,omitnested" hcl:"-"`

	tags, err := structtag.Parse(tag)
	if err != nil {
		t.Fatal(err)
	}

	tags.Delete("structs")
	if tags.Len() != 2 {
		t.Fatalf("tag length should be 2, have %d", tags.Len())
	}

	found, err := tags.Get("json")
	if err != nil {
		t.Fatal(err)
	}

	want := `json:"foo,omitempty"`
	if found.String() != want {
		t.Errorf("delete\n\twant: %#v\n\tgot : %#v", want, found.String())
	}

	wantFull := `json:"foo,omitempty" hcl:"-"`
	if tags.String() != wantFull {
		t.Errorf("delete\n\twant: %#v\n\tgot : %#v", wantFull, tags.String())
	}
}

func TestTags_DeleteOptions(t *testing.T) {
	tag := `json:"foo,omitempty" structs:"bar,omitnested,omitempty" hcl:"-"`

	tags, err := structtag.Parse(tag)
	if err != nil {
		t.Fatal(err)
	}

	tags.DeleteOptions("json", "omitempty")

	want := `json:"foo" structs:"bar,omitnested,omitempty" hcl:"-"`
	if tags.String() != want {
		t.Errorf("delete option\n\twant: %#v\n\tgot : %#v", want, tags.String())
	}

	tags.DeleteOptions("structs", "omitnested")
	want = `json:"foo" structs:"bar,omitempty" hcl:"-"`
	if tags.String() != want {
		t.Errorf("delete option\n\twant: %#v\n\tgot : %#v", want, tags.String())
	}
}

func TestTags_AddOption(t *testing.T) {
	tag := `json:"foo" structs:"bar,omitempty" hcl:"-"`

	tags, err := structtag.Parse(tag)
	if err != nil {
		t.Fatal(err)
	}

	tags.AddOptions("json", "omitempty")

	want := `json:"foo,omitempty" structs:"bar,omitempty" hcl:"-"`
	if tags.String() != want {
		t.Errorf("add options\n\twant: %#v\n\tgot : %#v", want, tags.String())
	}

	// this shouldn't change anything
	tags.AddOptions("structs", "omitempty")

	want = `json:"foo,omitempty" structs:"bar,omitempty" hcl:"-"`
	if tags.String() != want {
		t.Errorf("add options\n\twant: %#v\n\tgot : %#v", want, tags.String())
	}

	// this should append to the existing
	tags.AddOptions("structs", "omitnested", "flatten")
	want = `json:"foo,omitempty" structs:"bar,omitempty,omitnested,flatten" hcl:"-"`
	if tags.String() != want {
		t.Errorf("add options\n\twant: %#v\n\tgot : %#v", want, tags.String())
	}
}

func TestTags_String(t *testing.T) {
	tag := `json:"foo" structs:"bar,omitnested" hcl:"-"`

	tags, err := structtag.Parse(tag)
	if err != nil {
		t.Fatal(err)
	}

	if tags.String() != tag {
		t.Errorf("string\n\twant: %#v\n\tgot : %#v", tag, tags.String())
	}
}

func TestTags_Sort(t *testing.T) {
	tag := `json:"foo" structs:"bar,omitnested" hcl:"-"`

	tags, err := structtag.Parse(tag)
	if err != nil {
		t.Fatal(err)
	}

	sort.Sort(tags)

	want := `hcl:"-" json:"foo" structs:"bar,omitnested"`
	if tags.String() != want {
		t.Errorf("string\n\twant: %#v\n\tgot : %#v", want, tags.String())
	}
}

func TestTags_Function(t *testing.T) {
	tag := `js:"foo,rewrite('1,2,3', 'hi\\''),omit"`

	tags, err := structtag.Parse(tag)
	if err != nil {
		t.Fatal(err)
	}

	for _, tag := range tags.Tags() {
		for _, opt := range tag.Options {
			log.Infof("option=%s", opt)
		}
	}

	if tags.String() != tag {
		t.Errorf("string\n\twant: %#v\n\tgot : %#v", tag, tags.String())
	}
}
