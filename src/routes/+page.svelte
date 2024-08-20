<script lang="ts">
  import CodePlayer from "./CodePlayer.svelte";

  interface CodePlayerProps {
    title: string;
    description: string;
    code: string;
    highlightLines: string;
  }

  let codePlayers: CodePlayerProps[] = [
    {
      title: "slices.All()",
      description: `All returns an iterator over index-value pairs 
in the slice in the usual order.`,
      code: `package main

import (
  "fmt"
  "slices"
)

func main() {
  names := []string{
    "Alice",
    "Bob",
    "Vera",
  }
  for i, v := range slices.All(names) {
    fmt.Println(i, ":", v)
  }
}`,
      highlightLines: "8-15",
    },
    {
      title: "slices.Values()",
      description: `Values returns an iterator that yields the slice 
elements in order.`,
      code: `package main

import (
  "fmt"
  "slices"
)

func main() {
  names := []string{
    "Alice",
    "Bob",
    "Vera",
  }
  for v := range slices.Values(names) {
    fmt.Println(v)
  }
}`,
      highlightLines: "8-11",
    },
    {
      title: "slices.Backward()",
      description: `Backward returns an iterator over index-value pairs 
in the slice, traversing it backward with descending indices.`,
      code: `package main

import (
  "fmt"
  "slices"
)

func main() {
  names := []string{
    "Alice",
    "Bob",
    "Vera",
  }
  for i, v := range slices.Backward(names) {
    fmt.Println(i, ":", v)
  }
}`,
      highlightLines: "8-15",
    },
    {
      title: "slices.Collect()",
      description: `Collect collects values from seq into a new 
slice and returns it.`,
      code: `package main

import (
  "fmt"
  "slices"
)

func main() {
  seq := func(yield func(int) bool) {
    for i := 0; i < 10; i += 2 {
      if !yield(i) {
        return
      }
    }
  }

  s := slices.Collect(seq)
  fmt.Println(s)
}`,
      highlightLines: "8-17",
    },
    {
      title: "slices.AppendSeq()",
      description: `AppendSeq appends the values from seq to the 
slice and returns the extended slice.`,
      code: `package main
import (
  "fmt"
  "slices"
)

func main() {
  seq := func(yield func(int) bool) {
    for i := 0; i < 10; i += 2 {
      if !yield(i) {
        return
      }
    }
  }

  s := slices.AppendSeq([]int{1, 2}, seq)
  fmt.Println(s)
}`,
      highlightLines: "7-16",
    },
    {
      title: "slices.Sorted()",
      description: `Sorted collects values from seq into a new slice, 
sorts the slice, and returns it.`,
      code: `package main

import (
  "fmt"
  "slices"
)

func main() {
  seq := func(yield func(int) bool) {
    flag := -1
    for i := 0; i < 10; i += 2 {
      flag = -flag
      if !yield(i * flag) {
        return
      }
    }
  }

  s := slices.Sorted(seq)
  fmt.Println(s)
  fmt.Println(slices.IsSorted(s))
}`,
      highlightLines: "8-20",
    },
    {
      title: "slices.SortedFunc()",
      description: `SortedFunc collects values from seq into a new slice, 
sorts the slice using the comparison function, and returns it.`,
      code: `package main

import (
  "cmp"
  "fmt"
  "slices"
)

func main() {
  seq := func(yield func(int) bool) {
    flag := -1
    for i := 0; i < 10; i += 2 {
      flag = -flag
      if !yield(i * flag) {
        return
      }
    }
  }

  sortFunc := func(a, b int) int {
    // the comparison
    // is being done in reverse
    return cmp.Compare(b, a)
  }

  s := slices.SortedFunc(seq, sortFunc)
  fmt.Println(s)
}`,
      highlightLines: "9-26",
    },
    {
      title: "slices.SortedStableFunc()",
      description: `SortedStableFunc collects values from seq into 
a new slice. It then sorts the slice while keeping the 
original order of equal elements, using the comparison 
function to compare elements. It returns the new slice.`,
      code: `package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"Gopher", 13},
		{"Alice", 20},
		{"Bob", 5},
		{"Vera", 24},
		{"Zac", 20},
	}

	sortFunc := func(x, y Person) int {
		return cmp.Compare(x.Age, y.Age)
	}

	s := slices.SortedStableFunc(slices.Values(people), sortFunc)
	fmt.Println(s)
}`,
      highlightLines: "9-27",
    },
    {
      title: "slices.Chunk()",
      description: `Chunk returns an iterator over consecutive sub-slices of 
up to n elements of s. All but the last sub-slice will have 
size n. All sub-slices are clipped to have no capacity 
beyond the length. If s is empty, the sequence is empty: 
there is no empty slice in the sequence. Chunk panics if 
n is less than 1.`,
      code: `package main

import (
	"fmt"
	"slices"
)

func main() {
	type Person struct {
		Name string
		Age  int
	}

	type People []Person

	people := People{
		{"Gopher", 13},
		{"Alice", 20},
		{"Bob", 5},
		{"Vera", 24},
		{"Zac", 15},
	}

	// Chunk people into []Person 2 elements at a time.
	for c := range slices.Chunk(people, 2) {
		fmt.Println(c)
	}

}`,
      highlightLines: "8-26",
    },
    {
      title: "maps.All()",
      description: `All returns an iterator over key-value pairs from m. 
The iteration order is not specified and is not guaranteed to 
be the same from one call to the next.`,
      code: `package main

import (
	"fmt"
	"maps"
)

func main() {
	m1 := map[string]int{
		"one": 1,
		"two": 2,
	}
	m2 := map[string]int{
		"one": 10,
	}
	maps.Insert(m2, maps.All(m1))
	fmt.Println("m2 is:", m2)
}`,
      highlightLines: "8-16",
    },
    {
      title: "maps.Keys()",
      description: `Keys returns an iterator over keys in m. 
The iteration order is not specified and is not guaranteed to 
be the same from one call to the next.`,
      code: `package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	m1 := map[int]string{
		1:    "one",
		10:   "Ten",
		1000: "THOUSAND",
	}
	keys := slices.Sorted(maps.Keys(m1))
	fmt.Println(keys)
}
`,
      highlightLines: "9-15",
    },
    {
      title: "maps.Values()",
      description: `Values returns an iterator over values in m. 
The iteration order is not specified and is not guaranteed to 
be the same from one call to the next.`,
      code: `package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	m1 := map[int]string{
		1:    "one",
		10:   "Ten",
		1000: "THOUSAND",
	}
	values := slices.Sorted(maps.Values(m1))
	fmt.Println(values)
}
`,
      highlightLines: "9-15",
    },
    {
      title: "maps.Insert()",
      description: `Insert adds the key-value pairs from seq 
to m. If a key in seq already exists in m, its value will be 
overwritten.`,
      code: `package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	m1 := map[int]string{
		1000: "THOUSAND",
	}
	s1 := []string{"zero", "one", "two", "three"}
	maps.Insert(m1, slices.All(s1))
	fmt.Println("m1 is:", m1)
}
`,
      highlightLines: "9-14",
    },
    {
      title: "maps.Collect()",
      description: `Collect collects key-value pairs 
from seq into a new map and returns it.`,
      code: `package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	s1 := []string{"zero", "one", "two", "three"}
	m1 := maps.Collect(slices.All(s1))
	fmt.Println("m1 is:", m1)
}
`,
      highlightLines: "9-11",
    },
  ];
