# M3/W3/D4 - Thu, 19 Mar 2026 (WIB)

## Differential Equations — Full Curriculum

A structured learning path covering ordinary and partial differential equations,  
from first-order separable ODEs through Laplace transforms, phase plane analysis,  
and classical PDEs (heat, wave, Laplace).  
Each topic is a runnable Go simulation with mathematical derivations, worked examples,  
and rich inline comments. Topics build on each other — read in order.

Target audience: engineer or scientist who needs rigorous ODE/PDE fluency for  
numerical computing, control systems, physics simulation, or quant finance modelling.

---

## Curriculum Table

| # | Topic | Core Concept | Real World Application |
| --- | --- | --- | --- |
| 01 | First-Order Separable ODEs | Separation of variables, IVP, equilibria | Population growth, radioactive decay, Newton's cooling |
| 02 | Integrating Factor & Linear ODEs | μ(x) = exp(∫P dx), Bernoulli substitution | RC circuits, mixing problems, forced cooling |
| 03 | Exact Equations & Potential Functions | Exactness condition ∂M/∂y = ∂N/∂x, potential F | Conservative force fields, thermodynamic state functions |
| 04 | Second-Order Constant-Coefficient ODEs | Characteristic equation, three damping cases | Mass-spring-damper, RLC circuits, structural vibration |
| 05 | Undetermined Coefficients | y = y\_h + y\_p, resonance, superposition | Forced oscillations, AC circuit steady-state |
| 06 | Variation of Parameters | Wronskian, u₁/u₂ integration, Green's functions | Any driven oscillator with arbitrary forcing |
| 07 | Laplace Transform | L{f}, derivative rule, partial fractions, convolution | Control systems, signal processing, PID controllers |
| 08 | Systems of ODEs & Phase Plane | Eigenvalue classification, Jacobian linearization | Lotka-Volterra, SIR epidemic, coupled oscillators |
| 09 | Numerical Methods — Euler & RK4 | Euler O(h), RK4 O(h⁴), stiffness, adaptive step | GBM simulation, N-body physics, circuit simulation |
| 10 | Power Series & Frobenius Method | Recurrence relations, indicial equation, Bessel Jᵥ | Quantum mechanics, acoustics, cylindrical heat conduction |
| 11 | Heat Equation & Fourier Series | ∂u/∂t = α∂²u/∂x², separation of variables, Fourier modes | GPU/CPU thermal simulation, metallurgy slab cooling |
| 12 | Wave Equation & D'Alembert | D'Alembert formula, standing waves, normal modes | String vibration, acoustics, EM wave propagation |
| 13 | Laplace Equation & Harmonic Functions | ∇²u = 0, Dirichlet BVP, Poisson integral formula | Electrostatics, steady-state temperature, fluid potential |
| 14 | Stability & Lyapunov Methods | Lyapunov functions, bifurcations, Poincaré-Bendixson | Control design, power grid stability, epidemic equilibria |
| 15 | End-to-End ODE Solver Pipeline | Adaptive RK4, stiffness detection, SIR + Van der Pol | Production numerical library; quant finance PDE solvers |

---

## How to run

```
# run all topics
go run ./learn-differential-equations/

# run a specific topic: comment/uncomment the relevant line in main.go
```

---

## Function signatures

```
func Num01FirstOrderSeparableDemo()         // stub — implement when teaching
func Num02IntegratingFactorLinearDemo()     // stub — implement when teaching
func Num03ExactEquationsPotentialDemo()     // stub — implement when teaching
func Num04SecondOrderConstCoeffDemo()       // stub — implement when teaching
func Num05UndeterminedCoefficientsDemo()    // stub — implement when teaching
func Num06VariationOfParametersDemo()       // stub — implement when teaching
func Num07LaplaceTransformDemo()            // stub — implement when teaching
func Num08SystemsODEsPhsePlaneeDemo()       // stub — implement when teaching
func Num09NumericalMethodsEulerRK4Demo()    // stub — implement when teaching
func Num10SeriesSolutionsFrobeniusDemo()    // stub — implement when teaching
func Num11HeatEquationFourierDemo()         // stub — implement when teaching
func Num12WaveEquationDAlembertDemo()       // stub — implement when teaching
func Num13LaplaceEquationHarmonicDemo()     // stub — implement when teaching
func Num14StabilityLyapunovDemo()           // stub — implement when teaching
func Num15EndToEndODESolverDemo()           // stub — implement when teaching
```

