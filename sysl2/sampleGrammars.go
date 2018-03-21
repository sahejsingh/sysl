package main

import (
    "fmt"

    "anz-bank/sysl/sysl2/proto"
)

func makeQuantifierOptional() *sysl.Quantifier {
    return &sysl.Quantifier{Union: &sysl.Quantifier_Optional{}}
}

func makeQuantifierZeroPlus() *sysl.Quantifier {
    return &sysl.Quantifier{Union: &sysl.Quantifier_ZeroPlus{}}
}

func makeQuantifierOnePlus() *sysl.Quantifier {
    return &sysl.Quantifier{Union: &sysl.Quantifier_OnePlus{}}
}

func makeTerm(str string) *sysl.Term {
    return &sysl.Term{Atom: &sysl.Atom{Union: &sysl.Atom_Regexp{Regexp: str}}, Quantifier: nil}
}

func makeSequence(terms ...*sysl.Term) *sysl.Sequence {
    seq := sysl.Sequence{Term: terms}
    return &seq
}

func makeRule(name string) (*sysl.RuleName, *sysl.Term) {
    ruleName := sysl.RuleName{Name: name}
    ruleTerm := sysl.Term{Atom: &sysl.Atom{Union: &sysl.Atom_Rulename{Rulename: &ruleName}}, Quantifier: nil}
    return &ruleName, &ruleTerm
}

// S –> bab | bA
// A –> d | cA
func makeGrammar1() *sysl.Grammar {
    a := makeTerm("a")
    b := makeTerm("b")
    c := makeTerm("c")
    d := makeTerm("d")

    ruleNameA, A := makeRule("A")
    ruleNameS, _ := makeRule("S")

    return &sysl.Grammar{
        Name:  "test",
        Start: "S",
        Rules: map[string]*sysl.Rule{
            "S": &sysl.Rule{
                Name: ruleNameS,
                Choices: &sysl.Choice{
                    Sequence: []*sysl.Sequence{
                        makeSequence(b, a, b),
                        makeSequence(b, A)},
                },
            },
            "A": &sysl.Rule{
                Name: ruleNameA,
                Choices: &sysl.Choice{
                    Sequence: []*sysl.Sequence{
                        makeSequence(d),
                        makeSequence(c, A),
                    },
                },
            },
        },
    }
}

func makeEBNF() *sysl.Grammar {
    star := makeTerm("[*]")
    plus := makeTerm("[+]")
    qn := makeTerm("[?]")
    alt := makeTerm("[|]")
    colon := makeTerm("[:]")
    semiColon := makeTerm("[;]")
    openParen := makeTerm("[(]")
    closeParen := makeTerm("[)]")
    STRING := makeTerm(`["][^"]*["]`)

    tokenName := makeTerm("[A-Z][0-9A-Z_]*")
    lowercaseName := makeTerm("[a-z][0-9a-z_]*")

    lhsName, lhsTerm := makeRule("lhs")
    rhsName, rhsTerm := makeRule("rhs")
    ruleName, ruleTerm := makeRule("rule")
    grammarName, _ := makeRule("grammar")
    choiceName, choiceTerm := makeRule("choice")
    seqName, seqTerm := makeRule("seq")
    atomName, atomTerm := makeRule("atom")
    termName, termTerm := makeRule("term")
    termTerm.Quantifier = makeQuantifierOnePlus()
    quantifierName, quantifierTerm := makeRule("quantifier")
    quantifierTerm.Quantifier = makeQuantifierOptional()

    zeroPlusChoiceTerm := sysl.Term{
        Atom: &sysl.Atom{
            Union: &sysl.Atom_Choices{
                Choices: &sysl.Choice{
                    Sequence: []*sysl.Sequence{
                        makeSequence(alt, seqTerm),
                    },
                },
            },
        },
        Quantifier: makeQuantifierZeroPlus(),
    }

    ruleTerm.Quantifier = makeQuantifierOnePlus()

    return &sysl.Grammar{
        Name:  "EBNF",
        Start: "grammar",
        Rules: map[string]*sysl.Rule{
            "grammar": &sysl.Rule{
                Name: grammarName,
                Choices: &sysl.Choice{
                    Sequence: []*sysl.Sequence{
                        makeSequence(ruleTerm),
                    },
                },
            },
            "rule": &sysl.Rule{
                Name: ruleName,
                Choices: &sysl.Choice{
                    Sequence: []*sysl.Sequence{
                        makeSequence(lhsTerm, colon, rhsTerm, semiColon),
                    },
                },
            },
            "lhs": &sysl.Rule{
                Name: lhsName,
                Choices: &sysl.Choice{
                    Sequence: []*sysl.Sequence{
                        makeSequence(lowercaseName),
                    },
                },
            },
            "rhs": &sysl.Rule{
                Name: rhsName,
                Choices: &sysl.Choice{
                    Sequence: []*sysl.Sequence{
                        makeSequence(choiceTerm),
                    },
                },
            },
            "choice": &sysl.Rule{
                Name: choiceName,
                Choices: &sysl.Choice{
                    Sequence: []*sysl.Sequence{
                        makeSequence(seqTerm, &zeroPlusChoiceTerm),
                    },
                },
            },
            "seq": &sysl.Rule{
                Name: seqName,
                Choices: &sysl.Choice{
                    Sequence: []*sysl.Sequence{
                        makeSequence(termTerm),
                    },
                },
            },
            "term": &sysl.Rule{
                Name: termName,
                Choices: &sysl.Choice{
                    Sequence: []*sysl.Sequence{
                        makeSequence(atomTerm, quantifierTerm),
                    },
                },
            },
            "atom": &sysl.Rule{
                Name: atomName,
                Choices: &sysl.Choice{
                    Sequence: []*sysl.Sequence{
                        makeSequence(STRING),
                        makeSequence(tokenName),
                        makeSequence(lowercaseName),
                        makeSequence(openParen, choiceTerm, closeParen),
                    },
                },
            },
            "quantifier": &sysl.Rule{
                Name: quantifierName,
                Choices: &sysl.Choice{
                    Sequence: []*sysl.Sequence{
                        makeSequence(star),
                        makeSequence(plus),
                        makeSequence(qn),
                    },
                },
            },
        }}

}

