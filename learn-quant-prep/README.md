# M3/W3/D4 - Thu, 19 Mar 2026 (WIB)

## Quant Prep — Full Curriculum

A structured learning path for quantitative developer / SWE roles at firms like  
**Hudson River Trading, Jump Trading, Two Sigma, Jane Street, Citadel, D.E. Shaw**.  
Each topic is a runnable Go simulation with realistic financial scenarios and rich inline  
comments. Topics build on each other — read in order.

Target audience: strong software engineer who needs to level up on quant math, market  
structure, and the coding patterns that appear in quant firm interviews.

---

## Curriculum Table

| # | Topic | Core Concept | Real World at Quant Firms |
| --- | --- | --- | --- |
| 01 | Probability Foundations | Expected value, conditional prob, Bayes' theorem | Every quant phone screen; dice/card problems |
| 02 | Combinatorics & Brain Teasers | Counting, symmetry tricks, classic puzzle patterns | Jane Street, HRT, Susquehanna oral rounds |
| 03 | Statistics & Distributions | Normal, Poisson, log-normal, CLT, moments | Return distributions, risk modeling |
| 04 | Monte Carlo Simulation | Random sampling, convergence, variance reduction | Options pricing, VaR, portfolio simulation |
| 05 | Random Walk & GBM | Geometric Brownian Motion, drift, diffusion | Stock price modeling, Wiener process |
| 06 | Order Book Simulation | Bids, asks, FIFO matching engine, bid-ask spread | HRT, Jump — core market structure knowledge |
| 07 | Market Microstructure | Price impact, adverse selection, market maker logic | HFT firms, market making desks |
| 08 | Statistical Arbitrage | Mean reversion, cointegration, pairs trading z-score | Two Sigma, Citadel equity stat-arb strategies |
| 09 | Time Series Analysis | Autocorrelation, rolling stats, momentum signals | Signal research, alpha generation |
| 10 | Risk Metrics | VaR, CVaR, Sharpe ratio, max drawdown, correlation | Risk management, portfolio construction |
| 11 | Factor Models | CAPM, beta, alpha, Fama-French basics | Portfolio attribution, hedging, factor investing |
| 12 | Linear Algebra & PCA | Covariance matrix, eigendecomposition, PCA | Portfolio optimization, dimensionality reduction |
| 13 | Portfolio Optimization | Markowitz mean-variance, Kelly criterion | Allocation, bet sizing, efficient frontier |
| 14 | Coding: Numeric Patterns | Sliding window, prefix sums, two-pointer on prices | HRT / Jump coding rounds — finance-framed LC |
| 15 | Mental Math & Estimation | Fast arithmetic, Fermi estimation, approximation | Jane Street, Susquehanna timed oral interviews |

---

## How to run

```
# run all topics
go run ./learn-quant-prep/

# run a specific topic: comment/uncomment the relevant line in main.go
```

---

## Function signatures

```
func Num01ProbabilityFoundationsDemo()    // stub — implement when teaching
func Num02CombinatoricsBrainTeasersDemo() // stub — implement when teaching
func Num03StatisticsDistributionsDemo()   // stub — implement when teaching
func Num04MonteCarloSimulationDemo()      // stub — implement when teaching
func Num05RandomWalkGBMDemo()             // stub — implement when teaching
func Num06OrderBookSimulationDemo()       // stub — implement when teaching
func Num07MarketMicrostructureDemo()      // stub — implement when teaching
func Num08StatisticalArbitrageDemo()      // stub — implement when teaching
func Num09TimeSeriesAnalysisDemo()        // stub — implement when teaching
func Num10RiskMetricsDemo()               // stub — implement when teaching
func Num11FactorModelsDemo()              // stub — implement when teaching
func Num12LinearAlgebraPCADemo()          // stub — implement when teaching
func Num13PortfolioOptimizationDemo()     // stub — implement when teaching
func Num14CodingNumericPatternsDemo()     // stub — implement when teaching
func Num15MentalMathEstimationDemo()      // stub — implement when teaching
```

---

## Num01 — Probability Foundations

### The Problem

Quant firm phone screens almost always open with a probability question. Before any  
market structure or coding, you need to be fluent in expected value, conditional  
probability, and Bayes' theorem — and you need to explain your reasoning out loud.

### The Concept

**Expected Value (EV)**  
The probability-weighted average of all possible outcomes.

```
E[X] = Σ p(x) * x
```

Example: a fair die has E\[X\] = (1+2+3+4+5+6)/6 = 3.5

**Conditional Probability**

```
P(A | B) = P(A ∩ B) / P(B)
```

"Given that B happened, what is the probability of A?"

**Bayes' Theorem**

