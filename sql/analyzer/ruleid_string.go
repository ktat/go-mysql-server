// Code generated by "stringer -type=RuleId -linecomment"; DO NOT EDIT.

package analyzer

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[applyDefaultSelectLimitId-0]
	_ = x[validateOffsetAndLimitId-1]
	_ = x[validateStarExpressionsId-2]
	_ = x[validateCreateTableId-3]
	_ = x[validateExprSemId-4]
	_ = x[resolveVariablesId-5]
	_ = x[resolveNamedWindowsId-6]
	_ = x[resolveSetVariablesId-7]
	_ = x[resolveViewsId-8]
	_ = x[liftCtesId-9]
	_ = x[resolveCtesId-10]
	_ = x[liftRecursiveCtesId-11]
	_ = x[resolveDatabasesId-12]
	_ = x[resolveTablesId-13]
	_ = x[loadStoredProceduresId-14]
	_ = x[validateDropTablesId-15]
	_ = x[pruneDropTablesId-16]
	_ = x[setTargetSchemasId-17]
	_ = x[resolveCreateLikeId-18]
	_ = x[parseColumnDefaultsId-19]
	_ = x[resolveDropConstraintId-20]
	_ = x[validateDropConstraintId-21]
	_ = x[loadCheckConstraintsId-22]
	_ = x[assignCatalogId-23]
	_ = x[resolveAnalyzeTablesId-24]
	_ = x[resolveCreateSelectId-25]
	_ = x[resolveSubqueriesId-26]
	_ = x[setViewTargetSchemaId-27]
	_ = x[resolveUnionsId-28]
	_ = x[resolveDescribeQueryId-29]
	_ = x[checkUniqueTableNamesId-30]
	_ = x[resolveTableFunctionsId-31]
	_ = x[resolveDeclarationsId-32]
	_ = x[resolveColumnDefaultsId-33]
	_ = x[validateColumnDefaultsId-34]
	_ = x[validateCreateTriggerId-35]
	_ = x[validateCreateProcedureId-36]
	_ = x[resolveCreateProcedureId-37]
	_ = x[loadInfoSchemaId-38]
	_ = x[validateReadOnlyDatabaseId-39]
	_ = x[validateReadOnlyTransactionId-40]
	_ = x[validateDatabaseSetId-41]
	_ = x[validatePrivilegesId-42]
	_ = x[reresolveTablesId-43]
	_ = x[setInsertColumnsId-44]
	_ = x[validateJoinComplexityId-45]
	_ = x[applyBinlogReplicaControllerId-46]
	_ = x[resolveUsingJoinsId-47]
	_ = x[resolveOrderbyLiteralsId-48]
	_ = x[resolveFunctionsId-49]
	_ = x[flattenTableAliasesId-50]
	_ = x[pushdownSortId-51]
	_ = x[pushdownGroupbyAliasesId-52]
	_ = x[pushdownSubqueryAliasFiltersId-53]
	_ = x[qualifyColumnsId-54]
	_ = x[resolveColumnsId-55]
	_ = x[validateCheckConstraintId-56]
	_ = x[resolveBarewordSetVariablesId-57]
	_ = x[replaceCountStarId-58]
	_ = x[expandStarsId-59]
	_ = x[transposeRightJoinsId-60]
	_ = x[resolveHavingId-61]
	_ = x[mergeUnionSchemasId-62]
	_ = x[flattenAggregationExprsId-63]
	_ = x[reorderProjectionId-64]
	_ = x[resolveSubqueryExprsId-65]
	_ = x[replaceCrossJoinsId-66]
	_ = x[moveJoinCondsToFilterId-67]
	_ = x[evalFilterId-68]
	_ = x[optimizeDistinctId-69]
	_ = x[hoistOutOfScopeFiltersId-70]
	_ = x[transformJoinApplyId-71]
	_ = x[hoistSelectExistsId-72]
	_ = x[finalizeSubqueriesId-73]
	_ = x[finalizeUnionsId-74]
	_ = x[loadTriggersId-75]
	_ = x[loadEventsId-76]
	_ = x[processTruncateId-77]
	_ = x[resolveAlterColumnId-78]
	_ = x[resolveGeneratorsId-79]
	_ = x[removeUnnecessaryConvertsId-80]
	_ = x[pruneColumnsId-81]
	_ = x[stripTableNameInDefaultsId-82]
	_ = x[foldEmptyJoinsId-83]
	_ = x[optimizeJoinsId-84]
	_ = x[generateIndexScansId-85]
	_ = x[matchAgainstId-86]
	_ = x[pushFiltersId-87]
	_ = x[subqueryIndexesId-88]
	_ = x[pruneTablesId-89]
	_ = x[fixupAuxiliaryExprsId-90]
	_ = x[assignExecIndexesId-91]
	_ = x[inlineSubqueryAliasRefsId-92]
	_ = x[setJoinScopeLenId-93]
	_ = x[eraseProjectionId-94]
	_ = x[replaceSortPkId-95]
	_ = x[insertTopNId-96]
	_ = x[applyHashInId-97]
	_ = x[resolveInsertRowsId-98]
	_ = x[resolvePreparedInsertId-99]
	_ = x[applyTriggersId-100]
	_ = x[applyProceduresId-101]
	_ = x[assignRoutinesId-102]
	_ = x[modifyUpdateExprsForJoinId-103]
	_ = x[applyRowUpdateAccumulatorsId-104]
	_ = x[wrapWithRollbackId-105]
	_ = x[applyFKsId-106]
	_ = x[validateResolvedId-107]
	_ = x[validateOrderById-108]
	_ = x[validateGroupById-109]
	_ = x[validateSchemaSourceId-110]
	_ = x[validateIndexCreationId-111]
	_ = x[validateOperandsId-112]
	_ = x[validateCaseResultTypesId-113]
	_ = x[validateIntervalUsageId-114]
	_ = x[validateExplodeUsageId-115]
	_ = x[validateSubqueryColumnsId-116]
	_ = x[validateUnionSchemasMatchId-117]
	_ = x[validateAggregationsId-118]
	_ = x[validateDeleteFromId-119]
	_ = x[cacheSubqueryResultsId-120]
	_ = x[cacheSubqueryAliasesInJoinsId-121]
	_ = x[AutocommitId-122]
	_ = x[TrackProcessId-123]
	_ = x[parallelizeId-124]
	_ = x[clearWarningsId-125]
}