---

## Num01 — First-Order Separable ODEs

### The Problem

The simplest ODEs — and the ones that appear most often in applied science and  
engineering — can be solved by separating variables. Mastering this technique  
builds the intuition for all higher-order methods.

### The Concept

**Standard Separable Form**

```
dy/dx = f(x) · g(y)
```

Separate and integrate:

```
∫ dy/g(y) = ∫ f(x) dx + C
```

**Initial Value Problem (IVP)**  
Given y(x₀) = y₀, substitute to solve for the constant C.

**Autonomous ODEs**  
When f(x) = 1, the equation dy/dx = g(y) has no explicit x dependence.  
Equilibria (steady states): solve g(y\*) = 0.  
Stability: if g'(y\*) \< 0, the equilibrium is stable; if g'(y\*) > 0, unstable.

**Classic examples**

| Model | ODE | Solution |
| --- | --- | --- |
| Exponential growth | dy/dt = ky | y = y₀ e^(kt) |
| Logistic growth | dy/dt = ky(1 − y/K) | y = K / (1 + Ce^(-kt)) |
| Radioactive decay | dN/dt = −λN | N = N₀ e^(−λt) |
| Newton's cooling | dT/dt = −k(T − T\_env) | T = T\_env + (T₀ − T\_env)e^(−kt) |

### What the demo shows

*   Solve dy/dx = xy, verify with exact solution
*   Logistic growth: simulate population converging to carrying capacity K
*   Radioactive decay: compute half-life from λ programmatically
*   Newton's cooling: simulate a hot object cooling in ambient air

### Interview Tips

*   **Always check separability first.** Can you write the right-hand side as f(x)·g(y)?
*   **Equilibria from autonomous ODEs**: set g(y) = 0 before solving. Interviewers  
    often ask "what is the long-run behaviour?" — that's the stable equilibrium.
*   **Logistic model trap**: the solution saturates at K, not infinity. Many candidates  
    forget the denominator term.

---

## Num02 — Integrating Factor & Linear First-Order ODEs

### The Problem

Not every first-order ODE is separable. The integrating factor technique converts  
any linear first-order ODE into an exact derivative that can be integrated directly.

### The Concept

**Standard Linear Form**

```
dy/dx + P(x)·y = Q(x)
```

**Integrating Factor**

```
μ(x) = exp(∫ P(x) dx)
```

Multiplying both sides: d/dx\[μ(x)·y\] = μ(x)·Q(x)

**Solution**

```
y = (1/μ(x)) · [ ∫ μ(x)·Q(x) dx + C ]
```

**Bernoulli Equation** (nonlinear reduction)

```
dy/dx + P(x)·y = Q(x)·y^n    (n ≠ 0, 1)
```

Substitution v = y^(1-n) linearises the equation.

### Key Formulas

| Formula | Expression |
| --- | --- |
| Integrating factor | μ = exp(∫P dx) |
| General solution | y = (1/μ)∫μQ dx + C/μ |
| RC circuit charge | Q(t) = CE(1 − e^(−t/RC)) |
| Bernoulli substitution | v = y^(1−n), v' + (1−n)P·v = (1−n)Q |

### What the demo shows

*   Solve y' + 2y = 4x — integrating factor, then verify
*   RC circuit: charge build-up Q(t) as a function of R, C, and voltage E
*   Mixing problem: salt concentration in a tank with inflow/outflow
*   Bernoulli equation: logistic growth recast as Bernoulli

### Interview Tips

*   **The integrating factor is always an exponential.** If P(x) is constant,  
    μ = e^(Px) — a fact worth memorising for quick mental computation.
*   **Check your sign in P(x).** The standard form has +P(x)·y, not −P(x)·y.  
    Many errors come from misidentifying the sign.

---

## Num03 — Exact Equations & Potential Functions

### The Problem

Some first-order ODEs written as M dx + N dy = 0 are the total differential of a  
scalar function F(x,y) = C. Recognising and exploiting this structure gives an  
implicit solution directly — no integration by parts needed.

### The Concept

**Exactness Condition**

