package apiclient

import (
	"fmt"
	"strings"

	"github.com/junioryono/Riot-API-Golang/src/apiclient/ratelimiter"
	"github.com/junioryono/Riot-API-Golang/src/constants/queue"
	"github.com/junioryono/Riot-API-Golang/src/constants/region"
)

func (c *client) GetLeagueExpEntries(r region.Region, q queue.Queue, tier, division string, page int) ([]LeaguePosition, error) {
	var res []LeaguePosition
	_, err := c.dispatchAndUnmarshal(r, "/lol/league-exp/v4/entries", fmt.Sprintf("/%s/%s/%s?page=%d", q.String(), strings.ToUpper(tier), strings.ToUpper(division), page), nil, ratelimiter.GetLeagueExpEntries, &res)
	return res, err
}
