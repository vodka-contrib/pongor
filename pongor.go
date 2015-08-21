// Copyright 2015 ipfans
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package pongor

import (
	"io"
	"path"

	"github.com/flosch/pongo2"
)

type PongorOption struct {
	// Directory to load templates. Default is "templates"
	Directory string
}

type Renderer struct {
	Directory string
}

func perparOption(options []PongorOption) PongorOption {
	var opt PongorOption
	if len(options) > 0 {
		opt = options[0]
	}
	if len(opt.Directory) == 0 {
		opt.Directory = "templates"
	}
	return opt
}

func Renderor(opt ...PongorOption) *Renderer {
	o := perparOption(opt)
	r := &Renderer{
		Directory: o.Directory,
	}
	return r
}

func getContext(templateData interface{}) pongo2.Context {
	if templateData == nil {
		return nil
	}
	contextData, isMap := templateData.(map[string]interface{})
	if isMap {
		return contextData
	}
	return nil
}

func (r *Renderer) Render(w io.Writer, name string, data interface{}) error {
	template := pongo2.Must(pongo2.FromFile(path.Join(r.Directory, name)))
	err := template.ExecuteWriter(getContext(data), w)
	return err
}
