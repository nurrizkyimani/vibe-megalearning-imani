# M3/W3/D4 - Thu, 19 Mar 2026 (WIB)

## Linear Programming — Full Curriculum

A structured learning path for engineers who need rigorous LP fluency for optimization,  
operations research, machine learning, and quantitative finance.  
Each topic is a runnable Go simulation with mathematical derivations, worked examples,  
and rich inline comments. Topics build on each other — read in order.

Target audience: software engineer or data scientist who needs to understand how LP  
solvers work under the hood — from matrix algebra foundations through Simplex pivoting,  
duality, integer programming, and real-world applications in ML and RL.

---

## Curriculum Table

| # | Topic | Core Concept | Real World Application |
| --- | --- | --- | --- |
| 01 | Linear Equations & Systems | Gaussian elimination, rank, Ax = b, existence & uniqueness | Production planning, circuit analysis, supply-demand equilibrium |
| 02 | Geometric Interpretation | Half-spaces, feasible polyhedron, corner points, objective sweep | 2D diet/mix problems, portfolio efficient frontier |
| 03 | Standard & Canonical Forms | Slack variables, surplus, free variable splitting, max→min | Every LP solver (Gurobi, CPLEX) requires standard form |
| 04 | Simplex Method — Tableau | BFS, pivot selection, ratio test, reduced costs, Bland's rule | Resource allocation engines; airline/driver scheduling cores |
| 05 | Duality Theory | Weak/strong duality, complementary slackness, shadow prices | SVM dual = kernel trick; Jane Street constraint valuation |
| 06 | Sensitivity Analysis | Shadow prices, RHS ranging, objective ranging, 100% rule | "What if we add 1 more unit of capacity?" in supply chain |
| 07 | Integer Linear Programming | Branch-and-bound, Gomory cuts, LP relaxation, integrality gap | Nurse scheduling, ad auction allocation, neural architecture search |
| 08 | Network Flow as LP | Min-cost flow, max-flow = min-cut, incidence matrix, TU | Google Maps routing, logistics, GPU memory bandwidth |
| 09 | Assignment & Transportation | Hungarian algorithm, bipartite matching, northwest corner | Uber/Lyft driver-rider matching, GPU cluster allocation |
| 10 | Interior Point Methods | Log-barrier, central path, Newton step, Karmarkar O(√n·L) | Gurobi IPM mode, MOSEK, training SVMs at Google scale |
| 11 | LP in ML — SVM | Hard/soft margin SVM, dual SVM, kernel trick, L1-SVM as LP | DeepMind safety classifiers, feature selection in genomics |
| 12 | LP in ML — LASSO | Basis pursuit, L1 reformulation as LP, sparsity geometry, RIP | Compressed sensing, MRI reconstruction, LLM weight pruning |
| 13 | LP in RL & Planning | MDP LP formulation, ALP, occupancy measure, linear MPC | AlphaFold scheduling, robotics trajectory optimization |
| 14 | Column Generation | Dantzig-Wolfe decomposition, pricing subproblem, cutting stock | Airline crew scheduling, GPU bin packing, fleet routing |
| 15 | End-to-End LP Solver | Model → standard form → Phase I/II Simplex → sensitivity → ILP | Teaching-grade solver mirroring Gurobi/GLPK internals |

---

## How to run

```
# run all topics (only uncommented ones in main.go will execute)
go run ./learn-linear-programming/
```

To run a specific topic, comment/uncomment the relevant line in `main.go`.

---

## Function signatures

