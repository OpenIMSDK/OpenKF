// Copyright © 2023 OpenIM open source community. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package urltrie

import (
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

// Hook for interceptor.
type Hook interface {
	// Register with url pattern
	// Support * wildcard, you can use it like this:
	// /api/v1/*, and the url like /api/v1/123 will be matched
	//
	// Now we support multi urls.
	Patterns() []string

	// Priority will set the priority of the hook
	Priority() int64

	// Hooks
	// BeforeRun will be called before controller, you can do something here
	BeforeRun(c *gin.Context)
	// AfterRun will be called after controller, you can do something here
	AfterRun(c *gin.Context)
}

type sortedHook []Hook

func (sh sortedHook) Len() int { return len(sh) }

func (sh sortedHook) Less(i, j int) bool { return sh[i].Priority() > sh[j].Priority() }

func (sh sortedHook) Swap(i, j int) { sh[i], sh[j] = sh[j], sh[i] }

const _wildcard = "*"

type node struct {
	children map[string]*node
	hooks    []Hook
	data     string

	// for wildcard
	isWildcard bool
	isEnd      bool
}

// Trie is a tree for url.
type Trie struct {
	root *node
}

// NewTrie returns a new Trie.
func NewTrie() *Trie {
	return &Trie{
		root: &node{
			children: make(map[string]*node),
		},
	}
}

// InsertBatch insert urls with hooks.
func (t *Trie) InsertBatch(urls []string, hooks ...Hook) {
	for _, url := range urls {
		t.insert(url, hooks...)
	}
}

// Insert insert url with hooks.
func (t *Trie) insert(url string, hooks ...Hook) {
	current := t.root

	// split url with '/'
	parts := strings.Split(url, "/")
	for _, part := range parts {
		if part == "" {
			continue
		}

		child, exists := current.children[part]
		if !exists {
			child = &node{
				children: make(map[string]*node),
				data:     part,
			}
			// match wildcard
			if part == _wildcard {
				child.isWildcard = true
			}

			// set child node
			current.children[part] = child
		}

		current = child
	}

	// Set hooks to last node
	current.isEnd = true
	current.hooks = append(current.hooks, hooks...)
	sortedHooks := sortedHook(current.hooks)
	sort.Sort(sortedHooks)
	current.hooks = sortedHooks
}

// Match match url with hooks.
func (t *Trie) Match(url string) ([]Hook, bool) {
	current := t.root

	parts := strings.Split(url, "/")
	var matchedValues []Hook
	// use stack to save matched children nodes
	stack := make([]*node, 0)
	for _, c := range current.children {
		stack = append(stack, c)
	}

	for _, part := range parts {
		if part == "" {
			continue
		}

		if len(stack) == 0 {
			return matchedValues, len(matchedValues) > 0
		}

		// get current level length
		levelLen := len(stack)
		for i := 0; i < levelLen; i++ {
			// pop
			current, stack = stack[0], stack[1:]
			if current.isEnd {
				if current.isWildcard || current.data == part {
					// Match the last node, append the values
					matchedValues = append(matchedValues, current.hooks...)
				}

				continue
			}

			// find wildcard node or expect node
			if current.isWildcard || current.data == part {
				for _, child := range current.children {
					stack = append(stack, child)
				}
			}
		}
	}

	return matchedValues, len(matchedValues) > 0
}