// E  -> T E'
// E' -> + T E' | -TE' |epsilon
// T  -> F T'
// T' -> * F T' | /FT' |epsilon
// F  -> (E) | int
func makeEXPR() *sysl.Grammar {
    plus := makeTerm("[+]")
    minus := makeTerm("[-]")
    star := makeTerm("[*]")
    divide := makeTerm("[/]")
    openParen := makeTerm("[(]")
    closeParen := makeTerm("[)]")
    integer := makeTerm("[0-9]+")

    ERuleName, ETerm := makeRule("E")
    ETailRuleName, ETailTerm := makeRule("ETail")
    TRuleName, TTerm := makeRule("T")
    TTailRuleName, TTailTerm := makeRule("TTail")
    factorRuleName, factorTerm := makeRule("factor")

    return &sysl.Grammar{
        Name:  "EXPR",
        Start: "E",
        Rules: map[string]*sysl.Rule{
            "E": &sysl.Rule{
                Name: ERuleName,
                Choices: &sysl.Choice{
                    Sequence: []*sysl.Sequence{
                        makeSequence(TTerm, ETailTerm),
                    },
                },
            },
            "ETail": &sysl.Rule{
                Name: ETailRuleName,
                Choices: &sysl.Choice{
                    Sequence: []*sysl.Sequence{
                        makeSequence(plus, TTerm, ETailTerm),
                        makeSequence(minus, TTerm, ETailTerm),
                        makeSequence(nil),
                    },
                },
            },
            "T": &sysl.Rule{
                Name: TRuleName,
                Choices: &sysl.Choice{
                    Sequence: []*sysl.Sequence{
                        makeSequence(factorTerm, TTailTerm),
                    },
                },
            },
            "TTail": &sysl.Rule{
                Name: TTailRuleName,
                Choices: &sysl.Choice{
                    Sequence: []*sysl.Sequence{
                        makeSequence(star, factorTerm, TTailTerm),
                        makeSequence(divide, factorTerm, TTailTerm),
                        makeSequence(nil),
                    },
                },
            },
            "factor": &sysl.Rule{
                Name: factorRuleName,
                Choices: &sysl.Choice{
                    Sequence: []*sysl.Sequence{
                        makeSequence(openParen, ETerm, closeParen),
                        makeSequence(integer),
                    },
                },
            },
        },
    }

}

// obj
//    : '{' number (',' number)* '}'
//    | '{' '}'
//    ;
func makeRepeatSeq(quantifier *sysl.Quantifier) *sysl.Grammar {
    curlyOpen := makeTerm("[{]")
    curlyClosed := makeTerm("[}]")
    comma := makeTerm("[,]")
    number := makeTerm("[0-9]+")

    objRuleName, _ := makeRule("obj")
    obj2RuleName, obj2Term := makeRule("obj2")
    obj2Term.Quantifier = quantifier

    return &sysl.Grammar{
        Name:  "json",
        Start: "obj",
        Rules: map[string]*sysl.Rule{
            "obj": &sysl.Rule{
                Name: objRuleName,
                Choices: &sysl.Choice{
                    Sequence: []*sysl.Sequence{
                        makeSequence(curlyOpen, number, obj2Term, curlyClosed),
                        makeSequence(curlyOpen, curlyClosed),
                    },
                },
            },
            "obj2": &sysl.Rule{
                Name: obj2RuleName,
                Choices: &sysl.Choice{
                    Sequence: []*sysl.Sequence{
                        makeSequence(comma, number),
                    },
                },
            },
        },
    }

}

