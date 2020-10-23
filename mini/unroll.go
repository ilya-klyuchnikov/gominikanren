package mini

import (
	"github.com/awalterschulze/gominikanren/micro"
	"github.com/awalterschulze/gominikanren/sexpr/ast"
)

// defines two candidate functions for unrolling
// in order to check and apply parallel calls

/*
MemberO checks membership in a list

can we make a function that inspects and unrolls this?
i.e. given y = (1 2 3), (membero x y) unrolls to
(conde [(== x 1)][(== x 2)][(== x 3)])
(and of course then we can use a parallel version of conde)
unrolling when x is bound but y is unbound does not make sense
but we can optimise if we detect y is bound

(define (membero x y)
  (fresh (a d)
  (== y `(,a . ,d))
  (conde
    [(== x a)]
    [(membero x d)]
    )))
*/
func MemberO(x, y *ast.SExpr) micro.Goal {
	return micro.Zzz(micro.CallFresh(func(a *ast.SExpr) micro.Goal {
		return micro.CallFresh(func(d *ast.SExpr) micro.Goal {
			return micro.ConjunctionO(
				micro.EqualO(y, ast.Cons(a, d)),
				Conde(
					[]micro.Goal{micro.EqualO(x, a)},
					[]micro.Goal{MemberO(x, d)},
				),
			)
		})
	}))
}

/*
MapO defines a relation between inputs and outputs of a function f
inputs and outputs are given as equal-length lists

can we make a function that inspects and unrolls this?
i.e. given x = (1 2 3), (mapo f x y) unrolls to the following
(fresh (y1 y2 y3) (f 1 y1) (f 2 y2) (f 3 y3) (== x `(,y1 ,y2 ,y3)))
of which each subgoal can be evaluated in parallel
here unrolling makes sense whether x or y is bound
we can detect either and unroll before applying search

(define (mapo f x y)
  (conde
    [(== x '()) (== y '())]
    [(fresh (xa xd ya yd)
        (== `(,xa . ,xd) x)
        (== `(,ya . ,yd) y)
        (f xa ya)
        (mapo f xd yd)
      )]
  ))
*/
func MapO(f func(*ast.SExpr, *ast.SExpr) micro.Goal, x, y *ast.SExpr) micro.Goal {
	return Conde(
		[]micro.Goal{
			micro.EqualO(x, nil),
			micro.EqualO(y, nil),
		},
		[]micro.Goal{
			// really missing Fresh (xa xd ya yd) here
			micro.CallFresh(func(xa *ast.SExpr) micro.Goal {
				return micro.CallFresh(func(xd *ast.SExpr) micro.Goal {
					return micro.CallFresh(func(ya *ast.SExpr) micro.Goal {
						return micro.CallFresh(func(yd *ast.SExpr) micro.Goal {
							return ConjPlus(
								micro.EqualO(x, ast.Cons(xa, xd)),
								micro.EqualO(y, ast.Cons(ya, yd)),
								f(xa, ya),
								MapO(f, xd, yd),
							)
						})
					})
				})
			}),
		},
	)
}