```
P(A | B) = P(B | A) * P(A) / P(B)
```

The canonical application: you have a prior belief P(A), you observe evidence B,  
and you update to a posterior P(A|B).

**Common interview framing at quant firms:**

*   "A fair coin is flipped 3 times. Given at least one head, what is P(all heads)?" → Bayes
*   "You roll two dice. What is E\[sum | first die = 4\]?" → conditional EV
*   "A test for a rare disease is 99% accurate. The disease affects 1 in 10,000 people.  
    You test positive. What is the probability you have the disease?" → Bayes trap

### Key Formulas

| Formula | Expression |
| --- | --- |
| Expected Value | E\[X\] = Σ p\_i \* x\_i |
| Variance | Var\[X\] = E\[X²\] - (E\[X\])² |
| Conditional Prob | P(A |
| Bayes' Theorem | P(A |
| Law of Total Prob | P(B) = Σ P(B |

### What the demo shows

*   Simulate expected value of dice/card problems programmatically (verify by counting)
*   Bayes' theorem: rare disease problem, and a market-flavored version (signal accuracy)
*   Conditional EV problem: dice with constraints

### Real World Usage

Expected value is the foundation of every trading decision. A market maker quoting a  
two-sided market computes: E\[PnL per trade\] = p(fill) \* spread/2 - p(adverse selection) \* loss.  
Every bet sizing model (Kelly criterion, Num13) is built on EV.

### Interview Tips

*   **Always state your sample space first** before computing probabilities. Examiners at  
    Jane Street and HRT specifically look for rigorous framing.
*   **The Bayes trap**: most candidates compute P(disease | positive) as 99% — the correct  
    answer is ~1% because the prior is so low (1/10,000). Know this cold.
*   **Symmetry shortcut**: many dice/card problems have symmetric outcomes. Use symmetry  
    to avoid enumeration.
*   When asked "what is the expected number of flips to get HH?" — use the Markov chain  
    state method, not brute force.

---

## Num02 — Combinatorics & Brain Teasers

### The Problem

HRT, Jane Street, and Susquehanna are famous for oral puzzle rounds where you get a  
brain teaser and must reason through it live. These problems require combinatorics  
fluency and structured thinking — not memorization.

### The Concept

**Permutations vs Combinations**

```
P(n,k) = n! / (n-k)!    — ordered selection
C(n,k) = n! / (k!(n-k)!) — unordered selection (n choose k)
```

**Stars and Bars** — distributing k identical items into n bins:

```
C(k + n - 1, n - 1)
```

**Inclusion-Exclusion Principle**

```
|A ∪ B| = |A| + |B| - |A ∩ B|
```

**Classic quant brain teaser patterns:**

1.  **Symmetry trick**: "Two players alternately flip a coin. First to flip heads wins.  
    What is P(player 1 wins)?" → Geometric series: 1/2 + 1/8 + ... = 2/3
2.  **Invariant trick**: "You have a 8x8 chessboard, remove two opposite corners.  
    Can dominoes tile it?" → Color invariant (no, both removed squares are same color)
3.  **Expected steps (Markov)**: "Expected number of cards drawn until first ace?" → N/k
4.  **Handshake / graph problems**: n people, each shakes hands once with every other.  
    Total handshakes = C(n,2) = n(n-1)/2
5.  **The 100 prisoners problem**: optimal strategy uses cycle detection (permutation cycles)

### What the demo shows

*   Compute C(n,k) and P(n,k)
*   Solve 5 classic brain teasers programmatically (verify answers by simulation)
*   Geometric series sum for alternating-flip game
*   Expected card draw problem via simulation

### Interview Tips

*   **Think out loud.** Interviewers at quant firms evaluate your reasoning process,  
    not just the final answer. Narrate every step.
*   **Enumerate small cases first.** For n=2 or n=3, write out all possibilities.  
    The pattern usually becomes obvious.
*   **Use simulation as a sanity check.** Running 1 million trials in your head is not  
    possible, but knowing the simulation approach confirms your formula.
*   **Geometric series**: 1 + r + r² + ... = 1/(1-r) for |r| \< 1. Memorize this.

---

## Num03 — Statistics & Distributions

### The Problem

Quant developers need to reason about randomness in price returns: are they normal?  
Are they fat-tailed? What does a Poisson distribution model? These questions come up  
in both interviews and day-to-day signal research.

### The Concept

**Normal Distribution N(μ, σ²)**

```
f(x) = (1/σ√2π) * exp(-(x-μ)²/2σ²)
```

*   68% of data within 1σ, 95% within 2σ, 99.7% within 3σ (the "68-95-99.7 rule")
*   Financial returns are approximately normal on short timescales (but fat-tailed in reality)

**Log-Normal Distribution**  
If X ~ N(μ, σ²), then e^X is log-normal. Stock prices follow log-normal because  
log-returns are approximately normal and prices can't go negative.

**Poisson Distribution**  
Models count of rare events in a fixed interval: e.g., number of trades in 1 second.

```
P(X=k) = (λ^k * e^(-λ)) / k!     E[X] = Var[X] = λ
```

**Central Limit Theorem (CLT)**  
The sum (or mean) of n i.i.d. random variables converges to Normal as n → ∞,  
regardless of the underlying distribution.

```
X̄_n → N(μ, σ²/n)
```

**Moments**

| Moment | Meaning |
| --- | --- |
| Mean (1st) | Center of mass |
| Variance (2nd) | Spread |
| Skewness (3rd) | Asymmetry — negative skew is bad for risk |
| Kurtosis (4th) | Tail weight — fat tails = excess kurtosis |

### What the demo shows

*   Generate synthetic daily return series (normal + fat-tailed)
*   Compute mean, variance, skewness, kurtosis from scratch
*   Demonstrate CLT: show that sample means converge to normal
*   Poisson: simulate trade arrival counts per second

### Interview Tips

*   **Fat tails**: real financial returns have kurtosis > 3 (leptokurtic). Assuming normality  
    underestimates tail risk — this caused many hedge fund blowups (LTCM).
*   **Log-normal vs Normal**: prices are log-normal, returns are (approximately) normal.  
    Never say "stock prices follow a normal distribution."
*   **CLT trap**: CLT requires finite variance. Power-law distributions (like some HFT  
    microstructure data) may not satisfy this.

---

## Num04 — Monte Carlo Simulation

### The Problem

Many financial quantities have no closed-form solution (American options, path-dependent  
payoffs, portfolio VaR under complex correlations). Monte Carlo simulation is the  
universal numerical tool.

### The Concept

**Core Idea**  
Estimate E\[f(X)\] by averaging f over many random samples:

```
E[f(X)] ≈ (1/N) * Σ f(X_i)    where X_i ~ P(X)
```

Error decreases as O(1/√N) — doubling N halves the error.

**European Call Option Pricing via MC**  
Payoff at expiry: max(S\_T - K, 0)  
Under risk-neutral measure:

```
S_T = S_0 * exp((r - σ²/2)*T + σ*√T*Z)    Z ~ N(0,1)
```

Price = e^(-rT) \* E\[max(S\_T - K, 0)\]

**Variance Reduction Techniques**

*   **Antithetic variates**: pair each random draw Z with -Z → cuts variance in half
*   **Control variates**: use a known quantity (e.g., stock price mean) to reduce noise
*   **Importance sampling**: oversample the tails

**Convergence**

```
Standard Error = σ_payoff / √N
```

Plot error vs N on a log-log scale — should see slope = -0.5.

### What the demo shows

*   Price a European call option: MC vs Black-Scholes closed form
*   Show convergence: error vs N from 100 to 1,000,000 paths
*   Antithetic variates: same N, lower variance
*   Simulate portfolio VaR via Monte Carlo

### Interview Tips

*   **Know the error rate**: O(1/√N) is slow. 10x more paths = 3.16x more accuracy.
*   **When to use MC vs closed form**: MC is for path-dependent options (Asian, barrier,  
    lookback) and high-dimensional problems where analytical solutions don't exist.
*   **Seed your RNG for reproducibility**: in production, randomness must be auditable.

---

## Num05 — Random Walk & Geometric Brownian Motion

### The Problem

How do you model a stock price over time? The standard answer (and the foundation of  
all of quantitative finance since Black-Scholes) is Geometric Brownian Motion.  
Every quant dev is expected to be able to derive and simulate GBM.

### The Concept

**Discrete Random Walk**  
Each step: S\_{t+1} = S\_t + ε, ε ∈ {+1, -1} with equal probability  
Properties: E\[S\_t\] = S\_0, Var\[S\_t\] = t

**Brownian Motion (Wiener Process)**  
Continuous-time limit of the random walk:

*   W(0) = 0
*   W(t) - W(s) ~ N(0, t-s) for t > s
*   Independent increments
*   Continuous paths (but nowhere differentiable)

**Geometric Brownian Motion (GBM)**

```
dS = μS dt + σS dW
```

Solution (Itô's Lemma):

```
S(t) = S(0) * exp((μ - σ²/2)*t + σ*W(t))
```

*   μ = drift (expected return annualized)
*   σ = volatility (annualized)
*   The σ²/2 correction is the Itô term — do NOT forget it in interviews

**Log-returns**

```
log(S(t)/S(0)) = (μ - σ²/2)*t + σ*W(t)  ~ N((μ-σ²/2)*t, σ²*t)
```

### What the demo shows

*   Simulate a simple discrete random walk (N steps)
*   Simulate 1000 GBM paths with configurable μ and σ
*   Show the distribution of S(T) is log-normal
*   Compute realized volatility from simulated paths

### Interview Tips

*   **The Itô correction**: the drift of S(t) is μ, but the drift of log(S) is μ - σ²/2.  
    This is a classic interview trick. Interviewers will ask "what is E\[S(t)\]?" → S(0)\*e^(μt).
*   **Volatility scaling**: daily vol = annual vol / √252. For a 20% annual vol,  
    daily vol ≈ 20%/√252 ≈ 1.26% per day.
*   **GBM limitations**: assumes constant vol (violated in practice), no jumps,  
    continuous trading. Real prices have volatility clustering (GARCH) and jumps.

---

## Num06 — Order Book Simulation

### The Problem

Every trade at HRT, Jump, and Two Sigma goes through an order book. Understanding  
how a Limit Order Book (LOB) works — how orders are added, matched, and cancelled —  
is fundamental knowledge for any quant developer.

### The Concept

**Order Book Structure**

```
ASK side (sellers, ascending price)
  Ask[0]: 100.02  qty=500   ← best ask (lowest ask)
  Ask[1]: 100.03  qty=1200
  Ask[2]: 100.05  qty=800

  --- SPREAD = 100.02 - 100.01 = $0.01 ---

  Bid[0]: 100.01  qty=700   ← best bid (highest bid)
  Bid[1]: 100.00  qty=1500
  Bid[2]: 99.98   qty=300
BID side (buyers, descending price)
```

**Order Types**

*   **Limit Order**: "buy 100 shares at price ≤ $100.01" — rests in book, provides liquidity
*   **Market Order**: "buy 100 shares at best available price" — takes liquidity immediately
*   **Cancel**: remove a resting limit order

**Matching Engine (FIFO price-time priority)**

1.  New buy limit order at price P: match against all asks with price ≤ P, in price-time order
2.  Remaining unmatched quantity rests in the bid book
3.  This process determines the **fill price** and **executed quantity**

**Key Metrics**

*   **Mid price**: (best\_bid + best\_ask) / 2
*   **Spread**: best\_ask - best\_bid (in basis points: spread / mid \* 10000)
*   **Depth**: total volume within N ticks of mid
*   **VWAP**: volume-weighted average price of all fills

### What the demo shows

*   Build a Limit Order Book from scratch (sorted bid/ask sides)
*   Process a sequence of limit orders, market orders, and cancellations
*   Print book state after each event
*   Compute spread, mid price, and VWAP of executions

### Interview Tips

*   **Makers vs takers**: limit orders make liquidity (earn rebates at exchanges),  
    market orders take liquidity (pay fees). HFT firms are almost always makers.
*   **Price-time priority**: at the same price, earlier orders get filled first. This  
    is why HFT firms obsess over latency — being first in queue at a price level matters.
*   **Know the data structures**: bids are a max-heap or sorted map descending,  
    asks are a min-heap or sorted map ascending. Interview question: what is the  
    complexity of adding an order? Matching? → O(log N) with a sorted map.

---

## Num07 — Market Microstructure

### The Problem

Market microstructure is the study of how prices are formed at the microscopic level —  
tick by tick. Quant devs at HFT firms need to understand adverse selection, inventory  
risk, and the economics of market making.

### The Concept

**Market Making**  
A market maker simultaneously quotes a bid and an ask. They profit from the spread  
when uninformed traders hit their quotes, but lose when informed traders (who know  
the true price is moving) hit their quotes.

**Adverse Selection**  
If a trader hits your ask, it may be because they know good news is coming and the  
price will go up. Your fill was adversely selected — you sold low, they bought low.

```
E[PnL per trade] = p_uninformed * (spread/2) - p_informed * adverse_move
```

**Inventory Risk**  
As a market maker fills more buys than sells (or vice versa), they build up  
inventory. A long inventory position loses money if prices fall. Market makers  
skew their quotes to rebalance:

*   Long inventory → lower bid AND lower ask to attract sellers
*   Short inventory → raise bid AND raise ask to attract buyers

**Kyle's Lambda (Price Impact)**

```
ΔP = λ * Q
```

Larger orders move the price more. λ measures market impact per unit of order flow.

**The Roll Model**  
Observed spread S. True price change is εt ~ N(0,σ²). Observed returns have  
negative first-order autocorrelation = -S²/4 due to bid-ask bounce.

### What the demo shows

*   Simulate a market maker quoting around a random-walk true price
*   Track inventory, PnL, and spread over time
*   Show adverse selection: informed vs uninformed order flow
*   Compute realized spread (actual profit) vs quoted spread

### Interview Tips

*   **Realized spread \< quoted spread**: always. Adverse selection eats into it.
*   **Inventory management**: the most important practical skill in market making.  
    A market maker who ignores inventory will get picked off.
*   **Why do spreads widen before news?** Market makers fear adverse selection spikes,  
    so they widen quotes or pull them entirely. This causes flash crashes.

---

## Num08 — Statistical Arbitrage

### The Problem

Statistical arbitrage ("stat arb") is the strategy of exploiting mean-reverting  
relationships between related securities. Two Sigma, Renaissance, and D.E. Shaw  
made billions on stat arb. Quant devs build the pipelines that detect and trade  
these signals.

### The Concept

**Pairs Trading**  
Two stocks that are economically related (e.g., Coke and Pepsi) tend to move together.  
When they diverge, bet on convergence.

**Spread and Z-score**

```
spread(t) = price_A(t) - β * price_B(t)
z(t) = (spread(t) - mean(spread)) / std(spread)
```

Entry signal: |z(t)| > threshold (e.g., 2.0)  
Exit signal: |z(t)| \< exit\_threshold (e.g., 0.5)

**Cointegration**  
Two time series are cointegrated if their linear combination is stationary (mean-reverting).  
This is a stronger condition than correlation — it means there is a long-run equilibrium.  
Test: Engle-Granger two-step (regress, then ADF test on residuals).

**ADF Test (Augmented Dickey-Fuller)**  
Tests the null hypothesis that a time series has a unit root (non-stationary).  
Reject H0 → series is stationary → mean-reverting → tradeable.

**Half-life of mean reversion**

```
half_life = -ln(2) / β    where β from: Δspread_t = β * spread_{t-1} + ε
```

Short half-life = fast reversion = more trading opportunities.

### What the demo shows

*   Generate two cointegrated price series (synthetic)
*   Compute rolling spread and z-score
*   Simulate pairs trade with entry/exit rules, track PnL
*   Compute half-life of mean reversion from the series

### Interview Tips

*   **Correlation ≠ cointegration**: two random walks can be highly correlated without  
    being cointegrated. Stat arb requires cointegration (or at least stationarity of spread).
*   **Regime change risk**: the relationship between two stocks can break permanently  
    (e.g., a merger, bankruptcy). Always use a lookback window; don't over-fit.
*   **Transaction costs kill stat arb**: a z-score of 2 sounds big, but after commissions,  
    spread, and market impact, many signals are not profitable.

---

## Num09 — Time Series Analysis

### The Problem

Alpha signals in quantitative finance are time series. Understanding their statistical  
properties — autocorrelation, stationarity, momentum vs mean-reversion — is essential  
for building and evaluating trading strategies.

### The Concept

**Stationarity**  
A time series is stationary if its mean, variance, and autocovariance are constant over time.  
Required for most statistical models. Test with ADF test.

**Autocorrelation Function (ACF)**

```
ACF(k) = Cov(X_t, X_{t-k}) / Var(X_t)
```

*   ACF(1) > 0: momentum (trend-following)
*   ACF(1) \< 0: mean-reversion (bid-ask bounce, stat arb)

**Rolling Statistics**

*   Rolling mean: trend indicator
*   Rolling std: volatility estimator
*   Rolling Sharpe: signal quality over time

**ARIMA(p,d,q)**

*   AR(p): regress X\_t on past p values → captures momentum
*   MA(q): regress X\_t on past q errors → captures noise
*   I(d): differencing d times to achieve stationarity

**Momentum Signal**

```
signal(t) = mean(returns[t-window:t])  / std(returns[t-window:t])
```

Positive momentum: price went up recently, expect it to continue (short term).  
Negative autocorrelation at longer lags = mean reversion at longer timescales.

### What the demo shows

*   Generate a synthetic price series with known AR(1) coefficient
*   Compute and plot autocorrelation function (ACF)
*   Rolling mean, rolling std, rolling Sharpe over the series
*   Momentum signal: rolling z-score of recent returns

### Interview Tips

*   **Prices vs returns**: prices are almost never stationary. Returns often are.  
    Always check stationarity before applying AR/ARIMA models.
*   **Look-ahead bias**: the most common mistake in backtesting. Your rolling mean at  
    time t must only use data up to and including t. Never use future data.
*   **Sharpe ratio interpretation**: 1.0 = acceptable, 2.0 = good, 3.0+ = excellent  
    for a live strategy (annualized, after costs).

---

## Num10 — Risk Metrics

### The Problem

Every trading desk, portfolio manager, and risk system computes the same set of  
standard risk metrics. Quant devs are expected to implement these from scratch and  
understand their limitations.

### The Concept

**Sharpe Ratio**

```
Sharpe = (E[R] - R_f) / σ(R)    (annualized: multiply by √252 for daily returns)
```

Most commonly quoted measure of risk-adjusted return.

**Value at Risk (VaR)**  
The loss not exceeded with probability (1-α) over horizon T.

```
VaR_{α,T} = -quantile(returns, α)    e.g., 95% daily VaR
```

"With 95% probability, we will not lose more than $X today."

**CVaR (Conditional VaR / Expected Shortfall)**

```
CVaR_{α,T} = -E[R | R < -VaR_{α,T}]
```

Expected loss GIVEN that we're in the tail. More informative than VaR.  
Basel III requires CVaR (called ES) instead of VaR.

**Maximum Drawdown**

```
MDD = max over all t: (peak_value_before_t - trough_value_after_peak) / peak_value
```

The worst peak-to-trough loss. The most psychologically painful risk metric.

**Sortino Ratio**  
Like Sharpe but only penalizes downside volatility (σ of negative returns).

**Calmar Ratio**

```
Calmar = Annualized Return / |Maximum Drawdown|
```

### What the demo shows

*   Generate a synthetic return series (with a drawdown period)
*   Compute: Sharpe, Sortino, Calmar ratios
*   Compute VaR and CVaR at 95% and 99% confidence
*   Compute maximum drawdown with the peak-trough algorithm
*   Rolling Sharpe ratio over time

### Interview Tips

*   **VaR is NOT the worst loss**: it's the loss at the αth percentile. Actual losses  
    can be much worse. This is why CVaR (Expected Shortfall) is preferred.
*   **Sharpe assumes normality**: if returns are fat-tailed, Sharpe understates risk.  
    Use Sharpe alongside CVaR and drawdown for a complete picture.
*   **Annualization**: daily Sharpe \* √252 = annual Sharpe (√252 ≈ 15.87).

---

## Num11 — Factor Models

### The Problem

Why did our portfolio gain 5% yesterday? Was it a good bet, or did the whole market  
just go up? Factor models decompose returns into systematic (market-driven) and  
idiosyncratic (skill-driven) components.

### The Concept

**CAPM (Capital Asset Pricing Model)**

```
E[R_i] = R_f + β_i * (E[R_m] - R_f)
```

*   β\_i: sensitivity of asset i to the market
*   E\[R\_m\] - R\_f: market risk premium (~5-7% annually)
*   α: intercept above CAPM prediction → "alpha" (true skill)

**Computing Beta**

```
β = Cov(R_i, R_m) / Var(R_m)
```

Also the slope of OLS regression of R\_i on R\_m.

**Fama-French Three-Factor Model**  
Extends CAPM with two additional factors:

*   SMB (Small Minus Big): small-cap stocks outperform large-cap
*   HML (High Minus Low): value stocks outperform growth stocks

**Alpha vs Beta**

*   Beta returns are cheap (buy an index fund)
*   Alpha is what active managers charge for (and rarely deliver consistently)

### What the demo shows

*   Simulate market factor + asset returns with known β and α
*   Compute β via OLS regression (from scratch with matrix math)
*   Compute α (Jensen's alpha)
*   Extend to 3-factor model: decompose return into mkt/smb/hml/alpha

### Interview Tips

*   **Beta is backward-looking**: historical beta may not predict future beta.  
    Betas drift over time, especially for small-cap stocks.
*   **R² interpretation**: how much of asset variance is explained by the factor.  
    High R² = mostly systematic; low R² = mostly idiosyncratic.
*   **Factor crowding**: if everyone is long the same factors, adverse moves are  
    amplified. This caused the "quant quake" of August 2007.

---

## Num12 — Linear Algebra & PCA

### The Problem

Portfolio optimization and risk decomposition require matrix operations. PCA extracts  
the dominant risk factors from a covariance matrix. This is core to multi-factor risk  
models used at every major quant firm.

### The Concept

**Covariance Matrix**  
For N assets, the N×N covariance matrix Σ captures all pairwise return relationships.

```
Σ_ij = Cov(R_i, R_j) = E[(R_i - μ_i)(R_j - μ_j)]
```

Diagonal: variances. Off-diagonal: covariances.

**Portfolio Variance**

```
σ²_portfolio = w^T * Σ * w
```

where w is the vector of portfolio weights.

**PCA (Principal Component Analysis)**  
Decompose Σ into eigenvectors and eigenvalues:

```
Σ = Q * Λ * Q^T
```

*   Eigenvectors (Q): principal components — the "directions" of variance
*   Eigenvalues (Λ): amount of variance along each direction
*   PC1 typically explains ~40-60% of variance in an equity portfolio (market factor)
*   PC2 often captures sector effects
*   Truncate to k\<N components: dimensionality reduction

**Correlation vs Covariance**

```
ρ_ij = Σ_ij / (σ_i * σ_j)    (correlation, normalized to [-1,1])
```

Use correlation matrix for PCA (standardized) vs covariance for portfolio variance.

### What the demo shows

*   Build a synthetic 5-asset covariance matrix
*   Compute portfolio variance for various weight vectors
*   Implement power iteration to find the leading eigenvector (from scratch)
*   Run PCA: show variance explained by each PC
*   Minimum variance portfolio weights from Σ

### Interview Tips

*   **Condition number**: a poorly conditioned covariance matrix amplifies estimation error  
    in portfolio optimization. Shrinkage estimators (Ledoit-Wolf) help.
*   **PCA ≠ factor models**: PCA is data-driven (statistical factors). Fama-French uses  
    economic factors (SMB, HML). Both are used in practice.
*   **Positive semi-definiteness**: a valid covariance matrix must be PSD (all eigenvalues ≥ 0).  
    Estimation error can produce invalid (non-PSD) matrices — regularization needed.

---

## Num13 — Portfolio Optimization

### The Problem

Given N assets with expected returns and a covariance matrix, how do you allocate  
your capital to maximize return for a given level of risk? This is the Markowitz  
mean-variance optimization problem — the foundation of modern portfolio theory.

### The Concept

**Markowitz Mean-Variance Optimization**

```
maximize:   μ^T * w - (λ/2) * w^T * Σ * w
subject to: sum(w) = 1,  w_i >= 0  (optional)
```

*   μ: vector of expected returns
*   Σ: covariance matrix
*   λ: risk aversion parameter
*   w: portfolio weights (the solution)

**The Efficient Frontier**  
The set of portfolios that maximize return for each level of risk.  
Traced by sweeping λ from 0 (max return, ignores risk) to ∞ (min variance, ignores return).

**Minimum Variance Portfolio (MVP)**

```
w_MVP = Σ^(-1) * 1 / (1^T * Σ^(-1) * 1)
```

**Maximum Sharpe Ratio Portfolio (Tangency Portfolio)**  
On the efficient frontier, the portfolio with the highest Sharpe ratio.  
This is the "optimal" portfolio under the CAPM framework.

**Kelly Criterion**  
Optimal bet sizing to maximize long-run geometric growth:

```
f* = (p * b - q) / b    (binary bet: win $b with prob p, lose $1 with prob q=1-p)
```

Full Kelly is often too aggressive. Half-Kelly or fractional Kelly is used in practice.  
For continuous returns:

```
f* = μ / σ²    (fraction of capital to invest)
```

### What the demo shows

*   Sweep efficient frontier: generate 10,000 random portfolios, plot risk vs return
*   Compute minimum variance portfolio analytically
*   Compute maximum Sharpe ratio portfolio via grid search
*   Kelly criterion: binary bet + continuous version for a stock position

### Interview Tips

*   **Mean-variance is sensitive to expected returns**: small changes in μ lead to  
    large changes in optimal weights ("error maximization" problem). Use robust  
    estimators or Black-Litterman to stabilize.
*   **Concentration risk**: unconstrained Markowitz often produces extreme allocations  
    (100% in one asset). Add weight constraints in practice.
*   **Kelly vs fixed fractional**: Kelly maximizes growth but has huge variance. In  
    practice, most firms use 25-50% Kelly to reduce risk of ruin.

---

## Num14 — Coding: Numeric Patterns

### The Problem

HRT and Jump Trading coding interviews use LeetCode-style problems framed in financial  
contexts. The underlying patterns are standard DSA: sliding window, prefix sums,  
two pointers — but you need to recognize them in unfamiliar financial wrapping.

### The Concept

**Pattern 1: Sliding Window (max/min in window)**  
Use case: rolling max profit, rolling Sharpe, rolling maximum price in last N days.  
Monotonic deque gives O(1) per element for window max/min.

```
Time: O(N), Space: O(window)
```

**Pattern 2: Prefix Sums**  
Use case: sum of returns over any subarray \[i,j\] in O(1).

```
prefix[i] = prefix[i-1] + returns[i]
sum(i,j) = prefix[j] - prefix[i-1]
```

"Maximum subarray return" (Kadane's algorithm) is a prefix sum variant.

**Pattern 3: Two Pointers**  
Use case: finding pairs of prices with a given difference, or a range of days  
with sum of returns ≥ target.

**Pattern 4: Monotonic Stack**  
Use case: finding for each day the previous day with a higher/lower price ("next  
greater element"). Used in options payoff analysis and support/resistance levels.

**Finance-framed LeetCode problems:**

1.  "Best time to buy and sell stock" → max subarray / prefix approach
2.  "Max profit with cooldown" → DP on states
3.  "Rolling 30-day Sharpe of a return stream" → sliding window with online variance
4.  "Find all pairs of stocks with spread within $X" → two pointers on sorted prices
5.  "Longest profitable streak" → sliding window on sign of returns

### What the demo shows

*   Problem 1: best single buy/sell → O(N) one-pass
*   Problem 2: rolling 30-day Sharpe on return stream → sliding window + Welford's online variance
*   Problem 3: maximum subarray return (Kadane's)
*   Problem 4: pairs of assets with return difference ≤ threshold → two pointers
*   Problem 5: longest consecutive positive-return streak → sliding window

### Interview Tips

*   **Recognize the pattern fast**: in a 45-min coding round, you cannot afford to  
    rediscover sliding window from scratch. Practice until pattern recognition is instant.
*   **Edge cases in finance**: prices can be 0 (halted stock), negative returns are valid,  
    empty windows should return NaN not crash.
*   **Explain complexity**: always state O(N) time and O(window) space before coding.  
    Interviewers at HRT specifically probe this.

---

## Num15 — Mental Math & Estimation

### The Problem

Jane Street, Susquehanna, and Optiver are famous for asking oral mental math and  
Fermi estimation questions during interviews. Speed and accuracy under pressure are  
tested — not just the answer but the reasoning process.

### The Concept

**Fast Arithmetic Tricks**

| Operation | Trick | Example |
| --- | --- | --- |
| ×11 | Add digits, put sum in middle | 47×11 = 4(4+7)7 = 517 |
| ×25 | Divide by 4, multiply by 100 | 48×25 = 12×100 = 1200 |
| ×125 | Divide by 8, multiply by 1000 | 24×125 = 3×1000 = 3000 |
| Square near 50 | (50±n)² = 2500 ± 100n + n² | 47² = 2500-300+9 = 2209 |
| √ of non-perfect | Newton's method 1 step | √50 ≈ 7 + (50-49)/(2×7) ≈ 7.07 |
| % calculation | Use complementary % | 73% of 200 = 200 - 27% of 200 = 200-54 = 146 |

**Fermi Estimation Framework**

1.  Break into sub-problems
2.  Estimate each component from first principles
3.  Combine: multiply/add estimates
4.  Sanity check: is the order of magnitude right?

**Key numbers to memorize**

*   US population: ~330 million
*   World population: ~8 billion
*   NYSE daily volume: ~10 billion shares
*   S&P 500 market cap: ~$40 trillion
*   √252 ≈ 15.87 (annualization factor for daily returns)
*   ln(2) ≈ 0.693
*   e ≈ 2.718
*   1 basis point = 0.01%
*   Options: ATM call ≈ 0.4 \* σ \* √T (quick approximation)

**Classic Estimation Problems**

1.  "How many piano tuners in Chicago?" → Fermi
2.  "What is the bid-ask spread on a $1 stock vs $100 stock in bps?" → microstructure
3.  "If S&P 500 has daily vol of 1%, what is annual vol?" → × √252 ≈ 15.87%
4.  "Estimate the number of trades per second on NASDAQ" → ~1-2 million/day → ~12/sec
5.  "A stock moves $0.05 on 10,000 shares. What is Kyle's lambda?" → 0.05/10000 = 5e-6

### What the demo shows

*   20 timed estimation problems with solutions and worked reasoning
*   Mental math drills: ×25, ×125, squares, √ approximations
*   Finance-specific estimates: vol scaling, option premium approximation
*   Fermi chain: estimate NYSE daily notional volume from first principles

### Interview Tips

*   **Say your reasoning out loud**: at Jane Street especially, getting the wrong answer  
    but showing excellent reasoning is better than guessing the right answer silently.
*   **Order-of-magnitude first**: nail the exponent before worrying about the coefficient.  
    "About $10 billion" is more useful than a precise wrong answer.
*   **Practice daily**: 5 mental math problems per day for 2 weeks transforms your speed.  
    Apps like "Math Workout" or simple flashcards work well.
*   **Round aggressively**: π ≈ 3, √2 ≈ 1.4, ln(2) ≈ 0.7. Precision is the enemy of speed.

---

_Connection to other modules: Num04 (Monte Carlo) builds on Num01 (Probability) and Num03 (Distributions). Num08 (Stat Arb) builds on Num09 (Time Series). Num13 (Portfolio Opt) builds on Num12 (Linear Algebra) and Num10 (Risk Metrics)._

```
R_i = α + β_mkt*R_mkt + β_smb*SMB + β_hml*HML + ε
```