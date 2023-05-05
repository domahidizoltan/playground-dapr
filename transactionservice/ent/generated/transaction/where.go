// Code generated by ent, DO NOT EDIT.

package transaction

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/domahidizoltan/playground-dapr/transactionservice/ent/generated/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldID, id))
}

// Tnx applies equality check predicate on the "tnx" field. It's identical to TnxEQ.
func Tnx(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldTnx, v))
}

// SourceAcc applies equality check predicate on the "source_acc" field. It's identical to SourceAccEQ.
func SourceAcc(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldSourceAcc, v))
}

// DestAcc applies equality check predicate on the "dest_acc" field. It's identical to DestAccEQ.
func DestAcc(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldDestAcc, v))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldAmount, v))
}

// Datetime applies equality check predicate on the "datetime" field. It's identical to DatetimeEQ.
func Datetime(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldDatetime, v))
}

// TnxEQ applies the EQ predicate on the "tnx" field.
func TnxEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldTnx, v))
}

// TnxNEQ applies the NEQ predicate on the "tnx" field.
func TnxNEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldTnx, v))
}

// TnxIn applies the In predicate on the "tnx" field.
func TnxIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldTnx, vs...))
}

// TnxNotIn applies the NotIn predicate on the "tnx" field.
func TnxNotIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldTnx, vs...))
}

// TnxGT applies the GT predicate on the "tnx" field.
func TnxGT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldTnx, v))
}

// TnxGTE applies the GTE predicate on the "tnx" field.
func TnxGTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldTnx, v))
}

// TnxLT applies the LT predicate on the "tnx" field.
func TnxLT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldTnx, v))
}

// TnxLTE applies the LTE predicate on the "tnx" field.
func TnxLTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldTnx, v))
}

// TnxContains applies the Contains predicate on the "tnx" field.
func TnxContains(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContains(FieldTnx, v))
}

// TnxHasPrefix applies the HasPrefix predicate on the "tnx" field.
func TnxHasPrefix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasPrefix(FieldTnx, v))
}

// TnxHasSuffix applies the HasSuffix predicate on the "tnx" field.
func TnxHasSuffix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasSuffix(FieldTnx, v))
}

// TnxEqualFold applies the EqualFold predicate on the "tnx" field.
func TnxEqualFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEqualFold(FieldTnx, v))
}

// TnxContainsFold applies the ContainsFold predicate on the "tnx" field.
func TnxContainsFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContainsFold(FieldTnx, v))
}

// SourceAccEQ applies the EQ predicate on the "source_acc" field.
func SourceAccEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldSourceAcc, v))
}

// SourceAccNEQ applies the NEQ predicate on the "source_acc" field.
func SourceAccNEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldSourceAcc, v))
}

// SourceAccIn applies the In predicate on the "source_acc" field.
func SourceAccIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldSourceAcc, vs...))
}

// SourceAccNotIn applies the NotIn predicate on the "source_acc" field.
func SourceAccNotIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldSourceAcc, vs...))
}

// SourceAccGT applies the GT predicate on the "source_acc" field.
func SourceAccGT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldSourceAcc, v))
}

// SourceAccGTE applies the GTE predicate on the "source_acc" field.
func SourceAccGTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldSourceAcc, v))
}

// SourceAccLT applies the LT predicate on the "source_acc" field.
func SourceAccLT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldSourceAcc, v))
}

// SourceAccLTE applies the LTE predicate on the "source_acc" field.
func SourceAccLTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldSourceAcc, v))
}

// SourceAccContains applies the Contains predicate on the "source_acc" field.
func SourceAccContains(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContains(FieldSourceAcc, v))
}

// SourceAccHasPrefix applies the HasPrefix predicate on the "source_acc" field.
func SourceAccHasPrefix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasPrefix(FieldSourceAcc, v))
}

// SourceAccHasSuffix applies the HasSuffix predicate on the "source_acc" field.
func SourceAccHasSuffix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasSuffix(FieldSourceAcc, v))
}

// SourceAccEqualFold applies the EqualFold predicate on the "source_acc" field.
func SourceAccEqualFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEqualFold(FieldSourceAcc, v))
}

// SourceAccContainsFold applies the ContainsFold predicate on the "source_acc" field.
func SourceAccContainsFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContainsFold(FieldSourceAcc, v))
}

// DestAccEQ applies the EQ predicate on the "dest_acc" field.
func DestAccEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldDestAcc, v))
}

// DestAccNEQ applies the NEQ predicate on the "dest_acc" field.
func DestAccNEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldDestAcc, v))
}

// DestAccIn applies the In predicate on the "dest_acc" field.
func DestAccIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldDestAcc, vs...))
}

// DestAccNotIn applies the NotIn predicate on the "dest_acc" field.
func DestAccNotIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldDestAcc, vs...))
}

// DestAccGT applies the GT predicate on the "dest_acc" field.
func DestAccGT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldDestAcc, v))
}

// DestAccGTE applies the GTE predicate on the "dest_acc" field.
func DestAccGTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldDestAcc, v))
}

// DestAccLT applies the LT predicate on the "dest_acc" field.
func DestAccLT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldDestAcc, v))
}

// DestAccLTE applies the LTE predicate on the "dest_acc" field.
func DestAccLTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldDestAcc, v))
}

// DestAccContains applies the Contains predicate on the "dest_acc" field.
func DestAccContains(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContains(FieldDestAcc, v))
}

// DestAccHasPrefix applies the HasPrefix predicate on the "dest_acc" field.
func DestAccHasPrefix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasPrefix(FieldDestAcc, v))
}

// DestAccHasSuffix applies the HasSuffix predicate on the "dest_acc" field.
func DestAccHasSuffix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasSuffix(FieldDestAcc, v))
}

// DestAccEqualFold applies the EqualFold predicate on the "dest_acc" field.
func DestAccEqualFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEqualFold(FieldDestAcc, v))
}

// DestAccContainsFold applies the ContainsFold predicate on the "dest_acc" field.
func DestAccContainsFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContainsFold(FieldDestAcc, v))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldAmount, v))
}

// DatetimeEQ applies the EQ predicate on the "datetime" field.
func DatetimeEQ(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldDatetime, v))
}

// DatetimeNEQ applies the NEQ predicate on the "datetime" field.
func DatetimeNEQ(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldDatetime, v))
}

// DatetimeIn applies the In predicate on the "datetime" field.
func DatetimeIn(vs ...time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldDatetime, vs...))
}

// DatetimeNotIn applies the NotIn predicate on the "datetime" field.
func DatetimeNotIn(vs ...time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldDatetime, vs...))
}

// DatetimeGT applies the GT predicate on the "datetime" field.
func DatetimeGT(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldDatetime, v))
}

// DatetimeGTE applies the GTE predicate on the "datetime" field.
func DatetimeGTE(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldDatetime, v))
}

// DatetimeLT applies the LT predicate on the "datetime" field.
func DatetimeLT(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldDatetime, v))
}

// DatetimeLTE applies the LTE predicate on the "datetime" field.
func DatetimeLTE(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldDatetime, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Transaction) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Transaction) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Transaction) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		p(s.Not())
	})
}