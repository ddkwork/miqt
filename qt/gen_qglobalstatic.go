package qt

type QtGlobalStatic__GuardValues int

const (
	QtGlobalStatic__Destroyed     QtGlobalStatic__GuardValues = -2
	QtGlobalStatic__Initialized   QtGlobalStatic__GuardValues = -1
	QtGlobalStatic__Uninitialized QtGlobalStatic__GuardValues = 0
	QtGlobalStatic__Initializing  QtGlobalStatic__GuardValues = 1
)