# Implementation Plan: Alphametics Solver

## Proposal A: Column-based Backtracking with Constraint Propagation

**Role: Proponent**

### Approach

Process the puzzle column-by-column from right to left (like manual addition), assigning digits to letters as they are encountered. At each column, check if the constraint (sum of digits in that column + carry = result digit + 10 * new_carry) can be satisfied. Prune impossible assignments early.

### Files to Modify

- `go/exercises/practice/alphametics/alphametics.go` — sole implementation file

### Architecture

1. **Parse** the puzzle string into addend words and a result word
2. **Build column structure**: organize letters by their column position (units, tens, hundreds, etc.)
3. **Order letters** by column (right to left) for assignment priority
4. **Backtrack** through columns, assigning digits to unassigned letters, checking column constraints as each column becomes fully assigned
5. **Prune** via:
   - Leading-zero constraint: letters that start a word cannot be 0
   - Uniqueness: each digit used at most once
   - Column arithmetic: when all letters in a column are assigned, verify the column sum is consistent

### Rationale

- Column-based processing mirrors how humans solve these puzzles
- Early pruning at each column significantly reduces the search space vs. full permutation enumeration
- For the 199-addend puzzle with 10 letters, this approach can handle it because it prunes deeply at the rightmost columns before exploring further
- Natural handling of carry propagation

### Ordering of Changes

1. Write parser (split on `+` and `==`, extract words, identify unique letters and leading letters)
2. Build column data structure
3. Implement recursive backtracking solver with column constraints
4. Wire up `Solve` function
5. Test

---

## Proposal B: Permutation-based Brute Force with Algebraic Preprocessing

**Role: Opponent**

### Approach

Convert the puzzle into a single linear equation over the letter variables (each letter has a coefficient based on its positional value across all words), then enumerate digit permutations and check whether the equation is satisfied. Optimize by computing coefficients once and evaluating a simple sum.

### Files to Modify

- `go/exercises/practice/alphametics/alphametics.go` — sole implementation file

### Architecture

1. **Parse** the puzzle into addend words and result word
2. **Compute coefficients**: for each letter, compute its total weight. For addend words, each letter at position `i` from the right contributes `+10^i`. For the result word, each letter contributes `-10^i`. A valid solution makes the weighted sum equal to zero.
3. **Enumerate permutations** of digits for the unique letters. Use recursive generation of permutations, checking uniqueness and leading-zero constraints.
4. **Evaluate**: for each permutation, compute the weighted sum. If zero, return the mapping.

### Critique of Proposal A

- Column-based backtracking is more complex to implement correctly, especially handling carries across columns
- The column ordering creates subtle bugs around partially-assigned columns
- More code, more edge cases, harder to debug

### Rationale for Proposal B

- Simpler conceptually: reduce to a single equation, then search
- The coefficient computation is O(total_chars), done once
- Evaluation per permutation is O(num_letters), very fast
- For 10 letters: 10! = 3,628,800 permutations. Each permutation evaluation is ~10 multiplications and additions. On modern hardware this completes in well under a second.
- For fewer letters: search space is even smaller (e.g., 8 letters = 10*9*8*7*6*5*4*3 = 1,814,400)
- Much simpler to implement correctly
- Easy to add leading-zero pruning during permutation generation

### Ordering of Changes

1. Write parser
2. Compute letter coefficients
3. Implement permutation generator with leading-zero constraint
4. Evaluate equation for each permutation
5. Wire up `Solve` function
6. Test

---

## Selected Plan

**Role: Judge**

### Evaluation

| Criterion | Proposal A (Column Backtracking) | Proposal B (Algebraic Permutation) |
|---|---|---|
| **Correctness** | Correct but complex carry logic is error-prone | Correct; single-equation check is straightforward |
| **Risk** | Higher: column indexing, carry propagation, partial-column checks | Lower: simple arithmetic check per permutation |
| **Simplicity** | More complex: ~100+ lines, multiple data structures | Simpler: ~60-80 lines, one coefficient map |
| **Consistency** | Both fit codebase conventions equally | Both fit codebase conventions equally |
| **Performance** | Better asymptotic pruning | 10! worst case but evaluates in <1s on modern hardware |

### Decision: Proposal B wins

**Rationale**: The algebraic permutation approach is significantly simpler to implement correctly. The performance concern about 10! permutations is a non-issue — 3.6M iterations with a ~10-operation inner loop runs in well under a second. The simplicity advantage far outweighs the marginal performance benefit of column-based pruning. For a correct, maintainable solution, Proposal B is clearly superior.

However, I'll incorporate one optimization from Proposal A's spirit: during permutation generation, prune branches early when a leading letter is about to be assigned zero. This is trivially done in the recursive permutation generator.

### Final Implementation Plan

#### File: `go/exercises/practice/alphametics/alphametics.go`

**Package**: `alphametics`

**Imports**: `fmt`, `strings`

**Step 1: Parse the puzzle**

```go
func Solve(puzzle string) (map[string]int, error)
```

- Split `puzzle` on `"=="` to get left-hand side and right-hand side
- Split left-hand side on `"+"` to get addend words
- Trim whitespace from all words
- Collect unique letters (as a slice for stable ordering)
- Identify leading letters (first char of each word with len > 1) — these cannot be zero

**Step 2: Compute coefficients**

- For each addend word, iterate right-to-left, adding `10^position` to that letter's coefficient
- For the result word, iterate right-to-left, subtracting `10^position` from that letter's coefficient
- Store as `map[byte]int` (letter -> coefficient)
- Also create a `letters []byte` slice for iteration order

**Step 3: Recursive permutation search**

```go
func solve(letters []byte, coeffs map[byte]int, leading map[byte]bool, used [10]bool, assignment map[byte]int, idx int) bool
```

- Base case: `idx == len(letters)` → evaluate sum using coefficients; return true if sum == 0
- For each digit 0-9:
  - Skip if digit already used
  - Skip if digit == 0 and current letter is a leading letter
  - Assign digit, mark used, recurse
  - Unassign on backtrack

**Step 4: Build and return result**

- Convert `map[byte]int` to `map[string]int` for the return value
- Return error if no solution found

#### Performance Note

For the 199-addend, 10-letter puzzle: the coefficient precomputation reduces 199 words to 10 coefficients. Each of the 10! = 3,628,800 permutations requires only 10 multiply-add operations to evaluate. This will run in well under 1 second.

#### Additional Optimization

Order the letters so that letters with the largest absolute coefficients are assigned first. This doesn't change correctness but can speed up pruning when combined with a partial-sum check: at each recursive step, if the remaining unassigned letters can't possibly make the partial sum reach zero, prune. However, this optimization is likely unnecessary and adds complexity — implement only if the basic approach is too slow (it won't be).
