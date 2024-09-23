package applicationtemplate

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListApplicationTemplatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ApplicationTemplate
}

type ListApplicationTemplatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ApplicationTemplate
}

type ListApplicationTemplatesOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListApplicationTemplatesOperationOptions() ListApplicationTemplatesOperationOptions {
	return ListApplicationTemplatesOperationOptions{}
}

func (o ListApplicationTemplatesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListApplicationTemplatesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListApplicationTemplatesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListApplicationTemplatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListApplicationTemplatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListApplicationTemplates - List applicationTemplates. Retrieve a list of applicationTemplate objects from the
// Microsoft Entra application gallery.
func (c ApplicationTemplateClient) ListApplicationTemplates(ctx context.Context, options ListApplicationTemplatesOperationOptions) (result ListApplicationTemplatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListApplicationTemplatesCustomPager{},
		Path:          "/applicationTemplates",
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]stable.ApplicationTemplate `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListApplicationTemplatesComplete retrieves all the results into a single object
func (c ApplicationTemplateClient) ListApplicationTemplatesComplete(ctx context.Context, options ListApplicationTemplatesOperationOptions) (ListApplicationTemplatesCompleteResult, error) {
	return c.ListApplicationTemplatesCompleteMatchingPredicate(ctx, options, ApplicationTemplateOperationPredicate{})
}

// ListApplicationTemplatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ApplicationTemplateClient) ListApplicationTemplatesCompleteMatchingPredicate(ctx context.Context, options ListApplicationTemplatesOperationOptions, predicate ApplicationTemplateOperationPredicate) (result ListApplicationTemplatesCompleteResult, err error) {
	items := make([]stable.ApplicationTemplate, 0)

	resp, err := c.ListApplicationTemplates(ctx, options)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListApplicationTemplatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
