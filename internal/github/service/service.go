package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/invilliafelipeflores/notfiy/internal/github/model"
)

type GithubService struct {
	accessToken string
}

func NewGithubService(accessToken string) *GithubService {
	return &GithubService{
		accessToken: accessToken,
	}
}

func (g GithubService) GetPulls(repo string) ([]model.PullRequest, error) {

	c := http.Client{Timeout: time.Duration(15) * time.Second}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/repos/%s/pulls", repo), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", `application/vnd.github+json`)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", g.accessToken))

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	pullRequest := []model.PullRequest{}
	err = json.Unmarshal(body, &pullRequest)
	if err != nil {
		return nil, err
	}
	return pullRequest, nil
}
