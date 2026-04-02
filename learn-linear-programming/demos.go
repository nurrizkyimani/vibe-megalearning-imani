package main

import (
	"fmt"
	"math"
	"time"
)

// =============================================================================
// Num01LinearEquationsSystemsDemo
//
// TOPIC: Linear Equations & Systems
//   - Standard linear equation: a₁x₁ + a₂x₂ + ... + aₙxₙ = b
//   - Matrix form: Ax = b  (A ∈ ℝᵐˣⁿ, x ∈ ℝⁿ, b ∈ ℝᵐ)
//   - Gaussian elimination: O(n³) row reduction on augmented matrix [A | b]
//   - Existence: rank(A) == rank([A|b])  →  at least one solution
//   - Uniqueness: rank(A) == n           →  exactly one solution
//   - Underdetermined (m < n): infinite solutions — LP feasible region is born here
//
// Real world: supply-demand equilibrium, circuit analysis (KVL/KCL),
//
//	production planning, pricing model calibration
//
// =============================================================================
func Num01LinearEquationsSystemsDemo() {
	fmt.Println("============================================================")
	fmt.Println("  Num01 -- Linear Equations & Systems")
	fmt.Println("============================================================")
	fmt.Println()

	// ── Helper: print an augmented matrix [A | b] ────────────────────────────
	printAugmented := func(mat [][]float64, label string) {
		fmt.Printf("  %s\n", label)
		n := len(mat)
		for i := 0; i < n; i++ {
			fmt.Print("  [ ")
			cols := len(mat[i])
			for j := 0; j < cols-1; j++ {
				if j == cols-2 {
					fmt.Printf("%7.3f  |  %7.3f", mat[i][j], mat[i][cols-1])
				} else {
					fmt.Printf("%7.3f  ", mat[i][j])
				}
			}
			fmt.Println(" ]")
		}
		fmt.Println()
	}

	// ── Helper: Gaussian elimination with partial pivoting ───────────────────
	// Returns the row-echelon augmented matrix and the pivot column indices.
	// Operates on a deep copy so the original is unchanged.
	gaussianElim := func(mat [][]float64) ([][]float64, int) {
		m := len(mat)
		n := len(mat[0])
		// deep copy
		a := make([][]float64, m)
		for i := range mat {
			a[i] = make([]float64, n)
			copy(a[i], mat[i])
		}

		pivotRow := 0
		rank := 0
		for col := 0; col < n-1 && pivotRow < m; col++ {
			// find row with largest absolute value in this column (partial pivot)
			maxVal := math.Abs(a[pivotRow][col])
			maxRow := pivotRow
			for row := pivotRow + 1; row < m; row++ {
				if math.Abs(a[row][col]) > maxVal {
					maxVal = math.Abs(a[row][col])
					maxRow = row
				}
			}
			if maxVal < 1e-10 {
				continue // skip — no pivot in this column
			}
			// swap rows
			a[pivotRow], a[maxRow] = a[maxRow], a[pivotRow]
			// scale pivot row so pivot element = 1
			scale := a[pivotRow][col]
			for j := 0; j < n; j++ {
				a[pivotRow][j] /= scale
			}
			// eliminate all other rows
			for row := 0; row < m; row++ {
				if row == pivotRow {
					continue
				}
				factor := a[row][col]
				for j := 0; j < n; j++ {
					a[row][j] -= factor * a[pivotRow][j]
				}
			}
			rank++
			pivotRow++
		}
		return a, rank
	}

	// ── helper: check consistency (rank(A) == rank([A|b])) ───────────────────
	// After full row reduction, an inconsistent row looks like [0 0 ... 0 | c≠0]
	isConsistent := func(reduced [][]float64) bool {
		m := len(reduced)
		n := len(reduced[0])
		for i := 0; i < m; i++ {
			allZero := true
			for j := 0; j < n-1; j++ {
				if math.Abs(reduced[i][j]) > 1e-10 {
					allZero = false
					break
				}
			}
			if allZero && math.Abs(reduced[i][n-1]) > 1e-10 {
				return false // 0 = c  where c ≠ 0 → inconsistent
			}
		}
		return true
	}

	time.Sleep(300 * time.Millisecond)

	// =========================================================================
	// SECTION 1 — Matrix notation refresher
	// =========================================================================
	fmt.Println("── Section 1: Matrix Form of a Linear System ───────────────")
	fmt.Println()
	fmt.Println("  A linear system of m equations in n unknowns:")
	fmt.Println()
	fmt.Println("     a₁₁x₁ + a₁₂x₂ + ... + a₁ₙxₙ = b₁")
	fmt.Println("     a₂₁x₁ + a₂₂x₂ + ... + a₂ₙxₙ = b₂")
	fmt.Println("         ⋮")
	fmt.Println("     aₘ₁x₁ + aₘ₂x₂ + ... + aₘₙxₙ = bₘ")
	fmt.Println()
	fmt.Println("  Compact matrix form:   A x = b")
	fmt.Println()
	fmt.Println("     A ∈ ℝᵐˣⁿ   (coefficient matrix)")
	fmt.Println("     x ∈ ℝⁿ     (unknown vector)")
	fmt.Println("     b ∈ ℝᵐ     (right-hand side / constants)")
	fmt.Println()
	fmt.Println("  Augmented matrix [A | b] encodes everything needed for")
	fmt.Println("  Gaussian elimination.")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	// =========================================================================
	// SECTION 2 — Unique solution (3×3, full rank)
	// =========================================================================
	fmt.Println("── Section 2: Case 1 — Unique Solution (rank = n = 3) ──────")
	fmt.Println()
	fmt.Println("  Problem: 3 equations, 3 unknowns")
	fmt.Println()
	fmt.Println("     2x₁ +  1x₂ -  1x₃ =  8   (row 1)")
	fmt.Println("    -3x₁ + -1x₂ +  2x₃ = -11   (row 2)")
	fmt.Println("    -2x₁ +  1x₂ +  2x₃ = -3    (row 3)")
	fmt.Println()

	// Augmented matrix [A | b]  →  known solution: x₁=2, x₂=3, x₃=-1
	sys1 := [][]float64{
		{2, 1, -1, 8},
		{-3, -1, 2, -11},
		{-2, 1, 2, -3},
	}
	printAugmented(sys1, "Augmented matrix [A | b]:")

	fmt.Println("  Applying Gaussian elimination (partial pivoting)...")
	fmt.Println()
	time.Sleep(400 * time.Millisecond)

	// Step-by-step manual narration  (we run the actual code alongside)
	fmt.Println("  Step 1: Pivot on column 1")
	fmt.Println("          Largest |value| in col 1 is |-3| in row 2 → swap R1 ↔ R2")
	fmt.Println("          Scale R1 so pivot = 1:  R1 ← R1 / -3")
	fmt.Println("          Eliminate col 1 from R2 and R3")
	fmt.Println()
	time.Sleep(300 * time.Millisecond)

	fmt.Println("  Step 2: Pivot on column 2")
	fmt.Println("          Use the current leading row for column 2")
	fmt.Println("          Scale and eliminate col 2 from all other rows")
	fmt.Println()
	time.Sleep(300 * time.Millisecond)

	fmt.Println("  Step 3: Pivot on column 3")
	fmt.Println("          Scale and eliminate col 3 from all other rows")
	fmt.Println()
	time.Sleep(300 * time.Millisecond)

	reduced1, rank1 := gaussianElim(sys1)
	printAugmented(reduced1, "Row-reduced echelon form [I | x*]:")

	fmt.Printf("  rank(A) = %d   n = 3   →   unique solution\n", rank1)
	fmt.Println()
	fmt.Println("  Solution (read directly from last column):")
	fmt.Printf("    x₁ = %.3f\n", reduced1[0][3])
	fmt.Printf("    x₂ = %.3f\n", reduced1[1][3])
	fmt.Printf("    x₃ = %.3f\n", reduced1[2][3])
	fmt.Println()
	fmt.Println("  Verification:  Ax = b ?")
	x1 := []float64{reduced1[0][3], reduced1[1][3], reduced1[2][3]}
	A1 := [][]float64{{2, 1, -1}, {-3, -1, 2}, {-2, 1, 2}}
	b1 := []float64{8, -11, -3}
	allOK := true
	for i := 0; i < 3; i++ {
		sum := 0.0
		for j := 0; j < 3; j++ {
			sum += A1[i][j] * x1[j]
		}
		ok := math.Abs(sum-b1[i]) < 1e-9
		mark := "✓"
		if !ok {
			mark = "✗"
			allOK = false
		}
		fmt.Printf("    Row %d: %.3f = %.3f  %s\n", i+1, sum, b1[i], mark)
	}
	if allOK {
		fmt.Println("  All rows verified ✓")
	}
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	// =========================================================================
	// SECTION 3 — No solution (inconsistent)
	// =========================================================================
	fmt.Println("── Section 3: Case 2 — No Solution (Inconsistent System) ───")
	fmt.Println()
	fmt.Println("  Problem: 3 equations that are contradictory")
	fmt.Println()
	fmt.Println("      x₁ +   x₂ = 1   (row 1)")
	fmt.Println("     2x₁ +  2x₂ = 3   (row 2)  ← contradicts row 1 scaled by 2")
	fmt.Println("      x₁ +   x₂ = 5   (row 3)  ← contradicts row 1")
	fmt.Println()

	sys2 := [][]float64{
		{1, 1, 1},
		{2, 2, 3},
		{1, 1, 5},
	}
	printAugmented(sys2, "Augmented matrix [A | b]:")

	reduced2, rank2 := gaussianElim(sys2)
	printAugmented(reduced2, "After Gaussian elimination:")

	consistent2 := isConsistent(reduced2)
	fmt.Printf("  rank(A) = %d   Consistent = %v\n", rank2, consistent2)
	fmt.Println()
	fmt.Println("  Interpretation:")
	fmt.Println("  Row 2 reduced to [ 0  0 | 1 ]  →  0 = 1  — contradiction!")
	fmt.Println("  The system has NO solution (empty feasible set).")
	fmt.Println()
	fmt.Println("  In LP terms: an infeasible LP has an empty feasible polyhedron.")
	fmt.Println("  The constraints cannot all be satisfied simultaneously.")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	// =========================================================================
	// SECTION 4 — Infinite solutions (underdetermined)
	// =========================================================================
	fmt.Println("── Section 4: Case 3 — Infinite Solutions (rank < n) ───────")
	fmt.Println()
	fmt.Println("  Problem: 2 equations, 3 unknowns  (m=2 < n=3)")
	fmt.Println()
	fmt.Println("      x₁ + 2x₂ +  x₃ = 4   (row 1)")
	fmt.Println("     2x₁ + 4x₂ + 3x₃ = 9   (row 2)")
	fmt.Println()

	sys3 := [][]float64{
		{1, 2, 1, 4},
		{2, 4, 3, 9},
	}
	printAugmented(sys3, "Augmented matrix [A | b]:")

	reduced3, rank3 := gaussianElim(sys3)
	printAugmented(reduced3, "After Gaussian elimination:")

	fmt.Printf("  rank(A) = %d   n = 3   free variables = %d\n", rank3, 3-rank3)
	fmt.Println()
	fmt.Println("  Reduced form gives:")
	fmt.Println("    Row 1: x₁ + 2x₂       = 4 − x₃")
	fmt.Println("    Row 2:            x₃   = 1")
	fmt.Println()
	fmt.Println("  Pivot variables: x₁, x₃")
	fmt.Println("  Free variable:   x₂ = t  (any real number)")
	fmt.Println()
	fmt.Println("  Parametric solution (t ∈ ℝ):")
	fmt.Println("    x₁ = 3 − 2t")
	fmt.Println("    x₂ = t")
	fmt.Println("    x₃ = 1")
	fmt.Println()
	fmt.Println("  This is a 1-dimensional affine subspace (a LINE) in ℝ³.")
	fmt.Println("  Every point on that line is a valid solution.")
	fmt.Println()
	fmt.Println("  ┌─ LP Connection ──────────────────────────────────────────┐")
	fmt.Println("  │  When Ax = b is underdetermined, the feasible set is not │")
	fmt.Println("  │  a single point but an affine subspace.  Adding x ≥ 0    │")
	fmt.Println("  │  clips it to a polytope — the LP feasible region.        │")
	fmt.Println("  └──────────────────────────────────────────────────────────┘")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	// =========================================================================
	// SECTION 5 — Rank & existence/uniqueness theorem
	// =========================================================================
	fmt.Println("── Section 5: Rank & the Existence–Uniqueness Theorem ───────")
	fmt.Println()
	fmt.Println("  Given Ax = b  with A ∈ ℝᵐˣⁿ:")
	fmt.Println()
	fmt.Println("  ┌──────────────────────────────┬───────────────────────────┐")
	fmt.Println("  │  Condition                   │  Outcome                  │")
	fmt.Println("  ├──────────────────────────────┼───────────────────────────┤")
	fmt.Println("  │  rank(A) ≠ rank([A|b])       │  No solution (infeasible) │")
	fmt.Println("  │  rank(A) = rank([A|b]) = n   │  Unique solution          │")
	fmt.Println("  │  rank(A) = rank([A|b]) < n   │  Infinite solutions       │")
	fmt.Println("  └──────────────────────────────┴───────────────────────────┘")
	fmt.Println()
	fmt.Println("  Complexity of Gaussian elimination: O(n²m) ≈ O(n³) for square")
	fmt.Println()
	fmt.Println("  Numerical stability notes:")
	fmt.Println("    • Partial pivoting: swap rows to put largest |aᵢⱼ| in pivot")
	fmt.Println("      position — reduces floating-point error amplification")
	fmt.Println("    • Complete pivoting: also swap columns (more stable, slower)")
	fmt.Println("    • Ill-conditioned systems: condition number κ(A) >> 1")
	fmt.Println("      means small changes in b cause large changes in x")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	// =========================================================================
	// SECTION 6 — Real-world example: production planning
	// =========================================================================
	fmt.Println("── Section 6: Real-World Example — Production Planning ──────")
	fmt.Println()
	fmt.Println("  A factory produces 3 products (P1, P2, P3).")
	fmt.Println("  It uses 3 machines (M1, M2, M3) at full capacity each day.")
	fmt.Println()
	fmt.Println("  Machine hours consumed per unit:")
	fmt.Println("                P1    P2    P3    capacity")
	fmt.Println("    Machine M1:  2     1     3  =   60  hours/day")
	fmt.Println("    Machine M2:  1     3     2  =   70  hours/day")
	fmt.Println("    Machine M3:  3     2     1  =   50  hours/day")
	fmt.Println()
	fmt.Println("  Question: how many units of each product to make each day?")
	fmt.Println()
	fmt.Println("  System:   Ax = b")
	fmt.Println("            2x₁ +  x₂ + 3x₃ = 60")
	fmt.Println("             x₁ + 3x₂ + 2x₃ = 70")
	fmt.Println("            3x₁ + 2x₂ +  x₃ = 50")
	fmt.Println()

	sysProd := [][]float64{
		{2, 1, 3, 60},
		{1, 3, 2, 70},
		{3, 2, 1, 50},
	}
	printAugmented(sysProd, "Augmented matrix [A | b]:")

	reducedProd, rankProd := gaussianElim(sysProd)
	printAugmented(reducedProd, "After Gaussian elimination:")

	xProd := []float64{reducedProd[0][3], reducedProd[1][3], reducedProd[2][3]}
	fmt.Printf("  rank = %d  →  unique production plan:\n", rankProd)
	fmt.Printf("    P1: %.2f units/day\n", xProd[0])
	fmt.Printf("    P2: %.2f units/day\n", xProd[1])
	fmt.Printf("    P3: %.2f units/day\n", xProd[2])
	fmt.Println()

	// Verify
	AProd := [][]float64{{2, 1, 3}, {1, 3, 2}, {3, 2, 1}}
	bProd := []float64{60, 70, 50}
	machines := []string{"M1", "M2", "M3"}
	fmt.Println("  Verification — all machine capacities met?")
	for i := 0; i < 3; i++ {
		used := 0.0
		for j := 0; j < 3; j++ {
			used += AProd[i][j] * xProd[j]
		}
		ok := math.Abs(used-bProd[i]) < 1e-6
		mark := "✓"
		if !ok {
			mark = "✗"
		}
		fmt.Printf("    %s: %.2f / %.2f hours  %s\n", machines[i], used, bProd[i], mark)
	}
	fmt.Println()
	fmt.Println("  Key insight: this is exactly Ax = b at equality (binding")
	fmt.Println("  constraints). In LP, we'll relax = to ≤ and ask which")
	fmt.Println("  solution MAXIMISES profit — that's where we're headed.")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	// =========================================================================
	// SECTION 7 — Summary
	// =========================================================================
	fmt.Println("============================================================")
	fmt.Println("  Key Insights:")
	fmt.Println("  1. Every LP starts as Ax = b — understanding Gaussian")
	fmt.Println("     elimination is the mechanical foundation of LP solving.")
	fmt.Println("  2. Three solution types map directly to LP outcomes:")
	fmt.Println("       rank(A) < rank([A|b]) → infeasible LP")
	fmt.Println("       rank = n              → unique BFS (basic feasible solution)")
	fmt.Println("       rank < n              → infinite solutions → LP feasible region")
	fmt.Println("  3. Partial pivoting keeps elimination numerically stable —")
	fmt.Println("     real solvers (LAPACK / Gurobi) use this internally.")
	fmt.Println("  4. O(n³) cost of Gaussian elimination motivates sparse matrix")
	fmt.Println("     techniques used by every production LP solver.")
	fmt.Println("  5. The production-planning example is a toy LP — Topic 04")
	fmt.Println("     (Simplex) will show how to add an objective and optimise.")
	fmt.Println("============================================================")
	fmt.Println()
}

