package internal

import (
	"github.com/vikyd/zero"

	"nimona.io"

	"nimona.io/tilde"
)

var (
	_ = zero.IsZeroVal
	_ = tilde.NewScanner
)

func (t *Feed) Document() *nimona.Document {
	return nimona.NewDocument(t.Map())
}

func (t *Feed) Map() tilde.Map {
	m := tilde.Map{}

	// # t.$type
	//
	// Type: string, Kind: string, TildeKind: InvalidValueKind0
	// IsSlice: false, IsStruct: false, IsPointer: false
	{
		m.Set("$type", tilde.String("feed"))
	}

	// # t.Folowees
	//
	// Type: []*nimona.Follow, Kind: slice, TildeKind: List
	// IsSlice: true, IsStruct: false, IsPointer: false
	//
	// ElemType: nimona.Follow, ElemKind: struct
	// IsElemSlice: false, IsElemStruct: true, IsElemPointer: true
	{
		if !zero.IsZeroVal(t.Folowees) {
			sm := tilde.List{}
			for _, v := range t.Folowees {
				if !zero.IsZeroVal(t.Folowees) {
					sm = append(sm, v.Map())
				}
			}
			m.Set("folowees", sm)
		}
	}

	// # t.Metadata
	//
	// Type: nimona.Metadata, Kind: struct, TildeKind: Map
	// IsSlice: false, IsStruct: true, IsPointer: false
	{
		m.Set("$metadata", t.Metadata.Map())
	}

	// # t.Notes
	//
	// Type: []*nimona.Note, Kind: slice, TildeKind: List
	// IsSlice: true, IsStruct: false, IsPointer: false
	//
	// ElemType: nimona.Note, ElemKind: struct
	// IsElemSlice: false, IsElemStruct: true, IsElemPointer: true
	{
		if !zero.IsZeroVal(t.Notes) {
			sm := tilde.List{}
			for _, v := range t.Notes {
				if !zero.IsZeroVal(t.Notes) {
					sm = append(sm, v.Map())
				}
			}
			m.Set("posts", sm)
		}
	}

	// # t.Profile
	//
	// Type: nimona.Profile, Kind: struct, TildeKind: Map
	// IsSlice: false, IsStruct: true, IsPointer: false
	{
		if !zero.IsZeroVal(t.Profile) {
			m.Set("profile", t.Profile.Map())
		}
	}

	return m
}

func (t *Feed) FromDocument(d *nimona.Document) error {
	return t.FromMap(d.Map())
}

func (t *Feed) FromMap(d tilde.Map) error {
	*t = Feed{}

	// # t.Folowees
	//
	// Type: []*nimona.Follow, Kind: slice, TildeKind: List
	// IsSlice: true, IsStruct: false, IsPointer: false
	//
	// ElemType: nimona.Follow, ElemKind: struct, ElemTildeKind: Map
	// IsElemSlice: false, IsElemStruct: true, IsElemPointer: true
	{
		sm := []*Follow{} // Follow
		if vs, err := d.Get("folowees"); err == nil {
			if vs, ok := vs.(tilde.List); ok {
				for _, vi := range vs {
					if v, ok := vi.(tilde.Map); ok {
						e := &Follow{}
						d := nimona.NewDocument(v)
						e.FromDocument(d)
						sm = append(sm, e)
					}
				}
			}
		}
		if len(sm) > 0 {
			t.Folowees = sm
		}
	}

	// # t.Metadata
	//
	// Type: nimona.Metadata, Kind: struct, TildeKind: Map
	// IsSlice: false, IsStruct: true, IsPointer: false
	{
		if v, err := d.Get("$metadata"); err == nil {
			if v, ok := v.(tilde.Map); ok {
				e := nimona.Metadata{}
				d := nimona.NewDocument(v)
				e.FromDocument(d)
				t.Metadata = e
			}
		}
	}

	// # t.Notes
	//
	// Type: []*nimona.Note, Kind: slice, TildeKind: List
	// IsSlice: true, IsStruct: false, IsPointer: false
	//
	// ElemType: nimona.Note, ElemKind: struct, ElemTildeKind: Map
	// IsElemSlice: false, IsElemStruct: true, IsElemPointer: true
	{
		sm := []*Note{} // Note
		if vs, err := d.Get("posts"); err == nil {
			if vs, ok := vs.(tilde.List); ok {
				for _, vi := range vs {
					if v, ok := vi.(tilde.Map); ok {
						e := &Note{}
						d := nimona.NewDocument(v)
						e.FromDocument(d)
						sm = append(sm, e)
					}
				}
			}
		}
		if len(sm) > 0 {
			t.Notes = sm
		}
	}

	// # t.Profile
	//
	// Type: nimona.Profile, Kind: struct, TildeKind: Map
	// IsSlice: false, IsStruct: true, IsPointer: false
	{
		if v, err := d.Get("profile"); err == nil {
			if v, ok := v.(tilde.Map); ok {
				e := Profile{}
				d := nimona.NewDocument(v)
				e.FromDocument(d)
				t.Profile = e
			}
		}
	}

	return nil
}