```
M dx + N dy = 0  is exact  ⟺  ∂M/∂y = ∂N/∂x
```

**Potential Function**

```
Find F such that  ∂F/∂x = M  and  ∂F/∂y = N
Solution:  F(x,y) = C
```

**Making a Non-Exact Equation Exact**  
Integrating factor μ(x) or μ(y) such that μM dx + μN dy = 0 is exact.

### What the demo shows

*   Check exactness condition and build F(x,y) step by step
*   Solve a non-exact ODE by finding an integrating factor
*   Thermodynamics example: verify dU = T dS − P dV is an exact differential

### Interview Tips

*   **Exactness is a symmetry condition.** It says the mixed partials of F are equal —  
    this is Clairaut's theorem. Any time you see M dx + N dy, check ∂M/∂y vs ∂N/∂x first.
*   **Potential functions appear in physics** as conservative force fields.  
    The work done is path-independent exactly when the force is exact.

---

## Num04 — Second-Order Constant-Coefficient ODEs

### The Problem

The backbone of mechanical and electrical engineering: every vibrating structure,  
RLC circuit, and control loop reduces to a second-order linear ODE with constant  
coefficients. The three damping cases are fundamental engineering literacy.

### The Concept

**Homogeneous Form**

```
a y'' + b y' + c y = 0
```

**Characteristic Equation**

```
a r² + b r + c = 0    →    discriminant Δ = b² − 4ac
```

**Three Cases**

| Case | Discriminant | Roots | General Solution |
| --- | --- | --- | --- |
| Overdamped | Δ > 0 | r₁ ≠ r₂ real | C₁e^(r₁x) + C₂e^(r₂x) |
| Critically damped | Δ = 0 | r₁ = r₂ = r | (C₁ + C₂x)e^(rx) |
| Underdamped | Δ \< 0 | α ± βi | e^(αx)(C₁cos βx + C₂sin βx) |

**Damped Harmonic Oscillator** (mass m, damping b, spring k)

```
m x'' + b x' + k x = 0
ω₀ = √(k/m)  (natural frequency),   ζ = b/(2√(mk))  (damping ratio)
overdamped ζ>1,  critical ζ=1,  underdamped ζ<1
```

### What the demo shows

*   Simulate all three damping regimes; plot x(t) for each
*   RLC circuit: V'' + (R/L)V' + (1/LC)V = 0 — same structure
*   Compute natural frequency ω₀ and damping ratio ζ from m, b, k
*   Show energy decay: E(t) = ½kx² + ½mv² over time

### Interview Tips

*   **Underdamped is the most common exam case.** Know that α = −b/(2a) is the  
    decay rate and β = √(4ac−b²)/(2a) is the oscillation frequency.
*   **Critical damping** is the fastest non-oscillatory return to equilibrium —  
    important in car suspensions and precision instruments.

---

## Num05 — Undetermined Coefficients

### The Problem

For non-homogeneous ODEs ay'' + by' + cy = g(x), we need a particular solution y\_p.  
When g(x) has a recognisable form, guessing y\_p with undetermined coefficients  
is the fastest approach.

### The Concept

**General Solution Structure**

```
y = y_h  +  y_p
```

y\_h solves the homogeneous equation; y\_p is any particular solution.

**Guessing Table**

| g(x) | Trial y\_p | Modification if resonance |
| --- | --- | --- |
| Polynomial x^n | degree-n polynomial | — |
| e^(ax) | Ae^(ax) | multiply by x (if a is a root) |
| cos ωx or sin ωx | A cos ωx + B sin ωx | multiply by x (if ±ωi are roots) |
| e^(ax) cos ωx | e^(ax)(A cos ωx + B sin ωx) | multiply by x if needed |

**Superposition**  
If g(x) = g₁(x) + g₂(x), find y\_p1 and y\_p2 separately, then y\_p = y\_p1 + y\_p2.

### What the demo shows

*   Solve y'' + 4y = cos 2x (resonance: multiply trial by x)
*   Forced harmonic oscillator: amplitude diverges at resonance frequency ω₀
*   AC circuit: find steady-state current under sinusoidal voltage

### Interview Tips

*   **Resonance is the key trap.** If the forcing frequency matches a natural frequency  
    of y\_h, the standard trial fails — you must multiply by x.