// =============================================================================
// Num02GeometricInterpretationDemo
//
// TOPIC: Geometric Interpretation of LP
//   - Each linear constraint  aᵢᵀx ≤ bᵢ  defines a half-space in ℝⁿ
//   - Feasible region = intersection of all half-spaces (a convex polyhedron)
//   - Objective function cᵀx defines parallel hyperplanes; optimum at a vertex
//   - Extreme points (vertices) = basic feasible solutions (BFS)
//   - In 2D: visualising constraint lines, feasible polygon, corner enumeration
//   - Unbounded, infeasible, and degenerate cases — ASCII diagrams
//
// Real world: 2D diet / mix problems, Pareto frontier in multi-objective opt,
//
//	portfolio efficient frontier (mean–variance as LP)
//
// =============================================================================
func Num02GeometricInterpretationDemo() {
	// TODO: implement
}

// =============================================================================
// Num03StandardCanonicalFormsDemo
//
// TOPIC: Standard & Canonical Forms of LP
//   - Inequality form (canonical): min cᵀx  s.t. Ax ≤ b, x ≥ 0
//   - Standard form (equality):   min cᵀx  s.t. Ax = b, x ≥ 0
//   - Conversion rules:
//     Ax ≤ b  →  Ax + s = b,  s ≥ 0   (slack variable)
//     Ax ≥ b  →  Ax - e = b,  e ≥ 0   (excess / surplus variable)
//     free xᵢ →  xᵢ = xᵢ⁺ - xᵢ⁻,  xᵢ⁺,xᵢ⁻ ≥ 0
//     max cᵀx →  min -cᵀx
//   - Augmented matrix after conversion: [A | I | b]
//
// Real world: all LP solvers (CPLEX, Gurobi, GLPK) require standard form internally;
//
//	understanding this is prerequisite for implementing or debugging solvers
//
// =============================================================================
func Num03StandardCanonicalFormsDemo() {
	// TODO: implement
}

