// Encoding: utf-8
// Cloud Foundry Java Buildpack
// Copyright (c) 2015 the original author or authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package memory_test

import (
	"github.com/cloudfoundry/java-buildpack-memory-calculator/memory"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMemorySuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Memory Suite")
}

const (
	bYTE = 1
	kILO = 1024 * bYTE
	mEGA = 1024 * kILO
	gIGA = 1024 * mEGA
)

func getMs(msInt int64) memory.MemSize {
	return memory.NewMemSize(msInt)
}

func boundedMemoryRange(lo, hi int64) memory.Range {
	r, err := memory.NewRange(getMs(lo), getMs(hi))
	Ω(err).ShouldNot(HaveOccurred())
	return r
}

func unboundedMemoryRange(lo int64) memory.Range {
	r, err := memory.NewUnboundedRange(getMs(lo))
	Ω(err).ShouldNot(HaveOccurred())
	return r
}
