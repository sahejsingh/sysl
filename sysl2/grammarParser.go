package main

import (
	"sort"

	sysl "anz-bank/sysl/sysl2/proto"
)

type parser struct {
	g         *sysl.Grammar
	l         *lexer
	terminals []string
}

type builder struct {
	arr    []string
	tokens map[string]int32
}

const FIRST = 0
const EPSILON = -2
const EOF = -1

func (b *builder) handleChoice(choice *sysl.Choice) {
	for _, s := range choice.Sequence {
		for _, t := range s.Term {
			if t == nil {
				continue
			}
			var str string
			switch t.Atom.Union.(type) {
			case *sysl.Atom_String_:
				str = t.GetAtom().GetString_()
			case *sysl.Atom_Regexp:
				str = t.GetAtom().GetRegexp()
				str = `\A(?:` + str + `)`
			case *sysl.Atom_Choices:
				b.handleChoice(t.GetAtom().GetChoices())
			}
			if str != "" {
				_, has := b.tokens[str]
				if !has {
					b.tokens[str] = int32(len(b.arr))
					t.Atom.Id = int32(len(b.arr))
					b.arr = append(b.arr, str)
				}
			}
		}
	}
}

// assigns new value to Atom.Id
func (b *builder) buildTerminalsList(rules map[string]*sysl.Rule) []string {
	var ks []string
	for key := range rules {
		ks = append(ks, key)
	}
	sort.Strings(ks)
	for _, key := range ks {
		b.handleChoice(rules[key].Choices)
	}
	return b.arr
}

func makeBuilder() *builder {
	return &builder{
		arr:    make([]string, 0),
		tokens: make(map[string]int32),
	}
}

func makeParser(g *sysl.Grammar, text string) *parser {
	terms := makeBuilder().buildTerminalsList(g.Rules)
	lexer := makeLexer(text, terms)
	return &parser{
		g:         g,
		terminals: terms,
		l:         &lexer,
	}
}

func (p *parser) parse(tokens []int) (bool, []interface{}) {
	return parseGrammar(p.g, tokens, p.g.Start)
}

func setFromTerm(first map[string]*intSet, t *sysl.Term) *intSet {
	var firstSet *intSet
	switch t.Atom.Union.(type) {
	case *sysl.Atom_Choices:
		panic("not handled yet")
	case *sysl.Atom_Rulename:
		nt := t.GetAtom().GetRulename()
		firstSet = first[nt.Name]
	default: //Atom_String_ and Atom_Regexp
		s := make(intSet)
		s.add(int(t.GetAtom().Id))
		firstSet = &s
	}
	return firstSet
}

func hasEpsilon(t *sysl.Term, first map[string]*intSet) bool {
	has := false
	switch t.Atom.Union.(type) {
	// case *sysl.Atom_Choices:
	// firstChoice = gset.firstSetChoice(t.GetAtom().GetChoices())
	case *sysl.Atom_Rulename:
		has = setFromTerm(first, t).has(EPSILON)
	default: //Atom_String_ and Atom_Regexp
		has = false
	}
	return has
}

func isNonTerminal(t *sysl.Term) bool {
	switch t.Atom.Union.(type) {
	case *sysl.Atom_Choices:
		return true
	case *sysl.Atom_Rulename:
		return true
	}
	return false
}

func buildFirstFollowSet(g *sysl.Grammar) (map[string]*intSet, map[string]*intSet) {
	first := make(map[string]*intSet)
	follow := make(map[string]*intSet)

	for key := range g.Rules {
		s1 := make(intSet)
		s2 := make(intSet)
		first[key] = &s1
		follow[key] = &s2
	}

	// Follow Set Rule 1
	// 1. First put $ (the end of input marker) in Follow(S) (S is the start symbol)
	follow[g.Start].add(EOF)

	// check for epsilon
	for ruleName, rule := range g.Rules {
		for _, seq := range rule.Choices.Sequence {
			t := seq.Term[0]
			if t == nil {
				first[ruleName].add(EPSILON)
			}
		}
	}

	updated := true
	for updated {
		updated = false
		for ruleName, rule := range g.Rules {
			for _, seq := range rule.Choices.Sequence {
				numEpsilon := 0
				for _, t := range seq.Term {
					if t == nil {
						continue
					}
					hasEpsilon := hasEpsilon(t, first)
					if hasEpsilon == false {
						updated = first[ruleName].union(setFromTerm(first, t)) || updated
						break
					} else {
						y1 := setFromTerm(first, t).clone()
						y1.remove(EPSILON)
						updated = first[ruleName].union(y1) || updated
						numEpsilon++
					}
					if len(seq.Term) == numEpsilon {
						first[ruleName].add(EPSILON)
					}
				}
			}
		}
	}
	updated = true
	for updated {
		updated = false
		for ruleName, rule := range g.Rules {
			for _, seq := range rule.Choices.Sequence {
				last := len(seq.Term) - 1
				for i := range seq.Term {
					// follow
					if i < last {
						A := seq.Term[i]
						B := seq.Term[i+1]
						Bfirst := setFromTerm(first, B).clone()
						// Rule 4
						// If there is a production X → aAB, where FIRST(B) contains ε,
						// then everything in FOLLOW(X) is in FOLLOW(A)
						if isNonTerminal(B) {
							if Bfirst.has(EPSILON) {
								Afollow := setFromTerm(follow, A)
								updated = Afollow.union(follow[ruleName]) || updated
							}
						}

						// Rule 2.
						// If there is a production X → aAB,
						// (where a can be a whole string)
						// then everything in FIRST(B) except for ε is placed in FOLLOW(A).
						if isNonTerminal(A) {
							Bfirst.remove(EPSILON)
							Afollow := setFromTerm(follow, A)
							updated = Afollow.union(Bfirst) || updated
						}

					} else if i == last {
						// Rule 3
						// If there is a production X → aB, then everything in FOLLOW(X) is in FOLLOW(B)
						B := seq.Term[i]
						if B != nil && isNonTerminal(B) {
							Bfollow := setFromTerm(follow, B)
							updated = Bfollow.union(follow[ruleName]) || updated
						}
					}
				}
			}
		}
	}
	return first, follow
}
