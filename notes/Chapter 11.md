# Testing

Go's approach to testing relies on one command - go test - and a set of conventions for writing test functions that go test can run. 

## The go test Tool
3 special kind of functions:
- test functions: a function whose name begins with Test, exercises some program logic for correct behavior
- benchmark function: a function whose name begins with Benchmark, measures the performance of some operation
- example functions: whose name started with Examplel provides machine-checked documentation

The go test tool:
- scans the *_test.go files for these special functions
- generates a temporary main package 
- buils and runs it
- reports the results
- cleans up

## Test Functions
func TestName(t *testing.T) {
    //...
}

- must being with Test
- optional suffic Name must begin with a capital letter
- the t parameter provides methods for reporting test failures and logging additional information

Good practice: write the tests first and observe that it triggers the same failure described by the user's bug report. 

Running go test is usually quicker than manually going through the steps described in the bug report.

Table driven testing is very common in Go. 
```
func TestPalindrome(t *testing.T) {
    var tests = []struct {
        input string
        want bool
    } {
        {"", true},
        {"a", true},
        {"palindrome", false},
        {"desserts", false},
    }
    for _, test := range tests {
        if got := IsPalindrome(test.input); got != test.want {
            t.Errorf("IsPalindrome(%q) = %v", test.input, got)
        }
    }
}
```
Table-driven tests are convienent for checking that a function works on inputs carefully selected to exercise interesting cases in the logic. 

## Randomized Testing
Randomized testing explores a broader range of inputs by constructing inputs at random.
2 strats:
- write an alt implementation of the function that uses a less efficient but simpler and clearer algo and check that both implementations give the same result
- create input values according to a pattern so that we know what output to expect 

Since randomized tests are nondeterministic, it is critical that the log of the failing test reccord sufficient information to reproduce the failure.

## Testing a Command
Can use echo to drive this testing

## White Box Testing
A white-box test has privileged access to the internal functions and data structures of the package and can make observations and changes that an ordinary client cannot.
ex: a white-box test can check that the invariants of the package's data types are maintained after every operation. 

White box tests:
- can provide more detailed coverage of the trickier parts of the implementation

## Writing Effective Tests
- expects test authors to do most of the work themselves - defining functions to avoid repetition
- a good test does not explode on failure but prints a clear description of the symptoms of the problem
- Maintainer should not need to read the source code to decipher a test failure
- a good test should not give up after one failure but should try to report several errors in a single run

Assertion functions suffer from premature abstraction: by treating the failure of this particular tests as a mere difference of two integers, we forfeit the opportunity to provide meaninginful context. 

ex:
```
func TestSplit(t *testing.T) {
    s, sep := "a:b:c", ":"
    words := strings.Split(s, sep)
    if got, want := len(words), 3; got != want {
        t.Errorf("Split(%q, %q) returned %d words, want %d", s, sep, got, want)
    }
    // ...
}

```

## Avoiding Brittle Tests
A test that spuriously fails when a sound change was made to the program is called *brittle*.

The most brittle tests are sometimes called *change detector* or *status quo* test and the time spent dealing with them can quickly deplete any benefit they once seemed to provide. 

Easiest way to avoid brittle tests is to check only properties you care about.
- test your program's simpler and more stable interfaces in preference to its internal functions
- be selective in your assertions
- dont check for exact string matches but look for relevant substrings that will remain unchanged as the program evolves

## Coverage
The degree to which a test suite exercises the package under test is called the test's *coverage*.

*Statement coverage* is the simplest and most widely used of these heuristics. 

The go cover tool measures statement coverage and help identify obvious gaps in the tests.

## Benchmark Functions
Benchmarking, is the practice of measuring the performance of a program on a fixed workload.

ex:
```
import "testing"

func BenchmarkIsPalindrome(b *testing.B) {
    for i := 0; i < b.N; i++ {
        IsPalindrome("A man, a plan, a canal: Panama")
    }
}
```
Can run the above with '$ go test -bench=.'

The resulting report will tell us that each call to IsPalindrome took about 1.035 microseconds, average over 1,000,000 runs.

## Profiling 
Profiling is an automated approach to performance measurement based on sampling a number of profile events during execution, then extrapolating from them during a post-processing step; the resulting statistical summary is called a profile. 

The go test tool has several kinds of profiling:
- CPU profile identifies the functions whose executions requre the most CPU time
- heap profile identifies the operations responsible for blocking goroutines the longest, such as system calls, channel sends and receives, and acquisitions of locks.