// json
//    : value
//    ;

// obj
//    : '{' pair (',' pair)* '}'
//    | '{' '}'
//    ;

// pair
//    : STRING ':' value
//    ;

// array
//    : '[' value (',' value)* ']'
//    | '[' ']'
//    ;

// value
//    : STRING
//    | NUMBER
//    | obj
//      | array
func makeJSON(quantifier *sysl.Quantifier) *sysl.Grammar {
    // doubleQuote := makeTerm("\"")
    // singleQuote := makeTerm("'")
    curlyOpen := makeTerm("[{]")
    curlyClosed := makeTerm("[}]")
    comma := makeTerm("[,]")
    sqOpen := makeTerm("[[]")
    sqClose := makeTerm("[]]")
    colon := makeTerm("[:]")
    number := makeTerm("[0-9]+")
    STRING := makeTerm(`["][^"]*["]`)

    jsonRuleName, _ := makeRule("json")
    valueRuleName, valueTerm := makeRule("value")
    objRuleName, objTerm := makeRule("obj")
    pairRuleName, pairTerm := makeRule("pair")
    arrayRuleName, arrayTerm := makeRule("array")

    obj2Term := sysl.Term{
        Atom: &sysl.Atom{
            Union: &sysl.Atom_Choices{
                Choices: &sysl.Choice{
                    Sequence: []*sysl.Sequence{
                        makeSequence(comma, pairTerm),
                    },
                },
            },
        },
        Quantifier: quantifier,
    }

    array2Term := sysl.Term{
        Atom: &sysl.Atom{
            Union: &sysl.Atom_Choices{
                Choices: &sysl.Choice{
                    Sequence: []*sysl.Sequence{
                        makeSequence(comma, valueTerm),
                    },
                },
            },
        },
        Quantifier: quantifier,
    }

    return &sysl.Grammar{
        Name:  "json",
        Start: "json",
        Rules: map[string]*sysl.Rule{
            "obj": &sysl.Rule{
                Name: objRuleName,
                Choices: &sysl.Choice{
                    Sequence: []*sysl.Sequence{
                        makeSequence(curlyOpen, pairTerm, &obj2Term, curlyClosed),
                        makeSequence(curlyOpen, curlyClosed),
                    },
                },
            },
            "value": &sysl.Rule{
                Name: valueRuleName,
                Choices: &sysl.Choice{
                    Sequence: []*sysl.Sequence{
                        makeSequence(STRING),
                        makeSequence(number),
                        makeSequence(objTerm),
                        makeSequence(arrayTerm),
                    },
                },
            },
            "json": &sysl.Rule{
                Name: jsonRuleName,
                Choices: &sysl.Choice{
                    Sequence: []*sysl.Sequence{
                        makeSequence(valueTerm),
                    },
                },
            },
            "pair": &sysl.Rule{
                Name: pairRuleName,
                Choices: &sysl.Choice{
                    Sequence: []*sysl.Sequence{
                        makeSequence(STRING, colon, valueTerm),
                    },
                },
            },
            "array": &sysl.Rule{
                Name: arrayRuleName,
                Choices: &sysl.Choice{
                    Sequence: []*sysl.Sequence{
                        makeSequence(sqOpen, valueTerm, &array2Term, sqClose),
                        makeSequence(sqOpen, sqClose),
                    },
                },
            },
        },
    }

}

func makeG2() *sysl.Grammar {
    a := makeTerm("a")
    b := makeTerm("b")
    d := makeTerm("d")

    SruleName, _ := makeRule("S")
    AruleName, ATerm := makeRule("A")
    BruleName, BTerm := makeRule("B")
    DruleName, DTerm := makeRule("D")

    return &sysl.Grammar{
        Name:  "G2",
        Start: "S",
        Rules: map[string]*sysl.Rule{
            "S": &sysl.Rule{
                Name: SruleName,
                Choices: &sysl.Choice{
                    Sequence: []*sysl.Sequence{
                        makeSequence(ATerm, a),
                    },
                },
            },
            "A": &sysl.Rule{
                Name: AruleName,
                Choices: &sysl.Choice{
                    Sequence: []*sysl.Sequence{
                        makeSequence(BTerm, DTerm),
                    },
                },
            },
            "B": &sysl.Rule{
                Name: BruleName,
                Choices: &sysl.Choice{
                    Sequence: []*sysl.Sequence{
                        makeSequence(b), makeSequence(nil),
                    },
                },
            },
            "D": &sysl.Rule{
                Name: DruleName,
                Choices: &sysl.Choice{
                    Sequence: []*sysl.Sequence{
                        makeSequence(d), makeSequence(nil),
                    },
                },
            },
        },
    }
}

func main() {
    fmt.Println("parsing grammar")

}
