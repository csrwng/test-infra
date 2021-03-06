/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"fmt"
	"golang.org/x/tools/cover"
	"io"
	"k8s.io/test-infra/gopherage/pkg/cov"
	"os"
)

// DumpProfile dumps the profile to the given file destination.
// If the destination is "-", it instead writes to stdout.
func DumpProfile(destination string, profile []*cover.Profile) error {
	var output io.Writer
	if destination == "-" {
		output = os.Stdout
	} else {
		f, err := os.Create(destination)
		if err != nil {
			return fmt.Errorf("failed to open %s: %v", destination, err)
		}
		defer f.Close()
		output = f
	}
	err := cov.DumpProfile(profile, output)
	if err != nil {
		return fmt.Errorf("failed to dump profile: %v", err)
	}
	return nil
}
