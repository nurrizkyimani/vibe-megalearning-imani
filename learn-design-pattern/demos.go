package main

import (
	"fmt"
	"strings"
)

// =============================================================================
// Num01StrategyPatternDemo
//
// TOPIC: Strategy Pattern - Select behavior at runtime via interface composition
//   - Encapsulate interchangeable algorithms behind one contract
//   - Remove if-else explosion in business logic
//   - Make new behavior extensible without editing caller flow
//
// Real world: Dynamic pricing, shipping fee calculation, payment routing
// =============================================================================
func Num01StrategyPatternDemo() {
	fmt.Println("============================================================")
	fmt.Println("  Num01 -- Strategy Pattern (Pricing Engine)")
	fmt.Println("============================================================")
	fmt.Println()

	type order struct {
		customer  string
		tier      string
		unitPrice float64
		qty       int
	}

	orders := []order{
		{customer: "Alice", tier: "regular", unitPrice: 50, qty: 2},
		{customer: "Bob", tier: "regular", unitPrice: 50, qty: 12},
		{customer: "Carol", tier: "vip", unitPrice: 50, qty: 3},
		{customer: "Dedi", tier: "VIP", unitPrice: 80, qty: 15},
	}

	fmt.Println("Rule set:")
	fmt.Println("  - VIP tier gets 20% discount")
	fmt.Println("  - Regular tier with qty >= 10 gets 15% bulk discount")
	fmt.Println("  - Otherwise no discount")
	fmt.Println()

	fmt.Printf("%-8s %-8s %-5s %-10s %-28s %-10s %-10s\n",
		"Customer", "Tier", "Qty", "UnitPrice", "Strategy", "Total", "Savings")
	fmt.Println(strings.Repeat("-", 90))

	var grossRevenue float64
	var netRevenue float64

	for _, o := range orders {
		strategy := selectPricingStrategy(o.tier, o.qty)

		baseTotal := o.unitPrice * float64(o.qty)
		finalTotal := strategy.FinalTotal(o.unitPrice, o.qty)
		savings := baseTotal - finalTotal

		grossRevenue += baseTotal
		netRevenue += finalTotal

		fmt.Printf("%-8s %-8s %-5d $%-9.2f %-28s $%-9.2f $%-9.2f\n",
			o.customer,
			strings.ToLower(o.tier),
			o.qty,
			o.unitPrice,
			strategy.Name(),
			finalTotal,
			savings,
		)
	}

	fmt.Println(strings.Repeat("-", 90))
	fmt.Printf("Gross revenue (no discount): $%.2f\n", grossRevenue)
	fmt.Printf("Net revenue (with strategy): $%.2f\n", netRevenue)
	fmt.Printf("Total discount given:        $%.2f\n", grossRevenue-netRevenue)
	fmt.Println()
	fmt.Println("Key insight:")
	fmt.Println("  Caller code stays unchanged while pricing behavior swaps at runtime.")
	fmt.Println("  Add a new strategy type without modifying order processing flow.")
	fmt.Println()
}

// PricingStrategy defines a family of interchangeable pricing algorithms.
type PricingStrategy interface {
	Name() string
	FinalTotal(unitPrice float64, qty int) float64
}

type RegularPricing struct{}

func (RegularPricing) Name() string {
	return "RegularPricing (no discount)"
}

func (RegularPricing) FinalTotal(unitPrice float64, qty int) float64 {
	return unitPrice * float64(qty)
}

type BulkDiscountPricing struct{}

func (BulkDiscountPricing) Name() string {
	return "BulkDiscountPricing (15%)"
}

func (BulkDiscountPricing) FinalTotal(unitPrice float64, qty int) float64 {
	subtotal := unitPrice * float64(qty)
	if qty >= 10 {
		return subtotal * 0.85
	}
	return subtotal
}

type VIPDiscountPricing struct{}

func (VIPDiscountPricing) Name() string {
	return "VIPDiscountPricing (20%)"
}

func (VIPDiscountPricing) FinalTotal(unitPrice float64, qty int) float64 {
	return unitPrice * float64(qty) * 0.80
}

func selectPricingStrategy(customerTier string, qty int) PricingStrategy {
	if strings.EqualFold(customerTier, "vip") {
		return VIPDiscountPricing{}
	}
	if qty >= 10 {
		return BulkDiscountPricing{}
	}
	return RegularPricing{}
}

// =============================================================================
// Num02FactoryMethodPatternDemo
//
// TOPIC: Factory Method Pattern - Defer object creation to factory functions
//   - stub — implement when teaching
// =============================================================================
func Num02FactoryMethodPatternDemo() {
	// stub — implement when teaching
}

