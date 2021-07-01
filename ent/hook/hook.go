// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/a20070322/shop-go/ent"
)

// The BasicBannerFunc type is an adapter to allow the use of ordinary
// function as BasicBanner mutator.
type BasicBannerFunc func(context.Context, *ent.BasicBannerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BasicBannerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.BasicBannerMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BasicBannerMutation", m)
	}
	return f(ctx, mv)
}

// The BasicBannerPositionFunc type is an adapter to allow the use of ordinary
// function as BasicBannerPosition mutator.
type BasicBannerPositionFunc func(context.Context, *ent.BasicBannerPositionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BasicBannerPositionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.BasicBannerPositionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BasicBannerPositionMutation", m)
	}
	return f(ctx, mv)
}

// The BasicLinkFunc type is an adapter to allow the use of ordinary
// function as BasicLink mutator.
type BasicLinkFunc func(context.Context, *ent.BasicLinkMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BasicLinkFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.BasicLinkMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BasicLinkMutation", m)
	}
	return f(ctx, mv)
}

// The CustomerFunc type is an adapter to allow the use of ordinary
// function as Customer mutator.
type CustomerFunc func(context.Context, *ent.CustomerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CustomerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CustomerMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CustomerMutation", m)
	}
	return f(ctx, mv)
}

// The CustomerAddressFunc type is an adapter to allow the use of ordinary
// function as CustomerAddress mutator.
type CustomerAddressFunc func(context.Context, *ent.CustomerAddressMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CustomerAddressFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CustomerAddressMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CustomerAddressMutation", m)
	}
	return f(ctx, mv)
}

// The GoodsClassifyFunc type is an adapter to allow the use of ordinary
// function as GoodsClassify mutator.
type GoodsClassifyFunc func(context.Context, *ent.GoodsClassifyMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GoodsClassifyFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GoodsClassifyMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GoodsClassifyMutation", m)
	}
	return f(ctx, mv)
}

// The GoodsSkuFunc type is an adapter to allow the use of ordinary
// function as GoodsSku mutator.
type GoodsSkuFunc func(context.Context, *ent.GoodsSkuMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GoodsSkuFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GoodsSkuMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GoodsSkuMutation", m)
	}
	return f(ctx, mv)
}

// The GoodsSpecsFunc type is an adapter to allow the use of ordinary
// function as GoodsSpecs mutator.
type GoodsSpecsFunc func(context.Context, *ent.GoodsSpecsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GoodsSpecsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GoodsSpecsMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GoodsSpecsMutation", m)
	}
	return f(ctx, mv)
}

// The GoodsSpecsOptionFunc type is an adapter to allow the use of ordinary
// function as GoodsSpecsOption mutator.
type GoodsSpecsOptionFunc func(context.Context, *ent.GoodsSpecsOptionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GoodsSpecsOptionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GoodsSpecsOptionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GoodsSpecsOptionMutation", m)
	}
	return f(ctx, mv)
}

// The GoodsSpuFunc type is an adapter to allow the use of ordinary
// function as GoodsSpu mutator.
type GoodsSpuFunc func(context.Context, *ent.GoodsSpuMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GoodsSpuFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GoodsSpuMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GoodsSpuMutation", m)
	}
	return f(ctx, mv)
}

// The GoodsSpuImgsFunc type is an adapter to allow the use of ordinary
// function as GoodsSpuImgs mutator.
type GoodsSpuImgsFunc func(context.Context, *ent.GoodsSpuImgsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GoodsSpuImgsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GoodsSpuImgsMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GoodsSpuImgsMutation", m)
	}
	return f(ctx, mv)
}

// The OrderAddressFunc type is an adapter to allow the use of ordinary
// function as OrderAddress mutator.
type OrderAddressFunc func(context.Context, *ent.OrderAddressMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OrderAddressFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.OrderAddressMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OrderAddressMutation", m)
	}
	return f(ctx, mv)
}

// The OrderGoodsSkuFunc type is an adapter to allow the use of ordinary
// function as OrderGoodsSku mutator.
type OrderGoodsSkuFunc func(context.Context, *ent.OrderGoodsSkuMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OrderGoodsSkuFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.OrderGoodsSkuMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OrderGoodsSkuMutation", m)
	}
	return f(ctx, mv)
}

// The OrderInfoFunc type is an adapter to allow the use of ordinary
// function as OrderInfo mutator.
type OrderInfoFunc func(context.Context, *ent.OrderInfoMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OrderInfoFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.OrderInfoMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OrderInfoMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
//
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}