// Package tournament creates a standings chart from game results.
package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type stats struct {
	name                string
	wins, losses, draws int
	matches, points     int
}

type statsMap map[string]*stats

const (
	comment = '#'
)

func getstats(team string, standings statsMap) *stats {
	var s *stats
	s, ok := standings[team]
	if !ok {
		s = &stats{name: team}
		standings[team] = s
	}
	return s
}

func score(s1, s2 *stats, outcome string) error {
	var err error
	switch outcome {
	case "win":
		s1.wins++
		s2.losses++
	case "loss":
		s1.losses++
		s2.wins++
	case "draw":
		s1.draws++
		s2.draws++
	default:
		err = fmt.Errorf("bad outcome '%s'", outcome)
	}
	return err
}

func total(standings statsMap) []*stats {
	sl := make([]*stats, 0, len(standings))
	for _, s := range standings {
		s.matches = s.wins + s.losses + s.draws
		s.points = 3*s.wins + s.draws
		sl = append(sl, s)
	}
	compare := func(i, j int) bool {
		pi := sl[i].points
		pj := sl[j].points
		return pi > pj || pi == pj && sl[i].name < sl[j].name
	}
	sort.Slice(sl, compare)
	return sl
}

// Tally reads tournament results and writes tournament standings.
func Tally(in io.Reader, out io.Writer) error {
	var err error
	standings := make(statsMap)
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		if len(s) > 0 && s[0] != comment {
			f := strings.Split(s, ";")
			if len(f) != 3 {
				err = fmt.Errorf("bad input line '%s'", s)
				break
			}
			f[0] = strings.TrimSpace(f[0])
			f[1] = strings.TrimSpace(f[1])
			f[2] = strings.TrimSpace(f[2])
			if f[0] == f[1] {
				err = fmt.Errorf("duplicate team name %s'", f[0])
				break
			}
			s1 := getstats(f[0], standings)
			s2 := getstats(f[1], standings)
			if err = score(s1, s2, f[2]); err != nil {
				break
			}
		}
	}
	if err == nil {
		err = scanner.Err()
	}
	if err == nil {
		heading := "Team                           | MP |  W |  D |  L |  P"
		format := "%-31.31s|%3d |%3d |%3d |%3d |%3d\n"
		fmt.Fprintln(out, heading)
		for _, p := range total(standings) {
			fmt.Fprintf(out, format, p.name, p.matches, p.wins, p.draws,
				p.losses, p.points)
		}
	}
	return err
}
