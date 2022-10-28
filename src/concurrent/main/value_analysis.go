package main

import "unsafe"

type any = interface{}

type Value struct {
	v any
}
type ifaceWords struct {
	typ  unsafe.Pointer
	data unsafe.Pointer
}
