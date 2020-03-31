short_hand_0 = from( bucket: "thurston" )
	|> range( start: -1m )
	|> filter( fn: (r) => ( r._measurement == "cpu") )
	|> last()
	|> pivot(
		rowKey:["_time"],
		columnKey: ["_field"],
		valueColumn: "_value"
	)
	|> group( columns: [] )
	|> keep( columns: ["_time", "cpu", "usage_user", "usage_system", "usage_idle"])

short_hand_0