// =============================================================================
// Num04SimplexMethodTableauDemo
//
// TOPIC: Simplex Method — Tableau & Pivoting
//   - Basic feasible solution (BFS): n vars, m equalities → m basic, n-m non-basic
//   - Simplex tableau: augmented matrix with objective row
//   - Entering variable: most negative reduced cost in objective row
//   - Leaving variable: minimum ratio test  min{ bᵢ/aᵢⱼ : aᵢⱼ > 0 }
//   - Pivot operation: row reduction to make entering column a unit vector
//   - Termination: all reduced costs ≥ 0  →  optimal BFS found
//   - Degeneracy, cycling, Bland's rule
//   - Complexity: exponential worst case (Klee–Minty), polynomial in practice
//
// Real world: every LP solver core; MIP branching nodes; resource allocation engines
//
//	at Uber (driver dispatch), airlines (crew scheduling)
//
// =============================================================================
func Num04SimplexMethodTableauDemo() {
	// TODO: implement
}

// =============================================================================
// Num05DualityTheoryDemo
//
// TOPIC: Duality Theory
//   - Primal:  min cᵀx  s.t. Ax ≥ b, x ≥ 0
//   - Dual:    max bᵀy  s.t. Aᵀy ≤ c, y ≥ 0
//   - Weak duality:   cᵀx ≥ bᵀy  for any primal/dual feasible (x, y)
//   - Strong duality: optimal primal value == optimal dual value (Dantzig 1951)
//   - Complementary slackness:  xᵢ(cᵢ - Aᵀy)ᵢ = 0,  yⱼ(Axⱼ - bⱼ) = 0
//   - Economic interpretation: dual variables = shadow prices / marginal values
//   - Certificate of optimality: primal + dual feasible + equal objectives
//
// Real world: shadow prices in resource allocation; SVM dual = kernel trick entry point;
//
//	Jane Street uses dual prices to value trading constraints
//
// =============================================================================
func Num05DualityTheoryDemo() {
	// TODO: implement
}