```
func Num01LinearEquationsSystemsDemo()       // ✅ fully implemented
func Num02GeometricInterpretationDemo()      // 🔲 stub — to be implemented
func Num03StandardCanonicalFormsDemo()       // 🔲 stub — to be implemented
func Num04SimplexMethodTableauDemo()         // 🔲 stub — to be implemented
func Num05DualityTheoryDemo()                // 🔲 stub — to be implemented
func Num06SensitivityAnalysisDemo()          // 🔲 stub — to be implemented
func Num07IntegerLinearProgrammingDemo()     // 🔲 stub — to be implemented
func Num08NetworkFlowLPDemo()                // 🔲 stub — to be implemented
func Num09AssignmentTransportationDemo()     // 🔲 stub — to be implemented
func Num10InteriorPointMethodsDemo()         // 🔲 stub — to be implemented
func Num11LPMachineLearningsvmDemo()         // 🔲 stub — to be implemented
func Num12LPMachineLearningLassoDemo()       // 🔲 stub — to be implemented
func Num13LPRelaxationRLPlanningDemo()       // 🔲 stub — to be implemented
func Num14ColumnGenerationLargeScaleDemo()   // 🔲 stub — to be implemented
func Num15EndToEndLPSolverDemo()             // 🔲 stub — to be implemented
```

**Legend:** ✅ = Fully implemented | 🔲 = Skeleton stub (to be implemented)

---

## Num01 — Linear Equations & Systems

### The Problem

Every LP starts with a system of linear equations. Before you can understand the Simplex  
method, duality, or integer programming, you need fluency in `Ax = b` — what it means  
for a solution to exist, when it's unique, and what Gaussian elimination actually does.

Most engineers can solve a 2×2 system. Almost none can explain _why_ Gaussian elimination  
works, what rank tells you, or how the three solution types (unique / none / infinite)  
map directly to LP outcomes (unique BFS / infeasible / feasible region).

### The Concept

**Matrix Form**

A system of m equations in n unknowns:

```
a₁₁x₁ + a₁₂x₂ + ... + a₁ₙxₙ = b₁
a₂₁x₁ + a₂₂x₂ + ... + a₂ₙxₙ = b₂
    ⋮
aₘ₁x₁ + aₘ₂x₂ + ... + aₘₙxₙ = bₘ
```

Compact form: `Ax = b`, where `A ∈ ℝᵐˣⁿ`, `x ∈ ℝⁿ`, `b ∈ ℝᵐ`.

**Gaussian Elimination**

Row-reduce the augmented matrix `[A | b]` to row-echelon form using three operations:

1.  Swap two rows
2.  Scale a row by a non-zero constant
3.  Add a multiple of one row to another

Each pivot eliminates a variable from all other rows. With partial pivoting (always  
choose the row with the largest absolute value in the pivot column), the algorithm is  
numerically stable.

**Rank and Solution Types**

| Condition | Outcome |
| --- | --- |
| rank(A) ≠ rank(\[A|b\]) | No solution — system is inconsistent |
| rank(A) = rank(\[A|b\]) = n | Unique solution |
| rank(A) = rank(\[A|b\]) \< n | Infinite solutions (free variables) |

**Complexity:** O(n²m) ≈ O(n³) for square systems.

#### LP Connection

```
rank(A) < rank([A|b]) → infeasible LP (empty feasible set)
rank = n              → unique basic feasible solution (BFS)
rank < n              → infinite solutions → LP feasible region (polytope after x ≥ 0)
```

### What the demo shows

```
Section 1: Matrix notation refresher — Ax = b, augmented matrix [A|b]
Section 2: Unique solution — 3×3 full-rank system, step-by-step Gaussian elimination
Section 3: No solution — inconsistent system, 0 = 1 contradiction, maps to infeasible LP
Section 4: Infinite solutions — underdetermined system, free variable, parametric solution
Section 5: Rank theorem table — existence/uniqueness, O(n³) complexity, condition number
Section 6: Real-world example — factory production planning (3 products, 3 machines)
Section 7: Summary — 5 key insights connecting Gaussian elimination to LP topics 02–15
```

### Key Insight

```
Gaussian elimination is the mechanical foundation of LP solving.
Three solution types map directly to LP outcomes: infeasible / unique BFS / feasible region.
Partial pivoting keeps elimination numerically stable — all production LP solvers use this.
O(n³) cost motivates sparse matrix techniques in every production LP solver.
```

### Real World Usage