*   **For polynomials**: always use a full polynomial of the same degree even if  
    some coefficients look like they could be zero.

---

## Num06 — Variation of Parameters

### The Problem

Undetermined coefficients only works when g(x) is a polynomial, exponential,  
or sinusoid. For arbitrary g(x) — including 1/x, ln x, or a numerical function —  
variation of parameters is the universal fallback.

### The Concept

**Setup**  
Given y'' + p(x)y' + q(x)y = g(x) with homogeneous solutions y₁, y₂.  
Seek y\_p = u₁(x)y₁ + u₂(x)y₂ where u₁, u₂ are functions.

**Wronskian**

```
W(y₁,y₂) = y₁y₂' − y₂y₁'
```

**Formulas**

```
u₁' = −y₂·g / W        u₂' = y₁·g / W
u₁  = −∫ y₂·g/W dx     u₂  = ∫ y₁·g/W dx
```

**Green's Function Interpretation**  
The solution can be written as y\_p(x) = ∫ G(x, s) g(s) ds — this is the basis of  
Green's function methods used in mathematical physics.

### What the demo shows

*   Solve y'' + y = tan x (not a standard form for undetermined coefficients)
*   Numerically compute u₁, u₂ when g(x) has no closed-form antiderivative
*   Verify y\_p by substitution back into the ODE

### Interview Tips

*   **Variation of parameters always works** for linear ODEs with known y\_h.  
    It is the method of last resort — more general but more computational.
*   **The Wronskian must be non-zero.** W = 0 means y₁, y₂ are linearly dependent  
    (they are not a fundamental pair).

---

## Num07 — Laplace Transform

### The Problem

Laplace transforms convert an ODE (a problem in calculus) into an algebraic equation  
in the frequency domain. Initial conditions are automatically incorporated. This is  
the primary tool in control engineering.

### The Concept

**Definition**

```
L{f(t)} = F(s) = ∫₀^∞ e^(-st) f(t) dt
```

**Key Transform Table**

| f(t) | F(s) |
| --- | --- |
| 1 | 1/s |
| t^n | n!/s^(n+1) |
| e^(at) | 1/(s−a) |
| sin ωt | ω/(s²+ω²) |
| cos ωt | s/(s²+ω²) |
| u(t−a)f(t−a) | e^(−as) F(s) |

**Derivative Rule**

```
L{f'} = sF(s) − f(0)
L{f''} = s²F(s) − sf(0) − f'(0)
```

**IVP Solution Recipe**

```
1. Transform both sides: algebraic equation in F(s)
2. Solve for F(s)
3. Partial fractions decompose F(s)
4. Inverse transform: f(t) = L⁻¹{F(s)}
```

**Convolution Theorem**

```
L{f*g} = F(s) · G(s)        (f*g)(t) = ∫₀ᵗ f(τ)g(t−τ) dτ
```

### What the demo shows

*   Solve IVP y'' + 3y' + 2y = e^(−t), y(0)=0, y'(0)=1 using Laplace
*   Build a transform table and verify numerically
*   Simulate a PID controller transfer function in the s-domain
*   Demonstrate the time-delay property with Heaviside step input

### Interview Tips

*   **Partial fractions are the bottleneck.** Practice decomposing 1/((s+1)(s+2)) and  
    (s+3)/((s²+4)) quickly — this is 80% of the inverse transform work.
*   **Initial conditions go in automatically** in the Laplace method, unlike  
    undetermined coefficients where you apply them at the end.
*   **Transfer function H(s) = Y(s)/X(s)** is the Laplace of the impulse response.  
    This is the language of control engineering.

---

## Num08 — Systems of ODEs & Phase Plane

### The Problem

Many real systems have multiple coupled state variables — a predator and prey  
population, two coupled electrical loops, an epidemic with S, I, R compartments.  
Systems of ODEs are the natural language for all of these.

### The Concept

**Vector Form**

```
x' = Ax    x ∈ ℝⁿ,  A is n×n constant matrix
```

**Solution via Eigendecomposition**

```
x(t) = Σ cₖ vₖ e^(λₖt)
```

where λₖ, vₖ are eigenvalue-eigenvector pairs of A.

**Phase Plane Classification (2D)**

