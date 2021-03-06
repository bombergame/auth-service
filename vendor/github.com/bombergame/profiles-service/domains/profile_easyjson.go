// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package domains

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson521a5691DecodeGithubComBombergameProfilesServiceDomains(in *jlexer.Lexer, out *Profiles) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Profiles, 0, 1)
			} else {
				*out = Profiles{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 Profile
			if data := in.Raw(); in.Ok() {
				in.AddError((v1).UnmarshalJSON(data))
			}
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson521a5691EncodeGithubComBombergameProfilesServiceDomains(out *jwriter.Writer, in Profiles) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			out.Raw((v3).MarshalJSON())
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v Profiles) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson521a5691EncodeGithubComBombergameProfilesServiceDomains(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Profiles) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson521a5691EncodeGithubComBombergameProfilesServiceDomains(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Profiles) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson521a5691DecodeGithubComBombergameProfilesServiceDomains(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Profiles) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson521a5691DecodeGithubComBombergameProfilesServiceDomains(l, v)
}
func easyjson521a5691DecodeGithubComBombergameProfilesServiceDomains1(in *jlexer.Lexer, out *Profile) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int64(in.Int64())
		case "username":
			out.Username = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "score":
			out.Score = int32(in.Int32())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson521a5691EncodeGithubComBombergameProfilesServiceDomains1(out *jwriter.Writer, in Profile) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"username\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"score\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Score))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Profile) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson521a5691EncodeGithubComBombergameProfilesServiceDomains1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Profile) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson521a5691EncodeGithubComBombergameProfilesServiceDomains1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Profile) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson521a5691DecodeGithubComBombergameProfilesServiceDomains1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Profile) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson521a5691DecodeGithubComBombergameProfilesServiceDomains1(l, v)
}
