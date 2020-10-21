# jcheck [![Travis-CI](https://travis-ci.org/scrpi/jcheck.svg)](https://travis-ci.org/scrpi/jcheck) [![GoDoc](https://godoc.org/github.com/scrpi/jcheck?status.svg)](http://godoc.org/github.com/scrpi/jcheck) [![Report card](https://goreportcard.com/badge/github.com/scrpi/jcheck)](https://goreportcard.com/report/github.com/scrpi/jcheck)

Package jcheck provides content validation for json documents. This package is intended to be simpler and lighter to implement than something like json schema, and focuses more on validating content of json objects and values rather than the structure of a json document as a whole.

Each json object or value can be checked by defining a rule. A rule contains a pattern (see syntax below) and one or more `CheckFunc` checks.

Whitelist or blacklist behaviour can be described by passing `DefaultPermitted()` or `DefaultNotPermitted()` options to `NewJSONChecker()`. Rules can then be defined to either permit or not permit json objects based on a given pattern.

JSON strings can be tested for equality or whether they have a suffix or prefix.

JSON numbers have an assortment of equality, greater than or less than operators. There is support for parsing of numbers from strings with suffixes e.g. '200K' => 200,000.

See [Documentation](https://godoc.org/github.com/scrpi)

## Installation

```
go get github.com/scrpi/jcheck
```

## Usage

To construct a new JSONChecker:

```
jc, err := NewJSONChecker()
```

To construct a new JSONChecker with all nodes not permitted by default:

```
jc, err := NewJSONChecker(DefaultNotPermitted())
```

Add rules:

```
jc.AddRule("pattern.field?.*", ArrayLenGTE(4))
jc.AddRule("pattern.field?.*", ArrayLenLTE(10))
jc.AddRule("pattern.object.number", NumGT(0.2))
```

Run check:

```
results, ok := jc.Check()
if (!ok) {
    for _, r := range results {
        fmt.Println(r)
    }
}
```

Each check is executed independently once a pattern matches a node in the json document. A rule fails if any check contained by that rule fails - they are effectively combined as a logical 'OR'. e.g. (if !check1 || !check2 ...)

## Pattern syntax

The pattern syntax is as follows:

```
 pattern:
     { term }

 term:
     '#'     matches any sequence of characters
     '*'     matches any sequence of non-separator ('.') characters
     '?'     matches any single non-separator ('.') character
     c       matches character c

 Note: term must match the entire json node path, not just a substring
```