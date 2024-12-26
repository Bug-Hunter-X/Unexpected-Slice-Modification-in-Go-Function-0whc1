func myFunc(a []int) {for i := range a {a[i] = i}}

This function modifies the input slice directly, which can lead to unexpected behavior if the caller expects the original slice to remain unchanged.  The issue is subtle because the function doesn't explicitly return a modified slice; it changes the slice in place.