| Eigenvalues of A | Equilibrium Type |
| --- | --- |
| λ₁ \< λ₂ \< 0 (real) | Stable node |
| λ₁ > λ₂ > 0 (real) | Unstable node |
| λ₁ \< 0 \< λ₂ (real) | Saddle (always unstable) |
| α ± βi, α \< 0 | Stable spiral |
| α ± βi, α > 0 | Unstable spiral |
| ±βi (pure imaginary) | Centre (neutrally stable) |

**Nonlinear Systems**  
Linearise about equilibrium x\*: let u = x − x\*, then u' ≈ Df(x\*)·u.  
Stability follows from the Jacobian eigenvalues (Hartman-Grobman theorem).

### What the demo shows

*   Classify 6 different 2×2 matrices by eigenvalue type; show phase portrait sketch
*   Lotka-Volterra predator-prey: simulate oscillating populations
*   SIR epidemic model: simulate infection curve with RK4
*   Jacobian linearisation of the pendulum near equilibria θ=0 and θ=π

### Interview Tips

*   **Trace and determinant shortcut (2D):** tr(A) = λ₁+λ₂, det(A) = λ₁λ₂.  
    Stable ⟺ tr \< 0 and det > 0. Check this before computing eigenvalues.
*   **Saddle points are always unstable** regardless of other eigenvalues.
*   **Lotka-Volterra has no stable equilibrium** in the classic model — populations  
    orbit forever. Adding a carrying capacity creates a stable spiral.

---

## Num09 — Numerical Methods — Euler & RK4

### The Problem

Most ODEs that arise in practice have no closed-form solution. Numerical methods  
are the workhorse tool. RK4 is the standard choice: four function evaluations per  
step gives fourth-order accuracy — a massive improvement over Euler.

### The Concept

**Euler Method** (first-order, O(h) error per step)

```
y_{n+1} = y_n + h · f(t_n, y_n)
```

**Improved Euler / Heun** (second-order, O(h²))

```
ỹ_{n+1} = y_n + h · f(t_n, y_n)                   (predictor)
y_{n+1} = y_n + (h/2)[f(t_n, y_n) + f(t_{n+1}, ỹ_{n+1})]  (corrector)
```

**Runge-Kutta 4 (RK4)** (fourth-order, O(h⁴))

```
k₁ = h · f(t_n,       y_n)
k₂ = h · f(t_n + h/2, y_n + k₁/2)
k₃ = h · f(t_n + h/2, y_n + k₂/2)
k₄ = h · f(t_n + h,   y_n + k₃)
y_{n+1} = y_n + (k₁ + 2k₂ + 2k₃ + k₄) / 6
```

**Stiff Equations**  
When the solution has components with vastly different time scales, explicit methods  
require very small h to stay stable. Use implicit methods (backward Euler, trapezoidal).

**Euler-Maruyama** (stochastic ODEs — SDEs)

```
X_{n+1} = X_n + f(X_n)·h + g(X_n)·√h·Z_n    Z_n ~ N(0,1)
```

This is how GBM is simulated in quant finance.

### What the demo shows

*   Compare Euler, Heun, RK4 on y' = −2y, y(0)=1: error vs step size h
*   Show convergence order: log(error) vs log(h) slopes 1, 2, 4
*   SIR model solved with RK4 — same as Num08 but focus on solver implementation
*   Euler-Maruyama simulation of GBM: N paths, compute mean and variance

### Interview Tips

*   **RK4 is the default.** Unless told otherwise, use RK4 for any numerical ODE problem.
*   **Doubling h halves the error for Euler, quarters it for RK4.** This O(h⁴) advantage  
    means RK4 with h=0.1 is often more accurate than Euler with h=0.001.
*   **Stiffness detection**: if you need h \< 10⁻⁶ for stability but the solution  
    changes on scale 1, the equation is stiff — switch to an implicit method.

---

## Num10 — Power Series Solutions & Frobenius Method

### The Problem

When the ODE has variable coefficients, the characteristic equation trick no longer  
works. Power series methods find solutions as infinite polynomial expansions — and  
they are the origin of the special functions (Bessel, Legendre) that appear  
throughout physics and engineering.

### The Concept

**Ordinary Point — Power Series**  
If x₀ is an ordinary point of y'' + P(x)y' + Q(x)y = 0, assume:

