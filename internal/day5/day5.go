package day5

import (
	"errors"
	"fmt"
	"github.com/jhungerford/adventofcode-2023/internal/util"
	"math"
	"regexp"
	"sync"
)

// Part1 returns the lowest location number that corresponds to any of the initial seed numbers.
func Part1(almanac Almanac) int {
	lowestLocation := math.MaxInt

	for _, seed := range almanac.seeds {
		lowestLocation = min(lowestLocation, almanac.resolveLocation(seed))
	}

	return lowestLocation
}

// Part2 returns the lowest location number that corresponds to any of the initial seed numbers,
// where the seed like contains ranges of seeds.
func Part2(almanac Almanac) int {
	locs := make(chan int, len(almanac.seeds)/2)
	var wg sync.WaitGroup

	for i := 0; i < len(almanac.seeds); i += 2 {
		wg.Add(1)

		go func(seed, length int) {
			defer wg.Done()

			lowest := math.MaxInt

			for plusSeed := 0; plusSeed < length; plusSeed++ {
				lowest = min(lowest, almanac.resolveLocation(seed+plusSeed))
			}

			locs <- lowest
		}(almanac.seeds[i], almanac.seeds[i+1])
	}

	wg.Wait()
	close(locs)

	lowestLocation := math.MaxInt

	for loc := range locs {
		lowestLocation = min(lowestLocation, loc)
	}

	return lowestLocation
}

// LoadAlmanac loads an almanac from the given file.
func LoadAlmanac(filename string) (Almanac, error) {
	mapRE, err := regexp.Compile("^([a-z]+)-to-([a-z]+) map:$")
	if err != nil {
		return Almanac{}, fmt.Errorf("failed to compile map regex: %w", err)
	}

	type parseAlmanac struct {
		almanac         Almanac
		currentCategory string
	}

	sectionParsers := map[string]util.SectionLineParser[parseAlmanac]{
		"seeds": func(line string, parse *parseAlmanac) (string, error) {
			// First line looks like 'seeds: 79 14 55 13'
			seeds, serr := util.IntList(line[7:])
			if serr != nil {
				return "", fmt.Errorf("failed to parse seeds from seed like '%s': %w", line, serr)
			}

			parse.almanac.seeds = seeds

			return "maps", nil
		},
		"maps": func(line string, parse *parseAlmanac) (string, error) {
			if groups := mapRE.FindStringSubmatch(line); len(groups) == 3 {
				// Maps have a header line like 'seed-to-soil map:'
				source, dest := groups[1], groups[2]

				parse.almanac.categories[source] = category{
					source: source,
					dest:   dest,
				}

				parse.currentCategory = source

			} else {
				// Followed by '0 15 37', which is a dest range start, a source range start, and range length.
				rangeNums, serr := util.IntList(line)
				if serr != nil {
					return "", fmt.Errorf("failed to parse range from '%s': %w", line, serr)
				}

				if len(rangeNums) != 3 {
					return "", fmt.Errorf("incorrect range '%s': %w", line, errors.New("range malformed"))
				}

				cat := parse.almanac.categories[parse.currentCategory]

				cat.idRanges = append(cat.idRanges, idRange{
					sourceStart: rangeNums[1],
					destStart:   rangeNums[0],
					size:        rangeNums[2],
				})

				parse.almanac.categories[parse.currentCategory] = cat
			}

			return "maps", nil
		},
	}

	init := parseAlmanac{
		almanac: Almanac{
			seeds:      nil,
			categories: map[string]category{},
		},
		currentCategory: "",
	}

	parsed, err := util.ParseInputLinesSections(filename, "seeds", init, true, sectionParsers)
	if err != nil {
		return Almanac{}, fmt.Errorf("failed to parse almanac: %w", err)
	}

	return parsed.almanac, nil
}

// Almanac contains a list of seeds that need to be planted, and categories that map one resource to another
// with id ranges.
type Almanac struct {
	seeds      []int
	categories map[string]category
}

// category describes how to convert ids from a source category into a destination category.
type category struct {
	source, dest string
	idRanges     []idRange
}

// idRange maps ids from a source category to a destination category.
type idRange struct {
	sourceStart int
	destStart   int
	size        int
}

// resolveLocation resolves the location of the given seed.
func (a Almanac) resolveLocation(seed int) int {
	source := "seed"
	id := seed

	for source != "location" {
		cat := a.categories[source]

		source = cat.dest
		id = cat.mapID(id)
	}

	return id
}

// mapID maps the id from a source category to a destination category
func (c category) mapID(id int) int {
	for _, r := range c.idRanges {
		if id >= r.sourceStart && id < r.sourceStart+r.size {
			return r.destStart + id - r.sourceStart
		}
	}

	return id
}
