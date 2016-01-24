package github

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/gophergala2016/gophers/json"
)

func TestListOrgs(t *testing.T) {
	t.Parallel()

	j := Client.Get(t, "/user/orgs", 200).JSON(t)

	var found bool
	v := j.KeepFields("login")
	expect := JSON(`{"login": "gophergala2016"}`).String()
	for _, e := range v.(JSONArray) {
		if AsJSON(e).String() == expect {
			found = true
			break
		}
	}

	assert.True(t, found, "current user doesn't belong to gophergala2016 organization")
}