```
y = Σ aₙ (x − x₀)^n
```

Substitute into ODE → recurrence relation for aₙ.  
Radius of convergence R = distance to nearest singular point.

**Regular Singular Point — Frobenius Method**  
If x₀ = 0 is a regular singular point (xP(x) and x²Q(x) analytic at 0):

```
y = x^r Σ aₙ xⁿ
```

Indicial equation (from n=0 term): r(r−1) + p₀·r + q₀ = 0  
Two roots r₁ ≥ r₂.

| Case | Second Solution |
| --- | --- |
| r₁ − r₂ ∉ ℤ | x^(r₂) Σ bₙ xⁿ |
| r₁ = r₂ | y₁ ln x + x^r Σ cₙ xⁿ |
| r₁ − r₂ = integer | may involve ln x |

**Bessel's Equation**

```
x²y'' + xy' + (x² − ν²)y = 0     →     Jᵥ(x), Yᵥ(x)
```

### What the demo shows

*   Solve y'' − xy = 0 (Airy equation) via power series; compute 20 terms
*   Frobenius: solve 2xy'' + y' + y = 0 about the regular singular point x=0
*   Compute J₀(x) via its power series and compare to standard values

### Interview Tips

*   **Always identify the type of singular point first.** Ordinary → power series.  
    Regular singular → Frobenius. Irregular singular → asymptotic methods.
*   **Bessel functions are unavoidable** in cylindrical geometry problems  
    (heat in a pipe, drum vibration, waveguides). Know J₀ and J₁ at least by shape.

---

## Num11 — Heat Equation & Fourier Series

### The Problem

The heat equation ∂u/∂t = α∂²u/∂x² governs temperature diffusion in any solid.  
Its solution via Fourier series is the archetype for all PDE methods using  
separation of variables — and Fourier series underpin signal processing and  
spectral methods in numerical PDE solvers.

### The Concept

**Separation of Variables**  
Assume u(x,t) = X(x)·T(t):

```
T'/αT = X''/X = −λ    (separation constant)
```

With boundary conditions X(0)=X(L)=0:

```
λₙ = (nπ/L)²,    Xₙ = sin(nπx/L),    n = 1, 2, 3, ...
Tₙ(t) = e^(−α(nπ/L)²t)
```

**Fourier Sine Series**  
Expand initial condition f(x) = Σ bₙ sin(nπx/L):

```
bₙ = (2/L) ∫₀ᴸ f(x) sin(nπx/L) dx
```

**Solution**

```
u(x,t) = Σ bₙ sin(nπx/L) · e^(−α(nπ/L)²t)
```

High-frequency modes (large n) decay fastest: e^(−α n²π²t/L²) → 0.

**Full Fourier Series** (for periodic functions on \[−L, L\])

```
f(x) = a₀/2 + Σ [aₙ cos(nπx/L) + bₙ sin(nπx/L)]
aₙ = (1/L)∫ f cos,    bₙ = (1/L)∫ f sin
```

### What the demo shows

*   Solve heat equation on \[0,1\] with initial condition u(x,0) = sin(πx) + 0.5sin(3πx)
*   Compute Fourier coefficients of a square wave; show Gibbs phenomenon
*   Visualise u(x,t) at t=0, 0.01, 0.1, 1 — watch high modes decay first
*   Energy: E(t) = ∫u² dx decays exponentially; plot and verify

### Interview Tips

*   **High modes decay faster** — this is why smooth initial conditions  
    self-smooth over time. Intuition: fine-scale temperature variations equilibrate quickly.
*   **Gibbs phenomenon**: a truncated Fourier series overshoots at a jump  
    discontinuity by ~9% regardless of how many terms are used.
*   **Connection to signal processing**: the DFT/FFT is the discrete analogue of  
    the Fourier series. Spectral PDE solvers use FFT to compute ∂²u/∂x² in O(N log N).

---

## Num12 — Wave Equation & D'Alembert's Solution

### The Problem

The wave equation ∂²u/∂t² = c²∂²u/∂x² governs sound, light, vibrating strings,  
and elastic waves. Unlike the heat equation (parabolic, smoothing), the wave  
equation is hyperbolic: it propagates information at finite speed c.

### The Concept

