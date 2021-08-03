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
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var compNames = []string{"comp1", "comp2", "comp3"}
const compType = "test"

type fakeReaderPipe struct {
	userInput []string
	input     int
}

func (r *fakeReaderPipe) read(defaultVal string) string {
	out := r.userInput[r.input]
	if r.input < len(r.userInput)-1 {
		r.input++
	}
	return out
}

// Testing ComponentNameWizard()
func TestComponentNameWizardEmpty(t *testing.T) {
	w := fakeWriter{}
	r := fakeReader{}
	io := clio{w.write, r.read}
	pr := io.newIndentingPrinter(1)
	out, val := componentNameWizard(io, pr, compType, compNames)
	expected := buildNameWizard(1, "", compType, compNames)
	assert.Equal(t, "", out)
	assert.Equal(t, "", val)
	assert.Equal(t, expected, w.programOutput)
}

func TestComponentNameWizardExtended(t *testing.T) {
	w := fakeWriter{}
	r := fakeReader{"0"}
	io := clio{w.write, r.read}
	pr := io.newIndentingPrinter(1)
	out, val := componentNameWizard(io, pr, compType, compNames)
	expected := buildNameWizard(1, "", compType, compNames)
	tab := strings.Repeat(" ", 4)
	expected += fmt.Sprintf("%s%s %s extended name (optional) > ", tab, out, compType)
	assert.Equal(t, compNames[0], out)
	assert.Equal(t, val, "0")
	assert.Equal(t, expected, w.programOutput)
}

func TestComponentNameWizardError(t *testing.T) {
	w := fakeWriter{}
	r := fakeReaderPipe{[]string{"-1", ""}, 0}
	io := clio{w.write, r.read}
	pr := io.newIndentingPrinter(1)
	out, val := componentNameWizard(io, pr, compType, compNames)
	expected := buildNameWizard(1, "", compType, compNames)
	expected += "Invalid input. Try again.\n"
	expected += buildNameWizard(1, "", compType, compNames)
	assert.Equal(t, "", out)
	assert.Equal(t, "", val)
	assert.Equal(t, expected, w.programOutput)
}

// returns componentNameWizard() output, a list of all components
func buildNameWizard(indent int, prefix string, compType string, compNames []string) string {
	const tabSize = 4
	space := indent * tabSize
	tab := strings.Repeat(" ", space)
	expected := fmt.Sprintf("%sAdd %s (enter to skip)\n", tab, compType)
	for i := 0; i < len(compNames); i++ {
		expected += fmt.Sprintf("%s%d: %s\n", tab, i, compNames[i])
	}
	expected += tab + "> "
	return prefix + expected
}
