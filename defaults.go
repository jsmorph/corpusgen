package corpusgen

var (
	DefaultValue = &Value{
		Map:    0.5,
		Array:  0.5,
		Int:    0.4,
		String: 0.6,
		Strings: String{
			Length: Int{
				Min: 5,
				Max: 30,
			},
		},
		Arrays: Array{
			Length: Int{
				Min: 1,
				Max: 5,
			},
		},
		Ints: Int{
			Min: -100,
			Max: 100,
		},
		Maps: Map{
			NumProperties: Int{
				Min: 2,
				Max: 5,
			},
			Properties: String{
				Length: Int{
					Min: 3,
					Max: 20,
				},
			},
		},
		Decays: Decays{
			Map:   0.5,
			Array: 0.8,
		},
	}

	DefaultTrimmer = &Trimmer{
		Map:   0.5,
		Array: 0.5,
	}
)
