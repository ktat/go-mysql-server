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
	_ = x[validateAlterTableId-4]
	_ = x[validateExprSemId-5]
	_ = x[resolveVariablesId-6]
	_ = x[resolveNamedWindowsId-7]
	_ = x[resolveSetVariablesId-8]
	_ = x[resolveViewsId-9]
	_ = x[liftCtesId-10]
	_ = x[resolveCtesId-11]
	_ = x[liftRecursiveCtesId-12]
	_ = x[resolveDatabasesId-13]
	_ = x[resolveTablesId-14]
	_ = x[loadStoredProceduresId-15]
	_ = x[validateDropTablesId-16]
	_ = x[pruneDropTablesId-17]
	_ = x[setTargetSchemasId-18]
	_ = x[resolveCreateLikeId-19]
	_ = x[parseColumnDefaultsId-20]
	_ = x[resolveDropConstraintId-21]
	_ = x[validateDropConstraintId-22]
	_ = x[loadCheckConstraintsId-23]
	_ = x[assignCatalogId-24]
	_ = x[resolveAnalyzeTablesId-25]
	_ = x[resolveCreateSelectId-26]
	_ = x[resolveSubqueriesId-27]
	_ = x[setViewTargetSchemaId-28]
	_ = x[resolveUnionsId-29]
	_ = x[resolveDescribeQueryId-30]
	_ = x[checkUniqueTableNamesId-31]
	_ = x[resolveTableFunctionsId-32]
	_ = x[resolveDeclarationsId-33]
	_ = x[resolveColumnDefaultsId-34]
	_ = x[validateColumnDefaultsId-35]
	_ = x[validateCreateTriggerId-36]
	_ = x[validateCreateProcedureId-37]
	_ = x[resolveCreateProcedureId-38]
	_ = x[loadInfoSchemaId-39]
	_ = x[validateReadOnlyDatabaseId-40]
	_ = x[validateReadOnlyTransactionId-41]
	_ = x[validateDatabaseSetId-42]
	_ = x[validatePrivilegesId-43]
	_ = x[reresolveTablesId-44]
	_ = x[setInsertColumnsId-45]
	_ = x[validateJoinComplexityId-46]
	_ = x[applyBinlogReplicaControllerId-47]
	_ = x[applyEventSchedulerId-48]
	_ = x[resolveUsingJoinsId-49]
	_ = x[resolveOrderbyLiteralsId-50]
	_ = x[resolveFunctionsId-51]
	_ = x[flattenTableAliasesId-52]
	_ = x[pushdownSortId-53]
	_ = x[pushdownGroupbyAliasesId-54]
	_ = x[pushdownSubqueryAliasFiltersId-55]
	_ = x[qualifyColumnsId-56]
	_ = x[resolveColumnsId-57]
	_ = x[validateCheckConstraintId-58]
	_ = x[resolveBarewordSetVariablesId-59]
	_ = x[replaceCountStarId-60]
	_ = x[expandStarsId-61]
	_ = x[transposeRightJoinsId-62]
	_ = x[resolveHavingId-63]
	_ = x[mergeUnionSchemasId-64]
	_ = x[flattenAggregationExprsId-65]
	_ = x[reorderProjectionId-66]
	_ = x[resolveSubqueryExprsId-67]
	_ = x[replaceCrossJoinsId-68]
	_ = x[moveJoinCondsToFilterId-69]
	_ = x[evalFilterId-70]
	_ = x[optimizeDistinctId-71]
	_ = x[hoistOutOfScopeFiltersId-72]
	_ = x[transformJoinApplyId-73]
	_ = x[hoistSelectExistsId-74]
	_ = x[finalizeSubqueriesId-75]
	_ = x[finalizeUnionsId-76]
	_ = x[loadTriggersId-77]
	_ = x[loadEventsId-78]
	_ = x[processTruncateId-79]
	_ = x[resolveAlterColumnId-80]
	_ = x[resolveGeneratorsId-81]
	_ = x[removeUnnecessaryConvertsId-82]
	_ = x[pruneColumnsId-83]
	_ = x[stripTableNameInDefaultsId-84]
	_ = x[foldEmptyJoinsId-85]
	_ = x[optimizeJoinsId-86]
	_ = x[generateIndexScansId-87]
	_ = x[matchAgainstId-88]
	_ = x[pushFiltersId-89]
	_ = x[applyIndexesFromOuterScopeId-90]
	_ = x[pruneTablesId-91]
	_ = x[fixupAuxiliaryExprsId-92]
	_ = x[assignExecIndexesId-93]
	_ = x[inlineSubqueryAliasRefsId-94]
	_ = x[eraseProjectionId-95]
	_ = x[replaceAggId-96]
	_ = x[replaceSortPkId-97]
	_ = x[insertTopNId-98]
	_ = x[applyHashInId-99]
	_ = x[resolveInsertRowsId-100]
	_ = x[resolvePreparedInsertId-101]
	_ = x[applyTriggersId-102]
	_ = x[applyProceduresId-103]
	_ = x[assignRoutinesId-104]
	_ = x[modifyUpdateExprsForJoinId-105]
	_ = x[applyRowUpdateAccumulatorsId-106]
	_ = x[wrapWithRollbackId-107]
	_ = x[applyFKsId-108]
	_ = x[validateResolvedId-109]
	_ = x[validateOrderById-110]
	_ = x[validateGroupById-111]
	_ = x[validateSchemaSourceId-112]
	_ = x[validateIndexCreationId-113]
	_ = x[validateOperandsId-114]
	_ = x[validateCaseResultTypesId-115]
	_ = x[validateIntervalUsageId-116]
	_ = x[validateExplodeUsageId-117]
	_ = x[validateSubqueryColumnsId-118]
	_ = x[validateUnionSchemasMatchId-119]
	_ = x[validateAggregationsId-120]
	_ = x[validateDeleteFromId-121]
	_ = x[cacheSubqueryResultsId-122]
	_ = x[cacheSubqueryAliasesInJoinsId-123]
	_ = x[AutocommitId-124]
	_ = x[TrackProcessId-125]
	_ = x[parallelizeId-126]
	_ = x[clearWarningsId-127]
}