</script>

<div class="max-w-full sm:max-w-xl h-auto sm:h-100">
  <div class="flex flex-col space-y-4">
    <div>
      <h1
        class="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl"
      >
        <span>
          <img src="gopher.gif" alt="Gopher" />
        </span>
      </h1>
      <h1
        class="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl"
      >
        Go 1.23 Release Notes
      </h1>

      <p class="leading-7 [&:not(:first-child)]:mt-6">
        The latest Go release, version <a
          href="https://tip.golang.org/doc/go1.23"
          class="ext-primary font-medium underline underline-offset-4"
        >
          1.23
        </a>
        , arrives six months after
        <a
          href="https://tip.golang.org/doc/go1.22"
          class="text-primary font-medium underline underline-offset-4"
        >
          Go 1.22
        </a>. Most of its changes are in the implementation of the toolchain,
        runtime, and libraries. As always, the release maintains the Go 1
        <a
          href="https://tip.golang.org/doc/go1compat"
          class="text-primary font-medium underline underline-offset-4"
        >
          promise of compatibility
        </a>. We expect almost all Go programs to continue to compile and run as
        before.
      </p>
      <h2
        class="mt-10 scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0"
      >
        Visual notes for Go 1.23
      </h2>
      <p class="leading-7 [&:not(:first-child)]:mt-6">
        The functions introduced in slices and maps packages are described
        below:
      </p>
    </div>
    {#each codePlayers as codePlayer, index}
      <CodePlayer
        title={codePlayer.title}
        description={codePlayer.description}
        code={codePlayer.code}
        highlightLines={codePlayer.highlightLines}
      />
    {/each}
  </div>
</div>