// =============================================================================
// Num06SensitivityAnalysisDemo
//
// TOPIC: Sensitivity Analysis & Shadow Prices
//   - Shadow price of constraint i: ∂z*/∂bᵢ = yᵢ* (the i-th dual variable)
//   - RHS ranging: range of bᵢ over which current basis remains optimal
//   - Objective ranging: range of cⱼ over which current basis remains optimal
//   - Binding vs non-binding constraints: slacks = 0 vs slacks > 0
//   - 100% rule for simultaneous changes
//   - Interpreting the simplex final tableau: B⁻¹ encodes all ranging info
//
// Real world: "what if we add 1 more unit of capacity?" in supply chain;
//
//	sensitivity reports in Excel Solver / Gurobi; stress testing LP models
//
// =============================================================================
func Num06SensitivityAnalysisDemo() {
	// TODO: implement
}

// =============================================================================
// Num07IntegerLinearProgrammingDemo
//
// TOPIC: Integer Linear Programming (ILP / MIP)
//   - ILP: same Ax ≤ b, cᵀx objective, but x ∈ ℤⁿ (or {0,1}ⁿ for binary)
//   - LP relaxation: drop integrality → continuous lower bound
//   - Branch-and-bound: branch on fractional variable xⱼ*:
//     left:  xⱼ ≤ ⌊xⱼ*⌋,  right: xⱼ ≥ ⌈xⱼ*⌉
//   - Cutting planes (Gomory cuts): add valid inequalities to tighten LP relaxation
//   - Integrality gap: (ILP opt - LP opt) / LP opt
//   - Special structures: totally unimodular matrices → LP relaxation is integral
//
// Real world: nurse scheduling, ad auction allocation (Meta), feature selection,
//
//	neural architecture search (NAS) formulated as ILP at DeepMind
//
// =============================================================================
func Num07IntegerLinearProgrammingDemo() {
	// TODO: implement
}