const _RuleId_name = "applyDefaultSelectLimitvalidateOffsetAndLimitvalidateStarExpressionsvalidateCreateTablevalidateExprSemresolveVariablesresolveNamedWindowsresolveSetVariablesresolveViewsliftCtesresolveCtesliftRecursiveCtesresolveDatabasesresolveTablesloadStoredProceduresvalidateDropTablespruneDropTablessetTargetSchemasresolveCreateLikeparseColumnDefaultsresolveDropConstraintvalidateDropConstraintloadCheckConstraintsassignCatalogresolveAnalyzeTablesresolveCreateSelectresolveSubqueriessetViewTargetSchemaresolveUnionsresolveDescribeQuerycheckUniqueTableNamesresolveTableFunctionsresolveDeclarationsresolveColumnDefaultsvalidateColumnDefaultsvalidateCreateTriggervalidateCreateProcedureresolveCreateProcedureloadInfoSchemavalidateReadOnlyDatabasevalidateReadOnlyTransactionvalidateDatabaseSetvalidatePrivilegesreresolveTablessetInsertColumnsvalidateJoinComplexityapplyBinlogReplicaControllerresolveUsingJoinsresolveOrderbyLiteralsresolveFunctionsflattenTableAliasespushdownSortpushdownGroupbyAliasespushdownSubqueryAliasFiltersqualifyColumnsresolveColumnsvalidateCheckConstraintresolveBarewordSetVariablesreplaceCountStarexpandStarstransposeRightJoinsresolveHavingmergeUnionSchemasflattenAggregationExprsreorderProjectionresolveSubqueryExprsreplaceCrossJoinsmoveJoinCondsToFilterevalFilteroptimizeDistincthoistOutOfScopeFilterstransformJoinApplyhoistSelectExistsfinalizeSubqueriesfinalizeUnionsloadTriggersloadEventsprocessTruncateresolveAlterColumnresolveGeneratorsremoveUnnecessaryConvertspruneColumnsstripTableNamesFromColumnDefaultsfoldEmptyJoinsoptimizeJoinsgenerateIndexScansmatchAgainstpushFilterssubqueryIndexespruneTablesfixupAuxiliaryExprsfixupIndexesinlineSubqueryAliasRefssetJoinScopeLeneraseProjectionreplaceSortPkinsertTopNapplyHashInresolveInsertRowsresolvePreparedInsertapplyTriggersapplyProceduresassignRoutinesmodifyUpdateExprsForJoinapplyRowUpdateAccumulatorsrollback triggersapplyFKsvalidateResolvedvalidateOrderByvalidateGroupByvalidateSchemaSourcevalidateIndexCreationvalidateOperandsvalidateCaseResultTypesvalidateIntervalUsagevalidateExplodeUsagevalidateSubqueryColumnsvalidateUnionSchemasMatchvalidateAggregationsvalidateDeleteFromcacheSubqueryResultscacheSubqueryAliasesInJoinsaddAutocommitNodetrackProcessparallelizeclearWarnings"

var _RuleId_index = [...]uint16{0, 23, 45, 68, 87, 102, 118, 137, 156, 168, 176, 187, 204, 220, 233, 253, 271, 286, 302, 319, 338, 359, 381, 401, 414, 434, 453, 470, 489, 502, 522, 543, 564, 583, 604, 626, 647, 670, 692, 706, 730, 757, 776, 794, 809, 825, 847, 875, 892, 914, 930, 949, 961, 983, 1011, 1025, 1039, 1062, 1089, 1105, 1116, 1135, 1148, 1165, 1188, 1205, 1225, 1242, 1263, 1273, 1289, 1311, 1329, 1346, 1364, 1378, 1390, 1400, 1415, 1433, 1450, 1475, 1487, 1520, 1534, 1547, 1565, 1577, 1588, 1603, 1614, 1633, 1645, 1668, 1683, 1698, 1711, 1721, 1732, 1749, 1770, 1783, 1798, 1812, 1836, 1862, 1879, 1887, 1903, 1918, 1933, 1953, 1974, 1990, 2013, 2034, 2054, 2077, 2102, 2122, 2140, 2160, 2187, 2204, 2216, 2227, 2240}

func (i RuleId) String() string {
	if i < 0 || i >= RuleId(len(_RuleId_index)-1) {
		return "RuleId(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RuleId_name[_RuleId_index[i]:_RuleId_index[i+1]]
}