**D'Alembert's Formula**  
For the infinite line with initial data u(x,0)=f(x), u\_t(x,0)=g(x):

```
u(x,t) = ½[f(x+ct) + f(x−ct)] + (1/2c) ∫_{x−ct}^{x+ct} g(s) ds
```

Interpretation: the solution is two waves — one traveling right at speed c,  
one traveling left at speed c. They pass through each other unchanged.

**Standing Waves on \[0,L\]**  
With u(0,t) = u(L,t) = 0 (fixed endpoints):

```
u(x,t) = Σ [Aₙ cos(ωₙt) + Bₙ sin(ωₙt)] sin(nπx/L)
ωₙ = nπc/L    (normal mode frequencies)
```

The Aₙ, Bₙ come from the Fourier expansion of the initial conditions.

**Resonance and Normal Modes**  
The lowest frequency ω₁ = πc/L is the fundamental; higher modes are harmonics.  
Resonance occurs when a driving frequency matches ωₙ — energy builds without bound.

### What the demo shows

*   D'Alembert: simulate a plucked string (triangular initial shape, zero initial velocity)
*   Show the two traveling waves separating and recombining at the boundary
*   Compute the first 5 normal mode frequencies for a string of given length and tension
*   Energy conservation: kinetic + potential = constant over time

### Interview Tips

*   **Wave equation vs heat equation**: wave preserves information (time-reversible),  
    heat destroys it (time-irreversible). This is a fundamental physics distinction.
*   **Speed c = √(T/ρ)** for a string: T = tension, ρ = linear mass density.  
    Doubling tension raises frequency by √2 (one musical half-step is 2^(1/12) ≈ 1.059).
*   **D'Alembert is faster than Fourier** when you need the solution at one (x,t) point.  
    Fourier series is better for visualising the full spacetime picture.

---

## Num13 — Laplace Equation & Harmonic Functions

### The Problem

The Laplace equation ∇²u = 0 describes steady-state phenomena: temperature at  
equilibrium, electrostatic potential in charge-free space, and ideal fluid flow.  
Its solutions — harmonic functions — have the maximum principle and mean-value property.

### The Concept

**Laplace Equation in 2D**

```
∂²u/∂x² + ∂²u/∂y² = 0
```

**Key Properties of Harmonic Functions**

*   **Maximum principle**: u attains its max and min on the boundary, never in the interior
*   **Mean value property**: u(x₀,y₀) = average of u on any circle centred at (x₀,y₀)
*   **Uniqueness**: Dirichlet BVP (u prescribed on boundary) has at most one solution

**Separation in Cartesian**

```
X''/X = −Y''/Y = ±λ    →    X = e^(±√λ x),  Y = e^(±√λ y)
```

**Separation in Polar (r, θ)**  
For a disk of radius a with boundary condition u(a,θ) = f(θ):

```
u(r,θ) = a₀/2 + Σ (r/a)^n [aₙ cos nθ + bₙ sin nθ]
```

**Poisson Integral Formula**

```
u(r,θ) = (1/2π) ∫₀^{2π} [(a²−r²) / (a²−2ar cos(θ−φ)+r²)] f(φ) dφ
```

### What the demo shows

*   Solve Laplace on a rectangle with non-zero boundary condition on one edge
*   Verify mean-value property numerically on a computed harmonic function
*   Poisson integral formula: recover harmonic function from boundary values on disk
*   Fluid flow: potential flow around a cylinder (u and stream function ψ)

### Interview Tips

*   **Uniqueness is guaranteed** by the maximum principle for Dirichlet BC.  
    For Neumann BC (∂u/∂n prescribed), solution is unique up to an additive constant.
*   **Physical intuition**: a harmonic function is the "most uniform" function  
    consistent with its boundary values — no unnecessary oscillation inside the domain.

---

## Num14 — Stability & Lyapunov Methods

### The Problem

Beyond solving ODEs, we need to understand the long-run behaviour of dynamical  
systems: do trajectories converge to an equilibrium? Do they oscillate forever?  
Can a parameter change flip a stable system to an unstable one (bifurcation)?

### The Concept

**Stability Definitions** (equilibrium at x\*)

*   **Stable (Lyapunov)**: trajectories starting near x\* stay near x\*
*   **Asymptotically stable**: stable AND trajectories → x\* as t → ∞
*   **Unstable**: some trajectories near x\* escape