// =============================================================================
// Num08NetworkFlowLPDemo
//
// TOPIC: Network Flow as Linear Programming
//   - Graph G=(V,E): nodes V, directed edges E with capacities uₑ and costs cₑ
//   - Flow conservation at each node v:  Σ fₑ (in) - Σ fₑ (out) = bᵥ
//   - Written as: Ax = b  where A = incidence matrix ∈ {-1, 0, 1}^{|V|×|E|}
//   - Min-cost flow LP: min cᵀf  s.t. Af = b, 0 ≤ f ≤ u
//   - Max-flow as LP: max flow = min cut (LP strong duality proof)
//   - Shortest path = special case of min-cost flow (unit supply/demand)
//   - Totally unimodular incidence matrix → optimal LP solution always integral
//
// Real world: Google Maps routing, logistics (FedEx vehicle routing),
//
//	data center traffic engineering, GPU memory bandwidth allocation
//
// =============================================================================
func Num08NetworkFlowLPDemo() {
	// TODO: implement
}

// =============================================================================
// Num09AssignmentTransportationDemo
//
// TOPIC: Assignment Problem & Transportation LP
//   - Assignment: n workers, n jobs, cost matrix C ∈ ℝⁿˣⁿ
//     min  Σᵢ Σⱼ cᵢⱼ xᵢⱼ
//     s.t. Σⱼ xᵢⱼ = 1  ∀i  (each worker assigned once)
//     Σᵢ xᵢⱼ = 1  ∀j  (each job filled once)
//     xᵢⱼ ∈ {0,1}   → LP relaxation always integral (TU)
//   - Hungarian algorithm: O(n³) combinatorial solver
//   - Transportation: m suppliers, n consumers, supply sᵢ, demand dⱼ
//     min Σᵢⱼ cᵢⱼ xᵢⱼ  s.t. row sums = sᵢ, col sums = dⱼ, x ≥ 0
//   - Northwest corner method → MODI method for optimal transport
//
// Real world: Uber/Lyft driver-rider matching (Hungarian at scale),
//
//	OpenAI compute job scheduling, Anthropic GPU cluster allocation
//
// =============================================================================
func Num09AssignmentTransportationDemo() {
	// TODO: implement
}