// =============================================================================
// Num03AbstractFactoryPatternDemo
//
// TOPIC: Abstract Factory Pattern - Create related object families consistently
//   - stub — implement when teaching
// =============================================================================
func Num03AbstractFactoryPatternDemo() {
	// stub — implement when teaching
}

// =============================================================================
// Num04BuilderPatternDemo
//
// TOPIC: Builder Pattern - Build complex objects step-by-step safely
//   - stub — implement when teaching
// =============================================================================
func Num04BuilderPatternDemo() {
	// stub — implement when teaching
}

// =============================================================================
// Num05SingletonPatternDemo
//
// TOPIC: Singleton Pattern - One instance with global access
//   - stub — implement when teaching
// =============================================================================
func Num05SingletonPatternDemo() {
	// stub — implement when teaching
}

// =============================================================================
// Num06AdapterPatternDemo
//
// TOPIC: Adapter Pattern - Make incompatible interfaces work together
//   - stub — implement when teaching
// =============================================================================
func Num06AdapterPatternDemo() {
	// stub — implement when teaching
}

// =============================================================================
// Num07FacadePatternDemo
//
// TOPIC: Facade Pattern - Simplify complex subsystem usage
//   - stub — implement when teaching
// =============================================================================
func Num07FacadePatternDemo() {
	// stub — implement when teaching
}

// =============================================================================
// Num08DecoratorPatternDemo
//
// TOPIC: Decorator Pattern - Add behavior dynamically without subclassing
//   - stub — implement when teaching
// =============================================================================
func Num08DecoratorPatternDemo() {
	// stub — implement when teaching
}

// =============================================================================
// Num09ProxyPatternDemo
//
// TOPIC: Proxy Pattern - Intercept access for caching, auth, or logging
//   - stub — implement when teaching
// =============================================================================
func Num09ProxyPatternDemo() {
	// stub — implement when teaching
}

// =============================================================================
// Num10ObserverPatternDemo
//
// TOPIC: Observer Pattern - Publish changes to multiple subscribers
//   - stub — implement when teaching
// =============================================================================
func Num10ObserverPatternDemo() {
	// stub — implement when teaching
}

// =============================================================================
// Num11ChainOfResponsibilityDemo
//
// TOPIC: Chain of Responsibility - Pass request through handlers
//   - stub — implement when teaching
// =============================================================================
func Num11ChainOfResponsibilityDemo() {
	// stub — implement when teaching
}

// =============================================================================
// Num12CommandPatternDemo
//
// TOPIC: Command Pattern - Encapsulate operations as executable objects
//   - stub — implement when teaching
// =============================================================================
func Num12CommandPatternDemo() {
	// stub — implement when teaching
}

// =============================================================================
// Num13StatePatternDemo
//
// TOPIC: State Pattern - Behavior changes with internal state transitions
//   - stub — implement when teaching
// =============================================================================
func Num13StatePatternDemo() {
	// stub — implement when teaching
}

// =============================================================================
// Num14TemplateMethodPatternDemo
//
// TOPIC: Template Method Pattern - Fixed workflow with overridable steps
//   - stub — implement when teaching
// =============================================================================
func Num14TemplateMethodPatternDemo() {
	// stub — implement when teaching
}

// =============================================================================
// Num15RepositoryPatternDemo
//
// TOPIC: Repository Pattern - Decouple domain logic from data persistence
//   - stub — implement when teaching
// =============================================================================
func Num15RepositoryPatternDemo() {
	// stub — implement when teaching
}

// =============================================================================
// Num16DependencyInjectionDemo
//
// TOPIC: Dependency Injection - Wire dependencies explicitly for testability
//   - stub — implement when teaching
// =============================================================================
func Num16DependencyInjectionDemo() {
	// stub — implement when teaching
}

// =============================================================================
// Num17CircuitBreakerPatternDemo
//
// TOPIC: Circuit Breaker Pattern - Stop cascading failures on unstable services
//   - stub — implement when teaching
// =============================================================================
func Num17CircuitBreakerPatternDemo() {
	// stub — implement when teaching
}

// =============================================================================
// Num18PubSubEventBusDemo
//
// TOPIC: Pub/Sub Event Bus - Decouple producers and consumers via events
//   - stub — implement when teaching
// =============================================================================
func Num18PubSubEventBusDemo() {
	// stub — implement when teaching
}

// =============================================================================
// Num19PipelinePatternDemo
//
// TOPIC: Pipeline Pattern - Compose multi-stage processing flow
//   - stub — implement when teaching
// =============================================================================
func Num19PipelinePatternDemo() {
	// stub — implement when teaching
}

// =============================================================================
// Num20SagaOrchestrationDemo
//
// TOPIC: Saga Orchestration - Distributed transaction with compensating actions
//   - stub — implement when teaching
// =============================================================================
func Num20SagaOrchestrationDemo() {
	// stub — implement when teaching
}
