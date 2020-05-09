package uml

import (
	"strings"
	"testing"
)

func TestUML(t *testing.T) {
	tests := []struct {
		name     string
		src      []byte
		wantDist string
		wantErr  bool
	}{
		{
			name: "simple test",
			src: []byte(`
			@startuml
			Alice -> Bob: test
			@enduml
			`),
			wantDist: `<?xml version="1.0" encoding="UTF-8" standalone="no"?><svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" contentScriptType="application/ecmascript" contentStyleType="text/css" height="129px" preserveAspectRatio="none" style="width:120px;height:129px;" version="1.1" viewBox="0 0 120 129" width="120px" zoomAndPan="magnify"><defs><filter height="300%" id="fcxjnt4i49frl" width="300%" x="-1" y="-1"><feGaussianBlur result="blurOut" stdDeviation="2.0"/><feColorMatrix in="blurOut" result="blurOut2" type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 .4 0"/><feOffset dx="4.0" dy="4.0" in="blurOut2" result="blurOut3"/><feBlend in="SourceGraphic" in2="blurOut3" mode="normal"/></filter></defs><g><line style="stroke: #A80036; stroke-width: 1.0; stroke-dasharray: 5.0,5.0;" x1="32" x2="32" y1="39.0146" y2="88.814"/><line style="stroke: #A80036; stroke-width: 1.0; stroke-dasharray: 5.0,5.0;" x1="90" x2="90" y1="39.0146" y2="88.814"/><rect fill="#FEFECE" filter="url(#fcxjnt4i49frl)" height="31.0146" style="stroke: #A80036; stroke-width: 1.5;" width="45" x="8" y="3"/><text fill="#000000" font-family="sans-serif" font-size="14" lengthAdjust="spacingAndGlyphs" textLength="31" x="15" y="24.0752">Alice</text><rect fill="#FEFECE" filter="url(#fcxjnt4i49frl)" height="31.0146" style="stroke: #A80036; stroke-width: 1.5;" width="45" x="8" y="87.814"/><text fill="#000000" font-family="sans-serif" font-size="14" lengthAdjust="spacingAndGlyphs" textLength="31" x="15" y="108.8892">Alice</text><rect fill="#FEFECE" filter="url(#fcxjnt4i49frl)" height="31.0146" style="stroke: #A80036; stroke-width: 1.5;" width="42" x="67" y="3"/><text fill="#000000" font-family="sans-serif" font-size="14" lengthAdjust="spacingAndGlyphs" textLength="28" x="74" y="24.0752">Bob</text><rect fill="#FEFECE" filter="url(#fcxjnt4i49frl)" height="31.0146" style="stroke: #A80036; stroke-width: 1.5;" width="42" x="67" y="87.814"/><text fill="#000000" font-family="sans-serif" font-size="14" lengthAdjust="spacingAndGlyphs" textLength="28" x="74" y="108.8892">Bob</text><polygon fill="#A80036" points="78,66.814,88,70.814,78,74.814,82,70.814" style="stroke: #A80036; stroke-width: 1.0;"/><line style="stroke: #A80036; stroke-width: 1.0;" x1="32.5" x2="84" y1="70.814" y2="70.814"/><text fill="#000000" font-family="sans-serif" font-size="13" lengthAdjust="spacingAndGlyphs" textLength="27" x="39.5" y="66.0845">test</text>`,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDist, err := UML(tt.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("UML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && strings.HasPrefix(tt.wantDist, string(gotDist)) {
				t.Errorf("UML() = %s, want %v", gotDist, tt.wantDist)
			}
		})
	}
}
