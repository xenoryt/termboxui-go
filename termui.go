//termboxui adds extra functionality to the termbox package by nsf.
//This package is intended to be used with termbox and does not completely
//encapsulate and wrap all of its functions.
package termboxui

import "errors"

var ErrEOF = errors.New("EOF")