// =============================================================================
// Num10InteriorPointMethodsDemo
//
// TOPIC: Interior Point Methods (IPM)
//   - Barrier / log-barrier formulation:
//     min  cᵀx − μ Σⱼ ln(xⱼ)   s.t. Ax = b
//   - Central path: parameterized by μ > 0; as μ→0 path converges to LP optimum
//   - Newton step on the KKT system:  solve [A 0; 0 X][Δx; Δy] = [b; c-Aᵀy]
//     where X = diag(x)
//   - Polynomial-time complexity: O(√n · L) iterations (Karmarkar 1984)
//   - Warm-starting for sequential LP (vs cold simplex)
//   - When IPM beats simplex: very large dense LPs, SDPs, second-order cone programs
//
// Real world: all modern solvers (Gurobi interior point mode, IPOPT, MOSEK);
//
//	training convex ML models (LogReg, SVMs at Google scale);
//	xAI / Grok inference serving resource LP solved via IPM
//
// =============================================================================
func Num10InteriorPointMethodsDemo() {
	// TODO: implement
}

// =============================================================================
// Num11LPMachineLearningsvmDemo
//
// TOPIC: LP in Machine Learning — Support Vector Machines
//   - Hard-margin SVM:
//     min  ½‖w‖²   s.t. yᵢ(wᵀxᵢ + b) ≥ 1  ∀i
//   - Soft-margin SVM (introduces slack ξᵢ ≥ 0):
//     min  ½‖w‖² + C Σ ξᵢ   s.t. yᵢ(wᵀxᵢ + b) ≥ 1 − ξᵢ
//   - Dual SVM (kernel trick):
//     max  Σ αᵢ − ½ Σᵢⱼ αᵢαⱼyᵢyⱼ K(xᵢ,xⱼ)
//     s.t. 0 ≤ αᵢ ≤ C,  Σ αᵢyᵢ = 0
//   - Connection to LP: L1-SVM  (replace ½‖w‖² with ‖w‖₁) is a true LP
//   - KKT conditions → complementary slackness → support vectors
//   - Structural risk minimisation via margin maximisation
//
// Real world: DeepMind safety classifiers, Anthropic RLHF reward margin constraints,
//
//	feature selection in high-dimensional genomics (L1-SVM)
//
// =============================================================================
func Num11LPMachineLearningsvmDemo() {
	// TODO: implement
}

