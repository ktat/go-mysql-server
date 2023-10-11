// Code generated by plangen.

// Copyright 2023 Dolthub, Inc.
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

var TpccPlanTests = []QueryPlanTest{
	{
		Query: `
SELECT o_id, o_entry_d, COALESCE(o_carrier_id,0)
FROM orders2
WHERE
  o_w_id = 1 AND
  o_d_id = 3 AND
  o_c_id = 20001 AND
  o_id = (SELECT MAX(o_id) FROM orders2 WHERE o_w_id = 1 AND o_d_id = 3 AND o_c_id = 20001)`,
		ExpectedPlan: "Project\n" +
			" ├─ columns: [orders2.o_id:1!null, orders2.o_entry_d:5, coalesce(orders2.o_carrier_id,0) as COALESCE(o_carrier_id,0)]\n" +
			" └─ LookupJoin\n" +
			"     ├─ Eq\n" +
			"     │   ├─ orders2.o_id:1!null\n" +
			"     │   └─ scalarSubq0.MAX(o_id):0!null\n" +
			"     ├─ OrderedDistinct\n" +
			"     │   └─ Max1Row\n" +
			"     │       └─ SubqueryAlias\n" +
			"     │           ├─ name: scalarSubq0\n" +
			"     │           ├─ outerVisibility: false\n" +
			"     │           ├─ cacheable: true\n" +
			"     │           └─ Project\n" +
			"     │               ├─ columns: [max(orders2.o_id):0!null as MAX(o_id)]\n" +
			"     │               └─ GroupBy\n" +
			"     │                   ├─ select: MAX(orders2.o_id:0!null)\n" +
			"     │                   ├─ group: \n" +
			"     │                   └─ IndexedTableAccess(orders2)\n" +
			"     │                       ├─ index: [orders2.o_w_id,orders2.o_d_id,orders2.o_c_id,orders2.o_id]\n" +
			"     │                       ├─ static: [{[1, 1], [3, 3], [20001, 20001], [NULL, ∞)}]\n" +
			"     │                       └─ Table\n" +
			"     │                           ├─ name: orders2\n" +
			"     │                           └─ columns: [o_id o_d_id o_w_id o_c_id]\n" +
			"     └─ Filter\n" +
			"         ├─ AND\n" +
			"         │   ├─ AND\n" +
			"         │   │   ├─ Eq\n" +
			"         │   │   │   ├─ orders2.o_w_id:2!null\n" +
			"         │   │   │   └─ 1 (tinyint)\n" +
			"         │   │   └─ Eq\n" +
			"         │   │       ├─ orders2.o_d_id:1!null\n" +
			"         │   │       └─ 3 (tinyint)\n" +
			"         │   └─ Eq\n" +
			"         │       ├─ orders2.o_c_id:3\n" +
			"         │       └─ 20001 (smallint)\n" +
			"         └─ IndexedTableAccess(orders2)\n" +
			"             ├─ index: [orders2.o_w_id,orders2.o_d_id,orders2.o_id]\n" +
			"             └─ Table\n" +
			"                 ├─ name: orders2\n" +
			"                 └─ columns: [o_id o_d_id o_w_id o_c_id o_entry_d o_carrier_id o_ol_cnt o_all_local]\n" +
			"",
	},
	{
		Query: `
select o_id, o.o_d_id
from
  orders2 o,
  (
    select o_c_id, o_w_id, o_d_id, count(distinct o_id)
    from orders2
    where o_w_id=1  and o_id > 2100 and o_id < 11153
    group by o_c_id,o_d_id,o_w_id
    having count( distinct o_id) > 1
    limit 1
  ) t
  where
    t.o_w_id=o.o_w_id and
    t.o_d_id=o.o_d_id and
    t.o_c_id=o.o_c_id
  limit 1;`,
		ExpectedPlan: "Limit(1)\n" +
			" └─ Project\n" +
			"     ├─ columns: [o.o_id:0!null, o.o_d_id:1!null]\n" +
			"     └─ HashJoin\n" +
			"         ├─ AND\n" +
			"         │   ├─ AND\n" +
			"         │   │   ├─ Eq\n" +
			"         │   │   │   ├─ t.o_w_id:5!null\n" +
			"         │   │   │   └─ o.o_w_id:2!null\n" +
			"         │   │   └─ Eq\n" +
			"         │   │       ├─ t.o_d_id:6!null\n" +
			"         │   │       └─ o.o_d_id:1!null\n" +
			"         │   └─ Eq\n" +
			"         │       ├─ t.o_c_id:4\n" +
			"         │       └─ o.o_c_id:3\n" +
			"         ├─ TableAlias(o)\n" +
			"         │   └─ ProcessTable\n" +
			"         │       └─ Table\n" +
			"         │           ├─ name: orders2\n" +
			"         │           └─ columns: [o_id o_d_id o_w_id o_c_id]\n" +
			"         └─ HashLookup\n" +
			"             ├─ left-key: TUPLE(o.o_w_id:2!null, o.o_d_id:1!null, o.o_c_id:3)\n" +
			"             ├─ right-key: TUPLE(t.o_w_id:1!null, t.o_d_id:2!null, t.o_c_id:0)\n" +
			"             └─ SubqueryAlias\n" +
			"                 ├─ name: t\n" +
			"                 ├─ outerVisibility: false\n" +
			"                 ├─ cacheable: true\n" +
			"                 └─ Limit(1)\n" +
			"                     └─ Project\n" +
			"                         ├─ columns: [orders2.o_c_id:1, orders2.o_w_id:2!null, orders2.o_d_id:3!null, countdistinct([orders2.o_id]):0!null as count(distinct o_id)]\n" +
			"                         └─ Having\n" +
			"                             ├─ GreaterThan\n" +
			"                             │   ├─ countdistinct([orders2.o_id]):0!null\n" +
			"                             │   └─ 1 (tinyint)\n" +
			"                             └─ GroupBy\n" +
			"                                 ├─ select: COUNTDISTINCT([orders2.o_id]), orders2.o_c_id:3, orders2.o_w_id:2!null, orders2.o_d_id:1!null, orders2.o_id:0!null\n" +
			"                                 ├─ group: orders2.o_c_id:3, orders2.o_d_id:1!null, orders2.o_w_id:2!null\n" +
			"                                 └─ Filter\n" +
			"                                     ├─ AND\n" +
			"                                     │   ├─ GreaterThan\n" +
			"                                     │   │   ├─ orders2.o_id:0!null\n" +
			"                                     │   │   └─ 2100 (smallint)\n" +
			"                                     │   └─ LessThan\n" +
			"                                     │       ├─ orders2.o_id:0!null\n" +
			"                                     │       └─ 11153 (smallint)\n" +
			"                                     └─ IndexedTableAccess(orders2)\n" +
			"                                         ├─ index: [orders2.o_w_id,orders2.o_d_id,orders2.o_id]\n" +
			"                                         ├─ static: [{[1, 1], [NULL, ∞), [NULL, ∞)}]\n" +
			"                                         └─ Table\n" +
			"                                             ├─ name: orders2\n" +
			"                                             └─ columns: [o_id o_d_id o_w_id o_c_id o_entry_d o_carrier_id o_ol_cnt o_all_local]\n" +
			"",
	},
}