func (t *Follow) Document() *nimona.Document {
	return nimona.NewDocument(t.Map())
}

func (t *Follow) Map() tilde.Map {
	m := tilde.Map{}

	// # t.$type
	//
	// Type: string, Kind: string, TildeKind: InvalidValueKind0
	// IsSlice: false, IsStruct: false, IsPointer: false
	{
		m.Set("$type", tilde.String("follow"))
	}

	// # t.Alias
	//
	// Type: nimona.IdentityAlias, Kind: struct, TildeKind: Map
	// IsSlice: false, IsStruct: true, IsPointer: false
	{
		if !zero.IsZeroVal(t.Alias) {
			m.Set("alias", t.Alias.Map())
		}
	}

	// # t.Identity
	//
	// Type: nimona.Identity, Kind: struct, TildeKind: Map
	// IsSlice: false, IsStruct: true, IsPointer: false
	{
		if !zero.IsZeroVal(t.Identity) {
			m.Set("identity", t.Identity.Map())
		}
	}

	// # t.Metadata
	//
	// Type: nimona.Metadata, Kind: struct, TildeKind: Map
	// IsSlice: false, IsStruct: true, IsPointer: false
	{
		m.Set("$metadata", t.Metadata.Map())
	}

	// # t.Timestamp
	//
	// Type: string, Kind: string, TildeKind: String
	// IsSlice: false, IsStruct: false, IsPointer: false
	{
		if !zero.IsZeroVal(t.Timestamp) {
			m.Set("timestamp", tilde.String(t.Timestamp))
		}
	}

	return m
}

func (t *Follow) FromDocument(d *nimona.Document) error {
	return t.FromMap(d.Map())
}

func (t *Follow) FromMap(d tilde.Map) error {
	*t = Follow{}

	// # t.Alias
	//
	// Type: nimona.IdentityAlias, Kind: struct, TildeKind: Map
	// IsSlice: false, IsStruct: true, IsPointer: false
	{
		if v, err := d.Get("alias"); err == nil {
			if v, ok := v.(tilde.Map); ok {
				e := nimona.IdentityAlias{}
				d := nimona.NewDocument(v)
				e.FromDocument(d)
				t.Alias = e
			}
		}
	}

	// # t.Identity
	//
	// Type: nimona.Identity, Kind: struct, TildeKind: Map
	// IsSlice: false, IsStruct: true, IsPointer: false
	{
		if v, err := d.Get("identity"); err == nil {
			if v, ok := v.(tilde.Map); ok {
				e := nimona.Identity{}
				d := nimona.NewDocument(v)
				e.FromDocument(d)
				t.Identity = e
			}
		}
	}

	// # t.Metadata
	//
	// Type: nimona.Metadata, Kind: struct, TildeKind: Map
	// IsSlice: false, IsStruct: true, IsPointer: false
	{
		if v, err := d.Get("$metadata"); err == nil {
			if v, ok := v.(tilde.Map); ok {
				e := nimona.Metadata{}
				d := nimona.NewDocument(v)
				e.FromDocument(d)
				t.Metadata = e
			}
		}
	}

	// # t.Timestamp
	//
	// Type: string, Kind: string, TildeKind: String
	// IsSlice: false, IsStruct: false, IsPointer: false
	{
		if v, err := d.Get("timestamp"); err == nil {
			if v, ok := v.(tilde.String); ok {
				t.Timestamp = string(v)
			}
		}
	}

	return nil
}

func (t *Note) Document() *nimona.Document {
	return nimona.NewDocument(t.Map())
}

func (t *Note) Map() tilde.Map {
	m := tilde.Map{}

	// # t.$type
	//
	// Type: string, Kind: string, TildeKind: InvalidValueKind0
	// IsSlice: false, IsStruct: false, IsPointer: false
	{
		m.Set("$type", tilde.String("post"))
	}

	// # t.Content
	//
	// Type: string, Kind: string, TildeKind: String
	// IsSlice: false, IsStruct: false, IsPointer: false
	{
		m.Set("content", tilde.String(t.Content))
	}

	// # t.Key
	//
	// Type: string, Kind: string, TildeKind: String
	// IsSlice: false, IsStruct: false, IsPointer: false
	{
		if !zero.IsZeroVal(t.Key) {
			m.Set("_key", tilde.String(t.Key))
		}
	}

	// # t.Metadata
	//
	// Type: nimona.Metadata, Kind: struct, TildeKind: Map
	// IsSlice: false, IsStruct: true, IsPointer: false
	{
		m.Set("$metadata", t.Metadata.Map())
	}

	// # t.Partition
	//
	// Type: string, Kind: string, TildeKind: String
	// IsSlice: false, IsStruct: false, IsPointer: false
	{
		if !zero.IsZeroVal(t.Partition) {
			m.Set("_partition", tilde.String(t.Partition))
		}
	}

	return m
}