| System | Uses linear system solving for |
| --- | --- |
| **Gurobi / CPLEX** | Basis factorization (LU decomposition) at each Simplex iteration |
| **LAPACK** | Dense linear system routines underpinning all numerical solvers |
| **NumPy / SciPy** | `numpy.linalg.solve` — same Gaussian elimination under the hood |
| **FEM solvers** | Millions of equations from discretized PDEs solved as Ax = b |
| **Supply chain** | Solving production plans at exact capacity (binding constraints) |

### Interview Tips

**"What does rank tell you about a linear system?"**

*   rank(A) = number of linearly independent rows = number of pivots after elimination
*   If rank(A) \< n, you have n - rank(A) free variables → infinite solutions
*   If rank(A) ≠ rank(\[A|b\]), the system is inconsistent → no solution

**"Why does Gaussian elimination use partial pivoting?"**

*   Without pivoting, a near-zero pivot amplifies floating-point errors
*   Swapping to the row with the largest |value| in the pivot column bounds error growth
*   This is why `numpy.linalg.solve` is stable even for large matrices

**"How does Ax = b relate to LP?"**

*   An LP in standard form is min cᵀx s.t. Ax = b, x ≥ 0
*   The Simplex method is essentially solving a sequence of linear systems (basis inversions)
*   Understanding Gaussian elimination is prerequisite for understanding Simplex pivoting

---

## Num02 — Geometric Interpretation

### The Problem

**To be implemented** — this section will cover the geometry of LP in 2D and 3D:  
half-spaces from linear constraints, the feasible polyhedron as their intersection,  
corner points as basic feasible solutions, the objective function sweeping across  
the feasible region as a family of parallel hyperplanes, and unbounded/infeasible/  
degenerate cases with ASCII diagrams.

---

## Num03 — Standard & Canonical Forms

### The Problem

**To be implemented** — this section will cover converting any LP into the two standard  
representations: inequality form (canonical) and equality form (standard). Conversion  
rules for slack variables, surplus variables, free variable splitting, and max→min  
transformation. Understanding this is prerequisite for implementing any LP solver.

---

## Num04 — Simplex Method — Tableau

### The Problem

**To be implemented** — this section will cover the full Simplex algorithm: basic feasible  
solutions, the simplex tableau (augmented matrix with objective row), entering variable  
selection (most negative reduced cost), leaving variable selection (minimum ratio test),  
pivot operation, termination condition (all reduced costs ≥ 0), degeneracy, cycling,  
and Bland's rule to prevent cycling.

---

## Num05 — Duality Theory

### The Problem

**To be implemented** — this section will cover LP duality: constructing the dual from the  
primal, weak duality inequality, strong duality theorem (optimal primal = optimal dual),  
complementary slackness conditions, and the economic interpretation of dual variables as  
shadow prices. Duality is the bridge between LP and SVMs (Topic 11).

---

## Num06 — Sensitivity Analysis

### The Problem

**To be implemented** — this section will cover how the optimal solution changes as  
problem parameters vary: shadow prices (∂z\*/∂bᵢ = dual variable), right-hand-side  
ranging (how much can bᵢ change before the current basis changes), objective coefficient  
ranging, binding vs non-binding constraints, and the 100% rule for simultaneous changes.

---

## Num07 — Integer Linear Programming

### The Problem

**To be implemented** — this section will cover ILP (same LP structure but with integrality  
constraints x ∈ ℤⁿ or x ∈ {0,1}ⁿ), the LP relaxation as a lower bound, branch-and-bound  
tree search, Gomory cutting planes to tighten the LP relaxation, the integrality gap,  
and totally unimodular matrices (where LP relaxation is automatically integral).

---

## Num08 — Network Flow as LP

### The Problem

**To be implemented** — this section will cover min-cost flow formulated as an LP using  
the graph incidence matrix, the max-flow = min-cut duality result (as a consequence of  
LP strong duality), shortest path as a special case of min-cost flow, and why the  
totally unimodular incidence matrix guarantees integral optimal solutions from the LP.

---

## Num09 — Assignment & Transportation

### The Problem

**To be implemented** — this section will cover the assignment problem as a bipartite  
matching LP (totally unimodular, so LP relaxation is integral), the Hungarian algorithm  
as an O(n³) primal-dual method, and the transportation problem with supply/demand  
constraints solved by the northwest corner + MODI method.

