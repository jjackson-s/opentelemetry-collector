// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package configwiz

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetComponents(t *testing.T) {
	m := map[string]interface{}{
		"pipelines": map[string]interface{}{
			"metrics": rpe{
				Exporters: []string{"ccc", "bbb", "aaa"},
			},
		},
	}
	require.Equal(t, map[string][]string{
		"exporter": {"ccc", "bbb", "aaa"},
	}, serviceToComponentNames(m))
}

func TestCli(t *testing.T) {
	w := fakeWriter{}
	r := fakeReaderPipe{userInput: []string{"", "1", "", "0", "", "", "0", "", "", "0", "", "", "0", ""}}
	io := Clio{w.write, r.read}
	CLI(io, createTestFactories())
	outPut, err := os.ReadFile(defaultFileName)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "", outPut)
	assert.Equal(t, "",  w.programOutput)
	os.Remove(defaultFileName)

}