func (t *Note) FromDocument(d *nimona.Document) error {
	return t.FromMap(d.Map())
}

func (t *Note) FromMap(d tilde.Map) error {
	*t = Note{}

	// # t.Content
	//
	// Type: string, Kind: string, TildeKind: String
	// IsSlice: false, IsStruct: false, IsPointer: false
	{
		if v, err := d.Get("content"); err == nil {
			if v, ok := v.(tilde.String); ok {
				t.Content = string(v)
			}
		}
	}

	// # t.Key
	//
	// Type: string, Kind: string, TildeKind: String
	// IsSlice: false, IsStruct: false, IsPointer: false
	{
		if v, err := d.Get("_key"); err == nil {
			if v, ok := v.(tilde.String); ok {
				t.Key = string(v)
			}
		}
	}

	// # t.Metadata
	//
	// Type: nimona.Metadata, Kind: struct, TildeKind: Map
	// IsSlice: false, IsStruct: true, IsPointer: false
	{
		if v, err := d.Get("$metadata"); err == nil {
			if v, ok := v.(tilde.Map); ok {
				e := nimona.Metadata{}
				d := nimona.NewDocument(v)
				e.FromDocument(d)
				t.Metadata = e
			}
		}
	}

	// # t.Partition
	//
	// Type: string, Kind: string, TildeKind: String
	// IsSlice: false, IsStruct: false, IsPointer: false
	{
		if v, err := d.Get("_partition"); err == nil {
			if v, ok := v.(tilde.String); ok {
				t.Partition = string(v)
			}
		}
	}

	return nil
}

func (t *Profile) Document() *nimona.Document {
	return nimona.NewDocument(t.Map())
}

func (t *Profile) Map() tilde.Map {
	m := tilde.Map{}

	// # t.$type
	//
	// Type: string, Kind: string, TildeKind: InvalidValueKind0
	// IsSlice: false, IsStruct: false, IsPointer: false
	{
		m.Set("$type", tilde.String("profile"))
	}

	// # t.AvatarURL
	//
	// Type: string, Kind: string, TildeKind: String
	// IsSlice: false, IsStruct: false, IsPointer: false
	{
		if !zero.IsZeroVal(t.AvatarURL) {
			m.Set("avatarURL", tilde.String(t.AvatarURL))
		}
	}

	// # t.Description
	//
	// Type: string, Kind: string, TildeKind: String
	// IsSlice: false, IsStruct: false, IsPointer: false
	{
		if !zero.IsZeroVal(t.Description) {
			m.Set("description", tilde.String(t.Description))
		}
	}

	// # t.DisplayName
	//
	// Type: string, Kind: string, TildeKind: String
	// IsSlice: false, IsStruct: false, IsPointer: false
	{
		if !zero.IsZeroVal(t.DisplayName) {
			m.Set("displayName", tilde.String(t.DisplayName))
		}
	}

	// # t.Metadata
	//
	// Type: nimona.Metadata, Kind: struct, TildeKind: Map
	// IsSlice: false, IsStruct: true, IsPointer: false
	{
		m.Set("$metadata", t.Metadata.Map())
	}

	return m
}

func (t *Profile) FromDocument(d *nimona.Document) error {
	return t.FromMap(d.Map())
}

func (t *Profile) FromMap(d tilde.Map) error {
	*t = Profile{}

	// # t.AvatarURL
	//
	// Type: string, Kind: string, TildeKind: String
	// IsSlice: false, IsStruct: false, IsPointer: false
	{
		if v, err := d.Get("avatarURL"); err == nil {
			if v, ok := v.(tilde.String); ok {
				t.AvatarURL = string(v)
			}
		}
	}

	// # t.Description
	//
	// Type: string, Kind: string, TildeKind: String
	// IsSlice: false, IsStruct: false, IsPointer: false
	{
		if v, err := d.Get("description"); err == nil {
			if v, ok := v.(tilde.String); ok {
				t.Description = string(v)
			}
		}
	}

	// # t.DisplayName
	//
	// Type: string, Kind: string, TildeKind: String
	// IsSlice: false, IsStruct: false, IsPointer: false
	{
		if v, err := d.Get("displayName"); err == nil {
			if v, ok := v.(tilde.String); ok {
				t.DisplayName = string(v)
			}
		}
	}

	// # t.Metadata
	//
	// Type: nimona.Metadata, Kind: struct, TildeKind: Map
	// IsSlice: false, IsStruct: true, IsPointer: false
	{
		if v, err := d.Get("$metadata"); err == nil {
			if v, ok := v.(tilde.Map); ok {
				e := nimona.Metadata{}
				d := nimona.NewDocument(v)
				e.FromDocument(d)
				t.Metadata = e
			}
		}
	}

	return nil
}
