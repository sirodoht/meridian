// Code generated by nimona.io. DO NOT EDIT.

package internal

import (
	"github.com/vikyd/zero"

	"nimona.io"
	"nimona.io/tilde"
)

var _ = zero.IsZeroVal
var _ = tilde.NewScanner

func (t *NimonaFeed) Document() *nimona.Document {
	return nimona.NewDocument(t.Map())
}

func (t *NimonaFeed) Map() tilde.Map {
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
	// Type: []*internal.NimonaFollow, Kind: slice, TildeKind: List
	// IsSlice: true, IsStruct: false, IsPointer: false
	//
	// ElemType: internal.NimonaFollow, ElemKind: struct
	// IsElemSlice: false, IsElemStruct: true, IsElemPointer: true
	{
		if !zero.IsZeroVal(t.Folowees) {
			sm := tilde.List{}
			for i, _ := range t.Folowees {
				v := t.Folowees[i]
				if !zero.IsZeroVal(v) {
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
	// Type: []*internal.NimonaNote, Kind: slice, TildeKind: List
	// IsSlice: true, IsStruct: false, IsPointer: false
	//
	// ElemType: internal.NimonaNote, ElemKind: struct
	// IsElemSlice: false, IsElemStruct: true, IsElemPointer: true
	{
		if !zero.IsZeroVal(t.Notes) {
			sm := tilde.List{}
			for i, _ := range t.Notes {
				v := t.Notes[i]
				if !zero.IsZeroVal(v) {
					sm = append(sm, v.Map())
				}
			}
			m.Set("notes", sm)
		}
	}

	// # t.Profile
	//
	// Type: internal.NimonaProfile, Kind: struct, TildeKind: Map
	// IsSlice: false, IsStruct: true, IsPointer: false
	{
		if !zero.IsZeroVal(t.Profile) {
			m.Set("profile", t.Profile.Map())
		}
	}

	return m
}

func (t *NimonaFeed) FromDocument(d *nimona.Document) error {
	return t.FromMap(d.Map())
}

func (t *NimonaFeed) FromMap(d tilde.Map) error {
	*t = NimonaFeed{}

	// # t.Folowees
	//
	// Type: []*internal.NimonaFollow, Kind: slice, TildeKind: List
	// IsSlice: true, IsStruct: false, IsPointer: false
	//
	// ElemType: internal.NimonaFollow, ElemKind: struct, ElemTildeKind: Map
	// IsElemSlice: false, IsElemStruct: true, IsElemPointer: true
	{
		sm := []*NimonaFollow{} // NimonaFollow
		if vs, err := d.Get("folowees"); err == nil {
			if vs, ok := vs.(tilde.List); ok {
				for _, vi := range vs {
					if v, ok := vi.(tilde.Map); ok {
						e := &NimonaFollow{}
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
	// Type: []*internal.NimonaNote, Kind: slice, TildeKind: List
	// IsSlice: true, IsStruct: false, IsPointer: false
	//
	// ElemType: internal.NimonaNote, ElemKind: struct, ElemTildeKind: Map
	// IsElemSlice: false, IsElemStruct: true, IsElemPointer: true
	{
		sm := []*NimonaNote{} // NimonaNote
		if vs, err := d.Get("notes"); err == nil {
			if vs, ok := vs.(tilde.List); ok {
				for _, vi := range vs {
					if v, ok := vi.(tilde.Map); ok {
						e := &NimonaNote{}
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
	// Type: internal.NimonaProfile, Kind: struct, TildeKind: Map
	// IsSlice: false, IsStruct: true, IsPointer: false
	{
		if v, err := d.Get("profile"); err == nil {
			if v, ok := v.(tilde.Map); ok {
				e := NimonaProfile{}
				d := nimona.NewDocument(v)
				e.FromDocument(d)
				t.Profile = e
			}
		}
	}

	return nil
}
func (t *NimonaProfile) Document() *nimona.Document {
	return nimona.NewDocument(t.Map())
}

func (t *NimonaProfile) Map() tilde.Map {
	m := tilde.Map{}

	// # t.$type
	//
	// Type: string, Kind: string, TildeKind: InvalidValueKind0
	// IsSlice: false, IsStruct: false, IsPointer: false
	{
		m.Set("$type", tilde.String("profile"))
	}

	// # t.Alias
	//
	// Type: nimona.IdentityAlias, Kind: struct, TildeKind: Map
	// IsSlice: false, IsStruct: true, IsPointer: true
	{
		if !zero.IsZeroVal(t.Alias) {
			m.Set("alias", t.Alias.Map())
		}
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

func (t *NimonaProfile) FromDocument(d *nimona.Document) error {
	return t.FromMap(d.Map())
}

func (t *NimonaProfile) FromMap(d tilde.Map) error {
	*t = NimonaProfile{}

	// # t.Alias
	//
	// Type: nimona.IdentityAlias, Kind: struct, TildeKind: Map
	// IsSlice: false, IsStruct: true, IsPointer: true
	{
		if v, err := d.Get("alias"); err == nil {
			if v, ok := v.(tilde.Map); ok {
				e := nimona.IdentityAlias{}
				d := nimona.NewDocument(v)
				e.FromDocument(d)
				t.Alias = &e
			}
		}
	}

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
func (t *NimonaNote) Document() *nimona.Document {
	return nimona.NewDocument(t.Map())
}

func (t *NimonaNote) Map() tilde.Map {
	m := tilde.Map{}

	// # t.$type
	//
	// Type: string, Kind: string, TildeKind: InvalidValueKind0
	// IsSlice: false, IsStruct: false, IsPointer: false
	{
		m.Set("$type", tilde.String("note"))
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

func (t *NimonaNote) FromDocument(d *nimona.Document) error {
	return t.FromMap(d.Map())
}

func (t *NimonaNote) FromMap(d tilde.Map) error {
	*t = NimonaNote{}

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
func (t *NimonaFollow) Document() *nimona.Document {
	return nimona.NewDocument(t.Map())
}

func (t *NimonaFollow) Map() tilde.Map {
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

	// # t.KeygraphID
	//
	// Type: nimona.KeygraphID, Kind: array, TildeKind: Ref
	// IsSlice: false, IsStruct: false, IsPointer: false
	{
		if !zero.IsZeroVal(t.KeygraphID) {
			m.Set("identity", tilde.Ref(t.KeygraphID))
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

func (t *NimonaFollow) FromDocument(d *nimona.Document) error {
	return t.FromMap(d.Map())
}

func (t *NimonaFollow) FromMap(d tilde.Map) error {
	*t = NimonaFollow{}

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

	// # t.KeygraphID
	//
	// Type: nimona.KeygraphID, Kind: array, TildeKind: Ref
	// IsSlice: false, IsStruct: false, IsPointer: false
	{
		if v, err := d.Get("identity"); err == nil {
			if v, ok := v.(tilde.Ref); ok {
				t.KeygraphID = nimona.KeygraphID(v)
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
