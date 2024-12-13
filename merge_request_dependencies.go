//
// Copyright 2024, Lukas Lüdke
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
//

package gitlab

import (
	"fmt"
	"net/http"
	"time"
)

// MergeRequestDependenciesService handles communication with the merge request
// dependencies related methods of the GitLab API.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/merge_requests.html#get-merge-request-dependencies
type MergeRequestDependenciesService struct {
	client *Client
}

// MergeRequestDependency represents a GitLab merge request dependency.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/merge_requests.html#create-a-merge-request-dependency
type MergeRequestDependency struct {
	ID                   int                  `json:"id"`
	BlockingMergeRequest BlockingMergeRequest `json:"blocking_merge_request"`
	ProjectID            int                  `json:"project_id"`
}

// BlockingMergeRequest represents a GitLab merge request dependency.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/merge_requests.html#create-a-merge-request-dependency
type BlockingMergeRequest struct {
	ID                          int                    `json:"id"`
	Iid                         int                    `json:"iid"`
	TargetBranch                string                 `json:"target_branch"`
	SourceBranch                string                 `json:"source_branch"`
	ProjectID                   int                    `json:"project_id"`
	Title                       string                 `json:"title"`
	State                       string                 `json:"state"`
	CreatedAt                   time.Time              `json:"created_at"`
	UpdatedAt                   time.Time              `json:"updated_at"`
	Upvotes                     int                    `json:"upvotes"`
	Downvotes                   int                    `json:"downvotes"`
	Author                      *BasicUser             `json:"author"`
	Assignee                    *BasicUser             `json:"assignee"`
	Assignees                   []*BasicUser           `json:"assignees"`
	Reviewers                   []*BasicUser           `json:"reviewers"`
	SourceProjectID             int                    `json:"source_project_id"`
	TargetProjectID             int                    `json:"target_project_id"`
	Labels                      *LabelOptions          `json:"labels"`
	Description                 string                 `json:"description"`
	Draft                       bool                   `json:"draft"`
	WorkInProgress              bool                   `json:"work_in_progress"`
	Milestone                   *string                `json:"milestone"`
	MergeWhenPipelineSucceeds   bool                   `json:"merge_when_pipeline_succeeds"`
	DetailedMergeStatus         string                 `json:"detailed_merge_status"`
	MergedBy                    *BasicUser             `json:"merged_by"`
	MergedAt                    *time.Time             `json:"merged_at"`
	ClosedBy                    *BasicUser             `json:"closed_by"`
	ClosedAt                    *time.Time             `json:"closed_at"`
	Sha                         string                 `json:"sha"`
	MergeCommitSha              string                 `json:"merge_commit_sha"`
	SquashCommitSha             string                 `json:"squash_commit_sha"`
	UserNotesCount              int                    `json:"user_notes_count"`
	ShouldRemoveSourceBranch    *bool                  `json:"should_remove_source_branch"`
	ForceRemoveSourceBranch     bool                   `json:"force_remove_source_branch"`
	WebURL                      string                 `json:"web_url"`
	References                  *IssueReferences       `json:"references"`
	DiscussionLocked            *bool                  `json:"discussion_locked"`
	TimeStats                   *TimeStats             `json:"time_stats"`
	Squash                      bool                   `json:"squash"`
	ApprovalsBeforeMerge        *int                   `json:"approvals_before_merge"`
	Reference                   string                 `json:"reference"`
	TaskCompletionStatus        *TasksCompletionStatus `json:"task_completion_status"`
	HasConflicts                bool                   `json:"has_conflicts"`
	BlockingDiscussionsResolved bool                   `json:"blocking_discussions_resolved"`
	MergeStatus                 string                 `json:"merge_status"`
	MergeUser                   *BasicUser             `json:"merge_user"`
	MergeAfter                  time.Time              `json:"merge_after"`
	Imported                    bool                   `json:"imported"`
	ImportedFrom                string                 `json:"imported_from"`
	PreparedAt                  *time.Time             `json:"prepared_at"`
	SquashOnMerge               bool                   `json:"squash_on_merge"`
}

func (m MergeRequestDependency) String() string {
	return Stringify(m)
}

// CreateMergeRequestDependencyOptions represents the available CreateMergeRequestDependency()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/merge_requests.html#create-a-merge-request-dependency
type CreateMergeRequestDependencyOptions struct {
	BlockingMergeRequestID int `url:"blocking_merge_request_id,omitempty" json:"blocking_merge_request_id,omitempty"`
}

// CreateMergeRequestDependency creates a new merge request dependency for a given
// merge request.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/merge_requests.html#create-a-merge-request-dependency
func (s *MergeRequestDependenciesService) CreateMergeRequestDependency(pid interface{}, mergeRequest int, opts CreateMergeRequestDependencyOptions, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/merge_requests/%d/blocks", PathEscape(project), mergeRequest)

	req, err := s.client.NewRequest(http.MethodPost, u, opts, options)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// DeleteMergeRequestDependency deletes a merge request dependency for a given
// merge request.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/merge_requests.html#delete-a-merge-request-dependency
func (s *MergeRequestDependenciesService) DeleteMergeRequestDependency(pid interface{}, mergeRequest int, blockingMergeRequest int, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/merge_requests/%d/blocks/%d", PathEscape(project), mergeRequest, blockingMergeRequest)

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// GetMergeRequestDependencies gets a list of merge request dependencies.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/merge_requests.html#get-merge-request-dependencies
func (s *MergeRequestDependenciesService) GetMergeRequestDependencies(pid interface{}, mergeRequest int, options ...RequestOptionFunc) ([]MergeRequestDependency, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/merge_requests/%d/blocks", PathEscape(project), mergeRequest)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var mrd []MergeRequestDependency
	resp, err := s.client.Do(req, &mrd)
	if err != nil {
		return nil, resp, err
	}

	return mrd, resp, err
}
