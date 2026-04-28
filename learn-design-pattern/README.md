# M3/W5/D1 - Fri, 03 Apr 2026 (WIB)

## Design Patterns in Go - Full Curriculum

A structured learning path for writing maintainable, extensible Go systems using
proven design patterns. Each topic maps to one runnable `NumXX...` function,
so you can learn pattern-by-pattern and implement production-style architecture
without over-engineering.

Target audience: Go engineers who already write services, but want cleaner
abstractions, better testability, and scalable code organization.

---

## Curriculum Table

| # | Topic | Core Concept | Real World Usage | Demo |
| --- | --- | --- | --- | --- |
| 01 | Strategy Pattern | Swap business rules at runtime | Pricing, payment routing, recommendation logic | Num01 ✅ |
| 02 | Factory Method | Centralize object creation per use case | Driver/client creation by config | Num02 🔲 |
| 03 | Abstract Factory | Create families of related objects | Multi-cloud providers (AWS/GCP/Azure) | Num03 🔲 |
| 04 | Builder | Construct complex config safely step-by-step | HTTP client, query builder, SDK options | Num04 🔲 |
| 05 | Singleton | One shared instance with controlled access | Config manager, logger, metrics registry | Num05 🔲 |
| 06 | Adapter | Bridge incompatible interfaces | Third-party API wrapper normalization | Num06 🔲 |
| 07 | Facade | Simplify complex subsystem workflows | Checkout/service orchestration entrypoint | Num07 🔲 |
| 08 | Decorator | Add behavior dynamically around core logic | Middleware chains, retries, logging wrappers | Num08 🔲 |
| 09 | Proxy | Control access with lazy load/cache/auth | API gateway auth proxy, caching proxy | Num09 🔲 |
| 10 | Observer | Broadcast events to interested subscribers | Domain events, notification fan-out | Num10 🔲 |
| 11 | Chain of Responsibility | Pass request through handler chain | Validation and policy pipelines | Num11 🔲 |
| 12 | Command | Represent actions as objects | Job queues, undo/redo, task scheduling | Num12 🔲 |
| 13 | State | Behavior changes by internal state | Order lifecycle, workflow engines | Num13 🔲 |
| 14 | Template Method | Shared algorithm skeleton with custom steps | ETL and report generation templates | Num14 🔲 |
| 15 | Repository Pattern | Isolate persistence from domain logic | Swap DB implementations in tests | Num15 🔲 |
| 16 | Dependency Injection | Explicit dependency wiring | Testable services and clean constructors | Num16 🔲 |
| 17 | Circuit Breaker Pattern | Fail fast on unstable downstreams | Resilience in microservices | Num17 🔲 |
| 18 | Pub/Sub Event Bus | Loose coupling via events | Event-driven architecture, async workflows | Num18 🔲 |
| 19 | Pipeline Pattern | Multi-stage processing composition | Stream processing, ETL transformations | Num19 🔲 |
| 20 | Saga Orchestration | Distributed transaction with compensation | Order/payment/inventory consistency | Num20 🔲 |

Legend: ✅ = fully implemented | 🔲 = skeleton stub (to be implemented)

---

## How to run

```bash
# run all topics (only uncommented ones in main.go will execute)
go run ./learn-design-pattern/
```

To run a specific topic, comment/uncomment the relevant line in `main.go`.

---

## Function signatures

```go
func Num01StrategyPatternDemo()           // ✅ fully implemented
func Num02FactoryMethodPatternDemo()      // 🔲 stub — to be implemented
func Num03AbstractFactoryPatternDemo()    // 🔲 stub — to be implemented
func Num04BuilderPatternDemo()            // 🔲 stub — to be implemented
func Num05SingletonPatternDemo()          // 🔲 stub — to be implemented
func Num06AdapterPatternDemo()            // 🔲 stub — to be implemented
func Num07FacadePatternDemo()             // 🔲 stub — to be implemented
func Num08DecoratorPatternDemo()          // 🔲 stub — to be implemented
func Num09ProxyPatternDemo()              // 🔲 stub — to be implemented
func Num10ObserverPatternDemo()           // 🔲 stub — to be implemented
func Num11ChainOfResponsibilityDemo()     // 🔲 stub — to be implemented
func Num12CommandPatternDemo()            // 🔲 stub — to be implemented
func Num13StatePatternDemo()              // 🔲 stub — to be implemented
func Num14TemplateMethodPatternDemo()     // 🔲 stub — to be implemented
func Num15RepositoryPatternDemo()         // 🔲 stub — to be implemented
func Num16DependencyInjectionDemo()       // 🔲 stub — to be implemented
func Num17CircuitBreakerPatternDemo()     // 🔲 stub — to be implemented
func Num18PubSubEventBusDemo()            // 🔲 stub — to be implemented
func Num19PipelinePatternDemo()           // 🔲 stub — to be implemented
func Num20SagaOrchestrationDemo()         // 🔲 stub — to be implemented
```

---

## Num01 - Strategy Pattern (Pricing Engine)

### The Problem

Business rules change fast. If pricing logic is hardcoded with nested `if/else`,
the checkout code becomes brittle and hard to extend.

### The Concept

The Strategy pattern defines a shared interface for interchangeable algorithms.
Caller code depends on the interface, not concrete implementations.

In this demo:

* `PricingStrategy` is the contract.
* `RegularPricing`, `BulkDiscountPricing`, and `VIPDiscountPricing` are concrete strategies.
* `selectPricingStrategy(...)` chooses the right algorithm at runtime.

### What the demo shows

1. Reads multiple sample orders with different tiers and quantities.
2. Selects strategy dynamically (`VIP`, `bulk`, or `regular`).
3. Computes total and savings per order.
4. Prints gross revenue vs net revenue and total discount impact.

### Why this matters in production

* Add a new rule (for example, `FlashSalePricing`) without editing checkout flow.
* Keep business policy separate from transport/API/controller code.
* Unit test each strategy independently with deterministic inputs.

### Interview framing

When asked "When should we use Strategy?":

* Use it when behavior varies by context and changes frequently.
* Prefer it over long conditional branches spread across the codebase.
* Mention Open/Closed Principle: extend with new strategies, avoid modifying stable caller logic.