// =============================================================================
// Num12LPMachineLearningLassoDemo
//
// TOPIC: LP in Machine Learning — LASSO & Basis Pursuit
//   - LASSO regression:
//     min  ½‖Ax − b‖₂²  +  λ‖x‖₁
//   - Basis Pursuit (compressed sensing):
//     min  ‖x‖₁   s.t. Ax = b
//   - Reformulation as LP (introduce tᵢ ≥ |xᵢ|):
//     min  cᵀt   s.t.  -tᵢ ≤ xᵢ ≤ tᵢ,  Ax = b
//   - L1 norm geometry: sparsity-inducing corners of the ℓ₁ ball
//   - Restricted isometry property (RIP): sufficient condition for exact recovery
//   - Coordinate descent vs LP solver for LASSO: when each wins
//   - Elastic net: mix of L1 + L2 — only L1 part is LP-representable
//
// Real world: feature selection in biomedical AI (Perplexity data pipelines),
//
//	sparse signal recovery (MRI reconstruction, radar),
//	compressed sensing for LLM weight pruning / quantization research
//
// =============================================================================
func Num12LPMachineLearningLassoDemo() {
	// TODO: implement
}

// =============================================================================
// Num13LPRelaxationRLPlanningDemo
//
// TOPIC: LP Relaxation in Reinforcement Learning & Planning
//   - MDP: states S, actions A, transitions P(s'|s,a), rewards R(s,a), discount γ
//   - Bellman optimality equations (nonlinear fixed-point):
//     V*(s) = max_a [ R(s,a) + γ Σ P(s'|s,a) V*(s') ]
//   - LP formulation of MDP (de Farias & Van Roy 2003):
//     min  cᵀV   s.t. V(s) ≥ R(s,a) + γ Σ P(s'|s,a)V(s')  ∀s,a
//     where c ∈ Δ(S) is a state relevance weighting
//   - Linear programming approximation (ALP): project V onto Φθ, solve smaller LP
//   - Connection to policy gradient: dual of MDP LP = occupancy measure LP
//   - Linear–quadratic regulator (LQR): continuous-state LP analogue
//
// Real world: DeepMind AlphaFold scheduling, OpenAI o1 reasoning resource allocation,
//
//	robotics trajectory optimisation (linear MPC = LP at each step)
//
// =============================================================================
func Num13LPRelaxationRLPlanningDemo() {
	// TODO: implement
}