const _RuleId_name = "applyDefaultSelectLimitvalidateOffsetAndLimitvalidateStarExpressionsvalidateCreateTablevalidateAlterTablevalidateExprSemresolveVariablesresolveNamedWindowsresolveSetVariablesresolveViewsliftCtesresolveCtesliftRecursiveCtesresolveDatabasesresolveTablesloadStoredProceduresvalidateDropTablespruneDropTablessetTargetSchemasresolveCreateLikeparseColumnDefaultsresolveDropConstraintvalidateDropConstraintloadCheckConstraintsassignCatalogresolveAnalyzeTablesresolveCreateSelectresolveSubqueriessetViewTargetSchemaresolveUnionsresolveDescribeQuerycheckUniqueTableNamesresolveTableFunctionsresolveDeclarationsresolveColumnDefaultsvalidateColumnDefaultsvalidateCreateTriggervalidateCreateProcedureresolveCreateProcedureloadInfoSchemavalidateReadOnlyDatabasevalidateReadOnlyTransactionvalidateDatabaseSetvalidatePrivilegesreresolveTablessetInsertColumnsvalidateJoinComplexityapplyBinlogReplicaControllerapplyEventSchedulerresolveUsingJoinsresolveOrderbyLiteralsresolveFunctionsflattenTableAliasespushdownSortpushdownGroupbyAliasespushdownSubqueryAliasFiltersqualifyColumnsresolveColumnsvalidateCheckConstraintresolveBarewordSetVariablesreplaceCountStarexpandStarstransposeRightJoinsresolveHavingmergeUnionSchemasflattenAggregationExprsreorderProjectionresolveSubqueryExprsreplaceCrossJoinsmoveJoinCondsToFilterevalFilteroptimizeDistincthoistOutOfScopeFilterstransformJoinApplyhoistSelectExistsfinalizeSubqueriesfinalizeUnionsloadTriggersloadEventsprocessTruncateresolveAlterColumnresolveGeneratorsremoveUnnecessaryConvertspruneColumnsstripTableNamesFromColumnDefaultsfoldEmptyJoinsoptimizeJoinsgenerateIndexScansmatchAgainstpushFiltersapplyIndexesFromOuterScopepruneTablesfixupAuxiliaryExprsassignExecIndexesinlineSubqueryAliasRefseraseProjectionreplaceAggreplaceSortPkinsertTopNapplyHashInresolveInsertRowsresolvePreparedInsertapplyTriggersapplyProceduresassignRoutinesmodifyUpdateExprsForJoinapplyRowUpdateAccumulatorsrollback triggersapplyFKsvalidateResolvedvalidateOrderByvalidateGroupByvalidateSchemaSourcevalidateIndexCreationvalidateOperandsvalidateCaseResultTypesvalidateIntervalUsagevalidateExplodeUsagevalidateSubqueryColumnsvalidateUnionSchemasMatchvalidateAggregationsvalidateDeleteFromcacheSubqueryResultscacheSubqueryAliasesInJoinsaddAutocommitNodetrackProcessparallelizeclearWarnings"

var _RuleId_index = [...]uint16{0, 23, 45, 68, 87, 105, 120, 136, 155, 174, 186, 194, 205, 222, 238, 251, 271, 289, 304, 320, 337, 356, 377, 399, 419, 432, 452, 471, 488, 507, 520, 540, 561, 582, 601, 622, 644, 665, 688, 710, 724, 748, 775, 794, 812, 827, 843, 865, 893, 912, 929, 951, 967, 986, 998, 1020, 1048, 1062, 1076, 1099, 1126, 1142, 1153, 1172, 1185, 1202, 1225, 1242, 1262, 1279, 1300, 1310, 1326, 1348, 1366, 1383, 1401, 1415, 1427, 1437, 1452, 1470, 1487, 1512, 1524, 1557, 1571, 1584, 1602, 1614, 1625, 1651, 1662, 1681, 1698, 1721, 1736, 1746, 1759, 1769, 1780, 1797, 1818, 1831, 1846, 1860, 1884, 1910, 1927, 1935, 1951, 1966, 1981, 2001, 2022, 2038, 2061, 2082, 2102, 2125, 2150, 2170, 2188, 2208, 2235, 2252, 2264, 2275, 2288}

func (i RuleId) String() string {
	if i < 0 || i >= RuleId(len(_RuleId_index)-1) {
		return "RuleId(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RuleId_name[_RuleId_index[i]:_RuleId_index[i+1]]
}
