// Code generated by goderive DO NOT EDIT.

package ast

import (
	"bytes"
	"fmt"
)

// deriveGoString returns a recursive representation of this as a valid go string.
func deriveGoString(this *SExpr) string {
	buf := bytes.NewBuffer(nil)
	fmt.Fprintf(buf, "func() *ast.SExpr {\n")
	if this == nil {
		fmt.Fprintf(buf, "return nil\n")
	} else {
		fmt.Fprintf(buf, "this := &ast.SExpr{}\n")
		if this.List != nil {
			fmt.Fprintf(buf, "this.List = %s\n", deriveGoStringList(this.List))
		}
		if this.Atom != nil {
			fmt.Fprintf(buf, "this.Atom = %s\n", deriveGoStringAtom(this.Atom))
		}
		fmt.Fprintf(buf, "return this\n")
	}
	fmt.Fprintf(buf, "}()\n")
	return buf.String()
}

// deriveGoStringList returns a recursive representation of this as a valid go string.
func deriveGoStringList(this *List) string {
	buf := bytes.NewBuffer(nil)
	fmt.Fprintf(buf, "func() *ast.List {\n")
	if this == nil {
		fmt.Fprintf(buf, "return nil\n")
	} else {
		fmt.Fprintf(buf, "this := &ast.List{}\n")
		fmt.Fprintf(buf, "this.Quoted = %#v\n", this.Quoted)
		if this.Items != nil {
			fmt.Fprintf(buf, "this.Items = %s\n", deriveGoString_(this.Items))
		}
		fmt.Fprintf(buf, "return this\n")
	}
	fmt.Fprintf(buf, "}()\n")
	return buf.String()
}

// deriveGoStringAtom returns a recursive representation of this as a valid go string.
func deriveGoStringAtom(this *Atom) string {
	buf := bytes.NewBuffer(nil)
	fmt.Fprintf(buf, "func() *ast.Atom {\n")
	if this == nil {
		fmt.Fprintf(buf, "return nil\n")
	} else {
		fmt.Fprintf(buf, "this := &ast.Atom{}\n")
		if this.Str != nil {
			fmt.Fprintf(buf, "this.Str = func (v string) *string { return &v }(%#v)\n", *this.Str)
		}
		if this.Symbol != nil {
			fmt.Fprintf(buf, "this.Symbol = func (v string) *string { return &v }(%#v)\n", *this.Symbol)
		}
		if this.Float != nil {
			fmt.Fprintf(buf, "this.Float = func (v float64) *float64 { return &v }(%#v)\n", *this.Float)
		}
		if this.Int != nil {
			fmt.Fprintf(buf, "this.Int = func (v int64) *int64 { return &v }(%#v)\n", *this.Int)
		}
		if this.Var != nil {
			fmt.Fprintf(buf, "this.Var = %s\n", deriveGoStringVar(this.Var))
		}
		fmt.Fprintf(buf, "return this\n")
	}
	fmt.Fprintf(buf, "}()\n")
	return buf.String()
}

// deriveGoStringVar returns a recursive representation of this as a valid go string.
func deriveGoStringVar(this *Variable) string {
	buf := bytes.NewBuffer(nil)
	fmt.Fprintf(buf, "func() *ast.Variable {\n")
	if this == nil {
		fmt.Fprintf(buf, "return nil\n")
	} else {
		fmt.Fprintf(buf, "this := &ast.Variable{}\n")
		fmt.Fprintf(buf, "this.Name = %#v\n", this.Name)
		fmt.Fprintf(buf, "this.ID = %#v\n", this.ID)
		fmt.Fprintf(buf, "return this\n")
	}
	fmt.Fprintf(buf, "}()\n")
	return buf.String()
}

// deriveEqual returns whether this and that are equal.
func deriveEqual(this, that *SExpr) bool {
	return (this == nil && that == nil) ||
		this != nil && that != nil &&
			deriveEqual_(this.List, that.List) &&
			deriveEqual_1(this.Atom, that.Atom)
}

// deriveGoString_ returns a recursive representation of this as a valid go string.
func deriveGoString_(this []*SExpr) string {
	buf := bytes.NewBuffer(nil)
	fmt.Fprintf(buf, "func() []*ast.SExpr {\n")
	if this == nil {
		fmt.Fprintf(buf, "return nil\n")
	} else {
		fmt.Fprintf(buf, "this := make([]*ast.SExpr, %d)\n", len(this))
		for i := range this {
			fmt.Fprintf(buf, "this[%d] = %s\n", i, deriveGoString(this[i]))
		}
		fmt.Fprintf(buf, "return this\n")
	}
	fmt.Fprintf(buf, "}()\n")
	return buf.String()
}

// deriveEqual_ returns whether this and that are equal.
func deriveEqual_(this, that *List) bool {
	return (this == nil && that == nil) ||
		this != nil && that != nil &&
			this.Quoted == that.Quoted &&
			deriveEqual_2(this.Items, that.Items)
}

// deriveEqual_1 returns whether this and that are equal.
func deriveEqual_1(this, that *Atom) bool {
	return (this == nil && that == nil) ||
		this != nil && that != nil &&
			((this.Str == nil && that.Str == nil) || (this.Str != nil && that.Str != nil && *(this.Str) == *(that.Str))) &&
			((this.Symbol == nil && that.Symbol == nil) || (this.Symbol != nil && that.Symbol != nil && *(this.Symbol) == *(that.Symbol))) &&
			((this.Float == nil && that.Float == nil) || (this.Float != nil && that.Float != nil && *(this.Float) == *(that.Float))) &&
			((this.Int == nil && that.Int == nil) || (this.Int != nil && that.Int != nil && *(this.Int) == *(that.Int))) &&
			this.Var.Equal(that.Var)
}

// deriveEqual_2 returns whether this and that are equal.
func deriveEqual_2(this, that []*SExpr) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i].Equal(that[i])) {
			return false
		}
	}
	return true
}
