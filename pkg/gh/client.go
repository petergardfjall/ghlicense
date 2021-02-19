package gh

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-github/v32/github"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

var (
	ErrRepoNotFound    = errors.New("repository does not exist")
	ErrLicenseNotFound = errors.New("no license found for repository")
)

type Client struct {
	*github.Client
}

// NewClient creates a GitHub client using the supplied personal access token
// for authentication.
func NewClient(ctx context.Context, token string) *Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	c := github.NewClient(tc)

	return &Client{Client: c}
}

// GetLicense retrieves the repository license, as determined by GitHub (if one
// has been found).
func (c *Client) GetLicense(ctx context.Context, repoURL string) (*github.RepositoryLicense, error) {
	owner, repo, err := parseRepoURL(repoURL)
	if err != nil {
		return nil, err
	}

	_, r, err := c.Repositories.Get(ctx, owner, repo)
	if err != nil {
		if r.StatusCode == http.StatusNotFound {
			return nil, ErrRepoNotFound
		}
		return nil, err
	}

	lic, r, err := c.Repositories.License(ctx, owner, repo)
	if err != nil {
		if r.StatusCode == http.StatusNotFound {
			return nil, ErrLicenseNotFound
		}
		return nil, errors.Wrap(err, "retrieve license")
	}

	return lic, nil
}

func parseRepoURL(repoURL string) (owner, repo string, err error) {
	u, err := url.Parse(repoURL)
	if err != nil {
		return "", "", errors.Wrapf(err, "%s: malformed URL", repoURL)
	}
	if u.Hostname() != "github.com" {
		return "", "", errors.New("URL hostname is not github.com")
	}
	path := strings.Trim(u.Path, "/")
	// strip any trailing `.git`
	path = strings.TrimSuffix(path, ".git")

	parts := strings.Split(path, "/")
	if len(parts) < 2 || len(parts) > 2 {
		return "", "", fmt.Errorf("%s: URL does not look like a github repo: path %s does not contain <owner>/<repo>", repoURL, path)
	}
	owner = parts[0]
	repo = parts[1]
	return owner, repo, nil
}