// =============================================================================
// Num14ColumnGenerationLargeScaleDemo
//
// TOPIC: Column Generation & Large-Scale LP
//   - Motivation: LP with millions of variables — can't enumerate all columns
//   - Dantzig-Wolfe decomposition: master problem + subproblem structure
//   - Column generation loop:
//     1. Solve restricted master problem (RMP) with subset of columns
//     2. Compute dual variables y* from RMP
//     3. Solve pricing subproblem: find column with reduced cost < 0
//     min  cⱼ - yᵀAⱼ  over feasible xⱼ
//     4. Add new column to RMP; repeat until no improving column exists
//   - Cutting stock problem: canonical example — bin packing as LP
//   - Airline crew scheduling: millions of pairings, solved by column generation
//   - Branch-and-price: column generation inside branch-and-bound for ILP
//
// Real world: Google OR-Tools fleet routing, airline scheduling (IBM CPLEX),
//
//	Anthropic / OpenAI cluster job packing (GPU bin packing)
//
// =============================================================================
func Num14ColumnGenerationLargeScaleDemo() {
	// TODO: implement
}

// =============================================================================
// Num15EndToEndLPSolverDemo
//
// TOPIC: End-to-End LP Solver — Full Pipeline in Go
//   - Stage 1: Model definition (variables, constraints, objective as Go structs)
//   - Stage 2: Conversion to standard form (add slacks, negate max→min)
//   - Stage 3: Phase I — find initial BFS (big-M or two-phase method)
//   - Stage 4: Phase II — simplex iterations until optimality
//   - Stage 5: Extract primal solution x*, dual variables y*, reduced costs
//   - Stage 6: Generate sensitivity report (shadow prices, ranging intervals)
//   - Stage 7: ILP via branch-and-bound on fractional variables
//   - Full worked example: production mix problem (3 products, 4 constraints)
//     max  5x₁ + 4x₂ + 3x₃
//     s.t. 6x₁ + 4x₂ + 2x₃ ≤ 240   (machine hours)
//     3x₁ + 2x₂ + 5x₃ ≤ 270   (labour hours)
//     5x₁ + 6x₂ + 5x₃ ≤ 420   (materials)
//     x₁, x₂, x₃ ≥ 0
//
// Real world: implementing a teaching-grade solver mirrors Gurobi/GLPK internals;
//
//	understanding this is what separates LP users from LP engineers
//	at Two Sigma, Jane Street, DeepMind infrastructure teams
//
// =============================================================================
func Num15EndToEndLPSolverDemo() {
	// TODO: implement
}