---

## Num10 — Interior Point Methods

### The Problem

**To be implemented** — this section will cover the log-barrier formulation, the central  
path parameterized by μ, Newton steps on the KKT system, Karmarkar's polynomial-time  
complexity O(√n · L), and when interior point methods beat Simplex (large dense LPs,  
second-order cone programs, semidefinite programs).

---

## Num11 — LP in Machine Learning — SVM

### The Problem

**To be implemented** — this section will cover the hard-margin and soft-margin SVM as  
quadratic programs, their dual formulations (where the kernel trick enters), the  
connection between L1-SVM and a true LP, KKT conditions and complementary slackness  
identifying support vectors, and structural risk minimization via margin maximization.

---

## Num12 — LP in Machine Learning — LASSO

### The Problem

**To be implemented** — this section will cover LASSO regression and basis pursuit  
(compressed sensing) reformulated as LPs by introducing auxiliary variables for the  
L1 norm, the geometry of the ℓ₁ ball (sparsity-inducing corners), the restricted  
isometry property (RIP) as sufficient condition for exact sparse recovery, and when  
to use a coordinate descent solver vs an LP solver for LASSO.

---

## Num13 — LP in RL & Planning

### The Problem

**To be implemented** — this section will cover the LP formulation of Markov Decision  
Processes (de Farias & Van Roy), where the Bellman optimality equations become linear  
constraints and solving the LP gives the optimal value function V\*, the approximate  
LP (ALP) for large state spaces via basis function projection, and the connection  
between the dual of the MDP LP and occupancy measures (the foundation of policy  
gradient methods).

---

## Num14 — Column Generation & Large-Scale LP

### The Problem

**To be implemented** — this section will cover the Dantzig-Wolfe decomposition for  
block-structured LPs, the column generation loop (solve restricted master → compute  
duals → solve pricing subproblem → add improving column → repeat), the cutting stock  
problem as the canonical example, and branch-and-price for large-scale ILPs.

---

## Num15 — End-to-End LP Solver

### The Problem

**To be implemented** — this section will implement a complete teaching-grade LP solver  
in Go: model definition (variables, constraints, objective), conversion to standard form,  
Phase I (finding initial BFS via big-M or two-phase), Phase II (Simplex iterations to  
optimality), extraction of primal solution, dual variables and reduced costs, sensitivity  
report (shadow prices, ranging intervals), and ILP via branch-and-bound. Full worked  
example: 3-product production mix problem.

---

---

## What to Expect in a Staff/Principal Engineer Interview Involving LP

After mastering all 15 topics, here's how LP knowledge translates to interview performance  
at firms like Google, DeepMind, Two Sigma, Jane Street, Citadel, and Gurobi.

### System Design Round (60–90 min)

**Typical prompt:**

> _"Design a resource allocation system for a cloud platform that assigns GPU jobs to_  
> _machines to minimize cost while satisfying capacity and latency constraints."_

**What they're testing:**

Can you recognize this as an LP/ILP and decompose it correctly?

*   **Formulate as ILP (Num07):** binary x\_ij = 1 if job i is on machine j
*   **Constraints:** capacity per machine (Num03 standard form), deadline SLA (inequality constraints)
*   **Objective:** minimize total cost cᵀx (Num04 Simplex)
*   **Scale:** millions of jobs → column generation (Num14) to avoid enumerating all variables
*   **LP relaxation (Num07):** solve continuous relaxation first as lower bound
*   **Sensitivity (Num06):** shadow prices tell you which machine constraints are binding

**The answer that separates Staff from Senior:**

> "For the scale of millions of jobs, I'd use column generation (Num14). Rather than  
> enumerate all job-machine assignments, I start with a restricted master problem on a  
> small subset, compute dual variables (shadow prices, Num05/Num06) from that solution,  
> and use them to identify the most profitable new assignment via a pricing subproblem.  
> This is exactly how airline crew scheduling solvers handle millions of pairings —  
> the same decomposition applies here."

### Algorithm/Optimization Round (45–60 min)

