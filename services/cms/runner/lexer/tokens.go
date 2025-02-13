package lexer

import "fmt"

type TokenKind int

const (
	EOF TokenKind = iota
	STRING

	OPEN_BRACKE
	LINK

	HEADING_1
	HEADING_2
	HEADING_3
	HEADING_4
	HEADING_5

	DASH
	NUMBERED_LIST
)

type values []string

func (vals values) getString() string {
	if len(vals) == 1 {
		return vals[0]
	}

	str := ""

	for i, val := range vals {
		str += val
		if i < len(vals)-1 {
			str += " - "
		}
	}

	return str
}

type Location struct {
	start []int
	end   []int
}

func NewLoc(start, end []int) Location {
	return Location{
		start: start,
		end:   end,
	}
}

func (loc Location) Display() string {
	return fmt.Sprintf(
		"    [%d,%d] - [%d,%d]\n",
		loc.start[0],
		loc.start[1],
		loc.end[0],
		loc.end[1],
	)
}

type Token struct {
	Kind   TokenKind
	Values values
	Loc    Location
}

func NewToken(kind TokenKind, loc Location, values ...string) Token {
	return Token{
		Kind:   kind,
		Values: values,
		Loc:    loc,
	}
}

func (token Token) isOneOfKinds(kinds ...TokenKind) bool {
	for _, kind := range kinds {
		if token.Kind == kind {
			return true
		}
	}

	return false
}

func (token Token) Debug() {
	if token.isOneOfKinds(STRING, NUMBERED_LIST, LINK) {
		fmt.Printf(
			"%s (%s)",
			TokenKindString(token.Kind),
			token.Values.getString(),
		)
	} else if token.Kind == EOF {
		fmt.Println(TokenKindString(token.Kind))
		return
	} else {
		fmt.Printf("%s ()", TokenKindString(token.Kind))
	}

	fmt.Printf("%s", token.Loc.Display())
}

func TokenKindString(kind TokenKind) string {
	switch kind {
	case EOF:
		return "eof"
	case STRING:
		return "string"
	case LINK:
		return "link"
	case HEADING_1:
		return "heading_1"
	case HEADING_2:
		return "heading_2"
	case HEADING_3:
		return "heading_3"
	case HEADING_4:
		return "heading_4"
	case HEADING_5:
		return "heading_5"
	case DASH:
		return "dash"
	case NUMBERED_LIST:
		return "numbered_list"
	default:
		return ""
	}
}
