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
package ini

import "strings"

type Config struct {
	Sections map[string]*Section
}

func NewConfig() *Config {
	c := new(Config)
	c.Sections = make(map[string]*Section)
	return c
}

func (this *Config) Clear() {
	this.Sections = make(map[string]*Section)
}

func (this *Config) AddComment(section, val string) {
	section = strings.ToLower(section)

	if _, ok := this.Sections[section]; !ok {
		this.Sections[section] = NewSection(section)
	}

	this.Sections[section].AddComment(val)
}

func (this *Config) Set(section, key string, val interface{}) {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if _, ok := this.Sections[section]; !ok {
		this.Sections[section] = NewSection(section)
	}

	this.Sections[section].Set(key, val)
}

func (this *Config) S(section, key, defval string) string {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if s, ok := this.Sections[section]; ok {
		return s.S(key, defval)
	}

	return defval
}

func (this *Config) I(section, key string, defval int) int {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if s, ok := this.Sections[section]; ok {
		return s.I(key, defval)
	}

	return defval
}

func (this *Config) I8(section, key string, defval int8) int8 {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if s, ok := this.Sections[section]; ok {
		return s.I8(key, defval)
	}

	return defval
}

func (this *Config) I16(section, key string, defval int16) int16 {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if s, ok := this.Sections[section]; ok {
		return s.I16(key, defval)
	}

	return defval
}

func (this *Config) I32(section, key string, defval int32) int32 {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if s, ok := this.Sections[section]; ok {
		return s.I32(key, defval)
	}

	return defval
}

func (this *Config) I64(section, key string, defval int64) int64 {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if s, ok := this.Sections[section]; ok {
		return s.I64(key, defval)
	}

	return defval
}

func (this *Config) U(section, key string, defval uint) uint {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if s, ok := this.Sections[section]; ok {
		return s.U(key, defval)
	}

	return defval
}

func (this *Config) U8(section, key string, defval uint8) uint8 {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if s, ok := this.Sections[section]; ok {
		return s.U8(key, defval)
	}

	return defval
}

func (this *Config) U16(section, key string, defval uint16) uint16 {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if s, ok := this.Sections[section]; ok {
		return s.U16(key, defval)
	}

	return defval
}

func (this *Config) U32(section, key string, defval uint32) uint32 {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if s, ok := this.Sections[section]; ok {
		return s.U32(key, defval)
	}

	return defval
}

func (this *Config) U64(section, key string, defval uint64) uint64 {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if s, ok := this.Sections[section]; ok {
		return s.U64(key, defval)
	}

	return defval
}

func (this *Config) F32(section, key string, defval float32) float32 {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if s, ok := this.Sections[section]; ok {
		return s.F32(key, defval)
	}

	return defval
}

func (this *Config) F64(section, key string, defval float64) float64 {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if s, ok := this.Sections[section]; ok {
		return s.F64(key, defval)
	}

	return defval
}

func (this *Config) B(section, key string, defval bool) bool {
	section = strings.ToLower(section)
	key = strings.ToLower(key)

	if s, ok := this.Sections[section]; ok {
		return s.B(key, defval)
	}

	return defval
}
