// Code generated by plangen.

// Copyright 2024 Dolthub, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package queries

var SysbenchPlanTests = []QueryPlanTest{
	{
		Query: `select a.id, a.small_int_col from sbtest1 a, sbtest1 b where a.id = b.id limit 500`,
		ExpectedPlan: "Limit(500)\n" +
			" └─ Project\n" +
			"     ├─ columns: [a.id:1!null, a.small_int_col:2!null]\n" +
			"     └─ MergeJoin\n" +
			"         ├─ cmp: Eq\n" +
			"         │   ├─ b.id:0!null\n" +
			"         │   └─ a.id:1!null\n" +
			"         ├─ TableAlias(b)\n" +
			"         │   └─ IndexedTableAccess(sbtest1)\n" +
			"         │       ├─ index: [sbtest1.id]\n" +
			"         │       ├─ static: [{[NULL, ∞)}]\n" +
			"         │       ├─ colSet: (25-48)\n" +
			"         │       ├─ tableId: 2\n" +
			"         │       └─ Table\n" +
			"         │           ├─ name: sbtest1\n" +
			"         │           └─ columns: [id]\n" +
			"         └─ TableAlias(a)\n" +
			"             └─ IndexedTableAccess(sbtest1)\n" +
			"                 ├─ index: [sbtest1.id]\n" +
			"                 ├─ static: [{[NULL, ∞)}]\n" +
			"                 ├─ colSet: (1-24)\n" +
			"                 ├─ tableId: 1\n" +
			"                 └─ Table\n" +
			"                     ├─ name: sbtest1\n" +
			"                     └─ columns: [id small_int_col]\n" +
			"",
		ExpectedEstimates: "Limit(500)\n" +
			" └─ Project\n" +
			"     ├─ columns: [a.id:1!null, a.small_int_col:2!null]\n" +
			"     └─ MergeJoin\n" +
			"         ├─ cmp: Eq\n" +
			"         │   ├─ b.id:0!null\n" +
			"         │   └─ a.id:1!null\n" +
			"         ├─ TableAlias(b)\n" +
			"         │   └─ IndexedTableAccess(sbtest1)\n" +
			"         │       ├─ index: [sbtest1.id]\n" +
			"         │       ├─ static: [{[NULL, ∞)}]\n" +
			"         │       ├─ colSet: (25-48)\n" +
			"         │       ├─ tableId: 2\n" +
			"         │       └─ Table\n" +
			"         │           ├─ name: sbtest1\n" +
			"         │           └─ columns: [id]\n" +
			"         └─ TableAlias(a)\n" +
			"             └─ IndexedTableAccess(sbtest1)\n" +
			"                 ├─ index: [sbtest1.id]\n" +
			"                 ├─ static: [{[NULL, ∞)}]\n" +
			"                 ├─ colSet: (1-24)\n" +
			"                 ├─ tableId: 1\n" +
			"                 └─ Table\n" +
			"                     ├─ name: sbtest1\n" +
			"                     └─ columns: [id small_int_col]\n" +
			"",
		ExpectedAnalysis: "Limit(500)\n" +
			" └─ Project\n" +
			"     ├─ columns: [a.id:1!null, a.small_int_col:2!null]\n" +
			"     └─ MergeJoin\n" +
			"         ├─ cmp: Eq\n" +
			"         │   ├─ b.id:0!null\n" +
			"         │   └─ a.id:1!null\n" +
			"         ├─ TableAlias(b)\n" +
			"         │   └─ IndexedTableAccess(sbtest1)\n" +
			"         │       ├─ index: [sbtest1.id]\n" +
			"         │       ├─ static: [{[NULL, ∞)}]\n" +
			"         │       ├─ colSet: (25-48)\n" +
			"         │       ├─ tableId: 2\n" +
			"         │       └─ Table\n" +
			"         │           ├─ name: sbtest1\n" +
			"         │           └─ columns: [id]\n" +
			"         └─ TableAlias(a)\n" +
			"             └─ IndexedTableAccess(sbtest1)\n" +
			"                 ├─ index: [sbtest1.id]\n" +
			"                 ├─ static: [{[NULL, ∞)}]\n" +
			"                 ├─ colSet: (1-24)\n" +
			"                 ├─ tableId: 1\n" +
			"                 └─ Table\n" +
			"                     ├─ name: sbtest1\n" +
			"                     └─ columns: [id small_int_col]\n" +
			"",
	},
	{
		Query: `select a.id, a.small_int_col, b.id, b.int_col from sbtest1 a, sbtest2 b where a.id = b.int_col limit 500`,
		ExpectedPlan: "Limit(500)\n" +
			" └─ Project\n" +
			"     ├─ columns: [a.id:2!null, a.small_int_col:3!null, b.id:0!null, b.int_col:1!null]\n" +
			"     └─ LookupJoin\n" +
			"         ├─ TableAlias(b)\n" +
			"         │   └─ ProcessTable\n" +
			"         │       └─ Table\n" +
			"         │           ├─ name: sbtest2\n" +
			"         │           └─ columns: [id int_col]\n" +
			"         └─ TableAlias(a)\n" +
			"             └─ IndexedTableAccess(sbtest1)\n" +
			"                 ├─ index: [sbtest1.id]\n" +
			"                 ├─ keys: [b.int_col:1!null]\n" +
			"                 ├─ colSet: (1-24)\n" +
			"                 ├─ tableId: 1\n" +
			"                 └─ Table\n" +
			"                     ├─ name: sbtest1\n" +
			"                     └─ columns: [id small_int_col]\n" +
			"",
		ExpectedEstimates: "Limit(500)\n" +
			" └─ Project\n" +
			"     ├─ columns: [a.id:2!null, a.small_int_col:3!null, b.id:0!null, b.int_col:1!null]\n" +
			"     └─ LookupJoin\n" +
			"         ├─ TableAlias(b)\n" +
			"         │   └─ ProcessTable\n" +
			"         │       └─ Table\n" +
			"         │           ├─ name: sbtest2\n" +
			"         │           └─ columns: [id int_col]\n" +
			"         └─ TableAlias(a)\n" +
			"             └─ IndexedTableAccess(sbtest1)\n" +
			"                 ├─ index: [sbtest1.id]\n" +
			"                 ├─ keys: [b.int_col:1!null]\n" +
			"                 ├─ colSet: (1-24)\n" +
			"                 ├─ tableId: 1\n" +
			"                 └─ Table\n" +
			"                     ├─ name: sbtest1\n" +
			"                     └─ columns: [id small_int_col]\n" +
			"",
		ExpectedAnalysis: "Limit(500)\n" +
			" └─ Project\n" +
			"     ├─ columns: [a.id:2!null, a.small_int_col:3!null, b.id:0!null, b.int_col:1!null]\n" +
			"     └─ LookupJoin\n" +
			"         ├─ TableAlias(b)\n" +
			"         │   └─ ProcessTable\n" +
			"         │       └─ Table\n" +
			"         │           ├─ name: sbtest2\n" +
			"         │           └─ columns: [id int_col]\n" +
			"         └─ TableAlias(a)\n" +
			"             └─ IndexedTableAccess(sbtest1)\n" +
			"                 ├─ index: [sbtest1.id]\n" +
			"                 ├─ keys: [b.int_col:1!null]\n" +
			"                 ├─ colSet: (1-24)\n" +
			"                 ├─ tableId: 1\n" +
			"                 └─ Table\n" +
			"                     ├─ name: sbtest1\n" +
			"                     └─ columns: [id small_int_col]\n" +
			"",
	},
	{
		Query: `SELECT year_col, count(year_col), max(big_int_col), avg(small_int_col) FROM sbtest1 WHERE big_int_col > 0 GROUP BY set_col ORDER BY year_col`,
		ExpectedPlan: "Project\n" +
			" ├─ columns: [sbtest1.year_col:3!null, count(sbtest1.year_col):1!null as count(year_col), max(sbtest1.big_int_col):2!null as max(big_int_col), avg(sbtest1.small_int_col):0 as avg(small_int_col)]\n" +
			" └─ Sort(sbtest1.year_col:3!null ASC nullsFirst)\n" +
			"     └─ GroupBy\n" +
			"         ├─ select: AVG(sbtest1.small_int_col:0!null), COUNT(sbtest1.year_col:3!null), MAX(sbtest1.big_int_col:1!null), sbtest1.year_col:3!null\n" +
			"         ├─ group: sbtest1.set_col:2!null\n" +
			"         └─ IndexedTableAccess(sbtest1)\n" +
			"             ├─ index: [sbtest1.big_int_col]\n" +
			"             ├─ static: [{(0, ∞)}]\n" +
			"             ├─ colSet: (1-24)\n" +
			"             ├─ tableId: 1\n" +
			"             └─ Table\n" +
			"                 ├─ name: sbtest1\n" +
			"                 └─ columns: [small_int_col big_int_col set_col year_col]\n" +
			"",
		ExpectedEstimates: "Project\n" +
			" ├─ columns: [sbtest1.year_col:3!null, count(sbtest1.year_col):1!null as count(year_col), max(sbtest1.big_int_col):2!null as max(big_int_col), avg(sbtest1.small_int_col):0 as avg(small_int_col)]\n" +
			" └─ Sort(sbtest1.year_col:3!null ASC nullsFirst)\n" +
			"     └─ GroupBy\n" +
			"         ├─ select: AVG(sbtest1.small_int_col:0!null), COUNT(sbtest1.year_col:3!null), MAX(sbtest1.big_int_col:1!null), sbtest1.year_col:3!null\n" +
			"         ├─ group: sbtest1.set_col:2!null\n" +
			"         └─ IndexedTableAccess(sbtest1)\n" +
			"             ├─ index: [sbtest1.big_int_col]\n" +
			"             ├─ static: [{(0, ∞)}]\n" +
			"             ├─ colSet: (1-24)\n" +
			"             ├─ tableId: 1\n" +
			"             └─ Table\n" +
			"                 ├─ name: sbtest1\n" +
			"                 └─ columns: [small_int_col big_int_col set_col year_col]\n" +
			"",
		ExpectedAnalysis: "Project\n" +
			" ├─ columns: [sbtest1.year_col:3!null, count(sbtest1.year_col):1!null as count(year_col), max(sbtest1.big_int_col):2!null as max(big_int_col), avg(sbtest1.small_int_col):0 as avg(small_int_col)]\n" +
			" └─ Sort(sbtest1.year_col:3!null ASC nullsFirst)\n" +
			"     └─ GroupBy\n" +
			"         ├─ select: AVG(sbtest1.small_int_col:0!null), COUNT(sbtest1.year_col:3!null), MAX(sbtest1.big_int_col:1!null), sbtest1.year_col:3!null\n" +
			"         ├─ group: sbtest1.set_col:2!null\n" +
			"         └─ IndexedTableAccess(sbtest1)\n" +
			"             ├─ index: [sbtest1.big_int_col]\n" +
			"             ├─ static: [{(0, ∞)}]\n" +
			"             ├─ colSet: (1-24)\n" +
			"             ├─ tableId: 1\n" +
			"             └─ Table\n" +
			"                 ├─ name: sbtest1\n" +
			"                 └─ columns: [small_int_col big_int_col set_col year_col]\n" +
			"",
	},
	{
		Query: `SELECT count(id) FROM sbtest1 WHERE big_int_col > 0`,
		ExpectedPlan: "Project\n" +
			" ├─ columns: [count(sbtest1.id):0!null as count(id)]\n" +
			" └─ GroupBy\n" +
			"     ├─ select: COUNT(sbtest1.id:0!null)\n" +
			"     ├─ group: \n" +
			"     └─ IndexedTableAccess(sbtest1)\n" +
			"         ├─ index: [sbtest1.big_int_col]\n" +
			"         ├─ static: [{(0, ∞)}]\n" +
			"         ├─ colSet: (1-24)\n" +
			"         ├─ tableId: 1\n" +
			"         └─ Table\n" +
			"             ├─ name: sbtest1\n" +
			"             └─ columns: [id big_int_col]\n" +
			"",
		ExpectedEstimates: "Project\n" +
			" ├─ columns: [count(sbtest1.id):0!null as count(id)]\n" +
			" └─ GroupBy\n" +
			"     ├─ select: COUNT(sbtest1.id:0!null)\n" +
			"     ├─ group: \n" +
			"     └─ IndexedTableAccess(sbtest1)\n" +
			"         ├─ index: [sbtest1.big_int_col]\n" +
			"         ├─ static: [{(0, ∞)}]\n" +
			"         ├─ colSet: (1-24)\n" +
			"         ├─ tableId: 1\n" +
			"         └─ Table\n" +
			"             ├─ name: sbtest1\n" +
			"             └─ columns: [id big_int_col]\n" +
			"",
		ExpectedAnalysis: "Project\n" +
			" ├─ columns: [count(sbtest1.id):0!null as count(id)]\n" +
			" └─ GroupBy\n" +
			"     ├─ select: COUNT(sbtest1.id:0!null)\n" +
			"     ├─ group: \n" +
			"     └─ IndexedTableAccess(sbtest1)\n" +
			"         ├─ index: [sbtest1.big_int_col]\n" +
			"         ├─ static: [{(0, ∞)}]\n" +
			"         ├─ colSet: (1-24)\n" +
			"         ├─ tableId: 1\n" +
			"         └─ Table\n" +
			"             ├─ name: sbtest1\n" +
			"             └─ columns: [id big_int_col]\n" +
			"",
	},
	{
		Query: `SELECT * FROM sbtest1 WHERE big_int_col > 0`,
		ExpectedPlan: "IndexedTableAccess(sbtest1)\n" +
			" ├─ index: [sbtest1.big_int_col]\n" +
			" ├─ static: [{(0, ∞)}]\n" +
			" ├─ colSet: (1-24)\n" +
			" ├─ tableId: 1\n" +
			" └─ Table\n" +
			"     ├─ name: sbtest1\n" +
			"     └─ columns: [id tiny_int_col unsigned_tiny_int_col small_int_col unsigned_small_int_col medium_int_col unsigned_medium_int_col int_col unsigned_int_col big_int_col unsigned_big_int_col decimal_col float_col double_col bit_col char_col var_char_col enum_col set_col date_col time_col datetime_col timestamp_col year_col]\n" +
			"",
		ExpectedEstimates: "IndexedTableAccess(sbtest1)\n" +
			" ├─ index: [sbtest1.big_int_col]\n" +
			" ├─ static: [{(0, ∞)}]\n" +
			" ├─ colSet: (1-24)\n" +
			" ├─ tableId: 1\n" +
			" └─ Table\n" +
			"     ├─ name: sbtest1\n" +
			"     └─ columns: [id tiny_int_col unsigned_tiny_int_col small_int_col unsigned_small_int_col medium_int_col unsigned_medium_int_col int_col unsigned_int_col big_int_col unsigned_big_int_col decimal_col float_col double_col bit_col char_col var_char_col enum_col set_col date_col time_col datetime_col timestamp_col year_col]\n" +
			"",
		ExpectedAnalysis: "IndexedTableAccess(sbtest1)\n" +
			" ├─ index: [sbtest1.big_int_col]\n" +
			" ├─ static: [{(0, ∞)}]\n" +
			" ├─ colSet: (1-24)\n" +
			" ├─ tableId: 1\n" +
			" └─ Table\n" +
			"     ├─ name: sbtest1\n" +
			"     └─ columns: [id tiny_int_col unsigned_tiny_int_col small_int_col unsigned_small_int_col medium_int_col unsigned_medium_int_col int_col unsigned_int_col big_int_col unsigned_big_int_col decimal_col float_col double_col bit_col char_col var_char_col enum_col set_col date_col time_col datetime_col timestamp_col year_col]\n" +
			"",
	},
}
