package cmd

import (
	"bytes"
	"strings"
	"testing"
)

type byteBufferSeeker struct {
	buf *bytes.Buffer
}

func (b byteBufferSeeker) Read(p []byte) (int, error) {
	return b.buf.Read(p)
}

func (b byteBufferSeeker) Write(p []byte) (int, error) {
	return b.buf.Write(p)
}

func (b byteBufferSeeker) Seek(offset int64, whence int) (int64, error) {
	if offset != 0 && whence != 0 {
		panic("seek only supports 0, 0")
	}
	b.buf.Reset()
	return 0, nil
}

func TestRevertVersionFileContent(t *testing.T) {
	testData := []struct {
		name     string
		content  string
		expected string
	}{
		{
			name: "original",
			content: `package privatedns

import "github.com/Azure/azure-sdk-for-go/version"

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// UserAgent returns the UserAgent string to use when sending http.Requests.
func UserAgent() string {
	return "Azure-SDK-For-Go/" + version.Number + " privatedns/2018-09-01"
}

// Version returns the semantic version (see http://semver.org) of the client.
func Version() string {
	return version.Number
}
`,
			expected: `package privatedns

import "github.com/Azure/azure-sdk-for-go/version"

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// UserAgent returns the UserAgent string to use when sending http.Requests.
func UserAgent() string {
	return "Azure-SDK-For-Go/" + version.Number + " privatedns/2018-09-01"
}

// Version returns the semantic version (see http://semver.org) of the client.
func Version() string {
	return version.Number
}
`,
		},
		{
			name: "raw version.go from autorest",
			content: `package privatedns

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// UserAgent returns the UserAgent string to use when sending http.Requests.
func UserAgent() string {
	return "Azure-SDK-For-Go/" + Version() + " privatedns/2018-09-01"
}

// Version returns the semantic version (see http://semver.org) of the client.
func Version() string {
	return "0.0.0"
}
`,
			expected: `package privatedns

import "github.com/Azure/azure-sdk-for-go/version"
// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// UserAgent returns the UserAgent string to use when sending http.Requests.
func UserAgent() string {
	return "Azure-SDK-For-Go/" + Version() + " privatedns/2018-09-01"
}

// Version returns the semantic version (see http://semver.org) of the client.
func Version() string {
return version.Number
}
`,
		},
		{
			name: "moduled version.go",
			content: `package privatedns

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// UserAgent returns the UserAgent string to use when sending http.Requests.
func UserAgent() string {
	return "Azure-SDK-For-Go/" + Version() + " privatedns/2018-09-01"
}

// Version returns the semantic version (see http://semver.org) of the client.
func Version() string {
	return "v1.2.3"
}
`,
			expected: `package privatedns

import "github.com/Azure/azure-sdk-for-go/version"
// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// UserAgent returns the UserAgent string to use when sending http.Requests.
func UserAgent() string {
	return "Azure-SDK-For-Go/" + Version() + " privatedns/2018-09-01"
}

// Version returns the semantic version (see http://semver.org) of the client.
func Version() string {
return version.Number
}
`,
		},
	}

	repoRoot = "github.com/Azure/azure-sdk-for-go"

	for _, c := range testData {
		t.Logf("Testing %s", c.name)
		buf := byteBufferSeeker{
			buf: bytes.NewBuffer([]byte(c.content)),
		}
		err := revertVersionFileContent(buf)
		if err != nil {
			t.Fatalf("unexpected error: %+v", err)
		}
		if !strings.EqualFold(buf.buf.String(), c.expected) {
			t.Fatalf("expected '%s' but got '%s'", c.expected, buf.buf.String())
		}
	}
}
