package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
)

type OverwatchRates struct {
	Rates struct {
		Rates []struct {
			ID    string `json:"id"`
			Cells struct {
				Name     string  `json:"name"`
				Winrate  float64 `json:"winrate"`
				Pickrate float64 `json:"pickrate"`
				Banrate  float64 `json:"banrate"`
			} `json:"cells"`
			Hero struct {
				Color    string `json:"color"`
				Name     string `json:"name"`
				Portrait string `json:"portrait"`
				Subrole  string `json:"subrole"`
				Role     string `json:"role"`
				RoleIcon string `json:"roleIcon"`
			} `json:"hero"`
		} `json:"rates"`
		Extrema struct {
			All struct {
				Winrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"winrate"`
				Pickrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"pickrate"`
				Banrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"banrate"`
			} `json:"all"`
			Support struct {
				Winrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"winrate"`
				Pickrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"pickrate"`
				Banrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"banrate"`
			} `json:"support"`
			Tactician struct {
				Winrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"winrate"`
				Pickrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"pickrate"`
				Banrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"banrate"`
			} `json:"tactician"`
			Damage struct {
				Winrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"winrate"`
				Pickrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"pickrate"`
				Banrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"banrate"`
			} `json:"damage"`
			Flanker struct {
				Winrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"winrate"`
				Pickrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"pickrate"`
				Banrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"banrate"`
			} `json:"flanker"`
			Sharpshooter struct {
				Winrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"winrate"`
				Pickrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"pickrate"`
				Banrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"banrate"`
			} `json:"sharpshooter"`
			Specialist struct {
				Winrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"winrate"`
				Pickrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"pickrate"`
				Banrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"banrate"`
			} `json:"specialist"`
			Survivor struct {
				Winrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"winrate"`
				Pickrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"pickrate"`
				Banrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"banrate"`
			} `json:"survivor"`
			Tank struct {
				Winrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"winrate"`
				Pickrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"pickrate"`
				Banrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"banrate"`
			} `json:"tank"`
			Initiator struct {
				Winrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"winrate"`
				Pickrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"pickrate"`
				Banrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"banrate"`
			} `json:"initiator"`
			Stalwart struct {
				Winrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"winrate"`
				Pickrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"pickrate"`
				Banrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"banrate"`
			} `json:"stalwart"`
			Recon struct {
				Winrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"winrate"`
				Pickrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"pickrate"`
				Banrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"banrate"`
			} `json:"recon"`
			Medic struct {
				Winrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"winrate"`
				Pickrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"pickrate"`
				Banrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"banrate"`
			} `json:"medic"`
			Bruiser struct {
				Winrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"winrate"`
				Pickrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"pickrate"`
				Banrate struct {
					Max float64 `json:"max"`
					Min float64 `json:"min"`
				} `json:"banrate"`
			} `json:"bruiser"`
		} `json:"extrema"`
		Selected struct {
			Input  string `json:"input"`
			Map    string `json:"map"`
			Region string `json:"region"`
			Role   string `json:"role"`
			Rq     string `json:"rq"`
			Tier   string `json:"tier"`
		} `json:"selected"`
	} `json:"rates"`
	Columns []struct {
		ID            string `json:"id"`
		Text          string `json:"text"`
		Sortable      bool   `json:"sortable"`
		SortDirection string `json:"sortDirection,omitempty"`
	} `json:"columns"`
}

func main() {
	resp, err := http.Get("https://overwatch.blizzard.com/en-us/rates/data/?input=PC&map=all-maps&region=Americas&role=All&rq=1&tier=All")
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer resp.Body.Close()

	var rates OverwatchRates

	if err := json.NewDecoder(resp.Body).Decode(&rates); err != nil {
		log.Fatalf("Parsing failed: %v", err)
	}

	for _, item := range rates.Rates.Rates {
		heroName := item.Cells.Name
		heroWinRate := item.Cells.Winrate
		heroPickRate := item.Cells.Pickrate

		// sorting by win rate, compares the index using >
		sort.Slice(rates.Rates.Rates, func(i, j int) bool {
			return rates.Rates.Rates[i].Cells.Winrate > rates.Rates.Rates[j].Cells.Winrate
		})

		fmt.Println(heroName)
		fmt.Printf("Win Rate: %v \n\n", heroWinRate)
		fmt.Printf("Pick Rate: %v \n\n", heroPickRate)
	}

	fmt.Scanln()
}
