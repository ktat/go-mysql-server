// Code generated by "stringer -type=JoinType -linecomment"; DO NOT EDIT.

package plan

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[JoinTypeUnknown-0]
	_ = x[JoinTypeCross-1]
	_ = x[JoinTypeCrossHash-2]
	_ = x[JoinTypeInner-3]
	_ = x[JoinTypeSemi-4]
	_ = x[JoinTypeAnti-5]
	_ = x[JoinTypeLeftOuter-6]
	_ = x[JoinTypeLeftOuterExcludeNulls-7]
	_ = x[JoinTypeFullOuter-8]
	_ = x[JoinTypeGroupBy-9]
	_ = x[JoinTypeRightOuter-10]
	_ = x[JoinTypeLookup-11]
	_ = x[JoinTypeLeftOuterLookup-12]
	_ = x[JoinTypeHash-13]
	_ = x[JoinTypeLeftOuterHash-14]
	_ = x[JoinTypeLeftOuterHashExcludeNulls-15]
	_ = x[JoinTypeMerge-16]
	_ = x[JoinTypeLeftOuterMerge-17]
	_ = x[JoinTypeRangeHeap-18]
	_ = x[JoinTypeLeftOuterRangeHeap-19]
	_ = x[JoinTypeSemiHash-20]
	_ = x[JoinTypeAntiHash-21]
	_ = x[JoinTypeSemiLookup-22]
	_ = x[JoinTypeAntiLookup-23]
	_ = x[JoinTypeSemiMerge-24]
	_ = x[JoinTypeAntiMerge-25]
	_ = x[JoinTypeUsing-26]
	_ = x[JoinTypeUsingLeft-27]
	_ = x[JoinTypeUsingRight-28]
	_ = x[JoinTypeLateralCross-29]
	_ = x[JoinTypeLateralInner-30]
	_ = x[JoinTypeLateralLeft-31]
	_ = x[JoinTypeLateralRight-32]
}

const _JoinType_name = "UnknownJoinCrossJoinCrossHashJoinInnerJoinSemiJoinAntiJoinLeftOuterJoinLeftOuterJoinExcludingNullsFullOuterJoinGroupByJoinRightJoinLookupJoinLeftOuterLookupJoinHashJoinLeftOuterHashJoinLeftOuterHashJoinExcludeNullsMergeJoinLeftOuterMergeJoinRangeHeapJoinLeftOuterRangeHeapJoinSemiHashJoinAntiHashJoinSemiLookupJoinAntiLookupJoinSemiMergeJoinAntiMergeJoinNaturalJoinNaturalLeftJoinNaturalRightJoinLateralCrossJoinLateralInnerJoinLateralLeftJoinLateralLeftJoin"

var _JoinType_index = [...]uint16{0, 11, 20, 33, 42, 50, 58, 71, 98, 111, 122, 131, 141, 160, 168, 185, 214, 223, 241, 254, 276, 288, 300, 314, 328, 341, 354, 365, 380, 396, 412, 428, 443, 458}

func (i JoinType) String() string {
	if i >= JoinType(len(_JoinType_index)-1) {
		return "JoinType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _JoinType_name[_JoinType_index[i]:_JoinType_index[i+1]]
}