**Linearisation (Hartman-Grobman)**  
For x' = f(x), the stability of x\* is determined by eigenvalues of the Jacobian  
Df(x\*) — provided no eigenvalue has zero real part.

**Lyapunov Function**  
A scalar V(x) > 0 (positive definite) with V̇ = ∇V · f ≤ 0:

```
V̇ < 0  everywhere (except x*)  →  asymptotically stable
V̇ ≤ 0                          →  stable (Lyapunov stable)
```

**Bifurcations**

| Type | Description | Example |
| --- | --- | --- |
| Saddle-node | two equilibria collide and annihilate | x' = μ − x² |
| Pitchfork | one equilibrium splits into three | x' = μx − x³ |
| Hopf | spiral changes stability; limit cycle born | chemical oscillators |

**Poincaré-Bendixson Theorem** (2D)  
If a trajectory is bounded and has no equilibria, it must approach a limit cycle.

### What the demo shows

*   Build a Lyapunov function for the damped harmonic oscillator: V = ½kx² + ½mv²
*   Show V̇ = −bv² ≤ 0 — verifies asymptotic stability
*   Bifurcation diagram: equilibria of x' = μx − x³ as μ varies (pitchfork)
*   SIR epidemic: show the disease-free equilibrium is stable when R₀ \< 1

### Interview Tips

*   **The Jacobian eigenvalue test is your first tool.** If all eigenvalues have  
    negative real parts, the equilibrium is asymptotically stable.
*   **Lyapunov functions are not unique** — there is no algorithm to find them.  
    For physical systems, try energy. For others, quadratic forms V = xᵀPx work for  
    linear systems (Lyapunov equation AᵀP + PA = −Q).
*   **R₀ in epidemics** is a stability condition: R₀ = βS₀/γ.  
    R₀ > 1 → epidemic grows (unstable DFE). R₀ \< 1 → disease dies out (stable DFE).

---

## Num15 — End-to-End ODE Solver Pipeline

### The Problem

A production ODE solver must: choose the right algorithm, control step size  
adaptively, detect stiffness, and handle systems. This topic integrates  
everything into a complete, general-purpose solver pipeline.

### The Concept

**Solver Selection Logic**

```
1. Separable first-order?   → exact analytic solution
2. Linear with const coeff? → characteristic equation
3. General non-stiff ODE?   → adaptive RK4
4. Stiff ODE?               → implicit Euler or trapezoidal rule
```

**Adaptive Step Size (RK4/RK5 Doubling)**

```
Compute y_{n+1} with step h   (one step of size h)
Compute ỹ_{n+1} with step h/2 (two steps of size h/2)
error ≈ |y_{n+1} − ỹ_{n+1}| / 15
if error > tol: halve h, retry
if error < tol/32: double h for next step
```

**Stiffness Detection**  
Estimate Jacobian eigenvalue λ\_max. If |λ\_max| · h >> 1, switch to implicit method.

**Applied Problems in this Demo**

| Problem | Model | Method |
| --- | --- | --- |
| SIR epidemic | S'= −βSI, I'= βSI−γI, R'= γI | Adaptive RK4 |
| Van der Pol oscillator | x'' − μ(1−x²)x' + x = 0 (μ large → stiff) | Implicit trapezoidal |
| Black-Scholes PDE | Reduce to heat equation, then solve | Crank-Nicolson (finite difference) |

### What the demo shows

*   Full adaptive RK4 with step doubling; log step size vs time
*   SIR model: vary β, γ and observe R₀ crossing 1 (bifurcation)
*   Van der Pol: show Euler fails at μ=1000 but implicit method succeeds
*   Black-Scholes as heat equation: price a European call numerically, compare to analytic

### Interview Tips

*   **Know the Crank-Nicolson scheme** for parabolic PDEs: it is the standard  
    second-order implicit finite-difference method and appears in quant finance everywhere.
*   **Adaptive step control** is what makes production solvers robust. A solver  
    that uses fixed h will either be slow (h too small) or blow up (h too large).
*   **Black-Scholes is a heat equation** after a change of variables. This is the  
    most elegant connection between PDEs and quant finance — know the substitution cold.

---