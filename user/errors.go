package user

import "errors"

var (
	ErrOutOfBoundry = errors.New("out of boundry")
	ErrWall = errors.New("there is a wall in the way")
	
	ErrWrongLoc = errors.New("you are not in the required area")
	ErrNoAction = errors.New("no action present")
)