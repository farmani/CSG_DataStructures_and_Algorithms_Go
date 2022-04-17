package _2_WhyAlgorithmsMatter

import "errors"

func linearSearch(array []int, searchValue int) (int, error) {
	// We iterate through every element in the array:
	for index, element := range array {
		// If we find the value we're looking for, we return its index:
		if element == searchValue {
			return index, nil
		} else if element > searchValue {
			// If we reach an element that is greater than the value
			// we're looking for, we can exit the loop early:
			break
		}
	}

	// We return 0 if we do not find the value within the array together with a not nil error message
	return 0, errors.New("not found")
}

func binarySearch(array []int, searchValue int) (int, error) {
	// First, we establish the lower and upper bounds of where the value
	// we're searching for can be. To start, the lower bound is the first
	// value in the array, while the upper bound is the last value:
	lowerBound := 0
	upperBound := len(array) - 1

	// We begin a loop in which we keep inspecting the middlemost value
	// between the upper and lower bounds:
	for lowerBound <= upperBound {
		// We find the midpoint between the upper and lower bounds:
		// (We don't have to worry about the result being a non-integer
		// since in Ruby, the result of division of integers will always
		// be rounded down to the nearest integer.)
		midpoint := (upperBound + lowerBound) / 2
		// We inspect the value at the midpoint:
		valueAtMidpoint := array[midpoint]
		// If the value at the midpoint is the one we're looking for, we're done.
		// If not, we change the lower or upper bound based on whether we need
		// to guess higher or lower:
		if searchValue == valueAtMidpoint {
			return midpoint, nil
		} else if searchValue < valueAtMidpoint {
			upperBound = midpoint - 1
		} else if searchValue > valueAtMidpoint {
			lowerBound = midpoint + 1
		}
	}

	// If we've narrowed the bounds until they've reached each other, that
	// means that the value we're searching for is not contained within
	// this array:
	return 0, errors.New("not found")
}