**Typical prompts:**

_"How would you find the minimum-cost assignment of N workers to N jobs?"_

**Expected answer (Num09):**

*   Formulate as assignment LP: min Σᵢⱼ cᵢⱼ xᵢⱼ, Σⱼ xᵢⱼ = 1, Σᵢ xᵢⱼ = 1, x ≥ 0
*   The constraint matrix is totally unimodular (TU) → LP relaxation is automatically integral
*   Solve with Hungarian algorithm: O(n³), which is the primal-dual method on the LP

_"Explain the relationship between max-flow and min-cut."_

**Expected answer (Num08):**

*   Max-flow = min-cut is a consequence of LP strong duality (Num05)
*   Primal LP: maximize flow; Dual LP: minimize cut capacity
*   Strong duality says primal optimal = dual optimal → max-flow = min-cut

_"What is the kernel trick in SVMs and why does it work?"_

**Expected answer (Num11):**

*   Solve the dual SVM LP/QP instead of the primal
*   The dual only involves inner products xᵢᵀxⱼ (via complementary slackness)
*   Replace xᵢᵀxⱼ with K(xᵢ, xⱼ) — any valid kernel — without changing the dual structure
*   This works because duality (Num05) guarantees the dual solution is equally optimal

### The One Answer That Proves You Understand LP at Depth

**Question:** _"A LASSO regression is taking 10 hours to train on your dataset. How would you speed it up?"_

**Junior answer:** "Use a faster machine or parallelize."

**Staff/Principal answer:**

> "It depends on the regime (Num12):
> 
> **1\. Small n, large p (compressed sensing regime):** Reformulate LASSO as a true LP  
> (introduce tᵢ ≥ |xᵢ|, minimize Σtᵢ + λ\*residual). Use a warm-started interior  
> point solver (Num10) — Gurobi IPM is often 100x faster than coordinate descent here  
> because it exploits the LP structure globally rather than cycling through coordinates.
> 
> **2\. Large n, moderate p:** Coordinate descent is better — O(np) per pass, cache-friendly,  
> easily parallelized. The LP formulation becomes expensive because n constraints (one  
> per data point) make the LP matrix huge.
> 
> **3\. Very large n and p:** Active set methods: start with the LP relaxation on a small  
> subset of features (column generation, Num14), add features with non-zero reduced cost.  
> This is basis pursuit with column generation — used in signal processing pipelines.
> 
> The choice between LP solver vs coordinate descent vs column generation is exactly  
> the Num12 content — recognizing which structural regime you're in is the skill."

---

## Why This Curriculum Makes You a Better Engineer

### The Market Reality (2026)

Most engineers who "use LP" know:

*   Call `scipy.optimize.linprog` or `gurobipy.Model()`
*   Set up variables, constraints, objective
*   Read the solution

**They cannot:**

*   Explain why the solver returned "infeasible" (Num01 — rank deficiency)
*   Interpret a shadow price output (Num05/Num06)
*   Know when branch-and-bound will be slow and how to fix it (Num07)
*   Recognize that their LASSO problem is an LP in disguise (Num12)
*   Understand why their network flow solver always returns integer solutions (Num08/Num09)

These 15 topics cover **exactly the gaps** between "I call a solver" and "I understand  
what the solver is doing."

### Cross-references to other curricula

*   **Num01 (Linear Systems)** → See `learn-differential-equations` Num09 (Numerical Methods) — same Gaussian elimination for ODE solvers
*   **Num04 (Simplex)** → See `learn-quant-prep` Num12 (Linear Algebra) — matrix operations underpin both
*   **Num05 (Duality)** → See `learn-quant-prep` Num13 (Portfolio Optimization) — KKT conditions appear in Markowitz optimization
*   **Num11 (SVM)** → See `learn-quant-prep` Num11 (Factor Models) — both use dual formulations for insight
*   **Num13 (LP in RL)** → See `learn-differential-equations` Num08 (Systems of ODEs) — both analyze state-space dynamics

---

## License

This curriculum is open for educational use. Go build solvers. Go get Staff/Principal offers.