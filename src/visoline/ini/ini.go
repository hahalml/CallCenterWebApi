// Copyright (c) 2009-2010 Jim Teeuwen.
//
// This software is provided 'as-is', without any express or implied
// warranty. In no event will the authors be held liable for any damages
// arising from the use of this software.
//
// Permission is granted to anyone to use this software for any purpose,
// including commercial applications, and to alter it and redistribute it
// freely, subject to the following restrictions:
//
//     1. The origin of this software must not be misrepresented; you must not
//     claim that you wrote the original software. If you use this software
//     in a product, an acknowledgment in the product documentation would be
//     appreciated but is not required.
//
//     2. Altered source versions must be plainly marked as such, and must not be
//     misrepresented as being the original software.
//
//     3. This notice may not be removed or altered from any source distribution.

/*
This package parses run of the mill ini configutation files. It allows loading
and saving. It supports sections and preservation of comments.
*/
package ini

import "io/ioutil"
import "bytes"
import "strings"
import "visoline/mahonia"
import "io"

var strPairSeparator []byte = []byte{'='}
var strNewline []byte = []byte{'\n'}
var strComment []byte = []byte{';'}

func Load(file string, isGBK bool) (cfg *Config, err error) {
	var data []byte
	var pair [][]byte
	var n int

	if data, err = ioutil.ReadFile(file); err != nil {
		return
	}
	//读取GBK类型文件
	if isGBK {
		decode := mahonia.NewDecoder("GBK")
		encode := mahonia.NewEncoder("UTF-8")
		r := decode.NewReader(bytes.NewReader(data))
		buf := bytes.NewBuffer([]byte(""))
		w := encode.NewWriter(buf)
		io.Copy(w, r)
		data = buf.Bytes()
	}
	cfg = &Config{}
	cfg.Sections = make(map[string]*Section)

	section := "_"
	addSection(cfg, section)

	for _, line := range bytes.Split(data, strNewline) {
		if line = bytes.TrimSpace(line); len(line) == 0 {
			continue
		}

		pair = bytes.Split(line, strPairSeparator)

		for i, _ := range pair {
			pair[i] = bytes.TrimSpace(pair[i])
		}

		switch {
		case len(pair) == 1 && len(pair[0]) > 1:
			switch {
			case pair[0][0] == ';' || pair[0][0] == '#':
				cfg.Sections[section].AddComment(string(pair[0][1:]))
			case pair[0][0] == '[' && pair[0][len(pair[0])-1] == ']':
				section = string(pair[0][1 : len(pair[0])-1])
				addSection(cfg, section)
			}
		case len(pair) == 2:
			if n = bytes.Index(pair[1], strComment); n != -1 {
				cfg.Sections[section].AddComment(string(pair[1][n+1:]))
				pair[1] = pair[1][0:n]
			}
			addPair(cfg, section, pair)
		}
	}

	return
}

func Save(file string, isGBK bool, cfg *Config) (err error) {
	var data []byte
	buf := bytes.NewBuffer(data)

	// global section _ goes first
	if s, ok := cfg.Sections["_"]; ok {
		writeSection(buf, s)
	}

	for k, v := range cfg.Sections {
		if k == "_" {
			continue
		}
		writeSection(buf, v)
	}
	result := buf.Bytes()
	if isGBK {
		encode := mahonia.NewEncoder("GBK")
		encodeData := encode.ConvertString(string(result))
		result = []byte(encodeData)
	}
	return ioutil.WriteFile(file, result, 0600)
}

func writeSection(buf *bytes.Buffer, s *Section) {
	if s.Name != "_" {
		buf.WriteString(s.String())
		buf.WriteByte('\n')
	}

	for _, c := range s.Comments {
		buf.WriteString("; ")
		buf.WriteString(c)
		buf.WriteByte('\n')
	}

	for sk, sv := range s.Pairs {
		buf.WriteString(sk)
		buf.WriteString(" = ")
		buf.WriteString(sv)
		buf.WriteByte('\n')
	}
	buf.WriteByte('\n')
}

func addSection(cfg *Config, name string) {
	name = strings.ToLower(name)
	if _, ok := cfg.Sections[name]; !ok {
		cfg.Sections[name] = NewSection(name)
	}
}

func addPair(cfg *Config, name string, pair [][]byte) {
	if pair[0] = bytes.TrimSpace(pair[0]); len(pair[0]) == 0 {
		return
	}
	cfg.Sections[name].Pairs[string(pair[0])] = string(bytes.TrimSpace(pair[1]))
}
