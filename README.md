Compute the closure of given inputs: a set functional dependencies list and a set of attributes X.

Given a relation R, its set X of attributes, and a set F of functional dependencies in R, the largest set of attributes that can be derived from X using F, is called the closure of X (denoted X^+)

Inputs: set F of fds
        set X of attributes
Output: closure of X (i.e. X+)

X+ = X
stillChanging = true;
while (stillChanging) {
    stillChanging = false;
    for each W  Z in F {
        if (W is subset of X+) and not (Z is subset of X+) {
            X+ = X+  Z
            stillChanging = true;
        }
    }
}